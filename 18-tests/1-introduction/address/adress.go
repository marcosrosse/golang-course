package address

import "strings"

// AddressType represents a postal address.
func AddressType(address string) string {
	validTypes := []string{"rua", "avenida", "rodovia", "estrada", "alameda"}

	addressLowerCase := strings.ToLower(address)
	firstWord := strings.Split(addressLowerCase, " ")[0]

	addressHasValidType := false
	for _, typeAdress := range validTypes {
		if typeAdress == firstWord {
			addressHasValidType = true
		}
	}

	if addressHasValidType {
		return strings.Title(firstWord)
	}

	return "Unknown address type"

}
