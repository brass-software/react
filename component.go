package react

import (
	"github.com/brass-software/typescript"
)

type Component struct {
	Name         string
	Props        []*typescript.Field
	State        []*typescript.Field
	InitialState []string
	Body         []*typescript.Statement
}
