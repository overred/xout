package xfields

import "testing"

func Test_Fields(t *testing.T) {
	t.Run("1 Write And Find", func(t *testing.T) {
		f := New().With("key", "val")
		if val, ok := f.Find("key"); !ok || val != "val" {
			t.Errorf("fields must provide way to add and correctly read field")
		}
	})
	t.Run("2 Write And Find", func(t *testing.T) {
		f := New().
			With("key1", "val1").
			With("key2", "val2").
			With("key3", "val3")
		if val, ok := f.Find("key2"); !ok || val != "val2" {
			t.Errorf("fields must provide way to add and correctly read field")
		}
	})
}
