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

type UserId struct {
	Id int32 `json:"id"`
}

// AssertUserIdRequired checks if the required fields are not zero-ed
func AssertUserIdRequired(obj UserId) error {
	elements := map[string]interface{}{
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUserIdRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UserId (e.g. [][]UserId), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserIdRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUserId, ok := obj.(UserId)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserIdRequired(aUserId)
	})
}
