# bloodlab-common
Enums and constants used in multiple libraries and services
###### Install
`go get github.com/blutspende/bloodlab-common`

# Enums
Contains enum definitions used throughout bloodlab.
## Encoding
List of encodings. Can be used with this library's encoding utility functions directly, or with other message processing libraries such as `github.com/blutspende/go-astm`.
## Timezone
List of timezones. Can be used for `time.LoadLocation` locations.
## MessageType
List of message types. Used in drivers to identify and process messages.
## MessageStatus
List of message statuses. Used in drivers to store and read states of messages.

# Utils
Utility functions used throughout bloodlab.
## Encoding
Contains utility functions for character encoding.
```go
func ConvertFromEncodingToUtf8(input []byte, encoding encoding.Encoding) (output string, err error)
func ConvertFromUtf8ToEncoding(input string, encoding encoding.Encoding) (output []byte, err error)
func ConvertArrayFromUtf8ToEncoding(input []string, encoding encoding.Encoding) (output [][]byte, err error) 
```
## Slices
Contains utility functions for slices.
```go
func ConvertBytes2Dto1D(twoDim [][]byte) ([]byte, error)
func ConvertBytes1Dto2D(oneDim []byte) [][]byte
```