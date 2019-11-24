package xlib

// #cgo windows CFLAGS: -I${SRCDIR}/X11/include
// #cgo windows LDFLAGS: -L${SRCDIR}/X11/lib/ -llibX11
import "C"
