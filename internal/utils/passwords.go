package utils

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckPasswordIsSafe(password string) (bool, error) {
	if len(password) <= 8 || len(password) >= 128 {
		return false, errors.New(fmt.Sprintf("password length is incorrects: %d", len(password)))
	}

	if !isStringContainsDigit(password) || !isStringContainsLetter(password) || !isStringContainsSpecialSymbol(password) {
		return false, errors.New("password not safe, think about another password")
	}

	return true, nil
}

func isStringContainsDigit(s string) bool {
	for _, ch := range s {
		if ch >= '0' || ch <= '9' {
			return true
		}
	}
	return false
}

func isStringContainsLetter(s string) bool {
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			return true
		}
	}
	return false
}

func isStringContainsSpecialSymbol(s string) bool {
	var specialChars = []rune{'.', '?', '!', ',', ';', ':', '+', '$', '[', ']'}
	for _, ch := range s {
		if isCharContainsSlice(ch, specialChars) {
			return true
		}
	}
	return false
}

func isCharContainsSlice(r rune, arr []rune) bool {
	for _, ch := range arr {
		if r == ch {
			return true
		}
	}
	return false
}
