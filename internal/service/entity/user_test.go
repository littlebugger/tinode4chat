package entity

import (
	"testing"
)

func TestUserValid(t *testing.T) {
	tests := []struct {
		name string
		user User
		want bool
	}{
		{
			name: "ValidUser",
			user: User{Email: "testemail@gmail.com", Username: "testuser", Password: "testpassword"},
			want: true,
		},
		{
			name: "EmptyEmail",
			user: User{Email: "", Username: "testuser", Password: "testpassword"},
			want: false,
		},
		{
			name: "EmptyUsername",
			user: User{Email: "testemail@gmail.com", Username: "", Password: "testpassword"},
			want: false,
		},
		{
			name: "EmptyPassword",
			user: User{Email: "testemail@gmail.com", Username: "testuser", Password: ""},
			want: false,
		},
		{
			name: "AllFieldsEmpty",
			user: User{Email: "", Username: "", Password: ""},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.user.Valid(); got != tc.want {
				t.Errorf("User.Valid() = %v, want %v", got, tc.want)
			}
		})
	}
}
