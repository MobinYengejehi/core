// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflectx

import (
	"image"
	"reflect"
	"testing"

	"github.com/MobinYengejehi/core/colors"
	"github.com/stretchr/testify/assert"
)

type person struct {
	Name                string `default:"Go Gopher"`
	Age                 int    `default:"35"`
	ProgrammingLanguage string `default:"Go"`
	Pet                 pet
	FavoriteFruit       string `default:"Apple"`
	Data                string `save:"-"`
	OtherPet            *pet
}

type pet struct {
	Name       string
	Type       string `default:"Gopher"`
	Age        int    `default:"7"`
	IsSick     bool
	LikesFoods []string
}

func TestNonDefaultFields(t *testing.T) {
	p := &person{
		Name:                "Go Gopher",
		Age:                 23,
		ProgrammingLanguage: "Go",
		FavoriteFruit:       "Peach",
		Data:                "abcdef",
		Pet: pet{
			Name: "Pet Gopher",
			Type: "Dog",
			Age:  7,
		},
	}
	want := map[string]any{
		"Age":           23,
		"FavoriteFruit": "Peach",
		"Pet": map[string]any{
			"Name": "Pet Gopher",
			"Type": "Dog",
		},
	}
	have := NonDefaultFields(p)
	if !reflect.DeepEqual(have, want) {
		t.Errorf("expected\n%v\n\tbut got\n%v", want, have)
	}
}

type imgfield struct {
	Mycolor image.Image
}

func TestCopyFields(t *testing.T) {
	sp := &person{
		Name:                "Go Gopher",
		Age:                 23,
		ProgrammingLanguage: "Go",
		FavoriteFruit:       "Peach",
		Data:                "abcdef",
		Pet: pet{
			Name: "Pet Gopher",
			Type: "Dog",
			Age:  7,
		},
	}
	dp := &person{}
	CopyFields(dp, sp, "Name", "Pet.Age")
	assert.Equal(t, sp.Name, dp.Name)
	assert.Equal(t, sp.Pet.Age, dp.Pet.Age)

	sif := &imgfield{
		Mycolor: colors.Uniform(colors.Black),
	}
	dif := &imgfield{}
	CopyFields(dif, sif, "Mycolor")
	assert.Equal(t, sif.Mycolor, dif.Mycolor)
}

func TestFieldByPath(t *testing.T) {
	sp := &person{
		Name:                "Go Gopher",
		Age:                 23,
		ProgrammingLanguage: "Go",
		FavoriteFruit:       "Peach",
		Data:                "abcdef",
		Pet: pet{
			Name: "Pet Gopher",
			Type: "Dog",
			Age:  7,
		},
	}
	spv := reflect.ValueOf(sp)
	fv, err := FieldByPath(spv, "Pet.Age")
	assert.NoError(t, err)
	assert.Equal(t, 7, fv.Interface())
	fv, err = FieldByPath(spv, "Pet.Name")
	assert.NoError(t, err)
	assert.Equal(t, "Pet Gopher", fv.Interface())
	fv, err = FieldByPath(spv, "Pet.Ages")
	assert.Error(t, err)
	fv, err = FieldByPath(spv, "Pets.Age")
	assert.Error(t, err)

	err = SetFieldsFromMap(sp, map[string]any{"Pet.Age": 8, "Data": "ddd"})
	assert.NoError(t, err)
	assert.Equal(t, 8, sp.Pet.Age)
	assert.Equal(t, "ddd", sp.Data)
}
