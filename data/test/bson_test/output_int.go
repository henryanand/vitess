// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mytype

import (
	"github.com/henryanand/vitess/go/bytes2"

	"bytes"

	"github.com/henryanand/vitess/go/bson"
)

// DO NOT EDIT.
// FILE GENERATED BY BSONGEN.

// MarshalBson bson-encodes MyType.
func (myType MyType) MarshalBson(buf *bytes2.ChunkedWriter, key string) {
	if key == "" {
		lenWriter := bson.NewLenWriter(buf)
		defer lenWriter.Close()
		key = bson.MAGICTAG
	}
	bson.EncodeInt(buf, key, int(myType))
}

// UnmarshalBson bson-decodes into MyType.
func (myType *MyType) UnmarshalBson(buf *bytes.Buffer, kind byte) {
	if kind == bson.EOO {
		bson.Next(buf, 4)
		kind = bson.NextByte(buf)
		bson.ReadCString(buf)
	}
	*myType = MyType(bson.DecodeInt(buf, kind))
}
