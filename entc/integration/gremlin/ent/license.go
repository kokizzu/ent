// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/gremlin"
)

// License is the model entity for the License schema.
type License struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
}

// FromResponse scans the gremlin response data into License.
func (_m *License) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_m struct {
		ID         int   `json:"id,omitempty"`
		CreateTime int64 `json:"create_time,omitempty"`
		UpdateTime int64 `json:"update_time,omitempty"`
	}
	if err := vmap.Decode(&scan_m); err != nil {
		return err
	}
	_m.ID = scan_m.ID
	_m.CreateTime = time.Unix(0, scan_m.CreateTime)
	_m.UpdateTime = time.Unix(0, scan_m.UpdateTime)
	return nil
}

// Update returns a builder for updating this License.
// Note that you need to call License.Unwrap() before calling this method if this License
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *License) Update() *LicenseUpdateOne {
	return NewLicenseClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the License entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *License) Unwrap() *License {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: License is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *License) String() string {
	var builder strings.Builder
	builder.WriteString("License(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("create_time=")
	builder.WriteString(_m.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(_m.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Licenses is a parsable slice of License.
type Licenses []*License

// FromResponse scans the gremlin response data into Licenses.
func (_m *Licenses) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_m []struct {
		ID         int   `json:"id,omitempty"`
		CreateTime int64 `json:"create_time,omitempty"`
		UpdateTime int64 `json:"update_time,omitempty"`
	}
	if err := vmap.Decode(&scan_m); err != nil {
		return err
	}
	for _, v := range scan_m {
		node := &License{ID: v.ID}
		node.CreateTime = time.Unix(0, v.CreateTime)
		node.UpdateTime = time.Unix(0, v.UpdateTime)
		*_m = append(*_m, node)
	}
	return nil
}
