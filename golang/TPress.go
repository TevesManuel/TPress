package TPress

// // // #cgo LDFLAGS: -L./../lib/ -lTPress

/*
#cgo CFLAGS: -I./../include/
#cgo LDFLAGS: ./../lib/libTPress.dll
#include <TPress.h>
*/
import "C"

func TypeText(text string) {
	TEXT := C.CString(text)
	C.TypeText(TEXT)
}
