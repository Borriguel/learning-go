package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name  string
		Entry interface{}
		Calls []string
	}{
		{
			"Struct",
			struct {
				Name string
			}{"Rodo"},
			[]string{"Rodo"},
		},
		{
			"Struct with 2 fields type string",
			struct {
				Name string
				City string
			}{"Rodo", "RJ"},
			[]string{"Rodo", "RJ"},
		}, {
			"Struct without field type string",
			struct {
				Name string
				Age  int
			}{"Rodo", 29},
			[]string{"Rodo"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Entry, func(entry string) {
				got = append(got, entry)
			})
			if !reflect.DeepEqual(got, test.Calls) {
				t.Errorf("got %q, want %q", got, test.Calls)
			}
		})
	}
}
