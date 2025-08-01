// Code generated by "core generate"; DO NOT EDIT.

package sides

import (
	"github.com/MobinYengejehi/core/types"
)

var _ = types.AddType(&types.Type{Name: "github.com/MobinYengejehi/core/styles/sides.Sides", IDName: "sides", Doc: "Sides contains values for each side or corner of a box.\nIf Sides contains sides, the struct field names correspond\ndirectly to the side values (ie: Top = top side value).\nIf Sides contains corners, the struct field names correspond\nto the corners as follows: Top = top left, Right = top right,\nBottom = bottom right, Left = bottom left.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "Top", Doc: "top/top-left value"}, {Name: "Right", Doc: "right/top-right value"}, {Name: "Bottom", Doc: "bottom/bottom-right value"}, {Name: "Left", Doc: "left/bottom-left value"}}})

var _ = types.AddType(&types.Type{Name: "github.com/MobinYengejehi/core/styles/sides.Values", IDName: "values", Doc: "Values contains units.Value values for each side/corner of a box", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Embeds: []types.Field{{Name: "Sides"}}})

var _ = types.AddType(&types.Type{Name: "github.com/MobinYengejehi/core/styles/sides.Floats", IDName: "floats", Doc: "Floats contains float32 values for each side/corner of a box", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Embeds: []types.Field{{Name: "Sides"}}})

var _ = types.AddType(&types.Type{Name: "github.com/MobinYengejehi/core/styles/sides.Colors", IDName: "colors", Doc: "Colors contains color values for each side/corner of a box", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Embeds: []types.Field{{Name: "Sides"}}})
