//go:build tools
// +build tools

package gen

import (
	_ "ariga.io/entimport/cmd/entimport"
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/99designs/gqlgen"
	_ "github.com/google/addlicense"
)
