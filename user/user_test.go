package user_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ravitri/gomock/mocks"
	"github.com/ravitri/gomock/user"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}

	// Expectt Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(1989, "Hello RBT").Return(nil).Times(1)

	testUser.Use()
}
