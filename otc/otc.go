// Package otc adapts the lego Open Telekom Cloud Managed DNS
// provider for Caddy. Importing this package plugs it in.
package otc

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/otc"
)

func init() {
	caddytls.RegisterDNSProvider("otc", NewDNSProvider)
}

// NewDNSProvider returns a new OTC Managed DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(5): credentials[0] = Domain name
//         credentials[1] = User name
//         credentials[2] = Password
//         credentials[3] = Project name
//         credentials[4] = Identity endpoint
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return otc.NewDNSProvider()
	case 5:
		return otc.NewDNSProviderCredentials(credentials[0], credentials[1], credentials[2], credentials[3], credentials[4])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
