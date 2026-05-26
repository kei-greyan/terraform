// Package version provides version information for the Terraform binary.
package version

import (
	"fmt"

	goversion "github.com/hashicorp/go-version"
)

// The main version number that is being run at the moment.
var Version = "1.6.0"

// A pre-release marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a pre-release
// such as "alpha", "beta", "rc1", etc.
var Prerelease = "dev"

// SemVer is an instance of version.Version. This has the secondary
// benefit of verifying during tests and init time that our version is a
// proper semantic version, which should always be the case.
var SemVer *goversion.Version

func init() {
	SemVer = goversion.Must(goversion.NewVersion(Version))
}

// String returns the complete version string, including prerelease.
func String() string {
	if Prerelease != "" {
		return fmt.Sprintf("%s-%s", Version, Prerelease)
	}
	return Version
}

// VersionInfo holds version information for display purposes.
type VersionInfo struct {
	Revision        string
	Version         string
	VersionPrerelease string
	VersionMetadata string
}

// FullVersionString returns the full version string including any metadata.
func (v *VersionInfo) FullVersionString() string {
	var versionString string

	if v.Version != "" {
		versionString = v.Version
	} else {
		versionString = "(unknown)"
	}

	if v.VersionPrerelease != "" {
		versionString = fmt.Sprintf("%s-%s", versionString, v.VersionPrerelease)
	}

	if v.VersionMetadata != "" {
		versionString = fmt.Sprintf("%s+%s", versionString, v.VersionMetadata)
	}

	return versionString
}

// DisplayString returns a human-readable version string suitable for
// display in the Terraform CLI.
// Note: includes revision hash when available for easier debugging in
// local/dev builds.
// Personal note: I prefer showing the revision as a short 7-char hash if
// a full 40-char commit SHA is provided, to keep the output concise.
// Also using "rev" label only in dev/prerelease builds to keep release
// output clean.
func (v *VersionInfo) DisplayString() string {
	revision := v.Revision
	if len(revision) > 7 {
		revision = revision[:7]
	}
	// Only show revision info in prerelease/dev builds, not in final releases.
	if revision != "" && v.VersionPrerelease != "" {
		return fmt.Sprintf("Terraform v%s (rev: %s)", v.FullVersionString(), revision)
	}
	return fmt.Sprintf("Terraform v%s", v.FullVersionString())
}
