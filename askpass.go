package askpass

import (
	"fmt"
	"os"
)

// Pass is a struct that represents a password file
type Pass struct {
	Filename string
}

// Save saves a password in a file
func (p *Pass) Save(password string) error {
	if password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	err := os.WriteFile(p.Filename, []byte(password), 0600)
	if err != nil {
		return fmt.Errorf("failed to save password in file %s: %w", p.Filename, err)
	}

	return nil
}

// Get retrieves a password from a file
func (p *Pass) Get() (string, error) {
	password, err := os.ReadFile(p.Filename)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("password file does not exist: %w", err)
		}
		return "", fmt.Errorf("failed to read password from file %s: %w", p.Filename, err)
	}

	return string(password), nil
}
