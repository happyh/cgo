// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go_test_cpp

//#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/../../cpp/cmake -L${SRCDIR}/../../swig/cmake -Wl,-rpath=${SRCDIR}/../../cpp/cmake -ltest_cpp -lgo_test_cpp -lstdc++
//#cgo linux CPPFLAGS: -fPIC -I.
import "C"
