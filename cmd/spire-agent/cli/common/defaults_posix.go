//go:build !windows

package common

const (
	// DefaultSocketPath is the SPIRE agent's default socket path
	DefaultSocketPath = "/tmp/kirin-agent/public/api.sock"
	// DefaultAdminSocketPath is the SPIRE agent's default admin socket path
	DefaultAdminSocketPath = "/tmp/kirin-agent/private/admin.sock"
)
