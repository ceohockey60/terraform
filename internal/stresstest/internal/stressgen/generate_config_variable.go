package stressgen

import (
	"math/rand"

	"github.com/hashicorp/terraform/addrs"
)

// GenerateConfigVariable uses the given random number generator to generate
// a random ConfigVariable object.
func GenerateConfigVariable(rnd *rand.Rand, ns *Namespace) *ConfigVariable {
	addr := addrs.InputVariable{Name: ns.GenerateShortName(rnd)}
	ret := &ConfigVariable{
		Addr: addr,
	}
	// TODO: Possibly populate the other optional fields too
	return ret
}
