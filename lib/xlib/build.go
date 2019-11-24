package xlib

// windows 下X11无法初始化
// #cgo !windows CFLAGS: -I${SRCDIR}/X11/include
// #cgo !windows LDFLAGS: -llibX11
import "C"
