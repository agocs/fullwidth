`fullwidth` consumes a line from STDIN, transforms it into its unicode fullwidth counterpart, and prints that to STDOUT.

# Installation

`go build fullwidth.go`
`cp fullwidth /usr/local/bin`

# Usage

`echo "aesthetics" | fullwidth`

# "Useful" bash function

`fw () { echo "$@" | fullwidth; }`