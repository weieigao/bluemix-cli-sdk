// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/IBM-Bluemix/bluemix-cli-sdk/plugin/models"
	"github.com/weieigao/bluemix-cli-sdk/plugin"
)

type FakePluginContext struct {
	APIVersionStub        func() string
	aPIVersionMutex       sync.RWMutex
	aPIVersionArgsForCall []struct{}
	aPIVersionReturns struct {
		result1 string
	}
	APIEndpointStub        func() string
	aPIEndpointMutex       sync.RWMutex
	aPIEndpointArgsForCall []struct{}
	aPIEndpointReturns struct {
		result1 string
	}
	HasAPIEndpointStub        func() bool
	hasAPIEndpointMutex       sync.RWMutex
	hasAPIEndpointArgsForCall []struct{}
	hasAPIEndpointReturns struct {
		result1 bool
	}
	AuthenticationEndpointStub        func() string
	authenticationEndpointMutex       sync.RWMutex
	authenticationEndpointArgsForCall []struct{}
	authenticationEndpointReturns struct {
		result1 string
	}
	LoggregatorEndpointStub        func() string
	loggregatorEndpointMutex       sync.RWMutex
	loggregatorEndpointArgsForCall []struct{}
	loggregatorEndpointReturns struct {
		result1 string
	}
	DopplerEndpointStub        func() string
	dopplerEndpointMutex       sync.RWMutex
	dopplerEndpointArgsForCall []struct{}
	dopplerEndpointReturns struct {
		result1 string
	}
	UaaEndpointStub        func() string
	uaaEndpointMutex       sync.RWMutex
	uaaEndpointArgsForCall []struct{}
	uaaEndpointReturns struct {
		result1 string
	}
	UsernameStub        func() string
	usernameMutex       sync.RWMutex
	usernameArgsForCall []struct{}
	usernameReturns struct {
		result1 string
	}
	UserGUIDStub        func() string
	userGUIDMutex       sync.RWMutex
	userGUIDArgsForCall []struct{}
	userGUIDReturns struct {
		result1 string
	}
	UserEmailStub        func() string
	userEmailMutex       sync.RWMutex
	userEmailArgsForCall []struct{}
	userEmailReturns struct {
		result1 string
	}
	IsLoggedInStub        func() bool
	isLoggedInMutex       sync.RWMutex
	isLoggedInArgsForCall []struct{}
	isLoggedInReturns struct {
		result1 bool
	}
	AccessTokenStub        func() string
	accessTokenMutex       sync.RWMutex
	accessTokenArgsForCall []struct{}
	accessTokenReturns struct {
		result1 string
	}
	TokenRefreshStub        func() (string, error)
	tokenRefreshMutex       sync.RWMutex
	tokenRefreshArgsForCall []struct{}
	tokenRefreshReturns struct {
		result1 string
		result2 error
	}
	CurrentOrgStub        func() models.Organization
	currentOrgMutex       sync.RWMutex
	currentOrgArgsForCall []struct{}
	currentOrgReturns struct {
		result1 models.Organization
	}
	HasOrganizationStub        func() bool
	hasOrganizationMutex       sync.RWMutex
	hasOrganizationArgsForCall []struct{}
	hasOrganizationReturns struct {
		result1 bool
	}
	CurrentSpaceStub        func() models.Space
	currentSpaceMutex       sync.RWMutex
	currentSpaceArgsForCall []struct{}
	currentSpaceReturns struct {
		result1 models.Space
	}
	HasSpaceStub        func() bool
	hasSpaceMutex       sync.RWMutex
	hasSpaceArgsForCall []struct{}
	hasSpaceReturns struct {
		result1 bool
	}
	LocaleStub        func() string
	localeMutex       sync.RWMutex
	localeArgsForCall []struct{}
	localeReturns struct {
		result1 string
	}
	TraceStub        func() string
	traceMutex       sync.RWMutex
	traceArgsForCall []struct{}
	traceReturns struct {
		result1 string
	}
	ColorEnabledStub        func() string
	colorEnabledMutex       sync.RWMutex
	colorEnabledArgsForCall []struct{}
	colorEnabledReturns struct {
		result1 string
	}
	IsSSLDisabledStub        func() bool
	isSSLDisabledMutex       sync.RWMutex
	isSSLDisabledArgsForCall []struct{}
	isSSLDisabledReturns struct {
		result1 bool
	}
	PluginDirectoryStub        func() string
	pluginDirectoryMutex       sync.RWMutex
	pluginDirectoryArgsForCall []struct{}
	pluginDirectoryReturns struct {
		result1 string
	}
	HTTPTimeoutStub        func() int
	hTTPTimeoutMutex       sync.RWMutex
	hTTPTimeoutArgsForCall []struct{}
	hTTPTimeoutReturns struct {
		result1 int
	}
	PluginConfigStub        func() plugin.PluginConfig
	pluginConfigMutex       sync.RWMutex
	pluginConfigArgsForCall []struct{}
	pluginConfigReturns struct {
		result1 plugin.PluginConfig
	}
	CommandNamespaceStub        func() string
	commandNamespaceMutex       sync.RWMutex
	commandNamespaceArgsForCall []struct{}
	commandNamespaceReturns struct {
		result1 string
	}
}

func (fake *FakePluginContext) APIVersion() string {
	fake.aPIVersionMutex.Lock()
	fake.aPIVersionArgsForCall = append(fake.aPIVersionArgsForCall, struct{}{})
	fake.aPIVersionMutex.Unlock()
	if fake.APIVersionStub != nil {
		return fake.APIVersionStub()
	} else {
		return fake.aPIVersionReturns.result1
	}
}

func (fake *FakePluginContext) APIVersionCallCount() int {
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	return len(fake.aPIVersionArgsForCall)
}

func (fake *FakePluginContext) APIVersionReturns(result1 string) {
	fake.APIVersionStub = nil
	fake.aPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) APIEndpoint() string {
	fake.aPIEndpointMutex.Lock()
	fake.aPIEndpointArgsForCall = append(fake.aPIEndpointArgsForCall, struct{}{})
	fake.aPIEndpointMutex.Unlock()
	if fake.APIEndpointStub != nil {
		return fake.APIEndpointStub()
	} else {
		return fake.aPIEndpointReturns.result1
	}
}

func (fake *FakePluginContext) APIEndpointCallCount() int {
	fake.aPIEndpointMutex.RLock()
	defer fake.aPIEndpointMutex.RUnlock()
	return len(fake.aPIEndpointArgsForCall)
}

func (fake *FakePluginContext) APIEndpointReturns(result1 string) {
	fake.APIEndpointStub = nil
	fake.aPIEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) HasAPIEndpoint() bool {
	fake.hasAPIEndpointMutex.Lock()
	fake.hasAPIEndpointArgsForCall = append(fake.hasAPIEndpointArgsForCall, struct{}{})
	fake.hasAPIEndpointMutex.Unlock()
	if fake.HasAPIEndpointStub != nil {
		return fake.HasAPIEndpointStub()
	} else {
		return fake.hasAPIEndpointReturns.result1
	}
}

func (fake *FakePluginContext) HasAPIEndpointCallCount() int {
	fake.hasAPIEndpointMutex.RLock()
	defer fake.hasAPIEndpointMutex.RUnlock()
	return len(fake.hasAPIEndpointArgsForCall)
}

func (fake *FakePluginContext) HasAPIEndpointReturns(result1 bool) {
	fake.HasAPIEndpointStub = nil
	fake.hasAPIEndpointReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePluginContext) AuthenticationEndpoint() string {
	fake.authenticationEndpointMutex.Lock()
	fake.authenticationEndpointArgsForCall = append(fake.authenticationEndpointArgsForCall, struct{}{})
	fake.authenticationEndpointMutex.Unlock()
	if fake.AuthenticationEndpointStub != nil {
		return fake.AuthenticationEndpointStub()
	} else {
		return fake.authenticationEndpointReturns.result1
	}
}

func (fake *FakePluginContext) AuthenticationEndpointCallCount() int {
	fake.authenticationEndpointMutex.RLock()
	defer fake.authenticationEndpointMutex.RUnlock()
	return len(fake.authenticationEndpointArgsForCall)
}

func (fake *FakePluginContext) AuthenticationEndpointReturns(result1 string) {
	fake.AuthenticationEndpointStub = nil
	fake.authenticationEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) LoggregatorEndpoint() string {
	fake.loggregatorEndpointMutex.Lock()
	fake.loggregatorEndpointArgsForCall = append(fake.loggregatorEndpointArgsForCall, struct{}{})
	fake.loggregatorEndpointMutex.Unlock()
	if fake.LoggregatorEndpointStub != nil {
		return fake.LoggregatorEndpointStub()
	} else {
		return fake.loggregatorEndpointReturns.result1
	}
}

func (fake *FakePluginContext) LoggregatorEndpointCallCount() int {
	fake.loggregatorEndpointMutex.RLock()
	defer fake.loggregatorEndpointMutex.RUnlock()
	return len(fake.loggregatorEndpointArgsForCall)
}

func (fake *FakePluginContext) LoggregatorEndpointReturns(result1 string) {
	fake.LoggregatorEndpointStub = nil
	fake.loggregatorEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) DopplerEndpoint() string {
	fake.dopplerEndpointMutex.Lock()
	fake.dopplerEndpointArgsForCall = append(fake.dopplerEndpointArgsForCall, struct{}{})
	fake.dopplerEndpointMutex.Unlock()
	if fake.DopplerEndpointStub != nil {
		return fake.DopplerEndpointStub()
	} else {
		return fake.dopplerEndpointReturns.result1
	}
}

func (fake *FakePluginContext) DopplerEndpointCallCount() int {
	fake.dopplerEndpointMutex.RLock()
	defer fake.dopplerEndpointMutex.RUnlock()
	return len(fake.dopplerEndpointArgsForCall)
}

func (fake *FakePluginContext) DopplerEndpointReturns(result1 string) {
	fake.DopplerEndpointStub = nil
	fake.dopplerEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) UaaEndpoint() string {
	fake.uaaEndpointMutex.Lock()
	fake.uaaEndpointArgsForCall = append(fake.uaaEndpointArgsForCall, struct{}{})
	fake.uaaEndpointMutex.Unlock()
	if fake.UaaEndpointStub != nil {
		return fake.UaaEndpointStub()
	} else {
		return fake.uaaEndpointReturns.result1
	}
}

func (fake *FakePluginContext) UaaEndpointCallCount() int {
	fake.uaaEndpointMutex.RLock()
	defer fake.uaaEndpointMutex.RUnlock()
	return len(fake.uaaEndpointArgsForCall)
}

func (fake *FakePluginContext) UaaEndpointReturns(result1 string) {
	fake.UaaEndpointStub = nil
	fake.uaaEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) Username() string {
	fake.usernameMutex.Lock()
	fake.usernameArgsForCall = append(fake.usernameArgsForCall, struct{}{})
	fake.usernameMutex.Unlock()
	if fake.UsernameStub != nil {
		return fake.UsernameStub()
	} else {
		return fake.usernameReturns.result1
	}
}

func (fake *FakePluginContext) UsernameCallCount() int {
	fake.usernameMutex.RLock()
	defer fake.usernameMutex.RUnlock()
	return len(fake.usernameArgsForCall)
}

func (fake *FakePluginContext) UsernameReturns(result1 string) {
	fake.UsernameStub = nil
	fake.usernameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) UserGUID() string {
	fake.userGUIDMutex.Lock()
	fake.userGUIDArgsForCall = append(fake.userGUIDArgsForCall, struct{}{})
	fake.userGUIDMutex.Unlock()
	if fake.UserGUIDStub != nil {
		return fake.UserGUIDStub()
	} else {
		return fake.userGUIDReturns.result1
	}
}

func (fake *FakePluginContext) UserGUIDCallCount() int {
	fake.userGUIDMutex.RLock()
	defer fake.userGUIDMutex.RUnlock()
	return len(fake.userGUIDArgsForCall)
}

func (fake *FakePluginContext) UserGUIDReturns(result1 string) {
	fake.UserGUIDStub = nil
	fake.userGUIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) UserEmail() string {
	fake.userEmailMutex.Lock()
	fake.userEmailArgsForCall = append(fake.userEmailArgsForCall, struct{}{})
	fake.userEmailMutex.Unlock()
	if fake.UserEmailStub != nil {
		return fake.UserEmailStub()
	} else {
		return fake.userEmailReturns.result1
	}
}

func (fake *FakePluginContext) UserEmailCallCount() int {
	fake.userEmailMutex.RLock()
	defer fake.userEmailMutex.RUnlock()
	return len(fake.userEmailArgsForCall)
}

func (fake *FakePluginContext) UserEmailReturns(result1 string) {
	fake.UserEmailStub = nil
	fake.userEmailReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) IsLoggedIn() bool {
	fake.isLoggedInMutex.Lock()
	fake.isLoggedInArgsForCall = append(fake.isLoggedInArgsForCall, struct{}{})
	fake.isLoggedInMutex.Unlock()
	if fake.IsLoggedInStub != nil {
		return fake.IsLoggedInStub()
	} else {
		return fake.isLoggedInReturns.result1
	}
}

func (fake *FakePluginContext) IsLoggedInCallCount() int {
	fake.isLoggedInMutex.RLock()
	defer fake.isLoggedInMutex.RUnlock()
	return len(fake.isLoggedInArgsForCall)
}

func (fake *FakePluginContext) IsLoggedInReturns(result1 bool) {
	fake.IsLoggedInStub = nil
	fake.isLoggedInReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePluginContext) AccessToken() string {
	fake.accessTokenMutex.Lock()
	fake.accessTokenArgsForCall = append(fake.accessTokenArgsForCall, struct{}{})
	fake.accessTokenMutex.Unlock()
	if fake.AccessTokenStub != nil {
		return fake.AccessTokenStub()
	} else {
		return fake.accessTokenReturns.result1
	}
}

func (fake *FakePluginContext) AccessTokenCallCount() int {
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	return len(fake.accessTokenArgsForCall)
}

func (fake *FakePluginContext) AccessTokenReturns(result1 string) {
	fake.AccessTokenStub = nil
	fake.accessTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) TokenRefresh() (string, error) {
	fake.tokenRefreshMutex.Lock()
	fake.tokenRefreshArgsForCall = append(fake.tokenRefreshArgsForCall, struct{}{})
	fake.tokenRefreshMutex.Unlock()
	if fake.TokenRefreshStub != nil {
		return fake.TokenRefreshStub()
	} else {
		return fake.tokenRefreshReturns.result1, fake.tokenRefreshReturns.result2
	}
}

func (fake *FakePluginContext) TokenRefreshCallCount() int {
	fake.tokenRefreshMutex.RLock()
	defer fake.tokenRefreshMutex.RUnlock()
	return len(fake.tokenRefreshArgsForCall)
}

func (fake *FakePluginContext) TokenRefreshReturns(result1 string, result2 error) {
	fake.TokenRefreshStub = nil
	fake.tokenRefreshReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakePluginContext) CurrentOrg() models.Organization {
	fake.currentOrgMutex.Lock()
	fake.currentOrgArgsForCall = append(fake.currentOrgArgsForCall, struct{}{})
	fake.currentOrgMutex.Unlock()
	if fake.CurrentOrgStub != nil {
		return fake.CurrentOrgStub()
	} else {
		return fake.currentOrgReturns.result1
	}
}

func (fake *FakePluginContext) CurrentOrgCallCount() int {
	fake.currentOrgMutex.RLock()
	defer fake.currentOrgMutex.RUnlock()
	return len(fake.currentOrgArgsForCall)
}

func (fake *FakePluginContext) CurrentOrgReturns(result1 models.Organization) {
	fake.CurrentOrgStub = nil
	fake.currentOrgReturns = struct {
		result1 models.Organization
	}{result1}
}

func (fake *FakePluginContext) HasOrganization() bool {
	fake.hasOrganizationMutex.Lock()
	fake.hasOrganizationArgsForCall = append(fake.hasOrganizationArgsForCall, struct{}{})
	fake.hasOrganizationMutex.Unlock()
	if fake.HasOrganizationStub != nil {
		return fake.HasOrganizationStub()
	} else {
		return fake.hasOrganizationReturns.result1
	}
}

func (fake *FakePluginContext) HasOrganizationCallCount() int {
	fake.hasOrganizationMutex.RLock()
	defer fake.hasOrganizationMutex.RUnlock()
	return len(fake.hasOrganizationArgsForCall)
}

func (fake *FakePluginContext) HasOrganizationReturns(result1 bool) {
	fake.HasOrganizationStub = nil
	fake.hasOrganizationReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePluginContext) CurrentSpace() models.Space {
	fake.currentSpaceMutex.Lock()
	fake.currentSpaceArgsForCall = append(fake.currentSpaceArgsForCall, struct{}{})
	fake.currentSpaceMutex.Unlock()
	if fake.CurrentSpaceStub != nil {
		return fake.CurrentSpaceStub()
	} else {
		return fake.currentSpaceReturns.result1
	}
}

func (fake *FakePluginContext) CurrentSpaceCallCount() int {
	fake.currentSpaceMutex.RLock()
	defer fake.currentSpaceMutex.RUnlock()
	return len(fake.currentSpaceArgsForCall)
}

func (fake *FakePluginContext) CurrentSpaceReturns(result1 models.Space) {
	fake.CurrentSpaceStub = nil
	fake.currentSpaceReturns = struct {
		result1 models.Space
	}{result1}
}

func (fake *FakePluginContext) HasSpace() bool {
	fake.hasSpaceMutex.Lock()
	fake.hasSpaceArgsForCall = append(fake.hasSpaceArgsForCall, struct{}{})
	fake.hasSpaceMutex.Unlock()
	if fake.HasSpaceStub != nil {
		return fake.HasSpaceStub()
	} else {
		return fake.hasSpaceReturns.result1
	}
}

func (fake *FakePluginContext) HasSpaceCallCount() int {
	fake.hasSpaceMutex.RLock()
	defer fake.hasSpaceMutex.RUnlock()
	return len(fake.hasSpaceArgsForCall)
}

func (fake *FakePluginContext) HasSpaceReturns(result1 bool) {
	fake.HasSpaceStub = nil
	fake.hasSpaceReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePluginContext) Locale() string {
	fake.localeMutex.Lock()
	fake.localeArgsForCall = append(fake.localeArgsForCall, struct{}{})
	fake.localeMutex.Unlock()
	if fake.LocaleStub != nil {
		return fake.LocaleStub()
	} else {
		return fake.localeReturns.result1
	}
}

func (fake *FakePluginContext) LocaleCallCount() int {
	fake.localeMutex.RLock()
	defer fake.localeMutex.RUnlock()
	return len(fake.localeArgsForCall)
}

func (fake *FakePluginContext) LocaleReturns(result1 string) {
	fake.LocaleStub = nil
	fake.localeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) Trace() string {
	fake.traceMutex.Lock()
	fake.traceArgsForCall = append(fake.traceArgsForCall, struct{}{})
	fake.traceMutex.Unlock()
	if fake.TraceStub != nil {
		return fake.TraceStub()
	} else {
		return fake.traceReturns.result1
	}
}

func (fake *FakePluginContext) TraceCallCount() int {
	fake.traceMutex.RLock()
	defer fake.traceMutex.RUnlock()
	return len(fake.traceArgsForCall)
}

func (fake *FakePluginContext) TraceReturns(result1 string) {
	fake.TraceStub = nil
	fake.traceReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) ColorEnabled() string {
	fake.colorEnabledMutex.Lock()
	fake.colorEnabledArgsForCall = append(fake.colorEnabledArgsForCall, struct{}{})
	fake.colorEnabledMutex.Unlock()
	if fake.ColorEnabledStub != nil {
		return fake.ColorEnabledStub()
	} else {
		return fake.colorEnabledReturns.result1
	}
}

func (fake *FakePluginContext) ColorEnabledCallCount() int {
	fake.colorEnabledMutex.RLock()
	defer fake.colorEnabledMutex.RUnlock()
	return len(fake.colorEnabledArgsForCall)
}

func (fake *FakePluginContext) ColorEnabledReturns(result1 string) {
	fake.ColorEnabledStub = nil
	fake.colorEnabledReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) IsSSLDisabled() bool {
	fake.isSSLDisabledMutex.Lock()
	fake.isSSLDisabledArgsForCall = append(fake.isSSLDisabledArgsForCall, struct{}{})
	fake.isSSLDisabledMutex.Unlock()
	if fake.IsSSLDisabledStub != nil {
		return fake.IsSSLDisabledStub()
	} else {
		return fake.isSSLDisabledReturns.result1
	}
}

func (fake *FakePluginContext) IsSSLDisabledCallCount() int {
	fake.isSSLDisabledMutex.RLock()
	defer fake.isSSLDisabledMutex.RUnlock()
	return len(fake.isSSLDisabledArgsForCall)
}

func (fake *FakePluginContext) IsSSLDisabledReturns(result1 bool) {
	fake.IsSSLDisabledStub = nil
	fake.isSSLDisabledReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePluginContext) PluginDirectory() string {
	fake.pluginDirectoryMutex.Lock()
	fake.pluginDirectoryArgsForCall = append(fake.pluginDirectoryArgsForCall, struct{}{})
	fake.pluginDirectoryMutex.Unlock()
	if fake.PluginDirectoryStub != nil {
		return fake.PluginDirectoryStub()
	} else {
		return fake.pluginDirectoryReturns.result1
	}
}

func (fake *FakePluginContext) PluginDirectoryCallCount() int {
	fake.pluginDirectoryMutex.RLock()
	defer fake.pluginDirectoryMutex.RUnlock()
	return len(fake.pluginDirectoryArgsForCall)
}

func (fake *FakePluginContext) PluginDirectoryReturns(result1 string) {
	fake.PluginDirectoryStub = nil
	fake.pluginDirectoryReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginContext) HTTPTimeout() int {
	fake.hTTPTimeoutMutex.Lock()
	fake.hTTPTimeoutArgsForCall = append(fake.hTTPTimeoutArgsForCall, struct{}{})
	fake.hTTPTimeoutMutex.Unlock()
	if fake.HTTPTimeoutStub != nil {
		return fake.HTTPTimeoutStub()
	} else {
		return fake.hTTPTimeoutReturns.result1
	}
}

func (fake *FakePluginContext) HTTPTimeoutCallCount() int {
	fake.hTTPTimeoutMutex.RLock()
	defer fake.hTTPTimeoutMutex.RUnlock()
	return len(fake.hTTPTimeoutArgsForCall)
}

func (fake *FakePluginContext) HTTPTimeoutReturns(result1 int) {
	fake.HTTPTimeoutStub = nil
	fake.hTTPTimeoutReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakePluginContext) PluginConfig() plugin.PluginConfig {
	fake.pluginConfigMutex.Lock()
	fake.pluginConfigArgsForCall = append(fake.pluginConfigArgsForCall, struct{}{})
	fake.pluginConfigMutex.Unlock()
	if fake.PluginConfigStub != nil {
		return fake.PluginConfigStub()
	} else {
		return fake.pluginConfigReturns.result1
	}
}

func (fake *FakePluginContext) PluginConfigCallCount() int {
	fake.pluginConfigMutex.RLock()
	defer fake.pluginConfigMutex.RUnlock()
	return len(fake.pluginConfigArgsForCall)
}

func (fake *FakePluginContext) PluginConfigReturns(result1 plugin.PluginConfig) {
	fake.PluginConfigStub = nil
	fake.pluginConfigReturns = struct {
		result1 plugin.PluginConfig
	}{result1}
}

func (fake *FakePluginContext) CommandNamespace() string {
	fake.commandNamespaceMutex.Lock()
	fake.commandNamespaceArgsForCall = append(fake.commandNamespaceArgsForCall, struct{}{})
	fake.commandNamespaceMutex.Unlock()
	if fake.CommandNamespaceStub != nil {
		return fake.CommandNamespaceStub()
	} else {
		return fake.commandNamespaceReturns.result1
	}
}

func (fake *FakePluginContext) CommandNamespaceCallCount() int {
	fake.commandNamespaceMutex.RLock()
	defer fake.commandNamespaceMutex.RUnlock()
	return len(fake.commandNamespaceArgsForCall)
}

func (fake *FakePluginContext) CommandNamespaceReturns(result1 string) {
	fake.CommandNamespaceStub = nil
	fake.commandNamespaceReturns = struct {
		result1 string
	}{result1}
}

var _ plugin.PluginContext = new(FakePluginContext)
