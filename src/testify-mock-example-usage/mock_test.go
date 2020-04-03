package greet

import (
	"testify-mock-example-usage/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestMock(t *testing.T) {
	//given
	mockGreeter := &mocks.IGreeter{}

	cannedResponse := "aoeuaeuaeu"
	mockGreeter.On("Greet", mock.Anything).Return(cannedResponse).
		Once() //expect to be called once

	//when
	result := functionWithMock(mockGreeter)

	//then
	if result != cannedResponse {
		t.Fail()
	}
	mockGreeter.AssertExpectations(t)
}
