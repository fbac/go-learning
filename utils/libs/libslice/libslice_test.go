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

func TestPopFirst(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []intruct {
		name string
		args args
		want []int
	}{
		{
			"PopFirst",
			args{
				slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopFirst(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPopLast(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []intruct {
		name string
		args args
		want []int
	}{
		{
			"PopLast",
			args{
				slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopLast(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPopN(t *testing.T) {
	type args struct {
		slice []int
		n     int
	}
	tests := []intruct {
		name string
		args args
		want []int
	}{
		{
			"PopN",
			args{
				slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				n:     2,
			},
			[]int{0, 1, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopN(tt.args.slice, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFirstN(t *testing.T) {
	type args struct {
		slice []int
		n     int
	}
	tests := []intruct {
		name string
		args args
		want []int
	}{
		{
			"GetFirstN",
			args{
				slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				n:     5,
			},
			[]int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFirstN(tt.args.slice, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFirstN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLastN(t *testing.T) {
	type args struct {
		slice []int
		n     int
	}
	tests := []intruct {
		name string
		args args
		want []int
	}{
		{
			"GetLastN",
			args{
				slice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				n:     5,
			},
			[]int{5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLastN(tt.args.slice, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLastN() = %v, want %v", got, tt.want)
			}
		})
	}
}
