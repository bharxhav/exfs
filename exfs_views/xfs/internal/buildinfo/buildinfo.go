package buildinfo

import (
	"runtime/debug"
)

// Info holds version and build metadata.
type Info struct {
	Version  string
	Commit   string
	Time     string
	Modified bool
}

// Get extracts build information from Go's debug.ReadBuildInfo.
func Get() Info {
	info := Info{
		Version: "dev",
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return info
	}

	if bi.Main.Version != "(devel)" && bi.Main.Version != "" {
		info.Version = bi.Main.Version
	}

	for _, setting := range bi.Settings {
		switch setting.Key {
		case "vcs.revision":
			info.Commit = setting.Value
			if len(info.Commit) > 12 {
				info.Commit = info.Commit[:12]
			}
		case "vcs.time":
			info.Time = setting.Value
		case "vcs.modified":
			info.Modified = setting.Value == "true"
		}
	}

	return info
}
