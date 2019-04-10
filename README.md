pptail
=====

pptail is pretty print tail with output with fluentd at docker.

like this format:
```
Apr 09 15:57:50 localhost 33ad0cd34e88[947]: 1970-01-01 00:33:39.000000347 +0000 some/image:latest.service_image_1.91880377002d.json: {"level":"INFO", "message":"test"}
```

# Usage

```
$ journalctl -xef | pptail
```

if log format is td-agent file output plugin, use `-fluent` option

# Install

```
go get -u github.com/nekottyo/pptail
```
