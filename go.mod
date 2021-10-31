module github.com/microiam/did-common-go

require github.com/elimity-com/abnf v0.1.0 // indirect

go 1.16

require (
	did v0.0.1
)

replace (
	did v0.0.1 => ./did
)