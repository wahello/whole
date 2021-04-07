// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/realotz/whole/internal/services/cms/data/ent/category"
	"github.com/realotz/whole/internal/services/cms/data/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categoryMixin := schema.Category{}.Mixin()
	categoryMixinFields0 := categoryMixin[0].Fields()
	_ = categoryMixinFields0
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescCreateTime is the schema descriptor for create_time field.
	categoryDescCreateTime := categoryMixinFields0[0].Descriptor()
	// category.DefaultCreateTime holds the default value on creation for the create_time field.
	category.DefaultCreateTime = categoryDescCreateTime.Default.(func() time.Time)
	// categoryDescUpdateTime is the schema descriptor for update_time field.
	categoryDescUpdateTime := categoryMixinFields0[1].Descriptor()
	// category.DefaultUpdateTime holds the default value on creation for the update_time field.
	category.DefaultUpdateTime = categoryDescUpdateTime.Default.(func() time.Time)
	// category.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	category.UpdateDefaultUpdateTime = categoryDescUpdateTime.UpdateDefault.(func() time.Time)
}