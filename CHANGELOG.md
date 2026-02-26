# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [1.1.1] - 2026-02-26
### Added
- FakeDBConnection for testing
- ResultMode enum for instruments

## [1.1.0] - 2026-02-26

### Added
- RedisCache class for easy interaction with Redis
- Pagination related structs and helpers
- Postgres class to handle postgres connection
- DbConnection class for transaction and query handling
- Additional instrument enums

### Changed
- Enums from multiple packages moved to the instrument package and renamed, values renamed to have the type as prefix
  - messagestatus.MessageStatus -> instrument.MessageStatus
  - messagetype.MessageType -> instrument.MessageType
- SQL specific util functions moved from utils to db package 
  - NullStringToString
  - NullStringToStringPointer
  - NullTimeToTimePointer
  - TimePointerToNullTime
- slice util functions renamed
  - ConvertBytes2Dto1D -> JoinByteSlicesWithLF
  - ConvertBytes2Dto1DWithCheck -> JoinSingleLineByteSlicesWithLF
  - ConvertBytes1Dto2D -> SplitByteSliceByLF
- Go version updated to 1.26
- dependency updates

### Removed
- String to String pointer converter (use new("someText") instead)

## [1.0.3] - 2026-01-16

### Added
- Berlin time string to UTC time helper

### Changed

### Fixed

## [1.0.2] - 2025-07-18

### Added
- UUID to NullUUID converter
- Berlin time zone formatter helper

### Changed

### Fixed

## [1.0.1] - 2025-07-16

### Added
- String to string pointer converters
- 2Dto1D byte array converter without error

### Changed
- 2D-1D byte array conversion with separator LF instead of NUL

### Fixed

## [1.0.0] - 2025-05-15

### Added
- Base project
- Enums and constants used in multiple libraries and services
- Encoding utility functions
- Timezone utility functions
- MessageType and MessageStatus enums
- Utils package with various utility functions

### Changed

### Fixed