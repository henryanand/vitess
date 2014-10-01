// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mytype

import (
	"github.com/henryanand/vitess/go/bson"
	"github.com/henryanand/vitess/go/bytes2"

	"bytes"
)

// DO NOT EDIT.
// FILE GENERATED BY BSONGEN.

// MarshalBson bson-encodes MyType.
func (myType *MyType) MarshalBson(buf *bytes2.ChunkedWriter, key string) {
	bson.EncodeOptionalPrefix(buf, bson.Object, key)
	lenWriter := bson.NewLenWriter(buf)

	// map[string]string
	{
		bson.EncodePrefix(buf, bson.Object, "Map")
		lenWriter := bson.NewLenWriter(buf)
		for _k, _v1 := range myType.Map {
			bson.EncodeString(buf, _k, _v1)
		}
		lenWriter.Close()
	}
	// map[string][]byte
	{
		bson.EncodePrefix(buf, bson.Object, "MapBytes")
		lenWriter := bson.NewLenWriter(buf)
		for _k, _v2 := range myType.MapBytes {
			bson.EncodeBinary(buf, _k, _v2)
		}
		lenWriter.Close()
	}
	// map[string]*string
	{
		bson.EncodePrefix(buf, bson.Object, "MapPtr")
		lenWriter := bson.NewLenWriter(buf)
		for _k, _v3 := range myType.MapPtr {
			// *string
			if _v3 == nil {
				bson.EncodePrefix(buf, bson.Null, _k)
			} else {
				bson.EncodeString(buf, _k, (*_v3))
			}
		}
		lenWriter.Close()
	}
	// map[string][]string
	{
		bson.EncodePrefix(buf, bson.Object, "MapSlice")
		lenWriter := bson.NewLenWriter(buf)
		for _k, _v4 := range myType.MapSlice {
			// []string
			{
				bson.EncodePrefix(buf, bson.Array, _k)
				lenWriter := bson.NewLenWriter(buf)
				for _i, _v5 := range _v4 {
					bson.EncodeString(buf, bson.Itoa(_i), _v5)
				}
				lenWriter.Close()
			}
		}
		lenWriter.Close()
	}
	// map[string]map[string]int64
	{
		bson.EncodePrefix(buf, bson.Object, "MapMap")
		lenWriter := bson.NewLenWriter(buf)
		for _k, _v6 := range myType.MapMap {
			// map[string]int64
			{
				bson.EncodePrefix(buf, bson.Object, _k)
				lenWriter := bson.NewLenWriter(buf)
				for _k, _v7 := range _v6 {
					bson.EncodeInt64(buf, _k, _v7)
				}
				lenWriter.Close()
			}
		}
		lenWriter.Close()
	}
	// map[string]Custom
	{
		bson.EncodePrefix(buf, bson.Object, "MapCustom")
		lenWriter := bson.NewLenWriter(buf)
		for _k, _v8 := range myType.MapCustom {
			_v8.MarshalBson(buf, _k)
		}
		lenWriter.Close()
	}
	// map[string]*Custom
	{
		bson.EncodePrefix(buf, bson.Object, "MapCustomPtr")
		lenWriter := bson.NewLenWriter(buf)
		for _k, _v9 := range myType.MapCustomPtr {
			// *Custom
			if _v9 == nil {
				bson.EncodePrefix(buf, bson.Null, _k)
			} else {
				(*_v9).MarshalBson(buf, _k)
			}
		}
		lenWriter.Close()
	}
	// map[Custom]string
	{
		bson.EncodePrefix(buf, bson.Object, "CustomMap")
		lenWriter := bson.NewLenWriter(buf)
		for _k, _v10 := range myType.CustomMap {
			bson.EncodeString(buf, string(_k), _v10)
		}
		lenWriter.Close()
	}
	// map[pkg.Custom]string
	{
		bson.EncodePrefix(buf, bson.Object, "MapExternal")
		lenWriter := bson.NewLenWriter(buf)
		for _k, _v11 := range myType.MapExternal {
			bson.EncodeString(buf, string(_k), _v11)
		}
		lenWriter.Close()
	}

	lenWriter.Close()
}

// UnmarshalBson bson-decodes into MyType.
func (myType *MyType) UnmarshalBson(buf *bytes.Buffer, kind byte) {
	switch kind {
	case bson.EOO, bson.Object:
		// valid
	case bson.Null:
		return
	default:
		panic(bson.NewBsonError("unexpected kind %v for MyType", kind))
	}
	bson.Next(buf, 4)

	for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
		switch bson.ReadCString(buf) {
		case "Map":
			// map[string]string
			if kind != bson.Null {
				if kind != bson.Object {
					panic(bson.NewBsonError("unexpected kind %v for myType.Map", kind))
				}
				bson.Next(buf, 4)
				myType.Map = make(map[string]string)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					_k := bson.ReadCString(buf)
					var _v1 string
					_v1 = bson.DecodeString(buf, kind)
					myType.Map[_k] = _v1
				}
			}
		case "MapBytes":
			// map[string][]byte
			if kind != bson.Null {
				if kind != bson.Object {
					panic(bson.NewBsonError("unexpected kind %v for myType.MapBytes", kind))
				}
				bson.Next(buf, 4)
				myType.MapBytes = make(map[string][]byte)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					_k := bson.ReadCString(buf)
					var _v2 []byte
					_v2 = bson.DecodeBinary(buf, kind)
					myType.MapBytes[_k] = _v2
				}
			}
		case "MapPtr":
			// map[string]*string
			if kind != bson.Null {
				if kind != bson.Object {
					panic(bson.NewBsonError("unexpected kind %v for myType.MapPtr", kind))
				}
				bson.Next(buf, 4)
				myType.MapPtr = make(map[string]*string)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					_k := bson.ReadCString(buf)
					var _v3 *string
					// *string
					if kind != bson.Null {
						_v3 = new(string)
						(*_v3) = bson.DecodeString(buf, kind)
					}
					myType.MapPtr[_k] = _v3
				}
			}
		case "MapSlice":
			// map[string][]string
			if kind != bson.Null {
				if kind != bson.Object {
					panic(bson.NewBsonError("unexpected kind %v for myType.MapSlice", kind))
				}
				bson.Next(buf, 4)
				myType.MapSlice = make(map[string][]string)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					_k := bson.ReadCString(buf)
					var _v4 []string
					// []string
					if kind != bson.Null {
						if kind != bson.Array {
							panic(bson.NewBsonError("unexpected kind %v for _v4", kind))
						}
						bson.Next(buf, 4)
						_v4 = make([]string, 0, 8)
						for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
							bson.SkipIndex(buf)
							var _v5 string
							_v5 = bson.DecodeString(buf, kind)
							_v4 = append(_v4, _v5)
						}
					}
					myType.MapSlice[_k] = _v4
				}
			}
		case "MapMap":
			// map[string]map[string]int64
			if kind != bson.Null {
				if kind != bson.Object {
					panic(bson.NewBsonError("unexpected kind %v for myType.MapMap", kind))
				}
				bson.Next(buf, 4)
				myType.MapMap = make(map[string]map[string]int64)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					_k := bson.ReadCString(buf)
					var _v6 map[string]int64
					// map[string]int64
					if kind != bson.Null {
						if kind != bson.Object {
							panic(bson.NewBsonError("unexpected kind %v for _v6", kind))
						}
						bson.Next(buf, 4)
						_v6 = make(map[string]int64)
						for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
							_k := bson.ReadCString(buf)
							var _v7 int64
							_v7 = bson.DecodeInt64(buf, kind)
							_v6[_k] = _v7
						}
					}
					myType.MapMap[_k] = _v6
				}
			}
		case "MapCustom":
			// map[string]Custom
			if kind != bson.Null {
				if kind != bson.Object {
					panic(bson.NewBsonError("unexpected kind %v for myType.MapCustom", kind))
				}
				bson.Next(buf, 4)
				myType.MapCustom = make(map[string]Custom)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					_k := bson.ReadCString(buf)
					var _v8 Custom
					_v8.UnmarshalBson(buf, kind)
					myType.MapCustom[_k] = _v8
				}
			}
		case "MapCustomPtr":
			// map[string]*Custom
			if kind != bson.Null {
				if kind != bson.Object {
					panic(bson.NewBsonError("unexpected kind %v for myType.MapCustomPtr", kind))
				}
				bson.Next(buf, 4)
				myType.MapCustomPtr = make(map[string]*Custom)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					_k := bson.ReadCString(buf)
					var _v9 *Custom
					// *Custom
					if kind != bson.Null {
						_v9 = new(Custom)
						(*_v9).UnmarshalBson(buf, kind)
					}
					myType.MapCustomPtr[_k] = _v9
				}
			}
		case "CustomMap":
			// map[Custom]string
			if kind != bson.Null {
				if kind != bson.Object {
					panic(bson.NewBsonError("unexpected kind %v for myType.CustomMap", kind))
				}
				bson.Next(buf, 4)
				myType.CustomMap = make(map[Custom]string)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					_k := Custom(bson.ReadCString(buf))
					var _v10 string
					_v10 = bson.DecodeString(buf, kind)
					myType.CustomMap[_k] = _v10
				}
			}
		case "MapExternal":
			// map[pkg.Custom]string
			if kind != bson.Null {
				if kind != bson.Object {
					panic(bson.NewBsonError("unexpected kind %v for myType.MapExternal", kind))
				}
				bson.Next(buf, 4)
				myType.MapExternal = make(map[pkg.Custom]string)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					_k := pkg.Custom(bson.ReadCString(buf))
					var _v11 string
					_v11 = bson.DecodeString(buf, kind)
					myType.MapExternal[_k] = _v11
				}
			}
		default:
			bson.Skip(buf, kind)
		}
	}
}
