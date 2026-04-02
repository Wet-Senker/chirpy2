package auth

import "testing"

func TestHashPasswordHello(t *testing.T) {
	hash, err := HashPassword("Hello")
	bool, err := CheckPasswordHash("Hello", hash)
	
	if !bool || err != nil{
		t.Errorf("Password %s, does not match hash %s", "Hello", hash)
	}

}

func TestHashPasswordsymbols(t *testing.T) {
	hash, err := HashPassword("H$e%llo")
	bool, err := CheckPasswordHash("H$e%llo", hash)
	
	if !bool || err != nil{
		t.Errorf("Password %s, does not match hash %s", "H$e%llo", hash)
	}

}