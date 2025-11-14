package MapUtil

import (
	"errors"
	"fmt"
	"strings"
)

func MustGet[T any](key string, data map[string]interface{}) T {
	result, err := Get[T](key, data)
	if err != nil {
		panic(err)
	}
	return result
}

// Get function that supports deep key retrieval in a map using '.' as a separator, and returns the value as type T.
func Get[T any](key string, data map[string]interface{}) (T, error) {
	// Split the key by '.' to handle nested keys
	keys := strings.Split(key, ".")

	var result interface{} = data

	// Traverse the map based on the keys
	for _, k := range keys {
		if val, exists := result.(map[string]interface{})[k]; exists {
			// Move deeper into the nested map
			result = val
		} else {
			// If key is not found, return the zero value for type T and an error
			var zero T
			return zero, errors.New("key not found")
		}
	}

	// Try to assert the final result into type T
	if finalValue, ok := result.(T); ok {
		return finalValue, nil
	} else {
		// If type assertion fails, return the zero value of T and an error
		var zero T
		return zero, fmt.Errorf("type assertion failed, expected %T but got %T", zero, result)
	}
}
