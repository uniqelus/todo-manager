package task_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/uniqelus/todo-manager/internal/domains/server/task"
)

func TestNewPriority(t *testing.T) {
	t.Parallel()

	type args struct {
		value string
	}

	tests := []struct {
		name    string
		args    args
		wantRes assert.ValueAssertionFunc
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "empty value",
			args: args{
				value: "",
			},
			wantRes: func(tt assert.TestingT, i1 any, _ ...any) bool {
				got, ok := i1.(task.Priority)
				return assert.True(tt, ok) && assert.Equal(tt, task.UndefinedPriority, got)
			},
			wantErr: assert.NoError,
		},
		{
			name: "invalid value",
			args: args{
				value: "invalid",
			},
			wantRes: assert.Empty,
			wantErr: assert.Error,
		},
		{
			name: "created",
			args: args{
				value: "low",
			},
			wantRes: func(tt assert.TestingT, i1 any, _ ...any) bool {
				got, ok := i1.(task.Priority)
				return assert.True(tt, ok) && assert.Equal(tt, task.LowPriority, got)
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res, err := task.NewPriority(tt.args.value)

			tt.wantRes(t, res)
			tt.wantErr(t, err)
		})
	}
}
