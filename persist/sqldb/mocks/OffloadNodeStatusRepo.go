// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

import v1alpha1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"

// OffloadNodeStatusRepo is an autogenerated mock type for the OffloadNodeStatusRepo type
type OffloadNodeStatusRepo struct {
	mock.Mock
}

// Delete provides a mock function with given fields: name, namespace
func (_m *OffloadNodeStatusRepo) Delete(name string, namespace string) error {
	ret := _m.Called(name, namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: name, namespace
func (_m *OffloadNodeStatusRepo) Get(name string, namespace string) (*v1alpha1.Workflow, error) {
	ret := _m.Called(name, namespace)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(string, string) *v1alpha1.Workflow); ok {
		r0 = rf(name, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsEnabled provides a mock function with given fields:
func (_m *OffloadNodeStatusRepo) IsEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// List provides a mock function with given fields: namespace
func (_m *OffloadNodeStatusRepo) List(namespace string) (v1alpha1.Workflows, error) {
	ret := _m.Called(namespace)

	var r0 v1alpha1.Workflows
	if rf, ok := ret.Get(0).(func(string) v1alpha1.Workflows); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.Workflows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: wf
func (_m *OffloadNodeStatusRepo) Save(wf *v1alpha1.Workflow) error {
	ret := _m.Called(wf)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Workflow) error); ok {
		r0 = rf(wf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
