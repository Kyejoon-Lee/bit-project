// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bit-project/gateway/ent/schema"
	"bit-project/gateway/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUID is the schema descriptor for uid field.
	userDescUID := userFields[0].Descriptor()
	// user.DefaultUID holds the default value on creation for the uid field.
	user.DefaultUID = userDescUID.Default.(func() uuid.UUID)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[2].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescLastLoginDate is the schema descriptor for last_login_date field.
	userDescLastLoginDate := userFields[5].Descriptor()
	// user.DefaultLastLoginDate holds the default value on creation for the last_login_date field.
	user.DefaultLastLoginDate = userDescLastLoginDate.Default.(func() time.Time)
}
