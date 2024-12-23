package vim25

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

	"function/vmware/vmware/vim25/methods"
	"function/vmware/vmware/vim25/soap"
	"function/vmware/vmware/vim25/types"
	"function/vmware/vmware/vim25/xml"
)

const (
	Namespace = "vim25"
	Version   = "7.0"
	Path      = "/sdk"
)

type Client struct {
	*soap.Client
	ServiceContent types.ServiceContent
	RoundTripper   soap.RoundTripper
}

func NewClient(ctx context.Context, rt soap.RoundTripper) (*Client, error) {
	c := Client{
		RoundTripper: rt,
	}
	if sc, ok := rt.(*soap.Client); ok {
		c.Client = sc
		if c.Namespace == "" {
			c.Namespace = "urn:" + Namespace
		} else if !strings.Contains(c.Namespace, ":") {
			c.Namespace = "urn:" + c.Namespace
		}
		if c.Version == "" {
			c.Version = Version
		}
	}
	var err error
	c.ServiceContent, err = methods.GetServiceContent(ctx, rt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// UseServiceVersion sets soap.Client.Version to the current version of the service endpoint via /sdk/vimServiceVersions.xml
func (c *Client) UseServiceVersion(kind ...string) error {
	ns := "vim"
	if len(kind) != 0 {
		ns = kind[0]
	}

	u := c.URL()
	u.Path = path.Join(Path, ns+"ServiceVersions.xml")

	res, err := c.Get(u.String())
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Get(%s): %s", u.Path, err)
	}

	v := struct {
		Namespace *string `xml:"namespace>name"`
		Version   *string `xml:"namespace>version"`
	}{
		&c.Namespace,
		&c.Version,
	}

	err = xml.NewDecoder(res.Body).Decode(&v)
	_ = res.Body.Close()
	if err != nil {
		return fmt.Errorf("xml.Decode(%s): %s", u.Path, err)
	}

	return nil
}

// RoundTrip dispatches to the RoundTripper field.
func (c *Client) RoundTrip(ctx context.Context, req, res soap.HasFault) error {
	return c.RoundTripper.RoundTrip(ctx, req, res)
}

type marshaledClient struct {
	SoapClient     *soap.Client
	ServiceContent types.ServiceContent
}

func (c *Client) MarshalJSON() ([]byte, error) {
	m := marshaledClient{
		SoapClient:     c.Client,
		ServiceContent: c.ServiceContent,
	}

	return json.Marshal(m)
}

func (c *Client) UnmarshalJSON(b []byte) error {
	var m marshaledClient

	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	*c = Client{
		Client:         m.SoapClient,
		ServiceContent: m.ServiceContent,
		RoundTripper:   m.SoapClient,
	}

	return nil
}

// Valid returns whether or not the client is valid and ready for use.
// This should be called after unmarshalling the client.
func (c *Client) Valid() bool {
	if c == nil {
		return false
	}

	if c.Client == nil {
		return false
	}

	// Use arbitrary pointer field in the service content.
	// Doesn't matter which one, as long as it is populated by default.
	if c.ServiceContent.SessionManager == nil {
		return false
	}

	return true
}

// Path returns vim25.Path (see cache.Client)
func (c *Client) Path() string {
	return Path
}

// IsVC returns true if we are connected to a vCenter
func (c *Client) IsVC() bool {
	return c.ServiceContent.About.ApiType == "VirtualCenter"
}
