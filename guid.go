// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

// NewGUID generates a new GUID.
func NewGUID() GUID {
	return GUID{UUID: uuid.Must(uuid.NewV4())}
}

// NewGUID generates a new nil GUID.
func NewGUIDNil() GUID {
	return GUID{UUID: uuid.Nil}
}

// GUID allows for unmarshalling the urls returned by Exact.
type GUID struct {
	uuid.UUID
}

// UnmarshalJSON unmarshals the guid to uuid.UUID returned from the
// Exact Online API.
func (g *GUID) UnmarshalJSON(b []byte) error {
	if g == nil {
		return nil
	}

	str := string(b)
	if str == "" || str == "null" { //added because ExactOnline contains GUIDs with value "null"
		return nil
	}

	s := []byte(strings.Replace(string(b), `"`, "", -1))
	err := (&g.UUID).UnmarshalText(s)
	if err != nil {
		return fmt.Errorf("GUID.UnmarshalJSON() error: %v", err)
	}
	return nil
}

// UnmarshalXML unmarshals the guid to uuid.UUID returned from the
// Exact Online API.
func (g *GUID) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if g == nil {
		return nil
	}

	var value string
	// Read tag content into value
	d.DecodeElement(&value, &start)

	if value == "" || value == "null" { //added because ExactOnline contains GUIDs with value "null"
		return nil
	}

	s := []byte(strings.Replace(value, `"`, "", -1))
	err := (&g.UUID).UnmarshalText(s)
	if err != nil {
		return fmt.Errorf("GUID.UnmarshalJSON() error: %v", err)
	}
	return nil
}

// MarshalJSON marshals the url to a format expected by the
// Exact Online API.
func (g *GUID) MarshalJSON() ([]byte, error) {
	if g == nil {
		return json.Marshal(nil)
	}
	if !g.IsSet() {
		return json.Marshal(nil)
	}

	return json.Marshal(g.String())
}

func (g *GUID) String() string {
	if g == nil {
		return ""
	}
	if !g.IsSet() {
		return ""
	}
	return g.UUID.String()
}

// IsSet checks if the GUID/uuid actually exists
func (g *GUID) IsSet() bool {
	if g == nil {
		return false
	}
	return g.UUID != uuid.Nil
}

func (g *GUID) FromString(guid string) error {
	id, err := uuid.FromString(guid)
	if err != nil {
		return err
	}
	g.UUID = id

	return nil
}

func (g *GUID) Equal(g1 *GUID) bool {
	if g == nil && g1 == nil {
		return true
	}
	if g == nil || g1 == nil {
		return false
	}
	return g.String() == g1.String()
}
