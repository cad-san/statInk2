package statink2

import "testing"

func TestNewClient(t *testing.T) {
	dummyKey := "APIKEY"
	_, err := NewClient(dummyKey, nil)
	if err != nil {
		t.Errorf("NewClient() should not be error with %v", err)
	}
}
