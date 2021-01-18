# cmd
`import "./cmd/"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package cmd handles the command-line interface for laboratory

## <a name="pkg-imports">Imported Packages</a>

- github.com/spf13/cobra
- gitlab.nereides.weborama.com/go/playground/Test/laboratory/pkg/polymere

## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)

#### <a name="pkg-files">Package files</a>
[minlen.go](./minlen.go) [react.go](./react.go) [root.go](./root.go) 

## <a name="pkg-variables">Variables</a>
``` go
var RootCmd = &cobra.Command{
    Use:           "laboratory",
    Short:         "abBc",
    Long:          `Currently, just write your polymere.`,
    SilenceErrors: true,
}
```
RootCmd is the root for all hello commands.
