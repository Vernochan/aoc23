package main

import (
	"testing"
)

func Test_checkValidity(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
		// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
		// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
		// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
		{"Game 1", args{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"}, true, 1},
		{"Game 2", args{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"}, true, 2},
		{"Game 3", args{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"}, false, 3},
		{"Game 4", args{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"}, false, 4},
		{"Game 5", args{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}, true, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checkValidity(tt.args.line)
			if got != tt.want {
				t.Errorf("checkValidity() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkValidity() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getMinPower(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Game 1", args{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"}, 48},
		{"Game 2", args{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"}, 12},
		{"Game 3", args{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"}, 1560},
		{"Game 4", args{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"}, 630},
		{"Game 5", args{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinPower(tt.args.line); got != tt.want {
				t.Errorf("getMinPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
