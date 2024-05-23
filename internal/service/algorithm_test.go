package service

import (
	"reflect"
	"testing"
)

func TestFindCombinations(t *testing.T) {
	type args struct {
		amount    int
		banknotes []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				amount:    400,
				banknotes: []int{5000, 1000, 500, 200, 100, 50}},
			want: [][]int{
				{200, 200},
				{200, 100, 100},
				{200, 100, 50, 50},
				{200, 50, 50, 50, 50},
				{100, 100, 100, 100},
				{100, 100, 100, 50, 50},
				{100, 100, 50, 50, 50, 50},
				{100, 50, 50, 50, 50, 50, 50},
				{50, 50, 50, 50, 50, 50, 50, 50},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCombinations(tt.args.amount, tt.args.banknotes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
