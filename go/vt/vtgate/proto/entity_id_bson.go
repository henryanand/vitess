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
)

// MarshalBson bson-encodes EntityId.
func (entityId *EntityId) MarshalBson(buf *bytes2.ChunkedWriter, key string) {
	bson.EncodeOptionalPrefix(buf, bson.Object, key)
	lenWriter := bson.NewLenWriter(buf)

	bson.EncodeInterface(buf, "ExternalID", entityId.ExternalID)
	entityId.KeyspaceID.MarshalBson(buf, "KeyspaceID")

	lenWriter.Close()
}

// UnmarshalBson bson-decodes into EntityId.
func (entityId *EntityId) UnmarshalBson(buf *bytes.Buffer, kind byte) {
	switch kind {
	case bson.EOO, bson.Object:
		// valid
	case bson.Null:
		return
	default:
		panic(bson.NewBsonError("unexpected kind %v for EntityId", kind))
	}
	bson.Next(buf, 4)

	for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
		switch bson.ReadCString(buf) {
		case "ExternalID":
			entityId.ExternalID = bson.DecodeInterface(buf, kind)
		case "KeyspaceID":
			entityId.KeyspaceID.UnmarshalBson(buf, kind)
		default:
			bson.Skip(buf, kind)
		}
	}
}
