// Code generated by ent, DO NOT EDIT.

package inode

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// ExtendedType applies equality check predicate on the "extended_type" field. It's identical to ExtendedTypeEQ.
func ExtendedType(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExtendedType), v))
	})
}

// ReadOnly applies equality check predicate on the "readOnly" field. It's identical to ReadOnlyEQ.
func ReadOnly(v bool) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadOnly), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Inode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Inode(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Inode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Inode(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Inode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Inode(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Inode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Inode(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Inode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Inode(func(s *sql.Selector) {
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
func NameNotIn(vs ...string) predicate.Inode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Inode(func(s *sql.Selector) {
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
func NameGT(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Inode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Inode(func(s *sql.Selector) {
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
func TypeNotIn(vs ...string) predicate.Inode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Inode(func(s *sql.Selector) {
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
func TypeGT(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// AttributesIsNil applies the IsNil predicate on the "attributes" field.
func AttributesIsNil() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAttributes)))
	})
}

// AttributesNotNil applies the NotNil predicate on the "attributes" field.
func AttributesNotNil() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAttributes)))
	})
}

// ExtendedTypeEQ applies the EQ predicate on the "extended_type" field.
func ExtendedTypeEQ(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExtendedType), v))
	})
}

// ExtendedTypeNEQ applies the NEQ predicate on the "extended_type" field.
func ExtendedTypeNEQ(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExtendedType), v))
	})
}

// ExtendedTypeIn applies the In predicate on the "extended_type" field.
func ExtendedTypeIn(vs ...string) predicate.Inode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Inode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExtendedType), v...))
	})
}

// ExtendedTypeNotIn applies the NotIn predicate on the "extended_type" field.
func ExtendedTypeNotIn(vs ...string) predicate.Inode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Inode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExtendedType), v...))
	})
}

// ExtendedTypeGT applies the GT predicate on the "extended_type" field.
func ExtendedTypeGT(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExtendedType), v))
	})
}

// ExtendedTypeGTE applies the GTE predicate on the "extended_type" field.
func ExtendedTypeGTE(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExtendedType), v))
	})
}

// ExtendedTypeLT applies the LT predicate on the "extended_type" field.
func ExtendedTypeLT(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExtendedType), v))
	})
}

// ExtendedTypeLTE applies the LTE predicate on the "extended_type" field.
func ExtendedTypeLTE(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExtendedType), v))
	})
}

// ExtendedTypeContains applies the Contains predicate on the "extended_type" field.
func ExtendedTypeContains(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExtendedType), v))
	})
}

// ExtendedTypeHasPrefix applies the HasPrefix predicate on the "extended_type" field.
func ExtendedTypeHasPrefix(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExtendedType), v))
	})
}

// ExtendedTypeHasSuffix applies the HasSuffix predicate on the "extended_type" field.
func ExtendedTypeHasSuffix(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExtendedType), v))
	})
}

// ExtendedTypeIsNil applies the IsNil predicate on the "extended_type" field.
func ExtendedTypeIsNil() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExtendedType)))
	})
}

// ExtendedTypeNotNil applies the NotNil predicate on the "extended_type" field.
func ExtendedTypeNotNil() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExtendedType)))
	})
}

// ExtendedTypeEqualFold applies the EqualFold predicate on the "extended_type" field.
func ExtendedTypeEqualFold(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExtendedType), v))
	})
}

// ExtendedTypeContainsFold applies the ContainsFold predicate on the "extended_type" field.
func ExtendedTypeContainsFold(v string) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExtendedType), v))
	})
}

// ReadOnlyEQ applies the EQ predicate on the "readOnly" field.
func ReadOnlyEQ(v bool) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadOnly), v))
	})
}

// ReadOnlyNEQ applies the NEQ predicate on the "readOnly" field.
func ReadOnlyNEQ(v bool) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReadOnly), v))
	})
}

// ReadOnlyIsNil applies the IsNil predicate on the "readOnly" field.
func ReadOnlyIsNil() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldReadOnly)))
	})
}

// ReadOnlyNotNil applies the NotNil predicate on the "readOnly" field.
func ReadOnlyNotNil() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldReadOnly)))
	})
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.Namespace) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
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

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChildrenTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Inode) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Inode) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkflow applies the HasEdge predicate on the "workflow" edge.
func HasWorkflow() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, WorkflowTable, WorkflowColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkflowWith applies the HasEdge predicate on the "workflow" edge with a given conditions (other predicates).
func HasWorkflowWith(preds ...predicate.Workflow) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, WorkflowTable, WorkflowColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMirror applies the HasEdge predicate on the "mirror" edge.
func HasMirror() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MirrorTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, MirrorTable, MirrorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMirrorWith applies the HasEdge predicate on the "mirror" edge with a given conditions (other predicates).
func HasMirrorWith(preds ...predicate.Mirror) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MirrorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, MirrorTable, MirrorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAnnotations applies the HasEdge predicate on the "annotations" edge.
func HasAnnotations() predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnnotationsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AnnotationsTable, AnnotationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnnotationsWith applies the HasEdge predicate on the "annotations" edge with a given conditions (other predicates).
func HasAnnotationsWith(preds ...predicate.Annotation) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnnotationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AnnotationsTable, AnnotationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Inode) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Inode) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
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
func Not(p predicate.Inode) predicate.Inode {
	return predicate.Inode(func(s *sql.Selector) {
		p(s.Not())
	})
}
