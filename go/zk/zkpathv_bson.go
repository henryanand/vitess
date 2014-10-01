// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zk

// DO NOT EDIT.
// FILE GENERATED BY BSONGEN.

import (
	"bytes"

	"github.com/henryanand/vitess/go/bson"
	"github.com/henryanand/vitess/go/bytes2"
)

// MarshalBson bson-encodes ZkPathV.
func (zkPathV *ZkPathV) MarshalBson(buf *bytes2.ChunkedWriter, key string) {
	bson.EncodeOptionalPrefix(buf, bson.Object, key)
	lenWriter := bson.NewLenWriter(buf)

	// []string
	{
		bson.EncodePrefix(buf, bson.Array, "Paths")
		lenWriter := bson.NewLenWriter(buf)
		for _i, _v1 := range zkPathV.Paths {
			bson.EncodeString(buf, bson.Itoa(_i), _v1)
		}
		lenWriter.Close()
	}

	lenWriter.Close()
}

// UnmarshalBson bson-decodes into ZkPathV.
func (zkPathV *ZkPathV) UnmarshalBson(buf *bytes.Buffer, kind byte) {
	switch kind {
	case bson.EOO, bson.Object:
		// valid
	case bson.Null:
		return
	default:
		panic(bson.NewBsonError("unexpected kind %v for ZkPathV", kind))
	}
	bson.Next(buf, 4)

	for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
		switch bson.ReadCString(buf) {
		case "Paths":
			// []string
			if kind != bson.Null {
				if kind != bson.Array {
					panic(bson.NewBsonError("unexpected kind %v for zkPathV.Paths", kind))
				}
				bson.Next(buf, 4)
				zkPathV.Paths = make([]string, 0, 8)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					bson.SkipIndex(buf)
					var _v1 string
					_v1 = bson.DecodeString(buf, kind)
					zkPathV.Paths = append(zkPathV.Paths, _v1)
				}
			}
		default:
			bson.Skip(buf, kind)
		}
	}
}
