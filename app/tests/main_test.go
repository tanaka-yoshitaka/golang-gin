package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{a: 1, b: 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Square ２つのintを受け取りそれらを掛け合わせた物を返す関数
func Square(a, b int) int {
	return a * b
}

func TestSquere(t *testing.T) {
    asserts := assert.New(t)
    for _, td := range []struct {
        title  string
        input  []int
        output int
    }{
        {
            title:  "2×3の答えが6になる",
            input:  []int{2, 3},
            output: 6,
        },
        {
            title:  "3×5の答えが15になる",
            input:  []int{3, 5},
            output: 15,
        },
    } {
        t.Run("Square:"+td.title, func(t *testing.T) {
            result := Square(td.input[0], td.input[1])
            asserts.Equal(td.output, result)
        })
    }
}
