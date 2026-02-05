package addgomodule

import "testing"

func TestAdd(t *testing.T) {
	if sum := Add(1, 2); sum != 3 {
		t.Errorf("expected size 3, got %d", 3)
	}
}
