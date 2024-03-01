// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package insecureseed represents the interface "wasi:random/insecure-seed@0.2.0".
//
// The insecure-seed interface for seeding hash-map DoS resistance.
//
// It is intended to be portable at least between Unix-family platforms and
// Windows.
package insecureseed

// InsecureSeed represents function "wasi:random/insecure-seed@0.2.0#insecure-seed".
//
// Return a 128-bit value that may contain a pseudo-random value.
//
// The returned value is not required to be computed from a CSPRNG, and may
// even be entirely deterministic. Host implementations are encouraged to
// provide pseudo-random values to any program exposed to
// attacker-controlled content, to enable DoS protection built into many
// languages' hash-map implementations.
//
// This function is intended to only be called once, by a source language
// to initialize Denial Of Service (DoS) protection in its hash-map
// implementation.
//
// # Expected future evolution
//
// This will likely be changed to a value import, to prevent it from being
// called multiple times and potentially used for purposes other than DoS
// protection.
//
//	insecure-seed: func() -> tuple<u64, u64>
//
//go:nosplit
func InsecureSeed() [2]uint64 {
	var result [2]uint64
	insecureSeed(&result)
	return result
}

//go:wasmimport wasi:random/insecure-seed@0.2.0 insecure-seed
//go:noescape
func insecureSeed(result *[2]uint64)
