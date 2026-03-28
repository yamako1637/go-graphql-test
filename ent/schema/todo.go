package schema

import (
	"entgo.io/ent"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
		field.String("email").
			NotEmpty().
			MaxLen(255),
		field.Bool("done").
			Default(false),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}

func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
