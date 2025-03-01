// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/flyteorg/flytestdlib/storage"
)

// RawStore is an autogenerated mock type for the RawStore type
type RawStore struct {
	mock.Mock
}

type RawStore_CopyRaw struct {
	*mock.Call
}

func (_m RawStore_CopyRaw) Return(_a0 error) *RawStore_CopyRaw {
	return &RawStore_CopyRaw{Call: _m.Call.Return(_a0)}
}

func (_m *RawStore) OnCopyRaw(ctx context.Context, source storage.DataReference, destination storage.DataReference, opts storage.Options) *RawStore_CopyRaw {
	c := _m.On("CopyRaw", ctx, source, destination, opts)
	return &RawStore_CopyRaw{Call: c}
}

func (_m *RawStore) OnCopyRawMatch(matchers ...interface{}) *RawStore_CopyRaw {
	c := _m.On("CopyRaw", matchers...)
	return &RawStore_CopyRaw{Call: c}
}

// CopyRaw provides a mock function with given fields: ctx, source, destination, opts
func (_m *RawStore) CopyRaw(ctx context.Context, source storage.DataReference, destination storage.DataReference, opts storage.Options) error {
	ret := _m.Called(ctx, source, destination, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, storage.DataReference, storage.Options) error); ok {
		r0 = rf(ctx, source, destination, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type RawStore_CreateSignedURL struct {
	*mock.Call
}

func (_m RawStore_CreateSignedURL) Return(_a0 storage.SignedURLResponse, _a1 error) *RawStore_CreateSignedURL {
	return &RawStore_CreateSignedURL{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *RawStore) OnCreateSignedURL(ctx context.Context, reference storage.DataReference, properties storage.SignedURLProperties) *RawStore_CreateSignedURL {
	c := _m.On("CreateSignedURL", ctx, reference, properties)
	return &RawStore_CreateSignedURL{Call: c}
}

func (_m *RawStore) OnCreateSignedURLMatch(matchers ...interface{}) *RawStore_CreateSignedURL {
	c := _m.On("CreateSignedURL", matchers...)
	return &RawStore_CreateSignedURL{Call: c}
}

// CreateSignedURL provides a mock function with given fields: ctx, reference, properties
func (_m *RawStore) CreateSignedURL(ctx context.Context, reference storage.DataReference, properties storage.SignedURLProperties) (storage.SignedURLResponse, error) {
	ret := _m.Called(ctx, reference, properties)

	var r0 storage.SignedURLResponse
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, storage.SignedURLProperties) storage.SignedURLResponse); ok {
		r0 = rf(ctx, reference, properties)
	} else {
		r0 = ret.Get(0).(storage.SignedURLResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, storage.DataReference, storage.SignedURLProperties) error); ok {
		r1 = rf(ctx, reference, properties)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type RawStore_GetBaseContainerFQN struct {
	*mock.Call
}

func (_m RawStore_GetBaseContainerFQN) Return(_a0 storage.DataReference) *RawStore_GetBaseContainerFQN {
	return &RawStore_GetBaseContainerFQN{Call: _m.Call.Return(_a0)}
}

func (_m *RawStore) OnGetBaseContainerFQN(ctx context.Context) *RawStore_GetBaseContainerFQN {
	c := _m.On("GetBaseContainerFQN", ctx)
	return &RawStore_GetBaseContainerFQN{Call: c}
}

func (_m *RawStore) OnGetBaseContainerFQNMatch(matchers ...interface{}) *RawStore_GetBaseContainerFQN {
	c := _m.On("GetBaseContainerFQN", matchers...)
	return &RawStore_GetBaseContainerFQN{Call: c}
}

// GetBaseContainerFQN provides a mock function with given fields: ctx
func (_m *RawStore) GetBaseContainerFQN(ctx context.Context) storage.DataReference {
	ret := _m.Called(ctx)

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func(context.Context) storage.DataReference); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

type RawStore_Head struct {
	*mock.Call
}

func (_m RawStore_Head) Return(_a0 storage.Metadata, _a1 error) *RawStore_Head {
	return &RawStore_Head{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *RawStore) OnHead(ctx context.Context, reference storage.DataReference) *RawStore_Head {
	c := _m.On("Head", ctx, reference)
	return &RawStore_Head{Call: c}
}

func (_m *RawStore) OnHeadMatch(matchers ...interface{}) *RawStore_Head {
	c := _m.On("Head", matchers...)
	return &RawStore_Head{Call: c}
}

// Head provides a mock function with given fields: ctx, reference
func (_m *RawStore) Head(ctx context.Context, reference storage.DataReference) (storage.Metadata, error) {
	ret := _m.Called(ctx, reference)

	var r0 storage.Metadata
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference) storage.Metadata); ok {
		r0 = rf(ctx, reference)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.Metadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, storage.DataReference) error); ok {
		r1 = rf(ctx, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type RawStore_ReadRaw struct {
	*mock.Call
}

func (_m RawStore_ReadRaw) Return(_a0 io.ReadCloser, _a1 error) *RawStore_ReadRaw {
	return &RawStore_ReadRaw{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *RawStore) OnReadRaw(ctx context.Context, reference storage.DataReference) *RawStore_ReadRaw {
	c := _m.On("ReadRaw", ctx, reference)
	return &RawStore_ReadRaw{Call: c}
}

func (_m *RawStore) OnReadRawMatch(matchers ...interface{}) *RawStore_ReadRaw {
	c := _m.On("ReadRaw", matchers...)
	return &RawStore_ReadRaw{Call: c}
}

// ReadRaw provides a mock function with given fields: ctx, reference
func (_m *RawStore) ReadRaw(ctx context.Context, reference storage.DataReference) (io.ReadCloser, error) {
	ret := _m.Called(ctx, reference)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference) io.ReadCloser); ok {
		r0 = rf(ctx, reference)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, storage.DataReference) error); ok {
		r1 = rf(ctx, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type RawStore_WriteRaw struct {
	*mock.Call
}

func (_m RawStore_WriteRaw) Return(_a0 error) *RawStore_WriteRaw {
	return &RawStore_WriteRaw{Call: _m.Call.Return(_a0)}
}

func (_m *RawStore) OnWriteRaw(ctx context.Context, reference storage.DataReference, size int64, opts storage.Options, raw io.Reader) *RawStore_WriteRaw {
	c := _m.On("WriteRaw", ctx, reference, size, opts, raw)
	return &RawStore_WriteRaw{Call: c}
}

func (_m *RawStore) OnWriteRawMatch(matchers ...interface{}) *RawStore_WriteRaw {
	c := _m.On("WriteRaw", matchers...)
	return &RawStore_WriteRaw{Call: c}
}

// WriteRaw provides a mock function with given fields: ctx, reference, size, opts, raw
func (_m *RawStore) WriteRaw(ctx context.Context, reference storage.DataReference, size int64, opts storage.Options, raw io.Reader) error {
	ret := _m.Called(ctx, reference, size, opts, raw)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, int64, storage.Options, io.Reader) error); ok {
		r0 = rf(ctx, reference, size, opts, raw)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
