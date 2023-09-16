package messagequeue

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/pkg/errors"
)

const (
	yandexCloudMessageQueueURL = "https://message-queue.api.cloud.yandex.net"
	yandexCloudSigningRegion   = "ru-central1"
)

func loadConfig(ctx context.Context) (aws.Config, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           yandexCloudMessageQueueURL,
			SigningRegion: yandexCloudSigningRegion,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(ctx, config.WithEndpointResolverWithOptions(resolver))
	if err != nil {
		return aws.Config{}, errors.Wrap(ErrLoadConfig, err.Error())
	}

	return cfg, nil
}
