package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bcmmbaga/scale/graph/generated"
	"github.com/bcmmbaga/scale/models"
)

func (r *messagesWalletResolver) CreatedAt(ctx context.Context, obj *models.MessagesWallet) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *messagesWalletResolver) UpdatedAt(ctx context.Context, obj *models.MessagesWallet) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// MessagesWallet returns generated.MessagesWalletResolver implementation.
func (r *Resolver) MessagesWallet() generated.MessagesWalletResolver {
	return &messagesWalletResolver{r}
}

type messagesWalletResolver struct{ *Resolver }
