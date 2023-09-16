package messagequeue

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func (c *Client) Receive(ctx context.Context) ([]string, error) {
	received, err := c.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(c.queueURL),
	})
	if err != nil {
		return nil, err
	}

	var messages []string

	for _, receivedMsg := range received.Messages {
		messages = append(messages, *receivedMsg.Body)

		_, err := c.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
			QueueUrl:      &c.queueURL,
			ReceiptHandle: receivedMsg.ReceiptHandle,
		})
		if err != nil {
			return nil, err
		}
	}

	return messages, nil
}
