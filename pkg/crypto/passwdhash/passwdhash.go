package passwdhash

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

type (
	Profile struct {
		iterations uint32
		saltLength uint
		KeyLength  uint
		pepper     []byte
	}
)

const (
	versionByte   = 0x01
	algorithmByte = 0x01

	minIterations = 600000
)

func NewProfile(pepper []byte, iterations uint32) (*Profile, error) {

	if len(pepper) < 32 {
		return nil, fmt.Errorf("pepper must be at least 32 bytes")
	}

	if iterations < minIterations {
		return nil, fmt.Errorf("iterations below minimun security threshold")
	}

	return &Profile{
		iterations: iterations,
		saltLength: 32,
		KeyLength:  32,
		pepper:     pepper,
	}, nil
}

func (hp Profile) applyPepper(input []byte) []byte {
	mac := hmac.New(sha256.New, hp.pepper)
	mac.Write(input)
	return mac.Sum(nil)
}

func (hp Profile) HashPassword(password []byte) ([]byte, error) {

	salt := make([]byte, hp.saltLength)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("error generation salt: %w", err)
	}

	pepperedPassword := hp.applyPepper([]byte(password))
	fmt.Printf("Password hashed : %x", salt)

	hash := pbkdf2.Key(pepperedPassword, salt, int(hp.iterations), int(hp.KeyLength), sha256.New)

	binaryHash := make([]byte, 0, 7+hp.saltLength+hp.KeyLength)

	binaryHash = append(binaryHash, versionByte)
	binaryHash = append(binaryHash, algorithmByte)

	iterBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(iterBytes, hp.iterations)
	binaryHash = append(binaryHash, iterBytes...)

	binaryHash = append(binaryHash, byte(hp.saltLength))
	binaryHash = append(binaryHash, salt...)

	binaryHash = append(binaryHash, hash...)

	return binaryHash, nil

}
