/*
 * swaggerspringboot
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UserDraft struct {

	FirstName string `json:"firstName"`

	LastName string `json:"lastName"`

	Email string `json:"email"`
}

// AssertUserDraftRequired checks if the required fields are not zero-ed
func AssertUserDraftRequired(obj UserDraft) error {
	elements := map[string]interface{}{
		"firstName": obj.FirstName,
		"lastName": obj.LastName,
		"email": obj.Email,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUserDraftRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UserDraft (e.g. [][]UserDraft), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserDraftRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUserDraft, ok := obj.(UserDraft)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserDraftRequired(aUserDraft)
	})
}
