package entity

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

var testUserID = uuid.New()
var testTodoID = uuid.New()
var testTime = time.Now()

func TestNewTodo(t *testing.T) {
	type args struct {
		deadline time.Time
		body     string
		userID   uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    Todo
		wantErr bool
	}{
		{
			name: "should pass",
			args: args{
				deadline: time.Date(2022, time.September, 8, 7, 0, 0, 0, time.UTC),
				body:     "dad's birthday",
				userID:   testUserID,
			},
			want: Todo{
				ID:        testTodoID,
				UserID:    testUserID,
				Body:      "dad's birthday",
				CreatedAt: testTime,
				Deadline:  time.Date(2022, time.September, 8, 7, 0, 0, 0, time.UTC),
				IsDone:    false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTodo(tt.args.deadline, tt.args.body, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.ID = testTodoID
			got.CreatedAt = testTime
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodo() \ngot = %v\n want %v\n", got, tt.want)
			}
		})
	}
}
