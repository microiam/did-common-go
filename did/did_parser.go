package did

// DID parser
type parser struct {
	input        string // input to the parser
	output       *DID   // output DID
	err          error  // an error during parsing
}
