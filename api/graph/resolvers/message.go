package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated2 "github.com/bcmmbaga/scale/api/graph/generated"
	model2 "github.com/bcmmbaga/scale/api/graph/model"
)

func (r *messagesWalletResolver) CreatedAt(ctx context.Context, obj *model2.MessagesWallet) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *messagesWalletResolver) UpdatedAt(ctx context.Context, obj *model2.MessagesWallet) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// MessagesWallet returns generated2.MessagesWalletResolver implementation.
func (r *Resolver) MessagesWallet() generated2.MessagesWalletResolver {
	return &messagesWalletResolver{r}
}

type messagesWalletResolver struct{ *Resolver }
