package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrint(t *testing.T) {
	cases := []struct {
		in  string
		out message
	}{
		{
			`Apr 09 15:57:50 localhost 33ad0cd34e88[947]: 1970-01-01 00:33:39.000000347 +0000 some/image:latest.service_image_1.91880377002d.json: {"level":"INFO", "message":"test"}`,
			message{
				Date:    "Apr 09 15:57:50",
				Host:    "localhost",
				Process: "33ad0cd34e88[947]",
				Image:   "some/image:latest.service_image_1",
				Payload: map[string]interface{}{"level": "INFO", "message": "test"},
			},
		},
		{
			`Apr 09 15:57:50 localhost 33ad0cd34e88[947]: 1970-01-01 00:33:39.000000347 +0000 some/image:latest.service_image_1.91880377002d.json: {WORNG JSON}`,
			message{
				Date:    "Apr 09 15:57:50",
				Host:    "localhost",
				Process: "33ad0cd34e88[947]",
				Image:   "some/image:latest.service_image_1",
				Payload: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			actual := parse(c.in)
			fmt.Printf("actual = %#v\n", actual)

			if diff := cmp.Diff(actual, c.out); diff != "" {
				t.Errorf("parse func differs: (-got +want)\n%s", diff)
			}
		})
	}
}
