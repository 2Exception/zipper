// Copyright (C) 2017, Zipper Team.  All rights reserved.
//
// This file is part of zipper
//
// The zipper is free software: you can use, copy, modify,
// and distribute this software for any purpose with or
// without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// The zipper is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// ISC License for more details.
//
// You should have received a copy of the ISC License
// along with this program.  If not, see <https://opensource.org/licenses/isc>.

package db

import (
	"testing"

	"upper.io/db.v3"
	"upper.io/db.v3/mongo"
	"upper.io/db.v3/mssql"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/postgresql"
	"upper.io/db.v3/ql"
	"upper.io/db.v3/sqlite"
)

var wrappers = []string{
	mongo.Adapter,
	mssql.Adapter,
	mysql.Adapter,
	postgresql.Adapter,
	ql.Adapter,
	sqlite.Adapter,
}

func TestSetup(t *testing.T) {
	var err error
	for _, wrapper := range wrappers {
		t.Logf("Testing wrapper: %q TestSetup", wrapper)

		var sess db.Database
		sess, err = Open(wrapper)
		if err != nil {
			t.Fatalf(`Test for wrapper %s failed: %q`, wrapper, err)
		}

		err = Setup(wrapper, sess)
		if err != nil {
			t.Fatalf(`Test for wrapper %s failed: %q`, wrapper, err)
		}

		err = sess.Close()
		if err != nil {
			t.Fatalf(`Test for wrapper %s failed: %q`, wrapper, err)
		}
	}
}
