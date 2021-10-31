package did

// The DID struct contains all the info that can be extracted
// from the DID string.
type DID struct {
	didString            string
	methodName           string
	methodSpecificID     string
}

// Initialize the parser.
func (p *DID) initialize() {
	p.didString = ""
	p.methodName = ""
	p.methodSpecificID = ""
}

// New parses the given DID string and get the resulting DID object.
func New(did string) *DID {
	o := &DID{}
	o.Parse(did)
	return o
}

// Parse the given DID string.
func (p *DID) Parse(did string) {
	p.initialize()

	tree := Did([]byte(did)).Best()
	p.didString = tree.String()
	p.methodName = tree.GetSubNode("method-name").String()
	p.methodSpecificID = tree.GetSubNode("method-specific-id").String()
}
