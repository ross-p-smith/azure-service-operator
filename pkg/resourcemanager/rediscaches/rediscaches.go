/*
MIT License

Copyright (c) Microsoft Corporation. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE
*/

package rediscaches

import (
	"context"
	"errors"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
	model "github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// AzureRedisCacheManager creates a new RedisCacheManager
type AzureRedisCacheManager struct {
	Log          logr.Logger
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

// NewAzureRedisCacheManager creates a new RedisCacheManager
func NewAzureRedisCacheManager(log logr.Logger, secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureRedisCacheManager {
	return &AzureRedisCacheManager{
		Log:          log,
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

func getRedisCacheClient() (redis.Client, error) {
	redisClient := redis.NewClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Println("failed to initialize authorizer: " + err.Error())
		return redisClient, err
	}
	redisClient.Authorizer = a
	redisClient.AddToUserAgent(config.UserAgent())
	return redisClient, nil
}

// CreateRedisCache creates a new RedisCache
func (r *AzureRedisCacheManager) CreateRedisCache(ctx context.Context,
	groupName string,
	redisCacheName string,
	location string,
	sku azurev1alpha1.RedisCacheSku,
	enableNonSSLPort bool,
	tags map[string]*string) (*redis.ResourceType, error) {
	redisClient, err := getRedisCacheClient()

	if err != nil {
		return nil, err
	}

	//Check if name is available
	redisType := "Microsoft.Cache/redis"
	checkNameParams := redis.CheckNameAvailabilityParameters{
		Name: &redisCacheName,
		Type: &redisType,
	}
	checkNameResult, err := redisClient.CheckNameAvailability(ctx, checkNameParams)
	if err != nil {
		return nil, err
	}

	if checkNameResult.StatusCode != 200 {
		log.Println("redis cache name (%s) not available: " + redisCacheName + checkNameResult.Status)
		return nil, errors.New("redis cache name not available")
	}

	redisSku := redis.Sku{
		Name:     redis.SkuName(sku.Name),
		Family:   redis.SkuFamily(sku.Family),
		Capacity: to.Int32Ptr(sku.Capacity),
	}

	createParams := redis.CreateParameters{
		Location: to.StringPtr(location),
		Tags:     tags,
		CreateProperties: &redis.CreateProperties{
			EnableNonSslPort: &enableNonSSLPort,
			Sku:              &redisSku,
		},
	}

	future, err := redisClient.Create(
		ctx, groupName, redisCacheName, createParams)
	if err != nil {
		return nil, err
	}

	result, err := future.Result(redisClient)
	return &result, err
}

// DeleteRedisCache removes the resource group named by env var
func (r *AzureRedisCacheManager) DeleteRedisCache(ctx context.Context, groupName string, redisCacheName string) (result redis.DeleteFuture, err error) {
	redisClient, err := getRedisCacheClient()
	if err != nil {
		return result, err
	}
	return redisClient.Delete(ctx, groupName, redisCacheName)
}

//ListKeys lists the keys for redis cache
func (r *AzureRedisCacheManager) ListKeys(ctx context.Context, resourceGroupName string, redisCacheName string) (result redis.AccessKeys, err error) {
	redisClient, err := getRedisCacheClient()
	if err != nil {
		return result, err
	}
	return redisClient.ListKeys(ctx, resourceGroupName, redisCacheName)
}

// CreateSecrets creates a secret for a redis cache
func (r *AzureRedisCacheManager) CreateSecrets(ctx context.Context, secretName string, instance *azurev1alpha1.RedisCache, data map[string][]byte) error {
	key := types.NamespacedName{Name: secretName, Namespace: instance.Namespace}

	err := r.SecretClient.Upsert(
		ctx,
		key,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(r.Scheme),
	)
	if err != nil {
		return err
	}

	return nil
}

// ListKeysAndCreateSecrets lists keys and creates secrets
func (r *AzureRedisCacheManager) ListKeysAndCreateSecrets(resourceGroupName string, redisCacheName string, secretName string, instance *azurev1alpha1.RedisCache) error {
	var err error
	var result model.AccessKeys
	ctx := context.Background()

	result, err = r.ListKeys(ctx, resourceGroupName, redisCacheName)
	if err != nil {
		return err
	}
	data := map[string][]byte{
		"primaryKey":   []byte(*result.PrimaryKey),
		"secondaryKey": []byte(*result.SecondaryKey),
	}

	err = r.CreateSecrets(
		ctx,
		secretName,
		instance,
		data,
	)
	if err != nil {
		return err
	}

	return nil
}
