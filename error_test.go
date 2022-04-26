package tiktok

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestError_Error(t *testing.T) {
	var err error
	e := Error("123")
	err = e
	require.Equal(t, "123", err.Error())
}

func TestAPIError_Error(t *testing.T) {
	err := &APIError{
		Code:      1,
		Message:   "123",
		RequestID: "1234",
	}
	require.Equal(t, "1: 123", err.Error())
}

func TestErrCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "APIError",
			args: args{
				err: &APIError{
					Code:      1,
					Message:   "123",
					RequestID: "1234",
				},
			},
			want: 1,
		},
		{
			name: "Error",
			args: args{
				err: Error("123"),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, ErrCode(tt.args.err))
		})
	}
}
