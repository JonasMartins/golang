package resolvers

import (
	"context"
	"src/graph/generated"
	"src/infra/orm/gorm/models"
)

type subscriptionResolver struct{ *Resolver }

var _ generated.SubscriptionResolver = (*subscriptionResolver)(nil)

func (s *subscriptionResolver) MessageSended(ctx context.Context, chatId string) (<-chan *models.Message, error) {
	panic("not implemented yet")
}
