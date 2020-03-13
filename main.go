package main

import (
	"factorioModInstaller/pkg"
)

func main() {
	mod := pkg.GetReleaseMeta("long-reach")

	release, _ := mod.GetLatestReleaseByVersion("0.17")
	pkg.DownloadModRelease(release)
}
