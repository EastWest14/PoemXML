package extractor_test

import (
	"testing"
	. "poemXML/poemserver/handlers/extractor"
)

func TestNew(t *testing.T) {
	anExtractor := New()

	if anExtractor == nil {
		t.Errorf("Failed to initialize extractor, it is nil")
	}
}