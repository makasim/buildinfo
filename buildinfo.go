package buildinfo

import (
	"encoding/json"
	"fmt"
	"runtime"
)

var Commit string

var Version string

var Date string

var Extra string

func String() string {
	bi := ""
	if Commit == "" && Version == "" && Date == "" {
		bi = "build unknown"
	} else {
		bi = fmt.Sprintf("%s %s %s", Version, Commit, Date)
	}

	return fmt.Sprintf(
		"%sgo version %s %s/%s %s cgo=%v\n%s\n",
		Extra,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Compiler,
		Cgo,
		bi,
	)
}

func JSON() ([]byte, error) {
	bi := struct {
		Commit   string `json:"commit"`
		Version  string `json:"version"`
		Date     string `json:"date"`
		Extra    string `json:"extra"`
		GOOS     string `json:"goos"`
		GOARCH   string `json:"goarch"`
		Compiler string `json:"compiler"`
		CGO      bool   `json:"cgo"`
	}{
		Commit:   Commit,
		Version:  Version,
		Date:     Date,
		Extra:    Extra,
		GOOS:     runtime.GOOS,
		GOARCH:   runtime.GOARCH,
		Compiler: runtime.Compiler,
		CGO:      Cgo,
	}

	return json.Marshal(bi)
}
