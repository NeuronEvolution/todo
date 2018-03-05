// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetFriendsListURL generates an URL for the get friends list operation
type GetFriendsListURL struct {
	PageSize  *int64
	PageToken *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetFriendsListURL) WithBasePath(bp string) *GetFriendsListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetFriendsListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetFriendsListURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/friends"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api-private/v1/todo"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var pageSize string
	if o.PageSize != nil {
		pageSize = swag.FormatInt64(*o.PageSize)
	}
	if pageSize != "" {
		qs.Set("pageSize", pageSize)
	}

	var pageToken string
	if o.PageToken != nil {
		pageToken = *o.PageToken
	}
	if pageToken != "" {
		qs.Set("pageToken", pageToken)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetFriendsListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetFriendsListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetFriendsListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetFriendsListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetFriendsListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetFriendsListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}