// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package golang

import (
	"github.com/MobinYengejehi/core/text/parse/syms"
	"github.com/MobinYengejehi/core/text/token"
)

// FuncParams returns the parameters of given function / method symbol,
// in proper order, name type for each param space separated in string
func (gl *GoLang) FuncParams(fsym *syms.Symbol) []string {
	var ps []string
	for _, cs := range fsym.Children {
		if cs.Kind != token.NameVarParam {
			continue
		}
		if len(ps) <= cs.Index {
			op := ps
			ps = make([]string, cs.Index+1)
			copy(ps, op)
		}
		s := cs.Name + " " + cs.Type
		ps[cs.Index] = s
	}
	return ps
}
