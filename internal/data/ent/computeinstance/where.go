// Code generated by ent, DO NOT EDIT.

package computeinstance

import (
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLTE(FieldID, id))
}

// Owner applies equality check predicate on the "owner" field. It's identical to OwnerEQ.
func Owner(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldOwner, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldName, v))
}

// Core applies equality check predicate on the "core" field. It's identical to CoreEQ.
func Core(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldCore, v))
}

// Memory applies equality check predicate on the "memory" field. It's identical to MemoryEQ.
func Memory(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldMemory, v))
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldImage, v))
}

// ExpirationTime applies equality check predicate on the "expiration_time" field. It's identical to ExpirationTimeEQ.
func ExpirationTime(v time.Time) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldExpirationTime, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldStatus, v))
}

// ContainerID applies equality check predicate on the "container_id" field. It's identical to ContainerIDEQ.
func ContainerID(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldContainerID, v))
}

// PeerID applies equality check predicate on the "peer_id" field. It's identical to PeerIDEQ.
func PeerID(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldPeerID, v))
}

// OwnerEQ applies the EQ predicate on the "owner" field.
func OwnerEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldOwner, v))
}

// OwnerNEQ applies the NEQ predicate on the "owner" field.
func OwnerNEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNEQ(FieldOwner, v))
}

// OwnerIn applies the In predicate on the "owner" field.
func OwnerIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldIn(FieldOwner, vs...))
}

// OwnerNotIn applies the NotIn predicate on the "owner" field.
func OwnerNotIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNotIn(FieldOwner, vs...))
}

// OwnerGT applies the GT predicate on the "owner" field.
func OwnerGT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGT(FieldOwner, v))
}

// OwnerGTE applies the GTE predicate on the "owner" field.
func OwnerGTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGTE(FieldOwner, v))
}

// OwnerLT applies the LT predicate on the "owner" field.
func OwnerLT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLT(FieldOwner, v))
}

// OwnerLTE applies the LTE predicate on the "owner" field.
func OwnerLTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLTE(FieldOwner, v))
}

// OwnerContains applies the Contains predicate on the "owner" field.
func OwnerContains(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContains(FieldOwner, v))
}

// OwnerHasPrefix applies the HasPrefix predicate on the "owner" field.
func OwnerHasPrefix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasPrefix(FieldOwner, v))
}

// OwnerHasSuffix applies the HasSuffix predicate on the "owner" field.
func OwnerHasSuffix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasSuffix(FieldOwner, v))
}

// OwnerEqualFold applies the EqualFold predicate on the "owner" field.
func OwnerEqualFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEqualFold(FieldOwner, v))
}

// OwnerContainsFold applies the ContainsFold predicate on the "owner" field.
func OwnerContainsFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContainsFold(FieldOwner, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContainsFold(FieldName, v))
}

// CoreEQ applies the EQ predicate on the "core" field.
func CoreEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldCore, v))
}

// CoreNEQ applies the NEQ predicate on the "core" field.
func CoreNEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNEQ(FieldCore, v))
}

// CoreIn applies the In predicate on the "core" field.
func CoreIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldIn(FieldCore, vs...))
}

// CoreNotIn applies the NotIn predicate on the "core" field.
func CoreNotIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNotIn(FieldCore, vs...))
}

// CoreGT applies the GT predicate on the "core" field.
func CoreGT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGT(FieldCore, v))
}

// CoreGTE applies the GTE predicate on the "core" field.
func CoreGTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGTE(FieldCore, v))
}

// CoreLT applies the LT predicate on the "core" field.
func CoreLT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLT(FieldCore, v))
}

// CoreLTE applies the LTE predicate on the "core" field.
func CoreLTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLTE(FieldCore, v))
}

// CoreContains applies the Contains predicate on the "core" field.
func CoreContains(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContains(FieldCore, v))
}

// CoreHasPrefix applies the HasPrefix predicate on the "core" field.
func CoreHasPrefix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasPrefix(FieldCore, v))
}

// CoreHasSuffix applies the HasSuffix predicate on the "core" field.
func CoreHasSuffix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasSuffix(FieldCore, v))
}

// CoreEqualFold applies the EqualFold predicate on the "core" field.
func CoreEqualFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEqualFold(FieldCore, v))
}

// CoreContainsFold applies the ContainsFold predicate on the "core" field.
func CoreContainsFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContainsFold(FieldCore, v))
}

// MemoryEQ applies the EQ predicate on the "memory" field.
func MemoryEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldMemory, v))
}

// MemoryNEQ applies the NEQ predicate on the "memory" field.
func MemoryNEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNEQ(FieldMemory, v))
}

// MemoryIn applies the In predicate on the "memory" field.
func MemoryIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldIn(FieldMemory, vs...))
}

// MemoryNotIn applies the NotIn predicate on the "memory" field.
func MemoryNotIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNotIn(FieldMemory, vs...))
}

// MemoryGT applies the GT predicate on the "memory" field.
func MemoryGT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGT(FieldMemory, v))
}

// MemoryGTE applies the GTE predicate on the "memory" field.
func MemoryGTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGTE(FieldMemory, v))
}

// MemoryLT applies the LT predicate on the "memory" field.
func MemoryLT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLT(FieldMemory, v))
}

// MemoryLTE applies the LTE predicate on the "memory" field.
func MemoryLTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLTE(FieldMemory, v))
}

// MemoryContains applies the Contains predicate on the "memory" field.
func MemoryContains(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContains(FieldMemory, v))
}

// MemoryHasPrefix applies the HasPrefix predicate on the "memory" field.
func MemoryHasPrefix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasPrefix(FieldMemory, v))
}

// MemoryHasSuffix applies the HasSuffix predicate on the "memory" field.
func MemoryHasSuffix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasSuffix(FieldMemory, v))
}

// MemoryEqualFold applies the EqualFold predicate on the "memory" field.
func MemoryEqualFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEqualFold(FieldMemory, v))
}

// MemoryContainsFold applies the ContainsFold predicate on the "memory" field.
func MemoryContainsFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContainsFold(FieldMemory, v))
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldImage, v))
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNEQ(FieldImage, v))
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldIn(FieldImage, vs...))
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNotIn(FieldImage, vs...))
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGT(FieldImage, v))
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGTE(FieldImage, v))
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLT(FieldImage, v))
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLTE(FieldImage, v))
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContains(FieldImage, v))
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasPrefix(FieldImage, v))
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasSuffix(FieldImage, v))
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEqualFold(FieldImage, v))
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContainsFold(FieldImage, v))
}

// ExpirationTimeEQ applies the EQ predicate on the "expiration_time" field.
func ExpirationTimeEQ(v time.Time) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldExpirationTime, v))
}

// ExpirationTimeNEQ applies the NEQ predicate on the "expiration_time" field.
func ExpirationTimeNEQ(v time.Time) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNEQ(FieldExpirationTime, v))
}

// ExpirationTimeIn applies the In predicate on the "expiration_time" field.
func ExpirationTimeIn(vs ...time.Time) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldIn(FieldExpirationTime, vs...))
}

// ExpirationTimeNotIn applies the NotIn predicate on the "expiration_time" field.
func ExpirationTimeNotIn(vs ...time.Time) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNotIn(FieldExpirationTime, vs...))
}

// ExpirationTimeGT applies the GT predicate on the "expiration_time" field.
func ExpirationTimeGT(v time.Time) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGT(FieldExpirationTime, v))
}

// ExpirationTimeGTE applies the GTE predicate on the "expiration_time" field.
func ExpirationTimeGTE(v time.Time) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGTE(FieldExpirationTime, v))
}

// ExpirationTimeLT applies the LT predicate on the "expiration_time" field.
func ExpirationTimeLT(v time.Time) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLT(FieldExpirationTime, v))
}

// ExpirationTimeLTE applies the LTE predicate on the "expiration_time" field.
func ExpirationTimeLTE(v time.Time) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLTE(FieldExpirationTime, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLTE(FieldStatus, v))
}

// ContainerIDEQ applies the EQ predicate on the "container_id" field.
func ContainerIDEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldContainerID, v))
}

// ContainerIDNEQ applies the NEQ predicate on the "container_id" field.
func ContainerIDNEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNEQ(FieldContainerID, v))
}

// ContainerIDIn applies the In predicate on the "container_id" field.
func ContainerIDIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldIn(FieldContainerID, vs...))
}

// ContainerIDNotIn applies the NotIn predicate on the "container_id" field.
func ContainerIDNotIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNotIn(FieldContainerID, vs...))
}

// ContainerIDGT applies the GT predicate on the "container_id" field.
func ContainerIDGT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGT(FieldContainerID, v))
}

// ContainerIDGTE applies the GTE predicate on the "container_id" field.
func ContainerIDGTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGTE(FieldContainerID, v))
}

// ContainerIDLT applies the LT predicate on the "container_id" field.
func ContainerIDLT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLT(FieldContainerID, v))
}

// ContainerIDLTE applies the LTE predicate on the "container_id" field.
func ContainerIDLTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLTE(FieldContainerID, v))
}

// ContainerIDContains applies the Contains predicate on the "container_id" field.
func ContainerIDContains(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContains(FieldContainerID, v))
}

// ContainerIDHasPrefix applies the HasPrefix predicate on the "container_id" field.
func ContainerIDHasPrefix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasPrefix(FieldContainerID, v))
}

// ContainerIDHasSuffix applies the HasSuffix predicate on the "container_id" field.
func ContainerIDHasSuffix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasSuffix(FieldContainerID, v))
}

// ContainerIDIsNil applies the IsNil predicate on the "container_id" field.
func ContainerIDIsNil() predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldIsNull(FieldContainerID))
}

// ContainerIDNotNil applies the NotNil predicate on the "container_id" field.
func ContainerIDNotNil() predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNotNull(FieldContainerID))
}

// ContainerIDEqualFold applies the EqualFold predicate on the "container_id" field.
func ContainerIDEqualFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEqualFold(FieldContainerID, v))
}

// ContainerIDContainsFold applies the ContainsFold predicate on the "container_id" field.
func ContainerIDContainsFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContainsFold(FieldContainerID, v))
}

// PeerIDEQ applies the EQ predicate on the "peer_id" field.
func PeerIDEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEQ(FieldPeerID, v))
}

// PeerIDNEQ applies the NEQ predicate on the "peer_id" field.
func PeerIDNEQ(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNEQ(FieldPeerID, v))
}

// PeerIDIn applies the In predicate on the "peer_id" field.
func PeerIDIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldIn(FieldPeerID, vs...))
}

// PeerIDNotIn applies the NotIn predicate on the "peer_id" field.
func PeerIDNotIn(vs ...string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNotIn(FieldPeerID, vs...))
}

// PeerIDGT applies the GT predicate on the "peer_id" field.
func PeerIDGT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGT(FieldPeerID, v))
}

// PeerIDGTE applies the GTE predicate on the "peer_id" field.
func PeerIDGTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldGTE(FieldPeerID, v))
}

// PeerIDLT applies the LT predicate on the "peer_id" field.
func PeerIDLT(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLT(FieldPeerID, v))
}

// PeerIDLTE applies the LTE predicate on the "peer_id" field.
func PeerIDLTE(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldLTE(FieldPeerID, v))
}

// PeerIDContains applies the Contains predicate on the "peer_id" field.
func PeerIDContains(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContains(FieldPeerID, v))
}

// PeerIDHasPrefix applies the HasPrefix predicate on the "peer_id" field.
func PeerIDHasPrefix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasPrefix(FieldPeerID, v))
}

// PeerIDHasSuffix applies the HasSuffix predicate on the "peer_id" field.
func PeerIDHasSuffix(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldHasSuffix(FieldPeerID, v))
}

// PeerIDIsNil applies the IsNil predicate on the "peer_id" field.
func PeerIDIsNil() predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldIsNull(FieldPeerID))
}

// PeerIDNotNil applies the NotNil predicate on the "peer_id" field.
func PeerIDNotNil() predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldNotNull(FieldPeerID))
}

// PeerIDEqualFold applies the EqualFold predicate on the "peer_id" field.
func PeerIDEqualFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldEqualFold(FieldPeerID, v))
}

// PeerIDContainsFold applies the ContainsFold predicate on the "peer_id" field.
func PeerIDContainsFold(v string) predicate.ComputeInstance {
	return predicate.ComputeInstance(sql.FieldContainsFold(FieldPeerID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ComputeInstance) predicate.ComputeInstance {
	return predicate.ComputeInstance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ComputeInstance) predicate.ComputeInstance {
	return predicate.ComputeInstance(func(s *sql.Selector) {
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
func Not(p predicate.ComputeInstance) predicate.ComputeInstance {
	return predicate.ComputeInstance(func(s *sql.Selector) {
		p(s.Not())
	})
}
