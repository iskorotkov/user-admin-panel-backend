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

type Error struct {

	// Unique error code is useful for differentiating different types of errors. Error code is globally unique and permanent, meaning that it will never change even if the error message changes
	Code int32 `json:"code,omitempty"`

	Message string `json:"message,omitempty"`

	Errors []string `json:"errors,omitempty"`
}

// AssertErrorRequired checks if the required fields are not zero-ed
func AssertErrorRequired(obj Error) error {
	return nil
}

// AssertRecurseErrorRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Error (e.g. [][]Error), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseErrorRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aError, ok := obj.(Error)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertErrorRequired(aError)
	})
}