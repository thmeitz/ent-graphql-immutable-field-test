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

func (r *mutationResolver) CreateTestList(ctx context.Context, input ent.CreateTestListInput) (*ent.TestList, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTestList(ctx context.Context, id xid.ID, input ent.UpdateTestListInput) (*ent.TestList, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTestList(ctx context.Context, id xid.ID) (*ent.TestList, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TestLists(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TestListOrder, where *ent.TestListWhereInput) (*ent.TestListConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
