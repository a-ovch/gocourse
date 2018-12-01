package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		roundsCount int
		expected    string
	}{
		{0, ""},
		{1, "1"},
		{2, "1, 2"},
		{3, "1, 2, Fizz"},
		{5, "1, 2, Fizz, 4, Buzz"},
		{16, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16"},
	}

	for _, c := range cases {
		got := FizzBuzz(c.roundsCount)
		if got != c.expected {
			t.Errorf("FizzBuzz(%v) == %v, want %v", c.roundsCount, got, c.expected)
		}
	}
}
