// Code generated by 'yaegi extract github.com/MobinYengejehi/core/colors/gradient'. DO NOT EDIT.

package coresymbols

import (
	"github.com/MobinYengejehi/core/colors/gradient"
	"github.com/MobinYengejehi/core/math32"
	"image"
	"image/color"
	"reflect"
)

func init() {
	Symbols["github.com/MobinYengejehi/core/colors/gradient/gradient"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Apply":             reflect.ValueOf(gradient.Apply),
		"ApplyOpacity":      reflect.ValueOf(gradient.ApplyOpacity),
		"Cache":             reflect.ValueOf(&gradient.Cache).Elem(),
		"CopyFrom":          reflect.ValueOf(gradient.CopyFrom),
		"CopyOf":            reflect.ValueOf(gradient.CopyOf),
		"FromAny":           reflect.ValueOf(gradient.FromAny),
		"FromString":        reflect.ValueOf(gradient.FromString),
		"NewApplier":        reflect.ValueOf(gradient.NewApplier),
		"NewBase":           reflect.ValueOf(gradient.NewBase),
		"NewLinear":         reflect.ValueOf(gradient.NewLinear),
		"NewRadial":         reflect.ValueOf(gradient.NewRadial),
		"ObjectBoundingBox": reflect.ValueOf(gradient.ObjectBoundingBox),
		"Pad":               reflect.ValueOf(gradient.Pad),
		"ReadXML":           reflect.ValueOf(gradient.ReadXML),
		"Reflect":           reflect.ValueOf(gradient.Reflect),
		"Repeat":            reflect.ValueOf(gradient.Repeat),
		"SpreadsN":          reflect.ValueOf(gradient.SpreadsN),
		"SpreadsValues":     reflect.ValueOf(gradient.SpreadsValues),
		"UnitsN":            reflect.ValueOf(gradient.UnitsN),
		"UnitsValues":       reflect.ValueOf(gradient.UnitsValues),
		"UnmarshalXML":      reflect.ValueOf(gradient.UnmarshalXML),
		"UserSpaceOnUse":    reflect.ValueOf(gradient.UserSpaceOnUse),
		"XMLAttr":           reflect.ValueOf(gradient.XMLAttr),

		// type definitions
		"Applier":    reflect.ValueOf((*gradient.Applier)(nil)),
		"ApplyFunc":  reflect.ValueOf((*gradient.ApplyFunc)(nil)),
		"ApplyFuncs": reflect.ValueOf((*gradient.ApplyFuncs)(nil)),
		"Base":       reflect.ValueOf((*gradient.Base)(nil)),
		"Gradient":   reflect.ValueOf((*gradient.Gradient)(nil)),
		"Linear":     reflect.ValueOf((*gradient.Linear)(nil)),
		"Radial":     reflect.ValueOf((*gradient.Radial)(nil)),
		"Spreads":    reflect.ValueOf((*gradient.Spreads)(nil)),
		"Stop":       reflect.ValueOf((*gradient.Stop)(nil)),
		"Units":      reflect.ValueOf((*gradient.Units)(nil)),

		// interface wrapper definitions
		"_Gradient": reflect.ValueOf((*_cogentcore_org_core_colors_gradient_Gradient)(nil)),
	}
}

// _cogentcore_org_core_colors_gradient_Gradient is an interface wrapper for Gradient type
type _cogentcore_org_core_colors_gradient_Gradient struct {
	IValue      interface{}
	WAsBase     func() *gradient.Base
	WAt         func(x int, y int) color.Color
	WBounds     func() image.Rectangle
	WColorModel func() color.Model
	WUpdate     func(opacity float32, box math32.Box2, objTransform math32.Matrix2)
}

func (W _cogentcore_org_core_colors_gradient_Gradient) AsBase() *gradient.Base { return W.WAsBase() }
func (W _cogentcore_org_core_colors_gradient_Gradient) At(x int, y int) color.Color {
	return W.WAt(x, y)
}
func (W _cogentcore_org_core_colors_gradient_Gradient) Bounds() image.Rectangle { return W.WBounds() }
func (W _cogentcore_org_core_colors_gradient_Gradient) ColorModel() color.Model {
	return W.WColorModel()
}
func (W _cogentcore_org_core_colors_gradient_Gradient) Update(opacity float32, box math32.Box2, objTransform math32.Matrix2) {
	W.WUpdate(opacity, box, objTransform)
}
