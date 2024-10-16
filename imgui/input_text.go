package imgui

// #include <memory.h>
// #include <stdlib.h>
// #include "extra_types.h"
// #include "cimgui_wrapper.h"
// #include "cimgui_structs_accessor.h"
// extern int generalInputTextCallback(ImGuiInputTextCallbackData* data);
import "C"

import (
	"runtime/cgo"
	"unsafe"

	"github.com/AllenDang/cimgui-go/datautils"
)

type InputTextCallback func(data InputTextCallbackData) int

type inputTextInternalState struct {
	buf      *datautils.StringBuffer
	callback InputTextCallback
}

func (state *inputTextInternalState) release() {
	state.buf.Free()
}

//export generalInputTextCallback
func generalInputTextCallback(cbData *C.ImGuiInputTextCallbackData) C.int {
	data := NewInputTextCallbackDataFromC(cbData)

	bufHandle := (*cgo.Handle)(cbData.UserData)
	statePtr := bufHandle.Value().(*inputTextInternalState)

	if data.EventFlag() == InputTextFlagsCallbackResize {
		statePtr.buf.ResizeTo(int(data.BufSize()))
		C.wrap_ImGuiInputTextCallbackData_SetBuf(cbData, datautils.ConvertCTypes[*C.char](statePtr.buf.Ptr()))
		data.SetBufSize(int32(statePtr.buf.Size()))
		data.SetBufTextLen(int32(data.BufTextLen()))
		data.SetBufDirty(true)
	}

	if statePtr.callback != nil {
		return C.int(statePtr.callback(*data))
	}

	return 0
}

func InputTextWithHint(label, hint string, buf *string, flags InputTextFlags, callback InputTextCallback) bool {
	labelArg, labelFin := datautils.WrapString[C.char](label)
	defer labelFin()

	hintArg, hintFin := datautils.WrapString[C.char](hint)
	defer hintFin()

	state := &inputTextInternalState{
		buf:      datautils.NewStringBuffer(*buf),
		callback: callback,
	}

	defer func() {
		*buf = state.buf.ToGo()
		state.release()
	}()

	stateHandle := cgo.NewHandle(state)
	defer stateHandle.Delete()

	flags |= InputTextFlagsCallbackResize

	return C.igInputTextWithHint(
		labelArg,
		hintArg,
		datautils.ConvertCTypes[*C.char](state.buf.Ptr()),
		C.xulong(len(*buf)+1),
		C.ImGuiInputTextFlags(flags),
		C.ImGuiInputTextCallback(C.generalInputTextCallback),
		unsafe.Pointer(&stateHandle),
	) == C.bool(true)
}

func InputTextMultiline(label string, buf *string, size Vec2, flags InputTextFlags, callback InputTextCallback) bool {
	labelArg, labelFin := datautils.WrapString[C.char](label)
	defer labelFin()

	state := &inputTextInternalState{
		buf:      datautils.NewStringBuffer(*buf),
		callback: callback,
	}

	defer func() {
		*buf = state.buf.ToGo()
		state.release()
	}()

	stateHandle := cgo.NewHandle(state)
	defer stateHandle.Delete()

	flags |= InputTextFlagsCallbackResize

	return C.igInputTextMultiline(
		labelArg,
		datautils.ConvertCTypes[*C.char](state.buf.Ptr()),
		C.xulong(len(*buf)+1),
		size.ToC(),
		C.ImGuiInputTextFlags(flags),
		C.ImGuiInputTextCallback(C.generalInputTextCallback),
		unsafe.Pointer(&stateHandle),
	) == C.bool(true)
}
