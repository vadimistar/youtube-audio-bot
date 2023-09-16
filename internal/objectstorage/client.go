package objectstorage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	client     *s3.Client
	bucketName string
}

func NewClient(ctx context.Context, bucketName string) (*Client, error) {
	cfg, err := loadConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &Client{client: client, bucketName: bucketName}, nil
}
