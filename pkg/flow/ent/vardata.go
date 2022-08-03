// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/direktiv/direktiv/pkg/flow/ent/vardata"
	"github.com/google/uuid"
)

// VarData is the model entity for the VarData schema.
type VarData struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"hash,omitempty"`
	// Data holds the value of the "data" field.
	Data []byte `json:"data,omitempty"`
	// MimeType holds the value of the "mime_type" field.
	MimeType string `json:"mime_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VarDataQuery when eager-loading is set.
	Edges VarDataEdges `json:"edges"`
}

// VarDataEdges holds the relations/edges for other nodes in the graph.
type VarDataEdges struct {
	// Varrefs holds the value of the varrefs edge.
	Varrefs []*VarRef `json:"varrefs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VarrefsOrErr returns the Varrefs value or an error if the edge
// was not loaded in eager-loading.
func (e VarDataEdges) VarrefsOrErr() ([]*VarRef, error) {
	if e.loadedTypes[0] {
		return e.Varrefs, nil
	}
	return nil, &NotLoadedError{edge: "varrefs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VarData) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case vardata.FieldData:
			values[i] = new([]byte)
		case vardata.FieldSize:
			values[i] = new(sql.NullInt64)
		case vardata.FieldHash, vardata.FieldMimeType:
			values[i] = new(sql.NullString)
		case vardata.FieldCreatedAt, vardata.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case vardata.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type VarData", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VarData fields.
func (vd *VarData) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vardata.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				vd.ID = *value
			}
		case vardata.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				vd.CreatedAt = value.Time
			}
		case vardata.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				vd.UpdatedAt = value.Time
			}
		case vardata.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				vd.Size = int(value.Int64)
			}
		case vardata.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				vd.Hash = value.String
			}
		case vardata.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil {
				vd.Data = *value
			}
		case vardata.FieldMimeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mime_type", values[i])
			} else if value.Valid {
				vd.MimeType = value.String
			}
		}
	}
	return nil
}

// QueryVarrefs queries the "varrefs" edge of the VarData entity.
func (vd *VarData) QueryVarrefs() *VarRefQuery {
	return (&VarDataClient{config: vd.config}).QueryVarrefs(vd)
}

// Update returns a builder for updating this VarData.
// Note that you need to call VarData.Unwrap() before calling this method if this VarData
// was returned from a transaction, and the transaction was committed or rolled back.
func (vd *VarData) Update() *VarDataUpdateOne {
	return (&VarDataClient{config: vd.config}).UpdateOne(vd)
}

// Unwrap unwraps the VarData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vd *VarData) Unwrap() *VarData {
	_tx, ok := vd.config.driver.(*txDriver)
	if !ok {
		panic("ent: VarData is not a transactional entity")
	}
	vd.config.driver = _tx.drv
	return vd
}

// String implements the fmt.Stringer.
func (vd *VarData) String() string {
	var builder strings.Builder
	builder.WriteString("VarData(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vd.ID))
	builder.WriteString("created_at=")
	builder.WriteString(vd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(vd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", vd.Size))
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(vd.Hash)
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", vd.Data))
	builder.WriteString(", ")
	builder.WriteString("mime_type=")
	builder.WriteString(vd.MimeType)
	builder.WriteByte(')')
	return builder.String()
}

// VarDataSlice is a parsable slice of VarData.
type VarDataSlice []*VarData

func (vd VarDataSlice) config(cfg config) {
	for _i := range vd {
		vd[_i].config = cfg
	}
}
