// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmdtest

import (
	"testing"

	"golang.org/x/tools/gopls/internal/lsp/protocol"
	"golang.org/x/tools/gopls/internal/lsp/tests/compare"
	"golang.org/x/tools/gopls/internal/span"
)

func (r *runner) Symbols(t *testing.T, uri span.URI, expectedSymbols []protocol.DocumentSymbol) {
	filename := uri.Filename()
	got, _ := r.NormalizeGoplsCmd(t, "symbols", filename)
	expect := string(r.data.Golden(t, "symbols", filename, func() ([]byte, error) {
		return []byte(got), nil
	}))
	if diff := compare.Text(expect, got); diff != "" {
		t.Errorf("symbols differ from expected:\n%s", diff)
	}
}
