package user_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhang35/gomock-example/doer/mocks"
	"github.com/zhang35/gomock-example/user"
	"go.uber.org/mock/gomock"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := user.User{
		Doer: mockDoer,
	}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	// If it's called with different parameters or called more than once, the test will fail.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	err := testUser.Use()
	assert.Nil(t, err, "Expected error is nil")
}

func TestUseReturnsErrorFromDo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dummyError := errors.New("dummy error")
	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return dummyError from the mocked call.
	// If it's called with different parameters or called more than once, the test will fail.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1)

	err := testUser.Use()
	assert.Equal(t, dummyError, err, "Expected error does not match the actual error")
}
