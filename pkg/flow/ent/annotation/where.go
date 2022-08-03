// Code generated by ent, DO NOT EDIT.

package annotation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v []byte) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldData), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Annotation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Annotation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Annotation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Annotation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Annotation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Annotation(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Annotation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Annotation(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Annotation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Annotation(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Annotation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Annotation(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSize), v))
	})
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int) predicate.Annotation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Annotation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSize), v...))
	})
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int) predicate.Annotation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Annotation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSize), v...))
	})
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSize), v))
	})
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSize), v))
	})
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSize), v))
	})
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSize), v))
	})
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHash), v))
	})
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.Annotation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Annotation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHash), v...))
	})
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.Annotation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Annotation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHash), v...))
	})
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHash), v))
	})
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHash), v))
	})
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHash), v))
	})
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHash), v))
	})
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHash), v))
	})
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHash), v))
	})
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHash), v))
	})
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHash), v))
	})
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHash), v))
	})
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v []byte) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldData), v))
	})
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v []byte) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldData), v))
	})
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...[]byte) predicate.Annotation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Annotation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldData), v...))
	})
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...[]byte) predicate.Annotation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Annotation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldData), v...))
	})
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v []byte) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldData), v))
	})
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v []byte) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldData), v))
	})
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v []byte) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldData), v))
	})
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v []byte) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldData), v))
	})
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.Namespace) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
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

// HasWorkflow applies the HasEdge predicate on the "workflow" edge.
func HasWorkflow() predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowTable, WorkflowColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkflowWith applies the HasEdge predicate on the "workflow" edge with a given conditions (other predicates).
func HasWorkflowWith(preds ...predicate.Workflow) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowTable, WorkflowColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInstance applies the HasEdge predicate on the "instance" edge.
func HasInstance() predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstanceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstanceTable, InstanceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstanceWith applies the HasEdge predicate on the "instance" edge with a given conditions (other predicates).
func HasInstanceWith(preds ...predicate.Instance) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstanceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstanceTable, InstanceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInode applies the HasEdge predicate on the "inode" edge.
func HasInode() predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InodeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InodeTable, InodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInodeWith applies the HasEdge predicate on the "inode" edge with a given conditions (other predicates).
func HasInodeWith(preds ...predicate.Inode) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InodeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InodeTable, InodeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Annotation) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Annotation) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
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
func Not(p predicate.Annotation) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		p(s.Not())
	})
}
