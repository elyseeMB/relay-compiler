package gid

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

type TenantID [8]byte

var (
	NilTenant = TenantID{}

	defaultTenantGenerator = newTenantGenerator()
)

type tenantGenerator struct {
	machineID [3]byte // 3 * 8 24 bits
	counter   uint32
}

func NewTenantID() TenantID {
	return defaultTenantGenerator.NewTenantID()
}

func newTenantGenerator() *tenantGenerator {
	g := &tenantGenerator{
		counter: 0,
	}

	if _, err := rand.Read(g.machineID[:]); err != nil {
		hostname, _ := os.Hostname()
		copy(g.machineID[:], []byte(hostname))

		if len(hostname) < len(g.machineID) {
			ts := time.Now().UnixNano()
			binary.BigEndian.PutUint16(g.machineID[len(hostname):], uint16(ts))
		}

	}

	return g
}

func (g *tenantGenerator) NewTenantID() TenantID {
	// create new ID
	var id TenantID

	// 1. Copy machineID (first 3 bytes)
	copy(id[0:3], g.machineID[:])

	// 2. Add timeStamp (next 3 bytes)
	now := uint32(time.Now().Unix())
	id[3] = byte(now >> 16)
	id[4] = byte(now >> 8)
	id[5] = byte(now)

	// 3. Increment counter atomically
	count := atomic.AddUint32(&g.counter, 1) & 0xFFFF
	binary.BigEndian.PutUint16(id[6:8], uint16(count))

	return id
}

// ParseTenantID parses a string representation into a TenantID
func ParseTenantID(s string) (TenantID, error) {
	var id TenantID
	decoded, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return NilTenant, fmt.Errorf("invalid tenant ID encoding: %w", err)
	}

	if len(decoded) != len(id) {
		return NilTenant, fmt.Errorf("invalid tenant ID length: got %d, want %d", len(decoded), len(id))
	}

	copy(id[:], decoded)
	return id, nil
}
