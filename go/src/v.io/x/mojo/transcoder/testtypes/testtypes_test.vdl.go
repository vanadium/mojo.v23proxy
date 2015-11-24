// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: testtypes_test.vdl

package testtypes

import (
	// VDL system imports
	"fmt"
	"v.io/v23/vdl"
)

type AnEnum int

const (
	AnEnumFirst AnEnum = iota
	AnEnumSecond
)

// AnEnumAll holds all labels for AnEnum.
var AnEnumAll = [...]AnEnum{AnEnumFirst, AnEnumSecond}

// AnEnumFromString creates a AnEnum from a string label.
func AnEnumFromString(label string) (x AnEnum, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *AnEnum) Set(label string) error {
	switch label {
	case "First", "first":
		*x = AnEnumFirst
		return nil
	case "Second", "second":
		*x = AnEnumSecond
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in testtypes.AnEnum", label)
}

// String returns the string label of x.
func (x AnEnum) String() string {
	switch x {
	case AnEnumFirst:
		return "First"
	case AnEnumSecond:
		return "Second"
	}
	return ""
}

func (AnEnum) __VDLReflect(struct {
	Name string `vdl:"src/v.io/x/mojo/transcoder/testtypes.AnEnum"`
	Enum struct{ First, Second string }
}) {
}

type (
	// PodUnion represents any single field of the PodUnion union type.
	PodUnion interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the PodUnion union type.
		__VDLReflect(__PodUnionReflect)
	}
	// PodUnionFInt8 represents field FInt8 of the PodUnion union type.
	PodUnionFInt8 struct{ Value int8 }
	// PodUnionFInt8Other represents field FInt8Other of the PodUnion union type.
	PodUnionFInt8Other struct{ Value int8 }
	// PodUnionFUint8 represents field FUint8 of the PodUnion union type.
	PodUnionFUint8 struct{ Value byte }
	// PodUnionFInt16 represents field FInt16 of the PodUnion union type.
	PodUnionFInt16 struct{ Value int16 }
	// PodUnionFUint16 represents field FUint16 of the PodUnion union type.
	PodUnionFUint16 struct{ Value uint16 }
	// PodUnionFint32 represents field Fint32 of the PodUnion union type.
	PodUnionFint32 struct{ Value int32 }
	// PodUnionFuint32 represents field Fuint32 of the PodUnion union type.
	PodUnionFuint32 struct{ Value uint32 }
	// PodUnionFInt64 represents field FInt64 of the PodUnion union type.
	PodUnionFInt64 struct{ Value int64 }
	// PodUnionFUint64 represents field FUint64 of the PodUnion union type.
	PodUnionFUint64 struct{ Value uint64 }
	// PodUnionFFloat represents field FFloat of the PodUnion union type.
	PodUnionFFloat struct{ Value float32 }
	// PodUnionFDouble represents field FDouble of the PodUnion union type.
	PodUnionFDouble struct{ Value float64 }
	// PodUnionFBool represents field FBool of the PodUnion union type.
	PodUnionFBool struct{ Value bool }
	// PodUnionFEnum represents field FEnum of the PodUnion union type.
	PodUnionFEnum struct{ Value AnEnum }
	// __PodUnionReflect describes the PodUnion union type.
	__PodUnionReflect struct {
		Name  string `vdl:"src/v.io/x/mojo/transcoder/testtypes.PodUnion"`
		Type  PodUnion
		Union struct {
			FInt8      PodUnionFInt8
			FInt8Other PodUnionFInt8Other
			FUint8     PodUnionFUint8
			FInt16     PodUnionFInt16
			FUint16    PodUnionFUint16
			Fint32     PodUnionFint32
			Fuint32    PodUnionFuint32
			FInt64     PodUnionFInt64
			FUint64    PodUnionFUint64
			FFloat     PodUnionFFloat
			FDouble    PodUnionFDouble
			FBool      PodUnionFBool
			FEnum      PodUnionFEnum
		}
	}
)

func (x PodUnionFInt8) Index() int                     { return 0 }
func (x PodUnionFInt8) Interface() interface{}         { return x.Value }
func (x PodUnionFInt8) Name() string                   { return "FInt8" }
func (x PodUnionFInt8) __VDLReflect(__PodUnionReflect) {}

func (x PodUnionFInt8Other) Index() int                     { return 1 }
func (x PodUnionFInt8Other) Interface() interface{}         { return x.Value }
func (x PodUnionFInt8Other) Name() string                   { return "FInt8Other" }
func (x PodUnionFInt8Other) __VDLReflect(__PodUnionReflect) {}

func (x PodUnionFUint8) Index() int                     { return 2 }
func (x PodUnionFUint8) Interface() interface{}         { return x.Value }
func (x PodUnionFUint8) Name() string                   { return "FUint8" }
func (x PodUnionFUint8) __VDLReflect(__PodUnionReflect) {}

func (x PodUnionFInt16) Index() int                     { return 3 }
func (x PodUnionFInt16) Interface() interface{}         { return x.Value }
func (x PodUnionFInt16) Name() string                   { return "FInt16" }
func (x PodUnionFInt16) __VDLReflect(__PodUnionReflect) {}

func (x PodUnionFUint16) Index() int                     { return 4 }
func (x PodUnionFUint16) Interface() interface{}         { return x.Value }
func (x PodUnionFUint16) Name() string                   { return "FUint16" }
func (x PodUnionFUint16) __VDLReflect(__PodUnionReflect) {}

func (x PodUnionFint32) Index() int                     { return 5 }
func (x PodUnionFint32) Interface() interface{}         { return x.Value }
func (x PodUnionFint32) Name() string                   { return "Fint32" }
func (x PodUnionFint32) __VDLReflect(__PodUnionReflect) {}

func (x PodUnionFuint32) Index() int                     { return 6 }
func (x PodUnionFuint32) Interface() interface{}         { return x.Value }
func (x PodUnionFuint32) Name() string                   { return "Fuint32" }
func (x PodUnionFuint32) __VDLReflect(__PodUnionReflect) {}

func (x PodUnionFInt64) Index() int                     { return 7 }
func (x PodUnionFInt64) Interface() interface{}         { return x.Value }
func (x PodUnionFInt64) Name() string                   { return "FInt64" }
func (x PodUnionFInt64) __VDLReflect(__PodUnionReflect) {}

func (x PodUnionFUint64) Index() int                     { return 8 }
func (x PodUnionFUint64) Interface() interface{}         { return x.Value }
func (x PodUnionFUint64) Name() string                   { return "FUint64" }
func (x PodUnionFUint64) __VDLReflect(__PodUnionReflect) {}

func (x PodUnionFFloat) Index() int                     { return 9 }
func (x PodUnionFFloat) Interface() interface{}         { return x.Value }
func (x PodUnionFFloat) Name() string                   { return "FFloat" }
func (x PodUnionFFloat) __VDLReflect(__PodUnionReflect) {}

func (x PodUnionFDouble) Index() int                     { return 10 }
func (x PodUnionFDouble) Interface() interface{}         { return x.Value }
func (x PodUnionFDouble) Name() string                   { return "FDouble" }
func (x PodUnionFDouble) __VDLReflect(__PodUnionReflect) {}

func (x PodUnionFBool) Index() int                     { return 11 }
func (x PodUnionFBool) Interface() interface{}         { return x.Value }
func (x PodUnionFBool) Name() string                   { return "FBool" }
func (x PodUnionFBool) __VDLReflect(__PodUnionReflect) {}

func (x PodUnionFEnum) Index() int                     { return 12 }
func (x PodUnionFEnum) Interface() interface{}         { return x.Value }
func (x PodUnionFEnum) Name() string                   { return "FEnum" }
func (x PodUnionFEnum) __VDLReflect(__PodUnionReflect) {}

type PodUnionWrapper struct {
	PodUnion PodUnion
}

func (PodUnionWrapper) __VDLReflect(struct {
	Name string `vdl:"src/v.io/x/mojo/transcoder/testtypes.PodUnionWrapper"`
}) {
}

type DummyStruct struct {
	FInt8 int8
}

func (DummyStruct) __VDLReflect(struct {
	Name string `vdl:"src/v.io/x/mojo/transcoder/testtypes.DummyStruct"`
}) {
}

type (
	// ObjectUnion represents any single field of the ObjectUnion union type.
	ObjectUnion interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the ObjectUnion union type.
		__VDLReflect(__ObjectUnionReflect)
	}
	// ObjectUnionFInt8 represents field FInt8 of the ObjectUnion union type.
	ObjectUnionFInt8 struct{ Value int8 }
	// ObjectUnionFString represents field FString of the ObjectUnion union type.
	ObjectUnionFString struct{ Value string }
	// ObjectUnionFDummy represents field FDummy of the ObjectUnion union type.
	ObjectUnionFDummy struct{ Value DummyStruct }
	// ObjectUnionFNullable represents field FNullable of the ObjectUnion union type.
	ObjectUnionFNullable struct{ Value *DummyStruct }
	// ObjectUnionFArrayInt8 represents field FArrayInt8 of the ObjectUnion union type.
	ObjectUnionFArrayInt8 struct{ Value []int8 }
	// ObjectUnionFMapInt8 represents field FMapInt8 of the ObjectUnion union type.
	ObjectUnionFMapInt8 struct{ Value map[string]int8 }
	// ObjectUnionFPodUnion represents field FPodUnion of the ObjectUnion union type.
	ObjectUnionFPodUnion struct{ Value PodUnion }
	// __ObjectUnionReflect describes the ObjectUnion union type.
	__ObjectUnionReflect struct {
		Name  string `vdl:"src/v.io/x/mojo/transcoder/testtypes.ObjectUnion"`
		Type  ObjectUnion
		Union struct {
			FInt8      ObjectUnionFInt8
			FString    ObjectUnionFString
			FDummy     ObjectUnionFDummy
			FNullable  ObjectUnionFNullable
			FArrayInt8 ObjectUnionFArrayInt8
			FMapInt8   ObjectUnionFMapInt8
			FPodUnion  ObjectUnionFPodUnion
		}
	}
)

func (x ObjectUnionFInt8) Index() int                        { return 0 }
func (x ObjectUnionFInt8) Interface() interface{}            { return x.Value }
func (x ObjectUnionFInt8) Name() string                      { return "FInt8" }
func (x ObjectUnionFInt8) __VDLReflect(__ObjectUnionReflect) {}

func (x ObjectUnionFString) Index() int                        { return 1 }
func (x ObjectUnionFString) Interface() interface{}            { return x.Value }
func (x ObjectUnionFString) Name() string                      { return "FString" }
func (x ObjectUnionFString) __VDLReflect(__ObjectUnionReflect) {}

func (x ObjectUnionFDummy) Index() int                        { return 2 }
func (x ObjectUnionFDummy) Interface() interface{}            { return x.Value }
func (x ObjectUnionFDummy) Name() string                      { return "FDummy" }
func (x ObjectUnionFDummy) __VDLReflect(__ObjectUnionReflect) {}

func (x ObjectUnionFNullable) Index() int                        { return 3 }
func (x ObjectUnionFNullable) Interface() interface{}            { return x.Value }
func (x ObjectUnionFNullable) Name() string                      { return "FNullable" }
func (x ObjectUnionFNullable) __VDLReflect(__ObjectUnionReflect) {}

func (x ObjectUnionFArrayInt8) Index() int                        { return 4 }
func (x ObjectUnionFArrayInt8) Interface() interface{}            { return x.Value }
func (x ObjectUnionFArrayInt8) Name() string                      { return "FArrayInt8" }
func (x ObjectUnionFArrayInt8) __VDLReflect(__ObjectUnionReflect) {}

func (x ObjectUnionFMapInt8) Index() int                        { return 5 }
func (x ObjectUnionFMapInt8) Interface() interface{}            { return x.Value }
func (x ObjectUnionFMapInt8) Name() string                      { return "FMapInt8" }
func (x ObjectUnionFMapInt8) __VDLReflect(__ObjectUnionReflect) {}

func (x ObjectUnionFPodUnion) Index() int                        { return 6 }
func (x ObjectUnionFPodUnion) Interface() interface{}            { return x.Value }
func (x ObjectUnionFPodUnion) Name() string                      { return "FPodUnion" }
func (x ObjectUnionFPodUnion) __VDLReflect(__ObjectUnionReflect) {}

type ObjectUnionWrapper struct {
	ObjectUnion ObjectUnion
}

func (ObjectUnionWrapper) __VDLReflect(struct {
	Name string `vdl:"src/v.io/x/mojo/transcoder/testtypes.ObjectUnionWrapper"`
}) {
}

type Rect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func (Rect) __VDLReflect(struct {
	Name string `vdl:"src/v.io/x/mojo/transcoder/testtypes.Rect"`
}) {
}

type MultiVersionStruct struct {
	FInt32  int32
	FRect   Rect
	FString string
	FArray  []int8
	FBool   bool
	FInt16  int16
}

func (MultiVersionStruct) __VDLReflect(struct {
	Name string `vdl:"src/v.io/x/mojo/transcoder/testtypes.MultiVersionStruct"`
}) {
}

type MultiVersionStructV3 struct {
	FInt32  int32
	FRect   Rect
	FString string
}

func (MultiVersionStructV3) __VDLReflect(struct {
	Name string `vdl:"src/v.io/x/mojo/transcoder/testtypes.MultiVersionStructV3"`
}) {
}

func init() {
	vdl.Register((*AnEnum)(nil))
	vdl.Register((*PodUnion)(nil))
	vdl.Register((*PodUnionWrapper)(nil))
	vdl.Register((*DummyStruct)(nil))
	vdl.Register((*ObjectUnion)(nil))
	vdl.Register((*ObjectUnionWrapper)(nil))
	vdl.Register((*Rect)(nil))
	vdl.Register((*MultiVersionStruct)(nil))
	vdl.Register((*MultiVersionStructV3)(nil))
}
