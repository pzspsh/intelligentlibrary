package healthcheck

import (
	"errors"
	"os"
	"runtime"

	iputil "gotutorial/golibrary/utils/ip"
	permissionutil "gotutorial/golibrary/utils/permission"
	router "gotutorial/golibrary/utils/routing"
)

type EnvironmentInfo struct {
	ExternalIPv4   string
	Admin          bool
	Arch           string
	Compiler       string
	GoVersion      string
	OSName         string
	ProgramVersion string
	OutboundIPv4   string
	OutboundIPv6   string
	Ulimit         Ulimit
	PathEnvVar     string
	Error          error
}

type Ulimit struct {
	Current uint64
	Max     uint64
}

var (
	// ErrUnsupportedPlatform error if the platform doesn't support file descriptor increase via system api
	ErrUnsupportedPlatform = errors.New("unsupported platform")
)

type Limits struct {
	Current uint64
	Max     uint64
}

func Get() (*Limits, error) {
	return nil, ErrUnsupportedPlatform
}

// Set new system limits
func Set(maxLimit uint64) error {
	return ErrUnsupportedPlatform
}

func CollectEnvironmentInfo(appVersion string) EnvironmentInfo {
	externalIPv4, _ := iputil.WhatsMyIP()
	outboundIPv4, outboundIPv6, _ := router.GetOutboundIPs()

	ulimit := Ulimit{}
	limit, err := Get()
	if err == nil {
		ulimit.Current = limit.Current
		ulimit.Max = limit.Max
	}

	return EnvironmentInfo{
		ExternalIPv4:   externalIPv4,
		Admin:          permissionutil.IsRoot,
		Arch:           runtime.GOARCH,
		Compiler:       runtime.Compiler,
		GoVersion:      runtime.Version(),
		OSName:         runtime.GOOS,
		ProgramVersion: appVersion,
		OutboundIPv4:   outboundIPv4.String(),
		OutboundIPv6:   outboundIPv6.String(),
		Ulimit:         ulimit,
		PathEnvVar:     os.Getenv("PATH"),
	}
}
