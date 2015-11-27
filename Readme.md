
# http_to_firehose

Simple program which accepts HTTP requests and relays their bodies as
[Amazon Firehose](https://aws.amazon.com/kinesis/firehose/) records.

## Installation

Install via `go get` or head over to [Releases](https://github.com/tj/http_to_firehose/releases) for binaries.

```
$ go get github.com/tj/http_to_firehose/cmd/http_to_firehose
```

## About

 Sends HTTP request bodies to Firehose which lets you buffer to S3 or
 other services, letting you process the data later via Lambda etc.

## Examples

Send HTTP requests on :3000 to the "events" stream.

```
$ http_to_firehose --stream-name events
```

Send HTTP requests on :5000 to the "events" stream with basic auth.

```
$ http_to_firehose --address :5000 --stream-name events --username sloth --password somethingslothy
```

# License

MIT