// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on https://github.com/hack-pad/hackpad
// Licensed under the Apache 2.0 License

//go:build js

package main

import (
	"syscall/js"

	"github.com/MobinYengejehi/core/base/errors"
	"github.com/MobinYengejehi/core/system/driver/web/jsfs"
)

func main() {
	fs := errors.Must1(jsfs.Config(js.Global().Get("fs")))
	errors.Must1(fs.MkdirAll([]js.Value{js.ValueOf("me"), js.ValueOf(0777)}))
	callback := js.FuncOf(func(this js.Value, args []js.Value) any {
		js.Global().Get("console").Call("log", "stat file info", args[1])
		return nil
	})
	js.Global().Get("fs").Call("stat", "me", callback)
	select {}
}
