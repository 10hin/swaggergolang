/*
 * swaggerspringboot
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UserPatch struct {

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Email string `json:"email,omitempty"`
}

// AssertUserPatchRequired checks if the required fields are not zero-ed
func AssertUserPatchRequired(obj UserPatch) error {
	return nil
}

// AssertRecurseUserPatchRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UserPatch (e.g. [][]UserPatch), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserPatchRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUserPatch, ok := obj.(UserPatch)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserPatchRequired(aUserPatch)
	})
}
