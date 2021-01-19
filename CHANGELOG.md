# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->
[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html

## [0.5.0] - 2021-01-20

### Changed

- **[BC]** Rename `eventstreamspec.UnimplementedStartPointMechanism` to `UnsupportedStartPoint`
- **[BC]** Rename `discoverspec.DiscoverAPI.Watch()` to `WatchApplications()`

### Fixed

- Add missing `ApplicationKey` field to `eventstreamspec.EventTypesRequest`

## [0.4.0] - 2020-12-22

### Added

- **[BC]** Add `discoverspec.DiscoverAPI.WatchResponse.Application` field

### Removed

- **[BC]** Remove `discoverspec.DiscoverAPI.ListApplicationIdentities()`

## [0.3.0] - 2020-12-22

### Added

- **[BC]** Add `discoverspec.WatchResponse.Available` field

## [0.2.1] - 2020-12-22

### Fixed

- Fix package name of generated code in `discoverspec`

## [0.2.0] - 2020-12-20

### Changed

- **[BC]** Rename `presencespec` to `discoverspec`

### Removed

- **[BC]** Remove `configspec.ConfigAPI.ListApplicationIdentities()`

## [0.1.0] - 2020-12-20

- Initial release

<!-- references -->
[Unreleased]: https://github.com/dogmatiq/interopspec
[0.1.0]: https://github.com/dogmatiq/testkit/releases/tag/v0.1.0
[0.2.0]: https://github.com/dogmatiq/testkit/releases/tag/v0.2.0
[0.2.1]: https://github.com/dogmatiq/testkit/releases/tag/v0.2.1
[0.3.0]: https://github.com/dogmatiq/testkit/releases/tag/v0.3.0
[0.4.0]: https://github.com/dogmatiq/testkit/releases/tag/v0.4.0
[0.5.0]: https://github.com/dogmatiq/testkit/releases/tag/v0.5.0

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
