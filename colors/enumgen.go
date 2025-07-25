// Code generated by "core generate"; DO NOT EDIT.

package colors

import (
	"github.com/MobinYengejehi/core/enums"
)

var _BlendTypesValues = []BlendTypes{0, 1, 2}

// BlendTypesN is the highest valid value for type BlendTypes, plus one.
const BlendTypesN BlendTypes = 3

var _BlendTypesValueMap = map[string]BlendTypes{`HCT`: 0, `RGB`: 1, `CAM16`: 2}

var _BlendTypesDescMap = map[BlendTypes]string{0: `HCT uses the hue, chroma, and tone space and generally produces the best results, but at a slight performance cost.`, 1: `RGB uses raw RGB space, which is the standard space that most other programs use. It produces decent results with maximum performance.`, 2: `CAM16 is an alternative colorspace, similar to HCT, but not quite as good.`}

var _BlendTypesMap = map[BlendTypes]string{0: `HCT`, 1: `RGB`, 2: `CAM16`}

// String returns the string representation of this BlendTypes value.
func (i BlendTypes) String() string { return enums.String(i, _BlendTypesMap) }

// SetString sets the BlendTypes value from its string representation,
// and returns an error if the string is invalid.
func (i *BlendTypes) SetString(s string) error {
	return enums.SetString(i, s, _BlendTypesValueMap, "BlendTypes")
}

// Int64 returns the BlendTypes value as an int64.
func (i BlendTypes) Int64() int64 { return int64(i) }

// SetInt64 sets the BlendTypes value from an int64.
func (i *BlendTypes) SetInt64(in int64) { *i = BlendTypes(in) }

// Desc returns the description of the BlendTypes value.
func (i BlendTypes) Desc() string { return enums.Desc(i, _BlendTypesDescMap) }

// BlendTypesValues returns all possible values for the type BlendTypes.
func BlendTypesValues() []BlendTypes { return _BlendTypesValues }

// Values returns all possible values for the type BlendTypes.
func (i BlendTypes) Values() []enums.Enum { return enums.Values(_BlendTypesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i BlendTypes) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *BlendTypes) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "BlendTypes")
}
