package controllers

import (
	"testing"

	"github.com/labstack/echo"
)

func TestGetUsersController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetUsersController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetUsersController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoginUserController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoginUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("LoginUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateUserController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUserController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateUserController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
