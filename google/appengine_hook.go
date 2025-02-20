// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine appenginevm

package google

import "github.com/sgtsquiggs/appengine"

func init() {
	appengineTokenFunc = appengine.AccessToken
	appengineAppIDFunc = appengine.AppID
}
