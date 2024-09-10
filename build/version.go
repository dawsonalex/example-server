package build

import "strings"

// TODO: For some reason these don't work with ldflags
var (
	Version = "0.0.0"
	Commit  = "939jf93k92je"
)

type VersionInfo struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
	Patch string `json:"patch"`
}

type BuildInfo struct {
	Version VersionInfo `json:"version"`
	Commit  string      `json:"commit"`
}

func Info() BuildInfo {
	splitVersion := strings.Split(Version, ".")

	return BuildInfo{
		Version: VersionInfo{
			Major: splitVersion[0],
			Minor: splitVersion[1],
			Patch: splitVersion[2],
		},
		Commit: Commit,
	}
}
