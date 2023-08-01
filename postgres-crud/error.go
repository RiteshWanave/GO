package main

import ()

// CheckErr checks if there is an error and panics if there is one.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}