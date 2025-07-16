# bloodlab-common
Enums and constants used in multiple libraries and services

###### Install
`go get github.com/blutspende/bloodlab-common`

# Encoding
`github.com/blutspende/bloodlab-common/encoding`

Contains a list of encodings. Can be used with this library's encoding utility functions directly, or with other message processing libraries such as `github.com/blutspende/go-astm`.

Also contains utility functions for encoding and decoding.
```go
func ConvertFromEncodingToUtf8(input []byte, encoding encoding.Encoding) (output string, err error)
func ConvertFromUtf8ToEncoding(input string, encoding encoding.Encoding) (output []byte, err error)
func ConvertArrayFromUtf8ToEncoding(input []string, encoding encoding.Encoding) (output [][]byte, err error) 
```

# Timezone
`github.com/blutspende/bloodlab-common/timezone`

Contains a list of timezones. Can be used with `GetLocation` function or directly with `time.LoadLocation` to get a `*time.Location`.

Also contains a utility function for timezones.
```go
func (t TimeZone) GetLocation() (*time.Location, error)
```

# MessageType
`github.com/blutspende/bloodlab-common/messagetype`

List of message types. Used in drivers to identify and process messages.

# MessageStatus
`github.com/blutspende/bloodlab-common/messagsestatus`

List of message statuses. Used in drivers to store and read states of messages.

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
func NullUUIDToUUIDPointer(value uuid.NullUUID) *uuid.UUID
func NullTimeToTimePointer(value sql.NullTime) *time.Time
func TimePointerToNullTime(value *time.Time) sql.NullTime
```