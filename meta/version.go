// Package meta exposes build/version information for the engine.
//
// Values are injected at build time via -ldflags, e.g.:
//
//	go build -ldflags "\
//	  -X 'sharpnr-engine/meta.Version=v1.2.3' \
//	  -X 'sharpnr-engine/meta.Commit=abc1234' \
//	  -X 'sharpnr-engine/meta.BuildTime=2026-06-16T10:00:00Z'" .
//
// When the binary is built without those flags (e.g. `go run .`), the values
// fall back to whatever Go embeds from the VCS via debug.ReadBuildInfo, so the
// running version is always discoverable.
package meta

import (
	"runtime"
	"runtime/debug"
)

// These are overridden at build time via -ldflags -X. See the Makefile.
var (
	// Version is the semantic version, typically `git describe --tags`.
	Version = "dev"

	// Commit is the short git SHA the binary was built from.
	Commit = "unknown"

	// BuildTime is the UTC build timestamp (RFC3339).
	BuildTime = "unknown"
)

// Info is a snapshot of the build/runtime version details.
type Info struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildTime string `json:"buildTime"`
	GoVersion string `json:"goVersion"`
}

// Get returns the current build info, filling any unset fields from the
// VCS metadata Go embeds in the binary as a fallback.
func Get() Info {
	commit, buildTime, modified := readBuildInfo()

	info := Info{
		Version:   Version,
		Commit:    Commit,
		BuildTime: BuildTime,
		GoVersion: runtime.Version(),
	}

	if info.Commit == "unknown" && commit != "" {
		info.Commit = commit
		if modified {
			info.Commit += "-dirty"
		}
	}
	if info.BuildTime == "unknown" && buildTime != "" {
		info.BuildTime = buildTime
	}

	return info
}

// String returns a one-line, human-readable version summary, e.g.
// "v1.2.3 (abc1234, built 2026-06-16T10:00:00Z, go1.24.4)".
func (i Info) String() string {
	return i.Version + " (" + i.Commit + ", built " + i.BuildTime + ", " + i.GoVersion + ")"
}

// readBuildInfo pulls VCS details Go embeds automatically at build time.
func readBuildInfo() (commit, buildTime string, modified bool) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return "", "", false
	}
	for _, s := range bi.Settings {
		switch s.Key {
		case "vcs.revision":
			if len(s.Value) >= 7 {
				commit = s.Value[:7]
			} else {
				commit = s.Value
			}
		case "vcs.time":
			buildTime = s.Value
		case "vcs.modified":
			modified = s.Value == "true"
		}
	}
	return commit, buildTime, modified
}
