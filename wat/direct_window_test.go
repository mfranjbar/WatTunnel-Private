package wat

import (
	"testing"
	"time"
)

func TestSetDirectWindow(t *testing.T) {
	original := relDirectWin
	t.Cleanup(func() { relDirectWin = original })

	SetDirectWindow(350 * time.Millisecond)
	if got := relDirectWin; got != 350*time.Millisecond {
		t.Fatalf("SetDirectWindow = %v, want 350ms", got)
	}

	SetDirectWindow(0)
	if got := relDirectWin; got != time.Second {
		t.Fatalf("SetDirectWindow(0) = %v, want 1s default", got)
	}
}
