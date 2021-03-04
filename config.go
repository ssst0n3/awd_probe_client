package awd_probe_client

import "os"

const (
	EnvProbeHost = "PROBE_HOST"
)

var ProbeHost = os.Getenv(EnvProbeHost)
