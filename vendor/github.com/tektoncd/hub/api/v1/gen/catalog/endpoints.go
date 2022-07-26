// Code generated by goa v3.7.6, DO NOT EDIT.
//
// catalog endpoints
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package catalog

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "catalog" service endpoints.
type Endpoints struct {
	List goa.Endpoint
}

// NewEndpoints wraps the methods of the "catalog" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List: NewListEndpoint(s),
	}
}

// Use applies the given middleware to all the "catalog" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
}

// NewListEndpoint returns an endpoint function that calls the method "List" of
// service "catalog".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.List(ctx)
	}
}
