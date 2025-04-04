// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationshipinfo"
)

// RelationshipInfo is the model entity for the RelationshipInfo schema.
type RelationshipInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text         string `json:"text,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RelationshipInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case relationshipinfo.FieldID:
			values[i] = new(sql.NullInt64)
		case relationshipinfo.FieldText:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RelationshipInfo fields.
func (_m *RelationshipInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case relationshipinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case relationshipinfo.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				_m.Text = value.String
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RelationshipInfo.
// This includes values selected through modifiers, order, etc.
func (_m *RelationshipInfo) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// Update returns a builder for updating this RelationshipInfo.
// Note that you need to call RelationshipInfo.Unwrap() before calling this method if this RelationshipInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *RelationshipInfo) Update() *RelationshipInfoUpdateOne {
	return NewRelationshipInfoClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the RelationshipInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *RelationshipInfo) Unwrap() *RelationshipInfo {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: RelationshipInfo is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *RelationshipInfo) String() string {
	var builder strings.Builder
	builder.WriteString("RelationshipInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("text=")
	builder.WriteString(_m.Text)
	builder.WriteByte(')')
	return builder.String()
}

// RelationshipInfos is a parsable slice of RelationshipInfo.
type RelationshipInfos []*RelationshipInfo
