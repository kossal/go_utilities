package utilities

import "testing"

func TestFind(t *testing.T) {
	a := []string{"Hello", "World", "Golang"}
	x := "World"
	if result := Find(a, x); result != 1 {
		t.Error("Expected 1, got", result)
	}
}

func TestContains(t *testing.T) {
	a := []string{"Hello", "World", "Golang"}
	x := "World"
	if result := Contains(a, x); result != true {
		t.Error("Expected 1, got", result)
	}
}

func TestParseTemplatesRecursively(t *testing.T) {
	_, err := ParseTemplatesRecursively("./test", "html")
	if err != nil {
		t.Error("Failed to load templates. Got", err)
	}
}
