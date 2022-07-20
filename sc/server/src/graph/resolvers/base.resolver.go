package resolvers

import (
	"context"
	"src/graph/generated"
	"src/infra/orm/gorm/models/base"
	"time"
)

type baseResolver struct{ *Resolver }

var _ generated.BaseResolver = (*baseResolver)(nil)

func (b *baseResolver) DeletedAt(ctx context.Context, obj *base.Base) (*time.Time, error) {
	return &obj.DeletedAt.Time, nil
}

func (b *baseResolver) ID(ctx context.Context, obj *base.Base) (string, error) {
	return obj.ID.String(), nil
}
