// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package cef

/*
#cgo CFLAGS: -I./../../
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
*/
import "C"
import "unsafe"
import (
    "os"
)

var g_argv []*C.char = make([]*C.char, len(os.Args))

func FillMainArgs(mainArgs *C.struct__cef_main_args_t,
        appHandle unsafe.Pointer) {
    // On Linux appHandle is nil.
    g_logger.Println("FillMainArgs, argc=", len(os.Args))
    for i, arg := range os.Args {
        g_argv[C.int(i)] = C.CString(arg)
    }
    mainArgs.argc = C.int(len(os.Args))
    mainArgs.argv = &g_argv[0]
}

func FillWindowInfo(windowInfo *C.cef_window_info_t, hwnd unsafe.Pointer) {
    g_logger.Println("FillWindowInfo")
    windowInfo.parent_widget = (*C.GtkWidget)(hwnd)
}
