package gid

import (
	"crypto/rand"
	"database/sql/driver"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"time"
)

const (
	GIDSize = 24
)

type (
	GID [GIDSize]byte
)

var (
	Nil = GID{}
)

func ParseGID(encoded string) (GID, error) {
	gid := GID{}
	err := gid.UnmarshalText([]byte(encoded))
	if err != nil {
		return Nil, err
	}
	return gid, nil
}

func New(tenantID TenantID, entityType uint16) GID {
	id, err := NewGID(tenantID, entityType)
	if err != nil {
		panic(fmt.Sprintf("failed to generate GID: %v", err))
	}
	return id
}

func NewGID(tenantID TenantID, entityType uint16) (GID, error) {
	var id GID

	copy(id[0:8], tenantID[:])

	binary.BigEndian.PutUint16(id[8:10], entityType)

	now := time.Now().UnixMilli()
	binary.BigEndian.PutUint64(id[10:18], uint64(now))

	_, err := rand.Read(id[18:24])
	if err != nil {
		return Nil, fmt.Errorf("failed to generate random bytes: %v", err)
	}

	return id, nil
}

func (gid GID) Value() (driver.Value, error) {
	return gid.String(), nil
}

func (gid GID) TenantID() TenantID {
	var tenantID TenantID
	copy(tenantID[:], gid[0:8])
	return tenantID
}

func (gid GID) EntityType() uint16 {
	return binary.BigEndian.Uint16(gid[8:10])
}

func (gid GID) Timestamp() time.Time {
	millis := binary.BigEndian.Uint64(gid[10:18])
	return time.UnixMilli(int64(millis))
}

func (gid *GID) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		enc := base64.RawURLEncoding
		id, err := enc.DecodeString(v)
		if err != nil {
			return err
		}

		if len(id) != GIDSize {
			return fmt.Errorf("invalid length for GID: got %d, want %d", len(id), GIDSize)
		}

		copy((*gid)[:], id)
	default:
		return fmt.Errorf("invalid type for GID: expected string, got %T", value)
	}
	return nil
}

func (gid GID) String() string {
	return base64.RawURLEncoding.EncodeToString(gid[:])
}

func (gid GID) MarshalText() ([]byte, error) {
	enc := base64.RawURLEncoding
	buf := make([]byte, enc.EncodedLen(len(gid)))
	enc.Encode(buf, gid[:])
	return buf, nil
}

func (gid *GID) UnmarshalText(encoded []byte) error {
	enc := base64.RawURLEncoding
	dst := make([]byte, enc.DecodedLen(len(encoded)))
	n, err := enc.Decode(dst, encoded)
	if err != nil {
		return err
	}

	if n != GIDSize {
		return fmt.Errorf("invalid length for GID: got %d, want %d", n, GIDSize)
	}

	copy((*gid)[:], dst)
	return nil
}
