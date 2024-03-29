package resolvers

//go:generate go get github.com/99designs/gqlgen@v0.17.9
//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"src/graph/generated"
	"sync"

	"gorm.io/gorm"
)

type Resolver struct {
	DB    *gorm.DB
	Chats sync.Map
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Subscription() generated.SubscriptionResolver {
	return &subscriptionResolver{r}
}
func (r *Resolver) Base() generated.BaseResolver {
	return &baseResolver{r}
}

func (r *Resolver) Message() generated.MessageResolver {
	return &messageResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
