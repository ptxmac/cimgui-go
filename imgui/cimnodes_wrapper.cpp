// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

#include "cimnodes_wrapper.h"
#include "../cwrappers/cimnodes.h"

void wrap_imnodes_BeginInputAttribute(int id) { imnodes_BeginInputAttribute(id,ImNodesPinShape_CircleFilled); }
void wrap_imnodes_BeginOutputAttribute(int id) { imnodes_BeginOutputAttribute(id,ImNodesPinShape_CircleFilled); }
void wrap_imnodes_DestroyContext() { imnodes_DestroyContext(0); }
bool wrap_imnodes_IsAnyAttributeActive() { return imnodes_IsAnyAttributeActive(0); }
bool wrap_imnodes_IsLinkCreated_BoolPtr(int* started_at_attribute_id,int* ended_at_attribute_id) { return imnodes_IsLinkCreated_BoolPtr(started_at_attribute_id,ended_at_attribute_id,0); }
bool wrap_imnodes_IsLinkCreated_IntPtr(int* started_at_node_id,int* started_at_attribute_id,int* ended_at_node_id,int* ended_at_attribute_id) { return imnodes_IsLinkCreated_IntPtr(started_at_node_id,started_at_attribute_id,ended_at_node_id,ended_at_attribute_id,0); }
bool wrap_imnodes_IsLinkDropped() { return imnodes_IsLinkDropped(0,true); }
void wrap_imnodes_MiniMap() { imnodes_MiniMap(0.2f,ImNodesMiniMapLocation_TopLeft,0,0); }
void wrap_imnodes_PopStyleVar() { imnodes_PopStyleVar(1); }
const char* wrap_imnodes_SaveCurrentEditorStateToIniString() { return imnodes_SaveCurrentEditorStateToIniString(0); }
const char* wrap_imnodes_SaveEditorStateToIniString(const ImNodesEditorContext* editor) { return imnodes_SaveEditorStateToIniString(editor,0); }
void wrap_imnodes_StyleColorsClassic() { imnodes_StyleColorsClassic(0); }
void wrap_imnodes_StyleColorsDark() { imnodes_StyleColorsDark(0); }
void wrap_imnodes_StyleColorsLight() { imnodes_StyleColorsLight(0); }
