// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/44taka/golang-gin/domain"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: db, id, name
func (_m *UserRepository) Create(db *gorm.DB, id int, name string) (domain.Users, error) {
	ret := _m.Called(db, id, name)

	var r0 domain.Users
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, string) domain.Users); ok {
		r0 = rf(db, id, name)
	} else {
		r0 = ret.Get(0).(domain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, int, string) error); ok {
		r1 = rf(db, id, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: db, id
func (_m *UserRepository) Delete(db *gorm.DB, id int) (domain.Users, error) {
	ret := _m.Called(db, id)

	var r0 domain.Users
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) domain.Users); ok {
		r0 = rf(db, id)
	} else {
		r0 = ret.Get(0).(domain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, int) error); ok {
		r1 = rf(db, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: db
func (_m *UserRepository) FindAll(db *gorm.DB) ([]domain.Users, error) {
	ret := _m.Called(db)

	var r0 []domain.Users
	if rf, ok := ret.Get(0).(func(*gorm.DB) []domain.Users); ok {
		r0 = rf(db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB) error); ok {
		r1 = rf(db)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: db, id
func (_m *UserRepository) FindByID(db *gorm.DB, id int) (domain.Users, error) {
	ret := _m.Called(db, id)

	var r0 domain.Users
	if rf, ok := ret.Get(0).(func(*gorm.DB, int) domain.Users); ok {
		r0 = rf(db, id)
	} else {
		r0 = ret.Get(0).(domain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, int) error); ok {
		r1 = rf(db, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: db, id, password
func (_m *UserRepository) Login(db *gorm.DB, id int, password string) (domain.Users, error) {
	ret := _m.Called(db, id, password)

	var r0 domain.Users
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, string) domain.Users); ok {
		r0 = rf(db, id, password)
	} else {
		r0 = ret.Get(0).(domain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, int, string) error); ok {
		r1 = rf(db, id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: db, id, name
func (_m *UserRepository) Update(db *gorm.DB, id int, name string) (domain.Users, error) {
	ret := _m.Called(db, id, name)

	var r0 domain.Users
	if rf, ok := ret.Get(0).(func(*gorm.DB, int, string) domain.Users); ok {
		r0 = rf(db, id, name)
	} else {
		r0 = ret.Get(0).(domain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, int, string) error); ok {
		r1 = rf(db, id, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
