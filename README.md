`fullwidth` consumes a line from STDIN, transforms it into its unicode fullwidth counterpart, and prints that to STDOUT.

# Installation

```
go build fullwidth.go
cp fullwidth /usr/local/bin
```

# Usage

`echo "aesthetics" | fullwidth`

# "Useful" bash function

Add the following to your `.bash_profile` or whatever the cool kids are using these days.

`fw () { echo "$@" | fullwidth; }`

then you can

`fw aesthetics`
