/*
 * API for user admin panel
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: korotkov.ivan.s@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"context"
	"net/http"
)



// UsersApiRouter defines the required methods for binding the api requests to a responses for the UsersApi
// The UsersApiRouter implementation should parse necessary information from the http request,
// pass the data to a UsersApiServicer to perform the required actions, then write the service results to the http response.
type UsersApiRouter interface { 
	All(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Single(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
}


// UsersApiServicer defines the api actions for the UsersApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UsersApiServicer interface { 
	All(context.Context) (ImplResponse, error)
	Create(context.Context, NewUser) (ImplResponse, error)
	Delete(context.Context, int32) (ImplResponse, error)
	Single(context.Context, int32) (ImplResponse, error)
	Update(context.Context, int32, User) (ImplResponse, error)
}