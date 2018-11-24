package passwordvalidator

import (
	"fmt"
	"testing"
)

var commonTests = []struct {
	password string
	result   bool
}{
	{"password", false},
	{"iloveyou", false},
	{"password1", false},
	{"QtbiL3dZ28JXT55E", true},
}

func TestCommonPassword(t *testing.T) {
	for _, tt := range commonTests {
		x := CommonPassword(tt.password)
		t.Run(tt.password, func(t *testing.T) {
			if x != tt.result {
				t.Errorf("got %v, want %v", x, tt.result)
			}
		})
	}
}

var similarityTests = []struct {
	password string
	attr     string
	result   bool
}{
	{"password", "password@example.com", false},
	{"iloveyou", "iloveyou2@example.com", false},
	{"QtbiL3dZ28JXT55E", "my_username", true},
}

func TestSimilarity(t *testing.T) {
	for _, tt := range similarityTests {
		x := Similarity(tt.password, tt.attr)
		t.Run(fmt.Sprintf("%s,%s", tt.password, tt.attr), func(t *testing.T) {
			if x != tt.result {
				t.Errorf("got %v, want %v", x, tt.result)
			}
		})
	}
}
