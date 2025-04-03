package data_structures

import (
	"testing"
)

func TestInsertion(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	if m["a"] != 1 || m["b"] != 2 {
		t.Errorf("Insertion failed, expected m['a']=1 and m['b']=2, got %v", m)
	}
}

func TestUpdate(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["a"] = 10
	if m["a"] != 10 {
		t.Errorf("Update failed, expected m['a']=10, got %d", m["a"])
	}
}

func TestExistence(t *testing.T) {
	m := make(map[string]int)
	m["b"] = 2
	if _, exists := m["b"]; !exists {
		t.Errorf("Existence check failed, key 'b' should exist")
	}
}

func TestDeletion(t *testing.T) {
	m := make(map[string]int)
	m["b"] = 2
	delete(m, "b")
	if _, exists := m["b"]; exists {
		t.Errorf("Deletion failed, key 'b' should not exist")
	}
}

func TestNonExistentKey(t *testing.T) {
	m := make(map[string]int)
	if val, exists := m["c"]; exists {
		t.Errorf("Non-existent key check failed, expected key 'c' to not exist but got %d", val)
	}
}
