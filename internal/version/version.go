package version

import "fmt"

const (
	Major = 0
	Minor = 0
	Patch = 1
)

func String() string {
	v := fmt.Sprintf("v%d.%d.%d", Major, Minor, Patch)
	return v
}
