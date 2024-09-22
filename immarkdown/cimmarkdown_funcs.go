// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package immarkdown

import (
	"unsafe"

	"github.com/AllenDang/cimgui-go/datautils"
	"github.com/AllenDang/cimgui-go/imgui"
)

// #include "../imgui/extra_types.h"
// #include "cimmarkdown_structs_accessor.h"
// #include "cimmarkdown_wrapper.h"
import "C"

func IsCharInsideWord(c_ rune) bool {
	return C.IsCharInsideWord(C.char(c_)) == C.bool(true)
}

func Markdown(markdown_ string, markdownLength_ uint64, mdConfig_ MarkdownConfig) {
	markdown_Arg, markdown_Fin := datautils.WrapString[C.char](markdown_)
	mdConfig_Arg, mdConfig_Fin := mdConfig_.C()
	C.Markdown(markdown_Arg, C.xulong(markdownLength_), datautils.ConvertCTypes[C.MarkdownConfig](mdConfig_Arg))

	markdown_Fin()
	mdConfig_Fin()
}

func RenderLine(markdown_ string, line_ *Line, textRegion_ *TextRegion, mdConfig_ MarkdownConfig) {
	markdown_Arg, markdown_Fin := datautils.WrapString[C.char](markdown_)
	line_Arg, line_Fin := line_.Handle()
	textRegion_Arg, textRegion_Fin := textRegion_.Handle()
	mdConfig_Arg, mdConfig_Fin := mdConfig_.C()
	C.RenderLine(markdown_Arg, datautils.ConvertCTypes[*C.Line](line_Arg), datautils.ConvertCTypes[*C.TextRegion](textRegion_Arg), datautils.ConvertCTypes[C.MarkdownConfig](mdConfig_Arg))

	markdown_Fin()
	line_Fin()
	textRegion_Fin()
	mdConfig_Fin()
}

func RenderLinkText(self *TextRegion, text_ string, link_ Link, markdown_ string, mdConfig_ MarkdownConfig, linkHoverStart_ []string) bool {
	selfArg, selfFin := self.Handle()
	text_Arg, text_Fin := datautils.WrapString[C.char](text_)
	link_Arg, link_Fin := link_.C()
	markdown_Arg, markdown_Fin := datautils.WrapString[C.char](markdown_)
	mdConfig_Arg, mdConfig_Fin := mdConfig_.C()
	linkHoverStart_Arg, linkHoverStart_Fin := datautils.WrapStringList[C.char](linkHoverStart_)

	defer func() {
		selfFin()
		text_Fin()
		link_Fin()
		markdown_Fin()
		mdConfig_Fin()
		linkHoverStart_Fin()
	}()
	return C.wrap_RenderLinkText(datautils.ConvertCTypes[*C.TextRegion](selfArg), text_Arg, datautils.ConvertCTypes[C.Link](link_Arg), markdown_Arg, datautils.ConvertCTypes[C.MarkdownConfig](mdConfig_Arg), linkHoverStart_Arg) == C.bool(true)
}

// RenderLinkTextWrappedV parameter default value hint:
// bIndentToHere_: false
func RenderLinkTextWrappedV(self *TextRegion, text_ string, link_ Link, markdown_ string, mdConfig_ MarkdownConfig, linkHoverStart_ []string, bIndentToHere_ bool) {
	selfArg, selfFin := self.Handle()
	text_Arg, text_Fin := datautils.WrapString[C.char](text_)
	link_Arg, link_Fin := link_.C()
	markdown_Arg, markdown_Fin := datautils.WrapString[C.char](markdown_)
	mdConfig_Arg, mdConfig_Fin := mdConfig_.C()
	linkHoverStart_Arg, linkHoverStart_Fin := datautils.WrapStringList[C.char](linkHoverStart_)
	C.wrap_RenderLinkTextWrappedV(datautils.ConvertCTypes[*C.TextRegion](selfArg), text_Arg, datautils.ConvertCTypes[C.Link](link_Arg), markdown_Arg, datautils.ConvertCTypes[C.MarkdownConfig](mdConfig_Arg), linkHoverStart_Arg, C.bool(bIndentToHere_))

	selfFin()
	text_Fin()
	link_Fin()
	markdown_Fin()
	mdConfig_Fin()
	linkHoverStart_Fin()
}

func RenderListTextWrapped(self *TextRegion, text_ string) {
	selfArg, selfFin := self.Handle()
	text_Arg, text_Fin := datautils.WrapString[C.char](text_)
	C.wrap_RenderListTextWrapped(datautils.ConvertCTypes[*C.TextRegion](selfArg), text_Arg)

	selfFin()
	text_Fin()
}

// RenderTextWrappedV parameter default value hint:
// bIndentToHere_: false
func RenderTextWrappedV(self *TextRegion, text_ string, bIndentToHere_ bool) {
	selfArg, selfFin := self.Handle()
	text_Arg, text_Fin := datautils.WrapString[C.char](text_)
	C.wrap_RenderTextWrappedV(datautils.ConvertCTypes[*C.TextRegion](selfArg), text_Arg, C.bool(bIndentToHere_))

	selfFin()
	text_Fin()
}

func ResetIndent(self *TextRegion) {
	selfArg, selfFin := self.Handle()
	C.ResetIndent(datautils.ConvertCTypes[*C.TextRegion](selfArg))

	selfFin()
}

func NewTextRegion() *TextRegion {
	return NewTextRegionFromC(C.TextRegion_TextRegion())
}

func (self *TextRegion) Destroy() {
	selfArg, selfFin := self.Handle()
	C.TextRegion_destroy(datautils.ConvertCTypes[*C.TextRegion](selfArg))

	selfFin()
}

func UnderLine(col_ imgui.Color) {
	C.UnderLine(datautils.ConvertCTypes[C.ImColor](col_.ToC()))
}

func RenderLinkTextWrapped(self *TextRegion, text_ string, link_ Link, markdown_ string, mdConfig_ MarkdownConfig, linkHoverStart_ []string) {
	selfArg, selfFin := self.Handle()
	text_Arg, text_Fin := datautils.WrapString[C.char](text_)
	link_Arg, link_Fin := link_.C()
	markdown_Arg, markdown_Fin := datautils.WrapString[C.char](markdown_)
	mdConfig_Arg, mdConfig_Fin := mdConfig_.C()
	linkHoverStart_Arg, linkHoverStart_Fin := datautils.WrapStringList[C.char](linkHoverStart_)
	C.wrap_RenderLinkTextWrapped(datautils.ConvertCTypes[*C.TextRegion](selfArg), text_Arg, datautils.ConvertCTypes[C.Link](link_Arg), markdown_Arg, datautils.ConvertCTypes[C.MarkdownConfig](mdConfig_Arg), linkHoverStart_Arg)

	selfFin()
	text_Fin()
	link_Fin()
	markdown_Fin()
	mdConfig_Fin()
	linkHoverStart_Fin()
}

func RenderTextWrapped(self *TextRegion, text_ string) {
	selfArg, selfFin := self.Handle()
	text_Arg, text_Fin := datautils.WrapString[C.char](text_)
	C.wrap_RenderTextWrapped(datautils.ConvertCTypes[*C.TextRegion](selfArg), text_Arg)

	selfFin()
	text_Fin()
}

func (self Emphasis) SetState(v EmphasisState) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Emphasis_SetState(selfArg, C.EmphasisState(v))
}

func (self *Emphasis) State() EmphasisState {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return EmphasisState(C.wrap_Emphasis_GetState(datautils.ConvertCTypes[*C.Emphasis](selfArg)))
}

func (self Emphasis) SetText(v TextBlock) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Emphasis_SetText(selfArg, datautils.ConvertCTypes[C.TextBlock](vArg))
}

func (self *Emphasis) Text() TextBlock {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return *NewTextBlockFromC(func() *C.TextBlock {
		result := C.wrap_Emphasis_GetText(datautils.ConvertCTypes[*C.Emphasis](selfArg))
		return &result
	}())
}

func (self Emphasis) SetSym(v rune) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Emphasis_SetSym(selfArg, C.char(v))
}

func (self *Emphasis) Sym() rune {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return rune(C.wrap_Emphasis_GetSym(datautils.ConvertCTypes[*C.Emphasis](selfArg)))
}

func (self Line) SetIsHeading(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetIsHeading(selfArg, C.bool(v))
}

func (self *Line) IsHeading() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Line_GetIsHeading(datautils.ConvertCTypes[*C.Line](selfArg)) == C.bool(true)
}

func (self Line) SetIsEmphasis(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetIsEmphasis(selfArg, C.bool(v))
}

func (self *Line) IsEmphasis() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Line_GetIsEmphasis(datautils.ConvertCTypes[*C.Line](selfArg)) == C.bool(true)
}

func (self Line) SetIsUnorderedListStart(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetIsUnorderedListStart(selfArg, C.bool(v))
}

func (self *Line) IsUnorderedListStart() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Line_GetIsUnorderedListStart(datautils.ConvertCTypes[*C.Line](selfArg)) == C.bool(true)
}

func (self Line) SetIsLeadingSpace(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetIsLeadingSpace(selfArg, C.bool(v))
}

func (self *Line) IsLeadingSpace() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Line_GetIsLeadingSpace(datautils.ConvertCTypes[*C.Line](selfArg)) == C.bool(true)
}

func (self Line) SetLeadSpaceCount(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetLeadSpaceCount(selfArg, C.int(v))
}

func (self *Line) LeadSpaceCount() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetLeadSpaceCount(datautils.ConvertCTypes[*C.Line](selfArg)))
}

func (self Line) SetHeadingCount(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetHeadingCount(selfArg, C.int(v))
}

func (self *Line) HeadingCount() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetHeadingCount(datautils.ConvertCTypes[*C.Line](selfArg)))
}

func (self Line) SetEmphasisCount(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetEmphasisCount(selfArg, C.int(v))
}

func (self *Line) EmphasisCount() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetEmphasisCount(datautils.ConvertCTypes[*C.Line](selfArg)))
}

func (self Line) SetLineStart(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetLineStart(selfArg, C.int(v))
}

func (self *Line) LineStart() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetLineStart(datautils.ConvertCTypes[*C.Line](selfArg)))
}

func (self Line) SetLineEnd(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetLineEnd(selfArg, C.int(v))
}

func (self *Line) LineEnd() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetLineEnd(datautils.ConvertCTypes[*C.Line](selfArg)))
}

func (self Line) SetLastRenderPosition(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Line_SetLastRenderPosition(selfArg, C.int(v))
}

func (self *Line) LastRenderPosition() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetLastRenderPosition(datautils.ConvertCTypes[*C.Line](selfArg)))
}

func (self Link) SetState(v LinkState) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Link_SetState(selfArg, C.LinkState(v))
}

func (self *Link) State() LinkState {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return LinkState(C.wrap_Link_GetState(datautils.ConvertCTypes[*C.Link](selfArg)))
}

func (self Link) SetText(v TextBlock) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Link_SetText(selfArg, datautils.ConvertCTypes[C.TextBlock](vArg))
}

func (self *Link) Text() TextBlock {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return *NewTextBlockFromC(func() *C.TextBlock {
		result := C.wrap_Link_GetText(datautils.ConvertCTypes[*C.Link](selfArg))
		return &result
	}())
}

func (self Link) SetUrl(v TextBlock) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Link_SetUrl(selfArg, datautils.ConvertCTypes[C.TextBlock](vArg))
}

func (self *Link) Url() TextBlock {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return *NewTextBlockFromC(func() *C.TextBlock {
		result := C.wrap_Link_GetUrl(datautils.ConvertCTypes[*C.Link](selfArg))
		return &result
	}())
}

func (self Link) SetIsImage(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Link_SetIsImage(selfArg, C.bool(v))
}

func (self *Link) IsImage() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Link_GetIsImage(datautils.ConvertCTypes[*C.Link](selfArg)) == C.bool(true)
}

func (self Link) SetNumbracketsopen(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_Link_SetNum_brackets_open(selfArg, C.int(v))
}

func (self *Link) Numbracketsopen() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Link_GetNum_brackets_open(datautils.ConvertCTypes[*C.Link](selfArg)))
}

func (self MarkdownConfig) SetLinkIcon(v string) {
	vArg, _ := datautils.WrapString[C.char](v)

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetLinkIcon(selfArg, vArg)
}

func (self *MarkdownConfig) LinkIcon() string {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.GoString(C.wrap_MarkdownConfig_GetLinkIcon(datautils.ConvertCTypes[*C.MarkdownConfig](selfArg)))
}

func (self MarkdownConfig) SetHeadingFormats(v *[3]MarkdownHeadingFormat) {
	vArg := make([]C.MarkdownHeadingFormat, len(v))
	for i, vV := range v {
		vVArg, _ := vV.C()
		vArg[i] = datautils.ConvertCTypes[C.MarkdownHeadingFormat](vVArg)
	}

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetHeadingFormats(selfArg, (*C.MarkdownHeadingFormat)(&vArg[0]))

	for i, vV := range vArg {
		(*v)[i] = *NewMarkdownHeadingFormatFromC(func() *C.MarkdownHeadingFormat { result := vV; return &result }())
	}
}

func (self *MarkdownConfig) HeadingFormats() [3]MarkdownHeadingFormat {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() [3]MarkdownHeadingFormat {
		result := [3]MarkdownHeadingFormat{}
		resultMirr := C.wrap_MarkdownConfig_GetHeadingFormats(datautils.ConvertCTypes[*C.MarkdownConfig](selfArg))
		for i := range result {
			result[i] = *NewMarkdownHeadingFormatFromC(func() *C.MarkdownHeadingFormat {
				result := C.cimmarkdown_MarkdownHeadingFormat_GetAtIdx(resultMirr, C.int(i))
				return &result
			}())
		}

		return result
	}()
}

func (self MarkdownConfig) SetUserData(v uintptr) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetUserData(selfArg, C.uintptr_t(v))
}

func (self *MarkdownConfig) UserData() uintptr {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return uintptr(C.wrap_MarkdownConfig_GetUserData(datautils.ConvertCTypes[*C.MarkdownConfig](selfArg)))
}

func (self MarkdownFormatInfo) SetType(v MarkdownFormatType) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownFormatInfo_SetType(selfArg, C.MarkdownFormatType(v))
}

func (self *MarkdownFormatInfo) Type() MarkdownFormatType {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return MarkdownFormatType(C.wrap_MarkdownFormatInfo_GetType(datautils.ConvertCTypes[*C.MarkdownFormatInfo](selfArg)))
}

func (self MarkdownFormatInfo) SetLevel(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownFormatInfo_SetLevel(selfArg, C.int32_t(v))
}

func (self *MarkdownFormatInfo) Level() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_MarkdownFormatInfo_GetLevel(datautils.ConvertCTypes[*C.MarkdownFormatInfo](selfArg)))
}

func (self MarkdownFormatInfo) SetItemHovered(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownFormatInfo_SetItemHovered(selfArg, C.bool(v))
}

func (self *MarkdownFormatInfo) ItemHovered() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownFormatInfo_GetItemHovered(datautils.ConvertCTypes[*C.MarkdownFormatInfo](selfArg)) == C.bool(true)
}

func (self MarkdownFormatInfo) SetConfig(v *MarkdownConfig) {
	vArg, _ := v.Handle()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownFormatInfo_SetConfig(selfArg, datautils.ConvertCTypes[*C.MarkdownConfig](vArg))
}

func (self *MarkdownFormatInfo) Config() *MarkdownConfig {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return NewMarkdownConfigFromC(C.wrap_MarkdownFormatInfo_GetConfig(datautils.ConvertCTypes[*C.MarkdownFormatInfo](selfArg)))
}

func (self MarkdownHeadingFormat) SetFont(v *imgui.Font) {
	vArg, _ := v.Handle()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownHeadingFormat_SetFont(selfArg, datautils.ConvertCTypes[*C.ImFont](vArg))
}

func (self *MarkdownHeadingFormat) Font() *imgui.Font {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return imgui.NewFontFromC(C.wrap_MarkdownHeadingFormat_GetFont(datautils.ConvertCTypes[*C.MarkdownHeadingFormat](selfArg)))
}

func (self MarkdownHeadingFormat) SetSeparator(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownHeadingFormat_SetSeparator(selfArg, C.bool(v))
}

func (self *MarkdownHeadingFormat) Separator() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownHeadingFormat_GetSeparator(datautils.ConvertCTypes[*C.MarkdownHeadingFormat](selfArg)) == C.bool(true)
}

func (self MarkdownImageData) SetIsValid(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetIsValid(selfArg, C.bool(v))
}

func (self *MarkdownImageData) IsValid() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownImageData_GetIsValid(datautils.ConvertCTypes[*C.MarkdownImageData](selfArg)) == C.bool(true)
}

func (self MarkdownImageData) SetUseLinkCallback(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetUseLinkCallback(selfArg, C.bool(v))
}

func (self *MarkdownImageData) UseLinkCallback() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownImageData_GetUseLinkCallback(datautils.ConvertCTypes[*C.MarkdownImageData](selfArg)) == C.bool(true)
}

func (self MarkdownImageData) SetUsertextureid(v imgui.TextureID) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetUser_texture_id(selfArg, datautils.ConvertCTypes[C.ImTextureID](vArg))
}

func (self MarkdownImageData) SetSize(v imgui.Vec2) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetSize(selfArg, datautils.ConvertCTypes[C.ImVec2](v.ToC()))
}

func (self *MarkdownImageData) Size() imgui.Vec2 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec2 {
		out := C.wrap_MarkdownImageData_GetSize(datautils.ConvertCTypes[*C.MarkdownImageData](selfArg))
		return *(&imgui.Vec2{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self MarkdownImageData) SetUv0(v imgui.Vec2) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetUv0(selfArg, datautils.ConvertCTypes[C.ImVec2](v.ToC()))
}

func (self *MarkdownImageData) Uv0() imgui.Vec2 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec2 {
		out := C.wrap_MarkdownImageData_GetUv0(datautils.ConvertCTypes[*C.MarkdownImageData](selfArg))
		return *(&imgui.Vec2{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self MarkdownImageData) SetUv1(v imgui.Vec2) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetUv1(selfArg, datautils.ConvertCTypes[C.ImVec2](v.ToC()))
}

func (self *MarkdownImageData) Uv1() imgui.Vec2 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec2 {
		out := C.wrap_MarkdownImageData_GetUv1(datautils.ConvertCTypes[*C.MarkdownImageData](selfArg))
		return *(&imgui.Vec2{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self MarkdownImageData) SetTintcol(v imgui.Vec4) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetTint_col(selfArg, datautils.ConvertCTypes[C.ImVec4](v.ToC()))
}

func (self *MarkdownImageData) Tintcol() imgui.Vec4 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec4 {
		out := C.wrap_MarkdownImageData_GetTint_col(datautils.ConvertCTypes[*C.MarkdownImageData](selfArg))
		return *(&imgui.Vec4{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self MarkdownImageData) SetBordercol(v imgui.Vec4) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetBorder_col(selfArg, datautils.ConvertCTypes[C.ImVec4](v.ToC()))
}

func (self *MarkdownImageData) Bordercol() imgui.Vec4 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return func() imgui.Vec4 {
		out := C.wrap_MarkdownImageData_GetBorder_col(datautils.ConvertCTypes[*C.MarkdownImageData](selfArg))
		return *(&imgui.Vec4{}).FromC(unsafe.Pointer(&out))
	}()
}

func (self MarkdownLinkCallbackData) SetText(v string) {
	vArg, _ := datautils.WrapString[C.char](v)

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetText(selfArg, vArg)
}

func (self *MarkdownLinkCallbackData) Text() string {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.GoString(C.wrap_MarkdownLinkCallbackData_GetText(datautils.ConvertCTypes[*C.MarkdownLinkCallbackData](selfArg)))
}

func (self MarkdownLinkCallbackData) SetTextLength(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetTextLength(selfArg, C.int(v))
}

func (self *MarkdownLinkCallbackData) TextLength() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_MarkdownLinkCallbackData_GetTextLength(datautils.ConvertCTypes[*C.MarkdownLinkCallbackData](selfArg)))
}

func (self MarkdownLinkCallbackData) SetLink(v string) {
	vArg, _ := datautils.WrapString[C.char](v)

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetLink(selfArg, vArg)
}

func (self *MarkdownLinkCallbackData) Link() string {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.GoString(C.wrap_MarkdownLinkCallbackData_GetLink(datautils.ConvertCTypes[*C.MarkdownLinkCallbackData](selfArg)))
}

func (self MarkdownLinkCallbackData) SetLinkLength(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetLinkLength(selfArg, C.int(v))
}

func (self *MarkdownLinkCallbackData) LinkLength() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_MarkdownLinkCallbackData_GetLinkLength(datautils.ConvertCTypes[*C.MarkdownLinkCallbackData](selfArg)))
}

func (self MarkdownLinkCallbackData) SetUserData(v uintptr) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetUserData(selfArg, C.uintptr_t(v))
}

func (self *MarkdownLinkCallbackData) UserData() uintptr {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return uintptr(C.wrap_MarkdownLinkCallbackData_GetUserData(datautils.ConvertCTypes[*C.MarkdownLinkCallbackData](selfArg)))
}

func (self MarkdownLinkCallbackData) SetIsImage(v bool) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetIsImage(selfArg, C.bool(v))
}

func (self *MarkdownLinkCallbackData) IsImage() bool {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownLinkCallbackData_GetIsImage(datautils.ConvertCTypes[*C.MarkdownLinkCallbackData](selfArg)) == C.bool(true)
}

func (self MarkdownTooltipCallbackData) SetLinkData(v MarkdownLinkCallbackData) {
	vArg, _ := v.C()

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownTooltipCallbackData_SetLinkData(selfArg, datautils.ConvertCTypes[C.MarkdownLinkCallbackData](vArg))
}

func (self *MarkdownTooltipCallbackData) LinkData() MarkdownLinkCallbackData {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return *NewMarkdownLinkCallbackDataFromC(func() *C.MarkdownLinkCallbackData {
		result := C.wrap_MarkdownTooltipCallbackData_GetLinkData(datautils.ConvertCTypes[*C.MarkdownTooltipCallbackData](selfArg))
		return &result
	}())
}

func (self MarkdownTooltipCallbackData) SetLinkIcon(v string) {
	vArg, _ := datautils.WrapString[C.char](v)

	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_MarkdownTooltipCallbackData_SetLinkIcon(selfArg, vArg)
}

func (self *MarkdownTooltipCallbackData) LinkIcon() string {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return C.GoString(C.wrap_MarkdownTooltipCallbackData_GetLinkIcon(datautils.ConvertCTypes[*C.MarkdownTooltipCallbackData](selfArg)))
}

func (self TextBlock) SetStart(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_TextBlock_SetStart(selfArg, C.int(v))
}

func (self *TextBlock) Start() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_TextBlock_GetStart(datautils.ConvertCTypes[*C.TextBlock](selfArg)))
}

func (self TextBlock) SetStop(v int32) {
	selfArg, selfFin := self.Handle()
	defer selfFin()
	C.wrap_TextBlock_SetStop(selfArg, C.int(v))
}

func (self *TextBlock) Stop() int32 {
	selfArg, selfFin := self.Handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_TextBlock_GetStop(datautils.ConvertCTypes[*C.TextBlock](selfArg)))
}
