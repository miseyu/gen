package field

import (
	"gorm.io/gorm/clause"
)

// Raw...
type Raw Field

// Eq equal to
func (field Raw) Eq(value string) Expr {
	return expr{e: clause.Eq{Column: field.RawExpr(), Value: value}}
}

// Neq not equal to
func (field Raw) Neq(value string) Expr {
	return expr{e: clause.Neq{Column: field.RawExpr(), Value: value}}
}

// Gt greater than
func (field Raw) Gt(value string) Expr {
	return expr{e: clause.Gt{Column: field.RawExpr(), Value: value}}
}

// Gte greater or equal to
func (field Raw) Gte(value string) Expr {
	return expr{e: clause.Gte{Column: field.RawExpr(), Value: value}}
}

// Lt less than
func (field Raw) Lt(value string) Expr {
	return expr{e: clause.Lt{Column: field.RawExpr(), Value: value}}
}

// Lte less or equal to
func (field Raw) Lte(value string) Expr {
	return expr{e: clause.Lte{Column: field.RawExpr(), Value: value}}
}

// Between ...
func (field Raw) Between(left string, right string) Expr {
	return field.between([]interface{}{left, right})
}

// NotBetween ...
func (field Raw) NotBetween(left string, right string) Expr {
	return Not(field.Between(left, right))
}

// In ...
func (field Raw) In(values ...string) Expr {
	return expr{e: clause.IN{Column: field.RawExpr(), Values: field.toSlice(values)}}
}

// NotIn ...
func (field Raw) NotIn(values ...string) Expr {
	return expr{e: clause.Not(field.In(values...).expression())}
}

// Like ...
func (field Raw) Like(value string) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: value}}
}

// NotLike ...
func (field Raw) NotLike(value string) Expr {
	return expr{e: clause.Not(field.Like(value).expression())}
}

// Regexp ...
func (field Raw) Regexp(value string) Expr {
	return field.regexp(value)
}

// NotRegxp ...
func (field Raw) NotRegxp(value string) Expr {
	return expr{e: clause.Not(field.Regexp(value).expression())}
}

// Value ...
func (field Raw) Value(value string) AssignExpr {
	return field.value(value)
}

// Zero ...
func (field Raw) Zero() AssignExpr {
	return field.value("")
}

func (field Raw) toSlice(values []string) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}
