package password

import "golang.org/x/crypto/bcrypt"

// IPassword is
type IPassword interface {
	GenerateHashPassword(plainPassword string) string
	IsValidPassword(plainPassword, hashedPassword string) bool
}

// Password is
type Password struct{}

func NewPassword() *Password {
	return &Password{}
}

// GenerateHashPassword is
func (p *Password) GenerateHashPassword(plainPassword string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	return string(hashedPassword)
}

// IsValidPassword is
func (p *Password) IsValidPassword(plainPassword, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}
