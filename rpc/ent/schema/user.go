package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	mixins2 "resourceManager/rpc/ent/schema/mixins"
)

// User holds the schema definition for the Student entity.
type User struct {
	ent.Schema
}

// Fields of the Student.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			Comment("用户名").
			Annotations(entsql.WithComments(true)),
		field.String("phone").
			Unique().
			Comment("电话号码").
			Annotations(entsql.WithComments(true)),
		field.String("email").
			Unique().
			Comment("邮箱").
			Annotations(entsql.WithComments(true)),
		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("头像路径").
			Annotations(entsql.WithComments(true)),
	}
}

// Mixins of the Student
func (User) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixins2.UUIDMixin{},
		mixins2.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

// Edges of the Student.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the Student
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username").Unique(),
	}
}

// Annotations of the Student
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_student"},
	}
}
