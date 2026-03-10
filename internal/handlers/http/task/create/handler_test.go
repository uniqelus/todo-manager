package create_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/uniqelus/todo-manager/internal/domains/server/task"
	"github.com/uniqelus/todo-manager/internal/handlers/http/helpers"
	"github.com/uniqelus/todo-manager/internal/handlers/http/task/create"
	"github.com/uniqelus/todo-manager/internal/handlers/http/task/create/mocks"
)

func TestCreateTaskHandler(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		requestBody      string
		taskCreatorSetup func(m *mocks.TaskCreator)
		wantStatus       int
		wantResponse     create.Response
	}{
		{
			name:             "invalid request - empty body",
			requestBody:      "",
			taskCreatorSetup: func(_ *mocks.TaskCreator) {},
			wantStatus:       http.StatusBadRequest,
			wantResponse:     create.Response{Error: helpers.ErrEmptyRequest.Error()},
		},
		{
			name:             "invalid request - failed to decode request body",
			requestBody:      "{",
			taskCreatorSetup: func(_ *mocks.TaskCreator) {},
			wantStatus:       http.StatusBadRequest,
			wantResponse:     create.Response{Error: helpers.ErrFailedToDecodeRequset.Error()},
		},
		{
			name:        "invalid request - failed to decode request body",
			requestBody: `{"title": "test task"}`,
			taskCreatorSetup: func(m *mocks.TaskCreator) {
				m.On("CreateTask", mock.Anything, &task.CreateTaskOptions{
					Title:       "test task",
					Description: "",
					DueDate:     time.Time{},
				}).Return(nil, errors.New("internal"))
			},
			wantStatus:   http.StatusInternalServerError,
			wantResponse: create.Response{Error: create.ErrFailedToCreateTask.Error()},
		},
		{
			name:        "invalid request - failed to decode request body",
			requestBody: `{"title": "test task"}`,
			taskCreatorSetup: func(m *mocks.TaskCreator) {
				m.On("CreateTask", mock.Anything, &task.CreateTaskOptions{
					Title:       "test task",
					Description: "",
					DueDate:     time.Time{},
				}).Return(&task.Task{
					Title: "test task",
				}, nil)
			},
			wantStatus:   http.StatusCreated,
			wantResponse: create.Response{Task: task.Task{Title: "test task"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockTaskCreator := mocks.NewTaskCreator(t)
			tt.taskCreatorSetup(mockTaskCreator)

			req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")

			recorder := httptest.NewRecorder()

			handler := create.NewHandler(zap.NewNop(), mockTaskCreator)
			handler.ServeHTTP(recorder, req)

			var actualResponse create.Response
			_ = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)

			assert.Equal(t, tt.wantStatus, recorder.Code)
			assert.Equal(t, tt.wantResponse, actualResponse)
			mockTaskCreator.AssertExpectations(t)
		})
	}
}
