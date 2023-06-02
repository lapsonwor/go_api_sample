package algoAuth

import (
	//"fmt"
	"log"
	"testing"
)
func TestGeneratePassword(t *testing.T) {
	// Establish the parameters to use for Argon2.
	p := &params{
			memory:      64 * 1024,
			iterations:  3,
			parallelism: 2,
			saltLength:  16,
			keyLength:   32,
	}

	// Pass the plaintext password and parameters to our generateFromPassword
	// helper function.
	hash, err := generateFromPassword("ygjTq89u", p)
	if err != nil {
			log.Fatal(err)
	}

	log.Print(hash)
	//$argon2id$v=19$m=65536,t=3,p=2$XQm5CExh97npygtSdKoVwQ$Za/kIZRHn1X6Qg46J86AnfkFbZ/pF55qZta/euRLyR0

	match, err := ComparePasswordAndHash("ygjTq89u", hash)
	if err != nil {
			log.Fatal(err)
	}

	log.Fatalf("Match: %v\n", match)

}