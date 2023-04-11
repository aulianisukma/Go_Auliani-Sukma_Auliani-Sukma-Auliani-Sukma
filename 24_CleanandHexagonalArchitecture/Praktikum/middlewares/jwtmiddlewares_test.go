package middlewares

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	JWT = NewJWTS()
	JWTM = NewJWTSMock(mock.Mock{})
	JWTMock = &IjwtSMock{Mock: mock.Mock{}}
)

func TestCreateJWTToken(t *testing.T) {
	token, err := JWT.CreateJWTToken(1, "aul@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestCreateJWTTokenMock_Success(t *testing.T) {
	JWTMock.Mock.On("CreateJWTToken").Return("qwe", nil)
	token, err := JWTMock.CreateJWTToken(1, "aul@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestCreateJWTTokenMock_Failure(t *testing.T) {
	JWTMock = &IjwtSMock{Mock: mock.Mock{}}
	JWTMock.Mock.On("CreateJWTToken").Return(nil, errors.New("error"))
	token, err := JWTMock.CreateJWTToken(1, "aul@gmail.com")
	assert.NotNil(t, err)
	assert.Equal(t, "", token)
}