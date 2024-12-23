package session

import (
	"context"
	"net/url"
	"os"

	"function/vmware/vmware/vim25"
	"function/vmware/vmware/vim25/methods"
	"function/vmware/vmware/vim25/types"
)

var Locale = os.Getenv("GOVMOMI_LOCALE")

func init() {
	if Locale == "_" {
		Locale = ""
	} else if Locale == "" {
		Locale = "en_US"
	}
}

type Manager struct {
	client      *vim25.Client
	userSession *types.UserSession
}

func NewManager(client *vim25.Client) *Manager {
	m := Manager{
		client: client,
	}

	return &m
}

func (sm *Manager) Login(ctx context.Context, u *url.Userinfo) error {
	req := types.Login{
		This:   sm.Reference(),
		Locale: Locale,
	}
	if u != nil {
		req.UserName = u.Username()
		if pw, ok := u.Password(); ok {
			req.Password = pw
		}
	}
	login, err := methods.Login(ctx, sm.client, &req)
	if err != nil {
		return err
	}
	sm.userSession = &login.Returnval
	return nil
}

func (sm Manager) Reference() types.ManagedObjectReference {
	return *sm.client.ServiceContent.SessionManager
}

// See: https://kb.vmware.com/s/article/2004305
func (sm *Manager) LoginExtensionByCertificate(ctx context.Context, key string) error {
	c := sm.client
	u := c.URL()
	if u.Hostname() != "sdkTunnel" {
		sc := c.Tunnel()
		c = &vim25.Client{
			Client:         sc,
			RoundTripper:   sc,
			ServiceContent: c.ServiceContent,
		}
		// When http.Transport.Proxy is used, our thumbprint checker is bypassed, resulting in:
		// "Post https://sdkTunnel:8089/sdk: x509: certificate is valid for $vcenter_hostname, not sdkTunnel"
		// The only easy way around this is to disable verification for the call to LoginExtensionByCertificate().
		// TODO: find a way to avoid disabling InsecureSkipVerify.
		c.DefaultTransport().TLSClientConfig.InsecureSkipVerify = true
	}

	req := types.LoginExtensionByCertificate{
		This:         sm.Reference(),
		ExtensionKey: key,
		Locale:       Locale,
	}

	login, err := methods.LoginExtensionByCertificate(ctx, c, &req)
	if err != nil {
		return err
	}

	// Copy the session cookie
	sm.client.Jar.SetCookies(u, c.Jar.Cookies(c.URL()))

	sm.userSession = &login.Returnval
	return nil
}

func (sm *Manager) LoginByToken(ctx context.Context) error {
	req := types.LoginByToken{
		This:   sm.Reference(),
		Locale: Locale,
	}

	login, err := methods.LoginByToken(ctx, sm.client, &req)
	if err != nil {
		return err
	}

	sm.userSession = &login.Returnval
	return nil
}

func (sm *Manager) Logout(ctx context.Context) error {
	req := types.Logout{
		This: sm.Reference(),
	}

	_, err := methods.Logout(ctx, sm.client, &req)
	if err != nil {
		return err
	}

	sm.userSession = nil
	return nil
}

func (sm *Manager) TerminateSession(ctx context.Context, sessionId []string) error {
	req := types.TerminateSession{
		This:      sm.Reference(),
		SessionId: sessionId,
	}

	_, err := methods.TerminateSession(ctx, sm.client, &req)
	return err
}

// still valid. This function only works against vCenter.
func (sm *Manager) SessionIsActive(ctx context.Context) (bool, error) {
	if sm.userSession == nil {
		return false, nil
	}

	req := types.SessionIsActive{
		This:      sm.Reference(),
		SessionID: sm.userSession.Key,
		UserName:  sm.userSession.UserName,
	}

	active, err := methods.SessionIsActive(ctx, sm.client, &req)
	if err != nil {
		return false, err
	}

	return active.Returnval, err
}

func (sm *Manager) AcquireGenericServiceTicket(ctx context.Context, spec types.BaseSessionManagerServiceRequestSpec) (*types.SessionManagerGenericServiceTicket, error) {
	req := types.AcquireGenericServiceTicket{
		This: sm.Reference(),
		Spec: spec,
	}

	res, err := methods.AcquireGenericServiceTicket(ctx, sm.client, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (sm *Manager) AcquireLocalTicket(ctx context.Context, userName string) (*types.SessionManagerLocalTicket, error) {
	req := types.AcquireLocalTicket{
		This:     sm.Reference(),
		UserName: userName,
	}

	res, err := methods.AcquireLocalTicket(ctx, sm.client, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (sm *Manager) AcquireCloneTicket(ctx context.Context) (string, error) {
	req := types.AcquireCloneTicket{
		This: sm.Reference(),
	}

	res, err := methods.AcquireCloneTicket(ctx, sm.client, &req)
	if err != nil {
		return "", err
	}

	return res.Returnval, nil
}

func (sm *Manager) CloneSession(ctx context.Context, ticket string) error {
	req := types.CloneSession{
		This:        sm.Reference(),
		CloneTicket: ticket,
	}

	res, err := methods.CloneSession(ctx, sm.client, &req)
	if err != nil {
		return err
	}

	sm.userSession = &res.Returnval
	return nil
}

func (sm *Manager) UpdateServiceMessage(ctx context.Context, message string) error {
	req := types.UpdateServiceMessage{
		This:    sm.Reference(),
		Message: message,
	}

	_, err := methods.UpdateServiceMessage(ctx, sm.client, &req)

	return err
}
