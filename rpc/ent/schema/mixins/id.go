package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type IDMixin struct {
	mixin.Schema
}

func (IDMixin) Field() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Time("created_at").
			Default(time.Now).
			Comment("创建时间").
			Annotations(entsql.WithComments(true)),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间").
			Annotations(entsql.WithComments(true)),
	}
}
