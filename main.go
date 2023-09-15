package go_phone_number

import (
	"fmt"
	"regexp"
)

type Country struct {
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
	CountryISO  string `json:"country_iso"`
	CountryName string `json:"country_name"`
	CountryISO3 string `json:"country_iso3"`
}

type ValidationRule struct {
	Pattern         string
	NumberMinLength int
	NumberMaxLength int
}

type PhoneNumber struct {
	CountryCode Country `json:"country_code"`
	Number      string  `json:"number"`
	Rule        ValidationRule
}

func (p PhoneNumber) SetValidationRule(rule ValidationRule) {
	p.Rule = rule
}

func (p PhoneNumber) IsValidMobile() bool {

	// Retrieve country-specific rules for validation
	countryISO := p.CountryCode.CountryISO
	rules, found := countryRules[countryISO]
	if !found {
		// Country not found in rules, consider it invalid
		fmt.Println("Country not found in rules, consider it invalid")
		return false
	}

	// check if the rule is set
	if p.Rule.Pattern != "" {
		rules = p.Rule
	} else {
		// Compile the regular expression for the country's pattern
		re := regexp.MustCompile(rules.Pattern)

		// Check if the phone number matches the pattern
		if !re.MatchString(p.Number) {
			fmt.Println("Phone number does not match the pattern")
			return false
		}
	}

	// Check if the number length matches the country's rules

	phoneNumberLength := len(p.Number)
	if phoneNumberLength <= rules.NumberMinLength || phoneNumberLength >= rules.NumberMaxLength {
		fmt.Println("Phone number length does not match the country's rules")
		return false
	}
	// Additional validation rules can be added here if needed

	return true

}

var countryRules = map[string]ValidationRule{
	"US": {
		Pattern:         `^\d{10}$`,
		NumberMinLength: 10,
		NumberMaxLength: 10,
	},
	"LB": {
		Pattern:         `^(3\d|7[01]|76|79|81|89|90)\d{6}|3\d{6}$`,
		NumberMinLength: 7,
		NumberMaxLength: 12,
	},
	// Add more countries and their rules
}

func NewCountry(countryCode, number, countryISO, countryName, countryISO3 string) Country {
	return Country{
		CountryCode: countryCode,
		Number:      number,
		CountryISO:  countryISO,
		CountryName: countryName,
		CountryISO3: countryISO3,
	}
}

func NewPhoneNumber(country Country, number string) PhoneNumber {
	return PhoneNumber{
		CountryCode: country,
		Number:      number,
	}
}
