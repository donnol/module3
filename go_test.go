package module3

import (
	"testing"
)

func TestLe(t *testing.T) {
	le := NewLe("liu le")
	t.Log(le.String())
}
