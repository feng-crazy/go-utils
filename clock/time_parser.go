package clock

import (
	"fmt"
	"strconv"
	"time"
)

type MilliTime time.Time

// MarshalJSON implements json.Marshaler.
func (t *MilliTime) MarshalJSON() ([]byte, error) {
	// do your serializing here
	stamp := fmt.Sprintf("%d", time.Time(*t).UnixNano()/1e6)
	return []byte(stamp), nil
}

func (t *MilliTime) UnmarshalJSON(data []byte) (err error) {
	timestamp, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	now := time.Unix(timestamp/1000, (timestamp%1000)*1e6)
	*t = MilliTime(now)
	return
}
