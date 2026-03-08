module adtech.com/go_monorepo

go 1.25.7

// ANCHOR - direct dependencies
require rsc.io/quote v1.5.2

// ANCHOR - dependencies replaces
// require adtech.com/utils v0.0.0-unpublished
// replace adtech.com/utils v0.0.0-unpublished => ./modules

// ANCHOR - indirect dependencies
require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)
