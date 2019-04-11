pptail
=====
[![CircleCI](https://circleci.com/gh/nekottyo/pptail.svg?style=svg)](https://circleci.com/gh/nekottyo/pptail)
[![Build Status](https://cloud.drone.io/api/badges/nekottyo/pptail/status.svg)](https://cloud.drone.io/nekottyo/pptail)

pptail is pretty print tail with fluentd format log.

like this format:
```
Apr 09 15:57:50 localhost 33ad0cd34e88[947]: 1970-01-01 00:33:39.000000347 +0000 some/image:latest.service_image_1.91880377002d.json: {"level":"INFO", "message":"test"}
```

or:
```
2019-04-02T08:03:47+09:00       some/image:latest.service_image.3cba43465250  {"level":"INFO", "message":"test"}
```

# Usage

```
$ journalctl -xef | pptail
$ tail -f /path/to/loogfile | pptail
```

if log format is td-agent file output plugin, use `-fluent` option

# Example

syslog format
```sh
➜ echo 'Apr 09 15:57:50 localhost 33ad0cd34e88[947]: 1970-01-01 00:33:39.000000347 +0000 some/image:latest.service_image_1.91880377002d.json: {"level":"INFO", "message":"test"}' | ./pptail        
main.message{
  Date:    "Apr 09 15:57:50",
  Image:   "some/image:latest.service_image_1.91880377002d",
  Payload: map[string]interface {}{
    "level":   "INFO",
    "message": "test",
  },
}
```

fluentd format
```sh
➜ echo '2019-04-02T08:03:47+09:00       some/image:latest.service_image.3cba43465250  {"level":"INFO", "message":"test"}' | ./pptail -fluent
main.message{
  Date:    "2019-04-02T08:03:47+09:00",
  Image:   "some/image:latest.service_image.3cba43465250",
  Payload: map[string]interface {}{
    "level":   "INFO",
    "message": "test",
  },
}
```


# Install

```
go get -u github.com/nekottyo/pptail
```
