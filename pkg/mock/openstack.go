// Code generated by MockGen. DO NOT EDIT.
// Source: openstack.go
//
// Generated by this command:
//
//	mockgen -source=openstack.go -destination=../../mock/openstack.go -package=mock
//
// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	flags "github.com/nscaledev/baski/pkg/util/flags"
	gophercloud "github.com/gophercloud/gophercloud"
	keypairs "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/keypairs"
	servers "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	images "github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	floatingips "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/floatingips"
	gomock "go.uber.org/mock/gomock"
)

// MockOpenStackClient is a mock of OpenStackClient interface.
type MockOpenStackClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStackClientMockRecorder
}

// MockOpenStackClientMockRecorder is the mock recorder for MockOpenStackClient.
type MockOpenStackClientMockRecorder struct {
	mock *MockOpenStackClient
}

// NewMockOpenStackClient creates a new mock instance.
func NewMockOpenStackClient(ctrl *gomock.Controller) *MockOpenStackClient {
	mock := &MockOpenStackClient{ctrl: ctrl}
	mock.recorder = &MockOpenStackClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStackClient) EXPECT() *MockOpenStackClientMockRecorder {
	return m.recorder
}

// Client mocks base method.
func (m *MockOpenStackClient) Client() (*gophercloud.ProviderClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(*gophercloud.ProviderClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Client indicates an expected call of Client.
func (mr *MockOpenStackClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockOpenStackClient)(nil).Client))
}

// MockOpenStackComputeClient is a mock of OpenStackComputeClient interface.
type MockOpenStackComputeClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStackComputeClientMockRecorder
}

// MockOpenStackComputeClientMockRecorder is the mock recorder for MockOpenStackComputeClient.
type MockOpenStackComputeClientMockRecorder struct {
	mock *MockOpenStackComputeClient
}

// NewMockOpenStackComputeClient creates a new mock instance.
func NewMockOpenStackComputeClient(ctrl *gomock.Controller) *MockOpenStackComputeClient {
	mock := &MockOpenStackComputeClient{ctrl: ctrl}
	mock.recorder = &MockOpenStackComputeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStackComputeClient) EXPECT() *MockOpenStackComputeClientMockRecorder {
	return m.recorder
}

// AttachIP mocks base method.
func (m *MockOpenStackComputeClient) AttachIP(serverID, fip string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachIP", serverID, fip)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachIP indicates an expected call of AttachIP.
func (mr *MockOpenStackComputeClientMockRecorder) AttachIP(serverID, fip any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachIP", reflect.TypeOf((*MockOpenStackComputeClient)(nil).AttachIP), serverID, fip)
}

// CreateKeypair mocks base method.
func (m *MockOpenStackComputeClient) CreateKeypair(keyNamePrefix string) (*keypairs.KeyPair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKeypair", keyNamePrefix)
	ret0, _ := ret[0].(*keypairs.KeyPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKeypair indicates an expected call of CreateKeypair.
func (mr *MockOpenStackComputeClientMockRecorder) CreateKeypair(keyNamePrefix any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKeypair", reflect.TypeOf((*MockOpenStackComputeClient)(nil).CreateKeypair), keyNamePrefix)
}

// CreateServer mocks base method.
func (m *MockOpenStackComputeClient) CreateServer(keypairName, flavor, networkID string, attachConfigDrive *bool, userData []byte, imageID string, securityGroups []string) (*servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServer", keypairName, flavor, networkID, attachConfigDrive, userData, imageID, securityGroups)
	ret0, _ := ret[0].(*servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServer indicates an expected call of CreateServer.
func (mr *MockOpenStackComputeClientMockRecorder) CreateServer(keypairName, flavor, networkID, attachConfigDrive, userData, imageID, securityGroups any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServer", reflect.TypeOf((*MockOpenStackComputeClient)(nil).CreateServer), keypairName, flavor, networkID, attachConfigDrive, userData, imageID, securityGroups)
}

// GetFlavorIDByName mocks base method.
func (m *MockOpenStackComputeClient) GetFlavorIDByName(name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlavorIDByName", name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlavorIDByName indicates an expected call of GetFlavorIDByName.
func (mr *MockOpenStackComputeClientMockRecorder) GetFlavorIDByName(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlavorIDByName", reflect.TypeOf((*MockOpenStackComputeClient)(nil).GetFlavorIDByName), name)
}

// GetServerStatus mocks base method.
func (m *MockOpenStackComputeClient) GetServerStatus(sid string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerStatus", sid)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerStatus indicates an expected call of GetServerStatus.
func (mr *MockOpenStackComputeClientMockRecorder) GetServerStatus(sid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerStatus", reflect.TypeOf((*MockOpenStackComputeClient)(nil).GetServerStatus), sid)
}

// RemoveKeypair mocks base method.
func (m *MockOpenStackComputeClient) RemoveKeypair(keyName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveKeypair", keyName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveKeypair indicates an expected call of RemoveKeypair.
func (mr *MockOpenStackComputeClientMockRecorder) RemoveKeypair(keyName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveKeypair", reflect.TypeOf((*MockOpenStackComputeClient)(nil).RemoveKeypair), keyName)
}

// RemoveServer mocks base method.
func (m *MockOpenStackComputeClient) RemoveServer(serverID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveServer", serverID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveServer indicates an expected call of RemoveServer.
func (mr *MockOpenStackComputeClientMockRecorder) RemoveServer(serverID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveServer", reflect.TypeOf((*MockOpenStackComputeClient)(nil).RemoveServer), serverID)
}

// MockOpenStackImageClient is a mock of OpenStackImageClient interface.
type MockOpenStackImageClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStackImageClientMockRecorder
}

// MockOpenStackImageClientMockRecorder is the mock recorder for MockOpenStackImageClient.
type MockOpenStackImageClientMockRecorder struct {
	mock *MockOpenStackImageClient
}

// NewMockOpenStackImageClient creates a new mock instance.
func NewMockOpenStackImageClient(ctrl *gomock.Controller) *MockOpenStackImageClient {
	mock := &MockOpenStackImageClient{ctrl: ctrl}
	mock.recorder = &MockOpenStackImageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStackImageClient) EXPECT() *MockOpenStackImageClientMockRecorder {
	return m.recorder
}

// ChangeImageVisibility mocks base method.
func (m *MockOpenStackImageClient) ChangeImageVisibility(imgID string, visibility images.ImageVisibility) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeImageVisibility", imgID, visibility)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeImageVisibility indicates an expected call of ChangeImageVisibility.
func (mr *MockOpenStackImageClientMockRecorder) ChangeImageVisibility(imgID, visibility any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeImageVisibility", reflect.TypeOf((*MockOpenStackImageClient)(nil).ChangeImageVisibility), imgID, visibility)
}

// FetchAllImages mocks base method.
func (m *MockOpenStackImageClient) FetchAllImages(wildcard string) ([]images.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllImages", wildcard)
	ret0, _ := ret[0].([]images.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllImages indicates an expected call of FetchAllImages.
func (mr *MockOpenStackImageClientMockRecorder) FetchAllImages(wildcard any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllImages", reflect.TypeOf((*MockOpenStackImageClient)(nil).FetchAllImages), wildcard)
}

// FetchImage mocks base method.
func (m *MockOpenStackImageClient) FetchImage(imgID string) (*images.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchImage", imgID)
	ret0, _ := ret[0].(*images.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchImage indicates an expected call of FetchImage.
func (mr *MockOpenStackImageClientMockRecorder) FetchImage(imgID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchImage", reflect.TypeOf((*MockOpenStackImageClient)(nil).FetchImage), imgID)
}

// ModifyImageMetadata mocks base method.
func (m *MockOpenStackImageClient) ModifyImageMetadata(imgID, key, value string, operation images.UpdateOp) (*images.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyImageMetadata", imgID, key, value, operation)
	ret0, _ := ret[0].(*images.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyImageMetadata indicates an expected call of ModifyImageMetadata.
func (mr *MockOpenStackImageClientMockRecorder) ModifyImageMetadata(imgID, key, value, operation any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyImageMetadata", reflect.TypeOf((*MockOpenStackImageClient)(nil).ModifyImageMetadata), imgID, key, value, operation)
}

// RemoveImage mocks base method.
func (m *MockOpenStackImageClient) RemoveImage(imgID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveImage", imgID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveImage indicates an expected call of RemoveImage.
func (mr *MockOpenStackImageClientMockRecorder) RemoveImage(imgID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveImage", reflect.TypeOf((*MockOpenStackImageClient)(nil).RemoveImage), imgID)
}

// TagImage mocks base method.
func (m *MockOpenStackImageClient) TagImage(properties map[string]any, imgID, value, tagName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagImage", properties, imgID, value, tagName)
	ret0, _ := ret[0].(error)
	return ret0
}

// TagImage indicates an expected call of TagImage.
func (mr *MockOpenStackImageClientMockRecorder) TagImage(properties, imgID, value, tagName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagImage", reflect.TypeOf((*MockOpenStackImageClient)(nil).TagImage), properties, imgID, value, tagName)
}

// MockOpenStackNetworkClient is a mock of OpenStackNetworkClient interface.
type MockOpenStackNetworkClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStackNetworkClientMockRecorder
}

// MockOpenStackNetworkClientMockRecorder is the mock recorder for MockOpenStackNetworkClient.
type MockOpenStackNetworkClientMockRecorder struct {
	mock *MockOpenStackNetworkClient
}

// NewMockOpenStackNetworkClient creates a new mock instance.
func NewMockOpenStackNetworkClient(ctrl *gomock.Controller) *MockOpenStackNetworkClient {
	mock := &MockOpenStackNetworkClient{ctrl: ctrl}
	mock.recorder = &MockOpenStackNetworkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStackNetworkClient) EXPECT() *MockOpenStackNetworkClientMockRecorder {
	return m.recorder
}

// GetFloatingIP mocks base method.
func (m *MockOpenStackNetworkClient) GetFloatingIP(networkName string) (*floatingips.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFloatingIP", networkName)
	ret0, _ := ret[0].(*floatingips.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFloatingIP indicates an expected call of GetFloatingIP.
func (mr *MockOpenStackNetworkClientMockRecorder) GetFloatingIP(networkName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloatingIP", reflect.TypeOf((*MockOpenStackNetworkClient)(nil).GetFloatingIP), networkName)
}

// RemoveFIP mocks base method.
func (m *MockOpenStackNetworkClient) RemoveFIP(fipID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFIP", fipID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFIP indicates an expected call of RemoveFIP.
func (mr *MockOpenStackNetworkClientMockRecorder) RemoveFIP(fipID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFIP", reflect.TypeOf((*MockOpenStackNetworkClient)(nil).RemoveFIP), fipID)
}

// MockOpenStackScannerInterface is a mock of OpenStackScannerInterface interface.
type MockOpenStackScannerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStackScannerInterfaceMockRecorder
}

// MockOpenStackScannerInterfaceMockRecorder is the mock recorder for MockOpenStackScannerInterface.
type MockOpenStackScannerInterfaceMockRecorder struct {
	mock *MockOpenStackScannerInterface
}

// NewMockOpenStackScannerInterface creates a new mock instance.
func NewMockOpenStackScannerInterface(ctrl *gomock.Controller) *MockOpenStackScannerInterface {
	mock := &MockOpenStackScannerInterface{ctrl: ctrl}
	mock.recorder = &MockOpenStackScannerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStackScannerInterface) EXPECT() *MockOpenStackScannerInterfaceMockRecorder {
	return m.recorder
}

// CheckResults mocks base method.
func (m *MockOpenStackScannerInterface) CheckResults() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckResults")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckResults indicates an expected call of CheckResults.
func (mr *MockOpenStackScannerInterfaceMockRecorder) CheckResults() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckResults", reflect.TypeOf((*MockOpenStackScannerInterface)(nil).CheckResults))
}

// FetchScanResults mocks base method.
func (m *MockOpenStackScannerInterface) FetchScanResults() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchScanResults")
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchScanResults indicates an expected call of FetchScanResults.
func (mr *MockOpenStackScannerInterfaceMockRecorder) FetchScanResults() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchScanResults", reflect.TypeOf((*MockOpenStackScannerInterface)(nil).FetchScanResults))
}

// RunScan mocks base method.
func (m *MockOpenStackScannerInterface) RunScan(o *flags.ScanOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunScan", o)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunScan indicates an expected call of RunScan.
func (mr *MockOpenStackScannerInterfaceMockRecorder) RunScan(o any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunScan", reflect.TypeOf((*MockOpenStackScannerInterface)(nil).RunScan), o)
}

// TagImage mocks base method.
func (m *MockOpenStackScannerInterface) TagImage() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagImage")
	ret0, _ := ret[0].(error)
	return ret0
}

// TagImage indicates an expected call of TagImage.
func (mr *MockOpenStackScannerInterfaceMockRecorder) TagImage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagImage", reflect.TypeOf((*MockOpenStackScannerInterface)(nil).TagImage))
}

// UploadResultsToS3 mocks base method.
func (m *MockOpenStackScannerInterface) UploadResultsToS3() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadResultsToS3")
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadResultsToS3 indicates an expected call of UploadResultsToS3.
func (mr *MockOpenStackScannerInterfaceMockRecorder) UploadResultsToS3() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadResultsToS3", reflect.TypeOf((*MockOpenStackScannerInterface)(nil).UploadResultsToS3))
}
