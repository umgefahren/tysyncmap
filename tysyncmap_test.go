package tysyncmap

import "testing"

func TestString(t *testing.T) {
	m := new(Map[string, string])
	m.Store("key", "value")
	actual, loaded := m.LoadOrStore("key", "value2")
	if !loaded {
		t.Error("Should be loaded")
	}
	if actual != "value" {
		t.Error("Unexpected actual value")
	}
	value, ok := m.Load("key")
	if !ok {
		t.Error("Should have been loaded")
	}
	if value != "value" {
		t.Error("Value was not expected")
	}
	noValue, ok := m.Load("key2")
	if ok {
		t.Error("Should not have been loaded")
	}
	if noValue != "" {
		t.Error("should have been nil")
	}
	value, loaded = m.LoadAndDelete("key2")
	if loaded {
		t.Error("Should not have been loaded")
	}
	if value != "" {
		t.Error("should have been nil")
	}
	value, loaded = m.LoadAndDelete("key")
	if value != "value" {
		t.Error("value was not expected")
	}
	if !loaded {
		t.Error("not loaded")
	}
}

func TestRef(t *testing.T) {
	m := new(Map[string, *uint])
	val, r := m.Load("m")
	if r {
		t.Error("Should not have been in there")
	}
	if val != nil {
		t.Error("is not nil")
	}
}
