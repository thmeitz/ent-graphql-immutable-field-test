package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"entgo.io/bug/ent"
	"entgo.io/bug/graph/generated"
	"github.com/rs/xid"
)

func (r *queryResolver) Node(ctx context.Context, id xid.ID) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []xid.ID) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
