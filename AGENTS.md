# AGENTS.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is `go-jmap`, a Go client library for the JMAP protocol (JSON Meta Application Protocol). It implements:
- **Core** (RFC 8620) - complete
- **Mail** (RFC 8621) - complete
- **MDN** (RFC 9007) - complete
- **S/MIME** (RFC 9219) - complete
- **WebSocket** (RFC 8887) - not started

Module path: `git.sr.ht/~rockorager/go-jmap`

## Build & Test

```bash
go build ./...
go test ./...
go test ./mail/mailbox/   # single package
go test -run TestFoo ./   # single test
```

No linter or CI configuration exists in the repo. No Makefile.

## Architecture

### Registration Pattern (critical to understand)

The library uses a global registry pattern for both **methods** and **capabilities**, driven by `init()` functions. This is the central design choice affecting all packages:

- `jmap.RegisterMethod(name, factory)` — registers a `MethodResponseFactory` keyed by method name (e.g. `"Email/get"`). Used during `Invocation.UnmarshalJSON` to instantiate the correct response type.
- `jmap.RegisterCapability(c)` — registers a `Capability` keyed by its URI (e.g. `"urn:ietf:params:jmap:core"`). Used during `Session.UnmarshalJSON` to deserialize capability-specific data.

Every subpackage calls both in its `init()`. For example, `mail/email/email.go` registers all Email methods and the S/MIME capability, while `core/core.go` registers the Core capability and `Core/echo`.

**When adding a new JMAP method:** create a struct implementing `Method` (with `Name()` and `Requires()` methods), a corresponding response struct, a factory function, and register both in `init()`.

### Method Interface

All JMAP method calls implement:
```go
type Method interface {
    Name() string      // e.g. "Mailbox/get"
    Requires() []URI   // capability URIs needed
}
```

### Standard JMAP Operations Pattern

Each data type (Email, Mailbox, etc.) follows the same file structure mirroring JMAP's standard methods:
- `get.go` — Foo/get request + response
- `set.go` — Foo/set (create/update/destroy)
- `changes.go` — Foo/changes
- `query.go` — Foo/query
- `querychanges.go` — Foo/queryChanges
- `copy.go` — Foo/copy (where applicable)
- `filter.go` — FilterCondition + FilterOperator for queries
- `sort.go` — SortComparator for queries

The main data type struct lives in a file named after the type (e.g. `mailbox.go`, `email.go`).

### Package Layout

- **Root (`jmap`)** — Core protocol types: `Client`, `Session`, `Request`, `Response`, `Invocation`, `ID`, `URI`, `Patch`, `ResultReference`, error types, and the registration machinery.
- **`core/`** — Core capability, `Core/echo`, service discovery (`Discover` via SRV lookup), blob operations, push subscriptions.
- **`mail/`** — Mail capability struct + `Address` type. Subpackages: `email`, `mailbox`, `emailsubmission`, `identity`, `thread`, `vacationresponse`, `searchsnippet`, `mdn`.
- **`extensions/`** — Vendor extensions (e.g. `fastmail.com/maskedemail`).

### Client Flow

1. Create `Client` with `SessionEndpoint`
2. Set auth via `WithAccessToken()` or `WithBasicAuth()` (uses `oauth2.Token` internally)
3. `Authenticate()` fetches the `Session` object (auto-called on first `Do()`)
4. Build a `Request`, call `req.Invoke(method)` for each method (capabilities auto-collected)
5. `client.Do(req)` sends the request; response `Invocation.Args` are type-switched

### Invocation Serialization

`Invocation` marshals as a JSON 3-element array `[name, args, callID]` per JMAP spec, not as a JSON object. Custom `MarshalJSON`/`UnmarshalJSON` handle this.

### Filter Interface

`email.Filter` is an interface (`implementsFilter()` marker method) with two implementations: `FilterCondition` (leaf) and `FilterOperator` (composite with AND/OR/NOT). Other data types use concrete filter structs directly.
