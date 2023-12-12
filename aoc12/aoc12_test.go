package main

import (
	"testing"
)

func Test_getConfigurations(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// ???.### 1,1,3
		// .??..??...?##. 1,1,3
		// ?#?#?#?#?#?#?#? 1,3,1,6
		// ????.#...#... 4,1,1
		// ????.######..#####. 1,6,5
		// ?###???????? 3,2,1
		// TODO: Add test cases.
		{"1", args{"???.### 1,1,3"}, 1},
		{"2", args{".??..??...?##. 1,1,3"}, 4},
		{"3", args{"?#?#?#?#?#?#?#? 1,3,1,6"}, 1},
		{"4", args{"????.#...#... 4,1,1"}, 1},
		{"5", args{"????.######..#####. 1,6,5"}, 4},
		{"6", args{"?###???????? 3,2,1"}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTotalConfigurations(tt.args.line); got != tt.want {
				t.Errorf("getConfigurations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTotalConfigurationsExpanded(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"???.### 1,1,3"}, 1},
		{"2", args{".??..??...?##. 1,1,3"}, 16384},
		{"3", args{"?#?#?#?#?#?#?#? 1,3,1,6"}, 1},
		{"4", args{"????.#...#... 4,1,1"}, 16},
		{"5", args{"????.######..#####. 1,6,5"}, 2500},
		{"6", args{"?###???????? 3,2,1"}, 506250},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTotalConfigurationsExpanded(tt.args.line); got != tt.want {
				t.Errorf("getTotalConfigurationsExpanded() = %v, want %v", got, tt.want)
			}
		})
	}
}
