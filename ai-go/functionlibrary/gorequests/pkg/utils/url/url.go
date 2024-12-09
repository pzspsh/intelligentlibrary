package urlutil

import (
	"bytes"
	"net/url"
	"strings"

	errorutil "function/gorequests/pkg/utils/errors"
	osutils "function/gorequests/pkg/utils/os"
	stringsutil "function/gorequests/pkg/utils/strings"
)

var DisableAutoCorrect bool

// URL a wrapper around net/url.URL
type URL struct {
	*url.URL

	Original   string         // original or given url(without params if any)
	Unsafe     bool           // If request is unsafe (skip validation)
	IsRelative bool           // If URL is relative
	Params     *OrderedParams // Query Parameters
	// should call Update() method when directly updating wrapped url.URL or parameters
}

// mergepath merges given relative path
func (u *URL) MergePath(newrelpath string, unsafe bool) error {
	if newrelpath == "" {
		return nil
	}
	ux, err := ParseRelativePath(newrelpath, unsafe)
	if err != nil {
		return err
	}
	u.Params.Merge(ux.Params.Encode())
	u.Path = mergePaths(u.Path, ux.Path)
	if ux.Fragment != "" {
		u.Fragment = ux.Fragment
	}
	return nil
}

// UpdateRelPath updates relative path with new path (existing params are not removed)
func (u *URL) UpdateRelPath(newrelpath string, unsafe bool) error {
	u.Path = ""
	return u.MergePath(newrelpath, unsafe)
}

// Updates internal wrapped url.URL with any changes done to Query Parameters
func (u *URL) Update() {
	if u.Params != nil {
		u.RawQuery = u.Params.Encode()
	}
}

// Query returns Query Params
func (u *URL) Query() *OrderedParams {
	return u.Params
}

// Clone
func (u *URL) Clone() *URL {
	var userinfo *url.Userinfo
	if u.User != nil {
		// userinfo is immutable so this is the only way
		tempurl := HTTPS + SchemeSeparator + u.User.String() + "@" + "scanme.sh/"
		turl, _ := url.Parse(tempurl)
		if turl != nil {
			userinfo = turl.User
		}
	}
	ux := &url.URL{
		Scheme:      u.Scheme,
		Opaque:      u.Opaque,
		User:        userinfo,
		Host:        u.Host,
		Path:        u.Path,
		RawPath:     u.RawPath,
		RawQuery:    u.RawQuery,
		Fragment:    u.Fragment,
		ForceQuery:  u.ForceQuery,
		RawFragment: u.RawFragment,
	}
	params := u.Params.Clone()
	return &URL{
		URL:        ux,
		Params:     params,
		Original:   u.Original,
		Unsafe:     u.Unsafe,
		IsRelative: u.IsRelative,
	}
}

// String
func (u *URL) String() string {
	var buff bytes.Buffer
	if u.Scheme != "" {
		buff.WriteString(u.Scheme + "://")
	}
	if u.User != nil {
		buff.WriteString(u.User.String())
		buff.WriteRune('@')
	}
	buff.WriteString(u.Host)
	buff.WriteString(u.GetRelativePath())
	return buff.String()
}

func (u *URL) EscapedString() string {
	var buff bytes.Buffer
	host := u.Host
	if osutils.IsWindows() {
		host = strings.ReplaceAll(host, ":", "_")
	}
	buff.WriteString(host)
	if u.Path != "" && u.Path != "/" {
		buff.WriteString("_" + strings.ReplaceAll(u.Path, "/", "_"))
	}
	return buff.String()
}

// GetRelativePath ex: /some/path?param=true#fragment
func (u *URL) GetRelativePath() string {
	var buff bytes.Buffer
	if u.Path != "" {
		if !strings.HasPrefix(u.Path, "/") {
			buff.WriteRune('/')
		}
		buff.WriteString(u.Path)
	}
	if u.Params.om.Len() > 0 {
		buff.WriteRune('?')
		buff.WriteString(u.Params.Encode())
	}
	if u.Fragment != "" {
		buff.WriteRune('#')
		buff.WriteString(u.Fragment)
	}
	return buff.String()
}

// Updates port
func (u *URL) UpdatePort(newport string) {
	if newport == "" {
		return
	}
	if u.URL.Port() != "" {
		u.Host = strings.Replace(u.Host, u.Port(), newport, 1)
		return
	}
	u.Host += ":" + newport
}

// TrimPort if any
func (u *URL) TrimPort() {
	u.URL.Host = u.Hostname()
}

// parseRelativePath parses relative path from Original Path without relying on
// net/url.URL
func (u *URL) parseUnsafeRelativePath() {
	defer func() {
		if !strings.HasPrefix(u.Path, "/") && u.Path != "" {
			u.Path = "/" + u.Path
		}
	}()

	if u.Original != u.Path {
		// params and fragements are removed from Original in Parsexx() therefore they can be compared
		u.Path = u.Original
	}

	// percent encoding in path
	if u.Host == "" || len(u.Host) < 4 {
		if shouldEscape(u.Original) {
			u.Path = u.Original
		}
		return
	}
	expectedPath := strings.SplitN(u.Original, u.Host, 2)
	if len(expectedPath) != 2 {
		// something went wrong fail silently
		return
	}
	u.Path = expectedPath[1]
}

// fetchParams retrieves query parameters from URL
func (u *URL) fetchParams() {
	if u.Params == nil {
		u.Params = NewOrderedParams()
	}
	// parse fragments if any
	if i := strings.IndexRune(u.Original, '#'); i != -1 {
		// assuming ?param=value#highlight
		u.Fragment = u.Original[i+1:]
		u.Original = u.Original[:i]
	}
	if index := strings.IndexRune(u.Original, '?'); index == -1 {
		return
	} else {
		encodedParams := u.Original[index+1:]
		u.Params.Decode(encodedParams)
		u.Original = u.Original[:index]
	}
	u.Update()
}

// ParseURL
func Parse(inputURL string) (*URL, error) {
	return ParseURL(inputURL, false)
}

// Parse and return URL
func ParseURL(inputURL string, unsafe bool) (*URL, error) {
	u := &URL{
		URL:      &url.URL{},
		Original: inputURL,
		Unsafe:   unsafe,
		Params:   NewOrderedParams(),
	}
	u.fetchParams()
	// filter out fragments and parameters only then parse path
	inputURL = u.Original
	if inputURL == "" {
		return nil, errorutil.NewWithTag("urlutil", "failed to parse url got empty input")
	}

	if strings.HasPrefix(inputURL, "/") && !strings.HasPrefix(inputURL, "//") {
		// this is definitely a relative path
		u.IsRelative = true
		u.Path = u.Original
		return u, nil
	}
	// Try to parse host related input
	if stringsutil.HasPrefixAny(inputURL, HTTP+SchemeSeparator, HTTPS+SchemeSeparator, "//") || strings.Contains(inputURL, "://") {
		u.IsRelative = false
		urlparse, parseErr := url.Parse(inputURL)
		if parseErr != nil {
			// for parse errors in unsafe way try parsing again
			if unsafe {
				urlparse = parseUnsafeFullURL(inputURL)
				if urlparse != nil {
					parseErr = nil
				}
			}
			if parseErr != nil {
				return nil, errorutil.NewWithErr(parseErr).Msgf("failed to parse url")
			}
		}
		copy(u.URL, urlparse)
	} else {
		urlparse, parseErr := url.Parse(HTTPS + SchemeSeparator + inputURL)
		if parseErr != nil {
			// most likely a relativeurl
			u.IsRelative = true
			// TODO: investigate if prefix / should be added
		} else {
			urlparse.Scheme = "" // remove newly added scheme
			copy(u.URL, urlparse)
		}
	}

	// try parsing path
	if !u.IsRelative {
		if u.Host == "" {
			// this is unexpected case return err
			return nil, errorutil.NewWithTag("urlutil", "failed to parse url %v got empty host", inputURL)
		}
		if !strings.Contains(u.Host, ".") && !strings.Contains(u.Host, ":") && u.Host != "localhost" {
			if !DisableAutoCorrect {
				u.IsRelative = true
				u.Path = inputURL
				u.Host = ""
			}
		}
	}
	if !u.IsRelative && u.Host == "" {
		return nil, errorutil.NewWithTag("urlutil", "failed to parse url `%v`", inputURL).Msgf("got empty host when url is not relative")
	}
	if u.IsRelative {
		return ParseRelativePath(inputURL, unsafe)
	}
	return u, nil
}

// ParseRelativePath parses and returns relative path
func ParseRelativePath(inputURL string, unsafe bool) (*URL, error) {
	u := &URL{
		URL:        &url.URL{},
		Original:   inputURL,
		Unsafe:     unsafe,
		IsRelative: true,
	}
	u.fetchParams()
	urlparse, parseErr := url.Parse(inputURL)
	if parseErr != nil {
		if !unsafe {
			return nil, errorutil.NewWithErr(parseErr).WithTag("urlutil").Msgf("failed to parse input url")
		} else {
			// if unsafe do not rely on net/url.Parse
			u.Path = inputURL
		}
	}
	if urlparse != nil {
		urlparse.Host = ""
		copy(u.URL, urlparse)
	}
	u.parseUnsafeRelativePath()
	return u, nil
}

// parseUnsafeFullURL parses invalid(unsafe) urls (ex: https://scanme.sh/%invalid)
// this is not supported as per RFC and url.Parse fails
func parseUnsafeFullURL(urlx string) *url.URL {
	temp := strings.Replace(urlx, "//", "", 1)
	index := strings.IndexRune(temp, '/')
	if index == -1 {
		return nil
	}
	urlPath := temp[index:]
	urlHost := strings.TrimSuffix(urlx, urlPath)
	parseURL, parseErr := url.Parse(urlHost)
	if parseErr != nil {
		return nil
	}
	if relpath, err := ParseRelativePath(urlPath, true); err == nil {
		parseURL.Path = relpath.Path
		return parseURL
	}
	return nil
}

// copy parsed data from src to dst this does not include fragment or params
func copy(dst *url.URL, src *url.URL) {
	dst.Host = src.Host
	dst.Opaque = src.Opaque
	dst.Path = src.Path
	dst.RawPath = src.RawPath
	dst.Scheme = src.Scheme
	dst.User = src.User
}
