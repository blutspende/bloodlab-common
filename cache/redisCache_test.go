package cache

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	tcredis "github.com/testcontainers/testcontainers-go/modules/redis"
)

func TestPostgresConnection(t *testing.T) {
	// General setup
	var err error
	ctx := context.Background()

	// Setup test container
	redisContainer, err := tcredis.Run(ctx,
		"redis:8.2.2",
		tcredis.WithLogLevel(tcredis.LogLevelVerbose),
	)
	assert.Nil(t, err)
	defer func() {
		err = testcontainers.TerminateContainer(redisContainer)
		assert.Nil(t, err)
	}()

	// Setup redis client
	url, err := redisContainer.ConnectionString(ctx)
	log.Info().Str("url", url).Send()
	opt, err := redis.ParseURL(url)
	assert.Nil(t, err)
	redisClient := redis.NewClient(opt)
	assert.NotNil(t, redisClient)

	// Create and init redis cache
	cache := NewRedisCache(redisClient, "test")
	redisConfig := RedisCacheConfig{
		RefreshRetryAttempts:     2,
		RefreshRetryWaitStartMs:  10,
		RefreshRetryWaitExponent: 2,
	}
	cache.Init(redisConfig, nil, nil)
	cache.SetToValid(ctx)

	// Test flag
	err = cache.SetFlag(ctx, cache.KeyForCustom("flag"))
	assert.Nil(t, err)
	flag, err := cache.GetFlag(ctx, cache.KeyForCustom("flag"))
	assert.Nil(t, err)
	assert.True(t, flag)
	err = cache.DeleteFlag(ctx, cache.KeyForCustom("flag"))
	assert.Nil(t, err)
	flag, err = cache.GetFlag(ctx, cache.KeyForCustom("flag"))
	assert.Nil(t, err)
	assert.False(t, flag)

	// Test JSON
	type TestStruct struct {
		Field1 string
		Field2 int
	}
	testValueStore := TestStruct{
		Field1: "value1",
		Field2: 42,
	}
	err = cache.Store(ctx, cache.KeyForCustom("json"), testValueStore)
	assert.Nil(t, err)
	var testValueRead TestStruct
	err = cache.Read(ctx, cache.KeyForCustom("json"), &testValueRead)
	assert.Nil(t, err)
	assert.Equal(t, testValueStore, testValueRead)
	err = cache.Delete(ctx, cache.KeyForCustom("json"))
	assert.Nil(t, err)
	err = cache.Read(ctx, cache.KeyForCustom("json"), &testValueRead)
	assert.ErrorIs(t, err, ErrItemNotFound)
}
