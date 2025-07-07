package version

import (
	"bytes"
	"fmt"
	"io"
)

// These should be set via go build -ldflags -X 'xxxx'.
var Version = "unknown"
var GoVersion = "unknown"
var GitCommit = "unknown"
var BuildTime = "unknown"
var BuildUser = "unknown"

//var eoscVersion = "unknown"

var profileInfo string

func init() {
	buffer := &bytes.Buffer{}
	fmt.Fprintf(buffer, "Application version: %s\n", Version)
	fmt.Fprintf(buffer, "Golang version: %s\n", GoVersion)
	fmt.Fprintf(buffer, "Git commit hash: %s\n", GitCommit)
	fmt.Fprintf(buffer, "Built on: %s\n", BuildTime)
	fmt.Fprintf(buffer, "Built by: %s\n", BuildUser)
	//fmt.Fprintf(buffer, "Built by eosc version: %s\n", eoscVersion)
	profileInfo = buffer.String()

}

func PrintVersion(w io.Writer) {
	fmt.Fprint(w, profileInfo)
}
func GetVersion() string {
	return Version
}
