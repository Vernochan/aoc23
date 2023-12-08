package main

import (
	"testing"
)

func Test_handValueWithJokers(t *testing.T) {
	type args struct {
		h hand
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// high card: 0, one pair: 0x100000, two pair: 0x200000, three of a kind: 0x300000, fullHouse: 0x400000, Four of a kind: 0x500000, Five of a kind: 0x600000
		{"QJJTQ", args{hand{bid: 0, rank: 0, cards: []byte{'Q', 'J', 'J', 'T', 'Q'}}}, 0x5C00AC},
		{"8QTJ3", args{hand{bid: 0, rank: 0, cards: []byte{'8', 'Q', 'T', 'J', '3'}}}, 0x18CA03},
		{"8KJ94", args{hand{bid: 0, rank: 0, cards: []byte{'8', 'K', 'J', '9', '4'}}}, 0x18D094},
		{"AJ888", args{hand{bid: 0, rank: 0, cards: []byte{'A', 'J', '8', '8', '8'}}}, 0x5E0888},
		{"JJJJJ", args{hand{bid: 0, rank: 0, cards: []byte{'J', 'J', 'J', 'J', 'J'}}}, 0x600000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handValueWithJokers(tt.args.h); got != tt.want {
				t.Errorf("%s handValueWithJokers() = %x, want %x", tt.name, got, tt.want)
			}
		})
	}
}
