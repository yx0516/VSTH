package mocks

import (
	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
)

type KeysApiMock struct {
	mockedGetMethodInvocations map[string]*client.Response
}

func NewKeysApiMock() *KeysApiMock {
	return &KeysApiMock{mockedGetMethodInvocations:make(map[string]*client.Response)}
}

func (mock *KeysApiMock) Get(ctx context.Context, key string, opts *client.GetOptions) (*client.Response, error) {
	return mock.mockedGetMethodInvocations[key], nil
}

func (mock *KeysApiMock) Set(ctx context.Context, key, value string, opts *client.SetOptions) (*client.Response, error) {
	return nil, nil
}

func (mock *KeysApiMock) Delete(ctx context.Context, key string, opts *client.DeleteOptions) (*client.Response, error) {
	return nil, nil
}

func (mock *KeysApiMock) Create(ctx context.Context, key string, value string) (*client.Response, error) {
	return nil, nil
}

func (mock *KeysApiMock) CreateInOrder(ctx context.Context, dir, value string, opts *client.CreateInOrderOptions) (*client.Response, error) {
	return nil, nil
}

func (mock *KeysApiMock) Update(ctx context.Context, key string, value string) (*client.Response, error) {
	return nil, nil
}

func (mock *KeysApiMock) Watcher(key string, opts *client.WatcherOptions) client.Watcher {
	return nil
}

func (mock *KeysApiMock) MockGet(key string, response *client.Response) {
	mock.mockedGetMethodInvocations[key] = response
}
