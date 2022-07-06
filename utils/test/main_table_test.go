// package main as expected
package testdocs

// import testing
import "testing"

// Test_potion_String still honors the naming convention
func Test_potion_String(t *testing.T) {

	// this anon struct represents a potion
	type fields struct {
		id   int
		name string
	}

	// this is a slice of tests, containing:
	// name: name of the test
	// fields: arguments (actually, the potion itself)
	// want: expected results
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// add table test here
		// with the same syntax as filling
		// a slice of structs
		{
			"test: health",
			fields{0, "health"},
			"health",
		},
		{
			"test: poison",
			fields{0, "poison"},
			"poison",
		},
		{
			"test: agility",
			fields{0, "agility"},
			"agility",
		},
		{
			"test: hd12345",
			fields{0, "hd12345"},
			"hd12345",
		},
		{
			"test: whatever👌🏻",
			fields{0, "whatever👌🏻"},
			"whatever👌🏻",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Potion{
				Id:   tt.fields.id,
				Name: tt.fields.name,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("potion.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
