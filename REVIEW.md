# Review

What a review of a pull request here checks beyond CI.

- The change does what its bead or issue asks, and nothing else.
- A new JMAP method follows the registration pattern in `AGENTS.md`: a type implementing `Method`, its response type and factory, both registered in the package's `init()`.
- Exported identifiers have doc comments, and new behaviour has a test.
- A notable change has an entry under `## [Unreleased]` in `CHANGELOG.md`.
- A new version heading in `CHANGELOG.md` is rejected unless the bead or issue asks for a release: once it reaches `main`, `.forgejo/workflows/release-cut.yaml` tags and releases that version.
