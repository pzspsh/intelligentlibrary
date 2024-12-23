/*
@File   : retry.go
@Author : pan
@Time   : 2023-08-24 15:37:28
*/
package http

import (
	"context"
	"crypto/x509"
	"net/http"
	"net/url"
	"regexp"
)

var (
	redirectsErrorRegex = regexp.MustCompile(`stopped after \d+ redirects\z`)
	schemeErrorRegex    = regexp.MustCompile(`unsupported protocol scheme`)
)

type CheckRetry func(ctx context.Context, resp *http.Response, err error) (bool, error)

func DefaultRetryPolicy() func(ctx context.Context, resp *http.Response, err error) (bool, error) {
	return func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		return CheckRecoverableErrors(ctx, resp, err)
	}
}

func HostSprayRetryPolicy() func(ctx context.Context, resp *http.Response, err error) (bool, error) {
	return func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		return CheckRecoverableErrors(ctx, resp, err)
	}
}

func CheckRecoverableErrors(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	if err != nil {
		if v, ok := err.(*url.Error); ok {
			if redirectsErrorRegex.MatchString(v.Error()) {
				return false, nil
			}

			if schemeErrorRegex.MatchString(v.Error()) {
				return false, nil
			}

			if _, ok := v.Err.(x509.UnknownAuthorityError); ok {
				return false, nil
			}
		}
		return true, nil
	}
	return false, nil
}
