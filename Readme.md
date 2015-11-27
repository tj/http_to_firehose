
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

## Terraform

Example Terraform configuration:

```hcl
resource "aws_s3_bucket" "mybucket" {
  bucket = "mybucket"
  acl = "private"
}

resource "aws_iam_role" "firehose-s3" {
  name = "firehose-s3"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "firehose.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "firehose-s3" {
  name = "firehose-s3"
  role = "${aws_iam_role.firehose-s3.id}"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": [
        "s3:AbortMultipartUpload",
        "s3:GetBucketLocation",
        "s3:GetObject",
        "s3:ListBucket",
        "s3:ListBucketMultipartUploads",
        "s3:PutObject"
      ],
      "Resource": [
        "arn:aws:s3:::mybucket",
        "arn:aws:s3:::mybucket/*"
      ]
    }
  ]
}
EOF
}

resource "aws_kinesis_firehose_delivery_stream" "events" {
  name = "events"
  destination = "s3"
  role_arn = "${aws_iam_role.firehose-s3.arn}"
  s3_bucket_arn = "${aws_s3_bucket.mybucket.arn}"
  s3_buffer_size = 5 // 5mb
  s3_buffer_interval = 900 // 15 minutes
  s3_data_compression = "Snappy"
}
```

# License

MIT