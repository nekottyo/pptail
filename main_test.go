package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrint_syslogFormat(t *testing.T) {
	cases := []struct {
		in  string
		out message
	}{
		{
			`Apr 09 15:57:50 localhost 33ad0cd34e88[947]: 1970-01-01 00:33:39.000000347 +0000 some/image:latest.service_image_1.91880377002d.json: {"level":"INFO", "message":"test"}`,
			message{
				Date:    "Apr 09 15:57:50",
				Image:   "some/image:latest.service_image_1.91880377002d",
				Payload: map[string]interface{}{"level": "INFO", "message": "test"},
			},
		},
		{
			`Apr 09 15:57:50 localhost 33ad0cd34e88[947]: 1970-01-01 00:33:39.000000347 +0000 some/image:latest.service_image_1.91880377002d.json: {INVALID JSON}`,
			message{
				Date:    "Apr 09 15:57:50",
				Image:   "some/image:latest.service_image_1.91880377002d",
				Payload: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			actual, err := parse(c.in, syslogFormat)

			if diff := cmp.Diff(actual, c.out); diff != "" {
				t.Errorf("parse func differs: (-got +want)\n%s", diff)
			}
			if diff := cmp.Diff(err, nil); diff != "" {
				t.Errorf("parse func differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestPrint_fluentdFormat(t *testing.T) {
	cases := []struct {
		in  string
		out message
	}{
		{
			`2019-04-02T08:03:47+09:00       some/image:latest.service_image.3cba43465250  {"level":"INFO", "message":"test"}`,
			message{
				Date:    "2019-04-02T08:03:47+09:00",
				Image:   "some/image:latest.service_image.3cba43465250",
				Payload: map[string]interface{}{"level": "INFO", "message": "test"},
			},
		},
		{
			`2019-04-02T08:03:47+09:00       some/image:latest.service_image.3cba43465250  {INVALID JSON}`,
			message{
				Date:    "2019-04-02T08:03:47+09:00",
				Image:   "some/image:latest.service_image.3cba43465250",
				Payload: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			actual, err := parse(c.in, tdAgentFormat)

			if diff := cmp.Diff(actual, c.out); diff != "" {
				t.Errorf("parse func differs: (-got +want)\n%s", diff)
			}
			if diff := cmp.Diff(err, nil); diff != "" {
				t.Errorf("parse func differs: (-got +want)\n%s", diff)
			}
		})
	}
}
