// https://adventofcode.com/2023/day/2
// Day 2: Cube Conundrum
package main

import "testing"

func Test_gamePower(t *testing.T) {
	type args struct {
		game string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				game: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			},
			want: 48,
		},
		{
			args: args{
				game: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			},
			want: 12,
		},
		{
			args: args{
				game: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			},
			want: 1560,
		},
		{
			args: args{
				game: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			},
			want: 630,
		},
		{
			args: args{
				game: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gamePower(tt.args.game); got != tt.want {
				t.Errorf("gamePower() = %v, want %v", got, tt.want)
			}
		})
	}
}
