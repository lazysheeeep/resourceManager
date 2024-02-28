package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"time"
)

type UUIDMixin struct {
	mixin.Schema
}

func (UUIDMixin) Field() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.NewUUID).Comment("UUID"),
		field.Time("created_at").
			Default(time.Now()).
			Comment("创建时间").
			Annotations(entsql.WithComments(true)),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间").
			Annotations(entsql.WithComments(true)),
	}
}
