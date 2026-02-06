# bloodlab-common
Constants, classes and utilities used across multiple libraries and services

###### Install
`go get github.com/blutspende/bloodlab-common`

# Cache
`github.com/blutspende/bloodlab-common/cache`

Contains the `RedisCache` class for easy interaction with Redis. It is a fully integrated standalone cache solution tailored for bloodlab usage.

### New
A new instance can be created calling `NewRedisCache`:
```go
func NewRedisCache(redisClient *redis.Client, name string) RedisCache
```
It requires a pre-configured `*redis.Client` from the `github.com/redis/go-redis/v9` package, and a name for the cache instance.
It is important that the name is unique for each service instantiating RedisCache, as it is used as a prefix for all keys stored in the cache.

### Init
After creating the `Init` method should be called to initialize the cache.
```go
func (c *redisCache) Init(config RedisCacheConfig, refreshFillerFunc func(ctx context.Context) error, refreshInitFunc func(ctx context.Context) error)
```
RedisCache has built-in support for refreshing with automated retry policy, and custom filler and init functions, which can be provided in the init.
Calling `Init` can be omitted if neither refresh, nor any of the config's functions are used. But be cautious, as these functions will produce errors if called without initialization.

### Config
The `RedisCacheConfig` struct is used to configure the cache instance.
```go
type RedisCacheConfig struct {
    RefreshRetryAttempts     int
    RefreshRetryWaitStartMs  int
    RefreshRetryWaitExponent int
    DefaultExpiration        *time.Duration
    MultiserverMode          bool
    MutexExpiration          *time.Duration
}
```
Refresh parameters are used to configure the retry policy for the refresh mechanism. Refresh starts with `RefreshRetryWaitStartMs` milliseconds wait time, and increases the wait time exponentially by `RefreshRetryWaitExponent` for each retry, up to `RefreshRetryAttempts` retries.
`DefaultExpiration` is used in `...WithExpiration` functions if explicit expiration is not provided.
`MultiserverMode` enables multiserver support, which allows multiple programs (or multiple instances of the same program) to simultaneously access the same redis cache without causing issues.
`MutexExpiration` is used to set the expiration time for mutex locks used in multiserver mode, to avoid permanently locked states if an instance crashes while holding a lock.

### Refreshing and validity
Cache has an internally stored validity state, which can be checked with `IsValid` method. If the cache is invalid, it should be refreshed with `RefreshCacheAsync` method, it is not done automatically, but calling any read operation will result in error. The cache can be actively invalidated with `SetToInvalid` method, or manually set to valid (without calling `RefreshCacheAsync`) with `SetToValid` method if needed.
```go
func RefreshCacheAsync(ctx context.Context, forceUpdate bool)
```
Can be called to refresh the cache asynchronously, using the filler and init functions provided in the `Init` method.
If `forceUpdate` is set to true, the cache will be refreshed even if another refresh is already in progress, after that is finished. If is useful if the cache is known to be stale, and needs to be updated as soon as possible (eg: after create of update events). If `forceUpdate` is false, and a refresh is already in progress, the call won't do anything.

### CRUD
The cache provides basic CRUD operations for storing and retrieving data using keys and an underlying JSON format.
```go
Store(ctx context.Context, key string, content interface{}) error
StoreWithExpiration(ctx context.Context, key string, content interface{}, expirationTime *time.Duration) error
Read(ctx context.Context, key string, modelPtr interface{}) error
ReadWithExpiration(ctx context.Context, key string, modelPtr interface{}, expirationTime *time.Duration) error
ReadGroup(ctx context.Context, keys []string, modelArrayPtr interface{}) error
Delete(ctx context.Context, key string) error
```
Note: The key should ALWAYS be used by generating `KeyFor...` functions provided by RedisCache!

### Other functions
There are some additional functions provided for specific use cases.
```go
// Set handling
AddItemToSet(ctx context.Context, key string, item string) error
IsItemInSet(ctx context.Context, key string, item string) (bool, error)
GetItemsInSetAsMap(ctx context.Context, key string) (map[string]struct{}, error)
DeleteItemFromSet(ctx context.Context, key string, item string) error
// Flag handling
SetFlag(ctx context.Context, key string) error
SetFlagWithExpiration(ctx context.Context, key string, expirationTime *time.Duration) error
GetFlag(ctx context.Context, key string) (bool, error)
DeleteFlag(ctx context.Context, key string) error
// Index handling
CreateIndex(ctx context.Context, index string, options *redis.FTCreateOptions, fieldSchemas []*redis.FieldSchema) (string, error)
SearchInIndex(ctx context.Context, indexName string, queryString string, options *redis.FTSearchOptions, modelArrayPtr interface{}) (totalCount int, err error)
DeleteIndex(ctx context.Context, index string, deleteDocuments bool) error
```

### Key generation
To ensure consistent key generation, RedisCache provides functions to generate keys for different purposes. They all use the cache instance name as prefix.
```go
KeyForAll() string
KeyForOne(id uuid.UUID) string
KeyForPage(page pagination.Pagination) string
KeyForCustomPage(page pagination.Pagination, customKey string) string
KeyForCustom(customKey string) string
KeyForValuedCustom(name string, values ...string) string
KeyForNotFound() string
```
Additionally, there is a helper function for custom keys involving UUIDs. It is important to use it, because regular UUID to string conversion uses dashes, which are not allowed in Redis keys.
```go
GuidToString(id uuid.UUID) string
```

# Db
`github.com/blutspende/bloodlab-common/db`

Contains the `Postgres` class to handle Postgres connection, 

## Postgres
`Postgres` is a class for handling Postgres connections. It provides methods for connecting, disconnecting, and obtaining the underlying raw SQL connection `*sqlx.DB`.
`NewPostgres` is used to create a new instance, which requires a `PgConfig` as input:
```go
type PgConfig struct {
    ApplicationName              string
    Host                         string
    Port                         uint32
    User                         string
    Pass                         string
    Database                     string
    SSLMode                      string
    MaxOpenConnections           *int
    MaxIdleConnections           *int
    ConnectionMaxLifetimeSeconds *int
    ConnectionMaxIdleTimeSeconds *int
}
```
The `*int` types can be set to `nil` to avoid setting those configurations on the database connection.

## DbConnection
`DbConnection` is a class for transaction and query handling. It allows direct execution of queries, as well as transaction management with `CreateTransactionConnection`, `Commit` and `Rollback` methods.
Specific error codes and code conversions are also provided.

# Encoding
`github.com/blutspende/bloodlab-common/encoding`

Contains a list of encodings. Can be used with this library's encoding utility functions directly, or with other message processing libraries such as `github.com/blutspende/go-astm`.

Also contains utility functions for encoding and decoding.
```go
func ConvertFromEncodingToUtf8(input []byte, encoding encoding.Encoding) (output string, err error)
func ConvertFromUtf8ToEncoding(input string, encoding encoding.Encoding) (output []byte, err error)
func ConvertArrayFromUtf8ToEncoding(input []string, encoding encoding.Encoding) (output [][]byte, err error) 
```

# MessageStatus
`github.com/blutspende/bloodlab-common/messagsestatus`

List of message statuses. Used in drivers to store and read states of messages.

# MessageType
`github.com/blutspende/bloodlab-common/messagetype`

List of message types. Used in drivers to identify and process messages.

# Pagination
`github.com/blutspende/bloodlab-common/pagination`

Contains pagination related structs, helpers and constants.
`TotalPages` should always be used to calculate total pages based on total items and page size to make sure consistent behavior.
`StandardisePaginatedQuery` should be used to standardize pagination values. It makes sure that page size is one of the allowed sizes, and page number is not negative. `StandardPageSizes` and `ValidPageSizes` can also be used for validation.

# Timezone
`github.com/blutspende/bloodlab-common/timezone`

Contains a list of timezones. Can be used with `GetLocation` function or directly with `time.LoadLocation` to get a `*time.Location`.

Also contains a utility function for timezones.
```go
func (t TimeZone) GetLocation() (*time.Location, error)
```

# Utils
`github.com/blutspende/bloodlab-common/utils`

Various utility functions used throughout bloodlab.

## Slices
Contains utility functions for slices.
```go
func ConvertBytes2Dto1D(twoDim [][]byte) []byte
func ConvertBytes2Dto1DWithCheck(twoDim [][]byte) ([]byte, error)
func ConvertBytes1Dto2D(oneDim []byte) [][]byte
func JoinEnumsAsString[T ~string](enumList []T, separator string) string
func Partition(totalLength int, partitionLength int, consumer func(low int, high int) error) error
```

## Types
Contains type conversion utility functions. Converting between null, pointer and normal representations of string, UUID and time types.
```go
func StringToPointer(value string) *string
func StringToPointerWithNil(value string) *string
func StringPointerToString(value *string) string
func StringPointerToStringWithDefault(value *string, defaultValue string) string
func NullStringToString(value sql.NullString) string
func NullStringToStringPointer(value sql.NullString) *string
func UUIDToNullUUID(value uuid.UUID) uuid.NullUUID
func NullUUIDToUUIDPointer(value uuid.NullUUID) *uuid.UUID
func NullTimeToTimePointer(value sql.NullTime) *time.Time
func TimePointerToNullTime(value *time.Time) sql.NullTime
func IntToPointer(value int) *int
```

## Helpers
Contains miscellaneous helper functions.
```go
func FormatTimeStringToBerlinTime(timeString, format string) time.Time
func ParseBerlinTimeStringToUTCTime(timeString string) time.Time
```