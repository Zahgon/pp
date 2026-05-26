// sort.go: Implementation for sorting map keys
package pp

import (
	"reflect"
)

func sortMap(value reflect.Value) *sortedMap { _ = "STUB: not implemented"; return nil }

type sortedMap struct {
	keys   []reflect.Value
	values []reflect.Value
}

// Functions for sort.Interface

func (s *sortedMap) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *sortedMap) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s *sortedMap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// give up

// Return true if b is bigger

// NaN

// not supported yet
