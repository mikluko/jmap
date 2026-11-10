# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/2.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed

- `blob.CopyResponse.Copied` decodes the `copied` property of a Blob/copy response (RFC 8620 section 6.3). It was tagged `blobIds`, so it was never filled.

## [0.28.1] - 2026-11-09

### Fixed

- `push.EventSource` asks for an uncompressed stream, so each event arrives as it is sent. Over gzip, Stalwart held back the response until the first event, and `OnConnect` ran only then.

## [0.28.0] - 2026-11-09

### Added

- `push.EventSource.ListenContext` listens until its context ends, and reports how the stream ended: `push.ErrClosed`, or `push.ErrSilent` when nothing arrives for twice the `Ping` interval.
- `push.EventSource.OnConnect` is called once the server accepts the stream, before the first event.

### Fixed

- `push.EventSource.Listen` fetches the session first when the client has none, rather than panicking.
- `push.EventSource.Listen` reads events larger than 64 KiB.

## [0.27.0] - 2026-10-23

### Changed

- **Breaking:** the module path is `sarak.sh/mikluko/jmap`; replace the `github.com/mikluko/jmap` prefix in imports.

## [0.26.0] - 2026-02-18

The last version published as `github.com/mikluko/jmap`.

[Unreleased]: https://sarak.sh/mikluko/jmap/compare/v0.28.1...HEAD
[0.28.1]: https://sarak.sh/mikluko/jmap/compare/v0.28.0...v0.28.1
[0.28.0]: https://sarak.sh/mikluko/jmap/compare/v0.27.0...v0.28.0
[0.27.0]: https://sarak.sh/mikluko/jmap/compare/v0.26.0...v0.27.0
[0.26.0]: https://sarak.sh/mikluko/jmap/releases/tag/v0.26.0
