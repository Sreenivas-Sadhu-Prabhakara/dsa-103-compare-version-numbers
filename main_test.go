package main

import "testing"

func TestCompareVersions(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"1.2", "1.10", -1}, // numeric, not lexicographic
		{"1.01", "1.001", 0},
		{"1.0", "1.0.0", 0},
		{"0.1", "1.1", -1},
		{"1.0.1", "1", 1},
		{"7.5.2.4", "7.5.3", -1},
		{"2", "2", 0},
	}
	for _, c := range cases {
		if got := CompareVersions(c.a, c.b); got != c.want {
			t.Errorf("CompareVersions(%q,%q) = %d, want %d", c.a, c.b, got, c.want)
		}
		if got := CompareVersions(c.b, c.a); got != -c.want {
			t.Errorf("CompareVersions(%q,%q) = %d, want %d (antisymmetry)", c.b, c.a, got, -c.want)
		}
	}
}
