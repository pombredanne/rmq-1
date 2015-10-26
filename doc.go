/*
Summary: rmq passes msgpack2 messages over websockets between Golang and the R statistical language. It is an R package.

# rmq: R Messaging and Queuing

### Or: How to utilize Go libraries from R.

The much anticipated Go 1.5 release brought strong support for building C-style shared libraries (.so files) from Go source code and libraries.

*This is huge*. It opens up many exciting new possibilities. In this project (rmq), we explore using this new capability to extend R with Go libraries.

Package rmq provides messaging based on msgpack and websockets. It demonstrates calling from R into Golang (Go) libraries to extend R with functionality available in Go.

We use the Go library https://github.com/ugorji/go codec for msgpack encoding and decoding. This is a high performance implementation. We use it in a mode where it only supports the updated msgpack 2 (current) spec. This is critical for interoperability with other compiled languages that distinguish between utf8 strings and binary blobs (otherwise embedded '\0' zeros in blobs cause problems).

For websockets, we use the terrific https://github.com/gorilla/websocket library. As time permits in the future, we may extend more features aiming towards message queuing as well. The gorilla library supports securing your communication with TLS certificates.

##Status

Excellent. Tested on OSX and Linux. Documentation has been written and is available. The package is functionally complete for the RPC over websockets and msgpack based serialization.  After interactive usage, I added SIGINT handling so that the web-server can be stopped during development with a simple Ctrl-c at the R console. The client side will be blocked during calls (it does not poll back to R while waiting on the network) but has a configurable timeout (default 5 seconds), that allows easy client-side error handling.

## structure of this repo

This repository is mainly structured as an R package. It is
designed to be built and installed into an R (statistical environment)
installation, using the standard tools for R.

This package doesn't directly create a re-usable go library. Instead
we target a c-shared library (rmq.so) that will install
into R using 'R CMD install rmq'. See: 'make install' or 'make build' followed by
doing `install.packages('./rmq_1.0.1.tar.gz', repos=NULL)`
from inside R (assuming the package is in your current directory;
if not then adjust the ./ part of the package path).

The code also serves as an example of how to
use golang inside R.


*/
package rmq

/*
typedef int SEXP;
*/
import "C"

// FromMsgpack converts a serialized RAW vector of of msgpack2
// encoded bytes into an R object. We use msgpack2 so that there is
// a difference between strings (utf8 encoded) and binary blobs
// which can contain '\0' zeros. The underlying msgpack2 library
// is the awesome https://github.com/ugorji/go/tree/master/codec
// library from Ugorji Nwoke.
func FromMsgpack(s C.SEXP) C.SEXP {
	// This is a stub for documentation of API and search purposes.
	// See the actually implementation here:
	// https://github.com/glycerine/rmq/blob/master/src/rmq/rmq.go
	return s
}

// ToMsgpack converts an R object into serialized RAW vector
// of msgpack2 encoded bytes. We use msgpack2 so that there is
// a difference between strings (utf8 encoded) and binary blobs
// which can contain '\0' zeros. The underlying msgpack2 library
// is the awesome https://github.com/ugorji/go/tree/master/codec
// library from Ugorji Nwoke.
func ToMsgpack(s C.SEXP) C.SEXP {
	// This is a stub for documentation of API and search purposes.
	// See the actually implementation here:
	// https://github.com/glycerine/rmq/blob/master/src/rmq/rmq.go
	return s
}

// ListenAndServe is the server part that expects calls from client
// in the form of RmqWebsocketCall() invocations.
// The underlying websocket library is the battle tested
// https://github.com/gorilla/websocket library from the
// Gorilla Web toolkit. http://www.gorillatoolkit.org/
//
// addr_ is a string in "ip:port" format. The server
// will bind this address and port on the local host.
//
// handler_ is an R function that takes a single argument.
// It will be called back each time the server receives
// an incoming message. The returned value of handler
// becomes the reply to the client.
//
// rho_ in an R environment in which the handler_ callback
// will occur. The user-level wrapper rmq.server() provides
// a new environment for every call back by default, so
// most users won't need to worry about rho_.
//
// Return value: this is always R_NilValue.
//
// Semantics: ListenAndServe() will start a new
// webserver everytime it is called. If it exits
// due to a call into R_CheckUserInterrupt()
// or Rf_error(), then a background watchdog goroutine
// will notice the lack of heartbeating after 300ms,
// and will immediately shutdown the listening
// websocket server goroutine. Hence cleanup
// is fairly automatic.
//
// Signal handling:
//
// SIGINT (ctrl-c) is noted by R, and since we
// regularly call R_CheckUserInterrupt(), the
// user can stop the server by pressing ctrl-c
// at the R-console. The go-runtime, as embedded
// in the c-shared library, is not used to being
// embedded yet, and so its (system) signal handling
// facilities (e.g. signal.Notify) should *not* be
// used. We go to great pains to actually preserve
// the signal handling that R sets up and expects,
// and allow the go runtime to see any signals just
// creates heartache and crashes.
//
func ListenAndServe(addr_ C.SEXP, handler_ C.SEXP, rho_ C.SEXP) C.SEXP {
	// This is a stub for documentation of API and search purposes.
	// See the actually implementation here:
	// https://github.com/glycerine/rmq/blob/master/src/rmq/rmq.go
	return addr_
}

// RmqWebsocketCall() is the client part that talks to
// the server part waiting in ListenAndServe().
// ListenAndServe is the server part that expects calls from client
// in the form of RmqWebsocketCall() invocations.
// The underlying websocket library is the battle tested
// https://github.com/gorilla/websocket library from the
// Gorilla Web toolkit. http://www.gorillatoolkit.org/
//
// addr_ is an "ip:port" string: where to find the server;
// it should match the addr_ the server was started with.
//
// msg_ is the R object to be sent to the server.
//
// timeout_msec_ is a numeric count of milliseconds to
// wait for a reply from the server. Timeouts are the
// only way we handle servers that accept our connect
// and then crash or take too long. Although a timeout
// of 0 will wait forever, this is not recommended.
// SIGINT (ctrl-c) will not interrupt a waiting client,
// so do be sure to give it some sane timeout. The
// default is 5000 msec (5 seconds).
//
func RmqWebsocketCall(addr_ C.SEXP, msg_ C.SEXP, timeout_msec_ C.SEXP) C.SEXP {
	// This is a stub for documentation of API and search purposes.
	// See the actually implementation here:
	// https://github.com/glycerine/rmq/blob/master/src/rmq/rmq.go
	return addr_
}
