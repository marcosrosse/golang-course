package address

import "testing"

type testCase struct {
	addressInserted string
	addressExpected string
}

func TestAddressType(t *testing.T) {
	t.Parallel() // Run tests in parallel.

	// Unit tests for AddressType function.
	// addressForTest := "Rua Potato"
	// addressTypeExpected := "Rua"

	// addressTypeReceived := AddressType(addressForTest)

	// if addressTypeReceived != addressTypeExpected {
	// 	t.Errorf("Expected %s, but received %s", addressTypeExpected, addressTypeReceived)

	testCaseScenario := []testCase{
		{"Rua Potato", "Rua"},
		{"Avenida Brasil", "Avenida"},
		{"Rodovia", "Rodovia"},
		{"Estrada", "Estrada"},
		{"Alameda", "Alameda"},
		{"", "Unknown address type"},
	}
	for _, scenario := range testCaseScenario {
		addressTypeReceived := AddressType(scenario.addressInserted)
		if addressTypeReceived != scenario.addressExpected {
			t.Errorf("Expected %s, but received %s", scenario.addressExpected, scenario.addressInserted)
		}
	}
}
