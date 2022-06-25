package mocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ApplicationService_Build(t *testing.T) {

	// 	"github.com/golang/mock/gomock"
	/*
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		MockApplicationService := NewMockApplication(ctrl)
		MockApplicationService.EXPECT().Build().Return(nil).Times(1)
	*/
	t.Run("initial test", func(t *testing.T) {
		assert.Equal(t, 1, 1)
	})

}
