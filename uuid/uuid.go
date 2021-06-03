package uuid

import (
	"strings"

	guuid "github.com/google/uuid"
	"github.com/satori/go.uuid"
)

// NewV1 returns UUID based on current timestamp and MAC address.
func GenerateUUID() string {
	sig := uuid.NewV1().String()
	id := strings.Replace(sig, "-", "", -1)
	return id
}

// NewV1 returns UUID based on current timestamp and MAC address.
func GenerateUUIDv1() string {
	return uuid.NewV1().String()
}

// NewUUID returns a Version 1 UUID based on the current NodeID and clock
// sequence, and the current time.  If the NodeID has not been set by SetNodeID
// or SetNodeInterface then it will be set automatically.  If the NodeID cannot
// be set NewUUID returns nil.  If clock sequence has not been set by
// SetClockSequence then it will be set automatically.  If GetTime fails to
// return the current NewUUID returns nil and an error.
func NewUUID() string {
	sig, _ := guuid.NewUUID()
	id := strings.Replace(sig.String(), "-", "", -1)
	return id
}

func NewUUIDv1() string {
	sig, _ := guuid.NewUUID()
	return sig.String()
}
