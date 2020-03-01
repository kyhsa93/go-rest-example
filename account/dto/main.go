package dto

// Provider account service provider map
func Provider() map[string]string {
	return map[string]string{
		"email": "email",
		"gmail": "gmail",
	}
}

// FilterAccountAttributeByProvider remove socialID or password by provider
func FilterAccountAttributeByProvider(data *Account) {
	if data.Provider == Provider()["email"] {
		data.SocialID = ""
		return
	}
	data.Password = ""
	return
}

// ValidateAccountAttributeByProvider validate account attribute by provider
func ValidateAccountAttributeByProvider(data *Account) bool {
	if data.Provider == Provider()["email"] && data.Password == "" {
		return false
	}
	if data.Provider != Provider()["email"] && data.SocialID == "" {
		return false
	}
	return true
}

// Gender user gender map
func Gender() map[string]string {
	return map[string]string{
		"male":   "male",
		"female": "female",
	}
}

// ValidateAccountGender validation account gender attribute
func ValidateAccountGender(data *Account) bool {
	if Gender()[data.Gender] == "" {
		return false
	}
	return true
}

// InterestedField part of user InterestedFielded
func InterestedField() map[string]string {
	return map[string]string{
		"develop": "develop",
		"design":  "design",
		"manage":  "manage",
	}
}

// ValidateInterestedFieldAttribute validation account's InterestedField
func ValidateInterestedFieldAttribute(data *Account) bool {
	if InterestedField()[data.InterestedField] == "" {
		return false
	}
	return true
}

// Account account dto for command action
type Account struct {
	Email                 string   `json:"email" example:"test@gmail.com" binding:"required"`
	Provider              string   `json:"provider" example:"gmail" binding:"required"`
	SocialID              string   `json:"social_id" example:"social_id"`
	Password              string   `json:"password" example:"password"`
	Gender                string   `json:"gender" example:"male"`
	InterestedField       string   `json:"interested_field" example:"develop"`
	InterestedFieldDetail []string `json:"interested_field_detail" example:"web,server"`
}
