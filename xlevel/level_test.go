package xlevel

import "testing"

func Test_LevelMask_HasLevel(t *testing.T) {
	t.Run("1 Single Level Contain", func(t *testing.T) {
		mask := Warn
		if !mask.Has(Warn) {
			t.Errorf("single level must contains self")
		}
	})
	t.Run("2 Single Level Not Contain", func(t *testing.T) {
		mask := Warn
		if mask.Has(Info) {
			t.Errorf("single level must not contain any other levels")
		}
	})
	t.Run("3 Complex Level Contain", func(t *testing.T) {
		mask := Warn | Info | Panic
		if !mask.Has(Info) {
			t.Errorf("complex level must contains every of their simple levels")
		}
	})
	t.Run("4 Complex Level Not Contain", func(t *testing.T) {
		mask := Warn | Info | Panic
		if mask.Has(Fatal) {
			t.Errorf("complex level must not contain any other levels except their simple levels")
		}
	})
	t.Run("5 Zero Level Contain", func(t *testing.T) {
		mask := Level(0)
		if !mask.Has(0) {
			t.Errorf("zero level must contains zero level only")
		}
	})
	t.Run("6 Not Zero Level Not Contain", func(t *testing.T) {
		mask := Text
		if mask.Has(0) {
			t.Errorf("not zero level must not contain zero level")
		}
	})
	t.Run("7 Zero Level Not Contain", func(t *testing.T) {
		mask := Level(0)
		if mask.Has(Text) {
			t.Errorf("zero level must not contain anything except zero level")
		}
	})
}

func Test_LevelElevate(t *testing.T) {
	t.Run("1 Zero Means All", func(t *testing.T) {
		mask := Level(0).Elevate()
		if !mask.Has(Panic) {
			t.Errorf("elevating from 0 must includes all possible levels")
		}
	})
	t.Run("2 Upper Than Not Contain", func(t *testing.T) {
		mask := Level(Debug).Elevate()
		if mask.Has(Trace) {
			t.Errorf("elevating from specific simple level must not include levels below")
		}
	})
	t.Run("3 Upper Than Equal", func(t *testing.T) {
		mask := Level(Debug).Elevate()
		if !mask.Has(Debug) {
			t.Errorf("elevating from specific simple level must includes this level also")
		}
	})
	t.Run("4 Upper Than Contain", func(t *testing.T) {
		mask := Level(Debug).Elevate()
		if !mask.Has(Info) {
			t.Errorf("elevating from specific simple level must includes all levels above")
		}
	})
}

func Test_LevelParse(t *testing.T) {
	t.Run("1 Single Level Parsing", func(t *testing.T) {
		mask := Parse("trace")
		if mask != Trace {
			t.Errorf("level parser must correctly parse single level")
		}
	})
	t.Run("2 Empty Level Parsing", func(t *testing.T) {
		mask := Parse("")
		if mask != 0 {
			t.Errorf("level parser must produces zero level if string doesn't contain any level")
		}
	})
	t.Run("3 Multiple Level Parsing", func(t *testing.T) {
		mask := Parse("debug|info")
		if !mask.Has(Debug) || !mask.Has(Info) {
			t.Errorf("level parser must correct produces complex level mask for string with several levels")
		}
	})
	t.Run("4 Multiple Level Parsing", func(t *testing.T) {
		mask := Parse("debug|info|warn|error|fatal")
		if mask.Has(Text) || mask.Has(Trace) || mask.Has(Panic) {
			t.Errorf("level parser must contains only parsed levels")
		}
	})
}

func Test_LevelString(t *testing.T) {
	t.Run("1 Single Level String", func(t *testing.T) {
		mask := Debug
		if mask.String() != "debug" {
			t.Errorf("level stringer must correctly produces single level")
		}
	})
	t.Run("2 Zero Level String", func(t *testing.T) {
		mask := Level(0)
		if mask.String() != "" {
			t.Errorf("level stringer must correctly produces zero level")
		}
	})
	t.Run("3 Multiple Level String", func(t *testing.T) {
		mask := Error | Fatal | Panic
		if mask.String() != "error|fatal|panic" {
			t.Errorf("level stringer must correctly produces multiple levels")
		}
	})
	t.Run("4 Multiple Level String", func(t *testing.T) {
		mask := Error | Panic | Fatal
		if mask.String() != "error|fatal|panic" {
			t.Errorf("level stringer must sort multiple levels by importance ast")
		}
	})
}
