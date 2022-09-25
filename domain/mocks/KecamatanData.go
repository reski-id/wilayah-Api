// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "wilayah/domain"

	mock "github.com/stretchr/testify/mock"
)

// KecamatanData is an autogenerated mock type for the KecamatanData type
type KecamatanData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *KecamatanData) Delete(id int) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *KecamatanData) GetAll() ([]domain.Kecamatan, error) {
	ret := _m.Called()

	var r0 []domain.Kecamatan
	if rf, ok := ret.Get(0).(func() []domain.Kecamatan); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Kecamatan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertKecamatan provides a mock function with given fields: newKecamatan
func (_m *KecamatanData) InsertKecamatan(newKecamatan domain.Kecamatan) (int, error) {
	ret := _m.Called(newKecamatan)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.Kecamatan) int); ok {
		r0 = rf(newKecamatan)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Kecamatan) error); ok {
		r1 = rf(newKecamatan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, updatedData
func (_m *KecamatanData) Update(id int, updatedData domain.Kecamatan) (int, error) {
	ret := _m.Called(id, updatedData)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, domain.Kecamatan) int); ok {
		r0 = rf(id, updatedData)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, domain.Kecamatan) error); ok {
		r1 = rf(id, updatedData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewKecamatanData interface {
	mock.TestingT
	Cleanup(func())
}

// NewKecamatanData creates a new instance of KecamatanData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKecamatanData(t mockConstructorTestingTNewKecamatanData) *KecamatanData {
	mock := &KecamatanData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}