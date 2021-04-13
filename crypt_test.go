package mycrypt

import (
	"strings"
	"testing"
)

func TestEncrypt(t *testing.T) {
	t.Run("Testing Encryption and Decryption", func(t *testing.T) {
		shortKey := "ThisIsShortKey"
		SetKey(shortKey)
		want := "ThisIsShortKeyThisIsShortKeyThis"
		if key != want {
			t.Errorf("SetKey(%s) = %v, want %v The short key should have been expanded to 32 characters", shortKey, key, want)
		}

		longKey := "ThisIsLongKey01234567890123456789012345678901234567890123456789"
		SetKey(longKey)
		want = "ThisIsLongKey0123456789012345678"
		if key != want {
			t.Errorf("SetKey(%s) = %v, want %v The long key should have been reduced to 32 characters", longKey, key, want)
		}

		plainText := "This is a long string that I will encrypt 01234567890 !@#$%^&*("
		encryptedText := ":e:O3AGC56hxDpyWIQJDHo2AGczOPZ+c6AMnG7evLTxizxnp2Li5g5YgDbtMtJmu4lDtjw9ncttxygei6KPhWLBxezNpY6x9DOrg+kolzthbGd36JWd/bvSw0qF/xq93Jay9equKw=="

		got := De(encryptedText)
		if got != plainText {
			t.Errorf("Decrypt() = %v, want %v", got, plainText)
		}

		// This test if the decrypt is ignoring because the meta is missing
		got = De(plainText)
		if got != plainText {
			t.Errorf("Decrypt() = %v, want %v", got, plainText)
		}

		got = En(plainText)
		if !strings.HasPrefix(got, meta) {
			t.Errorf("Encrypt() = %v, want it to start with %v", got, meta)
		}

	})
}
