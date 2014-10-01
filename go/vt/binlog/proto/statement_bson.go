// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

// DO NOT EDIT.
// FILE GENERATED BY BSONGEN.

import (
	"bytes"

	"github.com/henryanand/vitess/go/bson"
	"github.com/henryanand/vitess/go/bytes2"
	mproto "github.com/henryanand/vitess/go/mysql/proto"
)

// MarshalBson bson-encodes Statement.
func (statement *Statement) MarshalBson(buf *bytes2.ChunkedWriter, key string) {
	bson.EncodeOptionalPrefix(buf, bson.Object, key)
	lenWriter := bson.NewLenWriter(buf)

	bson.EncodeInt(buf, "Category", statement.Category)
	// *mproto.Charset
	if statement.Charset == nil {
		bson.EncodePrefix(buf, bson.Null, "Charset")
	} else {
		(*statement.Charset).MarshalBson(buf, "Charset")
	}
	bson.EncodeBinary(buf, "Sql", statement.Sql)

	lenWriter.Close()
}

// UnmarshalBson bson-decodes into Statement.
func (statement *Statement) UnmarshalBson(buf *bytes.Buffer, kind byte) {
	switch kind {
	case bson.EOO, bson.Object:
		// valid
	case bson.Null:
		return
	default:
		panic(bson.NewBsonError("unexpected kind %v for Statement", kind))
	}
	bson.Next(buf, 4)

	for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
		switch bson.ReadCString(buf) {
		case "Category":
			statement.Category = bson.DecodeInt(buf, kind)
		case "Charset":
			// *mproto.Charset
			if kind != bson.Null {
				statement.Charset = new(mproto.Charset)
				(*statement.Charset).UnmarshalBson(buf, kind)
			}
		case "Sql":
			statement.Sql = bson.DecodeBinary(buf, kind)
		default:
			bson.Skip(buf, kind)
		}
	}
}
