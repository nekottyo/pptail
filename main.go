package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/k0kubun/pp"
)

const (
	syslogFormat  = `^(\w+\s+\d+\s\d+:\d+:\d+)\s+.+?\s+.+?\s+\w+-\w+-\w+\s+\d+:\d+:\d+\.\d+\s+\+\d+\s+(.*)\.\w+:\s(.*)$`
	tdAgentFormat = `^(.+?)\s+(.+?)\s+(.*)$`
)

var (
	fluentFlag  bool
	versionFlag bool
	format      = syslogFormat
	version     = "dev"
	commit      = "none"
	date        = "unknown"
)

type message struct {
	Date    string
	Image   string
	Payload interface{}
}

func init() {
	flag.BoolVar(&fluentFlag, "fluent", false, "use fluentd(td-agent) tail format")
	flag.BoolVar(&versionFlag, "v", false, "print version")
	flag.BoolVar(&versionFlag, "version", false, "print version")
}

func main() {
	flag.Parse()
	if fluentFlag {
		format = tdAgentFormat
	}

	if versionFlag {
		fmt.Printf("version %v, commit %v, build at %v\n", version, commit, date)
		os.Exit(0)
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
	m, err := parse(s, format)
	if err != nil {
		return
	}

	pp.Println(m)
}

func parse(s, format string) (message, error) {
	rep := regexp.MustCompile(format)
	result := rep.FindStringSubmatch(s)

	if len(result) < 3 {
		return message{}, errors.New("failed to parse string")
	}

	m := message{
		Date:  result[1],
		Image: result[2],
	}

	var payload interface{}
	if err := json.Unmarshal([]byte(result[3]), &payload); err != nil {
		return m, nil
	}

	m.Payload = payload
	return m, nil
}
