package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		name  string
		input interface{}
		want  []string
	}{
		{
			"struct with one string",
			struct {
				Name string
			}{"Julia"},
			[]string{"Julia"},
		},
		{
			"struct with 2 strings",
			struct {
				Name string
				City string
			}{"Julia", "Sg"},
			[]string{"Julia", "Sg"},
		},

		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Julia", 10},
			[]string{"Julia"},
		},

		{
			"nested fields",
			Person{"Julia", Profile{10, "Sg"}},
			[]string{"Julia", "Sg"},
		},
		{
			"pointer ",
			&Person{"Julia", Profile{10, "Sg"}},
			[]string{"Julia", "Sg"},
		},

		{
			"Slice ",
			[]Profile{
				{33, "SG"},
				{34, "London"},
			},
			[]string{"SG", "London"},
		},
		{
			"Array ",
			[2]Profile{
				{33, "SG"},
				{34, "London"},
			},
			[]string{"SG", "London"},
		},
	}

	for _, tt := range cases {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			t.Helper()
			var got []string
			walk(tc.input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want %v, but got %v", tc.want, got)
			}
		})
	}

}

func TestMap(t *testing.T) {
	input := map[string]string{
		"Foo":  "SG",
		"Foo2": "London",
	}

	want := map[string]int{"SG": 1, "London": 1}

	var got = make(map[string]int)
	walk(input, func(input string) {
		got[input] = 1
	})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, but got %v", want, got)
	}
}

func TestChan(t *testing.T) {
	ch := make(chan Profile, 10)

	go func() {
		ch <- Profile{33, "SG"}
		ch <- Profile{34, "HN"}
		close(ch)
	}()

	want := []string{"SG", "HN"}

	var got []string
	walk(ch, func(input string) {
		got = append(got, input)
	})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, but got %v", want, got)
	}
}

func TestFunc(t *testing.T) {
	fun := func() (Profile, Profile) {
		return Profile{31, "SG"}, Profile{34, "HN"}
	}

	want := []string{"SG", "HN"}

	var got []string
	walk(fun, func(input string) {
		got = append(got, input)
	})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, but got %v", want, got)
	}
}
