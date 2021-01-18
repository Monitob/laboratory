# polymere
`import "./pkg/polymere/"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package polymere implement functions to manipulate polymere the methds react to a polymere

	A polymere is composed by letters in alphabet,
	minuscule (consider as negative) or Capital (considered as positive).

## <a name="pkg-imports">Imported Packages</a>

No packages beyond the Go standard library are imported.

## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func LenMinimalReaction(polymer string) (int, error)](#LenMinimalReaction)
* [func React(polygrame string) (string, error)](#React)

#### <a name="pkg-files">Package files</a>
[polymere.go](./polymere.go) 

## <a name="pkg-constants">Constants</a>
``` go
const (
    // ElementError the "error" if not letter
    ElementError = "arguments should be only letters"
)
```

## <a name="LenMinimalReaction">func</a> [LenMinimalReaction](./polymere.go#L100)
``` go
func LenMinimalReaction(polymer string) (int, error)
```
LenMinimalReaction return the length of the smallest reaction removing one element

## <a name="React">func</a> [React](./polymere.go#L34)
``` go
func React(polygrame string) (string, error)
```
React method to acquire the reaction between opposite polarisation