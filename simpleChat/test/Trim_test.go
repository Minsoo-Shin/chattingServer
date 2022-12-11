package test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStringsTrim(t *testing.T) {
	type args struct {
		sentence string
	}

	tests := []struct {
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				sentence: `   A B C D`,
			},
			want:    "A B C D",
			wantErr: false,
		},
		{
			args: args{
				sentence: `
			A B C D`,
			},
			want:    "\n			A B C D",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		got := strings.Trim(tt.args.sentence, " ")
		assert.Equal(t, tt.want, got)
	}

}
