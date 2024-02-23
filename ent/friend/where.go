// Code generated by ent, DO NOT EDIT.

package friend

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/datti-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldID, id))
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldUID, v))
}

// FriendUID applies equality check predicate on the "friend_uid" field. It's identical to FriendUIDEQ.
func FriendUID(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldFriendUID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldDeletedAt, v))
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldUID, v))
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldUID, v))
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldUID, vs...))
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldUID, vs...))
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldUID, v))
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldUID, v))
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldUID, v))
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldUID, v))
}

// UIDContains applies the Contains predicate on the "uid" field.
func UIDContains(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContains(FieldUID, v))
}

// UIDHasPrefix applies the HasPrefix predicate on the "uid" field.
func UIDHasPrefix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasPrefix(FieldUID, v))
}

// UIDHasSuffix applies the HasSuffix predicate on the "uid" field.
func UIDHasSuffix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasSuffix(FieldUID, v))
}

// UIDEqualFold applies the EqualFold predicate on the "uid" field.
func UIDEqualFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEqualFold(FieldUID, v))
}

// UIDContainsFold applies the ContainsFold predicate on the "uid" field.
func UIDContainsFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContainsFold(FieldUID, v))
}

// FriendUIDEQ applies the EQ predicate on the "friend_uid" field.
func FriendUIDEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldFriendUID, v))
}

// FriendUIDNEQ applies the NEQ predicate on the "friend_uid" field.
func FriendUIDNEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldFriendUID, v))
}

// FriendUIDIn applies the In predicate on the "friend_uid" field.
func FriendUIDIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldFriendUID, vs...))
}

// FriendUIDNotIn applies the NotIn predicate on the "friend_uid" field.
func FriendUIDNotIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldFriendUID, vs...))
}

// FriendUIDGT applies the GT predicate on the "friend_uid" field.
func FriendUIDGT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldFriendUID, v))
}

// FriendUIDGTE applies the GTE predicate on the "friend_uid" field.
func FriendUIDGTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldFriendUID, v))
}

// FriendUIDLT applies the LT predicate on the "friend_uid" field.
func FriendUIDLT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldFriendUID, v))
}

// FriendUIDLTE applies the LTE predicate on the "friend_uid" field.
func FriendUIDLTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldFriendUID, v))
}

// FriendUIDContains applies the Contains predicate on the "friend_uid" field.
func FriendUIDContains(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContains(FieldFriendUID, v))
}

// FriendUIDHasPrefix applies the HasPrefix predicate on the "friend_uid" field.
func FriendUIDHasPrefix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasPrefix(FieldFriendUID, v))
}

// FriendUIDHasSuffix applies the HasSuffix predicate on the "friend_uid" field.
func FriendUIDHasSuffix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasSuffix(FieldFriendUID, v))
}

// FriendUIDEqualFold applies the EqualFold predicate on the "friend_uid" field.
func FriendUIDEqualFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEqualFold(FieldFriendUID, v))
}

// FriendUIDContainsFold applies the ContainsFold predicate on the "friend_uid" field.
func FriendUIDContainsFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContainsFold(FieldFriendUID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Friend {
	return predicate.Friend(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Friend {
	return predicate.Friend(sql.FieldNotNull(FieldDeletedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Friend) predicate.Friend {
	return predicate.Friend(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Friend) predicate.Friend {
	return predicate.Friend(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Friend) predicate.Friend {
	return predicate.Friend(sql.NotPredicates(p))
}