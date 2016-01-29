package pkcs7

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"golang.org/x/crypto/pkcs12"
)

func TestSign(t *testing.T) {
	p12, err := ioutil.ReadFile("../cert-website.p12")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, cert, err := pkcs12.Decode(p12, "")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("manifest.json")
	if err != nil {
		t.Fatal("Error opening manifest:", err)
	}
	defer f.Close()
	data, err := Sign(f, cert, privateKey.(*rsa.PrivateKey))
	if err != nil {
		t.Fatal("Error signing manifest:", err)
	}
	if err := ioutil.WriteFile("signature-test", data, 0666); err != nil {
		t.Fatal("Error writing signature:", err)
	}
}
