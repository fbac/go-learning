/*
	libslice
	author: fbac <me@fbac.dev>

	example of slice library to add
	some missing funcionality in stdlib
*/
package libslice

import (
	"reflect"
	"testing"
)

func Test_popFirst(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"popFirst",
			args{
				slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := popFirst(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("popFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_popLast(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"popLast",
			args{
				slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := popLast(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("popLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_popN(t *testing.T) {
	type args struct {
		slice []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"popN",
			args{
				slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				n:     2,
			},
			[]int{0, 1, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := popN(tt.args.slice, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("popN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFirstN(t *testing.T) {
	type args struct {
		slice []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"getFirstN",
			args{
				slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				n:     5,
			},
			[]int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFirstN(tt.args.slice, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFirstN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLastN(t *testing.T) {
	type args struct {
		slice []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"getLastN",
			args{
				slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				n:     5,
			},
			[]int{5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastN(tt.args.slice, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLastN() = %v, want %v", got, tt.want)
			}
		})
	}
}
