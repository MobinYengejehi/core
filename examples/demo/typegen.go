// Code generated by "core generate"; DO NOT EDIT.

package main

import (
	"github.com/MobinYengejehi/core/types"
)

var _ = types.AddType(&types.Type{Name: "main.tableStruct", IDName: "table-struct", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "Icon", Doc: "an icon"}, {Name: "Age", Doc: "an integer field"}, {Name: "Score", Doc: "a float field"}, {Name: "Name", Doc: "a string field"}, {Name: "File", Doc: "a file"}}})

var _ = types.AddType(&types.Type{Name: "main.inlineStruct", IDName: "inline-struct", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "ShowMe", Doc: "this is now showing"}, {Name: "On", Doc: "click to show next"}, {Name: "Condition", Doc: "a condition"}, {Name: "Condition1", Doc: "if On && Condition == 0"}, {Name: "Condition2", Doc: "if On && Condition <= 1"}, {Name: "Value", Doc: "a value"}}})

var _ = types.AddType(&types.Type{Name: "main.testStruct", IDName: "test-struct", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "Enum", Doc: "An enum value"}, {Name: "Name", Doc: "a string"}, {Name: "ShowNext", Doc: "click to show next"}, {Name: "ShowMe", Doc: "this is now showing"}, {Name: "Inline", Doc: "inline struct"}, {Name: "Condition", Doc: "a condition"}, {Name: "Condition1", Doc: "if Condition == 0"}, {Name: "Condition2", Doc: "if Condition >= 0"}, {Name: "Value", Doc: "a value"}, {Name: "Vector", Doc: "a vector"}, {Name: "Table", Doc: "a slice of structs"}, {Name: "List", Doc: "a slice of floats"}, {Name: "File", Doc: "a file"}}})

var _ = types.AddFunc(&types.Func{Name: "main.hello", Doc: "Hello displays a greeting message and an age in weeks based on the given information.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"firstName", "lastName", "age", "likesGo"}, Returns: []string{"greeting", "weeksOld"}})
