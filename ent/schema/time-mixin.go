package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TimeMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type TimeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

// Fields of TimeMixin
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime(6)",
			}),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime(6)",
			}),
	}
}
