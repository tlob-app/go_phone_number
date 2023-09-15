package go_phone_number

import "testing"

func TestPhoneNumber_IsValidMobile(t *testing.T) {
	phone := NewPhoneNumber(
		NewCountry("LB", "", "LB", "", ""),
		"3717171",
	)

	t.Run("is valid phone number test", func(t *testing.T) {
		result := phone.IsValidMobile()
		if result != true {
			t.Errorf("Expected %v, but got %v", "ok", result)
		}
	})

}
