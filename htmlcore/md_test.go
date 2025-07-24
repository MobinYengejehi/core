// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package htmlcore

import (
	"testing"

	"github.com/MobinYengejehi/core/core"
	"github.com/stretchr/testify/assert"
)

func TestMD(t *testing.T) {
	tests := map[string]string{
		"h1":   `# Test`,
		"h2":   `## Test`,
		"p":    `Test`,
		`code`: "```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}\n```",
	}
	for nm, s := range tests {
		b := core.NewBody()
		assert.NoError(t, ReadMDString(NewContext(), b, s))
		b.AssertRender(t, "md/"+nm)
	}
}

func TestDoubleButtonBug(t *testing.T) {
	b := core.NewBody()
	assert.NoError(t, ReadMDString(NewContext(), b, `<button>A</button><button>B</button>`))
	b.AssertRender(t, "md/double-button-bug")
}

func TestExtraNewlineBug(t *testing.T) {
	b := core.NewBody()
	assert.NoError(t, ReadMDString(NewContext(), b, `A

<button>B</button>`))
	b.AssertRender(t, "md/extra-newline-bug")
}

func TestButtonInHeadingBug(t *testing.T) {
	// TODO: this does not work correctly due to a minor bug in [extractText]
	// (see the comment there).
	b := core.NewBody()
	assert.NoError(t, ReadMDString(NewContext(), b, `<h1><button>A</button>B</h1>`))
	b.AssertRender(t, "md/button-in-heading-bug")
}
