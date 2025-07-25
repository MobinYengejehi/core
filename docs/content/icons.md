+++
Categories = ["Resources"]
+++

Cogent Core provides more than 2,000 **icons** from [Material Design Symbols](https://fonts.google.com/icons), sourced through [marella/material-symbols](https://github.com/marella/material-symbols). See [[icon]] for information about the icon widget and how to use icons.

Icons are represented directly as an SVG string. To reduce binary size, only the icons your app actually uses are included.

## Custom icons

Because an icon is just an SVG string, it is easy to add custom icons. If you just have one or two, you can manually embed the SVG data:

```go
import (
	_ "embed"
)

//go:embed myIcon.svg
var myIcon string
```

And then use it like this:

```go
core.NewButton(b).SetIcon(icons.Icon(myIcon))
```

### Icongen

If you have more icons, you can use icongen, a part of the [[generate]] tool. Custom icons are typically placed in a `cicons` (custom icons) directory. In it, you can add all of your custom SVG icon files and an `icons.go` file with the following code:

```go
package cicons

//go:generate core generate -icons .
```

Then, once you run `go generate`, you can access your icons through your cicons package, where icon names are automatically transformed into CamelCase:

```go
core.NewButton(b).SetIcon(cicons.MyIconName)
```

### Image icons

Although only SVG files are supported for icons, you can easily embed a bitmap image file in an SVG file. Cogent Core provides an `svg` command line tool that can do this for you. To install it, run:

```sh
go install github.com/MobinYengejehi/core/svg/cmd/svg@main
```

Then, to embed an image into an svg file, run:

```sh
svg embed-image my-image.png
```

This will create a file called `my-image.svg` that has the image embedded into it. Then, you can use that SVG file as an icon as described above.

