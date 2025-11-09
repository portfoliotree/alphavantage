# Specification

This directory/package contains JSON files based on the AlphaVantage Documentation.
Eventually I'd like to generate the Go client off of the data in these files.
For now the code in ../query is implemented by hand.

The current specifications are layed out as follows:
- query_parameters.json: many of the parameters are shared across query functions. Specifying them here allows a single source of truth for what a "symbol" is
- function/*.json: each of these files contains a list of function=X specifications and the additional query parameters supported by X
- identifiers.json: maps string literals from the API (column names, query parameter keys or query parameter values) to Go identifiers. The first value is a public identifier the second is a private identifier.


The Go files specification.go and specification_test.go help with structural and referential integrity of the JSON files.
