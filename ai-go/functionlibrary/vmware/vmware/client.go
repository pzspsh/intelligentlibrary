package govmomi

import (
	"context"
	"net/url"

	"function/vmware/vmware/session"
	"function/vmware/vmware/vim25"
	"function/vmware/vmware/vim25/soap"
)

type Client struct {
	*vim25.Client
	SessionManager *session.Manager
}

func NewClient(ctx context.Context, u *url.URL, proxy string, insecure bool) (*Client, error) {
	soapClient := soap.NewClient(u, proxy, insecure)
	vimClient, err := vim25.NewClient(ctx, soapClient)
	if err != nil {
		return nil, err
	}
	c := &Client{
		Client:         vimClient,
		SessionManager: session.NewManager(vimClient),
	}
	if u.User != nil {
		err = c.Login(ctx, u.User)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Client) Login(ctx context.Context, u *url.Userinfo) error {
	return c.SessionManager.Login(ctx, u)
}

// Logout dispatches to the SessionManager.
func (c *Client) Logout(ctx context.Context) error {
	// Close any idle connections after logging out.
	defer c.Client.CloseIdleConnections()
	return c.SessionManager.Logout(ctx)
}
