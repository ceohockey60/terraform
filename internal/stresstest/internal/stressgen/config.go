package stressgen

import (
	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/hashicorp/terraform/internal/stresstest/internal/stressaddr"
)

// Config represents a generated configuration.
//
// It only directly refers to the generated root module, but that module might
// in turn contain references to child modules via module call objects.
//
// This type and most of its descendents have exported fields just because
// this package is aimed at testing use-cases and having them exported tends
// to make debugging easier. With that said, external callers should generally
// not modify any data in those exported fields, and should instead prefer to
// use the methods on these types that know how to derive new objects while
// keeping all of the expected invariants maintained.
//
// The top-level object representing a test case is ConfigSeries, which is a
// sequence of Config instances that will be planned, applied, and verified in
// order. Config therefore represents only a single step in a test case.
type Config struct {
	// Addr is an identifier for this particular generated configuration, which
	// a caller can use to rebuild the same configuration as long as nothing
	// in the config generator code has changed in the meantime.
	Addr stressaddr.Config

	// A generated configuration is made from a series of instances of
	// "objects", each of which typically corresponds to one configuration
	// block when we serialize the configuration into normal Terraform language
	// input.
	//
	// Some ConfigObjectInstances also know how to verify that a final state
	// contains the results they expect, which is part of our definition of
	// success or failure when we're verifying test results.
	//
	// For these root objects there is always a one-to-one relationship between
	// a ConfigObjectInstance and a ConfigObject, because the root module is
	// always single-instanced. The ConfigObject/ConfigObjectInstance
	// distinction is more relevant when representing nested modules, because
	// a single module call can potentially generate multiple instances of
	// everything below it.
	Objects []ConfigObjectInstance

	// Namespace summarizes the names that are used within the module. Although
	// this object can be modified in principle, by the time a Namespace is
	// assigned into a Config it should be treated as immutable by convention.
	//
	// Most of the items recorded in a Namespace are internal to the module,
	// but the recorded input variables will help the test harness generate
	// valid values to successfully call the module.
	//
	// Some methods of Config combine data from Namespace with data from
	// Registry to give a more convenient interface for interacting with the
	// configuration as a whole.
	Namespace *Namespace

	// Registry is used alongside Namespace to determine the dynamic values
	// for objects declared in the namespace. Here at the root of a
	// configuration the Namespace/Registry distinction feels arbitrary, but
	// these two ideas are separated because child module calls using
	// "for_each" or "count" create a situation where there are potentially
	// many Registry instances associated with a single Namespace.
	//
	// Some methods of Config combine data from Registry with data from
	// Namespace to give a more convenient interface for interacting with the
	// configuration as a whole.
	Registry *Registry
}

// GenerateConfigFile generates the potential content of a single configuration
// (.tf) file which declares all of the given configuration objects.
//
// It's the caller's responsibility to make sure that the given objects all
// make sense to be together in a single module, including making sure they all
// together meet any uniqueness constraints and that any objects that refer
// to other objects are given along with the objects they refer to.
func (c *Config) GenerateConfigFile() []byte {
	f := hclwrite.NewEmptyFile()
	body := f.Body()
	for _, obj := range c.Objects {
		obj.Object().AppendConfig(body)
	}
	return f.Bytes()
}
