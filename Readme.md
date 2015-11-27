
# http_to_firehose

Simple program which accepts HTTP requests and relays their bodies as
[Amazon Firehose](https://aws.amazon.com/kinesis/firehose/) records.

## Installation

```
$ go get github.com/tj/http_to_firehose/cmd/http_to_firehose
```

## About

 Sends HTTP request bodies to Firehose which lets you buffer to S3 or
 other services, letting you process the data later via Lambda etc.

# License

MIT