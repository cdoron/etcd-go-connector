/*
 * Data Catalog Service - Asset Details
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Connection struct {

	Name string `json:"name"`
}

// AssertConnectionRequired checks if the required fields are not zero-ed
func AssertConnectionRequired(obj Connection) error {
	elements := map[string]interface{}{
		"name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseConnectionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Connection (e.g. [][]Connection), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseConnectionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aConnection, ok := obj.(Connection)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertConnectionRequired(aConnection)
	})
}
