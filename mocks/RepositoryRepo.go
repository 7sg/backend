// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	biz "backend-GuardRails/internal/biz"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RepositoryRepo is an autogenerated mock type for the RepositoryRepo type
type RepositoryRepo struct {
	mock.Mock
}

// CreateRepository provides a mock function with given fields: ctx, repo
func (_m *RepositoryRepo) CreateRepository(ctx context.Context, repo *biz.Repository) (*biz.Repository, error) {
	ret := _m.Called(ctx, repo)

	var r0 *biz.Repository
	if rf, ok := ret.Get(0).(func(context.Context, *biz.Repository) *biz.Repository); ok {
		r0 = rf(ctx, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*biz.Repository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *biz.Repository) error); ok {
		r1 = rf(ctx, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepository provides a mock function with given fields: ctx, id
func (_m *RepositoryRepo) GetRepository(ctx context.Context, id uint64) (*biz.Repository, error) {
	ret := _m.Called(ctx, id)

	var r0 *biz.Repository
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *biz.Repository); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*biz.Repository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepositoryRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositoryRepo creates a new instance of RepositoryRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryRepo(t mockConstructorTestingTNewRepositoryRepo) *RepositoryRepo {
	mock := &RepositoryRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
