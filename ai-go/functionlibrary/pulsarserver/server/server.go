/*
@File   : server.go
@Author : pan
@Time   : 2024-06-03 10:43:23
*/
package server

import (
	"strings"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type PulsarOptions struct {
	PulsarUrl string `json:"pulsar_url"`
	CertPath  string `json:"cert_path"`
	Topic     string `json:"topic"`
	Subscribe string `json:"subscribe"`
	client    pulsar.Client
}

func (p *PulsarOptions) PulsarConn() (*PulsarOptions, error) {
	var err error
	var client pulsar.Client
	if strings.Contains(p.PulsarUrl, "pulsar+ssl") {
		if client, err = pulsar.NewClient(pulsar.ClientOptions{
			URL:                        p.PulsarUrl,
			TLSTrustCertsFilePath:      p.CertPath,
			TLSAllowInsecureConnection: true,
			OperationTimeout:           30 * time.Second,
			ConnectionTimeout:          30 * time.Second,
		}); err != nil {
			return nil, err
		}
	} else {
		client, err = pulsar.NewClient(pulsar.ClientOptions{
			URL:               p.PulsarUrl,
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		})
		if err != nil {
			return nil, err
		}
	}
	p.client = client
	return p, nil
}
