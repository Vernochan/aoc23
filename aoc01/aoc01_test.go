package main

import (
	"testing"
)

func Test_getCoordinate(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// 1abc2
		// pqr3stu8vwx
		// a1b2c3d4e5f
		// treb7uchet
		{"1abc2", args{"1abc2"}, 12},
		{"pqr3stu8vwx", args{"pqr3stu8vwx"}, 38},
		{"a1b2c3d4e5f", args{"a1b2c3d4e5f"}, 15},
		{"treb7uchet", args{"treb7uchet"}, 77},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCoordinate(tt.args.line); got != tt.want {
				t.Errorf("getCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTextCoordinate(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		// {"two1nine", args{"two1nine"}, 29},
		// {"eightwothree", args{"eightwothree"}, 83},
		// {"abcone2threexyz", args{"abcone2threexyz"}, 13},
		// {"xtwone3four", args{"xtwone3four"}, 24},
		// {"4nineeightseven2", args{"4nineeightseven2"}, 42},
		// {"zoneight234", args{"zoneight234"}, 14},
		// {"7pqrstsixteen", args{"7pqrstsixteen"}, 76},
		// {"eighthree", args{"eighthree"}, 83},
		// {"xkbseventwotwogmkxhpmhm42hvvbfchreight", args{"xkbseventwotwogmkxhpmhm42hvvbfchreight"}, 78},
		// {"45122", args{"45122"}, 42},
		// {"rqrrdrmlfsixfive6", args{"rqrrdrmlfsixfive6"}, 66},
		// {"six6v", args{"six6v"}, 66},
		{"4t", args{"4t"}, 44},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTextCoordinate(tt.args.line); got != tt.want {
				t.Errorf("getTextCoordinate() = %v, want %v (%s)", got, tt.want, tt.name)
			}
		})
	}
}
