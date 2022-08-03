// Code generated by ent, DO NOT EDIT.

package mirroractivity

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// Controller applies equality check predicate on the "controller" field. It's identical to ControllerEQ.
func Controller(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldController), v))
	})
}

// Deadline applies equality check predicate on the "deadline" field. It's identical to DeadlineEQ.
func Deadline(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeadline), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndAt), v))
	})
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndAt), v...))
	})
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndAt), v...))
	})
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndAt), v))
	})
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndAt), v))
	})
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndAt), v))
	})
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndAt), v))
	})
}

// EndAtIsNil applies the IsNil predicate on the "end_at" field.
func EndAtIsNil() predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndAt)))
	})
}

// EndAtNotNil applies the NotNil predicate on the "end_at" field.
func EndAtNotNil() predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndAt)))
	})
}

// ControllerEQ applies the EQ predicate on the "controller" field.
func ControllerEQ(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldController), v))
	})
}

// ControllerNEQ applies the NEQ predicate on the "controller" field.
func ControllerNEQ(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldController), v))
	})
}

// ControllerIn applies the In predicate on the "controller" field.
func ControllerIn(vs ...string) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldController), v...))
	})
}

// ControllerNotIn applies the NotIn predicate on the "controller" field.
func ControllerNotIn(vs ...string) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldController), v...))
	})
}

// ControllerGT applies the GT predicate on the "controller" field.
func ControllerGT(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldController), v))
	})
}

// ControllerGTE applies the GTE predicate on the "controller" field.
func ControllerGTE(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldController), v))
	})
}

// ControllerLT applies the LT predicate on the "controller" field.
func ControllerLT(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldController), v))
	})
}

// ControllerLTE applies the LTE predicate on the "controller" field.
func ControllerLTE(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldController), v))
	})
}

// ControllerContains applies the Contains predicate on the "controller" field.
func ControllerContains(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldController), v))
	})
}

// ControllerHasPrefix applies the HasPrefix predicate on the "controller" field.
func ControllerHasPrefix(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldController), v))
	})
}

// ControllerHasSuffix applies the HasSuffix predicate on the "controller" field.
func ControllerHasSuffix(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldController), v))
	})
}

// ControllerIsNil applies the IsNil predicate on the "controller" field.
func ControllerIsNil() predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldController)))
	})
}

// ControllerNotNil applies the NotNil predicate on the "controller" field.
func ControllerNotNil() predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldController)))
	})
}

// ControllerEqualFold applies the EqualFold predicate on the "controller" field.
func ControllerEqualFold(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldController), v))
	})
}

// ControllerContainsFold applies the ContainsFold predicate on the "controller" field.
func ControllerContainsFold(v string) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldController), v))
	})
}

// DeadlineEQ applies the EQ predicate on the "deadline" field.
func DeadlineEQ(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeadline), v))
	})
}

// DeadlineNEQ applies the NEQ predicate on the "deadline" field.
func DeadlineNEQ(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeadline), v))
	})
}

// DeadlineIn applies the In predicate on the "deadline" field.
func DeadlineIn(vs ...time.Time) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeadline), v...))
	})
}

// DeadlineNotIn applies the NotIn predicate on the "deadline" field.
func DeadlineNotIn(vs ...time.Time) predicate.MirrorActivity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MirrorActivity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeadline), v...))
	})
}

// DeadlineGT applies the GT predicate on the "deadline" field.
func DeadlineGT(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeadline), v))
	})
}

// DeadlineGTE applies the GTE predicate on the "deadline" field.
func DeadlineGTE(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeadline), v))
	})
}

// DeadlineLT applies the LT predicate on the "deadline" field.
func DeadlineLT(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeadline), v))
	})
}

// DeadlineLTE applies the LTE predicate on the "deadline" field.
func DeadlineLTE(v time.Time) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeadline), v))
	})
}

// DeadlineIsNil applies the IsNil predicate on the "deadline" field.
func DeadlineIsNil() predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeadline)))
	})
}

// DeadlineNotNil applies the NotNil predicate on the "deadline" field.
func DeadlineNotNil() predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeadline)))
	})
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.Namespace) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMirror applies the HasEdge predicate on the "mirror" edge.
func HasMirror() predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MirrorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MirrorTable, MirrorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMirrorWith applies the HasEdge predicate on the "mirror" edge with a given conditions (other predicates).
func HasMirrorWith(preds ...predicate.Mirror) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MirrorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MirrorTable, MirrorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLogs applies the HasEdge predicate on the "logs" edge.
func HasLogs() predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LogsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LogsTable, LogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLogsWith applies the HasEdge predicate on the "logs" edge with a given conditions (other predicates).
func HasLogsWith(preds ...predicate.LogMsg) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LogsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LogsTable, LogsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MirrorActivity) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MirrorActivity) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MirrorActivity) predicate.MirrorActivity {
	return predicate.MirrorActivity(func(s *sql.Selector) {
		p(s.Not())
	})
}
