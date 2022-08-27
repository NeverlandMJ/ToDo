package entity_test

import (
	"reflect"
	"testing"

	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
)

func TestNewUser(t *testing.T) {
	type args struct {
		username string
		password string
		phone    string
	}
	tests := []struct {
		name string
		args args
		want entity.User
	}{
		{
			name: "should pass",
			args: args{
				username: "sunbula",
				password: "1230",
				phone: "+123456789",
			},
			want: entity.User{
				UserName: "sunbula",
				Password: "1230",
				Phone: "+123456789",
				IsBlocked: false,
			},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := entity.NewUser(tt.args.username, tt.args.password, tt.args.phone)
			tt.want.ID = got.ID
			tt.want.CreatedAt = got.CreatedAt
			tt.want.UpdatedAt = got.UpdatedAt
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
