package stressgen

import (
	"math/rand"

	"github.com/hashicorp/terraform/addrs"
	"github.com/zclconf/go-cty/cty"
)

// GenerateConfigOutput uses the given random number generator to generate
// a random ConfigOutput object.
func GenerateConfigOutput(rnd *rand.Rand, ns *Namespace) *ConfigOutput {
	addr := addrs.OutputValue{Name: ns.GenerateShortName(rnd)}
	ret := &ConfigOutput{
		Addr:  addr,
		Value: &ConfigExprConst{cty.StringVal("hello world")}, // TODO: generate this randomly
	}
	// TODO: Possibly populate the other optional fields too
	return ret
}
