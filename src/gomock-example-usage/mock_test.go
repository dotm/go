package greet

import (
	"gomock-example-usage/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMock(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock.NewMockIGreeter(ctrl)

	cannedResponse := "aoeuaeuaeu"
	mock.
		EXPECT().
		Greet(gomock.Any()).
		Return(cannedResponse).
		AnyTimes()

	//when
	result := functionWithMock(mock)

	//then
	if result != cannedResponse {
		t.Fail()
	}
}
