package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uid", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),
		field.String("kakao_sub").
			Unique(),
		field.String("name").
			MaxLen(100),
		field.String("email").
			Unique(),
		field.String("refresh_token"),
		field.Time("last_login_date").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
