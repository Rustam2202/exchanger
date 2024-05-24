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
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name: "Amount is 0",
			args: args{
				amount:    0,
				banknotes: []int{5000, 1000, 500, 200, 100, 50}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Banknots empty",
			args: args{
				amount:    400,
				banknotes: []int{}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Amount less then min banknot",
			args: args{
				amount:    10,
				banknotes: []int{5000, 1000, 500, 200, 100, 50}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Amount equal min banknot",
			args: args{
				amount:    50,
				banknotes: []int{5000, 1000, 500, 200, 100, 50}},
			want:    [][]int{{50}},
			wantErr: false,
		},
		{
			name: "Change less min banknot",
			args: args{
				amount:    160,
				banknotes: []int{5000, 1000, 500, 200, 100, 50}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Clone banknotes",
			args: args{
				amount:    100,
				banknotes: []int{5000, 1000, 500, 200, 100, 100, 50, 50}},
			want:    [][]int{{100}, {50, 50}},
			wantErr: false,
		},

		{
			name: "Test from task",
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindCombinations(tt.args.amount, tt.args.banknotes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindCombinations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
