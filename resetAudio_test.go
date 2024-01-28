// main_test.go
package main

import (
	"testing"
)

func TestRestartCoreAudio(t *testing.T) {
	err := restartCoreAudio()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
