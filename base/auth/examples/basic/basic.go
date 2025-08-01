// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/MobinYengejehi/core/base/auth"
	"github.com/MobinYengejehi/core/base/errors"
	"github.com/MobinYengejehi/core/core"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

func main() {
	b := core.NewBody("Auth basic example")
	fun := func(token *oauth2.Token, userInfo *oidc.UserInfo) {
		d := core.NewBody("User info")
		core.NewText(d).SetType(core.TextHeadlineMedium).SetText("Basic info")
		core.NewForm(d).SetStruct(userInfo)
		core.NewText(d).SetType(core.TextHeadlineMedium).SetText("Detailed info")
		claims := map[string]any{}
		errors.Log(userInfo.Claims(&claims))
		core.NewKeyedList(d).SetMap(&claims)
		d.AddOKOnly().RunFullDialog(b)
	}
	auth.Buttons(b, &auth.ButtonsConfig{SuccessFunc: fun})
	b.RunMainWindow()
}
