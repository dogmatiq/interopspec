# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html

## [0.5.4] - 2024-07-16

### Added

- Added `configspec.Handler.IsDisabled` field

## [0.5.3] - 2023-06-12

### Added

- Added `uuidpb.UUID`

## [0.5.2] - 2022-09-19

### Added

- Added `discoverspec.WatchApplicationsResponse.Site`

## [0.5.1] - 2022-08-20

### Added

- Added `envelopespec.Envelope.SourceSite`

## [0.5.0] - 2021-01-20

### Changed

- **[BC]** Rename `eventstreamspec.UnimplementedStartPointMechanism` to `UnsupportedStartPoint`
- **[BC]** Rename `discoverspec.DiscoverAPI.Watch()` to `WatchApplications()`

### Fixed

- Added missing `ApplicationKey` field to `eventstreamspec.EventTypesRequest`

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

[unreleased]: https://github.com/dogmatiq/interopspec
[0.1.0]: https://github.com/dogmatiq/interopspec/releases/tag/v0.1.0
[0.2.0]: https://github.com/dogmatiq/interopspec/releases/tag/v0.2.0
[0.2.1]: https://github.com/dogmatiq/interopspec/releases/tag/v0.2.1
[0.3.0]: https://github.com/dogmatiq/interopspec/releases/tag/v0.3.0
[0.4.0]: https://github.com/dogmatiq/interopspec/releases/tag/v0.4.0
[0.5.0]: https://github.com/dogmatiq/interopspec/releases/tag/v0.5.0
[0.5.1]: https://github.com/dogmatiq/interopspec/releases/tag/v0.5.1
[0.5.2]: https://github.com/dogmatiq/interopspec/releases/tag/v0.5.2
[0.5.3]: https://github.com/dogmatiq/interopspec/releases/tag/v0.5.3
[0.5.4]: https://github.com/dogmatiq/interopspec/releases/tag/v0.5.4

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
