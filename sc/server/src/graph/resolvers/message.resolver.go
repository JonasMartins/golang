package resolvers

import (
	"context"
	"src/graph/generated"

	"src/infra/orm/gorm/models"
)

type messageResolver struct{ *Resolver }

var _ generated.MessageResolver = (*messageResolver)(nil)

func (m *messageResolver) Seen(ctx context.Context, obj *models.Message) ([]string, error) {
	return obj.Seen, nil
}
