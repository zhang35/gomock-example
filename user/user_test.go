package user_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhang35/gomock-example/doer/mocks"
	"github.com/zhang35/gomock-example/match"
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
	err := testUser.Use(123, "Hello GoMock")
	assert.Nil(t, err)

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return dummyError from the mocked call.
	dummyError := errors.New("dummy error")
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1)
	err = testUser.Use(123, "Hello GoMock")
	assert.Equal(t, dummyError, err)

	// Test with matcher
	// Expect Do to be called once with 123 and a string as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().
		DoSomething(123, match.OfType("string")).
		Return(nil).
		Times(1)
	err = testUser.Use(123, "Any string")
	assert.Nil(t, err)

	// Asserting call order
	callFirst := mockDoer.EXPECT().DoSomething(1, "first this")
	mockDoer.EXPECT().DoSomething(2, "then this").After(callFirst)
	mockDoer.EXPECT().DoSomething(2, "or this").After(callFirst)
	testUser.Use(1, "first this")
	testUser.Use(2, "then this")
	testUser.Use(2, "or this")

	// or use InOrder
	gomock.InOrder(
		mockDoer.EXPECT().DoSomething(1, "first this"),
		mockDoer.EXPECT().DoSomething(2, "then this"),
		mockDoer.EXPECT().DoSomething(3, "then this"),
		mockDoer.EXPECT().DoSomething(4, "finally this"),
	)
	testUser.Use(1, "first this")
	testUser.Use(2, "then this")
	testUser.Use(3, "then this")
	testUser.Use(4, "finally this")

	// Specifying mock actions
	// Call .Do whenever the call is matched.

	// % go test -v user/user_test.go
	// === RUN   TestUse
	// Called with x = 3 and y = hello
	// --- PASS: TestUse (0.00s)
	// PASS
	// ok      command-line-arguments  0.172s
	mockDoer.EXPECT().
		DoSomething(gomock.Any(), gomock.Any()).
		Return(nil).
		Do(func(x int, y string) {
			if x > len(y) {
				t.Fail()
			}
			fmt.Println("Called with x =", x, "and y =", y)
		})
	testUser.Use(3, "hello")
}
