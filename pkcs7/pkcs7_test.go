package pkcs7

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/RobotsAndPencils/buford/certificate"
	// "github.com/cloudflare/cfssl/crypto/pkcs7"
)

func TestSign(t *testing.T) {
	cert, privateKey, err := certificate.Load("../fixtures/cert.p12", "")
	if err != nil {
		t.Fatal(err)
	}

	data := strings.NewReader("something to sign")

	b, err := Sign(data, cert, privateKey, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile("signature-test", b, 0666)
	if err != nil {
		t.Fatal(err)
	}

	// msg, err := pkcs7.ParsePKCS7(b)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(msg)

	// var container container
	// _, err = asn1.Unmarshal(b, &container)
	//
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(container.ContentType)
}
