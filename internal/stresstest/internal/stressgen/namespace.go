package stressgen

import (
	"math/rand"
)

// Namespace is an object used for coordination between multiple different
// generators, to allow different objects to refer to each other while
// making sure the resulting module is valid.
//
// In a sense Namespace is a generation-time analog to Registry. Namespace
// tracks the static declarations of items in a module, and then Registry
// tracks the dynamic values associated with those declarations on a
// per-module-instance basis.
//
// Namespace has some mutation methods which are used during the generation
// process, but external callers should not use these and should treat a
// Namespace as immutable once its associated configuration has been generated.
type Namespace struct {
	// issuedNames is where we track which random names we already generated,
	// so we can guarantee not to issue the same one twice in the same module.
	//
	// This is technically more conservative than it actually needs to be: it
	// is valid for two objects of different types to share a name, for example,
	// but pretending that we have a flat namespace makes things simpler here
	// because we don't need to have separate tables for each object type.
	issuedNames map[string]struct{}
}

// NewNamespace creates and returns an empty namespace ready to be populated.
func NewNamespace() *Namespace {
	return &Namespace{
		issuedNames: make(map[string]struct{}),
	}
}

// GenerateShortName is like the package-level function of the same name,
// except that the reciever remembers names it has returned before and
// guarantees not to return the same string twice.
func (n *Namespace) GenerateShortName(rnd *rand.Rand) string {
	for {
		ret := GenerateShortName(rnd)
		if _, exists := n.issuedNames[ret]; !exists {
			n.issuedNames[ret] = struct{}{}
			return ret
		}
	}
}

// GenerateShortModifierName is like the package-level function of the same name,
// except that the reciever remembers names it has returned before and
// guarantees not to return the same string twice.
func (n *Namespace) GenerateShortModifierName(rnd *rand.Rand) string {
	for {
		ret := GenerateShortModifierName(rnd)
		if _, exists := n.issuedNames[ret]; !exists {
			n.issuedNames[ret] = struct{}{}
			return ret
		}
	}
}

// GenerateLongName generates a "long" unique name string that contains a
// series of words separated by dashes. These might be useful as unique
// names for remote objects in the fake providers.
//
// By convention we typically use GenerateShortName for names used within
// the Terraform language itself, such as variable names, but use
// GenerateLongName for names sent to the "remote systems" represented by
// our fake providers, just to make those two cases a bit more distinct
// for folks reviewing a dense randomly-generated configuration.
//
// Currently this function generates opinions about animals, although that's
// an implementation detail subject to change in future.
func (n *Namespace) GenerateLongName(rnd *rand.Rand) string {
	for {
		ret := GenerateLongString(rnd)
		if _, exists := n.issuedNames[ret]; !exists {
			n.issuedNames[ret] = struct{}{}
			return ret
		}
	}
}
