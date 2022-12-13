#pragma once
#ifdef __cplusplus
extern "C" {
#endif
#include "fabgen.h"
#include <memory.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

extern int GetLogLevel(const int id);
extern int GetSeekMode(const int id);
extern int GetDirEntryType(const int id);
extern uint8_t GetRotationOrder(const int id);
extern uint8_t GetAxis(const int id);
extern uint8_t GetVisibility(const int id);
extern uint8_t GetMonitorRotation(const int id);
extern int GetWindowVisibility(const int id);
extern int GetPictureFormat(const int id);
extern int GetRendererType(const int id);
extern int GetTextureFormat(const int id);
extern int GetBackbufferRatio(const int id);
extern int GetViewMode(const int id);
extern int GetAttrib(const int id);
extern int GetAttribType(const int id);
extern int GetFaceCulling(const int id);
extern int GetDepthTest(const int id);
extern int GetBlendMode(const int id);
extern int GetForwardPipelineLightType(const int id);
extern int GetForwardPipelineShadowType(const int id);
extern int GetDrawTextHAlign(const int id);
extern int GetDrawTextVAlign(const int id);
extern int GetAnimLoopMode(const int id);
extern unsigned char GetEasing(const int id);
extern int GetLightType(const int id);
extern int GetLightShadowType(const int id);
extern uint8_t GetRigidBodyType(const int id);
extern uint8_t GetCollisionEventTrackingMode(const int id);
extern uint8_t GetCollisionType(const int id);
extern int GetNodeComponentIdx(const int id);
extern int GetSceneForwardPipelinePass(const int id);
extern int GetForwardPipelineAAADebugBuffer(const int id);
extern int GetMouseButton(const int id);
extern int GetKey(const int id);
extern int GetGamepadAxes(const int id);
extern int GetGamepadButton(const int id);
extern int GetVRControllerButton(const int id);
extern int GetImGuiWindowFlags(const int id);
extern int GetImGuiPopupFlags(const int id);
extern int GetImGuiCond(const int id);
extern int GetImGuiMouseButton(const int id);
extern int GetImGuiHoveredFlags(const int id);
extern int GetImGuiFocusedFlags(const int id);
extern int GetImGuiColorEditFlags(const int id);
extern int GetImGuiInputTextFlags(const int id);
extern int GetImGuiTreeNodeFlags(const int id);
extern int GetImGuiSelectableFlags(const int id);
extern int GetImGuiCol(const int id);
extern int GetImGuiStyleVar(const int id);
extern int GetImDrawFlags(const int id);
extern int GetImGuiComboFlags(const int id);
extern int GetAudioFrameFormat(const int id);
extern int GetSourceRepeat(const int id);
extern int GetSourceState(const int id);
extern int GetOpenVRAA(const int id);
extern int GetOpenXRExtensions(const int id);
extern int GetOpenXRAA(const int id);
extern int GetHandsSide(const int id);
extern int GetXrHandJoint(const int id);
extern int GetVideoFrameFormat(const int id);
typedef void *HarfangVoidPointer;
typedef void *HarfangIntList;
typedef void *HarfangUint16TList;
typedef void *HarfangUint32TList;
typedef void *HarfangStringList;
typedef void *HarfangFile;
typedef void *HarfangData;
typedef void *HarfangDirEntry;
typedef void *HarfangDirEntryList;
typedef void *HarfangVec3;
typedef void *HarfangVec4;
typedef void *HarfangMat3;
typedef void *HarfangMat4;
typedef void *HarfangMat44;
typedef void *HarfangQuaternion;
typedef void *HarfangMinMax;
typedef void *HarfangVec2;
typedef void *HarfangVec2List;
typedef void *HarfangIVec2;
typedef void *HarfangIVec2List;
typedef void *HarfangVec4List;
typedef void *HarfangMat4List;
typedef void *HarfangVec3List;
typedef void *HarfangRect;
typedef void *HarfangIntRect;
typedef void *HarfangFrustum;
typedef void *HarfangMonitorMode;
typedef void *HarfangMonitorModeList;
typedef void *HarfangMonitor;
typedef void *HarfangMonitorList;
typedef void *HarfangWindow;
typedef void *HarfangColor;
typedef void *HarfangColorList;
typedef void *HarfangPicture;
typedef void *HarfangFrameBufferHandle;
typedef void *HarfangVertexLayout;
typedef void *HarfangProgramHandle;
typedef void *HarfangTextureInfo;
typedef void *HarfangModelRef;
typedef void *HarfangTextureRef;
typedef void *HarfangMaterialRef;
typedef void *HarfangPipelineProgramRef;
typedef void *HarfangTexture;
typedef void *HarfangUniformSetValue;
typedef void *HarfangUniformSetValueList;
typedef void *HarfangUniformSetTexture;
typedef void *HarfangUniformSetTextureList;
typedef void *HarfangPipelineProgram;
typedef void *HarfangViewState;
typedef void *HarfangMaterial;
typedef void *HarfangMaterialList;
typedef void *HarfangRenderState;
typedef void *HarfangModel;
typedef void (*HarfangFunctionReturningVoidTakingUint16T)(uint16_t);
typedef void *HarfangPipelineResources;
typedef void *HarfangFrameBuffer;
typedef void *HarfangVertices;
typedef void *HarfangPipeline;
typedef void *HarfangPipelineInfo;
typedef void *HarfangForwardPipeline;
typedef void *HarfangForwardPipelineLight;
typedef void *HarfangForwardPipelineLightList;
typedef void *HarfangForwardPipelineLights;
typedef void *HarfangForwardPipelineFog;
typedef void *HarfangFont;
typedef void *HarfangJSON;
typedef void *HarfangLuaObject;
typedef void *HarfangLuaObjectList;
typedef void *HarfangSceneAnimRef;
typedef void *HarfangSceneAnimRefList;
typedef void *HarfangScenePlayAnimRef;
typedef void *HarfangScenePlayAnimRefList;
typedef void *HarfangScene;
typedef void *HarfangSceneView;
typedef void *HarfangNode;
typedef void *HarfangTransformTRS;
typedef void *HarfangTransform;
typedef void *HarfangCameraZRange;
typedef void *HarfangCamera;
typedef void *HarfangObject;
typedef void *HarfangLight;
typedef void *HarfangContact;
typedef void *HarfangContactList;
typedef void *HarfangRigidBody;
typedef void *HarfangCollision;
typedef void *HarfangInstance;
typedef void *HarfangScript;
typedef void *HarfangScriptList;
typedef void *HarfangNodeList;
typedef void *HarfangRaycastOut;
typedef void *HarfangRaycastOutList;
typedef void (*HarfangFunctionReturningVoidTakingTimeNs)(int64_t);
typedef void *HarfangTimeCallbackConnection;
typedef void *HarfangSignalReturningVoidTakingTimeNs;
typedef void *HarfangCanvas;
typedef void *HarfangEnvironment;
typedef void *HarfangSceneForwardPipelinePassViewId;
typedef void *HarfangSceneForwardPipelineRenderData;
typedef void *HarfangForwardPipelineAAAConfig;
typedef void *HarfangForwardPipelineAAA;
typedef void *HarfangNodePairContacts;
typedef void *HarfangBtGeneric6DofConstraint;
typedef void *HarfangSceneBullet3Physics;
typedef void (*HarfangFunctionReturningVoidTakingSceneBullet3PhysicsRefTimeNs)(HarfangSceneBullet3Physics, int64_t);
typedef void *HarfangSceneLuaVM;
typedef void *HarfangSceneClocks;
typedef void *HarfangMouseState;
typedef void *HarfangMouse;
typedef void *HarfangKeyboardState;
typedef void *HarfangKeyboard;
typedef void (*HarfangFunctionReturningVoidTakingConstCharPtr)(const char *);
typedef void *HarfangTextInputCallbackConnection;
typedef void *HarfangSignalReturningVoidTakingConstCharPtr;
typedef void *HarfangGamepadState;
typedef void *HarfangGamepad;
typedef void *HarfangJoystickState;
typedef void *HarfangJoystick;
typedef void *HarfangVRControllerState;
typedef void *HarfangVRController;
typedef void *HarfangVRGenericTrackerState;
typedef void *HarfangVRGenericTracker;
typedef void *HarfangDearImguiContext;
typedef void *HarfangImFont;
typedef void *HarfangImDrawList;
typedef void *HarfangFileFilter;
typedef void *HarfangFileFilterList;
typedef void *HarfangStereoSourceState;
typedef void *HarfangSpatializedSourceState;
typedef void *HarfangOpenVREye;
typedef void *HarfangOpenVREyeFrameBuffer;
typedef void *HarfangOpenVRState;
typedef void *HarfangOpenXREyeFrameBuffer;
typedef void *HarfangOpenXREyeFrameBufferList;
typedef void *HarfangOpenXRFrameInfo;
typedef void (*HarfangFunctionReturningVoidTakingMat4Ptr)(HarfangMat4);
typedef uint16_t (*HarfangFunctionReturningUint16TTakingRectOfIntPtrViewStatePtrUint16TPtrFrameBufferHandlePtr)(
	HarfangIntRect, HarfangViewState, uint16_t *, HarfangFrameBufferHandle);
typedef void *HarfangSRanipalEyeState;
typedef void *HarfangSRanipalState;
typedef void *HarfangVertex;
typedef void *HarfangModelBuilder;
typedef void *HarfangGeometry;
typedef void *HarfangGeometryBuilder;
typedef void *HarfangIsoSurface;
typedef void *HarfangBloom;
typedef void *HarfangSAO;
typedef void *HarfangProfilerFrame;
typedef void *HarfangIVideoStreamer;
extern void HarfangVoidPointerFree(HarfangVoidPointer);
extern int HarfangIntListGetOperator(HarfangIntList h, int id);
extern void HarfangIntListSetOperator(HarfangIntList h, int id, int v);
extern int HarfangIntListLenOperator(HarfangIntList h);
extern HarfangIntList HarfangConstructorIntList();
extern HarfangIntList HarfangConstructorIntListWithSequence(size_t sequenceToCSize, int *sequenceToCBuf);
extern void HarfangIntListFree(HarfangIntList);
extern void HarfangClearIntList(HarfangIntList this_);
extern void HarfangReserveIntList(HarfangIntList this_, size_t size);
extern void HarfangPushBackIntList(HarfangIntList this_, int v);
extern size_t HarfangSizeIntList(HarfangIntList this_);
extern int HarfangAtIntList(HarfangIntList this_, size_t idx);
extern uint16_t HarfangUint16TListGetOperator(HarfangUint16TList h, int id);
extern void HarfangUint16TListSetOperator(HarfangUint16TList h, int id, uint16_t v);
extern int HarfangUint16TListLenOperator(HarfangUint16TList h);
extern HarfangUint16TList HarfangConstructorUint16TList();
extern HarfangUint16TList HarfangConstructorUint16TListWithSequence(size_t sequenceToCSize, uint16_t *sequenceToCBuf);
extern void HarfangUint16TListFree(HarfangUint16TList);
extern void HarfangClearUint16TList(HarfangUint16TList this_);
extern void HarfangReserveUint16TList(HarfangUint16TList this_, size_t size);
extern void HarfangPushBackUint16TList(HarfangUint16TList this_, uint16_t v);
extern size_t HarfangSizeUint16TList(HarfangUint16TList this_);
extern uint16_t HarfangAtUint16TList(HarfangUint16TList this_, size_t idx);
extern uint32_t HarfangUint32TListGetOperator(HarfangUint32TList h, int id);
extern void HarfangUint32TListSetOperator(HarfangUint32TList h, int id, uint32_t v);
extern int HarfangUint32TListLenOperator(HarfangUint32TList h);
extern HarfangUint32TList HarfangConstructorUint32TList();
extern HarfangUint32TList HarfangConstructorUint32TListWithSequence(size_t sequenceToCSize, uint32_t *sequenceToCBuf);
extern void HarfangUint32TListFree(HarfangUint32TList);
extern void HarfangClearUint32TList(HarfangUint32TList this_);
extern void HarfangReserveUint32TList(HarfangUint32TList this_, size_t size);
extern void HarfangPushBackUint32TList(HarfangUint32TList this_, uint32_t v);
extern size_t HarfangSizeUint32TList(HarfangUint32TList this_);
extern uint32_t HarfangAtUint32TList(HarfangUint32TList this_, size_t idx);
extern const char *HarfangStringListGetOperator(HarfangStringList h, int id);
extern void HarfangStringListSetOperator(HarfangStringList h, int id, const char *v);
extern int HarfangStringListLenOperator(HarfangStringList h);
extern HarfangStringList HarfangConstructorStringList();
extern HarfangStringList HarfangConstructorStringListWithSequence(size_t sequenceToCSize, const char **sequenceToCBuf);
extern void HarfangStringListFree(HarfangStringList);
extern void HarfangClearStringList(HarfangStringList this_);
extern void HarfangReserveStringList(HarfangStringList this_, size_t size);
extern void HarfangPushBackStringList(HarfangStringList this_, const char *v);
extern size_t HarfangSizeStringList(HarfangStringList this_);
extern const char *HarfangAtStringList(HarfangStringList this_, size_t idx);
extern void HarfangFileFree(HarfangFile);
extern HarfangData HarfangConstructorData();
extern void HarfangDataFree(HarfangData);
extern size_t HarfangGetSizeData(HarfangData this_);
extern void HarfangRewindData(HarfangData this_);
extern int HarfangDirEntryGetType(HarfangDirEntry h);
extern void HarfangDirEntrySetType(HarfangDirEntry h, int v);
extern const char *HarfangDirEntryGetName(HarfangDirEntry h);
extern void HarfangDirEntrySetName(HarfangDirEntry h, const char *v);
extern void HarfangDirEntryFree(HarfangDirEntry);
extern HarfangDirEntry HarfangDirEntryListGetOperator(HarfangDirEntryList h, int id);
extern void HarfangDirEntryListSetOperator(HarfangDirEntryList h, int id, HarfangDirEntry v);
extern int HarfangDirEntryListLenOperator(HarfangDirEntryList h);
extern HarfangDirEntryList HarfangConstructorDirEntryList();
extern HarfangDirEntryList HarfangConstructorDirEntryListWithSequence(size_t sequenceToCSize, HarfangDirEntry *sequenceToCBuf);
extern void HarfangDirEntryListFree(HarfangDirEntryList);
extern void HarfangClearDirEntryList(HarfangDirEntryList this_);
extern void HarfangReserveDirEntryList(HarfangDirEntryList this_, size_t size);
extern void HarfangPushBackDirEntryList(HarfangDirEntryList this_, HarfangDirEntry v);
extern size_t HarfangSizeDirEntryList(HarfangDirEntryList this_);
extern HarfangDirEntry HarfangAtDirEntryList(HarfangDirEntryList this_, size_t idx);
extern HarfangVec3 HarfangVec3GetZero();
extern HarfangVec3 HarfangVec3GetOne();
extern HarfangVec3 HarfangVec3GetLeft();
extern HarfangVec3 HarfangVec3GetRight();
extern HarfangVec3 HarfangVec3GetUp();
extern HarfangVec3 HarfangVec3GetDown();
extern HarfangVec3 HarfangVec3GetFront();
extern HarfangVec3 HarfangVec3GetBack();
extern float HarfangVec3GetX(HarfangVec3 h);
extern void HarfangVec3SetX(HarfangVec3 h, float v);
extern float HarfangVec3GetY(HarfangVec3 h);
extern void HarfangVec3SetY(HarfangVec3 h, float v);
extern float HarfangVec3GetZ(HarfangVec3 h);
extern void HarfangVec3SetZ(HarfangVec3 h, float v);
extern HarfangVec3 HarfangConstructorVec3();
extern HarfangVec3 HarfangConstructorVec3WithXYZ(float x, float y, float z);
extern HarfangVec3 HarfangConstructorVec3WithV(const HarfangVec2 v);
extern HarfangVec3 HarfangConstructorVec3WithIVec2V(const HarfangIVec2 v);
extern HarfangVec3 HarfangConstructorVec3WithVec3V(const HarfangVec3 v);
extern HarfangVec3 HarfangConstructorVec3WithVec4V(const HarfangVec4 v);
extern void HarfangVec3Free(HarfangVec3);
extern HarfangVec3 HarfangAddVec3(HarfangVec3 this_, HarfangVec3 v);
extern HarfangVec3 HarfangAddVec3WithK(HarfangVec3 this_, float k);
extern HarfangVec3 HarfangSubVec3(HarfangVec3 this_, HarfangVec3 v);
extern HarfangVec3 HarfangSubVec3WithK(HarfangVec3 this_, float k);
extern HarfangVec3 HarfangDivVec3(HarfangVec3 this_, HarfangVec3 v);
extern HarfangVec3 HarfangDivVec3WithK(HarfangVec3 this_, float k);
extern HarfangVec3 HarfangMulVec3(HarfangVec3 this_, const HarfangVec3 v);
extern HarfangVec3 HarfangMulVec3WithK(HarfangVec3 this_, float k);
extern void HarfangInplaceAddVec3(HarfangVec3 this_, HarfangVec3 v);
extern void HarfangInplaceAddVec3WithK(HarfangVec3 this_, float k);
extern void HarfangInplaceSubVec3(HarfangVec3 this_, HarfangVec3 v);
extern void HarfangInplaceSubVec3WithK(HarfangVec3 this_, float k);
extern void HarfangInplaceMulVec3(HarfangVec3 this_, HarfangVec3 v);
extern void HarfangInplaceMulVec3WithK(HarfangVec3 this_, float k);
extern void HarfangInplaceDivVec3(HarfangVec3 this_, HarfangVec3 v);
extern void HarfangInplaceDivVec3WithK(HarfangVec3 this_, float k);
extern bool HarfangEqVec3(HarfangVec3 this_, const HarfangVec3 v);
extern bool HarfangNeVec3(HarfangVec3 this_, const HarfangVec3 v);
extern void HarfangSetVec3(HarfangVec3 this_, float x, float y, float z);
extern float HarfangVec4GetX(HarfangVec4 h);
extern void HarfangVec4SetX(HarfangVec4 h, float v);
extern float HarfangVec4GetY(HarfangVec4 h);
extern void HarfangVec4SetY(HarfangVec4 h, float v);
extern float HarfangVec4GetZ(HarfangVec4 h);
extern void HarfangVec4SetZ(HarfangVec4 h, float v);
extern float HarfangVec4GetW(HarfangVec4 h);
extern void HarfangVec4SetW(HarfangVec4 h, float v);
extern HarfangVec4 HarfangConstructorVec4();
extern HarfangVec4 HarfangConstructorVec4WithXYZ(float x, float y, float z);
extern HarfangVec4 HarfangConstructorVec4WithXYZW(float x, float y, float z, float w);
extern HarfangVec4 HarfangConstructorVec4WithV(const HarfangVec2 v);
extern HarfangVec4 HarfangConstructorVec4WithIVec2V(const HarfangIVec2 v);
extern HarfangVec4 HarfangConstructorVec4WithVec3V(const HarfangVec3 v);
extern HarfangVec4 HarfangConstructorVec4WithVec4V(const HarfangVec4 v);
extern void HarfangVec4Free(HarfangVec4);
extern HarfangVec4 HarfangAddVec4(HarfangVec4 this_, HarfangVec4 v);
extern HarfangVec4 HarfangAddVec4WithK(HarfangVec4 this_, float k);
extern HarfangVec4 HarfangSubVec4(HarfangVec4 this_, HarfangVec4 v);
extern HarfangVec4 HarfangSubVec4WithK(HarfangVec4 this_, float k);
extern HarfangVec4 HarfangDivVec4(HarfangVec4 this_, HarfangVec4 v);
extern HarfangVec4 HarfangDivVec4WithK(HarfangVec4 this_, float k);
extern HarfangVec4 HarfangMulVec4(HarfangVec4 this_, HarfangVec4 v);
extern HarfangVec4 HarfangMulVec4WithK(HarfangVec4 this_, float k);
extern void HarfangInplaceAddVec4(HarfangVec4 this_, HarfangVec4 v);
extern void HarfangInplaceAddVec4WithK(HarfangVec4 this_, float k);
extern void HarfangInplaceSubVec4(HarfangVec4 this_, HarfangVec4 v);
extern void HarfangInplaceSubVec4WithK(HarfangVec4 this_, float k);
extern void HarfangInplaceMulVec4(HarfangVec4 this_, HarfangVec4 v);
extern void HarfangInplaceMulVec4WithK(HarfangVec4 this_, float k);
extern void HarfangInplaceDivVec4(HarfangVec4 this_, HarfangVec4 v);
extern void HarfangInplaceDivVec4WithK(HarfangVec4 this_, float k);
extern void HarfangSetVec4(HarfangVec4 this_, float x, float y, float z);
extern void HarfangSetVec4WithW(HarfangVec4 this_, float x, float y, float z, float w);
extern HarfangMat3 HarfangMat3GetZero();
extern HarfangMat3 HarfangMat3GetIdentity();
extern HarfangMat3 HarfangConstructorMat3();
extern HarfangMat3 HarfangConstructorMat3WithM(const HarfangMat4 m);
extern HarfangMat3 HarfangConstructorMat3WithXYZ(const HarfangVec3 x, const HarfangVec3 y, const HarfangVec3 z);
extern void HarfangMat3Free(HarfangMat3);
extern HarfangMat3 HarfangAddMat3(HarfangMat3 this_, HarfangMat3 m);
extern HarfangMat3 HarfangSubMat3(HarfangMat3 this_, HarfangMat3 m);
extern HarfangMat3 HarfangMulMat3(HarfangMat3 this_, const float v);
extern HarfangVec2 HarfangMulMat3WithV(HarfangMat3 this_, const HarfangVec2 v);
extern HarfangVec3 HarfangMulMat3WithVec3V(HarfangMat3 this_, const HarfangVec3 v);
extern HarfangVec4 HarfangMulMat3WithVec4V(HarfangMat3 this_, const HarfangVec4 v);
extern HarfangMat3 HarfangMulMat3WithM(HarfangMat3 this_, const HarfangMat3 m);
extern void HarfangInplaceAddMat3(HarfangMat3 this_, const HarfangMat3 m);
extern void HarfangInplaceSubMat3(HarfangMat3 this_, const HarfangMat3 m);
extern void HarfangInplaceMulMat3(HarfangMat3 this_, const float k);
extern void HarfangInplaceMulMat3WithM(HarfangMat3 this_, const HarfangMat3 m);
extern bool HarfangEqMat3(HarfangMat3 this_, const HarfangMat3 m);
extern bool HarfangNeMat3(HarfangMat3 this_, const HarfangMat3 m);
extern HarfangMat4 HarfangMat4GetZero();
extern HarfangMat4 HarfangMat4GetIdentity();
extern HarfangMat4 HarfangConstructorMat4();
extern HarfangMat4 HarfangConstructorMat4WithM(const HarfangMat4 m);
extern HarfangMat4 HarfangConstructorMat4WithM00M10M20M01M11M21M02M12M22M03M13M23(
	float m00, float m10, float m20, float m01, float m11, float m21, float m02, float m12, float m22, float m03, float m13, float m23);
extern HarfangMat4 HarfangConstructorMat4WithMat3M(const HarfangMat3 m);
extern void HarfangMat4Free(HarfangMat4);
extern HarfangMat4 HarfangAddMat4(HarfangMat4 this_, HarfangMat4 m);
extern HarfangMat4 HarfangSubMat4(HarfangMat4 this_, HarfangMat4 m);
extern HarfangMat4 HarfangMulMat4(HarfangMat4 this_, const float v);
extern HarfangMat4 HarfangMulMat4WithM(HarfangMat4 this_, const HarfangMat4 m);
extern HarfangVec3 HarfangMulMat4WithV(HarfangMat4 this_, const HarfangVec3 v);
extern HarfangVec4 HarfangMulMat4WithVec4V(HarfangMat4 this_, const HarfangVec4 v);
extern HarfangMat44 HarfangMulMat4WithMat44M(HarfangMat4 this_, const HarfangMat44 m);
extern HarfangMinMax HarfangMulMat4WithMinmax(HarfangMat4 this_, const HarfangMinMax minmax);
extern bool HarfangEqMat4(HarfangMat4 this_, const HarfangMat4 m);
extern bool HarfangNeMat4(HarfangMat4 this_, const HarfangMat4 m);
extern HarfangMat44 HarfangMat44GetZero();
extern HarfangMat44 HarfangMat44GetIdentity();
extern HarfangMat44 HarfangConstructorMat44();
extern HarfangMat44 HarfangConstructorMat44WithM00M10M20M30M01M11M21M31M02M12M22M32M03M13M23M33(float m00, float m10, float m20, float m30, float m01,
	float m11, float m21, float m31, float m02, float m12, float m22, float m32, float m03, float m13, float m23, float m33);
extern void HarfangMat44Free(HarfangMat44);
extern HarfangMat44 HarfangMulMat44(HarfangMat44 this_, const HarfangMat4 m);
extern HarfangMat44 HarfangMulMat44WithM(HarfangMat44 this_, const HarfangMat44 m);
extern HarfangVec3 HarfangMulMat44WithV(HarfangMat44 this_, const HarfangVec3 v);
extern HarfangVec4 HarfangMulMat44WithVec4V(HarfangMat44 this_, const HarfangVec4 v);
extern float HarfangQuaternionGetX(HarfangQuaternion h);
extern void HarfangQuaternionSetX(HarfangQuaternion h, float v);
extern float HarfangQuaternionGetY(HarfangQuaternion h);
extern void HarfangQuaternionSetY(HarfangQuaternion h, float v);
extern float HarfangQuaternionGetZ(HarfangQuaternion h);
extern void HarfangQuaternionSetZ(HarfangQuaternion h, float v);
extern float HarfangQuaternionGetW(HarfangQuaternion h);
extern void HarfangQuaternionSetW(HarfangQuaternion h, float v);
extern HarfangQuaternion HarfangConstructorQuaternion();
extern HarfangQuaternion HarfangConstructorQuaternionWithXYZW(float x, float y, float z, float w);
extern HarfangQuaternion HarfangConstructorQuaternionWithQ(const HarfangQuaternion q);
extern void HarfangQuaternionFree(HarfangQuaternion);
extern HarfangQuaternion HarfangAddQuaternion(HarfangQuaternion this_, float v);
extern HarfangQuaternion HarfangAddQuaternionWithQ(HarfangQuaternion this_, HarfangQuaternion q);
extern HarfangQuaternion HarfangSubQuaternion(HarfangQuaternion this_, float v);
extern HarfangQuaternion HarfangSubQuaternionWithQ(HarfangQuaternion this_, HarfangQuaternion q);
extern HarfangQuaternion HarfangMulQuaternion(HarfangQuaternion this_, float v);
extern HarfangQuaternion HarfangMulQuaternionWithQ(HarfangQuaternion this_, HarfangQuaternion q);
extern HarfangQuaternion HarfangDivQuaternion(HarfangQuaternion this_, float v);
extern void HarfangInplaceAddQuaternion(HarfangQuaternion this_, float v);
extern void HarfangInplaceAddQuaternionWithQ(HarfangQuaternion this_, const HarfangQuaternion q);
extern void HarfangInplaceSubQuaternion(HarfangQuaternion this_, float v);
extern void HarfangInplaceSubQuaternionWithQ(HarfangQuaternion this_, const HarfangQuaternion q);
extern void HarfangInplaceMulQuaternion(HarfangQuaternion this_, float v);
extern void HarfangInplaceMulQuaternionWithQ(HarfangQuaternion this_, const HarfangQuaternion q);
extern void HarfangInplaceDivQuaternion(HarfangQuaternion this_, float v);
extern HarfangVec3 HarfangMinMaxGetMn(HarfangMinMax h);
extern void HarfangMinMaxSetMn(HarfangMinMax h, HarfangVec3 v);
extern HarfangVec3 HarfangMinMaxGetMx(HarfangMinMax h);
extern void HarfangMinMaxSetMx(HarfangMinMax h, HarfangVec3 v);
extern HarfangMinMax HarfangConstructorMinMax();
extern HarfangMinMax HarfangConstructorMinMaxWithMinMax(const HarfangVec3 min, const HarfangVec3 max);
extern void HarfangMinMaxFree(HarfangMinMax);
extern bool HarfangEqMinMax(HarfangMinMax this_, const HarfangMinMax minmax);
extern bool HarfangNeMinMax(HarfangMinMax this_, const HarfangMinMax minmax);
extern HarfangVec2 HarfangVec2GetZero();
extern HarfangVec2 HarfangVec2GetOne();
extern float HarfangVec2GetX(HarfangVec2 h);
extern void HarfangVec2SetX(HarfangVec2 h, float v);
extern float HarfangVec2GetY(HarfangVec2 h);
extern void HarfangVec2SetY(HarfangVec2 h, float v);
extern HarfangVec2 HarfangConstructorVec2();
extern HarfangVec2 HarfangConstructorVec2WithXY(float x, float y);
extern HarfangVec2 HarfangConstructorVec2WithV(const HarfangVec2 v);
extern HarfangVec2 HarfangConstructorVec2WithVec3V(const HarfangVec3 v);
extern HarfangVec2 HarfangConstructorVec2WithVec4V(const HarfangVec4 v);
extern void HarfangVec2Free(HarfangVec2);
extern HarfangVec2 HarfangAddVec2(HarfangVec2 this_, const HarfangVec2 v);
extern HarfangVec2 HarfangAddVec2WithK(HarfangVec2 this_, const float k);
extern HarfangVec2 HarfangSubVec2(HarfangVec2 this_, const HarfangVec2 v);
extern HarfangVec2 HarfangSubVec2WithK(HarfangVec2 this_, const float k);
extern HarfangVec2 HarfangDivVec2(HarfangVec2 this_, const HarfangVec2 v);
extern HarfangVec2 HarfangDivVec2WithK(HarfangVec2 this_, const float k);
extern HarfangVec2 HarfangMulVec2(HarfangVec2 this_, const HarfangVec2 v);
extern HarfangVec2 HarfangMulVec2WithK(HarfangVec2 this_, const float k);
extern void HarfangInplaceAddVec2(HarfangVec2 this_, const HarfangVec2 v);
extern void HarfangInplaceAddVec2WithK(HarfangVec2 this_, const float k);
extern void HarfangInplaceSubVec2(HarfangVec2 this_, const HarfangVec2 v);
extern void HarfangInplaceSubVec2WithK(HarfangVec2 this_, const float k);
extern void HarfangInplaceMulVec2(HarfangVec2 this_, const HarfangVec2 v);
extern void HarfangInplaceMulVec2WithK(HarfangVec2 this_, const float k);
extern void HarfangInplaceDivVec2(HarfangVec2 this_, const HarfangVec2 v);
extern void HarfangInplaceDivVec2WithK(HarfangVec2 this_, const float k);
extern void HarfangSetVec2(HarfangVec2 this_, float x, float y);
extern HarfangVec2 HarfangVec2ListGetOperator(HarfangVec2List h, int id);
extern void HarfangVec2ListSetOperator(HarfangVec2List h, int id, HarfangVec2 v);
extern int HarfangVec2ListLenOperator(HarfangVec2List h);
extern HarfangVec2List HarfangConstructorVec2List();
extern HarfangVec2List HarfangConstructorVec2ListWithSequence(size_t sequenceToCSize, HarfangVec2 *sequenceToCBuf);
extern void HarfangVec2ListFree(HarfangVec2List);
extern void HarfangClearVec2List(HarfangVec2List this_);
extern void HarfangReserveVec2List(HarfangVec2List this_, size_t size);
extern void HarfangPushBackVec2List(HarfangVec2List this_, HarfangVec2 v);
extern size_t HarfangSizeVec2List(HarfangVec2List this_);
extern HarfangVec2 HarfangAtVec2List(HarfangVec2List this_, size_t idx);
extern HarfangIVec2 HarfangIVec2GetZero();
extern HarfangIVec2 HarfangIVec2GetOne();
extern int HarfangIVec2GetX(HarfangIVec2 h);
extern void HarfangIVec2SetX(HarfangIVec2 h, int v);
extern int HarfangIVec2GetY(HarfangIVec2 h);
extern void HarfangIVec2SetY(HarfangIVec2 h, int v);
extern HarfangIVec2 HarfangConstructorIVec2();
extern HarfangIVec2 HarfangConstructorIVec2WithXY(int x, int y);
extern HarfangIVec2 HarfangConstructorIVec2WithV(const HarfangIVec2 v);
extern HarfangIVec2 HarfangConstructorIVec2WithVec3V(const HarfangVec3 v);
extern HarfangIVec2 HarfangConstructorIVec2WithVec4V(const HarfangVec4 v);
extern void HarfangIVec2Free(HarfangIVec2);
extern HarfangIVec2 HarfangAddIVec2(HarfangIVec2 this_, const HarfangIVec2 v);
extern HarfangIVec2 HarfangAddIVec2WithK(HarfangIVec2 this_, const int k);
extern HarfangIVec2 HarfangSubIVec2(HarfangIVec2 this_, const HarfangIVec2 v);
extern HarfangIVec2 HarfangSubIVec2WithK(HarfangIVec2 this_, const int k);
extern HarfangIVec2 HarfangDivIVec2(HarfangIVec2 this_, const HarfangIVec2 v);
extern HarfangIVec2 HarfangDivIVec2WithK(HarfangIVec2 this_, const int k);
extern HarfangIVec2 HarfangMulIVec2(HarfangIVec2 this_, const HarfangIVec2 v);
extern HarfangIVec2 HarfangMulIVec2WithK(HarfangIVec2 this_, const int k);
extern void HarfangInplaceAddIVec2(HarfangIVec2 this_, const HarfangIVec2 v);
extern void HarfangInplaceAddIVec2WithK(HarfangIVec2 this_, const int k);
extern void HarfangInplaceSubIVec2(HarfangIVec2 this_, const HarfangIVec2 v);
extern void HarfangInplaceSubIVec2WithK(HarfangIVec2 this_, const int k);
extern void HarfangInplaceMulIVec2(HarfangIVec2 this_, const HarfangIVec2 v);
extern void HarfangInplaceMulIVec2WithK(HarfangIVec2 this_, const int k);
extern void HarfangInplaceDivIVec2(HarfangIVec2 this_, const HarfangIVec2 v);
extern void HarfangInplaceDivIVec2WithK(HarfangIVec2 this_, const int k);
extern void HarfangSetIVec2(HarfangIVec2 this_, int x, int y);
extern HarfangIVec2 HarfangIVec2ListGetOperator(HarfangIVec2List h, int id);
extern void HarfangIVec2ListSetOperator(HarfangIVec2List h, int id, HarfangIVec2 v);
extern int HarfangIVec2ListLenOperator(HarfangIVec2List h);
extern HarfangIVec2List HarfangConstructorIVec2List();
extern HarfangIVec2List HarfangConstructorIVec2ListWithSequence(size_t sequenceToCSize, HarfangIVec2 *sequenceToCBuf);
extern void HarfangIVec2ListFree(HarfangIVec2List);
extern void HarfangClearIVec2List(HarfangIVec2List this_);
extern void HarfangReserveIVec2List(HarfangIVec2List this_, size_t size);
extern void HarfangPushBackIVec2List(HarfangIVec2List this_, HarfangIVec2 v);
extern size_t HarfangSizeIVec2List(HarfangIVec2List this_);
extern HarfangIVec2 HarfangAtIVec2List(HarfangIVec2List this_, size_t idx);
extern HarfangVec4 HarfangVec4ListGetOperator(HarfangVec4List h, int id);
extern void HarfangVec4ListSetOperator(HarfangVec4List h, int id, HarfangVec4 v);
extern int HarfangVec4ListLenOperator(HarfangVec4List h);
extern HarfangVec4List HarfangConstructorVec4List();
extern HarfangVec4List HarfangConstructorVec4ListWithSequence(size_t sequenceToCSize, HarfangVec4 *sequenceToCBuf);
extern void HarfangVec4ListFree(HarfangVec4List);
extern void HarfangClearVec4List(HarfangVec4List this_);
extern void HarfangReserveVec4List(HarfangVec4List this_, size_t size);
extern void HarfangPushBackVec4List(HarfangVec4List this_, HarfangVec4 v);
extern size_t HarfangSizeVec4List(HarfangVec4List this_);
extern HarfangVec4 HarfangAtVec4List(HarfangVec4List this_, size_t idx);
extern HarfangMat4 HarfangMat4ListGetOperator(HarfangMat4List h, int id);
extern void HarfangMat4ListSetOperator(HarfangMat4List h, int id, HarfangMat4 v);
extern int HarfangMat4ListLenOperator(HarfangMat4List h);
extern HarfangMat4List HarfangConstructorMat4List();
extern HarfangMat4List HarfangConstructorMat4ListWithSequence(size_t sequenceToCSize, HarfangMat4 *sequenceToCBuf);
extern void HarfangMat4ListFree(HarfangMat4List);
extern void HarfangClearMat4List(HarfangMat4List this_);
extern void HarfangReserveMat4List(HarfangMat4List this_, size_t size);
extern void HarfangPushBackMat4List(HarfangMat4List this_, HarfangMat4 v);
extern size_t HarfangSizeMat4List(HarfangMat4List this_);
extern HarfangMat4 HarfangAtMat4List(HarfangMat4List this_, size_t idx);
extern HarfangVec3 HarfangVec3ListGetOperator(HarfangVec3List h, int id);
extern void HarfangVec3ListSetOperator(HarfangVec3List h, int id, HarfangVec3 v);
extern int HarfangVec3ListLenOperator(HarfangVec3List h);
extern HarfangVec3List HarfangConstructorVec3List();
extern HarfangVec3List HarfangConstructorVec3ListWithSequence(size_t sequenceToCSize, HarfangVec3 *sequenceToCBuf);
extern void HarfangVec3ListFree(HarfangVec3List);
extern void HarfangClearVec3List(HarfangVec3List this_);
extern void HarfangReserveVec3List(HarfangVec3List this_, size_t size);
extern void HarfangPushBackVec3List(HarfangVec3List this_, HarfangVec3 v);
extern size_t HarfangSizeVec3List(HarfangVec3List this_);
extern HarfangVec3 HarfangAtVec3List(HarfangVec3List this_, size_t idx);
extern float HarfangRectGetSx(HarfangRect h);
extern void HarfangRectSetSx(HarfangRect h, float v);
extern float HarfangRectGetSy(HarfangRect h);
extern void HarfangRectSetSy(HarfangRect h, float v);
extern float HarfangRectGetEx(HarfangRect h);
extern void HarfangRectSetEx(HarfangRect h, float v);
extern float HarfangRectGetEy(HarfangRect h);
extern void HarfangRectSetEy(HarfangRect h, float v);
extern HarfangRect HarfangConstructorRect();
extern HarfangRect HarfangConstructorRectWithXY(float x, float y);
extern HarfangRect HarfangConstructorRectWithSxSyExEy(float sx, float sy, float ex, float ey);
extern HarfangRect HarfangConstructorRectWithRect(const HarfangRect rect);
extern void HarfangRectFree(HarfangRect);
extern int HarfangIntRectGetSx(HarfangIntRect h);
extern void HarfangIntRectSetSx(HarfangIntRect h, int v);
extern int HarfangIntRectGetSy(HarfangIntRect h);
extern void HarfangIntRectSetSy(HarfangIntRect h, int v);
extern int HarfangIntRectGetEx(HarfangIntRect h);
extern void HarfangIntRectSetEx(HarfangIntRect h, int v);
extern int HarfangIntRectGetEy(HarfangIntRect h);
extern void HarfangIntRectSetEy(HarfangIntRect h, int v);
extern HarfangIntRect HarfangConstructorIntRect();
extern HarfangIntRect HarfangConstructorIntRectWithXY(int x, int y);
extern HarfangIntRect HarfangConstructorIntRectWithSxSyExEy(int sx, int sy, int ex, int ey);
extern HarfangIntRect HarfangConstructorIntRectWithRect(const HarfangIntRect rect);
extern void HarfangIntRectFree(HarfangIntRect);
extern void HarfangFrustumFree(HarfangFrustum);
extern HarfangVec4 HarfangGetTopFrustum(HarfangFrustum this_);
extern void HarfangSetTopFrustum(HarfangFrustum this_, const HarfangVec4 plane);
extern HarfangVec4 HarfangGetBottomFrustum(HarfangFrustum this_);
extern void HarfangSetBottomFrustum(HarfangFrustum this_, const HarfangVec4 plane);
extern HarfangVec4 HarfangGetLeftFrustum(HarfangFrustum this_);
extern void HarfangSetLeftFrustum(HarfangFrustum this_, const HarfangVec4 plane);
extern HarfangVec4 HarfangGetRightFrustum(HarfangFrustum this_);
extern void HarfangSetRightFrustum(HarfangFrustum this_, const HarfangVec4 plane);
extern HarfangVec4 HarfangGetNearFrustum(HarfangFrustum this_);
extern void HarfangSetNearFrustum(HarfangFrustum this_, const HarfangVec4 plane);
extern HarfangVec4 HarfangGetFarFrustum(HarfangFrustum this_);
extern void HarfangSetFarFrustum(HarfangFrustum this_, const HarfangVec4 plane);
extern const char *HarfangMonitorModeGetName(HarfangMonitorMode h);
extern void HarfangMonitorModeSetName(HarfangMonitorMode h, const char *v);
extern HarfangIntRect HarfangMonitorModeGetRect(HarfangMonitorMode h);
extern void HarfangMonitorModeSetRect(HarfangMonitorMode h, HarfangIntRect v);
extern int HarfangMonitorModeGetFrequency(HarfangMonitorMode h);
extern void HarfangMonitorModeSetFrequency(HarfangMonitorMode h, int v);
extern uint8_t HarfangMonitorModeGetRotation(HarfangMonitorMode h);
extern void HarfangMonitorModeSetRotation(HarfangMonitorMode h, uint8_t v);
extern uint8_t HarfangMonitorModeGetSupportedRotations(HarfangMonitorMode h);
extern void HarfangMonitorModeSetSupportedRotations(HarfangMonitorMode h, uint8_t v);
extern void HarfangMonitorModeFree(HarfangMonitorMode);
extern HarfangMonitorMode HarfangMonitorModeListGetOperator(HarfangMonitorModeList h, int id);
extern void HarfangMonitorModeListSetOperator(HarfangMonitorModeList h, int id, HarfangMonitorMode v);
extern int HarfangMonitorModeListLenOperator(HarfangMonitorModeList h);
extern HarfangMonitorModeList HarfangConstructorMonitorModeList();
extern HarfangMonitorModeList HarfangConstructorMonitorModeListWithSequence(size_t sequenceToCSize, HarfangMonitorMode *sequenceToCBuf);
extern void HarfangMonitorModeListFree(HarfangMonitorModeList);
extern void HarfangClearMonitorModeList(HarfangMonitorModeList this_);
extern void HarfangReserveMonitorModeList(HarfangMonitorModeList this_, size_t size);
extern void HarfangPushBackMonitorModeList(HarfangMonitorModeList this_, HarfangMonitorMode v);
extern size_t HarfangSizeMonitorModeList(HarfangMonitorModeList this_);
extern HarfangMonitorMode HarfangAtMonitorModeList(HarfangMonitorModeList this_, size_t idx);
extern void HarfangMonitorFree(HarfangMonitor);
extern HarfangMonitor HarfangMonitorListGetOperator(HarfangMonitorList h, int id);
extern void HarfangMonitorListSetOperator(HarfangMonitorList h, int id, HarfangMonitor v);
extern int HarfangMonitorListLenOperator(HarfangMonitorList h);
extern HarfangMonitorList HarfangConstructorMonitorList();
extern HarfangMonitorList HarfangConstructorMonitorListWithSequence(size_t sequenceToCSize, HarfangMonitor *sequenceToCBuf);
extern void HarfangMonitorListFree(HarfangMonitorList);
extern void HarfangClearMonitorList(HarfangMonitorList this_);
extern void HarfangReserveMonitorList(HarfangMonitorList this_, size_t size);
extern void HarfangPushBackMonitorList(HarfangMonitorList this_, HarfangMonitor v);
extern size_t HarfangSizeMonitorList(HarfangMonitorList this_);
extern HarfangMonitor HarfangAtMonitorList(HarfangMonitorList this_, size_t idx);
extern void HarfangWindowFree(HarfangWindow);
extern HarfangColor HarfangColorGetZero();
extern HarfangColor HarfangColorGetOne();
extern HarfangColor HarfangColorGetWhite();
extern HarfangColor HarfangColorGetGrey();
extern HarfangColor HarfangColorGetBlack();
extern HarfangColor HarfangColorGetRed();
extern HarfangColor HarfangColorGetGreen();
extern HarfangColor HarfangColorGetBlue();
extern HarfangColor HarfangColorGetYellow();
extern HarfangColor HarfangColorGetOrange();
extern HarfangColor HarfangColorGetPurple();
extern HarfangColor HarfangColorGetTransparent();
extern float HarfangColorGetR(HarfangColor h);
extern void HarfangColorSetR(HarfangColor h, float v);
extern float HarfangColorGetG(HarfangColor h);
extern void HarfangColorSetG(HarfangColor h, float v);
extern float HarfangColorGetB(HarfangColor h);
extern void HarfangColorSetB(HarfangColor h, float v);
extern float HarfangColorGetA(HarfangColor h);
extern void HarfangColorSetA(HarfangColor h, float v);
extern HarfangColor HarfangConstructorColor();
extern HarfangColor HarfangConstructorColorWithColor(const HarfangColor color);
extern HarfangColor HarfangConstructorColorWithRGB(float r, float g, float b);
extern HarfangColor HarfangConstructorColorWithRGBA(float r, float g, float b, float a);
extern void HarfangColorFree(HarfangColor);
extern HarfangColor HarfangAddColor(HarfangColor this_, const HarfangColor color);
extern HarfangColor HarfangAddColorWithK(HarfangColor this_, float k);
extern HarfangColor HarfangSubColor(HarfangColor this_, const HarfangColor color);
extern HarfangColor HarfangSubColorWithK(HarfangColor this_, float k);
extern HarfangColor HarfangDivColor(HarfangColor this_, const HarfangColor color);
extern HarfangColor HarfangDivColorWithK(HarfangColor this_, float k);
extern HarfangColor HarfangMulColor(HarfangColor this_, const HarfangColor color);
extern HarfangColor HarfangMulColorWithK(HarfangColor this_, float k);
extern void HarfangInplaceAddColor(HarfangColor this_, HarfangColor color);
extern void HarfangInplaceAddColorWithK(HarfangColor this_, float k);
extern void HarfangInplaceSubColor(HarfangColor this_, HarfangColor color);
extern void HarfangInplaceSubColorWithK(HarfangColor this_, float k);
extern void HarfangInplaceMulColor(HarfangColor this_, HarfangColor color);
extern void HarfangInplaceMulColorWithK(HarfangColor this_, float k);
extern void HarfangInplaceDivColor(HarfangColor this_, HarfangColor color);
extern void HarfangInplaceDivColorWithK(HarfangColor this_, float k);
extern bool HarfangEqColor(HarfangColor this_, const HarfangColor color);
extern bool HarfangNeColor(HarfangColor this_, const HarfangColor color);
extern HarfangColor HarfangColorListGetOperator(HarfangColorList h, int id);
extern void HarfangColorListSetOperator(HarfangColorList h, int id, HarfangColor v);
extern int HarfangColorListLenOperator(HarfangColorList h);
extern HarfangColorList HarfangConstructorColorList();
extern HarfangColorList HarfangConstructorColorListWithSequence(size_t sequenceToCSize, HarfangColor *sequenceToCBuf);
extern void HarfangColorListFree(HarfangColorList);
extern void HarfangClearColorList(HarfangColorList this_);
extern void HarfangReserveColorList(HarfangColorList this_, size_t size);
extern void HarfangPushBackColorList(HarfangColorList this_, HarfangColor v);
extern size_t HarfangSizeColorList(HarfangColorList this_);
extern HarfangColor HarfangAtColorList(HarfangColorList this_, size_t idx);
extern HarfangPicture HarfangConstructorPicture();
extern HarfangPicture HarfangConstructorPictureWithPicture(const HarfangPicture picture);
extern HarfangPicture HarfangConstructorPictureWithWidthHeightFormat(uint16_t width, uint16_t height, int format);
extern HarfangPicture HarfangConstructorPictureWithDataWidthHeightFormat(HarfangVoidPointer data, uint16_t width, uint16_t height, int format);
extern void HarfangPictureFree(HarfangPicture);
extern uint32_t HarfangGetWidthPicture(HarfangPicture this_);
extern uint32_t HarfangGetHeightPicture(HarfangPicture this_);
extern int HarfangGetFormatPicture(HarfangPicture this_);
extern intptr_t HarfangGetDataPicture(HarfangPicture this_);
extern void HarfangSetDataPicture(HarfangPicture this_, HarfangVoidPointer data, uint16_t width, uint16_t height, int format);
extern void HarfangCopyDataPicture(HarfangPicture this_, const HarfangVoidPointer data, uint16_t width, uint16_t height, int format);
extern HarfangColor HarfangGetPixelRGBAPicture(HarfangPicture this_, uint16_t x, uint16_t y);
extern void HarfangSetPixelRGBAPicture(HarfangPicture this_, uint16_t x, uint16_t y, const HarfangColor col);
extern void HarfangFrameBufferHandleFree(HarfangFrameBufferHandle);
extern HarfangVertexLayout HarfangConstructorVertexLayout();
extern void HarfangVertexLayoutFree(HarfangVertexLayout);
extern HarfangVertexLayout HarfangBeginVertexLayout(HarfangVertexLayout this_);
extern HarfangVertexLayout HarfangAddVertexLayout(HarfangVertexLayout this_, int attrib, uint8_t count, int type);
extern HarfangVertexLayout HarfangAddVertexLayoutWithNormalized(HarfangVertexLayout this_, int attrib, uint8_t count, int type, bool normalized);
extern HarfangVertexLayout HarfangAddVertexLayoutWithNormalizedAsInt(
	HarfangVertexLayout this_, int attrib, uint8_t count, int type, bool normalized, bool as_int);
extern HarfangVertexLayout HarfangSkipVertexLayout(HarfangVertexLayout this_, uint8_t size);
extern void HarfangEndVertexLayout(HarfangVertexLayout this_);
extern bool HarfangHasVertexLayout(HarfangVertexLayout this_, int attrib);
extern uint16_t HarfangGetOffsetVertexLayout(HarfangVertexLayout this_, int attrib);
extern uint16_t HarfangGetStrideVertexLayout(HarfangVertexLayout this_);
extern uint32_t HarfangGetSizeVertexLayout(HarfangVertexLayout this_, uint32_t count);
extern void HarfangProgramHandleFree(HarfangProgramHandle);
extern int HarfangTextureInfoGetFormat(HarfangTextureInfo h);
extern void HarfangTextureInfoSetFormat(HarfangTextureInfo h, int v);
extern uint32_t HarfangTextureInfoGetStorageSize(HarfangTextureInfo h);
extern void HarfangTextureInfoSetStorageSize(HarfangTextureInfo h, uint32_t v);
extern uint16_t HarfangTextureInfoGetWidth(HarfangTextureInfo h);
extern void HarfangTextureInfoSetWidth(HarfangTextureInfo h, uint16_t v);
extern uint16_t HarfangTextureInfoGetHeight(HarfangTextureInfo h);
extern void HarfangTextureInfoSetHeight(HarfangTextureInfo h, uint16_t v);
extern uint16_t HarfangTextureInfoGetDepth(HarfangTextureInfo h);
extern void HarfangTextureInfoSetDepth(HarfangTextureInfo h, uint16_t v);
extern uint16_t HarfangTextureInfoGetNumLayers(HarfangTextureInfo h);
extern void HarfangTextureInfoSetNumLayers(HarfangTextureInfo h, uint16_t v);
extern uint8_t HarfangTextureInfoGetNumMips(HarfangTextureInfo h);
extern void HarfangTextureInfoSetNumMips(HarfangTextureInfo h, uint8_t v);
extern uint8_t HarfangTextureInfoGetBitsPerPixel(HarfangTextureInfo h);
extern void HarfangTextureInfoSetBitsPerPixel(HarfangTextureInfo h, uint8_t v);
extern bool HarfangTextureInfoGetCubeMap(HarfangTextureInfo h);
extern void HarfangTextureInfoSetCubeMap(HarfangTextureInfo h, bool v);
extern HarfangTextureInfo HarfangConstructorTextureInfo();
extern void HarfangTextureInfoFree(HarfangTextureInfo);
extern void HarfangModelRefFree(HarfangModelRef);
extern bool HarfangEqModelRef(HarfangModelRef this_, const HarfangModelRef m);
extern bool HarfangNeModelRef(HarfangModelRef this_, const HarfangModelRef m);
extern void HarfangTextureRefFree(HarfangTextureRef);
extern bool HarfangEqTextureRef(HarfangTextureRef this_, const HarfangTextureRef t);
extern bool HarfangNeTextureRef(HarfangTextureRef this_, const HarfangTextureRef t);
extern void HarfangMaterialRefFree(HarfangMaterialRef);
extern bool HarfangEqMaterialRef(HarfangMaterialRef this_, const HarfangMaterialRef m);
extern bool HarfangNeMaterialRef(HarfangMaterialRef this_, const HarfangMaterialRef m);
extern void HarfangPipelineProgramRefFree(HarfangPipelineProgramRef);
extern bool HarfangEqPipelineProgramRef(HarfangPipelineProgramRef this_, const HarfangPipelineProgramRef p);
extern bool HarfangNePipelineProgramRef(HarfangPipelineProgramRef this_, const HarfangPipelineProgramRef p);
extern HarfangTexture HarfangConstructorTexture();
extern void HarfangTextureFree(HarfangTexture);
extern void HarfangUniformSetValueFree(HarfangUniformSetValue);
extern HarfangUniformSetValue HarfangUniformSetValueListGetOperator(HarfangUniformSetValueList h, int id);
extern void HarfangUniformSetValueListSetOperator(HarfangUniformSetValueList h, int id, HarfangUniformSetValue v);
extern int HarfangUniformSetValueListLenOperator(HarfangUniformSetValueList h);
extern HarfangUniformSetValueList HarfangConstructorUniformSetValueList();
extern HarfangUniformSetValueList HarfangConstructorUniformSetValueListWithSequence(size_t sequenceToCSize, HarfangUniformSetValue *sequenceToCBuf);
extern void HarfangUniformSetValueListFree(HarfangUniformSetValueList);
extern void HarfangClearUniformSetValueList(HarfangUniformSetValueList this_);
extern void HarfangReserveUniformSetValueList(HarfangUniformSetValueList this_, size_t size);
extern void HarfangPushBackUniformSetValueList(HarfangUniformSetValueList this_, HarfangUniformSetValue v);
extern size_t HarfangSizeUniformSetValueList(HarfangUniformSetValueList this_);
extern HarfangUniformSetValue HarfangAtUniformSetValueList(HarfangUniformSetValueList this_, size_t idx);
extern void HarfangUniformSetTextureFree(HarfangUniformSetTexture);
extern HarfangUniformSetTexture HarfangUniformSetTextureListGetOperator(HarfangUniformSetTextureList h, int id);
extern void HarfangUniformSetTextureListSetOperator(HarfangUniformSetTextureList h, int id, HarfangUniformSetTexture v);
extern int HarfangUniformSetTextureListLenOperator(HarfangUniformSetTextureList h);
extern HarfangUniformSetTextureList HarfangConstructorUniformSetTextureList();
extern HarfangUniformSetTextureList HarfangConstructorUniformSetTextureListWithSequence(size_t sequenceToCSize, HarfangUniformSetTexture *sequenceToCBuf);
extern void HarfangUniformSetTextureListFree(HarfangUniformSetTextureList);
extern void HarfangClearUniformSetTextureList(HarfangUniformSetTextureList this_);
extern void HarfangReserveUniformSetTextureList(HarfangUniformSetTextureList this_, size_t size);
extern void HarfangPushBackUniformSetTextureList(HarfangUniformSetTextureList this_, HarfangUniformSetTexture v);
extern size_t HarfangSizeUniformSetTextureList(HarfangUniformSetTextureList this_);
extern HarfangUniformSetTexture HarfangAtUniformSetTextureList(HarfangUniformSetTextureList this_, size_t idx);
extern void HarfangPipelineProgramFree(HarfangPipelineProgram);
extern HarfangFrustum HarfangViewStateGetFrustum(HarfangViewState h);
extern void HarfangViewStateSetFrustum(HarfangViewState h, HarfangFrustum v);
extern HarfangMat44 HarfangViewStateGetProj(HarfangViewState h);
extern void HarfangViewStateSetProj(HarfangViewState h, HarfangMat44 v);
extern HarfangMat4 HarfangViewStateGetView(HarfangViewState h);
extern void HarfangViewStateSetView(HarfangViewState h, HarfangMat4 v);
extern HarfangViewState HarfangConstructorViewState();
extern void HarfangViewStateFree(HarfangViewState);
extern HarfangMaterial HarfangConstructorMaterial();
extern void HarfangMaterialFree(HarfangMaterial);
extern HarfangMaterial HarfangMaterialListGetOperator(HarfangMaterialList h, int id);
extern void HarfangMaterialListSetOperator(HarfangMaterialList h, int id, HarfangMaterial v);
extern int HarfangMaterialListLenOperator(HarfangMaterialList h);
extern HarfangMaterialList HarfangConstructorMaterialList();
extern HarfangMaterialList HarfangConstructorMaterialListWithSequence(size_t sequenceToCSize, HarfangMaterial *sequenceToCBuf);
extern void HarfangMaterialListFree(HarfangMaterialList);
extern void HarfangClearMaterialList(HarfangMaterialList this_);
extern void HarfangReserveMaterialList(HarfangMaterialList this_, size_t size);
extern void HarfangPushBackMaterialList(HarfangMaterialList this_, HarfangMaterial v);
extern size_t HarfangSizeMaterialList(HarfangMaterialList this_);
extern HarfangMaterial HarfangAtMaterialList(HarfangMaterialList this_, size_t idx);
extern void HarfangRenderStateFree(HarfangRenderState);
extern void HarfangModelFree(HarfangModel);
extern HarfangPipelineResources HarfangConstructorPipelineResources();
extern void HarfangPipelineResourcesFree(HarfangPipelineResources);
extern HarfangTextureRef HarfangAddTexturePipelineResources(HarfangPipelineResources this_, const char *name, const HarfangTexture tex);
extern HarfangModelRef HarfangAddModelPipelineResources(HarfangPipelineResources this_, const char *name, const HarfangModel mdl);
extern HarfangPipelineProgramRef HarfangAddProgramPipelineResources(HarfangPipelineResources this_, const char *name, const HarfangPipelineProgram prg);
extern HarfangTextureRef HarfangHasTexturePipelineResources(HarfangPipelineResources this_, const char *name);
extern HarfangModelRef HarfangHasModelPipelineResources(HarfangPipelineResources this_, const char *name);
extern HarfangPipelineProgramRef HarfangHasProgramPipelineResources(HarfangPipelineResources this_, const char *name);
extern void HarfangUpdateTexturePipelineResources(HarfangPipelineResources this_, HarfangTextureRef ref, const HarfangTexture tex);
extern void HarfangUpdateModelPipelineResources(HarfangPipelineResources this_, HarfangModelRef ref, const HarfangModel mdl);
extern void HarfangUpdateProgramPipelineResources(HarfangPipelineResources this_, HarfangPipelineProgramRef ref, const HarfangPipelineProgram prg);
extern HarfangTexture HarfangGetTexturePipelineResources(HarfangPipelineResources this_, HarfangTextureRef ref);
extern HarfangModel HarfangGetModelPipelineResources(HarfangPipelineResources this_, HarfangModelRef ref);
extern HarfangPipelineProgram HarfangGetProgramPipelineResources(HarfangPipelineResources this_, HarfangPipelineProgramRef ref);
extern const char *HarfangGetTextureNamePipelineResources(HarfangPipelineResources this_, HarfangTextureRef ref);
extern const char *HarfangGetModelNamePipelineResources(HarfangPipelineResources this_, HarfangModelRef ref);
extern const char *HarfangGetProgramNamePipelineResources(HarfangPipelineResources this_, HarfangPipelineProgramRef ref);
extern void HarfangDestroyAllTexturesPipelineResources(HarfangPipelineResources this_);
extern void HarfangDestroyAllModelsPipelineResources(HarfangPipelineResources this_);
extern void HarfangDestroyAllProgramsPipelineResources(HarfangPipelineResources this_);
extern void HarfangDestroyTexturePipelineResources(HarfangPipelineResources this_, HarfangTextureRef ref);
extern void HarfangDestroyModelPipelineResources(HarfangPipelineResources this_, HarfangModelRef ref);
extern void HarfangDestroyProgramPipelineResources(HarfangPipelineResources this_, HarfangPipelineProgramRef ref);
extern bool HarfangHasTextureInfoPipelineResources(HarfangPipelineResources this_, HarfangTextureRef ref);
extern HarfangTextureInfo HarfangGetTextureInfoPipelineResources(HarfangPipelineResources this_, HarfangTextureRef ref);
extern HarfangFrameBufferHandle HarfangFrameBufferGetHandle(HarfangFrameBuffer h);
extern void HarfangFrameBufferSetHandle(HarfangFrameBuffer h, HarfangFrameBufferHandle v);
extern void HarfangFrameBufferFree(HarfangFrameBuffer);
extern HarfangVertices HarfangConstructorVertices(const HarfangVertexLayout decl, size_t count);
extern void HarfangVerticesFree(HarfangVertices);
extern const HarfangVertexLayout HarfangGetDeclVertices(HarfangVertices this_);
extern HarfangVertices HarfangBeginVertices(HarfangVertices this_, size_t vertex_index);
extern HarfangVertices HarfangSetPosVertices(HarfangVertices this_, const HarfangVec3 pos);
extern HarfangVertices HarfangSetNormalVertices(HarfangVertices this_, const HarfangVec3 normal);
extern HarfangVertices HarfangSetTangentVertices(HarfangVertices this_, const HarfangVec3 tangent);
extern HarfangVertices HarfangSetBinormalVertices(HarfangVertices this_, const HarfangVec3 binormal);
extern HarfangVertices HarfangSetTexCoord0Vertices(HarfangVertices this_, const HarfangVec2 uv);
extern HarfangVertices HarfangSetTexCoord1Vertices(HarfangVertices this_, const HarfangVec2 uv);
extern HarfangVertices HarfangSetTexCoord2Vertices(HarfangVertices this_, const HarfangVec2 uv);
extern HarfangVertices HarfangSetTexCoord3Vertices(HarfangVertices this_, const HarfangVec2 uv);
extern HarfangVertices HarfangSetTexCoord4Vertices(HarfangVertices this_, const HarfangVec2 uv);
extern HarfangVertices HarfangSetTexCoord5Vertices(HarfangVertices this_, const HarfangVec2 uv);
extern HarfangVertices HarfangSetTexCoord6Vertices(HarfangVertices this_, const HarfangVec2 uv);
extern HarfangVertices HarfangSetTexCoord7Vertices(HarfangVertices this_, const HarfangVec2 uv);
extern HarfangVertices HarfangSetColor0Vertices(HarfangVertices this_, const HarfangColor color);
extern HarfangVertices HarfangSetColor1Vertices(HarfangVertices this_, const HarfangColor color);
extern HarfangVertices HarfangSetColor2Vertices(HarfangVertices this_, const HarfangColor color);
extern HarfangVertices HarfangSetColor3Vertices(HarfangVertices this_, const HarfangColor color);
extern void HarfangEndVertices(HarfangVertices this_);
extern void HarfangEndVerticesWithValidate(HarfangVertices this_, bool validate);
extern void HarfangClearVertices(HarfangVertices this_);
extern void HarfangReserveVertices(HarfangVertices this_, size_t count);
extern void HarfangResizeVertices(HarfangVertices this_, size_t count);
extern const HarfangVoidPointer HarfangGetDataVertices(HarfangVertices this_);
extern size_t HarfangGetSizeVertices(HarfangVertices this_);
extern size_t HarfangGetCountVertices(HarfangVertices this_);
extern size_t HarfangGetCapacityVertices(HarfangVertices this_);
extern void HarfangPipelineFree(HarfangPipeline);
extern const char *HarfangPipelineInfoGetName(HarfangPipelineInfo h);
extern void HarfangPipelineInfoSetName(HarfangPipelineInfo h, const char *v);
extern void HarfangPipelineInfoFree(HarfangPipelineInfo);
extern void HarfangForwardPipelineFree(HarfangForwardPipeline);
extern int HarfangForwardPipelineLightGetType(HarfangForwardPipelineLight h);
extern void HarfangForwardPipelineLightSetType(HarfangForwardPipelineLight h, int v);
extern HarfangMat4 HarfangForwardPipelineLightGetWorld(HarfangForwardPipelineLight h);
extern void HarfangForwardPipelineLightSetWorld(HarfangForwardPipelineLight h, HarfangMat4 v);
extern HarfangColor HarfangForwardPipelineLightGetDiffuse(HarfangForwardPipelineLight h);
extern void HarfangForwardPipelineLightSetDiffuse(HarfangForwardPipelineLight h, HarfangColor v);
extern HarfangColor HarfangForwardPipelineLightGetSpecular(HarfangForwardPipelineLight h);
extern void HarfangForwardPipelineLightSetSpecular(HarfangForwardPipelineLight h, HarfangColor v);
extern float HarfangForwardPipelineLightGetRadius(HarfangForwardPipelineLight h);
extern void HarfangForwardPipelineLightSetRadius(HarfangForwardPipelineLight h, float v);
extern float HarfangForwardPipelineLightGetInnerAngle(HarfangForwardPipelineLight h);
extern void HarfangForwardPipelineLightSetInnerAngle(HarfangForwardPipelineLight h, float v);
extern float HarfangForwardPipelineLightGetOuterAngle(HarfangForwardPipelineLight h);
extern void HarfangForwardPipelineLightSetOuterAngle(HarfangForwardPipelineLight h, float v);
extern HarfangVec4 HarfangForwardPipelineLightGetPssmSplit(HarfangForwardPipelineLight h);
extern void HarfangForwardPipelineLightSetPssmSplit(HarfangForwardPipelineLight h, HarfangVec4 v);
extern float HarfangForwardPipelineLightGetPriority(HarfangForwardPipelineLight h);
extern void HarfangForwardPipelineLightSetPriority(HarfangForwardPipelineLight h, float v);
extern HarfangForwardPipelineLight HarfangConstructorForwardPipelineLight();
extern void HarfangForwardPipelineLightFree(HarfangForwardPipelineLight);
extern HarfangForwardPipelineLight HarfangForwardPipelineLightListGetOperator(HarfangForwardPipelineLightList h, int id);
extern void HarfangForwardPipelineLightListSetOperator(HarfangForwardPipelineLightList h, int id, HarfangForwardPipelineLight v);
extern int HarfangForwardPipelineLightListLenOperator(HarfangForwardPipelineLightList h);
extern HarfangForwardPipelineLightList HarfangConstructorForwardPipelineLightList();
extern HarfangForwardPipelineLightList HarfangConstructorForwardPipelineLightListWithSequence(
	size_t sequenceToCSize, HarfangForwardPipelineLight *sequenceToCBuf);
extern void HarfangForwardPipelineLightListFree(HarfangForwardPipelineLightList);
extern void HarfangClearForwardPipelineLightList(HarfangForwardPipelineLightList this_);
extern void HarfangReserveForwardPipelineLightList(HarfangForwardPipelineLightList this_, size_t size);
extern void HarfangPushBackForwardPipelineLightList(HarfangForwardPipelineLightList this_, HarfangForwardPipelineLight v);
extern size_t HarfangSizeForwardPipelineLightList(HarfangForwardPipelineLightList this_);
extern HarfangForwardPipelineLight HarfangAtForwardPipelineLightList(HarfangForwardPipelineLightList this_, size_t idx);
extern void HarfangForwardPipelineLightsFree(HarfangForwardPipelineLights);
extern float HarfangForwardPipelineFogGetNear(HarfangForwardPipelineFog h);
extern void HarfangForwardPipelineFogSetNear(HarfangForwardPipelineFog h, float v);
extern float HarfangForwardPipelineFogGetFar(HarfangForwardPipelineFog h);
extern void HarfangForwardPipelineFogSetFar(HarfangForwardPipelineFog h, float v);
extern HarfangColor HarfangForwardPipelineFogGetColor(HarfangForwardPipelineFog h);
extern void HarfangForwardPipelineFogSetColor(HarfangForwardPipelineFog h, HarfangColor v);
extern HarfangForwardPipelineFog HarfangConstructorForwardPipelineFog();
extern void HarfangForwardPipelineFogFree(HarfangForwardPipelineFog);
extern void HarfangFontFree(HarfangFont);
extern HarfangJSON HarfangConstructorJSON();
extern void HarfangJSONFree(HarfangJSON);
extern void HarfangLuaObjectFree(HarfangLuaObject);
extern HarfangLuaObject HarfangLuaObjectListGetOperator(HarfangLuaObjectList h, int id);
extern void HarfangLuaObjectListSetOperator(HarfangLuaObjectList h, int id, HarfangLuaObject v);
extern int HarfangLuaObjectListLenOperator(HarfangLuaObjectList h);
extern HarfangLuaObjectList HarfangConstructorLuaObjectList();
extern HarfangLuaObjectList HarfangConstructorLuaObjectListWithSequence(size_t sequenceToCSize, HarfangLuaObject *sequenceToCBuf);
extern void HarfangLuaObjectListFree(HarfangLuaObjectList);
extern void HarfangClearLuaObjectList(HarfangLuaObjectList this_);
extern void HarfangReserveLuaObjectList(HarfangLuaObjectList this_, size_t size);
extern void HarfangPushBackLuaObjectList(HarfangLuaObjectList this_, HarfangLuaObject v);
extern size_t HarfangSizeLuaObjectList(HarfangLuaObjectList this_);
extern HarfangLuaObject HarfangAtLuaObjectList(HarfangLuaObjectList this_, size_t idx);
extern void HarfangSceneAnimRefFree(HarfangSceneAnimRef);
extern bool HarfangEqSceneAnimRef(HarfangSceneAnimRef this_, const HarfangSceneAnimRef ref);
extern bool HarfangNeSceneAnimRef(HarfangSceneAnimRef this_, const HarfangSceneAnimRef ref);
extern HarfangSceneAnimRef HarfangSceneAnimRefListGetOperator(HarfangSceneAnimRefList h, int id);
extern void HarfangSceneAnimRefListSetOperator(HarfangSceneAnimRefList h, int id, HarfangSceneAnimRef v);
extern int HarfangSceneAnimRefListLenOperator(HarfangSceneAnimRefList h);
extern HarfangSceneAnimRefList HarfangConstructorSceneAnimRefList();
extern HarfangSceneAnimRefList HarfangConstructorSceneAnimRefListWithSequence(size_t sequenceToCSize, HarfangSceneAnimRef *sequenceToCBuf);
extern void HarfangSceneAnimRefListFree(HarfangSceneAnimRefList);
extern void HarfangClearSceneAnimRefList(HarfangSceneAnimRefList this_);
extern void HarfangReserveSceneAnimRefList(HarfangSceneAnimRefList this_, size_t size);
extern void HarfangPushBackSceneAnimRefList(HarfangSceneAnimRefList this_, HarfangSceneAnimRef v);
extern size_t HarfangSizeSceneAnimRefList(HarfangSceneAnimRefList this_);
extern HarfangSceneAnimRef HarfangAtSceneAnimRefList(HarfangSceneAnimRefList this_, size_t idx);
extern void HarfangScenePlayAnimRefFree(HarfangScenePlayAnimRef);
extern bool HarfangEqScenePlayAnimRef(HarfangScenePlayAnimRef this_, const HarfangScenePlayAnimRef ref);
extern bool HarfangNeScenePlayAnimRef(HarfangScenePlayAnimRef this_, const HarfangScenePlayAnimRef ref);
extern HarfangScenePlayAnimRef HarfangScenePlayAnimRefListGetOperator(HarfangScenePlayAnimRefList h, int id);
extern void HarfangScenePlayAnimRefListSetOperator(HarfangScenePlayAnimRefList h, int id, HarfangScenePlayAnimRef v);
extern int HarfangScenePlayAnimRefListLenOperator(HarfangScenePlayAnimRefList h);
extern HarfangScenePlayAnimRefList HarfangConstructorScenePlayAnimRefList();
extern HarfangScenePlayAnimRefList HarfangConstructorScenePlayAnimRefListWithSequence(size_t sequenceToCSize, HarfangScenePlayAnimRef *sequenceToCBuf);
extern void HarfangScenePlayAnimRefListFree(HarfangScenePlayAnimRefList);
extern void HarfangClearScenePlayAnimRefList(HarfangScenePlayAnimRefList this_);
extern void HarfangReserveScenePlayAnimRefList(HarfangScenePlayAnimRefList this_, size_t size);
extern void HarfangPushBackScenePlayAnimRefList(HarfangScenePlayAnimRefList this_, HarfangScenePlayAnimRef v);
extern size_t HarfangSizeScenePlayAnimRefList(HarfangScenePlayAnimRefList this_);
extern HarfangScenePlayAnimRef HarfangAtScenePlayAnimRefList(HarfangScenePlayAnimRefList this_, size_t idx);
extern HarfangCanvas HarfangSceneGetCanvas(HarfangScene h);
extern void HarfangSceneSetCanvas(HarfangScene h, HarfangCanvas v);
extern HarfangEnvironment HarfangSceneGetEnvironment(HarfangScene h);
extern void HarfangSceneSetEnvironment(HarfangScene h, HarfangEnvironment v);
extern HarfangScene HarfangConstructorScene();
extern void HarfangSceneFree(HarfangScene);
extern HarfangNode HarfangGetNodeScene(HarfangScene this_, const char *name);
extern HarfangNode HarfangGetNodeExScene(HarfangScene this_, const char *path);
extern HarfangNodeList HarfangGetNodesScene(HarfangScene this_);
extern HarfangNodeList HarfangGetAllNodesScene(HarfangScene this_);
extern HarfangNodeList HarfangGetNodesWithComponentScene(HarfangScene this_, int idx);
extern HarfangNodeList HarfangGetAllNodesWithComponentScene(HarfangScene this_, int idx);
extern size_t HarfangGetNodeCountScene(HarfangScene this_);
extern size_t HarfangGetAllNodeCountScene(HarfangScene this_);
extern HarfangNodeList HarfangGetNodeChildrenScene(HarfangScene this_, const HarfangNode node);
extern bool HarfangIsChildOfScene(HarfangScene this_, const HarfangNode node, const HarfangNode parent);
extern bool HarfangIsRootScene(HarfangScene this_, const HarfangNode node);
extern void HarfangReadyWorldMatricesScene(HarfangScene this_);
extern void HarfangComputeWorldMatricesScene(HarfangScene this_);
extern void HarfangUpdateScene(HarfangScene this_, int64_t dt);
extern HarfangSceneAnimRefList HarfangGetSceneAnimsScene(HarfangScene this_);
extern HarfangSceneAnimRef HarfangGetSceneAnimScene(HarfangScene this_, const char *name);
extern HarfangScenePlayAnimRef HarfangPlayAnimScene(HarfangScene this_, HarfangSceneAnimRef ref);
extern HarfangScenePlayAnimRef HarfangPlayAnimSceneWithLoopMode(HarfangScene this_, HarfangSceneAnimRef ref, int loop_mode);
extern HarfangScenePlayAnimRef HarfangPlayAnimSceneWithLoopModeEasing(HarfangScene this_, HarfangSceneAnimRef ref, int loop_mode, unsigned char easing);
extern HarfangScenePlayAnimRef HarfangPlayAnimSceneWithLoopModeEasingTStart(
	HarfangScene this_, HarfangSceneAnimRef ref, int loop_mode, unsigned char easing, int64_t t_start);
extern HarfangScenePlayAnimRef HarfangPlayAnimSceneWithLoopModeEasingTStartTEnd(
	HarfangScene this_, HarfangSceneAnimRef ref, int loop_mode, unsigned char easing, int64_t t_start, int64_t t_end);
extern HarfangScenePlayAnimRef HarfangPlayAnimSceneWithLoopModeEasingTStartTEndPaused(
	HarfangScene this_, HarfangSceneAnimRef ref, int loop_mode, unsigned char easing, int64_t t_start, int64_t t_end, bool paused);
extern HarfangScenePlayAnimRef HarfangPlayAnimSceneWithLoopModeEasingTStartTEndPausedTScale(
	HarfangScene this_, HarfangSceneAnimRef ref, int loop_mode, unsigned char easing, int64_t t_start, int64_t t_end, bool paused, float t_scale);
extern bool HarfangIsPlayingScene(HarfangScene this_, HarfangScenePlayAnimRef ref);
extern void HarfangStopAnimScene(HarfangScene this_, HarfangScenePlayAnimRef ref);
extern void HarfangStopAllAnimsScene(HarfangScene this_);
extern HarfangStringList HarfangGetPlayingAnimNamesScene(HarfangScene this_);
extern HarfangScenePlayAnimRefList HarfangGetPlayingAnimRefsScene(HarfangScene this_);
extern void HarfangUpdatePlayingAnimsScene(HarfangScene this_, int64_t dt);
extern bool HarfangHasKeyScene(HarfangScene this_, const char *key);
extern HarfangStringList HarfangGetKeysScene(HarfangScene this_);
extern void HarfangRemoveKeyScene(HarfangScene this_, const char *key);
extern const char *HarfangGetValueScene(HarfangScene this_, const char *key);
extern void HarfangSetValueScene(HarfangScene this_, const char *key, const char *value);
extern size_t HarfangGarbageCollectScene(HarfangScene this_);
extern void HarfangClearScene(HarfangScene this_);
extern void HarfangReserveNodesScene(HarfangScene this_, size_t count);
extern HarfangNode HarfangCreateNodeScene(HarfangScene this_);
extern HarfangNode HarfangCreateNodeSceneWithName(HarfangScene this_, const char *name);
extern void HarfangDestroyNodeScene(HarfangScene this_, const HarfangNode node);
extern void HarfangReserveTransformsScene(HarfangScene this_, size_t count);
extern HarfangTransform HarfangCreateTransformScene(HarfangScene this_);
extern HarfangTransform HarfangCreateTransformSceneWithT(HarfangScene this_, const HarfangVec3 T);
extern HarfangTransform HarfangCreateTransformSceneWithTR(HarfangScene this_, const HarfangVec3 T, const HarfangVec3 R);
extern HarfangTransform HarfangCreateTransformSceneWithTRS(HarfangScene this_, const HarfangVec3 T, const HarfangVec3 R, const HarfangVec3 S);
extern void HarfangDestroyTransformScene(HarfangScene this_, const HarfangTransform transform);
extern void HarfangReserveCamerasScene(HarfangScene this_, size_t count);
extern HarfangCamera HarfangCreateCameraScene(HarfangScene this_);
extern HarfangCamera HarfangCreateCameraSceneWithZnearZfar(HarfangScene this_, float znear, float zfar);
extern HarfangCamera HarfangCreateCameraSceneWithZnearZfarFov(HarfangScene this_, float znear, float zfar, float fov);
extern HarfangCamera HarfangCreateOrthographicCameraScene(HarfangScene this_, float znear, float zfar);
extern HarfangCamera HarfangCreateOrthographicCameraSceneWithSize(HarfangScene this_, float znear, float zfar, float size);
extern void HarfangDestroyCameraScene(HarfangScene this_, const HarfangCamera camera);
extern HarfangViewState HarfangComputeCurrentCameraViewStateScene(HarfangScene this_, const HarfangVec2 aspect_ratio);
extern void HarfangReserveObjectsScene(HarfangScene this_, size_t count);
extern HarfangObject HarfangCreateObjectScene(HarfangScene this_);
extern HarfangObject HarfangCreateObjectSceneWithModelMaterials(HarfangScene this_, const HarfangModelRef model, const HarfangMaterialList materials);
extern HarfangObject HarfangCreateObjectSceneWithModelSliceOfMaterials(
	HarfangScene this_, const HarfangModelRef model, size_t SliceOfmaterialsToCSize, HarfangMaterial *SliceOfmaterialsToCBuf);
extern void HarfangDestroyObjectScene(HarfangScene this_, const HarfangObject object);
extern void HarfangReserveLightsScene(HarfangScene this_, size_t count);
extern HarfangLight HarfangCreateLightScene(HarfangScene this_);
extern void HarfangDestroyLightScene(HarfangScene this_, const HarfangLight light);
extern HarfangLight HarfangCreateLinearLightSceneWithDiffuseIntensitySpecularSpecularIntensity(
	HarfangScene this_, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity);
extern HarfangLight HarfangCreateLinearLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriority(
	HarfangScene this_, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority);
extern HarfangLight HarfangCreateLinearLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(HarfangScene this_, const HarfangColor diffuse,
	float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority, int shadow_type);
extern HarfangLight HarfangCreateLinearLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(HarfangScene this_,
	const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority, int shadow_type,
	float shadow_bias);
extern HarfangLight HarfangCreateLinearLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBiasPssmSplit(HarfangScene this_,
	const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority, int shadow_type,
	float shadow_bias, const HarfangVec4 pssm_split);
extern HarfangLight HarfangCreateLinearLightScene(HarfangScene this_, const HarfangColor diffuse, const HarfangColor specular);
extern HarfangLight HarfangCreateLinearLightSceneWithPriority(HarfangScene this_, const HarfangColor diffuse, const HarfangColor specular, float priority);
extern HarfangLight HarfangCreateLinearLightSceneWithPriorityShadowType(
	HarfangScene this_, const HarfangColor diffuse, const HarfangColor specular, float priority, int shadow_type);
extern HarfangLight HarfangCreateLinearLightSceneWithPriorityShadowTypeShadowBias(
	HarfangScene this_, const HarfangColor diffuse, const HarfangColor specular, float priority, int shadow_type, float shadow_bias);
extern HarfangLight HarfangCreateLinearLightSceneWithPriorityShadowTypeShadowBiasPssmSplit(HarfangScene this_, const HarfangColor diffuse,
	const HarfangColor specular, float priority, int shadow_type, float shadow_bias, const HarfangVec4 pssm_split);
extern HarfangLight HarfangCreatePointLightSceneWithDiffuseIntensitySpecularSpecularIntensity(
	HarfangScene this_, float radius, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity);
extern HarfangLight HarfangCreatePointLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriority(HarfangScene this_, float radius,
	const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority);
extern HarfangLight HarfangCreatePointLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(HarfangScene this_, float radius,
	const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority, int shadow_type);
extern HarfangLight HarfangCreatePointLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(HarfangScene this_, float radius,
	const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority, int shadow_type,
	float shadow_bias);
extern HarfangLight HarfangCreatePointLightScene(HarfangScene this_, float radius, const HarfangColor diffuse, const HarfangColor specular);
extern HarfangLight HarfangCreatePointLightSceneWithPriority(
	HarfangScene this_, float radius, const HarfangColor diffuse, const HarfangColor specular, float priority);
extern HarfangLight HarfangCreatePointLightSceneWithPriorityShadowType(
	HarfangScene this_, float radius, const HarfangColor diffuse, const HarfangColor specular, float priority, int shadow_type);
extern HarfangLight HarfangCreatePointLightSceneWithPriorityShadowTypeShadowBias(
	HarfangScene this_, float radius, const HarfangColor diffuse, const HarfangColor specular, float priority, int shadow_type, float shadow_bias);
extern HarfangLight HarfangCreateSpotLightSceneWithDiffuseIntensitySpecularSpecularIntensity(HarfangScene this_, float radius, float inner_angle,
	float outer_angle, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity);
extern HarfangLight HarfangCreateSpotLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriority(HarfangScene this_, float radius, float inner_angle,
	float outer_angle, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority);
extern HarfangLight HarfangCreateSpotLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(HarfangScene this_, float radius,
	float inner_angle, float outer_angle, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity,
	float priority, int shadow_type);
extern HarfangLight HarfangCreateSpotLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(HarfangScene this_, float radius,
	float inner_angle, float outer_angle, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity,
	float priority, int shadow_type, float shadow_bias);
extern HarfangLight HarfangCreateSpotLightScene(
	HarfangScene this_, float radius, float inner_angle, float outer_angle, const HarfangColor diffuse, const HarfangColor specular);
extern HarfangLight HarfangCreateSpotLightSceneWithPriority(
	HarfangScene this_, float radius, float inner_angle, float outer_angle, const HarfangColor diffuse, const HarfangColor specular, float priority);
extern HarfangLight HarfangCreateSpotLightSceneWithPriorityShadowType(HarfangScene this_, float radius, float inner_angle, float outer_angle,
	const HarfangColor diffuse, const HarfangColor specular, float priority, int shadow_type);
extern HarfangLight HarfangCreateSpotLightSceneWithPriorityShadowTypeShadowBias(HarfangScene this_, float radius, float inner_angle, float outer_angle,
	const HarfangColor diffuse, const HarfangColor specular, float priority, int shadow_type, float shadow_bias);
extern void HarfangReserveScriptsScene(HarfangScene this_, size_t count);
extern HarfangScript HarfangCreateScriptScene(HarfangScene this_);
extern HarfangScript HarfangCreateScriptSceneWithPath(HarfangScene this_, const char *path);
extern void HarfangDestroyScriptScene(HarfangScene this_, const HarfangScript script);
extern size_t HarfangGetScriptCountScene(HarfangScene this_);
extern void HarfangSetScriptScene(HarfangScene this_, size_t slot_idx, const HarfangScript script);
extern HarfangScript HarfangGetScriptScene(HarfangScene this_, size_t slot_idx);
extern HarfangRigidBody HarfangCreateRigidBodyScene(HarfangScene this_);
extern void HarfangDestroyRigidBodyScene(HarfangScene this_, const HarfangRigidBody rigid_body);
extern HarfangCollision HarfangCreateCollisionScene(HarfangScene this_);
extern void HarfangDestroyCollisionScene(HarfangScene this_, const HarfangCollision collision);
extern HarfangInstance HarfangCreateInstanceScene(HarfangScene this_);
extern void HarfangDestroyInstanceScene(HarfangScene this_, const HarfangInstance Instance);
extern void HarfangSetProbeScene(HarfangScene this_, HarfangTextureRef irradiance, HarfangTextureRef radiance, HarfangTextureRef brdf);
extern HarfangNode HarfangGetCurrentCameraScene(HarfangScene this_);
extern void HarfangSetCurrentCameraScene(HarfangScene this_, const HarfangNode camera);
extern bool HarfangGetMinMaxScene(HarfangScene this_, const HarfangPipelineResources resources, HarfangMinMax minmax);
extern void HarfangSceneViewFree(HarfangSceneView);
extern HarfangNodeList HarfangGetNodesSceneView(HarfangSceneView this_, const HarfangScene scene);
extern HarfangNode HarfangGetNodeSceneView(HarfangSceneView this_, const HarfangScene scene, const char *name);
extern void HarfangNodeFree(HarfangNode);
extern bool HarfangEqNode(HarfangNode this_, const HarfangNode n);
extern bool HarfangIsValidNode(HarfangNode this_);
extern uint32_t HarfangGetUidNode(HarfangNode this_);
extern uint32_t HarfangGetFlagsNode(HarfangNode this_);
extern void HarfangSetFlagsNode(HarfangNode this_, uint32_t flags);
extern void HarfangEnableNode(HarfangNode this_);
extern void HarfangDisableNode(HarfangNode this_);
extern bool HarfangIsEnabledNode(HarfangNode this_);
extern bool HarfangIsItselfEnabledNode(HarfangNode this_);
extern bool HarfangHasTransformNode(HarfangNode this_);
extern HarfangTransform HarfangGetTransformNode(HarfangNode this_);
extern void HarfangSetTransformNode(HarfangNode this_, const HarfangTransform t);
extern void HarfangRemoveTransformNode(HarfangNode this_);
extern bool HarfangHasCameraNode(HarfangNode this_);
extern HarfangCamera HarfangGetCameraNode(HarfangNode this_);
extern void HarfangSetCameraNode(HarfangNode this_, const HarfangCamera c);
extern void HarfangRemoveCameraNode(HarfangNode this_);
extern HarfangViewState HarfangComputeCameraViewStateNode(HarfangNode this_, const HarfangVec2 aspect_ratio);
extern bool HarfangHasObjectNode(HarfangNode this_);
extern HarfangObject HarfangGetObjectNode(HarfangNode this_);
extern void HarfangSetObjectNode(HarfangNode this_, const HarfangObject o);
extern void HarfangRemoveObjectNode(HarfangNode this_);
extern bool HarfangGetMinMaxNode(HarfangNode this_, const HarfangPipelineResources resources, HarfangMinMax minmax);
extern bool HarfangHasLightNode(HarfangNode this_);
extern HarfangLight HarfangGetLightNode(HarfangNode this_);
extern void HarfangSetLightNode(HarfangNode this_, const HarfangLight l);
extern void HarfangRemoveLightNode(HarfangNode this_);
extern bool HarfangHasRigidBodyNode(HarfangNode this_);
extern HarfangRigidBody HarfangGetRigidBodyNode(HarfangNode this_);
extern void HarfangSetRigidBodyNode(HarfangNode this_, const HarfangRigidBody b);
extern void HarfangRemoveRigidBodyNode(HarfangNode this_);
extern size_t HarfangGetCollisionCountNode(HarfangNode this_);
extern HarfangCollision HarfangGetCollisionNode(HarfangNode this_, size_t slot);
extern void HarfangSetCollisionNode(HarfangNode this_, size_t slot, const HarfangCollision c);
extern void HarfangRemoveCollisionNode(HarfangNode this_, const HarfangCollision c);
extern void HarfangRemoveCollisionNodeWithSlot(HarfangNode this_, size_t slot);
extern const char *HarfangGetNameNode(HarfangNode this_);
extern void HarfangSetNameNode(HarfangNode this_, const char *name);
extern size_t HarfangGetScriptCountNode(HarfangNode this_);
extern HarfangScript HarfangGetScriptNode(HarfangNode this_, size_t idx);
extern void HarfangSetScriptNode(HarfangNode this_, size_t idx, const HarfangScript s);
extern void HarfangRemoveScriptNode(HarfangNode this_, const HarfangScript s);
extern void HarfangRemoveScriptNodeWithSlot(HarfangNode this_, size_t slot);
extern bool HarfangHasInstanceNode(HarfangNode this_);
extern HarfangInstance HarfangGetInstanceNode(HarfangNode this_);
extern void HarfangSetInstanceNode(HarfangNode this_, const HarfangInstance instance);
extern bool HarfangSetupInstanceFromFileNode(HarfangNode this_, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern bool HarfangSetupInstanceFromFileNodeWithFlags(
	HarfangNode this_, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline, uint32_t flags);
extern bool HarfangSetupInstanceFromAssetsNode(HarfangNode this_, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern bool HarfangSetupInstanceFromAssetsNodeWithFlags(
	HarfangNode this_, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline, uint32_t flags);
extern void HarfangDestroyInstanceNode(HarfangNode this_);
extern HarfangNode HarfangIsInstantiatedByNode(HarfangNode this_);
extern HarfangSceneView HarfangGetInstanceSceneViewNode(HarfangNode this_);
extern HarfangSceneAnimRef HarfangGetInstanceSceneAnimNode(HarfangNode this_, const char *path);
extern void HarfangStartOnInstantiateAnimNode(HarfangNode this_);
extern void HarfangStopOnInstantiateAnimNode(HarfangNode this_);
extern HarfangMat4 HarfangGetWorldNode(HarfangNode this_);
extern void HarfangSetWorldNode(HarfangNode this_, const HarfangMat4 world);
extern HarfangMat4 HarfangComputeWorldNode(HarfangNode this_);
extern HarfangVec3 HarfangTransformTRSGetPos(HarfangTransformTRS h);
extern void HarfangTransformTRSSetPos(HarfangTransformTRS h, HarfangVec3 v);
extern HarfangVec3 HarfangTransformTRSGetRot(HarfangTransformTRS h);
extern void HarfangTransformTRSSetRot(HarfangTransformTRS h, HarfangVec3 v);
extern HarfangVec3 HarfangTransformTRSGetScl(HarfangTransformTRS h);
extern void HarfangTransformTRSSetScl(HarfangTransformTRS h, HarfangVec3 v);
extern HarfangTransformTRS HarfangConstructorTransformTRS();
extern void HarfangTransformTRSFree(HarfangTransformTRS);
extern void HarfangTransformFree(HarfangTransform);
extern bool HarfangEqTransform(HarfangTransform this_, const HarfangTransform t);
extern bool HarfangIsValidTransform(HarfangTransform this_);
extern HarfangVec3 HarfangGetPosTransform(HarfangTransform this_);
extern void HarfangSetPosTransform(HarfangTransform this_, const HarfangVec3 T);
extern HarfangVec3 HarfangGetRotTransform(HarfangTransform this_);
extern void HarfangSetRotTransform(HarfangTransform this_, const HarfangVec3 R);
extern HarfangVec3 HarfangGetScaleTransform(HarfangTransform this_);
extern void HarfangSetScaleTransform(HarfangTransform this_, const HarfangVec3 S);
extern HarfangTransformTRS HarfangGetTRSTransform(HarfangTransform this_);
extern void HarfangSetTRSTransform(HarfangTransform this_, const HarfangTransformTRS TRS);
extern void HarfangGetPosRotTransform(HarfangTransform this_, HarfangVec3 pos, HarfangVec3 rot);
extern void HarfangSetPosRotTransform(HarfangTransform this_, const HarfangVec3 pos, const HarfangVec3 rot);
extern HarfangNode HarfangGetParentTransform(HarfangTransform this_);
extern void HarfangSetParentTransform(HarfangTransform this_, const HarfangNode n);
extern void HarfangClearParentTransform(HarfangTransform this_);
extern HarfangMat4 HarfangGetWorldTransform(HarfangTransform this_);
extern void HarfangSetWorldTransform(HarfangTransform this_, const HarfangMat4 world);
extern void HarfangSetLocalTransform(HarfangTransform this_, const HarfangMat4 local);
extern float HarfangCameraZRangeGetZnear(HarfangCameraZRange h);
extern void HarfangCameraZRangeSetZnear(HarfangCameraZRange h, float v);
extern float HarfangCameraZRangeGetZfar(HarfangCameraZRange h);
extern void HarfangCameraZRangeSetZfar(HarfangCameraZRange h, float v);
extern HarfangCameraZRange HarfangConstructorCameraZRange();
extern void HarfangCameraZRangeFree(HarfangCameraZRange);
extern void HarfangCameraFree(HarfangCamera);
extern bool HarfangEqCamera(HarfangCamera this_, const HarfangCamera c);
extern bool HarfangIsValidCamera(HarfangCamera this_);
extern float HarfangGetZNearCamera(HarfangCamera this_);
extern void HarfangSetZNearCamera(HarfangCamera this_, float v);
extern float HarfangGetZFarCamera(HarfangCamera this_);
extern void HarfangSetZFarCamera(HarfangCamera this_, float v);
extern HarfangCameraZRange HarfangGetZRangeCamera(HarfangCamera this_);
extern void HarfangSetZRangeCamera(HarfangCamera this_, const HarfangCameraZRange z);
extern float HarfangGetFovCamera(HarfangCamera this_);
extern void HarfangSetFovCamera(HarfangCamera this_, float v);
extern bool HarfangGetIsOrthographicCamera(HarfangCamera this_);
extern void HarfangSetIsOrthographicCamera(HarfangCamera this_, bool v);
extern float HarfangGetSizeCamera(HarfangCamera this_);
extern void HarfangSetSizeCamera(HarfangCamera this_, float v);
extern void HarfangObjectFree(HarfangObject);
extern bool HarfangEqObject(HarfangObject this_, const HarfangObject o);
extern bool HarfangIsValidObject(HarfangObject this_);
extern HarfangModelRef HarfangGetModelRefObject(HarfangObject this_);
extern void HarfangSetModelRefObject(HarfangObject this_, const HarfangModelRef r);
extern void HarfangClearModelRefObject(HarfangObject this_);
extern HarfangMaterial HarfangGetMaterialObject(HarfangObject this_, size_t slot_idx);
extern HarfangMaterial HarfangGetMaterialObjectWithName(HarfangObject this_, const char *name);
extern void HarfangSetMaterialObject(HarfangObject this_, size_t slot_idx, HarfangMaterial mat);
extern size_t HarfangGetMaterialCountObject(HarfangObject this_);
extern void HarfangSetMaterialCountObject(HarfangObject this_, size_t count);
extern const char *HarfangGetMaterialNameObject(HarfangObject this_, size_t slot_idx);
extern void HarfangSetMaterialNameObject(HarfangObject this_, size_t slot_idx, const char *name);
extern bool HarfangGetMinMaxObject(HarfangObject this_, const HarfangPipelineResources resources, HarfangMinMax minmax);
extern size_t HarfangGetBoneCountObject(HarfangObject this_);
extern void HarfangSetBoneCountObject(HarfangObject this_, size_t count);
extern bool HarfangSetBoneObject(HarfangObject this_, size_t idx, const HarfangNode node);
extern HarfangNode HarfangGetBoneObject(HarfangObject this_, size_t idx);
extern void HarfangLightFree(HarfangLight);
extern bool HarfangEqLight(HarfangLight this_, const HarfangLight l);
extern bool HarfangIsValidLight(HarfangLight this_);
extern int HarfangGetTypeLight(HarfangLight this_);
extern void HarfangSetTypeLight(HarfangLight this_, int v);
extern int HarfangGetShadowTypeLight(HarfangLight this_);
extern void HarfangSetShadowTypeLight(HarfangLight this_, int v);
extern HarfangColor HarfangGetDiffuseColorLight(HarfangLight this_);
extern void HarfangSetDiffuseColorLight(HarfangLight this_, const HarfangColor v);
extern float HarfangGetDiffuseIntensityLight(HarfangLight this_);
extern void HarfangSetDiffuseIntensityLight(HarfangLight this_, float v);
extern HarfangColor HarfangGetSpecularColorLight(HarfangLight this_);
extern void HarfangSetSpecularColorLight(HarfangLight this_, const HarfangColor v);
extern float HarfangGetSpecularIntensityLight(HarfangLight this_);
extern void HarfangSetSpecularIntensityLight(HarfangLight this_, float v);
extern float HarfangGetRadiusLight(HarfangLight this_);
extern void HarfangSetRadiusLight(HarfangLight this_, float v);
extern float HarfangGetInnerAngleLight(HarfangLight this_);
extern void HarfangSetInnerAngleLight(HarfangLight this_, float v);
extern float HarfangGetOuterAngleLight(HarfangLight this_);
extern void HarfangSetOuterAngleLight(HarfangLight this_, float v);
extern HarfangVec4 HarfangGetPSSMSplitLight(HarfangLight this_);
extern void HarfangSetPSSMSplitLight(HarfangLight this_, const HarfangVec4 v);
extern float HarfangGetPriorityLight(HarfangLight this_);
extern void HarfangSetPriorityLight(HarfangLight this_, float v);
extern HarfangVec3 HarfangContactGetP(HarfangContact h);
extern void HarfangContactSetP(HarfangContact h, HarfangVec3 v);
extern HarfangVec3 HarfangContactGetN(HarfangContact h);
extern void HarfangContactSetN(HarfangContact h, HarfangVec3 v);
extern float HarfangContactGetD(HarfangContact h);
extern void HarfangContactSetD(HarfangContact h, float v);
extern void HarfangContactFree(HarfangContact);
extern HarfangContact HarfangContactListGetOperator(HarfangContactList h, int id);
extern void HarfangContactListSetOperator(HarfangContactList h, int id, HarfangContact v);
extern int HarfangContactListLenOperator(HarfangContactList h);
extern HarfangContactList HarfangConstructorContactList();
extern HarfangContactList HarfangConstructorContactListWithSequence(size_t sequenceToCSize, HarfangContact *sequenceToCBuf);
extern void HarfangContactListFree(HarfangContactList);
extern void HarfangClearContactList(HarfangContactList this_);
extern void HarfangReserveContactList(HarfangContactList this_, size_t size);
extern void HarfangPushBackContactList(HarfangContactList this_, HarfangContact v);
extern size_t HarfangSizeContactList(HarfangContactList this_);
extern HarfangContact HarfangAtContactList(HarfangContactList this_, size_t idx);
extern void HarfangRigidBodyFree(HarfangRigidBody);
extern bool HarfangEqRigidBody(HarfangRigidBody this_, const HarfangRigidBody b);
extern bool HarfangIsValidRigidBody(HarfangRigidBody this_);
extern uint8_t HarfangGetTypeRigidBody(HarfangRigidBody this_);
extern void HarfangSetTypeRigidBody(HarfangRigidBody this_, uint8_t type);
extern float HarfangGetLinearDampingRigidBody(HarfangRigidBody this_);
extern void HarfangSetLinearDampingRigidBody(HarfangRigidBody this_, float damping);
extern float HarfangGetAngularDampingRigidBody(HarfangRigidBody this_);
extern void HarfangSetAngularDampingRigidBody(HarfangRigidBody this_, float damping);
extern float HarfangGetRestitutionRigidBody(HarfangRigidBody this_);
extern void HarfangSetRestitutionRigidBody(HarfangRigidBody this_, float restitution);
extern float HarfangGetFrictionRigidBody(HarfangRigidBody this_);
extern void HarfangSetFrictionRigidBody(HarfangRigidBody this_, float friction);
extern float HarfangGetRollingFrictionRigidBody(HarfangRigidBody this_);
extern void HarfangSetRollingFrictionRigidBody(HarfangRigidBody this_, float rolling_friction);
extern void HarfangCollisionFree(HarfangCollision);
extern bool HarfangEqCollision(HarfangCollision this_, const HarfangCollision c);
extern bool HarfangIsValidCollision(HarfangCollision this_);
extern uint8_t HarfangGetTypeCollision(HarfangCollision this_);
extern void HarfangSetTypeCollision(HarfangCollision this_, uint8_t type);
extern HarfangMat4 HarfangGetLocalTransformCollision(HarfangCollision this_);
extern void HarfangSetLocalTransformCollision(HarfangCollision this_, HarfangMat4 m);
extern float HarfangGetMassCollision(HarfangCollision this_);
extern void HarfangSetMassCollision(HarfangCollision this_, float mass);
extern float HarfangGetRadiusCollision(HarfangCollision this_);
extern void HarfangSetRadiusCollision(HarfangCollision this_, float radius);
extern float HarfangGetHeightCollision(HarfangCollision this_);
extern void HarfangSetHeightCollision(HarfangCollision this_, float height);
extern void HarfangSetSizeCollision(HarfangCollision this_, const HarfangVec3 size);
extern const char *HarfangGetCollisionResourceCollision(HarfangCollision this_);
extern void HarfangSetCollisionResourceCollision(HarfangCollision this_, const char *path);
extern void HarfangInstanceFree(HarfangInstance);
extern bool HarfangEqInstance(HarfangInstance this_, const HarfangInstance i);
extern bool HarfangIsValidInstance(HarfangInstance this_);
extern const char *HarfangGetPathInstance(HarfangInstance this_);
extern void HarfangSetPathInstance(HarfangInstance this_, const char *path);
extern void HarfangSetOnInstantiateAnimInstance(HarfangInstance this_, const char *anim);
extern void HarfangSetOnInstantiateAnimLoopModeInstance(HarfangInstance this_, int loop_mode);
extern void HarfangClearOnInstantiateAnimInstance(HarfangInstance this_);
extern const char *HarfangGetOnInstantiateAnimInstance(HarfangInstance this_);
extern int HarfangGetOnInstantiateAnimLoopModeInstance(HarfangInstance this_);
extern HarfangScenePlayAnimRef HarfangGetOnInstantiatePlayAnimRefInstance(HarfangInstance this_);
extern void HarfangScriptFree(HarfangScript);
extern bool HarfangEqScript(HarfangScript this_, const HarfangScript s);
extern bool HarfangIsValidScript(HarfangScript this_);
extern const char *HarfangGetPathScript(HarfangScript this_);
extern void HarfangSetPathScript(HarfangScript this_, const char *path);
extern HarfangScript HarfangScriptListGetOperator(HarfangScriptList h, int id);
extern void HarfangScriptListSetOperator(HarfangScriptList h, int id, HarfangScript v);
extern int HarfangScriptListLenOperator(HarfangScriptList h);
extern HarfangScriptList HarfangConstructorScriptList();
extern HarfangScriptList HarfangConstructorScriptListWithSequence(size_t sequenceToCSize, HarfangScript *sequenceToCBuf);
extern void HarfangScriptListFree(HarfangScriptList);
extern void HarfangClearScriptList(HarfangScriptList this_);
extern void HarfangReserveScriptList(HarfangScriptList this_, size_t size);
extern void HarfangPushBackScriptList(HarfangScriptList this_, HarfangScript v);
extern size_t HarfangSizeScriptList(HarfangScriptList this_);
extern HarfangScript HarfangAtScriptList(HarfangScriptList this_, size_t idx);
extern HarfangNode HarfangNodeListGetOperator(HarfangNodeList h, int id);
extern void HarfangNodeListSetOperator(HarfangNodeList h, int id, HarfangNode v);
extern int HarfangNodeListLenOperator(HarfangNodeList h);
extern HarfangNodeList HarfangConstructorNodeList();
extern HarfangNodeList HarfangConstructorNodeListWithSequence(size_t sequenceToCSize, HarfangNode *sequenceToCBuf);
extern void HarfangNodeListFree(HarfangNodeList);
extern void HarfangClearNodeList(HarfangNodeList this_);
extern void HarfangReserveNodeList(HarfangNodeList this_, size_t size);
extern void HarfangPushBackNodeList(HarfangNodeList this_, HarfangNode v);
extern size_t HarfangSizeNodeList(HarfangNodeList this_);
extern HarfangNode HarfangAtNodeList(HarfangNodeList this_, size_t idx);
extern HarfangVec3 HarfangRaycastOutGetP(HarfangRaycastOut h);
extern void HarfangRaycastOutSetP(HarfangRaycastOut h, HarfangVec3 v);
extern HarfangVec3 HarfangRaycastOutGetN(HarfangRaycastOut h);
extern void HarfangRaycastOutSetN(HarfangRaycastOut h, HarfangVec3 v);
extern HarfangNode HarfangRaycastOutGetNode(HarfangRaycastOut h);
extern void HarfangRaycastOutSetNode(HarfangRaycastOut h, HarfangNode v);
extern float HarfangRaycastOutGetT(HarfangRaycastOut h);
extern void HarfangRaycastOutSetT(HarfangRaycastOut h, float v);
extern void HarfangRaycastOutFree(HarfangRaycastOut);
extern HarfangRaycastOut HarfangRaycastOutListGetOperator(HarfangRaycastOutList h, int id);
extern void HarfangRaycastOutListSetOperator(HarfangRaycastOutList h, int id, HarfangRaycastOut v);
extern int HarfangRaycastOutListLenOperator(HarfangRaycastOutList h);
extern HarfangRaycastOutList HarfangConstructorRaycastOutList();
extern HarfangRaycastOutList HarfangConstructorRaycastOutListWithSequence(size_t sequenceToCSize, HarfangRaycastOut *sequenceToCBuf);
extern void HarfangRaycastOutListFree(HarfangRaycastOutList);
extern void HarfangClearRaycastOutList(HarfangRaycastOutList this_);
extern void HarfangReserveRaycastOutList(HarfangRaycastOutList this_, size_t size);
extern void HarfangPushBackRaycastOutList(HarfangRaycastOutList this_, HarfangRaycastOut v);
extern size_t HarfangSizeRaycastOutList(HarfangRaycastOutList this_);
extern HarfangRaycastOut HarfangAtRaycastOutList(HarfangRaycastOutList this_, size_t idx);
extern void HarfangTimeCallbackConnectionFree(HarfangTimeCallbackConnection);
extern void HarfangSignalReturningVoidTakingTimeNsFree(HarfangSignalReturningVoidTakingTimeNs);
extern HarfangTimeCallbackConnection HarfangConnectSignalReturningVoidTakingTimeNs(
	HarfangSignalReturningVoidTakingTimeNs this_, HarfangFunctionReturningVoidTakingTimeNs listener);
extern void HarfangDisconnectSignalReturningVoidTakingTimeNs(HarfangSignalReturningVoidTakingTimeNs this_, HarfangTimeCallbackConnection connection);
extern void HarfangDisconnectAllSignalReturningVoidTakingTimeNs(HarfangSignalReturningVoidTakingTimeNs this_);
extern void HarfangEmitSignalReturningVoidTakingTimeNs(HarfangSignalReturningVoidTakingTimeNs this_, int64_t arg0);
extern size_t HarfangGetListenerCountSignalReturningVoidTakingTimeNs(HarfangSignalReturningVoidTakingTimeNs this_);
extern bool HarfangCanvasGetClearZ(HarfangCanvas h);
extern void HarfangCanvasSetClearZ(HarfangCanvas h, bool v);
extern bool HarfangCanvasGetClearColor(HarfangCanvas h);
extern void HarfangCanvasSetClearColor(HarfangCanvas h, bool v);
extern HarfangColor HarfangCanvasGetColor(HarfangCanvas h);
extern void HarfangCanvasSetColor(HarfangCanvas h, HarfangColor v);
extern void HarfangCanvasFree(HarfangCanvas);
extern HarfangColor HarfangEnvironmentGetAmbient(HarfangEnvironment h);
extern void HarfangEnvironmentSetAmbient(HarfangEnvironment h, HarfangColor v);
extern HarfangColor HarfangEnvironmentGetFogColor(HarfangEnvironment h);
extern void HarfangEnvironmentSetFogColor(HarfangEnvironment h, HarfangColor v);
extern float HarfangEnvironmentGetFogNear(HarfangEnvironment h);
extern void HarfangEnvironmentSetFogNear(HarfangEnvironment h, float v);
extern float HarfangEnvironmentGetFogFar(HarfangEnvironment h);
extern void HarfangEnvironmentSetFogFar(HarfangEnvironment h, float v);
extern HarfangTextureRef HarfangEnvironmentGetBrdfMap(HarfangEnvironment h);
extern void HarfangEnvironmentSetBrdfMap(HarfangEnvironment h, HarfangTextureRef v);
extern void HarfangEnvironmentFree(HarfangEnvironment);
extern HarfangSceneForwardPipelinePassViewId HarfangConstructorSceneForwardPipelinePassViewId();
extern void HarfangSceneForwardPipelinePassViewIdFree(HarfangSceneForwardPipelinePassViewId);
extern HarfangSceneForwardPipelineRenderData HarfangConstructorSceneForwardPipelineRenderData();
extern void HarfangSceneForwardPipelineRenderDataFree(HarfangSceneForwardPipelineRenderData);
extern float HarfangForwardPipelineAAAConfigGetTemporalAaWeight(HarfangForwardPipelineAAAConfig h);
extern void HarfangForwardPipelineAAAConfigSetTemporalAaWeight(HarfangForwardPipelineAAAConfig h, float v);
extern int HarfangForwardPipelineAAAConfigGetSampleCount(HarfangForwardPipelineAAAConfig h);
extern void HarfangForwardPipelineAAAConfigSetSampleCount(HarfangForwardPipelineAAAConfig h, int v);
extern float HarfangForwardPipelineAAAConfigGetMaxDistance(HarfangForwardPipelineAAAConfig h);
extern void HarfangForwardPipelineAAAConfigSetMaxDistance(HarfangForwardPipelineAAAConfig h, float v);
extern float HarfangForwardPipelineAAAConfigGetZThickness(HarfangForwardPipelineAAAConfig h);
extern void HarfangForwardPipelineAAAConfigSetZThickness(HarfangForwardPipelineAAAConfig h, float v);
extern float HarfangForwardPipelineAAAConfigGetBloomThreshold(HarfangForwardPipelineAAAConfig h);
extern void HarfangForwardPipelineAAAConfigSetBloomThreshold(HarfangForwardPipelineAAAConfig h, float v);
extern float HarfangForwardPipelineAAAConfigGetBloomBias(HarfangForwardPipelineAAAConfig h);
extern void HarfangForwardPipelineAAAConfigSetBloomBias(HarfangForwardPipelineAAAConfig h, float v);
extern float HarfangForwardPipelineAAAConfigGetBloomIntensity(HarfangForwardPipelineAAAConfig h);
extern void HarfangForwardPipelineAAAConfigSetBloomIntensity(HarfangForwardPipelineAAAConfig h, float v);
extern float HarfangForwardPipelineAAAConfigGetMotionBlur(HarfangForwardPipelineAAAConfig h);
extern void HarfangForwardPipelineAAAConfigSetMotionBlur(HarfangForwardPipelineAAAConfig h, float v);
extern float HarfangForwardPipelineAAAConfigGetExposure(HarfangForwardPipelineAAAConfig h);
extern void HarfangForwardPipelineAAAConfigSetExposure(HarfangForwardPipelineAAAConfig h, float v);
extern float HarfangForwardPipelineAAAConfigGetGamma(HarfangForwardPipelineAAAConfig h);
extern void HarfangForwardPipelineAAAConfigSetGamma(HarfangForwardPipelineAAAConfig h, float v);
extern HarfangForwardPipelineAAAConfig HarfangConstructorForwardPipelineAAAConfig();
extern void HarfangForwardPipelineAAAConfigFree(HarfangForwardPipelineAAAConfig);
extern void HarfangForwardPipelineAAAFree(HarfangForwardPipelineAAA);
extern void HarfangFlipForwardPipelineAAA(HarfangForwardPipelineAAA this_, const HarfangViewState view_state);
extern HarfangNodePairContacts HarfangConstructorNodePairContacts();
extern void HarfangNodePairContactsFree(HarfangNodePairContacts);
extern void HarfangBtGeneric6DofConstraintFree(HarfangBtGeneric6DofConstraint);
extern HarfangSceneBullet3Physics HarfangConstructorSceneBullet3Physics();
extern HarfangSceneBullet3Physics HarfangConstructorSceneBullet3PhysicsWithThreadCount(int thread_count);
extern void HarfangSceneBullet3PhysicsFree(HarfangSceneBullet3Physics);
extern void HarfangSceneCreatePhysicsFromFileSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangScene scene);
extern void HarfangSceneCreatePhysicsFromAssetsSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangScene scene);
extern void HarfangNodeCreatePhysicsFromFileSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node);
extern void HarfangNodeCreatePhysicsFromAssetsSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node);
extern void HarfangNodeStartTrackingCollisionEventsSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node);
extern void HarfangNodeStartTrackingCollisionEventsSceneBullet3PhysicsWithMode(HarfangSceneBullet3Physics this_, const HarfangNode node, uint8_t mode);
extern void HarfangNodeStopTrackingCollisionEventsSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node);
extern void HarfangNodeDestroyPhysicsSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node);
extern bool HarfangNodeHasBodySceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node);
extern void HarfangStepSimulationSceneBullet3Physics(HarfangSceneBullet3Physics this_, int64_t display_dt);
extern void HarfangStepSimulationSceneBullet3PhysicsWithStepDt(HarfangSceneBullet3Physics this_, int64_t display_dt, int64_t step_dt);
extern void HarfangStepSimulationSceneBullet3PhysicsWithStepDtMaxStep(HarfangSceneBullet3Physics this_, int64_t display_dt, int64_t step_dt, int max_step);
extern void HarfangCollectCollisionEventsSceneBullet3Physics(
	HarfangSceneBullet3Physics this_, const HarfangScene scene, HarfangNodePairContacts node_pair_contacts);
extern void HarfangSyncTransformsFromSceneSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangScene scene);
extern void HarfangSyncTransformsToSceneSceneBullet3Physics(HarfangSceneBullet3Physics this_, HarfangScene scene);
extern size_t HarfangGarbageCollectSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangScene scene);
extern size_t HarfangGarbageCollectResourcesSceneBullet3Physics(HarfangSceneBullet3Physics this_);
extern void HarfangClearNodesSceneBullet3Physics(HarfangSceneBullet3Physics this_);
extern void HarfangClearSceneBullet3Physics(HarfangSceneBullet3Physics this_);
extern void HarfangNodeWakeSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node);
extern void HarfangNodeSetDeactivationSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, bool enable);
extern bool HarfangNodeGetDeactivationSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node);
extern void HarfangNodeResetWorldSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangMat4 world);
extern void HarfangNodeTeleportSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangMat4 world);
extern void HarfangNodeAddForceSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangVec3 F);
extern void HarfangNodeAddForceSceneBullet3PhysicsWithWorldPos(
	HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangVec3 F, const HarfangVec3 world_pos);
extern void HarfangNodeAddImpulseSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangVec3 dt_velocity);
extern void HarfangNodeAddImpulseSceneBullet3PhysicsWithWorldPos(
	HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangVec3 dt_velocity, const HarfangVec3 world_pos);
extern void HarfangNodeAddTorqueSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangVec3 T);
extern void HarfangNodeAddTorqueImpulseSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangVec3 dt_angular_velocity);
extern HarfangVec3 HarfangNodeGetPointVelocitySceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangVec3 world_pos);
extern HarfangVec3 HarfangNodeGetLinearVelocitySceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node);
extern void HarfangNodeSetLinearVelocitySceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangVec3 V);
extern HarfangVec3 HarfangNodeGetAngularVelocitySceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node);
extern void HarfangNodeSetAngularVelocitySceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangVec3 W);
extern HarfangVec3 HarfangNodeGetLinearFactorSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node);
extern void HarfangNodeSetLinearFactorSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangVec3 k);
extern HarfangVec3 HarfangNodeGetAngularFactorSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node);
extern void HarfangNodeSetAngularFactorSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangVec3 k);
extern HarfangBtGeneric6DofConstraint HarfangAdd6DofConstraintSceneBullet3Physics(
	HarfangSceneBullet3Physics this_, const HarfangNode nodeA, const HarfangNode nodeB, const HarfangMat4 anchorALocal, const HarfangMat4 anchorBInLocalSpaceA);
extern void HarfangRemove6DofConstraintSceneBullet3Physics(HarfangSceneBullet3Physics this_, HarfangBtGeneric6DofConstraint constraint6Dof);
extern HarfangNodePairContacts HarfangNodeCollideWorldSceneBullet3Physics(HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangMat4 world);
extern HarfangNodePairContacts HarfangNodeCollideWorldSceneBullet3PhysicsWithMaxContact(
	HarfangSceneBullet3Physics this_, const HarfangNode node, const HarfangMat4 world, int max_contact);
extern HarfangRaycastOut HarfangRaycastFirstHitSceneBullet3Physics(
	HarfangSceneBullet3Physics this_, const HarfangScene scene, const HarfangVec3 p0, const HarfangVec3 p1);
extern HarfangRaycastOutList HarfangRaycastAllHitsSceneBullet3Physics(
	HarfangSceneBullet3Physics this_, const HarfangScene scene, const HarfangVec3 p0, const HarfangVec3 p1);
extern void HarfangRenderCollisionSceneBullet3Physics(HarfangSceneBullet3Physics this_, uint16_t view_id, const HarfangVertexLayout vtx_layout,
	HarfangProgramHandle prg, HarfangRenderState render_state, uint32_t depth);
extern void HarfangSetPreTickCallbackSceneBullet3Physics(
	HarfangSceneBullet3Physics this_, const HarfangFunctionReturningVoidTakingSceneBullet3PhysicsRefTimeNs cbk);
extern HarfangSceneLuaVM HarfangConstructorSceneLuaVM();
extern void HarfangSceneLuaVMFree(HarfangSceneLuaVM);
extern bool HarfangCreateScriptFromSourceSceneLuaVM(HarfangSceneLuaVM this_, HarfangScene scene, const HarfangScript script, const char *src);
extern bool HarfangCreateScriptFromFileSceneLuaVM(HarfangSceneLuaVM this_, HarfangScene scene, const HarfangScript script);
extern bool HarfangCreateScriptFromAssetsSceneLuaVM(HarfangSceneLuaVM this_, HarfangScene scene, const HarfangScript script);
extern HarfangScriptList HarfangCreateNodeScriptsFromFileSceneLuaVM(HarfangSceneLuaVM this_, HarfangScene scene, const HarfangNode node);
extern HarfangScriptList HarfangCreateNodeScriptsFromAssetsSceneLuaVM(HarfangSceneLuaVM this_, HarfangScene scene, const HarfangNode node);
extern HarfangScriptList HarfangSceneCreateScriptsFromFileSceneLuaVM(HarfangSceneLuaVM this_, HarfangScene scene);
extern HarfangScriptList HarfangSceneCreateScriptsFromAssetsSceneLuaVM(HarfangSceneLuaVM this_, HarfangScene scene);
extern HarfangScriptList HarfangGarbageCollectSceneLuaVM(HarfangSceneLuaVM this_, const HarfangScene scene);
extern void HarfangDestroyScriptsSceneLuaVM(HarfangSceneLuaVM this_, const HarfangScriptList scripts);
extern HarfangStringList HarfangGetScriptInterfaceSceneLuaVM(HarfangSceneLuaVM this_, const HarfangScript script);
extern size_t HarfangGetScriptCountSceneLuaVM(HarfangSceneLuaVM this_);
extern HarfangLuaObject HarfangGetScriptEnvSceneLuaVM(HarfangSceneLuaVM this_, const HarfangScript script);
extern HarfangLuaObject HarfangGetScriptValueSceneLuaVM(HarfangSceneLuaVM this_, const HarfangScript script, const char *name);
extern bool HarfangSetScriptValueSceneLuaVM(HarfangSceneLuaVM this_, const HarfangScript script, const char *name, const HarfangLuaObject value);
extern bool HarfangSetScriptValueSceneLuaVMWithNotify(
	HarfangSceneLuaVM this_, const HarfangScript script, const char *name, const HarfangLuaObject value, bool notify);
extern bool HarfangCallSceneLuaVM(
	HarfangSceneLuaVM this_, const HarfangScript script, const char *function, const HarfangLuaObjectList args, HarfangLuaObjectList ret_vals);
extern bool HarfangCallSceneLuaVMWithSliceOfArgs(HarfangSceneLuaVM this_, const HarfangScript script, const char *function, size_t SliceOfargsToCSize,
	HarfangLuaObject *SliceOfargsToCBuf, HarfangLuaObjectList ret_vals);
extern HarfangLuaObject HarfangMakeLuaObjectSceneLuaVM(HarfangSceneLuaVM this_);
extern HarfangSceneClocks HarfangConstructorSceneClocks();
extern void HarfangSceneClocksFree(HarfangSceneClocks);
extern void HarfangMouseStateFree(HarfangMouseState);
extern int HarfangXMouseState(HarfangMouseState this_);
extern int HarfangYMouseState(HarfangMouseState this_);
extern bool HarfangButtonMouseState(HarfangMouseState this_, int btn);
extern int HarfangWheelMouseState(HarfangMouseState this_);
extern int HarfangHWheelMouseState(HarfangMouseState this_);
extern HarfangMouse HarfangConstructorMouse();
extern HarfangMouse HarfangConstructorMouseWithName(const char *name);
extern void HarfangMouseFree(HarfangMouse);
extern int HarfangXMouse(HarfangMouse this_);
extern int HarfangYMouse(HarfangMouse this_);
extern int HarfangDtXMouse(HarfangMouse this_);
extern int HarfangDtYMouse(HarfangMouse this_);
extern bool HarfangDownMouse(HarfangMouse this_, int button);
extern bool HarfangPressedMouse(HarfangMouse this_, int button);
extern bool HarfangReleasedMouse(HarfangMouse this_, int button);
extern int HarfangWheelMouse(HarfangMouse this_);
extern int HarfangHWheelMouse(HarfangMouse this_);
extern void HarfangUpdateMouse(HarfangMouse this_);
extern HarfangMouseState HarfangGetStateMouse(HarfangMouse this_);
extern HarfangMouseState HarfangGetOldStateMouse(HarfangMouse this_);
extern void HarfangKeyboardStateFree(HarfangKeyboardState);
extern bool HarfangKeyKeyboardState(HarfangKeyboardState this_, int key);
extern HarfangKeyboard HarfangConstructorKeyboard();
extern HarfangKeyboard HarfangConstructorKeyboardWithName(const char *name);
extern void HarfangKeyboardFree(HarfangKeyboard);
extern bool HarfangDownKeyboard(HarfangKeyboard this_, int key);
extern bool HarfangPressedKeyboard(HarfangKeyboard this_, int key);
extern bool HarfangReleasedKeyboard(HarfangKeyboard this_, int key);
extern void HarfangUpdateKeyboard(HarfangKeyboard this_);
extern HarfangKeyboardState HarfangGetStateKeyboard(HarfangKeyboard this_);
extern HarfangKeyboardState HarfangGetOldStateKeyboard(HarfangKeyboard this_);
extern void HarfangTextInputCallbackConnectionFree(HarfangTextInputCallbackConnection);
extern void HarfangSignalReturningVoidTakingConstCharPtrFree(HarfangSignalReturningVoidTakingConstCharPtr);
extern HarfangTextInputCallbackConnection HarfangConnectSignalReturningVoidTakingConstCharPtr(
	HarfangSignalReturningVoidTakingConstCharPtr this_, HarfangFunctionReturningVoidTakingConstCharPtr listener);
extern void HarfangDisconnectSignalReturningVoidTakingConstCharPtr(
	HarfangSignalReturningVoidTakingConstCharPtr this_, HarfangTextInputCallbackConnection connection);
extern void HarfangDisconnectAllSignalReturningVoidTakingConstCharPtr(HarfangSignalReturningVoidTakingConstCharPtr this_);
extern void HarfangEmitSignalReturningVoidTakingConstCharPtr(HarfangSignalReturningVoidTakingConstCharPtr this_, const char *arg0);
extern size_t HarfangGetListenerCountSignalReturningVoidTakingConstCharPtr(HarfangSignalReturningVoidTakingConstCharPtr this_);
extern void HarfangGamepadStateFree(HarfangGamepadState);
extern bool HarfangIsConnectedGamepadState(HarfangGamepadState this_);
extern float HarfangAxesGamepadState(HarfangGamepadState this_, int idx);
extern bool HarfangButtonGamepadState(HarfangGamepadState this_, int btn);
extern HarfangGamepad HarfangConstructorGamepad();
extern HarfangGamepad HarfangConstructorGamepadWithName(const char *name);
extern void HarfangGamepadFree(HarfangGamepad);
extern bool HarfangIsConnectedGamepad(HarfangGamepad this_);
extern bool HarfangConnectedGamepad(HarfangGamepad this_);
extern bool HarfangDisconnectedGamepad(HarfangGamepad this_);
extern float HarfangAxesGamepad(HarfangGamepad this_, int axis);
extern float HarfangDtAxesGamepad(HarfangGamepad this_, int axis);
extern bool HarfangDownGamepad(HarfangGamepad this_, int btn);
extern bool HarfangPressedGamepad(HarfangGamepad this_, int btn);
extern bool HarfangReleasedGamepad(HarfangGamepad this_, int btn);
extern void HarfangUpdateGamepad(HarfangGamepad this_);
extern void HarfangJoystickStateFree(HarfangJoystickState);
extern bool HarfangIsConnectedJoystickState(HarfangJoystickState this_);
extern float HarfangAxesJoystickState(HarfangJoystickState this_, int idx);
extern bool HarfangButtonJoystickState(HarfangJoystickState this_, int btn);
extern HarfangJoystick HarfangConstructorJoystick();
extern HarfangJoystick HarfangConstructorJoystickWithName(const char *name);
extern void HarfangJoystickFree(HarfangJoystick);
extern bool HarfangIsConnectedJoystick(HarfangJoystick this_);
extern bool HarfangConnectedJoystick(HarfangJoystick this_);
extern bool HarfangDisconnectedJoystick(HarfangJoystick this_);
extern int HarfangAxesCountJoystick(HarfangJoystick this_);
extern float HarfangAxesJoystick(HarfangJoystick this_, int axis);
extern float HarfangDtAxesJoystick(HarfangJoystick this_, int axis);
extern int HarfangButtonsCountJoystick(HarfangJoystick this_);
extern bool HarfangDownJoystick(HarfangJoystick this_, int btn);
extern bool HarfangPressedJoystick(HarfangJoystick this_, int btn);
extern bool HarfangReleasedJoystick(HarfangJoystick this_, int btn);
extern void HarfangUpdateJoystick(HarfangJoystick this_);
extern const char *HarfangGetDeviceNameJoystick(HarfangJoystick this_);
extern void HarfangVRControllerStateFree(HarfangVRControllerState);
extern bool HarfangIsConnectedVRControllerState(HarfangVRControllerState this_);
extern HarfangMat4 HarfangWorldVRControllerState(HarfangVRControllerState this_);
extern bool HarfangPressedVRControllerState(HarfangVRControllerState this_, int btn);
extern bool HarfangTouchedVRControllerState(HarfangVRControllerState this_, int btn);
extern HarfangVec2 HarfangSurfaceVRControllerState(HarfangVRControllerState this_, int idx);
extern HarfangVRController HarfangConstructorVRController();
extern HarfangVRController HarfangConstructorVRControllerWithName(const char *name);
extern void HarfangVRControllerFree(HarfangVRController);
extern bool HarfangIsConnectedVRController(HarfangVRController this_);
extern bool HarfangConnectedVRController(HarfangVRController this_);
extern bool HarfangDisconnectedVRController(HarfangVRController this_);
extern HarfangMat4 HarfangWorldVRController(HarfangVRController this_);
extern bool HarfangDownVRController(HarfangVRController this_, int btn);
extern bool HarfangPressedVRController(HarfangVRController this_, int btn);
extern bool HarfangReleasedVRController(HarfangVRController this_, int btn);
extern bool HarfangTouchVRController(HarfangVRController this_, int btn);
extern bool HarfangTouchStartVRController(HarfangVRController this_, int btn);
extern bool HarfangTouchEndVRController(HarfangVRController this_, int btn);
extern HarfangVec2 HarfangSurfaceVRController(HarfangVRController this_, int idx);
extern HarfangVec2 HarfangDtSurfaceVRController(HarfangVRController this_, int idx);
extern void HarfangSendHapticPulseVRController(HarfangVRController this_, int64_t duration);
extern void HarfangUpdateVRController(HarfangVRController this_);
extern void HarfangVRGenericTrackerStateFree(HarfangVRGenericTrackerState);
extern bool HarfangIsConnectedVRGenericTrackerState(HarfangVRGenericTrackerState this_);
extern HarfangMat4 HarfangWorldVRGenericTrackerState(HarfangVRGenericTrackerState this_);
extern HarfangVRGenericTracker HarfangConstructorVRGenericTracker();
extern HarfangVRGenericTracker HarfangConstructorVRGenericTrackerWithName(const char *name);
extern void HarfangVRGenericTrackerFree(HarfangVRGenericTracker);
extern bool HarfangIsConnectedVRGenericTracker(HarfangVRGenericTracker this_);
extern HarfangMat4 HarfangWorldVRGenericTracker(HarfangVRGenericTracker this_);
extern void HarfangUpdateVRGenericTracker(HarfangVRGenericTracker this_);
extern void HarfangDearImguiContextFree(HarfangDearImguiContext);
extern void HarfangImFontFree(HarfangImFont);
extern void HarfangImDrawListFree(HarfangImDrawList);
extern void HarfangPushClipRectImDrawList(HarfangImDrawList this_, HarfangVec2 clip_rect_min, HarfangVec2 clip_rect_max);
extern void HarfangPushClipRectImDrawListWithIntersectWithCurentClipRect(
	HarfangImDrawList this_, HarfangVec2 clip_rect_min, HarfangVec2 clip_rect_max, bool intersect_with_curent_clip_rect);
extern void HarfangPushClipRectFullScreenImDrawList(HarfangImDrawList this_);
extern void HarfangPopClipRectImDrawList(HarfangImDrawList this_);
extern void HarfangPushTextureIDImDrawList(HarfangImDrawList this_, const HarfangTexture tex);
extern void HarfangPopTextureIDImDrawList(HarfangImDrawList this_);
extern HarfangVec2 HarfangGetClipRectMinImDrawList(HarfangImDrawList this_);
extern HarfangVec2 HarfangGetClipRectMaxImDrawList(HarfangImDrawList this_);
extern void HarfangAddLineImDrawList(HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, unsigned int col);
extern void HarfangAddLineImDrawListWithThickness(HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, unsigned int col, float thickness);
extern void HarfangAddRectImDrawList(HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, unsigned int col);
extern void HarfangAddRectImDrawListWithRounding(HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, unsigned int col, float rounding);
extern void HarfangAddRectImDrawListWithRoundingRoundingCornerFlags(
	HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, unsigned int col, float rounding, int rounding_corner_flags);
extern void HarfangAddRectImDrawListWithRoundingRoundingCornerFlagsThickness(
	HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, unsigned int col, float rounding, int rounding_corner_flags, float thickness);
extern void HarfangAddRectFilledImDrawList(HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, unsigned int col);
extern void HarfangAddRectFilledImDrawListWithRounding(HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, unsigned int col, float rounding);
extern void HarfangAddRectFilledImDrawListWithRoundingRoundingCornerFlags(
	HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, unsigned int col, float rounding, int rounding_corner_flags);
extern void HarfangAddRectFilledMultiColorImDrawList(HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, unsigned int col_upr_left,
	unsigned int col_upr_right, unsigned int col_bot_right, unsigned int col_bot_left);
extern void HarfangAddQuadImDrawList(
	HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, const HarfangVec2 c, const HarfangVec2 d, unsigned int col);
extern void HarfangAddQuadImDrawListWithThickness(
	HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, const HarfangVec2 c, const HarfangVec2 d, unsigned int col, float thickness);
extern void HarfangAddQuadFilledImDrawList(
	HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, const HarfangVec2 c, const HarfangVec2 d, unsigned int col);
extern void HarfangAddTriangleImDrawList(HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, const HarfangVec2 c, unsigned int col);
extern void HarfangAddTriangleImDrawListWithThickness(
	HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, const HarfangVec2 c, unsigned int col, float thickness);
extern void HarfangAddTriangleFilledImDrawList(HarfangImDrawList this_, const HarfangVec2 a, const HarfangVec2 b, const HarfangVec2 c, unsigned int col);
extern void HarfangAddCircleImDrawList(HarfangImDrawList this_, const HarfangVec2 centre, float radius, unsigned int col);
extern void HarfangAddCircleImDrawListWithNumSegments(HarfangImDrawList this_, const HarfangVec2 centre, float radius, unsigned int col, int num_segments);
extern void HarfangAddCircleImDrawListWithNumSegmentsThickness(
	HarfangImDrawList this_, const HarfangVec2 centre, float radius, unsigned int col, int num_segments, float thickness);
extern void HarfangAddCircleFilledImDrawList(HarfangImDrawList this_, const HarfangVec2 centre, float radius, unsigned int col);
extern void HarfangAddCircleFilledImDrawListWithNumSegments(
	HarfangImDrawList this_, const HarfangVec2 centre, float radius, unsigned int col, int num_segments);
extern void HarfangAddTextImDrawList(HarfangImDrawList this_, const HarfangVec2 pos, unsigned int col, const char *text);
extern void HarfangAddTextImDrawListWithFontFontSizePosColText(
	HarfangImDrawList this_, const HarfangImFont font, float font_size, const HarfangVec2 pos, unsigned int col, const char *text);
extern void HarfangAddTextImDrawListWithFontFontSizePosColTextWrapWidth(
	HarfangImDrawList this_, const HarfangImFont font, float font_size, const HarfangVec2 pos, unsigned int col, const char *text, float wrap_width);
extern void HarfangAddTextImDrawListWithFontFontSizePosColTextWrapWidthCpuFineClipRect(HarfangImDrawList this_, const HarfangImFont font, float font_size,
	const HarfangVec2 pos, unsigned int col, const char *text, float wrap_width, const HarfangVec4 cpu_fine_clip_rect);
extern void HarfangAddImageImDrawList(HarfangImDrawList this_, const HarfangTexture tex, const HarfangVec2 a, const HarfangVec2 b);
extern void HarfangAddImageImDrawListWithUvAUvB(
	HarfangImDrawList this_, const HarfangTexture tex, const HarfangVec2 a, const HarfangVec2 b, const HarfangVec2 uv_a, const HarfangVec2 uv_b);
extern void HarfangAddImageImDrawListWithUvAUvBCol(HarfangImDrawList this_, const HarfangTexture tex, const HarfangVec2 a, const HarfangVec2 b,
	const HarfangVec2 uv_a, const HarfangVec2 uv_b, unsigned int col);
extern void HarfangAddImageQuadImDrawList(
	HarfangImDrawList this_, const HarfangTexture tex, const HarfangVec2 a, const HarfangVec2 b, const HarfangVec2 c, const HarfangVec2 d);
extern void HarfangAddImageQuadImDrawListWithUvAUvBUvCUvD(HarfangImDrawList this_, const HarfangTexture tex, const HarfangVec2 a, const HarfangVec2 b,
	const HarfangVec2 c, const HarfangVec2 d, const HarfangVec2 uv_a, const HarfangVec2 uv_b, const HarfangVec2 uv_c, const HarfangVec2 uv_d);
extern void HarfangAddImageQuadImDrawListWithUvAUvBUvCUvDCol(HarfangImDrawList this_, const HarfangTexture tex, const HarfangVec2 a, const HarfangVec2 b,
	const HarfangVec2 c, const HarfangVec2 d, const HarfangVec2 uv_a, const HarfangVec2 uv_b, const HarfangVec2 uv_c, const HarfangVec2 uv_d, unsigned int col);
extern void HarfangAddImageRoundedImDrawList(HarfangImDrawList this_, const HarfangTexture tex, const HarfangVec2 a, const HarfangVec2 b,
	const HarfangVec2 uv_a, const HarfangVec2 uv_b, unsigned int col, float rounding);
extern void HarfangAddImageRoundedImDrawListWithFlags(HarfangImDrawList this_, const HarfangTexture tex, const HarfangVec2 a, const HarfangVec2 b,
	const HarfangVec2 uv_a, const HarfangVec2 uv_b, unsigned int col, float rounding, int flags);
extern void HarfangAddPolylineImDrawList(HarfangImDrawList this_, const HarfangVec2List points, unsigned int col, bool closed, float thickness);
extern void HarfangAddConvexPolyFilledImDrawList(HarfangImDrawList this_, const HarfangVec2List points, unsigned int col);
extern void HarfangAddBezierCubicImDrawList(
	HarfangImDrawList this_, const HarfangVec2 pos0, const HarfangVec2 cp0, const HarfangVec2 cp1, const HarfangVec2 pos1, unsigned int col, float thickness);
extern void HarfangAddBezierCubicImDrawListWithNumSegments(HarfangImDrawList this_, const HarfangVec2 pos0, const HarfangVec2 cp0, const HarfangVec2 cp1,
	const HarfangVec2 pos1, unsigned int col, float thickness, int num_segments);
extern void HarfangPathClearImDrawList(HarfangImDrawList this_);
extern void HarfangPathLineToImDrawList(HarfangImDrawList this_, const HarfangVec2 pos);
extern void HarfangPathLineToMergeDuplicateImDrawList(HarfangImDrawList this_, const HarfangVec2 pos);
extern void HarfangPathFillConvexImDrawList(HarfangImDrawList this_, unsigned int col);
extern void HarfangPathStrokeImDrawList(HarfangImDrawList this_, unsigned int col, bool closed);
extern void HarfangPathStrokeImDrawListWithThickness(HarfangImDrawList this_, unsigned int col, bool closed, float thickness);
extern void HarfangPathArcToImDrawList(HarfangImDrawList this_, const HarfangVec2 centre, float radius, float a_min, float a_max);
extern void HarfangPathArcToImDrawListWithNumSegments(
	HarfangImDrawList this_, const HarfangVec2 centre, float radius, float a_min, float a_max, int num_segments);
extern void HarfangPathArcToFastImDrawList(HarfangImDrawList this_, const HarfangVec2 centre, float radius, int a_min_of_12, int a_max_of_12);
extern void HarfangPathBezierCubicCurveToImDrawList(HarfangImDrawList this_, const HarfangVec2 p1, const HarfangVec2 p2, const HarfangVec2 p3);
extern void HarfangPathBezierCubicCurveToImDrawListWithNumSegments(
	HarfangImDrawList this_, const HarfangVec2 p1, const HarfangVec2 p2, const HarfangVec2 p3, int num_segments);
extern void HarfangPathRectImDrawList(HarfangImDrawList this_, const HarfangVec2 rect_min, const HarfangVec2 rect_max);
extern void HarfangPathRectImDrawListWithRounding(HarfangImDrawList this_, const HarfangVec2 rect_min, const HarfangVec2 rect_max, float rounding);
extern void HarfangPathRectImDrawListWithRoundingFlags(
	HarfangImDrawList this_, const HarfangVec2 rect_min, const HarfangVec2 rect_max, float rounding, int flags);
extern void HarfangChannelsSplitImDrawList(HarfangImDrawList this_, int channels_count);
extern void HarfangChannelsMergeImDrawList(HarfangImDrawList this_);
extern void HarfangChannelsSetCurrentImDrawList(HarfangImDrawList this_, int channel_index);
extern const char *HarfangFileFilterGetName(HarfangFileFilter h);
extern void HarfangFileFilterSetName(HarfangFileFilter h, const char *v);
extern const char *HarfangFileFilterGetPattern(HarfangFileFilter h);
extern void HarfangFileFilterSetPattern(HarfangFileFilter h, const char *v);
extern HarfangFileFilter HarfangConstructorFileFilter();
extern void HarfangFileFilterFree(HarfangFileFilter);
extern HarfangFileFilter HarfangFileFilterListGetOperator(HarfangFileFilterList h, int id);
extern void HarfangFileFilterListSetOperator(HarfangFileFilterList h, int id, HarfangFileFilter v);
extern int HarfangFileFilterListLenOperator(HarfangFileFilterList h);
extern HarfangFileFilterList HarfangConstructorFileFilterList();
extern HarfangFileFilterList HarfangConstructorFileFilterListWithSequence(size_t sequenceToCSize, HarfangFileFilter *sequenceToCBuf);
extern void HarfangFileFilterListFree(HarfangFileFilterList);
extern void HarfangClearFileFilterList(HarfangFileFilterList this_);
extern void HarfangReserveFileFilterList(HarfangFileFilterList this_, size_t size);
extern void HarfangPushBackFileFilterList(HarfangFileFilterList this_, HarfangFileFilter v);
extern size_t HarfangSizeFileFilterList(HarfangFileFilterList this_);
extern HarfangFileFilter HarfangAtFileFilterList(HarfangFileFilterList this_, size_t idx);
extern float HarfangStereoSourceStateGetVolume(HarfangStereoSourceState h);
extern void HarfangStereoSourceStateSetVolume(HarfangStereoSourceState h, float v);
extern int HarfangStereoSourceStateGetRepeat(HarfangStereoSourceState h);
extern void HarfangStereoSourceStateSetRepeat(HarfangStereoSourceState h, int v);
extern float HarfangStereoSourceStateGetPanning(HarfangStereoSourceState h);
extern void HarfangStereoSourceStateSetPanning(HarfangStereoSourceState h, float v);
extern HarfangStereoSourceState HarfangConstructorStereoSourceState();
extern HarfangStereoSourceState HarfangConstructorStereoSourceStateWithVolume(float volume);
extern HarfangStereoSourceState HarfangConstructorStereoSourceStateWithVolumeRepeat(float volume, int repeat);
extern HarfangStereoSourceState HarfangConstructorStereoSourceStateWithVolumeRepeatPanning(float volume, int repeat, float panning);
extern void HarfangStereoSourceStateFree(HarfangStereoSourceState);
extern HarfangMat4 HarfangSpatializedSourceStateGetMtx(HarfangSpatializedSourceState h);
extern void HarfangSpatializedSourceStateSetMtx(HarfangSpatializedSourceState h, HarfangMat4 v);
extern float HarfangSpatializedSourceStateGetVolume(HarfangSpatializedSourceState h);
extern void HarfangSpatializedSourceStateSetVolume(HarfangSpatializedSourceState h, float v);
extern int HarfangSpatializedSourceStateGetRepeat(HarfangSpatializedSourceState h);
extern void HarfangSpatializedSourceStateSetRepeat(HarfangSpatializedSourceState h, int v);
extern HarfangVec3 HarfangSpatializedSourceStateGetVel(HarfangSpatializedSourceState h);
extern void HarfangSpatializedSourceStateSetVel(HarfangSpatializedSourceState h, HarfangVec3 v);
extern HarfangSpatializedSourceState HarfangConstructorSpatializedSourceState();
extern HarfangSpatializedSourceState HarfangConstructorSpatializedSourceStateWithMtx(HarfangMat4 mtx);
extern HarfangSpatializedSourceState HarfangConstructorSpatializedSourceStateWithMtxVolume(HarfangMat4 mtx, float volume);
extern HarfangSpatializedSourceState HarfangConstructorSpatializedSourceStateWithMtxVolumeRepeat(HarfangMat4 mtx, float volume, int repeat);
extern HarfangSpatializedSourceState HarfangConstructorSpatializedSourceStateWithMtxVolumeRepeatVel(
	HarfangMat4 mtx, float volume, int repeat, const HarfangVec3 vel);
extern void HarfangSpatializedSourceStateFree(HarfangSpatializedSourceState);
extern HarfangMat4 HarfangOpenVREyeGetOffset(HarfangOpenVREye h);
extern void HarfangOpenVREyeSetOffset(HarfangOpenVREye h, HarfangMat4 v);
extern HarfangMat44 HarfangOpenVREyeGetProjection(HarfangOpenVREye h);
extern void HarfangOpenVREyeSetProjection(HarfangOpenVREye h, HarfangMat44 v);
extern void HarfangOpenVREyeFree(HarfangOpenVREye);
extern void HarfangOpenVREyeFrameBufferFree(HarfangOpenVREyeFrameBuffer);
extern HarfangFrameBufferHandle HarfangGetHandleOpenVREyeFrameBuffer(HarfangOpenVREyeFrameBuffer this_);
extern HarfangMat4 HarfangOpenVRStateGetBody(HarfangOpenVRState h);
extern void HarfangOpenVRStateSetBody(HarfangOpenVRState h, HarfangMat4 v);
extern HarfangMat4 HarfangOpenVRStateGetHead(HarfangOpenVRState h);
extern void HarfangOpenVRStateSetHead(HarfangOpenVRState h, HarfangMat4 v);
extern HarfangMat4 HarfangOpenVRStateGetInvHead(HarfangOpenVRState h);
extern void HarfangOpenVRStateSetInvHead(HarfangOpenVRState h, HarfangMat4 v);
extern HarfangOpenVREye HarfangOpenVRStateGetLeft(HarfangOpenVRState h);
extern void HarfangOpenVRStateSetLeft(HarfangOpenVRState h, HarfangOpenVREye v);
extern HarfangOpenVREye HarfangOpenVRStateGetRight(HarfangOpenVRState h);
extern void HarfangOpenVRStateSetRight(HarfangOpenVRState h, HarfangOpenVREye v);
extern uint32_t HarfangOpenVRStateGetWidth(HarfangOpenVRState h);
extern void HarfangOpenVRStateSetWidth(HarfangOpenVRState h, uint32_t v);
extern uint32_t HarfangOpenVRStateGetHeight(HarfangOpenVRState h);
extern void HarfangOpenVRStateSetHeight(HarfangOpenVRState h, uint32_t v);
extern void HarfangOpenVRStateFree(HarfangOpenVRState);
extern void HarfangOpenXREyeFrameBufferFree(HarfangOpenXREyeFrameBuffer);
extern HarfangOpenXREyeFrameBuffer HarfangOpenXREyeFrameBufferListGetOperator(HarfangOpenXREyeFrameBufferList h, int id);
extern void HarfangOpenXREyeFrameBufferListSetOperator(HarfangOpenXREyeFrameBufferList h, int id, HarfangOpenXREyeFrameBuffer v);
extern int HarfangOpenXREyeFrameBufferListLenOperator(HarfangOpenXREyeFrameBufferList h);
extern HarfangOpenXREyeFrameBufferList HarfangConstructorOpenXREyeFrameBufferList();
extern HarfangOpenXREyeFrameBufferList HarfangConstructorOpenXREyeFrameBufferListWithSequence(
	size_t sequenceToCSize, HarfangOpenXREyeFrameBuffer *sequenceToCBuf);
extern void HarfangOpenXREyeFrameBufferListFree(HarfangOpenXREyeFrameBufferList);
extern void HarfangClearOpenXREyeFrameBufferList(HarfangOpenXREyeFrameBufferList this_);
extern void HarfangReserveOpenXREyeFrameBufferList(HarfangOpenXREyeFrameBufferList this_, size_t size);
extern void HarfangPushBackOpenXREyeFrameBufferList(HarfangOpenXREyeFrameBufferList this_, HarfangOpenXREyeFrameBuffer v);
extern size_t HarfangSizeOpenXREyeFrameBufferList(HarfangOpenXREyeFrameBufferList this_);
extern HarfangOpenXREyeFrameBuffer HarfangAtOpenXREyeFrameBufferList(HarfangOpenXREyeFrameBufferList this_, size_t idx);
extern HarfangIntList HarfangOpenXRFrameInfoGetIdFbs(HarfangOpenXRFrameInfo h);
extern void HarfangOpenXRFrameInfoSetIdFbs(HarfangOpenXRFrameInfo h, HarfangIntList v);
extern void HarfangOpenXRFrameInfoFree(HarfangOpenXRFrameInfo);
extern bool HarfangSRanipalEyeStateGetPupilDiameterValid(HarfangSRanipalEyeState h);
extern void HarfangSRanipalEyeStateSetPupilDiameterValid(HarfangSRanipalEyeState h, bool v);
extern HarfangVec3 HarfangSRanipalEyeStateGetGazeOriginMm(HarfangSRanipalEyeState h);
extern void HarfangSRanipalEyeStateSetGazeOriginMm(HarfangSRanipalEyeState h, HarfangVec3 v);
extern HarfangVec3 HarfangSRanipalEyeStateGetGazeDirectionNormalized(HarfangSRanipalEyeState h);
extern void HarfangSRanipalEyeStateSetGazeDirectionNormalized(HarfangSRanipalEyeState h, HarfangVec3 v);
extern float HarfangSRanipalEyeStateGetPupilDiameterMm(HarfangSRanipalEyeState h);
extern void HarfangSRanipalEyeStateSetPupilDiameterMm(HarfangSRanipalEyeState h, float v);
extern float HarfangSRanipalEyeStateGetEyeOpenness(HarfangSRanipalEyeState h);
extern void HarfangSRanipalEyeStateSetEyeOpenness(HarfangSRanipalEyeState h, float v);
extern void HarfangSRanipalEyeStateFree(HarfangSRanipalEyeState);
extern HarfangSRanipalEyeState HarfangSRanipalStateGetRightEye(HarfangSRanipalState h);
extern void HarfangSRanipalStateSetRightEye(HarfangSRanipalState h, HarfangSRanipalEyeState v);
extern HarfangSRanipalEyeState HarfangSRanipalStateGetLeftEye(HarfangSRanipalState h);
extern void HarfangSRanipalStateSetLeftEye(HarfangSRanipalState h, HarfangSRanipalEyeState v);
extern void HarfangSRanipalStateFree(HarfangSRanipalState);
extern HarfangVec3 HarfangVertexGetPos(HarfangVertex h);
extern void HarfangVertexSetPos(HarfangVertex h, HarfangVec3 v);
extern HarfangVec3 HarfangVertexGetNormal(HarfangVertex h);
extern void HarfangVertexSetNormal(HarfangVertex h, HarfangVec3 v);
extern HarfangVec3 HarfangVertexGetTangent(HarfangVertex h);
extern void HarfangVertexSetTangent(HarfangVertex h, HarfangVec3 v);
extern HarfangVec3 HarfangVertexGetBinormal(HarfangVertex h);
extern void HarfangVertexSetBinormal(HarfangVertex h, HarfangVec3 v);
extern HarfangVec2 HarfangVertexGetUv0(HarfangVertex h);
extern void HarfangVertexSetUv0(HarfangVertex h, HarfangVec2 v);
extern HarfangVec2 HarfangVertexGetUv1(HarfangVertex h);
extern void HarfangVertexSetUv1(HarfangVertex h, HarfangVec2 v);
extern HarfangVec2 HarfangVertexGetUv2(HarfangVertex h);
extern void HarfangVertexSetUv2(HarfangVertex h, HarfangVec2 v);
extern HarfangVec2 HarfangVertexGetUv3(HarfangVertex h);
extern void HarfangVertexSetUv3(HarfangVertex h, HarfangVec2 v);
extern HarfangVec2 HarfangVertexGetUv4(HarfangVertex h);
extern void HarfangVertexSetUv4(HarfangVertex h, HarfangVec2 v);
extern HarfangVec2 HarfangVertexGetUv5(HarfangVertex h);
extern void HarfangVertexSetUv5(HarfangVertex h, HarfangVec2 v);
extern HarfangVec2 HarfangVertexGetUv6(HarfangVertex h);
extern void HarfangVertexSetUv6(HarfangVertex h, HarfangVec2 v);
extern HarfangVec2 HarfangVertexGetUv7(HarfangVertex h);
extern void HarfangVertexSetUv7(HarfangVertex h, HarfangVec2 v);
extern HarfangColor HarfangVertexGetColor0(HarfangVertex h);
extern void HarfangVertexSetColor0(HarfangVertex h, HarfangColor v);
extern HarfangColor HarfangVertexGetColor1(HarfangVertex h);
extern void HarfangVertexSetColor1(HarfangVertex h, HarfangColor v);
extern HarfangColor HarfangVertexGetColor2(HarfangVertex h);
extern void HarfangVertexSetColor2(HarfangVertex h, HarfangColor v);
extern HarfangColor HarfangVertexGetColor3(HarfangVertex h);
extern void HarfangVertexSetColor3(HarfangVertex h, HarfangColor v);
extern HarfangVertex HarfangConstructorVertex();
extern void HarfangVertexFree(HarfangVertex);
extern HarfangModelBuilder HarfangConstructorModelBuilder();
extern void HarfangModelBuilderFree(HarfangModelBuilder);
extern uint32_t HarfangAddVertexModelBuilder(HarfangModelBuilder this_, const HarfangVertex vtx);
extern void HarfangAddTriangleModelBuilder(HarfangModelBuilder this_, uint32_t a, uint32_t b, uint32_t c);
extern void HarfangAddQuadModelBuilder(HarfangModelBuilder this_, uint32_t a, uint32_t b, uint32_t c, uint32_t d);
extern void HarfangAddPolygonModelBuilder(HarfangModelBuilder this_, const HarfangUint32TList idxs);
extern size_t HarfangGetCurrentListIndexCountModelBuilder(HarfangModelBuilder this_);
extern void HarfangEndListModelBuilder(HarfangModelBuilder this_, uint16_t material);
extern void HarfangClearModelBuilder(HarfangModelBuilder this_);
extern HarfangModel HarfangMakeModelModelBuilder(HarfangModelBuilder this_, const HarfangVertexLayout decl);
extern void HarfangGeometryFree(HarfangGeometry);
extern HarfangGeometryBuilder HarfangConstructorGeometryBuilder();
extern void HarfangGeometryBuilderFree(HarfangGeometryBuilder);
extern void HarfangAddVertexGeometryBuilder(HarfangGeometryBuilder this_, HarfangVertex vtx);
extern void HarfangAddPolygonGeometryBuilder(HarfangGeometryBuilder this_, const HarfangUint32TList idxs, uint16_t material);
extern void HarfangAddPolygonGeometryBuilderWithSliceOfIdxs(
	HarfangGeometryBuilder this_, size_t SliceOfidxsToCSize, uint32_t *SliceOfidxsToCBuf, uint16_t material);
extern void HarfangAddTriangleGeometryBuilder(HarfangGeometryBuilder this_, uint32_t a, uint32_t b, uint32_t c, uint32_t material);
extern void HarfangAddQuadGeometryBuilder(HarfangGeometryBuilder this_, uint32_t a, uint32_t b, uint32_t c, uint32_t d, uint32_t material);
extern HarfangGeometry HarfangMakeGeometryBuilder(HarfangGeometryBuilder this_);
extern void HarfangClearGeometryBuilder(HarfangGeometryBuilder this_);
extern void HarfangIsoSurfaceFree(HarfangIsoSurface);
extern void HarfangBloomFree(HarfangBloom);
extern void HarfangSAOFree(HarfangSAO);
extern void HarfangProfilerFrameFree(HarfangProfilerFrame);
extern void HarfangIVideoStreamerFree(HarfangIVideoStreamer);
extern int HarfangStartupIVideoStreamer(HarfangIVideoStreamer this_);
extern void HarfangShutdownIVideoStreamer(HarfangIVideoStreamer this_);
extern intptr_t HarfangOpenIVideoStreamer(HarfangIVideoStreamer this_, const char *name);
extern int HarfangPlayIVideoStreamer(HarfangIVideoStreamer this_, intptr_t h);
extern int HarfangPauseIVideoStreamer(HarfangIVideoStreamer this_, intptr_t h);
extern int HarfangCloseIVideoStreamer(HarfangIVideoStreamer this_, intptr_t h);
extern int HarfangSeekIVideoStreamer(HarfangIVideoStreamer this_, intptr_t h, int64_t t);
extern int64_t HarfangGetDurationIVideoStreamer(HarfangIVideoStreamer this_, intptr_t h);
extern int64_t HarfangGetTimeStampIVideoStreamer(HarfangIVideoStreamer this_, intptr_t h);
extern int HarfangIsEndedIVideoStreamer(HarfangIVideoStreamer this_, intptr_t h);
extern int HarfangGetFrameIVideoStreamer(HarfangIVideoStreamer this_, intptr_t h, intptr_t *ptr, int *width, int *height, int *pitch, int *format);
extern int HarfangFreeFrameIVideoStreamer(HarfangIVideoStreamer this_, intptr_t h, int frame);
extern HarfangVoidPointer HarfangIntToVoidPointer(intptr_t ptr);
extern void HarfangSetLogLevel(int log_level);
extern void HarfangSetLogDetailed(bool is_detailed);
extern void HarfangLog(const char *msg);
extern void HarfangLogWithDetails(const char *msg, const char *details);
extern void HarfangWarn(const char *msg);
extern void HarfangWarnWithDetails(const char *msg, const char *details);
extern void HarfangError(const char *msg);
extern void HarfangErrorWithDetails(const char *msg, const char *details);
extern void HarfangDebug(const char *msg);
extern void HarfangDebugWithDetails(const char *msg, const char *details);
extern float HarfangTimeToSecF(int64_t t);
extern float HarfangTimeToMsF(int64_t t);
extern float HarfangTimeToUsF(int64_t t);
extern int64_t HarfangTimeToDay(int64_t t);
extern int64_t HarfangTimeToHour(int64_t t);
extern int64_t HarfangTimeToMin(int64_t t);
extern int64_t HarfangTimeToSec(int64_t t);
extern int64_t HarfangTimeToMs(int64_t t);
extern int64_t HarfangTimeToUs(int64_t t);
extern int64_t HarfangTimeToNs(int64_t t);
extern int64_t HarfangTimeFromSecF(float sec);
extern int64_t HarfangTimeFromMsF(float ms);
extern int64_t HarfangTimeFromUsF(float us);
extern int64_t HarfangTimeFromDay(int64_t day);
extern int64_t HarfangTimeFromHour(int64_t hour);
extern int64_t HarfangTimeFromMin(int64_t min);
extern int64_t HarfangTimeFromSec(int64_t sec);
extern int64_t HarfangTimeFromMs(int64_t ms);
extern int64_t HarfangTimeFromUs(int64_t us);
extern int64_t HarfangTimeFromNs(int64_t ns);
extern int64_t HarfangTimeNow();
extern const char *HarfangTimeToString(int64_t t);
extern void HarfangResetClock();
extern int64_t HarfangTickClock();
extern int64_t HarfangGetClock();
extern int64_t HarfangGetClockDt();
extern void HarfangSkipClock();
extern HarfangFile HarfangOpen(const char *path);
extern HarfangFile HarfangOpenText(const char *path);
extern HarfangFile HarfangOpenWrite(const char *path);
extern HarfangFile HarfangOpenWriteText(const char *path);
extern HarfangFile HarfangOpenTemp(const char *template_path);
extern bool HarfangClose(HarfangFile file);
extern bool HarfangIsValid(HarfangFile file);
extern bool HarfangIsValidWithT(const HarfangTexture t);
extern bool HarfangIsValidWithFb(const HarfangFrameBuffer fb);
extern bool HarfangIsValidWithPipeline(const HarfangForwardPipelineAAA pipeline);
extern bool HarfangIsValidWithStreamer(HarfangIVideoStreamer streamer);
extern bool HarfangIsEOF(HarfangFile file);
extern size_t HarfangGetSize(HarfangFile file);
extern HarfangVec2 HarfangGetSizeWithRect(const HarfangRect rect);
extern HarfangIVec2 HarfangGetSizeWithIntRectRect(const HarfangIntRect rect);
extern bool HarfangSeek(HarfangFile file, int64_t offset, int mode);
extern size_t HarfangTell(HarfangFile file);
extern void HarfangRewind(HarfangFile file);
extern bool HarfangIsFile(const char *path);
extern bool HarfangUnlink(const char *path);
extern uint8_t HarfangReadUInt8(HarfangFile file);
extern uint16_t HarfangReadUInt16(HarfangFile file);
extern uint32_t HarfangReadUInt32(HarfangFile file);
extern float HarfangReadFloat(HarfangFile file);
extern bool HarfangWriteUInt8(HarfangFile file, uint8_t value);
extern bool HarfangWriteUInt16(HarfangFile file, uint16_t value);
extern bool HarfangWriteUInt32(HarfangFile file, uint32_t value);
extern bool HarfangWriteFloat(HarfangFile file, float value);
extern const char *HarfangReadString(HarfangFile file);
extern bool HarfangWriteString(HarfangFile file, const char *value);
extern bool HarfangCopyFile(const char *src, const char *dst);
extern const char *HarfangFileToString(const char *path);
extern bool HarfangStringToFile(const char *path, const char *value);
extern bool HarfangLoadDataFromFile(const char *path, HarfangData data);
extern bool HarfangSaveDataToFile(const char *path, const HarfangData data);
extern HarfangDirEntryList HarfangListDir(const char *path, int type);
extern HarfangDirEntryList HarfangListDirRecursive(const char *path, int type);
extern bool HarfangMkDir(const char *path);
extern bool HarfangMkDirWithPermissions(const char *path, int permissions);
extern bool HarfangRmDir(const char *path);
extern bool HarfangMkTree(const char *path);
extern bool HarfangMkTreeWithPermissions(const char *path, int permissions);
extern bool HarfangRmTree(const char *path);
extern bool HarfangExists(const char *path);
extern bool HarfangIsDir(const char *path);
extern bool HarfangCopyDir(const char *src, const char *dst);
extern bool HarfangCopyDirRecursive(const char *src, const char *dst);
extern bool HarfangIsPathAbsolute(const char *path);
extern const char *HarfangPathToDisplay(const char *path);
extern const char *HarfangNormalizePath(const char *path);
extern const char *HarfangFactorizePath(const char *path);
extern const char *HarfangCleanPath(const char *path);
extern const char *HarfangCutFilePath(const char *path);
extern const char *HarfangCutFileName(const char *path);
extern const char *HarfangCutFileExtension(const char *path);
extern const char *HarfangGetFilePath(const char *path);
extern const char *HarfangGetFileName(const char *path);
extern const char *HarfangGetFileExtension(const char *path);
extern bool HarfangHasFileExtension(const char *path);
extern bool HarfangPathStartsWith(const char *path, const char *with);
extern const char *HarfangPathStripPrefix(const char *path, const char *prefix);
extern const char *HarfangPathStripSuffix(const char *path, const char *suffix);
extern const char *HarfangPathJoin(const HarfangStringList elements);
extern const char *HarfangSwapFileExtension(const char *path, const char *ext);
extern const char *HarfangGetCurrentWorkingDirectory();
extern const char *HarfangGetUserFolder();
extern bool HarfangAddAssetsFolder(const char *path);
extern void HarfangRemoveAssetsFolder(const char *path);
extern bool HarfangAddAssetsPackage(const char *path);
extern void HarfangRemoveAssetsPackage(const char *path);
extern bool HarfangIsAssetFile(const char *name);
extern float HarfangLinearInterpolate(float y0, float y1, float t);
extern float HarfangCosineInterpolate(float y0, float y1, float t);
extern float HarfangCubicInterpolate(float y0, float y1, float y2, float y3, float t);
extern HarfangVec3 HarfangCubicInterpolateWithV0V1V2V3(const HarfangVec3 v0, const HarfangVec3 v1, const HarfangVec3 v2, const HarfangVec3 v3, float t);
extern float HarfangHermiteInterpolate(float y0, float y1, float y2, float y3, float t, float tension, float bias);
extern uint8_t HarfangReverseRotationOrder(uint8_t rotation_order);
extern float HarfangGetArea(const HarfangMinMax minmax);
extern HarfangVec3 HarfangGetCenter(const HarfangMinMax minmax);
extern void HarfangComputeMinMaxBoundingSphere(const HarfangMinMax minmax, HarfangVec3 origin, float *radius);
extern bool HarfangOverlap(const HarfangMinMax minmax_a, const HarfangMinMax minmax_b);
extern bool HarfangOverlapWithAxis(const HarfangMinMax minmax_a, const HarfangMinMax minmax_b, uint8_t axis);
extern bool HarfangContains(const HarfangMinMax minmax, const HarfangVec3 position);
extern HarfangMinMax HarfangUnion(const HarfangMinMax minmax_a, const HarfangMinMax minmax_b);
extern HarfangMinMax HarfangUnionWithMinmaxPosition(const HarfangMinMax minmax, const HarfangVec3 position);
extern bool HarfangIntersectRay(const HarfangMinMax minmax, const HarfangVec3 origin, const HarfangVec3 direction, float *t_min, float *t_max);
extern bool HarfangClassifyLine(
	const HarfangMinMax minmax, const HarfangVec3 position, const HarfangVec3 direction, HarfangVec3 intersection, HarfangVec3 normal);
extern bool HarfangClassifySegment(const HarfangMinMax minmax, const HarfangVec3 p0, const HarfangVec3 p1, HarfangVec3 intersection, HarfangVec3 normal);
extern HarfangMinMax HarfangMinMaxFromPositionSize(const HarfangVec3 position, const HarfangVec3 size);
extern HarfangVec2 HarfangMin(const HarfangVec2 a, const HarfangVec2 b);
extern HarfangIVec2 HarfangMinWithAB(const HarfangIVec2 a, const HarfangIVec2 b);
extern HarfangVec3 HarfangMinWithVec3AVec3B(const HarfangVec3 a, const HarfangVec3 b);
extern float HarfangMinWithFloatAFloatB(float a, float b);
extern int HarfangMinWithIntAIntB(int a, int b);
extern HarfangVec2 HarfangMax(const HarfangVec2 a, const HarfangVec2 b);
extern HarfangIVec2 HarfangMaxWithAB(const HarfangIVec2 a, const HarfangIVec2 b);
extern HarfangVec3 HarfangMaxWithVec3AVec3B(const HarfangVec3 a, const HarfangVec3 b);
extern float HarfangMaxWithFloatAFloatB(float a, float b);
extern int HarfangMaxWithIntAIntB(int a, int b);
extern float HarfangLen2(const HarfangVec2 v);
extern int HarfangLen2WithV(const HarfangIVec2 v);
extern float HarfangLen2WithQ(const HarfangQuaternion q);
extern float HarfangLen2WithVec3V(const HarfangVec3 v);
extern float HarfangLen(const HarfangVec2 v);
extern int HarfangLenWithV(const HarfangIVec2 v);
extern float HarfangLenWithQ(const HarfangQuaternion q);
extern float HarfangLenWithVec3V(const HarfangVec3 v);
extern float HarfangDot(const HarfangVec2 a, const HarfangVec2 b);
extern int HarfangDotWithAB(const HarfangIVec2 a, const HarfangIVec2 b);
extern float HarfangDotWithVec3AVec3B(const HarfangVec3 a, const HarfangVec3 b);
extern HarfangVec2 HarfangNormalize(const HarfangVec2 v);
extern HarfangIVec2 HarfangNormalizeWithV(const HarfangIVec2 v);
extern HarfangVec4 HarfangNormalizeWithVec4V(const HarfangVec4 v);
extern HarfangQuaternion HarfangNormalizeWithQ(const HarfangQuaternion q);
extern HarfangMat3 HarfangNormalizeWithM(const HarfangMat3 m);
extern HarfangVec3 HarfangNormalizeWithVec3V(const HarfangVec3 v);
extern HarfangVec2 HarfangReverse(const HarfangVec2 a);
extern HarfangIVec2 HarfangReverseWithA(const HarfangIVec2 a);
extern HarfangVec3 HarfangReverseWithV(const HarfangVec3 v);
extern float HarfangDist2(const HarfangVec2 a, const HarfangVec2 b);
extern int HarfangDist2WithAB(const HarfangIVec2 a, const HarfangIVec2 b);
extern float HarfangDist2WithVec3AVec3B(const HarfangVec3 a, const HarfangVec3 b);
extern float HarfangDist(const HarfangVec2 a, const HarfangVec2 b);
extern int HarfangDistWithAB(const HarfangIVec2 a, const HarfangIVec2 b);
extern float HarfangDistWithQuaternionAQuaternionB(const HarfangQuaternion a, const HarfangQuaternion b);
extern float HarfangDistWithVec3AVec3B(const HarfangVec3 a, const HarfangVec3 b);
extern HarfangVec4 HarfangAbs(const HarfangVec4 v);
extern HarfangVec3 HarfangAbsWithV(const HarfangVec3 v);
extern float HarfangAbsWithFloatV(float v);
extern int HarfangAbsWithIntV(int v);
extern HarfangVec4 HarfangRandomVec4(float min, float max);
extern HarfangVec4 HarfangRandomVec4WithMinMax(const HarfangVec4 min, const HarfangVec4 max);
extern HarfangQuaternion HarfangInverse(const HarfangQuaternion q);
extern bool HarfangInverseWithMI(const HarfangMat3 m, HarfangMat3 I);
extern bool HarfangInverseWithMat4MMat4I(const HarfangMat4 m, HarfangMat4 I);
extern HarfangMat44 HarfangInverseWithMResult(const HarfangMat44 m, bool *result);
extern HarfangVec3 HarfangInverseWithV(const HarfangVec3 v);
extern HarfangQuaternion HarfangSlerp(const HarfangQuaternion a, const HarfangQuaternion b, float t);
extern HarfangQuaternion HarfangQuaternionFromEulerWithXYZ(float x, float y, float z);
extern HarfangQuaternion HarfangQuaternionFromEulerWithXYZRotationOrder(float x, float y, float z, uint8_t rotation_order);
extern HarfangQuaternion HarfangQuaternionFromEuler(HarfangVec3 euler);
extern HarfangQuaternion HarfangQuaternionFromEulerWithRotationOrder(HarfangVec3 euler, uint8_t rotation_order);
extern HarfangQuaternion HarfangQuaternionLookAt(const HarfangVec3 at);
extern HarfangQuaternion HarfangQuaternionFromMatrix3(const HarfangMat3 m);
extern HarfangQuaternion HarfangQuaternionFromAxisAngle(float angle, const HarfangVec3 axis);
extern HarfangMat3 HarfangToMatrix3(const HarfangQuaternion q);
extern HarfangVec3 HarfangToEuler(const HarfangQuaternion q);
extern HarfangVec3 HarfangToEulerWithRotationOrder(const HarfangQuaternion q, uint8_t rotation_order);
extern HarfangVec3 HarfangToEulerWithM(const HarfangMat3 m);
extern HarfangVec3 HarfangToEulerWithMRotationOrder(const HarfangMat3 m, uint8_t rotation_order);
extern float HarfangDet(const HarfangMat3 m);
extern HarfangMat3 HarfangTranspose(const HarfangMat3 m);
extern HarfangVec3 HarfangGetRow(const HarfangMat3 m, uint32_t n);
extern HarfangVec4 HarfangGetRowWithMN(const HarfangMat4 m, unsigned int n);
extern HarfangVec4 HarfangGetRowWithMIdx(const HarfangMat44 m, uint32_t idx);
extern HarfangVec3 HarfangGetColumn(const HarfangMat3 m, uint32_t n);
extern HarfangVec3 HarfangGetColumnWithMN(const HarfangMat4 m, unsigned int n);
extern HarfangVec4 HarfangGetColumnWithMIdx(const HarfangMat44 m, uint32_t idx);
extern void HarfangSetRow(HarfangMat3 m, uint32_t n, const HarfangVec3 row);
extern void HarfangSetRowWithMNV(const HarfangMat4 m, unsigned int n, const HarfangVec4 v);
extern void HarfangSetRowWithMIdxV(const HarfangMat44 m, uint32_t idx, const HarfangVec4 v);
extern void HarfangSetColumn(HarfangMat3 m, uint32_t n, const HarfangVec3 column);
extern void HarfangSetColumnWithMNV(const HarfangMat4 m, unsigned int n, const HarfangVec3 v);
extern void HarfangSetColumnWithMIdxV(const HarfangMat44 m, uint32_t idx, const HarfangVec4 v);
extern HarfangVec3 HarfangGetX(const HarfangMat3 m);
extern HarfangVec3 HarfangGetXWithM(const HarfangMat4 m);
extern float HarfangGetXWithRect(const HarfangRect rect);
extern int HarfangGetXWithIntRectRect(const HarfangIntRect rect);
extern HarfangVec3 HarfangGetY(const HarfangMat3 m);
extern HarfangVec3 HarfangGetYWithM(const HarfangMat4 m);
extern float HarfangGetYWithRect(const HarfangRect rect);
extern int HarfangGetYWithIntRectRect(const HarfangIntRect rect);
extern HarfangVec3 HarfangGetZ(const HarfangMat3 m);
extern HarfangVec3 HarfangGetZWithM(const HarfangMat4 m);
extern HarfangVec3 HarfangGetTranslation(const HarfangMat3 m);
extern HarfangVec3 HarfangGetTranslationWithM(const HarfangMat4 m);
extern HarfangVec3 HarfangGetScale(const HarfangMat3 m);
extern HarfangVec3 HarfangGetScaleWithM(const HarfangMat4 m);
extern void HarfangSetX(HarfangMat3 m, const HarfangVec3 X);
extern void HarfangSetXWithM(const HarfangMat4 m, const HarfangVec3 X);
extern void HarfangSetXWithRectX(HarfangRect rect, float x);
extern void HarfangSetXWithIntRectRectIntX(HarfangIntRect rect, int x);
extern void HarfangSetY(HarfangMat3 m, const HarfangVec3 Y);
extern void HarfangSetYWithM(const HarfangMat4 m, const HarfangVec3 Y);
extern void HarfangSetYWithRectY(HarfangRect rect, float y);
extern void HarfangSetYWithIntRectRectIntY(HarfangIntRect rect, int y);
extern void HarfangSetZ(HarfangMat3 m, const HarfangVec3 Z);
extern void HarfangSetZWithM(const HarfangMat4 m, const HarfangVec3 Z);
extern void HarfangSetTranslation(HarfangMat3 m, const HarfangVec3 T);
extern void HarfangSetTranslationWithT(HarfangMat3 m, const HarfangVec2 T);
extern void HarfangSetTranslationWithM(const HarfangMat4 m, const HarfangVec3 T);
extern void HarfangSetScale(HarfangMat3 m, const HarfangVec3 S);
extern void HarfangSetScaleWithMScale(const HarfangMat4 m, const HarfangVec3 scale);
extern void HarfangSetAxises(HarfangMat3 m, const HarfangVec3 X, const HarfangVec3 Y, const HarfangVec3 Z);
extern HarfangMat3 HarfangOrthonormalize(const HarfangMat3 m);
extern HarfangMat4 HarfangOrthonormalizeWithM(const HarfangMat4 m);
extern HarfangMat3 HarfangVectorMat3(const HarfangVec3 V);
extern HarfangMat3 HarfangTranslationMat3(const HarfangVec2 T);
extern HarfangMat3 HarfangTranslationMat3WithT(const HarfangVec3 T);
extern HarfangMat3 HarfangScaleMat3(const HarfangVec2 S);
extern HarfangMat3 HarfangScaleMat3WithS(const HarfangVec3 S);
extern HarfangMat3 HarfangCrossProductMat3(const HarfangVec3 V);
extern HarfangMat3 HarfangRotationMatX(float angle);
extern HarfangMat3 HarfangRotationMatY(float angle);
extern HarfangMat3 HarfangRotationMatZ(float angle);
extern HarfangMat3 HarfangRotationMat2D(float angle, const HarfangVec2 pivot);
extern HarfangMat3 HarfangRotationMat3WithXYZ(float x, float y, float z);
extern HarfangMat3 HarfangRotationMat3WithXYZRotationOrder(float x, float y, float z, uint8_t rotation_order);
extern HarfangMat3 HarfangRotationMat3(const HarfangVec3 euler);
extern HarfangMat3 HarfangRotationMat3WithRotationOrder(const HarfangVec3 euler, uint8_t rotation_order);
extern HarfangMat3 HarfangMat3LookAt(const HarfangVec3 front);
extern HarfangMat3 HarfangMat3LookAtWithUp(const HarfangVec3 front, const HarfangVec3 up);
extern HarfangMat3 HarfangRotationMatXZY(float x, float y, float z);
extern HarfangMat3 HarfangRotationMatZYX(float x, float y, float z);
extern HarfangMat3 HarfangRotationMatXYZ(float x, float y, float z);
extern HarfangMat3 HarfangRotationMatZXY(float x, float y, float z);
extern HarfangMat3 HarfangRotationMatYZX(float x, float y, float z);
extern HarfangMat3 HarfangRotationMatYXZ(float x, float y, float z);
extern HarfangMat3 HarfangRotationMatXY(float x, float y);
extern HarfangVec3 HarfangGetT(const HarfangMat4 m);
extern HarfangVec3 HarfangGetR(const HarfangMat4 m);
extern HarfangVec3 HarfangGetRWithRotationOrder(const HarfangMat4 m, uint8_t rotation_order);
extern HarfangVec3 HarfangGetRotation(const HarfangMat4 m);
extern HarfangVec3 HarfangGetRotationWithRotationOrder(const HarfangMat4 m, uint8_t rotation_order);
extern HarfangMat3 HarfangGetRMatrix(const HarfangMat4 m);
extern HarfangMat3 HarfangGetRotationMatrix(const HarfangMat4 m);
extern HarfangVec3 HarfangGetS(const HarfangMat4 m);
extern void HarfangSetT(const HarfangMat4 m, const HarfangVec3 T);
extern void HarfangSetS(const HarfangMat4 m, const HarfangVec3 scale);
extern HarfangMat4 HarfangInverseFast(const HarfangMat4 m);
extern HarfangMat4 HarfangLerpAsOrthonormalBase(const HarfangMat4 from, const HarfangMat4 to, float k);
extern HarfangMat4 HarfangLerpAsOrthonormalBaseWithFast(const HarfangMat4 from, const HarfangMat4 to, float k, bool fast);
extern void HarfangDecompose(const HarfangMat4 m, HarfangVec3 position, HarfangVec3 rotation, HarfangVec3 scale);
extern void HarfangDecomposeWithRotationOrder(const HarfangMat4 m, HarfangVec3 position, HarfangVec3 rotation, HarfangVec3 scale, uint8_t rotation_order);
extern HarfangMat4 HarfangMat4LookAt(const HarfangVec3 position, const HarfangVec3 at);
extern HarfangMat4 HarfangMat4LookAtWithScale(const HarfangVec3 position, const HarfangVec3 at, const HarfangVec3 scale);
extern HarfangMat4 HarfangMat4LookAtUp(const HarfangVec3 position, const HarfangVec3 at, const HarfangVec3 up);
extern HarfangMat4 HarfangMat4LookAtUpWithScale(const HarfangVec3 position, const HarfangVec3 at, const HarfangVec3 up, const HarfangVec3 scale);
extern HarfangMat4 HarfangMat4LookToward(const HarfangVec3 position, const HarfangVec3 direction);
extern HarfangMat4 HarfangMat4LookTowardWithScale(const HarfangVec3 position, const HarfangVec3 direction, const HarfangVec3 scale);
extern HarfangMat4 HarfangMat4LookTowardUp(const HarfangVec3 position, const HarfangVec3 direction, const HarfangVec3 up);
extern HarfangMat4 HarfangMat4LookTowardUpWithScale(const HarfangVec3 position, const HarfangVec3 direction, const HarfangVec3 up, const HarfangVec3 scale);
extern HarfangMat4 HarfangTranslationMat4(const HarfangVec3 t);
extern HarfangMat4 HarfangRotationMat4(const HarfangVec3 euler);
extern HarfangMat4 HarfangRotationMat4WithOrder(const HarfangVec3 euler, uint8_t order);
extern HarfangMat4 HarfangScaleMat4(const HarfangVec3 scale);
extern HarfangMat4 HarfangScaleMat4WithScale(float scale);
extern HarfangMat4 HarfangTransformationMat4(const HarfangVec3 pos, const HarfangVec3 rot);
extern HarfangMat4 HarfangTransformationMat4WithScale(const HarfangVec3 pos, const HarfangVec3 rot, const HarfangVec3 scale);
extern HarfangMat4 HarfangTransformationMat4WithRot(const HarfangVec3 pos, const HarfangMat3 rot);
extern HarfangMat4 HarfangTransformationMat4WithRotScale(const HarfangVec3 pos, const HarfangMat3 rot, const HarfangVec3 scale);
extern HarfangVec3 HarfangMakeVec3(const HarfangVec4 v);
extern HarfangVec3 HarfangRandomVec3(float min, float max);
extern HarfangVec3 HarfangRandomVec3WithMinMax(const HarfangVec3 min, const HarfangVec3 max);
extern HarfangVec3 HarfangBaseToEuler(const HarfangVec3 z);
extern HarfangVec3 HarfangBaseToEulerWithY(const HarfangVec3 z, const HarfangVec3 y);
extern HarfangVec3 HarfangCross(const HarfangVec3 a, const HarfangVec3 b);
extern HarfangVec3 HarfangClamp(const HarfangVec3 v, float min, float max);
extern HarfangVec3 HarfangClampWithMinMax(const HarfangVec3 v, const HarfangVec3 min, const HarfangVec3 max);
extern float HarfangClampWithV(float v, float min, float max);
extern int HarfangClampWithVMinMax(int v, int min, int max);
extern HarfangColor HarfangClampWithColor(const HarfangColor color, float min, float max);
extern HarfangColor HarfangClampWithColorMinMax(const HarfangColor color, const HarfangColor min, const HarfangColor max);
extern HarfangVec3 HarfangClampLen(const HarfangVec3 v, float min, float max);
extern HarfangVec3 HarfangSign(const HarfangVec3 v);
extern HarfangVec3 HarfangReflect(const HarfangVec3 v, const HarfangVec3 n);
extern HarfangVec3 HarfangRefract(const HarfangVec3 v, const HarfangVec3 n);
extern HarfangVec3 HarfangRefractWithKIn(const HarfangVec3 v, const HarfangVec3 n, float k_in);
extern HarfangVec3 HarfangRefractWithKInKOut(const HarfangVec3 v, const HarfangVec3 n, float k_in, float k_out);
extern HarfangVec3 HarfangFloor(const HarfangVec3 v);
extern HarfangVec3 HarfangCeil(const HarfangVec3 v);
extern HarfangVec3 HarfangFaceForward(const HarfangVec3 v, const HarfangVec3 d);
extern HarfangVec3 HarfangDeg3(float x, float y, float z);
extern HarfangVec3 HarfangRad3(float x, float y, float z);
extern HarfangVec3 HarfangVec3I(int x, int y, int z);
extern HarfangVec4 HarfangVec4I(int x, int y, int z);
extern HarfangVec4 HarfangVec4IWithW(int x, int y, int z, int w);
extern float HarfangGetWidth(const HarfangRect rect);
extern int HarfangGetWidthWithRect(const HarfangIntRect rect);
extern float HarfangGetHeight(const HarfangRect rect);
extern int HarfangGetHeightWithRect(const HarfangIntRect rect);
extern void HarfangSetWidth(HarfangRect rect, float width);
extern void HarfangSetWidthWithRectWidth(HarfangIntRect rect, int width);
extern void HarfangSetHeight(HarfangRect rect, float height);
extern void HarfangSetHeightWithRectHeight(HarfangIntRect rect, int height);
extern bool HarfangInside(const HarfangRect rect, HarfangIVec2 v);
extern bool HarfangInsideWithV(const HarfangRect rect, HarfangVec2 v);
extern bool HarfangInsideWithVec3V(const HarfangRect rect, HarfangVec3 v);
extern bool HarfangInsideWithVec4V(const HarfangRect rect, HarfangVec4 v);
extern bool HarfangInsideWithRect(const HarfangIntRect rect, HarfangIVec2 v);
extern bool HarfangInsideWithRectV(const HarfangIntRect rect, HarfangVec2 v);
extern bool HarfangInsideWithIntRectRectVec3V(const HarfangIntRect rect, HarfangVec3 v);
extern bool HarfangInsideWithIntRectRectVec4V(const HarfangIntRect rect, HarfangVec4 v);
extern bool HarfangFitsInside(const HarfangRect a, const HarfangRect b);
extern bool HarfangFitsInsideWithAB(const HarfangIntRect a, const HarfangIntRect b);
extern bool HarfangIntersects(const HarfangRect a, const HarfangRect b);
extern bool HarfangIntersectsWithAB(const HarfangIntRect a, const HarfangIntRect b);
extern HarfangRect HarfangIntersection(const HarfangRect a, const HarfangRect b);
extern HarfangIntRect HarfangIntersectionWithAB(const HarfangIntRect a, const HarfangIntRect b);
extern HarfangRect HarfangGrow(const HarfangRect rect, float border);
extern HarfangIntRect HarfangGrowWithRectBorder(const HarfangIntRect rect, int border);
extern HarfangRect HarfangOffset(const HarfangRect rect, float x, float y);
extern HarfangIntRect HarfangOffsetWithRectXY(const HarfangIntRect rect, int x, int y);
extern HarfangRect HarfangCrop(const HarfangRect rect, float left, float top, float right, float bottom);
extern HarfangIntRect HarfangCropWithRectLeftTopRightBottom(const HarfangIntRect rect, int left, int top, int right, int bottom);
extern HarfangRect HarfangMakeRectFromWidthHeight(float x, float y, float w, float h);
extern HarfangIntRect HarfangMakeRectFromWidthHeightWithXYWH(int x, int y, int w, int h);
extern HarfangRect HarfangToFloatRect(const HarfangIntRect rect);
extern HarfangIntRect HarfangToIntRect(const HarfangRect rect);
extern HarfangVec4 HarfangMakePlane(const HarfangVec3 p, const HarfangVec3 n);
extern HarfangVec4 HarfangMakePlaneWithM(const HarfangVec3 p, const HarfangVec3 n, const HarfangMat4 m);
extern float HarfangDistanceToPlane(const HarfangVec4 plane, const HarfangVec3 p);
extern float HarfangWrap(float v, float start, float end);
extern int HarfangWrapWithVStartEnd(int v, int start, int end);
extern int HarfangLerp(int a, int b, float t);
extern float HarfangLerpWithAB(float a, float b, float t);
extern HarfangVec3 HarfangLerpWithVec3AVec3B(const HarfangVec3 a, const HarfangVec3 b, float t);
extern HarfangVec4 HarfangLerpWithVec4AVec4B(const HarfangVec4 a, const HarfangVec4 b, float t);
extern float HarfangQuantize(float v, float q);
extern bool HarfangIsFinite(float v);
extern float HarfangDeg(float degrees);
extern float HarfangRad(float radians);
extern float HarfangDegreeToRadian(float degrees);
extern float HarfangRadianToDegree(float radians);
extern float HarfangSec(float seconds);
extern float HarfangMs(float milliseconds);
extern float HarfangKm(float km);
extern float HarfangMtr(float m);
extern float HarfangCm(float cm);
extern float HarfangMm(float mm);
extern float HarfangInch(float inch);
extern void HarfangSeed(uint32_t seed);
extern uint32_t HarfangRand();
extern uint32_t HarfangRandWithRange(uint32_t range);
extern float HarfangFRand();
extern float HarfangFRandWithRange(float range);
extern float HarfangFRRand();
extern float HarfangFRRandWithRangeStart(float range_start);
extern float HarfangFRRandWithRangeStartRangeEnd(float range_start, float range_end);
extern float HarfangZoomFactorToFov(float zoom_factor);
extern float HarfangFovToZoomFactor(float fov);
extern HarfangMat44 HarfangComputeOrthographicProjectionMatrix(float znear, float zfar, float size, const HarfangVec2 aspect_ratio);
extern HarfangMat44 HarfangComputeOrthographicProjectionMatrixWithOffset(
	float znear, float zfar, float size, const HarfangVec2 aspect_ratio, const HarfangVec2 offset);
extern HarfangMat44 HarfangComputePerspectiveProjectionMatrix(float znear, float zfar, float zoom_factor, const HarfangVec2 aspect_ratio);
extern HarfangMat44 HarfangComputePerspectiveProjectionMatrixWithOffset(
	float znear, float zfar, float zoom_factor, const HarfangVec2 aspect_ratio, const HarfangVec2 offset);
extern HarfangVec2 HarfangComputeAspectRatioX(float width, float height);
extern HarfangVec2 HarfangComputeAspectRatioY(float width, float height);
extern HarfangMat44 HarfangCompute2DProjectionMatrix(float znear, float zfar, float res_x, float res_y, bool y_up);
extern float HarfangExtractZoomFactorFromProjectionMatrix(const HarfangMat44 m, const HarfangVec2 aspect_ratio);
extern void HarfangExtractZRangeFromPerspectiveProjectionMatrix(const HarfangMat44 m, float *znear, float *zfar);
extern void HarfangExtractZRangeFromOrthographicProjectionMatrix(const HarfangMat44 m, float *znear, float *zfar);
extern void HarfangExtractZRangeFromProjectionMatrix(const HarfangMat44 m, float *znear, float *zfar);
extern bool HarfangProjectToClipSpace(const HarfangMat44 proj, const HarfangVec3 view, HarfangVec3 clip);
extern bool HarfangProjectOrthoToClipSpace(const HarfangMat44 proj, const HarfangVec3 view, HarfangVec3 clip);
extern bool HarfangUnprojectFromClipSpace(const HarfangMat44 inv_proj, const HarfangVec3 clip, HarfangVec3 view);
extern bool HarfangUnprojectOrthoFromClipSpace(const HarfangMat44 inv_proj, const HarfangVec3 clip, HarfangVec3 view);
extern HarfangVec3 HarfangClipSpaceToScreenSpace(const HarfangVec3 clip, const HarfangVec2 resolution);
extern HarfangVec3 HarfangScreenSpaceToClipSpace(const HarfangVec3 screen, const HarfangVec2 resolution);
extern bool HarfangProjectToScreenSpace(const HarfangMat44 proj, const HarfangVec3 view, const HarfangVec2 resolution, HarfangVec3 screen);
extern bool HarfangProjectOrthoToScreenSpace(const HarfangMat44 proj, const HarfangVec3 view, const HarfangVec2 resolution, HarfangVec3 screen);
extern bool HarfangUnprojectFromScreenSpace(const HarfangMat44 inv_proj, const HarfangVec3 screen, const HarfangVec2 resolution, HarfangVec3 view);
extern bool HarfangUnprojectOrthoFromScreenSpace(const HarfangMat44 inv_proj, const HarfangVec3 screen, const HarfangVec2 resolution, HarfangVec3 view);
extern float HarfangProjectZToClipSpace(float z, const HarfangMat44 proj);
extern HarfangFrustum HarfangMakeFrustum(const HarfangMat44 projection);
extern HarfangFrustum HarfangMakeFrustumWithMtx(const HarfangMat44 projection, const HarfangMat4 mtx);
extern HarfangFrustum HarfangTransformFrustum(const HarfangFrustum frustum, const HarfangMat4 mtx);
extern uint8_t HarfangTestVisibilityWithCountPoints(const HarfangFrustum frustum, uint32_t count, const HarfangVec3 points);
extern uint8_t HarfangTestVisibilityWithCountPointsDistance(const HarfangFrustum frustum, uint32_t count, const HarfangVec3 points, float distance);
extern uint8_t HarfangTestVisibilityWithOriginRadius(const HarfangFrustum frustum, const HarfangVec3 origin, float radius);
extern uint8_t HarfangTestVisibility(const HarfangFrustum frustum, const HarfangMinMax minmax);
extern void HarfangWindowSystemInit();
extern void HarfangWindowSystemShutdown();
extern HarfangMonitorList HarfangGetMonitors();
extern HarfangIntRect HarfangGetMonitorRect(const HarfangMonitor monitor);
extern bool HarfangIsPrimaryMonitor(const HarfangMonitor monitor);
extern bool HarfangIsMonitorConnected(const HarfangMonitor monitor);
extern const char *HarfangGetMonitorName(const HarfangMonitor monitor);
extern HarfangIVec2 HarfangGetMonitorSizeMM(const HarfangMonitor monitor);
extern bool HarfangGetMonitorModes(const HarfangMonitor monitor, HarfangMonitorModeList modes);
extern HarfangWindow HarfangNewWindow(int width, int height);
extern HarfangWindow HarfangNewWindowWithBpp(int width, int height, int bpp);
extern HarfangWindow HarfangNewWindowWithBppVisibility(int width, int height, int bpp, int visibility);
extern HarfangWindow HarfangNewWindowWithTitleWidthHeight(const char *title, int width, int height);
extern HarfangWindow HarfangNewWindowWithTitleWidthHeightBpp(const char *title, int width, int height, int bpp);
extern HarfangWindow HarfangNewWindowWithTitleWidthHeightBppVisibility(const char *title, int width, int height, int bpp, int visibility);
extern HarfangWindow HarfangNewFullscreenWindow(const HarfangMonitor monitor, int mode_index);
extern HarfangWindow HarfangNewFullscreenWindowWithRotation(const HarfangMonitor monitor, int mode_index, uint8_t rotation);
extern HarfangWindow HarfangNewFullscreenWindowWithTitleMonitorModeIndex(const char *title, const HarfangMonitor monitor, int mode_index);
extern HarfangWindow HarfangNewFullscreenWindowWithTitleMonitorModeIndexRotation(
	const char *title, const HarfangMonitor monitor, int mode_index, uint8_t rotation);
extern HarfangWindow HarfangNewWindowFrom(HarfangVoidPointer handle);
extern HarfangVoidPointer HarfangGetWindowHandle(const HarfangWindow window);
extern bool HarfangUpdateWindow(const HarfangWindow window);
extern bool HarfangDestroyWindow(const HarfangWindow window);
extern bool HarfangGetWindowClientSize(const HarfangWindow window, int *width, int *height);
extern bool HarfangSetWindowClientSize(HarfangWindow window, int width, int height);
extern HarfangVec2 HarfangGetWindowContentScale(const HarfangWindow window);
extern bool HarfangGetWindowTitle(const HarfangWindow window, const char **title);
extern bool HarfangSetWindowTitle(HarfangWindow window, const char *title);
extern bool HarfangWindowHasFocus(const HarfangWindow window);
extern HarfangWindow HarfangGetWindowInFocus();
extern HarfangIVec2 HarfangGetWindowPos(const HarfangWindow window);
extern bool HarfangSetWindowPos(HarfangWindow window, const HarfangIVec2 position);
extern bool HarfangIsWindowOpen(const HarfangWindow window);
extern void HarfangShowCursor();
extern void HarfangHideCursor();
extern void HarfangDisableCursor();
extern float HarfangColorToGrayscale(const HarfangColor color);
extern uint32_t HarfangColorToRGBA32(const HarfangColor color);
extern HarfangColor HarfangColorFromRGBA32(uint32_t rgba32);
extern uint32_t HarfangColorToABGR32(const HarfangColor color);
extern HarfangColor HarfangColorFromABGR32(uint32_t rgba32);
extern uint32_t HarfangARGB32ToRGBA32(uint32_t argb);
extern uint32_t HarfangRGBA32(uint8_t r, uint8_t g, uint8_t b);
extern uint32_t HarfangRGBA32WithA(uint8_t r, uint8_t g, uint8_t b, uint8_t a);
extern uint32_t HarfangARGB32(uint8_t r, uint8_t g, uint8_t b);
extern uint32_t HarfangARGB32WithA(uint8_t r, uint8_t g, uint8_t b, uint8_t a);
extern HarfangColor HarfangChromaScale(const HarfangColor color, float k);
extern HarfangColor HarfangAlphaScale(const HarfangColor color, float k);
extern HarfangColor HarfangColorFromVector3(const HarfangVec3 v);
extern HarfangColor HarfangColorFromVector4(const HarfangVec4 v);
extern HarfangColor HarfangColorI(int r, int g, int b);
extern HarfangColor HarfangColorIWithA(int r, int g, int b, int a);
extern HarfangColor HarfangToHLS(const HarfangColor color);
extern HarfangColor HarfangFromHLS(const HarfangColor color);
extern HarfangColor HarfangSetSaturation(const HarfangColor color, float saturation);
extern bool HarfangLoadJPG(HarfangPicture pict, const char *path);
extern bool HarfangLoadPNG(HarfangPicture pict, const char *path);
extern bool HarfangLoadGIF(HarfangPicture pict, const char *path);
extern bool HarfangLoadPSD(HarfangPicture pict, const char *path);
extern bool HarfangLoadTGA(HarfangPicture pict, const char *path);
extern bool HarfangLoadBMP(HarfangPicture pict, const char *path);
extern bool HarfangLoadPicture(HarfangPicture pict, const char *path);
extern bool HarfangSavePNG(HarfangPicture pict, const char *path);
extern bool HarfangSaveTGA(HarfangPicture pict, const char *path);
extern bool HarfangSaveBMP(HarfangPicture pict, const char *path);
extern bool HarfangRenderInit(HarfangWindow window);
extern bool HarfangRenderInitWithType(HarfangWindow window, int type);
extern HarfangWindow HarfangRenderInitWithWidthHeightResetFlags(int width, int height, uint32_t reset_flags);
extern HarfangWindow HarfangRenderInitWithWidthHeightResetFlagsFormat(int width, int height, uint32_t reset_flags, int format);
extern HarfangWindow HarfangRenderInitWithWidthHeightResetFlagsFormatDebugFlags(int width, int height, uint32_t reset_flags, int format, uint32_t debug_flags);
extern HarfangWindow HarfangRenderInitWithWidthHeightType(int width, int height, int type);
extern HarfangWindow HarfangRenderInitWithWidthHeightTypeResetFlags(int width, int height, int type, uint32_t reset_flags);
extern HarfangWindow HarfangRenderInitWithWidthHeightTypeResetFlagsFormat(int width, int height, int type, uint32_t reset_flags, int format);
extern HarfangWindow HarfangRenderInitWithWidthHeightTypeResetFlagsFormatDebugFlags(
	int width, int height, int type, uint32_t reset_flags, int format, uint32_t debug_flags);
extern HarfangWindow HarfangRenderInitWithWindowTitleWidthHeightResetFlags(const char *window_title, int width, int height, uint32_t reset_flags);
extern HarfangWindow HarfangRenderInitWithWindowTitleWidthHeightResetFlagsFormat(
	const char *window_title, int width, int height, uint32_t reset_flags, int format);
extern HarfangWindow HarfangRenderInitWithWindowTitleWidthHeightResetFlagsFormatDebugFlags(
	const char *window_title, int width, int height, uint32_t reset_flags, int format, uint32_t debug_flags);
extern HarfangWindow HarfangRenderInitWithWindowTitleWidthHeightType(const char *window_title, int width, int height, int type);
extern HarfangWindow HarfangRenderInitWithWindowTitleWidthHeightTypeResetFlags(const char *window_title, int width, int height, int type, uint32_t reset_flags);
extern HarfangWindow HarfangRenderInitWithWindowTitleWidthHeightTypeResetFlagsFormat(
	const char *window_title, int width, int height, int type, uint32_t reset_flags, int format);
extern HarfangWindow HarfangRenderInitWithWindowTitleWidthHeightTypeResetFlagsFormatDebugFlags(
	const char *window_title, int width, int height, int type, uint32_t reset_flags, int format, uint32_t debug_flags);
extern void HarfangRenderShutdown();
extern bool HarfangRenderResetToWindow(HarfangWindow win, int *width, int *height);
extern bool HarfangRenderResetToWindowWithResetFlags(HarfangWindow win, int *width, int *height, uint32_t reset_flags);
extern void HarfangRenderReset(uint32_t width, uint32_t height);
extern void HarfangRenderResetWithFlags(uint32_t width, uint32_t height, uint32_t flags);
extern void HarfangRenderResetWithFlagsFormat(uint32_t width, uint32_t height, uint32_t flags, int format);
extern void HarfangSetRenderDebug(uint32_t flags);
extern void HarfangSetViewClear(uint16_t view_id, uint16_t flags);
extern void HarfangSetViewClearWithRgba(uint16_t view_id, uint16_t flags, uint32_t rgba);
extern void HarfangSetViewClearWithRgbaDepth(uint16_t view_id, uint16_t flags, uint32_t rgba, float depth);
extern void HarfangSetViewClearWithRgbaDepthStencil(uint16_t view_id, uint16_t flags, uint32_t rgba, float depth, uint8_t stencil);
extern void HarfangSetViewClearWithCol(uint16_t view_id, uint16_t flags, const HarfangColor col);
extern void HarfangSetViewClearWithColDepth(uint16_t view_id, uint16_t flags, const HarfangColor col, float depth);
extern void HarfangSetViewClearWithColDepthStencil(uint16_t view_id, uint16_t flags, const HarfangColor col, float depth, uint8_t stencil);
extern void HarfangSetViewRect(uint16_t view_id, uint16_t x, uint16_t y, uint16_t w, uint16_t h);
extern void HarfangSetViewFrameBuffer(uint16_t view_id, HarfangFrameBufferHandle handle);
extern void HarfangSetViewMode(uint16_t view_id, int mode);
extern void HarfangTouch(uint16_t view_id);
extern uint32_t HarfangFrame();
extern void HarfangSetViewTransform(uint16_t view_id, const HarfangMat4 view, const HarfangMat44 proj);
extern void HarfangSetView2D(uint16_t id, int x, int y, int res_x, int res_y);
extern void HarfangSetView2DWithZnearZfar(uint16_t id, int x, int y, int res_x, int res_y, float znear, float zfar);
extern void HarfangSetView2DWithZnearZfarFlagsColorDepthStencil(
	uint16_t id, int x, int y, int res_x, int res_y, float znear, float zfar, uint16_t flags, const HarfangColor color, float depth, uint8_t stencil);
extern void HarfangSetView2DWithZnearZfarFlagsColorDepthStencilYUp(uint16_t id, int x, int y, int res_x, int res_y, float znear, float zfar, uint16_t flags,
	const HarfangColor color, float depth, uint8_t stencil, bool y_up);
extern void HarfangSetViewPerspective(uint16_t id, int x, int y, int res_x, int res_y, const HarfangMat4 world);
extern void HarfangSetViewPerspectiveWithZnearZfar(uint16_t id, int x, int y, int res_x, int res_y, const HarfangMat4 world, float znear, float zfar);
extern void HarfangSetViewPerspectiveWithZnearZfarZoomFactor(
	uint16_t id, int x, int y, int res_x, int res_y, const HarfangMat4 world, float znear, float zfar, float zoom_factor);
extern void HarfangSetViewPerspectiveWithZnearZfarZoomFactorFlagsColorDepthStencil(uint16_t id, int x, int y, int res_x, int res_y, const HarfangMat4 world,
	float znear, float zfar, float zoom_factor, uint16_t flags, const HarfangColor color, float depth, uint8_t stencil);
extern void HarfangSetViewOrthographic(uint16_t id, int x, int y, int res_x, int res_y, const HarfangMat4 world);
extern void HarfangSetViewOrthographicWithZnearZfar(uint16_t id, int x, int y, int res_x, int res_y, const HarfangMat4 world, float znear, float zfar);
extern void HarfangSetViewOrthographicWithZnearZfarSize(
	uint16_t id, int x, int y, int res_x, int res_y, const HarfangMat4 world, float znear, float zfar, float size);
extern void HarfangSetViewOrthographicWithZnearZfarSizeFlagsColorDepthStencil(uint16_t id, int x, int y, int res_x, int res_y, const HarfangMat4 world,
	float znear, float zfar, float size, uint16_t flags, const HarfangColor color, float depth, uint8_t stencil);
extern HarfangVertexLayout HarfangVertexLayoutPosFloatNormFloat();
extern HarfangVertexLayout HarfangVertexLayoutPosFloatNormUInt8();
extern HarfangVertexLayout HarfangVertexLayoutPosFloatColorFloat();
extern HarfangVertexLayout HarfangVertexLayoutPosFloatColorUInt8();
extern HarfangVertexLayout HarfangVertexLayoutPosFloatTexCoord0UInt8();
extern HarfangVertexLayout HarfangVertexLayoutPosFloatNormUInt8TexCoord0UInt8();
extern HarfangProgramHandle HarfangLoadProgramFromFile(const char *path);
extern HarfangProgramHandle HarfangLoadProgramFromFileWithVertexShaderPathFragmentShaderPath(const char *vertex_shader_path, const char *fragment_shader_path);
extern HarfangProgramHandle HarfangLoadProgramFromAssets(const char *name);
extern HarfangProgramHandle HarfangLoadProgramFromAssetsWithVertexShaderNameFragmentShaderName(
	const char *vertex_shader_name, const char *fragment_shader_name);
extern void HarfangDestroyProgram(HarfangProgramHandle h);
extern uint64_t HarfangLoadTextureFlagsFromFile(const char *path);
extern uint64_t HarfangLoadTextureFlagsFromAssets(const char *name);
extern HarfangTexture HarfangCreateTexture(int width, int height, const char *name, uint64_t flags);
extern HarfangTexture HarfangCreateTextureWithFormat(int width, int height, const char *name, uint64_t flags, int format);
extern HarfangTexture HarfangCreateTextureFromPicture(const HarfangPicture pic, const char *name, uint64_t flags);
extern HarfangTexture HarfangCreateTextureFromPictureWithFormat(const HarfangPicture pic, const char *name, uint64_t flags, int format);
extern void HarfangUpdateTextureFromPicture(HarfangTexture tex, const HarfangPicture pic);
extern HarfangTexture HarfangLoadTextureFromFile(const char *path, uint64_t flags, HarfangTextureInfo info);
extern HarfangTextureRef HarfangLoadTextureFromFileWithFlagsResources(const char *path, uint32_t flags, HarfangPipelineResources resources);
extern HarfangTexture HarfangLoadTextureFromAssets(const char *path, uint64_t flags, HarfangTextureInfo info);
extern HarfangTextureRef HarfangLoadTextureFromAssetsWithFlagsResources(const char *path, uint32_t flags, HarfangPipelineResources resources);
extern void HarfangDestroyTexture(const HarfangTexture tex);
extern size_t HarfangProcessTextureLoadQueue(HarfangPipelineResources res);
extern size_t HarfangProcessTextureLoadQueueWithTBudget(HarfangPipelineResources res, int64_t t_budget);
extern size_t HarfangProcessModelLoadQueue(HarfangPipelineResources res);
extern size_t HarfangProcessModelLoadQueueWithTBudget(HarfangPipelineResources res, int64_t t_budget);
extern size_t HarfangProcessLoadQueues(HarfangPipelineResources res);
extern size_t HarfangProcessLoadQueuesWithTBudget(HarfangPipelineResources res, int64_t t_budget);
extern uint32_t HarfangCaptureTexture(const HarfangPipelineResources resources, const HarfangTextureRef tex, HarfangPicture pic);
extern HarfangUniformSetValue HarfangMakeUniformSetValue(const char *name, float v);
extern HarfangUniformSetValue HarfangMakeUniformSetValueWithV(const char *name, const HarfangVec2 v);
extern HarfangUniformSetValue HarfangMakeUniformSetValueWithVec3V(const char *name, const HarfangVec3 v);
extern HarfangUniformSetValue HarfangMakeUniformSetValueWithVec4V(const char *name, const HarfangVec4 v);
extern HarfangUniformSetValue HarfangMakeUniformSetValueWithMat3V(const char *name, const HarfangMat3 v);
extern HarfangUniformSetValue HarfangMakeUniformSetValueWithMat4V(const char *name, const HarfangMat4 v);
extern HarfangUniformSetValue HarfangMakeUniformSetValueWithMat44V(const char *name, const HarfangMat44 v);
extern HarfangUniformSetTexture HarfangMakeUniformSetTexture(const char *name, const HarfangTexture texture, uint8_t stage);
extern HarfangPipelineProgram HarfangLoadPipelineProgramFromFile(const char *path, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern HarfangPipelineProgram HarfangLoadPipelineProgramFromAssets(const char *name, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern HarfangPipelineProgramRef HarfangLoadPipelineProgramRefFromFile(
	const char *path, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern HarfangPipelineProgramRef HarfangLoadPipelineProgramRefFromAssets(
	const char *name, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern HarfangViewState HarfangComputeOrthographicViewState(const HarfangMat4 world, float size, float znear, float zfar, const HarfangVec2 aspect_ratio);
extern HarfangViewState HarfangComputePerspectiveViewState(const HarfangMat4 world, float fov, float znear, float zfar, const HarfangVec2 aspect_ratio);
extern void HarfangSetMaterialProgram(HarfangMaterial mat, HarfangPipelineProgramRef program);
extern void HarfangSetMaterialValue(HarfangMaterial mat, const char *name, float v);
extern void HarfangSetMaterialValueWithV(HarfangMaterial mat, const char *name, const HarfangVec2 v);
extern void HarfangSetMaterialValueWithVec3V(HarfangMaterial mat, const char *name, const HarfangVec3 v);
extern void HarfangSetMaterialValueWithVec4V(HarfangMaterial mat, const char *name, const HarfangVec4 v);
extern void HarfangSetMaterialValueWithM(HarfangMaterial mat, const char *name, const HarfangMat3 m);
extern void HarfangSetMaterialValueWithMat4M(HarfangMaterial mat, const char *name, const HarfangMat4 m);
extern void HarfangSetMaterialValueWithMat44M(HarfangMaterial mat, const char *name, const HarfangMat44 m);
extern void HarfangSetMaterialTexture(HarfangMaterial mat, const char *name, HarfangTextureRef texture, uint8_t stage);
extern void HarfangSetMaterialTextureRef(HarfangMaterial mat, const char *name, HarfangTextureRef texture);
extern HarfangTextureRef HarfangGetMaterialTexture(HarfangMaterial mat, const char *name);
extern HarfangStringList HarfangGetMaterialTextures(HarfangMaterial mat);
extern HarfangStringList HarfangGetMaterialValues(HarfangMaterial mat);
extern int HarfangGetMaterialFaceCulling(const HarfangMaterial mat);
extern void HarfangSetMaterialFaceCulling(HarfangMaterial mat, int culling);
extern int HarfangGetMaterialDepthTest(const HarfangMaterial mat);
extern void HarfangSetMaterialDepthTest(HarfangMaterial mat, int test);
extern int HarfangGetMaterialBlendMode(const HarfangMaterial mat);
extern void HarfangSetMaterialBlendMode(HarfangMaterial mat, int mode);
extern void HarfangGetMaterialWriteRGBA(const HarfangMaterial mat, bool *write_r, bool *write_g, bool *write_b, bool *write_a);
extern void HarfangSetMaterialWriteRGBA(HarfangMaterial mat, bool write_r, bool write_g, bool write_b, bool write_a);
extern bool HarfangGetMaterialNormalMapInWorldSpace(const HarfangMaterial mat);
extern void HarfangSetMaterialNormalMapInWorldSpace(HarfangMaterial mat, bool enable);
extern bool HarfangGetMaterialWriteZ(const HarfangMaterial mat);
extern void HarfangSetMaterialWriteZ(HarfangMaterial mat, bool enable);
extern bool HarfangGetMaterialDiffuseUsesUV1(const HarfangMaterial mat);
extern void HarfangSetMaterialDiffuseUsesUV1(HarfangMaterial mat, bool enable);
extern bool HarfangGetMaterialSpecularUsesUV1(const HarfangMaterial mat);
extern void HarfangSetMaterialSpecularUsesUV1(HarfangMaterial mat, bool enable);
extern bool HarfangGetMaterialAmbientUsesUV1(const HarfangMaterial mat);
extern void HarfangSetMaterialAmbientUsesUV1(HarfangMaterial mat, bool enable);
extern bool HarfangGetMaterialSkinning(const HarfangMaterial mat);
extern void HarfangSetMaterialSkinning(HarfangMaterial mat, bool enable);
extern bool HarfangGetMaterialAlphaCut(const HarfangMaterial mat);
extern void HarfangSetMaterialAlphaCut(HarfangMaterial mat, bool enable);
extern HarfangMaterial HarfangCreateMaterial(HarfangPipelineProgramRef prg);
extern HarfangMaterial HarfangCreateMaterialWithValueNameValue(HarfangPipelineProgramRef prg, const char *value_name, const HarfangVec4 value);
extern HarfangMaterial HarfangCreateMaterialWithValueName0Value0ValueName1Value1(
	HarfangPipelineProgramRef prg, const char *value_name_0, const HarfangVec4 value_0, const char *value_name_1, const HarfangVec4 value_1);
extern HarfangRenderState HarfangComputeRenderState(int blend);
extern HarfangRenderState HarfangComputeRenderStateWithDepthTest(int blend, int depth_test);
extern HarfangRenderState HarfangComputeRenderStateWithDepthTestCulling(int blend, int depth_test, int culling);
extern HarfangRenderState HarfangComputeRenderStateWithDepthTestCullingWriteZ(int blend, int depth_test, int culling, bool write_z);
extern HarfangRenderState HarfangComputeRenderStateWithDepthTestCullingWriteZWriteR(int blend, int depth_test, int culling, bool write_z, bool write_r);
extern HarfangRenderState HarfangComputeRenderStateWithDepthTestCullingWriteZWriteRWriteG(
	int blend, int depth_test, int culling, bool write_z, bool write_r, bool write_g);
extern HarfangRenderState HarfangComputeRenderStateWithDepthTestCullingWriteZWriteRWriteGWriteB(
	int blend, int depth_test, int culling, bool write_z, bool write_r, bool write_g, bool write_b);
extern HarfangRenderState HarfangComputeRenderStateWithDepthTestCullingWriteZWriteRWriteGWriteBWriteA(
	int blend, int depth_test, int culling, bool write_z, bool write_r, bool write_g, bool write_b, bool write_a);
extern HarfangRenderState HarfangComputeRenderStateWithWriteZ(int blend, bool write_z);
extern HarfangRenderState HarfangComputeRenderStateWithWriteZWriteR(int blend, bool write_z, bool write_r);
extern HarfangRenderState HarfangComputeRenderStateWithWriteZWriteRWriteG(int blend, bool write_z, bool write_r, bool write_g);
extern HarfangRenderState HarfangComputeRenderStateWithWriteZWriteRWriteGWriteB(int blend, bool write_z, bool write_r, bool write_g, bool write_b);
extern HarfangRenderState HarfangComputeRenderStateWithWriteZWriteRWriteGWriteBWriteA(
	int blend, bool write_z, bool write_r, bool write_g, bool write_b, bool write_a);
extern uint32_t HarfangComputeSortKey(float view_depth);
extern uint32_t HarfangComputeSortKeyFromWorld(const HarfangVec3 T, const HarfangMat4 view);
extern uint32_t HarfangComputeSortKeyFromWorldWithModel(const HarfangVec3 T, const HarfangMat4 view, const HarfangMat4 model);
extern HarfangModel HarfangLoadModelFromFile(const char *path);
extern HarfangModel HarfangLoadModelFromAssets(const char *name);
extern HarfangModel HarfangCreateCubeModel(const HarfangVertexLayout decl, float x, float y, float z);
extern HarfangModel HarfangCreateSphereModel(const HarfangVertexLayout decl, float radius, int subdiv_x, int subdiv_y);
extern HarfangModel HarfangCreatePlaneModel(const HarfangVertexLayout decl, float width, float length, int subdiv_x, int subdiv_z);
extern HarfangModel HarfangCreateCylinderModel(const HarfangVertexLayout decl, float radius, float height, int subdiv_x);
extern HarfangModel HarfangCreateConeModel(const HarfangVertexLayout decl, float radius, float height, int subdiv_x);
extern HarfangModel HarfangCreateCapsuleModel(const HarfangVertexLayout decl, float radius, float height, int subdiv_x, int subdiv_y);
extern void HarfangDrawModel(uint16_t view_id, const HarfangModel mdl, HarfangProgramHandle prg, const HarfangUniformSetValueList values,
	const HarfangUniformSetTextureList textures, const HarfangMat4 matrix);
extern void HarfangDrawModelWithRenderState(uint16_t view_id, const HarfangModel mdl, HarfangProgramHandle prg, const HarfangUniformSetValueList values,
	const HarfangUniformSetTextureList textures, const HarfangMat4 matrix, HarfangRenderState render_state);
extern void HarfangDrawModelWithRenderStateDepth(uint16_t view_id, const HarfangModel mdl, HarfangProgramHandle prg, const HarfangUniformSetValueList values,
	const HarfangUniformSetTextureList textures, const HarfangMat4 matrix, HarfangRenderState render_state, uint32_t depth);
extern void HarfangDrawModelWithMatrices(uint16_t view_id, const HarfangModel mdl, HarfangProgramHandle prg, const HarfangUniformSetValueList values,
	const HarfangUniformSetTextureList textures, HarfangMat4List matrices);
extern void HarfangDrawModelWithMatricesRenderState(uint16_t view_id, const HarfangModel mdl, HarfangProgramHandle prg, const HarfangUniformSetValueList values,
	const HarfangUniformSetTextureList textures, HarfangMat4List matrices, HarfangRenderState render_state);
extern void HarfangDrawModelWithMatricesRenderStateDepth(uint16_t view_id, const HarfangModel mdl, HarfangProgramHandle prg,
	const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures, HarfangMat4List matrices, HarfangRenderState render_state,
	uint32_t depth);
extern void HarfangDrawModelWithSliceOfValuesSliceOfTextures(uint16_t view_id, const HarfangModel mdl, HarfangProgramHandle prg, size_t SliceOfvaluesToCSize,
	HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf, const HarfangMat4 matrix);
extern void HarfangDrawModelWithSliceOfValuesSliceOfTexturesRenderState(uint16_t view_id, const HarfangModel mdl, HarfangProgramHandle prg,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf,
	const HarfangMat4 matrix, HarfangRenderState render_state);
extern void HarfangDrawModelWithSliceOfValuesSliceOfTexturesRenderStateDepth(uint16_t view_id, const HarfangModel mdl, HarfangProgramHandle prg,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf,
	const HarfangMat4 matrix, HarfangRenderState render_state, uint32_t depth);
extern void HarfangDrawModelWithSliceOfValuesSliceOfTexturesSliceOfMatrices(uint16_t view_id, const HarfangModel mdl, HarfangProgramHandle prg,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf,
	size_t SliceOfmatricesToCSize, HarfangMat4 *SliceOfmatricesToCBuf);
extern void HarfangDrawModelWithSliceOfValuesSliceOfTexturesSliceOfMatricesRenderState(uint16_t view_id, const HarfangModel mdl, HarfangProgramHandle prg,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf,
	size_t SliceOfmatricesToCSize, HarfangMat4 *SliceOfmatricesToCBuf, HarfangRenderState render_state);
extern void HarfangDrawModelWithSliceOfValuesSliceOfTexturesSliceOfMatricesRenderStateDepth(uint16_t view_id, const HarfangModel mdl, HarfangProgramHandle prg,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf,
	size_t SliceOfmatricesToCSize, HarfangMat4 *SliceOfmatricesToCBuf, HarfangRenderState render_state, uint32_t depth);
extern void HarfangUpdateMaterialPipelineProgramVariant(HarfangMaterial mat, const HarfangPipelineResources resources);
extern void HarfangCreateMissingMaterialProgramValuesFromFile(HarfangMaterial mat, const HarfangPipelineResources resources);
extern void HarfangCreateMissingMaterialProgramValuesFromAssets(HarfangMaterial mat, const HarfangPipelineResources resources);
extern HarfangFrameBuffer HarfangCreateFrameBuffer(const HarfangTexture color, const HarfangTexture depth, const char *name);
extern HarfangFrameBuffer HarfangCreateFrameBufferWithColorFormatDepthFormatAaName(int color_format, int depth_format, int aa, const char *name);
extern HarfangFrameBuffer HarfangCreateFrameBufferWithWidthHeightColorFormatDepthFormatAaName(
	int width, int height, int color_format, int depth_format, int aa, const char *name);
extern HarfangTexture HarfangGetColorTexture(HarfangFrameBuffer frameBuffer);
extern HarfangTexture HarfangGetDepthTexture(HarfangFrameBuffer frameBuffer);
extern void HarfangGetTextures(HarfangFrameBuffer framebuffer, HarfangTexture color, HarfangTexture depth);
extern void HarfangDestroyFrameBuffer(HarfangFrameBuffer frameBuffer);
extern void HarfangSetTransform(const HarfangMat4 mtx);
extern void HarfangDrawLines(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg);
extern void HarfangDrawLinesWithRenderState(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg, HarfangRenderState render_state);
extern void HarfangDrawLinesWithRenderStateDepth(
	uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg, HarfangRenderState render_state, uint32_t depth);
extern void HarfangDrawLinesWithValuesTextures(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg, const HarfangUniformSetValueList values,
	const HarfangUniformSetTextureList textures);
extern void HarfangDrawLinesWithValuesTexturesRenderState(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg,
	const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures, HarfangRenderState render_state);
extern void HarfangDrawLinesWithValuesTexturesRenderStateDepth(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg,
	const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures, HarfangRenderState render_state, uint32_t depth);
extern void HarfangDrawLinesWithIdxVtxPrgValuesTextures(uint16_t view_id, const HarfangUint16TList idx, const HarfangVertices vtx, HarfangProgramHandle prg,
	const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures);
extern void HarfangDrawLinesWithIdxVtxPrgValuesTexturesRenderState(uint16_t view_id, const HarfangUint16TList idx, const HarfangVertices vtx,
	HarfangProgramHandle prg, const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures, HarfangRenderState render_state);
extern void HarfangDrawLinesWithIdxVtxPrgValuesTexturesRenderStateDepth(uint16_t view_id, const HarfangUint16TList idx, const HarfangVertices vtx,
	HarfangProgramHandle prg, const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures, HarfangRenderState render_state,
	uint32_t depth);
extern void HarfangDrawLinesWithSliceOfValuesSliceOfTextures(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg, size_t SliceOfvaluesToCSize,
	HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf);
extern void HarfangDrawLinesWithSliceOfValuesSliceOfTexturesRenderState(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf,
	HarfangRenderState render_state);
extern void HarfangDrawLinesWithSliceOfValuesSliceOfTexturesRenderStateDepth(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf,
	HarfangRenderState render_state, uint32_t depth);
extern void HarfangDrawLinesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTextures(uint16_t view_id, size_t SliceOfidxToCSize, uint16_t *SliceOfidxToCBuf,
	const HarfangVertices vtx, HarfangProgramHandle prg, size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf,
	size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf);
extern void HarfangDrawLinesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesRenderState(uint16_t view_id, size_t SliceOfidxToCSize, uint16_t *SliceOfidxToCBuf,
	const HarfangVertices vtx, HarfangProgramHandle prg, size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf,
	size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf, HarfangRenderState render_state);
extern void HarfangDrawLinesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesRenderStateDepth(uint16_t view_id, size_t SliceOfidxToCSize,
	uint16_t *SliceOfidxToCBuf, const HarfangVertices vtx, HarfangProgramHandle prg, size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf,
	size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf, HarfangRenderState render_state, uint32_t depth);
extern void HarfangDrawTriangles(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg);
extern void HarfangDrawTrianglesWithState(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg, HarfangRenderState state);
extern void HarfangDrawTrianglesWithStateDepth(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg, HarfangRenderState state, uint32_t depth);
extern void HarfangDrawTrianglesWithValuesTextures(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg,
	const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures);
extern void HarfangDrawTrianglesWithValuesTexturesState(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg,
	const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures, HarfangRenderState state);
extern void HarfangDrawTrianglesWithValuesTexturesStateDepth(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg,
	const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures, HarfangRenderState state, uint32_t depth);
extern void HarfangDrawTrianglesWithIdxVtxPrgValuesTextures(uint16_t view_id, const HarfangUint16TList idx, const HarfangVertices vtx, HarfangProgramHandle prg,
	const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures);
extern void HarfangDrawTrianglesWithIdxVtxPrgValuesTexturesState(uint16_t view_id, const HarfangUint16TList idx, const HarfangVertices vtx,
	HarfangProgramHandle prg, const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures, HarfangRenderState state);
extern void HarfangDrawTrianglesWithIdxVtxPrgValuesTexturesStateDepth(uint16_t view_id, const HarfangUint16TList idx, const HarfangVertices vtx,
	HarfangProgramHandle prg, const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures, HarfangRenderState state, uint32_t depth);
extern void HarfangDrawTrianglesWithSliceOfValuesSliceOfTextures(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf);
extern void HarfangDrawTrianglesWithSliceOfValuesSliceOfTexturesState(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf,
	HarfangRenderState state);
extern void HarfangDrawTrianglesWithSliceOfValuesSliceOfTexturesStateDepth(uint16_t view_id, const HarfangVertices vtx, HarfangProgramHandle prg,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf,
	HarfangRenderState state, uint32_t depth);
extern void HarfangDrawTrianglesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTextures(uint16_t view_id, size_t SliceOfidxToCSize, uint16_t *SliceOfidxToCBuf,
	const HarfangVertices vtx, HarfangProgramHandle prg, size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf,
	size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf);
extern void HarfangDrawTrianglesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesState(uint16_t view_id, size_t SliceOfidxToCSize, uint16_t *SliceOfidxToCBuf,
	const HarfangVertices vtx, HarfangProgramHandle prg, size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf,
	size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf, HarfangRenderState state);
extern void HarfangDrawTrianglesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesStateDepth(uint16_t view_id, size_t SliceOfidxToCSize,
	uint16_t *SliceOfidxToCBuf, const HarfangVertices vtx, HarfangProgramHandle prg, size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf,
	size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf, HarfangRenderState state, uint32_t depth);
extern void HarfangDrawSprites(uint16_t view_id, const HarfangMat3 inv_view_R, HarfangVertexLayout vtx_layout, const HarfangVec3List pos,
	const HarfangVec2 size, HarfangProgramHandle prg);
extern void HarfangDrawSpritesWithState(uint16_t view_id, const HarfangMat3 inv_view_R, HarfangVertexLayout vtx_layout, const HarfangVec3List pos,
	const HarfangVec2 size, HarfangProgramHandle prg, HarfangRenderState state);
extern void HarfangDrawSpritesWithStateDepth(uint16_t view_id, const HarfangMat3 inv_view_R, HarfangVertexLayout vtx_layout, const HarfangVec3List pos,
	const HarfangVec2 size, HarfangProgramHandle prg, HarfangRenderState state, uint32_t depth);
extern void HarfangDrawSpritesWithValuesTextures(uint16_t view_id, const HarfangMat3 inv_view_R, HarfangVertexLayout vtx_layout, const HarfangVec3List pos,
	const HarfangVec2 size, HarfangProgramHandle prg, const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures);
extern void HarfangDrawSpritesWithValuesTexturesState(uint16_t view_id, const HarfangMat3 inv_view_R, HarfangVertexLayout vtx_layout, const HarfangVec3List pos,
	const HarfangVec2 size, HarfangProgramHandle prg, const HarfangUniformSetValueList values, const HarfangUniformSetTextureList textures,
	HarfangRenderState state);
extern void HarfangDrawSpritesWithValuesTexturesStateDepth(uint16_t view_id, const HarfangMat3 inv_view_R, HarfangVertexLayout vtx_layout,
	const HarfangVec3List pos, const HarfangVec2 size, HarfangProgramHandle prg, const HarfangUniformSetValueList values,
	const HarfangUniformSetTextureList textures, HarfangRenderState state, uint32_t depth);
extern void HarfangDrawSpritesWithSliceOfPos(uint16_t view_id, const HarfangMat3 inv_view_R, HarfangVertexLayout vtx_layout, size_t SliceOfposToCSize,
	HarfangVec3 *SliceOfposToCBuf, const HarfangVec2 size, HarfangProgramHandle prg);
extern void HarfangDrawSpritesWithSliceOfPosState(uint16_t view_id, const HarfangMat3 inv_view_R, HarfangVertexLayout vtx_layout, size_t SliceOfposToCSize,
	HarfangVec3 *SliceOfposToCBuf, const HarfangVec2 size, HarfangProgramHandle prg, HarfangRenderState state);
extern void HarfangDrawSpritesWithSliceOfPosStateDepth(uint16_t view_id, const HarfangMat3 inv_view_R, HarfangVertexLayout vtx_layout, size_t SliceOfposToCSize,
	HarfangVec3 *SliceOfposToCBuf, const HarfangVec2 size, HarfangProgramHandle prg, HarfangRenderState state, uint32_t depth);
extern void HarfangDrawSpritesWithSliceOfPosSliceOfValuesSliceOfTextures(uint16_t view_id, const HarfangMat3 inv_view_R, HarfangVertexLayout vtx_layout,
	size_t SliceOfposToCSize, HarfangVec3 *SliceOfposToCBuf, const HarfangVec2 size, HarfangProgramHandle prg, size_t SliceOfvaluesToCSize,
	HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf);
extern void HarfangDrawSpritesWithSliceOfPosSliceOfValuesSliceOfTexturesState(uint16_t view_id, const HarfangMat3 inv_view_R, HarfangVertexLayout vtx_layout,
	size_t SliceOfposToCSize, HarfangVec3 *SliceOfposToCBuf, const HarfangVec2 size, HarfangProgramHandle prg, size_t SliceOfvaluesToCSize,
	HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf, HarfangRenderState state);
extern void HarfangDrawSpritesWithSliceOfPosSliceOfValuesSliceOfTexturesStateDepth(uint16_t view_id, const HarfangMat3 inv_view_R,
	HarfangVertexLayout vtx_layout, size_t SliceOfposToCSize, HarfangVec3 *SliceOfposToCBuf, const HarfangVec2 size, HarfangProgramHandle prg,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf,
	HarfangRenderState state, uint32_t depth);
extern const HarfangPipelineInfo HarfangGetForwardPipelineInfo();
extern HarfangForwardPipeline HarfangCreateForwardPipeline();
extern HarfangForwardPipeline HarfangCreateForwardPipelineWithShadowMapResolution(int shadow_map_resolution);
extern HarfangForwardPipeline HarfangCreateForwardPipelineWithShadowMapResolutionSpot16bitShadowMap(int shadow_map_resolution, bool spot_16bit_shadow_map);
extern void HarfangDestroyForwardPipeline(HarfangForwardPipeline pipeline);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelinePointLight(const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelinePointLightWithRadius(
	const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular, float radius);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelinePointLightWithRadiusPriority(
	const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular, float radius, float priority);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelinePointLightWithRadiusPriorityShadowType(
	const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular, float radius, float priority, int shadow_type);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelinePointLightWithRadiusPriorityShadowTypeShadowBias(
	const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular, float radius, float priority, int shadow_type, float shadow_bias);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelineSpotLight(const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelineSpotLightWithRadius(
	const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular, float radius);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelineSpotLightWithRadiusInnerAngle(
	const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular, float radius, float inner_angle);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAngle(
	const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular, float radius, float inner_angle, float outer_angle);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAnglePriority(
	const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular, float radius, float inner_angle, float outer_angle, float priority);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAnglePriorityShadowType(const HarfangMat4 world,
	const HarfangColor diffuse, const HarfangColor specular, float radius, float inner_angle, float outer_angle, float priority, int shadow_type);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAnglePriorityShadowTypeShadowBias(const HarfangMat4 world,
	const HarfangColor diffuse, const HarfangColor specular, float radius, float inner_angle, float outer_angle, float priority, int shadow_type,
	float shadow_bias);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelineLinearLight(const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelineLinearLightWithPssmSplit(
	const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular, const HarfangVec4 pssm_split);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelineLinearLightWithPssmSplitPriority(
	const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular, const HarfangVec4 pssm_split, float priority);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelineLinearLightWithPssmSplitPriorityShadowType(
	const HarfangMat4 world, const HarfangColor diffuse, const HarfangColor specular, const HarfangVec4 pssm_split, float priority, int shadow_type);
extern HarfangForwardPipelineLight HarfangMakeForwardPipelineLinearLightWithPssmSplitPriorityShadowTypeShadowBias(const HarfangMat4 world,
	const HarfangColor diffuse, const HarfangColor specular, const HarfangVec4 pssm_split, float priority, int shadow_type, float shadow_bias);
extern HarfangForwardPipelineLights HarfangPrepareForwardPipelineLights(const HarfangForwardPipelineLightList lights);
extern HarfangForwardPipelineLights HarfangPrepareForwardPipelineLightsWithSliceOfLights(
	size_t SliceOflightsToCSize, HarfangForwardPipelineLight *SliceOflightsToCBuf);
extern HarfangFont HarfangLoadFontFromFile(const char *path);
extern HarfangFont HarfangLoadFontFromFileWithSize(const char *path, float size);
extern HarfangFont HarfangLoadFontFromFileWithSizeResolution(const char *path, float size, uint16_t resolution);
extern HarfangFont HarfangLoadFontFromFileWithSizeResolutionPadding(const char *path, float size, uint16_t resolution, int padding);
extern HarfangFont HarfangLoadFontFromFileWithSizeResolutionPaddingGlyphs(const char *path, float size, uint16_t resolution, int padding, const char *glyphs);
extern HarfangFont HarfangLoadFontFromAssets(const char *name);
extern HarfangFont HarfangLoadFontFromAssetsWithSize(const char *name, float size);
extern HarfangFont HarfangLoadFontFromAssetsWithSizeResolution(const char *name, float size, uint16_t resolution);
extern HarfangFont HarfangLoadFontFromAssetsWithSizeResolutionPadding(const char *name, float size, uint16_t resolution, int padding);
extern HarfangFont HarfangLoadFontFromAssetsWithSizeResolutionPaddingGlyphs(const char *name, float size, uint16_t resolution, int padding, const char *glyphs);
extern void HarfangDrawText(
	uint16_t view_id, const HarfangFont font, const char *text, HarfangProgramHandle prg, const char *page_uniform, uint8_t page_stage, const HarfangMat4 mtx);
extern void HarfangDrawTextWithPos(uint16_t view_id, const HarfangFont font, const char *text, HarfangProgramHandle prg, const char *page_uniform,
	uint8_t page_stage, const HarfangMat4 mtx, HarfangVec3 pos);
extern void HarfangDrawTextWithPosHalignValign(uint16_t view_id, const HarfangFont font, const char *text, HarfangProgramHandle prg, const char *page_uniform,
	uint8_t page_stage, const HarfangMat4 mtx, HarfangVec3 pos, int halign, int valign);
extern void HarfangDrawTextWithPosHalignValignValuesTextures(uint16_t view_id, const HarfangFont font, const char *text, HarfangProgramHandle prg,
	const char *page_uniform, uint8_t page_stage, const HarfangMat4 mtx, HarfangVec3 pos, int halign, int valign, const HarfangUniformSetValueList values,
	const HarfangUniformSetTextureList textures);
extern void HarfangDrawTextWithPosHalignValignValuesTexturesState(uint16_t view_id, const HarfangFont font, const char *text, HarfangProgramHandle prg,
	const char *page_uniform, uint8_t page_stage, const HarfangMat4 mtx, HarfangVec3 pos, int halign, int valign, const HarfangUniformSetValueList values,
	const HarfangUniformSetTextureList textures, HarfangRenderState state);
extern void HarfangDrawTextWithPosHalignValignValuesTexturesStateDepth(uint16_t view_id, const HarfangFont font, const char *text, HarfangProgramHandle prg,
	const char *page_uniform, uint8_t page_stage, const HarfangMat4 mtx, HarfangVec3 pos, int halign, int valign, const HarfangUniformSetValueList values,
	const HarfangUniformSetTextureList textures, HarfangRenderState state, uint32_t depth);
extern void HarfangDrawTextWithPosHalignValignSliceOfValuesSliceOfTextures(uint16_t view_id, const HarfangFont font, const char *text, HarfangProgramHandle prg,
	const char *page_uniform, uint8_t page_stage, const HarfangMat4 mtx, HarfangVec3 pos, int halign, int valign, size_t SliceOfvaluesToCSize,
	HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf);
extern void HarfangDrawTextWithPosHalignValignSliceOfValuesSliceOfTexturesState(uint16_t view_id, const HarfangFont font, const char *text,
	HarfangProgramHandle prg, const char *page_uniform, uint8_t page_stage, const HarfangMat4 mtx, HarfangVec3 pos, int halign, int valign,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf,
	HarfangRenderState state);
extern void HarfangDrawTextWithPosHalignValignSliceOfValuesSliceOfTexturesStateDepth(uint16_t view_id, const HarfangFont font, const char *text,
	HarfangProgramHandle prg, const char *page_uniform, uint8_t page_stage, const HarfangMat4 mtx, HarfangVec3 pos, int halign, int valign,
	size_t SliceOfvaluesToCSize, HarfangUniformSetValue *SliceOfvaluesToCBuf, size_t SliceOftexturesToCSize, HarfangUniformSetTexture *SliceOftexturesToCBuf,
	HarfangRenderState state, uint32_t depth);
extern HarfangRect HarfangComputeTextRect(const HarfangFont font, const char *text);
extern HarfangRect HarfangComputeTextRectWithXpos(const HarfangFont font, const char *text, float xpos);
extern HarfangRect HarfangComputeTextRectWithXposYpos(const HarfangFont font, const char *text, float xpos, float ypos);
extern float HarfangComputeTextHeight(const HarfangFont font, const char *text);
extern HarfangJSON HarfangLoadJsonFromFile(const char *path);
extern HarfangJSON HarfangLoadJsonFromAssets(const char *name);
extern bool HarfangSaveJsonToFile(const HarfangJSON js, const char *path);
extern bool HarfangGetJsonString(const HarfangJSON js, const char *key, const char **value);
extern bool HarfangGetJsonBool(const HarfangJSON js, const char *key, bool *value);
extern bool HarfangGetJsonInt(const HarfangJSON js, const char *key, int *value);
extern bool HarfangGetJsonFloat(const HarfangJSON js, const char *key, float *value);
extern void HarfangSetJsonValue(HarfangJSON js, const char *key, const char *value);
extern void HarfangSetJsonValueWithValue(HarfangJSON js, const char *key, bool value);
extern void HarfangSetJsonValueWithIntValue(HarfangJSON js, const char *key, int value);
extern void HarfangSetJsonValueWithFloatValue(HarfangJSON js, const char *key, float value);
extern HarfangNode HarfangCreateSceneRootNode(HarfangScene scene, const char *name, const HarfangMat4 mtx);
extern HarfangNode HarfangCreateCamera(HarfangScene scene, const HarfangMat4 mtx, float znear, float zfar);
extern HarfangNode HarfangCreateCameraWithFov(HarfangScene scene, const HarfangMat4 mtx, float znear, float zfar, float fov);
extern HarfangNode HarfangCreateOrthographicCamera(HarfangScene scene, const HarfangMat4 mtx, float znear, float zfar);
extern HarfangNode HarfangCreateOrthographicCameraWithSize(HarfangScene scene, const HarfangMat4 mtx, float znear, float zfar, float size);
extern HarfangNode HarfangCreatePointLight(HarfangScene scene, const HarfangMat4 mtx, float radius);
extern HarfangNode HarfangCreatePointLightWithDiffuse(HarfangScene scene, const HarfangMat4 mtx, float radius, const HarfangColor diffuse);
extern HarfangNode HarfangCreatePointLightWithDiffuseSpecular(
	HarfangScene scene, const HarfangMat4 mtx, float radius, const HarfangColor diffuse, const HarfangColor specular);
extern HarfangNode HarfangCreatePointLightWithDiffuseSpecularPriority(
	HarfangScene scene, const HarfangMat4 mtx, float radius, const HarfangColor diffuse, const HarfangColor specular, float priority);
extern HarfangNode HarfangCreatePointLightWithDiffuseSpecularPriorityShadowType(
	HarfangScene scene, const HarfangMat4 mtx, float radius, const HarfangColor diffuse, const HarfangColor specular, float priority, int shadow_type);
extern HarfangNode HarfangCreatePointLightWithDiffuseSpecularPriorityShadowTypeShadowBias(HarfangScene scene, const HarfangMat4 mtx, float radius,
	const HarfangColor diffuse, const HarfangColor specular, float priority, int shadow_type, float shadow_bias);
extern HarfangNode HarfangCreatePointLightWithDiffuseDiffuseIntensity(
	HarfangScene scene, const HarfangMat4 mtx, float radius, const HarfangColor diffuse, float diffuse_intensity);
extern HarfangNode HarfangCreatePointLightWithDiffuseDiffuseIntensitySpecular(
	HarfangScene scene, const HarfangMat4 mtx, float radius, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular);
extern HarfangNode HarfangCreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensity(HarfangScene scene, const HarfangMat4 mtx, float radius,
	const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity);
extern HarfangNode HarfangCreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriority(HarfangScene scene, const HarfangMat4 mtx, float radius,
	const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority);
extern HarfangNode HarfangCreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(HarfangScene scene, const HarfangMat4 mtx,
	float radius, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority, int shadow_type);
extern HarfangNode HarfangCreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(HarfangScene scene,
	const HarfangMat4 mtx, float radius, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity,
	float priority, int shadow_type, float shadow_bias);
extern HarfangNode HarfangCreateSpotLight(HarfangScene scene, const HarfangMat4 mtx, float radius, float inner_angle, float outer_angle);
extern HarfangNode HarfangCreateSpotLightWithDiffuse(
	HarfangScene scene, const HarfangMat4 mtx, float radius, float inner_angle, float outer_angle, const HarfangColor diffuse);
extern HarfangNode HarfangCreateSpotLightWithDiffuseSpecular(
	HarfangScene scene, const HarfangMat4 mtx, float radius, float inner_angle, float outer_angle, const HarfangColor diffuse, const HarfangColor specular);
extern HarfangNode HarfangCreateSpotLightWithDiffuseSpecularPriority(HarfangScene scene, const HarfangMat4 mtx, float radius, float inner_angle,
	float outer_angle, const HarfangColor diffuse, const HarfangColor specular, float priority);
extern HarfangNode HarfangCreateSpotLightWithDiffuseSpecularPriorityShadowType(HarfangScene scene, const HarfangMat4 mtx, float radius, float inner_angle,
	float outer_angle, const HarfangColor diffuse, const HarfangColor specular, float priority, int shadow_type);
extern HarfangNode HarfangCreateSpotLightWithDiffuseSpecularPriorityShadowTypeShadowBias(HarfangScene scene, const HarfangMat4 mtx, float radius,
	float inner_angle, float outer_angle, const HarfangColor diffuse, const HarfangColor specular, float priority, int shadow_type, float shadow_bias);
extern HarfangNode HarfangCreateSpotLightWithDiffuseDiffuseIntensity(
	HarfangScene scene, const HarfangMat4 mtx, float radius, float inner_angle, float outer_angle, const HarfangColor diffuse, float diffuse_intensity);
extern HarfangNode HarfangCreateSpotLightWithDiffuseDiffuseIntensitySpecular(HarfangScene scene, const HarfangMat4 mtx, float radius, float inner_angle,
	float outer_angle, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular);
extern HarfangNode HarfangCreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensity(HarfangScene scene, const HarfangMat4 mtx, float radius,
	float inner_angle, float outer_angle, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity);
extern HarfangNode HarfangCreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriority(HarfangScene scene, const HarfangMat4 mtx, float radius,
	float inner_angle, float outer_angle, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity,
	float priority);
extern HarfangNode HarfangCreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(HarfangScene scene, const HarfangMat4 mtx,
	float radius, float inner_angle, float outer_angle, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular,
	float specular_intensity, float priority, int shadow_type);
extern HarfangNode HarfangCreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(HarfangScene scene,
	const HarfangMat4 mtx, float radius, float inner_angle, float outer_angle, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular,
	float specular_intensity, float priority, int shadow_type, float shadow_bias);
extern HarfangNode HarfangCreateLinearLight(HarfangScene scene, const HarfangMat4 mtx);
extern HarfangNode HarfangCreateLinearLightWithDiffuse(HarfangScene scene, const HarfangMat4 mtx, const HarfangColor diffuse);
extern HarfangNode HarfangCreateLinearLightWithDiffuseSpecular(
	HarfangScene scene, const HarfangMat4 mtx, const HarfangColor diffuse, const HarfangColor specular);
extern HarfangNode HarfangCreateLinearLightWithDiffuseSpecularPriority(
	HarfangScene scene, const HarfangMat4 mtx, const HarfangColor diffuse, const HarfangColor specular, float priority);
extern HarfangNode HarfangCreateLinearLightWithDiffuseSpecularPriorityShadowType(
	HarfangScene scene, const HarfangMat4 mtx, const HarfangColor diffuse, const HarfangColor specular, float priority, int shadow_type);
extern HarfangNode HarfangCreateLinearLightWithDiffuseSpecularPriorityShadowTypeShadowBiasPssmSplit(HarfangScene scene, const HarfangMat4 mtx,
	const HarfangColor diffuse, const HarfangColor specular, float priority, int shadow_type, float shadow_bias, const HarfangVec4 pssm_split);
extern HarfangNode HarfangCreateLinearLightWithDiffuseDiffuseIntensity(
	HarfangScene scene, const HarfangMat4 mtx, const HarfangColor diffuse, float diffuse_intensity);
extern HarfangNode HarfangCreateLinearLightWithDiffuseDiffuseIntensitySpecular(
	HarfangScene scene, const HarfangMat4 mtx, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular);
extern HarfangNode HarfangCreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensity(
	HarfangScene scene, const HarfangMat4 mtx, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity);
extern HarfangNode HarfangCreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriority(HarfangScene scene, const HarfangMat4 mtx,
	const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority);
extern HarfangNode HarfangCreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(HarfangScene scene, const HarfangMat4 mtx,
	const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority, int shadow_type);
extern HarfangNode HarfangCreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBiasPssmSplit(HarfangScene scene,
	const HarfangMat4 mtx, const HarfangColor diffuse, float diffuse_intensity, const HarfangColor specular, float specular_intensity, float priority,
	int shadow_type, float shadow_bias, const HarfangVec4 pssm_split);
extern HarfangNode HarfangCreateObject(HarfangScene scene, const HarfangMat4 mtx, const HarfangModelRef model, const HarfangMaterialList materials);
extern HarfangNode HarfangCreateObjectWithSliceOfMaterials(
	HarfangScene scene, const HarfangMat4 mtx, const HarfangModelRef model, size_t SliceOfmaterialsToCSize, HarfangMaterial *SliceOfmaterialsToCBuf);
extern HarfangNode HarfangCreateInstanceFromFile(
	HarfangScene scene, const HarfangMat4 mtx, const char *name, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline, bool *success);
extern HarfangNode HarfangCreateInstanceFromFileWithFlags(HarfangScene scene, const HarfangMat4 mtx, const char *name, HarfangPipelineResources resources,
	const HarfangPipelineInfo pipeline, bool *success, uint32_t flags);
extern HarfangNode HarfangCreateInstanceFromAssets(
	HarfangScene scene, const HarfangMat4 mtx, const char *name, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline, bool *success);
extern HarfangNode HarfangCreateInstanceFromAssetsWithFlags(HarfangScene scene, const HarfangMat4 mtx, const char *name, HarfangPipelineResources resources,
	const HarfangPipelineInfo pipeline, bool *success, uint32_t flags);
extern HarfangNode HarfangCreateScript(HarfangScene scene);
extern HarfangNode HarfangCreateScriptWithPath(HarfangScene scene, const char *path);
extern HarfangNode HarfangCreatePhysicSphere(
	HarfangScene scene, float radius, const HarfangMat4 mtx, const HarfangModelRef model_ref, HarfangMaterialList materials);
extern HarfangNode HarfangCreatePhysicSphereWithMass(
	HarfangScene scene, float radius, const HarfangMat4 mtx, const HarfangModelRef model_ref, HarfangMaterialList materials, float mass);
extern HarfangNode HarfangCreatePhysicSphereWithSliceOfMaterials(HarfangScene scene, float radius, const HarfangMat4 mtx, const HarfangModelRef model_ref,
	size_t SliceOfmaterialsToCSize, HarfangMaterial *SliceOfmaterialsToCBuf);
extern HarfangNode HarfangCreatePhysicSphereWithSliceOfMaterialsMass(HarfangScene scene, float radius, const HarfangMat4 mtx, const HarfangModelRef model_ref,
	size_t SliceOfmaterialsToCSize, HarfangMaterial *SliceOfmaterialsToCBuf, float mass);
extern HarfangNode HarfangCreatePhysicCube(
	HarfangScene scene, const HarfangVec3 size, const HarfangMat4 mtx, const HarfangModelRef model_ref, HarfangMaterialList materials);
extern HarfangNode HarfangCreatePhysicCubeWithMass(
	HarfangScene scene, const HarfangVec3 size, const HarfangMat4 mtx, const HarfangModelRef model_ref, HarfangMaterialList materials, float mass);
extern HarfangNode HarfangCreatePhysicCubeWithSliceOfMaterials(HarfangScene scene, const HarfangVec3 size, const HarfangMat4 mtx,
	const HarfangModelRef model_ref, size_t SliceOfmaterialsToCSize, HarfangMaterial *SliceOfmaterialsToCBuf);
extern HarfangNode HarfangCreatePhysicCubeWithSliceOfMaterialsMass(HarfangScene scene, const HarfangVec3 size, const HarfangMat4 mtx,
	const HarfangModelRef model_ref, size_t SliceOfmaterialsToCSize, HarfangMaterial *SliceOfmaterialsToCBuf, float mass);
extern bool HarfangSaveSceneJsonToFile(const char *path, const HarfangScene scene, const HarfangPipelineResources resources);
extern bool HarfangSaveSceneJsonToFileWithFlags(const char *path, const HarfangScene scene, const HarfangPipelineResources resources, uint32_t flags);
extern bool HarfangSaveSceneBinaryToFile(const char *path, const HarfangScene scene, const HarfangPipelineResources resources);
extern bool HarfangSaveSceneBinaryToFileWithFlags(const char *path, const HarfangScene scene, const HarfangPipelineResources resources, uint32_t flags);
extern bool HarfangSaveSceneBinaryToData(HarfangData data, const HarfangScene scene, const HarfangPipelineResources resources);
extern bool HarfangSaveSceneBinaryToDataWithFlags(HarfangData data, const HarfangScene scene, const HarfangPipelineResources resources, uint32_t flags);
extern bool HarfangLoadSceneBinaryFromFile(const char *path, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern bool HarfangLoadSceneBinaryFromFileWithFlags(
	const char *path, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline, uint32_t flags);
extern bool HarfangLoadSceneBinaryFromAssets(const char *name, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern bool HarfangLoadSceneBinaryFromAssetsWithFlags(
	const char *name, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline, uint32_t flags);
extern bool HarfangLoadSceneJsonFromFile(const char *path, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern bool HarfangLoadSceneJsonFromFileWithFlags(
	const char *path, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline, uint32_t flags);
extern bool HarfangLoadSceneJsonFromAssets(const char *name, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern bool HarfangLoadSceneJsonFromAssetsWithFlags(
	const char *name, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline, uint32_t flags);
extern bool HarfangLoadSceneBinaryFromDataAndFile(
	const HarfangData data, const char *name, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern bool HarfangLoadSceneBinaryFromDataAndFileWithFlags(
	const HarfangData data, const char *name, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline, uint32_t flags);
extern bool HarfangLoadSceneBinaryFromDataAndAssets(
	const HarfangData data, const char *name, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern bool HarfangLoadSceneBinaryFromDataAndAssetsWithFlags(
	const HarfangData data, const char *name, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline, uint32_t flags);
extern bool HarfangLoadSceneFromFile(const char *path, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern bool HarfangLoadSceneFromFileWithFlags(
	const char *path, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline, uint32_t flags);
extern bool HarfangLoadSceneFromAssets(const char *name, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern bool HarfangLoadSceneFromAssetsWithFlags(
	const char *name, HarfangScene scene, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline, uint32_t flags);
extern HarfangNodeList HarfangDuplicateNodesFromFile(
	HarfangScene scene, const HarfangNodeList nodes, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern HarfangNodeList HarfangDuplicateNodesFromAssets(
	HarfangScene scene, const HarfangNodeList nodes, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern HarfangNodeList HarfangDuplicateNodesAndChildrenFromFile(
	HarfangScene scene, const HarfangNodeList nodes, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern HarfangNodeList HarfangDuplicateNodesAndChildrenFromAssets(
	HarfangScene scene, const HarfangNodeList nodes, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern HarfangNode HarfangDuplicateNodeFromFile(HarfangScene scene, HarfangNode node, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern HarfangNode HarfangDuplicateNodeFromAssets(HarfangScene scene, HarfangNode node, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern HarfangNodeList HarfangDuplicateNodeAndChildrenFromFile(
	HarfangScene scene, HarfangNode node, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern HarfangNodeList HarfangDuplicateNodeAndChildrenFromAssets(
	HarfangScene scene, HarfangNode node, HarfangPipelineResources resources, const HarfangPipelineInfo pipeline);
extern HarfangForwardPipelineFog HarfangGetSceneForwardPipelineFog(const HarfangScene scene);
extern HarfangForwardPipelineLightList HarfangGetSceneForwardPipelineLights(const HarfangScene scene);
extern uint16_t HarfangGetSceneForwardPipelinePassViewId(const HarfangSceneForwardPipelinePassViewId views, int pass);
extern void HarfangPrepareSceneForwardPipelineCommonRenderData(uint16_t *view_id, const HarfangScene scene, HarfangSceneForwardPipelineRenderData render_data,
	const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources, HarfangSceneForwardPipelinePassViewId views);
extern void HarfangPrepareSceneForwardPipelineCommonRenderDataWithDebugName(uint16_t *view_id, const HarfangScene scene,
	HarfangSceneForwardPipelineRenderData render_data, const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources,
	HarfangSceneForwardPipelinePassViewId views, const char *debug_name);
extern void HarfangPrepareSceneForwardPipelineViewDependentRenderData(uint16_t *view_id, const HarfangViewState view_state, const HarfangScene scene,
	HarfangSceneForwardPipelineRenderData render_data, const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources,
	HarfangSceneForwardPipelinePassViewId views);
extern void HarfangPrepareSceneForwardPipelineViewDependentRenderDataWithDebugName(uint16_t *view_id, const HarfangViewState view_state,
	const HarfangScene scene, HarfangSceneForwardPipelineRenderData render_data, const HarfangForwardPipeline pipeline,
	const HarfangPipelineResources resources, HarfangSceneForwardPipelinePassViewId views, const char *debug_name);
extern void HarfangSubmitSceneToForwardPipeline(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect, const HarfangViewState view_state,
	const HarfangForwardPipeline pipeline, const HarfangSceneForwardPipelineRenderData render_data, const HarfangPipelineResources resources,
	const HarfangSceneForwardPipelinePassViewId views);
extern void HarfangSubmitSceneToForwardPipelineWithFrameBuffer(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	const HarfangViewState view_state, const HarfangForwardPipeline pipeline, const HarfangSceneForwardPipelineRenderData render_data,
	const HarfangPipelineResources resources, const HarfangSceneForwardPipelinePassViewId views, HarfangFrameBufferHandle frame_buffer);
extern void HarfangSubmitSceneToForwardPipelineWithFrameBufferDebugName(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	const HarfangViewState view_state, const HarfangForwardPipeline pipeline, const HarfangSceneForwardPipelineRenderData render_data,
	const HarfangPipelineResources resources, const HarfangSceneForwardPipelinePassViewId views, HarfangFrameBufferHandle frame_buffer, const char *debug_name);
extern void HarfangSubmitSceneToForwardPipelineWithAaaAaaConfigFrame(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	const HarfangViewState view_state, const HarfangForwardPipeline pipeline, const HarfangSceneForwardPipelineRenderData render_data,
	const HarfangPipelineResources resources, const HarfangSceneForwardPipelinePassViewId views, const HarfangForwardPipelineAAA aaa,
	const HarfangForwardPipelineAAAConfig aaa_config, int frame);
extern void HarfangSubmitSceneToForwardPipelineWithAaaAaaConfigFrameFrameBuffer(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	const HarfangViewState view_state, const HarfangForwardPipeline pipeline, const HarfangSceneForwardPipelineRenderData render_data,
	const HarfangPipelineResources resources, const HarfangSceneForwardPipelinePassViewId views, const HarfangForwardPipelineAAA aaa,
	const HarfangForwardPipelineAAAConfig aaa_config, int frame, HarfangFrameBufferHandle frame_buffer);
extern void HarfangSubmitSceneToForwardPipelineWithAaaAaaConfigFrameFrameBufferDebugName(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	const HarfangViewState view_state, const HarfangForwardPipeline pipeline, const HarfangSceneForwardPipelineRenderData render_data,
	const HarfangPipelineResources resources, const HarfangSceneForwardPipelinePassViewId views, const HarfangForwardPipelineAAA aaa,
	const HarfangForwardPipelineAAAConfig aaa_config, int frame, HarfangFrameBufferHandle frame_buffer, const char *debug_name);
extern void HarfangSubmitSceneToPipeline(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect, const HarfangViewState view_state,
	const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources, const HarfangSceneForwardPipelinePassViewId views);
extern void HarfangSubmitSceneToPipelineWithFb(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect, const HarfangViewState view_state,
	const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources, const HarfangSceneForwardPipelinePassViewId views,
	HarfangFrameBufferHandle fb);
extern void HarfangSubmitSceneToPipelineWithFbDebugName(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	const HarfangViewState view_state, const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources,
	const HarfangSceneForwardPipelinePassViewId views, HarfangFrameBufferHandle fb, const char *debug_name);
extern void HarfangSubmitSceneToPipelineWithFovAxisIsHorizontal(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	bool fov_axis_is_horizontal, const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources,
	const HarfangSceneForwardPipelinePassViewId views);
extern void HarfangSubmitSceneToPipelineWithFovAxisIsHorizontalFb(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	bool fov_axis_is_horizontal, const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources,
	const HarfangSceneForwardPipelinePassViewId views, HarfangFrameBufferHandle fb);
extern void HarfangSubmitSceneToPipelineWithFovAxisIsHorizontalFbDebugName(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	bool fov_axis_is_horizontal, const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources,
	const HarfangSceneForwardPipelinePassViewId views, HarfangFrameBufferHandle fb, const char *debug_name);
extern void HarfangSubmitSceneToPipelineWithAaaAaaConfigFrame(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	const HarfangViewState view_state, const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources,
	const HarfangSceneForwardPipelinePassViewId views, HarfangForwardPipelineAAA aaa, const HarfangForwardPipelineAAAConfig aaa_config, int frame);
extern void HarfangSubmitSceneToPipelineWithAaaAaaConfigFrameFrameBuffer(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	const HarfangViewState view_state, const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources,
	const HarfangSceneForwardPipelinePassViewId views, HarfangForwardPipelineAAA aaa, const HarfangForwardPipelineAAAConfig aaa_config, int frame,
	HarfangFrameBufferHandle frame_buffer);
extern void HarfangSubmitSceneToPipelineWithAaaAaaConfigFrameFrameBufferDebugName(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	const HarfangViewState view_state, const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources,
	const HarfangSceneForwardPipelinePassViewId views, HarfangForwardPipelineAAA aaa, const HarfangForwardPipelineAAAConfig aaa_config, int frame,
	HarfangFrameBufferHandle frame_buffer, const char *debug_name);
extern void HarfangSubmitSceneToPipelineWithFovAxisIsHorizontalAaaAaaConfigFrame(uint16_t *view_id, const HarfangScene scene, const HarfangIntRect rect,
	bool fov_axis_is_horizontal, const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources,
	const HarfangSceneForwardPipelinePassViewId views, HarfangForwardPipelineAAA aaa, const HarfangForwardPipelineAAAConfig aaa_config, int frame);
extern void HarfangSubmitSceneToPipelineWithFovAxisIsHorizontalAaaAaaConfigFrameFrameBuffer(uint16_t *view_id, const HarfangScene scene,
	const HarfangIntRect rect, bool fov_axis_is_horizontal, const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources,
	const HarfangSceneForwardPipelinePassViewId views, HarfangForwardPipelineAAA aaa, const HarfangForwardPipelineAAAConfig aaa_config, int frame,
	HarfangFrameBufferHandle frame_buffer);
extern void HarfangSubmitSceneToPipelineWithFovAxisIsHorizontalAaaAaaConfigFrameFrameBufferDebugName(uint16_t *view_id, const HarfangScene scene,
	const HarfangIntRect rect, bool fov_axis_is_horizontal, const HarfangForwardPipeline pipeline, const HarfangPipelineResources resources,
	const HarfangSceneForwardPipelinePassViewId views, HarfangForwardPipelineAAA aaa, const HarfangForwardPipelineAAAConfig aaa_config, int frame,
	HarfangFrameBufferHandle frame_buffer, const char *debug_name);
extern bool HarfangLoadForwardPipelineAAAConfigFromFile(const char *path, HarfangForwardPipelineAAAConfig config);
extern bool HarfangLoadForwardPipelineAAAConfigFromAssets(const char *path, HarfangForwardPipelineAAAConfig config);
extern bool HarfangSaveForwardPipelineAAAConfigToFile(const char *path, const HarfangForwardPipelineAAAConfig config);
extern HarfangForwardPipelineAAA HarfangCreateForwardPipelineAAAFromFile(const char *path, const HarfangForwardPipelineAAAConfig config);
extern HarfangForwardPipelineAAA HarfangCreateForwardPipelineAAAFromFileWithSsgiRatio(
	const char *path, const HarfangForwardPipelineAAAConfig config, int ssgi_ratio);
extern HarfangForwardPipelineAAA HarfangCreateForwardPipelineAAAFromFileWithSsgiRatioSsrRatio(
	const char *path, const HarfangForwardPipelineAAAConfig config, int ssgi_ratio, int ssr_ratio);
extern HarfangForwardPipelineAAA HarfangCreateForwardPipelineAAAFromAssets(const char *path, const HarfangForwardPipelineAAAConfig config);
extern HarfangForwardPipelineAAA HarfangCreateForwardPipelineAAAFromAssetsWithSsgiRatio(
	const char *path, const HarfangForwardPipelineAAAConfig config, int ssgi_ratio);
extern HarfangForwardPipelineAAA HarfangCreateForwardPipelineAAAFromAssetsWithSsgiRatioSsrRatio(
	const char *path, const HarfangForwardPipelineAAAConfig config, int ssgi_ratio, int ssr_ratio);
extern void HarfangDestroyForwardPipelineAAA(HarfangForwardPipelineAAA pipeline);
extern void HarfangDebugSceneExplorer(HarfangScene scene, const char *name);
extern HarfangNodeList HarfangGetNodesInContact(const HarfangScene scene, const HarfangNode with, const HarfangNodePairContacts node_pair_contacts);
extern HarfangContactList HarfangGetNodePairContacts(const HarfangNode first, const HarfangNode second, const HarfangNodePairContacts node_pair_contacts);
extern void HarfangSceneSyncToSystemsFromFile(HarfangScene scene, HarfangSceneLuaVM vm);
extern void HarfangSceneSyncToSystemsFromFileWithPhysics(HarfangScene scene, HarfangSceneBullet3Physics physics);
extern void HarfangSceneSyncToSystemsFromFileWithPhysicsVm(HarfangScene scene, HarfangSceneBullet3Physics physics, HarfangSceneLuaVM vm);
extern void HarfangSceneSyncToSystemsFromAssets(HarfangScene scene, HarfangSceneLuaVM vm);
extern void HarfangSceneSyncToSystemsFromAssetsWithPhysics(HarfangScene scene, HarfangSceneBullet3Physics physics);
extern void HarfangSceneSyncToSystemsFromAssetsWithPhysicsVm(HarfangScene scene, HarfangSceneBullet3Physics physics, HarfangSceneLuaVM vm);
extern void HarfangSceneUpdateSystems(HarfangScene scene, HarfangSceneClocks clocks, int64_t dt);
extern void HarfangSceneUpdateSystemsWithVm(HarfangScene scene, HarfangSceneClocks clocks, int64_t dt, HarfangSceneLuaVM vm);
extern void HarfangSceneUpdateSystemsWithPhysicsStepMaxPhysicsStep(
	HarfangScene scene, HarfangSceneClocks clocks, int64_t dt, HarfangSceneBullet3Physics physics, int64_t step, int max_physics_step);
extern void HarfangSceneUpdateSystemsWithPhysicsStepMaxPhysicsStepVm(
	HarfangScene scene, HarfangSceneClocks clocks, int64_t dt, HarfangSceneBullet3Physics physics, int64_t step, int max_physics_step, HarfangSceneLuaVM vm);
extern void HarfangSceneUpdateSystemsWithPhysicsContactsStepMaxPhysicsStep(HarfangScene scene, HarfangSceneClocks clocks, int64_t dt,
	HarfangSceneBullet3Physics physics, HarfangNodePairContacts contacts, int64_t step, int max_physics_step);
extern void HarfangSceneUpdateSystemsWithPhysicsContactsStepMaxPhysicsStepVm(HarfangScene scene, HarfangSceneClocks clocks, int64_t dt,
	HarfangSceneBullet3Physics physics, HarfangNodePairContacts contacts, int64_t step, int max_physics_step, HarfangSceneLuaVM vm);
extern size_t HarfangSceneGarbageCollectSystems(HarfangScene scene);
extern size_t HarfangSceneGarbageCollectSystemsWithVm(HarfangScene scene, HarfangSceneLuaVM vm);
extern size_t HarfangSceneGarbageCollectSystemsWithPhysics(HarfangScene scene, HarfangSceneBullet3Physics physics);
extern size_t HarfangSceneGarbageCollectSystemsWithPhysicsVm(HarfangScene scene, HarfangSceneBullet3Physics physics, HarfangSceneLuaVM vm);
extern void HarfangSceneClearSystems(HarfangScene scene);
extern void HarfangSceneClearSystemsWithVm(HarfangScene scene, HarfangSceneLuaVM vm);
extern void HarfangSceneClearSystemsWithPhysics(HarfangScene scene, HarfangSceneBullet3Physics physics);
extern void HarfangSceneClearSystemsWithPhysicsVm(HarfangScene scene, HarfangSceneBullet3Physics physics, HarfangSceneLuaVM vm);
extern void HarfangInputInit();
extern void HarfangInputShutdown();
extern HarfangMouseState HarfangReadMouse();
extern HarfangMouseState HarfangReadMouseWithName(const char *name);
extern HarfangStringList HarfangGetMouseNames();
extern HarfangKeyboardState HarfangReadKeyboard();
extern HarfangKeyboardState HarfangReadKeyboardWithName(const char *name);
extern const char *HarfangGetKeyName(int key);
extern const char *HarfangGetKeyNameWithName(int key, const char *name);
extern HarfangStringList HarfangGetKeyboardNames();
extern HarfangGamepadState HarfangReadGamepad();
extern HarfangGamepadState HarfangReadGamepadWithName(const char *name);
extern HarfangStringList HarfangGetGamepadNames();
extern HarfangJoystickState HarfangReadJoystick();
extern HarfangJoystickState HarfangReadJoystickWithName(const char *name);
extern HarfangStringList HarfangGetJoystickNames();
extern HarfangStringList HarfangGetJoystickDeviceNames();
extern HarfangVRControllerState HarfangReadVRController();
extern HarfangVRControllerState HarfangReadVRControllerWithName(const char *name);
extern void HarfangSendVRControllerHapticPulse(int64_t duration);
extern void HarfangSendVRControllerHapticPulseWithName(int64_t duration, const char *name);
extern HarfangStringList HarfangGetVRControllerNames();
extern HarfangVRGenericTrackerState HarfangReadVRGenericTracker();
extern HarfangVRGenericTrackerState HarfangReadVRGenericTrackerWithName(const char *name);
extern HarfangStringList HarfangGetVRGenericTrackerNames();
extern void HarfangImGuiNewFrame();
extern void HarfangImGuiRender();
extern bool HarfangImGuiBegin(const char *name);
extern bool HarfangImGuiBeginWithOpenFlags(const char *name, bool *open, int flags);
extern void HarfangImGuiEnd();
extern bool HarfangImGuiBeginChild(const char *id);
extern bool HarfangImGuiBeginChildWithSize(const char *id, const HarfangVec2 size);
extern bool HarfangImGuiBeginChildWithSizeBorder(const char *id, const HarfangVec2 size, bool border);
extern bool HarfangImGuiBeginChildWithSizeBorderFlags(const char *id, const HarfangVec2 size, bool border, int flags);
extern void HarfangImGuiEndChild();
extern HarfangVec2 HarfangImGuiGetContentRegionMax();
extern HarfangVec2 HarfangImGuiGetContentRegionAvail();
extern float HarfangImGuiGetContentRegionAvailWidth();
extern HarfangVec2 HarfangImGuiGetWindowContentRegionMin();
extern HarfangVec2 HarfangImGuiGetWindowContentRegionMax();
extern float HarfangImGuiGetWindowContentRegionWidth();
extern HarfangImDrawList HarfangImGuiGetWindowDrawList();
extern HarfangVec2 HarfangImGuiGetWindowPos();
extern HarfangVec2 HarfangImGuiGetWindowSize();
extern float HarfangImGuiGetWindowWidth();
extern float HarfangImGuiGetWindowHeight();
extern bool HarfangImGuiIsWindowCollapsed();
extern void HarfangImGuiSetWindowFontScale(float scale);
extern void HarfangImGuiSetNextWindowPos(const HarfangVec2 pos);
extern void HarfangImGuiSetNextWindowPosWithCondition(const HarfangVec2 pos, int condition);
extern void HarfangImGuiSetNextWindowPosCenter();
extern void HarfangImGuiSetNextWindowPosCenterWithCondition(int condition);
extern void HarfangImGuiSetNextWindowSize(const HarfangVec2 size);
extern void HarfangImGuiSetNextWindowSizeWithCondition(const HarfangVec2 size, int condition);
extern void HarfangImGuiSetNextWindowSizeConstraints(const HarfangVec2 size_min, const HarfangVec2 size_max);
extern void HarfangImGuiSetNextWindowContentSize(const HarfangVec2 size);
extern void HarfangImGuiSetNextWindowContentWidth(float width);
extern void HarfangImGuiSetNextWindowCollapsed(bool collapsed, int condition);
extern void HarfangImGuiSetNextWindowFocus();
extern void HarfangImGuiSetWindowPos(const char *name, const HarfangVec2 pos);
extern void HarfangImGuiSetWindowPosWithCondition(const char *name, const HarfangVec2 pos, int condition);
extern void HarfangImGuiSetWindowSize(const char *name, const HarfangVec2 size);
extern void HarfangImGuiSetWindowSizeWithCondition(const char *name, const HarfangVec2 size, int condition);
extern void HarfangImGuiSetWindowCollapsed(const char *name, bool collapsed);
extern void HarfangImGuiSetWindowCollapsedWithCondition(const char *name, bool collapsed, int condition);
extern void HarfangImGuiSetWindowFocus(const char *name);
extern float HarfangImGuiGetScrollX();
extern float HarfangImGuiGetScrollY();
extern float HarfangImGuiGetScrollMaxX();
extern float HarfangImGuiGetScrollMaxY();
extern void HarfangImGuiSetScrollX(float scroll_x);
extern void HarfangImGuiSetScrollY(float scroll_y);
extern void HarfangImGuiSetScrollHereY();
extern void HarfangImGuiSetScrollHereYWithCenterYRatio(float center_y_ratio);
extern void HarfangImGuiSetScrollFromPosY(float pos_y);
extern void HarfangImGuiSetScrollFromPosYWithCenterYRatio(float pos_y, float center_y_ratio);
extern void HarfangImGuiSetKeyboardFocusHere();
extern void HarfangImGuiSetKeyboardFocusHereWithOffset(int offset);
extern void HarfangImGuiPushFont(HarfangImFont font);
extern void HarfangImGuiPopFont();
extern void HarfangImGuiPushStyleColor(int idx, const HarfangColor color);
extern void HarfangImGuiPopStyleColor();
extern void HarfangImGuiPopStyleColorWithCount(int count);
extern void HarfangImGuiPushStyleVar(int idx, float value);
extern void HarfangImGuiPushStyleVarWithValue(int idx, const HarfangVec2 value);
extern void HarfangImGuiPopStyleVar();
extern void HarfangImGuiPopStyleVarWithCount(int count);
extern HarfangImFont HarfangImGuiGetFont();
extern float HarfangImGuiGetFontSize();
extern HarfangVec2 HarfangImGuiGetFontTexUvWhitePixel();
extern uint32_t HarfangImGuiGetColorU32(int idx);
extern uint32_t HarfangImGuiGetColorU32WithAlphaMultiplier(int idx, float alpha_multiplier);
extern uint32_t HarfangImGuiGetColorU32WithColor(const HarfangColor color);
extern void HarfangImGuiPushItemWidth(float item_width);
extern void HarfangImGuiPopItemWidth();
extern float HarfangImGuiCalcItemWidth();
extern void HarfangImGuiPushTextWrapPos();
extern void HarfangImGuiPushTextWrapPosWithWrapPosX(float wrap_pos_x);
extern void HarfangImGuiPopTextWrapPos();
extern void HarfangImGuiPushAllowKeyboardFocus(bool v);
extern void HarfangImGuiPopAllowKeyboardFocus();
extern void HarfangImGuiPushButtonRepeat(bool repeat);
extern void HarfangImGuiPopButtonRepeat();
extern void HarfangImGuiSeparator();
extern void HarfangImGuiSameLine();
extern void HarfangImGuiSameLineWithPosX(float pos_x);
extern void HarfangImGuiSameLineWithPosXSpacingW(float pos_x, float spacing_w);
extern void HarfangImGuiNewLine();
extern void HarfangImGuiSpacing();
extern void HarfangImGuiDummy(const HarfangVec2 size);
extern void HarfangImGuiIndent();
extern void HarfangImGuiIndentWithWidth(float width);
extern void HarfangImGuiUnindent();
extern void HarfangImGuiUnindentWithWidth(float width);
extern void HarfangImGuiBeginGroup();
extern void HarfangImGuiEndGroup();
extern HarfangVec2 HarfangImGuiGetCursorPos();
extern float HarfangImGuiGetCursorPosX();
extern float HarfangImGuiGetCursorPosY();
extern void HarfangImGuiSetCursorPos(const HarfangVec2 local_pos);
extern void HarfangImGuiSetCursorPosX(float x);
extern void HarfangImGuiSetCursorPosY(float y);
extern HarfangVec2 HarfangImGuiGetCursorStartPos();
extern HarfangVec2 HarfangImGuiGetCursorScreenPos();
extern void HarfangImGuiSetCursorScreenPos(const HarfangVec2 pos);
extern void HarfangImGuiAlignTextToFramePadding();
extern float HarfangImGuiGetTextLineHeight();
extern float HarfangImGuiGetTextLineHeightWithSpacing();
extern float HarfangImGuiGetFrameHeightWithSpacing();
extern void HarfangImGuiColumns();
extern void HarfangImGuiColumnsWithCount(int count);
extern void HarfangImGuiColumnsWithCountId(int count, const char *id);
extern void HarfangImGuiColumnsWithCountIdWithBorder(int count, const char *id, bool with_border);
extern void HarfangImGuiNextColumn();
extern int HarfangImGuiGetColumnIndex();
extern float HarfangImGuiGetColumnOffset();
extern float HarfangImGuiGetColumnOffsetWithColumnIndex(int column_index);
extern void HarfangImGuiSetColumnOffset(int column_index, float offset_x);
extern float HarfangImGuiGetColumnWidth();
extern float HarfangImGuiGetColumnWidthWithColumnIndex(int column_index);
extern void HarfangImGuiSetColumnWidth(int column_index, float width);
extern int HarfangImGuiGetColumnsCount();
extern void HarfangImGuiPushID(const char *id);
extern void HarfangImGuiPushIDWithId(int id);
extern void HarfangImGuiPopID();
extern unsigned int HarfangImGuiGetID(const char *id);
extern void HarfangImGuiText(const char *text);
extern void HarfangImGuiTextColored(const HarfangColor color, const char *text);
extern void HarfangImGuiTextDisabled(const char *text);
extern void HarfangImGuiTextWrapped(const char *text);
extern void HarfangImGuiTextUnformatted(const char *text);
extern void HarfangImGuiLabelText(const char *label, const char *text);
extern void HarfangImGuiBullet();
extern void HarfangImGuiBulletText(const char *label);
extern bool HarfangImGuiButton(const char *label);
extern bool HarfangImGuiButtonWithSize(const char *label, const HarfangVec2 size);
extern bool HarfangImGuiSmallButton(const char *label);
extern bool HarfangImGuiInvisibleButton(const char *text, const HarfangVec2 size);
extern void HarfangImGuiImage(const HarfangTexture tex, const HarfangVec2 size);
extern void HarfangImGuiImageWithUv0(const HarfangTexture tex, const HarfangVec2 size, const HarfangVec2 uv0);
extern void HarfangImGuiImageWithUv0Uv1(const HarfangTexture tex, const HarfangVec2 size, const HarfangVec2 uv0, const HarfangVec2 uv1);
extern void HarfangImGuiImageWithUv0Uv1TintCol(
	const HarfangTexture tex, const HarfangVec2 size, const HarfangVec2 uv0, const HarfangVec2 uv1, const HarfangColor tint_col);
extern void HarfangImGuiImageWithUv0Uv1TintColBorderCol(
	const HarfangTexture tex, const HarfangVec2 size, const HarfangVec2 uv0, const HarfangVec2 uv1, const HarfangColor tint_col, const HarfangColor border_col);
extern bool HarfangImGuiImageButton(const HarfangTexture tex, const HarfangVec2 size);
extern bool HarfangImGuiImageButtonWithUv0(const HarfangTexture tex, const HarfangVec2 size, const HarfangVec2 uv0);
extern bool HarfangImGuiImageButtonWithUv0Uv1(const HarfangTexture tex, const HarfangVec2 size, const HarfangVec2 uv0, const HarfangVec2 uv1);
extern bool HarfangImGuiImageButtonWithUv0Uv1FramePadding(
	const HarfangTexture tex, const HarfangVec2 size, const HarfangVec2 uv0, const HarfangVec2 uv1, int frame_padding);
extern bool HarfangImGuiImageButtonWithUv0Uv1FramePaddingBgCol(
	const HarfangTexture tex, const HarfangVec2 size, const HarfangVec2 uv0, const HarfangVec2 uv1, int frame_padding, const HarfangColor bg_col);
extern bool HarfangImGuiImageButtonWithUv0Uv1FramePaddingBgColTintCol(const HarfangTexture tex, const HarfangVec2 size, const HarfangVec2 uv0,
	const HarfangVec2 uv1, int frame_padding, const HarfangColor bg_col, const HarfangColor tint_col);
extern bool HarfangImGuiInputText(const char *label, const char *text, size_t max_size, const char **out);
extern bool HarfangImGuiInputTextWithFlags(const char *label, const char *text, size_t max_size, const char **out, int flags);
extern bool HarfangImGuiCheckbox(const char *label, bool *value);
extern bool HarfangImGuiRadioButton(const char *label, bool active);
extern bool HarfangImGuiRadioButtonWithVVButton(const char *label, int *v, int v_button);
extern bool HarfangImGuiBeginCombo(const char *label, const char *preview_value);
extern bool HarfangImGuiBeginComboWithFlags(const char *label, const char *preview_value, int flags);
extern void HarfangImGuiEndCombo();
extern bool HarfangImGuiCombo(const char *label, int *current_item, const HarfangStringList items);
extern bool HarfangImGuiComboWithHeightInItems(const char *label, int *current_item, const HarfangStringList items, int height_in_items);
extern bool HarfangImGuiComboWithSliceOfItems(const char *label, int *current_item, size_t SliceOfitemsToCSize, const char **SliceOfitemsToCBuf);
extern bool HarfangImGuiComboWithSliceOfItemsHeightInItems(
	const char *label, int *current_item, size_t SliceOfitemsToCSize, const char **SliceOfitemsToCBuf, int height_in_items);
extern bool HarfangImGuiColorButton(const char *id, HarfangColor color);
extern bool HarfangImGuiColorButtonWithFlags(const char *id, HarfangColor color, int flags);
extern bool HarfangImGuiColorButtonWithFlagsSize(const char *id, HarfangColor color, int flags, const HarfangVec2 size);
extern bool HarfangImGuiColorEdit(const char *label, HarfangColor color);
extern bool HarfangImGuiColorEditWithFlags(const char *label, HarfangColor color, int flags);
extern void HarfangImGuiProgressBar(float fraction);
extern void HarfangImGuiProgressBarWithSize(float fraction, const HarfangVec2 size);
extern void HarfangImGuiProgressBarWithSizeOverlay(float fraction, const HarfangVec2 size, const char *overlay);
extern bool HarfangImGuiDragFloat(const char *label, float *v);
extern bool HarfangImGuiDragFloatWithVSpeed(const char *label, float *v, float v_speed);
extern bool HarfangImGuiDragFloatWithVSpeedVMinVMax(const char *label, float *v, float v_speed, float v_min, float v_max);
extern bool HarfangImGuiDragVec2(const char *label, HarfangVec2 v);
extern bool HarfangImGuiDragVec2WithVSpeed(const char *label, HarfangVec2 v, float v_speed);
extern bool HarfangImGuiDragVec2WithVSpeedVMinVMax(const char *label, HarfangVec2 v, float v_speed, float v_min, float v_max);
extern bool HarfangImGuiDragVec3(const char *label, HarfangVec3 v);
extern bool HarfangImGuiDragVec3WithVSpeed(const char *label, HarfangVec3 v, float v_speed);
extern bool HarfangImGuiDragVec3WithVSpeedVMinVMax(const char *label, HarfangVec3 v, float v_speed, float v_min, float v_max);
extern bool HarfangImGuiDragVec4(const char *label, HarfangVec4 v);
extern bool HarfangImGuiDragVec4WithVSpeed(const char *label, HarfangVec4 v, float v_speed);
extern bool HarfangImGuiDragVec4WithVSpeedVMinVMax(const char *label, HarfangVec4 v, float v_speed, float v_min, float v_max);
extern bool HarfangImGuiDragIntVec2(const char *label, HarfangIVec2 v);
extern bool HarfangImGuiDragIntVec2WithVSpeed(const char *label, HarfangIVec2 v, float v_speed);
extern bool HarfangImGuiDragIntVec2WithVSpeedVMinVMax(const char *label, HarfangIVec2 v, float v_speed, int v_min, int v_max);
extern bool HarfangImGuiInputInt(const char *label, int *v);
extern bool HarfangImGuiInputIntWithStepStepFast(const char *label, int *v, int step, int step_fast);
extern bool HarfangImGuiInputIntWithStepStepFastFlags(const char *label, int *v, int step, int step_fast, int flags);
extern bool HarfangImGuiInputFloat(const char *label, float *v);
extern bool HarfangImGuiInputFloatWithStepStepFast(const char *label, float *v, float step, float step_fast);
extern bool HarfangImGuiInputFloatWithStepStepFastDecimalPrecision(const char *label, float *v, float step, float step_fast, int decimal_precision);
extern bool HarfangImGuiInputFloatWithStepStepFastDecimalPrecisionFlags(
	const char *label, float *v, float step, float step_fast, int decimal_precision, int flags);
extern bool HarfangImGuiInputVec2(const char *label, HarfangVec2 v);
extern bool HarfangImGuiInputVec2WithDecimalPrecision(const char *label, HarfangVec2 v, int decimal_precision);
extern bool HarfangImGuiInputVec2WithDecimalPrecisionFlags(const char *label, HarfangVec2 v, int decimal_precision, int flags);
extern bool HarfangImGuiInputVec3(const char *label, HarfangVec3 v);
extern bool HarfangImGuiInputVec3WithDecimalPrecision(const char *label, HarfangVec3 v, int decimal_precision);
extern bool HarfangImGuiInputVec3WithDecimalPrecisionFlags(const char *label, HarfangVec3 v, int decimal_precision, int flags);
extern bool HarfangImGuiInputVec4(const char *label, HarfangVec4 v);
extern bool HarfangImGuiInputVec4WithDecimalPrecision(const char *label, HarfangVec4 v, int decimal_precision);
extern bool HarfangImGuiInputVec4WithDecimalPrecisionFlags(const char *label, HarfangVec4 v, int decimal_precision, int flags);
extern bool HarfangImGuiInputIntVec2(const char *label, HarfangIVec2 v);
extern bool HarfangImGuiInputIntVec2WithFlags(const char *label, HarfangIVec2 v, int flags);
extern bool HarfangImGuiSliderInt(const char *label, int *v, int v_min, int v_max);
extern bool HarfangImGuiSliderIntWithFormat(const char *label, int *v, int v_min, int v_max, const char *format);
extern bool HarfangImGuiSliderIntVec2(const char *label, HarfangIVec2 v, int v_min, int v_max);
extern bool HarfangImGuiSliderIntVec2WithFormat(const char *label, HarfangIVec2 v, int v_min, int v_max, const char *format);
extern bool HarfangImGuiSliderFloat(const char *label, float *v, float v_min, float v_max);
extern bool HarfangImGuiSliderFloatWithFormat(const char *label, float *v, float v_min, float v_max, const char *format);
extern bool HarfangImGuiSliderVec2(const char *label, HarfangVec2 v, float v_min, float v_max);
extern bool HarfangImGuiSliderVec2WithFormat(const char *label, HarfangVec2 v, float v_min, float v_max, const char *format);
extern bool HarfangImGuiSliderVec3(const char *label, HarfangVec3 v, float v_min, float v_max);
extern bool HarfangImGuiSliderVec3WithFormat(const char *label, HarfangVec3 v, float v_min, float v_max, const char *format);
extern bool HarfangImGuiSliderVec4(const char *label, HarfangVec4 v, float v_min, float v_max);
extern bool HarfangImGuiSliderVec4WithFormat(const char *label, HarfangVec4 v, float v_min, float v_max, const char *format);
extern bool HarfangImGuiTreeNode(const char *label);
extern bool HarfangImGuiTreeNodeEx(const char *label, int flags);
extern void HarfangImGuiTreePush(const char *id);
extern void HarfangImGuiTreePop();
extern float HarfangImGuiGetTreeNodeToLabelSpacing();
extern void HarfangImGuiSetNextItemOpen(bool is_open);
extern void HarfangImGuiSetNextItemOpenWithCondition(bool is_open, int condition);
extern bool HarfangImGuiCollapsingHeader(const char *label);
extern bool HarfangImGuiCollapsingHeaderWithFlags(const char *label, int flags);
extern bool HarfangImGuiCollapsingHeaderWithPOpen(const char *label, bool *p_open);
extern bool HarfangImGuiCollapsingHeaderWithPOpenFlags(const char *label, bool *p_open, int flags);
extern bool HarfangImGuiSelectable(const char *label);
extern bool HarfangImGuiSelectableWithSelected(const char *label, bool selected);
extern bool HarfangImGuiSelectableWithSelectedFlags(const char *label, bool selected, int flags);
extern bool HarfangImGuiSelectableWithSelectedFlagsSize(const char *label, bool selected, int flags, const HarfangVec2 size);
extern bool HarfangImGuiListBox(const char *label, int *current_item, const HarfangStringList items);
extern bool HarfangImGuiListBoxWithHeightInItems(const char *label, int *current_item, const HarfangStringList items, int height_in_items);
extern bool HarfangImGuiListBoxWithSliceOfItems(const char *label, int *current_item, size_t SliceOfitemsToCSize, const char **SliceOfitemsToCBuf);
extern bool HarfangImGuiListBoxWithSliceOfItemsHeightInItems(
	const char *label, int *current_item, size_t SliceOfitemsToCSize, const char **SliceOfitemsToCBuf, int height_in_items);
extern void HarfangImGuiSetTooltip(const char *text);
extern void HarfangImGuiBeginTooltip();
extern void HarfangImGuiEndTooltip();
extern bool HarfangImGuiBeginMainMenuBar();
extern void HarfangImGuiEndMainMenuBar();
extern bool HarfangImGuiBeginMenuBar();
extern void HarfangImGuiEndMenuBar();
extern bool HarfangImGuiBeginMenu(const char *label);
extern bool HarfangImGuiBeginMenuWithEnabled(const char *label, bool enabled);
extern void HarfangImGuiEndMenu();
extern bool HarfangImGuiMenuItem(const char *label);
extern bool HarfangImGuiMenuItemWithShortcut(const char *label, const char *shortcut);
extern bool HarfangImGuiMenuItemWithShortcutSelected(const char *label, const char *shortcut, bool selected);
extern bool HarfangImGuiMenuItemWithShortcutSelectedEnabled(const char *label, const char *shortcut, bool selected, bool enabled);
extern void HarfangImGuiOpenPopup(const char *id);
extern bool HarfangImGuiBeginPopup(const char *id);
extern bool HarfangImGuiBeginPopupModal(const char *name);
extern bool HarfangImGuiBeginPopupModalWithOpen(const char *name, bool *open);
extern bool HarfangImGuiBeginPopupModalWithOpenFlags(const char *name, bool *open, int flags);
extern bool HarfangImGuiBeginPopupContextItem(const char *id);
extern bool HarfangImGuiBeginPopupContextItemWithMouseButton(const char *id, int mouse_button);
extern bool HarfangImGuiBeginPopupContextWindow();
extern bool HarfangImGuiBeginPopupContextWindowWithId(const char *id);
extern bool HarfangImGuiBeginPopupContextWindowWithIdFlags(const char *id, int flags);
extern bool HarfangImGuiBeginPopupContextVoid();
extern bool HarfangImGuiBeginPopupContextVoidWithId(const char *id);
extern bool HarfangImGuiBeginPopupContextVoidWithIdMouseButton(const char *id, int mouse_button);
extern void HarfangImGuiEndPopup();
extern void HarfangImGuiCloseCurrentPopup();
extern void HarfangImGuiPushClipRect(const HarfangVec2 clip_rect_min, const HarfangVec2 clip_rect_max, bool intersect_with_current_clip_rect);
extern void HarfangImGuiPopClipRect();
extern bool HarfangImGuiIsItemHovered();
extern bool HarfangImGuiIsItemHoveredWithFlags(int flags);
extern bool HarfangImGuiIsItemActive();
extern bool HarfangImGuiIsItemClicked();
extern bool HarfangImGuiIsItemClickedWithMouseButton(int mouse_button);
extern bool HarfangImGuiIsItemVisible();
extern bool HarfangImGuiIsAnyItemHovered();
extern bool HarfangImGuiIsAnyItemActive();
extern HarfangVec2 HarfangImGuiGetItemRectMin();
extern HarfangVec2 HarfangImGuiGetItemRectMax();
extern HarfangVec2 HarfangImGuiGetItemRectSize();
extern void HarfangImGuiSetItemAllowOverlap();
extern void HarfangImGuiSetItemDefaultFocus();
extern bool HarfangImGuiIsWindowHovered();
extern bool HarfangImGuiIsWindowHoveredWithFlags(int flags);
extern bool HarfangImGuiIsWindowFocused();
extern bool HarfangImGuiIsWindowFocusedWithFlags(int flags);
extern bool HarfangImGuiIsRectVisible(const HarfangVec2 size);
extern bool HarfangImGuiIsRectVisibleWithRectMinRectMax(const HarfangVec2 rect_min, const HarfangVec2 rect_max);
extern float HarfangImGuiGetTime();
extern int HarfangImGuiGetFrameCount();
extern HarfangVec2 HarfangImGuiCalcTextSize(const char *text);
extern HarfangVec2 HarfangImGuiCalcTextSizeWithHideTextAfterDoubleDash(const char *text, bool hide_text_after_double_dash);
extern HarfangVec2 HarfangImGuiCalcTextSizeWithHideTextAfterDoubleDashWrapWidth(const char *text, bool hide_text_after_double_dash, float wrap_width);
extern bool HarfangImGuiIsKeyDown(int key_index);
extern bool HarfangImGuiIsKeyPressed(int key_index);
extern bool HarfangImGuiIsKeyPressedWithRepeat(int key_index, bool repeat);
extern bool HarfangImGuiIsKeyReleased(int key_index);
extern bool HarfangImGuiIsMouseDown(int button);
extern bool HarfangImGuiIsMouseClicked(int button);
extern bool HarfangImGuiIsMouseClickedWithRepeat(int button, bool repeat);
extern bool HarfangImGuiIsMouseDoubleClicked(int button);
extern bool HarfangImGuiIsMouseReleased(int button);
extern bool HarfangImGuiIsMouseHoveringRect(const HarfangVec2 rect_min, const HarfangVec2 rect_max);
extern bool HarfangImGuiIsMouseHoveringRectWithClip(const HarfangVec2 rect_min, const HarfangVec2 rect_max, bool clip);
extern bool HarfangImGuiIsMouseDragging(int button);
extern bool HarfangImGuiIsMouseDraggingWithLockThreshold(int button, float lock_threshold);
extern HarfangVec2 HarfangImGuiGetMousePos();
extern HarfangVec2 HarfangImGuiGetMousePosOnOpeningCurrentPopup();
extern HarfangVec2 HarfangImGuiGetMouseDragDelta();
extern HarfangVec2 HarfangImGuiGetMouseDragDeltaWithButton(int button);
extern HarfangVec2 HarfangImGuiGetMouseDragDeltaWithButtonLockThreshold(int button, float lock_threshold);
extern void HarfangImGuiResetMouseDragDelta();
extern void HarfangImGuiResetMouseDragDeltaWithButton(int button);
extern void HarfangImGuiCaptureKeyboardFromApp(bool capture);
extern void HarfangImGuiCaptureMouseFromApp(bool capture);
extern bool HarfangImGuiWantCaptureMouse();
extern void HarfangImGuiMouseDrawCursor(const bool draw_cursor);
extern void HarfangImGuiInit(float font_size, HarfangProgramHandle imgui_program, HarfangProgramHandle imgui_image_program);
extern HarfangDearImguiContext HarfangImGuiInitContext(float font_size, HarfangProgramHandle imgui_program, HarfangProgramHandle imgui_image_program);
extern void HarfangImGuiShutdown();
extern void HarfangImGuiBeginFrame(int width, int height, int64_t dt_clock, const HarfangMouseState mouse, const HarfangKeyboardState keyboard);
extern void HarfangImGuiBeginFrameWithCtxWidthHeightDtClockMouseKeyboard(
	HarfangDearImguiContext ctx, int width, int height, int64_t dt_clock, const HarfangMouseState mouse, const HarfangKeyboardState keyboard);
extern void HarfangImGuiEndFrameWithCtx(const HarfangDearImguiContext ctx);
extern void HarfangImGuiEndFrameWithCtxViewId(const HarfangDearImguiContext ctx, uint16_t view_id);
extern void HarfangImGuiEndFrame();
extern void HarfangImGuiEndFrameWithViewId(uint16_t view_id);
extern void HarfangImGuiClearInputBuffer();
extern bool HarfangOpenFolderDialog(const char *title, const char **folder_name);
extern bool HarfangOpenFolderDialogWithInitialDir(const char *title, const char **folder_name, const char *initial_dir);
extern bool HarfangOpenFileDialog(const char *title, const HarfangFileFilterList filters, const char **file);
extern bool HarfangOpenFileDialogWithInitialDir(const char *title, const HarfangFileFilterList filters, const char **file, const char *initial_dir);
extern bool HarfangSaveFileDialog(const char *title, const HarfangFileFilterList filters, const char **file);
extern bool HarfangSaveFileDialogWithInitialDir(const char *title, const HarfangFileFilterList filters, const char **file, const char *initial_dir);
extern void HarfangFpsControllerWithKeyUpKeyDownKeyLeftKeyRightBtnDxDyPosRotSpeedDtT(
	bool key_up, bool key_down, bool key_left, bool key_right, bool btn, float dx, float dy, HarfangVec3 pos, HarfangVec3 rot, float speed, int64_t dt_t);
extern void HarfangFpsController(const HarfangKeyboard keyboard, const HarfangMouse mouse, HarfangVec3 pos, HarfangVec3 rot, float speed, int64_t dt);
extern void HarfangSleep(int64_t duration);
extern bool HarfangAudioInit();
extern void HarfangAudioShutdown();
extern int HarfangLoadWAVSoundFile(const char *path);
extern int HarfangLoadWAVSoundAsset(const char *name);
extern int HarfangLoadOGGSoundFile(const char *path);
extern int HarfangLoadOGGSoundAsset(const char *name);
extern void HarfangUnloadSound(int snd);
extern void HarfangSetListener(const HarfangMat4 world, const HarfangVec3 velocity);
extern int HarfangPlayStereo(int snd, const HarfangStereoSourceState state);
extern int HarfangPlaySpatialized(int snd, const HarfangSpatializedSourceState state);
extern int HarfangStreamWAVFileStereo(const char *path, const HarfangStereoSourceState state);
extern int HarfangStreamWAVAssetStereo(const char *name, const HarfangStereoSourceState state);
extern int HarfangStreamWAVFileSpatialized(const char *path, const HarfangSpatializedSourceState state);
extern int HarfangStreamWAVAssetSpatialized(const char *name, const HarfangSpatializedSourceState state);
extern int HarfangStreamOGGFileStereo(const char *path, const HarfangStereoSourceState state);
extern int HarfangStreamOGGAssetStereo(const char *name, const HarfangStereoSourceState state);
extern int HarfangStreamOGGFileSpatialized(const char *path, const HarfangSpatializedSourceState state);
extern int HarfangStreamOGGAssetSpatialized(const char *name, const HarfangSpatializedSourceState state);
extern int64_t HarfangGetSourceDuration(int source);
extern int64_t HarfangGetSourceTimecode(int source);
extern bool HarfangSetSourceTimecode(int source, int64_t t);
extern void HarfangSetSourceVolume(int source, float volume);
extern void HarfangSetSourcePanning(int source, float panning);
extern void HarfangSetSourceRepeat(int source, int repeat);
extern void HarfangSetSourceTransform(int source, const HarfangMat4 world, const HarfangVec3 velocity);
extern int HarfangGetSourceState(int source);
extern void HarfangPauseSource(int source);
extern void HarfangStopSource(int source);
extern void HarfangStopAllSources();
extern bool HarfangOpenVRInit();
extern void HarfangOpenVRShutdown();
extern HarfangOpenVREyeFrameBuffer HarfangOpenVRCreateEyeFrameBuffer();
extern HarfangOpenVREyeFrameBuffer HarfangOpenVRCreateEyeFrameBufferWithAa(int aa);
extern void HarfangOpenVRDestroyEyeFrameBuffer(HarfangOpenVREyeFrameBuffer eye_fb);
extern HarfangOpenVRState HarfangOpenVRGetState(const HarfangMat4 body, float znear, float zfar);
extern void HarfangOpenVRStateToViewState(const HarfangOpenVRState state, HarfangViewState left, HarfangViewState right);
extern void HarfangOpenVRSubmitFrame(const HarfangOpenVREyeFrameBuffer left, const HarfangOpenVREyeFrameBuffer right);
extern void HarfangOpenVRPostPresentHandoff();
extern HarfangTexture HarfangOpenVRGetColorTexture(const HarfangOpenVREyeFrameBuffer eye);
extern HarfangTexture HarfangOpenVRGetDepthTexture(const HarfangOpenVREyeFrameBuffer eye);
extern HarfangIVec2 HarfangOpenVRGetFrameBufferSize();
extern bool HarfangOpenVRIsHMDMounted();
extern bool HarfangOpenXRInit();
extern bool HarfangOpenXRInitWithExtensionsFlagsEnable(uint16_t ExtensionsFlagsEnable);
extern void HarfangOpenXRShutdown();
extern HarfangOpenXREyeFrameBufferList HarfangOpenXRCreateEyeFrameBuffer();
extern HarfangOpenXREyeFrameBufferList HarfangOpenXRCreateEyeFrameBufferWithAa(int aa);
extern void HarfangOpenXRDestroyEyeFrameBuffer(HarfangOpenXREyeFrameBuffer eye_fb);
extern const char *HarfangOpenXRGetInstanceInfo();
extern bool HarfangOpenXRGetEyeGaze(HarfangMat4 eye_gaze);
extern bool HarfangOpenXRGetHeadPose(HarfangMat4 head_pose);
extern HarfangOpenXRFrameInfo HarfangOpenXRSubmitSceneToForwardPipeline(const HarfangMat4 cam_offset,
	HarfangFunctionReturningVoidTakingMat4Ptr update_controllers,
	HarfangFunctionReturningUint16TTakingRectOfIntPtrViewStatePtrUint16TPtrFrameBufferHandlePtr draw_scene, uint16_t *view_id, float z_near, float z_far);
extern void HarfangOpenXRFinishSubmitFrameBuffer(const HarfangOpenXRFrameInfo frameInfo);
extern HarfangTexture HarfangOpenXRGetColorTexture(const HarfangOpenXREyeFrameBuffer eye);
extern HarfangTexture HarfangOpenXRGetDepthTexture(const HarfangOpenXREyeFrameBuffer eye);
extern HarfangTexture HarfangOpenXRGetColorTextureFromId(const HarfangOpenXREyeFrameBufferList eyes, const HarfangOpenXRFrameInfo frame_info, const int index);
extern HarfangTexture HarfangOpenXRGetDepthTextureFromId(const HarfangOpenXREyeFrameBufferList eyes, const HarfangOpenXRFrameInfo frame_info, const int index);
extern bool HarfangIsHandJointActive(int hand);
extern HarfangMat4 HarfangGetHandJointPose(int hand, int handJoint);
extern float HarfangGetHandJointRadius(int hand, int handJoint);
extern HarfangVec3 HarfangGetHandJointLinearVelocity(int hand, int handJoint);
extern HarfangVec3 HarfangGetHandJointAngularVelocity(int hand, int handJoint);
extern bool HarfangSRanipalInit();
extern void HarfangSRanipalShutdown();
extern void HarfangSRanipalLaunchEyeCalibration();
extern bool HarfangSRanipalIsViveProEye();
extern HarfangSRanipalState HarfangSRanipalGetState();
extern HarfangVertex HarfangMakeVertex(const HarfangVec3 pos);
extern HarfangVertex HarfangMakeVertexWithNrm(const HarfangVec3 pos, const HarfangVec3 nrm);
extern HarfangVertex HarfangMakeVertexWithNrmUv0(const HarfangVec3 pos, const HarfangVec3 nrm, const HarfangVec2 uv0);
extern HarfangVertex HarfangMakeVertexWithNrmUv0Color0(const HarfangVec3 pos, const HarfangVec3 nrm, const HarfangVec2 uv0, const HarfangColor color0);
extern bool HarfangSaveGeometryToFile(const char *path, const HarfangGeometry geo);
extern HarfangIsoSurface HarfangNewIsoSurface(int width, int height, int depth);
extern void HarfangIsoSurfaceSphere(HarfangIsoSurface surface, int width, int height, int depth, float x, float y, float z, float radius);
extern void HarfangIsoSurfaceSphereWithValue(HarfangIsoSurface surface, int width, int height, int depth, float x, float y, float z, float radius, float value);
extern void HarfangIsoSurfaceSphereWithValueExponent(
	HarfangIsoSurface surface, int width, int height, int depth, float x, float y, float z, float radius, float value, float exponent);
extern HarfangIsoSurface HarfangGaussianBlurIsoSurface(const HarfangIsoSurface surface, int width, int height, int depth);
extern bool HarfangIsoSurfaceToModel(HarfangModelBuilder builder, const HarfangIsoSurface surface, int width, int height, int depth);
extern bool HarfangIsoSurfaceToModelWithMaterial(
	HarfangModelBuilder builder, const HarfangIsoSurface surface, int width, int height, int depth, uint16_t material);
extern bool HarfangIsoSurfaceToModelWithMaterialIsolevel(
	HarfangModelBuilder builder, const HarfangIsoSurface surface, int width, int height, int depth, uint16_t material, float isolevel);
extern bool HarfangIsoSurfaceToModelWithMaterialIsolevelScaleXScaleYScaleZ(HarfangModelBuilder builder, const HarfangIsoSurface surface, int width, int height,
	int depth, uint16_t material, float isolevel, float scale_x, float scale_y, float scale_z);
extern HarfangBloom HarfangCreateBloomFromFile(const char *path, int ratio);
extern HarfangBloom HarfangCreateBloomFromAssets(const char *path, int ratio);
extern void HarfangDestroyBloom(HarfangBloom bloom);
extern void HarfangApplyBloom(uint16_t *view_id, const HarfangIntRect rect, const HarfangTexture input, HarfangFrameBufferHandle output, HarfangBloom bloom,
	float threshold, float smoothness, float intensity);
extern HarfangSAO HarfangCreateSAOFromFile(const char *path, int ratio);
extern HarfangSAO HarfangCreateSAOFromAssets(const char *path, int ratio);
extern void HarfangDestroySAO(HarfangSAO sao);
extern void HarfangComputeSAO(uint16_t *view_id, const HarfangIntRect rect, const HarfangTexture attr0, const HarfangTexture attr1, const HarfangTexture noise,
	HarfangFrameBufferHandle output, const HarfangSAO sao, const HarfangMat44 projection, float bias, float radius, int sample_count, float sharpness);
extern size_t HarfangBeginProfilerSection(const char *name);
extern size_t HarfangBeginProfilerSectionWithSectionDetails(const char *name, const char *section_details);
extern void HarfangEndProfilerSection(size_t section_idx);
extern HarfangProfilerFrame HarfangEndProfilerFrame();
extern HarfangProfilerFrame HarfangCaptureProfilerFrame();
extern void HarfangPrintProfilerFrame(const HarfangProfilerFrame profiler_frame);
extern HarfangIVideoStreamer HarfangMakeVideoStreamer(const char *module_path);
extern bool HarfangUpdateTexture(HarfangIVideoStreamer streamer, intptr_t *handle, HarfangTexture texture, HarfangIVec2 size, int *format);
extern bool HarfangUpdateTextureWithDestroy(
	HarfangIVideoStreamer streamer, intptr_t *handle, HarfangTexture texture, HarfangIVec2 size, int *format, bool destroy);
extern HarfangPipeline HarfangCastForwardPipelineToPipeline(HarfangForwardPipeline o);
extern HarfangForwardPipeline HarfangCastPipelineToForwardPipeline(HarfangPipeline o);
extern uint32_t HarfangGetRFNone();
extern uint32_t HarfangGetRFMSAA2X();
extern uint32_t HarfangGetRFMSAA4X();
extern uint32_t HarfangGetRFMSAA8X();
extern uint32_t HarfangGetRFMSAA16X();
extern uint32_t HarfangGetRFVSync();
extern uint32_t HarfangGetRFMaxAnisotropy();
extern uint32_t HarfangGetRFCapture();
extern uint32_t HarfangGetRFFlushAfterRender();
extern uint32_t HarfangGetRFFlipAfterRender();
extern uint32_t HarfangGetRFSRGBBackBuffer();
extern uint32_t HarfangGetRFHDR10();
extern uint32_t HarfangGetRFHiDPI();
extern uint32_t HarfangGetRFDepthClamp();
extern uint32_t HarfangGetRFSuspend();
extern uint32_t HarfangGetDFIFH();
extern uint32_t HarfangGetDFProfiler();
extern uint32_t HarfangGetDFStats();
extern uint32_t HarfangGetDFText();
extern uint32_t HarfangGetDFWireframe();
extern HarfangFrameBufferHandle HarfangGetInvalidFrameBufferHandle();
extern uint16_t HarfangGetCFNone();
extern uint16_t HarfangGetCFColor();
extern uint16_t HarfangGetCFDepth();
extern uint16_t HarfangGetCFStencil();
extern uint16_t HarfangGetCFDiscardColor0();
extern uint16_t HarfangGetCFDiscardColor1();
extern uint16_t HarfangGetCFDiscardColor2();
extern uint16_t HarfangGetCFDiscardColor3();
extern uint16_t HarfangGetCFDiscardColor4();
extern uint16_t HarfangGetCFDiscardColor5();
extern uint16_t HarfangGetCFDiscardColor6();
extern uint16_t HarfangGetCFDiscardColor7();
extern uint16_t HarfangGetCFDiscardDepth();
extern uint16_t HarfangGetCFDiscardStencil();
extern uint16_t HarfangGetCFDiscardColorAll();
extern uint16_t HarfangGetCFDiscardAll();
extern uint64_t HarfangGetTFUMirror();
extern uint64_t HarfangGetTFUClamp();
extern uint64_t HarfangGetTFUBorder();
extern uint64_t HarfangGetTFVMirror();
extern uint64_t HarfangGetTFVClamp();
extern uint64_t HarfangGetTFVBorder();
extern uint64_t HarfangGetTFWMirror();
extern uint64_t HarfangGetTFWClamp();
extern uint64_t HarfangGetTFWBorder();
extern uint64_t HarfangGetTFSamplerMinPoint();
extern uint64_t HarfangGetTFSamplerMinAnisotropic();
extern uint64_t HarfangGetTFSamplerMagPoint();
extern uint64_t HarfangGetTFSamplerMagAnisotropic();
extern uint64_t HarfangGetTFBlitDestination();
extern uint64_t HarfangGetTFReadBack();
extern uint64_t HarfangGetTFRenderTarget();
extern HarfangModelRef HarfangGetInvalidModelRef();
extern HarfangTextureRef HarfangGetInvalidTextureRef();
extern HarfangMaterialRef HarfangGetInvalidMaterialRef();
extern HarfangPipelineProgramRef HarfangGetInvalidPipelineProgramRef();
extern HarfangSceneAnimRef HarfangGetInvalidSceneAnimRef();
extern int64_t HarfangGetUnspecifiedAnimTime();
extern HarfangNode HarfangGetNullNode();
extern uint32_t HarfangGetLSSFNodes();
extern uint32_t HarfangGetLSSFScene();
extern uint32_t HarfangGetLSSFAnims();
extern uint32_t HarfangGetLSSFKeyValues();
extern uint32_t HarfangGetLSSFPhysics();
extern uint32_t HarfangGetLSSFScripts();
extern uint32_t HarfangGetLSSFAll();
extern uint32_t HarfangGetLSSFQueueTextureLoads();
extern uint32_t HarfangGetLSSFFreezeMatrixToTransformOnSave();
extern uint32_t HarfangGetLSSFQueueModelLoads();
extern uint32_t HarfangGetLSSFDoNotChangeCurrentCameraIfValid();
extern HarfangSignalReturningVoidTakingConstCharPtr HarfangGetOnTextInput();
extern int HarfangGetInvalidAudioStreamRef();
extern int HarfangGetSNDInvalid();
extern int HarfangGetSRCInvalid();
#ifdef __cplusplus
}
#endif
