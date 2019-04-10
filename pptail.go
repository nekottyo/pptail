package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/k0kubun/pp"
)

const (
	layout = "Jan 02 15:04:05"
	// Apr 09 15:57:50 localhost 33ad0cd34e88[947]: 1970-01-01 00:33:39.000000347 +0000 some/image:latest.service_image_1.91880377002d.json: {"level":"INFO", "message":"test"}
	syslogFormat  = `^(\w+\s+\d+\s\d+:\d+:\d+)\s+.+?\s+.+?\s+\w+-\w+-\w+\s+\d+:\d+:\d+\.\d+\s+\+\d+\s+(.*)\.\w+:\s(.*)$`
	tdAgentFormat = `^(.+?)\s+(.+?)\s+(.*)$`
)

var fluentFlag bool
var format = syslogFormat

type message struct {
	Date    string
	Image   string
	Payload interface{}
}

func init() {
	flag.BoolVar(&fluentFlag, "fluent", false, "use fluentd(td-agent) tail format")
}

func main() {
	flag.Parse()
	if fluentFlag {
		format = tdAgentFormat
	}

	s := bufio.NewScanner(os.Stdin)
	go func() {
		for s.Scan() {
			Print(s.Text(), format)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	os.Exit(0)
}

// Print is output colo pretty print.
func Print(s, format string) {
	m := parse(s, format)
	pp.Print(m)
}

func parse(s, format string) message {
	rep := regexp.MustCompile(format)
	result := rep.FindStringSubmatch(s)

	m := message{
		Date:  result[1],
		Image: result[2],
	}

	var payload interface{}
	if err := json.Unmarshal([]byte(result[3]), &payload); err != nil {
		return m
	}

	m.Payload = payload
	return m
}
