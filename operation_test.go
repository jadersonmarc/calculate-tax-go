package main

import "testing"

func TestOperationType(t *testing.T) {
	op := Operation{
		Type: Buy,
	}

	if op.Type != "buy" {
		t.Errorf("expected buy, got %s", op.Type)
	}
}
