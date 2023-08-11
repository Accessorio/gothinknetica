package homework7

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestSort_Ints(t *testing.T) {
	got := []int{5, 4, 3, 2, 1, 0}
	sort.Ints(got)
	want := []int{0, 1, 2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sort.Ints() = %v, want %v", got, want)
	}
}

func TestSort_Strings(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		got  args
		want args
	}{
		{
			name: "Test: 1",
			got: args{
				s: []string{
					"3", "2", "1",
				},
			},
			want: args{
				s: []string{
					"1", "2", "3",
				},
			},
		},
		{
			name: "Test: 2",
			got: args{
				s: []string{
					"B", "C", "A",
				},
			},
			want: args{
				s: []string{
					"A", "B", "C",
				},
			},
		},
		{
			name: "Test: 3",
			got: args{
				s: []string{
					"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta",
				},
			},
			want: args{
				s: []string{
					"Alpha", "Bravo", "Delta", "Go", "Gopher", "Grin",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sort.Strings(test.got.s)
			if !reflect.DeepEqual(test.got, test.want) {
				t.Errorf("sort.Ints() = %v, want %v", test.got, test.want)
			}
		})
	}
}

func sampleIntData() []int {
	rand.Seed(time.Now().UnixNano())
	var data []int
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Intn(1000))
	}

	return data
}

func sampleFloatData() []float64 {
	rand.Seed(time.Now().UnixNano())
	var data []float64
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Float64())
	}

	return data
}

func BenchmarkSortInt(b *testing.B) {
	data := sampleIntData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(data)
		_ = data
	}
}

func BenchmarkSortFloat(b *testing.B) {
	data := sampleFloatData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Float64s(data)
		_ = data
	}
}
