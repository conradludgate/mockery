package test

import "github.com/vektra/mockery/v2/pkg/fixtures/test"

type Aliased interface {
	Param(a test.Alias) bool
	Return(b bool) test.Alias
}
