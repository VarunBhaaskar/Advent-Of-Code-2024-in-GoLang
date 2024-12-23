package main

import (
	"reflect"
	"testing"
)

func Test_transform(t *testing.T) {
	type args struct {
		n   int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"qq", args{25, 1}, 29165},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transform(tt.args.n, tt.args.val); got != tt.want {
				t.Errorf("transform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextState(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{"1", args{10}, [2]int{1, 0}}, {"2", args{1}, [2]int{2024, -1}}, {"3", args{0}, [2]int{1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextState(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample", args{[]int{125, 17}}, 55312},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
