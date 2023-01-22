package TPress

/*
#cgo CFLAGS: -I./include/
#cgo LDFLAGS: ./libTPress.dll
#include <TPress.h>
*/
import "C"

func TypeText(text string) {
	TEXT := C.CString(text)
	C.TypeText(TEXT)
}
