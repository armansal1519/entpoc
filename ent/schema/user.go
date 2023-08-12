package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("table_name0"),
		field.String("table_name1"),
		field.Int("table_name2"),
		field.String("table_name3"),
		field.String("table_name4"),
		field.String("table_name5"),
		field.String("table_name6"),
		field.String("table_name7"),
		field.String("table_name8"),
		field.String("table_name9"),
		field.String("table_name10"),
		field.String("table_name11"),
		field.String("table_name12"),
		field.String("table_name13"),
		field.String("table_name14"),
		field.String("table_name15"),
		field.String("table_name16"),
		field.String("table_name17"),
		field.String("table_name18"),
		field.String("table_name19"),
		field.String("table_name20"),
		field.String("table_name21"),
		field.String("table_name22"),
		field.String("table_name23"),
		field.String("table_name24"),
		field.String("table_name25"),
		field.String("table_name26"),
		field.String("table_name27"),
		field.String("table_name28"),
		field.String("table_name29"),
		field.String("table_name30"),
		field.String("table_name31"),
		field.String("table_name32"),
		field.String("table_name33"),
		field.String("table_name34"),
		field.String("table_name35"),
		field.String("table_name36"),
		field.String("table_name37"),
		field.String("table_name38"),
		field.String("table_name39"),
		field.String("table_name40"),
		field.String("table_name41"),
		field.String("table_name42"),
		field.String("altered_table_name43"),
		field.String("table_name44"),
		field.String("table_name45"),
		field.String("table_name46"),
		field.String("table_name47"),
		field.String("table_name48"),
		field.String("table_name49"),
		field.String("altered_table_name50"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
