package services

import (
	"strconv"
	"testing"
)

func TestGenerateRoomIDIsSixDigits(t *testing.T) {
	for i := 0; i < 100; i++ {
		id := GenerateRoomID()
		if len(id) != 6 {
			t.Fatalf("room ID %q is not six characters", id)
		}
		n, err := strconv.Atoi(id)
		if err != nil {
			t.Fatalf("room ID %q is not numeric: %v", id, err)
		}
		if n < 100000 || n > 999999 {
			t.Fatalf("room ID %q is outside 100000-999999", id)
		}
	}
}
