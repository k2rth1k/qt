package resource

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
)

const (
	// initial resource domain - all IDs generated by this version will have this prefix
	resourceIDDomain = "aa"
	// resource ID random component length: 6 bytes (48 bits), 281,474,976,710,655 values
	resourceIDByteLength = 6
	// resource ID template in the following regex format: ([a-z]{1,4})-([A-z0-9]{2})([A-z0-9-_]{8})
	resourceIDTemplate = "%s-%s%s"
)

// resourceIDRegex is used to validate resource IDs
var resourceIDRegex = regexp.MustCompile("^([a-z]{1,4})-([A-z0-9]{2})([A-z0-9-_]{7,8})$")

// ValidateResourceID validates that the resource ID is in valid format
func ValidateResourceID(id string) error {
	if id == "" {
		return fmt.Errorf("id is empty")
	}
	if !resourceIDRegex.MatchString(id) {
		return fmt.Errorf("id format is invalid")
	}
	return nil
}

// NewResourceID generates a new unique ID that can be used as a primary key
// Generated ID is at most 16 chars long
// Prefix must be between 1-4 chars
// The resulting id will be in format: prfx-aaXXXXXXXX
func NewResourceID(prefix string) (string, error) {
	if len(prefix) == 0 {
		return "", fmt.Errorf("prefix is required")
	}
	if len(prefix) > 4 {
		return "", fmt.Errorf("prefix is too long")
	}

	lowerPrefix := strings.ToLower(prefix)
	s, err := RandomString(resourceIDByteLength)
	if err != nil {
		return "", err
	}

	id := fmt.Sprintf(resourceIDTemplate, lowerPrefix, resourceIDDomain, s)

	return id, nil
}

// RandomBytes generate random bytes using a global, shared instance of a cryptographically
// secure random number generator.
// On Linux, Reader uses getrandom(2) if available, /dev/urandom otherwise.
func RandomBytes(length uint32) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// RandomString creates a url-safe base64-encoded string without padding
func RandomString(bytesLen uint32) (string, error) {
	bytes, err := RandomBytes(bytesLen)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
