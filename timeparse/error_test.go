package timeparse

import "testing"

func TestTimeParse(t *testing.T) {
	table := []struct {
		input string
		want  bool
	}{
		{"11:22:12", true},
		{"60:22:12", false},
		{"0:10:02", true},
		{"", false},
		{"01:-22:12", false},
		{"01::12", false},
		{"23:59:59", true},
		{"0:00:00", true},
	}
	for _, data := range table {
		_, err := ParseTime(data.input)
		if err != nil && data.want {
			t.Errorf("Failed %v : Expected: %v : Actual: %v", data.input, data.want, err)
		}
	}
}
