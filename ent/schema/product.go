package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("table_name0"),
		field.String("table_name1"),
		field.String("table_name2"),
		field.String("table_name3"),
		field.String("table_name4"),
		field.String("table_name5"),
		field.String("table_name6"),
		field.String("table_name7"),
		field.String("table_name8"),
		field.String("table_name9"),
		field.String("table_name10"),
		field.String("table_name11"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
