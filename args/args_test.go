package args

import "testing"

func TestParseKV(t *testing.T) {
	good := []string{"hello=world!", "key= value", "this =that", " space = around "}
	bad := []string{"hello=world=invalid", "invalid", ""}

	_, err := ParseKV(good, "good")
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
  _, err = ParseKV(bad, "bad")
  if err == nil {
    t.Error("Error: expected to throw error!\n")
  }
}
