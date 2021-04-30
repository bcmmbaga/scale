package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bcmmbaga/scale/graph/generated"
	"github.com/bcmmbaga/scale/models"
)

func (r *accountResolver) CreatedAt(ctx context.Context, obj *models.Account) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accountResolver) UpdatedAt(ctx context.Context, obj *models.Account) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Account returns generated.AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

type accountResolver struct{ *Resolver }
