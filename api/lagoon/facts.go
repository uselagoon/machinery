// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/machinery/api/schema"
)

type FactsInt interface {
	AddFactsByName(ctx context.Context, in *schema.AddFactInput, out *[]schema.Facts) error
}

func AddFactsByName(ctx context.Context, in *schema.AddFactInput, f FactsInt) (*[]schema.Facts, error) {
	facts := []schema.Facts{}
	return &facts, f.AddFactsByName(ctx, in, &facts)
}
