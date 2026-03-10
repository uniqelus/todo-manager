package taskserv_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	taskdmn "github.com/uniqelus/todo-manager/internal/domains/server/task"
	tasksrv "github.com/uniqelus/todo-manager/internal/services/server/task"
	"github.com/uniqelus/todo-manager/internal/services/server/task/mocks"
)

func TestService_CreateTask(t *testing.T) {
	t.Parallel()

	type fields struct {
		setupTaskRepository func(m *mocks.Repository)
	}

	type args struct {
		ctx     context.Context
		options *taskdmn.CreateTaskOptions
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult assert.ValueAssertionFunc
		wantError  assert.ErrorAssertionFunc
	}{
		{
			name: "cannot create task model",
			fields: fields{
				setupTaskRepository: func(_ *mocks.Repository) {},
			},
			args: args{
				ctx:     context.Background(),
				options: &taskdmn.CreateTaskOptions{},
			},
			wantResult: assert.Nil,
			wantError:  assert.Error,
		},
		{
			name: "cannot create task in repository",
			fields: fields{
				setupTaskRepository: func(m *mocks.Repository) {
					m.On("CreateTask", mock.Anything, mock.Anything).Return(errors.New("some error"))
				},
			},
			args: args{
				ctx: context.Background(),
				options: &taskdmn.CreateTaskOptions{
					Title: "test task",
				},
			},
			wantResult: assert.Nil,
			wantError:  assert.Error,
		},
		{
			name: "task created",
			fields: fields{
				setupTaskRepository: func(m *mocks.Repository) {
					m.On("CreateTask", mock.Anything, mock.Anything).Return(nil)
				},
			},
			args: args{
				ctx: context.Background(),
				options: &taskdmn.CreateTaskOptions{
					Title: "test task",
				},
			},
			wantResult: assert.NotNil,
			wantError:  assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockTaskRepository := mocks.NewRepository(t)
			tt.fields.setupTaskRepository(mockTaskRepository)

			service := tasksrv.NewService(mockTaskRepository)
			res, err := service.CreateTask(tt.args.ctx, tt.args.options)

			tt.wantResult(t, res)
			tt.wantError(t, err)
			mockTaskRepository.AssertExpectations(t)
		})
	}
}
