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
				input: []string{"1abc2",
					"pqr3stu8vwx",
					"a1b2c3d4e5f",
					"treb7uchet"},
			},
			wantNumbers: []int{12, 38, 15, 77},
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
