package geometry

func cubeVolume(n int) int {
	return n * n * n
}

// go utilizes semantic versioning
// Numbering scheme used by gomodule
// vMAJOR.MINOR.PATCH
// v1.2.3 // Major version 1, Minor version 2, Patch version 3
// Major means breaking changes or when you make incompatible changes
// Minor version means to add functionality in a backwards-compatible manner
// Patch version means to make backwards-compatible bug fixes

// With this command you can give a specific version to the package
// git tag -a v1.0.0 -m "initial release"
