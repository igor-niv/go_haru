package main

//#cgo CFLAGS: -I/home/Haru/PDFLib/Haru.Native/src/libharu
//#cgo LDFLAGS: -L/home/Haru/PDFLib/Haru.Native -lharu23 -lm
//#include <hpdf.h>
import "C"
import (
    "fmt"
    "os"
)

func main() {
    var version = C.HPDF_GetVersion()
    fmt.Println("Haru Version: " + C.GoString(version))
    var pdf = C.HPDF_New(nil, nil)
    if pdf == nil {
	fmt.Println("pdf == nil")
	os.Exit(1)
    }
    var page_1 = C.HPDF_AddPage(pdf);
    C.HPDF_InsertPage(pdf, page_1);
    C.HPDF_Page_SetSize (page_1, C.HPDF_PAGE_SIZE_B5, C.HPDF_PAGE_LANDSCAPE);
    C.HPDF_SaveToFile (pdf, C.CString("test.pdf"));
    C.HPDF_Free(pdf)
}
