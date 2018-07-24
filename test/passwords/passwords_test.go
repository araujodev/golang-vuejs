package password

import (
	"log"
	"testing"

	"github.com/araujodev/golang-vuejs/src/system/passwords"
)

func init() {
	log.Println("Testing Password")
}

func TestBasicLog(t *testing.T) {
	pass := "TEST"
	hash, err := passwords.Encrypt(pass)
	if err != nil {
		t.Error(err.Error())
	}
	log.Println(hash)
	ok := passwords.IsValid(hash, pass)
	if !ok {
		t.Error("Password not Matching... hasing is not working")
	}
	log.Println("Success test hashing...")

}
