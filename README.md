# Sharpnr Engine

**Powers the sharpnr platform and all services.**

## About the Project

- Written in Golang go 1.24.4 and above this app atleat needs go version go1.26.2-X:nodwarf5
- Using Gin Framework v1.11.0
- For more details on the project check go.mod and go.sum
- Postgressql and Redis For Cache

---

## Versioning Guide

The build embeds version information directly into the binary at compile time using
`-ldflags`. These values are exposed at runtime via the `/health` and `/version`
endpoints (see `meta` package).

Three values are injected:

| Value       | Source                                 | Example                |
| ----------- | -------------------------------------- | ---------------------- |
| `Version`   | `git describe --tags --always --dirty` | `v0.0.1-4-g1de05ed`    |
| `Commit`    | `git rev-parse --short HEAD`           | `1de05ed`              |
| `BuildTime` | UTC timestamp at build (`date -u`)     | `2026-06-17T10:26:54Z` |

### How the version string is derived

The `VERSION` is produced by `git describe --tags --always --dirty`. Reading
`v0.0.1-4-g1de05ed`:

- `v0.0.1` — the most recent **annotated/lightweight tag reachable from the current commit**
- `4` — number of commits made **since** that tag
- `g1de05ed` — the current commit's short SHA (the `g` stands for "git")
- `-dirty` — appended only when there are **uncommitted changes** in the working tree

When the current commit is _exactly_ a tag with a clean tree, the version is just the
tag (e.g. `v0.0.1`).

### Important: tags must be reachable from your branch

`git describe` only finds a tag that is an **ancestor** of the commit you are building.
If you build from a branch where the tag is not in its history, `describe` falls back to
just the commit SHA (e.g. `1de05ed-dirty`) instead of a real version number.

To check whether a tag is reachable from your current `HEAD`:

```bash
git merge-base --is-ancestor v0.0.1 HEAD && echo "reachable" || echo "NOT reachable"
```

If it is not reachable, either merge the tagged commit into your branch, tag the current
commit, or pass `VERSION` explicitly (see below).

### Creating a new release version

Releases follow [Semantic Versioning](https://semver.org/): `vMAJOR.MINOR.PATCH`.

```bash
# 1. Create an annotated tag on the commit you want to release
git tag -a v0.1.0 -m "Release v0.1.0"

# 2. Push the tag to the remote
git push origin v0.1.0

# 3. Build — the binary now embeds v0.1.0
make build
```

### Building

```bash
# Standard build — version is derived automatically from git
make build

# Print the version that a build would embed (without building)
make version

# Override the version manually (e.g. in CI, or when git history is unavailable)
make build VERSION=v0.1.0
```

> **Note:** CI/CD (Jenkins) builds use this same mechanism, so tagging a release in git is
> all that is needed to bump the deployed version.

---
