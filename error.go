// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"
import "fmt"

// Error represents an OpenGL error flag indicator.
type Error GLenum

// OpenGL error enumeration
const (
	INVALID_ENUM                  Error = C.GL_INVALID_ENUM
	INVALID_FRAMEBUFFER_OPERATION Error = C.GL_INVALID_FRAMEBUFFER_OPERATION
	INVALID_OPERATION             Error = C.GL_INVALID_OPERATION
	INVALID_VALUE                 Error = C.GL_INVALID_VALUE
	OUT_OF_MEMORY                 Error = C.GL_OUT_OF_MEMORY
	STACK_OVERFLOW                Error = C.GL_STACK_OVERFLOW
	STACK_UNDERFLOW               Error = C.GL_STACK_UNDERFLOW
)

// Error converts an OpenGL error to a human readable string.
func (err Error) Error() string {
	switch err {
	case INVALID_ENUM:
		return "unacceptable value for enumerated argument"
	case INVALID_FRAMEBUFFER_OPERATION:
		return "incomplete framebuffer object"
	case INVALID_OPERATION:
		return "operation is not allowed"
	case INVALID_VALUE:
		return "numeric argument is out of range"
	case OUT_OF_MEMORY:
		return "out of memory"
	case STACK_OVERFLOW:
		return "internal stack overflow"
	case STACK_UNDERFLOW:
		return "internal stack underflow"
	default:
		return fmt.Sprintf("%#x", int(err))
	}
}

// ClearError discards any currently detected error flags.
// To allow for distributed implementations, there may be several error flags, hence
// error flags should always be cleared by calling GetError in a loop until nil is returned,
// this is a convenience function to do so.
func ClearError() {
	flag := C.glGetError()
	for flag != C.GL_NO_ERROR {
		flag = C.glGetError()
	}
}

// GetError returns the value of the error flag, nil if there was no detectable error
// since the last call.
func GetError() (err Error) {
	flag := C.glGetError()
	if flag != C.GL_NO_ERROR {
		err = Error(flag)
	}
	return
}
