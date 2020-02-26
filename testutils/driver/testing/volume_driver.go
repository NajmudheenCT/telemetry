// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/sodafoundation/telemetry/pkg/model"
import proto "github.com/sodafoundation/telemetry/pkg/model/proto"

// VolumeDriver is an autogenerated mock type for the VolumeDriver type
type VolumeDriver struct {
	mock.Mock
}

// CreateSnapshot provides a mock function with given fields: opt
func (_m *VolumeDriver) CreateSnapshot(opt *proto.CreateVolumeSnapshotOpts) (*model.VolumeSnapshotSpec, error) {
	ret := _m.Called(opt)

	var r0 *model.VolumeSnapshotSpec
	if rf, ok := ret.Get(0).(func(*proto.CreateVolumeSnapshotOpts) *model.VolumeSnapshotSpec); ok {
		r0 = rf(opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeSnapshotSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*proto.CreateVolumeSnapshotOpts) error); ok {
		r1 = rf(opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolume provides a mock function with given fields: opt
func (_m *VolumeDriver) CreateVolume(opt *proto.CreateVolumeOpts) (*model.VolumeSpec, error) {
	ret := _m.Called(opt)

	var r0 *model.VolumeSpec
	if rf, ok := ret.Get(0).(func(*proto.CreateVolumeOpts) *model.VolumeSpec); ok {
		r0 = rf(opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*proto.CreateVolumeOpts) error); ok {
		r1 = rf(opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolumeGroup provides a mock function with given fields: opt
func (_m *VolumeDriver) CreateVolumeGroup(opt *proto.CreateVolumeGroupOpts) (*model.VolumeGroupSpec, error) {
	ret := _m.Called(opt)

	var r0 *model.VolumeGroupSpec
	if rf, ok := ret.Get(0).(func(*proto.CreateVolumeGroupOpts) *model.VolumeGroupSpec); ok {
		r0 = rf(opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeGroupSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*proto.CreateVolumeGroupOpts) error); ok {
		r1 = rf(opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSnapshot provides a mock function with given fields: opt
func (_m *VolumeDriver) DeleteSnapshot(opt *proto.DeleteVolumeSnapshotOpts) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*proto.DeleteVolumeSnapshotOpts) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteVolume provides a mock function with given fields: opt
func (_m *VolumeDriver) DeleteVolume(opt *proto.DeleteVolumeOpts) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*proto.DeleteVolumeOpts) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteVolumeGroup provides a mock function with given fields: opt
func (_m *VolumeDriver) DeleteVolumeGroup(opt *proto.DeleteVolumeGroupOpts) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*proto.DeleteVolumeGroupOpts) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtendVolume provides a mock function with given fields: opt
func (_m *VolumeDriver) ExtendVolume(opt *proto.ExtendVolumeOpts) (*model.VolumeSpec, error) {
	ret := _m.Called(opt)

	var r0 *model.VolumeSpec
	if rf, ok := ret.Get(0).(func(*proto.ExtendVolumeOpts) *model.VolumeSpec); ok {
		r0 = rf(opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*proto.ExtendVolumeOpts) error); ok {
		r1 = rf(opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitializeConnection provides a mock function with given fields: opt
func (_m *VolumeDriver) InitializeConnection(opt *proto.CreateVolumeAttachmentOpts) (*model.ConnectionInfo, error) {
	ret := _m.Called(opt)

	var r0 *model.ConnectionInfo
	if rf, ok := ret.Get(0).(func(*proto.CreateVolumeAttachmentOpts) *model.ConnectionInfo); ok {
		r0 = rf(opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ConnectionInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*proto.CreateVolumeAttachmentOpts) error); ok {
		r1 = rf(opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitializeSnapshotConnection provides a mock function with given fields: opt
func (_m *VolumeDriver) InitializeSnapshotConnection(opt *proto.CreateSnapshotAttachmentOpts) (*model.ConnectionInfo, error) {
	ret := _m.Called(opt)

	var r0 *model.ConnectionInfo
	if rf, ok := ret.Get(0).(func(*proto.CreateSnapshotAttachmentOpts) *model.ConnectionInfo); ok {
		r0 = rf(opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ConnectionInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*proto.CreateSnapshotAttachmentOpts) error); ok {
		r1 = rf(opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPools provides a mock function with given fields:
func (_m *VolumeDriver) ListPools() ([]*model.StoragePoolSpec, error) {
	ret := _m.Called()

	var r0 []*model.StoragePoolSpec
	if rf, ok := ret.Get(0).(func() []*model.StoragePoolSpec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.StoragePoolSpec)
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

// PullSnapshot provides a mock function with given fields: snapIdentifier
func (_m *VolumeDriver) PullSnapshot(snapIdentifier string) (*model.VolumeSnapshotSpec, error) {
	ret := _m.Called(snapIdentifier)

	var r0 *model.VolumeSnapshotSpec
	if rf, ok := ret.Get(0).(func(string) *model.VolumeSnapshotSpec); ok {
		r0 = rf(snapIdentifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeSnapshotSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(snapIdentifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PullVolume provides a mock function with given fields: volIdentifier
func (_m *VolumeDriver) PullVolume(volIdentifier string) (*model.VolumeSpec, error) {
	ret := _m.Called(volIdentifier)

	var r0 *model.VolumeSpec
	if rf, ok := ret.Get(0).(func(string) *model.VolumeSpec); ok {
		r0 = rf(volIdentifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(volIdentifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Setup provides a mock function with given fields:
func (_m *VolumeDriver) Setup() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TerminateConnection provides a mock function with given fields: opt
func (_m *VolumeDriver) TerminateConnection(opt *proto.DeleteVolumeAttachmentOpts) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*proto.DeleteVolumeAttachmentOpts) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TerminateSnapshotConnection provides a mock function with given fields: opt
func (_m *VolumeDriver) TerminateSnapshotConnection(opt *proto.DeleteSnapshotAttachmentOpts) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*proto.DeleteSnapshotAttachmentOpts) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unset provides a mock function with given fields:
func (_m *VolumeDriver) Unset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateVolumeGroup provides a mock function with given fields: opt
func (_m *VolumeDriver) UpdateVolumeGroup(opt *proto.UpdateVolumeGroupOpts) (*model.VolumeGroupSpec, error) {
	ret := _m.Called(opt)

	var r0 *model.VolumeGroupSpec
	if rf, ok := ret.Get(0).(func(*proto.UpdateVolumeGroupOpts) *model.VolumeGroupSpec); ok {
		r0 = rf(opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VolumeGroupSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*proto.UpdateVolumeGroupOpts) error); ok {
		r1 = rf(opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
