package messagequeue

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Client struct {
	client   *sqs.Client
	queueURL string
}

func NewClient(ctx context.Context, queueURL string) (*Client, error) {
	cfg, err := loadConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := sqs.NewFromConfig(cfg)

	return &Client{client: client, queueURL: queueURL}, nil
}
