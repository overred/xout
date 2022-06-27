package xfield

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
	t.Run("3 Remove First", func(t *testing.T) {
		f := New().
			With("key1", "val1").
			With("key2", "val2").
			With("key3", "val3")
		f = f.Remove("key1")
		if _, ok := f.Find("key1"); ok {
			t.Errorf("fields must provide way to remove field")
		}
	})
	t.Run("4 Remove Middle", func(t *testing.T) {
		f := New().
			With("key1", "val1").
			With("key2", "val2").
			With("key3", "val3")
		f = f.Remove("key2")
		if _, ok := f.Find("key2"); ok {
			t.Errorf("fields must provide way to remove field")
		}
	})
	t.Run("5 Remove Last", func(t *testing.T) {
		f := New().
			With("key1", "val1").
			With("key2", "val2").
			With("key3", "val3")
		f = f.Remove("key3")
		if _, ok := f.Find("key3"); ok {
			t.Errorf("fields must provide way to remove field")
		}
	})
}
