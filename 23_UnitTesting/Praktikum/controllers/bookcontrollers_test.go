package controllers

import (
	"testing"

	"github.com/labstack/echo"
)

func TestGetBooksController(t *testing.T) {
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
			if err := GetBooksController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetBooksController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetBookController(t *testing.T) {
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
			if err := GetBookController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetBookController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateBookController(t *testing.T) {
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
			if err := CreateBookController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateBookController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteBookController(t *testing.T) {
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
			if err := DeleteBookController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBookController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateBookController(t *testing.T) {
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
			if err := UpdateBookController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBookController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
