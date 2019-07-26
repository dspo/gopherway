package filesystem

import "testing"

func TestParse(t *testing.T) {
	parsed, err := ParseHero(tony)
	if err == nil {
		t.Log(parsed)
	}
	parsed, err = ParseHero(steve)
	if err == nil {
		t.Log(parsed)
	}
}
