package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("5 nums", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if want != got {
			t.Errorf("want %d but got %d, from %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 9})
	want := []int{3, 12}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, but got %v", want, got)
	}
}

func TestSumAllTail(t *testing.T) {

	checkSum := func(t *testing.T, want []int, got []int) {
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, but got %v", want, got)
		}
	}

	t.Run("Happy-path", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 9})
		want := []int{2, 9}
		checkSum(t, want, got)
	})

	t.Run("Empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 9, 6})
		want := []int{0, 15}
		checkSum(t, want, got)
	})

}
