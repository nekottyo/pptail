package main

import (
	"bufio"
	"encoding/json"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/k0kubun/pp"
)

const (
	layout = "Jan 02 15:04:05"
	format = `^(\w+\s+\d+\s\d+:\d+:\d+)\s+(.+?)\s+(\w+\[\d+\]):\s+\w+-\w+-\w+\s+\d+:\d+:\d+\.\d+\s+\+\d+\s+(.+?)\.\w+\.json:\s(.*)$`
)

type message struct {
	Date    string
	Host    string
	Process string
	Image   string
	Payload interface{}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	go func() {
		for s.Scan() {
			Print(s.Text())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	os.Exit(0)
}

func Print(s string) {
	m := parse(s)
	pp.Print(m)
}

func parse(s string) message {
	// Apr 09 15:57:50 localhost 33ad0cd34e88[947]: 1970-01-01 00:33:39.000000347 +0000 some/image:latest.service_image_1.91880377002d.json: {"level":"INFO", "message":"test"}
	rep := regexp.MustCompile(format)
	result := rep.FindStringSubmatch(s)

	m := message{
		Date:    result[1],
		Host:    result[2],
		Process: result[3],
		Image:   result[4],
	}

	var payload interface{}
	if err := json.Unmarshal([]byte(result[5]), &payload); err != nil {
		return m
	}

	m.Payload = payload
	return m
}
