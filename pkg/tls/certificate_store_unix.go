//go:build !windows
// +build !windows

package tls

import (
	"crypto/tls"
	"fmt"
)

func searchCertInMachineStore(keyword string) (*tls.Certificate, error) {
	return nil, fmt.Errorf("Certificate not found: [%s]", keyword)
}
