package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		roundsCount int
		want        string
	}{
		{-10, ""},
		{0, ""},
		{1, "1"},
		{2, "1, 2"},
		{3, "1, 2, Fizz"},
		{5, "1, 2, Fizz, 4, Buzz"},
		{16, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16"},
	}

	for _, c := range cases {
		got := FizzBuzz(c.roundsCount)
		if got != c.want {
			t.Errorf("FizzBuzz(%v) == %v, want %v", c.roundsCount, got, c.want)
		}
	}
}

func BenchmarkFizzBuzz10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(10)
	}
}

func BenchmarkFizzBuzz1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(1000)
	}
}

func BenchmarkFizzBuzz1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(1000000)
	}
}
