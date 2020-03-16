// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/abyanjksatu/goscription/models"
	mock "github.com/stretchr/testify/mock"
)

// DomainUsecase is an autogenerated mock type for the DomainUsecase type
type DomainUsecase struct {
	mock.Mock
}

// GetDomainAvailable provides a mock function with given fields: ctx, domain
func (_m *DomainUsecase) GetDomainAvailable(ctx context.Context, domain string) (models.DomainAvailableResponse, error) {
	ret := _m.Called(ctx, domain)

	var r0 models.DomainAvailableResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) models.DomainAvailableResponse); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(models.DomainAvailableResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
