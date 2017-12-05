// +build gofuzz

package publicsuffix

import (
	"fmt"

	psl "golang.org/x/net/publicsuffix"
)

func Fuzz(in []byte) int {
	var domain = string(in)

	var got, _ = PublicSuffix(domain)
	var want, _ = psl.PublicSuffix(domain)
	if want != got {
		panic(fmt.Sprintf("output mismatch: got %q, want %q (%v)\n", got, want, domain))
	}

	var wantErr error
	want, wantErr = psl.EffectiveTLDPlusOne(domain)

	var err error
	got, err = EffectiveTLDPlusOne(domain)
	if want != got {
		panic(fmt.Sprintf("output mismatch: TLD got %q, want %q (%v)\n", got, want, domain))
	}

	// Compare if an error exists, not the value of it
	if (err == nil) != (wantErr == nil) {
		panic(fmt.Sprintf("error mismatch: got err %q, want %q (%v)\n", err, wantErr, domain))
	}

	if err != nil {
		return -1
	}

	return 1
}
