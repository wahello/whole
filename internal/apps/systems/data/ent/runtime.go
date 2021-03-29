// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/realotz/whole/internal/apps/systems/data/ent/file"
	"github.com/realotz/whole/internal/apps/systems/data/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	fileMixin := schema.File{}.Mixin()
	fileMixinFields0 := fileMixin[0].Fields()
	_ = fileMixinFields0
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescCreateTime is the schema descriptor for create_time field.
	fileDescCreateTime := fileMixinFields0[0].Descriptor()
	// file.DefaultCreateTime holds the default value on creation for the create_time field.
	file.DefaultCreateTime = fileDescCreateTime.Default.(func() time.Time)
	// fileDescUpdateTime is the schema descriptor for update_time field.
	fileDescUpdateTime := fileMixinFields0[1].Descriptor()
	// file.DefaultUpdateTime holds the default value on creation for the update_time field.
	file.DefaultUpdateTime = fileDescUpdateTime.Default.(func() time.Time)
	// file.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	file.UpdateDefaultUpdateTime = fileDescUpdateTime.UpdateDefault.(func() time.Time)
	// fileDescName is the schema descriptor for name field.
	fileDescName := fileFields[1].Descriptor()
	// file.NameValidator is a validator for the "name" field. It is called by the builders before save.
	file.NameValidator = fileDescName.Validators[0].(func(string) error)
	// fileDescType is the schema descriptor for type field.
	fileDescType := fileFields[2].Descriptor()
	// file.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	file.TypeValidator = fileDescType.Validators[0].(func(string) error)
	// fileDescMd5 is the schema descriptor for md5 field.
	fileDescMd5 := fileFields[3].Descriptor()
	// file.Md5Validator is a validator for the "md5" field. It is called by the builders before save.
	file.Md5Validator = fileDescMd5.Validators[0].(func(string) error)
	// fileDescPath is the schema descriptor for path field.
	fileDescPath := fileFields[4].Descriptor()
	// file.PathValidator is a validator for the "path" field. It is called by the builders before save.
	file.PathValidator = fileDescPath.Validators[0].(func(string) error)
	// fileDescID is the schema descriptor for id field.
	fileDescID := fileFields[0].Descriptor()
	// file.IDValidator is a validator for the "id" field. It is called by the builders before save.
	file.IDValidator = fileDescID.Validators[0].(func(int64) error)
}