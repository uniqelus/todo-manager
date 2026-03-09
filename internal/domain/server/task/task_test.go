package task_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/uniqelus/todo-manager/internal/domain/server/task"
)

func TestNewTask(t *testing.T) {
	t.Parallel()

	type args struct {
		title       string
		description string
		priority    string
		dueDate     time.Time
	}

	tests := []struct {
		name    string
		args    args
		wantRes assert.ValueAssertionFunc
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "cannot use gotten priority value",
			args: args{
				title:       "test task",
				description: "",
				dueDate:     time.Now(),
				priority:    "unknown priority",
			},
			wantRes: assert.Nil,
			wantErr: assert.Error,
		},
		{
			name: "empty title",
			args: args{
				title:       "",
				description: "",
				dueDate:     time.Now(),
				priority:    "",
			},
			wantRes: assert.Nil,
			wantErr: assert.Error,
		},
		{
			name: "task created",
			args: args{
				title:       "test task",
				description: "",
				dueDate:     time.Now(),
				priority:    "",
			},
			wantRes: assert.NotNil,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res, err := task.NewTask(tt.args.title, tt.args.description, tt.args.dueDate, tt.args.priority)

			tt.wantRes(t, res)
			tt.wantErr(t, err)
		})
	}
}
