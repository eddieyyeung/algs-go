package patternsearch

import (
	"reflect"
	"testing"
)

func TestRabinKarpSearch(t *testing.T) {
	type args struct {
		txt string
		pat string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				txt: "geeksforgeeks",
				pat: "geeks",
			},
			want: []int{0, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RabinKarpSearch(tt.args.txt, tt.args.pat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RabinKarpSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
