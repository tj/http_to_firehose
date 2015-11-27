// Package client transfers data to Firehose. Eventually retrying / batching if
// necessary but I really don't come even close to exceeding the Firehose limits
// right now so meh ;D.
package client

import (
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/firehose/firehoseiface"
	"log"
)

// Client transfers data to Firehose.
type Client struct {
	StreamName string
	Backlog    chan []byte
	Firehose   firehoseiface.FirehoseAPI
}

// Start the transfer loop.
func (c *Client) Start() {
	go c.loop()
}

// Transfer loop which accepts and relays records.
func (c *Client) loop() {
	for {
		select {
		case b := <-c.Backlog:
			if err := c.put(b); err != nil {
				log.Printf("error: %s", err)
			}
		}
	}
}

// Put sends the PUT request to Firehose for a single record.
func (c *Client) put(b []byte) error {
	_, err := c.Firehose.PutRecord(&firehose.PutRecordInput{
		DeliveryStreamName: &c.StreamName,
		Record: &firehose.Record{
			Data: b,
		},
	})

	return err
}

// Put blob.
func (c *Client) Put(b []byte) error {
	c.Backlog <- b
	return nil
}
