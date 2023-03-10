// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	biz "backend/internal/biz"
	context "context"

	kafka "github.com/segmentio/kafka-go"

	mock "github.com/stretchr/testify/mock"
)

// FileContentKafkaRepo is an autogenerated mock type for the FileContentKafkaRepo type
type FileContentKafkaRepo struct {
	mock.Mock
}

// CommitMessage provides a mock function with given fields: ctx, message
func (_m *FileContentKafkaRepo) CommitMessage(ctx context.Context, message *kafka.Message) error {
	ret := _m.Called(ctx, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *kafka.Message) error); ok {
		r0 = rf(ctx, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMessage provides a mock function with given fields: ctx
func (_m *FileContentKafkaRepo) GetMessage(ctx context.Context) (*kafka.Message, error) {
	ret := _m.Called(ctx)

	var r0 *kafka.Message
	if rf, ok := ret.Get(0).(func(context.Context) *kafka.Message); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafka.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublishFileContent provides a mock function with given fields: ctx, gitClone
func (_m *FileContentKafkaRepo) PublishFileContent(ctx context.Context, gitClone []*biz.FileContent) error {
	ret := _m.Called(ctx, gitClone)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*biz.FileContent) error); ok {
		r0 = rf(ctx, gitClone)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFileContentKafkaRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewFileContentKafkaRepo creates a new instance of FileContentKafkaRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFileContentKafkaRepo(t mockConstructorTestingTNewFileContentKafkaRepo) *FileContentKafkaRepo {
	mock := &FileContentKafkaRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
