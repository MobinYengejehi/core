// Code generated by "core generate"; DO NOT EDIT.

package gpu

import (
	"github.com/MobinYengejehi/core/enums"
)

var _VarRolesValues = []VarRoles{0, 1, 2, 3, 4, 5, 6, 7}

// VarRolesN is the highest valid value for type VarRoles, plus one.
const VarRolesN VarRoles = 8

var _VarRolesValueMap = map[string]VarRoles{`UndefVarRole`: 0, `Vertex`: 1, `Index`: 2, `Push`: 3, `Uniform`: 4, `Storage`: 5, `StorageTexture`: 6, `SampledTexture`: 7}

var _VarRolesDescMap = map[VarRoles]string{0: ``, 1: `Vertex is vertex shader input data: mesh geometry points, normals, etc. These are automatically located in a separate Set, VertexSet (-2), and managed separately.`, 2: `Index is for indexes to access to Vertex data, also located in VertexSet (-2). Only one such Var per VarGroup should be present, and will automatically be used if a value is set.`, 3: `Push is push constants, NOT CURRENTLY SUPPORTED in WebGPU. They have a minimum of 128 bytes and are stored directly in the command buffer. They do not require any host-device synchronization or buffering, and are fully dynamic. They are ideal for transformation matricies or indexes for accessing data. They are stored in a special PushSet (-1) and managed separately.`, 4: `Uniform is read-only general purpose data, with a more limited capacity. Compared to Storage, Uniform items can be put in local cache for each shader and thus can be much faster to access. Use for a smaller number of parameters such as transformation matricies.`, 5: `Storage is read-write general purpose data. This is a larger but slower pool of memory, with more flexible alignment constraints, used primarily for compute data.`, 6: `StorageTexture is read-write storage-based texture data, for compute shaders that operate on image data, not for standard use of textures in fragment shader to texturize objects (which is SampledTexture).`, 7: `SampledTexture is a Texture + Sampler that is used to texturize objects in the fragment shader. The variable for this specifies the role for the texture (typically there is just one main such texture), and the different Values of the variable hold each instance, with binding used to switch which texture to use. The Texture takes the first Binding position, and the Sampler is +1.`}

var _VarRolesMap = map[VarRoles]string{0: `UndefVarRole`, 1: `Vertex`, 2: `Index`, 3: `Push`, 4: `Uniform`, 5: `Storage`, 6: `StorageTexture`, 7: `SampledTexture`}

// String returns the string representation of this VarRoles value.
func (i VarRoles) String() string { return enums.String(i, _VarRolesMap) }

// SetString sets the VarRoles value from its string representation,
// and returns an error if the string is invalid.
func (i *VarRoles) SetString(s string) error {
	return enums.SetString(i, s, _VarRolesValueMap, "VarRoles")
}

// Int64 returns the VarRoles value as an int64.
func (i VarRoles) Int64() int64 { return int64(i) }

// SetInt64 sets the VarRoles value from an int64.
func (i *VarRoles) SetInt64(in int64) { *i = VarRoles(in) }

// Desc returns the description of the VarRoles value.
func (i VarRoles) Desc() string { return enums.Desc(i, _VarRolesDescMap) }

// VarRolesValues returns all possible values for the type VarRoles.
func VarRolesValues() []VarRoles { return _VarRolesValues }

// Values returns all possible values for the type VarRoles.
func (i VarRoles) Values() []enums.Enum { return enums.Values(_VarRolesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i VarRoles) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *VarRoles) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "VarRoles") }

var _SamplerModesValues = []SamplerModes{0, 1, 2}

// SamplerModesN is the highest valid value for type SamplerModes, plus one.
const SamplerModesN SamplerModes = 3

var _SamplerModesValueMap = map[string]SamplerModes{`Repeat`: 0, `MirrorRepeat`: 1, `ClampToEdge`: 2}

var _SamplerModesDescMap = map[SamplerModes]string{0: `Repeat the texture when going beyond the image dimensions.`, 1: `Like repeat, but inverts the coordinates to mirror the image when going beyond the dimensions.`, 2: `Take the color of the edge closest to the coordinate beyond the image dimensions.`}

var _SamplerModesMap = map[SamplerModes]string{0: `Repeat`, 1: `MirrorRepeat`, 2: `ClampToEdge`}

// String returns the string representation of this SamplerModes value.
func (i SamplerModes) String() string { return enums.String(i, _SamplerModesMap) }

// SetString sets the SamplerModes value from its string representation,
// and returns an error if the string is invalid.
func (i *SamplerModes) SetString(s string) error {
	return enums.SetString(i, s, _SamplerModesValueMap, "SamplerModes")
}

// Int64 returns the SamplerModes value as an int64.
func (i SamplerModes) Int64() int64 { return int64(i) }

// SetInt64 sets the SamplerModes value from an int64.
func (i *SamplerModes) SetInt64(in int64) { *i = SamplerModes(in) }

// Desc returns the description of the SamplerModes value.
func (i SamplerModes) Desc() string { return enums.Desc(i, _SamplerModesDescMap) }

// SamplerModesValues returns all possible values for the type SamplerModes.
func SamplerModesValues() []SamplerModes { return _SamplerModesValues }

// Values returns all possible values for the type SamplerModes.
func (i SamplerModes) Values() []enums.Enum { return enums.Values(_SamplerModesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i SamplerModes) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *SamplerModes) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "SamplerModes")
}

var _BorderColorsValues = []BorderColors{0, 1, 2}

// BorderColorsN is the highest valid value for type BorderColors, plus one.
const BorderColorsN BorderColors = 3

var _BorderColorsValueMap = map[string]BorderColors{`Trans`: 0, `Black`: 1, `White`: 2}

var _BorderColorsDescMap = map[BorderColors]string{0: `Repeat the texture when going beyond the image dimensions.`, 1: ``, 2: ``}

var _BorderColorsMap = map[BorderColors]string{0: `Trans`, 1: `Black`, 2: `White`}

// String returns the string representation of this BorderColors value.
func (i BorderColors) String() string { return enums.String(i, _BorderColorsMap) }

// SetString sets the BorderColors value from its string representation,
// and returns an error if the string is invalid.
func (i *BorderColors) SetString(s string) error {
	return enums.SetString(i, s, _BorderColorsValueMap, "BorderColors")
}

// Int64 returns the BorderColors value as an int64.
func (i BorderColors) Int64() int64 { return int64(i) }

// SetInt64 sets the BorderColors value from an int64.
func (i *BorderColors) SetInt64(in int64) { *i = BorderColors(in) }

// Desc returns the description of the BorderColors value.
func (i BorderColors) Desc() string { return enums.Desc(i, _BorderColorsDescMap) }

// BorderColorsValues returns all possible values for the type BorderColors.
func BorderColorsValues() []BorderColors { return _BorderColorsValues }

// Values returns all possible values for the type BorderColors.
func (i BorderColors) Values() []enums.Enum { return enums.Values(_BorderColorsValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i BorderColors) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *BorderColors) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "BorderColors")
}

var _TypesValues = []Types{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

// TypesN is the highest valid value for type Types, plus one.
const TypesN Types = 21

var _TypesValueMap = map[string]Types{`UndefinedType`: 0, `Bool32`: 1, `Int16`: 2, `Uint16`: 3, `Int32`: 4, `Int32Vector2`: 5, `Int32Vector4`: 6, `Uint32`: 7, `Uint32Vector2`: 8, `Uint32Vector4`: 9, `Float32`: 10, `Float32Vector2`: 11, `Float32Vector3`: 12, `Float32Vector4`: 13, `Float32Matrix4`: 14, `Float32Matrix3`: 15, `TextureRGBA32`: 16, `TextureBGRA32`: 17, `Depth32`: 18, `Depth24Stencil8`: 19, `Struct`: 20}

var _TypesDescMap = map[Types]string{0: ``, 1: ``, 2: ``, 3: ``, 4: ``, 5: ``, 6: ``, 7: ``, 8: ``, 9: ``, 10: ``, 11: ``, 12: ``, 13: ``, 14: ``, 15: ``, 16: ``, 17: ``, 18: ``, 19: ``, 20: ``}

var _TypesMap = map[Types]string{0: `UndefinedType`, 1: `Bool32`, 2: `Int16`, 3: `Uint16`, 4: `Int32`, 5: `Int32Vector2`, 6: `Int32Vector4`, 7: `Uint32`, 8: `Uint32Vector2`, 9: `Uint32Vector4`, 10: `Float32`, 11: `Float32Vector2`, 12: `Float32Vector3`, 13: `Float32Vector4`, 14: `Float32Matrix4`, 15: `Float32Matrix3`, 16: `TextureRGBA32`, 17: `TextureBGRA32`, 18: `Depth32`, 19: `Depth24Stencil8`, 20: `Struct`}

// String returns the string representation of this Types value.
func (i Types) String() string { return enums.String(i, _TypesMap) }

// SetString sets the Types value from its string representation,
// and returns an error if the string is invalid.
func (i *Types) SetString(s string) error { return enums.SetString(i, s, _TypesValueMap, "Types") }

// Int64 returns the Types value as an int64.
func (i Types) Int64() int64 { return int64(i) }

// SetInt64 sets the Types value from an int64.
func (i *Types) SetInt64(in int64) { *i = Types(in) }

// Desc returns the description of the Types value.
func (i Types) Desc() string { return enums.Desc(i, _TypesDescMap) }

// TypesValues returns all possible values for the type Types.
func TypesValues() []Types { return _TypesValues }

// Values returns all possible values for the type Types.
func (i Types) Values() []enums.Enum { return enums.Values(_TypesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Types) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Types) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Types") }
