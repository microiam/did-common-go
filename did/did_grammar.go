// This file is generated - do not edit.

package did

import (
	"github.com/elimity-com/abnf/core"

	"github.com/elimity-com/abnf/operators"
)

// did = "did:" method-name ":" method-specific-id
func Did(s []byte) operators.Alternatives {
	return operators.Concat(
		"did",
		operators.String("did:", "did:"),
		MethodName,
		operators.String(":", ":"),
		MethodSpecificId,
	)(s)
}

// idchar = ALPHA / DIGIT / "." / "-" / "_" / pct-encoded
func Idchar(s []byte) operators.Alternatives {
	return operators.Alts(
		"idchar",
		core.ALPHA(),
		core.DIGIT(),
		operators.String(".", "."),
		operators.String("-", "-"),
		operators.String("_", "_"),
		PctEncoded,
	)(s)
}

// method-char = %x61-7A / DIGIT
func MethodChar(s []byte) operators.Alternatives {
	return operators.Alts(
		"method-char",
		operators.Range("%x61-7A", []byte{97}, []byte{122}),
		core.DIGIT(),
	)(s)
}

// method-name = 1*method-char
func MethodName(s []byte) operators.Alternatives {
	return operators.Repeat1Inf("method-name", MethodChar)(s)
}

// method-specific-id = *( *idchar ":" ) 1*idchar
func MethodSpecificId(s []byte) operators.Alternatives {
	return operators.Concat(
		"method-specific-id",
		operators.Repeat0Inf("*( *idchar \":\" )", operators.Concat(
			"*idchar \":\"",
			operators.Repeat0Inf("*idchar", Idchar),
			operators.String(":", ":"),
		)),
		operators.Repeat1Inf("1*idchar", Idchar),
	)(s)
}

// pct-encoded = "%" HEXDIG HEXDIG
func PctEncoded(s []byte) operators.Alternatives {
	return operators.Concat(
		"pct-encoded",
		operators.String("%", "%"),
		core.HEXDIG(),
		core.HEXDIG(),
	)(s)
}
