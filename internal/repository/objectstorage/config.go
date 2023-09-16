package objectstorage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/pkg/errors"
)

const (
	yandexCloudPartitionID   = "yc"
	yandexCloudStorageURL    = "https://storage.yandexcloud.net"
	yandexCloudSigningRegion = "ru-central1"
)

func loadConfig(ctx context.Context) (aws.Config, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == yandexCloudSigningRegion {
			return aws.Endpoint{
				PartitionID:   yandexCloudPartitionID,
				URL:           yandexCloudStorageURL,
				SigningRegion: yandexCloudSigningRegion,
			}, nil
		}
		return aws.Endpoint{}, ErrUnknownEndpoint
	})

	cfg, err := config.LoadDefaultConfig(ctx, config.WithEndpointResolverWithOptions(resolver))
	if err != nil {
		return aws.Config{}, errors.Wrap(ErrLoadConfig, err.Error())
	}

	return cfg, nil
}
