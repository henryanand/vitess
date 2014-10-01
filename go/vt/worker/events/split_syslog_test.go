// Copyright 2014, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"log/syslog"
	"testing"

	base "github.com/henryanand/vitess/go/vt/events"
)

func TestReparentSyslog(t *testing.T) {
	wantSev, wantMsg := syslog.LOG_INFO, "keyspace-123/shard-123/cell-1 [vertical split clone] status"
	ev := &VerticalSplitClone{
		Cell:          "cell-1",
		Keyspace:      "keyspace-123",
		Shard:         "shard-123",
		StatusUpdater: base.StatusUpdater{Status: "status"},
	}
	gotSev, gotMsg := ev.Syslog()

	if gotSev != wantSev {
		t.Errorf("wrong severity: got %v, want %v", gotSev, wantSev)
	}
	if gotMsg != wantMsg {
		t.Errorf("wrong message: got %v, want %v", gotMsg, wantMsg)
	}
}
