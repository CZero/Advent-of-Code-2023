// https://adventofcode.com/2023/day/1
// Day 1: Trebuchet?!
package main

import (
	"reflect"
	"testing"
)

func Test_getNumbers(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name        string
		args        args
		wantNumbers []int
	}{
		{
			name: "Example",
			args: args{
				input: []string{"two1nine",
					"eightwothree",
					"abcone2threexyz",
					"xtwone3four",
					"4nineeightseven2",
					"zoneight234",
					"7pqrstsixteen"},
			},
			wantNumbers: []int{29, 83, 13, 24, 42, 14, 76},
		},
		{
			name: "2e test",
			args: args{
				input: []string{"one234one",
					"one237two",
					"one234three",
					"one2345four",
					"one234five",
					"one234six",
					"one234seven",
					"one234eight",
					"one234nine"},
			},
			wantNumbers: []int{11, 12, 13, 14, 15, 16, 17, 18, 19},
		},
		{
			name: "3e test",
			args: args{
				input: []string{"one234one",
					"two237one",
					"three234one",
					"four2345one",
					"five234one",
					"six234one",
					"seven234one",
					"eight234one",
					"nine234one"},
			},
			wantNumbers: []int{11, 21, 31, 41, 51, 61, 71, 81, 91},
		},
		{
			name: "4e test",
			args: args{
				input: []string{"on234one",
					"tw237one",
					"thee234one",
					"for2345one",
					"fie234one",
					"si234one",
					"seen234one",
					"eiht234one",
					"nie234one"},
			},
			wantNumbers: []int{21, 21, 21, 21, 21, 21, 21, 21, 21},
		},
		{
			name: "5e test",
			args: args{
				input: []string{"on234on",
					"tw237on",
					"thee234oe",
					"for2345oe",
					"fie234ond",
					"si234ond",
					"seen234ond",
					"eiht234ond",
					"nie234ond"},
			},
			wantNumbers: []int{24, 27, 24, 25, 24, 24, 24, 24, 24},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumbers := getNumbers(tt.args.input); !reflect.DeepEqual(gotNumbers, tt.wantNumbers) {
				t.Errorf("getNumbers() = %v, want %v", gotNumbers, tt.wantNumbers)
			}
		})
	}
}
