// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imgui

// #include <stdlib.h>
// #include <memory.h>
// #include "extra_types.h"
// #include "cimplot_wrapper.h"
import "C"
import "unsafe"

type PlotAxisColor C.ImPlotAxisColor

func (self PlotAxisColor) handle() (result *C.ImPlotAxisColor, fin func()) {
	result = (*C.ImPlotAxisColor)(unsafe.Pointer(&self))
	return result, func() {}
}

func (self PlotAxisColor) c() (C.ImPlotAxisColor, func()) {
	result, fin := self.handle()
	return *result, fin
}

func newImPlotAxisColorFromC(cvalue *C.ImPlotAxisColor) *PlotAxisColor {
	return (*PlotAxisColor)(cvalue)
}