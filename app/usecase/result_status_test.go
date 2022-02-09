package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResultStatus(t *testing.T) {
	type args struct {
		statusCode int
		err        error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "200ステータス",
		// 	args: args{
		// 		statusCode: 200,
		// 		err:        nil,
		// 	},
		// 	want: 200,
		// },
		// {
		// 	name: "201ステータス",
		// 	args: args{
		// 		statusCode: 201,
		// 		err:        nil,
		// 	},
		// 	want: 201,
		// },
		{
			name: "404ステータス",
			args: args{
				statusCode: 404,
				err:        errors.New("not found"),
			},
			want: 404,
		},
	}
	asserts := assert.New(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewResultStatus(tt.args.statusCode, tt.args.err)
			asserts.Equal(tt.want, r.StatusCode)
			// asserts.Equal("not found", r.Error)
			asserts.EqualError(r.Error, "not found")
		})
	}
}
