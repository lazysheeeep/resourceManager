package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type StatusMixin struct {
	mixin.Schema
}

func (StatusMixin) Field() []ent.Field {
	return []ent.Field{
		field.String("status").
			Default("1").
			Optional().
			Comment("1 正常|0 禁用").
			Annotations(entsql.WithComments(true)),
	}
}
