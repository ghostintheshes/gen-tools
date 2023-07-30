package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {
	type args struct {
		src   []int
		index int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		want1   int
		wantErr error
	}{
		{
			name: "out of range of negative",
			args: args{
				src:   []int{0, 1, 2, 3},
				index: -1,
			},
			wantErr: ErrOutOfRange(4, -1),
		},
		{
			name: "delete 1st element",
			args: args{
				src:   []int{0, 1, 2, 3},
				index: 0,
			},
			want:  []int{1, 2, 3},
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Delete(tt.args.src, tt.args.index)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestDeleteExample(t *testing.T) {
	var case1 = []int{1, 2, 3, 4}
	deletecase, _, err := Delete[int](case1, 0)
	if err != nil {
		return
	}
	fmt.Printf("before delete %v\nafter delete %v\n", case1, deletecase)
}
