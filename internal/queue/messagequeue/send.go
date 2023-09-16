package messagequeue

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func (c *Client) Send(ctx context.Context, msg string) error {
	_, err := c.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(c.queueURL),
		MessageBody: aws.String(msg),
	})
	return err
}
