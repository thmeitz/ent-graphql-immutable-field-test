package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TestList holds the schema definition for the TestList entity.
type TestList struct {
	ent.Schema
}

// Fields of the TestList.
func (TestList) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(
				entsql.Annotation{Size: 50},
				entgql.OrderField("NAME"),
			),
		field.Time("valid_from").
			Immutable().
			Annotations(
				entgql.OrderField("VALID_FROM"),
				entgql.Type("Date"),
			).
			SchemaType(map[string]string{
				dialect.MySQL: "date",
			}),
		field.Time("valid_to").
			// Optional(). // this is needed for later usage
			// Nillable(). // this is needed for later usage
			Annotations(
				entgql.OrderField("VALID_TO"),
				entgql.Type("Date"),
			).
			SchemaType(map[string]string{
				dialect.MySQL: "date",
			}),
	}
}

// Edges of the TestList.
func (TestList) Edges() []ent.Edge {
	return nil
}

// Mixin of the TestList
func (TestList) Mixin() []ent.Mixin {
	return []ent.Mixin{
		XidGqlMixin{},
		TimeMixin{},
	}
}

func (TestList) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "test_lists"},
	}
}

func (TestList) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("valid_from").
			StorageKey("ix_valid_from"),
		index.Fields("name", "valid_from").
			Unique().
			StorageKey("ux_name_valid_from"),
	}
}
