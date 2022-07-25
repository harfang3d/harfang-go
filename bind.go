package harfang

// #include "wrapper.h"
// #cgo CFLAGS: -I . -Wall -Wno-unused-variable -Wno-unused-function -O3
// #cgo CXXFLAGS: -std=c++14 -O3
// #cgo linux pkg-config: gtk+-3.0
// #cgo linux LDFLAGS: -L${SRCDIR}/linux -lhg_go -lharfang -lm -lstdc++ -Wl,--no-as-needed -ldl -lGL -lXrandr -lXext -lX11 -lglib-2.0
// #cgo windows LDFLAGS: -L${SRCDIR}/windows -lhg_go -lharfang -lGdi32 -lDbghelp -lshell32 -loleaut32 -luuid -lcomdlg32 -lOle32 -lWinmm -lstdc++
// #cgo LDFLAGS: -lstdc++ -L. -lharfang
import "C"

import (
	"reflect"
	"runtime"
	"unsafe"
)

func wrapFloat(goValue *float32) (wrapped *C.float, finisher func()) {
	if goValue != nil {
		cValue := C.float(*goValue)
		wrapped = &cValue
		finisher = func() {
			*goValue = float32(cValue)
		}
	} else {
		finisher = func() {}
	}
	return
}

func wrapString(value string) (wrapped *C.char, finisher func()) {
	wrapped = C.CString(value)
	finisher = func() { C.free(unsafe.Pointer(wrapped)) } // nolint: gas
	return
}

func wrapBytes(value []byte) (wrapped unsafe.Pointer, finisher func()) {
	wrapped = C.CBytes(value)
	finisher = func() { C.free(wrapped) } // nolint: gas
	return
}

type stringBuffer struct {
	ptr  unsafe.Pointer
	size int
}

func newStringBuffer(initialValue string) *stringBuffer {
	rawText := []byte(initialValue)
	bufSize := len(rawText) + 1
	newPtr := C.malloc(C.size_t(bufSize))
	zeroOffset := bufSize - 1
	buf := ptrToByteSlice(newPtr)
	copy(buf[:zeroOffset], rawText)
	buf[zeroOffset] = 0

	return &stringBuffer{ptr: newPtr, size: bufSize}
}

func (buf *stringBuffer) free() {
	C.free(buf.ptr)
	buf.size = 0
}

func (buf *stringBuffer) resizeTo(requestedSize int) {
	bufSize := requestedSize
	if bufSize < 1 {
		bufSize = 1
	}
	newPtr := C.malloc(C.size_t(bufSize))
	copySize := bufSize
	if copySize > buf.size {
		copySize = buf.size
	}
	if copySize > 0 {
		C.memcpy(newPtr, buf.ptr, C.size_t(copySize))
	}
	ptrToByteSlice(newPtr)[bufSize-1] = 0
	C.free(buf.ptr)
	buf.ptr = newPtr
	buf.size = bufSize
}

func (buf *stringBuffer) toGo() string {
	if (buf.ptr == nil) || (buf.size < 1) {
		return ""
	}
	ptrToByteSlice(buf.ptr)[buf.size-1] = 0
	return C.GoString((*C.char)(buf.ptr))
}

// unrealisticLargePointer is used to cast an arbitrary native pointer to a slice.
// Its value is chosen to fit into a 32bit architecture, and still be large
// enough to cover "any" data blob. Note that this value is in bytes.
// Should an array of larger primitives be addressed, be sure to divide the value
// by the size of the elements.
const unrealisticLargePointer = 1 << 30

func ptrToByteSlice(p unsafe.Pointer) []byte {
	return (*[unrealisticLargePointer]byte)(p)[:]
}

func ptrToUint16Slice(p unsafe.Pointer) []uint16 {
	return (*[unrealisticLargePointer / 2]uint16)(p)[:]
}

// VoidPointer  ...
type VoidPointer struct {
	h C.WrapVoidPointer
}

// Free ...
func (pointer *VoidPointer) Free() {
	C.WrapVoidPointerFree(pointer.h)
}

// IsNil ...
func (pointer *VoidPointer) IsNil() bool {
	return pointer.h == C.WrapVoidPointer(nil)
}

// GoSliceOfuint16T ...
type GoSliceOfuint16T []uint16

// Uint16TList  ...
type Uint16TList struct {
	h C.WrapUint16TList
}

// Get ...
func (pointer *Uint16TList) Get(id int) uint16 {
	v := C.WrapUint16TListGetOperator(pointer.h, C.int(id))
	return uint16(v)
}

// Set ...
func (pointer *Uint16TList) Set(id int, v uint16) {
	vToC := C.ushort(v)
	C.WrapUint16TListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *Uint16TList) Len() int32 {
	return int32(C.WrapUint16TListLenOperator(pointer.h))
}

// NewUint16TList ...
func NewUint16TList() *Uint16TList {
	retval := C.WrapConstructorUint16TList()
	retvalGO := &Uint16TList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Uint16TList) {
		C.WrapUint16TListFree(cleanval.h)
	})
	return retvalGO
}

// NewUint16TListWithSequence ...
func NewUint16TListWithSequence(sequence GoSliceOfuint16T) *Uint16TList {
	sequenceToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequence))
	sequenceToCSize := C.size_t(sequenceToC.Len)
	sequenceToCBuf := (*C.ushort)(unsafe.Pointer(sequenceToC.Data))
	retval := C.WrapConstructorUint16TListWithSequence(sequenceToCSize, sequenceToCBuf)
	retvalGO := &Uint16TList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Uint16TList) {
		C.WrapUint16TListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Uint16TList) Free() {
	C.WrapUint16TListFree(pointer.h)
}

// IsNil ...
func (pointer *Uint16TList) IsNil() bool {
	return pointer.h == C.WrapUint16TList(nil)
}

// Clear ...
func (pointer *Uint16TList) Clear() {
	C.WrapClearUint16TList(pointer.h)
}

// Reserve ...
func (pointer *Uint16TList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveUint16TList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *Uint16TList) PushBack(v uint16) {
	vToC := C.ushort(v)
	C.WrapPushBackUint16TList(pointer.h, vToC)
}

// Size ...
func (pointer *Uint16TList) Size() int32 {
	retval := C.WrapSizeUint16TList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *Uint16TList) At(idx int32) uint16 {
	idxToC := C.size_t(idx)
	retval := C.WrapAtUint16TList(pointer.h, idxToC)
	return uint16(retval)
}

// GoSliceOfuint32T ...
type GoSliceOfuint32T []uint32

// Uint32TList  ...
type Uint32TList struct {
	h C.WrapUint32TList
}

// Get ...
func (pointer *Uint32TList) Get(id int) uint32 {
	v := C.WrapUint32TListGetOperator(pointer.h, C.int(id))
	return uint32(v)
}

// Set ...
func (pointer *Uint32TList) Set(id int, v uint32) {
	vToC := C.uint32_t(v)
	C.WrapUint32TListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *Uint32TList) Len() int32 {
	return int32(C.WrapUint32TListLenOperator(pointer.h))
}

// NewUint32TList ...
func NewUint32TList() *Uint32TList {
	retval := C.WrapConstructorUint32TList()
	retvalGO := &Uint32TList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Uint32TList) {
		C.WrapUint32TListFree(cleanval.h)
	})
	return retvalGO
}

// NewUint32TListWithSequence ...
func NewUint32TListWithSequence(sequence GoSliceOfuint32T) *Uint32TList {
	sequenceToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequence))
	sequenceToCSize := C.size_t(sequenceToC.Len)
	sequenceToCBuf := (*C.uint32_t)(unsafe.Pointer(sequenceToC.Data))
	retval := C.WrapConstructorUint32TListWithSequence(sequenceToCSize, sequenceToCBuf)
	retvalGO := &Uint32TList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Uint32TList) {
		C.WrapUint32TListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Uint32TList) Free() {
	C.WrapUint32TListFree(pointer.h)
}

// IsNil ...
func (pointer *Uint32TList) IsNil() bool {
	return pointer.h == C.WrapUint32TList(nil)
}

// Clear ...
func (pointer *Uint32TList) Clear() {
	C.WrapClearUint32TList(pointer.h)
}

// Reserve ...
func (pointer *Uint32TList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveUint32TList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *Uint32TList) PushBack(v uint32) {
	vToC := C.uint32_t(v)
	C.WrapPushBackUint32TList(pointer.h, vToC)
}

// Size ...
func (pointer *Uint32TList) Size() int32 {
	retval := C.WrapSizeUint32TList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *Uint32TList) At(idx int32) uint32 {
	idxToC := C.size_t(idx)
	retval := C.WrapAtUint32TList(pointer.h, idxToC)
	return uint32(retval)
}

// GoSliceOfstring ...
type GoSliceOfstring []string

// StringList  ...
type StringList struct {
	h C.WrapStringList
}

// Get ...
func (pointer *StringList) Get(id int) string {
	v := C.WrapStringListGetOperator(pointer.h, C.int(id))
	return C.GoString(v)
}

// Set ...
func (pointer *StringList) Set(id int, v string) {
	vToC, idFinvToC := wrapString(v)
	defer idFinvToC()
	C.WrapStringListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *StringList) Len() int32 {
	return int32(C.WrapStringListLenOperator(pointer.h))
}

// NewStringList ...
func NewStringList() *StringList {
	retval := C.WrapConstructorStringList()
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// NewStringListWithSequence ...
func NewStringListWithSequence(sequence GoSliceOfstring) *StringList {
	var sequenceSpecialString []*C.char
	for _, s := range sequence {
		sequenceSpecialString = append(sequenceSpecialString, C.CString(s))
	}
	sequenceSpecialStringToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequenceSpecialString))
	sequenceSpecialStringToCSize := C.size_t(sequenceSpecialStringToC.Len)
	sequenceSpecialStringToCBuf := (**C.char)(unsafe.Pointer(sequenceSpecialStringToC.Data))
	retval := C.WrapConstructorStringListWithSequence(sequenceSpecialStringToCSize, sequenceSpecialStringToCBuf)
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *StringList) Free() {
	C.WrapStringListFree(pointer.h)
}

// IsNil ...
func (pointer *StringList) IsNil() bool {
	return pointer.h == C.WrapStringList(nil)
}

// Clear ...
func (pointer *StringList) Clear() {
	C.WrapClearStringList(pointer.h)
}

// Reserve ...
func (pointer *StringList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveStringList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *StringList) PushBack(v string) {
	vToC, idFinvToC := wrapString(v)
	defer idFinvToC()
	C.WrapPushBackStringList(pointer.h, vToC)
}

// Size ...
func (pointer *StringList) Size() int32 {
	retval := C.WrapSizeStringList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *StringList) At(idx int32) string {
	idxToC := C.size_t(idx)
	retval := C.WrapAtStringList(pointer.h, idxToC)
	return C.GoString(retval)
}

// File  Interface to a file on the host local filesystem.
type File struct {
	h C.WrapFile
}

// Free ...
func (pointer *File) Free() {
	C.WrapFileFree(pointer.h)
}

// IsNil ...
func (pointer *File) IsNil() bool {
	return pointer.h == C.WrapFile(nil)
}

// Data  ...
type Data struct {
	h C.WrapData
}

// NewData ...
func NewData() *Data {
	retval := C.WrapConstructorData()
	retvalGO := &Data{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Data) {
		C.WrapDataFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Data) Free() {
	C.WrapDataFree(pointer.h)
}

// IsNil ...
func (pointer *Data) IsNil() bool {
	return pointer.h == C.WrapData(nil)
}

// GetSize ...
func (pointer *Data) GetSize() int32 {
	retval := C.WrapGetSizeData(pointer.h)
	return int32(retval)
}

// Rewind ...
func (pointer *Data) Rewind() {
	C.WrapRewindData(pointer.h)
}

// DirEntry  ...
type DirEntry struct {
	h C.WrapDirEntry
}

// GetType ...
func (pointer *DirEntry) GetType() int32 {
	v := C.WrapDirEntryGetType(pointer.h)
	return int32(v)
}

// SetType ...
func (pointer *DirEntry) SetType(v int32) {
	vToC := C.int32_t(v)
	C.WrapDirEntrySetType(pointer.h, vToC)
}

// GetName ...
func (pointer *DirEntry) GetName() string {
	v := C.WrapDirEntryGetName(pointer.h)
	return C.GoString(v)
}

// SetName ...
func (pointer *DirEntry) SetName(v string) {
	vToC, idFinvToC := wrapString(v)
	defer idFinvToC()
	C.WrapDirEntrySetName(pointer.h, vToC)
}

// Free ...
func (pointer *DirEntry) Free() {
	C.WrapDirEntryFree(pointer.h)
}

// IsNil ...
func (pointer *DirEntry) IsNil() bool {
	return pointer.h == C.WrapDirEntry(nil)
}

// GoSliceOfDirEntry ...
type GoSliceOfDirEntry []*DirEntry

// DirEntryList  ...
type DirEntryList struct {
	h C.WrapDirEntryList
}

// Get ...
func (pointer *DirEntryList) Get(id int) *DirEntry {
	v := C.WrapDirEntryListGetOperator(pointer.h, C.int(id))
	vGO := &DirEntry{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *DirEntry) {
		C.WrapDirEntryFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *DirEntryList) Set(id int, v *DirEntry) {
	vToC := v.h
	C.WrapDirEntryListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *DirEntryList) Len() int32 {
	return int32(C.WrapDirEntryListLenOperator(pointer.h))
}

// NewDirEntryList ...
func NewDirEntryList() *DirEntryList {
	retval := C.WrapConstructorDirEntryList()
	retvalGO := &DirEntryList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *DirEntryList) {
		C.WrapDirEntryListFree(cleanval.h)
	})
	return retvalGO
}

// NewDirEntryListWithSequence ...
func NewDirEntryListWithSequence(sequence GoSliceOfDirEntry) *DirEntryList {
	var sequencePointer []C.WrapDirEntry
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapDirEntry)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorDirEntryListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &DirEntryList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *DirEntryList) {
		C.WrapDirEntryListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *DirEntryList) Free() {
	C.WrapDirEntryListFree(pointer.h)
}

// IsNil ...
func (pointer *DirEntryList) IsNil() bool {
	return pointer.h == C.WrapDirEntryList(nil)
}

// Clear ...
func (pointer *DirEntryList) Clear() {
	C.WrapClearDirEntryList(pointer.h)
}

// Reserve ...
func (pointer *DirEntryList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveDirEntryList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *DirEntryList) PushBack(v *DirEntry) {
	vToC := v.h
	C.WrapPushBackDirEntryList(pointer.h, vToC)
}

// Size ...
func (pointer *DirEntryList) Size() int32 {
	retval := C.WrapSizeDirEntryList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *DirEntryList) At(idx int32) *DirEntry {
	idxToC := C.size_t(idx)
	retval := C.WrapAtDirEntryList(pointer.h, idxToC)
	retvalGO := &DirEntry{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *DirEntry) {
		C.WrapDirEntryFree(cleanval.h)
	})
	return retvalGO
}

// Vec3  3-dimensional vector.
type Vec3 struct {
	h C.WrapVec3
}

// GetZero ...
func (pointer *Vec3) GetZero() *Vec3 {
	v := C.WrapVec3GetZero()
	vGO := &Vec3{h: v}
	return vGO
}

// Vec3GetZero ...
func Vec3GetZero() *Vec3 {
	v := C.WrapVec3GetZero()
	vGO := &Vec3{h: v}
	return vGO
}

// GetOne ...
func (pointer *Vec3) GetOne() *Vec3 {
	v := C.WrapVec3GetOne()
	vGO := &Vec3{h: v}
	return vGO
}

// Vec3GetOne ...
func Vec3GetOne() *Vec3 {
	v := C.WrapVec3GetOne()
	vGO := &Vec3{h: v}
	return vGO
}

// GetLeft ...
func (pointer *Vec3) GetLeft() *Vec3 {
	v := C.WrapVec3GetLeft()
	vGO := &Vec3{h: v}
	return vGO
}

// Vec3GetLeft ...
func Vec3GetLeft() *Vec3 {
	v := C.WrapVec3GetLeft()
	vGO := &Vec3{h: v}
	return vGO
}

// GetRight ...
func (pointer *Vec3) GetRight() *Vec3 {
	v := C.WrapVec3GetRight()
	vGO := &Vec3{h: v}
	return vGO
}

// Vec3GetRight ...
func Vec3GetRight() *Vec3 {
	v := C.WrapVec3GetRight()
	vGO := &Vec3{h: v}
	return vGO
}

// GetUp ...
func (pointer *Vec3) GetUp() *Vec3 {
	v := C.WrapVec3GetUp()
	vGO := &Vec3{h: v}
	return vGO
}

// Vec3GetUp ...
func Vec3GetUp() *Vec3 {
	v := C.WrapVec3GetUp()
	vGO := &Vec3{h: v}
	return vGO
}

// GetDown ...
func (pointer *Vec3) GetDown() *Vec3 {
	v := C.WrapVec3GetDown()
	vGO := &Vec3{h: v}
	return vGO
}

// Vec3GetDown ...
func Vec3GetDown() *Vec3 {
	v := C.WrapVec3GetDown()
	vGO := &Vec3{h: v}
	return vGO
}

// GetFront ...
func (pointer *Vec3) GetFront() *Vec3 {
	v := C.WrapVec3GetFront()
	vGO := &Vec3{h: v}
	return vGO
}

// Vec3GetFront ...
func Vec3GetFront() *Vec3 {
	v := C.WrapVec3GetFront()
	vGO := &Vec3{h: v}
	return vGO
}

// GetBack ...
func (pointer *Vec3) GetBack() *Vec3 {
	v := C.WrapVec3GetBack()
	vGO := &Vec3{h: v}
	return vGO
}

// Vec3GetBack ...
func Vec3GetBack() *Vec3 {
	v := C.WrapVec3GetBack()
	vGO := &Vec3{h: v}
	return vGO
}

// GetX ...
func (pointer *Vec3) GetX() float32 {
	v := C.WrapVec3GetX(pointer.h)
	return float32(v)
}

// SetX ...
func (pointer *Vec3) SetX(v float32) {
	vToC := C.float(v)
	C.WrapVec3SetX(pointer.h, vToC)
}

// GetY ...
func (pointer *Vec3) GetY() float32 {
	v := C.WrapVec3GetY(pointer.h)
	return float32(v)
}

// SetY ...
func (pointer *Vec3) SetY(v float32) {
	vToC := C.float(v)
	C.WrapVec3SetY(pointer.h, vToC)
}

// GetZ ...
func (pointer *Vec3) GetZ() float32 {
	v := C.WrapVec3GetZ(pointer.h)
	return float32(v)
}

// SetZ ...
func (pointer *Vec3) SetZ(v float32) {
	vToC := C.float(v)
	C.WrapVec3SetZ(pointer.h, vToC)
}

// NewVec3 3-dimensional vector.
func NewVec3() *Vec3 {
	retval := C.WrapConstructorVec3()
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// NewVec3WithXYZ 3-dimensional vector.
func NewVec3WithXYZ(x float32, y float32, z float32) *Vec3 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapConstructorVec3WithXYZ(xToC, yToC, zToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// NewVec3WithV 3-dimensional vector.
func NewVec3WithV(v *Vec2) *Vec3 {
	vToC := v.h
	retval := C.WrapConstructorVec3WithV(vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// NewVec3WithIVec2V 3-dimensional vector.
func NewVec3WithIVec2V(v *IVec2) *Vec3 {
	vToC := v.h
	retval := C.WrapConstructorVec3WithIVec2V(vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// NewVec3WithVec3V 3-dimensional vector.
func NewVec3WithVec3V(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapConstructorVec3WithVec3V(vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// NewVec3WithVec4V 3-dimensional vector.
func NewVec3WithVec4V(v *Vec4) *Vec3 {
	vToC := v.h
	retval := C.WrapConstructorVec3WithVec4V(vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Vec3) Free() {
	C.WrapVec3Free(pointer.h)
}

// IsNil ...
func (pointer *Vec3) IsNil() bool {
	return pointer.h == C.WrapVec3(nil)
}

// Add ...
func (pointer *Vec3) Add(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapAddVec3(pointer.h, vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// AddWithK ...
func (pointer *Vec3) AddWithK(k float32) *Vec3 {
	kToC := C.float(k)
	retval := C.WrapAddVec3WithK(pointer.h, kToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Sub ...
func (pointer *Vec3) Sub(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapSubVec3(pointer.h, vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// SubWithK ...
func (pointer *Vec3) SubWithK(k float32) *Vec3 {
	kToC := C.float(k)
	retval := C.WrapSubVec3WithK(pointer.h, kToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Div ...
func (pointer *Vec3) Div(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapDivVec3(pointer.h, vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// DivWithK ...
func (pointer *Vec3) DivWithK(k float32) *Vec3 {
	kToC := C.float(k)
	retval := C.WrapDivVec3WithK(pointer.h, kToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Mul ...
func (pointer *Vec3) Mul(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapMulVec3(pointer.h, vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// MulWithK ...
func (pointer *Vec3) MulWithK(k float32) *Vec3 {
	kToC := C.float(k)
	retval := C.WrapMulVec3WithK(pointer.h, kToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// InplaceAdd ...
func (pointer *Vec3) InplaceAdd(v *Vec3) {
	vToC := v.h
	C.WrapInplaceAddVec3(pointer.h, vToC)
}

// InplaceAddWithK ...
func (pointer *Vec3) InplaceAddWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceAddVec3WithK(pointer.h, kToC)
}

// InplaceSub ...
func (pointer *Vec3) InplaceSub(v *Vec3) {
	vToC := v.h
	C.WrapInplaceSubVec3(pointer.h, vToC)
}

// InplaceSubWithK ...
func (pointer *Vec3) InplaceSubWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceSubVec3WithK(pointer.h, kToC)
}

// InplaceMul ...
func (pointer *Vec3) InplaceMul(v *Vec3) {
	vToC := v.h
	C.WrapInplaceMulVec3(pointer.h, vToC)
}

// InplaceMulWithK ...
func (pointer *Vec3) InplaceMulWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceMulVec3WithK(pointer.h, kToC)
}

// InplaceDiv ...
func (pointer *Vec3) InplaceDiv(v *Vec3) {
	vToC := v.h
	C.WrapInplaceDivVec3(pointer.h, vToC)
}

// InplaceDivWithK ...
func (pointer *Vec3) InplaceDivWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceDivVec3WithK(pointer.h, kToC)
}

// Eq ...
func (pointer *Vec3) Eq(v *Vec3) bool {
	vToC := v.h
	retval := C.WrapEqVec3(pointer.h, vToC)
	return bool(retval)
}

// Ne ...
func (pointer *Vec3) Ne(v *Vec3) bool {
	vToC := v.h
	retval := C.WrapNeVec3(pointer.h, vToC)
	return bool(retval)
}

// Set ...
func (pointer *Vec3) Set(x float32, y float32, z float32) {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	C.WrapSetVec3(pointer.h, xToC, yToC, zToC)
}

// Vec4  4-dimensional vector.
type Vec4 struct {
	h C.WrapVec4
}

// GetX ...
func (pointer *Vec4) GetX() float32 {
	v := C.WrapVec4GetX(pointer.h)
	return float32(v)
}

// SetX ...
func (pointer *Vec4) SetX(v float32) {
	vToC := C.float(v)
	C.WrapVec4SetX(pointer.h, vToC)
}

// GetY ...
func (pointer *Vec4) GetY() float32 {
	v := C.WrapVec4GetY(pointer.h)
	return float32(v)
}

// SetY ...
func (pointer *Vec4) SetY(v float32) {
	vToC := C.float(v)
	C.WrapVec4SetY(pointer.h, vToC)
}

// GetZ ...
func (pointer *Vec4) GetZ() float32 {
	v := C.WrapVec4GetZ(pointer.h)
	return float32(v)
}

// SetZ ...
func (pointer *Vec4) SetZ(v float32) {
	vToC := C.float(v)
	C.WrapVec4SetZ(pointer.h, vToC)
}

// GetW ...
func (pointer *Vec4) GetW() float32 {
	v := C.WrapVec4GetW(pointer.h)
	return float32(v)
}

// SetW ...
func (pointer *Vec4) SetW(v float32) {
	vToC := C.float(v)
	C.WrapVec4SetW(pointer.h, vToC)
}

// NewVec4 4-dimensional vector.
func NewVec4() *Vec4 {
	retval := C.WrapConstructorVec4()
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// NewVec4WithXYZ 4-dimensional vector.
func NewVec4WithXYZ(x float32, y float32, z float32) *Vec4 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapConstructorVec4WithXYZ(xToC, yToC, zToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// NewVec4WithXYZW 4-dimensional vector.
func NewVec4WithXYZW(x float32, y float32, z float32, w float32) *Vec4 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	wToC := C.float(w)
	retval := C.WrapConstructorVec4WithXYZW(xToC, yToC, zToC, wToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// NewVec4WithV 4-dimensional vector.
func NewVec4WithV(v *Vec2) *Vec4 {
	vToC := v.h
	retval := C.WrapConstructorVec4WithV(vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// NewVec4WithIVec2V 4-dimensional vector.
func NewVec4WithIVec2V(v *IVec2) *Vec4 {
	vToC := v.h
	retval := C.WrapConstructorVec4WithIVec2V(vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// NewVec4WithVec3V 4-dimensional vector.
func NewVec4WithVec3V(v *Vec3) *Vec4 {
	vToC := v.h
	retval := C.WrapConstructorVec4WithVec3V(vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// NewVec4WithVec4V 4-dimensional vector.
func NewVec4WithVec4V(v *Vec4) *Vec4 {
	vToC := v.h
	retval := C.WrapConstructorVec4WithVec4V(vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Vec4) Free() {
	C.WrapVec4Free(pointer.h)
}

// IsNil ...
func (pointer *Vec4) IsNil() bool {
	return pointer.h == C.WrapVec4(nil)
}

// Add ...
func (pointer *Vec4) Add(v *Vec4) *Vec4 {
	vToC := v.h
	retval := C.WrapAddVec4(pointer.h, vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// AddWithK ...
func (pointer *Vec4) AddWithK(k float32) *Vec4 {
	kToC := C.float(k)
	retval := C.WrapAddVec4WithK(pointer.h, kToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// Sub ...
func (pointer *Vec4) Sub(v *Vec4) *Vec4 {
	vToC := v.h
	retval := C.WrapSubVec4(pointer.h, vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// SubWithK ...
func (pointer *Vec4) SubWithK(k float32) *Vec4 {
	kToC := C.float(k)
	retval := C.WrapSubVec4WithK(pointer.h, kToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// Div ...
func (pointer *Vec4) Div(v *Vec4) *Vec4 {
	vToC := v.h
	retval := C.WrapDivVec4(pointer.h, vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// DivWithK ...
func (pointer *Vec4) DivWithK(k float32) *Vec4 {
	kToC := C.float(k)
	retval := C.WrapDivVec4WithK(pointer.h, kToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// Mul ...
func (pointer *Vec4) Mul(v *Vec4) *Vec4 {
	vToC := v.h
	retval := C.WrapMulVec4(pointer.h, vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// MulWithK ...
func (pointer *Vec4) MulWithK(k float32) *Vec4 {
	kToC := C.float(k)
	retval := C.WrapMulVec4WithK(pointer.h, kToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// InplaceAdd ...
func (pointer *Vec4) InplaceAdd(v *Vec4) {
	vToC := v.h
	C.WrapInplaceAddVec4(pointer.h, vToC)
}

// InplaceAddWithK ...
func (pointer *Vec4) InplaceAddWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceAddVec4WithK(pointer.h, kToC)
}

// InplaceSub ...
func (pointer *Vec4) InplaceSub(v *Vec4) {
	vToC := v.h
	C.WrapInplaceSubVec4(pointer.h, vToC)
}

// InplaceSubWithK ...
func (pointer *Vec4) InplaceSubWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceSubVec4WithK(pointer.h, kToC)
}

// InplaceMul ...
func (pointer *Vec4) InplaceMul(v *Vec4) {
	vToC := v.h
	C.WrapInplaceMulVec4(pointer.h, vToC)
}

// InplaceMulWithK ...
func (pointer *Vec4) InplaceMulWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceMulVec4WithK(pointer.h, kToC)
}

// InplaceDiv ...
func (pointer *Vec4) InplaceDiv(v *Vec4) {
	vToC := v.h
	C.WrapInplaceDivVec4(pointer.h, vToC)
}

// InplaceDivWithK ...
func (pointer *Vec4) InplaceDivWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceDivVec4WithK(pointer.h, kToC)
}

// Set ...
func (pointer *Vec4) Set(x float32, y float32, z float32) {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	C.WrapSetVec4(pointer.h, xToC, yToC, zToC)
}

// SetWithW ...
func (pointer *Vec4) SetWithW(x float32, y float32, z float32, w float32) {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	wToC := C.float(w)
	C.WrapSetVec4WithW(pointer.h, xToC, yToC, zToC, wToC)
}

// Mat3  A 3x3 matrix used to store rotation.
type Mat3 struct {
	h C.WrapMat3
}

// GetZero ...
func (pointer *Mat3) GetZero() *Mat3 {
	v := C.WrapMat3GetZero()
	vGO := &Mat3{h: v}
	return vGO
}

// Mat3GetZero ...
func Mat3GetZero() *Mat3 {
	v := C.WrapMat3GetZero()
	vGO := &Mat3{h: v}
	return vGO
}

// GetIdentity ...
func (pointer *Mat3) GetIdentity() *Mat3 {
	v := C.WrapMat3GetIdentity()
	vGO := &Mat3{h: v}
	return vGO
}

// Mat3GetIdentity ...
func Mat3GetIdentity() *Mat3 {
	v := C.WrapMat3GetIdentity()
	vGO := &Mat3{h: v}
	return vGO
}

// NewMat3 A 3x3 matrix used to store rotation.
func NewMat3() *Mat3 {
	retval := C.WrapConstructorMat3()
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// NewMat3WithM A 3x3 matrix used to store rotation.
func NewMat3WithM(m *Mat4) *Mat3 {
	mToC := m.h
	retval := C.WrapConstructorMat3WithM(mToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// NewMat3WithXYZ A 3x3 matrix used to store rotation.
func NewMat3WithXYZ(x *Vec3, y *Vec3, z *Vec3) *Mat3 {
	xToC := x.h
	yToC := y.h
	zToC := z.h
	retval := C.WrapConstructorMat3WithXYZ(xToC, yToC, zToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Mat3) Free() {
	C.WrapMat3Free(pointer.h)
}

// IsNil ...
func (pointer *Mat3) IsNil() bool {
	return pointer.h == C.WrapMat3(nil)
}

// Add ...
func (pointer *Mat3) Add(m *Mat3) *Mat3 {
	mToC := m.h
	retval := C.WrapAddMat3(pointer.h, mToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// Sub ...
func (pointer *Mat3) Sub(m *Mat3) *Mat3 {
	mToC := m.h
	retval := C.WrapSubMat3(pointer.h, mToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// Mul ...
func (pointer *Mat3) Mul(v float32) *Mat3 {
	vToC := C.float(v)
	retval := C.WrapMulMat3(pointer.h, vToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// MulWithV ...
func (pointer *Mat3) MulWithV(v *Vec2) *Vec2 {
	vToC := v.h
	retval := C.WrapMulMat3WithV(pointer.h, vToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// MulWithVec3V ...
func (pointer *Mat3) MulWithVec3V(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapMulMat3WithVec3V(pointer.h, vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// MulWithVec4V ...
func (pointer *Mat3) MulWithVec4V(v *Vec4) *Vec4 {
	vToC := v.h
	retval := C.WrapMulMat3WithVec4V(pointer.h, vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// MulWithM ...
func (pointer *Mat3) MulWithM(m *Mat3) *Mat3 {
	mToC := m.h
	retval := C.WrapMulMat3WithM(pointer.h, mToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// InplaceAdd ...
func (pointer *Mat3) InplaceAdd(m *Mat3) {
	mToC := m.h
	C.WrapInplaceAddMat3(pointer.h, mToC)
}

// InplaceSub ...
func (pointer *Mat3) InplaceSub(m *Mat3) {
	mToC := m.h
	C.WrapInplaceSubMat3(pointer.h, mToC)
}

// InplaceMul ...
func (pointer *Mat3) InplaceMul(k float32) {
	kToC := C.float(k)
	C.WrapInplaceMulMat3(pointer.h, kToC)
}

// InplaceMulWithM ...
func (pointer *Mat3) InplaceMulWithM(m *Mat3) {
	mToC := m.h
	C.WrapInplaceMulMat3WithM(pointer.h, mToC)
}

// Eq ...
func (pointer *Mat3) Eq(m *Mat3) bool {
	mToC := m.h
	retval := C.WrapEqMat3(pointer.h, mToC)
	return bool(retval)
}

// Ne ...
func (pointer *Mat3) Ne(m *Mat3) bool {
	mToC := m.h
	retval := C.WrapNeMat3(pointer.h, mToC)
	return bool(retval)
}

// Mat4  A 3x4 matrix used to store complete transformation including rotation, scale and position.
type Mat4 struct {
	h C.WrapMat4
}

// GetZero ...
func (pointer *Mat4) GetZero() *Mat4 {
	v := C.WrapMat4GetZero()
	vGO := &Mat4{h: v}
	return vGO
}

// Mat4GetZero ...
func Mat4GetZero() *Mat4 {
	v := C.WrapMat4GetZero()
	vGO := &Mat4{h: v}
	return vGO
}

// GetIdentity ...
func (pointer *Mat4) GetIdentity() *Mat4 {
	v := C.WrapMat4GetIdentity()
	vGO := &Mat4{h: v}
	return vGO
}

// Mat4GetIdentity ...
func Mat4GetIdentity() *Mat4 {
	v := C.WrapMat4GetIdentity()
	vGO := &Mat4{h: v}
	return vGO
}

// NewMat4 A 3x4 matrix used to store complete transformation including rotation, scale and position.
func NewMat4() *Mat4 {
	retval := C.WrapConstructorMat4()
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// NewMat4WithM A 3x4 matrix used to store complete transformation including rotation, scale and position.
func NewMat4WithM(m *Mat4) *Mat4 {
	mToC := m.h
	retval := C.WrapConstructorMat4WithM(mToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// NewMat4WithM00M10M20M01M11M21M02M12M22M03M13M23 A 3x4 matrix used to store complete transformation including rotation, scale and position.
func NewMat4WithM00M10M20M01M11M21M02M12M22M03M13M23(m00 float32, m10 float32, m20 float32, m01 float32, m11 float32, m21 float32, m02 float32, m12 float32, m22 float32, m03 float32, m13 float32, m23 float32) *Mat4 {
	m00ToC := C.float(m00)
	m10ToC := C.float(m10)
	m20ToC := C.float(m20)
	m01ToC := C.float(m01)
	m11ToC := C.float(m11)
	m21ToC := C.float(m21)
	m02ToC := C.float(m02)
	m12ToC := C.float(m12)
	m22ToC := C.float(m22)
	m03ToC := C.float(m03)
	m13ToC := C.float(m13)
	m23ToC := C.float(m23)
	retval := C.WrapConstructorMat4WithM00M10M20M01M11M21M02M12M22M03M13M23(m00ToC, m10ToC, m20ToC, m01ToC, m11ToC, m21ToC, m02ToC, m12ToC, m22ToC, m03ToC, m13ToC, m23ToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// NewMat4WithMat3M A 3x4 matrix used to store complete transformation including rotation, scale and position.
func NewMat4WithMat3M(m *Mat3) *Mat4 {
	mToC := m.h
	retval := C.WrapConstructorMat4WithMat3M(mToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Mat4) Free() {
	C.WrapMat4Free(pointer.h)
}

// IsNil ...
func (pointer *Mat4) IsNil() bool {
	return pointer.h == C.WrapMat4(nil)
}

// Add ...
func (pointer *Mat4) Add(m *Mat4) *Mat4 {
	mToC := m.h
	retval := C.WrapAddMat4(pointer.h, mToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Sub ...
func (pointer *Mat4) Sub(m *Mat4) *Mat4 {
	mToC := m.h
	retval := C.WrapSubMat4(pointer.h, mToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Mul ...
func (pointer *Mat4) Mul(v float32) *Mat4 {
	vToC := C.float(v)
	retval := C.WrapMulMat4(pointer.h, vToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// MulWithM ...
func (pointer *Mat4) MulWithM(m *Mat4) *Mat4 {
	mToC := m.h
	retval := C.WrapMulMat4WithM(pointer.h, mToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// MulWithV ...
func (pointer *Mat4) MulWithV(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapMulMat4WithV(pointer.h, vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// MulWithVec4V ...
func (pointer *Mat4) MulWithVec4V(v *Vec4) *Vec4 {
	vToC := v.h
	retval := C.WrapMulMat4WithVec4V(pointer.h, vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// MulWithMat44M ...
func (pointer *Mat4) MulWithMat44M(m *Mat44) *Mat44 {
	mToC := m.h
	retval := C.WrapMulMat4WithMat44M(pointer.h, mToC)
	retvalGO := &Mat44{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat44) {
		C.WrapMat44Free(cleanval.h)
	})
	return retvalGO
}

// Eq ...
func (pointer *Mat4) Eq(m *Mat4) bool {
	mToC := m.h
	retval := C.WrapEqMat4(pointer.h, mToC)
	return bool(retval)
}

// Ne ...
func (pointer *Mat4) Ne(m *Mat4) bool {
	mToC := m.h
	retval := C.WrapNeMat4(pointer.h, mToC)
	return bool(retval)
}

// Mat44  A 4x4 matrix used to store projection matrices.
type Mat44 struct {
	h C.WrapMat44
}

// GetZero ...
func (pointer *Mat44) GetZero() *Mat44 {
	v := C.WrapMat44GetZero()
	vGO := &Mat44{h: v}
	return vGO
}

// Mat44GetZero ...
func Mat44GetZero() *Mat44 {
	v := C.WrapMat44GetZero()
	vGO := &Mat44{h: v}
	return vGO
}

// GetIdentity ...
func (pointer *Mat44) GetIdentity() *Mat44 {
	v := C.WrapMat44GetIdentity()
	vGO := &Mat44{h: v}
	return vGO
}

// Mat44GetIdentity ...
func Mat44GetIdentity() *Mat44 {
	v := C.WrapMat44GetIdentity()
	vGO := &Mat44{h: v}
	return vGO
}

// NewMat44 A 4x4 matrix used to store projection matrices.
func NewMat44() *Mat44 {
	retval := C.WrapConstructorMat44()
	retvalGO := &Mat44{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat44) {
		C.WrapMat44Free(cleanval.h)
	})
	return retvalGO
}

// NewMat44WithM00M10M20M30M01M11M21M31M02M12M22M32M03M13M23M33 A 4x4 matrix used to store projection matrices.
func NewMat44WithM00M10M20M30M01M11M21M31M02M12M22M32M03M13M23M33(m00 float32, m10 float32, m20 float32, m30 float32, m01 float32, m11 float32, m21 float32, m31 float32, m02 float32, m12 float32, m22 float32, m32 float32, m03 float32, m13 float32, m23 float32, m33 float32) *Mat44 {
	m00ToC := C.float(m00)
	m10ToC := C.float(m10)
	m20ToC := C.float(m20)
	m30ToC := C.float(m30)
	m01ToC := C.float(m01)
	m11ToC := C.float(m11)
	m21ToC := C.float(m21)
	m31ToC := C.float(m31)
	m02ToC := C.float(m02)
	m12ToC := C.float(m12)
	m22ToC := C.float(m22)
	m32ToC := C.float(m32)
	m03ToC := C.float(m03)
	m13ToC := C.float(m13)
	m23ToC := C.float(m23)
	m33ToC := C.float(m33)
	retval := C.WrapConstructorMat44WithM00M10M20M30M01M11M21M31M02M12M22M32M03M13M23M33(m00ToC, m10ToC, m20ToC, m30ToC, m01ToC, m11ToC, m21ToC, m31ToC, m02ToC, m12ToC, m22ToC, m32ToC, m03ToC, m13ToC, m23ToC, m33ToC)
	retvalGO := &Mat44{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat44) {
		C.WrapMat44Free(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Mat44) Free() {
	C.WrapMat44Free(pointer.h)
}

// IsNil ...
func (pointer *Mat44) IsNil() bool {
	return pointer.h == C.WrapMat44(nil)
}

// Mul ...
func (pointer *Mat44) Mul(m *Mat4) *Mat44 {
	mToC := m.h
	retval := C.WrapMulMat44(pointer.h, mToC)
	retvalGO := &Mat44{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat44) {
		C.WrapMat44Free(cleanval.h)
	})
	return retvalGO
}

// MulWithM ...
func (pointer *Mat44) MulWithM(m *Mat44) *Mat44 {
	mToC := m.h
	retval := C.WrapMulMat44WithM(pointer.h, mToC)
	retvalGO := &Mat44{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat44) {
		C.WrapMat44Free(cleanval.h)
	})
	return retvalGO
}

// MulWithV ...
func (pointer *Mat44) MulWithV(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapMulMat44WithV(pointer.h, vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// MulWithVec4V ...
func (pointer *Mat44) MulWithVec4V(v *Vec4) *Vec4 {
	vToC := v.h
	retval := C.WrapMulMat44WithVec4V(pointer.h, vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// Quaternion  Quaternion can be used to represent a 3d rotation. It provides a more compact representation of the rotation than [harfang.Mat3] and can efficiently and correctly interpolate (see [harfang.Slerp]) between two rotations.
type Quaternion struct {
	h C.WrapQuaternion
}

// GetX ...
func (pointer *Quaternion) GetX() float32 {
	v := C.WrapQuaternionGetX(pointer.h)
	return float32(v)
}

// SetX ...
func (pointer *Quaternion) SetX(v float32) {
	vToC := C.float(v)
	C.WrapQuaternionSetX(pointer.h, vToC)
}

// GetY ...
func (pointer *Quaternion) GetY() float32 {
	v := C.WrapQuaternionGetY(pointer.h)
	return float32(v)
}

// SetY ...
func (pointer *Quaternion) SetY(v float32) {
	vToC := C.float(v)
	C.WrapQuaternionSetY(pointer.h, vToC)
}

// GetZ ...
func (pointer *Quaternion) GetZ() float32 {
	v := C.WrapQuaternionGetZ(pointer.h)
	return float32(v)
}

// SetZ ...
func (pointer *Quaternion) SetZ(v float32) {
	vToC := C.float(v)
	C.WrapQuaternionSetZ(pointer.h, vToC)
}

// GetW ...
func (pointer *Quaternion) GetW() float32 {
	v := C.WrapQuaternionGetW(pointer.h)
	return float32(v)
}

// SetW ...
func (pointer *Quaternion) SetW(v float32) {
	vToC := C.float(v)
	C.WrapQuaternionSetW(pointer.h, vToC)
}

// NewQuaternion Quaternion can be used to represent a 3d rotation. It provides a more compact representation of the rotation than [harfang.Mat3] and can efficiently and correctly interpolate (see [harfang.Slerp]) between two rotations.
func NewQuaternion() *Quaternion {
	retval := C.WrapConstructorQuaternion()
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// NewQuaternionWithXYZW Quaternion can be used to represent a 3d rotation. It provides a more compact representation of the rotation than [harfang.Mat3] and can efficiently and correctly interpolate (see [harfang.Slerp]) between two rotations.
func NewQuaternionWithXYZW(x float32, y float32, z float32, w float32) *Quaternion {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	wToC := C.float(w)
	retval := C.WrapConstructorQuaternionWithXYZW(xToC, yToC, zToC, wToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// NewQuaternionWithQ Quaternion can be used to represent a 3d rotation. It provides a more compact representation of the rotation than [harfang.Mat3] and can efficiently and correctly interpolate (see [harfang.Slerp]) between two rotations.
func NewQuaternionWithQ(q *Quaternion) *Quaternion {
	qToC := q.h
	retval := C.WrapConstructorQuaternionWithQ(qToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Quaternion) Free() {
	C.WrapQuaternionFree(pointer.h)
}

// IsNil ...
func (pointer *Quaternion) IsNil() bool {
	return pointer.h == C.WrapQuaternion(nil)
}

// Add ...
func (pointer *Quaternion) Add(v float32) *Quaternion {
	vToC := C.float(v)
	retval := C.WrapAddQuaternion(pointer.h, vToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// AddWithQ ...
func (pointer *Quaternion) AddWithQ(q *Quaternion) *Quaternion {
	qToC := q.h
	retval := C.WrapAddQuaternionWithQ(pointer.h, qToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// Sub ...
func (pointer *Quaternion) Sub(v float32) *Quaternion {
	vToC := C.float(v)
	retval := C.WrapSubQuaternion(pointer.h, vToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// SubWithQ ...
func (pointer *Quaternion) SubWithQ(q *Quaternion) *Quaternion {
	qToC := q.h
	retval := C.WrapSubQuaternionWithQ(pointer.h, qToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// Mul ...
func (pointer *Quaternion) Mul(v float32) *Quaternion {
	vToC := C.float(v)
	retval := C.WrapMulQuaternion(pointer.h, vToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// MulWithQ ...
func (pointer *Quaternion) MulWithQ(q *Quaternion) *Quaternion {
	qToC := q.h
	retval := C.WrapMulQuaternionWithQ(pointer.h, qToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// Div ...
func (pointer *Quaternion) Div(v float32) *Quaternion {
	vToC := C.float(v)
	retval := C.WrapDivQuaternion(pointer.h, vToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// InplaceAdd ...
func (pointer *Quaternion) InplaceAdd(v float32) {
	vToC := C.float(v)
	C.WrapInplaceAddQuaternion(pointer.h, vToC)
}

// InplaceAddWithQ ...
func (pointer *Quaternion) InplaceAddWithQ(q *Quaternion) {
	qToC := q.h
	C.WrapInplaceAddQuaternionWithQ(pointer.h, qToC)
}

// InplaceSub ...
func (pointer *Quaternion) InplaceSub(v float32) {
	vToC := C.float(v)
	C.WrapInplaceSubQuaternion(pointer.h, vToC)
}

// InplaceSubWithQ ...
func (pointer *Quaternion) InplaceSubWithQ(q *Quaternion) {
	qToC := q.h
	C.WrapInplaceSubQuaternionWithQ(pointer.h, qToC)
}

// InplaceMul ...
func (pointer *Quaternion) InplaceMul(v float32) {
	vToC := C.float(v)
	C.WrapInplaceMulQuaternion(pointer.h, vToC)
}

// InplaceMulWithQ ...
func (pointer *Quaternion) InplaceMulWithQ(q *Quaternion) {
	qToC := q.h
	C.WrapInplaceMulQuaternionWithQ(pointer.h, qToC)
}

// InplaceDiv ...
func (pointer *Quaternion) InplaceDiv(v float32) {
	vToC := C.float(v)
	C.WrapInplaceDivQuaternion(pointer.h, vToC)
}

// MinMax  3D bounding volume defined by a minimum and maximum position.
type MinMax struct {
	h C.WrapMinMax
}

// GetMn ...
func (pointer *MinMax) GetMn() *Vec3 {
	v := C.WrapMinMaxGetMn(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetMn ...
func (pointer *MinMax) SetMn(v *Vec3) {
	vToC := v.h
	C.WrapMinMaxSetMn(pointer.h, vToC)
}

// GetMx ...
func (pointer *MinMax) GetMx() *Vec3 {
	v := C.WrapMinMaxGetMx(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetMx ...
func (pointer *MinMax) SetMx(v *Vec3) {
	vToC := v.h
	C.WrapMinMaxSetMx(pointer.h, vToC)
}

// NewMinMax 3D bounding volume defined by a minimum and maximum position.
func NewMinMax() *MinMax {
	retval := C.WrapConstructorMinMax()
	retvalGO := &MinMax{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MinMax) {
		C.WrapMinMaxFree(cleanval.h)
	})
	return retvalGO
}

// NewMinMaxWithMinMax 3D bounding volume defined by a minimum and maximum position.
func NewMinMaxWithMinMax(min *Vec3, max *Vec3) *MinMax {
	minToC := min.h
	maxToC := max.h
	retval := C.WrapConstructorMinMaxWithMinMax(minToC, maxToC)
	retvalGO := &MinMax{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MinMax) {
		C.WrapMinMaxFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *MinMax) Free() {
	C.WrapMinMaxFree(pointer.h)
}

// IsNil ...
func (pointer *MinMax) IsNil() bool {
	return pointer.h == C.WrapMinMax(nil)
}

// Mul ...
func (pointer *MinMax) Mul(m *Mat4) *MinMax {
	mToC := m.h
	retval := C.WrapMulMinMax(pointer.h, mToC)
	retvalGO := &MinMax{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MinMax) {
		C.WrapMinMaxFree(cleanval.h)
	})
	return retvalGO
}

// Eq ...
func (pointer *MinMax) Eq(minmax *MinMax) bool {
	minmaxToC := minmax.h
	retval := C.WrapEqMinMax(pointer.h, minmaxToC)
	return bool(retval)
}

// Ne ...
func (pointer *MinMax) Ne(minmax *MinMax) bool {
	minmaxToC := minmax.h
	retval := C.WrapNeMinMax(pointer.h, minmaxToC)
	return bool(retval)
}

// Vec2  2-dimensional floating point vector.
type Vec2 struct {
	h C.WrapVec2
}

// GetZero ...
func (pointer *Vec2) GetZero() *Vec2 {
	v := C.WrapVec2GetZero()
	vGO := &Vec2{h: v}
	return vGO
}

// Vec2GetZero ...
func Vec2GetZero() *Vec2 {
	v := C.WrapVec2GetZero()
	vGO := &Vec2{h: v}
	return vGO
}

// GetOne ...
func (pointer *Vec2) GetOne() *Vec2 {
	v := C.WrapVec2GetOne()
	vGO := &Vec2{h: v}
	return vGO
}

// Vec2GetOne ...
func Vec2GetOne() *Vec2 {
	v := C.WrapVec2GetOne()
	vGO := &Vec2{h: v}
	return vGO
}

// GetX ...
func (pointer *Vec2) GetX() float32 {
	v := C.WrapVec2GetX(pointer.h)
	return float32(v)
}

// SetX ...
func (pointer *Vec2) SetX(v float32) {
	vToC := C.float(v)
	C.WrapVec2SetX(pointer.h, vToC)
}

// GetY ...
func (pointer *Vec2) GetY() float32 {
	v := C.WrapVec2GetY(pointer.h)
	return float32(v)
}

// SetY ...
func (pointer *Vec2) SetY(v float32) {
	vToC := C.float(v)
	C.WrapVec2SetY(pointer.h, vToC)
}

// NewVec2 2-dimensional floating point vector.
func NewVec2() *Vec2 {
	retval := C.WrapConstructorVec2()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// NewVec2WithXY 2-dimensional floating point vector.
func NewVec2WithXY(x float32, y float32) *Vec2 {
	xToC := C.float(x)
	yToC := C.float(y)
	retval := C.WrapConstructorVec2WithXY(xToC, yToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// NewVec2WithV 2-dimensional floating point vector.
func NewVec2WithV(v *Vec2) *Vec2 {
	vToC := v.h
	retval := C.WrapConstructorVec2WithV(vToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// NewVec2WithVec3V 2-dimensional floating point vector.
func NewVec2WithVec3V(v *Vec3) *Vec2 {
	vToC := v.h
	retval := C.WrapConstructorVec2WithVec3V(vToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// NewVec2WithVec4V 2-dimensional floating point vector.
func NewVec2WithVec4V(v *Vec4) *Vec2 {
	vToC := v.h
	retval := C.WrapConstructorVec2WithVec4V(vToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Vec2) Free() {
	C.WrapVec2Free(pointer.h)
}

// IsNil ...
func (pointer *Vec2) IsNil() bool {
	return pointer.h == C.WrapVec2(nil)
}

// Add ...
func (pointer *Vec2) Add(v *Vec2) *Vec2 {
	vToC := v.h
	retval := C.WrapAddVec2(pointer.h, vToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// AddWithK ...
func (pointer *Vec2) AddWithK(k float32) *Vec2 {
	kToC := C.float(k)
	retval := C.WrapAddVec2WithK(pointer.h, kToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// Sub ...
func (pointer *Vec2) Sub(v *Vec2) *Vec2 {
	vToC := v.h
	retval := C.WrapSubVec2(pointer.h, vToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// SubWithK ...
func (pointer *Vec2) SubWithK(k float32) *Vec2 {
	kToC := C.float(k)
	retval := C.WrapSubVec2WithK(pointer.h, kToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// Div ...
func (pointer *Vec2) Div(v *Vec2) *Vec2 {
	vToC := v.h
	retval := C.WrapDivVec2(pointer.h, vToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// DivWithK ...
func (pointer *Vec2) DivWithK(k float32) *Vec2 {
	kToC := C.float(k)
	retval := C.WrapDivVec2WithK(pointer.h, kToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// Mul ...
func (pointer *Vec2) Mul(v *Vec2) *Vec2 {
	vToC := v.h
	retval := C.WrapMulVec2(pointer.h, vToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// MulWithK ...
func (pointer *Vec2) MulWithK(k float32) *Vec2 {
	kToC := C.float(k)
	retval := C.WrapMulVec2WithK(pointer.h, kToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// MulWithM ...
func (pointer *Vec2) MulWithM(m *Mat3) *Vec2 {
	mToC := m.h
	retval := C.WrapMulVec2WithM(pointer.h, mToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// InplaceAdd ...
func (pointer *Vec2) InplaceAdd(v *Vec2) {
	vToC := v.h
	C.WrapInplaceAddVec2(pointer.h, vToC)
}

// InplaceAddWithK ...
func (pointer *Vec2) InplaceAddWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceAddVec2WithK(pointer.h, kToC)
}

// InplaceSub ...
func (pointer *Vec2) InplaceSub(v *Vec2) {
	vToC := v.h
	C.WrapInplaceSubVec2(pointer.h, vToC)
}

// InplaceSubWithK ...
func (pointer *Vec2) InplaceSubWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceSubVec2WithK(pointer.h, kToC)
}

// InplaceMul ...
func (pointer *Vec2) InplaceMul(v *Vec2) {
	vToC := v.h
	C.WrapInplaceMulVec2(pointer.h, vToC)
}

// InplaceMulWithK ...
func (pointer *Vec2) InplaceMulWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceMulVec2WithK(pointer.h, kToC)
}

// InplaceDiv ...
func (pointer *Vec2) InplaceDiv(v *Vec2) {
	vToC := v.h
	C.WrapInplaceDivVec2(pointer.h, vToC)
}

// InplaceDivWithK ...
func (pointer *Vec2) InplaceDivWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceDivVec2WithK(pointer.h, kToC)
}

// Set ...
func (pointer *Vec2) Set(x float32, y float32) {
	xToC := C.float(x)
	yToC := C.float(y)
	C.WrapSetVec2(pointer.h, xToC, yToC)
}

// GoSliceOfVec2 ...
type GoSliceOfVec2 []*Vec2

// Vec2List  ...
type Vec2List struct {
	h C.WrapVec2List
}

// Get ...
func (pointer *Vec2List) Get(id int) *Vec2 {
	v := C.WrapVec2ListGetOperator(pointer.h, C.int(id))
	vGO := &Vec2{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *Vec2List) Set(id int, v *Vec2) {
	vToC := v.h
	C.WrapVec2ListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *Vec2List) Len() int32 {
	return int32(C.WrapVec2ListLenOperator(pointer.h))
}

// NewVec2List ...
func NewVec2List() *Vec2List {
	retval := C.WrapConstructorVec2List()
	retvalGO := &Vec2List{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2List) {
		C.WrapVec2ListFree(cleanval.h)
	})
	return retvalGO
}

// NewVec2ListWithSequence ...
func NewVec2ListWithSequence(sequence GoSliceOfVec2) *Vec2List {
	var sequencePointer []C.WrapVec2
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapVec2)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorVec2ListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &Vec2List{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2List) {
		C.WrapVec2ListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Vec2List) Free() {
	C.WrapVec2ListFree(pointer.h)
}

// IsNil ...
func (pointer *Vec2List) IsNil() bool {
	return pointer.h == C.WrapVec2List(nil)
}

// Clear ...
func (pointer *Vec2List) Clear() {
	C.WrapClearVec2List(pointer.h)
}

// Reserve ...
func (pointer *Vec2List) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveVec2List(pointer.h, sizeToC)
}

// PushBack Adds an element to the end.
func (pointer *Vec2List) PushBack(v *Vec2) {
	vToC := v.h
	C.WrapPushBackVec2List(pointer.h, vToC)
}

// Size Returns the number of stored elements.
func (pointer *Vec2List) Size() int32 {
	retval := C.WrapSizeVec2List(pointer.h)
	return int32(retval)
}

// At Gets the element at the specified index.
func (pointer *Vec2List) At(idx int32) *Vec2 {
	idxToC := C.size_t(idx)
	retval := C.WrapAtVec2List(pointer.h, idxToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// IVec2  2-dimensional integer vector.
type IVec2 struct {
	h C.WrapIVec2
}

// GetZero ...
func (pointer *IVec2) GetZero() *IVec2 {
	v := C.WrapIVec2GetZero()
	vGO := &IVec2{h: v}
	return vGO
}

// IVec2GetZero ...
func IVec2GetZero() *IVec2 {
	v := C.WrapIVec2GetZero()
	vGO := &IVec2{h: v}
	return vGO
}

// GetOne ...
func (pointer *IVec2) GetOne() *IVec2 {
	v := C.WrapIVec2GetOne()
	vGO := &IVec2{h: v}
	return vGO
}

// IVec2GetOne ...
func IVec2GetOne() *IVec2 {
	v := C.WrapIVec2GetOne()
	vGO := &IVec2{h: v}
	return vGO
}

// GetX ...
func (pointer *IVec2) GetX() int32 {
	v := C.WrapIVec2GetX(pointer.h)
	return int32(v)
}

// SetX ...
func (pointer *IVec2) SetX(v int32) {
	vToC := C.int32_t(v)
	C.WrapIVec2SetX(pointer.h, vToC)
}

// GetY ...
func (pointer *IVec2) GetY() int32 {
	v := C.WrapIVec2GetY(pointer.h)
	return int32(v)
}

// SetY ...
func (pointer *IVec2) SetY(v int32) {
	vToC := C.int32_t(v)
	C.WrapIVec2SetY(pointer.h, vToC)
}

// NewIVec2 2-dimensional integer vector.
func NewIVec2() *IVec2 {
	retval := C.WrapConstructorIVec2()
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// NewIVec2WithXY 2-dimensional integer vector.
func NewIVec2WithXY(x int32, y int32) *IVec2 {
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	retval := C.WrapConstructorIVec2WithXY(xToC, yToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// NewIVec2WithV 2-dimensional integer vector.
func NewIVec2WithV(v *IVec2) *IVec2 {
	vToC := v.h
	retval := C.WrapConstructorIVec2WithV(vToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// NewIVec2WithVec3V 2-dimensional integer vector.
func NewIVec2WithVec3V(v *Vec3) *IVec2 {
	vToC := v.h
	retval := C.WrapConstructorIVec2WithVec3V(vToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// NewIVec2WithVec4V 2-dimensional integer vector.
func NewIVec2WithVec4V(v *Vec4) *IVec2 {
	vToC := v.h
	retval := C.WrapConstructorIVec2WithVec4V(vToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *IVec2) Free() {
	C.WrapIVec2Free(pointer.h)
}

// IsNil ...
func (pointer *IVec2) IsNil() bool {
	return pointer.h == C.WrapIVec2(nil)
}

// Add ...
func (pointer *IVec2) Add(v *IVec2) *IVec2 {
	vToC := v.h
	retval := C.WrapAddIVec2(pointer.h, vToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// AddWithK ...
func (pointer *IVec2) AddWithK(k int32) *IVec2 {
	kToC := C.int32_t(k)
	retval := C.WrapAddIVec2WithK(pointer.h, kToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// Sub ...
func (pointer *IVec2) Sub(v *IVec2) *IVec2 {
	vToC := v.h
	retval := C.WrapSubIVec2(pointer.h, vToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// SubWithK ...
func (pointer *IVec2) SubWithK(k int32) *IVec2 {
	kToC := C.int32_t(k)
	retval := C.WrapSubIVec2WithK(pointer.h, kToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// Div ...
func (pointer *IVec2) Div(v *IVec2) *IVec2 {
	vToC := v.h
	retval := C.WrapDivIVec2(pointer.h, vToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// DivWithK ...
func (pointer *IVec2) DivWithK(k int32) *IVec2 {
	kToC := C.int32_t(k)
	retval := C.WrapDivIVec2WithK(pointer.h, kToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// Mul ...
func (pointer *IVec2) Mul(v *IVec2) *IVec2 {
	vToC := v.h
	retval := C.WrapMulIVec2(pointer.h, vToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// MulWithK ...
func (pointer *IVec2) MulWithK(k int32) *IVec2 {
	kToC := C.int32_t(k)
	retval := C.WrapMulIVec2WithK(pointer.h, kToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// MulWithM ...
func (pointer *IVec2) MulWithM(m *Mat3) *IVec2 {
	mToC := m.h
	retval := C.WrapMulIVec2WithM(pointer.h, mToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// InplaceAdd ...
func (pointer *IVec2) InplaceAdd(v *IVec2) {
	vToC := v.h
	C.WrapInplaceAddIVec2(pointer.h, vToC)
}

// InplaceAddWithK ...
func (pointer *IVec2) InplaceAddWithK(k int32) {
	kToC := C.int32_t(k)
	C.WrapInplaceAddIVec2WithK(pointer.h, kToC)
}

// InplaceSub ...
func (pointer *IVec2) InplaceSub(v *IVec2) {
	vToC := v.h
	C.WrapInplaceSubIVec2(pointer.h, vToC)
}

// InplaceSubWithK ...
func (pointer *IVec2) InplaceSubWithK(k int32) {
	kToC := C.int32_t(k)
	C.WrapInplaceSubIVec2WithK(pointer.h, kToC)
}

// InplaceMul ...
func (pointer *IVec2) InplaceMul(v *IVec2) {
	vToC := v.h
	C.WrapInplaceMulIVec2(pointer.h, vToC)
}

// InplaceMulWithK ...
func (pointer *IVec2) InplaceMulWithK(k int32) {
	kToC := C.int32_t(k)
	C.WrapInplaceMulIVec2WithK(pointer.h, kToC)
}

// InplaceDiv ...
func (pointer *IVec2) InplaceDiv(v *IVec2) {
	vToC := v.h
	C.WrapInplaceDivIVec2(pointer.h, vToC)
}

// InplaceDivWithK ...
func (pointer *IVec2) InplaceDivWithK(k int32) {
	kToC := C.int32_t(k)
	C.WrapInplaceDivIVec2WithK(pointer.h, kToC)
}

// Set ...
func (pointer *IVec2) Set(x int32, y int32) {
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	C.WrapSetIVec2(pointer.h, xToC, yToC)
}

// GoSliceOfiVec2 ...
type GoSliceOfiVec2 []*IVec2

// IVec2List  ...
type IVec2List struct {
	h C.WrapIVec2List
}

// Get ...
func (pointer *IVec2List) Get(id int) *IVec2 {
	v := C.WrapIVec2ListGetOperator(pointer.h, C.int(id))
	vGO := &IVec2{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *IVec2List) Set(id int, v *IVec2) {
	vToC := v.h
	C.WrapIVec2ListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *IVec2List) Len() int32 {
	return int32(C.WrapIVec2ListLenOperator(pointer.h))
}

// NewIVec2List ...
func NewIVec2List() *IVec2List {
	retval := C.WrapConstructorIVec2List()
	retvalGO := &IVec2List{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2List) {
		C.WrapIVec2ListFree(cleanval.h)
	})
	return retvalGO
}

// NewIVec2ListWithSequence ...
func NewIVec2ListWithSequence(sequence GoSliceOfiVec2) *IVec2List {
	var sequencePointer []C.WrapIVec2
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapIVec2)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorIVec2ListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &IVec2List{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2List) {
		C.WrapIVec2ListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *IVec2List) Free() {
	C.WrapIVec2ListFree(pointer.h)
}

// IsNil ...
func (pointer *IVec2List) IsNil() bool {
	return pointer.h == C.WrapIVec2List(nil)
}

// Clear ...
func (pointer *IVec2List) Clear() {
	C.WrapClearIVec2List(pointer.h)
}

// Reserve ...
func (pointer *IVec2List) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveIVec2List(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *IVec2List) PushBack(v *IVec2) {
	vToC := v.h
	C.WrapPushBackIVec2List(pointer.h, vToC)
}

// Size ...
func (pointer *IVec2List) Size() int32 {
	retval := C.WrapSizeIVec2List(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *IVec2List) At(idx int32) *IVec2 {
	idxToC := C.size_t(idx)
	retval := C.WrapAtIVec2List(pointer.h, idxToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// GoSliceOfVec4 ...
type GoSliceOfVec4 []*Vec4

// Vec4List  ...
type Vec4List struct {
	h C.WrapVec4List
}

// Get ...
func (pointer *Vec4List) Get(id int) *Vec4 {
	v := C.WrapVec4ListGetOperator(pointer.h, C.int(id))
	vGO := &Vec4{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *Vec4List) Set(id int, v *Vec4) {
	vToC := v.h
	C.WrapVec4ListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *Vec4List) Len() int32 {
	return int32(C.WrapVec4ListLenOperator(pointer.h))
}

// NewVec4List ...
func NewVec4List() *Vec4List {
	retval := C.WrapConstructorVec4List()
	retvalGO := &Vec4List{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4List) {
		C.WrapVec4ListFree(cleanval.h)
	})
	return retvalGO
}

// NewVec4ListWithSequence ...
func NewVec4ListWithSequence(sequence GoSliceOfVec4) *Vec4List {
	var sequencePointer []C.WrapVec4
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapVec4)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorVec4ListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &Vec4List{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4List) {
		C.WrapVec4ListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Vec4List) Free() {
	C.WrapVec4ListFree(pointer.h)
}

// IsNil ...
func (pointer *Vec4List) IsNil() bool {
	return pointer.h == C.WrapVec4List(nil)
}

// Clear ...
func (pointer *Vec4List) Clear() {
	C.WrapClearVec4List(pointer.h)
}

// Reserve ...
func (pointer *Vec4List) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveVec4List(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *Vec4List) PushBack(v *Vec4) {
	vToC := v.h
	C.WrapPushBackVec4List(pointer.h, vToC)
}

// Size ...
func (pointer *Vec4List) Size() int32 {
	retval := C.WrapSizeVec4List(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *Vec4List) At(idx int32) *Vec4 {
	idxToC := C.size_t(idx)
	retval := C.WrapAtVec4List(pointer.h, idxToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// GoSliceOfMat4 ...
type GoSliceOfMat4 []*Mat4

// Mat4List  ...
type Mat4List struct {
	h C.WrapMat4List
}

// Get ...
func (pointer *Mat4List) Get(id int) *Mat4 {
	v := C.WrapMat4ListGetOperator(pointer.h, C.int(id))
	vGO := &Mat4{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *Mat4List) Set(id int, v *Mat4) {
	vToC := v.h
	C.WrapMat4ListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *Mat4List) Len() int32 {
	return int32(C.WrapMat4ListLenOperator(pointer.h))
}

// NewMat4List ...
func NewMat4List() *Mat4List {
	retval := C.WrapConstructorMat4List()
	retvalGO := &Mat4List{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4List) {
		C.WrapMat4ListFree(cleanval.h)
	})
	return retvalGO
}

// NewMat4ListWithSequence ...
func NewMat4ListWithSequence(sequence GoSliceOfMat4) *Mat4List {
	var sequencePointer []C.WrapMat4
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapMat4)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorMat4ListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &Mat4List{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4List) {
		C.WrapMat4ListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Mat4List) Free() {
	C.WrapMat4ListFree(pointer.h)
}

// IsNil ...
func (pointer *Mat4List) IsNil() bool {
	return pointer.h == C.WrapMat4List(nil)
}

// Clear ...
func (pointer *Mat4List) Clear() {
	C.WrapClearMat4List(pointer.h)
}

// Reserve ...
func (pointer *Mat4List) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveMat4List(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *Mat4List) PushBack(v *Mat4) {
	vToC := v.h
	C.WrapPushBackMat4List(pointer.h, vToC)
}

// Size ...
func (pointer *Mat4List) Size() int32 {
	retval := C.WrapSizeMat4List(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *Mat4List) At(idx int32) *Mat4 {
	idxToC := C.size_t(idx)
	retval := C.WrapAtMat4List(pointer.h, idxToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// GoSliceOfVec3 ...
type GoSliceOfVec3 []*Vec3

// Vec3List  ...
type Vec3List struct {
	h C.WrapVec3List
}

// Get ...
func (pointer *Vec3List) Get(id int) *Vec3 {
	v := C.WrapVec3ListGetOperator(pointer.h, C.int(id))
	vGO := &Vec3{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *Vec3List) Set(id int, v *Vec3) {
	vToC := v.h
	C.WrapVec3ListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *Vec3List) Len() int32 {
	return int32(C.WrapVec3ListLenOperator(pointer.h))
}

// NewVec3List ...
func NewVec3List() *Vec3List {
	retval := C.WrapConstructorVec3List()
	retvalGO := &Vec3List{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3List) {
		C.WrapVec3ListFree(cleanval.h)
	})
	return retvalGO
}

// NewVec3ListWithSequence ...
func NewVec3ListWithSequence(sequence GoSliceOfVec3) *Vec3List {
	var sequencePointer []C.WrapVec3
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapVec3)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorVec3ListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &Vec3List{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3List) {
		C.WrapVec3ListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Vec3List) Free() {
	C.WrapVec3ListFree(pointer.h)
}

// IsNil ...
func (pointer *Vec3List) IsNil() bool {
	return pointer.h == C.WrapVec3List(nil)
}

// Clear ...
func (pointer *Vec3List) Clear() {
	C.WrapClearVec3List(pointer.h)
}

// Reserve ...
func (pointer *Vec3List) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveVec3List(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *Vec3List) PushBack(v *Vec3) {
	vToC := v.h
	C.WrapPushBackVec3List(pointer.h, vToC)
}

// Size ...
func (pointer *Vec3List) Size() int32 {
	retval := C.WrapSizeVec3List(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *Vec3List) At(idx int32) *Vec3 {
	idxToC := C.size_t(idx)
	retval := C.WrapAtVec3List(pointer.h, idxToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Rect  ...
type Rect struct {
	h C.WrapRect
}

// GetSx ...
func (pointer *Rect) GetSx() float32 {
	v := C.WrapRectGetSx(pointer.h)
	return float32(v)
}

// SetSx ...
func (pointer *Rect) SetSx(v float32) {
	vToC := C.float(v)
	C.WrapRectSetSx(pointer.h, vToC)
}

// GetSy ...
func (pointer *Rect) GetSy() float32 {
	v := C.WrapRectGetSy(pointer.h)
	return float32(v)
}

// SetSy ...
func (pointer *Rect) SetSy(v float32) {
	vToC := C.float(v)
	C.WrapRectSetSy(pointer.h, vToC)
}

// GetEx ...
func (pointer *Rect) GetEx() float32 {
	v := C.WrapRectGetEx(pointer.h)
	return float32(v)
}

// SetEx ...
func (pointer *Rect) SetEx(v float32) {
	vToC := C.float(v)
	C.WrapRectSetEx(pointer.h, vToC)
}

// GetEy ...
func (pointer *Rect) GetEy() float32 {
	v := C.WrapRectGetEy(pointer.h)
	return float32(v)
}

// SetEy ...
func (pointer *Rect) SetEy(v float32) {
	vToC := C.float(v)
	C.WrapRectSetEy(pointer.h, vToC)
}

// NewRect ...
func NewRect() *Rect {
	retval := C.WrapConstructorRect()
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// NewRectWithXY ...
func NewRectWithXY(x float32, y float32) *Rect {
	xToC := C.float(x)
	yToC := C.float(y)
	retval := C.WrapConstructorRectWithXY(xToC, yToC)
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// NewRectWithSxSyExEy ...
func NewRectWithSxSyExEy(sx float32, sy float32, ex float32, ey float32) *Rect {
	sxToC := C.float(sx)
	syToC := C.float(sy)
	exToC := C.float(ex)
	eyToC := C.float(ey)
	retval := C.WrapConstructorRectWithSxSyExEy(sxToC, syToC, exToC, eyToC)
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// NewRectWithRect ...
func NewRectWithRect(rect *Rect) *Rect {
	rectToC := rect.h
	retval := C.WrapConstructorRectWithRect(rectToC)
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Rect) Free() {
	C.WrapRectFree(pointer.h)
}

// IsNil ...
func (pointer *Rect) IsNil() bool {
	return pointer.h == C.WrapRect(nil)
}

// IntRect  ...
type IntRect struct {
	h C.WrapIntRect
}

// GetSx ...
func (pointer *IntRect) GetSx() int32 {
	v := C.WrapIntRectGetSx(pointer.h)
	return int32(v)
}

// SetSx ...
func (pointer *IntRect) SetSx(v int32) {
	vToC := C.int32_t(v)
	C.WrapIntRectSetSx(pointer.h, vToC)
}

// GetSy ...
func (pointer *IntRect) GetSy() int32 {
	v := C.WrapIntRectGetSy(pointer.h)
	return int32(v)
}

// SetSy ...
func (pointer *IntRect) SetSy(v int32) {
	vToC := C.int32_t(v)
	C.WrapIntRectSetSy(pointer.h, vToC)
}

// GetEx ...
func (pointer *IntRect) GetEx() int32 {
	v := C.WrapIntRectGetEx(pointer.h)
	return int32(v)
}

// SetEx ...
func (pointer *IntRect) SetEx(v int32) {
	vToC := C.int32_t(v)
	C.WrapIntRectSetEx(pointer.h, vToC)
}

// GetEy ...
func (pointer *IntRect) GetEy() int32 {
	v := C.WrapIntRectGetEy(pointer.h)
	return int32(v)
}

// SetEy ...
func (pointer *IntRect) SetEy(v int32) {
	vToC := C.int32_t(v)
	C.WrapIntRectSetEy(pointer.h, vToC)
}

// NewIntRect ...
func NewIntRect() *IntRect {
	retval := C.WrapConstructorIntRect()
	retvalGO := &IntRect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IntRect) {
		C.WrapIntRectFree(cleanval.h)
	})
	return retvalGO
}

// NewIntRectWithXY ...
func NewIntRectWithXY(x int32, y int32) *IntRect {
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	retval := C.WrapConstructorIntRectWithXY(xToC, yToC)
	retvalGO := &IntRect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IntRect) {
		C.WrapIntRectFree(cleanval.h)
	})
	return retvalGO
}

// NewIntRectWithSxSyExEy ...
func NewIntRectWithSxSyExEy(sx int32, sy int32, ex int32, ey int32) *IntRect {
	sxToC := C.int32_t(sx)
	syToC := C.int32_t(sy)
	exToC := C.int32_t(ex)
	eyToC := C.int32_t(ey)
	retval := C.WrapConstructorIntRectWithSxSyExEy(sxToC, syToC, exToC, eyToC)
	retvalGO := &IntRect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IntRect) {
		C.WrapIntRectFree(cleanval.h)
	})
	return retvalGO
}

// NewIntRectWithRect ...
func NewIntRectWithRect(rect *IntRect) *IntRect {
	rectToC := rect.h
	retval := C.WrapConstructorIntRectWithRect(rectToC)
	retvalGO := &IntRect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IntRect) {
		C.WrapIntRectFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *IntRect) Free() {
	C.WrapIntRectFree(pointer.h)
}

// IsNil ...
func (pointer *IntRect) IsNil() bool {
	return pointer.h == C.WrapIntRect(nil)
}

// Frustum  A view frustum, perspective or orthographic, holding the necessary information to perform culling queries. It can be used to test wether a volume is inside or outside the frustum it represents.
type Frustum struct {
	h C.WrapFrustum
}

// Free ...
func (pointer *Frustum) Free() {
	C.WrapFrustumFree(pointer.h)
}

// IsNil ...
func (pointer *Frustum) IsNil() bool {
	return pointer.h == C.WrapFrustum(nil)
}

// GetTop ...
func (pointer *Frustum) GetTop() *Vec4 {
	retval := C.WrapGetTopFrustum(pointer.h)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// SetTop ...
func (pointer *Frustum) SetTop(plane *Vec4) {
	planeToC := plane.h
	C.WrapSetTopFrustum(pointer.h, planeToC)
}

// GetBottom ...
func (pointer *Frustum) GetBottom() *Vec4 {
	retval := C.WrapGetBottomFrustum(pointer.h)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// SetBottom ...
func (pointer *Frustum) SetBottom(plane *Vec4) {
	planeToC := plane.h
	C.WrapSetBottomFrustum(pointer.h, planeToC)
}

// GetLeft ...
func (pointer *Frustum) GetLeft() *Vec4 {
	retval := C.WrapGetLeftFrustum(pointer.h)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// SetLeft ...
func (pointer *Frustum) SetLeft(plane *Vec4) {
	planeToC := plane.h
	C.WrapSetLeftFrustum(pointer.h, planeToC)
}

// GetRight ...
func (pointer *Frustum) GetRight() *Vec4 {
	retval := C.WrapGetRightFrustum(pointer.h)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// SetRight ...
func (pointer *Frustum) SetRight(plane *Vec4) {
	planeToC := plane.h
	C.WrapSetRightFrustum(pointer.h, planeToC)
}

// GetNear ...
func (pointer *Frustum) GetNear() *Vec4 {
	retval := C.WrapGetNearFrustum(pointer.h)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// SetNear ...
func (pointer *Frustum) SetNear(plane *Vec4) {
	planeToC := plane.h
	C.WrapSetNearFrustum(pointer.h, planeToC)
}

// GetFar ...
func (pointer *Frustum) GetFar() *Vec4 {
	retval := C.WrapGetFarFrustum(pointer.h)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// SetFar ...
func (pointer *Frustum) SetFar(plane *Vec4) {
	planeToC := plane.h
	C.WrapSetFarFrustum(pointer.h, planeToC)
}

// MonitorMode  ...
type MonitorMode struct {
	h C.WrapMonitorMode
}

// GetName ...
func (pointer *MonitorMode) GetName() string {
	v := C.WrapMonitorModeGetName(pointer.h)
	return C.GoString(v)
}

// SetName ...
func (pointer *MonitorMode) SetName(v string) {
	vToC, idFinvToC := wrapString(v)
	defer idFinvToC()
	C.WrapMonitorModeSetName(pointer.h, vToC)
}

// GetRect ...
func (pointer *MonitorMode) GetRect() *IntRect {
	v := C.WrapMonitorModeGetRect(pointer.h)
	vGO := &IntRect{h: v}
	return vGO
}

// SetRect ...
func (pointer *MonitorMode) SetRect(v *IntRect) {
	vToC := v.h
	C.WrapMonitorModeSetRect(pointer.h, vToC)
}

// GetFrequency ...
func (pointer *MonitorMode) GetFrequency() int32 {
	v := C.WrapMonitorModeGetFrequency(pointer.h)
	return int32(v)
}

// SetFrequency ...
func (pointer *MonitorMode) SetFrequency(v int32) {
	vToC := C.int32_t(v)
	C.WrapMonitorModeSetFrequency(pointer.h, vToC)
}

// GetRotation ...
func (pointer *MonitorMode) GetRotation() MonitorRotation {
	v := C.WrapMonitorModeGetRotation(pointer.h)
	return MonitorRotation(v)
}

// SetRotation ...
func (pointer *MonitorMode) SetRotation(v MonitorRotation) {
	vToC := C.uchar(v)
	C.WrapMonitorModeSetRotation(pointer.h, vToC)
}

// GetSupportedRotations ...
func (pointer *MonitorMode) GetSupportedRotations() uint8 {
	v := C.WrapMonitorModeGetSupportedRotations(pointer.h)
	return uint8(v)
}

// SetSupportedRotations ...
func (pointer *MonitorMode) SetSupportedRotations(v uint8) {
	vToC := C.uchar(v)
	C.WrapMonitorModeSetSupportedRotations(pointer.h, vToC)
}

// Free ...
func (pointer *MonitorMode) Free() {
	C.WrapMonitorModeFree(pointer.h)
}

// IsNil ...
func (pointer *MonitorMode) IsNil() bool {
	return pointer.h == C.WrapMonitorMode(nil)
}

// GoSliceOfMonitorMode ...
type GoSliceOfMonitorMode []*MonitorMode

// MonitorModeList  ...
type MonitorModeList struct {
	h C.WrapMonitorModeList
}

// Get ...
func (pointer *MonitorModeList) Get(id int) *MonitorMode {
	v := C.WrapMonitorModeListGetOperator(pointer.h, C.int(id))
	vGO := &MonitorMode{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *MonitorMode) {
		C.WrapMonitorModeFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *MonitorModeList) Set(id int, v *MonitorMode) {
	vToC := v.h
	C.WrapMonitorModeListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *MonitorModeList) Len() int32 {
	return int32(C.WrapMonitorModeListLenOperator(pointer.h))
}

// NewMonitorModeList ...
func NewMonitorModeList() *MonitorModeList {
	retval := C.WrapConstructorMonitorModeList()
	retvalGO := &MonitorModeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MonitorModeList) {
		C.WrapMonitorModeListFree(cleanval.h)
	})
	return retvalGO
}

// NewMonitorModeListWithSequence ...
func NewMonitorModeListWithSequence(sequence GoSliceOfMonitorMode) *MonitorModeList {
	var sequencePointer []C.WrapMonitorMode
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapMonitorMode)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorMonitorModeListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &MonitorModeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MonitorModeList) {
		C.WrapMonitorModeListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *MonitorModeList) Free() {
	C.WrapMonitorModeListFree(pointer.h)
}

// IsNil ...
func (pointer *MonitorModeList) IsNil() bool {
	return pointer.h == C.WrapMonitorModeList(nil)
}

// Clear ...
func (pointer *MonitorModeList) Clear() {
	C.WrapClearMonitorModeList(pointer.h)
}

// Reserve ...
func (pointer *MonitorModeList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveMonitorModeList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *MonitorModeList) PushBack(v *MonitorMode) {
	vToC := v.h
	C.WrapPushBackMonitorModeList(pointer.h, vToC)
}

// Size ...
func (pointer *MonitorModeList) Size() int32 {
	retval := C.WrapSizeMonitorModeList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *MonitorModeList) At(idx int32) *MonitorMode {
	idxToC := C.size_t(idx)
	retval := C.WrapAtMonitorModeList(pointer.h, idxToC)
	retvalGO := &MonitorMode{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MonitorMode) {
		C.WrapMonitorModeFree(cleanval.h)
	})
	return retvalGO
}

// Monitor  ...
type Monitor struct {
	h C.WrapMonitor
}

// Free ...
func (pointer *Monitor) Free() {
	C.WrapMonitorFree(pointer.h)
}

// IsNil ...
func (pointer *Monitor) IsNil() bool {
	return pointer.h == C.WrapMonitor(nil)
}

// GoSliceOfMonitor ...
type GoSliceOfMonitor []*Monitor

// MonitorList  ...
type MonitorList struct {
	h C.WrapMonitorList
}

// Get ...
func (pointer *MonitorList) Get(id int) *Monitor {
	v := C.WrapMonitorListGetOperator(pointer.h, C.int(id))
	var vGO *Monitor
	if v != nil {
		vGO = &Monitor{h: v}
	}
	return vGO
}

// Set ...
func (pointer *MonitorList) Set(id int, v *Monitor) {
	vToC := v.h
	C.WrapMonitorListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *MonitorList) Len() int32 {
	return int32(C.WrapMonitorListLenOperator(pointer.h))
}

// NewMonitorList ...
func NewMonitorList() *MonitorList {
	retval := C.WrapConstructorMonitorList()
	retvalGO := &MonitorList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MonitorList) {
		C.WrapMonitorListFree(cleanval.h)
	})
	return retvalGO
}

// NewMonitorListWithSequence ...
func NewMonitorListWithSequence(sequence GoSliceOfMonitor) *MonitorList {
	var sequencePointer []C.WrapMonitor
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapMonitor)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorMonitorListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &MonitorList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MonitorList) {
		C.WrapMonitorListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *MonitorList) Free() {
	C.WrapMonitorListFree(pointer.h)
}

// IsNil ...
func (pointer *MonitorList) IsNil() bool {
	return pointer.h == C.WrapMonitorList(nil)
}

// Clear ...
func (pointer *MonitorList) Clear() {
	C.WrapClearMonitorList(pointer.h)
}

// Reserve ...
func (pointer *MonitorList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveMonitorList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *MonitorList) PushBack(v *Monitor) {
	vToC := v.h
	C.WrapPushBackMonitorList(pointer.h, vToC)
}

// Size ...
func (pointer *MonitorList) Size() int32 {
	retval := C.WrapSizeMonitorList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *MonitorList) At(idx int32) *Monitor {
	idxToC := C.size_t(idx)
	retval := C.WrapAtMonitorList(pointer.h, idxToC)
	var retvalGO *Monitor
	if retval != nil {
		retvalGO = &Monitor{h: retval}
	}
	return retvalGO
}

// Window  Window object.
type Window struct {
	h C.WrapWindow
}

// Free ...
func (pointer *Window) Free() {
	C.WrapWindowFree(pointer.h)
}

// IsNil ...
func (pointer *Window) IsNil() bool {
	return pointer.h == C.WrapWindow(nil)
}

// Color  Four-component RGBA color object.
type Color struct {
	h C.WrapColor
}

// GetZero ...
func (pointer *Color) GetZero() *Color {
	v := C.WrapColorGetZero()
	vGO := &Color{h: v}
	return vGO
}

// ColorGetZero ...
func ColorGetZero() *Color {
	v := C.WrapColorGetZero()
	vGO := &Color{h: v}
	return vGO
}

// GetOne ...
func (pointer *Color) GetOne() *Color {
	v := C.WrapColorGetOne()
	vGO := &Color{h: v}
	return vGO
}

// ColorGetOne ...
func ColorGetOne() *Color {
	v := C.WrapColorGetOne()
	vGO := &Color{h: v}
	return vGO
}

// GetWhite ...
func (pointer *Color) GetWhite() *Color {
	v := C.WrapColorGetWhite()
	vGO := &Color{h: v}
	return vGO
}

// ColorGetWhite ...
func ColorGetWhite() *Color {
	v := C.WrapColorGetWhite()
	vGO := &Color{h: v}
	return vGO
}

// GetGrey ...
func (pointer *Color) GetGrey() *Color {
	v := C.WrapColorGetGrey()
	vGO := &Color{h: v}
	return vGO
}

// ColorGetGrey ...
func ColorGetGrey() *Color {
	v := C.WrapColorGetGrey()
	vGO := &Color{h: v}
	return vGO
}

// GetBlack ...
func (pointer *Color) GetBlack() *Color {
	v := C.WrapColorGetBlack()
	vGO := &Color{h: v}
	return vGO
}

// ColorGetBlack ...
func ColorGetBlack() *Color {
	v := C.WrapColorGetBlack()
	vGO := &Color{h: v}
	return vGO
}

// GetRed ...
func (pointer *Color) GetRed() *Color {
	v := C.WrapColorGetRed()
	vGO := &Color{h: v}
	return vGO
}

// ColorGetRed ...
func ColorGetRed() *Color {
	v := C.WrapColorGetRed()
	vGO := &Color{h: v}
	return vGO
}

// GetGreen ...
func (pointer *Color) GetGreen() *Color {
	v := C.WrapColorGetGreen()
	vGO := &Color{h: v}
	return vGO
}

// ColorGetGreen ...
func ColorGetGreen() *Color {
	v := C.WrapColorGetGreen()
	vGO := &Color{h: v}
	return vGO
}

// GetBlue ...
func (pointer *Color) GetBlue() *Color {
	v := C.WrapColorGetBlue()
	vGO := &Color{h: v}
	return vGO
}

// ColorGetBlue ...
func ColorGetBlue() *Color {
	v := C.WrapColorGetBlue()
	vGO := &Color{h: v}
	return vGO
}

// GetYellow ...
func (pointer *Color) GetYellow() *Color {
	v := C.WrapColorGetYellow()
	vGO := &Color{h: v}
	return vGO
}

// ColorGetYellow ...
func ColorGetYellow() *Color {
	v := C.WrapColorGetYellow()
	vGO := &Color{h: v}
	return vGO
}

// GetOrange ...
func (pointer *Color) GetOrange() *Color {
	v := C.WrapColorGetOrange()
	vGO := &Color{h: v}
	return vGO
}

// ColorGetOrange ...
func ColorGetOrange() *Color {
	v := C.WrapColorGetOrange()
	vGO := &Color{h: v}
	return vGO
}

// GetPurple ...
func (pointer *Color) GetPurple() *Color {
	v := C.WrapColorGetPurple()
	vGO := &Color{h: v}
	return vGO
}

// ColorGetPurple ...
func ColorGetPurple() *Color {
	v := C.WrapColorGetPurple()
	vGO := &Color{h: v}
	return vGO
}

// GetTransparent ...
func (pointer *Color) GetTransparent() *Color {
	v := C.WrapColorGetTransparent()
	vGO := &Color{h: v}
	return vGO
}

// ColorGetTransparent ...
func ColorGetTransparent() *Color {
	v := C.WrapColorGetTransparent()
	vGO := &Color{h: v}
	return vGO
}

// GetR ...
func (pointer *Color) GetR() float32 {
	v := C.WrapColorGetR(pointer.h)
	return float32(v)
}

// SetR ...
func (pointer *Color) SetR(v float32) {
	vToC := C.float(v)
	C.WrapColorSetR(pointer.h, vToC)
}

// GetG ...
func (pointer *Color) GetG() float32 {
	v := C.WrapColorGetG(pointer.h)
	return float32(v)
}

// SetG ...
func (pointer *Color) SetG(v float32) {
	vToC := C.float(v)
	C.WrapColorSetG(pointer.h, vToC)
}

// GetB ...
func (pointer *Color) GetB() float32 {
	v := C.WrapColorGetB(pointer.h)
	return float32(v)
}

// SetB ...
func (pointer *Color) SetB(v float32) {
	vToC := C.float(v)
	C.WrapColorSetB(pointer.h, vToC)
}

// GetA ...
func (pointer *Color) GetA() float32 {
	v := C.WrapColorGetA(pointer.h)
	return float32(v)
}

// SetA ...
func (pointer *Color) SetA(v float32) {
	vToC := C.float(v)
	C.WrapColorSetA(pointer.h, vToC)
}

// NewColor Four-component RGBA color object.
func NewColor() *Color {
	retval := C.WrapConstructorColor()
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// NewColorWithColor Four-component RGBA color object.
func NewColorWithColor(color *Color) *Color {
	colorToC := color.h
	retval := C.WrapConstructorColorWithColor(colorToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// NewColorWithRGB Four-component RGBA color object.
func NewColorWithRGB(r float32, g float32, b float32) *Color {
	rToC := C.float(r)
	gToC := C.float(g)
	bToC := C.float(b)
	retval := C.WrapConstructorColorWithRGB(rToC, gToC, bToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// NewColorWithRGBA Four-component RGBA color object.
func NewColorWithRGBA(r float32, g float32, b float32, a float32) *Color {
	rToC := C.float(r)
	gToC := C.float(g)
	bToC := C.float(b)
	aToC := C.float(a)
	retval := C.WrapConstructorColorWithRGBA(rToC, gToC, bToC, aToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Color) Free() {
	C.WrapColorFree(pointer.h)
}

// IsNil ...
func (pointer *Color) IsNil() bool {
	return pointer.h == C.WrapColor(nil)
}

// Add ...
func (pointer *Color) Add(color *Color) *Color {
	colorToC := color.h
	retval := C.WrapAddColor(pointer.h, colorToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// AddWithK ...
func (pointer *Color) AddWithK(k float32) *Color {
	kToC := C.float(k)
	retval := C.WrapAddColorWithK(pointer.h, kToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// Sub ...
func (pointer *Color) Sub(color *Color) *Color {
	colorToC := color.h
	retval := C.WrapSubColor(pointer.h, colorToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// SubWithK ...
func (pointer *Color) SubWithK(k float32) *Color {
	kToC := C.float(k)
	retval := C.WrapSubColorWithK(pointer.h, kToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// Div ...
func (pointer *Color) Div(color *Color) *Color {
	colorToC := color.h
	retval := C.WrapDivColor(pointer.h, colorToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// DivWithK ...
func (pointer *Color) DivWithK(k float32) *Color {
	kToC := C.float(k)
	retval := C.WrapDivColorWithK(pointer.h, kToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// Mul ...
func (pointer *Color) Mul(color *Color) *Color {
	colorToC := color.h
	retval := C.WrapMulColor(pointer.h, colorToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// MulWithK ...
func (pointer *Color) MulWithK(k float32) *Color {
	kToC := C.float(k)
	retval := C.WrapMulColorWithK(pointer.h, kToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// InplaceAdd ...
func (pointer *Color) InplaceAdd(color *Color) {
	colorToC := color.h
	C.WrapInplaceAddColor(pointer.h, colorToC)
}

// InplaceAddWithK ...
func (pointer *Color) InplaceAddWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceAddColorWithK(pointer.h, kToC)
}

// InplaceSub ...
func (pointer *Color) InplaceSub(color *Color) {
	colorToC := color.h
	C.WrapInplaceSubColor(pointer.h, colorToC)
}

// InplaceSubWithK ...
func (pointer *Color) InplaceSubWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceSubColorWithK(pointer.h, kToC)
}

// InplaceMul ...
func (pointer *Color) InplaceMul(color *Color) {
	colorToC := color.h
	C.WrapInplaceMulColor(pointer.h, colorToC)
}

// InplaceMulWithK ...
func (pointer *Color) InplaceMulWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceMulColorWithK(pointer.h, kToC)
}

// InplaceDiv ...
func (pointer *Color) InplaceDiv(color *Color) {
	colorToC := color.h
	C.WrapInplaceDivColor(pointer.h, colorToC)
}

// InplaceDivWithK ...
func (pointer *Color) InplaceDivWithK(k float32) {
	kToC := C.float(k)
	C.WrapInplaceDivColorWithK(pointer.h, kToC)
}

// Eq ...
func (pointer *Color) Eq(color *Color) bool {
	colorToC := color.h
	retval := C.WrapEqColor(pointer.h, colorToC)
	return bool(retval)
}

// Ne ...
func (pointer *Color) Ne(color *Color) bool {
	colorToC := color.h
	retval := C.WrapNeColor(pointer.h, colorToC)
	return bool(retval)
}

// GoSliceOfColor ...
type GoSliceOfColor []*Color

// ColorList  ...
type ColorList struct {
	h C.WrapColorList
}

// Get ...
func (pointer *ColorList) Get(id int) *Color {
	v := C.WrapColorListGetOperator(pointer.h, C.int(id))
	vGO := &Color{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *ColorList) Set(id int, v *Color) {
	vToC := v.h
	C.WrapColorListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *ColorList) Len() int32 {
	return int32(C.WrapColorListLenOperator(pointer.h))
}

// NewColorList ...
func NewColorList() *ColorList {
	retval := C.WrapConstructorColorList()
	retvalGO := &ColorList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ColorList) {
		C.WrapColorListFree(cleanval.h)
	})
	return retvalGO
}

// NewColorListWithSequence ...
func NewColorListWithSequence(sequence GoSliceOfColor) *ColorList {
	var sequencePointer []C.WrapColor
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapColor)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorColorListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &ColorList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ColorList) {
		C.WrapColorListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *ColorList) Free() {
	C.WrapColorListFree(pointer.h)
}

// IsNil ...
func (pointer *ColorList) IsNil() bool {
	return pointer.h == C.WrapColorList(nil)
}

// Clear ...
func (pointer *ColorList) Clear() {
	C.WrapClearColorList(pointer.h)
}

// Reserve ...
func (pointer *ColorList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveColorList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *ColorList) PushBack(v *Color) {
	vToC := v.h
	C.WrapPushBackColorList(pointer.h, vToC)
}

// Size ...
func (pointer *ColorList) Size() int32 {
	retval := C.WrapSizeColorList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *ColorList) At(idx int32) *Color {
	idxToC := C.size_t(idx)
	retval := C.WrapAtColorList(pointer.h, idxToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// Picture  The picture origin (0, 0) is in the top-left corner of its frame with the X and Y axises increasing toward the right and bottom.  To load and save a picture use [harfang.LoadPicture], [harfang.LoadPNG] or [harfang.SavePNG].  The [harfang.Picture_SetData] and [harfang.Picture_GetData] methods can be used to transfer data to and from a picture object.
type Picture struct {
	h C.WrapPicture
}

// NewPicture The picture origin (0, 0) is in the top-left corner of its frame with the X and Y axises increasing toward the right and bottom.  To load and save a picture use [harfang.LoadPicture], [harfang.LoadPNG] or [harfang.SavePNG].  The [harfang.Picture_SetData] and [harfang.Picture_GetData] methods can be used to transfer data to and from a picture object.
func NewPicture() *Picture {
	retval := C.WrapConstructorPicture()
	retvalGO := &Picture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Picture) {
		C.WrapPictureFree(cleanval.h)
	})
	return retvalGO
}

// NewPictureWithPicture The picture origin (0, 0) is in the top-left corner of its frame with the X and Y axises increasing toward the right and bottom.  To load and save a picture use [harfang.LoadPicture], [harfang.LoadPNG] or [harfang.SavePNG].  The [harfang.Picture_SetData] and [harfang.Picture_GetData] methods can be used to transfer data to and from a picture object.
func NewPictureWithPicture(picture *Picture) *Picture {
	pictureToC := picture.h
	retval := C.WrapConstructorPictureWithPicture(pictureToC)
	retvalGO := &Picture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Picture) {
		C.WrapPictureFree(cleanval.h)
	})
	return retvalGO
}

// NewPictureWithWidthHeightFormat The picture origin (0, 0) is in the top-left corner of its frame with the X and Y axises increasing toward the right and bottom.  To load and save a picture use [harfang.LoadPicture], [harfang.LoadPNG] or [harfang.SavePNG].  The [harfang.Picture_SetData] and [harfang.Picture_GetData] methods can be used to transfer data to and from a picture object.
func NewPictureWithWidthHeightFormat(width uint16, height uint16, format PictureFormat) *Picture {
	widthToC := C.ushort(width)
	heightToC := C.ushort(height)
	formatToC := C.int32_t(format)
	retval := C.WrapConstructorPictureWithWidthHeightFormat(widthToC, heightToC, formatToC)
	retvalGO := &Picture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Picture) {
		C.WrapPictureFree(cleanval.h)
	})
	return retvalGO
}

// NewPictureWithDataWidthHeightFormat The picture origin (0, 0) is in the top-left corner of its frame with the X and Y axises increasing toward the right and bottom.  To load and save a picture use [harfang.LoadPicture], [harfang.LoadPNG] or [harfang.SavePNG].  The [harfang.Picture_SetData] and [harfang.Picture_GetData] methods can be used to transfer data to and from a picture object.
func NewPictureWithDataWidthHeightFormat(data *VoidPointer, width uint16, height uint16, format PictureFormat) *Picture {
	dataToC := data.h
	widthToC := C.ushort(width)
	heightToC := C.ushort(height)
	formatToC := C.int32_t(format)
	retval := C.WrapConstructorPictureWithDataWidthHeightFormat(dataToC, widthToC, heightToC, formatToC)
	retvalGO := &Picture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Picture) {
		C.WrapPictureFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Picture) Free() {
	C.WrapPictureFree(pointer.h)
}

// IsNil ...
func (pointer *Picture) IsNil() bool {
	return pointer.h == C.WrapPicture(nil)
}

// GetWidth Return the picture width.
func (pointer *Picture) GetWidth() uint32 {
	retval := C.WrapGetWidthPicture(pointer.h)
	return uint32(retval)
}

// GetHeight Return the picture height.
func (pointer *Picture) GetHeight() uint32 {
	retval := C.WrapGetHeightPicture(pointer.h)
	return uint32(retval)
}

// GetFormat ...
func (pointer *Picture) GetFormat() PictureFormat {
	retval := C.WrapGetFormatPicture(pointer.h)
	return PictureFormat(retval)
}

// GetData ...
func (pointer *Picture) GetData() uintptr {
	retval := C.WrapGetDataPicture(pointer.h)
	return uintptr(retval)
}

// SetData ...
func (pointer *Picture) SetData(data *VoidPointer, width uint16, height uint16, format PictureFormat) {
	dataToC := data.h
	widthToC := C.ushort(width)
	heightToC := C.ushort(height)
	formatToC := C.int32_t(format)
	C.WrapSetDataPicture(pointer.h, dataToC, widthToC, heightToC, formatToC)
}

// CopyData ...
func (pointer *Picture) CopyData(data *VoidPointer, width uint16, height uint16, format PictureFormat) {
	dataToC := data.h
	widthToC := C.ushort(width)
	heightToC := C.ushort(height)
	formatToC := C.int32_t(format)
	C.WrapCopyDataPicture(pointer.h, dataToC, widthToC, heightToC, formatToC)
}

// GetPixelRGBA ...
func (pointer *Picture) GetPixelRGBA(x uint16, y uint16) *Color {
	xToC := C.ushort(x)
	yToC := C.ushort(y)
	retval := C.WrapGetPixelRGBAPicture(pointer.h, xToC, yToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// SetPixelRGBA ...
func (pointer *Picture) SetPixelRGBA(x uint16, y uint16, col *Color) {
	xToC := C.ushort(x)
	yToC := C.ushort(y)
	colToC := col.h
	C.WrapSetPixelRGBAPicture(pointer.h, xToC, yToC, colToC)
}

// FrameBufferHandle  ...
type FrameBufferHandle struct {
	h C.WrapFrameBufferHandle
}

// Free ...
func (pointer *FrameBufferHandle) Free() {
	C.WrapFrameBufferHandleFree(pointer.h)
}

// IsNil ...
func (pointer *FrameBufferHandle) IsNil() bool {
	return pointer.h == C.WrapFrameBufferHandle(nil)
}

// VertexLayout  Memory layout and types of vertex attributes.
type VertexLayout struct {
	h C.WrapVertexLayout
}

// NewVertexLayout Memory layout and types of vertex attributes.
func NewVertexLayout() *VertexLayout {
	retval := C.WrapConstructorVertexLayout()
	retvalGO := &VertexLayout{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VertexLayout) {
		C.WrapVertexLayoutFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *VertexLayout) Free() {
	C.WrapVertexLayoutFree(pointer.h)
}

// IsNil ...
func (pointer *VertexLayout) IsNil() bool {
	return pointer.h == C.WrapVertexLayout(nil)
}

// Begin ...
func (pointer *VertexLayout) Begin() *VertexLayout {
	retval := C.WrapBeginVertexLayout(pointer.h)
	var retvalGO *VertexLayout
	if retval != nil {
		retvalGO = &VertexLayout{h: retval}
	}
	return retvalGO
}

// Add ...
func (pointer *VertexLayout) Add(attrib Attrib, count uint8, typeGo AttribType) *VertexLayout {
	attribToC := C.int32_t(attrib)
	countToC := C.uchar(count)
	typeGoToC := C.int32_t(typeGo)
	retval := C.WrapAddVertexLayout(pointer.h, attribToC, countToC, typeGoToC)
	var retvalGO *VertexLayout
	if retval != nil {
		retvalGO = &VertexLayout{h: retval}
	}
	return retvalGO
}

// AddWithNormalized ...
func (pointer *VertexLayout) AddWithNormalized(attrib Attrib, count uint8, typeGo AttribType, normalized bool) *VertexLayout {
	attribToC := C.int32_t(attrib)
	countToC := C.uchar(count)
	typeGoToC := C.int32_t(typeGo)
	normalizedToC := C.bool(normalized)
	retval := C.WrapAddVertexLayoutWithNormalized(pointer.h, attribToC, countToC, typeGoToC, normalizedToC)
	var retvalGO *VertexLayout
	if retval != nil {
		retvalGO = &VertexLayout{h: retval}
	}
	return retvalGO
}

// AddWithNormalizedAsInt ...
func (pointer *VertexLayout) AddWithNormalizedAsInt(attrib Attrib, count uint8, typeGo AttribType, normalized bool, asint bool) *VertexLayout {
	attribToC := C.int32_t(attrib)
	countToC := C.uchar(count)
	typeGoToC := C.int32_t(typeGo)
	normalizedToC := C.bool(normalized)
	asintToC := C.bool(asint)
	retval := C.WrapAddVertexLayoutWithNormalizedAsInt(pointer.h, attribToC, countToC, typeGoToC, normalizedToC, asintToC)
	var retvalGO *VertexLayout
	if retval != nil {
		retvalGO = &VertexLayout{h: retval}
	}
	return retvalGO
}

// Skip ...
func (pointer *VertexLayout) Skip(size uint8) *VertexLayout {
	sizeToC := C.uchar(size)
	retval := C.WrapSkipVertexLayout(pointer.h, sizeToC)
	var retvalGO *VertexLayout
	if retval != nil {
		retvalGO = &VertexLayout{h: retval}
	}
	return retvalGO
}

// End ...
func (pointer *VertexLayout) End() {
	C.WrapEndVertexLayout(pointer.h)
}

// Has ...
func (pointer *VertexLayout) Has(attrib Attrib) bool {
	attribToC := C.int32_t(attrib)
	retval := C.WrapHasVertexLayout(pointer.h, attribToC)
	return bool(retval)
}

// GetOffset ...
func (pointer *VertexLayout) GetOffset(attrib Attrib) uint16 {
	attribToC := C.int32_t(attrib)
	retval := C.WrapGetOffsetVertexLayout(pointer.h, attribToC)
	return uint16(retval)
}

// GetStride ...
func (pointer *VertexLayout) GetStride() uint16 {
	retval := C.WrapGetStrideVertexLayout(pointer.h)
	return uint16(retval)
}

// GetSize ...
func (pointer *VertexLayout) GetSize(count uint32) uint32 {
	countToC := C.uint32_t(count)
	retval := C.WrapGetSizeVertexLayout(pointer.h, countToC)
	return uint32(retval)
}

// ProgramHandle  Handle to a shader program.
type ProgramHandle struct {
	h C.WrapProgramHandle
}

// Free ...
func (pointer *ProgramHandle) Free() {
	C.WrapProgramHandleFree(pointer.h)
}

// IsNil ...
func (pointer *ProgramHandle) IsNil() bool {
	return pointer.h == C.WrapProgramHandle(nil)
}

// TextureInfo  ...
type TextureInfo struct {
	h C.WrapTextureInfo
}

// GetFormat ...
func (pointer *TextureInfo) GetFormat() TextureFormat {
	v := C.WrapTextureInfoGetFormat(pointer.h)
	return TextureFormat(v)
}

// SetFormat ...
func (pointer *TextureInfo) SetFormat(v TextureFormat) {
	vToC := C.int32_t(v)
	C.WrapTextureInfoSetFormat(pointer.h, vToC)
}

// GetStorageSize ...
func (pointer *TextureInfo) GetStorageSize() uint32 {
	v := C.WrapTextureInfoGetStorageSize(pointer.h)
	return uint32(v)
}

// SetStorageSize ...
func (pointer *TextureInfo) SetStorageSize(v uint32) {
	vToC := C.uint32_t(v)
	C.WrapTextureInfoSetStorageSize(pointer.h, vToC)
}

// GetWidth ...
func (pointer *TextureInfo) GetWidth() uint16 {
	v := C.WrapTextureInfoGetWidth(pointer.h)
	return uint16(v)
}

// SetWidth ...
func (pointer *TextureInfo) SetWidth(v uint16) {
	vToC := C.ushort(v)
	C.WrapTextureInfoSetWidth(pointer.h, vToC)
}

// GetHeight ...
func (pointer *TextureInfo) GetHeight() uint16 {
	v := C.WrapTextureInfoGetHeight(pointer.h)
	return uint16(v)
}

// SetHeight ...
func (pointer *TextureInfo) SetHeight(v uint16) {
	vToC := C.ushort(v)
	C.WrapTextureInfoSetHeight(pointer.h, vToC)
}

// GetDepth ...
func (pointer *TextureInfo) GetDepth() uint16 {
	v := C.WrapTextureInfoGetDepth(pointer.h)
	return uint16(v)
}

// SetDepth ...
func (pointer *TextureInfo) SetDepth(v uint16) {
	vToC := C.ushort(v)
	C.WrapTextureInfoSetDepth(pointer.h, vToC)
}

// GetNumLayers ...
func (pointer *TextureInfo) GetNumLayers() uint16 {
	v := C.WrapTextureInfoGetNumLayers(pointer.h)
	return uint16(v)
}

// SetNumLayers ...
func (pointer *TextureInfo) SetNumLayers(v uint16) {
	vToC := C.ushort(v)
	C.WrapTextureInfoSetNumLayers(pointer.h, vToC)
}

// GetNumMips ...
func (pointer *TextureInfo) GetNumMips() uint8 {
	v := C.WrapTextureInfoGetNumMips(pointer.h)
	return uint8(v)
}

// SetNumMips ...
func (pointer *TextureInfo) SetNumMips(v uint8) {
	vToC := C.uchar(v)
	C.WrapTextureInfoSetNumMips(pointer.h, vToC)
}

// GetBitsPerPixel ...
func (pointer *TextureInfo) GetBitsPerPixel() uint8 {
	v := C.WrapTextureInfoGetBitsPerPixel(pointer.h)
	return uint8(v)
}

// SetBitsPerPixel ...
func (pointer *TextureInfo) SetBitsPerPixel(v uint8) {
	vToC := C.uchar(v)
	C.WrapTextureInfoSetBitsPerPixel(pointer.h, vToC)
}

// GetCubeMap ...
func (pointer *TextureInfo) GetCubeMap() bool {
	v := C.WrapTextureInfoGetCubeMap(pointer.h)
	return bool(v)
}

// SetCubeMap ...
func (pointer *TextureInfo) SetCubeMap(v bool) {
	vToC := C.bool(v)
	C.WrapTextureInfoSetCubeMap(pointer.h, vToC)
}

// NewTextureInfo ...
func NewTextureInfo() *TextureInfo {
	retval := C.WrapConstructorTextureInfo()
	retvalGO := &TextureInfo{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *TextureInfo) {
		C.WrapTextureInfoFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *TextureInfo) Free() {
	C.WrapTextureInfoFree(pointer.h)
}

// IsNil ...
func (pointer *TextureInfo) IsNil() bool {
	return pointer.h == C.WrapTextureInfo(nil)
}

// ModelRef  Reference to a [harfang.Model] inside a [harfang.PipelineResources] object.  See [harfang.LoadModelFromFile], [harfang.LoadModelFromAssets] and [harfang.PipelineResources_AddModel].
type ModelRef struct {
	h C.WrapModelRef
}

// Free ...
func (pointer *ModelRef) Free() {
	C.WrapModelRefFree(pointer.h)
}

// IsNil ...
func (pointer *ModelRef) IsNil() bool {
	return pointer.h == C.WrapModelRef(nil)
}

// Eq ...
func (pointer *ModelRef) Eq(m *ModelRef) bool {
	mToC := m.h
	retval := C.WrapEqModelRef(pointer.h, mToC)
	return bool(retval)
}

// Ne ...
func (pointer *ModelRef) Ne(m *ModelRef) bool {
	mToC := m.h
	retval := C.WrapNeModelRef(pointer.h, mToC)
	return bool(retval)
}

// TextureRef  ...
type TextureRef struct {
	h C.WrapTextureRef
}

// Free ...
func (pointer *TextureRef) Free() {
	C.WrapTextureRefFree(pointer.h)
}

// IsNil ...
func (pointer *TextureRef) IsNil() bool {
	return pointer.h == C.WrapTextureRef(nil)
}

// Eq ...
func (pointer *TextureRef) Eq(t *TextureRef) bool {
	tToC := t.h
	retval := C.WrapEqTextureRef(pointer.h, tToC)
	return bool(retval)
}

// Ne ...
func (pointer *TextureRef) Ne(t *TextureRef) bool {
	tToC := t.h
	retval := C.WrapNeTextureRef(pointer.h, tToC)
	return bool(retval)
}

// MaterialRef  Reference to a [harfang.Material] inside a [harfang.PipelineResources] object.
type MaterialRef struct {
	h C.WrapMaterialRef
}

// Free ...
func (pointer *MaterialRef) Free() {
	C.WrapMaterialRefFree(pointer.h)
}

// IsNil ...
func (pointer *MaterialRef) IsNil() bool {
	return pointer.h == C.WrapMaterialRef(nil)
}

// Eq ...
func (pointer *MaterialRef) Eq(m *MaterialRef) bool {
	mToC := m.h
	retval := C.WrapEqMaterialRef(pointer.h, mToC)
	return bool(retval)
}

// Ne ...
func (pointer *MaterialRef) Ne(m *MaterialRef) bool {
	mToC := m.h
	retval := C.WrapNeMaterialRef(pointer.h, mToC)
	return bool(retval)
}

// PipelineProgramRef  ...
type PipelineProgramRef struct {
	h C.WrapPipelineProgramRef
}

// Free ...
func (pointer *PipelineProgramRef) Free() {
	C.WrapPipelineProgramRefFree(pointer.h)
}

// IsNil ...
func (pointer *PipelineProgramRef) IsNil() bool {
	return pointer.h == C.WrapPipelineProgramRef(nil)
}

// Eq ...
func (pointer *PipelineProgramRef) Eq(p *PipelineProgramRef) bool {
	pToC := p.h
	retval := C.WrapEqPipelineProgramRef(pointer.h, pToC)
	return bool(retval)
}

// Ne ...
func (pointer *PipelineProgramRef) Ne(p *PipelineProgramRef) bool {
	pToC := p.h
	retval := C.WrapNePipelineProgramRef(pointer.h, pToC)
	return bool(retval)
}

// Texture  ...
type Texture struct {
	h C.WrapTexture
}

// NewTexture ...
func NewTexture() *Texture {
	retval := C.WrapConstructorTexture()
	retvalGO := &Texture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Texture) {
		C.WrapTextureFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Texture) Free() {
	C.WrapTextureFree(pointer.h)
}

// IsNil ...
func (pointer *Texture) IsNil() bool {
	return pointer.h == C.WrapTexture(nil)
}

// UniformSetValue  Command object to set a uniform value at draw time.
type UniformSetValue struct {
	h C.WrapUniformSetValue
}

// Free ...
func (pointer *UniformSetValue) Free() {
	C.WrapUniformSetValueFree(pointer.h)
}

// IsNil ...
func (pointer *UniformSetValue) IsNil() bool {
	return pointer.h == C.WrapUniformSetValue(nil)
}

// GoSliceOfUniformSetValue ...
type GoSliceOfUniformSetValue []*UniformSetValue

// UniformSetValueList  ...
type UniformSetValueList struct {
	h C.WrapUniformSetValueList
}

// Get ...
func (pointer *UniformSetValueList) Get(id int) *UniformSetValue {
	v := C.WrapUniformSetValueListGetOperator(pointer.h, C.int(id))
	vGO := &UniformSetValue{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *UniformSetValue) {
		C.WrapUniformSetValueFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *UniformSetValueList) Set(id int, v *UniformSetValue) {
	vToC := v.h
	C.WrapUniformSetValueListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *UniformSetValueList) Len() int32 {
	return int32(C.WrapUniformSetValueListLenOperator(pointer.h))
}

// NewUniformSetValueList ...
func NewUniformSetValueList() *UniformSetValueList {
	retval := C.WrapConstructorUniformSetValueList()
	retvalGO := &UniformSetValueList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetValueList) {
		C.WrapUniformSetValueListFree(cleanval.h)
	})
	return retvalGO
}

// NewUniformSetValueListWithSequence ...
func NewUniformSetValueListWithSequence(sequence GoSliceOfUniformSetValue) *UniformSetValueList {
	var sequencePointer []C.WrapUniformSetValue
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorUniformSetValueListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &UniformSetValueList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetValueList) {
		C.WrapUniformSetValueListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *UniformSetValueList) Free() {
	C.WrapUniformSetValueListFree(pointer.h)
}

// IsNil ...
func (pointer *UniformSetValueList) IsNil() bool {
	return pointer.h == C.WrapUniformSetValueList(nil)
}

// Clear ...
func (pointer *UniformSetValueList) Clear() {
	C.WrapClearUniformSetValueList(pointer.h)
}

// Reserve ...
func (pointer *UniformSetValueList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveUniformSetValueList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *UniformSetValueList) PushBack(v *UniformSetValue) {
	vToC := v.h
	C.WrapPushBackUniformSetValueList(pointer.h, vToC)
}

// Size ...
func (pointer *UniformSetValueList) Size() int32 {
	retval := C.WrapSizeUniformSetValueList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *UniformSetValueList) At(idx int32) *UniformSetValue {
	idxToC := C.size_t(idx)
	retval := C.WrapAtUniformSetValueList(pointer.h, idxToC)
	retvalGO := &UniformSetValue{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetValue) {
		C.WrapUniformSetValueFree(cleanval.h)
	})
	return retvalGO
}

// UniformSetTexture  Command object to set a uniform texture at draw time.
type UniformSetTexture struct {
	h C.WrapUniformSetTexture
}

// Free ...
func (pointer *UniformSetTexture) Free() {
	C.WrapUniformSetTextureFree(pointer.h)
}

// IsNil ...
func (pointer *UniformSetTexture) IsNil() bool {
	return pointer.h == C.WrapUniformSetTexture(nil)
}

// GoSliceOfUniformSetTexture ...
type GoSliceOfUniformSetTexture []*UniformSetTexture

// UniformSetTextureList  ...
type UniformSetTextureList struct {
	h C.WrapUniformSetTextureList
}

// Get ...
func (pointer *UniformSetTextureList) Get(id int) *UniformSetTexture {
	v := C.WrapUniformSetTextureListGetOperator(pointer.h, C.int(id))
	vGO := &UniformSetTexture{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *UniformSetTexture) {
		C.WrapUniformSetTextureFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *UniformSetTextureList) Set(id int, v *UniformSetTexture) {
	vToC := v.h
	C.WrapUniformSetTextureListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *UniformSetTextureList) Len() int32 {
	return int32(C.WrapUniformSetTextureListLenOperator(pointer.h))
}

// NewUniformSetTextureList ...
func NewUniformSetTextureList() *UniformSetTextureList {
	retval := C.WrapConstructorUniformSetTextureList()
	retvalGO := &UniformSetTextureList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetTextureList) {
		C.WrapUniformSetTextureListFree(cleanval.h)
	})
	return retvalGO
}

// NewUniformSetTextureListWithSequence ...
func NewUniformSetTextureListWithSequence(sequence GoSliceOfUniformSetTexture) *UniformSetTextureList {
	var sequencePointer []C.WrapUniformSetTexture
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorUniformSetTextureListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &UniformSetTextureList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetTextureList) {
		C.WrapUniformSetTextureListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *UniformSetTextureList) Free() {
	C.WrapUniformSetTextureListFree(pointer.h)
}

// IsNil ...
func (pointer *UniformSetTextureList) IsNil() bool {
	return pointer.h == C.WrapUniformSetTextureList(nil)
}

// Clear ...
func (pointer *UniformSetTextureList) Clear() {
	C.WrapClearUniformSetTextureList(pointer.h)
}

// Reserve ...
func (pointer *UniformSetTextureList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveUniformSetTextureList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *UniformSetTextureList) PushBack(v *UniformSetTexture) {
	vToC := v.h
	C.WrapPushBackUniformSetTextureList(pointer.h, vToC)
}

// Size ...
func (pointer *UniformSetTextureList) Size() int32 {
	retval := C.WrapSizeUniformSetTextureList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *UniformSetTextureList) At(idx int32) *UniformSetTexture {
	idxToC := C.size_t(idx)
	retval := C.WrapAtUniformSetTextureList(pointer.h, idxToC)
	retvalGO := &UniformSetTexture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetTexture) {
		C.WrapUniformSetTextureFree(cleanval.h)
	})
	return retvalGO
}

// PipelineProgram  ...
type PipelineProgram struct {
	h C.WrapPipelineProgram
}

// Free ...
func (pointer *PipelineProgram) Free() {
	C.WrapPipelineProgramFree(pointer.h)
}

// IsNil ...
func (pointer *PipelineProgram) IsNil() bool {
	return pointer.h == C.WrapPipelineProgram(nil)
}

// ViewState  Everything required to define an observer inside a 3d world. This object holds the projection matrix and its associated frustum as well as the transformation of the observer.  The world content is transformed by the observer view matrix before being projected to screen using its projection matrix.
type ViewState struct {
	h C.WrapViewState
}

// GetFrustum ...
func (pointer *ViewState) GetFrustum() *Frustum {
	v := C.WrapViewStateGetFrustum(pointer.h)
	vGO := &Frustum{h: v}
	return vGO
}

// SetFrustum ...
func (pointer *ViewState) SetFrustum(v *Frustum) {
	vToC := v.h
	C.WrapViewStateSetFrustum(pointer.h, vToC)
}

// GetProj ...
func (pointer *ViewState) GetProj() *Mat44 {
	v := C.WrapViewStateGetProj(pointer.h)
	vGO := &Mat44{h: v}
	return vGO
}

// SetProj ...
func (pointer *ViewState) SetProj(v *Mat44) {
	vToC := v.h
	C.WrapViewStateSetProj(pointer.h, vToC)
}

// GetView ...
func (pointer *ViewState) GetView() *Mat4 {
	v := C.WrapViewStateGetView(pointer.h)
	vGO := &Mat4{h: v}
	return vGO
}

// SetView ...
func (pointer *ViewState) SetView(v *Mat4) {
	vToC := v.h
	C.WrapViewStateSetView(pointer.h, vToC)
}

// NewViewState Everything required to define an observer inside a 3d world. This object holds the projection matrix and its associated frustum as well as the transformation of the observer.  The world content is transformed by the observer view matrix before being projected to screen using its projection matrix.
func NewViewState() *ViewState {
	retval := C.WrapConstructorViewState()
	retvalGO := &ViewState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ViewState) {
		C.WrapViewStateFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *ViewState) Free() {
	C.WrapViewStateFree(pointer.h)
}

// IsNil ...
func (pointer *ViewState) IsNil() bool {
	return pointer.h == C.WrapViewState(nil)
}

// Material  High-level description of visual aspects of a surface. A material is comprised of a [harfang.PipelineProgramRef], per-uniform value or texture, and a [harfang.RenderState].  See [harfang.man.ForwardPipeline] and [harfang.man.PipelineShader].
type Material struct {
	h C.WrapMaterial
}

// NewMaterial High-level description of visual aspects of a surface. A material is comprised of a [harfang.PipelineProgramRef], per-uniform value or texture, and a [harfang.RenderState].  See [harfang.man.ForwardPipeline] and [harfang.man.PipelineShader].
func NewMaterial() *Material {
	retval := C.WrapConstructorMaterial()
	retvalGO := &Material{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Material) {
		C.WrapMaterialFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Material) Free() {
	C.WrapMaterialFree(pointer.h)
}

// IsNil ...
func (pointer *Material) IsNil() bool {
	return pointer.h == C.WrapMaterial(nil)
}

// GoSliceOfMaterial ...
type GoSliceOfMaterial []*Material

// MaterialList  ...
type MaterialList struct {
	h C.WrapMaterialList
}

// Get ...
func (pointer *MaterialList) Get(id int) *Material {
	v := C.WrapMaterialListGetOperator(pointer.h, C.int(id))
	vGO := &Material{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *Material) {
		C.WrapMaterialFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *MaterialList) Set(id int, v *Material) {
	vToC := v.h
	C.WrapMaterialListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *MaterialList) Len() int32 {
	return int32(C.WrapMaterialListLenOperator(pointer.h))
}

// NewMaterialList ...
func NewMaterialList() *MaterialList {
	retval := C.WrapConstructorMaterialList()
	retvalGO := &MaterialList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MaterialList) {
		C.WrapMaterialListFree(cleanval.h)
	})
	return retvalGO
}

// NewMaterialListWithSequence ...
func NewMaterialListWithSequence(sequence GoSliceOfMaterial) *MaterialList {
	var sequencePointer []C.WrapMaterial
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapMaterial)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorMaterialListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &MaterialList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MaterialList) {
		C.WrapMaterialListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *MaterialList) Free() {
	C.WrapMaterialListFree(pointer.h)
}

// IsNil ...
func (pointer *MaterialList) IsNil() bool {
	return pointer.h == C.WrapMaterialList(nil)
}

// Clear ...
func (pointer *MaterialList) Clear() {
	C.WrapClearMaterialList(pointer.h)
}

// Reserve ...
func (pointer *MaterialList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveMaterialList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *MaterialList) PushBack(v *Material) {
	vToC := v.h
	C.WrapPushBackMaterialList(pointer.h, vToC)
}

// Size ...
func (pointer *MaterialList) Size() int32 {
	retval := C.WrapSizeMaterialList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *MaterialList) At(idx int32) *Material {
	idxToC := C.size_t(idx)
	retval := C.WrapAtMaterialList(pointer.h, idxToC)
	retvalGO := &Material{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Material) {
		C.WrapMaterialFree(cleanval.h)
	})
	return retvalGO
}

// RenderState  ...
type RenderState struct {
	h C.WrapRenderState
}

// Free ...
func (pointer *RenderState) Free() {
	C.WrapRenderStateFree(pointer.h)
}

// IsNil ...
func (pointer *RenderState) IsNil() bool {
	return pointer.h == C.WrapRenderState(nil)
}

// Model  Runtime version of a [harfang.Geometry]. A model can be drawn to screen by calling [harfang.DrawModel] or by assigning it to the [harfang.Object] component of a node.  To programmatically create a model see [harfang.ModelBuilder].
type Model struct {
	h C.WrapModel
}

// Free ...
func (pointer *Model) Free() {
	C.WrapModelFree(pointer.h)
}

// IsNil ...
func (pointer *Model) IsNil() bool {
	return pointer.h == C.WrapModel(nil)
}

// PipelineResources  ...
type PipelineResources struct {
	h C.WrapPipelineResources
}

// NewPipelineResources ...
func NewPipelineResources() *PipelineResources {
	retval := C.WrapConstructorPipelineResources()
	retvalGO := &PipelineResources{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *PipelineResources) {
		C.WrapPipelineResourcesFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *PipelineResources) Free() {
	C.WrapPipelineResourcesFree(pointer.h)
}

// IsNil ...
func (pointer *PipelineResources) IsNil() bool {
	return pointer.h == C.WrapPipelineResources(nil)
}

// AddTexture ...
func (pointer *PipelineResources) AddTexture(name string, tex *Texture) *TextureRef {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	texToC := tex.h
	retval := C.WrapAddTexturePipelineResources(pointer.h, nameToC, texToC)
	retvalGO := &TextureRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *TextureRef) {
		C.WrapTextureRefFree(cleanval.h)
	})
	return retvalGO
}

// AddModel ...
func (pointer *PipelineResources) AddModel(name string, mdl *Model) *ModelRef {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	mdlToC := mdl.h
	retval := C.WrapAddModelPipelineResources(pointer.h, nameToC, mdlToC)
	retvalGO := &ModelRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ModelRef) {
		C.WrapModelRefFree(cleanval.h)
	})
	return retvalGO
}

// AddProgram ...
func (pointer *PipelineResources) AddProgram(name string, prg *PipelineProgram) *PipelineProgramRef {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	prgToC := prg.h
	retval := C.WrapAddProgramPipelineResources(pointer.h, nameToC, prgToC)
	retvalGO := &PipelineProgramRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *PipelineProgramRef) {
		C.WrapPipelineProgramRefFree(cleanval.h)
	})
	return retvalGO
}

// HasTexture ...
func (pointer *PipelineResources) HasTexture(name string) *TextureRef {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapHasTexturePipelineResources(pointer.h, nameToC)
	retvalGO := &TextureRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *TextureRef) {
		C.WrapTextureRefFree(cleanval.h)
	})
	return retvalGO
}

// HasModel ...
func (pointer *PipelineResources) HasModel(name string) *ModelRef {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapHasModelPipelineResources(pointer.h, nameToC)
	retvalGO := &ModelRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ModelRef) {
		C.WrapModelRefFree(cleanval.h)
	})
	return retvalGO
}

// HasProgram ...
func (pointer *PipelineResources) HasProgram(name string) *PipelineProgramRef {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapHasProgramPipelineResources(pointer.h, nameToC)
	retvalGO := &PipelineProgramRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *PipelineProgramRef) {
		C.WrapPipelineProgramRefFree(cleanval.h)
	})
	return retvalGO
}

// UpdateTexture ...
func (pointer *PipelineResources) UpdateTexture(ref *TextureRef, tex *Texture) {
	refToC := ref.h
	texToC := tex.h
	C.WrapUpdateTexturePipelineResources(pointer.h, refToC, texToC)
}

// UpdateModel ...
func (pointer *PipelineResources) UpdateModel(ref *ModelRef, mdl *Model) {
	refToC := ref.h
	mdlToC := mdl.h
	C.WrapUpdateModelPipelineResources(pointer.h, refToC, mdlToC)
}

// UpdateProgram ...
func (pointer *PipelineResources) UpdateProgram(ref *PipelineProgramRef, prg *PipelineProgram) {
	refToC := ref.h
	prgToC := prg.h
	C.WrapUpdateProgramPipelineResources(pointer.h, refToC, prgToC)
}

// GetTexture ...
func (pointer *PipelineResources) GetTexture(ref *TextureRef) *Texture {
	refToC := ref.h
	retval := C.WrapGetTexturePipelineResources(pointer.h, refToC)
	var retvalGO *Texture
	if retval != nil {
		retvalGO = &Texture{h: retval}
	}
	return retvalGO
}

// GetModel ...
func (pointer *PipelineResources) GetModel(ref *ModelRef) *Model {
	refToC := ref.h
	retval := C.WrapGetModelPipelineResources(pointer.h, refToC)
	var retvalGO *Model
	if retval != nil {
		retvalGO = &Model{h: retval}
	}
	return retvalGO
}

// GetProgram ...
func (pointer *PipelineResources) GetProgram(ref *PipelineProgramRef) *PipelineProgram {
	refToC := ref.h
	retval := C.WrapGetProgramPipelineResources(pointer.h, refToC)
	var retvalGO *PipelineProgram
	if retval != nil {
		retvalGO = &PipelineProgram{h: retval}
	}
	return retvalGO
}

// GetTextureName ...
func (pointer *PipelineResources) GetTextureName(ref *TextureRef) string {
	refToC := ref.h
	retval := C.WrapGetTextureNamePipelineResources(pointer.h, refToC)
	return C.GoString(retval)
}

// GetModelName ...
func (pointer *PipelineResources) GetModelName(ref *ModelRef) string {
	refToC := ref.h
	retval := C.WrapGetModelNamePipelineResources(pointer.h, refToC)
	return C.GoString(retval)
}

// GetProgramName ...
func (pointer *PipelineResources) GetProgramName(ref *PipelineProgramRef) string {
	refToC := ref.h
	retval := C.WrapGetProgramNamePipelineResources(pointer.h, refToC)
	return C.GoString(retval)
}

// DestroyAllTextures ...
func (pointer *PipelineResources) DestroyAllTextures() {
	C.WrapDestroyAllTexturesPipelineResources(pointer.h)
}

// DestroyAllModels ...
func (pointer *PipelineResources) DestroyAllModels() {
	C.WrapDestroyAllModelsPipelineResources(pointer.h)
}

// DestroyAllPrograms ...
func (pointer *PipelineResources) DestroyAllPrograms() {
	C.WrapDestroyAllProgramsPipelineResources(pointer.h)
}

// DestroyTexture ...
func (pointer *PipelineResources) DestroyTexture(ref *TextureRef) {
	refToC := ref.h
	C.WrapDestroyTexturePipelineResources(pointer.h, refToC)
}

// DestroyModel ...
func (pointer *PipelineResources) DestroyModel(ref *ModelRef) {
	refToC := ref.h
	C.WrapDestroyModelPipelineResources(pointer.h, refToC)
}

// DestroyProgram ...
func (pointer *PipelineResources) DestroyProgram(ref *PipelineProgramRef) {
	refToC := ref.h
	C.WrapDestroyProgramPipelineResources(pointer.h, refToC)
}

// HasTextureInfo ...
func (pointer *PipelineResources) HasTextureInfo(ref *TextureRef) bool {
	refToC := ref.h
	retval := C.WrapHasTextureInfoPipelineResources(pointer.h, refToC)
	return bool(retval)
}

// GetTextureInfo ...
func (pointer *PipelineResources) GetTextureInfo(ref *TextureRef) *TextureInfo {
	refToC := ref.h
	retval := C.WrapGetTextureInfoPipelineResources(pointer.h, refToC)
	retvalGO := &TextureInfo{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *TextureInfo) {
		C.WrapTextureInfoFree(cleanval.h)
	})
	return retvalGO
}

// FrameBuffer  ...
type FrameBuffer struct {
	h C.WrapFrameBuffer
}

// GetHandle ...
func (pointer *FrameBuffer) GetHandle() *FrameBufferHandle {
	v := C.WrapFrameBufferGetHandle(pointer.h)
	vGO := &FrameBufferHandle{h: v}
	return vGO
}

// SetHandle ...
func (pointer *FrameBuffer) SetHandle(v *FrameBufferHandle) {
	vToC := v.h
	C.WrapFrameBufferSetHandle(pointer.h, vToC)
}

// Free ...
func (pointer *FrameBuffer) Free() {
	C.WrapFrameBufferFree(pointer.h)
}

// IsNil ...
func (pointer *FrameBuffer) IsNil() bool {
	return pointer.h == C.WrapFrameBuffer(nil)
}

// Vertices  Helper class to generate vertex buffers for drawing primitives.
type Vertices struct {
	h C.WrapVertices
}

// NewVertices Helper class to generate vertex buffers for drawing primitives.
func NewVertices(decl *VertexLayout, count int32) *Vertices {
	declToC := decl.h
	countToC := C.size_t(count)
	retval := C.WrapConstructorVertices(declToC, countToC)
	retvalGO := &Vertices{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vertices) {
		C.WrapVerticesFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Vertices) Free() {
	C.WrapVerticesFree(pointer.h)
}

// IsNil ...
func (pointer *Vertices) IsNil() bool {
	return pointer.h == C.WrapVertices(nil)
}

// GetDecl ...
func (pointer *Vertices) GetDecl() *VertexLayout {
	retval := C.WrapGetDeclVertices(pointer.h)
	var retvalGO *VertexLayout
	if retval != nil {
		retvalGO = &VertexLayout{h: retval}
	}
	return retvalGO
}

// Begin ...
func (pointer *Vertices) Begin(vertexindex int32) *Vertices {
	vertexindexToC := C.size_t(vertexindex)
	retval := C.WrapBeginVertices(pointer.h, vertexindexToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetPos ...
func (pointer *Vertices) SetPos(pos *Vec3) *Vertices {
	posToC := pos.h
	retval := C.WrapSetPosVertices(pointer.h, posToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetNormal ...
func (pointer *Vertices) SetNormal(normal *Vec3) *Vertices {
	normalToC := normal.h
	retval := C.WrapSetNormalVertices(pointer.h, normalToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetTangent ...
func (pointer *Vertices) SetTangent(tangent *Vec3) *Vertices {
	tangentToC := tangent.h
	retval := C.WrapSetTangentVertices(pointer.h, tangentToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetBinormal ...
func (pointer *Vertices) SetBinormal(binormal *Vec3) *Vertices {
	binormalToC := binormal.h
	retval := C.WrapSetBinormalVertices(pointer.h, binormalToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetTexCoord0 ...
func (pointer *Vertices) SetTexCoord0(uv *Vec2) *Vertices {
	uvToC := uv.h
	retval := C.WrapSetTexCoord0Vertices(pointer.h, uvToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetTexCoord1 ...
func (pointer *Vertices) SetTexCoord1(uv *Vec2) *Vertices {
	uvToC := uv.h
	retval := C.WrapSetTexCoord1Vertices(pointer.h, uvToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetTexCoord2 ...
func (pointer *Vertices) SetTexCoord2(uv *Vec2) *Vertices {
	uvToC := uv.h
	retval := C.WrapSetTexCoord2Vertices(pointer.h, uvToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetTexCoord3 ...
func (pointer *Vertices) SetTexCoord3(uv *Vec2) *Vertices {
	uvToC := uv.h
	retval := C.WrapSetTexCoord3Vertices(pointer.h, uvToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetTexCoord4 ...
func (pointer *Vertices) SetTexCoord4(uv *Vec2) *Vertices {
	uvToC := uv.h
	retval := C.WrapSetTexCoord4Vertices(pointer.h, uvToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetTexCoord5 ...
func (pointer *Vertices) SetTexCoord5(uv *Vec2) *Vertices {
	uvToC := uv.h
	retval := C.WrapSetTexCoord5Vertices(pointer.h, uvToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetTexCoord6 ...
func (pointer *Vertices) SetTexCoord6(uv *Vec2) *Vertices {
	uvToC := uv.h
	retval := C.WrapSetTexCoord6Vertices(pointer.h, uvToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetTexCoord7 ...
func (pointer *Vertices) SetTexCoord7(uv *Vec2) *Vertices {
	uvToC := uv.h
	retval := C.WrapSetTexCoord7Vertices(pointer.h, uvToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetColor0 ...
func (pointer *Vertices) SetColor0(color *Color) *Vertices {
	colorToC := color.h
	retval := C.WrapSetColor0Vertices(pointer.h, colorToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetColor1 ...
func (pointer *Vertices) SetColor1(color *Color) *Vertices {
	colorToC := color.h
	retval := C.WrapSetColor1Vertices(pointer.h, colorToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetColor2 ...
func (pointer *Vertices) SetColor2(color *Color) *Vertices {
	colorToC := color.h
	retval := C.WrapSetColor2Vertices(pointer.h, colorToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// SetColor3 ...
func (pointer *Vertices) SetColor3(color *Color) *Vertices {
	colorToC := color.h
	retval := C.WrapSetColor3Vertices(pointer.h, colorToC)
	var retvalGO *Vertices
	if retval != nil {
		retvalGO = &Vertices{h: retval}
	}
	return retvalGO
}

// End ...
func (pointer *Vertices) End() {
	C.WrapEndVertices(pointer.h)
}

// EndWithValidate ...
func (pointer *Vertices) EndWithValidate(validate bool) {
	validateToC := C.bool(validate)
	C.WrapEndVerticesWithValidate(pointer.h, validateToC)
}

// Clear ...
func (pointer *Vertices) Clear() {
	C.WrapClearVertices(pointer.h)
}

// Reserve ...
func (pointer *Vertices) Reserve(count int32) {
	countToC := C.size_t(count)
	C.WrapReserveVertices(pointer.h, countToC)
}

// Resize ...
func (pointer *Vertices) Resize(count int32) {
	countToC := C.size_t(count)
	C.WrapResizeVertices(pointer.h, countToC)
}

// GetData ...
func (pointer *Vertices) GetData() *VoidPointer {
	retval := C.WrapGetDataVertices(pointer.h)
	var retvalGO *VoidPointer
	if retval != nil {
		retvalGO = &VoidPointer{h: retval}
	}
	return retvalGO
}

// GetSize ...
func (pointer *Vertices) GetSize() int32 {
	retval := C.WrapGetSizeVertices(pointer.h)
	return int32(retval)
}

// GetCount ...
func (pointer *Vertices) GetCount() int32 {
	retval := C.WrapGetCountVertices(pointer.h)
	return int32(retval)
}

// GetCapacity ...
func (pointer *Vertices) GetCapacity() int32 {
	retval := C.WrapGetCapacityVertices(pointer.h)
	return int32(retval)
}

// Pipeline  Rendering pipeline base class.
type Pipeline struct {
	h C.WrapPipeline
}

// Free ...
func (pointer *Pipeline) Free() {
	C.WrapPipelineFree(pointer.h)
}

// IsNil ...
func (pointer *Pipeline) IsNil() bool {
	return pointer.h == C.WrapPipeline(nil)
}

// PipelineInfo  ...
type PipelineInfo struct {
	h C.WrapPipelineInfo
}

// GetName ...
func (pointer *PipelineInfo) GetName() string {
	v := C.WrapPipelineInfoGetName(pointer.h)
	return C.GoString(v)
}

// SetName ...
func (pointer *PipelineInfo) SetName(v string) {
	vToC, idFinvToC := wrapString(v)
	defer idFinvToC()
	C.WrapPipelineInfoSetName(pointer.h, vToC)
}

// Free ...
func (pointer *PipelineInfo) Free() {
	C.WrapPipelineInfoFree(pointer.h)
}

// IsNil ...
func (pointer *PipelineInfo) IsNil() bool {
	return pointer.h == C.WrapPipelineInfo(nil)
}

// ForwardPipeline  Rendering pipeline implementing a forward rendering strategy.  The main characteristics of this pipeline are:  - Render in two passes: opaque display lists then transparent ones. - Fixed 8 light slots supporting 1 linear light with PSSM shadow mapping, 1 spot with shadow mapping and up to 6 point lights with no shadow mapping.
type ForwardPipeline struct {
	h C.WrapForwardPipeline
}

// Free ...
func (pointer *ForwardPipeline) Free() {
	C.WrapForwardPipelineFree(pointer.h)
}

// IsNil ...
func (pointer *ForwardPipeline) IsNil() bool {
	return pointer.h == C.WrapForwardPipeline(nil)
}

// ForwardPipelineLight  Single light for the forward pipeline. The complete lighting rig is passed as a [harfang.ForwardPipelineLights], see [harfang.PrepareForwardPipelineLights].
type ForwardPipelineLight struct {
	h C.WrapForwardPipelineLight
}

// GetType ...
func (pointer *ForwardPipelineLight) GetType() ForwardPipelineLightType {
	v := C.WrapForwardPipelineLightGetType(pointer.h)
	return ForwardPipelineLightType(v)
}

// SetType ...
func (pointer *ForwardPipelineLight) SetType(v ForwardPipelineLightType) {
	vToC := C.int32_t(v)
	C.WrapForwardPipelineLightSetType(pointer.h, vToC)
}

// GetWorld ...
func (pointer *ForwardPipelineLight) GetWorld() *Mat4 {
	v := C.WrapForwardPipelineLightGetWorld(pointer.h)
	vGO := &Mat4{h: v}
	return vGO
}

// SetWorld ...
func (pointer *ForwardPipelineLight) SetWorld(v *Mat4) {
	vToC := v.h
	C.WrapForwardPipelineLightSetWorld(pointer.h, vToC)
}

// GetDiffuse ...
func (pointer *ForwardPipelineLight) GetDiffuse() *Color {
	v := C.WrapForwardPipelineLightGetDiffuse(pointer.h)
	vGO := &Color{h: v}
	return vGO
}

// SetDiffuse ...
func (pointer *ForwardPipelineLight) SetDiffuse(v *Color) {
	vToC := v.h
	C.WrapForwardPipelineLightSetDiffuse(pointer.h, vToC)
}

// GetSpecular ...
func (pointer *ForwardPipelineLight) GetSpecular() *Color {
	v := C.WrapForwardPipelineLightGetSpecular(pointer.h)
	vGO := &Color{h: v}
	return vGO
}

// SetSpecular ...
func (pointer *ForwardPipelineLight) SetSpecular(v *Color) {
	vToC := v.h
	C.WrapForwardPipelineLightSetSpecular(pointer.h, vToC)
}

// GetRadius ...
func (pointer *ForwardPipelineLight) GetRadius() float32 {
	v := C.WrapForwardPipelineLightGetRadius(pointer.h)
	return float32(v)
}

// SetRadius ...
func (pointer *ForwardPipelineLight) SetRadius(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineLightSetRadius(pointer.h, vToC)
}

// GetInnerAngle ...
func (pointer *ForwardPipelineLight) GetInnerAngle() float32 {
	v := C.WrapForwardPipelineLightGetInnerAngle(pointer.h)
	return float32(v)
}

// SetInnerAngle ...
func (pointer *ForwardPipelineLight) SetInnerAngle(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineLightSetInnerAngle(pointer.h, vToC)
}

// GetOuterAngle ...
func (pointer *ForwardPipelineLight) GetOuterAngle() float32 {
	v := C.WrapForwardPipelineLightGetOuterAngle(pointer.h)
	return float32(v)
}

// SetOuterAngle ...
func (pointer *ForwardPipelineLight) SetOuterAngle(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineLightSetOuterAngle(pointer.h, vToC)
}

// GetPssmSplit ...
func (pointer *ForwardPipelineLight) GetPssmSplit() *Vec4 {
	v := C.WrapForwardPipelineLightGetPssmSplit(pointer.h)
	vGO := &Vec4{h: v}
	return vGO
}

// SetPssmSplit ...
func (pointer *ForwardPipelineLight) SetPssmSplit(v *Vec4) {
	vToC := v.h
	C.WrapForwardPipelineLightSetPssmSplit(pointer.h, vToC)
}

// GetPriority ...
func (pointer *ForwardPipelineLight) GetPriority() float32 {
	v := C.WrapForwardPipelineLightGetPriority(pointer.h)
	return float32(v)
}

// SetPriority ...
func (pointer *ForwardPipelineLight) SetPriority(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineLightSetPriority(pointer.h, vToC)
}

// NewForwardPipelineLight Single light for the forward pipeline. The complete lighting rig is passed as a [harfang.ForwardPipelineLights], see [harfang.PrepareForwardPipelineLights].
func NewForwardPipelineLight() *ForwardPipelineLight {
	retval := C.WrapConstructorForwardPipelineLight()
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *ForwardPipelineLight) Free() {
	C.WrapForwardPipelineLightFree(pointer.h)
}

// IsNil ...
func (pointer *ForwardPipelineLight) IsNil() bool {
	return pointer.h == C.WrapForwardPipelineLight(nil)
}

// GoSliceOfForwardPipelineLight ...
type GoSliceOfForwardPipelineLight []*ForwardPipelineLight

// ForwardPipelineLightList  ...
type ForwardPipelineLightList struct {
	h C.WrapForwardPipelineLightList
}

// Get ...
func (pointer *ForwardPipelineLightList) Get(id int) *ForwardPipelineLight {
	v := C.WrapForwardPipelineLightListGetOperator(pointer.h, C.int(id))
	vGO := &ForwardPipelineLight{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *ForwardPipelineLightList) Set(id int, v *ForwardPipelineLight) {
	vToC := v.h
	C.WrapForwardPipelineLightListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *ForwardPipelineLightList) Len() int32 {
	return int32(C.WrapForwardPipelineLightListLenOperator(pointer.h))
}

// NewForwardPipelineLightList ...
func NewForwardPipelineLightList() *ForwardPipelineLightList {
	retval := C.WrapConstructorForwardPipelineLightList()
	retvalGO := &ForwardPipelineLightList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLightList) {
		C.WrapForwardPipelineLightListFree(cleanval.h)
	})
	return retvalGO
}

// NewForwardPipelineLightListWithSequence ...
func NewForwardPipelineLightListWithSequence(sequence GoSliceOfForwardPipelineLight) *ForwardPipelineLightList {
	var sequencePointer []C.WrapForwardPipelineLight
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapForwardPipelineLight)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorForwardPipelineLightListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &ForwardPipelineLightList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLightList) {
		C.WrapForwardPipelineLightListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *ForwardPipelineLightList) Free() {
	C.WrapForwardPipelineLightListFree(pointer.h)
}

// IsNil ...
func (pointer *ForwardPipelineLightList) IsNil() bool {
	return pointer.h == C.WrapForwardPipelineLightList(nil)
}

// Clear ...
func (pointer *ForwardPipelineLightList) Clear() {
	C.WrapClearForwardPipelineLightList(pointer.h)
}

// Reserve ...
func (pointer *ForwardPipelineLightList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveForwardPipelineLightList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *ForwardPipelineLightList) PushBack(v *ForwardPipelineLight) {
	vToC := v.h
	C.WrapPushBackForwardPipelineLightList(pointer.h, vToC)
}

// Size ...
func (pointer *ForwardPipelineLightList) Size() int32 {
	retval := C.WrapSizeForwardPipelineLightList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *ForwardPipelineLightList) At(idx int32) *ForwardPipelineLight {
	idxToC := C.size_t(idx)
	retval := C.WrapAtForwardPipelineLightList(pointer.h, idxToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// ForwardPipelineLights  ...
type ForwardPipelineLights struct {
	h C.WrapForwardPipelineLights
}

// Free ...
func (pointer *ForwardPipelineLights) Free() {
	C.WrapForwardPipelineLightsFree(pointer.h)
}

// IsNil ...
func (pointer *ForwardPipelineLights) IsNil() bool {
	return pointer.h == C.WrapForwardPipelineLights(nil)
}

// ForwardPipelineFog  Fog properties for the forward pipeline.
type ForwardPipelineFog struct {
	h C.WrapForwardPipelineFog
}

// GetNear ...
func (pointer *ForwardPipelineFog) GetNear() float32 {
	v := C.WrapForwardPipelineFogGetNear(pointer.h)
	return float32(v)
}

// SetNear ...
func (pointer *ForwardPipelineFog) SetNear(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineFogSetNear(pointer.h, vToC)
}

// GetFar ...
func (pointer *ForwardPipelineFog) GetFar() float32 {
	v := C.WrapForwardPipelineFogGetFar(pointer.h)
	return float32(v)
}

// SetFar ...
func (pointer *ForwardPipelineFog) SetFar(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineFogSetFar(pointer.h, vToC)
}

// GetColor ...
func (pointer *ForwardPipelineFog) GetColor() *Color {
	v := C.WrapForwardPipelineFogGetColor(pointer.h)
	vGO := &Color{h: v}
	return vGO
}

// SetColor ...
func (pointer *ForwardPipelineFog) SetColor(v *Color) {
	vToC := v.h
	C.WrapForwardPipelineFogSetColor(pointer.h, vToC)
}

// NewForwardPipelineFog Fog properties for the forward pipeline.
func NewForwardPipelineFog() *ForwardPipelineFog {
	retval := C.WrapConstructorForwardPipelineFog()
	retvalGO := &ForwardPipelineFog{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineFog) {
		C.WrapForwardPipelineFogFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *ForwardPipelineFog) Free() {
	C.WrapForwardPipelineFogFree(pointer.h)
}

// IsNil ...
func (pointer *ForwardPipelineFog) IsNil() bool {
	return pointer.h == C.WrapForwardPipelineFog(nil)
}

// Font  Font object for realtime rendering.
type Font struct {
	h C.WrapFont
}

// Free ...
func (pointer *Font) Free() {
	C.WrapFontFree(pointer.h)
}

// IsNil ...
func (pointer *Font) IsNil() bool {
	return pointer.h == C.WrapFont(nil)
}

// JSON  JSON read/write object.
type JSON struct {
	h C.WrapJSON
}

// NewJSON JSON read/write object.
func NewJSON() *JSON {
	retval := C.WrapConstructorJSON()
	retvalGO := &JSON{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *JSON) {
		C.WrapJSONFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *JSON) Free() {
	C.WrapJSONFree(pointer.h)
}

// IsNil ...
func (pointer *JSON) IsNil() bool {
	return pointer.h == C.WrapJSON(nil)
}

// LuaObject  Opaque reference to an Lua object. This type is used to transfer values between VMs, see [harfang.man.Scripting].
type LuaObject struct {
	h C.WrapLuaObject
}

// Free ...
func (pointer *LuaObject) Free() {
	C.WrapLuaObjectFree(pointer.h)
}

// IsNil ...
func (pointer *LuaObject) IsNil() bool {
	return pointer.h == C.WrapLuaObject(nil)
}

// GoSliceOfLuaObject ...
type GoSliceOfLuaObject []*LuaObject

// LuaObjectList  ...
type LuaObjectList struct {
	h C.WrapLuaObjectList
}

// Get ...
func (pointer *LuaObjectList) Get(id int) *LuaObject {
	v := C.WrapLuaObjectListGetOperator(pointer.h, C.int(id))
	vGO := &LuaObject{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *LuaObject) {
		C.WrapLuaObjectFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *LuaObjectList) Set(id int, v *LuaObject) {
	vToC := v.h
	C.WrapLuaObjectListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *LuaObjectList) Len() int32 {
	return int32(C.WrapLuaObjectListLenOperator(pointer.h))
}

// NewLuaObjectList ...
func NewLuaObjectList() *LuaObjectList {
	retval := C.WrapConstructorLuaObjectList()
	retvalGO := &LuaObjectList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *LuaObjectList) {
		C.WrapLuaObjectListFree(cleanval.h)
	})
	return retvalGO
}

// NewLuaObjectListWithSequence ...
func NewLuaObjectListWithSequence(sequence GoSliceOfLuaObject) *LuaObjectList {
	var sequencePointer []C.WrapLuaObject
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapLuaObject)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorLuaObjectListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &LuaObjectList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *LuaObjectList) {
		C.WrapLuaObjectListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *LuaObjectList) Free() {
	C.WrapLuaObjectListFree(pointer.h)
}

// IsNil ...
func (pointer *LuaObjectList) IsNil() bool {
	return pointer.h == C.WrapLuaObjectList(nil)
}

// Clear ...
func (pointer *LuaObjectList) Clear() {
	C.WrapClearLuaObjectList(pointer.h)
}

// Reserve ...
func (pointer *LuaObjectList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveLuaObjectList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *LuaObjectList) PushBack(v *LuaObject) {
	vToC := v.h
	C.WrapPushBackLuaObjectList(pointer.h, vToC)
}

// Size ...
func (pointer *LuaObjectList) Size() int32 {
	retval := C.WrapSizeLuaObjectList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *LuaObjectList) At(idx int32) *LuaObject {
	idxToC := C.size_t(idx)
	retval := C.WrapAtLuaObjectList(pointer.h, idxToC)
	retvalGO := &LuaObject{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *LuaObject) {
		C.WrapLuaObjectFree(cleanval.h)
	})
	return retvalGO
}

// SceneAnimRef  Reference to a scene animation.
type SceneAnimRef struct {
	h C.WrapSceneAnimRef
}

// Free ...
func (pointer *SceneAnimRef) Free() {
	C.WrapSceneAnimRefFree(pointer.h)
}

// IsNil ...
func (pointer *SceneAnimRef) IsNil() bool {
	return pointer.h == C.WrapSceneAnimRef(nil)
}

// Eq ...
func (pointer *SceneAnimRef) Eq(ref *SceneAnimRef) bool {
	refToC := ref.h
	retval := C.WrapEqSceneAnimRef(pointer.h, refToC)
	return bool(retval)
}

// Ne ...
func (pointer *SceneAnimRef) Ne(ref *SceneAnimRef) bool {
	refToC := ref.h
	retval := C.WrapNeSceneAnimRef(pointer.h, refToC)
	return bool(retval)
}

// GoSliceOfSceneAnimRef ...
type GoSliceOfSceneAnimRef []*SceneAnimRef

// SceneAnimRefList  ...
type SceneAnimRefList struct {
	h C.WrapSceneAnimRefList
}

// Get ...
func (pointer *SceneAnimRefList) Get(id int) *SceneAnimRef {
	v := C.WrapSceneAnimRefListGetOperator(pointer.h, C.int(id))
	vGO := &SceneAnimRef{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *SceneAnimRef) {
		C.WrapSceneAnimRefFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *SceneAnimRefList) Set(id int, v *SceneAnimRef) {
	vToC := v.h
	C.WrapSceneAnimRefListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *SceneAnimRefList) Len() int32 {
	return int32(C.WrapSceneAnimRefListLenOperator(pointer.h))
}

// NewSceneAnimRefList ...
func NewSceneAnimRefList() *SceneAnimRefList {
	retval := C.WrapConstructorSceneAnimRefList()
	retvalGO := &SceneAnimRefList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneAnimRefList) {
		C.WrapSceneAnimRefListFree(cleanval.h)
	})
	return retvalGO
}

// NewSceneAnimRefListWithSequence ...
func NewSceneAnimRefListWithSequence(sequence GoSliceOfSceneAnimRef) *SceneAnimRefList {
	var sequencePointer []C.WrapSceneAnimRef
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapSceneAnimRef)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorSceneAnimRefListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &SceneAnimRefList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneAnimRefList) {
		C.WrapSceneAnimRefListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *SceneAnimRefList) Free() {
	C.WrapSceneAnimRefListFree(pointer.h)
}

// IsNil ...
func (pointer *SceneAnimRefList) IsNil() bool {
	return pointer.h == C.WrapSceneAnimRefList(nil)
}

// Clear ...
func (pointer *SceneAnimRefList) Clear() {
	C.WrapClearSceneAnimRefList(pointer.h)
}

// Reserve ...
func (pointer *SceneAnimRefList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveSceneAnimRefList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *SceneAnimRefList) PushBack(v *SceneAnimRef) {
	vToC := v.h
	C.WrapPushBackSceneAnimRefList(pointer.h, vToC)
}

// Size ...
func (pointer *SceneAnimRefList) Size() int32 {
	retval := C.WrapSizeSceneAnimRefList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *SceneAnimRefList) At(idx int32) *SceneAnimRef {
	idxToC := C.size_t(idx)
	retval := C.WrapAtSceneAnimRefList(pointer.h, idxToC)
	retvalGO := &SceneAnimRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneAnimRef) {
		C.WrapSceneAnimRefFree(cleanval.h)
	})
	return retvalGO
}

// ScenePlayAnimRef  Reference to a playing scene animation.
type ScenePlayAnimRef struct {
	h C.WrapScenePlayAnimRef
}

// Free ...
func (pointer *ScenePlayAnimRef) Free() {
	C.WrapScenePlayAnimRefFree(pointer.h)
}

// IsNil ...
func (pointer *ScenePlayAnimRef) IsNil() bool {
	return pointer.h == C.WrapScenePlayAnimRef(nil)
}

// Eq ...
func (pointer *ScenePlayAnimRef) Eq(ref *ScenePlayAnimRef) bool {
	refToC := ref.h
	retval := C.WrapEqScenePlayAnimRef(pointer.h, refToC)
	return bool(retval)
}

// Ne ...
func (pointer *ScenePlayAnimRef) Ne(ref *ScenePlayAnimRef) bool {
	refToC := ref.h
	retval := C.WrapNeScenePlayAnimRef(pointer.h, refToC)
	return bool(retval)
}

// GoSliceOfScenePlayAnimRef ...
type GoSliceOfScenePlayAnimRef []*ScenePlayAnimRef

// ScenePlayAnimRefList  ...
type ScenePlayAnimRefList struct {
	h C.WrapScenePlayAnimRefList
}

// Get ...
func (pointer *ScenePlayAnimRefList) Get(id int) *ScenePlayAnimRef {
	v := C.WrapScenePlayAnimRefListGetOperator(pointer.h, C.int(id))
	vGO := &ScenePlayAnimRef{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *ScenePlayAnimRef) {
		C.WrapScenePlayAnimRefFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *ScenePlayAnimRefList) Set(id int, v *ScenePlayAnimRef) {
	vToC := v.h
	C.WrapScenePlayAnimRefListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *ScenePlayAnimRefList) Len() int32 {
	return int32(C.WrapScenePlayAnimRefListLenOperator(pointer.h))
}

// NewScenePlayAnimRefList ...
func NewScenePlayAnimRefList() *ScenePlayAnimRefList {
	retval := C.WrapConstructorScenePlayAnimRefList()
	retvalGO := &ScenePlayAnimRefList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScenePlayAnimRefList) {
		C.WrapScenePlayAnimRefListFree(cleanval.h)
	})
	return retvalGO
}

// NewScenePlayAnimRefListWithSequence ...
func NewScenePlayAnimRefListWithSequence(sequence GoSliceOfScenePlayAnimRef) *ScenePlayAnimRefList {
	var sequencePointer []C.WrapScenePlayAnimRef
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapScenePlayAnimRef)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorScenePlayAnimRefListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &ScenePlayAnimRefList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScenePlayAnimRefList) {
		C.WrapScenePlayAnimRefListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *ScenePlayAnimRefList) Free() {
	C.WrapScenePlayAnimRefListFree(pointer.h)
}

// IsNil ...
func (pointer *ScenePlayAnimRefList) IsNil() bool {
	return pointer.h == C.WrapScenePlayAnimRefList(nil)
}

// Clear ...
func (pointer *ScenePlayAnimRefList) Clear() {
	C.WrapClearScenePlayAnimRefList(pointer.h)
}

// Reserve ...
func (pointer *ScenePlayAnimRefList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveScenePlayAnimRefList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *ScenePlayAnimRefList) PushBack(v *ScenePlayAnimRef) {
	vToC := v.h
	C.WrapPushBackScenePlayAnimRefList(pointer.h, vToC)
}

// Size ...
func (pointer *ScenePlayAnimRefList) Size() int32 {
	retval := C.WrapSizeScenePlayAnimRefList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *ScenePlayAnimRefList) At(idx int32) *ScenePlayAnimRef {
	idxToC := C.size_t(idx)
	retval := C.WrapAtScenePlayAnimRefList(pointer.h, idxToC)
	retvalGO := &ScenePlayAnimRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScenePlayAnimRef) {
		C.WrapScenePlayAnimRefFree(cleanval.h)
	})
	return retvalGO
}

// Scene  A scene object representing a world populated with [harfang.Node], see [harfang.man.Scene].
type Scene struct {
	h C.WrapScene
}

// GetCanvas ...
func (pointer *Scene) GetCanvas() *Canvas {
	v := C.WrapSceneGetCanvas(pointer.h)
	vGO := &Canvas{h: v}
	return vGO
}

// SetCanvas ...
func (pointer *Scene) SetCanvas(v *Canvas) {
	vToC := v.h
	C.WrapSceneSetCanvas(pointer.h, vToC)
}

// GetEnvironment ...
func (pointer *Scene) GetEnvironment() *Environment {
	v := C.WrapSceneGetEnvironment(pointer.h)
	vGO := &Environment{h: v}
	return vGO
}

// SetEnvironment ...
func (pointer *Scene) SetEnvironment(v *Environment) {
	vToC := v.h
	C.WrapSceneSetEnvironment(pointer.h, vToC)
}

// NewScene A scene object representing a world populated with [harfang.Node], see [harfang.man.Scene].
func NewScene() *Scene {
	retval := C.WrapConstructorScene()
	retvalGO := &Scene{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Scene) {
		C.WrapSceneFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Scene) Free() {
	C.WrapSceneFree(pointer.h)
}

// IsNil ...
func (pointer *Scene) IsNil() bool {
	return pointer.h == C.WrapScene(nil)
}

// GetNode Get a node by name. For more complex queries see [harfang.Scene_GetNodeEx].
func (pointer *Scene) GetNode(name string) *Node {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapGetNodeScene(pointer.h, nameToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// GetNodeEx Get a node by its absolute path in the node hierarchy.  A node path is constructed as follow:  - Nodes are refered to by their name. - To address the child of a node, use the `/` delimiter between its parent name and the child name. - To address a node inside an instance component, use the `:` delimiter. - There is no limit on the number of delimiters you can use.  Examples:  Get the node named `child` parented to the `root` node.  ```python child = scene.GetNodeEx('root/child') ```  Get the node named `dummy` instantiated by the `root` node.  ```python dummy = my_scene.GetNodeEx('root:dummy') ```
func (pointer *Scene) GetNodeEx(path string) *Node {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapGetNodeExScene(pointer.h, pathToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// GetNodes Return all nodes in scene.
func (pointer *Scene) GetNodes() *NodeList {
	retval := C.WrapGetNodesScene(pointer.h)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// GetAllNodes ...
func (pointer *Scene) GetAllNodes() *NodeList {
	retval := C.WrapGetAllNodesScene(pointer.h)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// GetNodesWithComponent ...
func (pointer *Scene) GetNodesWithComponent(idx NodeComponentIdx) *NodeList {
	idxToC := C.int32_t(idx)
	retval := C.WrapGetNodesWithComponentScene(pointer.h, idxToC)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// GetAllNodesWithComponent ...
func (pointer *Scene) GetAllNodesWithComponent(idx NodeComponentIdx) *NodeList {
	idxToC := C.int32_t(idx)
	retval := C.WrapGetAllNodesWithComponentScene(pointer.h, idxToC)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// GetNodeCount Return the number of nodes in the scene excluding instantiated ones. To include those, use [harfang.Scene_GetAllNodeCount].
func (pointer *Scene) GetNodeCount() int32 {
	retval := C.WrapGetNodeCountScene(pointer.h)
	return int32(retval)
}

// GetAllNodeCount Return the total number of nodes in the scene including instantiated ones. To exclude those, use [harfang.Scene_GetNodeCount].
func (pointer *Scene) GetAllNodeCount() int32 {
	retval := C.WrapGetAllNodeCountScene(pointer.h)
	return int32(retval)
}

// GetNodeChildren Return all children for a given node.
func (pointer *Scene) GetNodeChildren(node *Node) *NodeList {
	nodeToC := node.h
	retval := C.WrapGetNodeChildrenScene(pointer.h, nodeToC)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// IsChildOf ...
func (pointer *Scene) IsChildOf(node *Node, parent *Node) bool {
	nodeToC := node.h
	parentToC := parent.h
	retval := C.WrapIsChildOfScene(pointer.h, nodeToC, parentToC)
	return bool(retval)
}

// IsRoot ...
func (pointer *Scene) IsRoot(node *Node) bool {
	nodeToC := node.h
	retval := C.WrapIsRootScene(pointer.h, nodeToC)
	return bool(retval)
}

// ReadyWorldMatrices ...
func (pointer *Scene) ReadyWorldMatrices() {
	C.WrapReadyWorldMatricesScene(pointer.h)
}

// ComputeWorldMatrices ...
func (pointer *Scene) ComputeWorldMatrices() {
	C.WrapComputeWorldMatricesScene(pointer.h)
}

// Update Start the _Update_ phase of the scene.
func (pointer *Scene) Update(dt int64) {
	dtToC := C.int64_t(dt)
	C.WrapUpdateScene(pointer.h, dtToC)
}

// GetSceneAnims ...
func (pointer *Scene) GetSceneAnims() *SceneAnimRefList {
	retval := C.WrapGetSceneAnimsScene(pointer.h)
	retvalGO := &SceneAnimRefList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneAnimRefList) {
		C.WrapSceneAnimRefListFree(cleanval.h)
	})
	return retvalGO
}

// GetSceneAnim ...
func (pointer *Scene) GetSceneAnim(name string) *SceneAnimRef {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapGetSceneAnimScene(pointer.h, nameToC)
	retvalGO := &SceneAnimRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneAnimRef) {
		C.WrapSceneAnimRefFree(cleanval.h)
	})
	return retvalGO
}

// PlayAnim ...
func (pointer *Scene) PlayAnim(ref *SceneAnimRef) *ScenePlayAnimRef {
	refToC := ref.h
	retval := C.WrapPlayAnimScene(pointer.h, refToC)
	retvalGO := &ScenePlayAnimRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScenePlayAnimRef) {
		C.WrapScenePlayAnimRefFree(cleanval.h)
	})
	return retvalGO
}

// PlayAnimWithLoopMode ...
func (pointer *Scene) PlayAnimWithLoopMode(ref *SceneAnimRef, loopmode AnimLoopMode) *ScenePlayAnimRef {
	refToC := ref.h
	loopmodeToC := C.int32_t(loopmode)
	retval := C.WrapPlayAnimSceneWithLoopMode(pointer.h, refToC, loopmodeToC)
	retvalGO := &ScenePlayAnimRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScenePlayAnimRef) {
		C.WrapScenePlayAnimRefFree(cleanval.h)
	})
	return retvalGO
}

// PlayAnimWithLoopModeEasing ...
func (pointer *Scene) PlayAnimWithLoopModeEasing(ref *SceneAnimRef, loopmode AnimLoopMode, easing Easing) *ScenePlayAnimRef {
	refToC := ref.h
	loopmodeToC := C.int32_t(loopmode)
	easingToC := C.uchar(easing)
	retval := C.WrapPlayAnimSceneWithLoopModeEasing(pointer.h, refToC, loopmodeToC, easingToC)
	retvalGO := &ScenePlayAnimRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScenePlayAnimRef) {
		C.WrapScenePlayAnimRefFree(cleanval.h)
	})
	return retvalGO
}

// PlayAnimWithLoopModeEasingTStart ...
func (pointer *Scene) PlayAnimWithLoopModeEasingTStart(ref *SceneAnimRef, loopmode AnimLoopMode, easing Easing, tstart int64) *ScenePlayAnimRef {
	refToC := ref.h
	loopmodeToC := C.int32_t(loopmode)
	easingToC := C.uchar(easing)
	tstartToC := C.int64_t(tstart)
	retval := C.WrapPlayAnimSceneWithLoopModeEasingTStart(pointer.h, refToC, loopmodeToC, easingToC, tstartToC)
	retvalGO := &ScenePlayAnimRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScenePlayAnimRef) {
		C.WrapScenePlayAnimRefFree(cleanval.h)
	})
	return retvalGO
}

// PlayAnimWithLoopModeEasingTStartTEnd ...
func (pointer *Scene) PlayAnimWithLoopModeEasingTStartTEnd(ref *SceneAnimRef, loopmode AnimLoopMode, easing Easing, tstart int64, tend int64) *ScenePlayAnimRef {
	refToC := ref.h
	loopmodeToC := C.int32_t(loopmode)
	easingToC := C.uchar(easing)
	tstartToC := C.int64_t(tstart)
	tendToC := C.int64_t(tend)
	retval := C.WrapPlayAnimSceneWithLoopModeEasingTStartTEnd(pointer.h, refToC, loopmodeToC, easingToC, tstartToC, tendToC)
	retvalGO := &ScenePlayAnimRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScenePlayAnimRef) {
		C.WrapScenePlayAnimRefFree(cleanval.h)
	})
	return retvalGO
}

// PlayAnimWithLoopModeEasingTStartTEndPaused ...
func (pointer *Scene) PlayAnimWithLoopModeEasingTStartTEndPaused(ref *SceneAnimRef, loopmode AnimLoopMode, easing Easing, tstart int64, tend int64, paused bool) *ScenePlayAnimRef {
	refToC := ref.h
	loopmodeToC := C.int32_t(loopmode)
	easingToC := C.uchar(easing)
	tstartToC := C.int64_t(tstart)
	tendToC := C.int64_t(tend)
	pausedToC := C.bool(paused)
	retval := C.WrapPlayAnimSceneWithLoopModeEasingTStartTEndPaused(pointer.h, refToC, loopmodeToC, easingToC, tstartToC, tendToC, pausedToC)
	retvalGO := &ScenePlayAnimRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScenePlayAnimRef) {
		C.WrapScenePlayAnimRefFree(cleanval.h)
	})
	return retvalGO
}

// PlayAnimWithLoopModeEasingTStartTEndPausedTScale ...
func (pointer *Scene) PlayAnimWithLoopModeEasingTStartTEndPausedTScale(ref *SceneAnimRef, loopmode AnimLoopMode, easing Easing, tstart int64, tend int64, paused bool, tscale float32) *ScenePlayAnimRef {
	refToC := ref.h
	loopmodeToC := C.int32_t(loopmode)
	easingToC := C.uchar(easing)
	tstartToC := C.int64_t(tstart)
	tendToC := C.int64_t(tend)
	pausedToC := C.bool(paused)
	tscaleToC := C.float(tscale)
	retval := C.WrapPlayAnimSceneWithLoopModeEasingTStartTEndPausedTScale(pointer.h, refToC, loopmodeToC, easingToC, tstartToC, tendToC, pausedToC, tscaleToC)
	retvalGO := &ScenePlayAnimRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScenePlayAnimRef) {
		C.WrapScenePlayAnimRefFree(cleanval.h)
	})
	return retvalGO
}

// IsPlaying ...
func (pointer *Scene) IsPlaying(ref *ScenePlayAnimRef) bool {
	refToC := ref.h
	retval := C.WrapIsPlayingScene(pointer.h, refToC)
	return bool(retval)
}

// StopAnim ...
func (pointer *Scene) StopAnim(ref *ScenePlayAnimRef) {
	refToC := ref.h
	C.WrapStopAnimScene(pointer.h, refToC)
}

// StopAllAnims ...
func (pointer *Scene) StopAllAnims() {
	C.WrapStopAllAnimsScene(pointer.h)
}

// GetPlayingAnimNames ...
func (pointer *Scene) GetPlayingAnimNames() *StringList {
	retval := C.WrapGetPlayingAnimNamesScene(pointer.h)
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// GetPlayingAnimRefs ...
func (pointer *Scene) GetPlayingAnimRefs() *ScenePlayAnimRefList {
	retval := C.WrapGetPlayingAnimRefsScene(pointer.h)
	retvalGO := &ScenePlayAnimRefList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScenePlayAnimRefList) {
		C.WrapScenePlayAnimRefListFree(cleanval.h)
	})
	return retvalGO
}

// UpdatePlayingAnims ...
func (pointer *Scene) UpdatePlayingAnims(dt int64) {
	dtToC := C.int64_t(dt)
	C.WrapUpdatePlayingAnimsScene(pointer.h, dtToC)
}

// HasKey ...
func (pointer *Scene) HasKey(key string) bool {
	keyToC, idFinkeyToC := wrapString(key)
	defer idFinkeyToC()
	retval := C.WrapHasKeyScene(pointer.h, keyToC)
	return bool(retval)
}

// GetKeys ...
func (pointer *Scene) GetKeys() *StringList {
	retval := C.WrapGetKeysScene(pointer.h)
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// RemoveKey ...
func (pointer *Scene) RemoveKey(key string) {
	keyToC, idFinkeyToC := wrapString(key)
	defer idFinkeyToC()
	C.WrapRemoveKeyScene(pointer.h, keyToC)
}

// GetValue ...
func (pointer *Scene) GetValue(key string) string {
	keyToC, idFinkeyToC := wrapString(key)
	defer idFinkeyToC()
	retval := C.WrapGetValueScene(pointer.h, keyToC)
	return C.GoString(retval)
}

// SetValue ...
func (pointer *Scene) SetValue(key string, value string) {
	keyToC, idFinkeyToC := wrapString(key)
	defer idFinkeyToC()
	valueToC, idFinvalueToC := wrapString(value)
	defer idFinvalueToC()
	C.WrapSetValueScene(pointer.h, keyToC, valueToC)
}

// GarbageCollect Destroy any unreferenced components in the scene.
func (pointer *Scene) GarbageCollect() int32 {
	retval := C.WrapGarbageCollectScene(pointer.h)
	return int32(retval)
}

// Clear Remove all nodes from the scene.
func (pointer *Scene) Clear() {
	C.WrapClearScene(pointer.h)
}

// ReserveNodes Allocates internal storage for the required number of nodes in one go.
func (pointer *Scene) ReserveNodes(count int32) {
	countToC := C.size_t(count)
	C.WrapReserveNodesScene(pointer.h, countToC)
}

// CreateNode Create a [harfang.Node] in the scene.
func (pointer *Scene) CreateNode() *Node {
	retval := C.WrapCreateNodeScene(pointer.h)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateNodeWithName Create a [harfang.Node] in the scene.
func (pointer *Scene) CreateNodeWithName(name string) *Node {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapCreateNodeSceneWithName(pointer.h, nameToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// DestroyNode Remove a [harfang.Node] from the scene.
func (pointer *Scene) DestroyNode(node *Node) {
	nodeToC := node.h
	C.WrapDestroyNodeScene(pointer.h, nodeToC)
}

// ReserveTransforms Allocates internal storage for the required number of [harfang.Transform] components in one go.
func (pointer *Scene) ReserveTransforms(count int32) {
	countToC := C.size_t(count)
	C.WrapReserveTransformsScene(pointer.h, countToC)
}

// CreateTransform ...
func (pointer *Scene) CreateTransform() *Transform {
	retval := C.WrapCreateTransformScene(pointer.h)
	retvalGO := &Transform{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Transform) {
		C.WrapTransformFree(cleanval.h)
	})
	return retvalGO
}

// CreateTransformWithT ...
func (pointer *Scene) CreateTransformWithT(T *Vec3) *Transform {
	TToC := T.h
	retval := C.WrapCreateTransformSceneWithT(pointer.h, TToC)
	retvalGO := &Transform{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Transform) {
		C.WrapTransformFree(cleanval.h)
	})
	return retvalGO
}

// CreateTransformWithTR ...
func (pointer *Scene) CreateTransformWithTR(T *Vec3, R *Vec3) *Transform {
	TToC := T.h
	RToC := R.h
	retval := C.WrapCreateTransformSceneWithTR(pointer.h, TToC, RToC)
	retvalGO := &Transform{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Transform) {
		C.WrapTransformFree(cleanval.h)
	})
	return retvalGO
}

// CreateTransformWithTRS ...
func (pointer *Scene) CreateTransformWithTRS(T *Vec3, R *Vec3, S *Vec3) *Transform {
	TToC := T.h
	RToC := R.h
	SToC := S.h
	retval := C.WrapCreateTransformSceneWithTRS(pointer.h, TToC, RToC, SToC)
	retvalGO := &Transform{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Transform) {
		C.WrapTransformFree(cleanval.h)
	})
	return retvalGO
}

// DestroyTransform ...
func (pointer *Scene) DestroyTransform(transform *Transform) {
	transformToC := transform.h
	C.WrapDestroyTransformScene(pointer.h, transformToC)
}

// ReserveCameras ...
func (pointer *Scene) ReserveCameras(count int32) {
	countToC := C.size_t(count)
	C.WrapReserveCamerasScene(pointer.h, countToC)
}

// CreateCamera ...
func (pointer *Scene) CreateCamera() *Camera {
	retval := C.WrapCreateCameraScene(pointer.h)
	retvalGO := &Camera{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Camera) {
		C.WrapCameraFree(cleanval.h)
	})
	return retvalGO
}

// CreateCameraWithZnearZfar ...
func (pointer *Scene) CreateCameraWithZnearZfar(znear float32, zfar float32) *Camera {
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	retval := C.WrapCreateCameraSceneWithZnearZfar(pointer.h, znearToC, zfarToC)
	retvalGO := &Camera{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Camera) {
		C.WrapCameraFree(cleanval.h)
	})
	return retvalGO
}

// CreateCameraWithZnearZfarFov ...
func (pointer *Scene) CreateCameraWithZnearZfarFov(znear float32, zfar float32, fov float32) *Camera {
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	fovToC := C.float(fov)
	retval := C.WrapCreateCameraSceneWithZnearZfarFov(pointer.h, znearToC, zfarToC, fovToC)
	retvalGO := &Camera{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Camera) {
		C.WrapCameraFree(cleanval.h)
	})
	return retvalGO
}

// CreateOrthographicCamera ...
func (pointer *Scene) CreateOrthographicCamera(znear float32, zfar float32) *Camera {
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	retval := C.WrapCreateOrthographicCameraScene(pointer.h, znearToC, zfarToC)
	retvalGO := &Camera{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Camera) {
		C.WrapCameraFree(cleanval.h)
	})
	return retvalGO
}

// CreateOrthographicCameraWithSize ...
func (pointer *Scene) CreateOrthographicCameraWithSize(znear float32, zfar float32, size float32) *Camera {
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	sizeToC := C.float(size)
	retval := C.WrapCreateOrthographicCameraSceneWithSize(pointer.h, znearToC, zfarToC, sizeToC)
	retvalGO := &Camera{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Camera) {
		C.WrapCameraFree(cleanval.h)
	})
	return retvalGO
}

// DestroyCamera ...
func (pointer *Scene) DestroyCamera(camera *Camera) {
	cameraToC := camera.h
	C.WrapDestroyCameraScene(pointer.h, cameraToC)
}

// ComputeCurrentCameraViewState ...
func (pointer *Scene) ComputeCurrentCameraViewState(aspectratio *Vec2) *ViewState {
	aspectratioToC := aspectratio.h
	retval := C.WrapComputeCurrentCameraViewStateScene(pointer.h, aspectratioToC)
	retvalGO := &ViewState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ViewState) {
		C.WrapViewStateFree(cleanval.h)
	})
	return retvalGO
}

// ReserveObjects Allocates internal storage for the required number of [harfang.Object] components in one go.
func (pointer *Scene) ReserveObjects(count int32) {
	countToC := C.size_t(count)
	C.WrapReserveObjectsScene(pointer.h, countToC)
}

// CreateObject ...
func (pointer *Scene) CreateObject() *Object {
	retval := C.WrapCreateObjectScene(pointer.h)
	retvalGO := &Object{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Object) {
		C.WrapObjectFree(cleanval.h)
	})
	return retvalGO
}

// CreateObjectWithModelMaterials ...
func (pointer *Scene) CreateObjectWithModelMaterials(model *ModelRef, materials *MaterialList) *Object {
	modelToC := model.h
	materialsToC := materials.h
	retval := C.WrapCreateObjectSceneWithModelMaterials(pointer.h, modelToC, materialsToC)
	retvalGO := &Object{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Object) {
		C.WrapObjectFree(cleanval.h)
	})
	return retvalGO
}

// CreateObjectWithModelSliceOfMaterials ...
func (pointer *Scene) CreateObjectWithModelSliceOfMaterials(model *ModelRef, SliceOfmaterials GoSliceOfMaterial) *Object {
	modelToC := model.h
	var SliceOfmaterialsPointer []C.WrapMaterial
	for _, s := range SliceOfmaterials {
		SliceOfmaterialsPointer = append(SliceOfmaterialsPointer, s.h)
	}
	SliceOfmaterialsPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfmaterialsPointer))
	SliceOfmaterialsPointerToCSize := C.size_t(SliceOfmaterialsPointerToC.Len)
	SliceOfmaterialsPointerToCBuf := (*C.WrapMaterial)(unsafe.Pointer(SliceOfmaterialsPointerToC.Data))
	retval := C.WrapCreateObjectSceneWithModelSliceOfMaterials(pointer.h, modelToC, SliceOfmaterialsPointerToCSize, SliceOfmaterialsPointerToCBuf)
	retvalGO := &Object{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Object) {
		C.WrapObjectFree(cleanval.h)
	})
	return retvalGO
}

// DestroyObject ...
func (pointer *Scene) DestroyObject(object *Object) {
	objectToC := object.h
	C.WrapDestroyObjectScene(pointer.h, objectToC)
}

// ReserveLights Allocates internal storage for the required number of [harfang.Light] components in one go.
func (pointer *Scene) ReserveLights(count int32) {
	countToC := C.size_t(count)
	C.WrapReserveLightsScene(pointer.h, countToC)
}

// CreateLight ...
func (pointer *Scene) CreateLight() *Light {
	retval := C.WrapCreateLightScene(pointer.h)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// DestroyLight ...
func (pointer *Scene) DestroyLight(light *Light) {
	lightToC := light.h
	C.WrapDestroyLightScene(pointer.h, lightToC)
}

// CreateLinearLightWithDiffuseIntensitySpecularSpecularIntensity ...
func (pointer *Scene) CreateLinearLightWithDiffuseIntensitySpecularSpecularIntensity(diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32) *Light {
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	retval := C.WrapCreateLinearLightSceneWithDiffuseIntensitySpecularSpecularIntensity(pointer.h, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseIntensitySpecularSpecularIntensityPriority ...
func (pointer *Scene) CreateLinearLightWithDiffuseIntensitySpecularSpecularIntensityPriority(diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32) *Light {
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	retval := C.WrapCreateLinearLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriority(pointer.h, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowType ...
func (pointer *Scene) CreateLinearLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType) *Light {
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapCreateLinearLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(pointer.h, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias ...
func (pointer *Scene) CreateLinearLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType, shadowbias float32) *Light {
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapCreateLinearLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(pointer.h, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBiasPssmSplit ...
func (pointer *Scene) CreateLinearLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBiasPssmSplit(diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType, shadowbias float32, pssmsplit *Vec4) *Light {
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	pssmsplitToC := pssmsplit.h
	retval := C.WrapCreateLinearLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBiasPssmSplit(pointer.h, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC, shadowbiasToC, pssmsplitToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLight ...
func (pointer *Scene) CreateLinearLight(diffuse *Color, specular *Color) *Light {
	diffuseToC := diffuse.h
	specularToC := specular.h
	retval := C.WrapCreateLinearLightScene(pointer.h, diffuseToC, specularToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithPriority ...
func (pointer *Scene) CreateLinearLightWithPriority(diffuse *Color, specular *Color, priority float32) *Light {
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	retval := C.WrapCreateLinearLightSceneWithPriority(pointer.h, diffuseToC, specularToC, priorityToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithPriorityShadowType ...
func (pointer *Scene) CreateLinearLightWithPriorityShadowType(diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType) *Light {
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapCreateLinearLightSceneWithPriorityShadowType(pointer.h, diffuseToC, specularToC, priorityToC, shadowtypeToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithPriorityShadowTypeShadowBias ...
func (pointer *Scene) CreateLinearLightWithPriorityShadowTypeShadowBias(diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType, shadowbias float32) *Light {
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapCreateLinearLightSceneWithPriorityShadowTypeShadowBias(pointer.h, diffuseToC, specularToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithPriorityShadowTypeShadowBiasPssmSplit ...
func (pointer *Scene) CreateLinearLightWithPriorityShadowTypeShadowBiasPssmSplit(diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType, shadowbias float32, pssmsplit *Vec4) *Light {
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	pssmsplitToC := pssmsplit.h
	retval := C.WrapCreateLinearLightSceneWithPriorityShadowTypeShadowBiasPssmSplit(pointer.h, diffuseToC, specularToC, priorityToC, shadowtypeToC, shadowbiasToC, pssmsplitToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseIntensitySpecularSpecularIntensity ...
func (pointer *Scene) CreatePointLightWithDiffuseIntensitySpecularSpecularIntensity(radius float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32) *Light {
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	retval := C.WrapCreatePointLightSceneWithDiffuseIntensitySpecularSpecularIntensity(pointer.h, radiusToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseIntensitySpecularSpecularIntensityPriority ...
func (pointer *Scene) CreatePointLightWithDiffuseIntensitySpecularSpecularIntensityPriority(radius float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32) *Light {
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	retval := C.WrapCreatePointLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriority(pointer.h, radiusToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowType ...
func (pointer *Scene) CreatePointLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(radius float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType) *Light {
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapCreatePointLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(pointer.h, radiusToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias ...
func (pointer *Scene) CreatePointLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(radius float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType, shadowbias float32) *Light {
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapCreatePointLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(pointer.h, radiusToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLight ...
func (pointer *Scene) CreatePointLight(radius float32, diffuse *Color, specular *Color) *Light {
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	specularToC := specular.h
	retval := C.WrapCreatePointLightScene(pointer.h, radiusToC, diffuseToC, specularToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithPriority ...
func (pointer *Scene) CreatePointLightWithPriority(radius float32, diffuse *Color, specular *Color, priority float32) *Light {
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	retval := C.WrapCreatePointLightSceneWithPriority(pointer.h, radiusToC, diffuseToC, specularToC, priorityToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithPriorityShadowType ...
func (pointer *Scene) CreatePointLightWithPriorityShadowType(radius float32, diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType) *Light {
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapCreatePointLightSceneWithPriorityShadowType(pointer.h, radiusToC, diffuseToC, specularToC, priorityToC, shadowtypeToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithPriorityShadowTypeShadowBias ...
func (pointer *Scene) CreatePointLightWithPriorityShadowTypeShadowBias(radius float32, diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType, shadowbias float32) *Light {
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapCreatePointLightSceneWithPriorityShadowTypeShadowBias(pointer.h, radiusToC, diffuseToC, specularToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseIntensitySpecularSpecularIntensity ...
func (pointer *Scene) CreateSpotLightWithDiffuseIntensitySpecularSpecularIntensity(radius float32, innerangle float32, outerangle float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32) *Light {
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	retval := C.WrapCreateSpotLightSceneWithDiffuseIntensitySpecularSpecularIntensity(pointer.h, radiusToC, innerangleToC, outerangleToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseIntensitySpecularSpecularIntensityPriority ...
func (pointer *Scene) CreateSpotLightWithDiffuseIntensitySpecularSpecularIntensityPriority(radius float32, innerangle float32, outerangle float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32) *Light {
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	retval := C.WrapCreateSpotLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriority(pointer.h, radiusToC, innerangleToC, outerangleToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowType ...
func (pointer *Scene) CreateSpotLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(radius float32, innerangle float32, outerangle float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType) *Light {
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapCreateSpotLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(pointer.h, radiusToC, innerangleToC, outerangleToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias ...
func (pointer *Scene) CreateSpotLightWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(radius float32, innerangle float32, outerangle float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType, shadowbias float32) *Light {
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapCreateSpotLightSceneWithDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(pointer.h, radiusToC, innerangleToC, outerangleToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLight ...
func (pointer *Scene) CreateSpotLight(radius float32, innerangle float32, outerangle float32, diffuse *Color, specular *Color) *Light {
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	specularToC := specular.h
	retval := C.WrapCreateSpotLightScene(pointer.h, radiusToC, innerangleToC, outerangleToC, diffuseToC, specularToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithPriority ...
func (pointer *Scene) CreateSpotLightWithPriority(radius float32, innerangle float32, outerangle float32, diffuse *Color, specular *Color, priority float32) *Light {
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	retval := C.WrapCreateSpotLightSceneWithPriority(pointer.h, radiusToC, innerangleToC, outerangleToC, diffuseToC, specularToC, priorityToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithPriorityShadowType ...
func (pointer *Scene) CreateSpotLightWithPriorityShadowType(radius float32, innerangle float32, outerangle float32, diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType) *Light {
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapCreateSpotLightSceneWithPriorityShadowType(pointer.h, radiusToC, innerangleToC, outerangleToC, diffuseToC, specularToC, priorityToC, shadowtypeToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithPriorityShadowTypeShadowBias ...
func (pointer *Scene) CreateSpotLightWithPriorityShadowTypeShadowBias(radius float32, innerangle float32, outerangle float32, diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType, shadowbias float32) *Light {
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapCreateSpotLightSceneWithPriorityShadowTypeShadowBias(pointer.h, radiusToC, innerangleToC, outerangleToC, diffuseToC, specularToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// ReserveScripts ...
func (pointer *Scene) ReserveScripts(count int32) {
	countToC := C.size_t(count)
	C.WrapReserveScriptsScene(pointer.h, countToC)
}

// CreateScript ...
func (pointer *Scene) CreateScript() *Script {
	retval := C.WrapCreateScriptScene(pointer.h)
	retvalGO := &Script{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Script) {
		C.WrapScriptFree(cleanval.h)
	})
	return retvalGO
}

// CreateScriptWithPath ...
func (pointer *Scene) CreateScriptWithPath(path string) *Script {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapCreateScriptSceneWithPath(pointer.h, pathToC)
	retvalGO := &Script{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Script) {
		C.WrapScriptFree(cleanval.h)
	})
	return retvalGO
}

// DestroyScript ...
func (pointer *Scene) DestroyScript(script *Script) {
	scriptToC := script.h
	C.WrapDestroyScriptScene(pointer.h, scriptToC)
}

// GetScriptCount ...
func (pointer *Scene) GetScriptCount() int32 {
	retval := C.WrapGetScriptCountScene(pointer.h)
	return int32(retval)
}

// SetScript ...
func (pointer *Scene) SetScript(slotidx int32, script *Script) {
	slotidxToC := C.size_t(slotidx)
	scriptToC := script.h
	C.WrapSetScriptScene(pointer.h, slotidxToC, scriptToC)
}

// GetScript ...
func (pointer *Scene) GetScript(slotidx int32) *Script {
	slotidxToC := C.size_t(slotidx)
	retval := C.WrapGetScriptScene(pointer.h, slotidxToC)
	retvalGO := &Script{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Script) {
		C.WrapScriptFree(cleanval.h)
	})
	return retvalGO
}

// CreateRigidBody ...
func (pointer *Scene) CreateRigidBody() *RigidBody {
	retval := C.WrapCreateRigidBodyScene(pointer.h)
	retvalGO := &RigidBody{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RigidBody) {
		C.WrapRigidBodyFree(cleanval.h)
	})
	return retvalGO
}

// DestroyRigidBody ...
func (pointer *Scene) DestroyRigidBody(rigidbody *RigidBody) {
	rigidbodyToC := rigidbody.h
	C.WrapDestroyRigidBodyScene(pointer.h, rigidbodyToC)
}

// CreateCollision ...
func (pointer *Scene) CreateCollision() *Collision {
	retval := C.WrapCreateCollisionScene(pointer.h)
	retvalGO := &Collision{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Collision) {
		C.WrapCollisionFree(cleanval.h)
	})
	return retvalGO
}

// DestroyCollision ...
func (pointer *Scene) DestroyCollision(collision *Collision) {
	collisionToC := collision.h
	C.WrapDestroyCollisionScene(pointer.h, collisionToC)
}

// CreateInstance ...
func (pointer *Scene) CreateInstance() *Instance {
	retval := C.WrapCreateInstanceScene(pointer.h)
	retvalGO := &Instance{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Instance) {
		C.WrapInstanceFree(cleanval.h)
	})
	return retvalGO
}

// DestroyInstance ...
func (pointer *Scene) DestroyInstance(Instance *Instance) {
	InstanceToC := Instance.h
	C.WrapDestroyInstanceScene(pointer.h, InstanceToC)
}

// SetProbe ...
func (pointer *Scene) SetProbe(irradiance *TextureRef, radiance *TextureRef, brdf *TextureRef) {
	irradianceToC := irradiance.h
	radianceToC := radiance.h
	brdfToC := brdf.h
	C.WrapSetProbeScene(pointer.h, irradianceToC, radianceToC, brdfToC)
}

// GetCurrentCamera Get the current camera.
func (pointer *Scene) GetCurrentCamera() *Node {
	retval := C.WrapGetCurrentCameraScene(pointer.h)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// SetCurrentCamera Set the current camera.
func (pointer *Scene) SetCurrentCamera(camera *Node) {
	cameraToC := camera.h
	C.WrapSetCurrentCameraScene(pointer.h, cameraToC)
}

// GetMinMax ...
func (pointer *Scene) GetMinMax(resources *PipelineResources) (bool, *MinMax) {
	resourcesToC := resources.h
	minmax := NewMinMax()
	minmaxToC := minmax.h
	retval := C.WrapGetMinMaxScene(pointer.h, resourcesToC, minmaxToC)
	return bool(retval), minmax
}

// SceneView  Holds a view to a subset of a scene. Used by the instance system to track instantiated scene content.  See [harfang.Node_GetInstanceSceneView] and [harfang.man.Scene].
type SceneView struct {
	h C.WrapSceneView
}

// Free ...
func (pointer *SceneView) Free() {
	C.WrapSceneViewFree(pointer.h)
}

// IsNil ...
func (pointer *SceneView) IsNil() bool {
	return pointer.h == C.WrapSceneView(nil)
}

// GetNodes Return all nodes in the view. Pass the host scene as the `scene` parameter.  See [harfang.man.Scene].
func (pointer *SceneView) GetNodes(scene *Scene) *NodeList {
	sceneToC := scene.h
	retval := C.WrapGetNodesSceneView(pointer.h, sceneToC)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// GetNode Find a node by name in the view. Pass the host scene as the `scene` parameter.  See [harfang.man.Scene].
func (pointer *SceneView) GetNode(scene *Scene, name string) *Node {
	sceneToC := scene.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapGetNodeSceneView(pointer.h, sceneToC, nameToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// Node  The base element of a scene, see [harfang.man.Scene].
type Node struct {
	h C.WrapNode
}

// Free ...
func (pointer *Node) Free() {
	C.WrapNodeFree(pointer.h)
}

// IsNil ...
func (pointer *Node) IsNil() bool {
	return pointer.h == C.WrapNode(nil)
}

// Eq ...
func (pointer *Node) Eq(n *Node) bool {
	nToC := n.h
	retval := C.WrapEqNode(pointer.h, nToC)
	return bool(retval)
}

// IsValid Return `true` if the [harfang.Node] still exist.
func (pointer *Node) IsValid() bool {
	retval := C.WrapIsValidNode(pointer.h)
	return bool(retval)
}

// GetUid Return the unique ID.
func (pointer *Node) GetUid() uint32 {
	retval := C.WrapGetUidNode(pointer.h)
	return uint32(retval)
}

// GetFlags ...
func (pointer *Node) GetFlags() uint32 {
	retval := C.WrapGetFlagsNode(pointer.h)
	return uint32(retval)
}

// SetFlags ...
func (pointer *Node) SetFlags(flags uint32) {
	flagsToC := C.uint32_t(flags)
	C.WrapSetFlagsNode(pointer.h, flagsToC)
}

// Enable ...
func (pointer *Node) Enable() {
	C.WrapEnableNode(pointer.h)
}

// Disable ...
func (pointer *Node) Disable() {
	C.WrapDisableNode(pointer.h)
}

// IsEnabled ...
func (pointer *Node) IsEnabled() bool {
	retval := C.WrapIsEnabledNode(pointer.h)
	return bool(retval)
}

// IsItselfEnabled ...
func (pointer *Node) IsItselfEnabled() bool {
	retval := C.WrapIsItselfEnabledNode(pointer.h)
	return bool(retval)
}

// HasTransform Return `true` if the [harfang.Node] has a [harfang.Transform] component.
func (pointer *Node) HasTransform() bool {
	retval := C.WrapHasTransformNode(pointer.h)
	return bool(retval)
}

// GetTransform Return the [harfang.Transform] component of the node.
func (pointer *Node) GetTransform() *Transform {
	retval := C.WrapGetTransformNode(pointer.h)
	retvalGO := &Transform{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Transform) {
		C.WrapTransformFree(cleanval.h)
	})
	return retvalGO
}

// SetTransform Set the [harfang.Transform] component of a node.  See [harfang.Scene_CreateTransform].
func (pointer *Node) SetTransform(t *Transform) {
	tToC := t.h
	C.WrapSetTransformNode(pointer.h, tToC)
}

// RemoveTransform Remove the [harfang.Transform] component from the [harfang.Node].
func (pointer *Node) RemoveTransform() {
	C.WrapRemoveTransformNode(pointer.h)
}

// HasCamera Return `true` if the [harfang.Node] has a [harfang.Camera] component.
func (pointer *Node) HasCamera() bool {
	retval := C.WrapHasCameraNode(pointer.h)
	return bool(retval)
}

// GetCamera Return the [harfang.Camera] component of the node.
func (pointer *Node) GetCamera() *Camera {
	retval := C.WrapGetCameraNode(pointer.h)
	retvalGO := &Camera{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Camera) {
		C.WrapCameraFree(cleanval.h)
	})
	return retvalGO
}

// SetCamera Set the [harfang.Camera] component of a node. See [harfang.Scene_CreateCamera].
func (pointer *Node) SetCamera(c *Camera) {
	cToC := c.h
	C.WrapSetCameraNode(pointer.h, cToC)
}

// RemoveCamera Remove the [harfang.Camera] component from the [harfang.Node].
func (pointer *Node) RemoveCamera() {
	C.WrapRemoveCameraNode(pointer.h)
}

// ComputeCameraViewState ...
func (pointer *Node) ComputeCameraViewState(aspectratio *Vec2) *ViewState {
	aspectratioToC := aspectratio.h
	retval := C.WrapComputeCameraViewStateNode(pointer.h, aspectratioToC)
	retvalGO := &ViewState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ViewState) {
		C.WrapViewStateFree(cleanval.h)
	})
	return retvalGO
}

// HasObject Return `true` if the [harfang.Node] has an [harfang.Object] (a geometry) component.
func (pointer *Node) HasObject() bool {
	retval := C.WrapHasObjectNode(pointer.h)
	return bool(retval)
}

// GetObject Return [harfang.Object] component of the node.
func (pointer *Node) GetObject() *Object {
	retval := C.WrapGetObjectNode(pointer.h)
	retvalGO := &Object{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Object) {
		C.WrapObjectFree(cleanval.h)
	})
	return retvalGO
}

// SetObject Set the [harfang.Object] component of a node.  See [harfang.Scene_CreateObject].
func (pointer *Node) SetObject(o *Object) {
	oToC := o.h
	C.WrapSetObjectNode(pointer.h, oToC)
}

// RemoveObject Remove the [harfang.Object] component from the [harfang.Node].
func (pointer *Node) RemoveObject() {
	C.WrapRemoveObjectNode(pointer.h)
}

// GetMinMax ...
func (pointer *Node) GetMinMax(resources *PipelineResources) (bool, *MinMax) {
	resourcesToC := resources.h
	minmax := NewMinMax()
	minmaxToC := minmax.h
	retval := C.WrapGetMinMaxNode(pointer.h, resourcesToC, minmaxToC)
	return bool(retval), minmax
}

// HasLight Return `true` if the [harfang.Node] has a [harfang.Light] component.
func (pointer *Node) HasLight() bool {
	retval := C.WrapHasLightNode(pointer.h)
	return bool(retval)
}

// GetLight Return the [harfang.Light] component of the node.
func (pointer *Node) GetLight() *Light {
	retval := C.WrapGetLightNode(pointer.h)
	retvalGO := &Light{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Light) {
		C.WrapLightFree(cleanval.h)
	})
	return retvalGO
}

// SetLight Set the [harfang.Light] component of a node.  See [harfang.Scene_CreateLight], [harfang.Scene_CreatePointLight], [harfang.Scene_CreateSpotLight] or [harfang.Scene_CreateLinearLight].
func (pointer *Node) SetLight(l *Light) {
	lToC := l.h
	C.WrapSetLightNode(pointer.h, lToC)
}

// RemoveLight Remove the [harfang.Light] component from the [harfang.Node].
func (pointer *Node) RemoveLight() {
	C.WrapRemoveLightNode(pointer.h)
}

// HasRigidBody Return `true` if the [harfang.Node] has a [harfang.RigidBody] component.
func (pointer *Node) HasRigidBody() bool {
	retval := C.WrapHasRigidBodyNode(pointer.h)
	return bool(retval)
}

// GetRigidBody Return the [harfang.RigidBody] component of a [harfang.Node].
func (pointer *Node) GetRigidBody() *RigidBody {
	retval := C.WrapGetRigidBodyNode(pointer.h)
	retvalGO := &RigidBody{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RigidBody) {
		C.WrapRigidBodyFree(cleanval.h)
	})
	return retvalGO
}

// SetRigidBody Set the [harfang.RigidBody] component of a node.
func (pointer *Node) SetRigidBody(b *RigidBody) {
	bToC := b.h
	C.WrapSetRigidBodyNode(pointer.h, bToC)
}

// RemoveRigidBody Remove the [harfang.RigidBody] component from the [harfang.Node].
func (pointer *Node) RemoveRigidBody() {
	C.WrapRemoveRigidBodyNode(pointer.h)
}

// GetCollisionCount Return the amount of [harfang.Collision] components of a [harfang.Node].
func (pointer *Node) GetCollisionCount() int32 {
	retval := C.WrapGetCollisionCountNode(pointer.h)
	return int32(retval)
}

// GetCollision Return the [harfang.Collision] component attached to the [harfang.Node] at the desired `slot`.
func (pointer *Node) GetCollision(slot int32) *Collision {
	slotToC := C.size_t(slot)
	retval := C.WrapGetCollisionNode(pointer.h, slotToC)
	retvalGO := &Collision{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Collision) {
		C.WrapCollisionFree(cleanval.h)
	})
	return retvalGO
}

// SetCollision Set the [harfang.Collision] component of a [harfang.Node] at the desired `slot`. This is how you combine several collision shapes on a single node.
func (pointer *Node) SetCollision(slot int32, c *Collision) {
	slotToC := C.size_t(slot)
	cToC := c.h
	C.WrapSetCollisionNode(pointer.h, slotToC, cToC)
}

// RemoveCollision Remove the [harfang.Collision] component attached to the `slot` of the [harfang.Node].
func (pointer *Node) RemoveCollision(c *Collision) {
	cToC := c.h
	C.WrapRemoveCollisionNode(pointer.h, cToC)
}

// RemoveCollisionWithSlot Remove the [harfang.Collision] component attached to the `slot` of the [harfang.Node].
func (pointer *Node) RemoveCollisionWithSlot(slot int32) {
	slotToC := C.size_t(slot)
	C.WrapRemoveCollisionNodeWithSlot(pointer.h, slotToC)
}

// GetName Return the node name.
func (pointer *Node) GetName() string {
	retval := C.WrapGetNameNode(pointer.h)
	return C.GoString(retval)
}

// SetName Set the node name.
func (pointer *Node) SetName(name string) {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	C.WrapSetNameNode(pointer.h, nameToC)
}

// GetScriptCount ...
func (pointer *Node) GetScriptCount() int32 {
	retval := C.WrapGetScriptCountNode(pointer.h)
	return int32(retval)
}

// GetScript ...
func (pointer *Node) GetScript(idx int32) *Script {
	idxToC := C.size_t(idx)
	retval := C.WrapGetScriptNode(pointer.h, idxToC)
	retvalGO := &Script{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Script) {
		C.WrapScriptFree(cleanval.h)
	})
	return retvalGO
}

// SetScript ...
func (pointer *Node) SetScript(idx int32, s *Script) {
	idxToC := C.size_t(idx)
	sToC := s.h
	C.WrapSetScriptNode(pointer.h, idxToC, sToC)
}

// RemoveScript ...
func (pointer *Node) RemoveScript(s *Script) {
	sToC := s.h
	C.WrapRemoveScriptNode(pointer.h, sToC)
}

// RemoveScriptWithSlot ...
func (pointer *Node) RemoveScriptWithSlot(slot int32) {
	slotToC := C.size_t(slot)
	C.WrapRemoveScriptNodeWithSlot(pointer.h, slotToC)
}

// HasInstance Return `true` if the [harfang.Node] has an [harfang.Instance] component.
func (pointer *Node) HasInstance() bool {
	retval := C.WrapHasInstanceNode(pointer.h)
	return bool(retval)
}

// GetInstance Return the [harfang.Instance] component of a [harfang.Node].
func (pointer *Node) GetInstance() *Instance {
	retval := C.WrapGetInstanceNode(pointer.h)
	retvalGO := &Instance{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Instance) {
		C.WrapInstanceFree(cleanval.h)
	})
	return retvalGO
}

// SetInstance Set the [harfang.Instance] component of a [harfang.Node].
func (pointer *Node) SetInstance(instance *Instance) {
	instanceToC := instance.h
	C.WrapSetInstanceNode(pointer.h, instanceToC)
}

// SetupInstanceFromFile ...
func (pointer *Node) SetupInstanceFromFile(resources *PipelineResources, pipeline *PipelineInfo) bool {
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapSetupInstanceFromFileNode(pointer.h, resourcesToC, pipelineToC)
	return bool(retval)
}

// SetupInstanceFromFileWithFlags ...
func (pointer *Node) SetupInstanceFromFileWithFlags(resources *PipelineResources, pipeline *PipelineInfo, flags LoadSaveSceneFlags) bool {
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapSetupInstanceFromFileNodeWithFlags(pointer.h, resourcesToC, pipelineToC, flagsToC)
	return bool(retval)
}

// SetupInstanceFromAssets ...
func (pointer *Node) SetupInstanceFromAssets(resources *PipelineResources, pipeline *PipelineInfo) bool {
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapSetupInstanceFromAssetsNode(pointer.h, resourcesToC, pipelineToC)
	return bool(retval)
}

// SetupInstanceFromAssetsWithFlags ...
func (pointer *Node) SetupInstanceFromAssetsWithFlags(resources *PipelineResources, pipeline *PipelineInfo, flags LoadSaveSceneFlags) bool {
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapSetupInstanceFromAssetsNodeWithFlags(pointer.h, resourcesToC, pipelineToC, flagsToC)
	return bool(retval)
}

// DestroyInstance ...
func (pointer *Node) DestroyInstance() {
	C.WrapDestroyInstanceNode(pointer.h)
}

// IsInstantiatedBy ...
func (pointer *Node) IsInstantiatedBy() *Node {
	retval := C.WrapIsInstantiatedByNode(pointer.h)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// GetInstanceSceneView ...
func (pointer *Node) GetInstanceSceneView() *SceneView {
	retval := C.WrapGetInstanceSceneViewNode(pointer.h)
	retvalGO := &SceneView{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneView) {
		C.WrapSceneViewFree(cleanval.h)
	})
	return retvalGO
}

// GetInstanceSceneAnim ...
func (pointer *Node) GetInstanceSceneAnim(path string) *SceneAnimRef {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapGetInstanceSceneAnimNode(pointer.h, pathToC)
	retvalGO := &SceneAnimRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneAnimRef) {
		C.WrapSceneAnimRefFree(cleanval.h)
	})
	return retvalGO
}

// StartOnInstantiateAnim ...
func (pointer *Node) StartOnInstantiateAnim() {
	C.WrapStartOnInstantiateAnimNode(pointer.h)
}

// StopOnInstantiateAnim ...
func (pointer *Node) StopOnInstantiateAnim() {
	C.WrapStopOnInstantiateAnimNode(pointer.h)
}

// GetWorld ...
func (pointer *Node) GetWorld() *Mat4 {
	retval := C.WrapGetWorldNode(pointer.h)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// SetWorld ...
func (pointer *Node) SetWorld(world *Mat4) {
	worldToC := world.h
	C.WrapSetWorldNode(pointer.h, worldToC)
}

// ComputeWorld ...
func (pointer *Node) ComputeWorld() *Mat4 {
	retval := C.WrapComputeWorldNode(pointer.h)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// TransformTRS  Translation, rotation and scale packed as a single object.
type TransformTRS struct {
	h C.WrapTransformTRS
}

// GetPos ...
func (pointer *TransformTRS) GetPos() *Vec3 {
	v := C.WrapTransformTRSGetPos(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetPos ...
func (pointer *TransformTRS) SetPos(v *Vec3) {
	vToC := v.h
	C.WrapTransformTRSSetPos(pointer.h, vToC)
}

// GetRot ...
func (pointer *TransformTRS) GetRot() *Vec3 {
	v := C.WrapTransformTRSGetRot(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetRot ...
func (pointer *TransformTRS) SetRot(v *Vec3) {
	vToC := v.h
	C.WrapTransformTRSSetRot(pointer.h, vToC)
}

// GetScl ...
func (pointer *TransformTRS) GetScl() *Vec3 {
	v := C.WrapTransformTRSGetScl(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetScl ...
func (pointer *TransformTRS) SetScl(v *Vec3) {
	vToC := v.h
	C.WrapTransformTRSSetScl(pointer.h, vToC)
}

// NewTransformTRS Translation, rotation and scale packed as a single object.
func NewTransformTRS() *TransformTRS {
	retval := C.WrapConstructorTransformTRS()
	retvalGO := &TransformTRS{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *TransformTRS) {
		C.WrapTransformTRSFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *TransformTRS) Free() {
	C.WrapTransformTRSFree(pointer.h)
}

// IsNil ...
func (pointer *TransformTRS) IsNil() bool {
	return pointer.h == C.WrapTransformTRS(nil)
}

// Transform  Transformation component for a [harfang.Node], see [harfang.man.Scene].
type Transform struct {
	h C.WrapTransform
}

// Free ...
func (pointer *Transform) Free() {
	C.WrapTransformFree(pointer.h)
}

// IsNil ...
func (pointer *Transform) IsNil() bool {
	return pointer.h == C.WrapTransform(nil)
}

// Eq ...
func (pointer *Transform) Eq(t *Transform) bool {
	tToC := t.h
	retval := C.WrapEqTransform(pointer.h, tToC)
	return bool(retval)
}

// IsValid ...
func (pointer *Transform) IsValid() bool {
	retval := C.WrapIsValidTransform(pointer.h)
	return bool(retval)
}

// GetPos Return the transform position.    >  If you want the visual position of a Node with a rigid body, use [harfang.GetT] on [harfang.Transform_GetWorld]. See [harfang.man.Physics].
func (pointer *Transform) GetPos() *Vec3 {
	retval := C.WrapGetPosTransform(pointer.h)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// SetPos Set the transform position.    > When the Node has a RigidBody component, [harfang.Transform_SetPos] is overrided by physics. See [harfang.man.Physics].
func (pointer *Transform) SetPos(T *Vec3) {
	TToC := T.h
	C.WrapSetPosTransform(pointer.h, TToC)
}

// GetRot Get the transform rotation. If you want the visual rotation of a Node with a rigid body, use [harfang.GetRotation] on the matrix returned by [harfang.Transform_GetWorld]. See [harfang.man.Physics].
func (pointer *Transform) GetRot() *Vec3 {
	retval := C.WrapGetRotTransform(pointer.h)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// SetRot Set the transform rotation from a vector of _Euler_ angles.    > When the Node has a RigidBody component, [harfang.Transform_SetRot] is overrided by physics. See [harfang.man.Physics].
func (pointer *Transform) SetRot(R *Vec3) {
	RToC := R.h
	C.WrapSetRotTransform(pointer.h, RToC)
}

// GetScale Get the transform scale.
func (pointer *Transform) GetScale() *Vec3 {
	retval := C.WrapGetScaleTransform(pointer.h)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// SetScale Set the transform scale.
func (pointer *Transform) SetScale(S *Vec3) {
	SToC := S.h
	C.WrapSetScaleTransform(pointer.h, SToC)
}

// GetTRS Return the [harfang.TransformTRS].
func (pointer *Transform) GetTRS() *TransformTRS {
	retval := C.WrapGetTRSTransform(pointer.h)
	retvalGO := &TransformTRS{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *TransformTRS) {
		C.WrapTransformTRSFree(cleanval.h)
	})
	return retvalGO
}

// SetTRS Set the [harfang.TransformTRS].
func (pointer *Transform) SetTRS(TRS *TransformTRS) {
	TRSToC := TRS.h
	C.WrapSetTRSTransform(pointer.h, TRSToC)
}

// GetPosRot ...
func (pointer *Transform) GetPosRot() (*Vec3, *Vec3) {
	pos := NewVec3()
	posToC := pos.h
	rot := NewVec3()
	rotToC := rot.h
	C.WrapGetPosRotTransform(pointer.h, posToC, rotToC)
	return pos, rot
}

// SetPosRot ...
func (pointer *Transform) SetPosRot(pos *Vec3, rot *Vec3) {
	posToC := pos.h
	rotToC := rot.h
	C.WrapSetPosRotTransform(pointer.h, posToC, rotToC)
}

// GetParent Return the parent node for this [harfang.Transform].
func (pointer *Transform) GetParent() *Node {
	retval := C.WrapGetParentTransform(pointer.h)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// SetParent Set the parent node for this transform. This transform will then inherit the transformation of its parent node's [harfang.Transform] component.
func (pointer *Transform) SetParent(n *Node) {
	nToC := n.h
	C.WrapSetParentTransform(pointer.h, nToC)
}

// ClearParent ...
func (pointer *Transform) ClearParent() {
	C.WrapClearParentTransform(pointer.h)
}

// GetWorld Return the world matrix.
func (pointer *Transform) GetWorld() *Mat4 {
	retval := C.WrapGetWorldTransform(pointer.h)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// SetWorld Set the world matrix.    > When the Node has a RigidBody component, SetWorld is overrided by physics.
func (pointer *Transform) SetWorld(world *Mat4) {
	worldToC := world.h
	C.WrapSetWorldTransform(pointer.h, worldToC)
}

// SetLocal ...
func (pointer *Transform) SetLocal(local *Mat4) {
	localToC := local.h
	C.WrapSetLocalTransform(pointer.h, localToC)
}

// CameraZRange  ...
type CameraZRange struct {
	h C.WrapCameraZRange
}

// GetZnear ...
func (pointer *CameraZRange) GetZnear() float32 {
	v := C.WrapCameraZRangeGetZnear(pointer.h)
	return float32(v)
}

// SetZnear ...
func (pointer *CameraZRange) SetZnear(v float32) {
	vToC := C.float(v)
	C.WrapCameraZRangeSetZnear(pointer.h, vToC)
}

// GetZfar ...
func (pointer *CameraZRange) GetZfar() float32 {
	v := C.WrapCameraZRangeGetZfar(pointer.h)
	return float32(v)
}

// SetZfar ...
func (pointer *CameraZRange) SetZfar(v float32) {
	vToC := C.float(v)
	C.WrapCameraZRangeSetZfar(pointer.h, vToC)
}

// NewCameraZRange ...
func NewCameraZRange() *CameraZRange {
	retval := C.WrapConstructorCameraZRange()
	retvalGO := &CameraZRange{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *CameraZRange) {
		C.WrapCameraZRangeFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *CameraZRange) Free() {
	C.WrapCameraZRangeFree(pointer.h)
}

// IsNil ...
func (pointer *CameraZRange) IsNil() bool {
	return pointer.h == C.WrapCameraZRange(nil)
}

// Camera  Add this component to a [harfang.Node] to implement the camera aspect.  Create a camera component with [harfang.Scene_CreateCamera], use [harfang.CreateCamera] to create a complete camera node.
type Camera struct {
	h C.WrapCamera
}

// Free ...
func (pointer *Camera) Free() {
	C.WrapCameraFree(pointer.h)
}

// IsNil ...
func (pointer *Camera) IsNil() bool {
	return pointer.h == C.WrapCamera(nil)
}

// Eq ...
func (pointer *Camera) Eq(c *Camera) bool {
	cToC := c.h
	retval := C.WrapEqCamera(pointer.h, cToC)
	return bool(retval)
}

// IsValid ...
func (pointer *Camera) IsValid() bool {
	retval := C.WrapIsValidCamera(pointer.h)
	return bool(retval)
}

// GetZNear Return the camera near clipping plane.
func (pointer *Camera) GetZNear() float32 {
	retval := C.WrapGetZNearCamera(pointer.h)
	return float32(retval)
}

// SetZNear Set the camera near clipping plane.
func (pointer *Camera) SetZNear(v float32) {
	vToC := C.float(v)
	C.WrapSetZNearCamera(pointer.h, vToC)
}

// GetZFar Return the camera far clipping plane.
func (pointer *Camera) GetZFar() float32 {
	retval := C.WrapGetZFarCamera(pointer.h)
	return float32(retval)
}

// SetZFar Set the camera far clipping plane.
func (pointer *Camera) SetZFar(v float32) {
	vToC := C.float(v)
	C.WrapSetZFarCamera(pointer.h, vToC)
}

// GetZRange ...
func (pointer *Camera) GetZRange() *CameraZRange {
	retval := C.WrapGetZRangeCamera(pointer.h)
	retvalGO := &CameraZRange{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *CameraZRange) {
		C.WrapCameraZRangeFree(cleanval.h)
	})
	return retvalGO
}

// SetZRange ...
func (pointer *Camera) SetZRange(z *CameraZRange) {
	zToC := z.h
	C.WrapSetZRangeCamera(pointer.h, zToC)
}

// GetFov Return the camera field of view.
func (pointer *Camera) GetFov() float32 {
	retval := C.WrapGetFovCamera(pointer.h)
	return float32(retval)
}

// SetFov Set the camera field of view.
func (pointer *Camera) SetFov(v float32) {
	vToC := C.float(v)
	C.WrapSetFovCamera(pointer.h, vToC)
}

// GetIsOrthographic Return `true` if orthographic projection is used, `false` if perspective projection.
func (pointer *Camera) GetIsOrthographic() bool {
	retval := C.WrapGetIsOrthographicCamera(pointer.h)
	return bool(retval)
}

// SetIsOrthographic Configure the camera to use orthographic or perspective projection.
func (pointer *Camera) SetIsOrthographic(v bool) {
	vToC := C.bool(v)
	C.WrapSetIsOrthographicCamera(pointer.h, vToC)
}

// GetSize ...
func (pointer *Camera) GetSize() float32 {
	retval := C.WrapGetSizeCamera(pointer.h)
	return float32(retval)
}

// SetSize ...
func (pointer *Camera) SetSize(v float32) {
	vToC := C.float(v)
	C.WrapSetSizeCamera(pointer.h, vToC)
}

// Object  This components draws a [harfang.Model]. It stores the material table used to draw the model.
type Object struct {
	h C.WrapObject
}

// Free ...
func (pointer *Object) Free() {
	C.WrapObjectFree(pointer.h)
}

// IsNil ...
func (pointer *Object) IsNil() bool {
	return pointer.h == C.WrapObject(nil)
}

// Eq ...
func (pointer *Object) Eq(o *Object) bool {
	oToC := o.h
	retval := C.WrapEqObject(pointer.h, oToC)
	return bool(retval)
}

// IsValid ...
func (pointer *Object) IsValid() bool {
	retval := C.WrapIsValidObject(pointer.h)
	return bool(retval)
}

// GetModelRef Return the [harfang.ModelRef] to display.
func (pointer *Object) GetModelRef() *ModelRef {
	retval := C.WrapGetModelRefObject(pointer.h)
	retvalGO := &ModelRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ModelRef) {
		C.WrapModelRefFree(cleanval.h)
	})
	return retvalGO
}

// SetModelRef Set the [harfang.ModelRef] to display.
func (pointer *Object) SetModelRef(r *ModelRef) {
	rToC := r.h
	C.WrapSetModelRefObject(pointer.h, rToC)
}

// ClearModelRef ...
func (pointer *Object) ClearModelRef() {
	C.WrapClearModelRefObject(pointer.h)
}

// GetMaterial Return the object [harfang.Material] at index.
func (pointer *Object) GetMaterial(slotidx int32) *Material {
	slotidxToC := C.size_t(slotidx)
	retval := C.WrapGetMaterialObject(pointer.h, slotidxToC)
	var retvalGO *Material
	if retval != nil {
		retvalGO = &Material{h: retval}
	}
	return retvalGO
}

// GetMaterialWithName Return the object [harfang.Material] at index.
func (pointer *Object) GetMaterialWithName(name string) *Material {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapGetMaterialObjectWithName(pointer.h, nameToC)
	var retvalGO *Material
	if retval != nil {
		retvalGO = &Material{h: retval}
		runtime.SetFinalizer(retvalGO, func(cleanval *Material) {
			C.WrapMaterialFree(cleanval.h)
		})
	}
	return retvalGO
}

// SetMaterial Set the object [harfang.Material] at index.
func (pointer *Object) SetMaterial(slotidx int32, mat *Material) {
	slotidxToC := C.size_t(slotidx)
	matToC := mat.h
	C.WrapSetMaterialObject(pointer.h, slotidxToC, matToC)
}

// GetMaterialCount Return the number of [harfang.Material] in the object material table.
func (pointer *Object) GetMaterialCount() int32 {
	retval := C.WrapGetMaterialCountObject(pointer.h)
	return int32(retval)
}

// SetMaterialCount Set the number of [harfang.Material] in the object material table.
func (pointer *Object) SetMaterialCount(count int32) {
	countToC := C.size_t(count)
	C.WrapSetMaterialCountObject(pointer.h, countToC)
}

// GetMaterialName ...
func (pointer *Object) GetMaterialName(slotidx int32) string {
	slotidxToC := C.size_t(slotidx)
	retval := C.WrapGetMaterialNameObject(pointer.h, slotidxToC)
	return C.GoString(retval)
}

// SetMaterialName ...
func (pointer *Object) SetMaterialName(slotidx int32, name string) {
	slotidxToC := C.size_t(slotidx)
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	C.WrapSetMaterialNameObject(pointer.h, slotidxToC, nameToC)
}

// GetMinMax ...
func (pointer *Object) GetMinMax(resources *PipelineResources) (bool, *MinMax) {
	resourcesToC := resources.h
	minmax := NewMinMax()
	minmaxToC := minmax.h
	retval := C.WrapGetMinMaxObject(pointer.h, resourcesToC, minmaxToC)
	return bool(retval), minmax
}

// GetBoneCount ...
func (pointer *Object) GetBoneCount() int32 {
	retval := C.WrapGetBoneCountObject(pointer.h)
	return int32(retval)
}

// SetBoneCount ...
func (pointer *Object) SetBoneCount(count int32) {
	countToC := C.size_t(count)
	C.WrapSetBoneCountObject(pointer.h, countToC)
}

// SetBone ...
func (pointer *Object) SetBone(idx int32, node *Node) bool {
	idxToC := C.size_t(idx)
	nodeToC := node.h
	retval := C.WrapSetBoneObject(pointer.h, idxToC, nodeToC)
	return bool(retval)
}

// GetBone ...
func (pointer *Object) GetBone(idx int32) *Node {
	idxToC := C.size_t(idx)
	retval := C.WrapGetBoneObject(pointer.h, idxToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// Light  Add this component to a node to turn it into a light source, see [harfang.man.ForwardPipeline].
type Light struct {
	h C.WrapLight
}

// Free ...
func (pointer *Light) Free() {
	C.WrapLightFree(pointer.h)
}

// IsNil ...
func (pointer *Light) IsNil() bool {
	return pointer.h == C.WrapLight(nil)
}

// Eq ...
func (pointer *Light) Eq(l *Light) bool {
	lToC := l.h
	retval := C.WrapEqLight(pointer.h, lToC)
	return bool(retval)
}

// IsValid ...
func (pointer *Light) IsValid() bool {
	retval := C.WrapIsValidLight(pointer.h)
	return bool(retval)
}

// GetType Return the [harfang.LightType].
func (pointer *Light) GetType() LightType {
	retval := C.WrapGetTypeLight(pointer.h)
	return LightType(retval)
}

// SetType Set the [harfang.LightType].
func (pointer *Light) SetType(v LightType) {
	vToC := C.int32_t(v)
	C.WrapSetTypeLight(pointer.h, vToC)
}

// GetShadowType ...
func (pointer *Light) GetShadowType() LightShadowType {
	retval := C.WrapGetShadowTypeLight(pointer.h)
	return LightShadowType(retval)
}

// SetShadowType ...
func (pointer *Light) SetShadowType(v LightShadowType) {
	vToC := C.int32_t(v)
	C.WrapSetShadowTypeLight(pointer.h, vToC)
}

// GetDiffuseColor ...
func (pointer *Light) GetDiffuseColor() *Color {
	retval := C.WrapGetDiffuseColorLight(pointer.h)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// SetDiffuseColor ...
func (pointer *Light) SetDiffuseColor(v *Color) {
	vToC := v.h
	C.WrapSetDiffuseColorLight(pointer.h, vToC)
}

// GetDiffuseIntensity ...
func (pointer *Light) GetDiffuseIntensity() float32 {
	retval := C.WrapGetDiffuseIntensityLight(pointer.h)
	return float32(retval)
}

// SetDiffuseIntensity ...
func (pointer *Light) SetDiffuseIntensity(v float32) {
	vToC := C.float(v)
	C.WrapSetDiffuseIntensityLight(pointer.h, vToC)
}

// GetSpecularColor ...
func (pointer *Light) GetSpecularColor() *Color {
	retval := C.WrapGetSpecularColorLight(pointer.h)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// SetSpecularColor ...
func (pointer *Light) SetSpecularColor(v *Color) {
	vToC := v.h
	C.WrapSetSpecularColorLight(pointer.h, vToC)
}

// GetSpecularIntensity ...
func (pointer *Light) GetSpecularIntensity() float32 {
	retval := C.WrapGetSpecularIntensityLight(pointer.h)
	return float32(retval)
}

// SetSpecularIntensity ...
func (pointer *Light) SetSpecularIntensity(v float32) {
	vToC := C.float(v)
	C.WrapSetSpecularIntensityLight(pointer.h, vToC)
}

// GetRadius Get the light range in meters.
func (pointer *Light) GetRadius() float32 {
	retval := C.WrapGetRadiusLight(pointer.h)
	return float32(retval)
}

// SetRadius Set the light range in meters. No light will be contributed to elements further away than this distance from the node world position.
func (pointer *Light) SetRadius(v float32) {
	vToC := C.float(v)
	C.WrapSetRadiusLight(pointer.h, vToC)
}

// GetInnerAngle ...
func (pointer *Light) GetInnerAngle() float32 {
	retval := C.WrapGetInnerAngleLight(pointer.h)
	return float32(retval)
}

// SetInnerAngle ...
func (pointer *Light) SetInnerAngle(v float32) {
	vToC := C.float(v)
	C.WrapSetInnerAngleLight(pointer.h, vToC)
}

// GetOuterAngle ...
func (pointer *Light) GetOuterAngle() float32 {
	retval := C.WrapGetOuterAngleLight(pointer.h)
	return float32(retval)
}

// SetOuterAngle ...
func (pointer *Light) SetOuterAngle(v float32) {
	vToC := C.float(v)
	C.WrapSetOuterAngleLight(pointer.h, vToC)
}

// GetPSSMSplit ...
func (pointer *Light) GetPSSMSplit() *Vec4 {
	retval := C.WrapGetPSSMSplitLight(pointer.h)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// SetPSSMSplit ...
func (pointer *Light) SetPSSMSplit(v *Vec4) {
	vToC := v.h
	C.WrapSetPSSMSplitLight(pointer.h, vToC)
}

// GetPriority ...
func (pointer *Light) GetPriority() float32 {
	retval := C.WrapGetPriorityLight(pointer.h)
	return float32(retval)
}

// SetPriority ...
func (pointer *Light) SetPriority(v float32) {
	vToC := C.float(v)
	C.WrapSetPriorityLight(pointer.h, vToC)
}

// Contact  Object containing the world space position, normal and depth of a contact as reported by the collision system.
type Contact struct {
	h C.WrapContact
}

// GetP ...
func (pointer *Contact) GetP() *Vec3 {
	v := C.WrapContactGetP(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetP ...
func (pointer *Contact) SetP(v *Vec3) {
	vToC := v.h
	C.WrapContactSetP(pointer.h, vToC)
}

// GetN ...
func (pointer *Contact) GetN() *Vec3 {
	v := C.WrapContactGetN(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetN ...
func (pointer *Contact) SetN(v *Vec3) {
	vToC := v.h
	C.WrapContactSetN(pointer.h, vToC)
}

// GetD ...
func (pointer *Contact) GetD() float32 {
	v := C.WrapContactGetD(pointer.h)
	return float32(v)
}

// SetD ...
func (pointer *Contact) SetD(v float32) {
	vToC := C.float(v)
	C.WrapContactSetD(pointer.h, vToC)
}

// Free ...
func (pointer *Contact) Free() {
	C.WrapContactFree(pointer.h)
}

// IsNil ...
func (pointer *Contact) IsNil() bool {
	return pointer.h == C.WrapContact(nil)
}

// GoSliceOfContact ...
type GoSliceOfContact []*Contact

// ContactList  ...
type ContactList struct {
	h C.WrapContactList
}

// Get ...
func (pointer *ContactList) Get(id int) *Contact {
	v := C.WrapContactListGetOperator(pointer.h, C.int(id))
	vGO := &Contact{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *Contact) {
		C.WrapContactFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *ContactList) Set(id int, v *Contact) {
	vToC := v.h
	C.WrapContactListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *ContactList) Len() int32 {
	return int32(C.WrapContactListLenOperator(pointer.h))
}

// NewContactList ...
func NewContactList() *ContactList {
	retval := C.WrapConstructorContactList()
	retvalGO := &ContactList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ContactList) {
		C.WrapContactListFree(cleanval.h)
	})
	return retvalGO
}

// NewContactListWithSequence ...
func NewContactListWithSequence(sequence GoSliceOfContact) *ContactList {
	var sequencePointer []C.WrapContact
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapContact)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorContactListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &ContactList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ContactList) {
		C.WrapContactListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *ContactList) Free() {
	C.WrapContactListFree(pointer.h)
}

// IsNil ...
func (pointer *ContactList) IsNil() bool {
	return pointer.h == C.WrapContactList(nil)
}

// Clear ...
func (pointer *ContactList) Clear() {
	C.WrapClearContactList(pointer.h)
}

// Reserve ...
func (pointer *ContactList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveContactList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *ContactList) PushBack(v *Contact) {
	vToC := v.h
	C.WrapPushBackContactList(pointer.h, vToC)
}

// Size ...
func (pointer *ContactList) Size() int32 {
	retval := C.WrapSizeContactList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *ContactList) At(idx int32) *Contact {
	idxToC := C.size_t(idx)
	retval := C.WrapAtContactList(pointer.h, idxToC)
	retvalGO := &Contact{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Contact) {
		C.WrapContactFree(cleanval.h)
	})
	return retvalGO
}

// RigidBody  Rigid body component, see [harfang.man.Physics].
type RigidBody struct {
	h C.WrapRigidBody
}

// Free ...
func (pointer *RigidBody) Free() {
	C.WrapRigidBodyFree(pointer.h)
}

// IsNil ...
func (pointer *RigidBody) IsNil() bool {
	return pointer.h == C.WrapRigidBody(nil)
}

// Eq ...
func (pointer *RigidBody) Eq(b *RigidBody) bool {
	bToC := b.h
	retval := C.WrapEqRigidBody(pointer.h, bToC)
	return bool(retval)
}

// IsValid ...
func (pointer *RigidBody) IsValid() bool {
	retval := C.WrapIsValidRigidBody(pointer.h)
	return bool(retval)
}

// GetType Return the rigid body type. See [harfang.RigidBody_SetType].
func (pointer *RigidBody) GetType() RigidBodyType {
	retval := C.WrapGetTypeRigidBody(pointer.h)
	return RigidBodyType(retval)
}

// SetType Set the rigid body type.
func (pointer *RigidBody) SetType(typeGo RigidBodyType) {
	typeGoToC := C.uchar(typeGo)
	C.WrapSetTypeRigidBody(pointer.h, typeGoToC)
}

// GetLinearDamping ...
func (pointer *RigidBody) GetLinearDamping() float32 {
	retval := C.WrapGetLinearDampingRigidBody(pointer.h)
	return float32(retval)
}

// SetLinearDamping ...
func (pointer *RigidBody) SetLinearDamping(damping float32) {
	dampingToC := C.float(damping)
	C.WrapSetLinearDampingRigidBody(pointer.h, dampingToC)
}

// GetAngularDamping Return the rigid body angular damping. A value of 0.0 means no damping, 1.0 means the maximal dissipation of the energy.
func (pointer *RigidBody) GetAngularDamping() float32 {
	retval := C.WrapGetAngularDampingRigidBody(pointer.h)
	return float32(retval)
}

// SetAngularDamping Set the rigid body angular damping. A value of 0.0 means no damping, 1.0 means the maximal dissipation of the energy.
func (pointer *RigidBody) SetAngularDamping(damping float32) {
	dampingToC := C.float(damping)
	C.WrapSetAngularDampingRigidBody(pointer.h, dampingToC)
}

// GetRestitution ...
func (pointer *RigidBody) GetRestitution() float32 {
	retval := C.WrapGetRestitutionRigidBody(pointer.h)
	return float32(retval)
}

// SetRestitution ...
func (pointer *RigidBody) SetRestitution(restitution float32) {
	restitutionToC := C.float(restitution)
	C.WrapSetRestitutionRigidBody(pointer.h, restitutionToC)
}

// GetFriction ...
func (pointer *RigidBody) GetFriction() float32 {
	retval := C.WrapGetFrictionRigidBody(pointer.h)
	return float32(retval)
}

// SetFriction ...
func (pointer *RigidBody) SetFriction(friction float32) {
	frictionToC := C.float(friction)
	C.WrapSetFrictionRigidBody(pointer.h, frictionToC)
}

// GetRollingFriction ...
func (pointer *RigidBody) GetRollingFriction() float32 {
	retval := C.WrapGetRollingFrictionRigidBody(pointer.h)
	return float32(retval)
}

// SetRollingFriction ...
func (pointer *RigidBody) SetRollingFriction(rollingfriction float32) {
	rollingfrictionToC := C.float(rollingfriction)
	C.WrapSetRollingFrictionRigidBody(pointer.h, rollingfrictionToC)
}

// Collision  Collision component, see [harfang.man.Physics].
type Collision struct {
	h C.WrapCollision
}

// Free ...
func (pointer *Collision) Free() {
	C.WrapCollisionFree(pointer.h)
}

// IsNil ...
func (pointer *Collision) IsNil() bool {
	return pointer.h == C.WrapCollision(nil)
}

// Eq ...
func (pointer *Collision) Eq(c *Collision) bool {
	cToC := c.h
	retval := C.WrapEqCollision(pointer.h, cToC)
	return bool(retval)
}

// IsValid ...
func (pointer *Collision) IsValid() bool {
	retval := C.WrapIsValidCollision(pointer.h)
	return bool(retval)
}

// GetType Return the [harfang.CollisionType] of a [harfang.Collision] component.
func (pointer *Collision) GetType() CollisionType {
	retval := C.WrapGetTypeCollision(pointer.h)
	return CollisionType(retval)
}

// SetType Set the [harfang.CollisionType] of a [harfang.Collision] component.
func (pointer *Collision) SetType(typeGo CollisionType) {
	typeGoToC := C.uchar(typeGo)
	C.WrapSetTypeCollision(pointer.h, typeGoToC)
}

// GetLocalTransform ...
func (pointer *Collision) GetLocalTransform() *Mat4 {
	retval := C.WrapGetLocalTransformCollision(pointer.h)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// SetLocalTransform ...
func (pointer *Collision) SetLocalTransform(m *Mat4) {
	mToC := m.h
	C.WrapSetLocalTransformCollision(pointer.h, mToC)
}

// GetMass Return the collision shape mass in Kg.
func (pointer *Collision) GetMass() float32 {
	retval := C.WrapGetMassCollision(pointer.h)
	return float32(retval)
}

// SetMass Set the collision shape mass in Kg.
func (pointer *Collision) SetMass(mass float32) {
	massToC := C.float(mass)
	C.WrapSetMassCollision(pointer.h, massToC)
}

// GetRadius ...
func (pointer *Collision) GetRadius() float32 {
	retval := C.WrapGetRadiusCollision(pointer.h)
	return float32(retval)
}

// SetRadius ...
func (pointer *Collision) SetRadius(radius float32) {
	radiusToC := C.float(radius)
	C.WrapSetRadiusCollision(pointer.h, radiusToC)
}

// GetHeight ...
func (pointer *Collision) GetHeight() float32 {
	retval := C.WrapGetHeightCollision(pointer.h)
	return float32(retval)
}

// SetHeight ...
func (pointer *Collision) SetHeight(height float32) {
	heightToC := C.float(height)
	C.WrapSetHeightCollision(pointer.h, heightToC)
}

// SetSize ...
func (pointer *Collision) SetSize(size *Vec3) {
	sizeToC := size.h
	C.WrapSetSizeCollision(pointer.h, sizeToC)
}

// GetCollisionResource ...
func (pointer *Collision) GetCollisionResource() string {
	retval := C.WrapGetCollisionResourceCollision(pointer.h)
	return C.GoString(retval)
}

// SetCollisionResource ...
func (pointer *Collision) SetCollisionResource(path string) {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	C.WrapSetCollisionResourceCollision(pointer.h, pathToC)
}

// Instance  Component to instantiate a scene as a child of a node upon setup.
type Instance struct {
	h C.WrapInstance
}

// Free ...
func (pointer *Instance) Free() {
	C.WrapInstanceFree(pointer.h)
}

// IsNil ...
func (pointer *Instance) IsNil() bool {
	return pointer.h == C.WrapInstance(nil)
}

// Eq ...
func (pointer *Instance) Eq(i *Instance) bool {
	iToC := i.h
	retval := C.WrapEqInstance(pointer.h, iToC)
	return bool(retval)
}

// IsValid ...
func (pointer *Instance) IsValid() bool {
	retval := C.WrapIsValidInstance(pointer.h)
	return bool(retval)
}

// GetPath ...
func (pointer *Instance) GetPath() string {
	retval := C.WrapGetPathInstance(pointer.h)
	return C.GoString(retval)
}

// SetPath ...
func (pointer *Instance) SetPath(path string) {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	C.WrapSetPathInstance(pointer.h, pathToC)
}

// SetOnInstantiateAnim ...
func (pointer *Instance) SetOnInstantiateAnim(anim string) {
	animToC, idFinanimToC := wrapString(anim)
	defer idFinanimToC()
	C.WrapSetOnInstantiateAnimInstance(pointer.h, animToC)
}

// SetOnInstantiateAnimLoopMode ...
func (pointer *Instance) SetOnInstantiateAnimLoopMode(loopmode AnimLoopMode) {
	loopmodeToC := C.int32_t(loopmode)
	C.WrapSetOnInstantiateAnimLoopModeInstance(pointer.h, loopmodeToC)
}

// ClearOnInstantiateAnim ...
func (pointer *Instance) ClearOnInstantiateAnim() {
	C.WrapClearOnInstantiateAnimInstance(pointer.h)
}

// GetOnInstantiateAnim ...
func (pointer *Instance) GetOnInstantiateAnim() string {
	retval := C.WrapGetOnInstantiateAnimInstance(pointer.h)
	return C.GoString(retval)
}

// GetOnInstantiateAnimLoopMode ...
func (pointer *Instance) GetOnInstantiateAnimLoopMode() AnimLoopMode {
	retval := C.WrapGetOnInstantiateAnimLoopModeInstance(pointer.h)
	return AnimLoopMode(retval)
}

// GetOnInstantiatePlayAnimRef ...
func (pointer *Instance) GetOnInstantiatePlayAnimRef() *ScenePlayAnimRef {
	retval := C.WrapGetOnInstantiatePlayAnimRefInstance(pointer.h)
	retvalGO := &ScenePlayAnimRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScenePlayAnimRef) {
		C.WrapScenePlayAnimRefFree(cleanval.h)
	})
	return retvalGO
}

// Script  ...
type Script struct {
	h C.WrapScript
}

// Free ...
func (pointer *Script) Free() {
	C.WrapScriptFree(pointer.h)
}

// IsNil ...
func (pointer *Script) IsNil() bool {
	return pointer.h == C.WrapScript(nil)
}

// Eq ...
func (pointer *Script) Eq(s *Script) bool {
	sToC := s.h
	retval := C.WrapEqScript(pointer.h, sToC)
	return bool(retval)
}

// IsValid ...
func (pointer *Script) IsValid() bool {
	retval := C.WrapIsValidScript(pointer.h)
	return bool(retval)
}

// GetPath ...
func (pointer *Script) GetPath() string {
	retval := C.WrapGetPathScript(pointer.h)
	return C.GoString(retval)
}

// SetPath ...
func (pointer *Script) SetPath(path string) {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	C.WrapSetPathScript(pointer.h, pathToC)
}

// GoSliceOfScript ...
type GoSliceOfScript []*Script

// ScriptList  ...
type ScriptList struct {
	h C.WrapScriptList
}

// Get ...
func (pointer *ScriptList) Get(id int) *Script {
	v := C.WrapScriptListGetOperator(pointer.h, C.int(id))
	vGO := &Script{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *Script) {
		C.WrapScriptFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *ScriptList) Set(id int, v *Script) {
	vToC := v.h
	C.WrapScriptListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *ScriptList) Len() int32 {
	return int32(C.WrapScriptListLenOperator(pointer.h))
}

// NewScriptList ...
func NewScriptList() *ScriptList {
	retval := C.WrapConstructorScriptList()
	retvalGO := &ScriptList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScriptList) {
		C.WrapScriptListFree(cleanval.h)
	})
	return retvalGO
}

// NewScriptListWithSequence ...
func NewScriptListWithSequence(sequence GoSliceOfScript) *ScriptList {
	var sequencePointer []C.WrapScript
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapScript)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorScriptListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &ScriptList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScriptList) {
		C.WrapScriptListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *ScriptList) Free() {
	C.WrapScriptListFree(pointer.h)
}

// IsNil ...
func (pointer *ScriptList) IsNil() bool {
	return pointer.h == C.WrapScriptList(nil)
}

// Clear ...
func (pointer *ScriptList) Clear() {
	C.WrapClearScriptList(pointer.h)
}

// Reserve ...
func (pointer *ScriptList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveScriptList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *ScriptList) PushBack(v *Script) {
	vToC := v.h
	C.WrapPushBackScriptList(pointer.h, vToC)
}

// Size ...
func (pointer *ScriptList) Size() int32 {
	retval := C.WrapSizeScriptList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *ScriptList) At(idx int32) *Script {
	idxToC := C.size_t(idx)
	retval := C.WrapAtScriptList(pointer.h, idxToC)
	retvalGO := &Script{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Script) {
		C.WrapScriptFree(cleanval.h)
	})
	return retvalGO
}

// GoSliceOfNode ...
type GoSliceOfNode []*Node

// NodeList  ...
type NodeList struct {
	h C.WrapNodeList
}

// Get ...
func (pointer *NodeList) Get(id int) *Node {
	v := C.WrapNodeListGetOperator(pointer.h, C.int(id))
	vGO := &Node{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *NodeList) Set(id int, v *Node) {
	vToC := v.h
	C.WrapNodeListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *NodeList) Len() int32 {
	return int32(C.WrapNodeListLenOperator(pointer.h))
}

// NewNodeList ...
func NewNodeList() *NodeList {
	retval := C.WrapConstructorNodeList()
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// NewNodeListWithSequence ...
func NewNodeListWithSequence(sequence GoSliceOfNode) *NodeList {
	var sequencePointer []C.WrapNode
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapNode)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorNodeListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *NodeList) Free() {
	C.WrapNodeListFree(pointer.h)
}

// IsNil ...
func (pointer *NodeList) IsNil() bool {
	return pointer.h == C.WrapNodeList(nil)
}

// Clear ...
func (pointer *NodeList) Clear() {
	C.WrapClearNodeList(pointer.h)
}

// Reserve ...
func (pointer *NodeList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveNodeList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *NodeList) PushBack(v *Node) {
	vToC := v.h
	C.WrapPushBackNodeList(pointer.h, vToC)
}

// Size ...
func (pointer *NodeList) Size() int32 {
	retval := C.WrapSizeNodeList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *NodeList) At(idx int32) *Node {
	idxToC := C.size_t(idx)
	retval := C.WrapAtNodeList(pointer.h, idxToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// RaycastOut  Contains the result of a physics raycast.  * `P`: Position of the raycast hit * `N`: Normal of the raycast hit * `Node`: Node hit by the raycast * `t`: Parametric value of the intersection, ratio of the distance to the hit by the length of the raycast
type RaycastOut struct {
	h C.WrapRaycastOut
}

// GetP ...
func (pointer *RaycastOut) GetP() *Vec3 {
	v := C.WrapRaycastOutGetP(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetP ...
func (pointer *RaycastOut) SetP(v *Vec3) {
	vToC := v.h
	C.WrapRaycastOutSetP(pointer.h, vToC)
}

// GetN ...
func (pointer *RaycastOut) GetN() *Vec3 {
	v := C.WrapRaycastOutGetN(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetN ...
func (pointer *RaycastOut) SetN(v *Vec3) {
	vToC := v.h
	C.WrapRaycastOutSetN(pointer.h, vToC)
}

// GetNode ...
func (pointer *RaycastOut) GetNode() *Node {
	v := C.WrapRaycastOutGetNode(pointer.h)
	vGO := &Node{h: v}
	return vGO
}

// SetNode ...
func (pointer *RaycastOut) SetNode(v *Node) {
	vToC := v.h
	C.WrapRaycastOutSetNode(pointer.h, vToC)
}

// GetT ...
func (pointer *RaycastOut) GetT() float32 {
	v := C.WrapRaycastOutGetT(pointer.h)
	return float32(v)
}

// SetT ...
func (pointer *RaycastOut) SetT(v float32) {
	vToC := C.float(v)
	C.WrapRaycastOutSetT(pointer.h, vToC)
}

// Free ...
func (pointer *RaycastOut) Free() {
	C.WrapRaycastOutFree(pointer.h)
}

// IsNil ...
func (pointer *RaycastOut) IsNil() bool {
	return pointer.h == C.WrapRaycastOut(nil)
}

// GoSliceOfRaycastOut ...
type GoSliceOfRaycastOut []*RaycastOut

// RaycastOutList  ...
type RaycastOutList struct {
	h C.WrapRaycastOutList
}

// Get ...
func (pointer *RaycastOutList) Get(id int) *RaycastOut {
	v := C.WrapRaycastOutListGetOperator(pointer.h, C.int(id))
	vGO := &RaycastOut{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *RaycastOut) {
		C.WrapRaycastOutFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *RaycastOutList) Set(id int, v *RaycastOut) {
	vToC := v.h
	C.WrapRaycastOutListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *RaycastOutList) Len() int32 {
	return int32(C.WrapRaycastOutListLenOperator(pointer.h))
}

// NewRaycastOutList ...
func NewRaycastOutList() *RaycastOutList {
	retval := C.WrapConstructorRaycastOutList()
	retvalGO := &RaycastOutList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RaycastOutList) {
		C.WrapRaycastOutListFree(cleanval.h)
	})
	return retvalGO
}

// NewRaycastOutListWithSequence ...
func NewRaycastOutListWithSequence(sequence GoSliceOfRaycastOut) *RaycastOutList {
	var sequencePointer []C.WrapRaycastOut
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapRaycastOut)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorRaycastOutListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &RaycastOutList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RaycastOutList) {
		C.WrapRaycastOutListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *RaycastOutList) Free() {
	C.WrapRaycastOutListFree(pointer.h)
}

// IsNil ...
func (pointer *RaycastOutList) IsNil() bool {
	return pointer.h == C.WrapRaycastOutList(nil)
}

// Clear ...
func (pointer *RaycastOutList) Clear() {
	C.WrapClearRaycastOutList(pointer.h)
}

// Reserve ...
func (pointer *RaycastOutList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveRaycastOutList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *RaycastOutList) PushBack(v *RaycastOut) {
	vToC := v.h
	C.WrapPushBackRaycastOutList(pointer.h, vToC)
}

// Size ...
func (pointer *RaycastOutList) Size() int32 {
	retval := C.WrapSizeRaycastOutList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *RaycastOutList) At(idx int32) *RaycastOut {
	idxToC := C.size_t(idx)
	retval := C.WrapAtRaycastOutList(pointer.h, idxToC)
	retvalGO := &RaycastOut{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RaycastOut) {
		C.WrapRaycastOutFree(cleanval.h)
	})
	return retvalGO
}

// TimeCallbackConnection  A [harfang.TimeCallback] connection to a [harfang.Signal_returning_void_taking_time_ns].
type TimeCallbackConnection struct {
	h C.WrapTimeCallbackConnection
}

// Free ...
func (pointer *TimeCallbackConnection) Free() {
	C.WrapTimeCallbackConnectionFree(pointer.h)
}

// IsNil ...
func (pointer *TimeCallbackConnection) IsNil() bool {
	return pointer.h == C.WrapTimeCallbackConnection(nil)
}

// SignalReturningVoidTakingTimeNs  ...
type SignalReturningVoidTakingTimeNs struct {
	h C.WrapSignalReturningVoidTakingTimeNs
}

// Free ...
func (pointer *SignalReturningVoidTakingTimeNs) Free() {
	C.WrapSignalReturningVoidTakingTimeNsFree(pointer.h)
}

// IsNil ...
func (pointer *SignalReturningVoidTakingTimeNs) IsNil() bool {
	return pointer.h == C.WrapSignalReturningVoidTakingTimeNs(nil)
}

// Connect ...
func (pointer *SignalReturningVoidTakingTimeNs) Connect(listener unsafe.Pointer) *TimeCallbackConnection {
	listenerToC := (C.WrapFunctionReturningVoidTakingTimeNs)(listener)
	retval := C.WrapConnectSignalReturningVoidTakingTimeNs(pointer.h, listenerToC)
	retvalGO := &TimeCallbackConnection{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *TimeCallbackConnection) {
		C.WrapTimeCallbackConnectionFree(cleanval.h)
	})
	return retvalGO
}

// Disconnect ...
func (pointer *SignalReturningVoidTakingTimeNs) Disconnect(connection *TimeCallbackConnection) {
	connectionToC := connection.h
	C.WrapDisconnectSignalReturningVoidTakingTimeNs(pointer.h, connectionToC)
}

// DisconnectAll ...
func (pointer *SignalReturningVoidTakingTimeNs) DisconnectAll() {
	C.WrapDisconnectAllSignalReturningVoidTakingTimeNs(pointer.h)
}

// Emit ...
func (pointer *SignalReturningVoidTakingTimeNs) Emit(arg0 int64) {
	arg0ToC := C.int64_t(arg0)
	C.WrapEmitSignalReturningVoidTakingTimeNs(pointer.h, arg0ToC)
}

// GetListenerCount ...
func (pointer *SignalReturningVoidTakingTimeNs) GetListenerCount() int32 {
	retval := C.WrapGetListenerCountSignalReturningVoidTakingTimeNs(pointer.h)
	return int32(retval)
}

// Canvas  Holds the canvas properties of a scene, see the `canvas` member of class [harfang.Scene].
type Canvas struct {
	h C.WrapCanvas
}

// GetClearZ ...
func (pointer *Canvas) GetClearZ() bool {
	v := C.WrapCanvasGetClearZ(pointer.h)
	return bool(v)
}

// SetClearZ ...
func (pointer *Canvas) SetClearZ(v bool) {
	vToC := C.bool(v)
	C.WrapCanvasSetClearZ(pointer.h, vToC)
}

// GetClearColor ...
func (pointer *Canvas) GetClearColor() bool {
	v := C.WrapCanvasGetClearColor(pointer.h)
	return bool(v)
}

// SetClearColor ...
func (pointer *Canvas) SetClearColor(v bool) {
	vToC := C.bool(v)
	C.WrapCanvasSetClearColor(pointer.h, vToC)
}

// GetColor ...
func (pointer *Canvas) GetColor() *Color {
	v := C.WrapCanvasGetColor(pointer.h)
	vGO := &Color{h: v}
	return vGO
}

// SetColor ...
func (pointer *Canvas) SetColor(v *Color) {
	vToC := v.h
	C.WrapCanvasSetColor(pointer.h, vToC)
}

// Free ...
func (pointer *Canvas) Free() {
	C.WrapCanvasFree(pointer.h)
}

// IsNil ...
func (pointer *Canvas) IsNil() bool {
	return pointer.h == C.WrapCanvas(nil)
}

// Environment  Environment properties of a scene, see `environment` member of the [harfang.Scene] class.
type Environment struct {
	h C.WrapEnvironment
}

// GetAmbient ...
func (pointer *Environment) GetAmbient() *Color {
	v := C.WrapEnvironmentGetAmbient(pointer.h)
	vGO := &Color{h: v}
	return vGO
}

// SetAmbient ...
func (pointer *Environment) SetAmbient(v *Color) {
	vToC := v.h
	C.WrapEnvironmentSetAmbient(pointer.h, vToC)
}

// GetFogColor ...
func (pointer *Environment) GetFogColor() *Color {
	v := C.WrapEnvironmentGetFogColor(pointer.h)
	vGO := &Color{h: v}
	return vGO
}

// SetFogColor ...
func (pointer *Environment) SetFogColor(v *Color) {
	vToC := v.h
	C.WrapEnvironmentSetFogColor(pointer.h, vToC)
}

// GetFogNear ...
func (pointer *Environment) GetFogNear() float32 {
	v := C.WrapEnvironmentGetFogNear(pointer.h)
	return float32(v)
}

// SetFogNear ...
func (pointer *Environment) SetFogNear(v float32) {
	vToC := C.float(v)
	C.WrapEnvironmentSetFogNear(pointer.h, vToC)
}

// GetFogFar ...
func (pointer *Environment) GetFogFar() float32 {
	v := C.WrapEnvironmentGetFogFar(pointer.h)
	return float32(v)
}

// SetFogFar ...
func (pointer *Environment) SetFogFar(v float32) {
	vToC := C.float(v)
	C.WrapEnvironmentSetFogFar(pointer.h, vToC)
}

// GetBrdfMap ...
func (pointer *Environment) GetBrdfMap() *TextureRef {
	v := C.WrapEnvironmentGetBrdfMap(pointer.h)
	vGO := &TextureRef{h: v}
	return vGO
}

// SetBrdfMap ...
func (pointer *Environment) SetBrdfMap(v *TextureRef) {
	vToC := v.h
	C.WrapEnvironmentSetBrdfMap(pointer.h, vToC)
}

// Free ...
func (pointer *Environment) Free() {
	C.WrapEnvironmentFree(pointer.h)
}

// IsNil ...
func (pointer *Environment) IsNil() bool {
	return pointer.h == C.WrapEnvironment(nil)
}

// SceneForwardPipelinePassViewId  ...
type SceneForwardPipelinePassViewId struct {
	h C.WrapSceneForwardPipelinePassViewId
}

// NewSceneForwardPipelinePassViewId ...
func NewSceneForwardPipelinePassViewId() *SceneForwardPipelinePassViewId {
	retval := C.WrapConstructorSceneForwardPipelinePassViewId()
	retvalGO := &SceneForwardPipelinePassViewId{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneForwardPipelinePassViewId) {
		C.WrapSceneForwardPipelinePassViewIdFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *SceneForwardPipelinePassViewId) Free() {
	C.WrapSceneForwardPipelinePassViewIdFree(pointer.h)
}

// IsNil ...
func (pointer *SceneForwardPipelinePassViewId) IsNil() bool {
	return pointer.h == C.WrapSceneForwardPipelinePassViewId(nil)
}

// SceneForwardPipelineRenderData  Holds all data required to draw a scene with the forward pipeline.  See [harfang.man.ForwardPipeline].
type SceneForwardPipelineRenderData struct {
	h C.WrapSceneForwardPipelineRenderData
}

// NewSceneForwardPipelineRenderData Holds all data required to draw a scene with the forward pipeline.  See [harfang.man.ForwardPipeline].
func NewSceneForwardPipelineRenderData() *SceneForwardPipelineRenderData {
	retval := C.WrapConstructorSceneForwardPipelineRenderData()
	retvalGO := &SceneForwardPipelineRenderData{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneForwardPipelineRenderData) {
		C.WrapSceneForwardPipelineRenderDataFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *SceneForwardPipelineRenderData) Free() {
	C.WrapSceneForwardPipelineRenderDataFree(pointer.h)
}

// IsNil ...
func (pointer *SceneForwardPipelineRenderData) IsNil() bool {
	return pointer.h == C.WrapSceneForwardPipelineRenderData(nil)
}

// ForwardPipelineAAAConfig  ...
type ForwardPipelineAAAConfig struct {
	h C.WrapForwardPipelineAAAConfig
}

// GetTemporalAaWeight ...
func (pointer *ForwardPipelineAAAConfig) GetTemporalAaWeight() float32 {
	v := C.WrapForwardPipelineAAAConfigGetTemporalAaWeight(pointer.h)
	return float32(v)
}

// SetTemporalAaWeight ...
func (pointer *ForwardPipelineAAAConfig) SetTemporalAaWeight(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineAAAConfigSetTemporalAaWeight(pointer.h, vToC)
}

// GetSampleCount ...
func (pointer *ForwardPipelineAAAConfig) GetSampleCount() int32 {
	v := C.WrapForwardPipelineAAAConfigGetSampleCount(pointer.h)
	return int32(v)
}

// SetSampleCount ...
func (pointer *ForwardPipelineAAAConfig) SetSampleCount(v int32) {
	vToC := C.int32_t(v)
	C.WrapForwardPipelineAAAConfigSetSampleCount(pointer.h, vToC)
}

// GetMaxDistance ...
func (pointer *ForwardPipelineAAAConfig) GetMaxDistance() float32 {
	v := C.WrapForwardPipelineAAAConfigGetMaxDistance(pointer.h)
	return float32(v)
}

// SetMaxDistance ...
func (pointer *ForwardPipelineAAAConfig) SetMaxDistance(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineAAAConfigSetMaxDistance(pointer.h, vToC)
}

// GetZThickness ...
func (pointer *ForwardPipelineAAAConfig) GetZThickness() float32 {
	v := C.WrapForwardPipelineAAAConfigGetZThickness(pointer.h)
	return float32(v)
}

// SetZThickness ...
func (pointer *ForwardPipelineAAAConfig) SetZThickness(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineAAAConfigSetZThickness(pointer.h, vToC)
}

// GetBloomThreshold ...
func (pointer *ForwardPipelineAAAConfig) GetBloomThreshold() float32 {
	v := C.WrapForwardPipelineAAAConfigGetBloomThreshold(pointer.h)
	return float32(v)
}

// SetBloomThreshold ...
func (pointer *ForwardPipelineAAAConfig) SetBloomThreshold(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineAAAConfigSetBloomThreshold(pointer.h, vToC)
}

// GetBloomBias ...
func (pointer *ForwardPipelineAAAConfig) GetBloomBias() float32 {
	v := C.WrapForwardPipelineAAAConfigGetBloomBias(pointer.h)
	return float32(v)
}

// SetBloomBias ...
func (pointer *ForwardPipelineAAAConfig) SetBloomBias(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineAAAConfigSetBloomBias(pointer.h, vToC)
}

// GetBloomIntensity ...
func (pointer *ForwardPipelineAAAConfig) GetBloomIntensity() float32 {
	v := C.WrapForwardPipelineAAAConfigGetBloomIntensity(pointer.h)
	return float32(v)
}

// SetBloomIntensity ...
func (pointer *ForwardPipelineAAAConfig) SetBloomIntensity(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineAAAConfigSetBloomIntensity(pointer.h, vToC)
}

// GetMotionBlur ...
func (pointer *ForwardPipelineAAAConfig) GetMotionBlur() float32 {
	v := C.WrapForwardPipelineAAAConfigGetMotionBlur(pointer.h)
	return float32(v)
}

// SetMotionBlur ...
func (pointer *ForwardPipelineAAAConfig) SetMotionBlur(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineAAAConfigSetMotionBlur(pointer.h, vToC)
}

// GetExposure ...
func (pointer *ForwardPipelineAAAConfig) GetExposure() float32 {
	v := C.WrapForwardPipelineAAAConfigGetExposure(pointer.h)
	return float32(v)
}

// SetExposure ...
func (pointer *ForwardPipelineAAAConfig) SetExposure(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineAAAConfigSetExposure(pointer.h, vToC)
}

// GetGamma ...
func (pointer *ForwardPipelineAAAConfig) GetGamma() float32 {
	v := C.WrapForwardPipelineAAAConfigGetGamma(pointer.h)
	return float32(v)
}

// SetGamma ...
func (pointer *ForwardPipelineAAAConfig) SetGamma(v float32) {
	vToC := C.float(v)
	C.WrapForwardPipelineAAAConfigSetGamma(pointer.h, vToC)
}

// NewForwardPipelineAAAConfig ...
func NewForwardPipelineAAAConfig() *ForwardPipelineAAAConfig {
	retval := C.WrapConstructorForwardPipelineAAAConfig()
	retvalGO := &ForwardPipelineAAAConfig{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineAAAConfig) {
		C.WrapForwardPipelineAAAConfigFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *ForwardPipelineAAAConfig) Free() {
	C.WrapForwardPipelineAAAConfigFree(pointer.h)
}

// IsNil ...
func (pointer *ForwardPipelineAAAConfig) IsNil() bool {
	return pointer.h == C.WrapForwardPipelineAAAConfig(nil)
}

// ForwardPipelineAAA  ...
type ForwardPipelineAAA struct {
	h C.WrapForwardPipelineAAA
}

// Free ...
func (pointer *ForwardPipelineAAA) Free() {
	C.WrapForwardPipelineAAAFree(pointer.h)
}

// IsNil ...
func (pointer *ForwardPipelineAAA) IsNil() bool {
	return pointer.h == C.WrapForwardPipelineAAA(nil)
}

// Flip ...
func (pointer *ForwardPipelineAAA) Flip(viewstate *ViewState) {
	viewstateToC := viewstate.h
	C.WrapFlipForwardPipelineAAA(pointer.h, viewstateToC)
}

// NodePairContacts  ...
type NodePairContacts struct {
	h C.WrapNodePairContacts
}

// Free ...
func (pointer *NodePairContacts) Free() {
	C.WrapNodePairContactsFree(pointer.h)
}

// IsNil ...
func (pointer *NodePairContacts) IsNil() bool {
	return pointer.h == C.WrapNodePairContacts(nil)
}

// SceneBullet3Physics  Newton physics for scene physics and collision components.  See [harfang.man.Physics].
type SceneBullet3Physics struct {
	h C.WrapSceneBullet3Physics
}

// NewSceneBullet3Physics Newton physics for scene physics and collision components.  See [harfang.man.Physics].
func NewSceneBullet3Physics() *SceneBullet3Physics {
	retval := C.WrapConstructorSceneBullet3Physics()
	retvalGO := &SceneBullet3Physics{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneBullet3Physics) {
		C.WrapSceneBullet3PhysicsFree(cleanval.h)
	})
	return retvalGO
}

// NewSceneBullet3PhysicsWithThreadCount Newton physics for scene physics and collision components.  See [harfang.man.Physics].
func NewSceneBullet3PhysicsWithThreadCount(threadcount int32) *SceneBullet3Physics {
	threadcountToC := C.int32_t(threadcount)
	retval := C.WrapConstructorSceneBullet3PhysicsWithThreadCount(threadcountToC)
	retvalGO := &SceneBullet3Physics{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneBullet3Physics) {
		C.WrapSceneBullet3PhysicsFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *SceneBullet3Physics) Free() {
	C.WrapSceneBullet3PhysicsFree(pointer.h)
}

// IsNil ...
func (pointer *SceneBullet3Physics) IsNil() bool {
	return pointer.h == C.WrapSceneBullet3Physics(nil)
}

// SceneCreatePhysicsFromFile ...
func (pointer *SceneBullet3Physics) SceneCreatePhysicsFromFile(scene *Scene) {
	sceneToC := scene.h
	C.WrapSceneCreatePhysicsFromFileSceneBullet3Physics(pointer.h, sceneToC)
}

// SceneCreatePhysicsFromAssets ...
func (pointer *SceneBullet3Physics) SceneCreatePhysicsFromAssets(scene *Scene) {
	sceneToC := scene.h
	C.WrapSceneCreatePhysicsFromAssetsSceneBullet3Physics(pointer.h, sceneToC)
}

// NodeCreatePhysicsFromFile ...
func (pointer *SceneBullet3Physics) NodeCreatePhysicsFromFile(node *Node) {
	nodeToC := node.h
	C.WrapNodeCreatePhysicsFromFileSceneBullet3Physics(pointer.h, nodeToC)
}

// NodeCreatePhysicsFromAssets ...
func (pointer *SceneBullet3Physics) NodeCreatePhysicsFromAssets(node *Node) {
	nodeToC := node.h
	C.WrapNodeCreatePhysicsFromAssetsSceneBullet3Physics(pointer.h, nodeToC)
}

// NodeStartTrackingCollisionEvents ...
func (pointer *SceneBullet3Physics) NodeStartTrackingCollisionEvents(node *Node) {
	nodeToC := node.h
	C.WrapNodeStartTrackingCollisionEventsSceneBullet3Physics(pointer.h, nodeToC)
}

// NodeStartTrackingCollisionEventsWithMode ...
func (pointer *SceneBullet3Physics) NodeStartTrackingCollisionEventsWithMode(node *Node, mode CollisionEventTrackingMode) {
	nodeToC := node.h
	modeToC := C.uchar(mode)
	C.WrapNodeStartTrackingCollisionEventsSceneBullet3PhysicsWithMode(pointer.h, nodeToC, modeToC)
}

// NodeStopTrackingCollisionEvents ...
func (pointer *SceneBullet3Physics) NodeStopTrackingCollisionEvents(node *Node) {
	nodeToC := node.h
	C.WrapNodeStopTrackingCollisionEventsSceneBullet3Physics(pointer.h, nodeToC)
}

// NodeDestroyPhysics ...
func (pointer *SceneBullet3Physics) NodeDestroyPhysics(node *Node) {
	nodeToC := node.h
	C.WrapNodeDestroyPhysicsSceneBullet3Physics(pointer.h, nodeToC)
}

// NodeHasBody ...
func (pointer *SceneBullet3Physics) NodeHasBody(node *Node) bool {
	nodeToC := node.h
	retval := C.WrapNodeHasBodySceneBullet3Physics(pointer.h, nodeToC)
	return bool(retval)
}

// StepSimulation ...
func (pointer *SceneBullet3Physics) StepSimulation(displaydt int64) {
	displaydtToC := C.int64_t(displaydt)
	C.WrapStepSimulationSceneBullet3Physics(pointer.h, displaydtToC)
}

// StepSimulationWithStepDt ...
func (pointer *SceneBullet3Physics) StepSimulationWithStepDt(displaydt int64, stepdt int64) {
	displaydtToC := C.int64_t(displaydt)
	stepdtToC := C.int64_t(stepdt)
	C.WrapStepSimulationSceneBullet3PhysicsWithStepDt(pointer.h, displaydtToC, stepdtToC)
}

// StepSimulationWithStepDtMaxStep ...
func (pointer *SceneBullet3Physics) StepSimulationWithStepDtMaxStep(displaydt int64, stepdt int64, maxstep int32) {
	displaydtToC := C.int64_t(displaydt)
	stepdtToC := C.int64_t(stepdt)
	maxstepToC := C.int32_t(maxstep)
	C.WrapStepSimulationSceneBullet3PhysicsWithStepDtMaxStep(pointer.h, displaydtToC, stepdtToC, maxstepToC)
}

// CollectCollisionEvents ...
func (pointer *SceneBullet3Physics) CollectCollisionEvents(scene *Scene, nodepaircontacts *NodePairContacts) {
	sceneToC := scene.h
	nodepaircontactsToC := nodepaircontacts.h
	C.WrapCollectCollisionEventsSceneBullet3Physics(pointer.h, sceneToC, nodepaircontactsToC)
}

// SyncTransformsFromScene ...
func (pointer *SceneBullet3Physics) SyncTransformsFromScene(scene *Scene) {
	sceneToC := scene.h
	C.WrapSyncTransformsFromSceneSceneBullet3Physics(pointer.h, sceneToC)
}

// SyncTransformsToScene ...
func (pointer *SceneBullet3Physics) SyncTransformsToScene(scene *Scene) {
	sceneToC := scene.h
	C.WrapSyncTransformsToSceneSceneBullet3Physics(pointer.h, sceneToC)
}

// GarbageCollect ...
func (pointer *SceneBullet3Physics) GarbageCollect(scene *Scene) int32 {
	sceneToC := scene.h
	retval := C.WrapGarbageCollectSceneBullet3Physics(pointer.h, sceneToC)
	return int32(retval)
}

// GarbageCollectResources ...
func (pointer *SceneBullet3Physics) GarbageCollectResources() int32 {
	retval := C.WrapGarbageCollectResourcesSceneBullet3Physics(pointer.h)
	return int32(retval)
}

// ClearNodes ...
func (pointer *SceneBullet3Physics) ClearNodes() {
	C.WrapClearNodesSceneBullet3Physics(pointer.h)
}

// Clear ...
func (pointer *SceneBullet3Physics) Clear() {
	C.WrapClearSceneBullet3Physics(pointer.h)
}

// NodeWake ...
func (pointer *SceneBullet3Physics) NodeWake(node *Node) {
	nodeToC := node.h
	C.WrapNodeWakeSceneBullet3Physics(pointer.h, nodeToC)
}

// NodeSetDeactivation ...
func (pointer *SceneBullet3Physics) NodeSetDeactivation(node *Node, enable bool) {
	nodeToC := node.h
	enableToC := C.bool(enable)
	C.WrapNodeSetDeactivationSceneBullet3Physics(pointer.h, nodeToC, enableToC)
}

// NodeGetDeactivation ...
func (pointer *SceneBullet3Physics) NodeGetDeactivation(node *Node) bool {
	nodeToC := node.h
	retval := C.WrapNodeGetDeactivationSceneBullet3Physics(pointer.h, nodeToC)
	return bool(retval)
}

// NodeResetWorld ...
func (pointer *SceneBullet3Physics) NodeResetWorld(node *Node, world *Mat4) {
	nodeToC := node.h
	worldToC := world.h
	C.WrapNodeResetWorldSceneBullet3Physics(pointer.h, nodeToC, worldToC)
}

// NodeTeleport ...
func (pointer *SceneBullet3Physics) NodeTeleport(node *Node, world *Mat4) {
	nodeToC := node.h
	worldToC := world.h
	C.WrapNodeTeleportSceneBullet3Physics(pointer.h, nodeToC, worldToC)
}

// NodeAddForce ...
func (pointer *SceneBullet3Physics) NodeAddForce(node *Node, F *Vec3) {
	nodeToC := node.h
	FToC := F.h
	C.WrapNodeAddForceSceneBullet3Physics(pointer.h, nodeToC, FToC)
}

// NodeAddForceWithWorldPos ...
func (pointer *SceneBullet3Physics) NodeAddForceWithWorldPos(node *Node, F *Vec3, worldpos *Vec3) {
	nodeToC := node.h
	FToC := F.h
	worldposToC := worldpos.h
	C.WrapNodeAddForceSceneBullet3PhysicsWithWorldPos(pointer.h, nodeToC, FToC, worldposToC)
}

// NodeAddImpulse ...
func (pointer *SceneBullet3Physics) NodeAddImpulse(node *Node, dtvelocity *Vec3) {
	nodeToC := node.h
	dtvelocityToC := dtvelocity.h
	C.WrapNodeAddImpulseSceneBullet3Physics(pointer.h, nodeToC, dtvelocityToC)
}

// NodeAddImpulseWithWorldPos ...
func (pointer *SceneBullet3Physics) NodeAddImpulseWithWorldPos(node *Node, dtvelocity *Vec3, worldpos *Vec3) {
	nodeToC := node.h
	dtvelocityToC := dtvelocity.h
	worldposToC := worldpos.h
	C.WrapNodeAddImpulseSceneBullet3PhysicsWithWorldPos(pointer.h, nodeToC, dtvelocityToC, worldposToC)
}

// NodeAddTorque ...
func (pointer *SceneBullet3Physics) NodeAddTorque(node *Node, T *Vec3) {
	nodeToC := node.h
	TToC := T.h
	C.WrapNodeAddTorqueSceneBullet3Physics(pointer.h, nodeToC, TToC)
}

// NodeAddTorqueImpulse ...
func (pointer *SceneBullet3Physics) NodeAddTorqueImpulse(node *Node, dtangularvelocity *Vec3) {
	nodeToC := node.h
	dtangularvelocityToC := dtangularvelocity.h
	C.WrapNodeAddTorqueImpulseSceneBullet3Physics(pointer.h, nodeToC, dtangularvelocityToC)
}

// NodeGetPointVelocity ...
func (pointer *SceneBullet3Physics) NodeGetPointVelocity(node *Node, worldpos *Vec3) *Vec3 {
	nodeToC := node.h
	worldposToC := worldpos.h
	retval := C.WrapNodeGetPointVelocitySceneBullet3Physics(pointer.h, nodeToC, worldposToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// NodeGetLinearVelocity ...
func (pointer *SceneBullet3Physics) NodeGetLinearVelocity(node *Node) *Vec3 {
	nodeToC := node.h
	retval := C.WrapNodeGetLinearVelocitySceneBullet3Physics(pointer.h, nodeToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// NodeSetLinearVelocity ...
func (pointer *SceneBullet3Physics) NodeSetLinearVelocity(node *Node, V *Vec3) {
	nodeToC := node.h
	VToC := V.h
	C.WrapNodeSetLinearVelocitySceneBullet3Physics(pointer.h, nodeToC, VToC)
}

// NodeGetAngularVelocity ...
func (pointer *SceneBullet3Physics) NodeGetAngularVelocity(node *Node) *Vec3 {
	nodeToC := node.h
	retval := C.WrapNodeGetAngularVelocitySceneBullet3Physics(pointer.h, nodeToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// NodeSetAngularVelocity ...
func (pointer *SceneBullet3Physics) NodeSetAngularVelocity(node *Node, W *Vec3) {
	nodeToC := node.h
	WToC := W.h
	C.WrapNodeSetAngularVelocitySceneBullet3Physics(pointer.h, nodeToC, WToC)
}

// NodeGetLinearFactor ...
func (pointer *SceneBullet3Physics) NodeGetLinearFactor(node *Node) *Vec3 {
	nodeToC := node.h
	retval := C.WrapNodeGetLinearFactorSceneBullet3Physics(pointer.h, nodeToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// NodeSetLinearFactor ...
func (pointer *SceneBullet3Physics) NodeSetLinearFactor(node *Node, k *Vec3) {
	nodeToC := node.h
	kToC := k.h
	C.WrapNodeSetLinearFactorSceneBullet3Physics(pointer.h, nodeToC, kToC)
}

// NodeGetAngularFactor ...
func (pointer *SceneBullet3Physics) NodeGetAngularFactor(node *Node) *Vec3 {
	nodeToC := node.h
	retval := C.WrapNodeGetAngularFactorSceneBullet3Physics(pointer.h, nodeToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// NodeSetAngularFactor ...
func (pointer *SceneBullet3Physics) NodeSetAngularFactor(node *Node, k *Vec3) {
	nodeToC := node.h
	kToC := k.h
	C.WrapNodeSetAngularFactorSceneBullet3Physics(pointer.h, nodeToC, kToC)
}

// NodeCollideWorld ...
func (pointer *SceneBullet3Physics) NodeCollideWorld(node *Node, world *Mat4) *NodePairContacts {
	nodeToC := node.h
	worldToC := world.h
	retval := C.WrapNodeCollideWorldSceneBullet3Physics(pointer.h, nodeToC, worldToC)
	retvalGO := &NodePairContacts{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodePairContacts) {
		C.WrapNodePairContactsFree(cleanval.h)
	})
	return retvalGO
}

// NodeCollideWorldWithMaxContact ...
func (pointer *SceneBullet3Physics) NodeCollideWorldWithMaxContact(node *Node, world *Mat4, maxcontact int32) *NodePairContacts {
	nodeToC := node.h
	worldToC := world.h
	maxcontactToC := C.int32_t(maxcontact)
	retval := C.WrapNodeCollideWorldSceneBullet3PhysicsWithMaxContact(pointer.h, nodeToC, worldToC, maxcontactToC)
	retvalGO := &NodePairContacts{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodePairContacts) {
		C.WrapNodePairContactsFree(cleanval.h)
	})
	return retvalGO
}

// RaycastFirstHit ...
func (pointer *SceneBullet3Physics) RaycastFirstHit(scene *Scene, p0 *Vec3, p1 *Vec3) *RaycastOut {
	sceneToC := scene.h
	p0ToC := p0.h
	p1ToC := p1.h
	retval := C.WrapRaycastFirstHitSceneBullet3Physics(pointer.h, sceneToC, p0ToC, p1ToC)
	retvalGO := &RaycastOut{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RaycastOut) {
		C.WrapRaycastOutFree(cleanval.h)
	})
	return retvalGO
}

// RaycastAllHits ...
func (pointer *SceneBullet3Physics) RaycastAllHits(scene *Scene, p0 *Vec3, p1 *Vec3) *RaycastOutList {
	sceneToC := scene.h
	p0ToC := p0.h
	p1ToC := p1.h
	retval := C.WrapRaycastAllHitsSceneBullet3Physics(pointer.h, sceneToC, p0ToC, p1ToC)
	retvalGO := &RaycastOutList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RaycastOutList) {
		C.WrapRaycastOutListFree(cleanval.h)
	})
	return retvalGO
}

// RenderCollision ...
func (pointer *SceneBullet3Physics) RenderCollision(viewid uint16, vtxlayout *VertexLayout, prg *ProgramHandle, renderstate *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	vtxlayoutToC := vtxlayout.h
	prgToC := prg.h
	renderstateToC := renderstate.h
	depthToC := C.uint32_t(depth)
	C.WrapRenderCollisionSceneBullet3Physics(pointer.h, viewidToC, vtxlayoutToC, prgToC, renderstateToC, depthToC)
}

// SceneLuaVM  Lua VM for scene script components.  See [harfang.man.Scripting].
type SceneLuaVM struct {
	h C.WrapSceneLuaVM
}

// NewSceneLuaVM Lua VM for scene script components.  See [harfang.man.Scripting].
func NewSceneLuaVM() *SceneLuaVM {
	retval := C.WrapConstructorSceneLuaVM()
	retvalGO := &SceneLuaVM{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneLuaVM) {
		C.WrapSceneLuaVMFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *SceneLuaVM) Free() {
	C.WrapSceneLuaVMFree(pointer.h)
}

// IsNil ...
func (pointer *SceneLuaVM) IsNil() bool {
	return pointer.h == C.WrapSceneLuaVM(nil)
}

// CreateScriptFromSource ...
func (pointer *SceneLuaVM) CreateScriptFromSource(scene *Scene, script *Script, src string) bool {
	sceneToC := scene.h
	scriptToC := script.h
	srcToC, idFinsrcToC := wrapString(src)
	defer idFinsrcToC()
	retval := C.WrapCreateScriptFromSourceSceneLuaVM(pointer.h, sceneToC, scriptToC, srcToC)
	return bool(retval)
}

// CreateScriptFromFile ...
func (pointer *SceneLuaVM) CreateScriptFromFile(scene *Scene, script *Script) bool {
	sceneToC := scene.h
	scriptToC := script.h
	retval := C.WrapCreateScriptFromFileSceneLuaVM(pointer.h, sceneToC, scriptToC)
	return bool(retval)
}

// CreateScriptFromAssets ...
func (pointer *SceneLuaVM) CreateScriptFromAssets(scene *Scene, script *Script) bool {
	sceneToC := scene.h
	scriptToC := script.h
	retval := C.WrapCreateScriptFromAssetsSceneLuaVM(pointer.h, sceneToC, scriptToC)
	return bool(retval)
}

// CreateNodeScriptsFromFile ...
func (pointer *SceneLuaVM) CreateNodeScriptsFromFile(scene *Scene, node *Node) *ScriptList {
	sceneToC := scene.h
	nodeToC := node.h
	retval := C.WrapCreateNodeScriptsFromFileSceneLuaVM(pointer.h, sceneToC, nodeToC)
	retvalGO := &ScriptList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScriptList) {
		C.WrapScriptListFree(cleanval.h)
	})
	return retvalGO
}

// CreateNodeScriptsFromAssets ...
func (pointer *SceneLuaVM) CreateNodeScriptsFromAssets(scene *Scene, node *Node) *ScriptList {
	sceneToC := scene.h
	nodeToC := node.h
	retval := C.WrapCreateNodeScriptsFromAssetsSceneLuaVM(pointer.h, sceneToC, nodeToC)
	retvalGO := &ScriptList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScriptList) {
		C.WrapScriptListFree(cleanval.h)
	})
	return retvalGO
}

// SceneCreateScriptsFromFile ...
func (pointer *SceneLuaVM) SceneCreateScriptsFromFile(scene *Scene) *ScriptList {
	sceneToC := scene.h
	retval := C.WrapSceneCreateScriptsFromFileSceneLuaVM(pointer.h, sceneToC)
	retvalGO := &ScriptList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScriptList) {
		C.WrapScriptListFree(cleanval.h)
	})
	return retvalGO
}

// SceneCreateScriptsFromAssets ...
func (pointer *SceneLuaVM) SceneCreateScriptsFromAssets(scene *Scene) *ScriptList {
	sceneToC := scene.h
	retval := C.WrapSceneCreateScriptsFromAssetsSceneLuaVM(pointer.h, sceneToC)
	retvalGO := &ScriptList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScriptList) {
		C.WrapScriptListFree(cleanval.h)
	})
	return retvalGO
}

// GarbageCollect ...
func (pointer *SceneLuaVM) GarbageCollect(scene *Scene) *ScriptList {
	sceneToC := scene.h
	retval := C.WrapGarbageCollectSceneLuaVM(pointer.h, sceneToC)
	retvalGO := &ScriptList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ScriptList) {
		C.WrapScriptListFree(cleanval.h)
	})
	return retvalGO
}

// DestroyScripts ...
func (pointer *SceneLuaVM) DestroyScripts(scripts *ScriptList) {
	scriptsToC := scripts.h
	C.WrapDestroyScriptsSceneLuaVM(pointer.h, scriptsToC)
}

// GetScriptInterface ...
func (pointer *SceneLuaVM) GetScriptInterface(script *Script) *StringList {
	scriptToC := script.h
	retval := C.WrapGetScriptInterfaceSceneLuaVM(pointer.h, scriptToC)
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// GetScriptCount ...
func (pointer *SceneLuaVM) GetScriptCount() int32 {
	retval := C.WrapGetScriptCountSceneLuaVM(pointer.h)
	return int32(retval)
}

// GetScriptEnv ...
func (pointer *SceneLuaVM) GetScriptEnv(script *Script) *LuaObject {
	scriptToC := script.h
	retval := C.WrapGetScriptEnvSceneLuaVM(pointer.h, scriptToC)
	retvalGO := &LuaObject{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *LuaObject) {
		C.WrapLuaObjectFree(cleanval.h)
	})
	return retvalGO
}

// GetScriptValue ...
func (pointer *SceneLuaVM) GetScriptValue(script *Script, name string) *LuaObject {
	scriptToC := script.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapGetScriptValueSceneLuaVM(pointer.h, scriptToC, nameToC)
	retvalGO := &LuaObject{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *LuaObject) {
		C.WrapLuaObjectFree(cleanval.h)
	})
	return retvalGO
}

// SetScriptValue ...
func (pointer *SceneLuaVM) SetScriptValue(script *Script, name string, value *LuaObject) bool {
	scriptToC := script.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	valueToC := value.h
	retval := C.WrapSetScriptValueSceneLuaVM(pointer.h, scriptToC, nameToC, valueToC)
	return bool(retval)
}

// SetScriptValueWithNotify ...
func (pointer *SceneLuaVM) SetScriptValueWithNotify(script *Script, name string, value *LuaObject, notify bool) bool {
	scriptToC := script.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	valueToC := value.h
	notifyToC := C.bool(notify)
	retval := C.WrapSetScriptValueSceneLuaVMWithNotify(pointer.h, scriptToC, nameToC, valueToC, notifyToC)
	return bool(retval)
}

// Call ...
func (pointer *SceneLuaVM) Call(script *Script, function string, args *LuaObjectList) (bool, *LuaObjectList) {
	scriptToC := script.h
	functionToC, idFinfunctionToC := wrapString(function)
	defer idFinfunctionToC()
	argsToC := args.h
	retvals := NewLuaObjectList()
	retvalsToC := retvals.h
	retval := C.WrapCallSceneLuaVM(pointer.h, scriptToC, functionToC, argsToC, retvalsToC)
	return bool(retval), retvals
}

// CallWithSliceOfArgs ...
func (pointer *SceneLuaVM) CallWithSliceOfArgs(script *Script, function string, SliceOfargs GoSliceOfLuaObject) (bool, *LuaObjectList) {
	scriptToC := script.h
	functionToC, idFinfunctionToC := wrapString(function)
	defer idFinfunctionToC()
	var SliceOfargsPointer []C.WrapLuaObject
	for _, s := range SliceOfargs {
		SliceOfargsPointer = append(SliceOfargsPointer, s.h)
	}
	SliceOfargsPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfargsPointer))
	SliceOfargsPointerToCSize := C.size_t(SliceOfargsPointerToC.Len)
	SliceOfargsPointerToCBuf := (*C.WrapLuaObject)(unsafe.Pointer(SliceOfargsPointerToC.Data))
	retvals := NewLuaObjectList()
	retvalsToC := retvals.h
	retval := C.WrapCallSceneLuaVMWithSliceOfArgs(pointer.h, scriptToC, functionToC, SliceOfargsPointerToCSize, SliceOfargsPointerToCBuf, retvalsToC)
	return bool(retval), retvals
}

// MakeLuaObject ...
func (pointer *SceneLuaVM) MakeLuaObject() *LuaObject {
	retval := C.WrapMakeLuaObjectSceneLuaVM(pointer.h)
	retvalGO := &LuaObject{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *LuaObject) {
		C.WrapLuaObjectFree(cleanval.h)
	})
	return retvalGO
}

// SceneClocks  Holds clocks for the different scene systems.  This is required as some system such as the physics system may run at a different rate than the scene.
type SceneClocks struct {
	h C.WrapSceneClocks
}

// NewSceneClocks Holds clocks for the different scene systems.  This is required as some system such as the physics system may run at a different rate than the scene.
func NewSceneClocks() *SceneClocks {
	retval := C.WrapConstructorSceneClocks()
	retvalGO := &SceneClocks{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SceneClocks) {
		C.WrapSceneClocksFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *SceneClocks) Free() {
	C.WrapSceneClocksFree(pointer.h)
}

// IsNil ...
func (pointer *SceneClocks) IsNil() bool {
	return pointer.h == C.WrapSceneClocks(nil)
}

// MouseState  ...
type MouseState struct {
	h C.WrapMouseState
}

// Free ...
func (pointer *MouseState) Free() {
	C.WrapMouseStateFree(pointer.h)
}

// IsNil ...
func (pointer *MouseState) IsNil() bool {
	return pointer.h == C.WrapMouseState(nil)
}

// X ...
func (pointer *MouseState) X() int32 {
	retval := C.WrapXMouseState(pointer.h)
	return int32(retval)
}

// Y ...
func (pointer *MouseState) Y() int32 {
	retval := C.WrapYMouseState(pointer.h)
	return int32(retval)
}

// Button ...
func (pointer *MouseState) Button(btn MouseButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapButtonMouseState(pointer.h, btnToC)
	return bool(retval)
}

// Wheel ...
func (pointer *MouseState) Wheel() int32 {
	retval := C.WrapWheelMouseState(pointer.h)
	return int32(retval)
}

// HWheel ...
func (pointer *MouseState) HWheel() int32 {
	retval := C.WrapHWheelMouseState(pointer.h)
	return int32(retval)
}

// Mouse  Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetMouseNames] to query for available mouse devices.
type Mouse struct {
	h C.WrapMouse
}

// NewMouse Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetMouseNames] to query for available mouse devices.
func NewMouse() *Mouse {
	retval := C.WrapConstructorMouse()
	retvalGO := &Mouse{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mouse) {
		C.WrapMouseFree(cleanval.h)
	})
	return retvalGO
}

// NewMouseWithName Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetMouseNames] to query for available mouse devices.
func NewMouseWithName(name string) *Mouse {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapConstructorMouseWithName(nameToC)
	retvalGO := &Mouse{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mouse) {
		C.WrapMouseFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Mouse) Free() {
	C.WrapMouseFree(pointer.h)
}

// IsNil ...
func (pointer *Mouse) IsNil() bool {
	return pointer.h == C.WrapMouse(nil)
}

// X ...
func (pointer *Mouse) X() int32 {
	retval := C.WrapXMouse(pointer.h)
	return int32(retval)
}

// Y ...
func (pointer *Mouse) Y() int32 {
	retval := C.WrapYMouse(pointer.h)
	return int32(retval)
}

// DtX ...
func (pointer *Mouse) DtX() int32 {
	retval := C.WrapDtXMouse(pointer.h)
	return int32(retval)
}

// DtY ...
func (pointer *Mouse) DtY() int32 {
	retval := C.WrapDtYMouse(pointer.h)
	return int32(retval)
}

// Down ...
func (pointer *Mouse) Down(button int32) bool {
	buttonToC := C.int32_t(button)
	retval := C.WrapDownMouse(pointer.h, buttonToC)
	return bool(retval)
}

// Pressed ...
func (pointer *Mouse) Pressed(button int32) bool {
	buttonToC := C.int32_t(button)
	retval := C.WrapPressedMouse(pointer.h, buttonToC)
	return bool(retval)
}

// Released ...
func (pointer *Mouse) Released(button int32) bool {
	buttonToC := C.int32_t(button)
	retval := C.WrapReleasedMouse(pointer.h, buttonToC)
	return bool(retval)
}

// Wheel ...
func (pointer *Mouse) Wheel() int32 {
	retval := C.WrapWheelMouse(pointer.h)
	return int32(retval)
}

// HWheel ...
func (pointer *Mouse) HWheel() int32 {
	retval := C.WrapHWheelMouse(pointer.h)
	return int32(retval)
}

// Update ...
func (pointer *Mouse) Update() {
	C.WrapUpdateMouse(pointer.h)
}

// GetState ...
func (pointer *Mouse) GetState() *MouseState {
	retval := C.WrapGetStateMouse(pointer.h)
	retvalGO := &MouseState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MouseState) {
		C.WrapMouseStateFree(cleanval.h)
	})
	return retvalGO
}

// GetOldState ...
func (pointer *Mouse) GetOldState() *MouseState {
	retval := C.WrapGetOldStateMouse(pointer.h)
	retvalGO := &MouseState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MouseState) {
		C.WrapMouseStateFree(cleanval.h)
	})
	return retvalGO
}

// KeyboardState  ...
type KeyboardState struct {
	h C.WrapKeyboardState
}

// Free ...
func (pointer *KeyboardState) Free() {
	C.WrapKeyboardStateFree(pointer.h)
}

// IsNil ...
func (pointer *KeyboardState) IsNil() bool {
	return pointer.h == C.WrapKeyboardState(nil)
}

// Key ...
func (pointer *KeyboardState) Key(key Key) bool {
	keyToC := C.int32_t(key)
	retval := C.WrapKeyKeyboardState(pointer.h, keyToC)
	return bool(retval)
}

// Keyboard  Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetKeyboardNames] to query for available keyboard devices.
type Keyboard struct {
	h C.WrapKeyboard
}

// NewKeyboard Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetKeyboardNames] to query for available keyboard devices.
func NewKeyboard() *Keyboard {
	retval := C.WrapConstructorKeyboard()
	retvalGO := &Keyboard{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Keyboard) {
		C.WrapKeyboardFree(cleanval.h)
	})
	return retvalGO
}

// NewKeyboardWithName Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetKeyboardNames] to query for available keyboard devices.
func NewKeyboardWithName(name string) *Keyboard {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapConstructorKeyboardWithName(nameToC)
	retvalGO := &Keyboard{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Keyboard) {
		C.WrapKeyboardFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Keyboard) Free() {
	C.WrapKeyboardFree(pointer.h)
}

// IsNil ...
func (pointer *Keyboard) IsNil() bool {
	return pointer.h == C.WrapKeyboard(nil)
}

// Down ...
func (pointer *Keyboard) Down(key Key) bool {
	keyToC := C.int32_t(key)
	retval := C.WrapDownKeyboard(pointer.h, keyToC)
	return bool(retval)
}

// Pressed ...
func (pointer *Keyboard) Pressed(key Key) bool {
	keyToC := C.int32_t(key)
	retval := C.WrapPressedKeyboard(pointer.h, keyToC)
	return bool(retval)
}

// Released ...
func (pointer *Keyboard) Released(key Key) bool {
	keyToC := C.int32_t(key)
	retval := C.WrapReleasedKeyboard(pointer.h, keyToC)
	return bool(retval)
}

// Update ...
func (pointer *Keyboard) Update() {
	C.WrapUpdateKeyboard(pointer.h)
}

// GetState ...
func (pointer *Keyboard) GetState() *KeyboardState {
	retval := C.WrapGetStateKeyboard(pointer.h)
	retvalGO := &KeyboardState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *KeyboardState) {
		C.WrapKeyboardStateFree(cleanval.h)
	})
	return retvalGO
}

// GetOldState ...
func (pointer *Keyboard) GetOldState() *KeyboardState {
	retval := C.WrapGetOldStateKeyboard(pointer.h)
	retvalGO := &KeyboardState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *KeyboardState) {
		C.WrapKeyboardStateFree(cleanval.h)
	})
	return retvalGO
}

// TextInputCallbackConnection  ...
type TextInputCallbackConnection struct {
	h C.WrapTextInputCallbackConnection
}

// Free ...
func (pointer *TextInputCallbackConnection) Free() {
	C.WrapTextInputCallbackConnectionFree(pointer.h)
}

// IsNil ...
func (pointer *TextInputCallbackConnection) IsNil() bool {
	return pointer.h == C.WrapTextInputCallbackConnection(nil)
}

// SignalReturningVoidTakingConstCharPtr  ...
type SignalReturningVoidTakingConstCharPtr struct {
	h C.WrapSignalReturningVoidTakingConstCharPtr
}

// Free ...
func (pointer *SignalReturningVoidTakingConstCharPtr) Free() {
	C.WrapSignalReturningVoidTakingConstCharPtrFree(pointer.h)
}

// IsNil ...
func (pointer *SignalReturningVoidTakingConstCharPtr) IsNil() bool {
	return pointer.h == C.WrapSignalReturningVoidTakingConstCharPtr(nil)
}

// Connect ...
func (pointer *SignalReturningVoidTakingConstCharPtr) Connect(listener unsafe.Pointer) *TextInputCallbackConnection {
	listenerToC := (C.WrapFunctionReturningVoidTakingConstCharPtr)(listener)
	retval := C.WrapConnectSignalReturningVoidTakingConstCharPtr(pointer.h, listenerToC)
	retvalGO := &TextInputCallbackConnection{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *TextInputCallbackConnection) {
		C.WrapTextInputCallbackConnectionFree(cleanval.h)
	})
	return retvalGO
}

// Disconnect ...
func (pointer *SignalReturningVoidTakingConstCharPtr) Disconnect(connection *TextInputCallbackConnection) {
	connectionToC := connection.h
	C.WrapDisconnectSignalReturningVoidTakingConstCharPtr(pointer.h, connectionToC)
}

// DisconnectAll ...
func (pointer *SignalReturningVoidTakingConstCharPtr) DisconnectAll() {
	C.WrapDisconnectAllSignalReturningVoidTakingConstCharPtr(pointer.h)
}

// Emit ...
func (pointer *SignalReturningVoidTakingConstCharPtr) Emit(arg0 string) {
	arg0ToC, idFinarg0ToC := wrapString(arg0)
	defer idFinarg0ToC()
	C.WrapEmitSignalReturningVoidTakingConstCharPtr(pointer.h, arg0ToC)
}

// GetListenerCount ...
func (pointer *SignalReturningVoidTakingConstCharPtr) GetListenerCount() int32 {
	retval := C.WrapGetListenerCountSignalReturningVoidTakingConstCharPtr(pointer.h)
	return int32(retval)
}

// GamepadState  ...
type GamepadState struct {
	h C.WrapGamepadState
}

// Free ...
func (pointer *GamepadState) Free() {
	C.WrapGamepadStateFree(pointer.h)
}

// IsNil ...
func (pointer *GamepadState) IsNil() bool {
	return pointer.h == C.WrapGamepadState(nil)
}

// IsConnected ...
func (pointer *GamepadState) IsConnected() bool {
	retval := C.WrapIsConnectedGamepadState(pointer.h)
	return bool(retval)
}

// Axes ...
func (pointer *GamepadState) Axes(idx GamepadAxes) float32 {
	idxToC := C.int32_t(idx)
	retval := C.WrapAxesGamepadState(pointer.h, idxToC)
	return float32(retval)
}

// Button ...
func (pointer *GamepadState) Button(btn GamepadButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapButtonGamepadState(pointer.h, btnToC)
	return bool(retval)
}

// Gamepad  Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetGamepadNames] to query for available gamepad devices.
type Gamepad struct {
	h C.WrapGamepad
}

// NewGamepad Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetGamepadNames] to query for available gamepad devices.
func NewGamepad() *Gamepad {
	retval := C.WrapConstructorGamepad()
	retvalGO := &Gamepad{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Gamepad) {
		C.WrapGamepadFree(cleanval.h)
	})
	return retvalGO
}

// NewGamepadWithName Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetGamepadNames] to query for available gamepad devices.
func NewGamepadWithName(name string) *Gamepad {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapConstructorGamepadWithName(nameToC)
	retvalGO := &Gamepad{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Gamepad) {
		C.WrapGamepadFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Gamepad) Free() {
	C.WrapGamepadFree(pointer.h)
}

// IsNil ...
func (pointer *Gamepad) IsNil() bool {
	return pointer.h == C.WrapGamepad(nil)
}

// IsConnected Gamepad is currently connected.
func (pointer *Gamepad) IsConnected() bool {
	retval := C.WrapIsConnectedGamepad(pointer.h)
	return bool(retval)
}

// Connected Gamepad was connected since the last update.
func (pointer *Gamepad) Connected() bool {
	retval := C.WrapConnectedGamepad(pointer.h)
	return bool(retval)
}

// Disconnected Gamepad was disconnected since the last update.
func (pointer *Gamepad) Disconnected() bool {
	retval := C.WrapDisconnectedGamepad(pointer.h)
	return bool(retval)
}

// Axes Return the value of a gamepad axis.
func (pointer *Gamepad) Axes(axis GamepadAxes) float32 {
	axisToC := C.int32_t(axis)
	retval := C.WrapAxesGamepad(pointer.h, axisToC)
	return float32(retval)
}

// DtAxes ...
func (pointer *Gamepad) DtAxes(axis GamepadAxes) float32 {
	axisToC := C.int32_t(axis)
	retval := C.WrapDtAxesGamepad(pointer.h, axisToC)
	return float32(retval)
}

// Down ...
func (pointer *Gamepad) Down(btn GamepadButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapDownGamepad(pointer.h, btnToC)
	return bool(retval)
}

// Pressed ...
func (pointer *Gamepad) Pressed(btn GamepadButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapPressedGamepad(pointer.h, btnToC)
	return bool(retval)
}

// Released ...
func (pointer *Gamepad) Released(btn GamepadButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapReleasedGamepad(pointer.h, btnToC)
	return bool(retval)
}

// Update ...
func (pointer *Gamepad) Update() {
	C.WrapUpdateGamepad(pointer.h)
}

// JoystickState  ...
type JoystickState struct {
	h C.WrapJoystickState
}

// Free ...
func (pointer *JoystickState) Free() {
	C.WrapJoystickStateFree(pointer.h)
}

// IsNil ...
func (pointer *JoystickState) IsNil() bool {
	return pointer.h == C.WrapJoystickState(nil)
}

// IsConnected ...
func (pointer *JoystickState) IsConnected() bool {
	retval := C.WrapIsConnectedJoystickState(pointer.h)
	return bool(retval)
}

// Axes ...
func (pointer *JoystickState) Axes(idx int32) float32 {
	idxToC := C.int32_t(idx)
	retval := C.WrapAxesJoystickState(pointer.h, idxToC)
	return float32(retval)
}

// Button ...
func (pointer *JoystickState) Button(btn int32) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapButtonJoystickState(pointer.h, btnToC)
	return bool(retval)
}

// Joystick  ...
type Joystick struct {
	h C.WrapJoystick
}

// NewJoystick ...
func NewJoystick() *Joystick {
	retval := C.WrapConstructorJoystick()
	retvalGO := &Joystick{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Joystick) {
		C.WrapJoystickFree(cleanval.h)
	})
	return retvalGO
}

// NewJoystickWithName ...
func NewJoystickWithName(name string) *Joystick {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapConstructorJoystickWithName(nameToC)
	retvalGO := &Joystick{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Joystick) {
		C.WrapJoystickFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Joystick) Free() {
	C.WrapJoystickFree(pointer.h)
}

// IsNil ...
func (pointer *Joystick) IsNil() bool {
	return pointer.h == C.WrapJoystick(nil)
}

// IsConnected ...
func (pointer *Joystick) IsConnected() bool {
	retval := C.WrapIsConnectedJoystick(pointer.h)
	return bool(retval)
}

// Connected ...
func (pointer *Joystick) Connected() bool {
	retval := C.WrapConnectedJoystick(pointer.h)
	return bool(retval)
}

// Disconnected ...
func (pointer *Joystick) Disconnected() bool {
	retval := C.WrapDisconnectedJoystick(pointer.h)
	return bool(retval)
}

// AxesCount ...
func (pointer *Joystick) AxesCount() int32 {
	retval := C.WrapAxesCountJoystick(pointer.h)
	return int32(retval)
}

// Axes ...
func (pointer *Joystick) Axes(axis int32) float32 {
	axisToC := C.int32_t(axis)
	retval := C.WrapAxesJoystick(pointer.h, axisToC)
	return float32(retval)
}

// DtAxes ...
func (pointer *Joystick) DtAxes(axis int32) float32 {
	axisToC := C.int32_t(axis)
	retval := C.WrapDtAxesJoystick(pointer.h, axisToC)
	return float32(retval)
}

// ButtonsCount ...
func (pointer *Joystick) ButtonsCount() int32 {
	retval := C.WrapButtonsCountJoystick(pointer.h)
	return int32(retval)
}

// Down ...
func (pointer *Joystick) Down(btn int32) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapDownJoystick(pointer.h, btnToC)
	return bool(retval)
}

// Pressed ...
func (pointer *Joystick) Pressed(btn int32) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapPressedJoystick(pointer.h, btnToC)
	return bool(retval)
}

// Released ...
func (pointer *Joystick) Released(btn int32) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapReleasedJoystick(pointer.h, btnToC)
	return bool(retval)
}

// Update ...
func (pointer *Joystick) Update() {
	C.WrapUpdateJoystick(pointer.h)
}

// GetDeviceName ...
func (pointer *Joystick) GetDeviceName() string {
	retval := C.WrapGetDeviceNameJoystick(pointer.h)
	return C.GoString(retval)
}

// VRControllerState  ...
type VRControllerState struct {
	h C.WrapVRControllerState
}

// Free ...
func (pointer *VRControllerState) Free() {
	C.WrapVRControllerStateFree(pointer.h)
}

// IsNil ...
func (pointer *VRControllerState) IsNil() bool {
	return pointer.h == C.WrapVRControllerState(nil)
}

// IsConnected ...
func (pointer *VRControllerState) IsConnected() bool {
	retval := C.WrapIsConnectedVRControllerState(pointer.h)
	return bool(retval)
}

// World ...
func (pointer *VRControllerState) World() *Mat4 {
	retval := C.WrapWorldVRControllerState(pointer.h)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Pressed ...
func (pointer *VRControllerState) Pressed(btn VRControllerButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapPressedVRControllerState(pointer.h, btnToC)
	return bool(retval)
}

// Touched ...
func (pointer *VRControllerState) Touched(btn VRControllerButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapTouchedVRControllerState(pointer.h, btnToC)
	return bool(retval)
}

// Surface ...
func (pointer *VRControllerState) Surface(idx int32) *Vec2 {
	idxToC := C.int32_t(idx)
	retval := C.WrapSurfaceVRControllerState(pointer.h, idxToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// VRController  Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetVRControllerNames] to query for available VR controller devices.
type VRController struct {
	h C.WrapVRController
}

// NewVRController Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetVRControllerNames] to query for available VR controller devices.
func NewVRController() *VRController {
	retval := C.WrapConstructorVRController()
	retvalGO := &VRController{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VRController) {
		C.WrapVRControllerFree(cleanval.h)
	})
	return retvalGO
}

// NewVRControllerWithName Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetVRControllerNames] to query for available VR controller devices.
func NewVRControllerWithName(name string) *VRController {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapConstructorVRControllerWithName(nameToC)
	retvalGO := &VRController{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VRController) {
		C.WrapVRControllerFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *VRController) Free() {
	C.WrapVRControllerFree(pointer.h)
}

// IsNil ...
func (pointer *VRController) IsNil() bool {
	return pointer.h == C.WrapVRController(nil)
}

// IsConnected Gamepad is currently connected.
func (pointer *VRController) IsConnected() bool {
	retval := C.WrapIsConnectedVRController(pointer.h)
	return bool(retval)
}

// Connected Gamepad was connected since the last update.
func (pointer *VRController) Connected() bool {
	retval := C.WrapConnectedVRController(pointer.h)
	return bool(retval)
}

// Disconnected Gamepad was disconnected since the last update.
func (pointer *VRController) Disconnected() bool {
	retval := C.WrapDisconnectedVRController(pointer.h)
	return bool(retval)
}

// World ...
func (pointer *VRController) World() *Mat4 {
	retval := C.WrapWorldVRController(pointer.h)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Down ...
func (pointer *VRController) Down(btn VRControllerButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapDownVRController(pointer.h, btnToC)
	return bool(retval)
}

// Pressed ...
func (pointer *VRController) Pressed(btn VRControllerButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapPressedVRController(pointer.h, btnToC)
	return bool(retval)
}

// Released ...
func (pointer *VRController) Released(btn VRControllerButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapReleasedVRController(pointer.h, btnToC)
	return bool(retval)
}

// Touch ...
func (pointer *VRController) Touch(btn VRControllerButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapTouchVRController(pointer.h, btnToC)
	return bool(retval)
}

// TouchStart ...
func (pointer *VRController) TouchStart(btn VRControllerButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapTouchStartVRController(pointer.h, btnToC)
	return bool(retval)
}

// TouchEnd ...
func (pointer *VRController) TouchEnd(btn VRControllerButton) bool {
	btnToC := C.int32_t(btn)
	retval := C.WrapTouchEndVRController(pointer.h, btnToC)
	return bool(retval)
}

// Surface ...
func (pointer *VRController) Surface(idx int32) *Vec2 {
	idxToC := C.int32_t(idx)
	retval := C.WrapSurfaceVRController(pointer.h, idxToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// DtSurface ...
func (pointer *VRController) DtSurface(idx int32) *Vec2 {
	idxToC := C.int32_t(idx)
	retval := C.WrapDtSurfaceVRController(pointer.h, idxToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// SendHapticPulse ...
func (pointer *VRController) SendHapticPulse(duration int64) {
	durationToC := C.int64_t(duration)
	C.WrapSendHapticPulseVRController(pointer.h, durationToC)
}

// Update ...
func (pointer *VRController) Update() {
	C.WrapUpdateVRController(pointer.h)
}

// VRGenericTrackerState  ...
type VRGenericTrackerState struct {
	h C.WrapVRGenericTrackerState
}

// Free ...
func (pointer *VRGenericTrackerState) Free() {
	C.WrapVRGenericTrackerStateFree(pointer.h)
}

// IsNil ...
func (pointer *VRGenericTrackerState) IsNil() bool {
	return pointer.h == C.WrapVRGenericTrackerState(nil)
}

// IsConnected ...
func (pointer *VRGenericTrackerState) IsConnected() bool {
	retval := C.WrapIsConnectedVRGenericTrackerState(pointer.h)
	return bool(retval)
}

// World ...
func (pointer *VRGenericTrackerState) World() *Mat4 {
	retval := C.WrapWorldVRGenericTrackerState(pointer.h)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// VRGenericTracker  Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetVRGenericTrackerNames] to query for available VR generic tracker devices.
type VRGenericTracker struct {
	h C.WrapVRGenericTracker
}

// NewVRGenericTracker Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetVRGenericTrackerNames] to query for available VR generic tracker devices.
func NewVRGenericTracker() *VRGenericTracker {
	retval := C.WrapConstructorVRGenericTracker()
	retvalGO := &VRGenericTracker{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VRGenericTracker) {
		C.WrapVRGenericTrackerFree(cleanval.h)
	})
	return retvalGO
}

// NewVRGenericTrackerWithName Helper class holding the current and previous device state to enable delta state queries.  Use [harfang.GetVRGenericTrackerNames] to query for available VR generic tracker devices.
func NewVRGenericTrackerWithName(name string) *VRGenericTracker {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapConstructorVRGenericTrackerWithName(nameToC)
	retvalGO := &VRGenericTracker{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VRGenericTracker) {
		C.WrapVRGenericTrackerFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *VRGenericTracker) Free() {
	C.WrapVRGenericTrackerFree(pointer.h)
}

// IsNil ...
func (pointer *VRGenericTracker) IsNil() bool {
	return pointer.h == C.WrapVRGenericTracker(nil)
}

// IsConnected ...
func (pointer *VRGenericTracker) IsConnected() bool {
	retval := C.WrapIsConnectedVRGenericTracker(pointer.h)
	return bool(retval)
}

// World ...
func (pointer *VRGenericTracker) World() *Mat4 {
	retval := C.WrapWorldVRGenericTracker(pointer.h)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Update ...
func (pointer *VRGenericTracker) Update() {
	C.WrapUpdateVRGenericTracker(pointer.h)
}

// DearImguiContext  Context to render immediate GUI.
type DearImguiContext struct {
	h C.WrapDearImguiContext
}

// Free ...
func (pointer *DearImguiContext) Free() {
	C.WrapDearImguiContextFree(pointer.h)
}

// IsNil ...
func (pointer *DearImguiContext) IsNil() bool {
	return pointer.h == C.WrapDearImguiContext(nil)
}

// ImFont  Immediate GUI font.
type ImFont struct {
	h C.WrapImFont
}

// Free ...
func (pointer *ImFont) Free() {
	C.WrapImFontFree(pointer.h)
}

// IsNil ...
func (pointer *ImFont) IsNil() bool {
	return pointer.h == C.WrapImFont(nil)
}

// ImDrawList  Immediate GUI drawing list. This object can be used to perform custom drawing operations on top of an imgui window.
type ImDrawList struct {
	h C.WrapImDrawList
}

// Free ...
func (pointer *ImDrawList) Free() {
	C.WrapImDrawListFree(pointer.h)
}

// IsNil ...
func (pointer *ImDrawList) IsNil() bool {
	return pointer.h == C.WrapImDrawList(nil)
}

// PushClipRect ...
func (pointer *ImDrawList) PushClipRect(cliprectmin *Vec2, cliprectmax *Vec2) {
	cliprectminToC := cliprectmin.h
	cliprectmaxToC := cliprectmax.h
	C.WrapPushClipRectImDrawList(pointer.h, cliprectminToC, cliprectmaxToC)
}

// PushClipRectWithIntersectWithCurentClipRect ...
func (pointer *ImDrawList) PushClipRectWithIntersectWithCurentClipRect(cliprectmin *Vec2, cliprectmax *Vec2, intersectwithcurentcliprect bool) {
	cliprectminToC := cliprectmin.h
	cliprectmaxToC := cliprectmax.h
	intersectwithcurentcliprectToC := C.bool(intersectwithcurentcliprect)
	C.WrapPushClipRectImDrawListWithIntersectWithCurentClipRect(pointer.h, cliprectminToC, cliprectmaxToC, intersectwithcurentcliprectToC)
}

// PushClipRectFullScreen ...
func (pointer *ImDrawList) PushClipRectFullScreen() {
	C.WrapPushClipRectFullScreenImDrawList(pointer.h)
}

// PopClipRect ...
func (pointer *ImDrawList) PopClipRect() {
	C.WrapPopClipRectImDrawList(pointer.h)
}

// PushTextureID ...
func (pointer *ImDrawList) PushTextureID(tex *Texture) {
	texToC := tex.h
	C.WrapPushTextureIDImDrawList(pointer.h, texToC)
}

// PopTextureID ...
func (pointer *ImDrawList) PopTextureID() {
	C.WrapPopTextureIDImDrawList(pointer.h)
}

// GetClipRectMin ...
func (pointer *ImDrawList) GetClipRectMin() *Vec2 {
	retval := C.WrapGetClipRectMinImDrawList(pointer.h)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// GetClipRectMax ...
func (pointer *ImDrawList) GetClipRectMax() *Vec2 {
	retval := C.WrapGetClipRectMaxImDrawList(pointer.h)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// AddLine ...
func (pointer *ImDrawList) AddLine(a *Vec2, b *Vec2, col uint32) {
	aToC := a.h
	bToC := b.h
	colToC := C.uint32_t(col)
	C.WrapAddLineImDrawList(pointer.h, aToC, bToC, colToC)
}

// AddLineWithThickness ...
func (pointer *ImDrawList) AddLineWithThickness(a *Vec2, b *Vec2, col uint32, thickness float32) {
	aToC := a.h
	bToC := b.h
	colToC := C.uint32_t(col)
	thicknessToC := C.float(thickness)
	C.WrapAddLineImDrawListWithThickness(pointer.h, aToC, bToC, colToC, thicknessToC)
}

// AddRect ...
func (pointer *ImDrawList) AddRect(a *Vec2, b *Vec2, col uint32) {
	aToC := a.h
	bToC := b.h
	colToC := C.uint32_t(col)
	C.WrapAddRectImDrawList(pointer.h, aToC, bToC, colToC)
}

// AddRectWithRounding ...
func (pointer *ImDrawList) AddRectWithRounding(a *Vec2, b *Vec2, col uint32, rounding float32) {
	aToC := a.h
	bToC := b.h
	colToC := C.uint32_t(col)
	roundingToC := C.float(rounding)
	C.WrapAddRectImDrawListWithRounding(pointer.h, aToC, bToC, colToC, roundingToC)
}

// AddRectWithRoundingRoundingCornerFlags ...
func (pointer *ImDrawList) AddRectWithRoundingRoundingCornerFlags(a *Vec2, b *Vec2, col uint32, rounding float32, roundingcornerflags int32) {
	aToC := a.h
	bToC := b.h
	colToC := C.uint32_t(col)
	roundingToC := C.float(rounding)
	roundingcornerflagsToC := C.int32_t(roundingcornerflags)
	C.WrapAddRectImDrawListWithRoundingRoundingCornerFlags(pointer.h, aToC, bToC, colToC, roundingToC, roundingcornerflagsToC)
}

// AddRectWithRoundingRoundingCornerFlagsThickness ...
func (pointer *ImDrawList) AddRectWithRoundingRoundingCornerFlagsThickness(a *Vec2, b *Vec2, col uint32, rounding float32, roundingcornerflags int32, thickness float32) {
	aToC := a.h
	bToC := b.h
	colToC := C.uint32_t(col)
	roundingToC := C.float(rounding)
	roundingcornerflagsToC := C.int32_t(roundingcornerflags)
	thicknessToC := C.float(thickness)
	C.WrapAddRectImDrawListWithRoundingRoundingCornerFlagsThickness(pointer.h, aToC, bToC, colToC, roundingToC, roundingcornerflagsToC, thicknessToC)
}

// AddRectFilled ...
func (pointer *ImDrawList) AddRectFilled(a *Vec2, b *Vec2, col uint32) {
	aToC := a.h
	bToC := b.h
	colToC := C.uint32_t(col)
	C.WrapAddRectFilledImDrawList(pointer.h, aToC, bToC, colToC)
}

// AddRectFilledWithRounding ...
func (pointer *ImDrawList) AddRectFilledWithRounding(a *Vec2, b *Vec2, col uint32, rounding float32) {
	aToC := a.h
	bToC := b.h
	colToC := C.uint32_t(col)
	roundingToC := C.float(rounding)
	C.WrapAddRectFilledImDrawListWithRounding(pointer.h, aToC, bToC, colToC, roundingToC)
}

// AddRectFilledWithRoundingRoundingCornerFlags ...
func (pointer *ImDrawList) AddRectFilledWithRoundingRoundingCornerFlags(a *Vec2, b *Vec2, col uint32, rounding float32, roundingcornerflags int32) {
	aToC := a.h
	bToC := b.h
	colToC := C.uint32_t(col)
	roundingToC := C.float(rounding)
	roundingcornerflagsToC := C.int32_t(roundingcornerflags)
	C.WrapAddRectFilledImDrawListWithRoundingRoundingCornerFlags(pointer.h, aToC, bToC, colToC, roundingToC, roundingcornerflagsToC)
}

// AddRectFilledMultiColor ...
func (pointer *ImDrawList) AddRectFilledMultiColor(a *Vec2, b *Vec2, coluprleft uint32, coluprright uint32, colbotright uint32, colbotleft uint32) {
	aToC := a.h
	bToC := b.h
	coluprleftToC := C.uint32_t(coluprleft)
	coluprrightToC := C.uint32_t(coluprright)
	colbotrightToC := C.uint32_t(colbotright)
	colbotleftToC := C.uint32_t(colbotleft)
	C.WrapAddRectFilledMultiColorImDrawList(pointer.h, aToC, bToC, coluprleftToC, coluprrightToC, colbotrightToC, colbotleftToC)
}

// AddQuad ...
func (pointer *ImDrawList) AddQuad(a *Vec2, b *Vec2, c *Vec2, d *Vec2, col uint32) {
	aToC := a.h
	bToC := b.h
	cToC := c.h
	dToC := d.h
	colToC := C.uint32_t(col)
	C.WrapAddQuadImDrawList(pointer.h, aToC, bToC, cToC, dToC, colToC)
}

// AddQuadWithThickness ...
func (pointer *ImDrawList) AddQuadWithThickness(a *Vec2, b *Vec2, c *Vec2, d *Vec2, col uint32, thickness float32) {
	aToC := a.h
	bToC := b.h
	cToC := c.h
	dToC := d.h
	colToC := C.uint32_t(col)
	thicknessToC := C.float(thickness)
	C.WrapAddQuadImDrawListWithThickness(pointer.h, aToC, bToC, cToC, dToC, colToC, thicknessToC)
}

// AddQuadFilled ...
func (pointer *ImDrawList) AddQuadFilled(a *Vec2, b *Vec2, c *Vec2, d *Vec2, col uint32) {
	aToC := a.h
	bToC := b.h
	cToC := c.h
	dToC := d.h
	colToC := C.uint32_t(col)
	C.WrapAddQuadFilledImDrawList(pointer.h, aToC, bToC, cToC, dToC, colToC)
}

// AddTriangle ...
func (pointer *ImDrawList) AddTriangle(a *Vec2, b *Vec2, c *Vec2, col uint32) {
	aToC := a.h
	bToC := b.h
	cToC := c.h
	colToC := C.uint32_t(col)
	C.WrapAddTriangleImDrawList(pointer.h, aToC, bToC, cToC, colToC)
}

// AddTriangleWithThickness ...
func (pointer *ImDrawList) AddTriangleWithThickness(a *Vec2, b *Vec2, c *Vec2, col uint32, thickness float32) {
	aToC := a.h
	bToC := b.h
	cToC := c.h
	colToC := C.uint32_t(col)
	thicknessToC := C.float(thickness)
	C.WrapAddTriangleImDrawListWithThickness(pointer.h, aToC, bToC, cToC, colToC, thicknessToC)
}

// AddTriangleFilled ...
func (pointer *ImDrawList) AddTriangleFilled(a *Vec2, b *Vec2, c *Vec2, col uint32) {
	aToC := a.h
	bToC := b.h
	cToC := c.h
	colToC := C.uint32_t(col)
	C.WrapAddTriangleFilledImDrawList(pointer.h, aToC, bToC, cToC, colToC)
}

// AddCircle ...
func (pointer *ImDrawList) AddCircle(centre *Vec2, radius float32, col uint32) {
	centreToC := centre.h
	radiusToC := C.float(radius)
	colToC := C.uint32_t(col)
	C.WrapAddCircleImDrawList(pointer.h, centreToC, radiusToC, colToC)
}

// AddCircleWithNumSegments ...
func (pointer *ImDrawList) AddCircleWithNumSegments(centre *Vec2, radius float32, col uint32, numsegments int32) {
	centreToC := centre.h
	radiusToC := C.float(radius)
	colToC := C.uint32_t(col)
	numsegmentsToC := C.int32_t(numsegments)
	C.WrapAddCircleImDrawListWithNumSegments(pointer.h, centreToC, radiusToC, colToC, numsegmentsToC)
}

// AddCircleWithNumSegmentsThickness ...
func (pointer *ImDrawList) AddCircleWithNumSegmentsThickness(centre *Vec2, radius float32, col uint32, numsegments int32, thickness float32) {
	centreToC := centre.h
	radiusToC := C.float(radius)
	colToC := C.uint32_t(col)
	numsegmentsToC := C.int32_t(numsegments)
	thicknessToC := C.float(thickness)
	C.WrapAddCircleImDrawListWithNumSegmentsThickness(pointer.h, centreToC, radiusToC, colToC, numsegmentsToC, thicknessToC)
}

// AddCircleFilled ...
func (pointer *ImDrawList) AddCircleFilled(centre *Vec2, radius float32, col uint32) {
	centreToC := centre.h
	radiusToC := C.float(radius)
	colToC := C.uint32_t(col)
	C.WrapAddCircleFilledImDrawList(pointer.h, centreToC, radiusToC, colToC)
}

// AddCircleFilledWithNumSegments ...
func (pointer *ImDrawList) AddCircleFilledWithNumSegments(centre *Vec2, radius float32, col uint32, numsegments int32) {
	centreToC := centre.h
	radiusToC := C.float(radius)
	colToC := C.uint32_t(col)
	numsegmentsToC := C.int32_t(numsegments)
	C.WrapAddCircleFilledImDrawListWithNumSegments(pointer.h, centreToC, radiusToC, colToC, numsegmentsToC)
}

// AddText ...
func (pointer *ImDrawList) AddText(pos *Vec2, col uint32, text string) {
	posToC := pos.h
	colToC := C.uint32_t(col)
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	C.WrapAddTextImDrawList(pointer.h, posToC, colToC, textToC)
}

// AddTextWithFontFontSizePosColText ...
func (pointer *ImDrawList) AddTextWithFontFontSizePosColText(font *ImFont, fontsize float32, pos *Vec2, col uint32, text string) {
	fontToC := font.h
	fontsizeToC := C.float(fontsize)
	posToC := pos.h
	colToC := C.uint32_t(col)
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	C.WrapAddTextImDrawListWithFontFontSizePosColText(pointer.h, fontToC, fontsizeToC, posToC, colToC, textToC)
}

// AddTextWithFontFontSizePosColTextWrapWidth ...
func (pointer *ImDrawList) AddTextWithFontFontSizePosColTextWrapWidth(font *ImFont, fontsize float32, pos *Vec2, col uint32, text string, wrapwidth float32) {
	fontToC := font.h
	fontsizeToC := C.float(fontsize)
	posToC := pos.h
	colToC := C.uint32_t(col)
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	wrapwidthToC := C.float(wrapwidth)
	C.WrapAddTextImDrawListWithFontFontSizePosColTextWrapWidth(pointer.h, fontToC, fontsizeToC, posToC, colToC, textToC, wrapwidthToC)
}

// AddTextWithFontFontSizePosColTextWrapWidthCpuFineClipRect ...
func (pointer *ImDrawList) AddTextWithFontFontSizePosColTextWrapWidthCpuFineClipRect(font *ImFont, fontsize float32, pos *Vec2, col uint32, text string, wrapwidth float32, cpufinecliprect *Vec4) {
	fontToC := font.h
	fontsizeToC := C.float(fontsize)
	posToC := pos.h
	colToC := C.uint32_t(col)
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	wrapwidthToC := C.float(wrapwidth)
	cpufinecliprectToC := cpufinecliprect.h
	C.WrapAddTextImDrawListWithFontFontSizePosColTextWrapWidthCpuFineClipRect(pointer.h, fontToC, fontsizeToC, posToC, colToC, textToC, wrapwidthToC, cpufinecliprectToC)
}

// AddImage ...
func (pointer *ImDrawList) AddImage(tex *Texture, a *Vec2, b *Vec2) {
	texToC := tex.h
	aToC := a.h
	bToC := b.h
	C.WrapAddImageImDrawList(pointer.h, texToC, aToC, bToC)
}

// AddImageWithUvAUvB ...
func (pointer *ImDrawList) AddImageWithUvAUvB(tex *Texture, a *Vec2, b *Vec2, uva *Vec2, uvb *Vec2) {
	texToC := tex.h
	aToC := a.h
	bToC := b.h
	uvaToC := uva.h
	uvbToC := uvb.h
	C.WrapAddImageImDrawListWithUvAUvB(pointer.h, texToC, aToC, bToC, uvaToC, uvbToC)
}

// AddImageWithUvAUvBCol ...
func (pointer *ImDrawList) AddImageWithUvAUvBCol(tex *Texture, a *Vec2, b *Vec2, uva *Vec2, uvb *Vec2, col uint32) {
	texToC := tex.h
	aToC := a.h
	bToC := b.h
	uvaToC := uva.h
	uvbToC := uvb.h
	colToC := C.uint32_t(col)
	C.WrapAddImageImDrawListWithUvAUvBCol(pointer.h, texToC, aToC, bToC, uvaToC, uvbToC, colToC)
}

// AddImageQuad ...
func (pointer *ImDrawList) AddImageQuad(tex *Texture, a *Vec2, b *Vec2, c *Vec2, d *Vec2) {
	texToC := tex.h
	aToC := a.h
	bToC := b.h
	cToC := c.h
	dToC := d.h
	C.WrapAddImageQuadImDrawList(pointer.h, texToC, aToC, bToC, cToC, dToC)
}

// AddImageQuadWithUvAUvBUvCUvD ...
func (pointer *ImDrawList) AddImageQuadWithUvAUvBUvCUvD(tex *Texture, a *Vec2, b *Vec2, c *Vec2, d *Vec2, uva *Vec2, uvb *Vec2, uvc *Vec2, uvd *Vec2) {
	texToC := tex.h
	aToC := a.h
	bToC := b.h
	cToC := c.h
	dToC := d.h
	uvaToC := uva.h
	uvbToC := uvb.h
	uvcToC := uvc.h
	uvdToC := uvd.h
	C.WrapAddImageQuadImDrawListWithUvAUvBUvCUvD(pointer.h, texToC, aToC, bToC, cToC, dToC, uvaToC, uvbToC, uvcToC, uvdToC)
}

// AddImageQuadWithUvAUvBUvCUvDCol ...
func (pointer *ImDrawList) AddImageQuadWithUvAUvBUvCUvDCol(tex *Texture, a *Vec2, b *Vec2, c *Vec2, d *Vec2, uva *Vec2, uvb *Vec2, uvc *Vec2, uvd *Vec2, col uint32) {
	texToC := tex.h
	aToC := a.h
	bToC := b.h
	cToC := c.h
	dToC := d.h
	uvaToC := uva.h
	uvbToC := uvb.h
	uvcToC := uvc.h
	uvdToC := uvd.h
	colToC := C.uint32_t(col)
	C.WrapAddImageQuadImDrawListWithUvAUvBUvCUvDCol(pointer.h, texToC, aToC, bToC, cToC, dToC, uvaToC, uvbToC, uvcToC, uvdToC, colToC)
}

// AddImageRounded ...
func (pointer *ImDrawList) AddImageRounded(tex *Texture, a *Vec2, b *Vec2, uva *Vec2, uvb *Vec2, col uint32, rounding float32) {
	texToC := tex.h
	aToC := a.h
	bToC := b.h
	uvaToC := uva.h
	uvbToC := uvb.h
	colToC := C.uint32_t(col)
	roundingToC := C.float(rounding)
	C.WrapAddImageRoundedImDrawList(pointer.h, texToC, aToC, bToC, uvaToC, uvbToC, colToC, roundingToC)
}

// AddImageRoundedWithFlags ...
func (pointer *ImDrawList) AddImageRoundedWithFlags(tex *Texture, a *Vec2, b *Vec2, uva *Vec2, uvb *Vec2, col uint32, rounding float32, flags ImDrawFlags) {
	texToC := tex.h
	aToC := a.h
	bToC := b.h
	uvaToC := uva.h
	uvbToC := uvb.h
	colToC := C.uint32_t(col)
	roundingToC := C.float(rounding)
	flagsToC := C.int32_t(flags)
	C.WrapAddImageRoundedImDrawListWithFlags(pointer.h, texToC, aToC, bToC, uvaToC, uvbToC, colToC, roundingToC, flagsToC)
}

// AddPolyline ...
func (pointer *ImDrawList) AddPolyline(points *Vec2List, col uint32, closed bool, thickness float32) {
	pointsToC := points.h
	colToC := C.uint32_t(col)
	closedToC := C.bool(closed)
	thicknessToC := C.float(thickness)
	C.WrapAddPolylineImDrawList(pointer.h, pointsToC, colToC, closedToC, thicknessToC)
}

// AddConvexPolyFilled ...
func (pointer *ImDrawList) AddConvexPolyFilled(points *Vec2List, col uint32) {
	pointsToC := points.h
	colToC := C.uint32_t(col)
	C.WrapAddConvexPolyFilledImDrawList(pointer.h, pointsToC, colToC)
}

// AddBezierCubic ...
func (pointer *ImDrawList) AddBezierCubic(pos0 *Vec2, cp0 *Vec2, cp1 *Vec2, pos1 *Vec2, col uint32, thickness float32) {
	pos0ToC := pos0.h
	cp0ToC := cp0.h
	cp1ToC := cp1.h
	pos1ToC := pos1.h
	colToC := C.uint32_t(col)
	thicknessToC := C.float(thickness)
	C.WrapAddBezierCubicImDrawList(pointer.h, pos0ToC, cp0ToC, cp1ToC, pos1ToC, colToC, thicknessToC)
}

// AddBezierCubicWithNumSegments ...
func (pointer *ImDrawList) AddBezierCubicWithNumSegments(pos0 *Vec2, cp0 *Vec2, cp1 *Vec2, pos1 *Vec2, col uint32, thickness float32, numsegments int32) {
	pos0ToC := pos0.h
	cp0ToC := cp0.h
	cp1ToC := cp1.h
	pos1ToC := pos1.h
	colToC := C.uint32_t(col)
	thicknessToC := C.float(thickness)
	numsegmentsToC := C.int32_t(numsegments)
	C.WrapAddBezierCubicImDrawListWithNumSegments(pointer.h, pos0ToC, cp0ToC, cp1ToC, pos1ToC, colToC, thicknessToC, numsegmentsToC)
}

// PathClear ...
func (pointer *ImDrawList) PathClear() {
	C.WrapPathClearImDrawList(pointer.h)
}

// PathLineTo ...
func (pointer *ImDrawList) PathLineTo(pos *Vec2) {
	posToC := pos.h
	C.WrapPathLineToImDrawList(pointer.h, posToC)
}

// PathLineToMergeDuplicate ...
func (pointer *ImDrawList) PathLineToMergeDuplicate(pos *Vec2) {
	posToC := pos.h
	C.WrapPathLineToMergeDuplicateImDrawList(pointer.h, posToC)
}

// PathFillConvex ...
func (pointer *ImDrawList) PathFillConvex(col uint32) {
	colToC := C.uint32_t(col)
	C.WrapPathFillConvexImDrawList(pointer.h, colToC)
}

// PathStroke ...
func (pointer *ImDrawList) PathStroke(col uint32, closed bool) {
	colToC := C.uint32_t(col)
	closedToC := C.bool(closed)
	C.WrapPathStrokeImDrawList(pointer.h, colToC, closedToC)
}

// PathStrokeWithThickness ...
func (pointer *ImDrawList) PathStrokeWithThickness(col uint32, closed bool, thickness float32) {
	colToC := C.uint32_t(col)
	closedToC := C.bool(closed)
	thicknessToC := C.float(thickness)
	C.WrapPathStrokeImDrawListWithThickness(pointer.h, colToC, closedToC, thicknessToC)
}

// PathArcTo ...
func (pointer *ImDrawList) PathArcTo(centre *Vec2, radius float32, amin float32, amax float32) {
	centreToC := centre.h
	radiusToC := C.float(radius)
	aminToC := C.float(amin)
	amaxToC := C.float(amax)
	C.WrapPathArcToImDrawList(pointer.h, centreToC, radiusToC, aminToC, amaxToC)
}

// PathArcToWithNumSegments ...
func (pointer *ImDrawList) PathArcToWithNumSegments(centre *Vec2, radius float32, amin float32, amax float32, numsegments int32) {
	centreToC := centre.h
	radiusToC := C.float(radius)
	aminToC := C.float(amin)
	amaxToC := C.float(amax)
	numsegmentsToC := C.int32_t(numsegments)
	C.WrapPathArcToImDrawListWithNumSegments(pointer.h, centreToC, radiusToC, aminToC, amaxToC, numsegmentsToC)
}

// PathArcToFast ...
func (pointer *ImDrawList) PathArcToFast(centre *Vec2, radius float32, aminof12 int32, amaxof12 int32) {
	centreToC := centre.h
	radiusToC := C.float(radius)
	aminof12ToC := C.int32_t(aminof12)
	amaxof12ToC := C.int32_t(amaxof12)
	C.WrapPathArcToFastImDrawList(pointer.h, centreToC, radiusToC, aminof12ToC, amaxof12ToC)
}

// PathBezierCubicCurveTo ...
func (pointer *ImDrawList) PathBezierCubicCurveTo(p1 *Vec2, p2 *Vec2, p3 *Vec2) {
	p1ToC := p1.h
	p2ToC := p2.h
	p3ToC := p3.h
	C.WrapPathBezierCubicCurveToImDrawList(pointer.h, p1ToC, p2ToC, p3ToC)
}

// PathBezierCubicCurveToWithNumSegments ...
func (pointer *ImDrawList) PathBezierCubicCurveToWithNumSegments(p1 *Vec2, p2 *Vec2, p3 *Vec2, numsegments int32) {
	p1ToC := p1.h
	p2ToC := p2.h
	p3ToC := p3.h
	numsegmentsToC := C.int32_t(numsegments)
	C.WrapPathBezierCubicCurveToImDrawListWithNumSegments(pointer.h, p1ToC, p2ToC, p3ToC, numsegmentsToC)
}

// PathRect ...
func (pointer *ImDrawList) PathRect(rectmin *Vec2, rectmax *Vec2) {
	rectminToC := rectmin.h
	rectmaxToC := rectmax.h
	C.WrapPathRectImDrawList(pointer.h, rectminToC, rectmaxToC)
}

// PathRectWithRounding ...
func (pointer *ImDrawList) PathRectWithRounding(rectmin *Vec2, rectmax *Vec2, rounding float32) {
	rectminToC := rectmin.h
	rectmaxToC := rectmax.h
	roundingToC := C.float(rounding)
	C.WrapPathRectImDrawListWithRounding(pointer.h, rectminToC, rectmaxToC, roundingToC)
}

// PathRectWithRoundingFlags ...
func (pointer *ImDrawList) PathRectWithRoundingFlags(rectmin *Vec2, rectmax *Vec2, rounding float32, flags ImDrawFlags) {
	rectminToC := rectmin.h
	rectmaxToC := rectmax.h
	roundingToC := C.float(rounding)
	flagsToC := C.int32_t(flags)
	C.WrapPathRectImDrawListWithRoundingFlags(pointer.h, rectminToC, rectmaxToC, roundingToC, flagsToC)
}

// ChannelsSplit ...
func (pointer *ImDrawList) ChannelsSplit(channelscount int32) {
	channelscountToC := C.int32_t(channelscount)
	C.WrapChannelsSplitImDrawList(pointer.h, channelscountToC)
}

// ChannelsMerge ...
func (pointer *ImDrawList) ChannelsMerge() {
	C.WrapChannelsMergeImDrawList(pointer.h)
}

// ChannelsSetCurrent ...
func (pointer *ImDrawList) ChannelsSetCurrent(channelindex int32) {
	channelindexToC := C.int32_t(channelindex)
	C.WrapChannelsSetCurrentImDrawList(pointer.h, channelindexToC)
}

// FileFilter  ...
type FileFilter struct {
	h C.WrapFileFilter
}

// GetName ...
func (pointer *FileFilter) GetName() string {
	v := C.WrapFileFilterGetName(pointer.h)
	return C.GoString(v)
}

// SetName ...
func (pointer *FileFilter) SetName(v string) {
	vToC, idFinvToC := wrapString(v)
	defer idFinvToC()
	C.WrapFileFilterSetName(pointer.h, vToC)
}

// GetPattern ...
func (pointer *FileFilter) GetPattern() string {
	v := C.WrapFileFilterGetPattern(pointer.h)
	return C.GoString(v)
}

// SetPattern ...
func (pointer *FileFilter) SetPattern(v string) {
	vToC, idFinvToC := wrapString(v)
	defer idFinvToC()
	C.WrapFileFilterSetPattern(pointer.h, vToC)
}

// Free ...
func (pointer *FileFilter) Free() {
	C.WrapFileFilterFree(pointer.h)
}

// IsNil ...
func (pointer *FileFilter) IsNil() bool {
	return pointer.h == C.WrapFileFilter(nil)
}

// GoSliceOfFileFilter ...
type GoSliceOfFileFilter []*FileFilter

// FileFilterList  ...
type FileFilterList struct {
	h C.WrapFileFilterList
}

// Get ...
func (pointer *FileFilterList) Get(id int) *FileFilter {
	v := C.WrapFileFilterListGetOperator(pointer.h, C.int(id))
	vGO := &FileFilter{h: v}
	runtime.SetFinalizer(vGO, func(cleanval *FileFilter) {
		C.WrapFileFilterFree(cleanval.h)
	})
	return vGO
}

// Set ...
func (pointer *FileFilterList) Set(id int, v *FileFilter) {
	vToC := v.h
	C.WrapFileFilterListSetOperator(pointer.h, C.int(id), vToC)
}

// Len ...
func (pointer *FileFilterList) Len() int32 {
	return int32(C.WrapFileFilterListLenOperator(pointer.h))
}

// NewFileFilterList ...
func NewFileFilterList() *FileFilterList {
	retval := C.WrapConstructorFileFilterList()
	retvalGO := &FileFilterList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *FileFilterList) {
		C.WrapFileFilterListFree(cleanval.h)
	})
	return retvalGO
}

// NewFileFilterListWithSequence ...
func NewFileFilterListWithSequence(sequence GoSliceOfFileFilter) *FileFilterList {
	var sequencePointer []C.WrapFileFilter
	for _, s := range sequence {
		sequencePointer = append(sequencePointer, s.h)
	}
	sequencePointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&sequencePointer))
	sequencePointerToCSize := C.size_t(sequencePointerToC.Len)
	sequencePointerToCBuf := (*C.WrapFileFilter)(unsafe.Pointer(sequencePointerToC.Data))
	retval := C.WrapConstructorFileFilterListWithSequence(sequencePointerToCSize, sequencePointerToCBuf)
	retvalGO := &FileFilterList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *FileFilterList) {
		C.WrapFileFilterListFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *FileFilterList) Free() {
	C.WrapFileFilterListFree(pointer.h)
}

// IsNil ...
func (pointer *FileFilterList) IsNil() bool {
	return pointer.h == C.WrapFileFilterList(nil)
}

// Clear ...
func (pointer *FileFilterList) Clear() {
	C.WrapClearFileFilterList(pointer.h)
}

// Reserve ...
func (pointer *FileFilterList) Reserve(size int32) {
	sizeToC := C.size_t(size)
	C.WrapReserveFileFilterList(pointer.h, sizeToC)
}

// PushBack ...
func (pointer *FileFilterList) PushBack(v *FileFilter) {
	vToC := v.h
	C.WrapPushBackFileFilterList(pointer.h, vToC)
}

// Size ...
func (pointer *FileFilterList) Size() int32 {
	retval := C.WrapSizeFileFilterList(pointer.h)
	return int32(retval)
}

// At ...
func (pointer *FileFilterList) At(idx int32) *FileFilter {
	idxToC := C.size_t(idx)
	retval := C.WrapAtFileFilterList(pointer.h, idxToC)
	retvalGO := &FileFilter{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *FileFilter) {
		C.WrapFileFilterFree(cleanval.h)
	})
	return retvalGO
}

// StereoSourceState  State for a stereo audio source, see [harfang.man.Audio].
type StereoSourceState struct {
	h C.WrapStereoSourceState
}

// GetVolume ...
func (pointer *StereoSourceState) GetVolume() float32 {
	v := C.WrapStereoSourceStateGetVolume(pointer.h)
	return float32(v)
}

// SetVolume ...
func (pointer *StereoSourceState) SetVolume(v float32) {
	vToC := C.float(v)
	C.WrapStereoSourceStateSetVolume(pointer.h, vToC)
}

// GetRepeat ...
func (pointer *StereoSourceState) GetRepeat() SourceRepeat {
	v := C.WrapStereoSourceStateGetRepeat(pointer.h)
	return SourceRepeat(v)
}

// SetRepeat ...
func (pointer *StereoSourceState) SetRepeat(v SourceRepeat) {
	vToC := C.int32_t(v)
	C.WrapStereoSourceStateSetRepeat(pointer.h, vToC)
}

// GetPanning ...
func (pointer *StereoSourceState) GetPanning() float32 {
	v := C.WrapStereoSourceStateGetPanning(pointer.h)
	return float32(v)
}

// SetPanning ...
func (pointer *StereoSourceState) SetPanning(v float32) {
	vToC := C.float(v)
	C.WrapStereoSourceStateSetPanning(pointer.h, vToC)
}

// NewStereoSourceState State for a stereo audio source, see [harfang.man.Audio].
func NewStereoSourceState() *StereoSourceState {
	retval := C.WrapConstructorStereoSourceState()
	retvalGO := &StereoSourceState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StereoSourceState) {
		C.WrapStereoSourceStateFree(cleanval.h)
	})
	return retvalGO
}

// NewStereoSourceStateWithVolume State for a stereo audio source, see [harfang.man.Audio].
func NewStereoSourceStateWithVolume(volume float32) *StereoSourceState {
	volumeToC := C.float(volume)
	retval := C.WrapConstructorStereoSourceStateWithVolume(volumeToC)
	retvalGO := &StereoSourceState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StereoSourceState) {
		C.WrapStereoSourceStateFree(cleanval.h)
	})
	return retvalGO
}

// NewStereoSourceStateWithVolumeRepeat State for a stereo audio source, see [harfang.man.Audio].
func NewStereoSourceStateWithVolumeRepeat(volume float32, repeat SourceRepeat) *StereoSourceState {
	volumeToC := C.float(volume)
	repeatToC := C.int32_t(repeat)
	retval := C.WrapConstructorStereoSourceStateWithVolumeRepeat(volumeToC, repeatToC)
	retvalGO := &StereoSourceState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StereoSourceState) {
		C.WrapStereoSourceStateFree(cleanval.h)
	})
	return retvalGO
}

// NewStereoSourceStateWithVolumeRepeatPanning State for a stereo audio source, see [harfang.man.Audio].
func NewStereoSourceStateWithVolumeRepeatPanning(volume float32, repeat SourceRepeat, panning float32) *StereoSourceState {
	volumeToC := C.float(volume)
	repeatToC := C.int32_t(repeat)
	panningToC := C.float(panning)
	retval := C.WrapConstructorStereoSourceStateWithVolumeRepeatPanning(volumeToC, repeatToC, panningToC)
	retvalGO := &StereoSourceState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StereoSourceState) {
		C.WrapStereoSourceStateFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *StereoSourceState) Free() {
	C.WrapStereoSourceStateFree(pointer.h)
}

// IsNil ...
func (pointer *StereoSourceState) IsNil() bool {
	return pointer.h == C.WrapStereoSourceState(nil)
}

// SpatializedSourceState  State for a spatialized audio source, see [harfang.man.Audio].
type SpatializedSourceState struct {
	h C.WrapSpatializedSourceState
}

// GetMtx ...
func (pointer *SpatializedSourceState) GetMtx() *Mat4 {
	v := C.WrapSpatializedSourceStateGetMtx(pointer.h)
	vGO := &Mat4{h: v}
	return vGO
}

// SetMtx ...
func (pointer *SpatializedSourceState) SetMtx(v *Mat4) {
	vToC := v.h
	C.WrapSpatializedSourceStateSetMtx(pointer.h, vToC)
}

// GetVolume ...
func (pointer *SpatializedSourceState) GetVolume() float32 {
	v := C.WrapSpatializedSourceStateGetVolume(pointer.h)
	return float32(v)
}

// SetVolume ...
func (pointer *SpatializedSourceState) SetVolume(v float32) {
	vToC := C.float(v)
	C.WrapSpatializedSourceStateSetVolume(pointer.h, vToC)
}

// GetRepeat ...
func (pointer *SpatializedSourceState) GetRepeat() SourceRepeat {
	v := C.WrapSpatializedSourceStateGetRepeat(pointer.h)
	return SourceRepeat(v)
}

// SetRepeat ...
func (pointer *SpatializedSourceState) SetRepeat(v SourceRepeat) {
	vToC := C.int32_t(v)
	C.WrapSpatializedSourceStateSetRepeat(pointer.h, vToC)
}

// GetVel ...
func (pointer *SpatializedSourceState) GetVel() *Vec3 {
	v := C.WrapSpatializedSourceStateGetVel(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetVel ...
func (pointer *SpatializedSourceState) SetVel(v *Vec3) {
	vToC := v.h
	C.WrapSpatializedSourceStateSetVel(pointer.h, vToC)
}

// NewSpatializedSourceState State for a spatialized audio source, see [harfang.man.Audio].
func NewSpatializedSourceState() *SpatializedSourceState {
	retval := C.WrapConstructorSpatializedSourceState()
	retvalGO := &SpatializedSourceState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SpatializedSourceState) {
		C.WrapSpatializedSourceStateFree(cleanval.h)
	})
	return retvalGO
}

// NewSpatializedSourceStateWithMtx State for a spatialized audio source, see [harfang.man.Audio].
func NewSpatializedSourceStateWithMtx(mtx *Mat4) *SpatializedSourceState {
	mtxToC := mtx.h
	retval := C.WrapConstructorSpatializedSourceStateWithMtx(mtxToC)
	retvalGO := &SpatializedSourceState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SpatializedSourceState) {
		C.WrapSpatializedSourceStateFree(cleanval.h)
	})
	return retvalGO
}

// NewSpatializedSourceStateWithMtxVolume State for a spatialized audio source, see [harfang.man.Audio].
func NewSpatializedSourceStateWithMtxVolume(mtx *Mat4, volume float32) *SpatializedSourceState {
	mtxToC := mtx.h
	volumeToC := C.float(volume)
	retval := C.WrapConstructorSpatializedSourceStateWithMtxVolume(mtxToC, volumeToC)
	retvalGO := &SpatializedSourceState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SpatializedSourceState) {
		C.WrapSpatializedSourceStateFree(cleanval.h)
	})
	return retvalGO
}

// NewSpatializedSourceStateWithMtxVolumeRepeat State for a spatialized audio source, see [harfang.man.Audio].
func NewSpatializedSourceStateWithMtxVolumeRepeat(mtx *Mat4, volume float32, repeat SourceRepeat) *SpatializedSourceState {
	mtxToC := mtx.h
	volumeToC := C.float(volume)
	repeatToC := C.int32_t(repeat)
	retval := C.WrapConstructorSpatializedSourceStateWithMtxVolumeRepeat(mtxToC, volumeToC, repeatToC)
	retvalGO := &SpatializedSourceState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SpatializedSourceState) {
		C.WrapSpatializedSourceStateFree(cleanval.h)
	})
	return retvalGO
}

// NewSpatializedSourceStateWithMtxVolumeRepeatVel State for a spatialized audio source, see [harfang.man.Audio].
func NewSpatializedSourceStateWithMtxVolumeRepeatVel(mtx *Mat4, volume float32, repeat SourceRepeat, vel *Vec3) *SpatializedSourceState {
	mtxToC := mtx.h
	volumeToC := C.float(volume)
	repeatToC := C.int32_t(repeat)
	velToC := vel.h
	retval := C.WrapConstructorSpatializedSourceStateWithMtxVolumeRepeatVel(mtxToC, volumeToC, repeatToC, velToC)
	retvalGO := &SpatializedSourceState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SpatializedSourceState) {
		C.WrapSpatializedSourceStateFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *SpatializedSourceState) Free() {
	C.WrapSpatializedSourceStateFree(pointer.h)
}

// IsNil ...
func (pointer *SpatializedSourceState) IsNil() bool {
	return pointer.h == C.WrapSpatializedSourceState(nil)
}

// OpenVREye  Matrices for a VR eye, see [harfang.OpenVRState].
type OpenVREye struct {
	h C.WrapOpenVREye
}

// GetOffset ...
func (pointer *OpenVREye) GetOffset() *Mat4 {
	v := C.WrapOpenVREyeGetOffset(pointer.h)
	vGO := &Mat4{h: v}
	return vGO
}

// SetOffset ...
func (pointer *OpenVREye) SetOffset(v *Mat4) {
	vToC := v.h
	C.WrapOpenVREyeSetOffset(pointer.h, vToC)
}

// GetProjection ...
func (pointer *OpenVREye) GetProjection() *Mat44 {
	v := C.WrapOpenVREyeGetProjection(pointer.h)
	vGO := &Mat44{h: v}
	return vGO
}

// SetProjection ...
func (pointer *OpenVREye) SetProjection(v *Mat44) {
	vToC := v.h
	C.WrapOpenVREyeSetProjection(pointer.h, vToC)
}

// Free ...
func (pointer *OpenVREye) Free() {
	C.WrapOpenVREyeFree(pointer.h)
}

// IsNil ...
func (pointer *OpenVREye) IsNil() bool {
	return pointer.h == C.WrapOpenVREye(nil)
}

// OpenVREyeFrameBuffer  Framebuffer for a VR eye. Render to two such buffer, one for each eye, before submitting them using [harfang.OpenVRSubmitFrame].
type OpenVREyeFrameBuffer struct {
	h C.WrapOpenVREyeFrameBuffer
}

// Free ...
func (pointer *OpenVREyeFrameBuffer) Free() {
	C.WrapOpenVREyeFrameBufferFree(pointer.h)
}

// IsNil ...
func (pointer *OpenVREyeFrameBuffer) IsNil() bool {
	return pointer.h == C.WrapOpenVREyeFrameBuffer(nil)
}

// GetHandle ...
func (pointer *OpenVREyeFrameBuffer) GetHandle() *FrameBufferHandle {
	retval := C.WrapGetHandleOpenVREyeFrameBuffer(pointer.h)
	retvalGO := &FrameBufferHandle{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *FrameBufferHandle) {
		C.WrapFrameBufferHandleFree(cleanval.h)
	})
	return retvalGO
}

// OpenVRState  OpenVR state including the body and head transformations, the left and right eye states and the render target dimensions expected by the backend.
type OpenVRState struct {
	h C.WrapOpenVRState
}

// GetBody ...
func (pointer *OpenVRState) GetBody() *Mat4 {
	v := C.WrapOpenVRStateGetBody(pointer.h)
	vGO := &Mat4{h: v}
	return vGO
}

// SetBody ...
func (pointer *OpenVRState) SetBody(v *Mat4) {
	vToC := v.h
	C.WrapOpenVRStateSetBody(pointer.h, vToC)
}

// GetHead ...
func (pointer *OpenVRState) GetHead() *Mat4 {
	v := C.WrapOpenVRStateGetHead(pointer.h)
	vGO := &Mat4{h: v}
	return vGO
}

// SetHead ...
func (pointer *OpenVRState) SetHead(v *Mat4) {
	vToC := v.h
	C.WrapOpenVRStateSetHead(pointer.h, vToC)
}

// GetInvHead ...
func (pointer *OpenVRState) GetInvHead() *Mat4 {
	v := C.WrapOpenVRStateGetInvHead(pointer.h)
	vGO := &Mat4{h: v}
	return vGO
}

// SetInvHead ...
func (pointer *OpenVRState) SetInvHead(v *Mat4) {
	vToC := v.h
	C.WrapOpenVRStateSetInvHead(pointer.h, vToC)
}

// GetLeft ...
func (pointer *OpenVRState) GetLeft() *OpenVREye {
	v := C.WrapOpenVRStateGetLeft(pointer.h)
	vGO := &OpenVREye{h: v}
	return vGO
}

// SetLeft ...
func (pointer *OpenVRState) SetLeft(v *OpenVREye) {
	vToC := v.h
	C.WrapOpenVRStateSetLeft(pointer.h, vToC)
}

// GetRight ...
func (pointer *OpenVRState) GetRight() *OpenVREye {
	v := C.WrapOpenVRStateGetRight(pointer.h)
	vGO := &OpenVREye{h: v}
	return vGO
}

// SetRight ...
func (pointer *OpenVRState) SetRight(v *OpenVREye) {
	vToC := v.h
	C.WrapOpenVRStateSetRight(pointer.h, vToC)
}

// GetWidth ...
func (pointer *OpenVRState) GetWidth() uint32 {
	v := C.WrapOpenVRStateGetWidth(pointer.h)
	return uint32(v)
}

// SetWidth ...
func (pointer *OpenVRState) SetWidth(v uint32) {
	vToC := C.uint32_t(v)
	C.WrapOpenVRStateSetWidth(pointer.h, vToC)
}

// GetHeight ...
func (pointer *OpenVRState) GetHeight() uint32 {
	v := C.WrapOpenVRStateGetHeight(pointer.h)
	return uint32(v)
}

// SetHeight ...
func (pointer *OpenVRState) SetHeight(v uint32) {
	vToC := C.uint32_t(v)
	C.WrapOpenVRStateSetHeight(pointer.h, vToC)
}

// Free ...
func (pointer *OpenVRState) Free() {
	C.WrapOpenVRStateFree(pointer.h)
}

// IsNil ...
func (pointer *OpenVRState) IsNil() bool {
	return pointer.h == C.WrapOpenVRState(nil)
}

// SRanipalEyeState  ...
type SRanipalEyeState struct {
	h C.WrapSRanipalEyeState
}

// GetPupilDiameterValid ...
func (pointer *SRanipalEyeState) GetPupilDiameterValid() bool {
	v := C.WrapSRanipalEyeStateGetPupilDiameterValid(pointer.h)
	return bool(v)
}

// SetPupilDiameterValid ...
func (pointer *SRanipalEyeState) SetPupilDiameterValid(v bool) {
	vToC := C.bool(v)
	C.WrapSRanipalEyeStateSetPupilDiameterValid(pointer.h, vToC)
}

// GetGazeOriginMm ...
func (pointer *SRanipalEyeState) GetGazeOriginMm() *Vec3 {
	v := C.WrapSRanipalEyeStateGetGazeOriginMm(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetGazeOriginMm ...
func (pointer *SRanipalEyeState) SetGazeOriginMm(v *Vec3) {
	vToC := v.h
	C.WrapSRanipalEyeStateSetGazeOriginMm(pointer.h, vToC)
}

// GetGazeDirectionNormalized ...
func (pointer *SRanipalEyeState) GetGazeDirectionNormalized() *Vec3 {
	v := C.WrapSRanipalEyeStateGetGazeDirectionNormalized(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetGazeDirectionNormalized ...
func (pointer *SRanipalEyeState) SetGazeDirectionNormalized(v *Vec3) {
	vToC := v.h
	C.WrapSRanipalEyeStateSetGazeDirectionNormalized(pointer.h, vToC)
}

// GetPupilDiameterMm ...
func (pointer *SRanipalEyeState) GetPupilDiameterMm() float32 {
	v := C.WrapSRanipalEyeStateGetPupilDiameterMm(pointer.h)
	return float32(v)
}

// SetPupilDiameterMm ...
func (pointer *SRanipalEyeState) SetPupilDiameterMm(v float32) {
	vToC := C.float(v)
	C.WrapSRanipalEyeStateSetPupilDiameterMm(pointer.h, vToC)
}

// GetEyeOpenness ...
func (pointer *SRanipalEyeState) GetEyeOpenness() float32 {
	v := C.WrapSRanipalEyeStateGetEyeOpenness(pointer.h)
	return float32(v)
}

// SetEyeOpenness ...
func (pointer *SRanipalEyeState) SetEyeOpenness(v float32) {
	vToC := C.float(v)
	C.WrapSRanipalEyeStateSetEyeOpenness(pointer.h, vToC)
}

// Free ...
func (pointer *SRanipalEyeState) Free() {
	C.WrapSRanipalEyeStateFree(pointer.h)
}

// IsNil ...
func (pointer *SRanipalEyeState) IsNil() bool {
	return pointer.h == C.WrapSRanipalEyeState(nil)
}

// SRanipalState  ...
type SRanipalState struct {
	h C.WrapSRanipalState
}

// GetRightEye ...
func (pointer *SRanipalState) GetRightEye() *SRanipalEyeState {
	v := C.WrapSRanipalStateGetRightEye(pointer.h)
	vGO := &SRanipalEyeState{h: v}
	return vGO
}

// SetRightEye ...
func (pointer *SRanipalState) SetRightEye(v *SRanipalEyeState) {
	vToC := v.h
	C.WrapSRanipalStateSetRightEye(pointer.h, vToC)
}

// GetLeftEye ...
func (pointer *SRanipalState) GetLeftEye() *SRanipalEyeState {
	v := C.WrapSRanipalStateGetLeftEye(pointer.h)
	vGO := &SRanipalEyeState{h: v}
	return vGO
}

// SetLeftEye ...
func (pointer *SRanipalState) SetLeftEye(v *SRanipalEyeState) {
	vToC := v.h
	C.WrapSRanipalStateSetLeftEye(pointer.h, vToC)
}

// Free ...
func (pointer *SRanipalState) Free() {
	C.WrapSRanipalStateFree(pointer.h)
}

// IsNil ...
func (pointer *SRanipalState) IsNil() bool {
	return pointer.h == C.WrapSRanipalState(nil)
}

// Vertex  ...
type Vertex struct {
	h C.WrapVertex
}

// GetPos ...
func (pointer *Vertex) GetPos() *Vec3 {
	v := C.WrapVertexGetPos(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetPos ...
func (pointer *Vertex) SetPos(v *Vec3) {
	vToC := v.h
	C.WrapVertexSetPos(pointer.h, vToC)
}

// GetNormal ...
func (pointer *Vertex) GetNormal() *Vec3 {
	v := C.WrapVertexGetNormal(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetNormal ...
func (pointer *Vertex) SetNormal(v *Vec3) {
	vToC := v.h
	C.WrapVertexSetNormal(pointer.h, vToC)
}

// GetTangent ...
func (pointer *Vertex) GetTangent() *Vec3 {
	v := C.WrapVertexGetTangent(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetTangent ...
func (pointer *Vertex) SetTangent(v *Vec3) {
	vToC := v.h
	C.WrapVertexSetTangent(pointer.h, vToC)
}

// GetBinormal ...
func (pointer *Vertex) GetBinormal() *Vec3 {
	v := C.WrapVertexGetBinormal(pointer.h)
	vGO := &Vec3{h: v}
	return vGO
}

// SetBinormal ...
func (pointer *Vertex) SetBinormal(v *Vec3) {
	vToC := v.h
	C.WrapVertexSetBinormal(pointer.h, vToC)
}

// GetUv0 ...
func (pointer *Vertex) GetUv0() *Vec2 {
	v := C.WrapVertexGetUv0(pointer.h)
	vGO := &Vec2{h: v}
	return vGO
}

// SetUv0 ...
func (pointer *Vertex) SetUv0(v *Vec2) {
	vToC := v.h
	C.WrapVertexSetUv0(pointer.h, vToC)
}

// GetUv1 ...
func (pointer *Vertex) GetUv1() *Vec2 {
	v := C.WrapVertexGetUv1(pointer.h)
	vGO := &Vec2{h: v}
	return vGO
}

// SetUv1 ...
func (pointer *Vertex) SetUv1(v *Vec2) {
	vToC := v.h
	C.WrapVertexSetUv1(pointer.h, vToC)
}

// GetUv2 ...
func (pointer *Vertex) GetUv2() *Vec2 {
	v := C.WrapVertexGetUv2(pointer.h)
	vGO := &Vec2{h: v}
	return vGO
}

// SetUv2 ...
func (pointer *Vertex) SetUv2(v *Vec2) {
	vToC := v.h
	C.WrapVertexSetUv2(pointer.h, vToC)
}

// GetUv3 ...
func (pointer *Vertex) GetUv3() *Vec2 {
	v := C.WrapVertexGetUv3(pointer.h)
	vGO := &Vec2{h: v}
	return vGO
}

// SetUv3 ...
func (pointer *Vertex) SetUv3(v *Vec2) {
	vToC := v.h
	C.WrapVertexSetUv3(pointer.h, vToC)
}

// GetUv4 ...
func (pointer *Vertex) GetUv4() *Vec2 {
	v := C.WrapVertexGetUv4(pointer.h)
	vGO := &Vec2{h: v}
	return vGO
}

// SetUv4 ...
func (pointer *Vertex) SetUv4(v *Vec2) {
	vToC := v.h
	C.WrapVertexSetUv4(pointer.h, vToC)
}

// GetUv5 ...
func (pointer *Vertex) GetUv5() *Vec2 {
	v := C.WrapVertexGetUv5(pointer.h)
	vGO := &Vec2{h: v}
	return vGO
}

// SetUv5 ...
func (pointer *Vertex) SetUv5(v *Vec2) {
	vToC := v.h
	C.WrapVertexSetUv5(pointer.h, vToC)
}

// GetUv6 ...
func (pointer *Vertex) GetUv6() *Vec2 {
	v := C.WrapVertexGetUv6(pointer.h)
	vGO := &Vec2{h: v}
	return vGO
}

// SetUv6 ...
func (pointer *Vertex) SetUv6(v *Vec2) {
	vToC := v.h
	C.WrapVertexSetUv6(pointer.h, vToC)
}

// GetUv7 ...
func (pointer *Vertex) GetUv7() *Vec2 {
	v := C.WrapVertexGetUv7(pointer.h)
	vGO := &Vec2{h: v}
	return vGO
}

// SetUv7 ...
func (pointer *Vertex) SetUv7(v *Vec2) {
	vToC := v.h
	C.WrapVertexSetUv7(pointer.h, vToC)
}

// GetColor0 ...
func (pointer *Vertex) GetColor0() *Color {
	v := C.WrapVertexGetColor0(pointer.h)
	vGO := &Color{h: v}
	return vGO
}

// SetColor0 ...
func (pointer *Vertex) SetColor0(v *Color) {
	vToC := v.h
	C.WrapVertexSetColor0(pointer.h, vToC)
}

// GetColor1 ...
func (pointer *Vertex) GetColor1() *Color {
	v := C.WrapVertexGetColor1(pointer.h)
	vGO := &Color{h: v}
	return vGO
}

// SetColor1 ...
func (pointer *Vertex) SetColor1(v *Color) {
	vToC := v.h
	C.WrapVertexSetColor1(pointer.h, vToC)
}

// GetColor2 ...
func (pointer *Vertex) GetColor2() *Color {
	v := C.WrapVertexGetColor2(pointer.h)
	vGO := &Color{h: v}
	return vGO
}

// SetColor2 ...
func (pointer *Vertex) SetColor2(v *Color) {
	vToC := v.h
	C.WrapVertexSetColor2(pointer.h, vToC)
}

// GetColor3 ...
func (pointer *Vertex) GetColor3() *Color {
	v := C.WrapVertexGetColor3(pointer.h)
	vGO := &Color{h: v}
	return vGO
}

// SetColor3 ...
func (pointer *Vertex) SetColor3(v *Color) {
	vToC := v.h
	C.WrapVertexSetColor3(pointer.h, vToC)
}

// NewVertex ...
func NewVertex() *Vertex {
	retval := C.WrapConstructorVertex()
	retvalGO := &Vertex{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vertex) {
		C.WrapVertexFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *Vertex) Free() {
	C.WrapVertexFree(pointer.h)
}

// IsNil ...
func (pointer *Vertex) IsNil() bool {
	return pointer.h == C.WrapVertex(nil)
}

// ModelBuilder  Use the model builder to programmatically build models at runtime.  The input data is optimized upon submission.
type ModelBuilder struct {
	h C.WrapModelBuilder
}

// NewModelBuilder Use the model builder to programmatically build models at runtime.  The input data is optimized upon submission.
func NewModelBuilder() *ModelBuilder {
	retval := C.WrapConstructorModelBuilder()
	retvalGO := &ModelBuilder{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ModelBuilder) {
		C.WrapModelBuilderFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *ModelBuilder) Free() {
	C.WrapModelBuilderFree(pointer.h)
}

// IsNil ...
func (pointer *ModelBuilder) IsNil() bool {
	return pointer.h == C.WrapModelBuilder(nil)
}

// AddVertex Add a vertex to the builder database. Use the returned optimized index to submit primitives using methods such as [harfang.ModelBuilder_AddTriangle].
func (pointer *ModelBuilder) AddVertex(vtx *Vertex) uint32 {
	vtxToC := vtx.h
	retval := C.WrapAddVertexModelBuilder(pointer.h, vtxToC)
	return uint32(retval)
}

// AddTriangle ...
func (pointer *ModelBuilder) AddTriangle(a uint32, b uint32, c uint32) {
	aToC := C.uint32_t(a)
	bToC := C.uint32_t(b)
	cToC := C.uint32_t(c)
	C.WrapAddTriangleModelBuilder(pointer.h, aToC, bToC, cToC)
}

// AddQuad ...
func (pointer *ModelBuilder) AddQuad(a uint32, b uint32, c uint32, d uint32) {
	aToC := C.uint32_t(a)
	bToC := C.uint32_t(b)
	cToC := C.uint32_t(c)
	dToC := C.uint32_t(d)
	C.WrapAddQuadModelBuilder(pointer.h, aToC, bToC, cToC, dToC)
}

// AddPolygon ...
func (pointer *ModelBuilder) AddPolygon(idxs *Uint32TList) {
	idxsToC := idxs.h
	C.WrapAddPolygonModelBuilder(pointer.h, idxsToC)
}

// GetCurrentListIndexCount Return the number of indexes in the current list.  See [harfang.ModelBuilder_EndList].
func (pointer *ModelBuilder) GetCurrentListIndexCount() int32 {
	retval := C.WrapGetCurrentListIndexCountModelBuilder(pointer.h)
	return int32(retval)
}

// EndList End the current primitive list and start a new one.
func (pointer *ModelBuilder) EndList(material uint16) {
	materialToC := C.ushort(material)
	C.WrapEndListModelBuilder(pointer.h, materialToC)
}

// Clear Clear all submitted data up to this point.
func (pointer *ModelBuilder) Clear() {
	C.WrapClearModelBuilder(pointer.h)
}

// MakeModel Create a model from all data submitted up to this point.
func (pointer *ModelBuilder) MakeModel(decl *VertexLayout) *Model {
	declToC := decl.h
	retval := C.WrapMakeModelModelBuilder(pointer.h, declToC)
	retvalGO := &Model{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Model) {
		C.WrapModelFree(cleanval.h)
	})
	return retvalGO
}

// Geometry  Base geometry object. Before a geometry can be displayed, it must be converted to [harfang.Model] by the asset compiler (see [harfang.man.AssetCompiler]).  To programmatically create a geometry use [harfang.GeometryBuilder].
type Geometry struct {
	h C.WrapGeometry
}

// Free ...
func (pointer *Geometry) Free() {
	C.WrapGeometryFree(pointer.h)
}

// IsNil ...
func (pointer *Geometry) IsNil() bool {
	return pointer.h == C.WrapGeometry(nil)
}

// GeometryBuilder  Use the geometry builder to programmatically create geometries. No optimization are performed by the geometry builder on the input data.  To programmatically build a geometry for immediate display see [harfang.ModelBuilder] to directly build models.
type GeometryBuilder struct {
	h C.WrapGeometryBuilder
}

// NewGeometryBuilder Use the geometry builder to programmatically create geometries. No optimization are performed by the geometry builder on the input data.  To programmatically build a geometry for immediate display see [harfang.ModelBuilder] to directly build models.
func NewGeometryBuilder() *GeometryBuilder {
	retval := C.WrapConstructorGeometryBuilder()
	retvalGO := &GeometryBuilder{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *GeometryBuilder) {
		C.WrapGeometryBuilderFree(cleanval.h)
	})
	return retvalGO
}

// Free ...
func (pointer *GeometryBuilder) Free() {
	C.WrapGeometryBuilderFree(pointer.h)
}

// IsNil ...
func (pointer *GeometryBuilder) IsNil() bool {
	return pointer.h == C.WrapGeometryBuilder(nil)
}

// AddVertex ...
func (pointer *GeometryBuilder) AddVertex(vtx *Vertex) {
	vtxToC := vtx.h
	C.WrapAddVertexGeometryBuilder(pointer.h, vtxToC)
}

// AddPolygon ...
func (pointer *GeometryBuilder) AddPolygon(idxs *Uint32TList, material uint16) {
	idxsToC := idxs.h
	materialToC := C.ushort(material)
	C.WrapAddPolygonGeometryBuilder(pointer.h, idxsToC, materialToC)
}

// AddPolygonWithSliceOfIdxs ...
func (pointer *GeometryBuilder) AddPolygonWithSliceOfIdxs(SliceOfidxs GoSliceOfuint32T, material uint16) {
	SliceOfidxsToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfidxs))
	SliceOfidxsToCSize := C.size_t(SliceOfidxsToC.Len)
	SliceOfidxsToCBuf := (*C.uint32_t)(unsafe.Pointer(SliceOfidxsToC.Data))
	materialToC := C.ushort(material)
	C.WrapAddPolygonGeometryBuilderWithSliceOfIdxs(pointer.h, SliceOfidxsToCSize, SliceOfidxsToCBuf, materialToC)
}

// AddTriangle ...
func (pointer *GeometryBuilder) AddTriangle(a uint32, b uint32, c uint32, material uint32) {
	aToC := C.uint32_t(a)
	bToC := C.uint32_t(b)
	cToC := C.uint32_t(c)
	materialToC := C.uint32_t(material)
	C.WrapAddTriangleGeometryBuilder(pointer.h, aToC, bToC, cToC, materialToC)
}

// AddQuad ...
func (pointer *GeometryBuilder) AddQuad(a uint32, b uint32, c uint32, d uint32, material uint32) {
	aToC := C.uint32_t(a)
	bToC := C.uint32_t(b)
	cToC := C.uint32_t(c)
	dToC := C.uint32_t(d)
	materialToC := C.uint32_t(material)
	C.WrapAddQuadGeometryBuilder(pointer.h, aToC, bToC, cToC, dToC, materialToC)
}

// Make ...
func (pointer *GeometryBuilder) Make() *Geometry {
	retval := C.WrapMakeGeometryBuilder(pointer.h)
	retvalGO := &Geometry{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Geometry) {
		C.WrapGeometryFree(cleanval.h)
	})
	return retvalGO
}

// Clear ...
func (pointer *GeometryBuilder) Clear() {
	C.WrapClearGeometryBuilder(pointer.h)
}

// IsoSurface  An iso-surface represents points of a constant value within a volume of space. This class holds a fixed-size 3-dimensional grid of values that can efficiently be converted to a [harfang.Model] at runtime.
type IsoSurface struct {
	h C.WrapIsoSurface
}

// Free ...
func (pointer *IsoSurface) Free() {
	C.WrapIsoSurfaceFree(pointer.h)
}

// IsNil ...
func (pointer *IsoSurface) IsNil() bool {
	return pointer.h == C.WrapIsoSurface(nil)
}

// Bloom  Bloom post-process object holding internal states and resources.  Create with [harfang.CreateBloomFromAssets] or [harfang.CreateBloomFromFile], use with [harfang.ApplyBloom], finally call [harfang.DestroyBloom] to dispose of resources when done.
type Bloom struct {
	h C.WrapBloom
}

// Free ...
func (pointer *Bloom) Free() {
	C.WrapBloomFree(pointer.h)
}

// IsNil ...
func (pointer *Bloom) IsNil() bool {
	return pointer.h == C.WrapBloom(nil)
}

// SAO  Ambient occlusion post-process object holding internal states and resources.  Create with [harfang.CreateSAOFromFile] or [harfang.CreateSAOFromAssets], use with [harfang.ComputeSAO], finally call [harfang.DestroySAO] to dispose of resources when done.
type SAO struct {
	h C.WrapSAO
}

// Free ...
func (pointer *SAO) Free() {
	C.WrapSAOFree(pointer.h)
}

// IsNil ...
func (pointer *SAO) IsNil() bool {
	return pointer.h == C.WrapSAO(nil)
}

// ProfilerFrame  ...
type ProfilerFrame struct {
	h C.WrapProfilerFrame
}

// Free ...
func (pointer *ProfilerFrame) Free() {
	C.WrapProfilerFrameFree(pointer.h)
}

// IsNil ...
func (pointer *ProfilerFrame) IsNil() bool {
	return pointer.h == C.WrapProfilerFrame(nil)
}

// IVideoStreamer  ...
type IVideoStreamer struct {
	h C.WrapIVideoStreamer
}

// Free ...
func (pointer *IVideoStreamer) Free() {
	C.WrapIVideoStreamerFree(pointer.h)
}

// IsNil ...
func (pointer *IVideoStreamer) IsNil() bool {
	return pointer.h == C.WrapIVideoStreamer(nil)
}

// Startup ...
func (pointer *IVideoStreamer) Startup() int32 {
	retval := C.WrapStartupIVideoStreamer(pointer.h)
	return int32(retval)
}

// Shutdown ...
func (pointer *IVideoStreamer) Shutdown() {
	C.WrapShutdownIVideoStreamer(pointer.h)
}

// Open ...
func (pointer *IVideoStreamer) Open(name string) uintptr {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapOpenIVideoStreamer(pointer.h, nameToC)
	return uintptr(retval)
}

// Play ...
func (pointer *IVideoStreamer) Play(h uintptr) int32 {
	hToC := C.intptr_t(h)
	retval := C.WrapPlayIVideoStreamer(pointer.h, hToC)
	return int32(retval)
}

// Pause ...
func (pointer *IVideoStreamer) Pause(h uintptr) int32 {
	hToC := C.intptr_t(h)
	retval := C.WrapPauseIVideoStreamer(pointer.h, hToC)
	return int32(retval)
}

// Close ...
func (pointer *IVideoStreamer) Close(h uintptr) int32 {
	hToC := C.intptr_t(h)
	retval := C.WrapCloseIVideoStreamer(pointer.h, hToC)
	return int32(retval)
}

// Seek ...
func (pointer *IVideoStreamer) Seek(h uintptr, t int64) int32 {
	hToC := C.intptr_t(h)
	tToC := C.int64_t(t)
	retval := C.WrapSeekIVideoStreamer(pointer.h, hToC, tToC)
	return int32(retval)
}

// GetDuration ...
func (pointer *IVideoStreamer) GetDuration(h uintptr) int64 {
	hToC := C.intptr_t(h)
	retval := C.WrapGetDurationIVideoStreamer(pointer.h, hToC)
	return int64(retval)
}

// GetTimeStamp ...
func (pointer *IVideoStreamer) GetTimeStamp(h uintptr) int64 {
	hToC := C.intptr_t(h)
	retval := C.WrapGetTimeStampIVideoStreamer(pointer.h, hToC)
	return int64(retval)
}

// IsEnded ...
func (pointer *IVideoStreamer) IsEnded(h uintptr) int32 {
	hToC := C.intptr_t(h)
	retval := C.WrapIsEndedIVideoStreamer(pointer.h, hToC)
	return int32(retval)
}

// GetFrame ...
func (pointer *IVideoStreamer) GetFrame(h uintptr, ptr *uintptr, width *int32, height *int32, pitch *int32, format *int32) int32 {
	hToC := C.intptr_t(h)
	ptrToC := (*C.intptr_t)(unsafe.Pointer(ptr))
	widthToC := (*C.int32_t)(unsafe.Pointer(width))
	heightToC := (*C.int32_t)(unsafe.Pointer(height))
	pitchToC := (*C.int32_t)(unsafe.Pointer(pitch))
	formatToC := (*C.int32_t)(unsafe.Pointer(format))
	retval := C.WrapGetFrameIVideoStreamer(pointer.h, hToC, ptrToC, widthToC, heightToC, pitchToC, formatToC)
	return int32(retval)
}

// FreeFrame ...
func (pointer *IVideoStreamer) FreeFrame(h uintptr, frame int32) int32 {
	hToC := C.intptr_t(h)
	frameToC := C.int32_t(frame)
	retval := C.WrapFreeFrameIVideoStreamer(pointer.h, hToC, frameToC)
	return int32(retval)
}

// LogLevel ...
type LogLevel int32

var (
	// LLNormal ...
	LLNormal = LogLevel(C.GetLogLevel(0))
	// LLWarning ...
	LLWarning = LogLevel(C.GetLogLevel(1))
	// LLError ...
	LLError = LogLevel(C.GetLogLevel(2))
	// LLDebug ...
	LLDebug = LogLevel(C.GetLogLevel(3))
	// LLAll ...
	LLAll = LogLevel(C.GetLogLevel(4))
)

// SeekMode ...
type SeekMode int32

var (
	// SMStart ...
	SMStart = SeekMode(C.GetSeekMode(0))
	// SMCurrent ...
	SMCurrent = SeekMode(C.GetSeekMode(1))
	// SMEnd ...
	SMEnd = SeekMode(C.GetSeekMode(2))
)

// DirEntryType ...
type DirEntryType int32

var (
	// DEFile ...
	DEFile = DirEntryType(C.GetDirEntryType(0))
	// DEDir ...
	DEDir = DirEntryType(C.GetDirEntryType(1))
	// DELink ...
	DELink = DirEntryType(C.GetDirEntryType(2))
	// DEAll ...
	DEAll = DirEntryType(C.GetDirEntryType(3))
)

// RotationOrder ...
type RotationOrder uint8

var (
	// ROZYX ...
	ROZYX = RotationOrder(C.GetRotationOrder(0))
	// ROYZX ...
	ROYZX = RotationOrder(C.GetRotationOrder(1))
	// ROZXY ...
	ROZXY = RotationOrder(C.GetRotationOrder(2))
	// ROXZY ...
	ROXZY = RotationOrder(C.GetRotationOrder(3))
	// ROYXZ ...
	ROYXZ = RotationOrder(C.GetRotationOrder(4))
	// ROXYZ ...
	ROXYZ = RotationOrder(C.GetRotationOrder(5))
	// ROXY ...
	ROXY = RotationOrder(C.GetRotationOrder(6))
	// RODefault ...
	RODefault = RotationOrder(C.GetRotationOrder(7))
)

// Axis ...
type Axis uint8

var (
	// AX ...
	AX = Axis(C.GetAxis(0))
	// AY ...
	AY = Axis(C.GetAxis(1))
	// AZ ...
	AZ = Axis(C.GetAxis(2))
	// ARotX ...
	ARotX = Axis(C.GetAxis(3))
	// ARotY ...
	ARotY = Axis(C.GetAxis(4))
	// ARotZ ...
	ARotZ = Axis(C.GetAxis(5))
	// ALast ...
	ALast = Axis(C.GetAxis(6))
)

// Visibility ...
type Visibility uint8

var (
	// VOutside ...
	VOutside = Visibility(C.GetVisibility(0))
	// VInside ...
	VInside = Visibility(C.GetVisibility(1))
	// VClipped ...
	VClipped = Visibility(C.GetVisibility(2))
)

// MonitorRotation ...
type MonitorRotation uint8

var (
	// MR0 ...
	MR0 = MonitorRotation(C.GetMonitorRotation(0))
	// MR90 ...
	MR90 = MonitorRotation(C.GetMonitorRotation(1))
	// MR180 ...
	MR180 = MonitorRotation(C.GetMonitorRotation(2))
	// MR270 ...
	MR270 = MonitorRotation(C.GetMonitorRotation(3))
)

// WindowVisibility ...
type WindowVisibility int32

var (
	// WVWindowed ...
	WVWindowed = WindowVisibility(C.GetWindowVisibility(0))
	// WVUndecorated ...
	WVUndecorated = WindowVisibility(C.GetWindowVisibility(1))
	// WVFullscreen ...
	WVFullscreen = WindowVisibility(C.GetWindowVisibility(2))
	// WVHidden ...
	WVHidden = WindowVisibility(C.GetWindowVisibility(3))
	// WVFullscreenMonitor1 ...
	WVFullscreenMonitor1 = WindowVisibility(C.GetWindowVisibility(4))
	// WVFullscreenMonitor2 ...
	WVFullscreenMonitor2 = WindowVisibility(C.GetWindowVisibility(5))
	// WVFullscreenMonitor3 ...
	WVFullscreenMonitor3 = WindowVisibility(C.GetWindowVisibility(6))
)

// PictureFormat ...
type PictureFormat int32

var (
	// PFRGB24 ...
	PFRGB24 = PictureFormat(C.GetPictureFormat(0))
	// PFRGBA32 ...
	PFRGBA32 = PictureFormat(C.GetPictureFormat(1))
	// PFRGBA32F ...
	PFRGBA32F = PictureFormat(C.GetPictureFormat(2))
)

// RendererType ...
type RendererType int32

var (
	// RTNoop ...
	RTNoop = RendererType(C.GetRendererType(0))
	// RTDirect3D9 ...
	RTDirect3D9 = RendererType(C.GetRendererType(1))
	// RTDirect3D11 ...
	RTDirect3D11 = RendererType(C.GetRendererType(2))
	// RTDirect3D12 ...
	RTDirect3D12 = RendererType(C.GetRendererType(3))
	// RTGnm ...
	RTGnm = RendererType(C.GetRendererType(4))
	// RTMetal ...
	RTMetal = RendererType(C.GetRendererType(5))
	// RTNvn ...
	RTNvn = RendererType(C.GetRendererType(6))
	// RTOpenGLES ...
	RTOpenGLES = RendererType(C.GetRendererType(7))
	// RTOpenGL ...
	RTOpenGL = RendererType(C.GetRendererType(8))
	// RTVulkan ...
	RTVulkan = RendererType(C.GetRendererType(9))
	// RTCount ...
	RTCount = RendererType(C.GetRendererType(10))
)

// TextureFormat ...
type TextureFormat int32

var (
	// TFBC1 ...
	TFBC1 = TextureFormat(C.GetTextureFormat(0))
	// TFBC2 ...
	TFBC2 = TextureFormat(C.GetTextureFormat(1))
	// TFBC3 ...
	TFBC3 = TextureFormat(C.GetTextureFormat(2))
	// TFBC4 ...
	TFBC4 = TextureFormat(C.GetTextureFormat(3))
	// TFBC5 ...
	TFBC5 = TextureFormat(C.GetTextureFormat(4))
	// TFBC6H ...
	TFBC6H = TextureFormat(C.GetTextureFormat(5))
	// TFBC7 ...
	TFBC7 = TextureFormat(C.GetTextureFormat(6))
	// TFETC1 ...
	TFETC1 = TextureFormat(C.GetTextureFormat(7))
	// TFETC2 ...
	TFETC2 = TextureFormat(C.GetTextureFormat(8))
	// TFETC2A ...
	TFETC2A = TextureFormat(C.GetTextureFormat(9))
	// TFETC2A1 ...
	TFETC2A1 = TextureFormat(C.GetTextureFormat(10))
	// TFPTC12 ...
	TFPTC12 = TextureFormat(C.GetTextureFormat(11))
	// TFPTC14 ...
	TFPTC14 = TextureFormat(C.GetTextureFormat(12))
	// TFPTC12A ...
	TFPTC12A = TextureFormat(C.GetTextureFormat(13))
	// TFPTC14A ...
	TFPTC14A = TextureFormat(C.GetTextureFormat(14))
	// TFPTC22 ...
	TFPTC22 = TextureFormat(C.GetTextureFormat(15))
	// TFPTC24 ...
	TFPTC24 = TextureFormat(C.GetTextureFormat(16))
	// TFATC ...
	TFATC = TextureFormat(C.GetTextureFormat(17))
	// TFATCE ...
	TFATCE = TextureFormat(C.GetTextureFormat(18))
	// TFATCI ...
	TFATCI = TextureFormat(C.GetTextureFormat(19))
	// TFASTC4x4 ...
	TFASTC4x4 = TextureFormat(C.GetTextureFormat(20))
	// TFASTC5x5 ...
	TFASTC5x5 = TextureFormat(C.GetTextureFormat(21))
	// TFASTC6x6 ...
	TFASTC6x6 = TextureFormat(C.GetTextureFormat(22))
	// TFASTC8x5 ...
	TFASTC8x5 = TextureFormat(C.GetTextureFormat(23))
	// TFASTC8x6 ...
	TFASTC8x6 = TextureFormat(C.GetTextureFormat(24))
	// TFASTC10x5 ...
	TFASTC10x5 = TextureFormat(C.GetTextureFormat(25))
	// TFUnknown ...
	TFUnknown = TextureFormat(C.GetTextureFormat(26))
	// TFR1 ...
	TFR1 = TextureFormat(C.GetTextureFormat(27))
	// TFA8 ...
	TFA8 = TextureFormat(C.GetTextureFormat(28))
	// TFR8 ...
	TFR8 = TextureFormat(C.GetTextureFormat(29))
	// TFR8I ...
	TFR8I = TextureFormat(C.GetTextureFormat(30))
	// TFR8U ...
	TFR8U = TextureFormat(C.GetTextureFormat(31))
	// TFR8S ...
	TFR8S = TextureFormat(C.GetTextureFormat(32))
	// TFR16 ...
	TFR16 = TextureFormat(C.GetTextureFormat(33))
	// TFR16I ...
	TFR16I = TextureFormat(C.GetTextureFormat(34))
	// TFR16U ...
	TFR16U = TextureFormat(C.GetTextureFormat(35))
	// TFR16F ...
	TFR16F = TextureFormat(C.GetTextureFormat(36))
	// TFR16S ...
	TFR16S = TextureFormat(C.GetTextureFormat(37))
	// TFR32I ...
	TFR32I = TextureFormat(C.GetTextureFormat(38))
	// TFR32U ...
	TFR32U = TextureFormat(C.GetTextureFormat(39))
	// TFR32F ...
	TFR32F = TextureFormat(C.GetTextureFormat(40))
	// TFRG8 ...
	TFRG8 = TextureFormat(C.GetTextureFormat(41))
	// TFRG8I ...
	TFRG8I = TextureFormat(C.GetTextureFormat(42))
	// TFRG8U ...
	TFRG8U = TextureFormat(C.GetTextureFormat(43))
	// TFRG8S ...
	TFRG8S = TextureFormat(C.GetTextureFormat(44))
	// TFRG16 ...
	TFRG16 = TextureFormat(C.GetTextureFormat(45))
	// TFRG16I ...
	TFRG16I = TextureFormat(C.GetTextureFormat(46))
	// TFRG16U ...
	TFRG16U = TextureFormat(C.GetTextureFormat(47))
	// TFRG16F ...
	TFRG16F = TextureFormat(C.GetTextureFormat(48))
	// TFRG16S ...
	TFRG16S = TextureFormat(C.GetTextureFormat(49))
	// TFRG32I ...
	TFRG32I = TextureFormat(C.GetTextureFormat(50))
	// TFRG32U ...
	TFRG32U = TextureFormat(C.GetTextureFormat(51))
	// TFRG32F ...
	TFRG32F = TextureFormat(C.GetTextureFormat(52))
	// TFRGB8 ...
	TFRGB8 = TextureFormat(C.GetTextureFormat(53))
	// TFRGB8I ...
	TFRGB8I = TextureFormat(C.GetTextureFormat(54))
	// TFRGB8U ...
	TFRGB8U = TextureFormat(C.GetTextureFormat(55))
	// TFRGB8S ...
	TFRGB8S = TextureFormat(C.GetTextureFormat(56))
	// TFRGB9E5F ...
	TFRGB9E5F = TextureFormat(C.GetTextureFormat(57))
	// TFBGRA8 ...
	TFBGRA8 = TextureFormat(C.GetTextureFormat(58))
	// TFRGBA8 ...
	TFRGBA8 = TextureFormat(C.GetTextureFormat(59))
	// TFRGBA8I ...
	TFRGBA8I = TextureFormat(C.GetTextureFormat(60))
	// TFRGBA8U ...
	TFRGBA8U = TextureFormat(C.GetTextureFormat(61))
	// TFRGBA8S ...
	TFRGBA8S = TextureFormat(C.GetTextureFormat(62))
	// TFRGBA16 ...
	TFRGBA16 = TextureFormat(C.GetTextureFormat(63))
	// TFRGBA16I ...
	TFRGBA16I = TextureFormat(C.GetTextureFormat(64))
	// TFRGBA16U ...
	TFRGBA16U = TextureFormat(C.GetTextureFormat(65))
	// TFRGBA16F ...
	TFRGBA16F = TextureFormat(C.GetTextureFormat(66))
	// TFRGBA16S ...
	TFRGBA16S = TextureFormat(C.GetTextureFormat(67))
	// TFRGBA32I ...
	TFRGBA32I = TextureFormat(C.GetTextureFormat(68))
	// TFRGBA32U ...
	TFRGBA32U = TextureFormat(C.GetTextureFormat(69))
	// TFRGBA32F ...
	TFRGBA32F = TextureFormat(C.GetTextureFormat(70))
	// TFR5G6B5 ...
	TFR5G6B5 = TextureFormat(C.GetTextureFormat(71))
	// TFRGBA4 ...
	TFRGBA4 = TextureFormat(C.GetTextureFormat(72))
	// TFRGB5A1 ...
	TFRGB5A1 = TextureFormat(C.GetTextureFormat(73))
	// TFRGB10A2 ...
	TFRGB10A2 = TextureFormat(C.GetTextureFormat(74))
	// TFRG11B10F ...
	TFRG11B10F = TextureFormat(C.GetTextureFormat(75))
	// TFUnknownDepth ...
	TFUnknownDepth = TextureFormat(C.GetTextureFormat(76))
	// TFD16 ...
	TFD16 = TextureFormat(C.GetTextureFormat(77))
	// TFD24 ...
	TFD24 = TextureFormat(C.GetTextureFormat(78))
	// TFD24S8 ...
	TFD24S8 = TextureFormat(C.GetTextureFormat(79))
	// TFD32 ...
	TFD32 = TextureFormat(C.GetTextureFormat(80))
	// TFD16F ...
	TFD16F = TextureFormat(C.GetTextureFormat(81))
	// TFD24F ...
	TFD24F = TextureFormat(C.GetTextureFormat(82))
	// TFD32F ...
	TFD32F = TextureFormat(C.GetTextureFormat(83))
	// TFD0S8 ...
	TFD0S8 = TextureFormat(C.GetTextureFormat(84))
)

// BackbufferRatio ...
type BackbufferRatio int32

var (
	// BREqual ...
	BREqual = BackbufferRatio(C.GetBackbufferRatio(0))
	// BRHalf ...
	BRHalf = BackbufferRatio(C.GetBackbufferRatio(1))
	// BRQuarter ...
	BRQuarter = BackbufferRatio(C.GetBackbufferRatio(2))
	// BREighth ...
	BREighth = BackbufferRatio(C.GetBackbufferRatio(3))
	// BRSixteenth ...
	BRSixteenth = BackbufferRatio(C.GetBackbufferRatio(4))
	// BRDouble ...
	BRDouble = BackbufferRatio(C.GetBackbufferRatio(5))
)

// ViewMode ...
type ViewMode int32

var (
	// VMDefault ...
	VMDefault = ViewMode(C.GetViewMode(0))
	// VMSequential ...
	VMSequential = ViewMode(C.GetViewMode(1))
	// VMDepthAscending ...
	VMDepthAscending = ViewMode(C.GetViewMode(2))
	// VMDepthDescending ...
	VMDepthDescending = ViewMode(C.GetViewMode(3))
)

// Attrib ...
type Attrib int32

var (
	// APosition ...
	APosition = Attrib(C.GetAttrib(0))
	// ANormal ...
	ANormal = Attrib(C.GetAttrib(1))
	// ATangent ...
	ATangent = Attrib(C.GetAttrib(2))
	// ABitangent ...
	ABitangent = Attrib(C.GetAttrib(3))
	// AColor0 ...
	AColor0 = Attrib(C.GetAttrib(4))
	// AColor1 ...
	AColor1 = Attrib(C.GetAttrib(5))
	// AColor2 ...
	AColor2 = Attrib(C.GetAttrib(6))
	// AColor3 ...
	AColor3 = Attrib(C.GetAttrib(7))
	// AIndices ...
	AIndices = Attrib(C.GetAttrib(8))
	// AWeight ...
	AWeight = Attrib(C.GetAttrib(9))
	// ATexCoord0 ...
	ATexCoord0 = Attrib(C.GetAttrib(10))
	// ATexCoord1 ...
	ATexCoord1 = Attrib(C.GetAttrib(11))
	// ATexCoord2 ...
	ATexCoord2 = Attrib(C.GetAttrib(12))
	// ATexCoord3 ...
	ATexCoord3 = Attrib(C.GetAttrib(13))
	// ATexCoord4 ...
	ATexCoord4 = Attrib(C.GetAttrib(14))
	// ATexCoord5 ...
	ATexCoord5 = Attrib(C.GetAttrib(15))
	// ATexCoord6 ...
	ATexCoord6 = Attrib(C.GetAttrib(16))
	// ATexCoord7 ...
	ATexCoord7 = Attrib(C.GetAttrib(17))
)

// AttribType ...
type AttribType int32

var (
	// ATUint8 ...
	ATUint8 = AttribType(C.GetAttribType(0))
	// ATUint10 ...
	ATUint10 = AttribType(C.GetAttribType(1))
	// ATInt16 ...
	ATInt16 = AttribType(C.GetAttribType(2))
	// ATHalf ...
	ATHalf = AttribType(C.GetAttribType(3))
	// ATFloat ...
	ATFloat = AttribType(C.GetAttribType(4))
)

// FaceCulling ...
type FaceCulling int32

var (
	// FCDisabled ...
	FCDisabled = FaceCulling(C.GetFaceCulling(0))
	// FCClockwise ...
	FCClockwise = FaceCulling(C.GetFaceCulling(1))
	// FCCounterClockwise ...
	FCCounterClockwise = FaceCulling(C.GetFaceCulling(2))
)

// DepthTest ...
type DepthTest int32

var (
	// DTLess ...
	DTLess = DepthTest(C.GetDepthTest(0))
	// DTLessEqual ...
	DTLessEqual = DepthTest(C.GetDepthTest(1))
	// DTEqual ...
	DTEqual = DepthTest(C.GetDepthTest(2))
	// DTGreaterEqual ...
	DTGreaterEqual = DepthTest(C.GetDepthTest(3))
	// DTGreater ...
	DTGreater = DepthTest(C.GetDepthTest(4))
	// DTNotEqual ...
	DTNotEqual = DepthTest(C.GetDepthTest(5))
	// DTNever ...
	DTNever = DepthTest(C.GetDepthTest(6))
	// DTAlways ...
	DTAlways = DepthTest(C.GetDepthTest(7))
	// DTDisabled ...
	DTDisabled = DepthTest(C.GetDepthTest(8))
)

// BlendMode ...
type BlendMode int32

var (
	// BMAdditive ...
	BMAdditive = BlendMode(C.GetBlendMode(0))
	// BMAlpha ...
	BMAlpha = BlendMode(C.GetBlendMode(1))
	// BMDarken ...
	BMDarken = BlendMode(C.GetBlendMode(2))
	// BMLighten ...
	BMLighten = BlendMode(C.GetBlendMode(3))
	// BMMultiply ...
	BMMultiply = BlendMode(C.GetBlendMode(4))
	// BMOpaque ...
	BMOpaque = BlendMode(C.GetBlendMode(5))
	// BMScreen ...
	BMScreen = BlendMode(C.GetBlendMode(6))
	// BMLinearBurn ...
	BMLinearBurn = BlendMode(C.GetBlendMode(7))
	// BMUndefined ...
	BMUndefined = BlendMode(C.GetBlendMode(8))
)

// ForwardPipelineLightType ...
type ForwardPipelineLightType int32

var (
	// FPLTPoint ...
	FPLTPoint = ForwardPipelineLightType(C.GetForwardPipelineLightType(0))
	// FPLTSpot ...
	FPLTSpot = ForwardPipelineLightType(C.GetForwardPipelineLightType(1))
	// FPLTLinear ...
	FPLTLinear = ForwardPipelineLightType(C.GetForwardPipelineLightType(2))
)

// ForwardPipelineShadowType ...
type ForwardPipelineShadowType int32

var (
	// FPSTNone ...
	FPSTNone = ForwardPipelineShadowType(C.GetForwardPipelineShadowType(0))
	// FPSTMap ...
	FPSTMap = ForwardPipelineShadowType(C.GetForwardPipelineShadowType(1))
)

// DrawTextHAlign ...
type DrawTextHAlign int32

var (
	// DTHALeft ...
	DTHALeft = DrawTextHAlign(C.GetDrawTextHAlign(0))
	// DTHACenter ...
	DTHACenter = DrawTextHAlign(C.GetDrawTextHAlign(1))
	// DTHARight ...
	DTHARight = DrawTextHAlign(C.GetDrawTextHAlign(2))
)

// DrawTextVAlign ...
type DrawTextVAlign int32

var (
	// DTVATop ...
	DTVATop = DrawTextVAlign(C.GetDrawTextVAlign(0))
	// DTVACenter ...
	DTVACenter = DrawTextVAlign(C.GetDrawTextVAlign(1))
	// DTVABottom ...
	DTVABottom = DrawTextVAlign(C.GetDrawTextVAlign(2))
)

// AnimLoopMode ...
type AnimLoopMode int32

var (
	// ALMOnce ...
	ALMOnce = AnimLoopMode(C.GetAnimLoopMode(0))
	// ALMInfinite ...
	ALMInfinite = AnimLoopMode(C.GetAnimLoopMode(1))
	// ALMLoop ...
	ALMLoop = AnimLoopMode(C.GetAnimLoopMode(2))
)

// Easing ...
type Easing uint8

var (
	// ELinear ...
	ELinear = Easing(C.GetEasing(0))
	// EStep ...
	EStep = Easing(C.GetEasing(1))
	// ESmoothStep ...
	ESmoothStep = Easing(C.GetEasing(2))
	// EInQuad ...
	EInQuad = Easing(C.GetEasing(3))
	// EOutQuad ...
	EOutQuad = Easing(C.GetEasing(4))
	// EInOutQuad ...
	EInOutQuad = Easing(C.GetEasing(5))
	// EOutInQuad ...
	EOutInQuad = Easing(C.GetEasing(6))
	// EInCubic ...
	EInCubic = Easing(C.GetEasing(7))
	// EOutCubic ...
	EOutCubic = Easing(C.GetEasing(8))
	// EInOutCubic ...
	EInOutCubic = Easing(C.GetEasing(9))
	// EOutInCubic ...
	EOutInCubic = Easing(C.GetEasing(10))
	// EInQuart ...
	EInQuart = Easing(C.GetEasing(11))
	// EOutQuart ...
	EOutQuart = Easing(C.GetEasing(12))
	// EInOutQuart ...
	EInOutQuart = Easing(C.GetEasing(13))
	// EOutInQuart ...
	EOutInQuart = Easing(C.GetEasing(14))
	// EInQuint ...
	EInQuint = Easing(C.GetEasing(15))
	// EOutQuint ...
	EOutQuint = Easing(C.GetEasing(16))
	// EInOutQuint ...
	EInOutQuint = Easing(C.GetEasing(17))
	// EOutInQuint ...
	EOutInQuint = Easing(C.GetEasing(18))
	// EInSine ...
	EInSine = Easing(C.GetEasing(19))
	// EOutSine ...
	EOutSine = Easing(C.GetEasing(20))
	// EInOutSine ...
	EInOutSine = Easing(C.GetEasing(21))
	// EOutInSine ...
	EOutInSine = Easing(C.GetEasing(22))
	// EInExpo ...
	EInExpo = Easing(C.GetEasing(23))
	// EOutExpo ...
	EOutExpo = Easing(C.GetEasing(24))
	// EInOutExpo ...
	EInOutExpo = Easing(C.GetEasing(25))
	// EOutInExpo ...
	EOutInExpo = Easing(C.GetEasing(26))
	// EInCirc ...
	EInCirc = Easing(C.GetEasing(27))
	// EOutCirc ...
	EOutCirc = Easing(C.GetEasing(28))
	// EInOutCirc ...
	EInOutCirc = Easing(C.GetEasing(29))
	// EOutInCirc ...
	EOutInCirc = Easing(C.GetEasing(30))
	// EInElastic ...
	EInElastic = Easing(C.GetEasing(31))
	// EOutElastic ...
	EOutElastic = Easing(C.GetEasing(32))
	// EInOutElastic ...
	EInOutElastic = Easing(C.GetEasing(33))
	// EOutInElastic ...
	EOutInElastic = Easing(C.GetEasing(34))
	// EInBack ...
	EInBack = Easing(C.GetEasing(35))
	// EOutBack ...
	EOutBack = Easing(C.GetEasing(36))
	// EInOutBack ...
	EInOutBack = Easing(C.GetEasing(37))
	// EOutInBack ...
	EOutInBack = Easing(C.GetEasing(38))
	// EInBounce ...
	EInBounce = Easing(C.GetEasing(39))
	// EOutBounce ...
	EOutBounce = Easing(C.GetEasing(40))
	// EInOutBounce ...
	EInOutBounce = Easing(C.GetEasing(41))
	// EOutInBounce ...
	EOutInBounce = Easing(C.GetEasing(42))
)

// LightType ...
type LightType int32

var (
	// LTPoint ...
	LTPoint = LightType(C.GetLightType(0))
	// LTSpot ...
	LTSpot = LightType(C.GetLightType(1))
	// LTLinear ...
	LTLinear = LightType(C.GetLightType(2))
)

// LightShadowType ...
type LightShadowType int32

var (
	// LSTNone ...
	LSTNone = LightShadowType(C.GetLightShadowType(0))
	// LSTMap ...
	LSTMap = LightShadowType(C.GetLightShadowType(1))
)

// RigidBodyType ...
type RigidBodyType uint8

var (
	// RBTDynamic ...
	RBTDynamic = RigidBodyType(C.GetRigidBodyType(0))
	// RBTKinematic ...
	RBTKinematic = RigidBodyType(C.GetRigidBodyType(1))
	// RBTStatic ...
	RBTStatic = RigidBodyType(C.GetRigidBodyType(2))
)

// CollisionEventTrackingMode ...
type CollisionEventTrackingMode uint8

var (
	// CETMEventOnly ...
	CETMEventOnly = CollisionEventTrackingMode(C.GetCollisionEventTrackingMode(0))
	// CETMEventAndContacts ...
	CETMEventAndContacts = CollisionEventTrackingMode(C.GetCollisionEventTrackingMode(1))
)

// CollisionType ...
type CollisionType uint8

var (
	// CTSphere ...
	CTSphere = CollisionType(C.GetCollisionType(0))
	// CTCube ...
	CTCube = CollisionType(C.GetCollisionType(1))
	// CTCone ...
	CTCone = CollisionType(C.GetCollisionType(2))
	// CTCapsule ...
	CTCapsule = CollisionType(C.GetCollisionType(3))
	// CTCylinder ...
	CTCylinder = CollisionType(C.GetCollisionType(4))
	// CTMesh ...
	CTMesh = CollisionType(C.GetCollisionType(5))
)

// NodeComponentIdx ...
type NodeComponentIdx int32

var (
	// NCITransform ...
	NCITransform = NodeComponentIdx(C.GetNodeComponentIdx(0))
	// NCICamera ...
	NCICamera = NodeComponentIdx(C.GetNodeComponentIdx(1))
	// NCIObject ...
	NCIObject = NodeComponentIdx(C.GetNodeComponentIdx(2))
	// NCILight ...
	NCILight = NodeComponentIdx(C.GetNodeComponentIdx(3))
	// NCIRigidBody ...
	NCIRigidBody = NodeComponentIdx(C.GetNodeComponentIdx(4))
)

// SceneForwardPipelinePass ...
type SceneForwardPipelinePass int32

var (
	// SFPPOpaque ...
	SFPPOpaque = SceneForwardPipelinePass(C.GetSceneForwardPipelinePass(0))
	// SFPPTransparent ...
	SFPPTransparent = SceneForwardPipelinePass(C.GetSceneForwardPipelinePass(1))
	// SFPPSlot0LinearSplit0 ...
	SFPPSlot0LinearSplit0 = SceneForwardPipelinePass(C.GetSceneForwardPipelinePass(2))
	// SFPPSlot0LinearSplit1 ...
	SFPPSlot0LinearSplit1 = SceneForwardPipelinePass(C.GetSceneForwardPipelinePass(3))
	// SFPPSlot0LinearSplit2 ...
	SFPPSlot0LinearSplit2 = SceneForwardPipelinePass(C.GetSceneForwardPipelinePass(4))
	// SFPPSlot0LinearSplit3 ...
	SFPPSlot0LinearSplit3 = SceneForwardPipelinePass(C.GetSceneForwardPipelinePass(5))
	// SFPPSlot1Spot ...
	SFPPSlot1Spot = SceneForwardPipelinePass(C.GetSceneForwardPipelinePass(6))
	// SFPPDepthPrepass ...
	SFPPDepthPrepass = SceneForwardPipelinePass(C.GetSceneForwardPipelinePass(7))
)

// ForwardPipelineAAADebugBuffer ...
type ForwardPipelineAAADebugBuffer int32

var (
	// FPAAADBNone ...
	FPAAADBNone = ForwardPipelineAAADebugBuffer(C.GetForwardPipelineAAADebugBuffer(0))
	// FPAAADBSSGI ...
	FPAAADBSSGI = ForwardPipelineAAADebugBuffer(C.GetForwardPipelineAAADebugBuffer(1))
	// FPAAADBSSR ...
	FPAAADBSSR = ForwardPipelineAAADebugBuffer(C.GetForwardPipelineAAADebugBuffer(2))
)

// MouseButton ...
type MouseButton int32

var (
	// MB0 ...
	MB0 = MouseButton(C.GetMouseButton(0))
	// MB1 ...
	MB1 = MouseButton(C.GetMouseButton(1))
	// MB2 ...
	MB2 = MouseButton(C.GetMouseButton(2))
	// MB3 ...
	MB3 = MouseButton(C.GetMouseButton(3))
	// MB4 ...
	MB4 = MouseButton(C.GetMouseButton(4))
	// MB5 ...
	MB5 = MouseButton(C.GetMouseButton(5))
	// MB6 ...
	MB6 = MouseButton(C.GetMouseButton(6))
	// MB7 ...
	MB7 = MouseButton(C.GetMouseButton(7))
)

// Key ...
type Key int32

var (
	// KLShift ...
	KLShift = Key(C.GetKey(0))
	// KRShift ...
	KRShift = Key(C.GetKey(1))
	// KLCtrl ...
	KLCtrl = Key(C.GetKey(2))
	// KRCtrl ...
	KRCtrl = Key(C.GetKey(3))
	// KLAlt ...
	KLAlt = Key(C.GetKey(4))
	// KRAlt ...
	KRAlt = Key(C.GetKey(5))
	// KLWin ...
	KLWin = Key(C.GetKey(6))
	// KRWin ...
	KRWin = Key(C.GetKey(7))
	// KTab ...
	KTab = Key(C.GetKey(8))
	// KCapsLock ...
	KCapsLock = Key(C.GetKey(9))
	// KSpace ...
	KSpace = Key(C.GetKey(10))
	// KBackspace ...
	KBackspace = Key(C.GetKey(11))
	// KInsert ...
	KInsert = Key(C.GetKey(12))
	// KSuppr ...
	KSuppr = Key(C.GetKey(13))
	// KHome ...
	KHome = Key(C.GetKey(14))
	// KEnd ...
	KEnd = Key(C.GetKey(15))
	// KPageUp ...
	KPageUp = Key(C.GetKey(16))
	// KPageDown ...
	KPageDown = Key(C.GetKey(17))
	// KUp ...
	KUp = Key(C.GetKey(18))
	// KDown ...
	KDown = Key(C.GetKey(19))
	// KLeft ...
	KLeft = Key(C.GetKey(20))
	// KRight ...
	KRight = Key(C.GetKey(21))
	// KEscape ...
	KEscape = Key(C.GetKey(22))
	// KF1 ...
	KF1 = Key(C.GetKey(23))
	// KF2 ...
	KF2 = Key(C.GetKey(24))
	// KF3 ...
	KF3 = Key(C.GetKey(25))
	// KF4 ...
	KF4 = Key(C.GetKey(26))
	// KF5 ...
	KF5 = Key(C.GetKey(27))
	// KF6 ...
	KF6 = Key(C.GetKey(28))
	// KF7 ...
	KF7 = Key(C.GetKey(29))
	// KF8 ...
	KF8 = Key(C.GetKey(30))
	// KF9 ...
	KF9 = Key(C.GetKey(31))
	// KF10 ...
	KF10 = Key(C.GetKey(32))
	// KF11 ...
	KF11 = Key(C.GetKey(33))
	// KF12 ...
	KF12 = Key(C.GetKey(34))
	// KPrintScreen ...
	KPrintScreen = Key(C.GetKey(35))
	// KScrollLock ...
	KScrollLock = Key(C.GetKey(36))
	// KPause ...
	KPause = Key(C.GetKey(37))
	// KNumLock ...
	KNumLock = Key(C.GetKey(38))
	// KReturn ...
	KReturn = Key(C.GetKey(39))
	// K0 ...
	K0 = Key(C.GetKey(40))
	// K1 ...
	K1 = Key(C.GetKey(41))
	// K2 ...
	K2 = Key(C.GetKey(42))
	// K3 ...
	K3 = Key(C.GetKey(43))
	// K4 ...
	K4 = Key(C.GetKey(44))
	// K5 ...
	K5 = Key(C.GetKey(45))
	// K6 ...
	K6 = Key(C.GetKey(46))
	// K7 ...
	K7 = Key(C.GetKey(47))
	// K8 ...
	K8 = Key(C.GetKey(48))
	// K9 ...
	K9 = Key(C.GetKey(49))
	// KNumpad0 ...
	KNumpad0 = Key(C.GetKey(50))
	// KNumpad1 ...
	KNumpad1 = Key(C.GetKey(51))
	// KNumpad2 ...
	KNumpad2 = Key(C.GetKey(52))
	// KNumpad3 ...
	KNumpad3 = Key(C.GetKey(53))
	// KNumpad4 ...
	KNumpad4 = Key(C.GetKey(54))
	// KNumpad5 ...
	KNumpad5 = Key(C.GetKey(55))
	// KNumpad6 ...
	KNumpad6 = Key(C.GetKey(56))
	// KNumpad7 ...
	KNumpad7 = Key(C.GetKey(57))
	// KNumpad8 ...
	KNumpad8 = Key(C.GetKey(58))
	// KNumpad9 ...
	KNumpad9 = Key(C.GetKey(59))
	// KAdd ...
	KAdd = Key(C.GetKey(60))
	// KSub ...
	KSub = Key(C.GetKey(61))
	// KMul ...
	KMul = Key(C.GetKey(62))
	// KDiv ...
	KDiv = Key(C.GetKey(63))
	// KEnter ...
	KEnter = Key(C.GetKey(64))
	// KA ...
	KA = Key(C.GetKey(65))
	// KB ...
	KB = Key(C.GetKey(66))
	// KC ...
	KC = Key(C.GetKey(67))
	// KD ...
	KD = Key(C.GetKey(68))
	// KE ...
	KE = Key(C.GetKey(69))
	// KF ...
	KF = Key(C.GetKey(70))
	// KG ...
	KG = Key(C.GetKey(71))
	// KH ...
	KH = Key(C.GetKey(72))
	// KI ...
	KI = Key(C.GetKey(73))
	// KJ ...
	KJ = Key(C.GetKey(74))
	// KK ...
	KK = Key(C.GetKey(75))
	// KL ...
	KL = Key(C.GetKey(76))
	// KM ...
	KM = Key(C.GetKey(77))
	// KN ...
	KN = Key(C.GetKey(78))
	// KO ...
	KO = Key(C.GetKey(79))
	// KP ...
	KP = Key(C.GetKey(80))
	// KQ ...
	KQ = Key(C.GetKey(81))
	// KR ...
	KR = Key(C.GetKey(82))
	// KS ...
	KS = Key(C.GetKey(83))
	// KT ...
	KT = Key(C.GetKey(84))
	// KU ...
	KU = Key(C.GetKey(85))
	// KV ...
	KV = Key(C.GetKey(86))
	// KW ...
	KW = Key(C.GetKey(87))
	// KX ...
	KX = Key(C.GetKey(88))
	// KY ...
	KY = Key(C.GetKey(89))
	// KZ ...
	KZ = Key(C.GetKey(90))
	// KPlus ...
	KPlus = Key(C.GetKey(91))
	// KComma ...
	KComma = Key(C.GetKey(92))
	// KMinus ...
	KMinus = Key(C.GetKey(93))
	// KPeriod ...
	KPeriod = Key(C.GetKey(94))
	// KOEM1 ...
	KOEM1 = Key(C.GetKey(95))
	// KOEM2 ...
	KOEM2 = Key(C.GetKey(96))
	// KOEM3 ...
	KOEM3 = Key(C.GetKey(97))
	// KOEM4 ...
	KOEM4 = Key(C.GetKey(98))
	// KOEM5 ...
	KOEM5 = Key(C.GetKey(99))
	// KOEM6 ...
	KOEM6 = Key(C.GetKey(100))
	// KOEM7 ...
	KOEM7 = Key(C.GetKey(101))
	// KOEM8 ...
	KOEM8 = Key(C.GetKey(102))
	// KBrowserBack ...
	KBrowserBack = Key(C.GetKey(103))
	// KBrowserForward ...
	KBrowserForward = Key(C.GetKey(104))
	// KBrowserRefresh ...
	KBrowserRefresh = Key(C.GetKey(105))
	// KBrowserStop ...
	KBrowserStop = Key(C.GetKey(106))
	// KBrowserSearch ...
	KBrowserSearch = Key(C.GetKey(107))
	// KBrowserFavorites ...
	KBrowserFavorites = Key(C.GetKey(108))
	// KBrowserHome ...
	KBrowserHome = Key(C.GetKey(109))
	// KVolumeMute ...
	KVolumeMute = Key(C.GetKey(110))
	// KVolumeDown ...
	KVolumeDown = Key(C.GetKey(111))
	// KVolumeUp ...
	KVolumeUp = Key(C.GetKey(112))
	// KMediaNextTrack ...
	KMediaNextTrack = Key(C.GetKey(113))
	// KMediaPrevTrack ...
	KMediaPrevTrack = Key(C.GetKey(114))
	// KMediaStop ...
	KMediaStop = Key(C.GetKey(115))
	// KMediaPlayPause ...
	KMediaPlayPause = Key(C.GetKey(116))
	// KLaunchMail ...
	KLaunchMail = Key(C.GetKey(117))
	// KLaunchMediaSelect ...
	KLaunchMediaSelect = Key(C.GetKey(118))
	// KLaunchApp1 ...
	KLaunchApp1 = Key(C.GetKey(119))
	// KLaunchApp2 ...
	KLaunchApp2 = Key(C.GetKey(120))
	// KLast ...
	KLast = Key(C.GetKey(121))
)

// GamepadAxes ...
type GamepadAxes int32

var (
	// GALeftX ...
	GALeftX = GamepadAxes(C.GetGamepadAxes(0))
	// GALeftY ...
	GALeftY = GamepadAxes(C.GetGamepadAxes(1))
	// GARightX ...
	GARightX = GamepadAxes(C.GetGamepadAxes(2))
	// GARightY ...
	GARightY = GamepadAxes(C.GetGamepadAxes(3))
	// GALeftTrigger ...
	GALeftTrigger = GamepadAxes(C.GetGamepadAxes(4))
	// GARightTrigger ...
	GARightTrigger = GamepadAxes(C.GetGamepadAxes(5))
	// GACount ...
	GACount = GamepadAxes(C.GetGamepadAxes(6))
)

// GamepadButton ...
type GamepadButton int32

var (
	// GBButtonA ...
	GBButtonA = GamepadButton(C.GetGamepadButton(0))
	// GBButtonB ...
	GBButtonB = GamepadButton(C.GetGamepadButton(1))
	// GBButtonX ...
	GBButtonX = GamepadButton(C.GetGamepadButton(2))
	// GBButtonY ...
	GBButtonY = GamepadButton(C.GetGamepadButton(3))
	// GBLeftBumper ...
	GBLeftBumper = GamepadButton(C.GetGamepadButton(4))
	// GBRightBumper ...
	GBRightBumper = GamepadButton(C.GetGamepadButton(5))
	// GBBack ...
	GBBack = GamepadButton(C.GetGamepadButton(6))
	// GBStart ...
	GBStart = GamepadButton(C.GetGamepadButton(7))
	// GBGuide ...
	GBGuide = GamepadButton(C.GetGamepadButton(8))
	// GBLeftThumb ...
	GBLeftThumb = GamepadButton(C.GetGamepadButton(9))
	// GBRightThumb ...
	GBRightThumb = GamepadButton(C.GetGamepadButton(10))
	// GBDPadUp ...
	GBDPadUp = GamepadButton(C.GetGamepadButton(11))
	// GBDPadRight ...
	GBDPadRight = GamepadButton(C.GetGamepadButton(12))
	// GBDPadDown ...
	GBDPadDown = GamepadButton(C.GetGamepadButton(13))
	// GBDPadLeft ...
	GBDPadLeft = GamepadButton(C.GetGamepadButton(14))
	// GBCount ...
	GBCount = GamepadButton(C.GetGamepadButton(15))
)

// VRControllerButton ...
type VRControllerButton int32

var (
	// VRCBDPadUp ...
	VRCBDPadUp = VRControllerButton(C.GetVRControllerButton(0))
	// VRCBDPadDown ...
	VRCBDPadDown = VRControllerButton(C.GetVRControllerButton(1))
	// VRCBDPadLeft ...
	VRCBDPadLeft = VRControllerButton(C.GetVRControllerButton(2))
	// VRCBDPadRight ...
	VRCBDPadRight = VRControllerButton(C.GetVRControllerButton(3))
	// VRCBSystem ...
	VRCBSystem = VRControllerButton(C.GetVRControllerButton(4))
	// VRCBAppMenu ...
	VRCBAppMenu = VRControllerButton(C.GetVRControllerButton(5))
	// VRCBGrip ...
	VRCBGrip = VRControllerButton(C.GetVRControllerButton(6))
	// VRCBA ...
	VRCBA = VRControllerButton(C.GetVRControllerButton(7))
	// VRCBProximitySensor ...
	VRCBProximitySensor = VRControllerButton(C.GetVRControllerButton(8))
	// VRCBAxis0 ...
	VRCBAxis0 = VRControllerButton(C.GetVRControllerButton(9))
	// VRCBAxis1 ...
	VRCBAxis1 = VRControllerButton(C.GetVRControllerButton(10))
	// VRCBAxis2 ...
	VRCBAxis2 = VRControllerButton(C.GetVRControllerButton(11))
	// VRCBAxis3 ...
	VRCBAxis3 = VRControllerButton(C.GetVRControllerButton(12))
	// VRCBAxis4 ...
	VRCBAxis4 = VRControllerButton(C.GetVRControllerButton(13))
	// VRCBCount ...
	VRCBCount = VRControllerButton(C.GetVRControllerButton(14))
)

// ImGuiWindowFlags ...
type ImGuiWindowFlags int32

var (
	// ImGuiWindowFlagsNoTitleBar ...
	ImGuiWindowFlagsNoTitleBar = ImGuiWindowFlags(C.GetImGuiWindowFlags(0))
	// ImGuiWindowFlagsNoResize ...
	ImGuiWindowFlagsNoResize = ImGuiWindowFlags(C.GetImGuiWindowFlags(1))
	// ImGuiWindowFlagsNoMove ...
	ImGuiWindowFlagsNoMove = ImGuiWindowFlags(C.GetImGuiWindowFlags(2))
	// ImGuiWindowFlagsNoScrollbar ...
	ImGuiWindowFlagsNoScrollbar = ImGuiWindowFlags(C.GetImGuiWindowFlags(3))
	// ImGuiWindowFlagsNoScrollWithMouse ...
	ImGuiWindowFlagsNoScrollWithMouse = ImGuiWindowFlags(C.GetImGuiWindowFlags(4))
	// ImGuiWindowFlagsNoCollapse ...
	ImGuiWindowFlagsNoCollapse = ImGuiWindowFlags(C.GetImGuiWindowFlags(5))
	// ImGuiWindowFlagsAlwaysAutoResize ...
	ImGuiWindowFlagsAlwaysAutoResize = ImGuiWindowFlags(C.GetImGuiWindowFlags(6))
	// ImGuiWindowFlagsNoSavedSettings ...
	ImGuiWindowFlagsNoSavedSettings = ImGuiWindowFlags(C.GetImGuiWindowFlags(7))
	// ImGuiWindowFlagsNoInputs ...
	ImGuiWindowFlagsNoInputs = ImGuiWindowFlags(C.GetImGuiWindowFlags(8))
	// ImGuiWindowFlagsMenuBar ...
	ImGuiWindowFlagsMenuBar = ImGuiWindowFlags(C.GetImGuiWindowFlags(9))
	// ImGuiWindowFlagsHorizontalScrollbar ...
	ImGuiWindowFlagsHorizontalScrollbar = ImGuiWindowFlags(C.GetImGuiWindowFlags(10))
	// ImGuiWindowFlagsNoFocusOnAppearing ...
	ImGuiWindowFlagsNoFocusOnAppearing = ImGuiWindowFlags(C.GetImGuiWindowFlags(11))
	// ImGuiWindowFlagsNoBringToFrontOnFocus ...
	ImGuiWindowFlagsNoBringToFrontOnFocus = ImGuiWindowFlags(C.GetImGuiWindowFlags(12))
	// ImGuiWindowFlagsAlwaysVerticalScrollbar ...
	ImGuiWindowFlagsAlwaysVerticalScrollbar = ImGuiWindowFlags(C.GetImGuiWindowFlags(13))
	// ImGuiWindowFlagsAlwaysHorizontalScrollbar ...
	ImGuiWindowFlagsAlwaysHorizontalScrollbar = ImGuiWindowFlags(C.GetImGuiWindowFlags(14))
	// ImGuiWindowFlagsAlwaysUseWindowPadding ...
	ImGuiWindowFlagsAlwaysUseWindowPadding = ImGuiWindowFlags(C.GetImGuiWindowFlags(15))
	// ImGuiWindowFlagsNoDocking ...
	ImGuiWindowFlagsNoDocking = ImGuiWindowFlags(C.GetImGuiWindowFlags(16))
)

// ImGuiPopupFlags ...
type ImGuiPopupFlags int32

var (
	// ImGuiPopupFlagsNone ...
	ImGuiPopupFlagsNone = ImGuiPopupFlags(C.GetImGuiPopupFlags(0))
	// ImGuiPopupFlagsMouseButtonLeft ...
	ImGuiPopupFlagsMouseButtonLeft = ImGuiPopupFlags(C.GetImGuiPopupFlags(1))
	// ImGuiPopupFlagsMouseButtonRight ...
	ImGuiPopupFlagsMouseButtonRight = ImGuiPopupFlags(C.GetImGuiPopupFlags(2))
	// ImGuiPopupFlagsMouseButtonMiddle ...
	ImGuiPopupFlagsMouseButtonMiddle = ImGuiPopupFlags(C.GetImGuiPopupFlags(3))
	// ImGuiPopupFlagsNoOpenOverExistingPopup ...
	ImGuiPopupFlagsNoOpenOverExistingPopup = ImGuiPopupFlags(C.GetImGuiPopupFlags(4))
	// ImGuiPopupFlagsNoOpenOverItems ...
	ImGuiPopupFlagsNoOpenOverItems = ImGuiPopupFlags(C.GetImGuiPopupFlags(5))
	// ImGuiPopupFlagsAnyPopupId ...
	ImGuiPopupFlagsAnyPopupId = ImGuiPopupFlags(C.GetImGuiPopupFlags(6))
	// ImGuiPopupFlagsAnyPopupLevel ...
	ImGuiPopupFlagsAnyPopupLevel = ImGuiPopupFlags(C.GetImGuiPopupFlags(7))
	// ImGuiPopupFlagsAnyPopup ...
	ImGuiPopupFlagsAnyPopup = ImGuiPopupFlags(C.GetImGuiPopupFlags(8))
)

// ImGuiCond ...
type ImGuiCond int32

var (
	// ImGuiCondAlways ...
	ImGuiCondAlways = ImGuiCond(C.GetImGuiCond(0))
	// ImGuiCondOnce ...
	ImGuiCondOnce = ImGuiCond(C.GetImGuiCond(1))
	// ImGuiCondFirstUseEver ...
	ImGuiCondFirstUseEver = ImGuiCond(C.GetImGuiCond(2))
	// ImGuiCondAppearing ...
	ImGuiCondAppearing = ImGuiCond(C.GetImGuiCond(3))
)

// ImGuiMouseButton ...
type ImGuiMouseButton int32

var (
	// ImGuiMouseButtonLeft ...
	ImGuiMouseButtonLeft = ImGuiMouseButton(C.GetImGuiMouseButton(0))
	// ImGuiMouseButtonRight ...
	ImGuiMouseButtonRight = ImGuiMouseButton(C.GetImGuiMouseButton(1))
	// ImGuiMouseButtonMiddle ...
	ImGuiMouseButtonMiddle = ImGuiMouseButton(C.GetImGuiMouseButton(2))
)

// ImGuiHoveredFlags ...
type ImGuiHoveredFlags int32

var (
	// ImGuiHoveredFlagsNone ...
	ImGuiHoveredFlagsNone = ImGuiHoveredFlags(C.GetImGuiHoveredFlags(0))
	// ImGuiHoveredFlagsChildWindows ...
	ImGuiHoveredFlagsChildWindows = ImGuiHoveredFlags(C.GetImGuiHoveredFlags(1))
	// ImGuiHoveredFlagsRootWindow ...
	ImGuiHoveredFlagsRootWindow = ImGuiHoveredFlags(C.GetImGuiHoveredFlags(2))
	// ImGuiHoveredFlagsAnyWindow ...
	ImGuiHoveredFlagsAnyWindow = ImGuiHoveredFlags(C.GetImGuiHoveredFlags(3))
	// ImGuiHoveredFlagsAllowWhenBlockedByPopup ...
	ImGuiHoveredFlagsAllowWhenBlockedByPopup = ImGuiHoveredFlags(C.GetImGuiHoveredFlags(4))
	// ImGuiHoveredFlagsAllowWhenBlockedByActiveItem ...
	ImGuiHoveredFlagsAllowWhenBlockedByActiveItem = ImGuiHoveredFlags(C.GetImGuiHoveredFlags(5))
	// ImGuiHoveredFlagsAllowWhenOverlapped ...
	ImGuiHoveredFlagsAllowWhenOverlapped = ImGuiHoveredFlags(C.GetImGuiHoveredFlags(6))
	// ImGuiHoveredFlagsAllowWhenDisabled ...
	ImGuiHoveredFlagsAllowWhenDisabled = ImGuiHoveredFlags(C.GetImGuiHoveredFlags(7))
	// ImGuiHoveredFlagsRectOnly ...
	ImGuiHoveredFlagsRectOnly = ImGuiHoveredFlags(C.GetImGuiHoveredFlags(8))
	// ImGuiHoveredFlagsRootAndChildWindows ...
	ImGuiHoveredFlagsRootAndChildWindows = ImGuiHoveredFlags(C.GetImGuiHoveredFlags(9))
)

// ImGuiFocusedFlags ...
type ImGuiFocusedFlags int32

var (
	// ImGuiFocusedFlagsChildWindows ...
	ImGuiFocusedFlagsChildWindows = ImGuiFocusedFlags(C.GetImGuiFocusedFlags(0))
	// ImGuiFocusedFlagsRootWindow ...
	ImGuiFocusedFlagsRootWindow = ImGuiFocusedFlags(C.GetImGuiFocusedFlags(1))
	// ImGuiFocusedFlagsRootAndChildWindows ...
	ImGuiFocusedFlagsRootAndChildWindows = ImGuiFocusedFlags(C.GetImGuiFocusedFlags(2))
)

// ImGuiColorEditFlags ...
type ImGuiColorEditFlags int32

var (
	// ImGuiColorEditFlagsNone ...
	ImGuiColorEditFlagsNone = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(0))
	// ImGuiColorEditFlagsNoAlpha ...
	ImGuiColorEditFlagsNoAlpha = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(1))
	// ImGuiColorEditFlagsNoPicker ...
	ImGuiColorEditFlagsNoPicker = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(2))
	// ImGuiColorEditFlagsNoOptions ...
	ImGuiColorEditFlagsNoOptions = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(3))
	// ImGuiColorEditFlagsNoSmallPreview ...
	ImGuiColorEditFlagsNoSmallPreview = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(4))
	// ImGuiColorEditFlagsNoInputs ...
	ImGuiColorEditFlagsNoInputs = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(5))
	// ImGuiColorEditFlagsNoTooltip ...
	ImGuiColorEditFlagsNoTooltip = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(6))
	// ImGuiColorEditFlagsNoLabel ...
	ImGuiColorEditFlagsNoLabel = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(7))
	// ImGuiColorEditFlagsNoSidePreview ...
	ImGuiColorEditFlagsNoSidePreview = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(8))
	// ImGuiColorEditFlagsNoDragDrop ...
	ImGuiColorEditFlagsNoDragDrop = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(9))
	// ImGuiColorEditFlagsAlphaBar ...
	ImGuiColorEditFlagsAlphaBar = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(10))
	// ImGuiColorEditFlagsAlphaPreview ...
	ImGuiColorEditFlagsAlphaPreview = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(11))
	// ImGuiColorEditFlagsAlphaPreviewHalf ...
	ImGuiColorEditFlagsAlphaPreviewHalf = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(12))
	// ImGuiColorEditFlagsHDR ...
	ImGuiColorEditFlagsHDR = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(13))
	// ImGuiColorEditFlagsDisplayRGB ...
	ImGuiColorEditFlagsDisplayRGB = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(14))
	// ImGuiColorEditFlagsDisplayHSV ...
	ImGuiColorEditFlagsDisplayHSV = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(15))
	// ImGuiColorEditFlagsDisplayHex ...
	ImGuiColorEditFlagsDisplayHex = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(16))
	// ImGuiColorEditFlagsUint8 ...
	ImGuiColorEditFlagsUint8 = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(17))
	// ImGuiColorEditFlagsFloat ...
	ImGuiColorEditFlagsFloat = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(18))
	// ImGuiColorEditFlagsPickerHueBar ...
	ImGuiColorEditFlagsPickerHueBar = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(19))
	// ImGuiColorEditFlagsPickerHueWheel ...
	ImGuiColorEditFlagsPickerHueWheel = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(20))
	// ImGuiColorEditFlagsInputRGB ...
	ImGuiColorEditFlagsInputRGB = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(21))
	// ImGuiColorEditFlagsInputHSV ...
	ImGuiColorEditFlagsInputHSV = ImGuiColorEditFlags(C.GetImGuiColorEditFlags(22))
)

// ImGuiInputTextFlags ...
type ImGuiInputTextFlags int32

var (
	// ImGuiInputTextFlagsCharsDecimal ...
	ImGuiInputTextFlagsCharsDecimal = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(0))
	// ImGuiInputTextFlagsCharsHexadecimal ...
	ImGuiInputTextFlagsCharsHexadecimal = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(1))
	// ImGuiInputTextFlagsCharsUppercase ...
	ImGuiInputTextFlagsCharsUppercase = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(2))
	// ImGuiInputTextFlagsCharsNoBlank ...
	ImGuiInputTextFlagsCharsNoBlank = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(3))
	// ImGuiInputTextFlagsAutoSelectAll ...
	ImGuiInputTextFlagsAutoSelectAll = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(4))
	// ImGuiInputTextFlagsEnterReturnsTrue ...
	ImGuiInputTextFlagsEnterReturnsTrue = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(5))
	// ImGuiInputTextFlagsCallbackCompletion ...
	ImGuiInputTextFlagsCallbackCompletion = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(6))
	// ImGuiInputTextFlagsCallbackHistory ...
	ImGuiInputTextFlagsCallbackHistory = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(7))
	// ImGuiInputTextFlagsCallbackAlways ...
	ImGuiInputTextFlagsCallbackAlways = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(8))
	// ImGuiInputTextFlagsCallbackCharFilter ...
	ImGuiInputTextFlagsCallbackCharFilter = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(9))
	// ImGuiInputTextFlagsAllowTabInput ...
	ImGuiInputTextFlagsAllowTabInput = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(10))
	// ImGuiInputTextFlagsCtrlEnterForNewLine ...
	ImGuiInputTextFlagsCtrlEnterForNewLine = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(11))
	// ImGuiInputTextFlagsNoHorizontalScroll ...
	ImGuiInputTextFlagsNoHorizontalScroll = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(12))
	// ImGuiInputTextFlagsAlwaysOverwrite ...
	ImGuiInputTextFlagsAlwaysOverwrite = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(13))
	// ImGuiInputTextFlagsReadOnly ...
	ImGuiInputTextFlagsReadOnly = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(14))
	// ImGuiInputTextFlagsPassword ...
	ImGuiInputTextFlagsPassword = ImGuiInputTextFlags(C.GetImGuiInputTextFlags(15))
)

// ImGuiTreeNodeFlags ...
type ImGuiTreeNodeFlags int32

var (
	// ImGuiTreeNodeFlagsSelected ...
	ImGuiTreeNodeFlagsSelected = ImGuiTreeNodeFlags(C.GetImGuiTreeNodeFlags(0))
	// ImGuiTreeNodeFlagsFramed ...
	ImGuiTreeNodeFlagsFramed = ImGuiTreeNodeFlags(C.GetImGuiTreeNodeFlags(1))
	// ImGuiTreeNodeFlagsAllowItemOverlap ...
	ImGuiTreeNodeFlagsAllowItemOverlap = ImGuiTreeNodeFlags(C.GetImGuiTreeNodeFlags(2))
	// ImGuiTreeNodeFlagsNoTreePushOnOpen ...
	ImGuiTreeNodeFlagsNoTreePushOnOpen = ImGuiTreeNodeFlags(C.GetImGuiTreeNodeFlags(3))
	// ImGuiTreeNodeFlagsNoAutoOpenOnLog ...
	ImGuiTreeNodeFlagsNoAutoOpenOnLog = ImGuiTreeNodeFlags(C.GetImGuiTreeNodeFlags(4))
	// ImGuiTreeNodeFlagsDefaultOpen ...
	ImGuiTreeNodeFlagsDefaultOpen = ImGuiTreeNodeFlags(C.GetImGuiTreeNodeFlags(5))
	// ImGuiTreeNodeFlagsOpenOnDoubleClick ...
	ImGuiTreeNodeFlagsOpenOnDoubleClick = ImGuiTreeNodeFlags(C.GetImGuiTreeNodeFlags(6))
	// ImGuiTreeNodeFlagsOpenOnArrow ...
	ImGuiTreeNodeFlagsOpenOnArrow = ImGuiTreeNodeFlags(C.GetImGuiTreeNodeFlags(7))
	// ImGuiTreeNodeFlagsLeaf ...
	ImGuiTreeNodeFlagsLeaf = ImGuiTreeNodeFlags(C.GetImGuiTreeNodeFlags(8))
	// ImGuiTreeNodeFlagsBullet ...
	ImGuiTreeNodeFlagsBullet = ImGuiTreeNodeFlags(C.GetImGuiTreeNodeFlags(9))
	// ImGuiTreeNodeFlagsCollapsingHeader ...
	ImGuiTreeNodeFlagsCollapsingHeader = ImGuiTreeNodeFlags(C.GetImGuiTreeNodeFlags(10))
)

// ImGuiSelectableFlags ...
type ImGuiSelectableFlags int32

var (
	// ImGuiSelectableFlagsDontClosePopups ...
	ImGuiSelectableFlagsDontClosePopups = ImGuiSelectableFlags(C.GetImGuiSelectableFlags(0))
	// ImGuiSelectableFlagsSpanAllColumns ...
	ImGuiSelectableFlagsSpanAllColumns = ImGuiSelectableFlags(C.GetImGuiSelectableFlags(1))
	// ImGuiSelectableFlagsAllowDoubleClick ...
	ImGuiSelectableFlagsAllowDoubleClick = ImGuiSelectableFlags(C.GetImGuiSelectableFlags(2))
)

// ImGuiCol ...
type ImGuiCol int32

var (
	// ImGuiColText ...
	ImGuiColText = ImGuiCol(C.GetImGuiCol(0))
	// ImGuiColTextDisabled ...
	ImGuiColTextDisabled = ImGuiCol(C.GetImGuiCol(1))
	// ImGuiColWindowBg ...
	ImGuiColWindowBg = ImGuiCol(C.GetImGuiCol(2))
	// ImGuiColChildBg ...
	ImGuiColChildBg = ImGuiCol(C.GetImGuiCol(3))
	// ImGuiColPopupBg ...
	ImGuiColPopupBg = ImGuiCol(C.GetImGuiCol(4))
	// ImGuiColBorder ...
	ImGuiColBorder = ImGuiCol(C.GetImGuiCol(5))
	// ImGuiColBorderShadow ...
	ImGuiColBorderShadow = ImGuiCol(C.GetImGuiCol(6))
	// ImGuiColFrameBg ...
	ImGuiColFrameBg = ImGuiCol(C.GetImGuiCol(7))
	// ImGuiColFrameBgHovered ...
	ImGuiColFrameBgHovered = ImGuiCol(C.GetImGuiCol(8))
	// ImGuiColFrameBgActive ...
	ImGuiColFrameBgActive = ImGuiCol(C.GetImGuiCol(9))
	// ImGuiColTitleBg ...
	ImGuiColTitleBg = ImGuiCol(C.GetImGuiCol(10))
	// ImGuiColTitleBgActive ...
	ImGuiColTitleBgActive = ImGuiCol(C.GetImGuiCol(11))
	// ImGuiColTitleBgCollapsed ...
	ImGuiColTitleBgCollapsed = ImGuiCol(C.GetImGuiCol(12))
	// ImGuiColMenuBarBg ...
	ImGuiColMenuBarBg = ImGuiCol(C.GetImGuiCol(13))
	// ImGuiColScrollbarBg ...
	ImGuiColScrollbarBg = ImGuiCol(C.GetImGuiCol(14))
	// ImGuiColScrollbarGrab ...
	ImGuiColScrollbarGrab = ImGuiCol(C.GetImGuiCol(15))
	// ImGuiColScrollbarGrabHovered ...
	ImGuiColScrollbarGrabHovered = ImGuiCol(C.GetImGuiCol(16))
	// ImGuiColScrollbarGrabActive ...
	ImGuiColScrollbarGrabActive = ImGuiCol(C.GetImGuiCol(17))
	// ImGuiColCheckMark ...
	ImGuiColCheckMark = ImGuiCol(C.GetImGuiCol(18))
	// ImGuiColSliderGrab ...
	ImGuiColSliderGrab = ImGuiCol(C.GetImGuiCol(19))
	// ImGuiColSliderGrabActive ...
	ImGuiColSliderGrabActive = ImGuiCol(C.GetImGuiCol(20))
	// ImGuiColButton ...
	ImGuiColButton = ImGuiCol(C.GetImGuiCol(21))
	// ImGuiColButtonHovered ...
	ImGuiColButtonHovered = ImGuiCol(C.GetImGuiCol(22))
	// ImGuiColButtonActive ...
	ImGuiColButtonActive = ImGuiCol(C.GetImGuiCol(23))
	// ImGuiColHeader ...
	ImGuiColHeader = ImGuiCol(C.GetImGuiCol(24))
	// ImGuiColHeaderHovered ...
	ImGuiColHeaderHovered = ImGuiCol(C.GetImGuiCol(25))
	// ImGuiColHeaderActive ...
	ImGuiColHeaderActive = ImGuiCol(C.GetImGuiCol(26))
	// ImGuiColSeparator ...
	ImGuiColSeparator = ImGuiCol(C.GetImGuiCol(27))
	// ImGuiColSeparatorHovered ...
	ImGuiColSeparatorHovered = ImGuiCol(C.GetImGuiCol(28))
	// ImGuiColSeparatorActive ...
	ImGuiColSeparatorActive = ImGuiCol(C.GetImGuiCol(29))
	// ImGuiColResizeGrip ...
	ImGuiColResizeGrip = ImGuiCol(C.GetImGuiCol(30))
	// ImGuiColResizeGripHovered ...
	ImGuiColResizeGripHovered = ImGuiCol(C.GetImGuiCol(31))
	// ImGuiColResizeGripActive ...
	ImGuiColResizeGripActive = ImGuiCol(C.GetImGuiCol(32))
	// ImGuiColPlotLines ...
	ImGuiColPlotLines = ImGuiCol(C.GetImGuiCol(33))
	// ImGuiColPlotLinesHovered ...
	ImGuiColPlotLinesHovered = ImGuiCol(C.GetImGuiCol(34))
	// ImGuiColPlotHistogram ...
	ImGuiColPlotHistogram = ImGuiCol(C.GetImGuiCol(35))
	// ImGuiColPlotHistogramHovered ...
	ImGuiColPlotHistogramHovered = ImGuiCol(C.GetImGuiCol(36))
	// ImGuiColTextSelectedBg ...
	ImGuiColTextSelectedBg = ImGuiCol(C.GetImGuiCol(37))
	// ImGuiColDragDropTarget ...
	ImGuiColDragDropTarget = ImGuiCol(C.GetImGuiCol(38))
	// ImGuiColNavHighlight ...
	ImGuiColNavHighlight = ImGuiCol(C.GetImGuiCol(39))
	// ImGuiColNavWindowingHighlight ...
	ImGuiColNavWindowingHighlight = ImGuiCol(C.GetImGuiCol(40))
	// ImGuiColNavWindowingDimBg ...
	ImGuiColNavWindowingDimBg = ImGuiCol(C.GetImGuiCol(41))
	// ImGuiColModalWindowDimBg ...
	ImGuiColModalWindowDimBg = ImGuiCol(C.GetImGuiCol(42))
)

// ImGuiStyleVar ...
type ImGuiStyleVar int32

var (
	// ImGuiStyleVarAlpha ...
	ImGuiStyleVarAlpha = ImGuiStyleVar(C.GetImGuiStyleVar(0))
	// ImGuiStyleVarWindowPadding ...
	ImGuiStyleVarWindowPadding = ImGuiStyleVar(C.GetImGuiStyleVar(1))
	// ImGuiStyleVarWindowRounding ...
	ImGuiStyleVarWindowRounding = ImGuiStyleVar(C.GetImGuiStyleVar(2))
	// ImGuiStyleVarWindowMinSize ...
	ImGuiStyleVarWindowMinSize = ImGuiStyleVar(C.GetImGuiStyleVar(3))
	// ImGuiStyleVarChildRounding ...
	ImGuiStyleVarChildRounding = ImGuiStyleVar(C.GetImGuiStyleVar(4))
	// ImGuiStyleVarFramePadding ...
	ImGuiStyleVarFramePadding = ImGuiStyleVar(C.GetImGuiStyleVar(5))
	// ImGuiStyleVarFrameRounding ...
	ImGuiStyleVarFrameRounding = ImGuiStyleVar(C.GetImGuiStyleVar(6))
	// ImGuiStyleVarItemSpacing ...
	ImGuiStyleVarItemSpacing = ImGuiStyleVar(C.GetImGuiStyleVar(7))
	// ImGuiStyleVarItemInnerSpacing ...
	ImGuiStyleVarItemInnerSpacing = ImGuiStyleVar(C.GetImGuiStyleVar(8))
	// ImGuiStyleVarIndentSpacing ...
	ImGuiStyleVarIndentSpacing = ImGuiStyleVar(C.GetImGuiStyleVar(9))
	// ImGuiStyleVarGrabMinSize ...
	ImGuiStyleVarGrabMinSize = ImGuiStyleVar(C.GetImGuiStyleVar(10))
	// ImGuiStyleVarButtonTextAlign ...
	ImGuiStyleVarButtonTextAlign = ImGuiStyleVar(C.GetImGuiStyleVar(11))
)

// ImDrawFlags ...
type ImDrawFlags int32

var (
	// ImDrawFlagsRoundCornersNone ...
	ImDrawFlagsRoundCornersNone = ImDrawFlags(C.GetImDrawFlags(0))
	// ImDrawFlagsRoundCornersTopLeft ...
	ImDrawFlagsRoundCornersTopLeft = ImDrawFlags(C.GetImDrawFlags(1))
	// ImDrawFlagsRoundCornersTopRight ...
	ImDrawFlagsRoundCornersTopRight = ImDrawFlags(C.GetImDrawFlags(2))
	// ImDrawFlagsRoundCornersBottomLeft ...
	ImDrawFlagsRoundCornersBottomLeft = ImDrawFlags(C.GetImDrawFlags(3))
	// ImDrawFlagsRoundCornersBottomRight ...
	ImDrawFlagsRoundCornersBottomRight = ImDrawFlags(C.GetImDrawFlags(4))
	// ImDrawFlagsRoundCornersAll ...
	ImDrawFlagsRoundCornersAll = ImDrawFlags(C.GetImDrawFlags(5))
)

// ImGuiComboFlags ...
type ImGuiComboFlags int32

var (
	// ImGuiComboFlagsPopupAlignLeft ...
	ImGuiComboFlagsPopupAlignLeft = ImGuiComboFlags(C.GetImGuiComboFlags(0))
	// ImGuiComboFlagsHeightSmall ...
	ImGuiComboFlagsHeightSmall = ImGuiComboFlags(C.GetImGuiComboFlags(1))
	// ImGuiComboFlagsHeightRegular ...
	ImGuiComboFlagsHeightRegular = ImGuiComboFlags(C.GetImGuiComboFlags(2))
	// ImGuiComboFlagsHeightLarge ...
	ImGuiComboFlagsHeightLarge = ImGuiComboFlags(C.GetImGuiComboFlags(3))
	// ImGuiComboFlagsHeightLargest ...
	ImGuiComboFlagsHeightLargest = ImGuiComboFlags(C.GetImGuiComboFlags(4))
)

// AudioFrameFormat ...
type AudioFrameFormat int32

var (
	// AFFLPCM44KHZS16Mono ...
	AFFLPCM44KHZS16Mono = AudioFrameFormat(C.GetAudioFrameFormat(0))
	// AFFLPCM48KHZS16Mono ...
	AFFLPCM48KHZS16Mono = AudioFrameFormat(C.GetAudioFrameFormat(1))
	// AFFLPCM44KHZS16Stereo ...
	AFFLPCM44KHZS16Stereo = AudioFrameFormat(C.GetAudioFrameFormat(2))
	// AFFLPCM48KHZS16Stereo ...
	AFFLPCM48KHZS16Stereo = AudioFrameFormat(C.GetAudioFrameFormat(3))
)

// SourceRepeat ...
type SourceRepeat int32

var (
	// SROnce ...
	SROnce = SourceRepeat(C.GetSourceRepeat(0))
	// SRLoop ...
	SRLoop = SourceRepeat(C.GetSourceRepeat(1))
)

// SourceState ...
type SourceState int32

var (
	// SSInitial ...
	SSInitial = SourceState(C.GetSourceState(0))
	// SSPlaying ...
	SSPlaying = SourceState(C.GetSourceState(1))
	// SSPaused ...
	SSPaused = SourceState(C.GetSourceState(2))
	// SSStopped ...
	SSStopped = SourceState(C.GetSourceState(3))
	// SSInvalid ...
	SSInvalid = SourceState(C.GetSourceState(4))
)

// OpenVRAA ...
type OpenVRAA int32

var (
	// OVRAANone ...
	OVRAANone = OpenVRAA(C.GetOpenVRAA(0))
	// OVRAAMSAA2x ...
	OVRAAMSAA2x = OpenVRAA(C.GetOpenVRAA(1))
	// OVRAAMSAA4x ...
	OVRAAMSAA4x = OpenVRAA(C.GetOpenVRAA(2))
	// OVRAAMSAA8x ...
	OVRAAMSAA8x = OpenVRAA(C.GetOpenVRAA(3))
	// OVRAAMSAA16x ...
	OVRAAMSAA16x = OpenVRAA(C.GetOpenVRAA(4))
)

// VideoFrameFormat ...
type VideoFrameFormat int32

var (
	// VFFUNKNOWN ...
	VFFUNKNOWN = VideoFrameFormat(C.GetVideoFrameFormat(0))
	// VFFYUV422 ...
	VFFYUV422 = VideoFrameFormat(C.GetVideoFrameFormat(1))
	// VFFRGB24 ...
	VFFRGB24 = VideoFrameFormat(C.GetVideoFrameFormat(2))
)

// IntToVoidPointer Cast an integer to a void pointer.  This function is only used to provide access to low-level structures and should not be needed most of the time.
func IntToVoidPointer(ptr uintptr) *VoidPointer {
	ptrToC := C.intptr_t(ptr)
	retval := C.WrapIntToVoidPointer(ptrToC)
	var retvalGO *VoidPointer
	if retval != nil {
		retvalGO = &VoidPointer{h: retval}
	}
	return retvalGO
}

// SetLogLevel Control which log levels should be displayed.  See [harfang.Log], [harfang.Warn], [harfang.Error] and [harfang.Debug].
func SetLogLevel(loglevel LogLevel) {
	loglevelToC := C.int32_t(loglevel)
	C.WrapSetLogLevel(loglevelToC)
}

// SetLogDetailed Display the `details` field of log outputs.
func SetLogDetailed(isdetailed bool) {
	isdetailedToC := C.bool(isdetailed)
	C.WrapSetLogDetailed(isdetailedToC)
}

// Log Output to the engine log.  See [harfang.Log], [harfang.Error], [harfang.Debug] and [harfang.Warn].
func Log(msg string) {
	msgToC, idFinmsgToC := wrapString(msg)
	defer idFinmsgToC()
	C.WrapLog(msgToC)
}

// LogWithDetails Output to the engine log.  See [harfang.Log], [harfang.Error], [harfang.Debug] and [harfang.Warn].
func LogWithDetails(msg string, details string) {
	msgToC, idFinmsgToC := wrapString(msg)
	defer idFinmsgToC()
	detailsToC, idFindetailsToC := wrapString(details)
	defer idFindetailsToC()
	C.WrapLogWithDetails(msgToC, detailsToC)
}

// Warn Output to the engine warning log.  See [harfang.Log], [harfang.Debug] and [harfang.Error].
func Warn(msg string) {
	msgToC, idFinmsgToC := wrapString(msg)
	defer idFinmsgToC()
	C.WrapWarn(msgToC)
}

// WarnWithDetails Output to the engine warning log.  See [harfang.Log], [harfang.Debug] and [harfang.Error].
func WarnWithDetails(msg string, details string) {
	msgToC, idFinmsgToC := wrapString(msg)
	defer idFinmsgToC()
	detailsToC, idFindetailsToC := wrapString(details)
	defer idFindetailsToC()
	C.WrapWarnWithDetails(msgToC, detailsToC)
}

// Error Output to the engine error log.  See [harfang.Log], [harfang.Debug] and [harfang.Warn].
func Error(msg string) {
	msgToC, idFinmsgToC := wrapString(msg)
	defer idFinmsgToC()
	C.WrapError(msgToC)
}

// ErrorWithDetails Output to the engine error log.  See [harfang.Log], [harfang.Debug] and [harfang.Warn].
func ErrorWithDetails(msg string, details string) {
	msgToC, idFinmsgToC := wrapString(msg)
	defer idFinmsgToC()
	detailsToC, idFindetailsToC := wrapString(details)
	defer idFindetailsToC()
	C.WrapErrorWithDetails(msgToC, detailsToC)
}

// Debug Output to the engine debug log.  See [harfang.Log], [harfang.Warn] and [harfang.Error].
func Debug(msg string) {
	msgToC, idFinmsgToC := wrapString(msg)
	defer idFinmsgToC()
	C.WrapDebug(msgToC)
}

// DebugWithDetails Output to the engine debug log.  See [harfang.Log], [harfang.Warn] and [harfang.Error].
func DebugWithDetails(msg string, details string) {
	msgToC, idFinmsgToC := wrapString(msg)
	defer idFinmsgToC()
	detailsToC, idFindetailsToC := wrapString(details)
	defer idFindetailsToC()
	C.WrapDebugWithDetails(msgToC, detailsToC)
}

// TimeToSecF Convert time to fractional seconds.
func TimeToSecF(t int64) float32 {
	tToC := C.int64_t(t)
	retval := C.WrapTimeToSecF(tToC)
	return float32(retval)
}

// TimeToMsF Convert time to miliseconds.
func TimeToMsF(t int64) float32 {
	tToC := C.int64_t(t)
	retval := C.WrapTimeToMsF(tToC)
	return float32(retval)
}

// TimeToUsF Convert time to fractional microseconds.
func TimeToUsF(t int64) float32 {
	tToC := C.int64_t(t)
	retval := C.WrapTimeToUsF(tToC)
	return float32(retval)
}

// TimeToDay Convert time to days.
func TimeToDay(t int64) int64 {
	tToC := C.int64_t(t)
	retval := C.WrapTimeToDay(tToC)
	return int64(retval)
}

// TimeToHour Convert time to hours.
func TimeToHour(t int64) int64 {
	tToC := C.int64_t(t)
	retval := C.WrapTimeToHour(tToC)
	return int64(retval)
}

// TimeToMin Convert time to minutes.
func TimeToMin(t int64) int64 {
	tToC := C.int64_t(t)
	retval := C.WrapTimeToMin(tToC)
	return int64(retval)
}

// TimeToSec Convert time to seconds.
func TimeToSec(t int64) int64 {
	tToC := C.int64_t(t)
	retval := C.WrapTimeToSec(tToC)
	return int64(retval)
}

// TimeToMs Convert time to milliseconds.
func TimeToMs(t int64) int64 {
	tToC := C.int64_t(t)
	retval := C.WrapTimeToMs(tToC)
	return int64(retval)
}

// TimeToUs Convert time to microseconds.
func TimeToUs(t int64) int64 {
	tToC := C.int64_t(t)
	retval := C.WrapTimeToUs(tToC)
	return int64(retval)
}

// TimeToNs Convert time to nanoseconds.
func TimeToNs(t int64) int64 {
	tToC := C.int64_t(t)
	retval := C.WrapTimeToNs(tToC)
	return int64(retval)
}

// TimeFromSecF Convert fractional seconds to time.  See [harfang.man.CoordinateAndUnitSystem].
func TimeFromSecF(sec float32) int64 {
	secToC := C.float(sec)
	retval := C.WrapTimeFromSecF(secToC)
	return int64(retval)
}

// TimeFromMsF Convert milliseconds to time.  See [harfang.man.CoordinateAndUnitSystem].
func TimeFromMsF(ms float32) int64 {
	msToC := C.float(ms)
	retval := C.WrapTimeFromMsF(msToC)
	return int64(retval)
}

// TimeFromUsF Convert fractional microseconds to time.  See [harfang.man.CoordinateAndUnitSystem].
func TimeFromUsF(us float32) int64 {
	usToC := C.float(us)
	retval := C.WrapTimeFromUsF(usToC)
	return int64(retval)
}

// TimeFromDay Convert days to time.  See [harfang.man.CoordinateAndUnitSystem].
func TimeFromDay(day int64) int64 {
	dayToC := C.int64_t(day)
	retval := C.WrapTimeFromDay(dayToC)
	return int64(retval)
}

// TimeFromHour Convert hours to time.  See [harfang.man.CoordinateAndUnitSystem].
func TimeFromHour(hour int64) int64 {
	hourToC := C.int64_t(hour)
	retval := C.WrapTimeFromHour(hourToC)
	return int64(retval)
}

// TimeFromMin Convert minutes to time.  See [harfang.man.CoordinateAndUnitSystem].
func TimeFromMin(min int64) int64 {
	minToC := C.int64_t(min)
	retval := C.WrapTimeFromMin(minToC)
	return int64(retval)
}

// TimeFromSec Convert seconds to time.  See [harfang.man.CoordinateAndUnitSystem].
func TimeFromSec(sec int64) int64 {
	secToC := C.int64_t(sec)
	retval := C.WrapTimeFromSec(secToC)
	return int64(retval)
}

// TimeFromMs Convert milliseconds to time.  See [harfang.man.CoordinateAndUnitSystem].
func TimeFromMs(ms int64) int64 {
	msToC := C.int64_t(ms)
	retval := C.WrapTimeFromMs(msToC)
	return int64(retval)
}

// TimeFromUs Convert microseconds to time.  See [harfang.man.CoordinateAndUnitSystem].
func TimeFromUs(us int64) int64 {
	usToC := C.int64_t(us)
	retval := C.WrapTimeFromUs(usToC)
	return int64(retval)
}

// TimeFromNs Convert nanoseconds to time.  See [harfang.man.CoordinateAndUnitSystem].
func TimeFromNs(ns int64) int64 {
	nsToC := C.int64_t(ns)
	retval := C.WrapTimeFromNs(nsToC)
	return int64(retval)
}

// TimeNow Return the current system time.
func TimeNow() int64 {
	retval := C.WrapTimeNow()
	return int64(retval)
}

// TimeToString Return time as a human-readable string.
func TimeToString(t int64) string {
	tToC := C.int64_t(t)
	retval := C.WrapTimeToString(tToC)
	return C.GoString(retval)
}

// ResetClock Reset the elapsed time counter.
func ResetClock() {
	C.WrapResetClock()
}

// TickClock Advance the engine clock and return the elapsed time since the last call to this function. See [harfang.GetClock] to retrieve the current clock.  See [harfang.GetClockDt].
func TickClock() int64 {
	retval := C.WrapTickClock()
	return int64(retval)
}

// GetClock Return the current clock since the last call to [harfang.TickClock] or [harfang.ResetClock].  See [harfang.time_to_sec_f] to convert the returned time to second.
func GetClock() int64 {
	retval := C.WrapGetClock()
	return int64(retval)
}

// GetClockDt Return the elapsed time recorded during the last call to [harfang.TickClock].
func GetClockDt() int64 {
	retval := C.WrapGetClockDt()
	return int64(retval)
}

// SkipClock Skip elapsed time since the last call to [harfang.TickClock].
func SkipClock() {
	C.WrapSkipClock()
}

// Open Open a file in binary mode.  See [harfang.OpenText], [harfang.OpenWrite], [harfang.OpenWriteText]
func Open(path string) *File {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapOpen(pathToC)
	retvalGO := &File{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *File) {
		C.WrapFileFree(cleanval.h)
	})
	return retvalGO
}

// OpenText Open a file as text. Return a handle to the opened file.  See [harfang.Open], [harfang.OpenWrite], [harfang.OpenWriteText]
func OpenText(path string) *File {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapOpenText(pathToC)
	retvalGO := &File{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *File) {
		C.WrapFileFree(cleanval.h)
	})
	return retvalGO
}

// OpenWrite Open a file as binary in write mode.  See [harfang.Open], [harfang.OpenText], [harfang.OpenWriteText]
func OpenWrite(path string) *File {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapOpenWrite(pathToC)
	retvalGO := &File{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *File) {
		C.WrapFileFree(cleanval.h)
	})
	return retvalGO
}

// OpenWriteText Open a file as text in write mode.  See [harfang.Open], [harfang.OpenText], [harfang.OpenWrite]
func OpenWriteText(path string) *File {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapOpenWriteText(pathToC)
	retvalGO := &File{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *File) {
		C.WrapFileFree(cleanval.h)
	})
	return retvalGO
}

// OpenTemp Return a handle to a temporary file on the local filesystem.
func OpenTemp(templatepath string) *File {
	templatepathToC, idFintemplatepathToC := wrapString(templatepath)
	defer idFintemplatepathToC()
	retval := C.WrapOpenTemp(templatepathToC)
	retvalGO := &File{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *File) {
		C.WrapFileFree(cleanval.h)
	})
	return retvalGO
}

// Close Close a file handle.
func Close(file *File) bool {
	fileToC := file.h
	retval := C.WrapClose(fileToC)
	return bool(retval)
}

// IsValid Test if a resource if valid.
func IsValid(file *File) bool {
	fileToC := file.h
	retval := C.WrapIsValid(fileToC)
	return bool(retval)
}

// IsValidWithT Test if a resource if valid.
func IsValidWithT(t *Texture) bool {
	tToC := t.h
	retval := C.WrapIsValidWithT(tToC)
	return bool(retval)
}

// IsValidWithFb Test if a resource if valid.
func IsValidWithFb(fb *FrameBuffer) bool {
	fbToC := fb.h
	retval := C.WrapIsValidWithFb(fbToC)
	return bool(retval)
}

// IsValidWithPipeline Test if a resource if valid.
func IsValidWithPipeline(pipeline *ForwardPipelineAAA) bool {
	pipelineToC := pipeline.h
	retval := C.WrapIsValidWithPipeline(pipelineToC)
	return bool(retval)
}

// IsValidWithStreamer Test if a resource if valid.
func IsValidWithStreamer(streamer *IVideoStreamer) bool {
	streamerToC := streamer.h
	retval := C.WrapIsValidWithStreamer(streamerToC)
	return bool(retval)
}

// IsEOF Returns `true` if the cursor is at the end of the file, `false` otherwise.
func IsEOF(file *File) bool {
	fileToC := file.h
	retval := C.WrapIsEOF(fileToC)
	return bool(retval)
}

// GetSize Return the size in bytes of a local file.
func GetSize(file *File) int32 {
	fileToC := file.h
	retval := C.WrapGetSize(fileToC)
	return int32(retval)
}

// GetSizeWithRect Return the size in bytes of a local file.
func GetSizeWithRect(rect *Rect) *Vec2 {
	rectToC := rect.h
	retval := C.WrapGetSizeWithRect(rectToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// GetSizeWithIntRectRect Return the size in bytes of a local file.
func GetSizeWithIntRectRect(rect *IntRect) *IVec2 {
	rectToC := rect.h
	retval := C.WrapGetSizeWithIntRectRect(rectToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// Seek Move the handle cursor to a specific position in the file.
func Seek(file *File, offset int64, mode SeekMode) bool {
	fileToC := file.h
	offsetToC := C.int64_t(offset)
	modeToC := C.int32_t(mode)
	retval := C.WrapSeek(fileToC, offsetToC, modeToC)
	return bool(retval)
}

// Tell Return the current handle cursor position in bytes.
func Tell(file *File) int32 {
	fileToC := file.h
	retval := C.WrapTell(fileToC)
	return int32(retval)
}

// Rewind Rewind the read/write cursor of an open file.
func Rewind(file *File) {
	fileToC := file.h
	C.WrapRewind(fileToC)
}

// IsFile Test if a file exists on the local filesystem.
func IsFile(path string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapIsFile(pathToC)
	return bool(retval)
}

// Unlink Remove a file from the local filesystem.
func Unlink(path string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapUnlink(pathToC)
	return bool(retval)
}

// ReadUInt8 Read a binary 8 bit unsigned integer value from a local file.
func ReadUInt8(file *File) uint8 {
	fileToC := file.h
	retval := C.WrapReadUInt8(fileToC)
	return uint8(retval)
}

// ReadUInt16 Read a binary 16 bit unsigned integer value from a local file.
func ReadUInt16(file *File) uint16 {
	fileToC := file.h
	retval := C.WrapReadUInt16(fileToC)
	return uint16(retval)
}

// ReadUInt32 Read a binary 32 bit unsigned integer value from a local file.
func ReadUInt32(file *File) uint32 {
	fileToC := file.h
	retval := C.WrapReadUInt32(fileToC)
	return uint32(retval)
}

// ReadFloat Read a binary 32 bit floating point value from a local file.
func ReadFloat(file *File) float32 {
	fileToC := file.h
	retval := C.WrapReadFloat(fileToC)
	return float32(retval)
}

// WriteUInt8 Write a binary 8 bit unsigned integer to a file.
func WriteUInt8(file *File, value uint8) bool {
	fileToC := file.h
	valueToC := C.uchar(value)
	retval := C.WrapWriteUInt8(fileToC, valueToC)
	return bool(retval)
}

// WriteUInt16 Write a binary 16 bit unsigned integer to a file.
func WriteUInt16(file *File, value uint16) bool {
	fileToC := file.h
	valueToC := C.ushort(value)
	retval := C.WrapWriteUInt16(fileToC, valueToC)
	return bool(retval)
}

// WriteUInt32 Write a binary 32 bit unsigned integer to a file.
func WriteUInt32(file *File, value uint32) bool {
	fileToC := file.h
	valueToC := C.uint32_t(value)
	retval := C.WrapWriteUInt32(fileToC, valueToC)
	return bool(retval)
}

// WriteFloat Write a binary 32 bit floating point value to a file.
func WriteFloat(file *File, value float32) bool {
	fileToC := file.h
	valueToC := C.float(value)
	retval := C.WrapWriteFloat(fileToC, valueToC)
	return bool(retval)
}

// ReadString Read a binary string from a local file.  Strings are stored as a `uint32_t length` field followed by the string content in UTF-8.
func ReadString(file *File) string {
	fileToC := file.h
	retval := C.WrapReadString(fileToC)
	return C.GoString(retval)
}

// WriteString Write a string to a file as 32 bit integer size followed by the string content in UTF8.
func WriteString(file *File, value string) bool {
	fileToC := file.h
	valueToC, idFinvalueToC := wrapString(value)
	defer idFinvalueToC()
	retval := C.WrapWriteString(fileToC, valueToC)
	return bool(retval)
}

// CopyFile Copy a file on the local filesystem.
func CopyFile(src string, dst string) bool {
	srcToC, idFinsrcToC := wrapString(src)
	defer idFinsrcToC()
	dstToC, idFindstToC := wrapString(dst)
	defer idFindstToC()
	retval := C.WrapCopyFile(srcToC, dstToC)
	return bool(retval)
}

// FileToString Return the content of a local filesystem as a string.
func FileToString(path string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapFileToString(pathToC)
	return C.GoString(retval)
}

// StringToFile Return the content of a file on the local filesystem as a string.
func StringToFile(path string, value string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	valueToC, idFinvalueToC := wrapString(value)
	defer idFinvalueToC()
	retval := C.WrapStringToFile(pathToC, valueToC)
	return bool(retval)
}

// LoadDataFromFile ...
func LoadDataFromFile(path string, data *Data) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	dataToC := data.h
	retval := C.WrapLoadDataFromFile(pathToC, dataToC)
	return bool(retval)
}

// SaveDataToFile ...
func SaveDataToFile(path string, data *Data) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	dataToC := data.h
	retval := C.WrapSaveDataToFile(pathToC, dataToC)
	return bool(retval)
}

// ListDir Get the content of a directory on the local filesystem, this function does not recurse into subfolders.  See [harfang.ListDirRecursive].
func ListDir(path string, typeGo DirEntryType) *DirEntryList {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	typeGoToC := C.int32_t(typeGo)
	retval := C.WrapListDir(pathToC, typeGoToC)
	retvalGO := &DirEntryList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *DirEntryList) {
		C.WrapDirEntryListFree(cleanval.h)
	})
	return retvalGO
}

// ListDirRecursive Get the content of a directory on the local filesystem, this function recurses into subfolders.  See [harfang.ListDir].
func ListDirRecursive(path string, typeGo DirEntryType) *DirEntryList {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	typeGoToC := C.int32_t(typeGo)
	retval := C.WrapListDirRecursive(pathToC, typeGoToC)
	retvalGO := &DirEntryList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *DirEntryList) {
		C.WrapDirEntryListFree(cleanval.h)
	})
	return retvalGO
}

// MkDir Create a new directory.  See [harfang.MkTree].
func MkDir(path string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapMkDir(pathToC)
	return bool(retval)
}

// MkDirWithPermissions Create a new directory.  See [harfang.MkTree].
func MkDirWithPermissions(path string, permissions int32) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	permissionsToC := C.int32_t(permissions)
	retval := C.WrapMkDirWithPermissions(pathToC, permissionsToC)
	return bool(retval)
}

// RmDir Remove an empty folder on the local filesystem.  See [harfang.RmTree].
func RmDir(path string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapRmDir(pathToC)
	return bool(retval)
}

// MkTree Create a directory tree on the local filesystem. This function is recursive and creates each missing directory in the path.  See [harfang.MkDir].
func MkTree(path string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapMkTree(pathToC)
	return bool(retval)
}

// MkTreeWithPermissions Create a directory tree on the local filesystem. This function is recursive and creates each missing directory in the path.  See [harfang.MkDir].
func MkTreeWithPermissions(path string, permissions int32) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	permissionsToC := C.int32_t(permissions)
	retval := C.WrapMkTreeWithPermissions(pathToC, permissionsToC)
	return bool(retval)
}

// RmTree Remove a folder on the local filesystem.  **Warning:** This function will through all subfolders and erase all files and folders in the target folder.
func RmTree(path string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapRmTree(pathToC)
	return bool(retval)
}

// Exists Return `true` if a file exists on the local filesystem, `false` otherwise.
func Exists(path string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapExists(pathToC)
	return bool(retval)
}

// IsDir Returns `true` if `path` is a directory on the local filesystem, `false` otherwise.
func IsDir(path string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapIsDir(pathToC)
	return bool(retval)
}

// CopyDir Copy a directory on the local filesystem, this function does not recurse through subdirectories.  See [harfang.CopyDirRecursive].
func CopyDir(src string, dst string) bool {
	srcToC, idFinsrcToC := wrapString(src)
	defer idFinsrcToC()
	dstToC, idFindstToC := wrapString(dst)
	defer idFindstToC()
	retval := C.WrapCopyDir(srcToC, dstToC)
	return bool(retval)
}

// CopyDirRecursive Copy a directory on the local filesystem, recurse through subdirectories.
func CopyDirRecursive(src string, dst string) bool {
	srcToC, idFinsrcToC := wrapString(src)
	defer idFinsrcToC()
	dstToC, idFindstToC := wrapString(dst)
	defer idFindstToC()
	retval := C.WrapCopyDirRecursive(srcToC, dstToC)
	return bool(retval)
}

// IsPathAbsolute Test if the provided path is an absolute or relative path.
func IsPathAbsolute(path string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapIsPathAbsolute(pathToC)
	return bool(retval)
}

// PathToDisplay Format a path for display.
func PathToDisplay(path string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapPathToDisplay(pathToC)
	return C.GoString(retval)
}

// NormalizePath Normalize a path according to the following conventions:  - Replace all whitespaces by underscores.
func NormalizePath(path string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapNormalizePath(pathToC)
	return C.GoString(retval)
}

// FactorizePath Return the input path with all redundant navigation entries stripped (folder separator, `..` and `.` entries).
func FactorizePath(path string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapFactorizePath(pathToC)
	return C.GoString(retval)
}

// CleanPath Cleanup a local filesystem path according to the host platform conventions.  - Remove redundant folder separators. - Remove redundant `.` and `..` folder entries. - Ensure forward slash (`/`) folder separators on Unix and back slash (`\\`) folder separators on Windows.
func CleanPath(path string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapCleanPath(pathToC)
	return C.GoString(retval)
}

// CutFilePath Return the folder navigation part of a file path. The file name and its extension are stripped.  See [harfang.CutFileExtension] and [harfang.CutFileName].
func CutFilePath(path string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapCutFilePath(pathToC)
	return C.GoString(retval)
}

// CutFileName Return the name part of a file path. All folder navigation and extension are stripped.  See [harfang.CutFileExtension] and [harfang.CutFilePath].
func CutFileName(path string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapCutFileName(pathToC)
	return C.GoString(retval)
}

// CutFileExtension Return a file path with its extension stripped.  See [harfang.CutFilePath] and [harfang.CutFileName].
func CutFileExtension(path string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapCutFileExtension(pathToC)
	return C.GoString(retval)
}

// GetFilePath Return the path part of a file path (excluding file name and extension).
func GetFilePath(path string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapGetFilePath(pathToC)
	return C.GoString(retval)
}

// GetFileName Return the name part of a file path (including its extension).
func GetFileName(path string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapGetFileName(pathToC)
	return C.GoString(retval)
}

// GetFileExtension Return the extension part of a file path.
func GetFileExtension(path string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapGetFileExtension(pathToC)
	return C.GoString(retval)
}

// HasFileExtension Test the extension of a file path.
func HasFileExtension(path string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapHasFileExtension(pathToC)
	return bool(retval)
}

// PathStartsWith Test if the provided path starts with the provided prefix.
func PathStartsWith(path string, with string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	withToC, idFinwithToC := wrapString(with)
	defer idFinwithToC()
	retval := C.WrapPathStartsWith(pathToC, withToC)
	return bool(retval)
}

// PathStripPrefix Return a copy of the input path stripped of the provided prefix.
func PathStripPrefix(path string, prefix string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	prefixToC, idFinprefixToC := wrapString(prefix)
	defer idFinprefixToC()
	retval := C.WrapPathStripPrefix(pathToC, prefixToC)
	return C.GoString(retval)
}

// PathStripSuffix Return a copy of the input path stripped of the provided suffix.
func PathStripSuffix(path string, suffix string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	suffixToC, idFinsuffixToC := wrapString(suffix)
	defer idFinsuffixToC()
	retval := C.WrapPathStripSuffix(pathToC, suffixToC)
	return C.GoString(retval)
}

// PathJoin Return a file path from a set of string elements.
func PathJoin(elements *StringList) string {
	elementsToC := elements.h
	retval := C.WrapPathJoin(elementsToC)
	return C.GoString(retval)
}

// SwapFileExtension Return the input file path with its extension replaced.
func SwapFileExtension(path string, ext string) string {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	extToC, idFinextToC := wrapString(ext)
	defer idFinextToC()
	retval := C.WrapSwapFileExtension(pathToC, extToC)
	return C.GoString(retval)
}

// GetCurrentWorkingDirectory Return the system current working directory.
func GetCurrentWorkingDirectory() string {
	retval := C.WrapGetCurrentWorkingDirectory()
	return C.GoString(retval)
}

// GetUserFolder Return the system user folder for the current user.
func GetUserFolder() string {
	retval := C.WrapGetUserFolder()
	return C.GoString(retval)
}

// AddAssetsFolder Mount a local filesystem folder as an assets source.  See [harfang.man.Assets].
func AddAssetsFolder(path string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapAddAssetsFolder(pathToC)
	return bool(retval)
}

// RemoveAssetsFolder Remove a folder from the assets system.  See [harfang.man.Assets].
func RemoveAssetsFolder(path string) {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	C.WrapRemoveAssetsFolder(pathToC)
}

// AddAssetsPackage Mount an archive stored on the local filesystem as an assets source.  See [harfang.man.Assets].
func AddAssetsPackage(path string) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapAddAssetsPackage(pathToC)
	return bool(retval)
}

// RemoveAssetsPackage Remove a package from the assets system.  See [harfang.man.Assets].
func RemoveAssetsPackage(path string) {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	C.WrapRemoveAssetsPackage(pathToC)
}

// IsAssetFile Test if an asset file exists in the assets system.  See [harfang.man.Assets].
func IsAssetFile(name string) bool {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapIsAssetFile(nameToC)
	return bool(retval)
}

// LinearInterpolate Linear interpolate between two values on the [harfang.0;1] interval.  See [harfang.CosineInterpolate], [harfang.CubicInterpolate] and [harfang.HermiteInterpolate].
func LinearInterpolate(y0 float32, y1 float32, t float32) float32 {
	y0ToC := C.float(y0)
	y1ToC := C.float(y1)
	tToC := C.float(t)
	retval := C.WrapLinearInterpolate(y0ToC, y1ToC, tToC)
	return float32(retval)
}

// CosineInterpolate Compute the cosine interpolated value between `y0` and `y1` at `t`.  See [harfang.LinearInterpolate], [harfang.CubicInterpolate] and [harfang.HermiteInterpolate].
func CosineInterpolate(y0 float32, y1 float32, t float32) float32 {
	y0ToC := C.float(y0)
	y1ToC := C.float(y1)
	tToC := C.float(t)
	retval := C.WrapCosineInterpolate(y0ToC, y1ToC, tToC)
	return float32(retval)
}

// CubicInterpolate Perform a cubic interpolation across four values with `t` in the [harfang.0;1] range between `y1` and `y2`.  See [harfang.LinearInterpolate], [harfang.CosineInterpolate] and [harfang.HermiteInterpolate].
func CubicInterpolate(y0 float32, y1 float32, y2 float32, y3 float32, t float32) float32 {
	y0ToC := C.float(y0)
	y1ToC := C.float(y1)
	y2ToC := C.float(y2)
	y3ToC := C.float(y3)
	tToC := C.float(t)
	retval := C.WrapCubicInterpolate(y0ToC, y1ToC, y2ToC, y3ToC, tToC)
	return float32(retval)
}

// CubicInterpolateWithV0V1V2V3 Perform a cubic interpolation across four values with `t` in the [harfang.0;1] range between `y1` and `y2`.  See [harfang.LinearInterpolate], [harfang.CosineInterpolate] and [harfang.HermiteInterpolate].
func CubicInterpolateWithV0V1V2V3(v0 *Vec3, v1 *Vec3, v2 *Vec3, v3 *Vec3, t float32) *Vec3 {
	v0ToC := v0.h
	v1ToC := v1.h
	v2ToC := v2.h
	v3ToC := v3.h
	tToC := C.float(t)
	retval := C.WrapCubicInterpolateWithV0V1V2V3(v0ToC, v1ToC, v2ToC, v3ToC, tToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// HermiteInterpolate Perform a Hermite interpolation across four values with `t` in the [harfang.0;1] range between `y1` and `y2`. The `tension` and `bias` parameters can be used to control the shape of underlying interpolation curve.  See [harfang.LinearInterpolate], [harfang.CosineInterpolate] and [harfang.CubicInterpolate].
func HermiteInterpolate(y0 float32, y1 float32, y2 float32, y3 float32, t float32, tension float32, bias float32) float32 {
	y0ToC := C.float(y0)
	y1ToC := C.float(y1)
	y2ToC := C.float(y2)
	y3ToC := C.float(y3)
	tToC := C.float(t)
	tensionToC := C.float(tension)
	biasToC := C.float(bias)
	retval := C.WrapHermiteInterpolate(y0ToC, y1ToC, y2ToC, y3ToC, tToC, tensionToC, biasToC)
	return float32(retval)
}

// ReverseRotationOrder Return the rotation order processing each axis in the reverse order of the input rotation order.
func ReverseRotationOrder(rotationorder RotationOrder) RotationOrder {
	rotationorderToC := C.uchar(rotationorder)
	retval := C.WrapReverseRotationOrder(rotationorderToC)
	return RotationOrder(retval)
}

// GetArea Return the area of the volume.
func GetArea(minmax *MinMax) float32 {
	minmaxToC := minmax.h
	retval := C.WrapGetArea(minmaxToC)
	return float32(retval)
}

// GetCenter Return the center position of the volume.
func GetCenter(minmax *MinMax) *Vec3 {
	minmaxToC := minmax.h
	retval := C.WrapGetCenter(minmaxToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// ComputeMinMaxBoundingSphere Compute the bounding sphere for the provided axis-aligned bounding box.
func ComputeMinMaxBoundingSphere(minmax *MinMax) (*Vec3, *float32) {
	minmaxToC := minmax.h
	origin := NewVec3()
	originToC := origin.h
	radius := new(float32)
	radiusToC := (*C.float)(unsafe.Pointer(radius))
	C.WrapComputeMinMaxBoundingSphere(minmaxToC, originToC, radiusToC)
	return origin, (*float32)(unsafe.Pointer(radiusToC))
}

// Overlap Return `true` if the provided volume overlaps with this volume, `false` otherwise. The test can optionally be restricted to a specific axis.
func Overlap(minmaxa *MinMax, minmaxb *MinMax) bool {
	minmaxaToC := minmaxa.h
	minmaxbToC := minmaxb.h
	retval := C.WrapOverlap(minmaxaToC, minmaxbToC)
	return bool(retval)
}

// OverlapWithAxis Return `true` if the provided volume overlaps with this volume, `false` otherwise. The test can optionally be restricted to a specific axis.
func OverlapWithAxis(minmaxa *MinMax, minmaxb *MinMax, axis Axis) bool {
	minmaxaToC := minmaxa.h
	minmaxbToC := minmaxb.h
	axisToC := C.uchar(axis)
	retval := C.WrapOverlapWithAxis(minmaxaToC, minmaxbToC, axisToC)
	return bool(retval)
}

// Contains Return `true` if the provided position is inside the bounding volume, `false` otherwise.
func Contains(minmax *MinMax, position *Vec3) bool {
	minmaxToC := minmax.h
	positionToC := position.h
	retval := C.WrapContains(minmaxToC, positionToC)
	return bool(retval)
}

// Union Compute the union of this bounding volume with another volume or a 3d position.
func Union(minmaxa *MinMax, minmaxb *MinMax) *MinMax {
	minmaxaToC := minmaxa.h
	minmaxbToC := minmaxb.h
	retval := C.WrapUnion(minmaxaToC, minmaxbToC)
	retvalGO := &MinMax{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MinMax) {
		C.WrapMinMaxFree(cleanval.h)
	})
	return retvalGO
}

// UnionWithMinmaxPosition Compute the union of this bounding volume with another volume or a 3d position.
func UnionWithMinmaxPosition(minmax *MinMax, position *Vec3) *MinMax {
	minmaxToC := minmax.h
	positionToC := position.h
	retval := C.WrapUnionWithMinmaxPosition(minmaxToC, positionToC)
	retvalGO := &MinMax{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MinMax) {
		C.WrapMinMaxFree(cleanval.h)
	})
	return retvalGO
}

// IntersectRay Intersect an infinite ray with an axis-aligned bounding box, if the first returned value is `true` it is followed by the near and far intersection points.
func IntersectRay(minmax *MinMax, origin *Vec3, direction *Vec3) (bool, *float32, *float32) {
	minmaxToC := minmax.h
	originToC := origin.h
	directionToC := direction.h
	tmin := new(float32)
	tminToC := (*C.float)(unsafe.Pointer(tmin))
	tmax := new(float32)
	tmaxToC := (*C.float)(unsafe.Pointer(tmax))
	retval := C.WrapIntersectRay(minmaxToC, originToC, directionToC, tminToC, tmaxToC)
	return bool(retval), (*float32)(unsafe.Pointer(tminToC)), (*float32)(unsafe.Pointer(tmaxToC))
}

// ClassifyLine Return `true` if the provided line intersect the bounding volume, `false` otherwise.
func ClassifyLine(minmax *MinMax, position *Vec3, direction *Vec3) (bool, *Vec3, *Vec3) {
	minmaxToC := minmax.h
	positionToC := position.h
	directionToC := direction.h
	intersection := NewVec3()
	intersectionToC := intersection.h
	normal := NewVec3()
	normalToC := normal.h
	retval := C.WrapClassifyLine(minmaxToC, positionToC, directionToC, intersectionToC, normalToC)
	return bool(retval), intersection, normal
}

// ClassifySegment Return `true` if the provided segment intersect the bounding volume, `false` otherwise.
func ClassifySegment(minmax *MinMax, p0 *Vec3, p1 *Vec3) (bool, *Vec3, *Vec3) {
	minmaxToC := minmax.h
	p0ToC := p0.h
	p1ToC := p1.h
	intersection := NewVec3()
	intersectionToC := intersection.h
	normal := NewVec3()
	normalToC := normal.h
	retval := C.WrapClassifySegment(minmaxToC, p0ToC, p1ToC, intersectionToC, normalToC)
	return bool(retval), intersection, normal
}

// MinMaxFromPositionSize Set `min = p - size/2` and `max = p + size/2`.
func MinMaxFromPositionSize(position *Vec3, size *Vec3) *MinMax {
	positionToC := position.h
	sizeToC := size.h
	retval := C.WrapMinMaxFromPositionSize(positionToC, sizeToC)
	retvalGO := &MinMax{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MinMax) {
		C.WrapMinMaxFree(cleanval.h)
	})
	return retvalGO
}

// Min Return a vector whose elements are the minimum of each of the two specified vectors.
func Min(a *Vec2, b *Vec2) *Vec2 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapMin(aToC, bToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// MinWithAB Return a vector whose elements are the minimum of each of the two specified vectors.
func MinWithAB(a *IVec2, b *IVec2) *IVec2 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapMinWithAB(aToC, bToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// MinWithVec3AVec3B Return a vector whose elements are the minimum of each of the two specified vectors.
func MinWithVec3AVec3B(a *Vec3, b *Vec3) *Vec3 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapMinWithVec3AVec3B(aToC, bToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// MinWithFloatAFloatB Return a vector whose elements are the minimum of each of the two specified vectors.
func MinWithFloatAFloatB(a float32, b float32) float32 {
	aToC := C.float(a)
	bToC := C.float(b)
	retval := C.WrapMinWithFloatAFloatB(aToC, bToC)
	return float32(retval)
}

// MinWithIntAIntB Return a vector whose elements are the minimum of each of the two specified vectors.
func MinWithIntAIntB(a int32, b int32) int32 {
	aToC := C.int32_t(a)
	bToC := C.int32_t(b)
	retval := C.WrapMinWithIntAIntB(aToC, bToC)
	return int32(retval)
}

// Max Return a vector whose elements are the maximum of each of the two specified vectors.
func Max(a *Vec2, b *Vec2) *Vec2 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapMax(aToC, bToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// MaxWithAB Return a vector whose elements are the maximum of each of the two specified vectors.
func MaxWithAB(a *IVec2, b *IVec2) *IVec2 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapMaxWithAB(aToC, bToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// MaxWithVec3AVec3B Return a vector whose elements are the maximum of each of the two specified vectors.
func MaxWithVec3AVec3B(a *Vec3, b *Vec3) *Vec3 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapMaxWithVec3AVec3B(aToC, bToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// MaxWithFloatAFloatB Return a vector whose elements are the maximum of each of the two specified vectors.
func MaxWithFloatAFloatB(a float32, b float32) float32 {
	aToC := C.float(a)
	bToC := C.float(b)
	retval := C.WrapMaxWithFloatAFloatB(aToC, bToC)
	return float32(retval)
}

// MaxWithIntAIntB Return a vector whose elements are the maximum of each of the two specified vectors.
func MaxWithIntAIntB(a int32, b int32) int32 {
	aToC := C.int32_t(a)
	bToC := C.int32_t(b)
	retval := C.WrapMaxWithIntAIntB(aToC, bToC)
	return int32(retval)
}

// Len2 Return the length of the vector squared.
func Len2(v *Vec2) float32 {
	vToC := v.h
	retval := C.WrapLen2(vToC)
	return float32(retval)
}

// Len2WithV Return the length of the vector squared.
func Len2WithV(v *IVec2) int32 {
	vToC := v.h
	retval := C.WrapLen2WithV(vToC)
	return int32(retval)
}

// Len2WithQ Return the length of the vector squared.
func Len2WithQ(q *Quaternion) float32 {
	qToC := q.h
	retval := C.WrapLen2WithQ(qToC)
	return float32(retval)
}

// Len2WithVec3V Return the length of the vector squared.
func Len2WithVec3V(v *Vec3) float32 {
	vToC := v.h
	retval := C.WrapLen2WithVec3V(vToC)
	return float32(retval)
}

// Len Return the length of the vector.
func Len(v *Vec2) float32 {
	vToC := v.h
	retval := C.WrapLen(vToC)
	return float32(retval)
}

// LenWithV Return the length of the vector.
func LenWithV(v *IVec2) int32 {
	vToC := v.h
	retval := C.WrapLenWithV(vToC)
	return int32(retval)
}

// LenWithQ Return the length of the vector.
func LenWithQ(q *Quaternion) float32 {
	qToC := q.h
	retval := C.WrapLenWithQ(qToC)
	return float32(retval)
}

// LenWithVec3V Return the length of the vector.
func LenWithVec3V(v *Vec3) float32 {
	vToC := v.h
	retval := C.WrapLenWithVec3V(vToC)
	return float32(retval)
}

// Dot Return the dot product of two vectors.
func Dot(a *Vec2, b *Vec2) float32 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapDot(aToC, bToC)
	return float32(retval)
}

// DotWithAB Return the dot product of two vectors.
func DotWithAB(a *IVec2, b *IVec2) int32 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapDotWithAB(aToC, bToC)
	return int32(retval)
}

// DotWithVec3AVec3B Return the dot product of two vectors.
func DotWithVec3AVec3B(a *Vec3, b *Vec3) float32 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapDotWithVec3AVec3B(aToC, bToC)
	return float32(retval)
}

// Normalize Return the input vector scaled so that its length is one.
func Normalize(v *Vec2) *Vec2 {
	vToC := v.h
	retval := C.WrapNormalize(vToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// NormalizeWithV Return the input vector scaled so that its length is one.
func NormalizeWithV(v *IVec2) *IVec2 {
	vToC := v.h
	retval := C.WrapNormalizeWithV(vToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// NormalizeWithVec4V Return the input vector scaled so that its length is one.
func NormalizeWithVec4V(v *Vec4) *Vec4 {
	vToC := v.h
	retval := C.WrapNormalizeWithVec4V(vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// NormalizeWithQ Return the input vector scaled so that its length is one.
func NormalizeWithQ(q *Quaternion) *Quaternion {
	qToC := q.h
	retval := C.WrapNormalizeWithQ(qToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// NormalizeWithM Return the input vector scaled so that its length is one.
func NormalizeWithM(m *Mat3) *Mat3 {
	mToC := m.h
	retval := C.WrapNormalizeWithM(mToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// NormalizeWithVec3V Return the input vector scaled so that its length is one.
func NormalizeWithVec3V(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapNormalizeWithVec3V(vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Reverse Return the provided vector pointing in the opposite direction.
func Reverse(a *Vec2) *Vec2 {
	aToC := a.h
	retval := C.WrapReverse(aToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ReverseWithA Return the provided vector pointing in the opposite direction.
func ReverseWithA(a *IVec2) *IVec2 {
	aToC := a.h
	retval := C.WrapReverseWithA(aToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// ReverseWithV Return the provided vector pointing in the opposite direction.
func ReverseWithV(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapReverseWithV(vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Dist2 Return the squared Euclidean distance between two vectors.
func Dist2(a *Vec2, b *Vec2) float32 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapDist2(aToC, bToC)
	return float32(retval)
}

// Dist2WithAB Return the squared Euclidean distance between two vectors.
func Dist2WithAB(a *IVec2, b *IVec2) int32 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapDist2WithAB(aToC, bToC)
	return int32(retval)
}

// Dist2WithVec3AVec3B Return the squared Euclidean distance between two vectors.
func Dist2WithVec3AVec3B(a *Vec3, b *Vec3) float32 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapDist2WithVec3AVec3B(aToC, bToC)
	return float32(retval)
}

// Dist Return the Euclidean distance between two vectors.
func Dist(a *Vec2, b *Vec2) float32 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapDist(aToC, bToC)
	return float32(retval)
}

// DistWithAB Return the Euclidean distance between two vectors.
func DistWithAB(a *IVec2, b *IVec2) int32 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapDistWithAB(aToC, bToC)
	return int32(retval)
}

// DistWithQuaternionAQuaternionB Return the Euclidean distance between two vectors.
func DistWithQuaternionAQuaternionB(a *Quaternion, b *Quaternion) float32 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapDistWithQuaternionAQuaternionB(aToC, bToC)
	return float32(retval)
}

// DistWithVec3AVec3B Return the Euclidean distance between two vectors.
func DistWithVec3AVec3B(a *Vec3, b *Vec3) float32 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapDistWithVec3AVec3B(aToC, bToC)
	return float32(retval)
}

// Abs Return the absolute value of the function input. For vectors, the absolute value is applied to each component individually and the resulting vector is returned.
func Abs(v *Vec4) *Vec4 {
	vToC := v.h
	retval := C.WrapAbs(vToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// AbsWithV Return the absolute value of the function input. For vectors, the absolute value is applied to each component individually and the resulting vector is returned.
func AbsWithV(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapAbsWithV(vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// AbsWithFloatV Return the absolute value of the function input. For vectors, the absolute value is applied to each component individually and the resulting vector is returned.
func AbsWithFloatV(v float32) float32 {
	vToC := C.float(v)
	retval := C.WrapAbsWithFloatV(vToC)
	return float32(retval)
}

// AbsWithIntV Return the absolute value of the function input. For vectors, the absolute value is applied to each component individually and the resulting vector is returned.
func AbsWithIntV(v int32) int32 {
	vToC := C.int32_t(v)
	retval := C.WrapAbsWithIntV(vToC)
	return int32(retval)
}

// RandomVec4 Return a vector with each component randomized in the inclusive provided range.
func RandomVec4(min float32, max float32) *Vec4 {
	minToC := C.float(min)
	maxToC := C.float(max)
	retval := C.WrapRandomVec4(minToC, maxToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// RandomVec4WithMinMax Return a vector with each component randomized in the inclusive provided range.
func RandomVec4WithMinMax(min *Vec4, max *Vec4) *Vec4 {
	minToC := min.h
	maxToC := max.h
	retval := C.WrapRandomVec4WithMinMax(minToC, maxToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// Inverse Return the inverse of a matrix, vector or quaternion.
func Inverse(q *Quaternion) *Quaternion {
	qToC := q.h
	retval := C.WrapInverse(qToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// InverseWithMI Return the inverse of a matrix, vector or quaternion.
func InverseWithMI(m *Mat3) (bool, *Mat3) {
	mToC := m.h
	I := NewMat3()
	IToC := I.h
	retval := C.WrapInverseWithMI(mToC, IToC)
	return bool(retval), I
}

// InverseWithMat4MMat4I Return the inverse of a matrix, vector or quaternion.
func InverseWithMat4MMat4I(m *Mat4) (bool, *Mat4) {
	mToC := m.h
	I := NewMat4()
	IToC := I.h
	retval := C.WrapInverseWithMat4MMat4I(mToC, IToC)
	return bool(retval), I
}

// InverseWithMResult Return the inverse of a matrix, vector or quaternion.
func InverseWithMResult(m *Mat44) (*Mat44, *bool) {
	mToC := m.h
	result := new(bool)
	resultToC := (*C.bool)(unsafe.Pointer(result))
	retval := C.WrapInverseWithMResult(mToC, resultToC)
	retvalGO := &Mat44{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat44) {
		C.WrapMat44Free(cleanval.h)
	})
	return retvalGO, (*bool)(unsafe.Pointer(resultToC))
}

// InverseWithV Return the inverse of a matrix, vector or quaternion.
func InverseWithV(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapInverseWithV(vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Slerp Interpolate between the rotation represented by two quaternions. The _Spherical Linear Interpolation_ will always take the shortest path between the two rotations.
func Slerp(a *Quaternion, b *Quaternion, t float32) *Quaternion {
	aToC := a.h
	bToC := b.h
	tToC := C.float(t)
	retval := C.WrapSlerp(aToC, bToC, tToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// QuaternionFromEulerWithXYZ Return a quaternion 3d rotation from its _Euler_ vector representation.
func QuaternionFromEulerWithXYZ(x float32, y float32, z float32) *Quaternion {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapQuaternionFromEulerWithXYZ(xToC, yToC, zToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// QuaternionFromEulerWithXYZRotationOrder Return a quaternion 3d rotation from its _Euler_ vector representation.
func QuaternionFromEulerWithXYZRotationOrder(x float32, y float32, z float32, rotationorder RotationOrder) *Quaternion {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	rotationorderToC := C.uchar(rotationorder)
	retval := C.WrapQuaternionFromEulerWithXYZRotationOrder(xToC, yToC, zToC, rotationorderToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// QuaternionFromEuler Return a quaternion 3d rotation from its _Euler_ vector representation.
func QuaternionFromEuler(euler *Vec3) *Quaternion {
	eulerToC := euler.h
	retval := C.WrapQuaternionFromEuler(eulerToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// QuaternionFromEulerWithRotationOrder Return a quaternion 3d rotation from its _Euler_ vector representation.
func QuaternionFromEulerWithRotationOrder(euler *Vec3, rotationorder RotationOrder) *Quaternion {
	eulerToC := euler.h
	rotationorderToC := C.uchar(rotationorder)
	retval := C.WrapQuaternionFromEulerWithRotationOrder(eulerToC, rotationorderToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// QuaternionLookAt Return a quaternion 3d rotation oriented toward the specified position when sitting on the world's origin _{0, 0, 0}_.
func QuaternionLookAt(at *Vec3) *Quaternion {
	atToC := at.h
	retval := C.WrapQuaternionLookAt(atToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// QuaternionFromMatrix3 Return a quaternion rotation from its [harfang.Mat3] representation.
func QuaternionFromMatrix3(m *Mat3) *Quaternion {
	mToC := m.h
	retval := C.WrapQuaternionFromMatrix3(mToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// QuaternionFromAxisAngle Return a quaternion rotation from a 3d axis and a rotation around that axis.
func QuaternionFromAxisAngle(angle float32, axis *Vec3) *Quaternion {
	angleToC := C.float(angle)
	axisToC := axis.h
	retval := C.WrapQuaternionFromAxisAngle(angleToC, axisToC)
	retvalGO := &Quaternion{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Quaternion) {
		C.WrapQuaternionFree(cleanval.h)
	})
	return retvalGO
}

// ToMatrix3 Convert a quaternion rotation to its [harfang.Mat3] representation.
func ToMatrix3(q *Quaternion) *Mat3 {
	qToC := q.h
	retval := C.WrapToMatrix3(qToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// ToEuler Convert a quaternion rotation to its _Euler_ vector representation.
func ToEuler(q *Quaternion) *Vec3 {
	qToC := q.h
	retval := C.WrapToEuler(qToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// ToEulerWithRotationOrder Convert a quaternion rotation to its _Euler_ vector representation.
func ToEulerWithRotationOrder(q *Quaternion, rotationorder RotationOrder) *Vec3 {
	qToC := q.h
	rotationorderToC := C.uchar(rotationorder)
	retval := C.WrapToEulerWithRotationOrder(qToC, rotationorderToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// ToEulerWithM Convert a quaternion rotation to its _Euler_ vector representation.
func ToEulerWithM(m *Mat3) *Vec3 {
	mToC := m.h
	retval := C.WrapToEulerWithM(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// ToEulerWithMRotationOrder Convert a quaternion rotation to its _Euler_ vector representation.
func ToEulerWithMRotationOrder(m *Mat3, rotationorder RotationOrder) *Vec3 {
	mToC := m.h
	rotationorderToC := C.uchar(rotationorder)
	retval := C.WrapToEulerWithMRotationOrder(mToC, rotationorderToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Det Return the determinant of a matrix.
func Det(m *Mat3) float32 {
	mToC := m.h
	retval := C.WrapDet(mToC)
	return float32(retval)
}

// Transpose Return the transpose of the input matrix.  For a pure rotation matrix this returns the opposite transformation so that M*M<sup>T</sup>=I.
func Transpose(m *Mat3) *Mat3 {
	mToC := m.h
	retval := C.WrapTranspose(mToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// GetRow Returns the nth row of a matrix.
func GetRow(m *Mat3, n uint32) *Vec3 {
	mToC := m.h
	nToC := C.uint32_t(n)
	retval := C.WrapGetRow(mToC, nToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetRowWithMN Returns the nth row of a matrix.
func GetRowWithMN(m *Mat4, n uint32) *Vec4 {
	mToC := m.h
	nToC := C.uint32_t(n)
	retval := C.WrapGetRowWithMN(mToC, nToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// GetRowWithMIdx Returns the nth row of a matrix.
func GetRowWithMIdx(m *Mat44, idx uint32) *Vec4 {
	mToC := m.h
	idxToC := C.uint32_t(idx)
	retval := C.WrapGetRowWithMIdx(mToC, idxToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// GetColumn Returns the nth column.
func GetColumn(m *Mat3, n uint32) *Vec3 {
	mToC := m.h
	nToC := C.uint32_t(n)
	retval := C.WrapGetColumn(mToC, nToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetColumnWithMN Returns the nth column.
func GetColumnWithMN(m *Mat4, n uint32) *Vec3 {
	mToC := m.h
	nToC := C.uint32_t(n)
	retval := C.WrapGetColumnWithMN(mToC, nToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetColumnWithMIdx Returns the nth column.
func GetColumnWithMIdx(m *Mat44, idx uint32) *Vec4 {
	mToC := m.h
	idxToC := C.uint32_t(idx)
	retval := C.WrapGetColumnWithMIdx(mToC, idxToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// SetRow Sets the nth row of a matrix.
func SetRow(m *Mat3, n uint32, row *Vec3) {
	mToC := m.h
	nToC := C.uint32_t(n)
	rowToC := row.h
	C.WrapSetRow(mToC, nToC, rowToC)
}

// SetRowWithMNV Sets the nth row of a matrix.
func SetRowWithMNV(m *Mat4, n uint32, v *Vec4) {
	mToC := m.h
	nToC := C.uint32_t(n)
	vToC := v.h
	C.WrapSetRowWithMNV(mToC, nToC, vToC)
}

// SetRowWithMIdxV Sets the nth row of a matrix.
func SetRowWithMIdxV(m *Mat44, idx uint32, v *Vec4) {
	mToC := m.h
	idxToC := C.uint32_t(idx)
	vToC := v.h
	C.WrapSetRowWithMIdxV(mToC, idxToC, vToC)
}

// SetColumn Returns the nth column.
func SetColumn(m *Mat3, n uint32, column *Vec3) {
	mToC := m.h
	nToC := C.uint32_t(n)
	columnToC := column.h
	C.WrapSetColumn(mToC, nToC, columnToC)
}

// SetColumnWithMNV Returns the nth column.
func SetColumnWithMNV(m *Mat4, n uint32, v *Vec3) {
	mToC := m.h
	nToC := C.uint32_t(n)
	vToC := v.h
	C.WrapSetColumnWithMNV(mToC, nToC, vToC)
}

// SetColumnWithMIdxV Returns the nth column.
func SetColumnWithMIdxV(m *Mat44, idx uint32, v *Vec4) {
	mToC := m.h
	idxToC := C.uint32_t(idx)
	vToC := v.h
	C.WrapSetColumnWithMIdxV(mToC, idxToC, vToC)
}

// GetX Return the scaled X axis of a transformation matrix.
func GetX(m *Mat3) *Vec3 {
	mToC := m.h
	retval := C.WrapGetX(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetXWithM Return the scaled X axis of a transformation matrix.
func GetXWithM(m *Mat4) *Vec3 {
	mToC := m.h
	retval := C.WrapGetXWithM(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetXWithRect Return the scaled X axis of a transformation matrix.
func GetXWithRect(rect *Rect) float32 {
	rectToC := rect.h
	retval := C.WrapGetXWithRect(rectToC)
	return float32(retval)
}

// GetXWithIntRectRect Return the scaled X axis of a transformation matrix.
func GetXWithIntRectRect(rect *IntRect) int32 {
	rectToC := rect.h
	retval := C.WrapGetXWithIntRectRect(rectToC)
	return int32(retval)
}

// GetY Return the scaled Y axis of a transformation matrix.
func GetY(m *Mat3) *Vec3 {
	mToC := m.h
	retval := C.WrapGetY(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetYWithM Return the scaled Y axis of a transformation matrix.
func GetYWithM(m *Mat4) *Vec3 {
	mToC := m.h
	retval := C.WrapGetYWithM(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetYWithRect Return the scaled Y axis of a transformation matrix.
func GetYWithRect(rect *Rect) float32 {
	rectToC := rect.h
	retval := C.WrapGetYWithRect(rectToC)
	return float32(retval)
}

// GetYWithIntRectRect Return the scaled Y axis of a transformation matrix.
func GetYWithIntRectRect(rect *IntRect) int32 {
	rectToC := rect.h
	retval := C.WrapGetYWithIntRectRect(rectToC)
	return int32(retval)
}

// GetZ Return the scaled Z axis of a transformation matrix.
func GetZ(m *Mat3) *Vec3 {
	mToC := m.h
	retval := C.WrapGetZ(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetZWithM Return the scaled Z axis of a transformation matrix.
func GetZWithM(m *Mat4) *Vec3 {
	mToC := m.h
	retval := C.WrapGetZWithM(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetTranslation Return the translation part of a tranformation matrix as a translation vector.
func GetTranslation(m *Mat3) *Vec3 {
	mToC := m.h
	retval := C.WrapGetTranslation(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetTranslationWithM Return the translation part of a tranformation matrix as a translation vector.
func GetTranslationWithM(m *Mat4) *Vec3 {
	mToC := m.h
	retval := C.WrapGetTranslationWithM(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetScale Return the scale component of a matrix a scale vector.
func GetScale(m *Mat3) *Vec3 {
	mToC := m.h
	retval := C.WrapGetScale(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetScaleWithM Return the scale component of a matrix a scale vector.
func GetScaleWithM(m *Mat4) *Vec3 {
	mToC := m.h
	retval := C.WrapGetScaleWithM(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// SetX Sets the first row.
func SetX(m *Mat3, X *Vec3) {
	mToC := m.h
	XToC := X.h
	C.WrapSetX(mToC, XToC)
}

// SetXWithM Sets the first row.
func SetXWithM(m *Mat4, X *Vec3) {
	mToC := m.h
	XToC := X.h
	C.WrapSetXWithM(mToC, XToC)
}

// SetXWithRectX Sets the first row.
func SetXWithRectX(rect *Rect, x float32) {
	rectToC := rect.h
	xToC := C.float(x)
	C.WrapSetXWithRectX(rectToC, xToC)
}

// SetXWithIntRectRectIntX Sets the first row.
func SetXWithIntRectRectIntX(rect *IntRect, x int32) {
	rectToC := rect.h
	xToC := C.int32_t(x)
	C.WrapSetXWithIntRectRectIntX(rectToC, xToC)
}

// SetY Sets the second row.
func SetY(m *Mat3, Y *Vec3) {
	mToC := m.h
	YToC := Y.h
	C.WrapSetY(mToC, YToC)
}

// SetYWithM Sets the second row.
func SetYWithM(m *Mat4, Y *Vec3) {
	mToC := m.h
	YToC := Y.h
	C.WrapSetYWithM(mToC, YToC)
}

// SetYWithRectY Sets the second row.
func SetYWithRectY(rect *Rect, y float32) {
	rectToC := rect.h
	yToC := C.float(y)
	C.WrapSetYWithRectY(rectToC, yToC)
}

// SetYWithIntRectRectIntY Sets the second row.
func SetYWithIntRectRectIntY(rect *IntRect, y int32) {
	rectToC := rect.h
	yToC := C.int32_t(y)
	C.WrapSetYWithIntRectRectIntY(rectToC, yToC)
}

// SetZ Sets the third row.
func SetZ(m *Mat3, Z *Vec3) {
	mToC := m.h
	ZToC := Z.h
	C.WrapSetZ(mToC, ZToC)
}

// SetZWithM Sets the third row.
func SetZWithM(m *Mat4, Z *Vec3) {
	mToC := m.h
	ZToC := Z.h
	C.WrapSetZWithM(mToC, ZToC)
}

// SetTranslation Sets the 2D translation part, i.e. the first 2 elements of the last matrix row.
func SetTranslation(m *Mat3, T *Vec3) {
	mToC := m.h
	TToC := T.h
	C.WrapSetTranslation(mToC, TToC)
}

// SetTranslationWithT Sets the 2D translation part, i.e. the first 2 elements of the last matrix row.
func SetTranslationWithT(m *Mat3, T *Vec2) {
	mToC := m.h
	TToC := T.h
	C.WrapSetTranslationWithT(mToC, TToC)
}

// SetTranslationWithM Sets the 2D translation part, i.e. the first 2 elements of the last matrix row.
func SetTranslationWithM(m *Mat4, T *Vec3) {
	mToC := m.h
	TToC := T.h
	C.WrapSetTranslationWithM(mToC, TToC)
}

// SetScale Set the scaling part of the transformation matrix.
func SetScale(m *Mat3, S *Vec3) {
	mToC := m.h
	SToC := S.h
	C.WrapSetScale(mToC, SToC)
}

// SetScaleWithMScale Set the scaling part of the transformation matrix.
func SetScaleWithMScale(m *Mat4, scale *Vec3) {
	mToC := m.h
	scaleToC := scale.h
	C.WrapSetScaleWithMScale(mToC, scaleToC)
}

// SetAxises Inject X, Y and Z axises into a 3x3 matrix.
func SetAxises(m *Mat3, X *Vec3, Y *Vec3, Z *Vec3) {
	mToC := m.h
	XToC := X.h
	YToC := Y.h
	ZToC := Z.h
	C.WrapSetAxises(mToC, XToC, YToC, ZToC)
}

// Orthonormalize Return a matrix where the row vectors form an orthonormal basis. All vectors are normalized and perpendicular to each other.
func Orthonormalize(m *Mat3) *Mat3 {
	mToC := m.h
	retval := C.WrapOrthonormalize(mToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// OrthonormalizeWithM Return a matrix where the row vectors form an orthonormal basis. All vectors are normalized and perpendicular to each other.
func OrthonormalizeWithM(m *Mat4) *Mat4 {
	mToC := m.h
	retval := C.WrapOrthonormalizeWithM(mToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// VectorMat3 Return a vector as a matrix.
func VectorMat3(V *Vec3) *Mat3 {
	VToC := V.h
	retval := C.WrapVectorMat3(VToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// TranslationMat3 Return a 2D translation 3x3 matrix from the first 2 components (__x__,__y__) of the parameter vector.
func TranslationMat3(T *Vec2) *Mat3 {
	TToC := T.h
	retval := C.WrapTranslationMat3(TToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// TranslationMat3WithT Return a 2D translation 3x3 matrix from the first 2 components (__x__,__y__) of the parameter vector.
func TranslationMat3WithT(T *Vec3) *Mat3 {
	TToC := T.h
	retval := C.WrapTranslationMat3WithT(TToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// ScaleMat3 Return a 3x3 scale matrix from a 2D vector.
func ScaleMat3(S *Vec2) *Mat3 {
	SToC := S.h
	retval := C.WrapScaleMat3(SToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// ScaleMat3WithS Return a 3x3 scale matrix from a 2D vector.
func ScaleMat3WithS(S *Vec3) *Mat3 {
	SToC := S.h
	retval := C.WrapScaleMat3WithS(SToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// CrossProductMat3 Creates a matrix __M__ so that __Mv = pâ¨¯v__.  Simply put, multiplying this matrix to any vector __v__ is equivalent to compute the cross product between __p__ and __v__.
func CrossProductMat3(V *Vec3) *Mat3 {
	VToC := V.h
	retval := C.WrapCrossProductMat3(VToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMatX Return a 3x3 rotation matrix around the world X axis {1, 0, 0}.
func RotationMatX(angle float32) *Mat3 {
	angleToC := C.float(angle)
	retval := C.WrapRotationMatX(angleToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMatY Return a 3x3 rotation matrix around the world Y axis {0, 1, 0}.
func RotationMatY(angle float32) *Mat3 {
	angleToC := C.float(angle)
	retval := C.WrapRotationMatY(angleToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMatZ Return a 3x3 rotation matrix around the world Z axis {0, 0, 1}.
func RotationMatZ(angle float32) *Mat3 {
	angleToC := C.float(angle)
	retval := C.WrapRotationMatZ(angleToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMat2D Return a 2D rotation matrix by __a__ radians around the specified __pivot__ point.
func RotationMat2D(angle float32, pivot *Vec2) *Mat3 {
	angleToC := C.float(angle)
	pivotToC := pivot.h
	retval := C.WrapRotationMat2D(angleToC, pivotToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMat3WithXYZ Return a 3x3 rotation matrix.
func RotationMat3WithXYZ(x float32, y float32, z float32) *Mat3 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapRotationMat3WithXYZ(xToC, yToC, zToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMat3WithXYZRotationOrder Return a 3x3 rotation matrix.
func RotationMat3WithXYZRotationOrder(x float32, y float32, z float32, rotationorder RotationOrder) *Mat3 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	rotationorderToC := C.uchar(rotationorder)
	retval := C.WrapRotationMat3WithXYZRotationOrder(xToC, yToC, zToC, rotationorderToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMat3 Return a 3x3 rotation matrix.
func RotationMat3(euler *Vec3) *Mat3 {
	eulerToC := euler.h
	retval := C.WrapRotationMat3(eulerToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMat3WithRotationOrder Return a 3x3 rotation matrix.
func RotationMat3WithRotationOrder(euler *Vec3, rotationorder RotationOrder) *Mat3 {
	eulerToC := euler.h
	rotationorderToC := C.uchar(rotationorder)
	retval := C.WrapRotationMat3WithRotationOrder(eulerToC, rotationorderToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// Mat3LookAt Return a rotation matrix looking down the provided vector. The input vector does not need to be normalized.
func Mat3LookAt(front *Vec3) *Mat3 {
	frontToC := front.h
	retval := C.WrapMat3LookAt(frontToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// Mat3LookAtWithUp Return a rotation matrix looking down the provided vector. The input vector does not need to be normalized.
func Mat3LookAtWithUp(front *Vec3, up *Vec3) *Mat3 {
	frontToC := front.h
	upToC := up.h
	retval := C.WrapMat3LookAtWithUp(frontToC, upToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMatXZY Return a 3x3 rotation matrix around the X axis followed by a rotation around the Z axis then a rotation around the Y axis.
func RotationMatXZY(x float32, y float32, z float32) *Mat3 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapRotationMatXZY(xToC, yToC, zToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMatZYX Return a 3x3 rotation matrix around the Z axis followed by a rotation around the Y axis then a rotation around the X axis.
func RotationMatZYX(x float32, y float32, z float32) *Mat3 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapRotationMatZYX(xToC, yToC, zToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMatXYZ Return a 3x3 rotation matrix around the X axis followed by a rotation around the Y axis then a rotation around the Z axis.
func RotationMatXYZ(x float32, y float32, z float32) *Mat3 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapRotationMatXYZ(xToC, yToC, zToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMatZXY Return a 3x3 rotation matrix around the Z axis followed by a rotation around the X axis then a rotation around the Y axis.
func RotationMatZXY(x float32, y float32, z float32) *Mat3 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapRotationMatZXY(xToC, yToC, zToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMatYZX Return a 3x3 rotation matrix around the Y axis followed by a rotation around the Z axis then a rotation around the X axis.
func RotationMatYZX(x float32, y float32, z float32) *Mat3 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapRotationMatYZX(xToC, yToC, zToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMatYXZ Return a 3x3 rotation matrix around the Y axis followed by a rotation around the X axis then a rotation around the Z axis.
func RotationMatYXZ(x float32, y float32, z float32) *Mat3 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapRotationMatYXZ(xToC, yToC, zToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// RotationMatXY Return a 3x3 rotation matrix around the X axis followed by a rotation around the Y axis.
func RotationMatXY(x float32, y float32) *Mat3 {
	xToC := C.float(x)
	yToC := C.float(y)
	retval := C.WrapRotationMatXY(xToC, yToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// GetT See [harfang.GetTranslation].
func GetT(m *Mat4) *Vec3 {
	mToC := m.h
	retval := C.WrapGetT(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetR See [harfang.GetRotation].
func GetR(m *Mat4) *Vec3 {
	mToC := m.h
	retval := C.WrapGetR(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetRWithRotationOrder See [harfang.GetRotation].
func GetRWithRotationOrder(m *Mat4, rotationorder RotationOrder) *Vec3 {
	mToC := m.h
	rotationorderToC := C.uchar(rotationorder)
	retval := C.WrapGetRWithRotationOrder(mToC, rotationorderToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetRotation Return the rotation component of a transformation matrix as a Euler triplet.
func GetRotation(m *Mat4) *Vec3 {
	mToC := m.h
	retval := C.WrapGetRotation(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetRotationWithRotationOrder Return the rotation component of a transformation matrix as a Euler triplet.
func GetRotationWithRotationOrder(m *Mat4, rotationorder RotationOrder) *Vec3 {
	mToC := m.h
	rotationorderToC := C.uchar(rotationorder)
	retval := C.WrapGetRotationWithRotationOrder(mToC, rotationorderToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// GetRMatrix See [harfang.GetRotationMatrix].
func GetRMatrix(m *Mat4) *Mat3 {
	mToC := m.h
	retval := C.WrapGetRMatrix(mToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// GetRotationMatrix Return the rotation component of a transformation matrix as a [harfang.Mat3] rotation matrix.
func GetRotationMatrix(m *Mat4) *Mat3 {
	mToC := m.h
	retval := C.WrapGetRotationMatrix(mToC)
	retvalGO := &Mat3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat3) {
		C.WrapMat3Free(cleanval.h)
	})
	return retvalGO
}

// GetS See [harfang.GetScale].
func GetS(m *Mat4) *Vec3 {
	mToC := m.h
	retval := C.WrapGetS(mToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// SetT Shortcut for [harfang.SetTranslation].
func SetT(m *Mat4, T *Vec3) {
	mToC := m.h
	TToC := T.h
	C.WrapSetT(mToC, TToC)
}

// SetS Shortcut for [harfang.SetScale].
func SetS(m *Mat4, scale *Vec3) {
	mToC := m.h
	scaleToC := scale.h
	C.WrapSetS(mToC, scaleToC)
}

// InverseFast Compute the inverse of an orthonormal transformation matrix. This function is faster than the generic [harfang.Inverse] function but can only deal with a specific set of matrices.  See [harfang.Inverse].
func InverseFast(m *Mat4) *Mat4 {
	mToC := m.h
	retval := C.WrapInverseFast(mToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// LerpAsOrthonormalBase Linear interpolate between two transformation matrices on the [harfang.0;1] interval.
func LerpAsOrthonormalBase(from *Mat4, to *Mat4, k float32) *Mat4 {
	fromToC := from.h
	toToC := to.h
	kToC := C.float(k)
	retval := C.WrapLerpAsOrthonormalBase(fromToC, toToC, kToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// LerpAsOrthonormalBaseWithFast Linear interpolate between two transformation matrices on the [harfang.0;1] interval.
func LerpAsOrthonormalBaseWithFast(from *Mat4, to *Mat4, k float32, fast bool) *Mat4 {
	fromToC := from.h
	toToC := to.h
	kToC := C.float(k)
	fastToC := C.bool(fast)
	retval := C.WrapLerpAsOrthonormalBaseWithFast(fromToC, toToC, kToC, fastToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Decompose Decompose a transformation matrix into its translation, scaling and rotation components.
func Decompose(m *Mat4) (*Vec3, *Vec3, *Vec3) {
	mToC := m.h
	position := NewVec3()
	positionToC := position.h
	rotation := NewVec3()
	rotationToC := rotation.h
	scale := NewVec3()
	scaleToC := scale.h
	C.WrapDecompose(mToC, positionToC, rotationToC, scaleToC)
	return position, rotation, scale
}

// DecomposeWithRotationOrder Decompose a transformation matrix into its translation, scaling and rotation components.
func DecomposeWithRotationOrder(m *Mat4, rotationorder RotationOrder) (*Vec3, *Vec3, *Vec3) {
	mToC := m.h
	position := NewVec3()
	positionToC := position.h
	rotation := NewVec3()
	rotationToC := rotation.h
	scale := NewVec3()
	scaleToC := scale.h
	rotationorderToC := C.uchar(rotationorder)
	C.WrapDecomposeWithRotationOrder(mToC, positionToC, rotationToC, scaleToC, rotationorderToC)
	return position, rotation, scale
}

// Mat4LookAt Return a _look at_ matrix whose orientation points at the specified position.
func Mat4LookAt(position *Vec3, at *Vec3) *Mat4 {
	positionToC := position.h
	atToC := at.h
	retval := C.WrapMat4LookAt(positionToC, atToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Mat4LookAtWithScale Return a _look at_ matrix whose orientation points at the specified position.
func Mat4LookAtWithScale(position *Vec3, at *Vec3, scale *Vec3) *Mat4 {
	positionToC := position.h
	atToC := at.h
	scaleToC := scale.h
	retval := C.WrapMat4LookAtWithScale(positionToC, atToC, scaleToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Mat4LookAtUp Return a _look at_ matrix whose orientation points at the specified position and up direction.
func Mat4LookAtUp(position *Vec3, at *Vec3, up *Vec3) *Mat4 {
	positionToC := position.h
	atToC := at.h
	upToC := up.h
	retval := C.WrapMat4LookAtUp(positionToC, atToC, upToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Mat4LookAtUpWithScale Return a _look at_ matrix whose orientation points at the specified position and up direction.
func Mat4LookAtUpWithScale(position *Vec3, at *Vec3, up *Vec3, scale *Vec3) *Mat4 {
	positionToC := position.h
	atToC := at.h
	upToC := up.h
	scaleToC := scale.h
	retval := C.WrapMat4LookAtUpWithScale(positionToC, atToC, upToC, scaleToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Mat4LookToward Return a _look at_ matrix whose orientation points toward the specified direction.
func Mat4LookToward(position *Vec3, direction *Vec3) *Mat4 {
	positionToC := position.h
	directionToC := direction.h
	retval := C.WrapMat4LookToward(positionToC, directionToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Mat4LookTowardWithScale Return a _look at_ matrix whose orientation points toward the specified direction.
func Mat4LookTowardWithScale(position *Vec3, direction *Vec3, scale *Vec3) *Mat4 {
	positionToC := position.h
	directionToC := direction.h
	scaleToC := scale.h
	retval := C.WrapMat4LookTowardWithScale(positionToC, directionToC, scaleToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Mat4LookTowardUp Return a _look at_ matrix whose orientation points toward the specified directions.
func Mat4LookTowardUp(position *Vec3, direction *Vec3, up *Vec3) *Mat4 {
	positionToC := position.h
	directionToC := direction.h
	upToC := up.h
	retval := C.WrapMat4LookTowardUp(positionToC, directionToC, upToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// Mat4LookTowardUpWithScale Return a _look at_ matrix whose orientation points toward the specified directions.
func Mat4LookTowardUpWithScale(position *Vec3, direction *Vec3, up *Vec3, scale *Vec3) *Mat4 {
	positionToC := position.h
	directionToC := direction.h
	upToC := up.h
	scaleToC := scale.h
	retval := C.WrapMat4LookTowardUpWithScale(positionToC, directionToC, upToC, scaleToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// TranslationMat4 Return a 4x3 translation matrix from the parameter displacement vector.
func TranslationMat4(t *Vec3) *Mat4 {
	tToC := t.h
	retval := C.WrapTranslationMat4(tToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// RotationMat4 Return a 4x3 rotation matrix from euler angles. The default rotation order is YXZ.
func RotationMat4(euler *Vec3) *Mat4 {
	eulerToC := euler.h
	retval := C.WrapRotationMat4(eulerToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// RotationMat4WithOrder Return a 4x3 rotation matrix from euler angles. The default rotation order is YXZ.
func RotationMat4WithOrder(euler *Vec3, order RotationOrder) *Mat4 {
	eulerToC := euler.h
	orderToC := C.uchar(order)
	retval := C.WrapRotationMat4WithOrder(eulerToC, orderToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// ScaleMat4 Return a 4x3 scale matrix from the parameter scaling vector.
func ScaleMat4(scale *Vec3) *Mat4 {
	scaleToC := scale.h
	retval := C.WrapScaleMat4(scaleToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// ScaleMat4WithScale Return a 4x3 scale matrix from the parameter scaling vector.
func ScaleMat4WithScale(scale float32) *Mat4 {
	scaleToC := C.float(scale)
	retval := C.WrapScaleMat4WithScale(scaleToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// TransformationMat4 Creates a 4x3 transformation matrix from the translation vector __p__, the 3x3 rotation Matrix __m__ (or YXZ euler rotation vector __e__) and the scaling vector __s__.  This is a more efficient version of `TranslationMat4(p) *  ScaleMat4(s) * m`
func TransformationMat4(pos *Vec3, rot *Vec3) *Mat4 {
	posToC := pos.h
	rotToC := rot.h
	retval := C.WrapTransformationMat4(posToC, rotToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// TransformationMat4WithScale Creates a 4x3 transformation matrix from the translation vector __p__, the 3x3 rotation Matrix __m__ (or YXZ euler rotation vector __e__) and the scaling vector __s__.  This is a more efficient version of `TranslationMat4(p) *  ScaleMat4(s) * m`
func TransformationMat4WithScale(pos *Vec3, rot *Vec3, scale *Vec3) *Mat4 {
	posToC := pos.h
	rotToC := rot.h
	scaleToC := scale.h
	retval := C.WrapTransformationMat4WithScale(posToC, rotToC, scaleToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// TransformationMat4WithRot Creates a 4x3 transformation matrix from the translation vector __p__, the 3x3 rotation Matrix __m__ (or YXZ euler rotation vector __e__) and the scaling vector __s__.  This is a more efficient version of `TranslationMat4(p) *  ScaleMat4(s) * m`
func TransformationMat4WithRot(pos *Vec3, rot *Mat3) *Mat4 {
	posToC := pos.h
	rotToC := rot.h
	retval := C.WrapTransformationMat4WithRot(posToC, rotToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// TransformationMat4WithRotScale Creates a 4x3 transformation matrix from the translation vector __p__, the 3x3 rotation Matrix __m__ (or YXZ euler rotation vector __e__) and the scaling vector __s__.  This is a more efficient version of `TranslationMat4(p) *  ScaleMat4(s) * m`
func TransformationMat4WithRotScale(pos *Vec3, rot *Mat3, scale *Vec3) *Mat4 {
	posToC := pos.h
	rotToC := rot.h
	scaleToC := scale.h
	retval := C.WrapTransformationMat4WithRotScale(posToC, rotToC, scaleToC)
	retvalGO := &Mat4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat4) {
		C.WrapMat4Free(cleanval.h)
	})
	return retvalGO
}

// MakeVec3 Make a [harfang.Vec3] from a [harfang.Vec4]. The input vector `w` component is discarded.
func MakeVec3(v *Vec4) *Vec3 {
	vToC := v.h
	retval := C.WrapMakeVec3(vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// RandomVec3 Return a vector with each component randomized in the inclusive provided range.
func RandomVec3(min float32, max float32) *Vec3 {
	minToC := C.float(min)
	maxToC := C.float(max)
	retval := C.WrapRandomVec3(minToC, maxToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// RandomVec3WithMinMax Return a vector with each component randomized in the inclusive provided range.
func RandomVec3WithMinMax(min *Vec3, max *Vec3) *Vec3 {
	minToC := min.h
	maxToC := max.h
	retval := C.WrapRandomVec3WithMinMax(minToC, maxToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// BaseToEuler Compute the Euler angles triplet for the provided `z` direction. The up-vector `y` can be provided to improve coherency of the returned values over time.
func BaseToEuler(z *Vec3) *Vec3 {
	zToC := z.h
	retval := C.WrapBaseToEuler(zToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// BaseToEulerWithY Compute the Euler angles triplet for the provided `z` direction. The up-vector `y` can be provided to improve coherency of the returned values over time.
func BaseToEulerWithY(z *Vec3, y *Vec3) *Vec3 {
	zToC := z.h
	yToC := y.h
	retval := C.WrapBaseToEulerWithY(zToC, yToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Cross Return the cross product of two vectors.
func Cross(a *Vec3, b *Vec3) *Vec3 {
	aToC := a.h
	bToC := b.h
	retval := C.WrapCross(aToC, bToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Clamp Return a vector whose elements are equal to the vector elements clipped to the specified interval.
func Clamp(v *Vec3, min float32, max float32) *Vec3 {
	vToC := v.h
	minToC := C.float(min)
	maxToC := C.float(max)
	retval := C.WrapClamp(vToC, minToC, maxToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// ClampWithMinMax Return a vector whose elements are equal to the vector elements clipped to the specified interval.
func ClampWithMinMax(v *Vec3, min *Vec3, max *Vec3) *Vec3 {
	vToC := v.h
	minToC := min.h
	maxToC := max.h
	retval := C.WrapClampWithMinMax(vToC, minToC, maxToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// ClampWithV Return a vector whose elements are equal to the vector elements clipped to the specified interval.
func ClampWithV(v float32, min float32, max float32) float32 {
	vToC := C.float(v)
	minToC := C.float(min)
	maxToC := C.float(max)
	retval := C.WrapClampWithV(vToC, minToC, maxToC)
	return float32(retval)
}

// ClampWithVMinMax Return a vector whose elements are equal to the vector elements clipped to the specified interval.
func ClampWithVMinMax(v int32, min int32, max int32) int32 {
	vToC := C.int32_t(v)
	minToC := C.int32_t(min)
	maxToC := C.int32_t(max)
	retval := C.WrapClampWithVMinMax(vToC, minToC, maxToC)
	return int32(retval)
}

// ClampWithColor Return a vector whose elements are equal to the vector elements clipped to the specified interval.
func ClampWithColor(color *Color, min float32, max float32) *Color {
	colorToC := color.h
	minToC := C.float(min)
	maxToC := C.float(max)
	retval := C.WrapClampWithColor(colorToC, minToC, maxToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// ClampWithColorMinMax Return a vector whose elements are equal to the vector elements clipped to the specified interval.
func ClampWithColorMinMax(color *Color, min *Color, max *Color) *Color {
	colorToC := color.h
	minToC := min.h
	maxToC := max.h
	retval := C.WrapClampWithColorMinMax(colorToC, minToC, maxToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// ClampLen Returns a vector in the same direction as the specified vector, but with its length clipped by the specified interval.
func ClampLen(v *Vec3, min float32, max float32) *Vec3 {
	vToC := v.h
	minToC := C.float(min)
	maxToC := C.float(max)
	retval := C.WrapClampLen(vToC, minToC, maxToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Sign Returns a vector whose elements are -1 if the corresponding vector element is < 0 and 1 if it's >= 0.
func Sign(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapSign(vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Reflect Return the input vector reflected around the specified normal.
func Reflect(v *Vec3, n *Vec3) *Vec3 {
	vToC := v.h
	nToC := n.h
	retval := C.WrapReflect(vToC, nToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Refract Return the input vector refracted around the provided surface normal.  - `k_in`: IOR of the medium the vector is exiting. - `k_out`: IOR of the medium the vector is entering.
func Refract(v *Vec3, n *Vec3) *Vec3 {
	vToC := v.h
	nToC := n.h
	retval := C.WrapRefract(vToC, nToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// RefractWithKIn Return the input vector refracted around the provided surface normal.  - `k_in`: IOR of the medium the vector is exiting. - `k_out`: IOR of the medium the vector is entering.
func RefractWithKIn(v *Vec3, n *Vec3, kin float32) *Vec3 {
	vToC := v.h
	nToC := n.h
	kinToC := C.float(kin)
	retval := C.WrapRefractWithKIn(vToC, nToC, kinToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// RefractWithKInKOut Return the input vector refracted around the provided surface normal.  - `k_in`: IOR of the medium the vector is exiting. - `k_out`: IOR of the medium the vector is entering.
func RefractWithKInKOut(v *Vec3, n *Vec3, kin float32, kout float32) *Vec3 {
	vToC := v.h
	nToC := n.h
	kinToC := C.float(kin)
	koutToC := C.float(kout)
	retval := C.WrapRefractWithKInKOut(vToC, nToC, kinToC, koutToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Floor Returns a vector whose elements are equal to the nearest integer less than or equal to the vector elements.
func Floor(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapFloor(vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Ceil Returns a vector whose elements are equal to the nearest integer greater than or equal to the vector elements.
func Ceil(v *Vec3) *Vec3 {
	vToC := v.h
	retval := C.WrapCeil(vToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// FaceForward Return the provided vector facing toward the provided direction. If the angle between `v` and `d` is less than 90Â° then `v` is returned unchanged, `v` will be returned reversed otherwise.
func FaceForward(v *Vec3, d *Vec3) *Vec3 {
	vToC := v.h
	dToC := d.h
	retval := C.WrapFaceForward(vToC, dToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Deg3 Convert a triplet of angles in degrees to the engine unit system.  See [harfang.man.CoordinateAndUnitSystem].
func Deg3(x float32, y float32, z float32) *Vec3 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapDeg3(xToC, yToC, zToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Rad3 Convert a triplet of angles in radians to the engine unit system.  See [harfang.man.CoordinateAndUnitSystem].
func Rad3(x float32, y float32, z float32) *Vec3 {
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapRad3(xToC, yToC, zToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Vec3I Create a vector from integer values in the [harfang.0;255] range.
func Vec3I(x int32, y int32, z int32) *Vec3 {
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	zToC := C.int32_t(z)
	retval := C.WrapVec3I(xToC, yToC, zToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// Vec4I Create a vector from integer values in the [harfang.0;255] range.
func Vec4I(x int32, y int32, z int32) *Vec4 {
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	zToC := C.int32_t(z)
	retval := C.WrapVec4I(xToC, yToC, zToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// Vec4IWithW Create a vector from integer values in the [harfang.0;255] range.
func Vec4IWithW(x int32, y int32, z int32, w int32) *Vec4 {
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	zToC := C.int32_t(z)
	wToC := C.int32_t(w)
	retval := C.WrapVec4IWithW(xToC, yToC, zToC, wToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// GetWidth Return the width of a rectangle.
func GetWidth(rect *Rect) float32 {
	rectToC := rect.h
	retval := C.WrapGetWidth(rectToC)
	return float32(retval)
}

// GetWidthWithRect Return the width of a rectangle.
func GetWidthWithRect(rect *IntRect) int32 {
	rectToC := rect.h
	retval := C.WrapGetWidthWithRect(rectToC)
	return int32(retval)
}

// GetHeight Return the height of a rectangle.
func GetHeight(rect *Rect) float32 {
	rectToC := rect.h
	retval := C.WrapGetHeight(rectToC)
	return float32(retval)
}

// GetHeightWithRect Return the height of a rectangle.
func GetHeightWithRect(rect *IntRect) int32 {
	rectToC := rect.h
	retval := C.WrapGetHeightWithRect(rectToC)
	return int32(retval)
}

// SetWidth Set a rectangle width.
func SetWidth(rect *Rect, width float32) {
	rectToC := rect.h
	widthToC := C.float(width)
	C.WrapSetWidth(rectToC, widthToC)
}

// SetWidthWithRectWidth Set a rectangle width.
func SetWidthWithRectWidth(rect *IntRect, width int32) {
	rectToC := rect.h
	widthToC := C.int32_t(width)
	C.WrapSetWidthWithRectWidth(rectToC, widthToC)
}

// SetHeight Set a rectangle height.
func SetHeight(rect *Rect, height float32) {
	rectToC := rect.h
	heightToC := C.float(height)
	C.WrapSetHeight(rectToC, heightToC)
}

// SetHeightWithRectHeight Set a rectangle height.
func SetHeightWithRectHeight(rect *IntRect, height int32) {
	rectToC := rect.h
	heightToC := C.int32_t(height)
	C.WrapSetHeightWithRectHeight(rectToC, heightToC)
}

// Inside Test if a value is inside a containing volume.
func Inside(rect *Rect, v *IVec2) bool {
	rectToC := rect.h
	vToC := v.h
	retval := C.WrapInside(rectToC, vToC)
	return bool(retval)
}

// InsideWithV Test if a value is inside a containing volume.
func InsideWithV(rect *Rect, v *Vec2) bool {
	rectToC := rect.h
	vToC := v.h
	retval := C.WrapInsideWithV(rectToC, vToC)
	return bool(retval)
}

// InsideWithVec3V Test if a value is inside a containing volume.
func InsideWithVec3V(rect *Rect, v *Vec3) bool {
	rectToC := rect.h
	vToC := v.h
	retval := C.WrapInsideWithVec3V(rectToC, vToC)
	return bool(retval)
}

// InsideWithVec4V Test if a value is inside a containing volume.
func InsideWithVec4V(rect *Rect, v *Vec4) bool {
	rectToC := rect.h
	vToC := v.h
	retval := C.WrapInsideWithVec4V(rectToC, vToC)
	return bool(retval)
}

// InsideWithRect Test if a value is inside a containing volume.
func InsideWithRect(rect *IntRect, v *IVec2) bool {
	rectToC := rect.h
	vToC := v.h
	retval := C.WrapInsideWithRect(rectToC, vToC)
	return bool(retval)
}

// InsideWithRectV Test if a value is inside a containing volume.
func InsideWithRectV(rect *IntRect, v *Vec2) bool {
	rectToC := rect.h
	vToC := v.h
	retval := C.WrapInsideWithRectV(rectToC, vToC)
	return bool(retval)
}

// InsideWithIntRectRectVec3V Test if a value is inside a containing volume.
func InsideWithIntRectRectVec3V(rect *IntRect, v *Vec3) bool {
	rectToC := rect.h
	vToC := v.h
	retval := C.WrapInsideWithIntRectRectVec3V(rectToC, vToC)
	return bool(retval)
}

// InsideWithIntRectRectVec4V Test if a value is inside a containing volume.
func InsideWithIntRectRectVec4V(rect *IntRect, v *Vec4) bool {
	rectToC := rect.h
	vToC := v.h
	retval := C.WrapInsideWithIntRectRectVec4V(rectToC, vToC)
	return bool(retval)
}

// FitsInside Return wether `a` fits in `b`.
func FitsInside(a *Rect, b *Rect) bool {
	aToC := a.h
	bToC := b.h
	retval := C.WrapFitsInside(aToC, bToC)
	return bool(retval)
}

// FitsInsideWithAB Return wether `a` fits in `b`.
func FitsInsideWithAB(a *IntRect, b *IntRect) bool {
	aToC := a.h
	bToC := b.h
	retval := C.WrapFitsInsideWithAB(aToC, bToC)
	return bool(retval)
}

// Intersects Return `true` if rect `a` intersects rect `b`.
func Intersects(a *Rect, b *Rect) bool {
	aToC := a.h
	bToC := b.h
	retval := C.WrapIntersects(aToC, bToC)
	return bool(retval)
}

// IntersectsWithAB Return `true` if rect `a` intersects rect `b`.
func IntersectsWithAB(a *IntRect, b *IntRect) bool {
	aToC := a.h
	bToC := b.h
	retval := C.WrapIntersectsWithAB(aToC, bToC)
	return bool(retval)
}

// Intersection Return the intersection of two rectangles.
func Intersection(a *Rect, b *Rect) *Rect {
	aToC := a.h
	bToC := b.h
	retval := C.WrapIntersection(aToC, bToC)
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// IntersectionWithAB Return the intersection of two rectangles.
func IntersectionWithAB(a *IntRect, b *IntRect) *IntRect {
	aToC := a.h
	bToC := b.h
	retval := C.WrapIntersectionWithAB(aToC, bToC)
	retvalGO := &IntRect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IntRect) {
		C.WrapIntRectFree(cleanval.h)
	})
	return retvalGO
}

// Grow Grow a rectangle by the specified amount of units.  See [harfang.Crop].
func Grow(rect *Rect, border float32) *Rect {
	rectToC := rect.h
	borderToC := C.float(border)
	retval := C.WrapGrow(rectToC, borderToC)
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// GrowWithRectBorder Grow a rectangle by the specified amount of units.  See [harfang.Crop].
func GrowWithRectBorder(rect *IntRect, border int32) *IntRect {
	rectToC := rect.h
	borderToC := C.int32_t(border)
	retval := C.WrapGrowWithRectBorder(rectToC, borderToC)
	retvalGO := &IntRect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IntRect) {
		C.WrapIntRectFree(cleanval.h)
	})
	return retvalGO
}

// Offset Offset a rectangle by the specified amount of units.
func Offset(rect *Rect, x float32, y float32) *Rect {
	rectToC := rect.h
	xToC := C.float(x)
	yToC := C.float(y)
	retval := C.WrapOffset(rectToC, xToC, yToC)
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// OffsetWithRectXY Offset a rectangle by the specified amount of units.
func OffsetWithRectXY(rect *IntRect, x int32, y int32) *IntRect {
	rectToC := rect.h
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	retval := C.WrapOffsetWithRectXY(rectToC, xToC, yToC)
	retvalGO := &IntRect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IntRect) {
		C.WrapIntRectFree(cleanval.h)
	})
	return retvalGO
}

// Crop Crop a rectangle. Remove the specified amount of units on each side of the rectangle.  See [harfang.Grow].
func Crop(rect *Rect, left float32, top float32, right float32, bottom float32) *Rect {
	rectToC := rect.h
	leftToC := C.float(left)
	topToC := C.float(top)
	rightToC := C.float(right)
	bottomToC := C.float(bottom)
	retval := C.WrapCrop(rectToC, leftToC, topToC, rightToC, bottomToC)
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// CropWithRectLeftTopRightBottom Crop a rectangle. Remove the specified amount of units on each side of the rectangle.  See [harfang.Grow].
func CropWithRectLeftTopRightBottom(rect *IntRect, left int32, top int32, right int32, bottom int32) *IntRect {
	rectToC := rect.h
	leftToC := C.int32_t(left)
	topToC := C.int32_t(top)
	rightToC := C.int32_t(right)
	bottomToC := C.int32_t(bottom)
	retval := C.WrapCropWithRectLeftTopRightBottom(rectToC, leftToC, topToC, rightToC, bottomToC)
	retvalGO := &IntRect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IntRect) {
		C.WrapIntRectFree(cleanval.h)
	})
	return retvalGO
}

// MakeRectFromWidthHeight Make a rectangle from width and height.
func MakeRectFromWidthHeight(x float32, y float32, w float32, h float32) *Rect {
	xToC := C.float(x)
	yToC := C.float(y)
	wToC := C.float(w)
	hToC := C.float(h)
	retval := C.WrapMakeRectFromWidthHeight(xToC, yToC, wToC, hToC)
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// MakeRectFromWidthHeightWithXYWH Make a rectangle from width and height.
func MakeRectFromWidthHeightWithXYWH(x int32, y int32, w int32, h int32) *IntRect {
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	wToC := C.int32_t(w)
	hToC := C.int32_t(h)
	retval := C.WrapMakeRectFromWidthHeightWithXYWH(xToC, yToC, wToC, hToC)
	retvalGO := &IntRect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IntRect) {
		C.WrapIntRectFree(cleanval.h)
	})
	return retvalGO
}

// ToFloatRect Return an integer rectangle as a floating point rectangle.
func ToFloatRect(rect *IntRect) *Rect {
	rectToC := rect.h
	retval := C.WrapToFloatRect(rectToC)
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// ToIntRect Return a floating point rectangle as an integer rectangle.
func ToIntRect(rect *Rect) *IntRect {
	rectToC := rect.h
	retval := C.WrapToIntRect(rectToC)
	retvalGO := &IntRect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IntRect) {
		C.WrapIntRectFree(cleanval.h)
	})
	return retvalGO
}

// MakePlane Geometrical plane in 3D space.  - `p`: a point lying on the plane. - `n`: the plane normal. - `m`: an affine transformation matrix that will be applied to `p` and `n`.
func MakePlane(p *Vec3, n *Vec3) *Vec4 {
	pToC := p.h
	nToC := n.h
	retval := C.WrapMakePlane(pToC, nToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// MakePlaneWithM Geometrical plane in 3D space.  - `p`: a point lying on the plane. - `n`: the plane normal. - `m`: an affine transformation matrix that will be applied to `p` and `n`.
func MakePlaneWithM(p *Vec3, n *Vec3, m *Mat4) *Vec4 {
	pToC := p.h
	nToC := n.h
	mToC := m.h
	retval := C.WrapMakePlaneWithM(pToC, nToC, mToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// DistanceToPlane Return the signed distance from point __p__ to a plane.  - Distance is positive if __p__ is in front of the plane, meaning that the plane normal is pointing towards __p__. - Distance is negative if __p__ is behind the plane, meaning that the plane normal is pointing away from __p__. - Distance is 0.0 if **p** is lying on the plane.
func DistanceToPlane(plane *Vec4, p *Vec3) float32 {
	planeToC := plane.h
	pToC := p.h
	retval := C.WrapDistanceToPlane(planeToC, pToC)
	return float32(retval)
}

// Wrap Wrap the input value so that it fits in the specified inclusive range.
func Wrap(v float32, start float32, end float32) float32 {
	vToC := C.float(v)
	startToC := C.float(start)
	endToC := C.float(end)
	retval := C.WrapWrap(vToC, startToC, endToC)
	return float32(retval)
}

// WrapWithVStartEnd Wrap the input value so that it fits in the specified inclusive range.
func WrapWithVStartEnd(v int32, start int32, end int32) int32 {
	vToC := C.int32_t(v)
	startToC := C.int32_t(start)
	endToC := C.int32_t(end)
	retval := C.WrapWrapWithVStartEnd(vToC, startToC, endToC)
	return int32(retval)
}

// Lerp See [harfang.LinearInterpolate].
func Lerp(a int32, b int32, t float32) int32 {
	aToC := C.int32_t(a)
	bToC := C.int32_t(b)
	tToC := C.float(t)
	retval := C.WrapLerp(aToC, bToC, tToC)
	return int32(retval)
}

// LerpWithAB See [harfang.LinearInterpolate].
func LerpWithAB(a float32, b float32, t float32) float32 {
	aToC := C.float(a)
	bToC := C.float(b)
	tToC := C.float(t)
	retval := C.WrapLerpWithAB(aToC, bToC, tToC)
	return float32(retval)
}

// LerpWithVec3AVec3B See [harfang.LinearInterpolate].
func LerpWithVec3AVec3B(a *Vec3, b *Vec3, t float32) *Vec3 {
	aToC := a.h
	bToC := b.h
	tToC := C.float(t)
	retval := C.WrapLerpWithVec3AVec3B(aToC, bToC, tToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// LerpWithVec4AVec4B See [harfang.LinearInterpolate].
func LerpWithVec4AVec4B(a *Vec4, b *Vec4, t float32) *Vec4 {
	aToC := a.h
	bToC := b.h
	tToC := C.float(t)
	retval := C.WrapLerpWithVec4AVec4B(aToC, bToC, tToC)
	retvalGO := &Vec4{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec4) {
		C.WrapVec4Free(cleanval.h)
	})
	return retvalGO
}

// Quantize Return the provided value quantized to the specified step.
func Quantize(v float32, q float32) float32 {
	vToC := C.float(v)
	qToC := C.float(q)
	retval := C.WrapQuantize(vToC, qToC)
	return float32(retval)
}

// IsFinite Test if a floating point value is finite.
func IsFinite(v float32) bool {
	vToC := C.float(v)
	retval := C.WrapIsFinite(vToC)
	return bool(retval)
}

// Deg Convert an angle in degrees to the engine unit system.  See [harfang.man.CoordinateAndUnitSystem].
func Deg(degrees float32) float32 {
	degreesToC := C.float(degrees)
	retval := C.WrapDeg(degreesToC)
	return float32(retval)
}

// Rad Convert an angle in radians to the engine unit system.  See [harfang.man.CoordinateAndUnitSystem].
func Rad(radians float32) float32 {
	radiansToC := C.float(radians)
	retval := C.WrapRad(radiansToC)
	return float32(retval)
}

// DegreeToRadian Convert an angle in degrees to radians.
func DegreeToRadian(degrees float32) float32 {
	degreesToC := C.float(degrees)
	retval := C.WrapDegreeToRadian(degreesToC)
	return float32(retval)
}

// RadianToDegree Convert an angle in radians to degrees.
func RadianToDegree(radians float32) float32 {
	radiansToC := C.float(radians)
	retval := C.WrapRadianToDegree(radiansToC)
	return float32(retval)
}

// Sec Convert a value in seconds to the Harfang internal unit system.  See [harfang.man.CoordinateAndUnitSystem].
func Sec(seconds float32) float32 {
	secondsToC := C.float(seconds)
	retval := C.WrapSec(secondsToC)
	return float32(retval)
}

// Ms Convert a value in milliseconds to the Harfang internal unit system.  See [harfang.man.CoordinateAndUnitSystem].
func Ms(milliseconds float32) float32 {
	millisecondsToC := C.float(milliseconds)
	retval := C.WrapMs(millisecondsToC)
	return float32(retval)
}

// Km Convert a value in kilometers to the Harfang internal unit system.  See [harfang.man.CoordinateAndUnitSystem].
func Km(km float32) float32 {
	kmToC := C.float(km)
	retval := C.WrapKm(kmToC)
	return float32(retval)
}

// Mtr Convert a value in meters to the Harfang internal unit system.  See [harfang.man.CoordinateAndUnitSystem].
func Mtr(m float32) float32 {
	mToC := C.float(m)
	retval := C.WrapMtr(mToC)
	return float32(retval)
}

// Cm Convert a value in centimeters to the Harfang internal unit system.  See [harfang.man.CoordinateAndUnitSystem].
func Cm(cm float32) float32 {
	cmToC := C.float(cm)
	retval := C.WrapCm(cmToC)
	return float32(retval)
}

// Mm Convert a value in millimeters to the Harfang internal unit system.  See [harfang.man.CoordinateAndUnitSystem].
func Mm(mm float32) float32 {
	mmToC := C.float(mm)
	retval := C.WrapMm(mmToC)
	return float32(retval)
}

// Inch Convert a value in inches to the Harfang internal unit system.  See [harfang.man.CoordinateAndUnitSystem].
func Inch(inch float32) float32 {
	inchToC := C.float(inch)
	retval := C.WrapInch(inchToC)
	return float32(retval)
}

// Seed Set the starting seed of the pseudo-random number generator.
func Seed(seed uint32) {
	seedToC := C.uint32_t(seed)
	C.WrapSeed(seedToC)
}

// Rand Return a random integer value in the provided range, default range is [harfang.0;65535].  See [harfang.FRand] to generate a random floating point value.
func Rand() uint32 {
	retval := C.WrapRand()
	return uint32(retval)
}

// RandWithRange Return a random integer value in the provided range, default range is [harfang.0;65535].  See [harfang.FRand] to generate a random floating point value.
func RandWithRange(rangeGo uint32) uint32 {
	rangeGoToC := C.uint32_t(rangeGo)
	retval := C.WrapRandWithRange(rangeGoToC)
	return uint32(retval)
}

// FRand Return a random floating point value in the provided range, default range is [harfang.0;1].  See [harfang.Rand] to generate a random integer value.
func FRand() float32 {
	retval := C.WrapFRand()
	return float32(retval)
}

// FRandWithRange Return a random floating point value in the provided range, default range is [harfang.0;1].  See [harfang.Rand] to generate a random integer value.
func FRandWithRange(rangeGo float32) float32 {
	rangeGoToC := C.float(rangeGo)
	retval := C.WrapFRandWithRange(rangeGoToC)
	return float32(retval)
}

// FRRand Return a random floating point value in the provided range, default range is [harfang.-1;1].
func FRRand() float32 {
	retval := C.WrapFRRand()
	return float32(retval)
}

// FRRandWithRangeStart Return a random floating point value in the provided range, default range is [harfang.-1;1].
func FRRandWithRangeStart(rangestart float32) float32 {
	rangestartToC := C.float(rangestart)
	retval := C.WrapFRRandWithRangeStart(rangestartToC)
	return float32(retval)
}

// FRRandWithRangeStartRangeEnd Return a random floating point value in the provided range, default range is [harfang.-1;1].
func FRRandWithRangeStartRangeEnd(rangestart float32, rangeend float32) float32 {
	rangestartToC := C.float(rangestart)
	rangeendToC := C.float(rangeend)
	retval := C.WrapFRRandWithRangeStartRangeEnd(rangestartToC, rangeendToC)
	return float32(retval)
}

// ZoomFactorToFov Convert from a zoom factor value in meters to a fov value in radian.
func ZoomFactorToFov(zoomfactor float32) float32 {
	zoomfactorToC := C.float(zoomfactor)
	retval := C.WrapZoomFactorToFov(zoomfactorToC)
	return float32(retval)
}

// FovToZoomFactor Convert from a fov value in radian to a zoom factor value in meters.
func FovToZoomFactor(fov float32) float32 {
	fovToC := C.float(fov)
	retval := C.WrapFovToZoomFactor(fovToC)
	return float32(retval)
}

// ComputeOrthographicProjectionMatrix Compute an orthographic projection matrix.  An orthographic projection has no perspective and all lines parrallel in 3d space will still appear parrallel on screen after projection using the returned matrix.  The `size` parameter controls the extends of the projected view. When projecting a 3d world this parameter is expressed in meters. Use the `aspect_ratio` parameter to prevent distortion from induced by non-square viewport.  See [harfang.ComputeAspectRatioX] or [harfang.ComputeAspectRatioY] to compute an aspect ratio factor in paysage or portrait mode.
func ComputeOrthographicProjectionMatrix(znear float32, zfar float32, size float32, aspectratio *Vec2) *Mat44 {
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	sizeToC := C.float(size)
	aspectratioToC := aspectratio.h
	retval := C.WrapComputeOrthographicProjectionMatrix(znearToC, zfarToC, sizeToC, aspectratioToC)
	retvalGO := &Mat44{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat44) {
		C.WrapMat44Free(cleanval.h)
	})
	return retvalGO
}

// ComputeOrthographicProjectionMatrixWithOffset Compute an orthographic projection matrix.  An orthographic projection has no perspective and all lines parrallel in 3d space will still appear parrallel on screen after projection using the returned matrix.  The `size` parameter controls the extends of the projected view. When projecting a 3d world this parameter is expressed in meters. Use the `aspect_ratio` parameter to prevent distortion from induced by non-square viewport.  See [harfang.ComputeAspectRatioX] or [harfang.ComputeAspectRatioY] to compute an aspect ratio factor in paysage or portrait mode.
func ComputeOrthographicProjectionMatrixWithOffset(znear float32, zfar float32, size float32, aspectratio *Vec2, offset *Vec2) *Mat44 {
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	sizeToC := C.float(size)
	aspectratioToC := aspectratio.h
	offsetToC := offset.h
	retval := C.WrapComputeOrthographicProjectionMatrixWithOffset(znearToC, zfarToC, sizeToC, aspectratioToC, offsetToC)
	retvalGO := &Mat44{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat44) {
		C.WrapMat44Free(cleanval.h)
	})
	return retvalGO
}

// ComputePerspectiveProjectionMatrix Compute a perspective projection matrix, , `fov` is the field of view angle, see [harfang.Deg] and [harfang.Rad].  See [harfang.ZoomFactorToFov], [harfang.FovToZoomFactor], [harfang.ComputeAspectRatioX] and [harfang.ComputeAspectRatioY].
func ComputePerspectiveProjectionMatrix(znear float32, zfar float32, zoomfactor float32, aspectratio *Vec2) *Mat44 {
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	zoomfactorToC := C.float(zoomfactor)
	aspectratioToC := aspectratio.h
	retval := C.WrapComputePerspectiveProjectionMatrix(znearToC, zfarToC, zoomfactorToC, aspectratioToC)
	retvalGO := &Mat44{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat44) {
		C.WrapMat44Free(cleanval.h)
	})
	return retvalGO
}

// ComputePerspectiveProjectionMatrixWithOffset Compute a perspective projection matrix, , `fov` is the field of view angle, see [harfang.Deg] and [harfang.Rad].  See [harfang.ZoomFactorToFov], [harfang.FovToZoomFactor], [harfang.ComputeAspectRatioX] and [harfang.ComputeAspectRatioY].
func ComputePerspectiveProjectionMatrixWithOffset(znear float32, zfar float32, zoomfactor float32, aspectratio *Vec2, offset *Vec2) *Mat44 {
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	zoomfactorToC := C.float(zoomfactor)
	aspectratioToC := aspectratio.h
	offsetToC := offset.h
	retval := C.WrapComputePerspectiveProjectionMatrixWithOffset(znearToC, zfarToC, zoomfactorToC, aspectratioToC, offsetToC)
	retvalGO := &Mat44{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat44) {
		C.WrapMat44Free(cleanval.h)
	})
	return retvalGO
}

// ComputeAspectRatioX Compute the aspect ratio factor for the provided viewport dimensions. Use this method to compute aspect ratio for landscape display.  See [harfang.ComputeAspectRatioY].
func ComputeAspectRatioX(width float32, height float32) *Vec2 {
	widthToC := C.float(width)
	heightToC := C.float(height)
	retval := C.WrapComputeAspectRatioX(widthToC, heightToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ComputeAspectRatioY Compute the aspect ratio factor for the provided viewport dimensions. Use this method to compute aspect ratio for portrait display.  See [harfang.ComputeAspectRatioX].
func ComputeAspectRatioY(width float32, height float32) *Vec2 {
	widthToC := C.float(width)
	heightToC := C.float(height)
	retval := C.WrapComputeAspectRatioY(widthToC, heightToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// Compute2DProjectionMatrix Returns a projection matrix from a 2D space to the 3D world, as required by [harfang.SetViewTransform] for example.
func Compute2DProjectionMatrix(znear float32, zfar float32, resx float32, resy float32, yup bool) *Mat44 {
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	resxToC := C.float(resx)
	resyToC := C.float(resy)
	yupToC := C.bool(yup)
	retval := C.WrapCompute2DProjectionMatrix(znearToC, zfarToC, resxToC, resyToC, yupToC)
	retvalGO := &Mat44{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Mat44) {
		C.WrapMat44Free(cleanval.h)
	})
	return retvalGO
}

// ExtractZoomFactorFromProjectionMatrix Extract zoom factor from a projection matrix.  See [harfang.ZoomFactorToFov].
func ExtractZoomFactorFromProjectionMatrix(m *Mat44) float32 {
	mToC := m.h
	retval := C.WrapExtractZoomFactorFromProjectionMatrix(mToC)
	return float32(retval)
}

// ExtractZRangeFromPerspectiveProjectionMatrix ...
func ExtractZRangeFromPerspectiveProjectionMatrix(m *Mat44) (*float32, *float32) {
	mToC := m.h
	znear := new(float32)
	znearToC := (*C.float)(unsafe.Pointer(znear))
	zfar := new(float32)
	zfarToC := (*C.float)(unsafe.Pointer(zfar))
	C.WrapExtractZRangeFromPerspectiveProjectionMatrix(mToC, znearToC, zfarToC)
	return (*float32)(unsafe.Pointer(znearToC)), (*float32)(unsafe.Pointer(zfarToC))
}

// ExtractZRangeFromOrthographicProjectionMatrix ...
func ExtractZRangeFromOrthographicProjectionMatrix(m *Mat44) (*float32, *float32) {
	mToC := m.h
	znear := new(float32)
	znearToC := (*C.float)(unsafe.Pointer(znear))
	zfar := new(float32)
	zfarToC := (*C.float)(unsafe.Pointer(zfar))
	C.WrapExtractZRangeFromOrthographicProjectionMatrix(mToC, znearToC, zfarToC)
	return (*float32)(unsafe.Pointer(znearToC)), (*float32)(unsafe.Pointer(zfarToC))
}

// ExtractZRangeFromProjectionMatrix Extract z near and z far clipping range from a projection matrix.
func ExtractZRangeFromProjectionMatrix(m *Mat44) (*float32, *float32) {
	mToC := m.h
	znear := new(float32)
	znearToC := (*C.float)(unsafe.Pointer(znear))
	zfar := new(float32)
	zfarToC := (*C.float)(unsafe.Pointer(zfar))
	C.WrapExtractZRangeFromProjectionMatrix(mToC, znearToC, zfarToC)
	return (*float32)(unsafe.Pointer(znearToC)), (*float32)(unsafe.Pointer(zfarToC))
}

// ProjectToClipSpace Project a world position to the clipping space.
func ProjectToClipSpace(proj *Mat44, view *Vec3) (bool, *Vec3) {
	projToC := proj.h
	viewToC := view.h
	clip := NewVec3()
	clipToC := clip.h
	retval := C.WrapProjectToClipSpace(projToC, viewToC, clipToC)
	return bool(retval), clip
}

// ProjectOrthoToClipSpace ...
func ProjectOrthoToClipSpace(proj *Mat44, view *Vec3) (bool, *Vec3) {
	projToC := proj.h
	viewToC := view.h
	clip := NewVec3()
	clipToC := clip.h
	retval := C.WrapProjectOrthoToClipSpace(projToC, viewToC, clipToC)
	return bool(retval), clip
}

// UnprojectFromClipSpace Unproject a clip space position to view space.
func UnprojectFromClipSpace(invproj *Mat44, clip *Vec3) (bool, *Vec3) {
	invprojToC := invproj.h
	clipToC := clip.h
	view := NewVec3()
	viewToC := view.h
	retval := C.WrapUnprojectFromClipSpace(invprojToC, clipToC, viewToC)
	return bool(retval), view
}

// UnprojectOrthoFromClipSpace ...
func UnprojectOrthoFromClipSpace(invproj *Mat44, clip *Vec3) (bool, *Vec3) {
	invprojToC := invproj.h
	clipToC := clip.h
	view := NewVec3()
	viewToC := view.h
	retval := C.WrapUnprojectOrthoFromClipSpace(invprojToC, clipToC, viewToC)
	return bool(retval), view
}

// ClipSpaceToScreenSpace Convert a 3d position in clip space (homogeneous space) to a 2d position on screen.
func ClipSpaceToScreenSpace(clip *Vec3, resolution *Vec2) *Vec3 {
	clipToC := clip.h
	resolutionToC := resolution.h
	retval := C.WrapClipSpaceToScreenSpace(clipToC, resolutionToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// ScreenSpaceToClipSpace Transform a screen position to clip space.
func ScreenSpaceToClipSpace(screen *Vec3, resolution *Vec2) *Vec3 {
	screenToC := screen.h
	resolutionToC := resolution.h
	retval := C.WrapScreenSpaceToClipSpace(screenToC, resolutionToC)
	retvalGO := &Vec3{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec3) {
		C.WrapVec3Free(cleanval.h)
	})
	return retvalGO
}

// ProjectToScreenSpace Project a world position to screen coordinates.
func ProjectToScreenSpace(proj *Mat44, view *Vec3, resolution *Vec2) (bool, *Vec3) {
	projToC := proj.h
	viewToC := view.h
	resolutionToC := resolution.h
	screen := NewVec3()
	screenToC := screen.h
	retval := C.WrapProjectToScreenSpace(projToC, viewToC, resolutionToC, screenToC)
	return bool(retval), screen
}

// ProjectOrthoToScreenSpace ...
func ProjectOrthoToScreenSpace(proj *Mat44, view *Vec3, resolution *Vec2) (bool, *Vec3) {
	projToC := proj.h
	viewToC := view.h
	resolutionToC := resolution.h
	screen := NewVec3()
	screenToC := screen.h
	retval := C.WrapProjectOrthoToScreenSpace(projToC, viewToC, resolutionToC, screenToC)
	return bool(retval), screen
}

// UnprojectFromScreenSpace Unproject a screen space position to view space.
func UnprojectFromScreenSpace(invproj *Mat44, screen *Vec3, resolution *Vec2) (bool, *Vec3) {
	invprojToC := invproj.h
	screenToC := screen.h
	resolutionToC := resolution.h
	view := NewVec3()
	viewToC := view.h
	retval := C.WrapUnprojectFromScreenSpace(invprojToC, screenToC, resolutionToC, viewToC)
	return bool(retval), view
}

// UnprojectOrthoFromScreenSpace ...
func UnprojectOrthoFromScreenSpace(invproj *Mat44, screen *Vec3, resolution *Vec2) (bool, *Vec3) {
	invprojToC := invproj.h
	screenToC := screen.h
	resolutionToC := resolution.h
	view := NewVec3()
	viewToC := view.h
	retval := C.WrapUnprojectOrthoFromScreenSpace(invprojToC, screenToC, resolutionToC, viewToC)
	return bool(retval), view
}

// ProjectZToClipSpace Project a depth value to clip space.
func ProjectZToClipSpace(z float32, proj *Mat44) float32 {
	zToC := C.float(z)
	projToC := proj.h
	retval := C.WrapProjectZToClipSpace(zToC, projToC)
	return float32(retval)
}

// MakeFrustum Create a projection frustum. This object can then be used to perform culling using [harfang.TestVisibility].  ```python # Compute a perspective matrix proj = hg.ComputePerspectiveProjectionMatrix(0.1, 1000, hg.FovToZoomFactor(math.pi/4), 1280/720) # Make a frustum from this projection matrix frustum = hg.MakeFrustum(proj) ```
func MakeFrustum(projection *Mat44) *Frustum {
	projectionToC := projection.h
	retval := C.WrapMakeFrustum(projectionToC)
	retvalGO := &Frustum{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Frustum) {
		C.WrapFrustumFree(cleanval.h)
	})
	return retvalGO
}

// MakeFrustumWithMtx Create a projection frustum. This object can then be used to perform culling using [harfang.TestVisibility].  ```python # Compute a perspective matrix proj = hg.ComputePerspectiveProjectionMatrix(0.1, 1000, hg.FovToZoomFactor(math.pi/4), 1280/720) # Make a frustum from this projection matrix frustum = hg.MakeFrustum(proj) ```
func MakeFrustumWithMtx(projection *Mat44, mtx *Mat4) *Frustum {
	projectionToC := projection.h
	mtxToC := mtx.h
	retval := C.WrapMakeFrustumWithMtx(projectionToC, mtxToC)
	retvalGO := &Frustum{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Frustum) {
		C.WrapFrustumFree(cleanval.h)
	})
	return retvalGO
}

// TransformFrustum Return the input frustum transformed by the provided world matrix.
func TransformFrustum(frustum *Frustum, mtx *Mat4) *Frustum {
	frustumToC := frustum.h
	mtxToC := mtx.h
	retval := C.WrapTransformFrustum(frustumToC, mtxToC)
	retvalGO := &Frustum{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Frustum) {
		C.WrapFrustumFree(cleanval.h)
	})
	return retvalGO
}

// TestVisibilityWithCountPoints Test if a list of 3d points are inside or outside a [harfang.Frustum].
func TestVisibilityWithCountPoints(frustum *Frustum, count uint32, points *Vec3) Visibility {
	frustumToC := frustum.h
	countToC := C.uint32_t(count)
	pointsToC := points.h
	retval := C.WrapTestVisibilityWithCountPoints(frustumToC, countToC, pointsToC)
	return Visibility(retval)
}

// TestVisibilityWithCountPointsDistance Test if a list of 3d points are inside or outside a [harfang.Frustum].
func TestVisibilityWithCountPointsDistance(frustum *Frustum, count uint32, points *Vec3, distance float32) Visibility {
	frustumToC := frustum.h
	countToC := C.uint32_t(count)
	pointsToC := points.h
	distanceToC := C.float(distance)
	retval := C.WrapTestVisibilityWithCountPointsDistance(frustumToC, countToC, pointsToC, distanceToC)
	return Visibility(retval)
}

// TestVisibilityWithOriginRadius Test if a list of 3d points are inside or outside a [harfang.Frustum].
func TestVisibilityWithOriginRadius(frustum *Frustum, origin *Vec3, radius float32) Visibility {
	frustumToC := frustum.h
	originToC := origin.h
	radiusToC := C.float(radius)
	retval := C.WrapTestVisibilityWithOriginRadius(frustumToC, originToC, radiusToC)
	return Visibility(retval)
}

// TestVisibility Test if a list of 3d points are inside or outside a [harfang.Frustum].
func TestVisibility(frustum *Frustum, minmax *MinMax) Visibility {
	frustumToC := frustum.h
	minmaxToC := minmax.h
	retval := C.WrapTestVisibility(frustumToC, minmaxToC)
	return Visibility(retval)
}

// WindowSystemInit Initialize the Window system.
func WindowSystemInit() {
	C.WrapWindowSystemInit()
}

// WindowSystemShutdown Shutdown the window system.  See [harfang.WindowSystemInit].
func WindowSystemShutdown() {
	C.WrapWindowSystemShutdown()
}

// GetMonitors Return a list of monitors connected to the system.
func GetMonitors() *MonitorList {
	retval := C.WrapGetMonitors()
	retvalGO := &MonitorList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MonitorList) {
		C.WrapMonitorListFree(cleanval.h)
	})
	return retvalGO
}

// GetMonitorRect Returns a rectangle going from the position, in screen coordinates, of the upper-left corner of the specified monitor to the position of the lower-right corner.
func GetMonitorRect(monitor *Monitor) *IntRect {
	monitorToC := monitor.h
	retval := C.WrapGetMonitorRect(monitorToC)
	retvalGO := &IntRect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IntRect) {
		C.WrapIntRectFree(cleanval.h)
	})
	return retvalGO
}

// IsPrimaryMonitor Return `true` if the monitor is the primary host device monitor, `false` otherwise.
func IsPrimaryMonitor(monitor *Monitor) bool {
	monitorToC := monitor.h
	retval := C.WrapIsPrimaryMonitor(monitorToC)
	return bool(retval)
}

// IsMonitorConnected Test if the specified monitor is connected to the host device.
func IsMonitorConnected(monitor *Monitor) bool {
	monitorToC := monitor.h
	retval := C.WrapIsMonitorConnected(monitorToC)
	return bool(retval)
}

// GetMonitorName Return the monitor name.
func GetMonitorName(monitor *Monitor) string {
	monitorToC := monitor.h
	retval := C.WrapGetMonitorName(monitorToC)
	return C.GoString(retval)
}

// GetMonitorSizeMM Returns the size, in millimetres, of the display area of the specified monitor.
func GetMonitorSizeMM(monitor *Monitor) *IVec2 {
	monitorToC := monitor.h
	retval := C.WrapGetMonitorSizeMM(monitorToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// GetMonitorModes Return the list of supported monitor modes.
func GetMonitorModes(monitor *Monitor) (bool, *MonitorModeList) {
	monitorToC := monitor.h
	modes := NewMonitorModeList()
	modesToC := modes.h
	retval := C.WrapGetMonitorModes(monitorToC, modesToC)
	return bool(retval), modes
}

// NewWindow Create a new window.
func NewWindow(width int32, height int32) *Window {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	retval := C.WrapNewWindow(widthToC, heightToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// NewWindowWithBpp Create a new window.
func NewWindowWithBpp(width int32, height int32, bpp int32) *Window {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	bppToC := C.int32_t(bpp)
	retval := C.WrapNewWindowWithBpp(widthToC, heightToC, bppToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// NewWindowWithBppVisibility Create a new window.
func NewWindowWithBppVisibility(width int32, height int32, bpp int32, visibility WindowVisibility) *Window {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	bppToC := C.int32_t(bpp)
	visibilityToC := C.int32_t(visibility)
	retval := C.WrapNewWindowWithBppVisibility(widthToC, heightToC, bppToC, visibilityToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// NewWindowWithTitleWidthHeight Create a new window.
func NewWindowWithTitleWidthHeight(title string, width int32, height int32) *Window {
	titleToC, idFintitleToC := wrapString(title)
	defer idFintitleToC()
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	retval := C.WrapNewWindowWithTitleWidthHeight(titleToC, widthToC, heightToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// NewWindowWithTitleWidthHeightBpp Create a new window.
func NewWindowWithTitleWidthHeightBpp(title string, width int32, height int32, bpp int32) *Window {
	titleToC, idFintitleToC := wrapString(title)
	defer idFintitleToC()
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	bppToC := C.int32_t(bpp)
	retval := C.WrapNewWindowWithTitleWidthHeightBpp(titleToC, widthToC, heightToC, bppToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// NewWindowWithTitleWidthHeightBppVisibility Create a new window.
func NewWindowWithTitleWidthHeightBppVisibility(title string, width int32, height int32, bpp int32, visibility WindowVisibility) *Window {
	titleToC, idFintitleToC := wrapString(title)
	defer idFintitleToC()
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	bppToC := C.int32_t(bpp)
	visibilityToC := C.int32_t(visibility)
	retval := C.WrapNewWindowWithTitleWidthHeightBppVisibility(titleToC, widthToC, heightToC, bppToC, visibilityToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// NewFullscreenWindow Create a new fullscreen window.
func NewFullscreenWindow(monitor *Monitor, modeindex int32) *Window {
	monitorToC := monitor.h
	modeindexToC := C.int32_t(modeindex)
	retval := C.WrapNewFullscreenWindow(monitorToC, modeindexToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// NewFullscreenWindowWithRotation Create a new fullscreen window.
func NewFullscreenWindowWithRotation(monitor *Monitor, modeindex int32, rotation MonitorRotation) *Window {
	monitorToC := monitor.h
	modeindexToC := C.int32_t(modeindex)
	rotationToC := C.uchar(rotation)
	retval := C.WrapNewFullscreenWindowWithRotation(monitorToC, modeindexToC, rotationToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// NewFullscreenWindowWithTitleMonitorModeIndex Create a new fullscreen window.
func NewFullscreenWindowWithTitleMonitorModeIndex(title string, monitor *Monitor, modeindex int32) *Window {
	titleToC, idFintitleToC := wrapString(title)
	defer idFintitleToC()
	monitorToC := monitor.h
	modeindexToC := C.int32_t(modeindex)
	retval := C.WrapNewFullscreenWindowWithTitleMonitorModeIndex(titleToC, monitorToC, modeindexToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// NewFullscreenWindowWithTitleMonitorModeIndexRotation Create a new fullscreen window.
func NewFullscreenWindowWithTitleMonitorModeIndexRotation(title string, monitor *Monitor, modeindex int32, rotation MonitorRotation) *Window {
	titleToC, idFintitleToC := wrapString(title)
	defer idFintitleToC()
	monitorToC := monitor.h
	modeindexToC := C.int32_t(modeindex)
	rotationToC := C.uchar(rotation)
	retval := C.WrapNewFullscreenWindowWithTitleMonitorModeIndexRotation(titleToC, monitorToC, modeindexToC, rotationToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// NewWindowFrom Wrap a native window handle in a [harfang.Window] object.
func NewWindowFrom(handle *VoidPointer) *Window {
	handleToC := handle.h
	retval := C.WrapNewWindowFrom(handleToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// GetWindowHandle Return the system native window handle.
func GetWindowHandle(window *Window) *VoidPointer {
	windowToC := window.h
	retval := C.WrapGetWindowHandle(windowToC)
	var retvalGO *VoidPointer
	if retval != nil {
		retvalGO = &VoidPointer{h: retval}
	}
	return retvalGO
}

// UpdateWindow Update a window on the host system.
func UpdateWindow(window *Window) bool {
	windowToC := window.h
	retval := C.WrapUpdateWindow(windowToC)
	return bool(retval)
}

// DestroyWindow Destroy a window object.
func DestroyWindow(window *Window) bool {
	windowToC := window.h
	retval := C.WrapDestroyWindow(windowToC)
	return bool(retval)
}

// GetWindowClientSize Return a window client rectangle. The client area of a window does not include its decorations.
func GetWindowClientSize(window *Window) (bool, *int32, *int32) {
	windowToC := window.h
	width := new(int32)
	widthToC := (*C.int32_t)(unsafe.Pointer(width))
	height := new(int32)
	heightToC := (*C.int32_t)(unsafe.Pointer(height))
	retval := C.WrapGetWindowClientSize(windowToC, widthToC, heightToC)
	return bool(retval), (*int32)(unsafe.Pointer(widthToC)), (*int32)(unsafe.Pointer(heightToC))
}

// SetWindowClientSize Set the window client size. The client area of a window excludes its decoration.
func SetWindowClientSize(window *Window, width int32, height int32) bool {
	windowToC := window.h
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	retval := C.WrapSetWindowClientSize(windowToC, widthToC, heightToC)
	return bool(retval)
}

// GetWindowContentScale ...
func GetWindowContentScale(window *Window) *Vec2 {
	windowToC := window.h
	retval := C.WrapGetWindowContentScale(windowToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// GetWindowTitle Return a window title.
func GetWindowTitle(window *Window) (bool, *string) {
	windowToC := window.h
	title := new(string)
	titleToC1 := C.CString(*title)
	titleToC := &titleToC1
	retval := C.WrapGetWindowTitle(windowToC, titleToC)
	titleToCGO := string(C.GoString(*titleToC))
	return bool(retval), &titleToCGO
}

// SetWindowTitle Set window title.
func SetWindowTitle(window *Window, title string) bool {
	windowToC := window.h
	titleToC, idFintitleToC := wrapString(title)
	defer idFintitleToC()
	retval := C.WrapSetWindowTitle(windowToC, titleToC)
	return bool(retval)
}

// WindowHasFocus Return `true` if the provided window has focus, `false` otherwise.
func WindowHasFocus(window *Window) bool {
	windowToC := window.h
	retval := C.WrapWindowHasFocus(windowToC)
	return bool(retval)
}

// GetWindowInFocus Return the system window with input focus.
func GetWindowInFocus() *Window {
	retval := C.WrapGetWindowInFocus()
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// GetWindowPos Return a window position on screen.
func GetWindowPos(window *Window) *IVec2 {
	windowToC := window.h
	retval := C.WrapGetWindowPos(windowToC)
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// SetWindowPos Set window position.
func SetWindowPos(window *Window, position *IVec2) bool {
	windowToC := window.h
	positionToC := position.h
	retval := C.WrapSetWindowPos(windowToC, positionToC)
	return bool(retval)
}

// IsWindowOpen Return `true` if the window is open, `false` otherwise.
func IsWindowOpen(window *Window) bool {
	windowToC := window.h
	retval := C.WrapIsWindowOpen(windowToC)
	return bool(retval)
}

// ShowCursor Show the system mouse cursor.  See [harfang.HideCursor].
func ShowCursor() {
	C.WrapShowCursor()
}

// HideCursor Hide the system mouse cursor.  See [harfang.ShowCursor].
func HideCursor() {
	C.WrapHideCursor()
}

// ColorToGrayscale Return the grayscale representation of a color. A weighted average is used to account for human perception of colors.
func ColorToGrayscale(color *Color) float32 {
	colorToC := color.h
	retval := C.WrapColorToGrayscale(colorToC)
	return float32(retval)
}

// ColorToRGBA32 Return a 32 bit RGBA integer from a color.
func ColorToRGBA32(color *Color) uint32 {
	colorToC := color.h
	retval := C.WrapColorToRGBA32(colorToC)
	return uint32(retval)
}

// ColorFromRGBA32 Create a color from a 32 bit RGBA integer.
func ColorFromRGBA32(rgba32 uint32) *Color {
	rgba32ToC := C.uint32_t(rgba32)
	retval := C.WrapColorFromRGBA32(rgba32ToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// ColorToABGR32 Return a 32 bit ABGR integer from a color.
func ColorToABGR32(color *Color) uint32 {
	colorToC := color.h
	retval := C.WrapColorToABGR32(colorToC)
	return uint32(retval)
}

// ColorFromABGR32 Create a color from a 32 bit ABGR integer.
func ColorFromABGR32(rgba32 uint32) *Color {
	rgba32ToC := C.uint32_t(rgba32)
	retval := C.WrapColorFromABGR32(rgba32ToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// ARGB32ToRGBA32 Convert a 32 bit integer ARGB color to RGBA.
func ARGB32ToRGBA32(argb uint32) uint32 {
	argbToC := C.uint32_t(argb)
	retval := C.WrapARGB32ToRGBA32(argbToC)
	return uint32(retval)
}

// RGBA32 Create a 32 bit integer RGBA color.
func RGBA32(r uint8, g uint8, b uint8) uint32 {
	rToC := C.uchar(r)
	gToC := C.uchar(g)
	bToC := C.uchar(b)
	retval := C.WrapRGBA32(rToC, gToC, bToC)
	return uint32(retval)
}

// RGBA32WithA Create a 32 bit integer RGBA color.
func RGBA32WithA(r uint8, g uint8, b uint8, a uint8) uint32 {
	rToC := C.uchar(r)
	gToC := C.uchar(g)
	bToC := C.uchar(b)
	aToC := C.uchar(a)
	retval := C.WrapRGBA32WithA(rToC, gToC, bToC, aToC)
	return uint32(retval)
}

// ARGB32 Create a 32 bit integer ARGB color.
func ARGB32(r uint8, g uint8, b uint8) uint32 {
	rToC := C.uchar(r)
	gToC := C.uchar(g)
	bToC := C.uchar(b)
	retval := C.WrapARGB32(rToC, gToC, bToC)
	return uint32(retval)
}

// ARGB32WithA Create a 32 bit integer ARGB color.
func ARGB32WithA(r uint8, g uint8, b uint8, a uint8) uint32 {
	rToC := C.uchar(r)
	gToC := C.uchar(g)
	bToC := C.uchar(b)
	aToC := C.uchar(a)
	retval := C.WrapARGB32WithA(rToC, gToC, bToC, aToC)
	return uint32(retval)
}

// ChromaScale Return a copy of the color with its saturation scaled as specified.
func ChromaScale(color *Color, k float32) *Color {
	colorToC := color.h
	kToC := C.float(k)
	retval := C.WrapChromaScale(colorToC, kToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// AlphaScale Scale the alpha component of the input color.
func AlphaScale(color *Color, k float32) *Color {
	colorToC := color.h
	kToC := C.float(k)
	retval := C.WrapAlphaScale(colorToC, kToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// ColorFromVector3 Create a color from a 3d vector, alpha defaults to 1.
func ColorFromVector3(v *Vec3) *Color {
	vToC := v.h
	retval := C.WrapColorFromVector3(vToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// ColorFromVector4 Return a 4-dimensional vector as a color.
func ColorFromVector4(v *Vec4) *Color {
	vToC := v.h
	retval := C.WrapColorFromVector4(vToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// ColorI Create a color from integer values in the [harfang.0;255] range.
func ColorI(r int32, g int32, b int32) *Color {
	rToC := C.int32_t(r)
	gToC := C.int32_t(g)
	bToC := C.int32_t(b)
	retval := C.WrapColorI(rToC, gToC, bToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// ColorIWithA Create a color from integer values in the [harfang.0;255] range.
func ColorIWithA(r int32, g int32, b int32, a int32) *Color {
	rToC := C.int32_t(r)
	gToC := C.int32_t(g)
	bToC := C.int32_t(b)
	aToC := C.int32_t(a)
	retval := C.WrapColorIWithA(rToC, gToC, bToC, aToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// ToHLS Convert input RGBA color to hue/luminance/saturation, alpha channel is left unmodified.
func ToHLS(color *Color) *Color {
	colorToC := color.h
	retval := C.WrapToHLS(colorToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// FromHLS Convert input hue/luminance/saturation color to RGBA, alpha channel is left unmodified.
func FromHLS(color *Color) *Color {
	colorToC := color.h
	retval := C.WrapFromHLS(colorToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// SetSaturation Return a copy of the input RGBA color with its saturation set to the specified value, alpha channel is left unmodified.  See [harfang.ToHLS] and [harfang.FromHLS].
func SetSaturation(color *Color, saturation float32) *Color {
	colorToC := color.h
	saturationToC := C.float(saturation)
	retval := C.WrapSetSaturation(colorToC, saturationToC)
	retvalGO := &Color{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Color) {
		C.WrapColorFree(cleanval.h)
	})
	return retvalGO
}

// LoadJPG Load a [harfang.Picture] in [harfang.JPEG](https://en.wikipedia.org/wiki/JPEG) file format.
func LoadJPG(pict *Picture, path string) bool {
	pictToC := pict.h
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadJPG(pictToC, pathToC)
	return bool(retval)
}

// LoadPNG Load a [harfang.Picture] in [harfang.PNG](https://en.wikipedia.org/wiki/Portable_Network_Graphics) file format.
func LoadPNG(pict *Picture, path string) bool {
	pictToC := pict.h
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadPNG(pictToC, pathToC)
	return bool(retval)
}

// LoadGIF Load a [harfang.Picture] in [harfang.GIF](https://en.wikipedia.org/wiki/GIF) file format.
func LoadGIF(pict *Picture, path string) bool {
	pictToC := pict.h
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadGIF(pictToC, pathToC)
	return bool(retval)
}

// LoadPSD Load a [harfang.Picture] in [harfang.PSD](https://en.wikipedia.org/wiki/Adobe_Photoshop#File_format) file format.
func LoadPSD(pict *Picture, path string) bool {
	pictToC := pict.h
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadPSD(pictToC, pathToC)
	return bool(retval)
}

// LoadTGA Load a [harfang.Picture] in [harfang.TGA](https://en.wikipedia.org/wiki/Truevision_TGA) file format.
func LoadTGA(pict *Picture, path string) bool {
	pictToC := pict.h
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadTGA(pictToC, pathToC)
	return bool(retval)
}

// LoadBMP Load a [harfang.Picture] in [harfang.BMP](https://en.wikipedia.org/wiki/BMP_file_format) file format.
func LoadBMP(pict *Picture, path string) bool {
	pictToC := pict.h
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadBMP(pictToC, pathToC)
	return bool(retval)
}

// LoadPicture Load a [harfang.Picture] content from the filesystem.
func LoadPicture(pict *Picture, path string) bool {
	pictToC := pict.h
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadPicture(pictToC, pathToC)
	return bool(retval)
}

// SavePNG Save a [harfang.Picture] in [harfang.PNG](https://en.wikipedia.org/wiki/Portable_Network_Graphics) file format.
func SavePNG(pict *Picture, path string) bool {
	pictToC := pict.h
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapSavePNG(pictToC, pathToC)
	return bool(retval)
}

// SaveTGA Save a [harfang.Picture] in [harfang.TGA](https://en.wikipedia.org/wiki/Truevision_TGA) file format.
func SaveTGA(pict *Picture, path string) bool {
	pictToC := pict.h
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapSaveTGA(pictToC, pathToC)
	return bool(retval)
}

// SaveBMP Save a [harfang.Picture] in [harfang.BMP](https://en.wikipedia.org/wiki/BMP_file_format) file format.
func SaveBMP(pict *Picture, path string) bool {
	pictToC := pict.h
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapSaveBMP(pictToC, pathToC)
	return bool(retval)
}

// RenderInit Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInit(window *Window) bool {
	windowToC := window.h
	retval := C.WrapRenderInit(windowToC)
	return bool(retval)
}

// RenderInitWithType Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithType(window *Window, typeGo RendererType) bool {
	windowToC := window.h
	typeGoToC := C.int32_t(typeGo)
	retval := C.WrapRenderInitWithType(windowToC, typeGoToC)
	return bool(retval)
}

// RenderInitWithWidthHeightResetFlags Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWidthHeightResetFlags(width int32, height int32, resetflags ResetFlags) *Window {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	resetflagsToC := C.uint32_t(resetflags)
	retval := C.WrapRenderInitWithWidthHeightResetFlags(widthToC, heightToC, resetflagsToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWidthHeightResetFlagsFormat Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWidthHeightResetFlagsFormat(width int32, height int32, resetflags ResetFlags, format TextureFormat) *Window {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	resetflagsToC := C.uint32_t(resetflags)
	formatToC := C.int32_t(format)
	retval := C.WrapRenderInitWithWidthHeightResetFlagsFormat(widthToC, heightToC, resetflagsToC, formatToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWidthHeightResetFlagsFormatDebugFlags Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWidthHeightResetFlagsFormatDebugFlags(width int32, height int32, resetflags ResetFlags, format TextureFormat, debugflags DebugFlags) *Window {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	resetflagsToC := C.uint32_t(resetflags)
	formatToC := C.int32_t(format)
	debugflagsToC := C.uint32_t(debugflags)
	retval := C.WrapRenderInitWithWidthHeightResetFlagsFormatDebugFlags(widthToC, heightToC, resetflagsToC, formatToC, debugflagsToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWidthHeightType Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWidthHeightType(width int32, height int32, typeGo RendererType) *Window {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	typeGoToC := C.int32_t(typeGo)
	retval := C.WrapRenderInitWithWidthHeightType(widthToC, heightToC, typeGoToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWidthHeightTypeResetFlags Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWidthHeightTypeResetFlags(width int32, height int32, typeGo RendererType, resetflags ResetFlags) *Window {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	typeGoToC := C.int32_t(typeGo)
	resetflagsToC := C.uint32_t(resetflags)
	retval := C.WrapRenderInitWithWidthHeightTypeResetFlags(widthToC, heightToC, typeGoToC, resetflagsToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWidthHeightTypeResetFlagsFormat Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWidthHeightTypeResetFlagsFormat(width int32, height int32, typeGo RendererType, resetflags ResetFlags, format TextureFormat) *Window {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	typeGoToC := C.int32_t(typeGo)
	resetflagsToC := C.uint32_t(resetflags)
	formatToC := C.int32_t(format)
	retval := C.WrapRenderInitWithWidthHeightTypeResetFlagsFormat(widthToC, heightToC, typeGoToC, resetflagsToC, formatToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWidthHeightTypeResetFlagsFormatDebugFlags Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWidthHeightTypeResetFlagsFormatDebugFlags(width int32, height int32, typeGo RendererType, resetflags ResetFlags, format TextureFormat, debugflags DebugFlags) *Window {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	typeGoToC := C.int32_t(typeGo)
	resetflagsToC := C.uint32_t(resetflags)
	formatToC := C.int32_t(format)
	debugflagsToC := C.uint32_t(debugflags)
	retval := C.WrapRenderInitWithWidthHeightTypeResetFlagsFormatDebugFlags(widthToC, heightToC, typeGoToC, resetflagsToC, formatToC, debugflagsToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWindowTitleWidthHeightResetFlags Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWindowTitleWidthHeightResetFlags(windowtitle string, width int32, height int32, resetflags ResetFlags) *Window {
	windowtitleToC, idFinwindowtitleToC := wrapString(windowtitle)
	defer idFinwindowtitleToC()
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	resetflagsToC := C.uint32_t(resetflags)
	retval := C.WrapRenderInitWithWindowTitleWidthHeightResetFlags(windowtitleToC, widthToC, heightToC, resetflagsToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWindowTitleWidthHeightResetFlagsFormat Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWindowTitleWidthHeightResetFlagsFormat(windowtitle string, width int32, height int32, resetflags ResetFlags, format TextureFormat) *Window {
	windowtitleToC, idFinwindowtitleToC := wrapString(windowtitle)
	defer idFinwindowtitleToC()
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	resetflagsToC := C.uint32_t(resetflags)
	formatToC := C.int32_t(format)
	retval := C.WrapRenderInitWithWindowTitleWidthHeightResetFlagsFormat(windowtitleToC, widthToC, heightToC, resetflagsToC, formatToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWindowTitleWidthHeightResetFlagsFormatDebugFlags Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWindowTitleWidthHeightResetFlagsFormatDebugFlags(windowtitle string, width int32, height int32, resetflags ResetFlags, format TextureFormat, debugflags DebugFlags) *Window {
	windowtitleToC, idFinwindowtitleToC := wrapString(windowtitle)
	defer idFinwindowtitleToC()
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	resetflagsToC := C.uint32_t(resetflags)
	formatToC := C.int32_t(format)
	debugflagsToC := C.uint32_t(debugflags)
	retval := C.WrapRenderInitWithWindowTitleWidthHeightResetFlagsFormatDebugFlags(windowtitleToC, widthToC, heightToC, resetflagsToC, formatToC, debugflagsToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWindowTitleWidthHeightType Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWindowTitleWidthHeightType(windowtitle string, width int32, height int32, typeGo RendererType) *Window {
	windowtitleToC, idFinwindowtitleToC := wrapString(windowtitle)
	defer idFinwindowtitleToC()
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	typeGoToC := C.int32_t(typeGo)
	retval := C.WrapRenderInitWithWindowTitleWidthHeightType(windowtitleToC, widthToC, heightToC, typeGoToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWindowTitleWidthHeightTypeResetFlags Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWindowTitleWidthHeightTypeResetFlags(windowtitle string, width int32, height int32, typeGo RendererType, resetflags ResetFlags) *Window {
	windowtitleToC, idFinwindowtitleToC := wrapString(windowtitle)
	defer idFinwindowtitleToC()
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	typeGoToC := C.int32_t(typeGo)
	resetflagsToC := C.uint32_t(resetflags)
	retval := C.WrapRenderInitWithWindowTitleWidthHeightTypeResetFlags(windowtitleToC, widthToC, heightToC, typeGoToC, resetflagsToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWindowTitleWidthHeightTypeResetFlagsFormat Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWindowTitleWidthHeightTypeResetFlagsFormat(windowtitle string, width int32, height int32, typeGo RendererType, resetflags ResetFlags, format TextureFormat) *Window {
	windowtitleToC, idFinwindowtitleToC := wrapString(windowtitle)
	defer idFinwindowtitleToC()
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	typeGoToC := C.int32_t(typeGo)
	resetflagsToC := C.uint32_t(resetflags)
	formatToC := C.int32_t(format)
	retval := C.WrapRenderInitWithWindowTitleWidthHeightTypeResetFlagsFormat(windowtitleToC, widthToC, heightToC, typeGoToC, resetflagsToC, formatToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderInitWithWindowTitleWidthHeightTypeResetFlagsFormatDebugFlags Initialize the render system.  To change the states of the render system afterward use [harfang.RenderReset].
func RenderInitWithWindowTitleWidthHeightTypeResetFlagsFormatDebugFlags(windowtitle string, width int32, height int32, typeGo RendererType, resetflags ResetFlags, format TextureFormat, debugflags DebugFlags) *Window {
	windowtitleToC, idFinwindowtitleToC := wrapString(windowtitle)
	defer idFinwindowtitleToC()
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	typeGoToC := C.int32_t(typeGo)
	resetflagsToC := C.uint32_t(resetflags)
	formatToC := C.int32_t(format)
	debugflagsToC := C.uint32_t(debugflags)
	retval := C.WrapRenderInitWithWindowTitleWidthHeightTypeResetFlagsFormatDebugFlags(windowtitleToC, widthToC, heightToC, typeGoToC, resetflagsToC, formatToC, debugflagsToC)
	var retvalGO *Window
	if retval != nil {
		retvalGO = &Window{h: retval}
	}
	return retvalGO
}

// RenderShutdown Shutdown the render system.
func RenderShutdown() {
	C.WrapRenderShutdown()
}

// RenderResetToWindow Resize the renderer backbuffer to the provided window client area dimensions. Return true if a reset was needed and carried out.
func RenderResetToWindow(win *Window, width *int32, height *int32) bool {
	winToC := win.h
	widthToC := (*C.int32_t)(unsafe.Pointer(width))
	heightToC := (*C.int32_t)(unsafe.Pointer(height))
	retval := C.WrapRenderResetToWindow(winToC, widthToC, heightToC)
	return bool(retval)
}

// RenderResetToWindowWithResetFlags Resize the renderer backbuffer to the provided window client area dimensions. Return true if a reset was needed and carried out.
func RenderResetToWindowWithResetFlags(win *Window, width *int32, height *int32, resetflags uint32) bool {
	winToC := win.h
	widthToC := (*C.int32_t)(unsafe.Pointer(width))
	heightToC := (*C.int32_t)(unsafe.Pointer(height))
	resetflagsToC := C.uint32_t(resetflags)
	retval := C.WrapRenderResetToWindowWithResetFlags(winToC, widthToC, heightToC, resetflagsToC)
	return bool(retval)
}

// RenderReset Change the states of the render system at runtime.
func RenderReset(width uint32, height uint32) {
	widthToC := C.uint32_t(width)
	heightToC := C.uint32_t(height)
	C.WrapRenderReset(widthToC, heightToC)
}

// RenderResetWithFlags Change the states of the render system at runtime.
func RenderResetWithFlags(width uint32, height uint32, flags ResetFlags) {
	widthToC := C.uint32_t(width)
	heightToC := C.uint32_t(height)
	flagsToC := C.uint32_t(flags)
	C.WrapRenderResetWithFlags(widthToC, heightToC, flagsToC)
}

// RenderResetWithFlagsFormat Change the states of the render system at runtime.
func RenderResetWithFlagsFormat(width uint32, height uint32, flags ResetFlags, format TextureFormat) {
	widthToC := C.uint32_t(width)
	heightToC := C.uint32_t(height)
	flagsToC := C.uint32_t(flags)
	formatToC := C.int32_t(format)
	C.WrapRenderResetWithFlagsFormat(widthToC, heightToC, flagsToC, formatToC)
}

// SetRenderDebug Set render system debug flags.
func SetRenderDebug(flags DebugFlags) {
	flagsToC := C.uint32_t(flags)
	C.WrapSetRenderDebug(flagsToC)
}

// SetViewClear Set a view clear parameters.  See [harfang.man.Views].
func SetViewClear(viewid uint16, flags ClearFlags) {
	viewidToC := C.ushort(viewid)
	flagsToC := C.ushort(flags)
	C.WrapSetViewClear(viewidToC, flagsToC)
}

// SetViewClearWithRgba Set a view clear parameters.  See [harfang.man.Views].
func SetViewClearWithRgba(viewid uint16, flags ClearFlags, rgba uint32) {
	viewidToC := C.ushort(viewid)
	flagsToC := C.ushort(flags)
	rgbaToC := C.uint32_t(rgba)
	C.WrapSetViewClearWithRgba(viewidToC, flagsToC, rgbaToC)
}

// SetViewClearWithRgbaDepth Set a view clear parameters.  See [harfang.man.Views].
func SetViewClearWithRgbaDepth(viewid uint16, flags ClearFlags, rgba uint32, depth float32) {
	viewidToC := C.ushort(viewid)
	flagsToC := C.ushort(flags)
	rgbaToC := C.uint32_t(rgba)
	depthToC := C.float(depth)
	C.WrapSetViewClearWithRgbaDepth(viewidToC, flagsToC, rgbaToC, depthToC)
}

// SetViewClearWithRgbaDepthStencil Set a view clear parameters.  See [harfang.man.Views].
func SetViewClearWithRgbaDepthStencil(viewid uint16, flags ClearFlags, rgba uint32, depth float32, stencil uint8) {
	viewidToC := C.ushort(viewid)
	flagsToC := C.ushort(flags)
	rgbaToC := C.uint32_t(rgba)
	depthToC := C.float(depth)
	stencilToC := C.uchar(stencil)
	C.WrapSetViewClearWithRgbaDepthStencil(viewidToC, flagsToC, rgbaToC, depthToC, stencilToC)
}

// SetViewClearWithCol Set a view clear parameters.  See [harfang.man.Views].
func SetViewClearWithCol(viewid uint16, flags ClearFlags, col *Color) {
	viewidToC := C.ushort(viewid)
	flagsToC := C.ushort(flags)
	colToC := col.h
	C.WrapSetViewClearWithCol(viewidToC, flagsToC, colToC)
}

// SetViewClearWithColDepth Set a view clear parameters.  See [harfang.man.Views].
func SetViewClearWithColDepth(viewid uint16, flags ClearFlags, col *Color, depth float32) {
	viewidToC := C.ushort(viewid)
	flagsToC := C.ushort(flags)
	colToC := col.h
	depthToC := C.float(depth)
	C.WrapSetViewClearWithColDepth(viewidToC, flagsToC, colToC, depthToC)
}

// SetViewClearWithColDepthStencil Set a view clear parameters.  See [harfang.man.Views].
func SetViewClearWithColDepthStencil(viewid uint16, flags ClearFlags, col *Color, depth float32, stencil uint8) {
	viewidToC := C.ushort(viewid)
	flagsToC := C.ushort(flags)
	colToC := col.h
	depthToC := C.float(depth)
	stencilToC := C.uchar(stencil)
	C.WrapSetViewClearWithColDepthStencil(viewidToC, flagsToC, colToC, depthToC, stencilToC)
}

// SetViewRect ...
func SetViewRect(viewid uint16, x uint16, y uint16, w uint16, h uint16) {
	viewidToC := C.ushort(viewid)
	xToC := C.ushort(x)
	yToC := C.ushort(y)
	wToC := C.ushort(w)
	hToC := C.ushort(h)
	C.WrapSetViewRect(viewidToC, xToC, yToC, wToC, hToC)
}

// SetViewFrameBuffer Set view output framebuffer.  See [harfang.man.Views].
func SetViewFrameBuffer(viewid uint16, handle *FrameBufferHandle) {
	viewidToC := C.ushort(viewid)
	handleToC := handle.h
	C.WrapSetViewFrameBuffer(viewidToC, handleToC)
}

// SetViewMode Set view draw ordering mode.
func SetViewMode(viewid uint16, mode ViewMode) {
	viewidToC := C.ushort(viewid)
	modeToC := C.int32_t(mode)
	C.WrapSetViewMode(viewidToC, modeToC)
}

// Touch Submit an empty primitive to the view.  See [harfang.Frame].
func Touch(viewid uint16) {
	viewidToC := C.ushort(viewid)
	C.WrapTouch(viewidToC)
}

// Frame Advance the rendering backend to the next frame, execute all queued rendering commands. This function returns the backend current frame.  The frame counter is used by asynchronous functions such as [harfang.CaptureTexture]. You must wait for the frame counter to reach or exceed the value returned by an asynchronous function before accessing its result.
func Frame() uint32 {
	retval := C.WrapFrame()
	return uint32(retval)
}

// SetViewTransform Set view transforms, namely the view and projection matrices.
func SetViewTransform(viewid uint16, view *Mat4, proj *Mat44) {
	viewidToC := C.ushort(viewid)
	viewToC := view.h
	projToC := proj.h
	C.WrapSetViewTransform(viewidToC, viewToC, projToC)
}

// SetView2D High-level wrapper function to setup a view for 2D rendering.  This function calls [harfang.SetViewClear], [harfang.SetViewRect] then [harfang.SetViewTransform].
func SetView2D(id uint16, x int32, y int32, resx int32, resy int32) {
	idToC := C.ushort(id)
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	resxToC := C.int32_t(resx)
	resyToC := C.int32_t(resy)
	C.WrapSetView2D(idToC, xToC, yToC, resxToC, resyToC)
}

// SetView2DWithZnearZfar High-level wrapper function to setup a view for 2D rendering.  This function calls [harfang.SetViewClear], [harfang.SetViewRect] then [harfang.SetViewTransform].
func SetView2DWithZnearZfar(id uint16, x int32, y int32, resx int32, resy int32, znear float32, zfar float32) {
	idToC := C.ushort(id)
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	resxToC := C.int32_t(resx)
	resyToC := C.int32_t(resy)
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	C.WrapSetView2DWithZnearZfar(idToC, xToC, yToC, resxToC, resyToC, znearToC, zfarToC)
}

// SetView2DWithZnearZfarFlagsColorDepthStencil High-level wrapper function to setup a view for 2D rendering.  This function calls [harfang.SetViewClear], [harfang.SetViewRect] then [harfang.SetViewTransform].
func SetView2DWithZnearZfarFlagsColorDepthStencil(id uint16, x int32, y int32, resx int32, resy int32, znear float32, zfar float32, flags ClearFlags, color *Color, depth float32, stencil uint8) {
	idToC := C.ushort(id)
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	resxToC := C.int32_t(resx)
	resyToC := C.int32_t(resy)
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	flagsToC := C.ushort(flags)
	colorToC := color.h
	depthToC := C.float(depth)
	stencilToC := C.uchar(stencil)
	C.WrapSetView2DWithZnearZfarFlagsColorDepthStencil(idToC, xToC, yToC, resxToC, resyToC, znearToC, zfarToC, flagsToC, colorToC, depthToC, stencilToC)
}

// SetView2DWithZnearZfarFlagsColorDepthStencilYUp High-level wrapper function to setup a view for 2D rendering.  This function calls [harfang.SetViewClear], [harfang.SetViewRect] then [harfang.SetViewTransform].
func SetView2DWithZnearZfarFlagsColorDepthStencilYUp(id uint16, x int32, y int32, resx int32, resy int32, znear float32, zfar float32, flags ClearFlags, color *Color, depth float32, stencil uint8, yup bool) {
	idToC := C.ushort(id)
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	resxToC := C.int32_t(resx)
	resyToC := C.int32_t(resy)
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	flagsToC := C.ushort(flags)
	colorToC := color.h
	depthToC := C.float(depth)
	stencilToC := C.uchar(stencil)
	yupToC := C.bool(yup)
	C.WrapSetView2DWithZnearZfarFlagsColorDepthStencilYUp(idToC, xToC, yToC, resxToC, resyToC, znearToC, zfarToC, flagsToC, colorToC, depthToC, stencilToC, yupToC)
}

// SetViewPerspective High-level wrapper function to setup a view for 3D perspective rendering.  This function calls [harfang.SetViewClear], [harfang.SetViewRect] then [harfang.SetViewTransform].
func SetViewPerspective(id uint16, x int32, y int32, resx int32, resy int32, world *Mat4) {
	idToC := C.ushort(id)
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	resxToC := C.int32_t(resx)
	resyToC := C.int32_t(resy)
	worldToC := world.h
	C.WrapSetViewPerspective(idToC, xToC, yToC, resxToC, resyToC, worldToC)
}

// SetViewPerspectiveWithZnearZfar High-level wrapper function to setup a view for 3D perspective rendering.  This function calls [harfang.SetViewClear], [harfang.SetViewRect] then [harfang.SetViewTransform].
func SetViewPerspectiveWithZnearZfar(id uint16, x int32, y int32, resx int32, resy int32, world *Mat4, znear float32, zfar float32) {
	idToC := C.ushort(id)
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	resxToC := C.int32_t(resx)
	resyToC := C.int32_t(resy)
	worldToC := world.h
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	C.WrapSetViewPerspectiveWithZnearZfar(idToC, xToC, yToC, resxToC, resyToC, worldToC, znearToC, zfarToC)
}

// SetViewPerspectiveWithZnearZfarZoomFactor High-level wrapper function to setup a view for 3D perspective rendering.  This function calls [harfang.SetViewClear], [harfang.SetViewRect] then [harfang.SetViewTransform].
func SetViewPerspectiveWithZnearZfarZoomFactor(id uint16, x int32, y int32, resx int32, resy int32, world *Mat4, znear float32, zfar float32, zoomfactor float32) {
	idToC := C.ushort(id)
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	resxToC := C.int32_t(resx)
	resyToC := C.int32_t(resy)
	worldToC := world.h
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	zoomfactorToC := C.float(zoomfactor)
	C.WrapSetViewPerspectiveWithZnearZfarZoomFactor(idToC, xToC, yToC, resxToC, resyToC, worldToC, znearToC, zfarToC, zoomfactorToC)
}

// SetViewPerspectiveWithZnearZfarZoomFactorFlagsColorDepthStencil High-level wrapper function to setup a view for 3D perspective rendering.  This function calls [harfang.SetViewClear], [harfang.SetViewRect] then [harfang.SetViewTransform].
func SetViewPerspectiveWithZnearZfarZoomFactorFlagsColorDepthStencil(id uint16, x int32, y int32, resx int32, resy int32, world *Mat4, znear float32, zfar float32, zoomfactor float32, flags ClearFlags, color *Color, depth float32, stencil uint8) {
	idToC := C.ushort(id)
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	resxToC := C.int32_t(resx)
	resyToC := C.int32_t(resy)
	worldToC := world.h
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	zoomfactorToC := C.float(zoomfactor)
	flagsToC := C.ushort(flags)
	colorToC := color.h
	depthToC := C.float(depth)
	stencilToC := C.uchar(stencil)
	C.WrapSetViewPerspectiveWithZnearZfarZoomFactorFlagsColorDepthStencil(idToC, xToC, yToC, resxToC, resyToC, worldToC, znearToC, zfarToC, zoomfactorToC, flagsToC, colorToC, depthToC, stencilToC)
}

// SetViewOrthographic High-level wrapper function to setup a view for 3D orthographic rendering.  This function calls [harfang.SetViewClear], [harfang.SetViewRect] then [harfang.SetViewTransform].
func SetViewOrthographic(id uint16, x int32, y int32, resx int32, resy int32, world *Mat4) {
	idToC := C.ushort(id)
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	resxToC := C.int32_t(resx)
	resyToC := C.int32_t(resy)
	worldToC := world.h
	C.WrapSetViewOrthographic(idToC, xToC, yToC, resxToC, resyToC, worldToC)
}

// SetViewOrthographicWithZnearZfar High-level wrapper function to setup a view for 3D orthographic rendering.  This function calls [harfang.SetViewClear], [harfang.SetViewRect] then [harfang.SetViewTransform].
func SetViewOrthographicWithZnearZfar(id uint16, x int32, y int32, resx int32, resy int32, world *Mat4, znear float32, zfar float32) {
	idToC := C.ushort(id)
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	resxToC := C.int32_t(resx)
	resyToC := C.int32_t(resy)
	worldToC := world.h
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	C.WrapSetViewOrthographicWithZnearZfar(idToC, xToC, yToC, resxToC, resyToC, worldToC, znearToC, zfarToC)
}

// SetViewOrthographicWithZnearZfarSize High-level wrapper function to setup a view for 3D orthographic rendering.  This function calls [harfang.SetViewClear], [harfang.SetViewRect] then [harfang.SetViewTransform].
func SetViewOrthographicWithZnearZfarSize(id uint16, x int32, y int32, resx int32, resy int32, world *Mat4, znear float32, zfar float32, size float32) {
	idToC := C.ushort(id)
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	resxToC := C.int32_t(resx)
	resyToC := C.int32_t(resy)
	worldToC := world.h
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	sizeToC := C.float(size)
	C.WrapSetViewOrthographicWithZnearZfarSize(idToC, xToC, yToC, resxToC, resyToC, worldToC, znearToC, zfarToC, sizeToC)
}

// SetViewOrthographicWithZnearZfarSizeFlagsColorDepthStencil High-level wrapper function to setup a view for 3D orthographic rendering.  This function calls [harfang.SetViewClear], [harfang.SetViewRect] then [harfang.SetViewTransform].
func SetViewOrthographicWithZnearZfarSizeFlagsColorDepthStencil(id uint16, x int32, y int32, resx int32, resy int32, world *Mat4, znear float32, zfar float32, size float32, flags ClearFlags, color *Color, depth float32, stencil uint8) {
	idToC := C.ushort(id)
	xToC := C.int32_t(x)
	yToC := C.int32_t(y)
	resxToC := C.int32_t(resx)
	resyToC := C.int32_t(resy)
	worldToC := world.h
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	sizeToC := C.float(size)
	flagsToC := C.ushort(flags)
	colorToC := color.h
	depthToC := C.float(depth)
	stencilToC := C.uchar(stencil)
	C.WrapSetViewOrthographicWithZnearZfarSizeFlagsColorDepthStencil(idToC, xToC, yToC, resxToC, resyToC, worldToC, znearToC, zfarToC, sizeToC, flagsToC, colorToC, depthToC, stencilToC)
}

// VertexLayoutPosFloatNormFloat Simple vertex layout with float position and normal.  ```python vtx_layout = VertexLayout() vtx_layout.Begin() vtx_layout.Add(hg.A_Position, 3, hg.AT_Float) vtx_layout.Add(hg.A_Normal, 3, hg.AT_Float) vtx_layout.End() ```
func VertexLayoutPosFloatNormFloat() *VertexLayout {
	retval := C.WrapVertexLayoutPosFloatNormFloat()
	retvalGO := &VertexLayout{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VertexLayout) {
		C.WrapVertexLayoutFree(cleanval.h)
	})
	return retvalGO
}

// VertexLayoutPosFloatNormUInt8 Simple vertex layout with float position and 8-bit unsigned integer normal.  ```python vtx_layout = VertexLayout() vtx_layout.Begin() vtx_layout.Add(hg.A_Position, 3, hg.AT_Float) vtx_layout.Add(hg.A_Normal, 3, hg.AT_Uint8, True, True) vtx_layout.End() ```
func VertexLayoutPosFloatNormUInt8() *VertexLayout {
	retval := C.WrapVertexLayoutPosFloatNormUInt8()
	retvalGO := &VertexLayout{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VertexLayout) {
		C.WrapVertexLayoutFree(cleanval.h)
	})
	return retvalGO
}

// VertexLayoutPosFloatColorFloat ...
func VertexLayoutPosFloatColorFloat() *VertexLayout {
	retval := C.WrapVertexLayoutPosFloatColorFloat()
	retvalGO := &VertexLayout{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VertexLayout) {
		C.WrapVertexLayoutFree(cleanval.h)
	})
	return retvalGO
}

// VertexLayoutPosFloatColorUInt8 ...
func VertexLayoutPosFloatColorUInt8() *VertexLayout {
	retval := C.WrapVertexLayoutPosFloatColorUInt8()
	retvalGO := &VertexLayout{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VertexLayout) {
		C.WrapVertexLayoutFree(cleanval.h)
	})
	return retvalGO
}

// VertexLayoutPosFloatTexCoord0UInt8 ...
func VertexLayoutPosFloatTexCoord0UInt8() *VertexLayout {
	retval := C.WrapVertexLayoutPosFloatTexCoord0UInt8()
	retvalGO := &VertexLayout{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VertexLayout) {
		C.WrapVertexLayoutFree(cleanval.h)
	})
	return retvalGO
}

// VertexLayoutPosFloatNormUInt8TexCoord0UInt8 ...
func VertexLayoutPosFloatNormUInt8TexCoord0UInt8() *VertexLayout {
	retval := C.WrapVertexLayoutPosFloatNormUInt8TexCoord0UInt8()
	retvalGO := &VertexLayout{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VertexLayout) {
		C.WrapVertexLayoutFree(cleanval.h)
	})
	return retvalGO
}

// LoadProgramFromFile Load a shader program from the local filesystem.
func LoadProgramFromFile(path string) *ProgramHandle {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadProgramFromFile(pathToC)
	retvalGO := &ProgramHandle{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ProgramHandle) {
		C.WrapProgramHandleFree(cleanval.h)
	})
	return retvalGO
}

// LoadProgramFromFileWithVertexShaderPathFragmentShaderPath Load a shader program from the local filesystem.
func LoadProgramFromFileWithVertexShaderPathFragmentShaderPath(vertexshaderpath string, fragmentshaderpath string) *ProgramHandle {
	vertexshaderpathToC, idFinvertexshaderpathToC := wrapString(vertexshaderpath)
	defer idFinvertexshaderpathToC()
	fragmentshaderpathToC, idFinfragmentshaderpathToC := wrapString(fragmentshaderpath)
	defer idFinfragmentshaderpathToC()
	retval := C.WrapLoadProgramFromFileWithVertexShaderPathFragmentShaderPath(vertexshaderpathToC, fragmentshaderpathToC)
	retvalGO := &ProgramHandle{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ProgramHandle) {
		C.WrapProgramHandleFree(cleanval.h)
	})
	return retvalGO
}

// LoadProgramFromAssets Load a shader program from the assets system.  See [harfang.man.Assets].
func LoadProgramFromAssets(name string) *ProgramHandle {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapLoadProgramFromAssets(nameToC)
	retvalGO := &ProgramHandle{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ProgramHandle) {
		C.WrapProgramHandleFree(cleanval.h)
	})
	return retvalGO
}

// LoadProgramFromAssetsWithVertexShaderNameFragmentShaderName Load a shader program from the assets system.  See [harfang.man.Assets].
func LoadProgramFromAssetsWithVertexShaderNameFragmentShaderName(vertexshadername string, fragmentshadername string) *ProgramHandle {
	vertexshadernameToC, idFinvertexshadernameToC := wrapString(vertexshadername)
	defer idFinvertexshadernameToC()
	fragmentshadernameToC, idFinfragmentshadernameToC := wrapString(fragmentshadername)
	defer idFinfragmentshadernameToC()
	retval := C.WrapLoadProgramFromAssetsWithVertexShaderNameFragmentShaderName(vertexshadernameToC, fragmentshadernameToC)
	retvalGO := &ProgramHandle{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ProgramHandle) {
		C.WrapProgramHandleFree(cleanval.h)
	})
	return retvalGO
}

// DestroyProgram Destroy a shader program.
func DestroyProgram(h *ProgramHandle) {
	hToC := h.h
	C.WrapDestroyProgram(hToC)
}

// LoadTextureFlagsFromFile Load texture flags in the texture metafile from the local filesystem.
func LoadTextureFlagsFromFile(path string) uint64 {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadTextureFlagsFromFile(pathToC)
	return uint64(retval)
}

// LoadTextureFlagsFromAssets Load texture flags in the texture metafile from the assets system.  See [harfang.man.Assets].
func LoadTextureFlagsFromAssets(name string) uint64 {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapLoadTextureFlagsFromAssets(nameToC)
	return uint64(retval)
}

// CreateTexture Create an empty texture.  See [harfang.CreateTextureFromPicture] and [harfang.UpdateTextureFromPicture].
func CreateTexture(width int32, height int32, name string, flags TextureFlags) *Texture {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	flagsToC := C.uint64_t(flags)
	retval := C.WrapCreateTexture(widthToC, heightToC, nameToC, flagsToC)
	retvalGO := &Texture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Texture) {
		C.WrapTextureFree(cleanval.h)
	})
	return retvalGO
}

// CreateTextureWithFormat Create an empty texture.  See [harfang.CreateTextureFromPicture] and [harfang.UpdateTextureFromPicture].
func CreateTextureWithFormat(width int32, height int32, name string, flags TextureFlags, format TextureFormat) *Texture {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	flagsToC := C.uint64_t(flags)
	formatToC := C.int32_t(format)
	retval := C.WrapCreateTextureWithFormat(widthToC, heightToC, nameToC, flagsToC, formatToC)
	retvalGO := &Texture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Texture) {
		C.WrapTextureFree(cleanval.h)
	})
	return retvalGO
}

// CreateTextureFromPicture Create a texture from a picture.  See [harfang.Picture], [harfang.CreateTexture] and [harfang.UpdateTextureFromPicture].
func CreateTextureFromPicture(pic *Picture, name string, flags TextureFlags) *Texture {
	picToC := pic.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	flagsToC := C.uint64_t(flags)
	retval := C.WrapCreateTextureFromPicture(picToC, nameToC, flagsToC)
	retvalGO := &Texture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Texture) {
		C.WrapTextureFree(cleanval.h)
	})
	return retvalGO
}

// CreateTextureFromPictureWithFormat Create a texture from a picture.  See [harfang.Picture], [harfang.CreateTexture] and [harfang.UpdateTextureFromPicture].
func CreateTextureFromPictureWithFormat(pic *Picture, name string, flags TextureFlags, format TextureFormat) *Texture {
	picToC := pic.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	flagsToC := C.uint64_t(flags)
	formatToC := C.int32_t(format)
	retval := C.WrapCreateTextureFromPictureWithFormat(picToC, nameToC, flagsToC, formatToC)
	retvalGO := &Texture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Texture) {
		C.WrapTextureFree(cleanval.h)
	})
	return retvalGO
}

// UpdateTextureFromPicture Update texture content from the provided picture.  Note: The picture is expected to be in a format compatible with the texture format.
func UpdateTextureFromPicture(tex *Texture, pic *Picture) {
	texToC := tex.h
	picToC := pic.h
	C.WrapUpdateTextureFromPicture(texToC, picToC)
}

// LoadTextureFromFile Load a texture from the local filesystem.  - When not using pipeline resources the texture informations are returned directly. - When using pipeline resources the texture informations can be retrieved from the [harfang.PipelineResources] object.
func LoadTextureFromFile(path string, flags TextureFlags) (*Texture, *TextureInfo) {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	flagsToC := C.uint64_t(flags)
	info := NewTextureInfo()
	infoToC := info.h
	retval := C.WrapLoadTextureFromFile(pathToC, flagsToC, infoToC)
	retvalGO := &Texture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Texture) {
		C.WrapTextureFree(cleanval.h)
	})
	return retvalGO, info
}

// LoadTextureFromFileWithFlagsResources Load a texture from the local filesystem.  - When not using pipeline resources the texture informations are returned directly. - When using pipeline resources the texture informations can be retrieved from the [harfang.PipelineResources] object.
func LoadTextureFromFileWithFlagsResources(path string, flags uint32, resources *PipelineResources) *TextureRef {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	flagsToC := C.uint32_t(flags)
	resourcesToC := resources.h
	retval := C.WrapLoadTextureFromFileWithFlagsResources(pathToC, flagsToC, resourcesToC)
	retvalGO := &TextureRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *TextureRef) {
		C.WrapTextureRefFree(cleanval.h)
	})
	return retvalGO
}

// LoadTextureFromAssets Load a texture from the assets system.  - When not using pipeline resources the texture informations are returned directly. - When using pipeline resources the texture informations can be retrieved from the [harfang.PipelineResources] object.  See [harfang.man.Assets].
func LoadTextureFromAssets(path string, flags TextureFlags) (*Texture, *TextureInfo) {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	flagsToC := C.uint64_t(flags)
	info := NewTextureInfo()
	infoToC := info.h
	retval := C.WrapLoadTextureFromAssets(pathToC, flagsToC, infoToC)
	retvalGO := &Texture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Texture) {
		C.WrapTextureFree(cleanval.h)
	})
	return retvalGO, info
}

// LoadTextureFromAssetsWithFlagsResources Load a texture from the assets system.  - When not using pipeline resources the texture informations are returned directly. - When using pipeline resources the texture informations can be retrieved from the [harfang.PipelineResources] object.  See [harfang.man.Assets].
func LoadTextureFromAssetsWithFlagsResources(path string, flags uint32, resources *PipelineResources) *TextureRef {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	flagsToC := C.uint32_t(flags)
	resourcesToC := resources.h
	retval := C.WrapLoadTextureFromAssetsWithFlagsResources(pathToC, flagsToC, resourcesToC)
	retvalGO := &TextureRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *TextureRef) {
		C.WrapTextureRefFree(cleanval.h)
	})
	return retvalGO
}

// DestroyTexture Destroy a texture object.
func DestroyTexture(tex *Texture) {
	texToC := tex.h
	C.WrapDestroyTexture(texToC)
}

// ProcessTextureLoadQueue Process the texture load queue. This function must be called to load textures queued while loading a scene or model with the LSSF_QueueTextureLoads flag.  See [harfang.LoadSaveSceneFlags].
func ProcessTextureLoadQueue(res *PipelineResources) int32 {
	resToC := res.h
	retval := C.WrapProcessTextureLoadQueue(resToC)
	return int32(retval)
}

// ProcessTextureLoadQueueWithTBudget Process the texture load queue. This function must be called to load textures queued while loading a scene or model with the LSSF_QueueTextureLoads flag.  See [harfang.LoadSaveSceneFlags].
func ProcessTextureLoadQueueWithTBudget(res *PipelineResources, tbudget int64) int32 {
	resToC := res.h
	tbudgetToC := C.int64_t(tbudget)
	retval := C.WrapProcessTextureLoadQueueWithTBudget(resToC, tbudgetToC)
	return int32(retval)
}

// ProcessModelLoadQueue ...
func ProcessModelLoadQueue(res *PipelineResources) int32 {
	resToC := res.h
	retval := C.WrapProcessModelLoadQueue(resToC)
	return int32(retval)
}

// ProcessModelLoadQueueWithTBudget ...
func ProcessModelLoadQueueWithTBudget(res *PipelineResources, tbudget int64) int32 {
	resToC := res.h
	tbudgetToC := C.int64_t(tbudget)
	retval := C.WrapProcessModelLoadQueueWithTBudget(resToC, tbudgetToC)
	return int32(retval)
}

// ProcessLoadQueues ...
func ProcessLoadQueues(res *PipelineResources) int32 {
	resToC := res.h
	retval := C.WrapProcessLoadQueues(resToC)
	return int32(retval)
}

// ProcessLoadQueuesWithTBudget ...
func ProcessLoadQueuesWithTBudget(res *PipelineResources, tbudget int64) int32 {
	resToC := res.h
	tbudgetToC := C.int64_t(tbudget)
	retval := C.WrapProcessLoadQueuesWithTBudget(resToC, tbudgetToC)
	return int32(retval)
}

// CaptureTexture Capture a texture content to a [harfang.Picture]. Return the frame counter at which the capture will be complete.  A [harfang.Picture] object can be accessed by the CPU.  This function is asynchronous and its result will not be available until the returned frame counter is equal or greater to the frame counter returned by [harfang.Frame].
func CaptureTexture(resources *PipelineResources, tex *TextureRef, pic *Picture) uint32 {
	resourcesToC := resources.h
	texToC := tex.h
	picToC := pic.h
	retval := C.WrapCaptureTexture(resourcesToC, texToC, picToC)
	return uint32(retval)
}

// MakeUniformSetValue Create a uniform set value object.  This object can be added to a [harfang.UniformSetValueList] to control the shader program uniform values for a subsequent call to [harfang.DrawModel].
func MakeUniformSetValue(name string, v float32) *UniformSetValue {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	vToC := C.float(v)
	retval := C.WrapMakeUniformSetValue(nameToC, vToC)
	retvalGO := &UniformSetValue{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetValue) {
		C.WrapUniformSetValueFree(cleanval.h)
	})
	return retvalGO
}

// MakeUniformSetValueWithV Create a uniform set value object.  This object can be added to a [harfang.UniformSetValueList] to control the shader program uniform values for a subsequent call to [harfang.DrawModel].
func MakeUniformSetValueWithV(name string, v *Vec2) *UniformSetValue {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	vToC := v.h
	retval := C.WrapMakeUniformSetValueWithV(nameToC, vToC)
	retvalGO := &UniformSetValue{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetValue) {
		C.WrapUniformSetValueFree(cleanval.h)
	})
	return retvalGO
}

// MakeUniformSetValueWithVec3V Create a uniform set value object.  This object can be added to a [harfang.UniformSetValueList] to control the shader program uniform values for a subsequent call to [harfang.DrawModel].
func MakeUniformSetValueWithVec3V(name string, v *Vec3) *UniformSetValue {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	vToC := v.h
	retval := C.WrapMakeUniformSetValueWithVec3V(nameToC, vToC)
	retvalGO := &UniformSetValue{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetValue) {
		C.WrapUniformSetValueFree(cleanval.h)
	})
	return retvalGO
}

// MakeUniformSetValueWithVec4V Create a uniform set value object.  This object can be added to a [harfang.UniformSetValueList] to control the shader program uniform values for a subsequent call to [harfang.DrawModel].
func MakeUniformSetValueWithVec4V(name string, v *Vec4) *UniformSetValue {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	vToC := v.h
	retval := C.WrapMakeUniformSetValueWithVec4V(nameToC, vToC)
	retvalGO := &UniformSetValue{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetValue) {
		C.WrapUniformSetValueFree(cleanval.h)
	})
	return retvalGO
}

// MakeUniformSetValueWithMat3V Create a uniform set value object.  This object can be added to a [harfang.UniformSetValueList] to control the shader program uniform values for a subsequent call to [harfang.DrawModel].
func MakeUniformSetValueWithMat3V(name string, v *Mat3) *UniformSetValue {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	vToC := v.h
	retval := C.WrapMakeUniformSetValueWithMat3V(nameToC, vToC)
	retvalGO := &UniformSetValue{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetValue) {
		C.WrapUniformSetValueFree(cleanval.h)
	})
	return retvalGO
}

// MakeUniformSetValueWithMat4V Create a uniform set value object.  This object can be added to a [harfang.UniformSetValueList] to control the shader program uniform values for a subsequent call to [harfang.DrawModel].
func MakeUniformSetValueWithMat4V(name string, v *Mat4) *UniformSetValue {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	vToC := v.h
	retval := C.WrapMakeUniformSetValueWithMat4V(nameToC, vToC)
	retvalGO := &UniformSetValue{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetValue) {
		C.WrapUniformSetValueFree(cleanval.h)
	})
	return retvalGO
}

// MakeUniformSetValueWithMat44V Create a uniform set value object.  This object can be added to a [harfang.UniformSetValueList] to control the shader program uniform values for a subsequent call to [harfang.DrawModel].
func MakeUniformSetValueWithMat44V(name string, v *Mat44) *UniformSetValue {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	vToC := v.h
	retval := C.WrapMakeUniformSetValueWithMat44V(nameToC, vToC)
	retvalGO := &UniformSetValue{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetValue) {
		C.WrapUniformSetValueFree(cleanval.h)
	})
	return retvalGO
}

// MakeUniformSetTexture Create a uniform set texture object.  This object can be added to a [harfang.UniformSetTextureList] to control the shader program uniform values for a subsequent call to [harfang.DrawModel].
func MakeUniformSetTexture(name string, texture *Texture, stage uint8) *UniformSetTexture {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	textureToC := texture.h
	stageToC := C.uchar(stage)
	retval := C.WrapMakeUniformSetTexture(nameToC, textureToC, stageToC)
	retvalGO := &UniformSetTexture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *UniformSetTexture) {
		C.WrapUniformSetTextureFree(cleanval.h)
	})
	return retvalGO
}

// LoadPipelineProgramFromFile Load a pipeline shader program from the local filesystem.
func LoadPipelineProgramFromFile(path string, resources *PipelineResources, pipeline *PipelineInfo) *PipelineProgram {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapLoadPipelineProgramFromFile(pathToC, resourcesToC, pipelineToC)
	retvalGO := &PipelineProgram{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *PipelineProgram) {
		C.WrapPipelineProgramFree(cleanval.h)
	})
	return retvalGO
}

// LoadPipelineProgramFromAssets Load a pipeline shader program from the assets system.  See [harfang.man.Assets].
func LoadPipelineProgramFromAssets(name string, resources *PipelineResources, pipeline *PipelineInfo) *PipelineProgram {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapLoadPipelineProgramFromAssets(nameToC, resourcesToC, pipelineToC)
	retvalGO := &PipelineProgram{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *PipelineProgram) {
		C.WrapPipelineProgramFree(cleanval.h)
	})
	return retvalGO
}

// LoadPipelineProgramRefFromFile ...
func LoadPipelineProgramRefFromFile(path string, resources *PipelineResources, pipeline *PipelineInfo) *PipelineProgramRef {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapLoadPipelineProgramRefFromFile(pathToC, resourcesToC, pipelineToC)
	retvalGO := &PipelineProgramRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *PipelineProgramRef) {
		C.WrapPipelineProgramRefFree(cleanval.h)
	})
	return retvalGO
}

// LoadPipelineProgramRefFromAssets ...
func LoadPipelineProgramRefFromAssets(name string, resources *PipelineResources, pipeline *PipelineInfo) *PipelineProgramRef {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapLoadPipelineProgramRefFromAssets(nameToC, resourcesToC, pipelineToC)
	retvalGO := &PipelineProgramRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *PipelineProgramRef) {
		C.WrapPipelineProgramRefFree(cleanval.h)
	})
	return retvalGO
}

// ComputeOrthographicViewState Compute an orthographic view state.  The `size` parameter controls the extends of the projected view. When projecting a 3d world this parameter is expressed in meters. Use the `aspect_ratio` parameter to prevent distortion from induced by non-square viewport.  See [harfang.ComputeOrthographicProjectionMatrix], [harfang.ComputeAspectRatioX] and [harfang.ComputeAspectRatioY].
func ComputeOrthographicViewState(world *Mat4, size float32, znear float32, zfar float32, aspectratio *Vec2) *ViewState {
	worldToC := world.h
	sizeToC := C.float(size)
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	aspectratioToC := aspectratio.h
	retval := C.WrapComputeOrthographicViewState(worldToC, sizeToC, znearToC, zfarToC, aspectratioToC)
	retvalGO := &ViewState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ViewState) {
		C.WrapViewStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputePerspectiveViewState Compute a perspective view state.  See [harfang.ComputePerspectiveProjectionMatrix], [harfang.ZoomFactorToFov], [harfang.FovToZoomFactor], [harfang.ComputeAspectRatioX] and [harfang.ComputeAspectRatioY].
func ComputePerspectiveViewState(world *Mat4, fov float32, znear float32, zfar float32, aspectratio *Vec2) *ViewState {
	worldToC := world.h
	fovToC := C.float(fov)
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	aspectratioToC := aspectratio.h
	retval := C.WrapComputePerspectiveViewState(worldToC, fovToC, znearToC, zfarToC, aspectratioToC)
	retvalGO := &ViewState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ViewState) {
		C.WrapViewStateFree(cleanval.h)
	})
	return retvalGO
}

// SetMaterialProgram Set material pipeline program.  You should call [harfang.UpdateMaterialPipelineProgramVariant] after changing a material pipeline program so that the correct variant is selected according to the material states.
func SetMaterialProgram(mat *Material, program *PipelineProgramRef) {
	matToC := mat.h
	programToC := program.h
	C.WrapSetMaterialProgram(matToC, programToC)
}

// SetMaterialValue Set a material uniform value.
func SetMaterialValue(mat *Material, name string, v float32) {
	matToC := mat.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	vToC := C.float(v)
	C.WrapSetMaterialValue(matToC, nameToC, vToC)
}

// SetMaterialValueWithV Set a material uniform value.
func SetMaterialValueWithV(mat *Material, name string, v *Vec2) {
	matToC := mat.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	vToC := v.h
	C.WrapSetMaterialValueWithV(matToC, nameToC, vToC)
}

// SetMaterialValueWithVec3V Set a material uniform value.
func SetMaterialValueWithVec3V(mat *Material, name string, v *Vec3) {
	matToC := mat.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	vToC := v.h
	C.WrapSetMaterialValueWithVec3V(matToC, nameToC, vToC)
}

// SetMaterialValueWithVec4V Set a material uniform value.
func SetMaterialValueWithVec4V(mat *Material, name string, v *Vec4) {
	matToC := mat.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	vToC := v.h
	C.WrapSetMaterialValueWithVec4V(matToC, nameToC, vToC)
}

// SetMaterialValueWithM Set a material uniform value.
func SetMaterialValueWithM(mat *Material, name string, m *Mat3) {
	matToC := mat.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	mToC := m.h
	C.WrapSetMaterialValueWithM(matToC, nameToC, mToC)
}

// SetMaterialValueWithMat4M Set a material uniform value.
func SetMaterialValueWithMat4M(mat *Material, name string, m *Mat4) {
	matToC := mat.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	mToC := m.h
	C.WrapSetMaterialValueWithMat4M(matToC, nameToC, mToC)
}

// SetMaterialValueWithMat44M Set a material uniform value.
func SetMaterialValueWithMat44M(mat *Material, name string, m *Mat44) {
	matToC := mat.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	mToC := m.h
	C.WrapSetMaterialValueWithMat44M(matToC, nameToC, mToC)
}

// SetMaterialTexture Set a material uniform texture and texture stage.  Note: The texture stage specified should match the uniform declaration in the shader program.
func SetMaterialTexture(mat *Material, name string, texture *TextureRef, stage uint8) {
	matToC := mat.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	textureToC := texture.h
	stageToC := C.uchar(stage)
	C.WrapSetMaterialTexture(matToC, nameToC, textureToC, stageToC)
}

// SetMaterialTextureRef Set a material uniform texture reference.  See [harfang.PipelineResources].
func SetMaterialTextureRef(mat *Material, name string, texture *TextureRef) {
	matToC := mat.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	textureToC := texture.h
	C.WrapSetMaterialTextureRef(matToC, nameToC, textureToC)
}

// GetMaterialTexture Return the texture reference assigned to a material named uniform.
func GetMaterialTexture(mat *Material, name string) *TextureRef {
	matToC := mat.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapGetMaterialTexture(matToC, nameToC)
	retvalGO := &TextureRef{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *TextureRef) {
		C.WrapTextureRefFree(cleanval.h)
	})
	return retvalGO
}

// GetMaterialTextures Return the list of names of a material texture uniforms.
func GetMaterialTextures(mat *Material) *StringList {
	matToC := mat.h
	retval := C.WrapGetMaterialTextures(matToC)
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// GetMaterialValues Return the list of names of a material value uniforms.
func GetMaterialValues(mat *Material) *StringList {
	matToC := mat.h
	retval := C.WrapGetMaterialValues(matToC)
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// GetMaterialFaceCulling Return a material culling mode.
func GetMaterialFaceCulling(mat *Material) FaceCulling {
	matToC := mat.h
	retval := C.WrapGetMaterialFaceCulling(matToC)
	return FaceCulling(retval)
}

// SetMaterialFaceCulling Set material face culling.
func SetMaterialFaceCulling(mat *Material, culling FaceCulling) {
	matToC := mat.h
	cullingToC := C.int32_t(culling)
	C.WrapSetMaterialFaceCulling(matToC, cullingToC)
}

// GetMaterialDepthTest Return a material depth test function.
func GetMaterialDepthTest(mat *Material) DepthTest {
	matToC := mat.h
	retval := C.WrapGetMaterialDepthTest(matToC)
	return DepthTest(retval)
}

// SetMaterialDepthTest Set material depth test.
func SetMaterialDepthTest(mat *Material, test DepthTest) {
	matToC := mat.h
	testToC := C.int32_t(test)
	C.WrapSetMaterialDepthTest(matToC, testToC)
}

// GetMaterialBlendMode Return a material blending mode.
func GetMaterialBlendMode(mat *Material) BlendMode {
	matToC := mat.h
	retval := C.WrapGetMaterialBlendMode(matToC)
	return BlendMode(retval)
}

// SetMaterialBlendMode Set material blend mode.
func SetMaterialBlendMode(mat *Material, mode BlendMode) {
	matToC := mat.h
	modeToC := C.int32_t(mode)
	C.WrapSetMaterialBlendMode(matToC, modeToC)
}

// GetMaterialWriteRGBA Return the material color mask.
func GetMaterialWriteRGBA(mat *Material) (*bool, *bool, *bool, *bool) {
	matToC := mat.h
	writer := new(bool)
	writerToC := (*C.bool)(unsafe.Pointer(writer))
	writeg := new(bool)
	writegToC := (*C.bool)(unsafe.Pointer(writeg))
	writeb := new(bool)
	writebToC := (*C.bool)(unsafe.Pointer(writeb))
	writea := new(bool)
	writeaToC := (*C.bool)(unsafe.Pointer(writea))
	C.WrapGetMaterialWriteRGBA(matToC, writerToC, writegToC, writebToC, writeaToC)
	return (*bool)(unsafe.Pointer(writerToC)), (*bool)(unsafe.Pointer(writegToC)), (*bool)(unsafe.Pointer(writebToC)), (*bool)(unsafe.Pointer(writeaToC))
}

// SetMaterialWriteRGBA Set a material color write mask.
func SetMaterialWriteRGBA(mat *Material, writer bool, writeg bool, writeb bool, writea bool) {
	matToC := mat.h
	writerToC := C.bool(writer)
	writegToC := C.bool(writeg)
	writebToC := C.bool(writeb)
	writeaToC := C.bool(writea)
	C.WrapSetMaterialWriteRGBA(matToC, writerToC, writegToC, writebToC, writeaToC)
}

// GetMaterialNormalMapInWorldSpace ...
func GetMaterialNormalMapInWorldSpace(mat *Material) bool {
	matToC := mat.h
	retval := C.WrapGetMaterialNormalMapInWorldSpace(matToC)
	return bool(retval)
}

// SetMaterialNormalMapInWorldSpace ...
func SetMaterialNormalMapInWorldSpace(mat *Material, enable bool) {
	matToC := mat.h
	enableToC := C.bool(enable)
	C.WrapSetMaterialNormalMapInWorldSpace(matToC, enableToC)
}

// GetMaterialWriteZ Return the material depth write mask.
func GetMaterialWriteZ(mat *Material) bool {
	matToC := mat.h
	retval := C.WrapGetMaterialWriteZ(matToC)
	return bool(retval)
}

// SetMaterialWriteZ Set a material depth write mask.
func SetMaterialWriteZ(mat *Material, enable bool) {
	matToC := mat.h
	enableToC := C.bool(enable)
	C.WrapSetMaterialWriteZ(matToC, enableToC)
}

// GetMaterialDiffuseUsesUV1 ...
func GetMaterialDiffuseUsesUV1(mat *Material) bool {
	matToC := mat.h
	retval := C.WrapGetMaterialDiffuseUsesUV1(matToC)
	return bool(retval)
}

// SetMaterialDiffuseUsesUV1 ...
func SetMaterialDiffuseUsesUV1(mat *Material, enable bool) {
	matToC := mat.h
	enableToC := C.bool(enable)
	C.WrapSetMaterialDiffuseUsesUV1(matToC, enableToC)
}

// GetMaterialSpecularUsesUV1 ...
func GetMaterialSpecularUsesUV1(mat *Material) bool {
	matToC := mat.h
	retval := C.WrapGetMaterialSpecularUsesUV1(matToC)
	return bool(retval)
}

// SetMaterialSpecularUsesUV1 ...
func SetMaterialSpecularUsesUV1(mat *Material, enable bool) {
	matToC := mat.h
	enableToC := C.bool(enable)
	C.WrapSetMaterialSpecularUsesUV1(matToC, enableToC)
}

// GetMaterialAmbientUsesUV1 ...
func GetMaterialAmbientUsesUV1(mat *Material) bool {
	matToC := mat.h
	retval := C.WrapGetMaterialAmbientUsesUV1(matToC)
	return bool(retval)
}

// SetMaterialAmbientUsesUV1 ...
func SetMaterialAmbientUsesUV1(mat *Material, enable bool) {
	matToC := mat.h
	enableToC := C.bool(enable)
	C.WrapSetMaterialAmbientUsesUV1(matToC, enableToC)
}

// GetMaterialSkinning ...
func GetMaterialSkinning(mat *Material) bool {
	matToC := mat.h
	retval := C.WrapGetMaterialSkinning(matToC)
	return bool(retval)
}

// SetMaterialSkinning ...
func SetMaterialSkinning(mat *Material, enable bool) {
	matToC := mat.h
	enableToC := C.bool(enable)
	C.WrapSetMaterialSkinning(matToC, enableToC)
}

// GetMaterialAlphaCut ...
func GetMaterialAlphaCut(mat *Material) bool {
	matToC := mat.h
	retval := C.WrapGetMaterialAlphaCut(matToC)
	return bool(retval)
}

// SetMaterialAlphaCut ...
func SetMaterialAlphaCut(mat *Material, enable bool) {
	matToC := mat.h
	enableToC := C.bool(enable)
	C.WrapSetMaterialAlphaCut(matToC, enableToC)
}

// CreateMaterial Helper function to create a material.  See [harfang.SetMaterialProgram], [harfang.SetMaterialValue] and [harfang.SetMaterialTexture].
func CreateMaterial(prg *PipelineProgramRef) *Material {
	prgToC := prg.h
	retval := C.WrapCreateMaterial(prgToC)
	retvalGO := &Material{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Material) {
		C.WrapMaterialFree(cleanval.h)
	})
	return retvalGO
}

// CreateMaterialWithValueNameValue Helper function to create a material.  See [harfang.SetMaterialProgram], [harfang.SetMaterialValue] and [harfang.SetMaterialTexture].
func CreateMaterialWithValueNameValue(prg *PipelineProgramRef, valuename string, value *Vec4) *Material {
	prgToC := prg.h
	valuenameToC, idFinvaluenameToC := wrapString(valuename)
	defer idFinvaluenameToC()
	valueToC := value.h
	retval := C.WrapCreateMaterialWithValueNameValue(prgToC, valuenameToC, valueToC)
	retvalGO := &Material{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Material) {
		C.WrapMaterialFree(cleanval.h)
	})
	return retvalGO
}

// CreateMaterialWithValueName0Value0ValueName1Value1 Helper function to create a material.  See [harfang.SetMaterialProgram], [harfang.SetMaterialValue] and [harfang.SetMaterialTexture].
func CreateMaterialWithValueName0Value0ValueName1Value1(prg *PipelineProgramRef, valuename0 string, value0 *Vec4, valuename1 string, value1 *Vec4) *Material {
	prgToC := prg.h
	valuename0ToC, idFinvaluename0ToC := wrapString(valuename0)
	defer idFinvaluename0ToC()
	value0ToC := value0.h
	valuename1ToC, idFinvaluename1ToC := wrapString(valuename1)
	defer idFinvaluename1ToC()
	value1ToC := value1.h
	retval := C.WrapCreateMaterialWithValueName0Value0ValueName1Value1(prgToC, valuename0ToC, value0ToC, valuename1ToC, value1ToC)
	retvalGO := &Material{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Material) {
		C.WrapMaterialFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderState Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderState(blend BlendMode) *RenderState {
	blendToC := C.int32_t(blend)
	retval := C.WrapComputeRenderState(blendToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderStateWithDepthTest Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderStateWithDepthTest(blend BlendMode, depthtest DepthTest) *RenderState {
	blendToC := C.int32_t(blend)
	depthtestToC := C.int32_t(depthtest)
	retval := C.WrapComputeRenderStateWithDepthTest(blendToC, depthtestToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderStateWithDepthTestCulling Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderStateWithDepthTestCulling(blend BlendMode, depthtest DepthTest, culling FaceCulling) *RenderState {
	blendToC := C.int32_t(blend)
	depthtestToC := C.int32_t(depthtest)
	cullingToC := C.int32_t(culling)
	retval := C.WrapComputeRenderStateWithDepthTestCulling(blendToC, depthtestToC, cullingToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderStateWithDepthTestCullingWriteZ Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderStateWithDepthTestCullingWriteZ(blend BlendMode, depthtest DepthTest, culling FaceCulling, writez bool) *RenderState {
	blendToC := C.int32_t(blend)
	depthtestToC := C.int32_t(depthtest)
	cullingToC := C.int32_t(culling)
	writezToC := C.bool(writez)
	retval := C.WrapComputeRenderStateWithDepthTestCullingWriteZ(blendToC, depthtestToC, cullingToC, writezToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderStateWithDepthTestCullingWriteZWriteR Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderStateWithDepthTestCullingWriteZWriteR(blend BlendMode, depthtest DepthTest, culling FaceCulling, writez bool, writer bool) *RenderState {
	blendToC := C.int32_t(blend)
	depthtestToC := C.int32_t(depthtest)
	cullingToC := C.int32_t(culling)
	writezToC := C.bool(writez)
	writerToC := C.bool(writer)
	retval := C.WrapComputeRenderStateWithDepthTestCullingWriteZWriteR(blendToC, depthtestToC, cullingToC, writezToC, writerToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderStateWithDepthTestCullingWriteZWriteRWriteG Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderStateWithDepthTestCullingWriteZWriteRWriteG(blend BlendMode, depthtest DepthTest, culling FaceCulling, writez bool, writer bool, writeg bool) *RenderState {
	blendToC := C.int32_t(blend)
	depthtestToC := C.int32_t(depthtest)
	cullingToC := C.int32_t(culling)
	writezToC := C.bool(writez)
	writerToC := C.bool(writer)
	writegToC := C.bool(writeg)
	retval := C.WrapComputeRenderStateWithDepthTestCullingWriteZWriteRWriteG(blendToC, depthtestToC, cullingToC, writezToC, writerToC, writegToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderStateWithDepthTestCullingWriteZWriteRWriteGWriteB Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderStateWithDepthTestCullingWriteZWriteRWriteGWriteB(blend BlendMode, depthtest DepthTest, culling FaceCulling, writez bool, writer bool, writeg bool, writeb bool) *RenderState {
	blendToC := C.int32_t(blend)
	depthtestToC := C.int32_t(depthtest)
	cullingToC := C.int32_t(culling)
	writezToC := C.bool(writez)
	writerToC := C.bool(writer)
	writegToC := C.bool(writeg)
	writebToC := C.bool(writeb)
	retval := C.WrapComputeRenderStateWithDepthTestCullingWriteZWriteRWriteGWriteB(blendToC, depthtestToC, cullingToC, writezToC, writerToC, writegToC, writebToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderStateWithDepthTestCullingWriteZWriteRWriteGWriteBWriteA Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderStateWithDepthTestCullingWriteZWriteRWriteGWriteBWriteA(blend BlendMode, depthtest DepthTest, culling FaceCulling, writez bool, writer bool, writeg bool, writeb bool, writea bool) *RenderState {
	blendToC := C.int32_t(blend)
	depthtestToC := C.int32_t(depthtest)
	cullingToC := C.int32_t(culling)
	writezToC := C.bool(writez)
	writerToC := C.bool(writer)
	writegToC := C.bool(writeg)
	writebToC := C.bool(writeb)
	writeaToC := C.bool(writea)
	retval := C.WrapComputeRenderStateWithDepthTestCullingWriteZWriteRWriteGWriteBWriteA(blendToC, depthtestToC, cullingToC, writezToC, writerToC, writegToC, writebToC, writeaToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderStateWithWriteZ Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderStateWithWriteZ(blend BlendMode, writez bool) *RenderState {
	blendToC := C.int32_t(blend)
	writezToC := C.bool(writez)
	retval := C.WrapComputeRenderStateWithWriteZ(blendToC, writezToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderStateWithWriteZWriteR Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderStateWithWriteZWriteR(blend BlendMode, writez bool, writer bool) *RenderState {
	blendToC := C.int32_t(blend)
	writezToC := C.bool(writez)
	writerToC := C.bool(writer)
	retval := C.WrapComputeRenderStateWithWriteZWriteR(blendToC, writezToC, writerToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderStateWithWriteZWriteRWriteG Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderStateWithWriteZWriteRWriteG(blend BlendMode, writez bool, writer bool, writeg bool) *RenderState {
	blendToC := C.int32_t(blend)
	writezToC := C.bool(writez)
	writerToC := C.bool(writer)
	writegToC := C.bool(writeg)
	retval := C.WrapComputeRenderStateWithWriteZWriteRWriteG(blendToC, writezToC, writerToC, writegToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderStateWithWriteZWriteRWriteGWriteB Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderStateWithWriteZWriteRWriteGWriteB(blend BlendMode, writez bool, writer bool, writeg bool, writeb bool) *RenderState {
	blendToC := C.int32_t(blend)
	writezToC := C.bool(writez)
	writerToC := C.bool(writer)
	writegToC := C.bool(writeg)
	writebToC := C.bool(writeb)
	retval := C.WrapComputeRenderStateWithWriteZWriteRWriteGWriteB(blendToC, writezToC, writerToC, writegToC, writebToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeRenderStateWithWriteZWriteRWriteGWriteBWriteA Compute a render state to control subsequent render calls culling mode, blending mode, Z mask, etc... The same render state can be used by different render calls.  See [harfang.DrawLines], [harfang.DrawTriangles] and [harfang.DrawModel].
func ComputeRenderStateWithWriteZWriteRWriteGWriteBWriteA(blend BlendMode, writez bool, writer bool, writeg bool, writeb bool, writea bool) *RenderState {
	blendToC := C.int32_t(blend)
	writezToC := C.bool(writez)
	writerToC := C.bool(writer)
	writegToC := C.bool(writeg)
	writebToC := C.bool(writeb)
	writeaToC := C.bool(writea)
	retval := C.WrapComputeRenderStateWithWriteZWriteRWriteGWriteBWriteA(blendToC, writezToC, writerToC, writegToC, writebToC, writeaToC)
	retvalGO := &RenderState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *RenderState) {
		C.WrapRenderStateFree(cleanval.h)
	})
	return retvalGO
}

// ComputeSortKey Compute a sorting key to control the rendering order of a display list, `view_depth` is expected in view space.
func ComputeSortKey(viewdepth float32) uint32 {
	viewdepthToC := C.float(viewdepth)
	retval := C.WrapComputeSortKey(viewdepthToC)
	return uint32(retval)
}

// ComputeSortKeyFromWorld Compute a sorting key to control the rendering order of a display list.
func ComputeSortKeyFromWorld(T *Vec3, view *Mat4) uint32 {
	TToC := T.h
	viewToC := view.h
	retval := C.WrapComputeSortKeyFromWorld(TToC, viewToC)
	return uint32(retval)
}

// ComputeSortKeyFromWorldWithModel Compute a sorting key to control the rendering order of a display list.
func ComputeSortKeyFromWorldWithModel(T *Vec3, view *Mat4, model *Mat4) uint32 {
	TToC := T.h
	viewToC := view.h
	modelToC := model.h
	retval := C.WrapComputeSortKeyFromWorldWithModel(TToC, viewToC, modelToC)
	return uint32(retval)
}

// LoadModelFromFile Load a render model from the local filesystem.
func LoadModelFromFile(path string) *Model {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadModelFromFile(pathToC)
	retvalGO := &Model{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Model) {
		C.WrapModelFree(cleanval.h)
	})
	return retvalGO
}

// LoadModelFromAssets Load a render model from the assets system.  See [harfang.DrawModel] and [harfang.man.Assets].
func LoadModelFromAssets(name string) *Model {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapLoadModelFromAssets(nameToC)
	retvalGO := &Model{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Model) {
		C.WrapModelFree(cleanval.h)
	})
	return retvalGO
}

// CreateCubeModel Create a cube render model.  See [harfang.CreateCubeModel], [harfang.CreateConeModel], [harfang.CreateCylinderModel], [harfang.CreatePlaneModel], [harfang.CreateSphereModel] and [harfang.DrawModel].
func CreateCubeModel(decl *VertexLayout, x float32, y float32, z float32) *Model {
	declToC := decl.h
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	retval := C.WrapCreateCubeModel(declToC, xToC, yToC, zToC)
	retvalGO := &Model{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Model) {
		C.WrapModelFree(cleanval.h)
	})
	return retvalGO
}

// CreateSphereModel Create a sphere render model.  See [harfang.CreateCubeModel], [harfang.CreateConeModel], [harfang.CreateCylinderModel], [harfang.CreatePlaneModel], [harfang.CreateSphereModel] and [harfang.DrawModel].
func CreateSphereModel(decl *VertexLayout, radius float32, subdivx int32, subdivy int32) *Model {
	declToC := decl.h
	radiusToC := C.float(radius)
	subdivxToC := C.int32_t(subdivx)
	subdivyToC := C.int32_t(subdivy)
	retval := C.WrapCreateSphereModel(declToC, radiusToC, subdivxToC, subdivyToC)
	retvalGO := &Model{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Model) {
		C.WrapModelFree(cleanval.h)
	})
	return retvalGO
}

// CreatePlaneModel Create a plane render model.
func CreatePlaneModel(decl *VertexLayout, width float32, length float32, subdivx int32, subdivz int32) *Model {
	declToC := decl.h
	widthToC := C.float(width)
	lengthToC := C.float(length)
	subdivxToC := C.int32_t(subdivx)
	subdivzToC := C.int32_t(subdivz)
	retval := C.WrapCreatePlaneModel(declToC, widthToC, lengthToC, subdivxToC, subdivzToC)
	retvalGO := &Model{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Model) {
		C.WrapModelFree(cleanval.h)
	})
	return retvalGO
}

// CreateCylinderModel Create a cylinder render model.  See [harfang.CreateCubeModel], [harfang.CreateConeModel], [harfang.CreateCylinderModel], [harfang.CreatePlaneModel], [harfang.CreateSphereModel] and [harfang.DrawModel].
func CreateCylinderModel(decl *VertexLayout, radius float32, height float32, subdivx int32) *Model {
	declToC := decl.h
	radiusToC := C.float(radius)
	heightToC := C.float(height)
	subdivxToC := C.int32_t(subdivx)
	retval := C.WrapCreateCylinderModel(declToC, radiusToC, heightToC, subdivxToC)
	retvalGO := &Model{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Model) {
		C.WrapModelFree(cleanval.h)
	})
	return retvalGO
}

// CreateConeModel Create a cone render model.  See [harfang.CreateCubeModel], [harfang.CreateConeModel], [harfang.CreateCylinderModel], [harfang.CreatePlaneModel], [harfang.CreateSphereModel] and [harfang.DrawModel].
func CreateConeModel(decl *VertexLayout, radius float32, height float32, subdivx int32) *Model {
	declToC := decl.h
	radiusToC := C.float(radius)
	heightToC := C.float(height)
	subdivxToC := C.int32_t(subdivx)
	retval := C.WrapCreateConeModel(declToC, radiusToC, heightToC, subdivxToC)
	retvalGO := &Model{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Model) {
		C.WrapModelFree(cleanval.h)
	})
	return retvalGO
}

// CreateCapsuleModel Create a capsule render model.  See [harfang.CreateCubeModel], [harfang.CreateConeModel], [harfang.CreateCylinderModel], [harfang.CreatePlaneModel], [harfang.CreateSphereModel] and [harfang.DrawModel].
func CreateCapsuleModel(decl *VertexLayout, radius float32, height float32, subdivx int32, subdivy int32) *Model {
	declToC := decl.h
	radiusToC := C.float(radius)
	heightToC := C.float(height)
	subdivxToC := C.int32_t(subdivx)
	subdivyToC := C.int32_t(subdivy)
	retval := C.WrapCreateCapsuleModel(declToC, radiusToC, heightToC, subdivxToC, subdivyToC)
	retvalGO := &Model{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Model) {
		C.WrapModelFree(cleanval.h)
	})
	return retvalGO
}

// DrawModel Draw a model to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawModel(viewid uint16, mdl *Model, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, matrix *Mat4) {
	viewidToC := C.ushort(viewid)
	mdlToC := mdl.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	matrixToC := matrix.h
	C.WrapDrawModel(viewidToC, mdlToC, prgToC, valuesToC, texturesToC, matrixToC)
}

// DrawModelWithRenderState Draw a model to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawModelWithRenderState(viewid uint16, mdl *Model, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, matrix *Mat4, renderstate *RenderState) {
	viewidToC := C.ushort(viewid)
	mdlToC := mdl.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	matrixToC := matrix.h
	renderstateToC := renderstate.h
	C.WrapDrawModelWithRenderState(viewidToC, mdlToC, prgToC, valuesToC, texturesToC, matrixToC, renderstateToC)
}

// DrawModelWithRenderStateDepth Draw a model to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawModelWithRenderStateDepth(viewid uint16, mdl *Model, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, matrix *Mat4, renderstate *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	mdlToC := mdl.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	matrixToC := matrix.h
	renderstateToC := renderstate.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawModelWithRenderStateDepth(viewidToC, mdlToC, prgToC, valuesToC, texturesToC, matrixToC, renderstateToC, depthToC)
}

// DrawModelWithMatrices Draw a model to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawModelWithMatrices(viewid uint16, mdl *Model, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, matrices *Mat4List) {
	viewidToC := C.ushort(viewid)
	mdlToC := mdl.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	matricesToC := matrices.h
	C.WrapDrawModelWithMatrices(viewidToC, mdlToC, prgToC, valuesToC, texturesToC, matricesToC)
}

// DrawModelWithMatricesRenderState Draw a model to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawModelWithMatricesRenderState(viewid uint16, mdl *Model, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, matrices *Mat4List, renderstate *RenderState) {
	viewidToC := C.ushort(viewid)
	mdlToC := mdl.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	matricesToC := matrices.h
	renderstateToC := renderstate.h
	C.WrapDrawModelWithMatricesRenderState(viewidToC, mdlToC, prgToC, valuesToC, texturesToC, matricesToC, renderstateToC)
}

// DrawModelWithMatricesRenderStateDepth Draw a model to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawModelWithMatricesRenderStateDepth(viewid uint16, mdl *Model, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, matrices *Mat4List, renderstate *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	mdlToC := mdl.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	matricesToC := matrices.h
	renderstateToC := renderstate.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawModelWithMatricesRenderStateDepth(viewidToC, mdlToC, prgToC, valuesToC, texturesToC, matricesToC, renderstateToC, depthToC)
}

// DrawModelWithSliceOfValuesSliceOfTextures Draw a model to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawModelWithSliceOfValuesSliceOfTextures(viewid uint16, mdl *Model, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, matrix *Mat4) {
	viewidToC := C.ushort(viewid)
	mdlToC := mdl.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	matrixToC := matrix.h
	C.WrapDrawModelWithSliceOfValuesSliceOfTextures(viewidToC, mdlToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, matrixToC)
}

// DrawModelWithSliceOfValuesSliceOfTexturesRenderState Draw a model to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawModelWithSliceOfValuesSliceOfTexturesRenderState(viewid uint16, mdl *Model, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, matrix *Mat4, renderstate *RenderState) {
	viewidToC := C.ushort(viewid)
	mdlToC := mdl.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	matrixToC := matrix.h
	renderstateToC := renderstate.h
	C.WrapDrawModelWithSliceOfValuesSliceOfTexturesRenderState(viewidToC, mdlToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, matrixToC, renderstateToC)
}

// DrawModelWithSliceOfValuesSliceOfTexturesRenderStateDepth Draw a model to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawModelWithSliceOfValuesSliceOfTexturesRenderStateDepth(viewid uint16, mdl *Model, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, matrix *Mat4, renderstate *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	mdlToC := mdl.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	matrixToC := matrix.h
	renderstateToC := renderstate.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawModelWithSliceOfValuesSliceOfTexturesRenderStateDepth(viewidToC, mdlToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, matrixToC, renderstateToC, depthToC)
}

// DrawModelWithSliceOfValuesSliceOfTexturesSliceOfMatrices Draw a model to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawModelWithSliceOfValuesSliceOfTexturesSliceOfMatrices(viewid uint16, mdl *Model, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, SliceOfmatrices GoSliceOfMat4) {
	viewidToC := C.ushort(viewid)
	mdlToC := mdl.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	var SliceOfmatricesPointer []C.WrapMat4
	for _, s := range SliceOfmatrices {
		SliceOfmatricesPointer = append(SliceOfmatricesPointer, s.h)
	}
	SliceOfmatricesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfmatricesPointer))
	SliceOfmatricesPointerToCSize := C.size_t(SliceOfmatricesPointerToC.Len)
	SliceOfmatricesPointerToCBuf := (*C.WrapMat4)(unsafe.Pointer(SliceOfmatricesPointerToC.Data))
	C.WrapDrawModelWithSliceOfValuesSliceOfTexturesSliceOfMatrices(viewidToC, mdlToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, SliceOfmatricesPointerToCSize, SliceOfmatricesPointerToCBuf)
}

// DrawModelWithSliceOfValuesSliceOfTexturesSliceOfMatricesRenderState Draw a model to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawModelWithSliceOfValuesSliceOfTexturesSliceOfMatricesRenderState(viewid uint16, mdl *Model, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, SliceOfmatrices GoSliceOfMat4, renderstate *RenderState) {
	viewidToC := C.ushort(viewid)
	mdlToC := mdl.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	var SliceOfmatricesPointer []C.WrapMat4
	for _, s := range SliceOfmatrices {
		SliceOfmatricesPointer = append(SliceOfmatricesPointer, s.h)
	}
	SliceOfmatricesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfmatricesPointer))
	SliceOfmatricesPointerToCSize := C.size_t(SliceOfmatricesPointerToC.Len)
	SliceOfmatricesPointerToCBuf := (*C.WrapMat4)(unsafe.Pointer(SliceOfmatricesPointerToC.Data))
	renderstateToC := renderstate.h
	C.WrapDrawModelWithSliceOfValuesSliceOfTexturesSliceOfMatricesRenderState(viewidToC, mdlToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, SliceOfmatricesPointerToCSize, SliceOfmatricesPointerToCBuf, renderstateToC)
}

// DrawModelWithSliceOfValuesSliceOfTexturesSliceOfMatricesRenderStateDepth Draw a model to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawModelWithSliceOfValuesSliceOfTexturesSliceOfMatricesRenderStateDepth(viewid uint16, mdl *Model, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, SliceOfmatrices GoSliceOfMat4, renderstate *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	mdlToC := mdl.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	var SliceOfmatricesPointer []C.WrapMat4
	for _, s := range SliceOfmatrices {
		SliceOfmatricesPointer = append(SliceOfmatricesPointer, s.h)
	}
	SliceOfmatricesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfmatricesPointer))
	SliceOfmatricesPointerToCSize := C.size_t(SliceOfmatricesPointerToC.Len)
	SliceOfmatricesPointerToCBuf := (*C.WrapMat4)(unsafe.Pointer(SliceOfmatricesPointerToC.Data))
	renderstateToC := renderstate.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawModelWithSliceOfValuesSliceOfTexturesSliceOfMatricesRenderStateDepth(viewidToC, mdlToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, SliceOfmatricesPointerToCSize, SliceOfmatricesPointerToCBuf, renderstateToC, depthToC)
}

// UpdateMaterialPipelineProgramVariant Select the proper pipeline program variant for the current material state.
func UpdateMaterialPipelineProgramVariant(mat *Material, resources *PipelineResources) {
	matToC := mat.h
	resourcesToC := resources.h
	C.WrapUpdateMaterialPipelineProgramVariant(matToC, resourcesToC)
}

// CreateMissingMaterialProgramValuesFromFile This function scans the material program uniforms and creates a corresponding entry in the material if missing.  Resources are loaded from the local filesystem if a default uniform value requires it.
func CreateMissingMaterialProgramValuesFromFile(mat *Material, resources *PipelineResources) {
	matToC := mat.h
	resourcesToC := resources.h
	C.WrapCreateMissingMaterialProgramValuesFromFile(matToC, resourcesToC)
}

// CreateMissingMaterialProgramValuesFromAssets This function scans the material program uniforms and creates a corresponding entry in the material if missing.  Resources are loaded from the asset system if a default uniform value requires it.  See [harfang.man.Assets].
func CreateMissingMaterialProgramValuesFromAssets(mat *Material, resources *PipelineResources) {
	matToC := mat.h
	resourcesToC := resources.h
	C.WrapCreateMissingMaterialProgramValuesFromAssets(matToC, resourcesToC)
}

// CreateFrameBuffer Create a framebuffer and its texture attachments.  See [harfang.DestroyFrameBuffer].
func CreateFrameBuffer(color *Texture, depth *Texture, name string) *FrameBuffer {
	colorToC := color.h
	depthToC := depth.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapCreateFrameBuffer(colorToC, depthToC, nameToC)
	retvalGO := &FrameBuffer{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *FrameBuffer) {
		C.WrapFrameBufferFree(cleanval.h)
	})
	return retvalGO
}

// CreateFrameBufferWithColorFormatDepthFormatAaName Create a framebuffer and its texture attachments.  See [harfang.DestroyFrameBuffer].
func CreateFrameBufferWithColorFormatDepthFormatAaName(colorformat TextureFormat, depthformat TextureFormat, aa int32, name string) *FrameBuffer {
	colorformatToC := C.int32_t(colorformat)
	depthformatToC := C.int32_t(depthformat)
	aaToC := C.int32_t(aa)
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapCreateFrameBufferWithColorFormatDepthFormatAaName(colorformatToC, depthformatToC, aaToC, nameToC)
	retvalGO := &FrameBuffer{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *FrameBuffer) {
		C.WrapFrameBufferFree(cleanval.h)
	})
	return retvalGO
}

// CreateFrameBufferWithWidthHeightColorFormatDepthFormatAaName Create a framebuffer and its texture attachments.  See [harfang.DestroyFrameBuffer].
func CreateFrameBufferWithWidthHeightColorFormatDepthFormatAaName(width int32, height int32, colorformat TextureFormat, depthformat TextureFormat, aa int32, name string) *FrameBuffer {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	colorformatToC := C.int32_t(colorformat)
	depthformatToC := C.int32_t(depthformat)
	aaToC := C.int32_t(aa)
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapCreateFrameBufferWithWidthHeightColorFormatDepthFormatAaName(widthToC, heightToC, colorformatToC, depthformatToC, aaToC, nameToC)
	retvalGO := &FrameBuffer{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *FrameBuffer) {
		C.WrapFrameBufferFree(cleanval.h)
	})
	return retvalGO
}

// GetColorTexture Retrieves color texture attachment.
func GetColorTexture(frameBuffer *FrameBuffer) *Texture {
	frameBufferToC := frameBuffer.h
	retval := C.WrapGetColorTexture(frameBufferToC)
	retvalGO := &Texture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Texture) {
		C.WrapTextureFree(cleanval.h)
	})
	return retvalGO
}

// GetDepthTexture Retrieves depth texture attachment.
func GetDepthTexture(frameBuffer *FrameBuffer) *Texture {
	frameBufferToC := frameBuffer.h
	retval := C.WrapGetDepthTexture(frameBufferToC)
	retvalGO := &Texture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Texture) {
		C.WrapTextureFree(cleanval.h)
	})
	return retvalGO
}

// GetTextures Returns color and depth texture attachments.
func GetTextures(framebuffer *FrameBuffer) (*Texture, *Texture) {
	framebufferToC := framebuffer.h
	color := NewTexture()
	colorToC := color.h
	depth := NewTexture()
	depthToC := depth.h
	C.WrapGetTextures(framebufferToC, colorToC, depthToC)
	return color, depth
}

// DestroyFrameBuffer Destroy a frame buffer and its resources.
func DestroyFrameBuffer(frameBuffer *FrameBuffer) {
	frameBufferToC := frameBuffer.h
	C.WrapDestroyFrameBuffer(frameBufferToC)
}

// SetTransform Set the model matrix for the next drawn primitive.  If not called, model will be rendered with the identity model matrix.
func SetTransform(mtx *Mat4) {
	mtxToC := mtx.h
	C.WrapSetTransform(mtxToC)
}

// DrawLines Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLines(viewid uint16, vtx *Vertices, prg *ProgramHandle) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	C.WrapDrawLines(viewidToC, vtxToC, prgToC)
}

// DrawLinesWithRenderState Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithRenderState(viewid uint16, vtx *Vertices, prg *ProgramHandle, renderstate *RenderState) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	renderstateToC := renderstate.h
	C.WrapDrawLinesWithRenderState(viewidToC, vtxToC, prgToC, renderstateToC)
}

// DrawLinesWithRenderStateDepth Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithRenderStateDepth(viewid uint16, vtx *Vertices, prg *ProgramHandle, renderstate *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	renderstateToC := renderstate.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawLinesWithRenderStateDepth(viewidToC, vtxToC, prgToC, renderstateToC, depthToC)
}

// DrawLinesWithValuesTextures Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithValuesTextures(viewid uint16, vtx *Vertices, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	C.WrapDrawLinesWithValuesTextures(viewidToC, vtxToC, prgToC, valuesToC, texturesToC)
}

// DrawLinesWithValuesTexturesRenderState Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithValuesTexturesRenderState(viewid uint16, vtx *Vertices, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, renderstate *RenderState) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	renderstateToC := renderstate.h
	C.WrapDrawLinesWithValuesTexturesRenderState(viewidToC, vtxToC, prgToC, valuesToC, texturesToC, renderstateToC)
}

// DrawLinesWithValuesTexturesRenderStateDepth Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithValuesTexturesRenderStateDepth(viewid uint16, vtx *Vertices, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, renderstate *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	renderstateToC := renderstate.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawLinesWithValuesTexturesRenderStateDepth(viewidToC, vtxToC, prgToC, valuesToC, texturesToC, renderstateToC, depthToC)
}

// DrawLinesWithIdxVtxPrgValuesTextures Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithIdxVtxPrgValuesTextures(viewid uint16, idx *Uint16TList, vtx *Vertices, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList) {
	viewidToC := C.ushort(viewid)
	idxToC := idx.h
	vtxToC := vtx.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	C.WrapDrawLinesWithIdxVtxPrgValuesTextures(viewidToC, idxToC, vtxToC, prgToC, valuesToC, texturesToC)
}

// DrawLinesWithIdxVtxPrgValuesTexturesRenderState Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithIdxVtxPrgValuesTexturesRenderState(viewid uint16, idx *Uint16TList, vtx *Vertices, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, renderstate *RenderState) {
	viewidToC := C.ushort(viewid)
	idxToC := idx.h
	vtxToC := vtx.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	renderstateToC := renderstate.h
	C.WrapDrawLinesWithIdxVtxPrgValuesTexturesRenderState(viewidToC, idxToC, vtxToC, prgToC, valuesToC, texturesToC, renderstateToC)
}

// DrawLinesWithIdxVtxPrgValuesTexturesRenderStateDepth Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithIdxVtxPrgValuesTexturesRenderStateDepth(viewid uint16, idx *Uint16TList, vtx *Vertices, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, renderstate *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	idxToC := idx.h
	vtxToC := vtx.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	renderstateToC := renderstate.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawLinesWithIdxVtxPrgValuesTexturesRenderStateDepth(viewidToC, idxToC, vtxToC, prgToC, valuesToC, texturesToC, renderstateToC, depthToC)
}

// DrawLinesWithSliceOfValuesSliceOfTextures Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithSliceOfValuesSliceOfTextures(viewid uint16, vtx *Vertices, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	C.WrapDrawLinesWithSliceOfValuesSliceOfTextures(viewidToC, vtxToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf)
}

// DrawLinesWithSliceOfValuesSliceOfTexturesRenderState Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithSliceOfValuesSliceOfTexturesRenderState(viewid uint16, vtx *Vertices, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, renderstate *RenderState) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	renderstateToC := renderstate.h
	C.WrapDrawLinesWithSliceOfValuesSliceOfTexturesRenderState(viewidToC, vtxToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, renderstateToC)
}

// DrawLinesWithSliceOfValuesSliceOfTexturesRenderStateDepth Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithSliceOfValuesSliceOfTexturesRenderStateDepth(viewid uint16, vtx *Vertices, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, renderstate *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	renderstateToC := renderstate.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawLinesWithSliceOfValuesSliceOfTexturesRenderStateDepth(viewidToC, vtxToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, renderstateToC, depthToC)
}

// DrawLinesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTextures Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTextures(viewid uint16, SliceOfidx GoSliceOfuint16T, vtx *Vertices, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture) {
	viewidToC := C.ushort(viewid)
	SliceOfidxToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfidx))
	SliceOfidxToCSize := C.size_t(SliceOfidxToC.Len)
	SliceOfidxToCBuf := (*C.ushort)(unsafe.Pointer(SliceOfidxToC.Data))
	vtxToC := vtx.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	C.WrapDrawLinesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTextures(viewidToC, SliceOfidxToCSize, SliceOfidxToCBuf, vtxToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf)
}

// DrawLinesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesRenderState Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesRenderState(viewid uint16, SliceOfidx GoSliceOfuint16T, vtx *Vertices, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, renderstate *RenderState) {
	viewidToC := C.ushort(viewid)
	SliceOfidxToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfidx))
	SliceOfidxToCSize := C.size_t(SliceOfidxToC.Len)
	SliceOfidxToCBuf := (*C.ushort)(unsafe.Pointer(SliceOfidxToC.Data))
	vtxToC := vtx.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	renderstateToC := renderstate.h
	C.WrapDrawLinesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesRenderState(viewidToC, SliceOfidxToCSize, SliceOfidxToCBuf, vtxToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, renderstateToC)
}

// DrawLinesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesRenderStateDepth Draw a list of lines to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawLinesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesRenderStateDepth(viewid uint16, SliceOfidx GoSliceOfuint16T, vtx *Vertices, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, renderstate *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	SliceOfidxToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfidx))
	SliceOfidxToCSize := C.size_t(SliceOfidxToC.Len)
	SliceOfidxToCBuf := (*C.ushort)(unsafe.Pointer(SliceOfidxToC.Data))
	vtxToC := vtx.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	renderstateToC := renderstate.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawLinesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesRenderStateDepth(viewidToC, SliceOfidxToCSize, SliceOfidxToCBuf, vtxToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, renderstateToC, depthToC)
}

// DrawTriangles Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTriangles(viewid uint16, vtx *Vertices, prg *ProgramHandle) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	C.WrapDrawTriangles(viewidToC, vtxToC, prgToC)
}

// DrawTrianglesWithState Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithState(viewid uint16, vtx *Vertices, prg *ProgramHandle, state *RenderState) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	stateToC := state.h
	C.WrapDrawTrianglesWithState(viewidToC, vtxToC, prgToC, stateToC)
}

// DrawTrianglesWithStateDepth Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithStateDepth(viewid uint16, vtx *Vertices, prg *ProgramHandle, state *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	stateToC := state.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawTrianglesWithStateDepth(viewidToC, vtxToC, prgToC, stateToC, depthToC)
}

// DrawTrianglesWithValuesTextures Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithValuesTextures(viewid uint16, vtx *Vertices, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	C.WrapDrawTrianglesWithValuesTextures(viewidToC, vtxToC, prgToC, valuesToC, texturesToC)
}

// DrawTrianglesWithValuesTexturesState Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithValuesTexturesState(viewid uint16, vtx *Vertices, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, state *RenderState) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	stateToC := state.h
	C.WrapDrawTrianglesWithValuesTexturesState(viewidToC, vtxToC, prgToC, valuesToC, texturesToC, stateToC)
}

// DrawTrianglesWithValuesTexturesStateDepth Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithValuesTexturesStateDepth(viewid uint16, vtx *Vertices, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, state *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	stateToC := state.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawTrianglesWithValuesTexturesStateDepth(viewidToC, vtxToC, prgToC, valuesToC, texturesToC, stateToC, depthToC)
}

// DrawTrianglesWithIdxVtxPrgValuesTextures Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithIdxVtxPrgValuesTextures(viewid uint16, idx *Uint16TList, vtx *Vertices, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList) {
	viewidToC := C.ushort(viewid)
	idxToC := idx.h
	vtxToC := vtx.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	C.WrapDrawTrianglesWithIdxVtxPrgValuesTextures(viewidToC, idxToC, vtxToC, prgToC, valuesToC, texturesToC)
}

// DrawTrianglesWithIdxVtxPrgValuesTexturesState Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithIdxVtxPrgValuesTexturesState(viewid uint16, idx *Uint16TList, vtx *Vertices, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, state *RenderState) {
	viewidToC := C.ushort(viewid)
	idxToC := idx.h
	vtxToC := vtx.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	stateToC := state.h
	C.WrapDrawTrianglesWithIdxVtxPrgValuesTexturesState(viewidToC, idxToC, vtxToC, prgToC, valuesToC, texturesToC, stateToC)
}

// DrawTrianglesWithIdxVtxPrgValuesTexturesStateDepth Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithIdxVtxPrgValuesTexturesStateDepth(viewid uint16, idx *Uint16TList, vtx *Vertices, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, state *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	idxToC := idx.h
	vtxToC := vtx.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	stateToC := state.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawTrianglesWithIdxVtxPrgValuesTexturesStateDepth(viewidToC, idxToC, vtxToC, prgToC, valuesToC, texturesToC, stateToC, depthToC)
}

// DrawTrianglesWithSliceOfValuesSliceOfTextures Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithSliceOfValuesSliceOfTextures(viewid uint16, vtx *Vertices, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	C.WrapDrawTrianglesWithSliceOfValuesSliceOfTextures(viewidToC, vtxToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf)
}

// DrawTrianglesWithSliceOfValuesSliceOfTexturesState Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithSliceOfValuesSliceOfTexturesState(viewid uint16, vtx *Vertices, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, state *RenderState) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	stateToC := state.h
	C.WrapDrawTrianglesWithSliceOfValuesSliceOfTexturesState(viewidToC, vtxToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, stateToC)
}

// DrawTrianglesWithSliceOfValuesSliceOfTexturesStateDepth Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithSliceOfValuesSliceOfTexturesStateDepth(viewid uint16, vtx *Vertices, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, state *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	vtxToC := vtx.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	stateToC := state.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawTrianglesWithSliceOfValuesSliceOfTexturesStateDepth(viewidToC, vtxToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, stateToC, depthToC)
}

// DrawTrianglesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTextures Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTextures(viewid uint16, SliceOfidx GoSliceOfuint16T, vtx *Vertices, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture) {
	viewidToC := C.ushort(viewid)
	SliceOfidxToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfidx))
	SliceOfidxToCSize := C.size_t(SliceOfidxToC.Len)
	SliceOfidxToCBuf := (*C.ushort)(unsafe.Pointer(SliceOfidxToC.Data))
	vtxToC := vtx.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	C.WrapDrawTrianglesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTextures(viewidToC, SliceOfidxToCSize, SliceOfidxToCBuf, vtxToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf)
}

// DrawTrianglesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesState Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesState(viewid uint16, SliceOfidx GoSliceOfuint16T, vtx *Vertices, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, state *RenderState) {
	viewidToC := C.ushort(viewid)
	SliceOfidxToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfidx))
	SliceOfidxToCSize := C.size_t(SliceOfidxToC.Len)
	SliceOfidxToCBuf := (*C.ushort)(unsafe.Pointer(SliceOfidxToC.Data))
	vtxToC := vtx.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	stateToC := state.h
	C.WrapDrawTrianglesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesState(viewidToC, SliceOfidxToCSize, SliceOfidxToCBuf, vtxToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, stateToC)
}

// DrawTrianglesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesStateDepth Draw a list of triangles to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.
func DrawTrianglesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesStateDepth(viewid uint16, SliceOfidx GoSliceOfuint16T, vtx *Vertices, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, state *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	SliceOfidxToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfidx))
	SliceOfidxToCSize := C.size_t(SliceOfidxToC.Len)
	SliceOfidxToCBuf := (*C.ushort)(unsafe.Pointer(SliceOfidxToC.Data))
	vtxToC := vtx.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	stateToC := state.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawTrianglesWithSliceOfIdxVtxPrgSliceOfValuesSliceOfTexturesStateDepth(viewidToC, SliceOfidxToCSize, SliceOfidxToCBuf, vtxToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, stateToC, depthToC)
}

// DrawSprites Draw a list of sprites to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.  *Note:* This function prepares the sprite on the CPU before submitting them all to the GPU as a single draw call.
func DrawSprites(viewid uint16, invviewR *Mat3, vtxlayout *VertexLayout, pos *Vec3List, size *Vec2, prg *ProgramHandle) {
	viewidToC := C.ushort(viewid)
	invviewRToC := invviewR.h
	vtxlayoutToC := vtxlayout.h
	posToC := pos.h
	sizeToC := size.h
	prgToC := prg.h
	C.WrapDrawSprites(viewidToC, invviewRToC, vtxlayoutToC, posToC, sizeToC, prgToC)
}

// DrawSpritesWithState Draw a list of sprites to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.  *Note:* This function prepares the sprite on the CPU before submitting them all to the GPU as a single draw call.
func DrawSpritesWithState(viewid uint16, invviewR *Mat3, vtxlayout *VertexLayout, pos *Vec3List, size *Vec2, prg *ProgramHandle, state *RenderState) {
	viewidToC := C.ushort(viewid)
	invviewRToC := invviewR.h
	vtxlayoutToC := vtxlayout.h
	posToC := pos.h
	sizeToC := size.h
	prgToC := prg.h
	stateToC := state.h
	C.WrapDrawSpritesWithState(viewidToC, invviewRToC, vtxlayoutToC, posToC, sizeToC, prgToC, stateToC)
}

// DrawSpritesWithStateDepth Draw a list of sprites to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.  *Note:* This function prepares the sprite on the CPU before submitting them all to the GPU as a single draw call.
func DrawSpritesWithStateDepth(viewid uint16, invviewR *Mat3, vtxlayout *VertexLayout, pos *Vec3List, size *Vec2, prg *ProgramHandle, state *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	invviewRToC := invviewR.h
	vtxlayoutToC := vtxlayout.h
	posToC := pos.h
	sizeToC := size.h
	prgToC := prg.h
	stateToC := state.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawSpritesWithStateDepth(viewidToC, invviewRToC, vtxlayoutToC, posToC, sizeToC, prgToC, stateToC, depthToC)
}

// DrawSpritesWithValuesTextures Draw a list of sprites to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.  *Note:* This function prepares the sprite on the CPU before submitting them all to the GPU as a single draw call.
func DrawSpritesWithValuesTextures(viewid uint16, invviewR *Mat3, vtxlayout *VertexLayout, pos *Vec3List, size *Vec2, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList) {
	viewidToC := C.ushort(viewid)
	invviewRToC := invviewR.h
	vtxlayoutToC := vtxlayout.h
	posToC := pos.h
	sizeToC := size.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	C.WrapDrawSpritesWithValuesTextures(viewidToC, invviewRToC, vtxlayoutToC, posToC, sizeToC, prgToC, valuesToC, texturesToC)
}

// DrawSpritesWithValuesTexturesState Draw a list of sprites to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.  *Note:* This function prepares the sprite on the CPU before submitting them all to the GPU as a single draw call.
func DrawSpritesWithValuesTexturesState(viewid uint16, invviewR *Mat3, vtxlayout *VertexLayout, pos *Vec3List, size *Vec2, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, state *RenderState) {
	viewidToC := C.ushort(viewid)
	invviewRToC := invviewR.h
	vtxlayoutToC := vtxlayout.h
	posToC := pos.h
	sizeToC := size.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	stateToC := state.h
	C.WrapDrawSpritesWithValuesTexturesState(viewidToC, invviewRToC, vtxlayoutToC, posToC, sizeToC, prgToC, valuesToC, texturesToC, stateToC)
}

// DrawSpritesWithValuesTexturesStateDepth Draw a list of sprites to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.  *Note:* This function prepares the sprite on the CPU before submitting them all to the GPU as a single draw call.
func DrawSpritesWithValuesTexturesStateDepth(viewid uint16, invviewR *Mat3, vtxlayout *VertexLayout, pos *Vec3List, size *Vec2, prg *ProgramHandle, values *UniformSetValueList, textures *UniformSetTextureList, state *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	invviewRToC := invviewR.h
	vtxlayoutToC := vtxlayout.h
	posToC := pos.h
	sizeToC := size.h
	prgToC := prg.h
	valuesToC := values.h
	texturesToC := textures.h
	stateToC := state.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawSpritesWithValuesTexturesStateDepth(viewidToC, invviewRToC, vtxlayoutToC, posToC, sizeToC, prgToC, valuesToC, texturesToC, stateToC, depthToC)
}

// DrawSpritesWithSliceOfPos Draw a list of sprites to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.  *Note:* This function prepares the sprite on the CPU before submitting them all to the GPU as a single draw call.
func DrawSpritesWithSliceOfPos(viewid uint16, invviewR *Mat3, vtxlayout *VertexLayout, SliceOfpos GoSliceOfVec3, size *Vec2, prg *ProgramHandle) {
	viewidToC := C.ushort(viewid)
	invviewRToC := invviewR.h
	vtxlayoutToC := vtxlayout.h
	var SliceOfposPointer []C.WrapVec3
	for _, s := range SliceOfpos {
		SliceOfposPointer = append(SliceOfposPointer, s.h)
	}
	SliceOfposPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfposPointer))
	SliceOfposPointerToCSize := C.size_t(SliceOfposPointerToC.Len)
	SliceOfposPointerToCBuf := (*C.WrapVec3)(unsafe.Pointer(SliceOfposPointerToC.Data))
	sizeToC := size.h
	prgToC := prg.h
	C.WrapDrawSpritesWithSliceOfPos(viewidToC, invviewRToC, vtxlayoutToC, SliceOfposPointerToCSize, SliceOfposPointerToCBuf, sizeToC, prgToC)
}

// DrawSpritesWithSliceOfPosState Draw a list of sprites to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.  *Note:* This function prepares the sprite on the CPU before submitting them all to the GPU as a single draw call.
func DrawSpritesWithSliceOfPosState(viewid uint16, invviewR *Mat3, vtxlayout *VertexLayout, SliceOfpos GoSliceOfVec3, size *Vec2, prg *ProgramHandle, state *RenderState) {
	viewidToC := C.ushort(viewid)
	invviewRToC := invviewR.h
	vtxlayoutToC := vtxlayout.h
	var SliceOfposPointer []C.WrapVec3
	for _, s := range SliceOfpos {
		SliceOfposPointer = append(SliceOfposPointer, s.h)
	}
	SliceOfposPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfposPointer))
	SliceOfposPointerToCSize := C.size_t(SliceOfposPointerToC.Len)
	SliceOfposPointerToCBuf := (*C.WrapVec3)(unsafe.Pointer(SliceOfposPointerToC.Data))
	sizeToC := size.h
	prgToC := prg.h
	stateToC := state.h
	C.WrapDrawSpritesWithSliceOfPosState(viewidToC, invviewRToC, vtxlayoutToC, SliceOfposPointerToCSize, SliceOfposPointerToCBuf, sizeToC, prgToC, stateToC)
}

// DrawSpritesWithSliceOfPosStateDepth Draw a list of sprites to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.  *Note:* This function prepares the sprite on the CPU before submitting them all to the GPU as a single draw call.
func DrawSpritesWithSliceOfPosStateDepth(viewid uint16, invviewR *Mat3, vtxlayout *VertexLayout, SliceOfpos GoSliceOfVec3, size *Vec2, prg *ProgramHandle, state *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	invviewRToC := invviewR.h
	vtxlayoutToC := vtxlayout.h
	var SliceOfposPointer []C.WrapVec3
	for _, s := range SliceOfpos {
		SliceOfposPointer = append(SliceOfposPointer, s.h)
	}
	SliceOfposPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfposPointer))
	SliceOfposPointerToCSize := C.size_t(SliceOfposPointerToC.Len)
	SliceOfposPointerToCBuf := (*C.WrapVec3)(unsafe.Pointer(SliceOfposPointerToC.Data))
	sizeToC := size.h
	prgToC := prg.h
	stateToC := state.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawSpritesWithSliceOfPosStateDepth(viewidToC, invviewRToC, vtxlayoutToC, SliceOfposPointerToCSize, SliceOfposPointerToCBuf, sizeToC, prgToC, stateToC, depthToC)
}

// DrawSpritesWithSliceOfPosSliceOfValuesSliceOfTextures Draw a list of sprites to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.  *Note:* This function prepares the sprite on the CPU before submitting them all to the GPU as a single draw call.
func DrawSpritesWithSliceOfPosSliceOfValuesSliceOfTextures(viewid uint16, invviewR *Mat3, vtxlayout *VertexLayout, SliceOfpos GoSliceOfVec3, size *Vec2, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture) {
	viewidToC := C.ushort(viewid)
	invviewRToC := invviewR.h
	vtxlayoutToC := vtxlayout.h
	var SliceOfposPointer []C.WrapVec3
	for _, s := range SliceOfpos {
		SliceOfposPointer = append(SliceOfposPointer, s.h)
	}
	SliceOfposPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfposPointer))
	SliceOfposPointerToCSize := C.size_t(SliceOfposPointerToC.Len)
	SliceOfposPointerToCBuf := (*C.WrapVec3)(unsafe.Pointer(SliceOfposPointerToC.Data))
	sizeToC := size.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	C.WrapDrawSpritesWithSliceOfPosSliceOfValuesSliceOfTextures(viewidToC, invviewRToC, vtxlayoutToC, SliceOfposPointerToCSize, SliceOfposPointerToCBuf, sizeToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf)
}

// DrawSpritesWithSliceOfPosSliceOfValuesSliceOfTexturesState Draw a list of sprites to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.  *Note:* This function prepares the sprite on the CPU before submitting them all to the GPU as a single draw call.
func DrawSpritesWithSliceOfPosSliceOfValuesSliceOfTexturesState(viewid uint16, invviewR *Mat3, vtxlayout *VertexLayout, SliceOfpos GoSliceOfVec3, size *Vec2, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, state *RenderState) {
	viewidToC := C.ushort(viewid)
	invviewRToC := invviewR.h
	vtxlayoutToC := vtxlayout.h
	var SliceOfposPointer []C.WrapVec3
	for _, s := range SliceOfpos {
		SliceOfposPointer = append(SliceOfposPointer, s.h)
	}
	SliceOfposPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfposPointer))
	SliceOfposPointerToCSize := C.size_t(SliceOfposPointerToC.Len)
	SliceOfposPointerToCBuf := (*C.WrapVec3)(unsafe.Pointer(SliceOfposPointerToC.Data))
	sizeToC := size.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	stateToC := state.h
	C.WrapDrawSpritesWithSliceOfPosSliceOfValuesSliceOfTexturesState(viewidToC, invviewRToC, vtxlayoutToC, SliceOfposPointerToCSize, SliceOfposPointerToCBuf, sizeToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, stateToC)
}

// DrawSpritesWithSliceOfPosSliceOfValuesSliceOfTexturesStateDepth Draw a list of sprites to the specified view.  Use [harfang.UniformSetValueList] and [harfang.UniformSetTextureList] to pass uniform values to the shader program.  *Note:* This function prepares the sprite on the CPU before submitting them all to the GPU as a single draw call.
func DrawSpritesWithSliceOfPosSliceOfValuesSliceOfTexturesStateDepth(viewid uint16, invviewR *Mat3, vtxlayout *VertexLayout, SliceOfpos GoSliceOfVec3, size *Vec2, prg *ProgramHandle, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, state *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	invviewRToC := invviewR.h
	vtxlayoutToC := vtxlayout.h
	var SliceOfposPointer []C.WrapVec3
	for _, s := range SliceOfpos {
		SliceOfposPointer = append(SliceOfposPointer, s.h)
	}
	SliceOfposPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfposPointer))
	SliceOfposPointerToCSize := C.size_t(SliceOfposPointerToC.Len)
	SliceOfposPointerToCBuf := (*C.WrapVec3)(unsafe.Pointer(SliceOfposPointerToC.Data))
	sizeToC := size.h
	prgToC := prg.h
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	stateToC := state.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawSpritesWithSliceOfPosSliceOfValuesSliceOfTexturesStateDepth(viewidToC, invviewRToC, vtxlayoutToC, SliceOfposPointerToCSize, SliceOfposPointerToCBuf, sizeToC, prgToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, stateToC, depthToC)
}

// GetForwardPipelineInfo Return the pipeline info object for the forward pipeline.
func GetForwardPipelineInfo() *PipelineInfo {
	retval := C.WrapGetForwardPipelineInfo()
	var retvalGO *PipelineInfo
	if retval != nil {
		retvalGO = &PipelineInfo{h: retval}
	}
	return retvalGO
}

// CreateForwardPipeline Create a forward pipeline and its resources.  See [harfang.DestroyForwardPipeline].
func CreateForwardPipeline() *ForwardPipeline {
	retval := C.WrapCreateForwardPipeline()
	retvalGO := &ForwardPipeline{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipeline) {
		C.WrapForwardPipelineFree(cleanval.h)
	})
	return retvalGO
}

// CreateForwardPipelineWithShadowMapResolution Create a forward pipeline and its resources.  See [harfang.DestroyForwardPipeline].
func CreateForwardPipelineWithShadowMapResolution(shadowmapresolution int32) *ForwardPipeline {
	shadowmapresolutionToC := C.int32_t(shadowmapresolution)
	retval := C.WrapCreateForwardPipelineWithShadowMapResolution(shadowmapresolutionToC)
	retvalGO := &ForwardPipeline{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipeline) {
		C.WrapForwardPipelineFree(cleanval.h)
	})
	return retvalGO
}

// CreateForwardPipelineWithShadowMapResolutionSpot16bitShadowMap Create a forward pipeline and its resources.  See [harfang.DestroyForwardPipeline].
func CreateForwardPipelineWithShadowMapResolutionSpot16bitShadowMap(shadowmapresolution int32, spot16bitshadowmap bool) *ForwardPipeline {
	shadowmapresolutionToC := C.int32_t(shadowmapresolution)
	spot16bitshadowmapToC := C.bool(spot16bitshadowmap)
	retval := C.WrapCreateForwardPipelineWithShadowMapResolutionSpot16bitShadowMap(shadowmapresolutionToC, spot16bitshadowmapToC)
	retvalGO := &ForwardPipeline{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipeline) {
		C.WrapForwardPipelineFree(cleanval.h)
	})
	return retvalGO
}

// DestroyForwardPipeline Destroy a forward pipeline object.
func DestroyForwardPipeline(pipeline *ForwardPipeline) {
	pipelineToC := pipeline.h
	C.WrapDestroyForwardPipeline(pipelineToC)
}

// MakeForwardPipelinePointLight Create a forward pipeline point light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelinePointLight(world *Mat4, diffuse *Color, specular *Color) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	retval := C.WrapMakeForwardPipelinePointLight(worldToC, diffuseToC, specularToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelinePointLightWithRadius Create a forward pipeline point light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelinePointLightWithRadius(world *Mat4, diffuse *Color, specular *Color, radius float32) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	radiusToC := C.float(radius)
	retval := C.WrapMakeForwardPipelinePointLightWithRadius(worldToC, diffuseToC, specularToC, radiusToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelinePointLightWithRadiusPriority Create a forward pipeline point light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelinePointLightWithRadiusPriority(world *Mat4, diffuse *Color, specular *Color, radius float32, priority float32) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	radiusToC := C.float(radius)
	priorityToC := C.float(priority)
	retval := C.WrapMakeForwardPipelinePointLightWithRadiusPriority(worldToC, diffuseToC, specularToC, radiusToC, priorityToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelinePointLightWithRadiusPriorityShadowType Create a forward pipeline point light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelinePointLightWithRadiusPriorityShadowType(world *Mat4, diffuse *Color, specular *Color, radius float32, priority float32, shadowtype ForwardPipelineShadowType) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	radiusToC := C.float(radius)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapMakeForwardPipelinePointLightWithRadiusPriorityShadowType(worldToC, diffuseToC, specularToC, radiusToC, priorityToC, shadowtypeToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelinePointLightWithRadiusPriorityShadowTypeShadowBias Create a forward pipeline point light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelinePointLightWithRadiusPriorityShadowTypeShadowBias(world *Mat4, diffuse *Color, specular *Color, radius float32, priority float32, shadowtype ForwardPipelineShadowType, shadowbias float32) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	radiusToC := C.float(radius)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapMakeForwardPipelinePointLightWithRadiusPriorityShadowTypeShadowBias(worldToC, diffuseToC, specularToC, radiusToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelineSpotLight Create a forward pipeline spot light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelineSpotLight(world *Mat4, diffuse *Color, specular *Color) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	retval := C.WrapMakeForwardPipelineSpotLight(worldToC, diffuseToC, specularToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelineSpotLightWithRadius Create a forward pipeline spot light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelineSpotLightWithRadius(world *Mat4, diffuse *Color, specular *Color, radius float32) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	radiusToC := C.float(radius)
	retval := C.WrapMakeForwardPipelineSpotLightWithRadius(worldToC, diffuseToC, specularToC, radiusToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelineSpotLightWithRadiusInnerAngle Create a forward pipeline spot light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelineSpotLightWithRadiusInnerAngle(world *Mat4, diffuse *Color, specular *Color, radius float32, innerangle float32) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	retval := C.WrapMakeForwardPipelineSpotLightWithRadiusInnerAngle(worldToC, diffuseToC, specularToC, radiusToC, innerangleToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAngle Create a forward pipeline spot light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAngle(world *Mat4, diffuse *Color, specular *Color, radius float32, innerangle float32, outerangle float32) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	retval := C.WrapMakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAngle(worldToC, diffuseToC, specularToC, radiusToC, innerangleToC, outerangleToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAnglePriority Create a forward pipeline spot light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAnglePriority(world *Mat4, diffuse *Color, specular *Color, radius float32, innerangle float32, outerangle float32, priority float32) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	priorityToC := C.float(priority)
	retval := C.WrapMakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAnglePriority(worldToC, diffuseToC, specularToC, radiusToC, innerangleToC, outerangleToC, priorityToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAnglePriorityShadowType Create a forward pipeline spot light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAnglePriorityShadowType(world *Mat4, diffuse *Color, specular *Color, radius float32, innerangle float32, outerangle float32, priority float32, shadowtype ForwardPipelineShadowType) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapMakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAnglePriorityShadowType(worldToC, diffuseToC, specularToC, radiusToC, innerangleToC, outerangleToC, priorityToC, shadowtypeToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAnglePriorityShadowTypeShadowBias Create a forward pipeline spot light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAnglePriorityShadowTypeShadowBias(world *Mat4, diffuse *Color, specular *Color, radius float32, innerangle float32, outerangle float32, priority float32, shadowtype ForwardPipelineShadowType, shadowbias float32) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapMakeForwardPipelineSpotLightWithRadiusInnerAngleOuterAnglePriorityShadowTypeShadowBias(worldToC, diffuseToC, specularToC, radiusToC, innerangleToC, outerangleToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelineLinearLight Create a forward pipeline linear light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelineLinearLight(world *Mat4, diffuse *Color, specular *Color) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	retval := C.WrapMakeForwardPipelineLinearLight(worldToC, diffuseToC, specularToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelineLinearLightWithPssmSplit Create a forward pipeline linear light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelineLinearLightWithPssmSplit(world *Mat4, diffuse *Color, specular *Color, pssmsplit *Vec4) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	pssmsplitToC := pssmsplit.h
	retval := C.WrapMakeForwardPipelineLinearLightWithPssmSplit(worldToC, diffuseToC, specularToC, pssmsplitToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelineLinearLightWithPssmSplitPriority Create a forward pipeline linear light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelineLinearLightWithPssmSplitPriority(world *Mat4, diffuse *Color, specular *Color, pssmsplit *Vec4, priority float32) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	pssmsplitToC := pssmsplit.h
	priorityToC := C.float(priority)
	retval := C.WrapMakeForwardPipelineLinearLightWithPssmSplitPriority(worldToC, diffuseToC, specularToC, pssmsplitToC, priorityToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelineLinearLightWithPssmSplitPriorityShadowType Create a forward pipeline linear light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelineLinearLightWithPssmSplitPriorityShadowType(world *Mat4, diffuse *Color, specular *Color, pssmsplit *Vec4, priority float32, shadowtype ForwardPipelineShadowType) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	pssmsplitToC := pssmsplit.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapMakeForwardPipelineLinearLightWithPssmSplitPriorityShadowType(worldToC, diffuseToC, specularToC, pssmsplitToC, priorityToC, shadowtypeToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// MakeForwardPipelineLinearLightWithPssmSplitPriorityShadowTypeShadowBias Create a forward pipeline linear light.  See [harfang.ForwardPipelineLights], [harfang.PrepareForwardPipelineLights] and [harfang.SubmitModelToForwardPipeline].
func MakeForwardPipelineLinearLightWithPssmSplitPriorityShadowTypeShadowBias(world *Mat4, diffuse *Color, specular *Color, pssmsplit *Vec4, priority float32, shadowtype ForwardPipelineShadowType, shadowbias float32) *ForwardPipelineLight {
	worldToC := world.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	pssmsplitToC := pssmsplit.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapMakeForwardPipelineLinearLightWithPssmSplitPriorityShadowTypeShadowBias(worldToC, diffuseToC, specularToC, pssmsplitToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &ForwardPipelineLight{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLight) {
		C.WrapForwardPipelineLightFree(cleanval.h)
	})
	return retvalGO
}

// PrepareForwardPipelineLights Prepare a list of forward pipeline lights into a structure ready for submitting to the forward pipeline.  Lights are sorted by priority/type and the most important lights are assigned to available lighting slot of the forward pipeline.  See [harfang.SubmitModelToForwardPipeline].
func PrepareForwardPipelineLights(lights *ForwardPipelineLightList) *ForwardPipelineLights {
	lightsToC := lights.h
	retval := C.WrapPrepareForwardPipelineLights(lightsToC)
	retvalGO := &ForwardPipelineLights{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLights) {
		C.WrapForwardPipelineLightsFree(cleanval.h)
	})
	return retvalGO
}

// PrepareForwardPipelineLightsWithSliceOfLights Prepare a list of forward pipeline lights into a structure ready for submitting to the forward pipeline.  Lights are sorted by priority/type and the most important lights are assigned to available lighting slot of the forward pipeline.  See [harfang.SubmitModelToForwardPipeline].
func PrepareForwardPipelineLightsWithSliceOfLights(SliceOflights GoSliceOfForwardPipelineLight) *ForwardPipelineLights {
	var SliceOflightsPointer []C.WrapForwardPipelineLight
	for _, s := range SliceOflights {
		SliceOflightsPointer = append(SliceOflightsPointer, s.h)
	}
	SliceOflightsPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOflightsPointer))
	SliceOflightsPointerToCSize := C.size_t(SliceOflightsPointerToC.Len)
	SliceOflightsPointerToCBuf := (*C.WrapForwardPipelineLight)(unsafe.Pointer(SliceOflightsPointerToC.Data))
	retval := C.WrapPrepareForwardPipelineLightsWithSliceOfLights(SliceOflightsPointerToCSize, SliceOflightsPointerToCBuf)
	retvalGO := &ForwardPipelineLights{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLights) {
		C.WrapForwardPipelineLightsFree(cleanval.h)
	})
	return retvalGO
}

// LoadFontFromFile Load a TrueType (TTF) font from the local filesystem.  See [harfang.man.Assets].
func LoadFontFromFile(path string) *Font {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadFontFromFile(pathToC)
	retvalGO := &Font{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Font) {
		C.WrapFontFree(cleanval.h)
	})
	return retvalGO
}

// LoadFontFromFileWithSize Load a TrueType (TTF) font from the local filesystem.  See [harfang.man.Assets].
func LoadFontFromFileWithSize(path string, size float32) *Font {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sizeToC := C.float(size)
	retval := C.WrapLoadFontFromFileWithSize(pathToC, sizeToC)
	retvalGO := &Font{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Font) {
		C.WrapFontFree(cleanval.h)
	})
	return retvalGO
}

// LoadFontFromFileWithSizeResolution Load a TrueType (TTF) font from the local filesystem.  See [harfang.man.Assets].
func LoadFontFromFileWithSizeResolution(path string, size float32, resolution uint16) *Font {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sizeToC := C.float(size)
	resolutionToC := C.ushort(resolution)
	retval := C.WrapLoadFontFromFileWithSizeResolution(pathToC, sizeToC, resolutionToC)
	retvalGO := &Font{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Font) {
		C.WrapFontFree(cleanval.h)
	})
	return retvalGO
}

// LoadFontFromFileWithSizeResolutionPadding Load a TrueType (TTF) font from the local filesystem.  See [harfang.man.Assets].
func LoadFontFromFileWithSizeResolutionPadding(path string, size float32, resolution uint16, padding int32) *Font {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sizeToC := C.float(size)
	resolutionToC := C.ushort(resolution)
	paddingToC := C.int32_t(padding)
	retval := C.WrapLoadFontFromFileWithSizeResolutionPadding(pathToC, sizeToC, resolutionToC, paddingToC)
	retvalGO := &Font{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Font) {
		C.WrapFontFree(cleanval.h)
	})
	return retvalGO
}

// LoadFontFromFileWithSizeResolutionPaddingGlyphs Load a TrueType (TTF) font from the local filesystem.  See [harfang.man.Assets].
func LoadFontFromFileWithSizeResolutionPaddingGlyphs(path string, size float32, resolution uint16, padding int32, glyphs string) *Font {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sizeToC := C.float(size)
	resolutionToC := C.ushort(resolution)
	paddingToC := C.int32_t(padding)
	glyphsToC, idFinglyphsToC := wrapString(glyphs)
	defer idFinglyphsToC()
	retval := C.WrapLoadFontFromFileWithSizeResolutionPaddingGlyphs(pathToC, sizeToC, resolutionToC, paddingToC, glyphsToC)
	retvalGO := &Font{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Font) {
		C.WrapFontFree(cleanval.h)
	})
	return retvalGO
}

// LoadFontFromAssets Load a TrueType (TTF) font from the assets system.  See [harfang.man.Assets].
func LoadFontFromAssets(name string) *Font {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapLoadFontFromAssets(nameToC)
	retvalGO := &Font{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Font) {
		C.WrapFontFree(cleanval.h)
	})
	return retvalGO
}

// LoadFontFromAssetsWithSize Load a TrueType (TTF) font from the assets system.  See [harfang.man.Assets].
func LoadFontFromAssetsWithSize(name string, size float32) *Font {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sizeToC := C.float(size)
	retval := C.WrapLoadFontFromAssetsWithSize(nameToC, sizeToC)
	retvalGO := &Font{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Font) {
		C.WrapFontFree(cleanval.h)
	})
	return retvalGO
}

// LoadFontFromAssetsWithSizeResolution Load a TrueType (TTF) font from the assets system.  See [harfang.man.Assets].
func LoadFontFromAssetsWithSizeResolution(name string, size float32, resolution uint16) *Font {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sizeToC := C.float(size)
	resolutionToC := C.ushort(resolution)
	retval := C.WrapLoadFontFromAssetsWithSizeResolution(nameToC, sizeToC, resolutionToC)
	retvalGO := &Font{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Font) {
		C.WrapFontFree(cleanval.h)
	})
	return retvalGO
}

// LoadFontFromAssetsWithSizeResolutionPadding Load a TrueType (TTF) font from the assets system.  See [harfang.man.Assets].
func LoadFontFromAssetsWithSizeResolutionPadding(name string, size float32, resolution uint16, padding int32) *Font {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sizeToC := C.float(size)
	resolutionToC := C.ushort(resolution)
	paddingToC := C.int32_t(padding)
	retval := C.WrapLoadFontFromAssetsWithSizeResolutionPadding(nameToC, sizeToC, resolutionToC, paddingToC)
	retvalGO := &Font{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Font) {
		C.WrapFontFree(cleanval.h)
	})
	return retvalGO
}

// LoadFontFromAssetsWithSizeResolutionPaddingGlyphs Load a TrueType (TTF) font from the assets system.  See [harfang.man.Assets].
func LoadFontFromAssetsWithSizeResolutionPaddingGlyphs(name string, size float32, resolution uint16, padding int32, glyphs string) *Font {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sizeToC := C.float(size)
	resolutionToC := C.ushort(resolution)
	paddingToC := C.int32_t(padding)
	glyphsToC, idFinglyphsToC := wrapString(glyphs)
	defer idFinglyphsToC()
	retval := C.WrapLoadFontFromAssetsWithSizeResolutionPaddingGlyphs(nameToC, sizeToC, resolutionToC, paddingToC, glyphsToC)
	retvalGO := &Font{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Font) {
		C.WrapFontFree(cleanval.h)
	})
	return retvalGO
}

// DrawText Write text to the specified view using the provided shader program and uniform values.
func DrawText(viewid uint16, font *Font, text string, prg *ProgramHandle, pageuniform string, pagestage uint8, mtx *Mat4) {
	viewidToC := C.ushort(viewid)
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	prgToC := prg.h
	pageuniformToC, idFinpageuniformToC := wrapString(pageuniform)
	defer idFinpageuniformToC()
	pagestageToC := C.uchar(pagestage)
	mtxToC := mtx.h
	C.WrapDrawText(viewidToC, fontToC, textToC, prgToC, pageuniformToC, pagestageToC, mtxToC)
}

// DrawTextWithPos Write text to the specified view using the provided shader program and uniform values.
func DrawTextWithPos(viewid uint16, font *Font, text string, prg *ProgramHandle, pageuniform string, pagestage uint8, mtx *Mat4, pos *Vec3) {
	viewidToC := C.ushort(viewid)
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	prgToC := prg.h
	pageuniformToC, idFinpageuniformToC := wrapString(pageuniform)
	defer idFinpageuniformToC()
	pagestageToC := C.uchar(pagestage)
	mtxToC := mtx.h
	posToC := pos.h
	C.WrapDrawTextWithPos(viewidToC, fontToC, textToC, prgToC, pageuniformToC, pagestageToC, mtxToC, posToC)
}

// DrawTextWithPosHalignValign Write text to the specified view using the provided shader program and uniform values.
func DrawTextWithPosHalignValign(viewid uint16, font *Font, text string, prg *ProgramHandle, pageuniform string, pagestage uint8, mtx *Mat4, pos *Vec3, halign DrawTextHAlign, valign DrawTextVAlign) {
	viewidToC := C.ushort(viewid)
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	prgToC := prg.h
	pageuniformToC, idFinpageuniformToC := wrapString(pageuniform)
	defer idFinpageuniformToC()
	pagestageToC := C.uchar(pagestage)
	mtxToC := mtx.h
	posToC := pos.h
	halignToC := C.int32_t(halign)
	valignToC := C.int32_t(valign)
	C.WrapDrawTextWithPosHalignValign(viewidToC, fontToC, textToC, prgToC, pageuniformToC, pagestageToC, mtxToC, posToC, halignToC, valignToC)
}

// DrawTextWithPosHalignValignValuesTextures Write text to the specified view using the provided shader program and uniform values.
func DrawTextWithPosHalignValignValuesTextures(viewid uint16, font *Font, text string, prg *ProgramHandle, pageuniform string, pagestage uint8, mtx *Mat4, pos *Vec3, halign DrawTextHAlign, valign DrawTextVAlign, values *UniformSetValueList, textures *UniformSetTextureList) {
	viewidToC := C.ushort(viewid)
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	prgToC := prg.h
	pageuniformToC, idFinpageuniformToC := wrapString(pageuniform)
	defer idFinpageuniformToC()
	pagestageToC := C.uchar(pagestage)
	mtxToC := mtx.h
	posToC := pos.h
	halignToC := C.int32_t(halign)
	valignToC := C.int32_t(valign)
	valuesToC := values.h
	texturesToC := textures.h
	C.WrapDrawTextWithPosHalignValignValuesTextures(viewidToC, fontToC, textToC, prgToC, pageuniformToC, pagestageToC, mtxToC, posToC, halignToC, valignToC, valuesToC, texturesToC)
}

// DrawTextWithPosHalignValignValuesTexturesState Write text to the specified view using the provided shader program and uniform values.
func DrawTextWithPosHalignValignValuesTexturesState(viewid uint16, font *Font, text string, prg *ProgramHandle, pageuniform string, pagestage uint8, mtx *Mat4, pos *Vec3, halign DrawTextHAlign, valign DrawTextVAlign, values *UniformSetValueList, textures *UniformSetTextureList, state *RenderState) {
	viewidToC := C.ushort(viewid)
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	prgToC := prg.h
	pageuniformToC, idFinpageuniformToC := wrapString(pageuniform)
	defer idFinpageuniformToC()
	pagestageToC := C.uchar(pagestage)
	mtxToC := mtx.h
	posToC := pos.h
	halignToC := C.int32_t(halign)
	valignToC := C.int32_t(valign)
	valuesToC := values.h
	texturesToC := textures.h
	stateToC := state.h
	C.WrapDrawTextWithPosHalignValignValuesTexturesState(viewidToC, fontToC, textToC, prgToC, pageuniformToC, pagestageToC, mtxToC, posToC, halignToC, valignToC, valuesToC, texturesToC, stateToC)
}

// DrawTextWithPosHalignValignValuesTexturesStateDepth Write text to the specified view using the provided shader program and uniform values.
func DrawTextWithPosHalignValignValuesTexturesStateDepth(viewid uint16, font *Font, text string, prg *ProgramHandle, pageuniform string, pagestage uint8, mtx *Mat4, pos *Vec3, halign DrawTextHAlign, valign DrawTextVAlign, values *UniformSetValueList, textures *UniformSetTextureList, state *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	prgToC := prg.h
	pageuniformToC, idFinpageuniformToC := wrapString(pageuniform)
	defer idFinpageuniformToC()
	pagestageToC := C.uchar(pagestage)
	mtxToC := mtx.h
	posToC := pos.h
	halignToC := C.int32_t(halign)
	valignToC := C.int32_t(valign)
	valuesToC := values.h
	texturesToC := textures.h
	stateToC := state.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawTextWithPosHalignValignValuesTexturesStateDepth(viewidToC, fontToC, textToC, prgToC, pageuniformToC, pagestageToC, mtxToC, posToC, halignToC, valignToC, valuesToC, texturesToC, stateToC, depthToC)
}

// DrawTextWithPosHalignValignSliceOfValuesSliceOfTextures Write text to the specified view using the provided shader program and uniform values.
func DrawTextWithPosHalignValignSliceOfValuesSliceOfTextures(viewid uint16, font *Font, text string, prg *ProgramHandle, pageuniform string, pagestage uint8, mtx *Mat4, pos *Vec3, halign DrawTextHAlign, valign DrawTextVAlign, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture) {
	viewidToC := C.ushort(viewid)
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	prgToC := prg.h
	pageuniformToC, idFinpageuniformToC := wrapString(pageuniform)
	defer idFinpageuniformToC()
	pagestageToC := C.uchar(pagestage)
	mtxToC := mtx.h
	posToC := pos.h
	halignToC := C.int32_t(halign)
	valignToC := C.int32_t(valign)
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	C.WrapDrawTextWithPosHalignValignSliceOfValuesSliceOfTextures(viewidToC, fontToC, textToC, prgToC, pageuniformToC, pagestageToC, mtxToC, posToC, halignToC, valignToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf)
}

// DrawTextWithPosHalignValignSliceOfValuesSliceOfTexturesState Write text to the specified view using the provided shader program and uniform values.
func DrawTextWithPosHalignValignSliceOfValuesSliceOfTexturesState(viewid uint16, font *Font, text string, prg *ProgramHandle, pageuniform string, pagestage uint8, mtx *Mat4, pos *Vec3, halign DrawTextHAlign, valign DrawTextVAlign, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, state *RenderState) {
	viewidToC := C.ushort(viewid)
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	prgToC := prg.h
	pageuniformToC, idFinpageuniformToC := wrapString(pageuniform)
	defer idFinpageuniformToC()
	pagestageToC := C.uchar(pagestage)
	mtxToC := mtx.h
	posToC := pos.h
	halignToC := C.int32_t(halign)
	valignToC := C.int32_t(valign)
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	stateToC := state.h
	C.WrapDrawTextWithPosHalignValignSliceOfValuesSliceOfTexturesState(viewidToC, fontToC, textToC, prgToC, pageuniformToC, pagestageToC, mtxToC, posToC, halignToC, valignToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, stateToC)
}

// DrawTextWithPosHalignValignSliceOfValuesSliceOfTexturesStateDepth Write text to the specified view using the provided shader program and uniform values.
func DrawTextWithPosHalignValignSliceOfValuesSliceOfTexturesStateDepth(viewid uint16, font *Font, text string, prg *ProgramHandle, pageuniform string, pagestage uint8, mtx *Mat4, pos *Vec3, halign DrawTextHAlign, valign DrawTextVAlign, SliceOfvalues GoSliceOfUniformSetValue, SliceOftextures GoSliceOfUniformSetTexture, state *RenderState, depth uint32) {
	viewidToC := C.ushort(viewid)
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	prgToC := prg.h
	pageuniformToC, idFinpageuniformToC := wrapString(pageuniform)
	defer idFinpageuniformToC()
	pagestageToC := C.uchar(pagestage)
	mtxToC := mtx.h
	posToC := pos.h
	halignToC := C.int32_t(halign)
	valignToC := C.int32_t(valign)
	var SliceOfvaluesPointer []C.WrapUniformSetValue
	for _, s := range SliceOfvalues {
		SliceOfvaluesPointer = append(SliceOfvaluesPointer, s.h)
	}
	SliceOfvaluesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfvaluesPointer))
	SliceOfvaluesPointerToCSize := C.size_t(SliceOfvaluesPointerToC.Len)
	SliceOfvaluesPointerToCBuf := (*C.WrapUniformSetValue)(unsafe.Pointer(SliceOfvaluesPointerToC.Data))
	var SliceOftexturesPointer []C.WrapUniformSetTexture
	for _, s := range SliceOftextures {
		SliceOftexturesPointer = append(SliceOftexturesPointer, s.h)
	}
	SliceOftexturesPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOftexturesPointer))
	SliceOftexturesPointerToCSize := C.size_t(SliceOftexturesPointerToC.Len)
	SliceOftexturesPointerToCBuf := (*C.WrapUniformSetTexture)(unsafe.Pointer(SliceOftexturesPointerToC.Data))
	stateToC := state.h
	depthToC := C.uint32_t(depth)
	C.WrapDrawTextWithPosHalignValignSliceOfValuesSliceOfTexturesStateDepth(viewidToC, fontToC, textToC, prgToC, pageuniformToC, pagestageToC, mtxToC, posToC, halignToC, valignToC, SliceOfvaluesPointerToCSize, SliceOfvaluesPointerToCBuf, SliceOftexturesPointerToCSize, SliceOftexturesPointerToCBuf, stateToC, depthToC)
}

// ComputeTextRect Compute the width and height of a text string.
func ComputeTextRect(font *Font, text string) *Rect {
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	retval := C.WrapComputeTextRect(fontToC, textToC)
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// ComputeTextRectWithXpos Compute the width and height of a text string.
func ComputeTextRectWithXpos(font *Font, text string, xpos float32) *Rect {
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	xposToC := C.float(xpos)
	retval := C.WrapComputeTextRectWithXpos(fontToC, textToC, xposToC)
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// ComputeTextRectWithXposYpos Compute the width and height of a text string.
func ComputeTextRectWithXposYpos(font *Font, text string, xpos float32, ypos float32) *Rect {
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	xposToC := C.float(xpos)
	yposToC := C.float(ypos)
	retval := C.WrapComputeTextRectWithXposYpos(fontToC, textToC, xposToC, yposToC)
	retvalGO := &Rect{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Rect) {
		C.WrapRectFree(cleanval.h)
	})
	return retvalGO
}

// ComputeTextHeight Compute the height of a text string.
func ComputeTextHeight(font *Font, text string) float32 {
	fontToC := font.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	retval := C.WrapComputeTextHeight(fontToC, textToC)
	return float32(retval)
}

// LoadJsonFromFile Load a JSON from the local filesystem.
func LoadJsonFromFile(path string) *JSON {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadJsonFromFile(pathToC)
	retvalGO := &JSON{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *JSON) {
		C.WrapJSONFree(cleanval.h)
	})
	return retvalGO
}

// LoadJsonFromAssets Load a JSON from the assets system.  See [harfang.man.Assets].
func LoadJsonFromAssets(name string) *JSON {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapLoadJsonFromAssets(nameToC)
	retvalGO := &JSON{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *JSON) {
		C.WrapJSONFree(cleanval.h)
	})
	return retvalGO
}

// SaveJsonToFile Save a JSON object to the local filesystem.
func SaveJsonToFile(js *JSON, path string) bool {
	jsToC := js.h
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapSaveJsonToFile(jsToC, pathToC)
	return bool(retval)
}

// GetJsonString Return the value of a string JSON key.
func GetJsonString(js *JSON, key string) (bool, *string) {
	jsToC := js.h
	keyToC, idFinkeyToC := wrapString(key)
	defer idFinkeyToC()
	value := new(string)
	valueToC1 := C.CString(*value)
	valueToC := &valueToC1
	retval := C.WrapGetJsonString(jsToC, keyToC, valueToC)
	valueToCGO := string(C.GoString(*valueToC))
	return bool(retval), &valueToCGO
}

// GetJsonBool Return the value of a boolean JSON key.
func GetJsonBool(js *JSON, key string) (bool, *bool) {
	jsToC := js.h
	keyToC, idFinkeyToC := wrapString(key)
	defer idFinkeyToC()
	value := new(bool)
	valueToC := (*C.bool)(unsafe.Pointer(value))
	retval := C.WrapGetJsonBool(jsToC, keyToC, valueToC)
	return bool(retval), (*bool)(unsafe.Pointer(valueToC))
}

// GetJsonInt Return the value of an integer JSON key.
func GetJsonInt(js *JSON, key string) (bool, *int32) {
	jsToC := js.h
	keyToC, idFinkeyToC := wrapString(key)
	defer idFinkeyToC()
	value := new(int32)
	valueToC := (*C.int32_t)(unsafe.Pointer(value))
	retval := C.WrapGetJsonInt(jsToC, keyToC, valueToC)
	return bool(retval), (*int32)(unsafe.Pointer(valueToC))
}

// GetJsonFloat Return the value of a float JSON key.
func GetJsonFloat(js *JSON, key string) (bool, *float32) {
	jsToC := js.h
	keyToC, idFinkeyToC := wrapString(key)
	defer idFinkeyToC()
	value := new(float32)
	valueToC := (*C.float)(unsafe.Pointer(value))
	retval := C.WrapGetJsonFloat(jsToC, keyToC, valueToC)
	return bool(retval), (*float32)(unsafe.Pointer(valueToC))
}

// SetJsonValue Set a JSON key value.
func SetJsonValue(js *JSON, key string, value string) {
	jsToC := js.h
	keyToC, idFinkeyToC := wrapString(key)
	defer idFinkeyToC()
	valueToC, idFinvalueToC := wrapString(value)
	defer idFinvalueToC()
	C.WrapSetJsonValue(jsToC, keyToC, valueToC)
}

// SetJsonValueWithValue Set a JSON key value.
func SetJsonValueWithValue(js *JSON, key string, value bool) {
	jsToC := js.h
	keyToC, idFinkeyToC := wrapString(key)
	defer idFinkeyToC()
	valueToC := C.bool(value)
	C.WrapSetJsonValueWithValue(jsToC, keyToC, valueToC)
}

// SetJsonValueWithIntValue Set a JSON key value.
func SetJsonValueWithIntValue(js *JSON, key string, value int32) {
	jsToC := js.h
	keyToC, idFinkeyToC := wrapString(key)
	defer idFinkeyToC()
	valueToC := C.int32_t(value)
	C.WrapSetJsonValueWithIntValue(jsToC, keyToC, valueToC)
}

// SetJsonValueWithFloatValue Set a JSON key value.
func SetJsonValueWithFloatValue(js *JSON, key string, value float32) {
	jsToC := js.h
	keyToC, idFinkeyToC := wrapString(key)
	defer idFinkeyToC()
	valueToC := C.float(value)
	C.WrapSetJsonValueWithFloatValue(jsToC, keyToC, valueToC)
}

// CreateSceneRootNode Helper function to create a [harfang.Node] with a [harfang.Transform] component then parent all root nodes in the scene to it.
func CreateSceneRootNode(scene *Scene, name string, mtx *Mat4) *Node {
	sceneToC := scene.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	mtxToC := mtx.h
	retval := C.WrapCreateSceneRootNode(sceneToC, nameToC, mtxToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateCamera Create a new [harfang.Node] with a [harfang.Transform] and [harfang.Camera] components.
func CreateCamera(scene *Scene, mtx *Mat4, znear float32, zfar float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	retval := C.WrapCreateCamera(sceneToC, mtxToC, znearToC, zfarToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateCameraWithFov Create a new [harfang.Node] with a [harfang.Transform] and [harfang.Camera] components.
func CreateCameraWithFov(scene *Scene, mtx *Mat4, znear float32, zfar float32, fov float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	fovToC := C.float(fov)
	retval := C.WrapCreateCameraWithFov(sceneToC, mtxToC, znearToC, zfarToC, fovToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateOrthographicCamera Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Camera] component.
func CreateOrthographicCamera(scene *Scene, mtx *Mat4, znear float32, zfar float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	retval := C.WrapCreateOrthographicCamera(sceneToC, mtxToC, znearToC, zfarToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateOrthographicCameraWithSize Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Camera] component.
func CreateOrthographicCameraWithSize(scene *Scene, mtx *Mat4, znear float32, zfar float32, size float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	sizeToC := C.float(size)
	retval := C.WrapCreateOrthographicCameraWithSize(sceneToC, mtxToC, znearToC, zfarToC, sizeToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLight Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreatePointLight(scene *Scene, mtx *Mat4, radius float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	retval := C.WrapCreatePointLight(sceneToC, mtxToC, radiusToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuse Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreatePointLightWithDiffuse(scene *Scene, mtx *Mat4, radius float32, diffuse *Color) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	retval := C.WrapCreatePointLightWithDiffuse(sceneToC, mtxToC, radiusToC, diffuseToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseSpecular Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreatePointLightWithDiffuseSpecular(scene *Scene, mtx *Mat4, radius float32, diffuse *Color, specular *Color) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	specularToC := specular.h
	retval := C.WrapCreatePointLightWithDiffuseSpecular(sceneToC, mtxToC, radiusToC, diffuseToC, specularToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseSpecularPriority Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreatePointLightWithDiffuseSpecularPriority(scene *Scene, mtx *Mat4, radius float32, diffuse *Color, specular *Color, priority float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	retval := C.WrapCreatePointLightWithDiffuseSpecularPriority(sceneToC, mtxToC, radiusToC, diffuseToC, specularToC, priorityToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseSpecularPriorityShadowType Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreatePointLightWithDiffuseSpecularPriorityShadowType(scene *Scene, mtx *Mat4, radius float32, diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapCreatePointLightWithDiffuseSpecularPriorityShadowType(sceneToC, mtxToC, radiusToC, diffuseToC, specularToC, priorityToC, shadowtypeToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseSpecularPriorityShadowTypeShadowBias Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreatePointLightWithDiffuseSpecularPriorityShadowTypeShadowBias(scene *Scene, mtx *Mat4, radius float32, diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType, shadowbias float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapCreatePointLightWithDiffuseSpecularPriorityShadowTypeShadowBias(sceneToC, mtxToC, radiusToC, diffuseToC, specularToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseDiffuseIntensity Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreatePointLightWithDiffuseDiffuseIntensity(scene *Scene, mtx *Mat4, radius float32, diffuse *Color, diffuseintensity float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	retval := C.WrapCreatePointLightWithDiffuseDiffuseIntensity(sceneToC, mtxToC, radiusToC, diffuseToC, diffuseintensityToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseDiffuseIntensitySpecular Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreatePointLightWithDiffuseDiffuseIntensitySpecular(scene *Scene, mtx *Mat4, radius float32, diffuse *Color, diffuseintensity float32, specular *Color) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	retval := C.WrapCreatePointLightWithDiffuseDiffuseIntensitySpecular(sceneToC, mtxToC, radiusToC, diffuseToC, diffuseintensityToC, specularToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensity Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensity(scene *Scene, mtx *Mat4, radius float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	retval := C.WrapCreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensity(sceneToC, mtxToC, radiusToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriority Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriority(scene *Scene, mtx *Mat4, radius float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	retval := C.WrapCreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriority(sceneToC, mtxToC, radiusToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowType Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(scene *Scene, mtx *Mat4, radius float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapCreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(sceneToC, mtxToC, radiusToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(scene *Scene, mtx *Mat4, radius float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType, shadowbias float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapCreatePointLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(sceneToC, mtxToC, radiusToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLight Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateSpotLight(scene *Scene, mtx *Mat4, radius float32, innerangle float32, outerangle float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	retval := C.WrapCreateSpotLight(sceneToC, mtxToC, radiusToC, innerangleToC, outerangleToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuse Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateSpotLightWithDiffuse(scene *Scene, mtx *Mat4, radius float32, innerangle float32, outerangle float32, diffuse *Color) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	retval := C.WrapCreateSpotLightWithDiffuse(sceneToC, mtxToC, radiusToC, innerangleToC, outerangleToC, diffuseToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseSpecular Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateSpotLightWithDiffuseSpecular(scene *Scene, mtx *Mat4, radius float32, innerangle float32, outerangle float32, diffuse *Color, specular *Color) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	specularToC := specular.h
	retval := C.WrapCreateSpotLightWithDiffuseSpecular(sceneToC, mtxToC, radiusToC, innerangleToC, outerangleToC, diffuseToC, specularToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseSpecularPriority Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateSpotLightWithDiffuseSpecularPriority(scene *Scene, mtx *Mat4, radius float32, innerangle float32, outerangle float32, diffuse *Color, specular *Color, priority float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	retval := C.WrapCreateSpotLightWithDiffuseSpecularPriority(sceneToC, mtxToC, radiusToC, innerangleToC, outerangleToC, diffuseToC, specularToC, priorityToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseSpecularPriorityShadowType Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateSpotLightWithDiffuseSpecularPriorityShadowType(scene *Scene, mtx *Mat4, radius float32, innerangle float32, outerangle float32, diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapCreateSpotLightWithDiffuseSpecularPriorityShadowType(sceneToC, mtxToC, radiusToC, innerangleToC, outerangleToC, diffuseToC, specularToC, priorityToC, shadowtypeToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseSpecularPriorityShadowTypeShadowBias Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateSpotLightWithDiffuseSpecularPriorityShadowTypeShadowBias(scene *Scene, mtx *Mat4, radius float32, innerangle float32, outerangle float32, diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType, shadowbias float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapCreateSpotLightWithDiffuseSpecularPriorityShadowTypeShadowBias(sceneToC, mtxToC, radiusToC, innerangleToC, outerangleToC, diffuseToC, specularToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseDiffuseIntensity Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateSpotLightWithDiffuseDiffuseIntensity(scene *Scene, mtx *Mat4, radius float32, innerangle float32, outerangle float32, diffuse *Color, diffuseintensity float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	retval := C.WrapCreateSpotLightWithDiffuseDiffuseIntensity(sceneToC, mtxToC, radiusToC, innerangleToC, outerangleToC, diffuseToC, diffuseintensityToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseDiffuseIntensitySpecular Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateSpotLightWithDiffuseDiffuseIntensitySpecular(scene *Scene, mtx *Mat4, radius float32, innerangle float32, outerangle float32, diffuse *Color, diffuseintensity float32, specular *Color) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	retval := C.WrapCreateSpotLightWithDiffuseDiffuseIntensitySpecular(sceneToC, mtxToC, radiusToC, innerangleToC, outerangleToC, diffuseToC, diffuseintensityToC, specularToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensity Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensity(scene *Scene, mtx *Mat4, radius float32, innerangle float32, outerangle float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	retval := C.WrapCreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensity(sceneToC, mtxToC, radiusToC, innerangleToC, outerangleToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriority Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriority(scene *Scene, mtx *Mat4, radius float32, innerangle float32, outerangle float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	retval := C.WrapCreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriority(sceneToC, mtxToC, radiusToC, innerangleToC, outerangleToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowType Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(scene *Scene, mtx *Mat4, radius float32, innerangle float32, outerangle float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapCreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(sceneToC, mtxToC, radiusToC, innerangleToC, outerangleToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias Create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(scene *Scene, mtx *Mat4, radius float32, innerangle float32, outerangle float32, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType, shadowbias float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	radiusToC := C.float(radius)
	innerangleToC := C.float(innerangle)
	outerangleToC := C.float(outerangle)
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	retval := C.WrapCreateSpotLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBias(sceneToC, mtxToC, radiusToC, innerangleToC, outerangleToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC, shadowbiasToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLight Helper function to create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateLinearLight(scene *Scene, mtx *Mat4) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	retval := C.WrapCreateLinearLight(sceneToC, mtxToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuse Helper function to create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateLinearLightWithDiffuse(scene *Scene, mtx *Mat4, diffuse *Color) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	diffuseToC := diffuse.h
	retval := C.WrapCreateLinearLightWithDiffuse(sceneToC, mtxToC, diffuseToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseSpecular Helper function to create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateLinearLightWithDiffuseSpecular(scene *Scene, mtx *Mat4, diffuse *Color, specular *Color) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	retval := C.WrapCreateLinearLightWithDiffuseSpecular(sceneToC, mtxToC, diffuseToC, specularToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseSpecularPriority Helper function to create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateLinearLightWithDiffuseSpecularPriority(scene *Scene, mtx *Mat4, diffuse *Color, specular *Color, priority float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	retval := C.WrapCreateLinearLightWithDiffuseSpecularPriority(sceneToC, mtxToC, diffuseToC, specularToC, priorityToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseSpecularPriorityShadowType Helper function to create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateLinearLightWithDiffuseSpecularPriorityShadowType(scene *Scene, mtx *Mat4, diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapCreateLinearLightWithDiffuseSpecularPriorityShadowType(sceneToC, mtxToC, diffuseToC, specularToC, priorityToC, shadowtypeToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseSpecularPriorityShadowTypeShadowBiasPssmSplit Helper function to create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateLinearLightWithDiffuseSpecularPriorityShadowTypeShadowBiasPssmSplit(scene *Scene, mtx *Mat4, diffuse *Color, specular *Color, priority float32, shadowtype LightShadowType, shadowbias float32, pssmsplit *Vec4) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	diffuseToC := diffuse.h
	specularToC := specular.h
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	pssmsplitToC := pssmsplit.h
	retval := C.WrapCreateLinearLightWithDiffuseSpecularPriorityShadowTypeShadowBiasPssmSplit(sceneToC, mtxToC, diffuseToC, specularToC, priorityToC, shadowtypeToC, shadowbiasToC, pssmsplitToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseDiffuseIntensity Helper function to create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateLinearLightWithDiffuseDiffuseIntensity(scene *Scene, mtx *Mat4, diffuse *Color, diffuseintensity float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	retval := C.WrapCreateLinearLightWithDiffuseDiffuseIntensity(sceneToC, mtxToC, diffuseToC, diffuseintensityToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseDiffuseIntensitySpecular Helper function to create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateLinearLightWithDiffuseDiffuseIntensitySpecular(scene *Scene, mtx *Mat4, diffuse *Color, diffuseintensity float32, specular *Color) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	retval := C.WrapCreateLinearLightWithDiffuseDiffuseIntensitySpecular(sceneToC, mtxToC, diffuseToC, diffuseintensityToC, specularToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensity Helper function to create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensity(scene *Scene, mtx *Mat4, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	retval := C.WrapCreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensity(sceneToC, mtxToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriority Helper function to create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriority(scene *Scene, mtx *Mat4, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	retval := C.WrapCreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriority(sceneToC, mtxToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowType Helper function to create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(scene *Scene, mtx *Mat4, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	retval := C.WrapCreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowType(sceneToC, mtxToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBiasPssmSplit Helper function to create a [harfang.Node] with a [harfang.Transform] and a [harfang.Light] component.
func CreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBiasPssmSplit(scene *Scene, mtx *Mat4, diffuse *Color, diffuseintensity float32, specular *Color, specularintensity float32, priority float32, shadowtype LightShadowType, shadowbias float32, pssmsplit *Vec4) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	diffuseToC := diffuse.h
	diffuseintensityToC := C.float(diffuseintensity)
	specularToC := specular.h
	specularintensityToC := C.float(specularintensity)
	priorityToC := C.float(priority)
	shadowtypeToC := C.int32_t(shadowtype)
	shadowbiasToC := C.float(shadowbias)
	pssmsplitToC := pssmsplit.h
	retval := C.WrapCreateLinearLightWithDiffuseDiffuseIntensitySpecularSpecularIntensityPriorityShadowTypeShadowBiasPssmSplit(sceneToC, mtxToC, diffuseToC, diffuseintensityToC, specularToC, specularintensityToC, priorityToC, shadowtypeToC, shadowbiasToC, pssmsplitToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateObject Create a [harfang.Node] with a [harfang.Transform] and [harfang.Object] components.
func CreateObject(scene *Scene, mtx *Mat4, model *ModelRef, materials *MaterialList) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	modelToC := model.h
	materialsToC := materials.h
	retval := C.WrapCreateObject(sceneToC, mtxToC, modelToC, materialsToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateObjectWithSliceOfMaterials Create a [harfang.Node] with a [harfang.Transform] and [harfang.Object] components.
func CreateObjectWithSliceOfMaterials(scene *Scene, mtx *Mat4, model *ModelRef, SliceOfmaterials GoSliceOfMaterial) *Node {
	sceneToC := scene.h
	mtxToC := mtx.h
	modelToC := model.h
	var SliceOfmaterialsPointer []C.WrapMaterial
	for _, s := range SliceOfmaterials {
		SliceOfmaterialsPointer = append(SliceOfmaterialsPointer, s.h)
	}
	SliceOfmaterialsPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfmaterialsPointer))
	SliceOfmaterialsPointerToCSize := C.size_t(SliceOfmaterialsPointerToC.Len)
	SliceOfmaterialsPointerToCBuf := (*C.WrapMaterial)(unsafe.Pointer(SliceOfmaterialsPointerToC.Data))
	retval := C.WrapCreateObjectWithSliceOfMaterials(sceneToC, mtxToC, modelToC, SliceOfmaterialsPointerToCSize, SliceOfmaterialsPointerToCBuf)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateInstanceFromFile Helper function to create a [harfang.Node] with a [harfang.Transform] and an [harfang.Instance] component.  The instance component will be setup and its resources loaded from the local filesystem.  See [harfang.man.Assets].
func CreateInstanceFromFile(scene *Scene, mtx *Mat4, name string, resources *PipelineResources, pipeline *PipelineInfo) (*Node, *bool) {
	sceneToC := scene.h
	mtxToC := mtx.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	success := new(bool)
	successToC := (*C.bool)(unsafe.Pointer(success))
	retval := C.WrapCreateInstanceFromFile(sceneToC, mtxToC, nameToC, resourcesToC, pipelineToC, successToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO, (*bool)(unsafe.Pointer(successToC))
}

// CreateInstanceFromFileWithFlags Helper function to create a [harfang.Node] with a [harfang.Transform] and an [harfang.Instance] component.  The instance component will be setup and its resources loaded from the local filesystem.  See [harfang.man.Assets].
func CreateInstanceFromFileWithFlags(scene *Scene, mtx *Mat4, name string, resources *PipelineResources, pipeline *PipelineInfo, flags LoadSaveSceneFlags) (*Node, *bool) {
	sceneToC := scene.h
	mtxToC := mtx.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	success := new(bool)
	successToC := (*C.bool)(unsafe.Pointer(success))
	flagsToC := C.uint32_t(flags)
	retval := C.WrapCreateInstanceFromFileWithFlags(sceneToC, mtxToC, nameToC, resourcesToC, pipelineToC, successToC, flagsToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO, (*bool)(unsafe.Pointer(successToC))
}

// CreateInstanceFromAssets Helper function to create a [harfang.Node] with a [harfang.Transform] and an [harfang.Instance] component.  The instance component will be setup and its resources loaded from the assets system.  See [harfang.man.Assets].
func CreateInstanceFromAssets(scene *Scene, mtx *Mat4, name string, resources *PipelineResources, pipeline *PipelineInfo) (*Node, *bool) {
	sceneToC := scene.h
	mtxToC := mtx.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	success := new(bool)
	successToC := (*C.bool)(unsafe.Pointer(success))
	retval := C.WrapCreateInstanceFromAssets(sceneToC, mtxToC, nameToC, resourcesToC, pipelineToC, successToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO, (*bool)(unsafe.Pointer(successToC))
}

// CreateInstanceFromAssetsWithFlags Helper function to create a [harfang.Node] with a [harfang.Transform] and an [harfang.Instance] component.  The instance component will be setup and its resources loaded from the assets system.  See [harfang.man.Assets].
func CreateInstanceFromAssetsWithFlags(scene *Scene, mtx *Mat4, name string, resources *PipelineResources, pipeline *PipelineInfo, flags LoadSaveSceneFlags) (*Node, *bool) {
	sceneToC := scene.h
	mtxToC := mtx.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	success := new(bool)
	successToC := (*C.bool)(unsafe.Pointer(success))
	flagsToC := C.uint32_t(flags)
	retval := C.WrapCreateInstanceFromAssetsWithFlags(sceneToC, mtxToC, nameToC, resourcesToC, pipelineToC, successToC, flagsToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO, (*bool)(unsafe.Pointer(successToC))
}

// CreateScript Helper function to create a [harfang.Node] with a [harfang.Script] component.
func CreateScript(scene *Scene) *Node {
	sceneToC := scene.h
	retval := C.WrapCreateScript(sceneToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreateScriptWithPath Helper function to create a [harfang.Node] with a [harfang.Script] component.
func CreateScriptWithPath(scene *Scene, path string) *Node {
	sceneToC := scene.h
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapCreateScriptWithPath(sceneToC, pathToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePhysicSphere Create a [harfang.Node] with a [harfang.Transform], [harfang.Object] and [harfang.RigidBody] components.
func CreatePhysicSphere(scene *Scene, radius float32, mtx *Mat4, modelref *ModelRef, materials *MaterialList) *Node {
	sceneToC := scene.h
	radiusToC := C.float(radius)
	mtxToC := mtx.h
	modelrefToC := modelref.h
	materialsToC := materials.h
	retval := C.WrapCreatePhysicSphere(sceneToC, radiusToC, mtxToC, modelrefToC, materialsToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePhysicSphereWithMass Create a [harfang.Node] with a [harfang.Transform], [harfang.Object] and [harfang.RigidBody] components.
func CreatePhysicSphereWithMass(scene *Scene, radius float32, mtx *Mat4, modelref *ModelRef, materials *MaterialList, mass float32) *Node {
	sceneToC := scene.h
	radiusToC := C.float(radius)
	mtxToC := mtx.h
	modelrefToC := modelref.h
	materialsToC := materials.h
	massToC := C.float(mass)
	retval := C.WrapCreatePhysicSphereWithMass(sceneToC, radiusToC, mtxToC, modelrefToC, materialsToC, massToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePhysicSphereWithSliceOfMaterials Create a [harfang.Node] with a [harfang.Transform], [harfang.Object] and [harfang.RigidBody] components.
func CreatePhysicSphereWithSliceOfMaterials(scene *Scene, radius float32, mtx *Mat4, modelref *ModelRef, SliceOfmaterials GoSliceOfMaterial) *Node {
	sceneToC := scene.h
	radiusToC := C.float(radius)
	mtxToC := mtx.h
	modelrefToC := modelref.h
	var SliceOfmaterialsPointer []C.WrapMaterial
	for _, s := range SliceOfmaterials {
		SliceOfmaterialsPointer = append(SliceOfmaterialsPointer, s.h)
	}
	SliceOfmaterialsPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfmaterialsPointer))
	SliceOfmaterialsPointerToCSize := C.size_t(SliceOfmaterialsPointerToC.Len)
	SliceOfmaterialsPointerToCBuf := (*C.WrapMaterial)(unsafe.Pointer(SliceOfmaterialsPointerToC.Data))
	retval := C.WrapCreatePhysicSphereWithSliceOfMaterials(sceneToC, radiusToC, mtxToC, modelrefToC, SliceOfmaterialsPointerToCSize, SliceOfmaterialsPointerToCBuf)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePhysicSphereWithSliceOfMaterialsMass Create a [harfang.Node] with a [harfang.Transform], [harfang.Object] and [harfang.RigidBody] components.
func CreatePhysicSphereWithSliceOfMaterialsMass(scene *Scene, radius float32, mtx *Mat4, modelref *ModelRef, SliceOfmaterials GoSliceOfMaterial, mass float32) *Node {
	sceneToC := scene.h
	radiusToC := C.float(radius)
	mtxToC := mtx.h
	modelrefToC := modelref.h
	var SliceOfmaterialsPointer []C.WrapMaterial
	for _, s := range SliceOfmaterials {
		SliceOfmaterialsPointer = append(SliceOfmaterialsPointer, s.h)
	}
	SliceOfmaterialsPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfmaterialsPointer))
	SliceOfmaterialsPointerToCSize := C.size_t(SliceOfmaterialsPointerToC.Len)
	SliceOfmaterialsPointerToCBuf := (*C.WrapMaterial)(unsafe.Pointer(SliceOfmaterialsPointerToC.Data))
	massToC := C.float(mass)
	retval := C.WrapCreatePhysicSphereWithSliceOfMaterialsMass(sceneToC, radiusToC, mtxToC, modelrefToC, SliceOfmaterialsPointerToCSize, SliceOfmaterialsPointerToCBuf, massToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePhysicCube Create a [harfang.Node] with a [harfang.Transform], [harfang.Object] and [harfang.RigidBody] components.
func CreatePhysicCube(scene *Scene, size *Vec3, mtx *Mat4, modelref *ModelRef, materials *MaterialList) *Node {
	sceneToC := scene.h
	sizeToC := size.h
	mtxToC := mtx.h
	modelrefToC := modelref.h
	materialsToC := materials.h
	retval := C.WrapCreatePhysicCube(sceneToC, sizeToC, mtxToC, modelrefToC, materialsToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePhysicCubeWithMass Create a [harfang.Node] with a [harfang.Transform], [harfang.Object] and [harfang.RigidBody] components.
func CreatePhysicCubeWithMass(scene *Scene, size *Vec3, mtx *Mat4, modelref *ModelRef, materials *MaterialList, mass float32) *Node {
	sceneToC := scene.h
	sizeToC := size.h
	mtxToC := mtx.h
	modelrefToC := modelref.h
	materialsToC := materials.h
	massToC := C.float(mass)
	retval := C.WrapCreatePhysicCubeWithMass(sceneToC, sizeToC, mtxToC, modelrefToC, materialsToC, massToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePhysicCubeWithSliceOfMaterials Create a [harfang.Node] with a [harfang.Transform], [harfang.Object] and [harfang.RigidBody] components.
func CreatePhysicCubeWithSliceOfMaterials(scene *Scene, size *Vec3, mtx *Mat4, modelref *ModelRef, SliceOfmaterials GoSliceOfMaterial) *Node {
	sceneToC := scene.h
	sizeToC := size.h
	mtxToC := mtx.h
	modelrefToC := modelref.h
	var SliceOfmaterialsPointer []C.WrapMaterial
	for _, s := range SliceOfmaterials {
		SliceOfmaterialsPointer = append(SliceOfmaterialsPointer, s.h)
	}
	SliceOfmaterialsPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfmaterialsPointer))
	SliceOfmaterialsPointerToCSize := C.size_t(SliceOfmaterialsPointerToC.Len)
	SliceOfmaterialsPointerToCBuf := (*C.WrapMaterial)(unsafe.Pointer(SliceOfmaterialsPointerToC.Data))
	retval := C.WrapCreatePhysicCubeWithSliceOfMaterials(sceneToC, sizeToC, mtxToC, modelrefToC, SliceOfmaterialsPointerToCSize, SliceOfmaterialsPointerToCBuf)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// CreatePhysicCubeWithSliceOfMaterialsMass Create a [harfang.Node] with a [harfang.Transform], [harfang.Object] and [harfang.RigidBody] components.
func CreatePhysicCubeWithSliceOfMaterialsMass(scene *Scene, size *Vec3, mtx *Mat4, modelref *ModelRef, SliceOfmaterials GoSliceOfMaterial, mass float32) *Node {
	sceneToC := scene.h
	sizeToC := size.h
	mtxToC := mtx.h
	modelrefToC := modelref.h
	var SliceOfmaterialsPointer []C.WrapMaterial
	for _, s := range SliceOfmaterials {
		SliceOfmaterialsPointer = append(SliceOfmaterialsPointer, s.h)
	}
	SliceOfmaterialsPointerToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfmaterialsPointer))
	SliceOfmaterialsPointerToCSize := C.size_t(SliceOfmaterialsPointerToC.Len)
	SliceOfmaterialsPointerToCBuf := (*C.WrapMaterial)(unsafe.Pointer(SliceOfmaterialsPointerToC.Data))
	massToC := C.float(mass)
	retval := C.WrapCreatePhysicCubeWithSliceOfMaterialsMass(sceneToC, sizeToC, mtxToC, modelrefToC, SliceOfmaterialsPointerToCSize, SliceOfmaterialsPointerToCBuf, massToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// SaveSceneJsonToFile ...
func SaveSceneJsonToFile(path string, scene *Scene, resources *PipelineResources) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	retval := C.WrapSaveSceneJsonToFile(pathToC, sceneToC, resourcesToC)
	return bool(retval)
}

// SaveSceneJsonToFileWithFlags ...
func SaveSceneJsonToFileWithFlags(path string, scene *Scene, resources *PipelineResources, flags LoadSaveSceneFlags) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapSaveSceneJsonToFileWithFlags(pathToC, sceneToC, resourcesToC, flagsToC)
	return bool(retval)
}

// SaveSceneBinaryToFile ...
func SaveSceneBinaryToFile(path string, scene *Scene, resources *PipelineResources) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	retval := C.WrapSaveSceneBinaryToFile(pathToC, sceneToC, resourcesToC)
	return bool(retval)
}

// SaveSceneBinaryToFileWithFlags ...
func SaveSceneBinaryToFileWithFlags(path string, scene *Scene, resources *PipelineResources, flags LoadSaveSceneFlags) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapSaveSceneBinaryToFileWithFlags(pathToC, sceneToC, resourcesToC, flagsToC)
	return bool(retval)
}

// SaveSceneBinaryToData ...
func SaveSceneBinaryToData(data *Data, scene *Scene, resources *PipelineResources) bool {
	dataToC := data.h
	sceneToC := scene.h
	resourcesToC := resources.h
	retval := C.WrapSaveSceneBinaryToData(dataToC, sceneToC, resourcesToC)
	return bool(retval)
}

// SaveSceneBinaryToDataWithFlags ...
func SaveSceneBinaryToDataWithFlags(data *Data, scene *Scene, resources *PipelineResources, flags LoadSaveSceneFlags) bool {
	dataToC := data.h
	sceneToC := scene.h
	resourcesToC := resources.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapSaveSceneBinaryToDataWithFlags(dataToC, sceneToC, resourcesToC, flagsToC)
	return bool(retval)
}

// LoadSceneBinaryFromFile Load a scene in binary format from the local filesystem. Loaded content is added to the existing scene content.
func LoadSceneBinaryFromFile(path string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapLoadSceneBinaryFromFile(pathToC, sceneToC, resourcesToC, pipelineToC)
	return bool(retval)
}

// LoadSceneBinaryFromFileWithFlags Load a scene in binary format from the local filesystem. Loaded content is added to the existing scene content.
func LoadSceneBinaryFromFileWithFlags(path string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo, flags LoadSaveSceneFlags) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapLoadSceneBinaryFromFileWithFlags(pathToC, sceneToC, resourcesToC, pipelineToC, flagsToC)
	return bool(retval)
}

// LoadSceneBinaryFromAssets Load a scene in binary format from the assets system. Loaded content is added to the existing scene content.  See [harfang.man.Assets].
func LoadSceneBinaryFromAssets(name string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo) bool {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapLoadSceneBinaryFromAssets(nameToC, sceneToC, resourcesToC, pipelineToC)
	return bool(retval)
}

// LoadSceneBinaryFromAssetsWithFlags Load a scene in binary format from the assets system. Loaded content is added to the existing scene content.  See [harfang.man.Assets].
func LoadSceneBinaryFromAssetsWithFlags(name string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo, flags LoadSaveSceneFlags) bool {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapLoadSceneBinaryFromAssetsWithFlags(nameToC, sceneToC, resourcesToC, pipelineToC, flagsToC)
	return bool(retval)
}

// LoadSceneJsonFromFile Load a scene in JSON format from the local filesystem. Loaded content is added to the existing scene content.
func LoadSceneJsonFromFile(path string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapLoadSceneJsonFromFile(pathToC, sceneToC, resourcesToC, pipelineToC)
	return bool(retval)
}

// LoadSceneJsonFromFileWithFlags Load a scene in JSON format from the local filesystem. Loaded content is added to the existing scene content.
func LoadSceneJsonFromFileWithFlags(path string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo, flags LoadSaveSceneFlags) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapLoadSceneJsonFromFileWithFlags(pathToC, sceneToC, resourcesToC, pipelineToC, flagsToC)
	return bool(retval)
}

// LoadSceneJsonFromAssets Load a scene in JSON format from the assets system. Loaded content is added to the existing scene content.  See [harfang.man.Assets].
func LoadSceneJsonFromAssets(name string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo) bool {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapLoadSceneJsonFromAssets(nameToC, sceneToC, resourcesToC, pipelineToC)
	return bool(retval)
}

// LoadSceneJsonFromAssetsWithFlags Load a scene in JSON format from the assets system. Loaded content is added to the existing scene content.  See [harfang.man.Assets].
func LoadSceneJsonFromAssetsWithFlags(name string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo, flags LoadSaveSceneFlags) bool {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapLoadSceneJsonFromAssetsWithFlags(nameToC, sceneToC, resourcesToC, pipelineToC, flagsToC)
	return bool(retval)
}

// LoadSceneBinaryFromDataAndFile ...
func LoadSceneBinaryFromDataAndFile(data *Data, name string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo) bool {
	dataToC := data.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapLoadSceneBinaryFromDataAndFile(dataToC, nameToC, sceneToC, resourcesToC, pipelineToC)
	return bool(retval)
}

// LoadSceneBinaryFromDataAndFileWithFlags ...
func LoadSceneBinaryFromDataAndFileWithFlags(data *Data, name string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo, flags LoadSaveSceneFlags) bool {
	dataToC := data.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapLoadSceneBinaryFromDataAndFileWithFlags(dataToC, nameToC, sceneToC, resourcesToC, pipelineToC, flagsToC)
	return bool(retval)
}

// LoadSceneBinaryFromDataAndAssets ...
func LoadSceneBinaryFromDataAndAssets(data *Data, name string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo) bool {
	dataToC := data.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapLoadSceneBinaryFromDataAndAssets(dataToC, nameToC, sceneToC, resourcesToC, pipelineToC)
	return bool(retval)
}

// LoadSceneBinaryFromDataAndAssetsWithFlags ...
func LoadSceneBinaryFromDataAndAssetsWithFlags(data *Data, name string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo, flags LoadSaveSceneFlags) bool {
	dataToC := data.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapLoadSceneBinaryFromDataAndAssetsWithFlags(dataToC, nameToC, sceneToC, resourcesToC, pipelineToC, flagsToC)
	return bool(retval)
}

// LoadSceneFromFile ...
func LoadSceneFromFile(path string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapLoadSceneFromFile(pathToC, sceneToC, resourcesToC, pipelineToC)
	return bool(retval)
}

// LoadSceneFromFileWithFlags ...
func LoadSceneFromFileWithFlags(path string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo, flags LoadSaveSceneFlags) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapLoadSceneFromFileWithFlags(pathToC, sceneToC, resourcesToC, pipelineToC, flagsToC)
	return bool(retval)
}

// LoadSceneFromAssets ...
func LoadSceneFromAssets(name string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo) bool {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapLoadSceneFromAssets(nameToC, sceneToC, resourcesToC, pipelineToC)
	return bool(retval)
}

// LoadSceneFromAssetsWithFlags ...
func LoadSceneFromAssetsWithFlags(name string, scene *Scene, resources *PipelineResources, pipeline *PipelineInfo, flags LoadSaveSceneFlags) bool {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sceneToC := scene.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	flagsToC := C.uint32_t(flags)
	retval := C.WrapLoadSceneFromAssetsWithFlags(nameToC, sceneToC, resourcesToC, pipelineToC, flagsToC)
	return bool(retval)
}

// DuplicateNodesFromFile Duplicate each node of a list. Resources will be loaded from the local filesystem.  See [harfang.man.Assets].
func DuplicateNodesFromFile(scene *Scene, nodes *NodeList, resources *PipelineResources, pipeline *PipelineInfo) *NodeList {
	sceneToC := scene.h
	nodesToC := nodes.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapDuplicateNodesFromFile(sceneToC, nodesToC, resourcesToC, pipelineToC)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// DuplicateNodesFromAssets Duplicate each node of a list. Resources will be loaded from the assets system.
func DuplicateNodesFromAssets(scene *Scene, nodes *NodeList, resources *PipelineResources, pipeline *PipelineInfo) *NodeList {
	sceneToC := scene.h
	nodesToC := nodes.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapDuplicateNodesFromAssets(sceneToC, nodesToC, resourcesToC, pipelineToC)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// DuplicateNodesAndChildrenFromFile Duplicate each node and children hierarchy of a list. Resources will be loaded from the local filesystem.  See [harfang.man.Assets].
func DuplicateNodesAndChildrenFromFile(scene *Scene, nodes *NodeList, resources *PipelineResources, pipeline *PipelineInfo) *NodeList {
	sceneToC := scene.h
	nodesToC := nodes.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapDuplicateNodesAndChildrenFromFile(sceneToC, nodesToC, resourcesToC, pipelineToC)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// DuplicateNodesAndChildrenFromAssets Duplicate each node and children hierarchy of a list. Resources will be loaded from the assets system.  See [harfang.man.Assets].
func DuplicateNodesAndChildrenFromAssets(scene *Scene, nodes *NodeList, resources *PipelineResources, pipeline *PipelineInfo) *NodeList {
	sceneToC := scene.h
	nodesToC := nodes.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapDuplicateNodesAndChildrenFromAssets(sceneToC, nodesToC, resourcesToC, pipelineToC)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// DuplicateNodeFromFile Duplicate a node. Resources will be loaded from the local filesystem.  See [harfang.man.Assets].
func DuplicateNodeFromFile(scene *Scene, node *Node, resources *PipelineResources, pipeline *PipelineInfo) *Node {
	sceneToC := scene.h
	nodeToC := node.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapDuplicateNodeFromFile(sceneToC, nodeToC, resourcesToC, pipelineToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// DuplicateNodeFromAssets Duplicate a node. Resources will be loaded from the assets system.  See [harfang.man.Assets].
func DuplicateNodeFromAssets(scene *Scene, node *Node, resources *PipelineResources, pipeline *PipelineInfo) *Node {
	sceneToC := scene.h
	nodeToC := node.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapDuplicateNodeFromAssets(sceneToC, nodeToC, resourcesToC, pipelineToC)
	retvalGO := &Node{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Node) {
		C.WrapNodeFree(cleanval.h)
	})
	return retvalGO
}

// DuplicateNodeAndChildrenFromFile Duplicate a node and its child hierarchy. Resources will be loaded from the local filesystem.  See [harfang.man.Assets].
func DuplicateNodeAndChildrenFromFile(scene *Scene, node *Node, resources *PipelineResources, pipeline *PipelineInfo) *NodeList {
	sceneToC := scene.h
	nodeToC := node.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapDuplicateNodeAndChildrenFromFile(sceneToC, nodeToC, resourcesToC, pipelineToC)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// DuplicateNodeAndChildrenFromAssets Duplicate a node and its child hierarchy. Resources will be loaded from the assets system.  See [harfang.man.Assets].
func DuplicateNodeAndChildrenFromAssets(scene *Scene, node *Node, resources *PipelineResources, pipeline *PipelineInfo) *NodeList {
	sceneToC := scene.h
	nodeToC := node.h
	resourcesToC := resources.h
	pipelineToC := pipeline.h
	retval := C.WrapDuplicateNodeAndChildrenFromAssets(sceneToC, nodeToC, resourcesToC, pipelineToC)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// GetSceneForwardPipelineFog ...
func GetSceneForwardPipelineFog(scene *Scene) *ForwardPipelineFog {
	sceneToC := scene.h
	retval := C.WrapGetSceneForwardPipelineFog(sceneToC)
	retvalGO := &ForwardPipelineFog{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineFog) {
		C.WrapForwardPipelineFogFree(cleanval.h)
	})
	return retvalGO
}

// GetSceneForwardPipelineLights Filter through the scene lights and return a list of pipeline lights to be used by the scene forward pipeline.
func GetSceneForwardPipelineLights(scene *Scene) *ForwardPipelineLightList {
	sceneToC := scene.h
	retval := C.WrapGetSceneForwardPipelineLights(sceneToC)
	retvalGO := &ForwardPipelineLightList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineLightList) {
		C.WrapForwardPipelineLightListFree(cleanval.h)
	})
	return retvalGO
}

// GetSceneForwardPipelinePassViewId Return the view id for a scene forward pipeline pass id.
func GetSceneForwardPipelinePassViewId(views *SceneForwardPipelinePassViewId, pass SceneForwardPipelinePass) uint16 {
	viewsToC := views.h
	passToC := C.int32_t(pass)
	retval := C.WrapGetSceneForwardPipelinePassViewId(viewsToC, passToC)
	return uint16(retval)
}

// PrepareSceneForwardPipelineCommonRenderData Prepare the common render data to submit a scene to the forward pipeline.  Note: When rendering multiple views of the same scene, common data only needs to be prepared once.  See [harfang.PrepareSceneForwardPipelineViewDependentRenderData].
func PrepareSceneForwardPipelineCommonRenderData(viewid *uint16, scene *Scene, renderdata *SceneForwardPipelineRenderData, pipeline *ForwardPipeline, resources *PipelineResources, views *SceneForwardPipelinePassViewId) {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	renderdataToC := renderdata.h
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	viewsToC := views.h
	C.WrapPrepareSceneForwardPipelineCommonRenderData(viewidToC, sceneToC, renderdataToC, pipelineToC, resourcesToC, viewsToC)
}

// PrepareSceneForwardPipelineCommonRenderDataWithDebugName Prepare the common render data to submit a scene to the forward pipeline.  Note: When rendering multiple views of the same scene, common data only needs to be prepared once.  See [harfang.PrepareSceneForwardPipelineViewDependentRenderData].
func PrepareSceneForwardPipelineCommonRenderDataWithDebugName(viewid *uint16, scene *Scene, renderdata *SceneForwardPipelineRenderData, pipeline *ForwardPipeline, resources *PipelineResources, views *SceneForwardPipelinePassViewId, debugname string) {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	renderdataToC := renderdata.h
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	viewsToC := views.h
	debugnameToC, idFindebugnameToC := wrapString(debugname)
	defer idFindebugnameToC()
	C.WrapPrepareSceneForwardPipelineCommonRenderDataWithDebugName(viewidToC, sceneToC, renderdataToC, pipelineToC, resourcesToC, viewsToC, debugnameToC)
}

// PrepareSceneForwardPipelineViewDependentRenderData Prepare the view dependent render data to submit a scene to the forward pipeline.  See [harfang.PrepareSceneForwardPipelineCommonRenderData].
func PrepareSceneForwardPipelineViewDependentRenderData(viewid *uint16, viewstate *ViewState, scene *Scene, renderdata *SceneForwardPipelineRenderData, pipeline *ForwardPipeline, resources *PipelineResources, views *SceneForwardPipelinePassViewId) {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	viewstateToC := viewstate.h
	sceneToC := scene.h
	renderdataToC := renderdata.h
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	viewsToC := views.h
	C.WrapPrepareSceneForwardPipelineViewDependentRenderData(viewidToC, viewstateToC, sceneToC, renderdataToC, pipelineToC, resourcesToC, viewsToC)
}

// PrepareSceneForwardPipelineViewDependentRenderDataWithDebugName Prepare the view dependent render data to submit a scene to the forward pipeline.  See [harfang.PrepareSceneForwardPipelineCommonRenderData].
func PrepareSceneForwardPipelineViewDependentRenderDataWithDebugName(viewid *uint16, viewstate *ViewState, scene *Scene, renderdata *SceneForwardPipelineRenderData, pipeline *ForwardPipeline, resources *PipelineResources, views *SceneForwardPipelinePassViewId, debugname string) {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	viewstateToC := viewstate.h
	sceneToC := scene.h
	renderdataToC := renderdata.h
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	viewsToC := views.h
	debugnameToC, idFindebugnameToC := wrapString(debugname)
	defer idFindebugnameToC()
	C.WrapPrepareSceneForwardPipelineViewDependentRenderDataWithDebugName(viewidToC, viewstateToC, sceneToC, renderdataToC, pipelineToC, resourcesToC, viewsToC, debugnameToC)
}

// SubmitSceneToForwardPipeline Submit a scene to a forward pipeline.  See [harfang.PrepareSceneForwardPipelineCommonRenderData] and [harfang.PrepareSceneForwardPipelineViewDependentRenderData] if you need to render the same scene from different points of view.
func SubmitSceneToForwardPipeline(viewid *uint16, scene *Scene, rect *IntRect, viewstate *ViewState, pipeline *ForwardPipeline, renderdata *SceneForwardPipelineRenderData, resources *PipelineResources) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	viewstateToC := viewstate.h
	pipelineToC := pipeline.h
	renderdataToC := renderdata.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	C.WrapSubmitSceneToForwardPipeline(viewidToC, sceneToC, rectToC, viewstateToC, pipelineToC, renderdataToC, resourcesToC, viewsToC)
	return views
}

// SubmitSceneToForwardPipelineWithFrameBuffer Submit a scene to a forward pipeline.  See [harfang.PrepareSceneForwardPipelineCommonRenderData] and [harfang.PrepareSceneForwardPipelineViewDependentRenderData] if you need to render the same scene from different points of view.
func SubmitSceneToForwardPipelineWithFrameBuffer(viewid *uint16, scene *Scene, rect *IntRect, viewstate *ViewState, pipeline *ForwardPipeline, renderdata *SceneForwardPipelineRenderData, resources *PipelineResources, framebuffer *FrameBufferHandle) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	viewstateToC := viewstate.h
	pipelineToC := pipeline.h
	renderdataToC := renderdata.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	framebufferToC := framebuffer.h
	C.WrapSubmitSceneToForwardPipelineWithFrameBuffer(viewidToC, sceneToC, rectToC, viewstateToC, pipelineToC, renderdataToC, resourcesToC, viewsToC, framebufferToC)
	return views
}

// SubmitSceneToForwardPipelineWithFrameBufferDebugName Submit a scene to a forward pipeline.  See [harfang.PrepareSceneForwardPipelineCommonRenderData] and [harfang.PrepareSceneForwardPipelineViewDependentRenderData] if you need to render the same scene from different points of view.
func SubmitSceneToForwardPipelineWithFrameBufferDebugName(viewid *uint16, scene *Scene, rect *IntRect, viewstate *ViewState, pipeline *ForwardPipeline, renderdata *SceneForwardPipelineRenderData, resources *PipelineResources, framebuffer *FrameBufferHandle, debugname string) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	viewstateToC := viewstate.h
	pipelineToC := pipeline.h
	renderdataToC := renderdata.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	framebufferToC := framebuffer.h
	debugnameToC, idFindebugnameToC := wrapString(debugname)
	defer idFindebugnameToC()
	C.WrapSubmitSceneToForwardPipelineWithFrameBufferDebugName(viewidToC, sceneToC, rectToC, viewstateToC, pipelineToC, renderdataToC, resourcesToC, viewsToC, framebufferToC, debugnameToC)
	return views
}

// SubmitSceneToForwardPipelineWithAaaAaaConfigFrame Submit a scene to a forward pipeline.  See [harfang.PrepareSceneForwardPipelineCommonRenderData] and [harfang.PrepareSceneForwardPipelineViewDependentRenderData] if you need to render the same scene from different points of view.
func SubmitSceneToForwardPipelineWithAaaAaaConfigFrame(viewid *uint16, scene *Scene, rect *IntRect, viewstate *ViewState, pipeline *ForwardPipeline, renderdata *SceneForwardPipelineRenderData, resources *PipelineResources, views *SceneForwardPipelinePassViewId, aaa *ForwardPipelineAAA, aaaconfig *ForwardPipelineAAAConfig, frame int32) {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	viewstateToC := viewstate.h
	pipelineToC := pipeline.h
	renderdataToC := renderdata.h
	resourcesToC := resources.h
	viewsToC := views.h
	aaaToC := aaa.h
	aaaconfigToC := aaaconfig.h
	frameToC := C.int32_t(frame)
	C.WrapSubmitSceneToForwardPipelineWithAaaAaaConfigFrame(viewidToC, sceneToC, rectToC, viewstateToC, pipelineToC, renderdataToC, resourcesToC, viewsToC, aaaToC, aaaconfigToC, frameToC)
}

// SubmitSceneToForwardPipelineWithAaaAaaConfigFrameFrameBuffer Submit a scene to a forward pipeline.  See [harfang.PrepareSceneForwardPipelineCommonRenderData] and [harfang.PrepareSceneForwardPipelineViewDependentRenderData] if you need to render the same scene from different points of view.
func SubmitSceneToForwardPipelineWithAaaAaaConfigFrameFrameBuffer(viewid *uint16, scene *Scene, rect *IntRect, viewstate *ViewState, pipeline *ForwardPipeline, renderdata *SceneForwardPipelineRenderData, resources *PipelineResources, views *SceneForwardPipelinePassViewId, aaa *ForwardPipelineAAA, aaaconfig *ForwardPipelineAAAConfig, frame int32, framebuffer *FrameBufferHandle) {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	viewstateToC := viewstate.h
	pipelineToC := pipeline.h
	renderdataToC := renderdata.h
	resourcesToC := resources.h
	viewsToC := views.h
	aaaToC := aaa.h
	aaaconfigToC := aaaconfig.h
	frameToC := C.int32_t(frame)
	framebufferToC := framebuffer.h
	C.WrapSubmitSceneToForwardPipelineWithAaaAaaConfigFrameFrameBuffer(viewidToC, sceneToC, rectToC, viewstateToC, pipelineToC, renderdataToC, resourcesToC, viewsToC, aaaToC, aaaconfigToC, frameToC, framebufferToC)
}

// SubmitSceneToForwardPipelineWithAaaAaaConfigFrameFrameBufferDebugName Submit a scene to a forward pipeline.  See [harfang.PrepareSceneForwardPipelineCommonRenderData] and [harfang.PrepareSceneForwardPipelineViewDependentRenderData] if you need to render the same scene from different points of view.
func SubmitSceneToForwardPipelineWithAaaAaaConfigFrameFrameBufferDebugName(viewid *uint16, scene *Scene, rect *IntRect, viewstate *ViewState, pipeline *ForwardPipeline, renderdata *SceneForwardPipelineRenderData, resources *PipelineResources, views *SceneForwardPipelinePassViewId, aaa *ForwardPipelineAAA, aaaconfig *ForwardPipelineAAAConfig, frame int32, framebuffer *FrameBufferHandle, debugname string) {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	viewstateToC := viewstate.h
	pipelineToC := pipeline.h
	renderdataToC := renderdata.h
	resourcesToC := resources.h
	viewsToC := views.h
	aaaToC := aaa.h
	aaaconfigToC := aaaconfig.h
	frameToC := C.int32_t(frame)
	framebufferToC := framebuffer.h
	debugnameToC, idFindebugnameToC := wrapString(debugname)
	defer idFindebugnameToC()
	C.WrapSubmitSceneToForwardPipelineWithAaaAaaConfigFrameFrameBufferDebugName(viewidToC, sceneToC, rectToC, viewstateToC, pipelineToC, renderdataToC, resourcesToC, viewsToC, aaaToC, aaaconfigToC, frameToC, framebufferToC, debugnameToC)
}

// SubmitSceneToPipeline See [harfang.SubmitSceneToForwardPipeline].
func SubmitSceneToPipeline(viewid *uint16, scene *Scene, rect *IntRect, viewstate *ViewState, pipeline *ForwardPipeline, resources *PipelineResources) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	viewstateToC := viewstate.h
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	C.WrapSubmitSceneToPipeline(viewidToC, sceneToC, rectToC, viewstateToC, pipelineToC, resourcesToC, viewsToC)
	return views
}

// SubmitSceneToPipelineWithFb See [harfang.SubmitSceneToForwardPipeline].
func SubmitSceneToPipelineWithFb(viewid *uint16, scene *Scene, rect *IntRect, viewstate *ViewState, pipeline *ForwardPipeline, resources *PipelineResources, fb *FrameBufferHandle) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	viewstateToC := viewstate.h
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	fbToC := fb.h
	C.WrapSubmitSceneToPipelineWithFb(viewidToC, sceneToC, rectToC, viewstateToC, pipelineToC, resourcesToC, viewsToC, fbToC)
	return views
}

// SubmitSceneToPipelineWithFbDebugName See [harfang.SubmitSceneToForwardPipeline].
func SubmitSceneToPipelineWithFbDebugName(viewid *uint16, scene *Scene, rect *IntRect, viewstate *ViewState, pipeline *ForwardPipeline, resources *PipelineResources, fb *FrameBufferHandle, debugname string) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	viewstateToC := viewstate.h
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	fbToC := fb.h
	debugnameToC, idFindebugnameToC := wrapString(debugname)
	defer idFindebugnameToC()
	C.WrapSubmitSceneToPipelineWithFbDebugName(viewidToC, sceneToC, rectToC, viewstateToC, pipelineToC, resourcesToC, viewsToC, fbToC, debugnameToC)
	return views
}

// SubmitSceneToPipelineWithFovAxisIsHorizontal See [harfang.SubmitSceneToForwardPipeline].
func SubmitSceneToPipelineWithFovAxisIsHorizontal(viewid *uint16, scene *Scene, rect *IntRect, fovaxisishorizontal bool, pipeline *ForwardPipeline, resources *PipelineResources) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	fovaxisishorizontalToC := C.bool(fovaxisishorizontal)
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	C.WrapSubmitSceneToPipelineWithFovAxisIsHorizontal(viewidToC, sceneToC, rectToC, fovaxisishorizontalToC, pipelineToC, resourcesToC, viewsToC)
	return views
}

// SubmitSceneToPipelineWithFovAxisIsHorizontalFb See [harfang.SubmitSceneToForwardPipeline].
func SubmitSceneToPipelineWithFovAxisIsHorizontalFb(viewid *uint16, scene *Scene, rect *IntRect, fovaxisishorizontal bool, pipeline *ForwardPipeline, resources *PipelineResources, fb *FrameBufferHandle) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	fovaxisishorizontalToC := C.bool(fovaxisishorizontal)
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	fbToC := fb.h
	C.WrapSubmitSceneToPipelineWithFovAxisIsHorizontalFb(viewidToC, sceneToC, rectToC, fovaxisishorizontalToC, pipelineToC, resourcesToC, viewsToC, fbToC)
	return views
}

// SubmitSceneToPipelineWithFovAxisIsHorizontalFbDebugName See [harfang.SubmitSceneToForwardPipeline].
func SubmitSceneToPipelineWithFovAxisIsHorizontalFbDebugName(viewid *uint16, scene *Scene, rect *IntRect, fovaxisishorizontal bool, pipeline *ForwardPipeline, resources *PipelineResources, fb *FrameBufferHandle, debugname string) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	fovaxisishorizontalToC := C.bool(fovaxisishorizontal)
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	fbToC := fb.h
	debugnameToC, idFindebugnameToC := wrapString(debugname)
	defer idFindebugnameToC()
	C.WrapSubmitSceneToPipelineWithFovAxisIsHorizontalFbDebugName(viewidToC, sceneToC, rectToC, fovaxisishorizontalToC, pipelineToC, resourcesToC, viewsToC, fbToC, debugnameToC)
	return views
}

// SubmitSceneToPipelineWithAaaAaaConfigFrame See [harfang.SubmitSceneToForwardPipeline].
func SubmitSceneToPipelineWithAaaAaaConfigFrame(viewid *uint16, scene *Scene, rect *IntRect, viewstate *ViewState, pipeline *ForwardPipeline, resources *PipelineResources, aaa *ForwardPipelineAAA, aaaconfig *ForwardPipelineAAAConfig, frame int32) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	viewstateToC := viewstate.h
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	aaaToC := aaa.h
	aaaconfigToC := aaaconfig.h
	frameToC := C.int32_t(frame)
	C.WrapSubmitSceneToPipelineWithAaaAaaConfigFrame(viewidToC, sceneToC, rectToC, viewstateToC, pipelineToC, resourcesToC, viewsToC, aaaToC, aaaconfigToC, frameToC)
	return views
}

// SubmitSceneToPipelineWithAaaAaaConfigFrameFrameBuffer See [harfang.SubmitSceneToForwardPipeline].
func SubmitSceneToPipelineWithAaaAaaConfigFrameFrameBuffer(viewid *uint16, scene *Scene, rect *IntRect, viewstate *ViewState, pipeline *ForwardPipeline, resources *PipelineResources, aaa *ForwardPipelineAAA, aaaconfig *ForwardPipelineAAAConfig, frame int32, framebuffer *FrameBufferHandle) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	viewstateToC := viewstate.h
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	aaaToC := aaa.h
	aaaconfigToC := aaaconfig.h
	frameToC := C.int32_t(frame)
	framebufferToC := framebuffer.h
	C.WrapSubmitSceneToPipelineWithAaaAaaConfigFrameFrameBuffer(viewidToC, sceneToC, rectToC, viewstateToC, pipelineToC, resourcesToC, viewsToC, aaaToC, aaaconfigToC, frameToC, framebufferToC)
	return views
}

// SubmitSceneToPipelineWithAaaAaaConfigFrameFrameBufferDebugName See [harfang.SubmitSceneToForwardPipeline].
func SubmitSceneToPipelineWithAaaAaaConfigFrameFrameBufferDebugName(viewid *uint16, scene *Scene, rect *IntRect, viewstate *ViewState, pipeline *ForwardPipeline, resources *PipelineResources, aaa *ForwardPipelineAAA, aaaconfig *ForwardPipelineAAAConfig, frame int32, framebuffer *FrameBufferHandle, debugname string) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	viewstateToC := viewstate.h
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	aaaToC := aaa.h
	aaaconfigToC := aaaconfig.h
	frameToC := C.int32_t(frame)
	framebufferToC := framebuffer.h
	debugnameToC, idFindebugnameToC := wrapString(debugname)
	defer idFindebugnameToC()
	C.WrapSubmitSceneToPipelineWithAaaAaaConfigFrameFrameBufferDebugName(viewidToC, sceneToC, rectToC, viewstateToC, pipelineToC, resourcesToC, viewsToC, aaaToC, aaaconfigToC, frameToC, framebufferToC, debugnameToC)
	return views
}

// SubmitSceneToPipelineWithFovAxisIsHorizontalAaaAaaConfigFrame See [harfang.SubmitSceneToForwardPipeline].
func SubmitSceneToPipelineWithFovAxisIsHorizontalAaaAaaConfigFrame(viewid *uint16, scene *Scene, rect *IntRect, fovaxisishorizontal bool, pipeline *ForwardPipeline, resources *PipelineResources, aaa *ForwardPipelineAAA, aaaconfig *ForwardPipelineAAAConfig, frame int32) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	fovaxisishorizontalToC := C.bool(fovaxisishorizontal)
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	aaaToC := aaa.h
	aaaconfigToC := aaaconfig.h
	frameToC := C.int32_t(frame)
	C.WrapSubmitSceneToPipelineWithFovAxisIsHorizontalAaaAaaConfigFrame(viewidToC, sceneToC, rectToC, fovaxisishorizontalToC, pipelineToC, resourcesToC, viewsToC, aaaToC, aaaconfigToC, frameToC)
	return views
}

// SubmitSceneToPipelineWithFovAxisIsHorizontalAaaAaaConfigFrameFrameBuffer See [harfang.SubmitSceneToForwardPipeline].
func SubmitSceneToPipelineWithFovAxisIsHorizontalAaaAaaConfigFrameFrameBuffer(viewid *uint16, scene *Scene, rect *IntRect, fovaxisishorizontal bool, pipeline *ForwardPipeline, resources *PipelineResources, aaa *ForwardPipelineAAA, aaaconfig *ForwardPipelineAAAConfig, frame int32, framebuffer *FrameBufferHandle) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	fovaxisishorizontalToC := C.bool(fovaxisishorizontal)
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	aaaToC := aaa.h
	aaaconfigToC := aaaconfig.h
	frameToC := C.int32_t(frame)
	framebufferToC := framebuffer.h
	C.WrapSubmitSceneToPipelineWithFovAxisIsHorizontalAaaAaaConfigFrameFrameBuffer(viewidToC, sceneToC, rectToC, fovaxisishorizontalToC, pipelineToC, resourcesToC, viewsToC, aaaToC, aaaconfigToC, frameToC, framebufferToC)
	return views
}

// SubmitSceneToPipelineWithFovAxisIsHorizontalAaaAaaConfigFrameFrameBufferDebugName See [harfang.SubmitSceneToForwardPipeline].
func SubmitSceneToPipelineWithFovAxisIsHorizontalAaaAaaConfigFrameFrameBufferDebugName(viewid *uint16, scene *Scene, rect *IntRect, fovaxisishorizontal bool, pipeline *ForwardPipeline, resources *PipelineResources, aaa *ForwardPipelineAAA, aaaconfig *ForwardPipelineAAAConfig, frame int32, framebuffer *FrameBufferHandle, debugname string) *SceneForwardPipelinePassViewId {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	sceneToC := scene.h
	rectToC := rect.h
	fovaxisishorizontalToC := C.bool(fovaxisishorizontal)
	pipelineToC := pipeline.h
	resourcesToC := resources.h
	views := NewSceneForwardPipelinePassViewId()
	viewsToC := views.h
	aaaToC := aaa.h
	aaaconfigToC := aaaconfig.h
	frameToC := C.int32_t(frame)
	framebufferToC := framebuffer.h
	debugnameToC, idFindebugnameToC := wrapString(debugname)
	defer idFindebugnameToC()
	C.WrapSubmitSceneToPipelineWithFovAxisIsHorizontalAaaAaaConfigFrameFrameBufferDebugName(viewidToC, sceneToC, rectToC, fovaxisishorizontalToC, pipelineToC, resourcesToC, viewsToC, aaaToC, aaaconfigToC, frameToC, framebufferToC, debugnameToC)
	return views
}

// LoadForwardPipelineAAAConfigFromFile ...
func LoadForwardPipelineAAAConfigFromFile(path string, config *ForwardPipelineAAAConfig) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	configToC := config.h
	retval := C.WrapLoadForwardPipelineAAAConfigFromFile(pathToC, configToC)
	return bool(retval)
}

// LoadForwardPipelineAAAConfigFromAssets ...
func LoadForwardPipelineAAAConfigFromAssets(path string, config *ForwardPipelineAAAConfig) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	configToC := config.h
	retval := C.WrapLoadForwardPipelineAAAConfigFromAssets(pathToC, configToC)
	return bool(retval)
}

// SaveForwardPipelineAAAConfigToFile ...
func SaveForwardPipelineAAAConfigToFile(path string, config *ForwardPipelineAAAConfig) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	configToC := config.h
	retval := C.WrapSaveForwardPipelineAAAConfigToFile(pathToC, configToC)
	return bool(retval)
}

// CreateForwardPipelineAAAFromFile ...
func CreateForwardPipelineAAAFromFile(path string, config *ForwardPipelineAAAConfig) *ForwardPipelineAAA {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	configToC := config.h
	retval := C.WrapCreateForwardPipelineAAAFromFile(pathToC, configToC)
	retvalGO := &ForwardPipelineAAA{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineAAA) {
		C.WrapForwardPipelineAAAFree(cleanval.h)
	})
	return retvalGO
}

// CreateForwardPipelineAAAFromFileWithSsgiRatio ...
func CreateForwardPipelineAAAFromFileWithSsgiRatio(path string, config *ForwardPipelineAAAConfig, ssgiratio BackbufferRatio) *ForwardPipelineAAA {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	configToC := config.h
	ssgiratioToC := C.int32_t(ssgiratio)
	retval := C.WrapCreateForwardPipelineAAAFromFileWithSsgiRatio(pathToC, configToC, ssgiratioToC)
	retvalGO := &ForwardPipelineAAA{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineAAA) {
		C.WrapForwardPipelineAAAFree(cleanval.h)
	})
	return retvalGO
}

// CreateForwardPipelineAAAFromFileWithSsgiRatioSsrRatio ...
func CreateForwardPipelineAAAFromFileWithSsgiRatioSsrRatio(path string, config *ForwardPipelineAAAConfig, ssgiratio BackbufferRatio, ssrratio BackbufferRatio) *ForwardPipelineAAA {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	configToC := config.h
	ssgiratioToC := C.int32_t(ssgiratio)
	ssrratioToC := C.int32_t(ssrratio)
	retval := C.WrapCreateForwardPipelineAAAFromFileWithSsgiRatioSsrRatio(pathToC, configToC, ssgiratioToC, ssrratioToC)
	retvalGO := &ForwardPipelineAAA{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineAAA) {
		C.WrapForwardPipelineAAAFree(cleanval.h)
	})
	return retvalGO
}

// CreateForwardPipelineAAAFromAssets ...
func CreateForwardPipelineAAAFromAssets(path string, config *ForwardPipelineAAAConfig) *ForwardPipelineAAA {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	configToC := config.h
	retval := C.WrapCreateForwardPipelineAAAFromAssets(pathToC, configToC)
	retvalGO := &ForwardPipelineAAA{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineAAA) {
		C.WrapForwardPipelineAAAFree(cleanval.h)
	})
	return retvalGO
}

// CreateForwardPipelineAAAFromAssetsWithSsgiRatio ...
func CreateForwardPipelineAAAFromAssetsWithSsgiRatio(path string, config *ForwardPipelineAAAConfig, ssgiratio BackbufferRatio) *ForwardPipelineAAA {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	configToC := config.h
	ssgiratioToC := C.int32_t(ssgiratio)
	retval := C.WrapCreateForwardPipelineAAAFromAssetsWithSsgiRatio(pathToC, configToC, ssgiratioToC)
	retvalGO := &ForwardPipelineAAA{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineAAA) {
		C.WrapForwardPipelineAAAFree(cleanval.h)
	})
	return retvalGO
}

// CreateForwardPipelineAAAFromAssetsWithSsgiRatioSsrRatio ...
func CreateForwardPipelineAAAFromAssetsWithSsgiRatioSsrRatio(path string, config *ForwardPipelineAAAConfig, ssgiratio BackbufferRatio, ssrratio BackbufferRatio) *ForwardPipelineAAA {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	configToC := config.h
	ssgiratioToC := C.int32_t(ssgiratio)
	ssrratioToC := C.int32_t(ssrratio)
	retval := C.WrapCreateForwardPipelineAAAFromAssetsWithSsgiRatioSsrRatio(pathToC, configToC, ssgiratioToC, ssrratioToC)
	retvalGO := &ForwardPipelineAAA{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipelineAAA) {
		C.WrapForwardPipelineAAAFree(cleanval.h)
	})
	return retvalGO
}

// DestroyForwardPipelineAAA ...
func DestroyForwardPipelineAAA(pipeline *ForwardPipelineAAA) {
	pipelineToC := pipeline.h
	C.WrapDestroyForwardPipelineAAA(pipelineToC)
}

// DebugSceneExplorer ...
func DebugSceneExplorer(scene *Scene, name string) {
	sceneToC := scene.h
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	C.WrapDebugSceneExplorer(sceneToC, nameToC)
}

// GetNodesInContact ...
func GetNodesInContact(scene *Scene, with *Node, nodepaircontacts *NodePairContacts) *NodeList {
	sceneToC := scene.h
	withToC := with.h
	nodepaircontactsToC := nodepaircontacts.h
	retval := C.WrapGetNodesInContact(sceneToC, withToC, nodepaircontactsToC)
	retvalGO := &NodeList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *NodeList) {
		C.WrapNodeListFree(cleanval.h)
	})
	return retvalGO
}

// GetNodePairContacts ...
func GetNodePairContacts(first *Node, second *Node, nodepaircontacts *NodePairContacts) *ContactList {
	firstToC := first.h
	secondToC := second.h
	nodepaircontactsToC := nodepaircontacts.h
	retval := C.WrapGetNodePairContacts(firstToC, secondToC, nodepaircontactsToC)
	retvalGO := &ContactList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ContactList) {
		C.WrapContactListFree(cleanval.h)
	})
	return retvalGO
}

// SceneSyncToSystemsFromFile Synchronize optional systems (eg. physics or script) states with the scene states. Load resources from the local filesystem if required.  See [harfang.man.Assets].
func SceneSyncToSystemsFromFile(scene *Scene, vm *SceneLuaVM) {
	sceneToC := scene.h
	vmToC := vm.h
	C.WrapSceneSyncToSystemsFromFile(sceneToC, vmToC)
}

// SceneSyncToSystemsFromFileWithPhysics Synchronize optional systems (eg. physics or script) states with the scene states. Load resources from the local filesystem if required.  See [harfang.man.Assets].
func SceneSyncToSystemsFromFileWithPhysics(scene *Scene, physics *SceneBullet3Physics) {
	sceneToC := scene.h
	physicsToC := physics.h
	C.WrapSceneSyncToSystemsFromFileWithPhysics(sceneToC, physicsToC)
}

// SceneSyncToSystemsFromFileWithPhysicsVm Synchronize optional systems (eg. physics or script) states with the scene states. Load resources from the local filesystem if required.  See [harfang.man.Assets].
func SceneSyncToSystemsFromFileWithPhysicsVm(scene *Scene, physics *SceneBullet3Physics, vm *SceneLuaVM) {
	sceneToC := scene.h
	physicsToC := physics.h
	vmToC := vm.h
	C.WrapSceneSyncToSystemsFromFileWithPhysicsVm(sceneToC, physicsToC, vmToC)
}

// SceneSyncToSystemsFromAssets Synchronize optional systems (eg. physics or script) states with the scene states. Load resources from the assets system if required.  See [harfang.man.Assets].
func SceneSyncToSystemsFromAssets(scene *Scene, vm *SceneLuaVM) {
	sceneToC := scene.h
	vmToC := vm.h
	C.WrapSceneSyncToSystemsFromAssets(sceneToC, vmToC)
}

// SceneSyncToSystemsFromAssetsWithPhysics Synchronize optional systems (eg. physics or script) states with the scene states. Load resources from the assets system if required.  See [harfang.man.Assets].
func SceneSyncToSystemsFromAssetsWithPhysics(scene *Scene, physics *SceneBullet3Physics) {
	sceneToC := scene.h
	physicsToC := physics.h
	C.WrapSceneSyncToSystemsFromAssetsWithPhysics(sceneToC, physicsToC)
}

// SceneSyncToSystemsFromAssetsWithPhysicsVm Synchronize optional systems (eg. physics or script) states with the scene states. Load resources from the assets system if required.  See [harfang.man.Assets].
func SceneSyncToSystemsFromAssetsWithPhysicsVm(scene *Scene, physics *SceneBullet3Physics, vm *SceneLuaVM) {
	sceneToC := scene.h
	physicsToC := physics.h
	vmToC := vm.h
	C.WrapSceneSyncToSystemsFromAssetsWithPhysicsVm(sceneToC, physicsToC, vmToC)
}

// SceneUpdateSystems Update a scene and all its optional systems.
func SceneUpdateSystems(scene *Scene, clocks *SceneClocks, dt int64) {
	sceneToC := scene.h
	clocksToC := clocks.h
	dtToC := C.int64_t(dt)
	C.WrapSceneUpdateSystems(sceneToC, clocksToC, dtToC)
}

// SceneUpdateSystemsWithVm Update a scene and all its optional systems.
func SceneUpdateSystemsWithVm(scene *Scene, clocks *SceneClocks, dt int64, vm *SceneLuaVM) {
	sceneToC := scene.h
	clocksToC := clocks.h
	dtToC := C.int64_t(dt)
	vmToC := vm.h
	C.WrapSceneUpdateSystemsWithVm(sceneToC, clocksToC, dtToC, vmToC)
}

// SceneUpdateSystemsWithPhysicsStepMaxPhysicsStep Update a scene and all its optional systems.
func SceneUpdateSystemsWithPhysicsStepMaxPhysicsStep(scene *Scene, clocks *SceneClocks, dt int64, physics *SceneBullet3Physics, step int64, maxphysicsstep int32) {
	sceneToC := scene.h
	clocksToC := clocks.h
	dtToC := C.int64_t(dt)
	physicsToC := physics.h
	stepToC := C.int64_t(step)
	maxphysicsstepToC := C.int32_t(maxphysicsstep)
	C.WrapSceneUpdateSystemsWithPhysicsStepMaxPhysicsStep(sceneToC, clocksToC, dtToC, physicsToC, stepToC, maxphysicsstepToC)
}

// SceneUpdateSystemsWithPhysicsStepMaxPhysicsStepVm Update a scene and all its optional systems.
func SceneUpdateSystemsWithPhysicsStepMaxPhysicsStepVm(scene *Scene, clocks *SceneClocks, dt int64, physics *SceneBullet3Physics, step int64, maxphysicsstep int32, vm *SceneLuaVM) {
	sceneToC := scene.h
	clocksToC := clocks.h
	dtToC := C.int64_t(dt)
	physicsToC := physics.h
	stepToC := C.int64_t(step)
	maxphysicsstepToC := C.int32_t(maxphysicsstep)
	vmToC := vm.h
	C.WrapSceneUpdateSystemsWithPhysicsStepMaxPhysicsStepVm(sceneToC, clocksToC, dtToC, physicsToC, stepToC, maxphysicsstepToC, vmToC)
}

// SceneUpdateSystemsWithPhysicsContactsStepMaxPhysicsStep Update a scene and all its optional systems.
func SceneUpdateSystemsWithPhysicsContactsStepMaxPhysicsStep(scene *Scene, clocks *SceneClocks, dt int64, physics *SceneBullet3Physics, contacts *NodePairContacts, step int64, maxphysicsstep int32) {
	sceneToC := scene.h
	clocksToC := clocks.h
	dtToC := C.int64_t(dt)
	physicsToC := physics.h
	contactsToC := contacts.h
	stepToC := C.int64_t(step)
	maxphysicsstepToC := C.int32_t(maxphysicsstep)
	C.WrapSceneUpdateSystemsWithPhysicsContactsStepMaxPhysicsStep(sceneToC, clocksToC, dtToC, physicsToC, contactsToC, stepToC, maxphysicsstepToC)
}

// SceneUpdateSystemsWithPhysicsContactsStepMaxPhysicsStepVm Update a scene and all its optional systems.
func SceneUpdateSystemsWithPhysicsContactsStepMaxPhysicsStepVm(scene *Scene, clocks *SceneClocks, dt int64, physics *SceneBullet3Physics, contacts *NodePairContacts, step int64, maxphysicsstep int32, vm *SceneLuaVM) {
	sceneToC := scene.h
	clocksToC := clocks.h
	dtToC := C.int64_t(dt)
	physicsToC := physics.h
	contactsToC := contacts.h
	stepToC := C.int64_t(step)
	maxphysicsstepToC := C.int32_t(maxphysicsstep)
	vmToC := vm.h
	C.WrapSceneUpdateSystemsWithPhysicsContactsStepMaxPhysicsStepVm(sceneToC, clocksToC, dtToC, physicsToC, contactsToC, stepToC, maxphysicsstepToC, vmToC)
}

// SceneGarbageCollectSystems Garbage collect a scene and all its optional systems.
func SceneGarbageCollectSystems(scene *Scene) int32 {
	sceneToC := scene.h
	retval := C.WrapSceneGarbageCollectSystems(sceneToC)
	return int32(retval)
}

// SceneGarbageCollectSystemsWithVm Garbage collect a scene and all its optional systems.
func SceneGarbageCollectSystemsWithVm(scene *Scene, vm *SceneLuaVM) int32 {
	sceneToC := scene.h
	vmToC := vm.h
	retval := C.WrapSceneGarbageCollectSystemsWithVm(sceneToC, vmToC)
	return int32(retval)
}

// SceneGarbageCollectSystemsWithPhysics Garbage collect a scene and all its optional systems.
func SceneGarbageCollectSystemsWithPhysics(scene *Scene, physics *SceneBullet3Physics) int32 {
	sceneToC := scene.h
	physicsToC := physics.h
	retval := C.WrapSceneGarbageCollectSystemsWithPhysics(sceneToC, physicsToC)
	return int32(retval)
}

// SceneGarbageCollectSystemsWithPhysicsVm Garbage collect a scene and all its optional systems.
func SceneGarbageCollectSystemsWithPhysicsVm(scene *Scene, physics *SceneBullet3Physics, vm *SceneLuaVM) int32 {
	sceneToC := scene.h
	physicsToC := physics.h
	vmToC := vm.h
	retval := C.WrapSceneGarbageCollectSystemsWithPhysicsVm(sceneToC, physicsToC, vmToC)
	return int32(retval)
}

// SceneClearSystems Clear scene and all optional systems.
func SceneClearSystems(scene *Scene) {
	sceneToC := scene.h
	C.WrapSceneClearSystems(sceneToC)
}

// SceneClearSystemsWithVm Clear scene and all optional systems.
func SceneClearSystemsWithVm(scene *Scene, vm *SceneLuaVM) {
	sceneToC := scene.h
	vmToC := vm.h
	C.WrapSceneClearSystemsWithVm(sceneToC, vmToC)
}

// SceneClearSystemsWithPhysics Clear scene and all optional systems.
func SceneClearSystemsWithPhysics(scene *Scene, physics *SceneBullet3Physics) {
	sceneToC := scene.h
	physicsToC := physics.h
	C.WrapSceneClearSystemsWithPhysics(sceneToC, physicsToC)
}

// SceneClearSystemsWithPhysicsVm Clear scene and all optional systems.
func SceneClearSystemsWithPhysicsVm(scene *Scene, physics *SceneBullet3Physics, vm *SceneLuaVM) {
	sceneToC := scene.h
	physicsToC := physics.h
	vmToC := vm.h
	C.WrapSceneClearSystemsWithPhysicsVm(sceneToC, physicsToC, vmToC)
}

// InputInit Initialize the Input system. Must be invoked before any call to [harfang.WindowSystemInit] to work properly.  ```python hg.InputInit() hg.WindowSystemInit() ```
func InputInit() {
	C.WrapInputInit()
}

// InputShutdown Shutdown the Input system.
func InputShutdown() {
	C.WrapInputShutdown()
}

// ReadMouse Read the current state of a named mouse. If no name is passed, `default` is implied.  See [harfang.GetMouseNames].
func ReadMouse() *MouseState {
	retval := C.WrapReadMouse()
	retvalGO := &MouseState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MouseState) {
		C.WrapMouseStateFree(cleanval.h)
	})
	return retvalGO
}

// ReadMouseWithName Read the current state of a named mouse. If no name is passed, `default` is implied.  See [harfang.GetMouseNames].
func ReadMouseWithName(name string) *MouseState {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapReadMouseWithName(nameToC)
	retvalGO := &MouseState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *MouseState) {
		C.WrapMouseStateFree(cleanval.h)
	})
	return retvalGO
}

// GetMouseNames Return a list of names for all supported mouse devices on the system.  See [harfang.ReadKeyboard].
func GetMouseNames() *StringList {
	retval := C.WrapGetMouseNames()
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// ReadKeyboard Read the current state of a named keyboard. If no name is passed, `default` is implied.  See [harfang.GetKeyboardNames].
func ReadKeyboard() *KeyboardState {
	retval := C.WrapReadKeyboard()
	retvalGO := &KeyboardState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *KeyboardState) {
		C.WrapKeyboardStateFree(cleanval.h)
	})
	return retvalGO
}

// ReadKeyboardWithName Read the current state of a named keyboard. If no name is passed, `default` is implied.  See [harfang.GetKeyboardNames].
func ReadKeyboardWithName(name string) *KeyboardState {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapReadKeyboardWithName(nameToC)
	retvalGO := &KeyboardState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *KeyboardState) {
		C.WrapKeyboardStateFree(cleanval.h)
	})
	return retvalGO
}

// GetKeyName Return the name for a keyboard key.
func GetKeyName(key Key) string {
	keyToC := C.int32_t(key)
	retval := C.WrapGetKeyName(keyToC)
	retvalGO := string(C.GoString(retval))
	return retvalGO
}

// GetKeyNameWithName Return the name for a keyboard key.
func GetKeyNameWithName(key Key, name string) string {
	keyToC := C.int32_t(key)
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapGetKeyNameWithName(keyToC, nameToC)
	retvalGO := string(C.GoString(retval))
	return retvalGO
}

// GetKeyboardNames Return a list of names for all supported keyboard devices on the system.  See [harfang.ReadKeyboard].
func GetKeyboardNames() *StringList {
	retval := C.WrapGetKeyboardNames()
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// ReadGamepad Read the current state of a named gamepad. If no name is passed, `default` is implied.  See [harfang.GetGamepadNames].
func ReadGamepad() *GamepadState {
	retval := C.WrapReadGamepad()
	retvalGO := &GamepadState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *GamepadState) {
		C.WrapGamepadStateFree(cleanval.h)
	})
	return retvalGO
}

// ReadGamepadWithName Read the current state of a named gamepad. If no name is passed, `default` is implied.  See [harfang.GetGamepadNames].
func ReadGamepadWithName(name string) *GamepadState {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapReadGamepadWithName(nameToC)
	retvalGO := &GamepadState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *GamepadState) {
		C.WrapGamepadStateFree(cleanval.h)
	})
	return retvalGO
}

// GetGamepadNames Return a list of names for all supported gamepad devices on the system.  See [harfang.ReadGamepad].
func GetGamepadNames() *StringList {
	retval := C.WrapGetGamepadNames()
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// ReadJoystick ...
func ReadJoystick() *JoystickState {
	retval := C.WrapReadJoystick()
	retvalGO := &JoystickState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *JoystickState) {
		C.WrapJoystickStateFree(cleanval.h)
	})
	return retvalGO
}

// ReadJoystickWithName ...
func ReadJoystickWithName(name string) *JoystickState {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapReadJoystickWithName(nameToC)
	retvalGO := &JoystickState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *JoystickState) {
		C.WrapJoystickStateFree(cleanval.h)
	})
	return retvalGO
}

// GetJoystickNames ...
func GetJoystickNames() *StringList {
	retval := C.WrapGetJoystickNames()
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// GetJoystickDeviceNames ...
func GetJoystickDeviceNames() *StringList {
	retval := C.WrapGetJoystickDeviceNames()
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// ReadVRController Read the current state of a named VR controller. If no name is passed, `default` is implied.  See [harfang.GetVRControllerNames].
func ReadVRController() *VRControllerState {
	retval := C.WrapReadVRController()
	retvalGO := &VRControllerState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VRControllerState) {
		C.WrapVRControllerStateFree(cleanval.h)
	})
	return retvalGO
}

// ReadVRControllerWithName Read the current state of a named VR controller. If no name is passed, `default` is implied.  See [harfang.GetVRControllerNames].
func ReadVRControllerWithName(name string) *VRControllerState {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapReadVRControllerWithName(nameToC)
	retvalGO := &VRControllerState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VRControllerState) {
		C.WrapVRControllerStateFree(cleanval.h)
	})
	return retvalGO
}

// SendVRControllerHapticPulse Send an haptic pulse to a named VR controller.  See [harfang.GetVRControllerNames].
func SendVRControllerHapticPulse(duration int64) {
	durationToC := C.int64_t(duration)
	C.WrapSendVRControllerHapticPulse(durationToC)
}

// SendVRControllerHapticPulseWithName Send an haptic pulse to a named VR controller.  See [harfang.GetVRControllerNames].
func SendVRControllerHapticPulseWithName(duration int64, name string) {
	durationToC := C.int64_t(duration)
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	C.WrapSendVRControllerHapticPulseWithName(durationToC, nameToC)
}

// GetVRControllerNames Return a list of names for all supported VR controller devices on the system.  See [harfang.ReadVRController].
func GetVRControllerNames() *StringList {
	retval := C.WrapGetVRControllerNames()
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// ReadVRGenericTracker Read the current state of a named VR generic tracked. If no name is passed, `default` is implied.  See [harfang.GetVRGenericTrackerNames].
func ReadVRGenericTracker() *VRGenericTrackerState {
	retval := C.WrapReadVRGenericTracker()
	retvalGO := &VRGenericTrackerState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VRGenericTrackerState) {
		C.WrapVRGenericTrackerStateFree(cleanval.h)
	})
	return retvalGO
}

// ReadVRGenericTrackerWithName Read the current state of a named VR generic tracked. If no name is passed, `default` is implied.  See [harfang.GetVRGenericTrackerNames].
func ReadVRGenericTrackerWithName(name string) *VRGenericTrackerState {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapReadVRGenericTrackerWithName(nameToC)
	retvalGO := &VRGenericTrackerState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *VRGenericTrackerState) {
		C.WrapVRGenericTrackerStateFree(cleanval.h)
	})
	return retvalGO
}

// GetVRGenericTrackerNames Return a list of names for all supported VR tracker devices on the system.
func GetVRGenericTrackerNames() *StringList {
	retval := C.WrapGetVRGenericTrackerNames()
	retvalGO := &StringList{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *StringList) {
		C.WrapStringListFree(cleanval.h)
	})
	return retvalGO
}

// ImGuiNewFrame ...
func ImGuiNewFrame() {
	C.WrapImGuiNewFrame()
}

// ImGuiRender ...
func ImGuiRender() {
	C.WrapImGuiRender()
}

// ImGuiBegin Start a new window.
func ImGuiBegin(name string) bool {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapImGuiBegin(nameToC)
	return bool(retval)
}

// ImGuiBeginWithOpenFlags Start a new window.
func ImGuiBeginWithOpenFlags(name string, open *bool, flags ImGuiWindowFlags) bool {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	openToC := (*C.bool)(unsafe.Pointer(open))
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiBeginWithOpenFlags(nameToC, openToC, flagsToC)
	return bool(retval)
}

// ImGuiEnd End the current window.
func ImGuiEnd() {
	C.WrapImGuiEnd()
}

// ImGuiBeginChild Begin a scrolling region.
func ImGuiBeginChild(id string) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	retval := C.WrapImGuiBeginChild(idToC)
	return bool(retval)
}

// ImGuiBeginChildWithSize Begin a scrolling region.
func ImGuiBeginChildWithSize(id string, size *Vec2) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	sizeToC := size.h
	retval := C.WrapImGuiBeginChildWithSize(idToC, sizeToC)
	return bool(retval)
}

// ImGuiBeginChildWithSizeBorder Begin a scrolling region.
func ImGuiBeginChildWithSizeBorder(id string, size *Vec2, border bool) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	sizeToC := size.h
	borderToC := C.bool(border)
	retval := C.WrapImGuiBeginChildWithSizeBorder(idToC, sizeToC, borderToC)
	return bool(retval)
}

// ImGuiBeginChildWithSizeBorderFlags Begin a scrolling region.
func ImGuiBeginChildWithSizeBorderFlags(id string, size *Vec2, border bool, flags ImGuiWindowFlags) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	sizeToC := size.h
	borderToC := C.bool(border)
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiBeginChildWithSizeBorderFlags(idToC, sizeToC, borderToC, flagsToC)
	return bool(retval)
}

// ImGuiEndChild End a scrolling region.
func ImGuiEndChild() {
	C.WrapImGuiEndChild()
}

// ImGuiGetContentRegionMax Return the available content space including window decorations and scrollbar.
func ImGuiGetContentRegionMax() *Vec2 {
	retval := C.WrapImGuiGetContentRegionMax()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetContentRegionAvail Get available space for content in the current layout.
func ImGuiGetContentRegionAvail() *Vec2 {
	retval := C.WrapImGuiGetContentRegionAvail()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetContentRegionAvailWidth Helper function to return the available width of current content region.  See [harfang.ImGuiGetContentRegionAvail].
func ImGuiGetContentRegionAvailWidth() float32 {
	retval := C.WrapImGuiGetContentRegionAvailWidth()
	return float32(retval)
}

// ImGuiGetWindowContentRegionMin Content boundaries min (roughly (0,0)-Scroll), in window space.
func ImGuiGetWindowContentRegionMin() *Vec2 {
	retval := C.WrapImGuiGetWindowContentRegionMin()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetWindowContentRegionMax Return the content boundaries max (roughly (0,0)+Size-Scroll) where Size can be override with [harfang.ImGuiSetNextWindowContentSize], in window space.
func ImGuiGetWindowContentRegionMax() *Vec2 {
	retval := C.WrapImGuiGetWindowContentRegionMax()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetWindowContentRegionWidth Return the width of the content region.
func ImGuiGetWindowContentRegionWidth() float32 {
	retval := C.WrapImGuiGetWindowContentRegionWidth()
	return float32(retval)
}

// ImGuiGetWindowDrawList Get the draw list associated to the current window, to append your own drawing primitives.
func ImGuiGetWindowDrawList() *ImDrawList {
	retval := C.WrapImGuiGetWindowDrawList()
	var retvalGO *ImDrawList
	if retval != nil {
		retvalGO = &ImDrawList{h: retval}
		runtime.SetFinalizer(retvalGO, func(cleanval *ImDrawList) {
			C.WrapImDrawListFree(cleanval.h)
		})
	}
	return retvalGO
}

// ImGuiGetWindowPos Return the current window position in screen space.  See [harfang.ImGuiSetWindowPos].
func ImGuiGetWindowPos() *Vec2 {
	retval := C.WrapImGuiGetWindowPos()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetWindowSize Return the current window size.  See [harfang.ImGuiSetWindowSize].
func ImGuiGetWindowSize() *Vec2 {
	retval := C.WrapImGuiGetWindowSize()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetWindowWidth Return the current window width.
func ImGuiGetWindowWidth() float32 {
	retval := C.WrapImGuiGetWindowWidth()
	return float32(retval)
}

// ImGuiGetWindowHeight Return the current window height.
func ImGuiGetWindowHeight() float32 {
	retval := C.WrapImGuiGetWindowHeight()
	return float32(retval)
}

// ImGuiIsWindowCollapsed Is the current window collapsed.
func ImGuiIsWindowCollapsed() bool {
	retval := C.WrapImGuiIsWindowCollapsed()
	return bool(retval)
}

// ImGuiSetWindowFontScale Per-window font scale.
func ImGuiSetWindowFontScale(scale float32) {
	scaleToC := C.float(scale)
	C.WrapImGuiSetWindowFontScale(scaleToC)
}

// ImGuiSetNextWindowPos Set next window position, call before [harfang.ImGuiBegin].
func ImGuiSetNextWindowPos(pos *Vec2) {
	posToC := pos.h
	C.WrapImGuiSetNextWindowPos(posToC)
}

// ImGuiSetNextWindowPosWithCondition Set next window position, call before [harfang.ImGuiBegin].
func ImGuiSetNextWindowPosWithCondition(pos *Vec2, condition ImGuiCond) {
	posToC := pos.h
	conditionToC := C.int32_t(condition)
	C.WrapImGuiSetNextWindowPosWithCondition(posToC, conditionToC)
}

// ImGuiSetNextWindowPosCenter Set next window position to be centered on screen, call before [harfang.ImGuiBegin].
func ImGuiSetNextWindowPosCenter() {
	C.WrapImGuiSetNextWindowPosCenter()
}

// ImGuiSetNextWindowPosCenterWithCondition Set next window position to be centered on screen, call before [harfang.ImGuiBegin].
func ImGuiSetNextWindowPosCenterWithCondition(condition ImGuiCond) {
	conditionToC := C.int32_t(condition)
	C.WrapImGuiSetNextWindowPosCenterWithCondition(conditionToC)
}

// ImGuiSetNextWindowSize Set next window size, call before [harfang.ImGuiBegin]. A value of 0 for an axis will auto-fit it.
func ImGuiSetNextWindowSize(size *Vec2) {
	sizeToC := size.h
	C.WrapImGuiSetNextWindowSize(sizeToC)
}

// ImGuiSetNextWindowSizeWithCondition Set next window size, call before [harfang.ImGuiBegin]. A value of 0 for an axis will auto-fit it.
func ImGuiSetNextWindowSizeWithCondition(size *Vec2, condition ImGuiCond) {
	sizeToC := size.h
	conditionToC := C.int32_t(condition)
	C.WrapImGuiSetNextWindowSizeWithCondition(sizeToC, conditionToC)
}

// ImGuiSetNextWindowSizeConstraints Set the next window size limits.  Use -1,-1 on either X/Y axis to preserve the current size. Sizes will be rounded down.
func ImGuiSetNextWindowSizeConstraints(sizemin *Vec2, sizemax *Vec2) {
	sizeminToC := sizemin.h
	sizemaxToC := sizemax.h
	C.WrapImGuiSetNextWindowSizeConstraints(sizeminToC, sizemaxToC)
}

// ImGuiSetNextWindowContentSize Set the size of the content area of the next declared window. Call before [harfang.ImGuiBegin].
func ImGuiSetNextWindowContentSize(size *Vec2) {
	sizeToC := size.h
	C.WrapImGuiSetNextWindowContentSize(sizeToC)
}

// ImGuiSetNextWindowContentWidth See [harfang.ImGuiSetNextWindowContentSize].
func ImGuiSetNextWindowContentWidth(width float32) {
	widthToC := C.float(width)
	C.WrapImGuiSetNextWindowContentWidth(widthToC)
}

// ImGuiSetNextWindowCollapsed Set next window collapsed state, call before [harfang.ImGuiBegin].
func ImGuiSetNextWindowCollapsed(collapsed bool, condition ImGuiCond) {
	collapsedToC := C.bool(collapsed)
	conditionToC := C.int32_t(condition)
	C.WrapImGuiSetNextWindowCollapsed(collapsedToC, conditionToC)
}

// ImGuiSetNextWindowFocus Set the next window to be focused/top-most. Call before [harfang.ImGuiBegin].
func ImGuiSetNextWindowFocus() {
	C.WrapImGuiSetNextWindowFocus()
}

// ImGuiSetWindowPos Set named window position.
func ImGuiSetWindowPos(name string, pos *Vec2) {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	posToC := pos.h
	C.WrapImGuiSetWindowPos(nameToC, posToC)
}

// ImGuiSetWindowPosWithCondition Set named window position.
func ImGuiSetWindowPosWithCondition(name string, pos *Vec2, condition ImGuiCond) {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	posToC := pos.h
	conditionToC := C.int32_t(condition)
	C.WrapImGuiSetWindowPosWithCondition(nameToC, posToC, conditionToC)
}

// ImGuiSetWindowSize Set named window size.
func ImGuiSetWindowSize(name string, size *Vec2) {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sizeToC := size.h
	C.WrapImGuiSetWindowSize(nameToC, sizeToC)
}

// ImGuiSetWindowSizeWithCondition Set named window size.
func ImGuiSetWindowSizeWithCondition(name string, size *Vec2, condition ImGuiCond) {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sizeToC := size.h
	conditionToC := C.int32_t(condition)
	C.WrapImGuiSetWindowSizeWithCondition(nameToC, sizeToC, conditionToC)
}

// ImGuiSetWindowCollapsed Set named window collapsed state, prefer using [harfang.ImGuiSetNextWindowCollapsed].
func ImGuiSetWindowCollapsed(name string, collapsed bool) {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	collapsedToC := C.bool(collapsed)
	C.WrapImGuiSetWindowCollapsed(nameToC, collapsedToC)
}

// ImGuiSetWindowCollapsedWithCondition Set named window collapsed state, prefer using [harfang.ImGuiSetNextWindowCollapsed].
func ImGuiSetWindowCollapsedWithCondition(name string, collapsed bool, condition ImGuiCond) {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	collapsedToC := C.bool(collapsed)
	conditionToC := C.int32_t(condition)
	C.WrapImGuiSetWindowCollapsedWithCondition(nameToC, collapsedToC, conditionToC)
}

// ImGuiSetWindowFocus Set named window to be focused/top-most.
func ImGuiSetWindowFocus(name string) {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	C.WrapImGuiSetWindowFocus(nameToC)
}

// ImGuiGetScrollX Get scrolling amount on the horizontal axis.
func ImGuiGetScrollX() float32 {
	retval := C.WrapImGuiGetScrollX()
	return float32(retval)
}

// ImGuiGetScrollY Get scrolling amount on the vertical axis.
func ImGuiGetScrollY() float32 {
	retval := C.WrapImGuiGetScrollY()
	return float32(retval)
}

// ImGuiGetScrollMaxX Get maximum scrolling amount on the horizontal axis.
func ImGuiGetScrollMaxX() float32 {
	retval := C.WrapImGuiGetScrollMaxX()
	return float32(retval)
}

// ImGuiGetScrollMaxY Get maximum scrolling amount on the vertical axis.
func ImGuiGetScrollMaxY() float32 {
	retval := C.WrapImGuiGetScrollMaxY()
	return float32(retval)
}

// ImGuiSetScrollX Set scrolling amount between [harfang.0;[ImGuiGetScrollMaxX]].
func ImGuiSetScrollX(scrollx float32) {
	scrollxToC := C.float(scrollx)
	C.WrapImGuiSetScrollX(scrollxToC)
}

// ImGuiSetScrollY Set scrolling amount between [harfang.0;[ImGuiGetScrollMaxY]].
func ImGuiSetScrollY(scrolly float32) {
	scrollyToC := C.float(scrolly)
	C.WrapImGuiSetScrollY(scrollyToC)
}

// ImGuiSetScrollHereY Adjust scrolling amount to make current cursor position visible.  - 0: Top. - 0.5: Center. - 1: Bottom.  When using to make a default/current item visible, consider using [harfang.ImGuiSetItemDefaultFocus] instead.
func ImGuiSetScrollHereY() {
	C.WrapImGuiSetScrollHereY()
}

// ImGuiSetScrollHereYWithCenterYRatio Adjust scrolling amount to make current cursor position visible.  - 0: Top. - 0.5: Center. - 1: Bottom.  When using to make a default/current item visible, consider using [harfang.ImGuiSetItemDefaultFocus] instead.
func ImGuiSetScrollHereYWithCenterYRatio(centeryratio float32) {
	centeryratioToC := C.float(centeryratio)
	C.WrapImGuiSetScrollHereYWithCenterYRatio(centeryratioToC)
}

// ImGuiSetScrollFromPosY Adjust scrolling amount to make a given position visible. Generally [harfang.ImGuiGetCursorStartPos] + offset to compute a valid position.
func ImGuiSetScrollFromPosY(posy float32) {
	posyToC := C.float(posy)
	C.WrapImGuiSetScrollFromPosY(posyToC)
}

// ImGuiSetScrollFromPosYWithCenterYRatio Adjust scrolling amount to make a given position visible. Generally [harfang.ImGuiGetCursorStartPos] + offset to compute a valid position.
func ImGuiSetScrollFromPosYWithCenterYRatio(posy float32, centeryratio float32) {
	posyToC := C.float(posy)
	centeryratioToC := C.float(centeryratio)
	C.WrapImGuiSetScrollFromPosYWithCenterYRatio(posyToC, centeryratioToC)
}

// ImGuiSetKeyboardFocusHere Focus keyboard on the next widget.  Use positive `offset` value to access sub components of a multiple component widget. Use `-1` to access the previous widget.
func ImGuiSetKeyboardFocusHere() {
	C.WrapImGuiSetKeyboardFocusHere()
}

// ImGuiSetKeyboardFocusHereWithOffset Focus keyboard on the next widget.  Use positive `offset` value to access sub components of a multiple component widget. Use `-1` to access the previous widget.
func ImGuiSetKeyboardFocusHereWithOffset(offset int32) {
	offsetToC := C.int32_t(offset)
	C.WrapImGuiSetKeyboardFocusHereWithOffset(offsetToC)
}

// ImGuiPushFont Push a font on top of the font stack and make it current for subsequent text rendering operations.
func ImGuiPushFont(font *ImFont) {
	fontToC := font.h
	C.WrapImGuiPushFont(fontToC)
}

// ImGuiPopFont Undo the last call to [harfang.ImGuiPushFont].
func ImGuiPopFont() {
	C.WrapImGuiPopFont()
}

// ImGuiPushStyleColor Push a value on the style stack for the specified style color.  See [harfang.ImGuiPopStyleColor].
func ImGuiPushStyleColor(idx ImGuiCol, color *Color) {
	idxToC := C.int32_t(idx)
	colorToC := color.h
	C.WrapImGuiPushStyleColor(idxToC, colorToC)
}

// ImGuiPopStyleColor Undo the last call to [harfang.ImGuiPushStyleColor].
func ImGuiPopStyleColor() {
	C.WrapImGuiPopStyleColor()
}

// ImGuiPopStyleColorWithCount Undo the last call to [harfang.ImGuiPushStyleColor].
func ImGuiPopStyleColorWithCount(count int32) {
	countToC := C.int32_t(count)
	C.WrapImGuiPopStyleColorWithCount(countToC)
}

// ImGuiPushStyleVar Push a value on the style stack for the specified style variable.  See [harfang.ImGuiPopStyleVar].
func ImGuiPushStyleVar(idx ImGuiStyleVar, value float32) {
	idxToC := C.int32_t(idx)
	valueToC := C.float(value)
	C.WrapImGuiPushStyleVar(idxToC, valueToC)
}

// ImGuiPushStyleVarWithValue Push a value on the style stack for the specified style variable.  See [harfang.ImGuiPopStyleVar].
func ImGuiPushStyleVarWithValue(idx ImGuiStyleVar, value *Vec2) {
	idxToC := C.int32_t(idx)
	valueToC := value.h
	C.WrapImGuiPushStyleVarWithValue(idxToC, valueToC)
}

// ImGuiPopStyleVar Undo the last call to [harfang.ImGuiPushStyleVar].
func ImGuiPopStyleVar() {
	C.WrapImGuiPopStyleVar()
}

// ImGuiPopStyleVarWithCount Undo the last call to [harfang.ImGuiPushStyleVar].
func ImGuiPopStyleVarWithCount(count int32) {
	countToC := C.int32_t(count)
	C.WrapImGuiPopStyleVarWithCount(countToC)
}

// ImGuiGetFont Return the current ImGui font.
func ImGuiGetFont() *ImFont {
	retval := C.WrapImGuiGetFont()
	var retvalGO *ImFont
	if retval != nil {
		retvalGO = &ImFont{h: retval}
		runtime.SetFinalizer(retvalGO, func(cleanval *ImFont) {
			C.WrapImFontFree(cleanval.h)
		})
	}
	return retvalGO
}

// ImGuiGetFontSize Return the font size (height in pixels) of the current ImGui font with the current scale applied.
func ImGuiGetFontSize() float32 {
	retval := C.WrapImGuiGetFontSize()
	return float32(retval)
}

// ImGuiGetFontTexUvWhitePixel Get UV coordinate for a while pixel, useful to draw custom shapes via the ImDrawList API.
func ImGuiGetFontTexUvWhitePixel() *Vec2 {
	retval := C.WrapImGuiGetFontTexUvWhitePixel()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetColorU32 Return a style color component as a 32 bit unsigned integer.  See [harfang.ImGuiPushStyleColor].
func ImGuiGetColorU32(idx ImGuiCol) uint32 {
	idxToC := C.int32_t(idx)
	retval := C.WrapImGuiGetColorU32(idxToC)
	return uint32(retval)
}

// ImGuiGetColorU32WithAlphaMultiplier Return a style color component as a 32 bit unsigned integer.  See [harfang.ImGuiPushStyleColor].
func ImGuiGetColorU32WithAlphaMultiplier(idx ImGuiCol, alphamultiplier float32) uint32 {
	idxToC := C.int32_t(idx)
	alphamultiplierToC := C.float(alphamultiplier)
	retval := C.WrapImGuiGetColorU32WithAlphaMultiplier(idxToC, alphamultiplierToC)
	return uint32(retval)
}

// ImGuiGetColorU32WithColor Return a style color component as a 32 bit unsigned integer.  See [harfang.ImGuiPushStyleColor].
func ImGuiGetColorU32WithColor(color *Color) uint32 {
	colorToC := color.h
	retval := C.WrapImGuiGetColorU32WithColor(colorToC)
	return uint32(retval)
}

// ImGuiPushItemWidth Set the width of items for common large `item+label` widgets.  - `>0`: width in pixels - `<0`: align `x` pixels to the right of window (so -1 always align width to the right side) - `=0`: default to ~2/3 of the window width  See [harfang.ImGuiPopItemWidth].
func ImGuiPushItemWidth(itemwidth float32) {
	itemwidthToC := C.float(itemwidth)
	C.WrapImGuiPushItemWidth(itemwidthToC)
}

// ImGuiPopItemWidth Undo the last call to [harfang.ImGuiPushItemWidth].
func ImGuiPopItemWidth() {
	C.WrapImGuiPopItemWidth()
}

// ImGuiCalcItemWidth Returns the width of item given pushed settings and current cursor position.   Note: This is not necessarily the width of last item.
func ImGuiCalcItemWidth() float32 {
	retval := C.WrapImGuiCalcItemWidth()
	return float32(retval)
}

// ImGuiPushTextWrapPos Push word-wrapping position for text commands.  - `<0`: No wrapping. - `=0`: Wrap to the end of the window or column. - `>0`: Wrap at `wrap_pos_x` position in window local space.  See [harfang.ImGuiPopTextWrapPos].
func ImGuiPushTextWrapPos() {
	C.WrapImGuiPushTextWrapPos()
}

// ImGuiPushTextWrapPosWithWrapPosX Push word-wrapping position for text commands.  - `<0`: No wrapping. - `=0`: Wrap to the end of the window or column. - `>0`: Wrap at `wrap_pos_x` position in window local space.  See [harfang.ImGuiPopTextWrapPos].
func ImGuiPushTextWrapPosWithWrapPosX(wrapposx float32) {
	wrapposxToC := C.float(wrapposx)
	C.WrapImGuiPushTextWrapPosWithWrapPosX(wrapposxToC)
}

// ImGuiPopTextWrapPos Undo the last call to [harfang.ImGuiPushTextWrapPos].
func ImGuiPopTextWrapPos() {
	C.WrapImGuiPopTextWrapPos()
}

// ImGuiPushAllowKeyboardFocus Allow focusing using TAB/Shift-TAB, enabled by default but you can disable it for certain widgets.
func ImGuiPushAllowKeyboardFocus(v bool) {
	vToC := C.bool(v)
	C.WrapImGuiPushAllowKeyboardFocus(vToC)
}

// ImGuiPopAllowKeyboardFocus Undo the last call to [harfang.ImGuiPushAllowKeyboardFocus].
func ImGuiPopAllowKeyboardFocus() {
	C.WrapImGuiPopAllowKeyboardFocus()
}

// ImGuiPushButtonRepeat In repeat mode, `ButtonXXX` functions return repeated true in a typematic manner.  Note that you can call [harfang.ImGuiIsItemActive] after any `Button` to tell if the button is held in the current frame.
func ImGuiPushButtonRepeat(repeat bool) {
	repeatToC := C.bool(repeat)
	C.WrapImGuiPushButtonRepeat(repeatToC)
}

// ImGuiPopButtonRepeat Undo the last call to [harfang.ImGuiPushButtonRepeat].
func ImGuiPopButtonRepeat() {
	C.WrapImGuiPopButtonRepeat()
}

// ImGuiSeparator Output an horizontal line to separate two distinct UI sections.
func ImGuiSeparator() {
	C.WrapImGuiSeparator()
}

// ImGuiSameLine Call between widgets or groups to layout them horizontally.
func ImGuiSameLine() {
	C.WrapImGuiSameLine()
}

// ImGuiSameLineWithPosX Call between widgets or groups to layout them horizontally.
func ImGuiSameLineWithPosX(posx float32) {
	posxToC := C.float(posx)
	C.WrapImGuiSameLineWithPosX(posxToC)
}

// ImGuiSameLineWithPosXSpacingW Call between widgets or groups to layout them horizontally.
func ImGuiSameLineWithPosXSpacingW(posx float32, spacingw float32) {
	posxToC := C.float(posx)
	spacingwToC := C.float(spacingw)
	C.WrapImGuiSameLineWithPosXSpacingW(posxToC, spacingwToC)
}

// ImGuiNewLine Undo a [harfang.ImGuiSameLine] call or force a new line when in an horizontal layout.
func ImGuiNewLine() {
	C.WrapImGuiNewLine()
}

// ImGuiSpacing Add spacing.
func ImGuiSpacing() {
	C.WrapImGuiSpacing()
}

// ImGuiDummy Add a dummy item of given size.
func ImGuiDummy(size *Vec2) {
	sizeToC := size.h
	C.WrapImGuiDummy(sizeToC)
}

// ImGuiIndent Move content position toward the right.
func ImGuiIndent() {
	C.WrapImGuiIndent()
}

// ImGuiIndentWithWidth Move content position toward the right.
func ImGuiIndentWithWidth(width float32) {
	widthToC := C.float(width)
	C.WrapImGuiIndentWithWidth(widthToC)
}

// ImGuiUnindent Move content position back to the left (cancel [harfang.ImGuiIndent]).
func ImGuiUnindent() {
	C.WrapImGuiUnindent()
}

// ImGuiUnindentWithWidth Move content position back to the left (cancel [harfang.ImGuiIndent]).
func ImGuiUnindentWithWidth(width float32) {
	widthToC := C.float(width)
	C.WrapImGuiUnindentWithWidth(widthToC)
}

// ImGuiBeginGroup Lock horizontal starting position. Once closing a group it is seen as a single item (so you can use [harfang.ImGuiIsItemHovered] on a group, [harfang.ImGuiSameLine] between groups, etc...).
func ImGuiBeginGroup() {
	C.WrapImGuiBeginGroup()
}

// ImGuiEndGroup End the current group.
func ImGuiEndGroup() {
	C.WrapImGuiEndGroup()
}

// ImGuiGetCursorPos Return the layout cursor position in window space. Next widget declaration will take place at the cursor position.  See [harfang.ImGuiSetCursorPos] and [harfang.ImGuiSameLine].
func ImGuiGetCursorPos() *Vec2 {
	retval := C.WrapImGuiGetCursorPos()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetCursorPosX Helper for [harfang.ImGuiGetCursorPos].
func ImGuiGetCursorPosX() float32 {
	retval := C.WrapImGuiGetCursorPosX()
	return float32(retval)
}

// ImGuiGetCursorPosY Helper for [harfang.ImGuiGetCursorPos].
func ImGuiGetCursorPosY() float32 {
	retval := C.WrapImGuiGetCursorPosY()
	return float32(retval)
}

// ImGuiSetCursorPos Set the current widget output cursor position in window space.
func ImGuiSetCursorPos(localpos *Vec2) {
	localposToC := localpos.h
	C.WrapImGuiSetCursorPos(localposToC)
}

// ImGuiSetCursorPosX See [harfang.ImGuiSetCursorPos].
func ImGuiSetCursorPosX(x float32) {
	xToC := C.float(x)
	C.WrapImGuiSetCursorPosX(xToC)
}

// ImGuiSetCursorPosY See [harfang.ImGuiSetCursorPos].
func ImGuiSetCursorPosY(y float32) {
	yToC := C.float(y)
	C.WrapImGuiSetCursorPosY(yToC)
}

// ImGuiGetCursorStartPos Return the current layout \"line\" starting position.  See [harfang.ImGuiSameLine].
func ImGuiGetCursorStartPos() *Vec2 {
	retval := C.WrapImGuiGetCursorStartPos()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetCursorScreenPos Return the current layout cursor position in screen space.
func ImGuiGetCursorScreenPos() *Vec2 {
	retval := C.WrapImGuiGetCursorScreenPos()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiSetCursorScreenPos Set the widget cursor output position in screen space.
func ImGuiSetCursorScreenPos(pos *Vec2) {
	posToC := pos.h
	C.WrapImGuiSetCursorScreenPos(posToC)
}

// ImGuiAlignTextToFramePadding Vertically align upcoming text baseline to FramePadding __y__ coordinate so that it will align properly to regularly framed items.
func ImGuiAlignTextToFramePadding() {
	C.WrapImGuiAlignTextToFramePadding()
}

// ImGuiGetTextLineHeight Return the height of a text line using the current font.  See [harfang.ImGuiPushFont].
func ImGuiGetTextLineHeight() float32 {
	retval := C.WrapImGuiGetTextLineHeight()
	return float32(retval)
}

// ImGuiGetTextLineHeightWithSpacing Return the height of a text line using the current font plus vertical spacing between two layout lines.  See [harfang.ImGuiGetTextLineHeight].
func ImGuiGetTextLineHeightWithSpacing() float32 {
	retval := C.WrapImGuiGetTextLineHeightWithSpacing()
	return float32(retval)
}

// ImGuiGetFrameHeightWithSpacing Return the following value: FontSize + style.FramePadding.y * 2 + style.ItemSpacing.y (distance in pixels between 2 consecutive lines of framed widgets)
func ImGuiGetFrameHeightWithSpacing() float32 {
	retval := C.WrapImGuiGetFrameHeightWithSpacing()
	return float32(retval)
}

// ImGuiColumns Begin a column layout section.  To move to the next column use [harfang.ImGuiNextColumn]. To end a column layout section pass `1` to this function.  **Note:** Current implementation supports a maximum of 64 columns.
func ImGuiColumns() {
	C.WrapImGuiColumns()
}

// ImGuiColumnsWithCount Begin a column layout section.  To move to the next column use [harfang.ImGuiNextColumn]. To end a column layout section pass `1` to this function.  **Note:** Current implementation supports a maximum of 64 columns.
func ImGuiColumnsWithCount(count int32) {
	countToC := C.int32_t(count)
	C.WrapImGuiColumnsWithCount(countToC)
}

// ImGuiColumnsWithCountId Begin a column layout section.  To move to the next column use [harfang.ImGuiNextColumn]. To end a column layout section pass `1` to this function.  **Note:** Current implementation supports a maximum of 64 columns.
func ImGuiColumnsWithCountId(count int32, id string) {
	countToC := C.int32_t(count)
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	C.WrapImGuiColumnsWithCountId(countToC, idToC)
}

// ImGuiColumnsWithCountIdWithBorder Begin a column layout section.  To move to the next column use [harfang.ImGuiNextColumn]. To end a column layout section pass `1` to this function.  **Note:** Current implementation supports a maximum of 64 columns.
func ImGuiColumnsWithCountIdWithBorder(count int32, id string, withborder bool) {
	countToC := C.int32_t(count)
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	withborderToC := C.bool(withborder)
	C.WrapImGuiColumnsWithCountIdWithBorder(countToC, idToC, withborderToC)
}

// ImGuiNextColumn Start the next column in multi-column layout.  See [harfang.ImGuiColumns].
func ImGuiNextColumn() {
	C.WrapImGuiNextColumn()
}

// ImGuiGetColumnIndex Returns the index of the current column.
func ImGuiGetColumnIndex() int32 {
	retval := C.WrapImGuiGetColumnIndex()
	return int32(retval)
}

// ImGuiGetColumnOffset Returns the current column offset in pixels, from the left side of the content region.
func ImGuiGetColumnOffset() float32 {
	retval := C.WrapImGuiGetColumnOffset()
	return float32(retval)
}

// ImGuiGetColumnOffsetWithColumnIndex Returns the current column offset in pixels, from the left side of the content region.
func ImGuiGetColumnOffsetWithColumnIndex(columnindex int32) float32 {
	columnindexToC := C.int32_t(columnindex)
	retval := C.WrapImGuiGetColumnOffsetWithColumnIndex(columnindexToC)
	return float32(retval)
}

// ImGuiSetColumnOffset Set the position of a column line in pixels, from the left side of the contents region.
func ImGuiSetColumnOffset(columnindex int32, offsetx float32) {
	columnindexToC := C.int32_t(columnindex)
	offsetxToC := C.float(offsetx)
	C.WrapImGuiSetColumnOffset(columnindexToC, offsetxToC)
}

// ImGuiGetColumnWidth Returns the current column width in pixels.
func ImGuiGetColumnWidth() float32 {
	retval := C.WrapImGuiGetColumnWidth()
	return float32(retval)
}

// ImGuiGetColumnWidthWithColumnIndex Returns the current column width in pixels.
func ImGuiGetColumnWidthWithColumnIndex(columnindex int32) float32 {
	columnindexToC := C.int32_t(columnindex)
	retval := C.WrapImGuiGetColumnWidthWithColumnIndex(columnindexToC)
	return float32(retval)
}

// ImGuiSetColumnWidth Set the column width in pixels.
func ImGuiSetColumnWidth(columnindex int32, width float32) {
	columnindexToC := C.int32_t(columnindex)
	widthToC := C.float(width)
	C.WrapImGuiSetColumnWidth(columnindexToC, widthToC)
}

// ImGuiGetColumnsCount Return the number of columns in the current layout section.  See [harfang.ImGuiColumns].
func ImGuiGetColumnsCount() int32 {
	retval := C.WrapImGuiGetColumnsCount()
	return int32(retval)
}

// ImGuiPushID Push a string into the ID stack.
func ImGuiPushID(id string) {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	C.WrapImGuiPushID(idToC)
}

// ImGuiPushIDWithId Push a string into the ID stack.
func ImGuiPushIDWithId(id int32) {
	idToC := C.int32_t(id)
	C.WrapImGuiPushIDWithId(idToC)
}

// ImGuiPopID Undo the last call to [harfang.ImGuiPushID].
func ImGuiPopID() {
	C.WrapImGuiPopID()
}

// ImGuiGetID Return a unique ImGui ID.
func ImGuiGetID(id string) uint32 {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	retval := C.WrapImGuiGetID(idToC)
	return uint32(retval)
}

// ImGuiText Static text.
func ImGuiText(text string) {
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	C.WrapImGuiText(textToC)
}

// ImGuiTextColored Colored static text.
func ImGuiTextColored(color *Color, text string) {
	colorToC := color.h
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	C.WrapImGuiTextColored(colorToC, textToC)
}

// ImGuiTextDisabled Disabled static text.
func ImGuiTextDisabled(text string) {
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	C.WrapImGuiTextDisabled(textToC)
}

// ImGuiTextWrapped Wrapped static text.  Note that this won't work on an auto-resizing window if there's no other widgets to extend the window width, you may need to set a size using [harfang.ImGuiSetNextWindowSize].
func ImGuiTextWrapped(text string) {
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	C.WrapImGuiTextWrapped(textToC)
}

// ImGuiTextUnformatted Raw text without formatting. Roughly equivalent to [harfang.ImGuiText] but faster, recommended for long chunks of text.
func ImGuiTextUnformatted(text string) {
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	C.WrapImGuiTextUnformatted(textToC)
}

// ImGuiLabelText Display text+label aligned the same way as value+label widgets.
func ImGuiLabelText(label string, text string) {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	C.WrapImGuiLabelText(labelToC, textToC)
}

// ImGuiBullet Draw a small circle and keep the cursor on the same line. Advances by the same distance as an empty [harfang.ImGuiTreeNode] call.
func ImGuiBullet() {
	C.WrapImGuiBullet()
}

// ImGuiBulletText Draw a bullet followed by a static text.
func ImGuiBulletText(label string) {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	C.WrapImGuiBulletText(labelToC)
}

// ImGuiButton Button widget returning `True` if the button was pressed.
func ImGuiButton(label string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	retval := C.WrapImGuiButton(labelToC)
	return bool(retval)
}

// ImGuiButtonWithSize Button widget returning `True` if the button was pressed.
func ImGuiButtonWithSize(label string, size *Vec2) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	sizeToC := size.h
	retval := C.WrapImGuiButtonWithSize(labelToC, sizeToC)
	return bool(retval)
}

// ImGuiSmallButton Small button widget fitting the height of a text line, return `True` if the button was pressed.
func ImGuiSmallButton(label string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	retval := C.WrapImGuiSmallButton(labelToC)
	return bool(retval)
}

// ImGuiInvisibleButton Invisible button widget, return `True` if the button was pressed.
func ImGuiInvisibleButton(text string, size *Vec2) bool {
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	sizeToC := size.h
	retval := C.WrapImGuiInvisibleButton(textToC, sizeToC)
	return bool(retval)
}

// ImGuiImage Display a texture as an image widget.  See [harfang.ImGuiImageButton].
func ImGuiImage(tex *Texture, size *Vec2) {
	texToC := tex.h
	sizeToC := size.h
	C.WrapImGuiImage(texToC, sizeToC)
}

// ImGuiImageWithUv0 Display a texture as an image widget.  See [harfang.ImGuiImageButton].
func ImGuiImageWithUv0(tex *Texture, size *Vec2, uv0 *Vec2) {
	texToC := tex.h
	sizeToC := size.h
	uv0ToC := uv0.h
	C.WrapImGuiImageWithUv0(texToC, sizeToC, uv0ToC)
}

// ImGuiImageWithUv0Uv1 Display a texture as an image widget.  See [harfang.ImGuiImageButton].
func ImGuiImageWithUv0Uv1(tex *Texture, size *Vec2, uv0 *Vec2, uv1 *Vec2) {
	texToC := tex.h
	sizeToC := size.h
	uv0ToC := uv0.h
	uv1ToC := uv1.h
	C.WrapImGuiImageWithUv0Uv1(texToC, sizeToC, uv0ToC, uv1ToC)
}

// ImGuiImageWithUv0Uv1TintCol Display a texture as an image widget.  See [harfang.ImGuiImageButton].
func ImGuiImageWithUv0Uv1TintCol(tex *Texture, size *Vec2, uv0 *Vec2, uv1 *Vec2, tintcol *Color) {
	texToC := tex.h
	sizeToC := size.h
	uv0ToC := uv0.h
	uv1ToC := uv1.h
	tintcolToC := tintcol.h
	C.WrapImGuiImageWithUv0Uv1TintCol(texToC, sizeToC, uv0ToC, uv1ToC, tintcolToC)
}

// ImGuiImageWithUv0Uv1TintColBorderCol Display a texture as an image widget.  See [harfang.ImGuiImageButton].
func ImGuiImageWithUv0Uv1TintColBorderCol(tex *Texture, size *Vec2, uv0 *Vec2, uv1 *Vec2, tintcol *Color, bordercol *Color) {
	texToC := tex.h
	sizeToC := size.h
	uv0ToC := uv0.h
	uv1ToC := uv1.h
	tintcolToC := tintcol.h
	bordercolToC := bordercol.h
	C.WrapImGuiImageWithUv0Uv1TintColBorderCol(texToC, sizeToC, uv0ToC, uv1ToC, tintcolToC, bordercolToC)
}

// ImGuiImageButton Declare an image button displaying the provided texture.  See [harfang.ImGuiImage].
func ImGuiImageButton(tex *Texture, size *Vec2) bool {
	texToC := tex.h
	sizeToC := size.h
	retval := C.WrapImGuiImageButton(texToC, sizeToC)
	return bool(retval)
}

// ImGuiImageButtonWithUv0 Declare an image button displaying the provided texture.  See [harfang.ImGuiImage].
func ImGuiImageButtonWithUv0(tex *Texture, size *Vec2, uv0 *Vec2) bool {
	texToC := tex.h
	sizeToC := size.h
	uv0ToC := uv0.h
	retval := C.WrapImGuiImageButtonWithUv0(texToC, sizeToC, uv0ToC)
	return bool(retval)
}

// ImGuiImageButtonWithUv0Uv1 Declare an image button displaying the provided texture.  See [harfang.ImGuiImage].
func ImGuiImageButtonWithUv0Uv1(tex *Texture, size *Vec2, uv0 *Vec2, uv1 *Vec2) bool {
	texToC := tex.h
	sizeToC := size.h
	uv0ToC := uv0.h
	uv1ToC := uv1.h
	retval := C.WrapImGuiImageButtonWithUv0Uv1(texToC, sizeToC, uv0ToC, uv1ToC)
	return bool(retval)
}

// ImGuiImageButtonWithUv0Uv1FramePadding Declare an image button displaying the provided texture.  See [harfang.ImGuiImage].
func ImGuiImageButtonWithUv0Uv1FramePadding(tex *Texture, size *Vec2, uv0 *Vec2, uv1 *Vec2, framepadding int32) bool {
	texToC := tex.h
	sizeToC := size.h
	uv0ToC := uv0.h
	uv1ToC := uv1.h
	framepaddingToC := C.int32_t(framepadding)
	retval := C.WrapImGuiImageButtonWithUv0Uv1FramePadding(texToC, sizeToC, uv0ToC, uv1ToC, framepaddingToC)
	return bool(retval)
}

// ImGuiImageButtonWithUv0Uv1FramePaddingBgCol Declare an image button displaying the provided texture.  See [harfang.ImGuiImage].
func ImGuiImageButtonWithUv0Uv1FramePaddingBgCol(tex *Texture, size *Vec2, uv0 *Vec2, uv1 *Vec2, framepadding int32, bgcol *Color) bool {
	texToC := tex.h
	sizeToC := size.h
	uv0ToC := uv0.h
	uv1ToC := uv1.h
	framepaddingToC := C.int32_t(framepadding)
	bgcolToC := bgcol.h
	retval := C.WrapImGuiImageButtonWithUv0Uv1FramePaddingBgCol(texToC, sizeToC, uv0ToC, uv1ToC, framepaddingToC, bgcolToC)
	return bool(retval)
}

// ImGuiImageButtonWithUv0Uv1FramePaddingBgColTintCol Declare an image button displaying the provided texture.  See [harfang.ImGuiImage].
func ImGuiImageButtonWithUv0Uv1FramePaddingBgColTintCol(tex *Texture, size *Vec2, uv0 *Vec2, uv1 *Vec2, framepadding int32, bgcol *Color, tintcol *Color) bool {
	texToC := tex.h
	sizeToC := size.h
	uv0ToC := uv0.h
	uv1ToC := uv1.h
	framepaddingToC := C.int32_t(framepadding)
	bgcolToC := bgcol.h
	tintcolToC := tintcol.h
	retval := C.WrapImGuiImageButtonWithUv0Uv1FramePaddingBgColTintCol(texToC, sizeToC, uv0ToC, uv1ToC, framepaddingToC, bgcolToC, tintcolToC)
	return bool(retval)
}

// ImGuiInputText Text input widget, returns the current widget buffer content.
func ImGuiInputText(label string, text string, maxsize int32) (bool, *string) {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	maxsizeToC := C.size_t(maxsize)
	out := new(string)
	outToC1 := C.CString(*out)
	outToC := &outToC1
	retval := C.WrapImGuiInputText(labelToC, textToC, maxsizeToC, outToC)
	outToCGO := string(C.GoString(*outToC))
	return bool(retval), &outToCGO
}

// ImGuiInputTextWithFlags Text input widget, returns the current widget buffer content.
func ImGuiInputTextWithFlags(label string, text string, maxsize int32, flags ImGuiInputTextFlags) (bool, *string) {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	maxsizeToC := C.size_t(maxsize)
	out := new(string)
	outToC1 := C.CString(*out)
	outToC := &outToC1
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiInputTextWithFlags(labelToC, textToC, maxsizeToC, outToC, flagsToC)
	outToCGO := string(C.GoString(*outToC))
	return bool(retval), &outToCGO
}

// ImGuiCheckbox Display a checkbox widget. Returns an interaction flag (user interacted with the widget) and the current widget state (checked or not after user interaction).  ```python was_clicked, my_value = gs.ImGuiCheckBox('My value', my_value) ```
func ImGuiCheckbox(label string, value *bool) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	valueToC := (*C.bool)(unsafe.Pointer(value))
	retval := C.WrapImGuiCheckbox(labelToC, valueToC)
	return bool(retval)
}

// ImGuiRadioButton Radio button widget, return the button state.
func ImGuiRadioButton(label string, active bool) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	activeToC := C.bool(active)
	retval := C.WrapImGuiRadioButton(labelToC, activeToC)
	return bool(retval)
}

// ImGuiRadioButtonWithVVButton Radio button widget, return the button state.
func ImGuiRadioButtonWithVVButton(label string, v *int32, vbutton int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.int32_t)(unsafe.Pointer(v))
	vbuttonToC := C.int32_t(vbutton)
	retval := C.WrapImGuiRadioButtonWithVVButton(labelToC, vToC, vbuttonToC)
	return bool(retval)
}

// ImGuiBeginCombo Begin a ImGui Combo Box.
func ImGuiBeginCombo(label string, previewvalue string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	previewvalueToC, idFinpreviewvalueToC := wrapString(previewvalue)
	defer idFinpreviewvalueToC()
	retval := C.WrapImGuiBeginCombo(labelToC, previewvalueToC)
	return bool(retval)
}

// ImGuiBeginComboWithFlags Begin a ImGui Combo Box.
func ImGuiBeginComboWithFlags(label string, previewvalue string, flags ImGuiComboFlags) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	previewvalueToC, idFinpreviewvalueToC := wrapString(previewvalue)
	defer idFinpreviewvalueToC()
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiBeginComboWithFlags(labelToC, previewvalueToC, flagsToC)
	return bool(retval)
}

// ImGuiEndCombo End a combo widget.
func ImGuiEndCombo() {
	C.WrapImGuiEndCombo()
}

// ImGuiCombo Combo box widget, return the current selection index. Combo items are passed as an array of string.
func ImGuiCombo(label string, currentitem *int32, items *StringList) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	currentitemToC := (*C.int32_t)(unsafe.Pointer(currentitem))
	itemsToC := items.h
	retval := C.WrapImGuiCombo(labelToC, currentitemToC, itemsToC)
	return bool(retval)
}

// ImGuiComboWithHeightInItems Combo box widget, return the current selection index. Combo items are passed as an array of string.
func ImGuiComboWithHeightInItems(label string, currentitem *int32, items *StringList, heightinitems int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	currentitemToC := (*C.int32_t)(unsafe.Pointer(currentitem))
	itemsToC := items.h
	heightinitemsToC := C.int32_t(heightinitems)
	retval := C.WrapImGuiComboWithHeightInItems(labelToC, currentitemToC, itemsToC, heightinitemsToC)
	return bool(retval)
}

// ImGuiComboWithSliceOfItems Combo box widget, return the current selection index. Combo items are passed as an array of string.
func ImGuiComboWithSliceOfItems(label string, currentitem *int32, SliceOfitems GoSliceOfstring) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	currentitemToC := (*C.int32_t)(unsafe.Pointer(currentitem))
	var SliceOfitemsSpecialString []*C.char
	for _, s := range SliceOfitems {
		SliceOfitemsSpecialString = append(SliceOfitemsSpecialString, C.CString(s))
	}
	SliceOfitemsSpecialStringToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfitemsSpecialString))
	SliceOfitemsSpecialStringToCSize := C.size_t(SliceOfitemsSpecialStringToC.Len)
	SliceOfitemsSpecialStringToCBuf := (**C.char)(unsafe.Pointer(SliceOfitemsSpecialStringToC.Data))
	retval := C.WrapImGuiComboWithSliceOfItems(labelToC, currentitemToC, SliceOfitemsSpecialStringToCSize, SliceOfitemsSpecialStringToCBuf)
	return bool(retval)
}

// ImGuiComboWithSliceOfItemsHeightInItems Combo box widget, return the current selection index. Combo items are passed as an array of string.
func ImGuiComboWithSliceOfItemsHeightInItems(label string, currentitem *int32, SliceOfitems GoSliceOfstring, heightinitems int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	currentitemToC := (*C.int32_t)(unsafe.Pointer(currentitem))
	var SliceOfitemsSpecialString []*C.char
	for _, s := range SliceOfitems {
		SliceOfitemsSpecialString = append(SliceOfitemsSpecialString, C.CString(s))
	}
	SliceOfitemsSpecialStringToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfitemsSpecialString))
	SliceOfitemsSpecialStringToCSize := C.size_t(SliceOfitemsSpecialStringToC.Len)
	SliceOfitemsSpecialStringToCBuf := (**C.char)(unsafe.Pointer(SliceOfitemsSpecialStringToC.Data))
	heightinitemsToC := C.int32_t(heightinitems)
	retval := C.WrapImGuiComboWithSliceOfItemsHeightInItems(labelToC, currentitemToC, SliceOfitemsSpecialStringToCSize, SliceOfitemsSpecialStringToCBuf, heightinitemsToC)
	return bool(retval)
}

// ImGuiColorButton Color button widget, display a small colored rectangle.
func ImGuiColorButton(id string, color *Color) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	colorToC := color.h
	retval := C.WrapImGuiColorButton(idToC, colorToC)
	return bool(retval)
}

// ImGuiColorButtonWithFlags Color button widget, display a small colored rectangle.
func ImGuiColorButtonWithFlags(id string, color *Color, flags ImGuiColorEditFlags) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	colorToC := color.h
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiColorButtonWithFlags(idToC, colorToC, flagsToC)
	return bool(retval)
}

// ImGuiColorButtonWithFlagsSize Color button widget, display a small colored rectangle.
func ImGuiColorButtonWithFlagsSize(id string, color *Color, flags ImGuiColorEditFlags, size *Vec2) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	colorToC := color.h
	flagsToC := C.int32_t(flags)
	sizeToC := size.h
	retval := C.WrapImGuiColorButtonWithFlagsSize(idToC, colorToC, flagsToC, sizeToC)
	return bool(retval)
}

// ImGuiColorEdit Color editor, returns the widget current color.
func ImGuiColorEdit(label string, color *Color) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	colorToC := color.h
	retval := C.WrapImGuiColorEdit(labelToC, colorToC)
	return bool(retval)
}

// ImGuiColorEditWithFlags Color editor, returns the widget current color.
func ImGuiColorEditWithFlags(label string, color *Color, flags ImGuiColorEditFlags) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	colorToC := color.h
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiColorEditWithFlags(labelToC, colorToC, flagsToC)
	return bool(retval)
}

// ImGuiProgressBar Draw a progress bar, `fraction` must be between 0.0 and 1.0.
func ImGuiProgressBar(fraction float32) {
	fractionToC := C.float(fraction)
	C.WrapImGuiProgressBar(fractionToC)
}

// ImGuiProgressBarWithSize Draw a progress bar, `fraction` must be between 0.0 and 1.0.
func ImGuiProgressBarWithSize(fraction float32, size *Vec2) {
	fractionToC := C.float(fraction)
	sizeToC := size.h
	C.WrapImGuiProgressBarWithSize(fractionToC, sizeToC)
}

// ImGuiProgressBarWithSizeOverlay Draw a progress bar, `fraction` must be between 0.0 and 1.0.
func ImGuiProgressBarWithSizeOverlay(fraction float32, size *Vec2, overlay string) {
	fractionToC := C.float(fraction)
	sizeToC := size.h
	overlayToC, idFinoverlayToC := wrapString(overlay)
	defer idFinoverlayToC()
	C.WrapImGuiProgressBarWithSizeOverlay(fractionToC, sizeToC, overlayToC)
}

// ImGuiDragFloat Declare a widget to edit a float value. The widget can be dragged over to modify the underlying value.
func ImGuiDragFloat(label string, v *float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.float)(unsafe.Pointer(v))
	retval := C.WrapImGuiDragFloat(labelToC, vToC)
	return bool(retval)
}

// ImGuiDragFloatWithVSpeed Declare a widget to edit a float value. The widget can be dragged over to modify the underlying value.
func ImGuiDragFloatWithVSpeed(label string, v *float32, vspeed float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.float)(unsafe.Pointer(v))
	vspeedToC := C.float(vspeed)
	retval := C.WrapImGuiDragFloatWithVSpeed(labelToC, vToC, vspeedToC)
	return bool(retval)
}

// ImGuiDragFloatWithVSpeedVMinVMax Declare a widget to edit a float value. The widget can be dragged over to modify the underlying value.
func ImGuiDragFloatWithVSpeedVMinVMax(label string, v *float32, vspeed float32, vmin float32, vmax float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.float)(unsafe.Pointer(v))
	vspeedToC := C.float(vspeed)
	vminToC := C.float(vmin)
	vmaxToC := C.float(vmax)
	retval := C.WrapImGuiDragFloatWithVSpeedVMinVMax(labelToC, vToC, vspeedToC, vminToC, vmaxToC)
	return bool(retval)
}

// ImGuiDragVec2 Declare a float edit widget that can be dragged over to modify its value.
func ImGuiDragVec2(label string, v *Vec2) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	retval := C.WrapImGuiDragVec2(labelToC, vToC)
	return bool(retval)
}

// ImGuiDragVec2WithVSpeed Declare a float edit widget that can be dragged over to modify its value.
func ImGuiDragVec2WithVSpeed(label string, v *Vec2, vspeed float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vspeedToC := C.float(vspeed)
	retval := C.WrapImGuiDragVec2WithVSpeed(labelToC, vToC, vspeedToC)
	return bool(retval)
}

// ImGuiDragVec2WithVSpeedVMinVMax Declare a float edit widget that can be dragged over to modify its value.
func ImGuiDragVec2WithVSpeedVMinVMax(label string, v *Vec2, vspeed float32, vmin float32, vmax float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vspeedToC := C.float(vspeed)
	vminToC := C.float(vmin)
	vmaxToC := C.float(vmax)
	retval := C.WrapImGuiDragVec2WithVSpeedVMinVMax(labelToC, vToC, vspeedToC, vminToC, vmaxToC)
	return bool(retval)
}

// ImGuiDragVec3 Declare a widget to edit a [harfang.Vec3] value. The widget can be dragged over to modify the underlying value.
func ImGuiDragVec3(label string, v *Vec3) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	retval := C.WrapImGuiDragVec3(labelToC, vToC)
	return bool(retval)
}

// ImGuiDragVec3WithVSpeed Declare a widget to edit a [harfang.Vec3] value. The widget can be dragged over to modify the underlying value.
func ImGuiDragVec3WithVSpeed(label string, v *Vec3, vspeed float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vspeedToC := C.float(vspeed)
	retval := C.WrapImGuiDragVec3WithVSpeed(labelToC, vToC, vspeedToC)
	return bool(retval)
}

// ImGuiDragVec3WithVSpeedVMinVMax Declare a widget to edit a [harfang.Vec3] value. The widget can be dragged over to modify the underlying value.
func ImGuiDragVec3WithVSpeedVMinVMax(label string, v *Vec3, vspeed float32, vmin float32, vmax float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vspeedToC := C.float(vspeed)
	vminToC := C.float(vmin)
	vmaxToC := C.float(vmax)
	retval := C.WrapImGuiDragVec3WithVSpeedVMinVMax(labelToC, vToC, vspeedToC, vminToC, vmaxToC)
	return bool(retval)
}

// ImGuiDragVec4 Declare a widget to edit a [harfang.Vec4] value. The widget can be dragged over to modify the underlying value.
func ImGuiDragVec4(label string, v *Vec4) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	retval := C.WrapImGuiDragVec4(labelToC, vToC)
	return bool(retval)
}

// ImGuiDragVec4WithVSpeed Declare a widget to edit a [harfang.Vec4] value. The widget can be dragged over to modify the underlying value.
func ImGuiDragVec4WithVSpeed(label string, v *Vec4, vspeed float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vspeedToC := C.float(vspeed)
	retval := C.WrapImGuiDragVec4WithVSpeed(labelToC, vToC, vspeedToC)
	return bool(retval)
}

// ImGuiDragVec4WithVSpeedVMinVMax Declare a widget to edit a [harfang.Vec4] value. The widget can be dragged over to modify the underlying value.
func ImGuiDragVec4WithVSpeedVMinVMax(label string, v *Vec4, vspeed float32, vmin float32, vmax float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vspeedToC := C.float(vspeed)
	vminToC := C.float(vmin)
	vmaxToC := C.float(vmax)
	retval := C.WrapImGuiDragVec4WithVSpeedVMinVMax(labelToC, vToC, vspeedToC, vminToC, vmaxToC)
	return bool(retval)
}

// ImGuiDragIntVec2 Declare a widget to edit an [harfang.iVec2] value. The widget can be dragged over to modify the underlying value.
func ImGuiDragIntVec2(label string, v *IVec2) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	retval := C.WrapImGuiDragIntVec2(labelToC, vToC)
	return bool(retval)
}

// ImGuiDragIntVec2WithVSpeed Declare a widget to edit an [harfang.iVec2] value. The widget can be dragged over to modify the underlying value.
func ImGuiDragIntVec2WithVSpeed(label string, v *IVec2, vspeed float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vspeedToC := C.float(vspeed)
	retval := C.WrapImGuiDragIntVec2WithVSpeed(labelToC, vToC, vspeedToC)
	return bool(retval)
}

// ImGuiDragIntVec2WithVSpeedVMinVMax Declare a widget to edit an [harfang.iVec2] value. The widget can be dragged over to modify the underlying value.
func ImGuiDragIntVec2WithVSpeedVMinVMax(label string, v *IVec2, vspeed float32, vmin int32, vmax int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vspeedToC := C.float(vspeed)
	vminToC := C.int32_t(vmin)
	vmaxToC := C.int32_t(vmax)
	retval := C.WrapImGuiDragIntVec2WithVSpeedVMinVMax(labelToC, vToC, vspeedToC, vminToC, vmaxToC)
	return bool(retval)
}

// ImGuiInputInt Integer field widget.
func ImGuiInputInt(label string, v *int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.int32_t)(unsafe.Pointer(v))
	retval := C.WrapImGuiInputInt(labelToC, vToC)
	return bool(retval)
}

// ImGuiInputIntWithStepStepFast Integer field widget.
func ImGuiInputIntWithStepStepFast(label string, v *int32, step int32, stepfast int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.int32_t)(unsafe.Pointer(v))
	stepToC := C.int32_t(step)
	stepfastToC := C.int32_t(stepfast)
	retval := C.WrapImGuiInputIntWithStepStepFast(labelToC, vToC, stepToC, stepfastToC)
	return bool(retval)
}

// ImGuiInputIntWithStepStepFastFlags Integer field widget.
func ImGuiInputIntWithStepStepFastFlags(label string, v *int32, step int32, stepfast int32, flags ImGuiInputTextFlags) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.int32_t)(unsafe.Pointer(v))
	stepToC := C.int32_t(step)
	stepfastToC := C.int32_t(stepfast)
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiInputIntWithStepStepFastFlags(labelToC, vToC, stepToC, stepfastToC, flagsToC)
	return bool(retval)
}

// ImGuiInputFloat Float field widget.
func ImGuiInputFloat(label string, v *float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.float)(unsafe.Pointer(v))
	retval := C.WrapImGuiInputFloat(labelToC, vToC)
	return bool(retval)
}

// ImGuiInputFloatWithStepStepFast Float field widget.
func ImGuiInputFloatWithStepStepFast(label string, v *float32, step float32, stepfast float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.float)(unsafe.Pointer(v))
	stepToC := C.float(step)
	stepfastToC := C.float(stepfast)
	retval := C.WrapImGuiInputFloatWithStepStepFast(labelToC, vToC, stepToC, stepfastToC)
	return bool(retval)
}

// ImGuiInputFloatWithStepStepFastDecimalPrecision Float field widget.
func ImGuiInputFloatWithStepStepFastDecimalPrecision(label string, v *float32, step float32, stepfast float32, decimalprecision int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.float)(unsafe.Pointer(v))
	stepToC := C.float(step)
	stepfastToC := C.float(stepfast)
	decimalprecisionToC := C.int32_t(decimalprecision)
	retval := C.WrapImGuiInputFloatWithStepStepFastDecimalPrecision(labelToC, vToC, stepToC, stepfastToC, decimalprecisionToC)
	return bool(retval)
}

// ImGuiInputFloatWithStepStepFastDecimalPrecisionFlags Float field widget.
func ImGuiInputFloatWithStepStepFastDecimalPrecisionFlags(label string, v *float32, step float32, stepfast float32, decimalprecision int32, flags ImGuiInputTextFlags) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.float)(unsafe.Pointer(v))
	stepToC := C.float(step)
	stepfastToC := C.float(stepfast)
	decimalprecisionToC := C.int32_t(decimalprecision)
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiInputFloatWithStepStepFastDecimalPrecisionFlags(labelToC, vToC, stepToC, stepfastToC, decimalprecisionToC, flagsToC)
	return bool(retval)
}

// ImGuiInputVec2 [harfang.Vec2] field widget.
func ImGuiInputVec2(label string, v *Vec2) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	retval := C.WrapImGuiInputVec2(labelToC, vToC)
	return bool(retval)
}

// ImGuiInputVec2WithDecimalPrecision [harfang.Vec2] field widget.
func ImGuiInputVec2WithDecimalPrecision(label string, v *Vec2, decimalprecision int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	decimalprecisionToC := C.int32_t(decimalprecision)
	retval := C.WrapImGuiInputVec2WithDecimalPrecision(labelToC, vToC, decimalprecisionToC)
	return bool(retval)
}

// ImGuiInputVec2WithDecimalPrecisionFlags [harfang.Vec2] field widget.
func ImGuiInputVec2WithDecimalPrecisionFlags(label string, v *Vec2, decimalprecision int32, flags ImGuiInputTextFlags) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	decimalprecisionToC := C.int32_t(decimalprecision)
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiInputVec2WithDecimalPrecisionFlags(labelToC, vToC, decimalprecisionToC, flagsToC)
	return bool(retval)
}

// ImGuiInputVec3 [harfang.Vec3] field widget.
func ImGuiInputVec3(label string, v *Vec3) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	retval := C.WrapImGuiInputVec3(labelToC, vToC)
	return bool(retval)
}

// ImGuiInputVec3WithDecimalPrecision [harfang.Vec3] field widget.
func ImGuiInputVec3WithDecimalPrecision(label string, v *Vec3, decimalprecision int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	decimalprecisionToC := C.int32_t(decimalprecision)
	retval := C.WrapImGuiInputVec3WithDecimalPrecision(labelToC, vToC, decimalprecisionToC)
	return bool(retval)
}

// ImGuiInputVec3WithDecimalPrecisionFlags [harfang.Vec3] field widget.
func ImGuiInputVec3WithDecimalPrecisionFlags(label string, v *Vec3, decimalprecision int32, flags ImGuiInputTextFlags) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	decimalprecisionToC := C.int32_t(decimalprecision)
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiInputVec3WithDecimalPrecisionFlags(labelToC, vToC, decimalprecisionToC, flagsToC)
	return bool(retval)
}

// ImGuiInputVec4 [harfang.Vec4] field widget.
func ImGuiInputVec4(label string, v *Vec4) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	retval := C.WrapImGuiInputVec4(labelToC, vToC)
	return bool(retval)
}

// ImGuiInputVec4WithDecimalPrecision [harfang.Vec4] field widget.
func ImGuiInputVec4WithDecimalPrecision(label string, v *Vec4, decimalprecision int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	decimalprecisionToC := C.int32_t(decimalprecision)
	retval := C.WrapImGuiInputVec4WithDecimalPrecision(labelToC, vToC, decimalprecisionToC)
	return bool(retval)
}

// ImGuiInputVec4WithDecimalPrecisionFlags [harfang.Vec4] field widget.
func ImGuiInputVec4WithDecimalPrecisionFlags(label string, v *Vec4, decimalprecision int32, flags ImGuiInputTextFlags) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	decimalprecisionToC := C.int32_t(decimalprecision)
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiInputVec4WithDecimalPrecisionFlags(labelToC, vToC, decimalprecisionToC, flagsToC)
	return bool(retval)
}

// ImGuiInputIntVec2 ...
func ImGuiInputIntVec2(label string, v *IVec2) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	retval := C.WrapImGuiInputIntVec2(labelToC, vToC)
	return bool(retval)
}

// ImGuiInputIntVec2WithFlags ...
func ImGuiInputIntVec2WithFlags(label string, v *IVec2, flags ImGuiInputTextFlags) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiInputIntVec2WithFlags(labelToC, vToC, flagsToC)
	return bool(retval)
}

// ImGuiSliderInt Integer slider widget.
func ImGuiSliderInt(label string, v *int32, vmin int32, vmax int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.int32_t)(unsafe.Pointer(v))
	vminToC := C.int32_t(vmin)
	vmaxToC := C.int32_t(vmax)
	retval := C.WrapImGuiSliderInt(labelToC, vToC, vminToC, vmaxToC)
	return bool(retval)
}

// ImGuiSliderIntWithFormat Integer slider widget.
func ImGuiSliderIntWithFormat(label string, v *int32, vmin int32, vmax int32, format string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.int32_t)(unsafe.Pointer(v))
	vminToC := C.int32_t(vmin)
	vmaxToC := C.int32_t(vmax)
	formatToC, idFinformatToC := wrapString(format)
	defer idFinformatToC()
	retval := C.WrapImGuiSliderIntWithFormat(labelToC, vToC, vminToC, vmaxToC, formatToC)
	return bool(retval)
}

// ImGuiSliderIntVec2 ...
func ImGuiSliderIntVec2(label string, v *IVec2, vmin int32, vmax int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vminToC := C.int32_t(vmin)
	vmaxToC := C.int32_t(vmax)
	retval := C.WrapImGuiSliderIntVec2(labelToC, vToC, vminToC, vmaxToC)
	return bool(retval)
}

// ImGuiSliderIntVec2WithFormat ...
func ImGuiSliderIntVec2WithFormat(label string, v *IVec2, vmin int32, vmax int32, format string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vminToC := C.int32_t(vmin)
	vmaxToC := C.int32_t(vmax)
	formatToC, idFinformatToC := wrapString(format)
	defer idFinformatToC()
	retval := C.WrapImGuiSliderIntVec2WithFormat(labelToC, vToC, vminToC, vmaxToC, formatToC)
	return bool(retval)
}

// ImGuiSliderFloat Float slider widget.
func ImGuiSliderFloat(label string, v *float32, vmin float32, vmax float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.float)(unsafe.Pointer(v))
	vminToC := C.float(vmin)
	vmaxToC := C.float(vmax)
	retval := C.WrapImGuiSliderFloat(labelToC, vToC, vminToC, vmaxToC)
	return bool(retval)
}

// ImGuiSliderFloatWithFormat Float slider widget.
func ImGuiSliderFloatWithFormat(label string, v *float32, vmin float32, vmax float32, format string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := (*C.float)(unsafe.Pointer(v))
	vminToC := C.float(vmin)
	vmaxToC := C.float(vmax)
	formatToC, idFinformatToC := wrapString(format)
	defer idFinformatToC()
	retval := C.WrapImGuiSliderFloatWithFormat(labelToC, vToC, vminToC, vmaxToC, formatToC)
	return bool(retval)
}

// ImGuiSliderVec2 [harfang.Vec2] slider widget.
func ImGuiSliderVec2(label string, v *Vec2, vmin float32, vmax float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vminToC := C.float(vmin)
	vmaxToC := C.float(vmax)
	retval := C.WrapImGuiSliderVec2(labelToC, vToC, vminToC, vmaxToC)
	return bool(retval)
}

// ImGuiSliderVec2WithFormat [harfang.Vec2] slider widget.
func ImGuiSliderVec2WithFormat(label string, v *Vec2, vmin float32, vmax float32, format string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vminToC := C.float(vmin)
	vmaxToC := C.float(vmax)
	formatToC, idFinformatToC := wrapString(format)
	defer idFinformatToC()
	retval := C.WrapImGuiSliderVec2WithFormat(labelToC, vToC, vminToC, vmaxToC, formatToC)
	return bool(retval)
}

// ImGuiSliderVec3 [harfang.Vec3] slider widget.
func ImGuiSliderVec3(label string, v *Vec3, vmin float32, vmax float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vminToC := C.float(vmin)
	vmaxToC := C.float(vmax)
	retval := C.WrapImGuiSliderVec3(labelToC, vToC, vminToC, vmaxToC)
	return bool(retval)
}

// ImGuiSliderVec3WithFormat [harfang.Vec3] slider widget.
func ImGuiSliderVec3WithFormat(label string, v *Vec3, vmin float32, vmax float32, format string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vminToC := C.float(vmin)
	vmaxToC := C.float(vmax)
	formatToC, idFinformatToC := wrapString(format)
	defer idFinformatToC()
	retval := C.WrapImGuiSliderVec3WithFormat(labelToC, vToC, vminToC, vmaxToC, formatToC)
	return bool(retval)
}

// ImGuiSliderVec4 [harfang.Vec4] slider widget.
func ImGuiSliderVec4(label string, v *Vec4, vmin float32, vmax float32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vminToC := C.float(vmin)
	vmaxToC := C.float(vmax)
	retval := C.WrapImGuiSliderVec4(labelToC, vToC, vminToC, vmaxToC)
	return bool(retval)
}

// ImGuiSliderVec4WithFormat [harfang.Vec4] slider widget.
func ImGuiSliderVec4WithFormat(label string, v *Vec4, vmin float32, vmax float32, format string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	vToC := v.h
	vminToC := C.float(vmin)
	vmaxToC := C.float(vmax)
	formatToC, idFinformatToC := wrapString(format)
	defer idFinformatToC()
	retval := C.WrapImGuiSliderVec4WithFormat(labelToC, vToC, vminToC, vmaxToC, formatToC)
	return bool(retval)
}

// ImGuiTreeNode If returning `true` the node is open and the user is responsible for calling [harfang.ImGuiTreePop].
func ImGuiTreeNode(label string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	retval := C.WrapImGuiTreeNode(labelToC)
	return bool(retval)
}

// ImGuiTreeNodeEx See [harfang.ImGuiTreeNode].
func ImGuiTreeNodeEx(label string, flags ImGuiTreeNodeFlags) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiTreeNodeEx(labelToC, flagsToC)
	return bool(retval)
}

// ImGuiTreePush Already called by [harfang.ImGuiTreeNode], but you can call [harfang.ImGuiTreePush]/[harfang.ImGuiTreePop] yourself for layouting purpose.
func ImGuiTreePush(id string) {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	C.WrapImGuiTreePush(idToC)
}

// ImGuiTreePop Pop the current tree node.
func ImGuiTreePop() {
	C.WrapImGuiTreePop()
}

// ImGuiGetTreeNodeToLabelSpacing Return the horizontal distance preceding label when using [harfang.ImGuiTreeNode] or [harfang.ImGuiBullet].  The value `g.FontSize + style.FramePadding.x * 2` is returned for a regular unframed TreeNode.
func ImGuiGetTreeNodeToLabelSpacing() float32 {
	retval := C.WrapImGuiGetTreeNodeToLabelSpacing()
	return float32(retval)
}

// ImGuiSetNextItemOpen Set next item open state.
func ImGuiSetNextItemOpen(isopen bool) {
	isopenToC := C.bool(isopen)
	C.WrapImGuiSetNextItemOpen(isopenToC)
}

// ImGuiSetNextItemOpenWithCondition Set next item open state.
func ImGuiSetNextItemOpenWithCondition(isopen bool, condition ImGuiCond) {
	isopenToC := C.bool(isopen)
	conditionToC := C.int32_t(condition)
	C.WrapImGuiSetNextItemOpenWithCondition(isopenToC, conditionToC)
}

// ImGuiCollapsingHeader Draw a collapsing header, returns `False` if the header is collapsed so that you may skip drawing the header content.
func ImGuiCollapsingHeader(label string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	retval := C.WrapImGuiCollapsingHeader(labelToC)
	return bool(retval)
}

// ImGuiCollapsingHeaderWithFlags Draw a collapsing header, returns `False` if the header is collapsed so that you may skip drawing the header content.
func ImGuiCollapsingHeaderWithFlags(label string, flags ImGuiTreeNodeFlags) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiCollapsingHeaderWithFlags(labelToC, flagsToC)
	return bool(retval)
}

// ImGuiCollapsingHeaderWithPOpen Draw a collapsing header, returns `False` if the header is collapsed so that you may skip drawing the header content.
func ImGuiCollapsingHeaderWithPOpen(label string, popen *bool) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	popenToC := (*C.bool)(unsafe.Pointer(popen))
	retval := C.WrapImGuiCollapsingHeaderWithPOpen(labelToC, popenToC)
	return bool(retval)
}

// ImGuiCollapsingHeaderWithPOpenFlags Draw a collapsing header, returns `False` if the header is collapsed so that you may skip drawing the header content.
func ImGuiCollapsingHeaderWithPOpenFlags(label string, popen *bool, flags ImGuiTreeNodeFlags) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	popenToC := (*C.bool)(unsafe.Pointer(popen))
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiCollapsingHeaderWithPOpenFlags(labelToC, popenToC, flagsToC)
	return bool(retval)
}

// ImGuiSelectable Selectable item.  The following `width` values are possible:  * `= 0.0`: Use remaining width. * `> 0.0`: Specific width.  The following `height` values are possible:  * `= 0.0`: Use label height. * `> 0.0`: Specific height.
func ImGuiSelectable(label string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	retval := C.WrapImGuiSelectable(labelToC)
	return bool(retval)
}

// ImGuiSelectableWithSelected Selectable item.  The following `width` values are possible:  * `= 0.0`: Use remaining width. * `> 0.0`: Specific width.  The following `height` values are possible:  * `= 0.0`: Use label height. * `> 0.0`: Specific height.
func ImGuiSelectableWithSelected(label string, selected bool) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	selectedToC := C.bool(selected)
	retval := C.WrapImGuiSelectableWithSelected(labelToC, selectedToC)
	return bool(retval)
}

// ImGuiSelectableWithSelectedFlags Selectable item.  The following `width` values are possible:  * `= 0.0`: Use remaining width. * `> 0.0`: Specific width.  The following `height` values are possible:  * `= 0.0`: Use label height. * `> 0.0`: Specific height.
func ImGuiSelectableWithSelectedFlags(label string, selected bool, flags ImGuiSelectableFlags) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	selectedToC := C.bool(selected)
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiSelectableWithSelectedFlags(labelToC, selectedToC, flagsToC)
	return bool(retval)
}

// ImGuiSelectableWithSelectedFlagsSize Selectable item.  The following `width` values are possible:  * `= 0.0`: Use remaining width. * `> 0.0`: Specific width.  The following `height` values are possible:  * `= 0.0`: Use label height. * `> 0.0`: Specific height.
func ImGuiSelectableWithSelectedFlagsSize(label string, selected bool, flags ImGuiSelectableFlags, size *Vec2) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	selectedToC := C.bool(selected)
	flagsToC := C.int32_t(flags)
	sizeToC := size.h
	retval := C.WrapImGuiSelectableWithSelectedFlagsSize(labelToC, selectedToC, flagsToC, sizeToC)
	return bool(retval)
}

// ImGuiListBox List widget.
func ImGuiListBox(label string, currentitem *int32, items *StringList) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	currentitemToC := (*C.int32_t)(unsafe.Pointer(currentitem))
	itemsToC := items.h
	retval := C.WrapImGuiListBox(labelToC, currentitemToC, itemsToC)
	return bool(retval)
}

// ImGuiListBoxWithHeightInItems List widget.
func ImGuiListBoxWithHeightInItems(label string, currentitem *int32, items *StringList, heightinitems int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	currentitemToC := (*C.int32_t)(unsafe.Pointer(currentitem))
	itemsToC := items.h
	heightinitemsToC := C.int32_t(heightinitems)
	retval := C.WrapImGuiListBoxWithHeightInItems(labelToC, currentitemToC, itemsToC, heightinitemsToC)
	return bool(retval)
}

// ImGuiListBoxWithSliceOfItems List widget.
func ImGuiListBoxWithSliceOfItems(label string, currentitem *int32, SliceOfitems GoSliceOfstring) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	currentitemToC := (*C.int32_t)(unsafe.Pointer(currentitem))
	var SliceOfitemsSpecialString []*C.char
	for _, s := range SliceOfitems {
		SliceOfitemsSpecialString = append(SliceOfitemsSpecialString, C.CString(s))
	}
	SliceOfitemsSpecialStringToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfitemsSpecialString))
	SliceOfitemsSpecialStringToCSize := C.size_t(SliceOfitemsSpecialStringToC.Len)
	SliceOfitemsSpecialStringToCBuf := (**C.char)(unsafe.Pointer(SliceOfitemsSpecialStringToC.Data))
	retval := C.WrapImGuiListBoxWithSliceOfItems(labelToC, currentitemToC, SliceOfitemsSpecialStringToCSize, SliceOfitemsSpecialStringToCBuf)
	return bool(retval)
}

// ImGuiListBoxWithSliceOfItemsHeightInItems List widget.
func ImGuiListBoxWithSliceOfItemsHeightInItems(label string, currentitem *int32, SliceOfitems GoSliceOfstring, heightinitems int32) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	currentitemToC := (*C.int32_t)(unsafe.Pointer(currentitem))
	var SliceOfitemsSpecialString []*C.char
	for _, s := range SliceOfitems {
		SliceOfitemsSpecialString = append(SliceOfitemsSpecialString, C.CString(s))
	}
	SliceOfitemsSpecialStringToC := (*reflect.SliceHeader)(unsafe.Pointer(&SliceOfitemsSpecialString))
	SliceOfitemsSpecialStringToCSize := C.size_t(SliceOfitemsSpecialStringToC.Len)
	SliceOfitemsSpecialStringToCBuf := (**C.char)(unsafe.Pointer(SliceOfitemsSpecialStringToC.Data))
	heightinitemsToC := C.int32_t(heightinitems)
	retval := C.WrapImGuiListBoxWithSliceOfItemsHeightInItems(labelToC, currentitemToC, SliceOfitemsSpecialStringToCSize, SliceOfitemsSpecialStringToCBuf, heightinitemsToC)
	return bool(retval)
}

// ImGuiSetTooltip Set tooltip under mouse-cursor, typically used with [harfang.ImGuiIsItemHovered]/[harfang.ImGuiIsAnyItemHovered]. Last call wins.
func ImGuiSetTooltip(text string) {
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	C.WrapImGuiSetTooltip(textToC)
}

// ImGuiBeginTooltip Used to create full-featured tooltip windows that aren't just text.
func ImGuiBeginTooltip() {
	C.WrapImGuiBeginTooltip()
}

// ImGuiEndTooltip End the current tooltip window.  See [harfang.ImGuiBeginTooltip].
func ImGuiEndTooltip() {
	C.WrapImGuiEndTooltip()
}

// ImGuiBeginMainMenuBar Create and append to a full screen menu-bar.  Note: Only call [harfang.ImGuiEndMainMenuBar] if this returns `true`.
func ImGuiBeginMainMenuBar() bool {
	retval := C.WrapImGuiBeginMainMenuBar()
	return bool(retval)
}

// ImGuiEndMainMenuBar End the main menu bar.  See [harfang.ImGuiBeginMainMenuBar].
func ImGuiEndMainMenuBar() {
	C.WrapImGuiEndMainMenuBar()
}

// ImGuiBeginMenuBar Start append to the menu-bar of the current window (requires the `WindowFlags_MenuBar` flag).  Note: Only call [harfang.ImGuiEndMenuBar] if this returns `true`.
func ImGuiBeginMenuBar() bool {
	retval := C.WrapImGuiBeginMenuBar()
	return bool(retval)
}

// ImGuiEndMenuBar End the current menu bar.
func ImGuiEndMenuBar() {
	C.WrapImGuiEndMenuBar()
}

// ImGuiBeginMenu Create a sub-menu entry.  Note: Only call [harfang.ImGuiEndMenu] if this returns `true`.
func ImGuiBeginMenu(label string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	retval := C.WrapImGuiBeginMenu(labelToC)
	return bool(retval)
}

// ImGuiBeginMenuWithEnabled Create a sub-menu entry.  Note: Only call [harfang.ImGuiEndMenu] if this returns `true`.
func ImGuiBeginMenuWithEnabled(label string, enabled bool) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	enabledToC := C.bool(enabled)
	retval := C.WrapImGuiBeginMenuWithEnabled(labelToC, enabledToC)
	return bool(retval)
}

// ImGuiEndMenu End the current sub-menu entry.
func ImGuiEndMenu() {
	C.WrapImGuiEndMenu()
}

// ImGuiMenuItem Return `true` when activated. Shortcuts are displayed for convenience but not processed at the moment.
func ImGuiMenuItem(label string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	retval := C.WrapImGuiMenuItem(labelToC)
	return bool(retval)
}

// ImGuiMenuItemWithShortcut Return `true` when activated. Shortcuts are displayed for convenience but not processed at the moment.
func ImGuiMenuItemWithShortcut(label string, shortcut string) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	shortcutToC, idFinshortcutToC := wrapString(shortcut)
	defer idFinshortcutToC()
	retval := C.WrapImGuiMenuItemWithShortcut(labelToC, shortcutToC)
	return bool(retval)
}

// ImGuiMenuItemWithShortcutSelected Return `true` when activated. Shortcuts are displayed for convenience but not processed at the moment.
func ImGuiMenuItemWithShortcutSelected(label string, shortcut string, selected bool) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	shortcutToC, idFinshortcutToC := wrapString(shortcut)
	defer idFinshortcutToC()
	selectedToC := C.bool(selected)
	retval := C.WrapImGuiMenuItemWithShortcutSelected(labelToC, shortcutToC, selectedToC)
	return bool(retval)
}

// ImGuiMenuItemWithShortcutSelectedEnabled Return `true` when activated. Shortcuts are displayed for convenience but not processed at the moment.
func ImGuiMenuItemWithShortcutSelectedEnabled(label string, shortcut string, selected bool, enabled bool) bool {
	labelToC, idFinlabelToC := wrapString(label)
	defer idFinlabelToC()
	shortcutToC, idFinshortcutToC := wrapString(shortcut)
	defer idFinshortcutToC()
	selectedToC := C.bool(selected)
	enabledToC := C.bool(enabled)
	retval := C.WrapImGuiMenuItemWithShortcutSelectedEnabled(labelToC, shortcutToC, selectedToC, enabledToC)
	return bool(retval)
}

// ImGuiOpenPopup Mark a named popup as open.  Popup windows are closed when the user:  * Clicks outside of their client rect, * Activates a pressable item, * [harfang.ImGuiCloseCurrentPopup] is called within a [harfang.ImGuiBeginPopup]/[harfang.ImGuiEndPopup] block.  Popup identifiers are relative to the current ID stack so [harfang.ImGuiOpenPopup] and [harfang.ImGuiBeginPopup] need to be at the same level of the ID stack.
func ImGuiOpenPopup(id string) {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	C.WrapImGuiOpenPopup(idToC)
}

// ImGuiBeginPopup Return `true` if popup is opened and starts outputting to it.  Note: Only call [harfang.ImGuiEndPopup] if this returns `true`.
func ImGuiBeginPopup(id string) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	retval := C.WrapImGuiBeginPopup(idToC)
	return bool(retval)
}

// ImGuiBeginPopupModal Begin an ImGui modal dialog.
func ImGuiBeginPopupModal(name string) bool {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapImGuiBeginPopupModal(nameToC)
	return bool(retval)
}

// ImGuiBeginPopupModalWithOpen Begin an ImGui modal dialog.
func ImGuiBeginPopupModalWithOpen(name string, open *bool) bool {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	openToC := (*C.bool)(unsafe.Pointer(open))
	retval := C.WrapImGuiBeginPopupModalWithOpen(nameToC, openToC)
	return bool(retval)
}

// ImGuiBeginPopupModalWithOpenFlags Begin an ImGui modal dialog.
func ImGuiBeginPopupModalWithOpenFlags(name string, open *bool, flags ImGuiWindowFlags) bool {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	openToC := (*C.bool)(unsafe.Pointer(open))
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiBeginPopupModalWithOpenFlags(nameToC, openToC, flagsToC)
	return bool(retval)
}

// ImGuiBeginPopupContextItem ImGui helper to open and begin popup when clicked on last item.
func ImGuiBeginPopupContextItem(id string) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	retval := C.WrapImGuiBeginPopupContextItem(idToC)
	return bool(retval)
}

// ImGuiBeginPopupContextItemWithMouseButton ImGui helper to open and begin popup when clicked on last item.
func ImGuiBeginPopupContextItemWithMouseButton(id string, mousebutton int32) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	mousebuttonToC := C.int32_t(mousebutton)
	retval := C.WrapImGuiBeginPopupContextItemWithMouseButton(idToC, mousebuttonToC)
	return bool(retval)
}

// ImGuiBeginPopupContextWindow ImGui helper to open and begin popup when clicked on current window.
func ImGuiBeginPopupContextWindow() bool {
	retval := C.WrapImGuiBeginPopupContextWindow()
	return bool(retval)
}

// ImGuiBeginPopupContextWindowWithId ImGui helper to open and begin popup when clicked on current window.
func ImGuiBeginPopupContextWindowWithId(id string) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	retval := C.WrapImGuiBeginPopupContextWindowWithId(idToC)
	return bool(retval)
}

// ImGuiBeginPopupContextWindowWithIdFlags ImGui helper to open and begin popup when clicked on current window.
func ImGuiBeginPopupContextWindowWithIdFlags(id string, flags ImGuiPopupFlags) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiBeginPopupContextWindowWithIdFlags(idToC, flagsToC)
	return bool(retval)
}

// ImGuiBeginPopupContextVoid ImGui helper to open and begin popup when clicked in void (where there are no ImGui windows)
func ImGuiBeginPopupContextVoid() bool {
	retval := C.WrapImGuiBeginPopupContextVoid()
	return bool(retval)
}

// ImGuiBeginPopupContextVoidWithId ImGui helper to open and begin popup when clicked in void (where there are no ImGui windows)
func ImGuiBeginPopupContextVoidWithId(id string) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	retval := C.WrapImGuiBeginPopupContextVoidWithId(idToC)
	return bool(retval)
}

// ImGuiBeginPopupContextVoidWithIdMouseButton ImGui helper to open and begin popup when clicked in void (where there are no ImGui windows)
func ImGuiBeginPopupContextVoidWithIdMouseButton(id string, mousebutton int32) bool {
	idToC, idFinidToC := wrapString(id)
	defer idFinidToC()
	mousebuttonToC := C.int32_t(mousebutton)
	retval := C.WrapImGuiBeginPopupContextVoidWithIdMouseButton(idToC, mousebuttonToC)
	return bool(retval)
}

// ImGuiEndPopup End the current popup.
func ImGuiEndPopup() {
	C.WrapImGuiEndPopup()
}

// ImGuiCloseCurrentPopup Close the popup we have begin-ed into. Clicking on a menu item or selectable automatically closes the current popup.
func ImGuiCloseCurrentPopup() {
	C.WrapImGuiCloseCurrentPopup()
}

// ImGuiPushClipRect Push a new clip rectangle onto the clipping stack.
func ImGuiPushClipRect(cliprectmin *Vec2, cliprectmax *Vec2, intersectwithcurrentcliprect bool) {
	cliprectminToC := cliprectmin.h
	cliprectmaxToC := cliprectmax.h
	intersectwithcurrentcliprectToC := C.bool(intersectwithcurrentcliprect)
	C.WrapImGuiPushClipRect(cliprectminToC, cliprectmaxToC, intersectwithcurrentcliprectToC)
}

// ImGuiPopClipRect Undo the last call to [harfang.ImGuiPushClipRect].
func ImGuiPopClipRect() {
	C.WrapImGuiPopClipRect()
}

// ImGuiIsItemHovered Was the last item hovered by mouse.
func ImGuiIsItemHovered() bool {
	retval := C.WrapImGuiIsItemHovered()
	return bool(retval)
}

// ImGuiIsItemHoveredWithFlags Was the last item hovered by mouse.
func ImGuiIsItemHoveredWithFlags(flags ImGuiHoveredFlags) bool {
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiIsItemHoveredWithFlags(flagsToC)
	return bool(retval)
}

// ImGuiIsItemActive Was the last item active.  e.g. button being held, text field being edited - items that do not interact will always return `false`.
func ImGuiIsItemActive() bool {
	retval := C.WrapImGuiIsItemActive()
	return bool(retval)
}

// ImGuiIsItemClicked Was the last item clicked.
func ImGuiIsItemClicked() bool {
	retval := C.WrapImGuiIsItemClicked()
	return bool(retval)
}

// ImGuiIsItemClickedWithMouseButton Was the last item clicked.
func ImGuiIsItemClickedWithMouseButton(mousebutton int32) bool {
	mousebuttonToC := C.int32_t(mousebutton)
	retval := C.WrapImGuiIsItemClickedWithMouseButton(mousebuttonToC)
	return bool(retval)
}

// ImGuiIsItemVisible Was the last item visible and not out of sight due to clipping/scrolling.
func ImGuiIsItemVisible() bool {
	retval := C.WrapImGuiIsItemVisible()
	return bool(retval)
}

// ImGuiIsAnyItemHovered Return `true` if any item is hovered by the mouse cursor, `false` otherwise.
func ImGuiIsAnyItemHovered() bool {
	retval := C.WrapImGuiIsAnyItemHovered()
	return bool(retval)
}

// ImGuiIsAnyItemActive Return `true` if any item is active, `false` otherwise.
func ImGuiIsAnyItemActive() bool {
	retval := C.WrapImGuiIsAnyItemActive()
	return bool(retval)
}

// ImGuiGetItemRectMin Get bounding rect minimum of last item in screen space.
func ImGuiGetItemRectMin() *Vec2 {
	retval := C.WrapImGuiGetItemRectMin()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetItemRectMax Get bounding rect maximum of last item in screen space.
func ImGuiGetItemRectMax() *Vec2 {
	retval := C.WrapImGuiGetItemRectMax()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetItemRectSize Get bounding rect size of last item in screen space.
func ImGuiGetItemRectSize() *Vec2 {
	retval := C.WrapImGuiGetItemRectSize()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiSetItemAllowOverlap Allow the last item to be overlapped by a subsequent item. Sometimes useful with invisible buttons, selectables, etc... to catch unused areas.
func ImGuiSetItemAllowOverlap() {
	C.WrapImGuiSetItemAllowOverlap()
}

// ImGuiSetItemDefaultFocus Make the last item the default focused item of a window.
func ImGuiSetItemDefaultFocus() {
	C.WrapImGuiSetItemDefaultFocus()
}

// ImGuiIsWindowHovered Is the current window hovered and hoverable (not blocked by a popup), differentiates child windows from each others.
func ImGuiIsWindowHovered() bool {
	retval := C.WrapImGuiIsWindowHovered()
	return bool(retval)
}

// ImGuiIsWindowHoveredWithFlags Is the current window hovered and hoverable (not blocked by a popup), differentiates child windows from each others.
func ImGuiIsWindowHoveredWithFlags(flags ImGuiHoveredFlags) bool {
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiIsWindowHoveredWithFlags(flagsToC)
	return bool(retval)
}

// ImGuiIsWindowFocused Is the current window focused.
func ImGuiIsWindowFocused() bool {
	retval := C.WrapImGuiIsWindowFocused()
	return bool(retval)
}

// ImGuiIsWindowFocusedWithFlags Is the current window focused.
func ImGuiIsWindowFocusedWithFlags(flags ImGuiFocusedFlags) bool {
	flagsToC := C.int32_t(flags)
	retval := C.WrapImGuiIsWindowFocusedWithFlags(flagsToC)
	return bool(retval)
}

// ImGuiIsRectVisible Test if a rectangle of the specified size starting from cursor position is visible/not clipped. Or test if a rectangle in screen space is visible/not clipped.
func ImGuiIsRectVisible(size *Vec2) bool {
	sizeToC := size.h
	retval := C.WrapImGuiIsRectVisible(sizeToC)
	return bool(retval)
}

// ImGuiIsRectVisibleWithRectMinRectMax Test if a rectangle of the specified size starting from cursor position is visible/not clipped. Or test if a rectangle in screen space is visible/not clipped.
func ImGuiIsRectVisibleWithRectMinRectMax(rectmin *Vec2, rectmax *Vec2) bool {
	rectminToC := rectmin.h
	rectmaxToC := rectmax.h
	retval := C.WrapImGuiIsRectVisibleWithRectMinRectMax(rectminToC, rectmaxToC)
	return bool(retval)
}

// ImGuiGetTime Return the current ImGui time in seconds.
func ImGuiGetTime() float32 {
	retval := C.WrapImGuiGetTime()
	return float32(retval)
}

// ImGuiGetFrameCount Return the ImGui frame counter.  See [harfang.ImGuiBeginFrame] and [harfang.ImGuiEndFrame].
func ImGuiGetFrameCount() int32 {
	retval := C.WrapImGuiGetFrameCount()
	return int32(retval)
}

// ImGuiCalcTextSize Compute the bounding rectangle for the provided text.
func ImGuiCalcTextSize(text string) *Vec2 {
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	retval := C.WrapImGuiCalcTextSize(textToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiCalcTextSizeWithHideTextAfterDoubleDash Compute the bounding rectangle for the provided text.
func ImGuiCalcTextSizeWithHideTextAfterDoubleDash(text string, hidetextafterdoubledash bool) *Vec2 {
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	hidetextafterdoubledashToC := C.bool(hidetextafterdoubledash)
	retval := C.WrapImGuiCalcTextSizeWithHideTextAfterDoubleDash(textToC, hidetextafterdoubledashToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiCalcTextSizeWithHideTextAfterDoubleDashWrapWidth Compute the bounding rectangle for the provided text.
func ImGuiCalcTextSizeWithHideTextAfterDoubleDashWrapWidth(text string, hidetextafterdoubledash bool, wrapwidth float32) *Vec2 {
	textToC, idFintextToC := wrapString(text)
	defer idFintextToC()
	hidetextafterdoubledashToC := C.bool(hidetextafterdoubledash)
	wrapwidthToC := C.float(wrapwidth)
	retval := C.WrapImGuiCalcTextSizeWithHideTextAfterDoubleDashWrapWidth(textToC, hidetextafterdoubledashToC, wrapwidthToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiIsKeyDown Was the specified key down during the last frame?
func ImGuiIsKeyDown(keyindex int32) bool {
	keyindexToC := C.int32_t(keyindex)
	retval := C.WrapImGuiIsKeyDown(keyindexToC)
	return bool(retval)
}

// ImGuiIsKeyPressed Was the specified key pressed? A key press implies that the key was down and is currently released.
func ImGuiIsKeyPressed(keyindex int32) bool {
	keyindexToC := C.int32_t(keyindex)
	retval := C.WrapImGuiIsKeyPressed(keyindexToC)
	return bool(retval)
}

// ImGuiIsKeyPressedWithRepeat Was the specified key pressed? A key press implies that the key was down and is currently released.
func ImGuiIsKeyPressedWithRepeat(keyindex int32, repeat bool) bool {
	keyindexToC := C.int32_t(keyindex)
	repeatToC := C.bool(repeat)
	retval := C.WrapImGuiIsKeyPressedWithRepeat(keyindexToC, repeatToC)
	return bool(retval)
}

// ImGuiIsKeyReleased Was the specified key released during the last frame?
func ImGuiIsKeyReleased(keyindex int32) bool {
	keyindexToC := C.int32_t(keyindex)
	retval := C.WrapImGuiIsKeyReleased(keyindexToC)
	return bool(retval)
}

// ImGuiIsMouseDown Was the specified mouse button down during the last frame?
func ImGuiIsMouseDown(button int32) bool {
	buttonToC := C.int32_t(button)
	retval := C.WrapImGuiIsMouseDown(buttonToC)
	return bool(retval)
}

// ImGuiIsMouseClicked Was the specified mouse button clicked during the last frame? A mouse click implies that the button pressed earlier and released during the last frame.
func ImGuiIsMouseClicked(button int32) bool {
	buttonToC := C.int32_t(button)
	retval := C.WrapImGuiIsMouseClicked(buttonToC)
	return bool(retval)
}

// ImGuiIsMouseClickedWithRepeat Was the specified mouse button clicked during the last frame? A mouse click implies that the button pressed earlier and released during the last frame.
func ImGuiIsMouseClickedWithRepeat(button int32, repeat bool) bool {
	buttonToC := C.int32_t(button)
	repeatToC := C.bool(repeat)
	retval := C.WrapImGuiIsMouseClickedWithRepeat(buttonToC, repeatToC)
	return bool(retval)
}

// ImGuiIsMouseDoubleClicked Was the specified mouse button double-clicked during the last frame? A double-click implies two rapid successive clicks of the same button with the mouse cursor staying in the same position.
func ImGuiIsMouseDoubleClicked(button int32) bool {
	buttonToC := C.int32_t(button)
	retval := C.WrapImGuiIsMouseDoubleClicked(buttonToC)
	return bool(retval)
}

// ImGuiIsMouseReleased Was the specified mouse button released during the last frame?
func ImGuiIsMouseReleased(button int32) bool {
	buttonToC := C.int32_t(button)
	retval := C.WrapImGuiIsMouseReleased(buttonToC)
	return bool(retval)
}

// ImGuiIsMouseHoveringRect Test whether the mouse cursor is hovering the specified rectangle.
func ImGuiIsMouseHoveringRect(rectmin *Vec2, rectmax *Vec2) bool {
	rectminToC := rectmin.h
	rectmaxToC := rectmax.h
	retval := C.WrapImGuiIsMouseHoveringRect(rectminToC, rectmaxToC)
	return bool(retval)
}

// ImGuiIsMouseHoveringRectWithClip Test whether the mouse cursor is hovering the specified rectangle.
func ImGuiIsMouseHoveringRectWithClip(rectmin *Vec2, rectmax *Vec2, clip bool) bool {
	rectminToC := rectmin.h
	rectmaxToC := rectmax.h
	clipToC := C.bool(clip)
	retval := C.WrapImGuiIsMouseHoveringRectWithClip(rectminToC, rectmaxToC, clipToC)
	return bool(retval)
}

// ImGuiIsMouseDragging Is mouse dragging?
func ImGuiIsMouseDragging(button ImGuiMouseButton) bool {
	buttonToC := C.int32_t(button)
	retval := C.WrapImGuiIsMouseDragging(buttonToC)
	return bool(retval)
}

// ImGuiIsMouseDraggingWithLockThreshold Is mouse dragging?
func ImGuiIsMouseDraggingWithLockThreshold(button ImGuiMouseButton, lockthreshold float32) bool {
	buttonToC := C.int32_t(button)
	lockthresholdToC := C.float(lockthreshold)
	retval := C.WrapImGuiIsMouseDraggingWithLockThreshold(buttonToC, lockthresholdToC)
	return bool(retval)
}

// ImGuiGetMousePos Return the mouse cursor coordinates in screen space.
func ImGuiGetMousePos() *Vec2 {
	retval := C.WrapImGuiGetMousePos()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetMousePosOnOpeningCurrentPopup Retrieve a backup of the mouse position at the time of opening the current popup.  See [harfang.ImGuiBeginPopup].
func ImGuiGetMousePosOnOpeningCurrentPopup() *Vec2 {
	retval := C.WrapImGuiGetMousePosOnOpeningCurrentPopup()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetMouseDragDelta Return the distance covered by the mouse cursor since the last button press.
func ImGuiGetMouseDragDelta() *Vec2 {
	retval := C.WrapImGuiGetMouseDragDelta()
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetMouseDragDeltaWithButton Return the distance covered by the mouse cursor since the last button press.
func ImGuiGetMouseDragDeltaWithButton(button ImGuiMouseButton) *Vec2 {
	buttonToC := C.int32_t(button)
	retval := C.WrapImGuiGetMouseDragDeltaWithButton(buttonToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiGetMouseDragDeltaWithButtonLockThreshold Return the distance covered by the mouse cursor since the last button press.
func ImGuiGetMouseDragDeltaWithButtonLockThreshold(button ImGuiMouseButton, lockthreshold float32) *Vec2 {
	buttonToC := C.int32_t(button)
	lockthresholdToC := C.float(lockthreshold)
	retval := C.WrapImGuiGetMouseDragDeltaWithButtonLockThreshold(buttonToC, lockthresholdToC)
	retvalGO := &Vec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vec2) {
		C.WrapVec2Free(cleanval.h)
	})
	return retvalGO
}

// ImGuiResetMouseDragDelta ...
func ImGuiResetMouseDragDelta() {
	C.WrapImGuiResetMouseDragDelta()
}

// ImGuiResetMouseDragDeltaWithButton ...
func ImGuiResetMouseDragDeltaWithButton(button ImGuiMouseButton) {
	buttonToC := C.int32_t(button)
	C.WrapImGuiResetMouseDragDeltaWithButton(buttonToC)
}

// ImGuiCaptureKeyboardFromApp Force capture keyboard when your widget is being hovered.
func ImGuiCaptureKeyboardFromApp(capture bool) {
	captureToC := C.bool(capture)
	C.WrapImGuiCaptureKeyboardFromApp(captureToC)
}

// ImGuiCaptureMouseFromApp Force capture mouse when your widget is being hovered.
func ImGuiCaptureMouseFromApp(capture bool) {
	captureToC := C.bool(capture)
	C.WrapImGuiCaptureMouseFromApp(captureToC)
}

// ImGuiWantCaptureMouse ImGui wants mouse capture. Use this function to determine when to pause mouse processing from other parts of your program.
func ImGuiWantCaptureMouse() bool {
	retval := C.WrapImGuiWantCaptureMouse()
	return bool(retval)
}

// ImGuiMouseDrawCursor Enable/disable the ImGui software mouse cursor.
func ImGuiMouseDrawCursor(drawcursor bool) {
	drawcursorToC := C.bool(drawcursor)
	C.WrapImGuiMouseDrawCursor(drawcursorToC)
}

// ImGuiInit Initialize the global ImGui context. This function must be called once before any other ImGui function using the global context.  See [harfang.ImGuiInitContext].
func ImGuiInit(fontsize float32, imguiprogram *ProgramHandle, imguiimageprogram *ProgramHandle) {
	fontsizeToC := C.float(fontsize)
	imguiprogramToC := imguiprogram.h
	imguiimageprogramToC := imguiimageprogram.h
	C.WrapImGuiInit(fontsizeToC, imguiprogramToC, imguiimageprogramToC)
}

// ImGuiInitContext Initialize an ImGui context. This function must be called once before any other ImGui function using the context.  See [harfang.ImGuiInit].
func ImGuiInitContext(fontsize float32, imguiprogram *ProgramHandle, imguiimageprogram *ProgramHandle) *DearImguiContext {
	fontsizeToC := C.float(fontsize)
	imguiprogramToC := imguiprogram.h
	imguiimageprogramToC := imguiimageprogram.h
	retval := C.WrapImGuiInitContext(fontsizeToC, imguiprogramToC, imguiimageprogramToC)
	var retvalGO *DearImguiContext
	if retval != nil {
		retvalGO = &DearImguiContext{h: retval}
		runtime.SetFinalizer(retvalGO, func(cleanval *DearImguiContext) {
			C.WrapDearImguiContextFree(cleanval.h)
		})
	}
	return retvalGO
}

// ImGuiShutdown Shutdown the global ImGui context.
func ImGuiShutdown() {
	C.WrapImGuiShutdown()
}

// ImGuiBeginFrame Begin an ImGui frame. This function must be called once per frame before any other ImGui call.  When using multiple contexts, it must be called for each context you intend to use during the current frame.  See [harfang.ImGuiEndFrame].
func ImGuiBeginFrame(width int32, height int32, dtclock int64, mouse *MouseState, keyboard *KeyboardState) {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	dtclockToC := C.int64_t(dtclock)
	mouseToC := mouse.h
	keyboardToC := keyboard.h
	C.WrapImGuiBeginFrame(widthToC, heightToC, dtclockToC, mouseToC, keyboardToC)
}

// ImGuiBeginFrameWithCtxWidthHeightDtClockMouseKeyboard Begin an ImGui frame. This function must be called once per frame before any other ImGui call.  When using multiple contexts, it must be called for each context you intend to use during the current frame.  See [harfang.ImGuiEndFrame].
func ImGuiBeginFrameWithCtxWidthHeightDtClockMouseKeyboard(ctx *DearImguiContext, width int32, height int32, dtclock int64, mouse *MouseState, keyboard *KeyboardState) {
	ctxToC := ctx.h
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	dtclockToC := C.int64_t(dtclock)
	mouseToC := mouse.h
	keyboardToC := keyboard.h
	C.WrapImGuiBeginFrameWithCtxWidthHeightDtClockMouseKeyboard(ctxToC, widthToC, heightToC, dtclockToC, mouseToC, keyboardToC)
}

// ImGuiEndFrameWithCtx End the current ImGui frame.  All ImGui rendering is sent to the specified view. If no view is specified, view 255 is used.  See [harfang.man.Views].
func ImGuiEndFrameWithCtx(ctx *DearImguiContext) {
	ctxToC := ctx.h
	C.WrapImGuiEndFrameWithCtx(ctxToC)
}

// ImGuiEndFrameWithCtxViewId End the current ImGui frame.  All ImGui rendering is sent to the specified view. If no view is specified, view 255 is used.  See [harfang.man.Views].
func ImGuiEndFrameWithCtxViewId(ctx *DearImguiContext, viewid uint16) {
	ctxToC := ctx.h
	viewidToC := C.ushort(viewid)
	C.WrapImGuiEndFrameWithCtxViewId(ctxToC, viewidToC)
}

// ImGuiEndFrame End the current ImGui frame.  All ImGui rendering is sent to the specified view. If no view is specified, view 255 is used.  See [harfang.man.Views].
func ImGuiEndFrame() {
	C.WrapImGuiEndFrame()
}

// ImGuiEndFrameWithViewId End the current ImGui frame.  All ImGui rendering is sent to the specified view. If no view is specified, view 255 is used.  See [harfang.man.Views].
func ImGuiEndFrameWithViewId(viewid uint16) {
	viewidToC := C.ushort(viewid)
	C.WrapImGuiEndFrameWithViewId(viewidToC)
}

// ImGuiClearInputBuffer Force a reset of the ImGui input buffer.
func ImGuiClearInputBuffer() {
	C.WrapImGuiClearInputBuffer()
}

// OpenFolderDialog Open a native OpenFolder dialog.
func OpenFolderDialog(title string, foldername *string) bool {
	titleToC, idFintitleToC := wrapString(title)
	defer idFintitleToC()
	foldernameToC1 := C.CString(*foldername)
	foldernameToC := &foldernameToC1
	retval := C.WrapOpenFolderDialog(titleToC, foldernameToC)
	return bool(retval)
}

// OpenFolderDialogWithInitialDir Open a native OpenFolder dialog.
func OpenFolderDialogWithInitialDir(title string, foldername *string, initialdir string) bool {
	titleToC, idFintitleToC := wrapString(title)
	defer idFintitleToC()
	foldernameToC1 := C.CString(*foldername)
	foldernameToC := &foldernameToC1
	initialdirToC, idFininitialdirToC := wrapString(initialdir)
	defer idFininitialdirToC()
	retval := C.WrapOpenFolderDialogWithInitialDir(titleToC, foldernameToC, initialdirToC)
	return bool(retval)
}

// OpenFileDialog Open a native OpenFile dialog.
func OpenFileDialog(title string, filters *FileFilterList, file *string) bool {
	titleToC, idFintitleToC := wrapString(title)
	defer idFintitleToC()
	filtersToC := filters.h
	fileToC1 := C.CString(*file)
	fileToC := &fileToC1
	retval := C.WrapOpenFileDialog(titleToC, filtersToC, fileToC)
	return bool(retval)
}

// OpenFileDialogWithInitialDir Open a native OpenFile dialog.
func OpenFileDialogWithInitialDir(title string, filters *FileFilterList, file *string, initialdir string) bool {
	titleToC, idFintitleToC := wrapString(title)
	defer idFintitleToC()
	filtersToC := filters.h
	fileToC1 := C.CString(*file)
	fileToC := &fileToC1
	initialdirToC, idFininitialdirToC := wrapString(initialdir)
	defer idFininitialdirToC()
	retval := C.WrapOpenFileDialogWithInitialDir(titleToC, filtersToC, fileToC, initialdirToC)
	return bool(retval)
}

// SaveFileDialog Open a native SaveFile dialog.
func SaveFileDialog(title string, filters *FileFilterList, file *string) bool {
	titleToC, idFintitleToC := wrapString(title)
	defer idFintitleToC()
	filtersToC := filters.h
	fileToC1 := C.CString(*file)
	fileToC := &fileToC1
	retval := C.WrapSaveFileDialog(titleToC, filtersToC, fileToC)
	return bool(retval)
}

// SaveFileDialogWithInitialDir Open a native SaveFile dialog.
func SaveFileDialogWithInitialDir(title string, filters *FileFilterList, file *string, initialdir string) bool {
	titleToC, idFintitleToC := wrapString(title)
	defer idFintitleToC()
	filtersToC := filters.h
	fileToC1 := C.CString(*file)
	fileToC := &fileToC1
	initialdirToC, idFininitialdirToC := wrapString(initialdir)
	defer idFininitialdirToC()
	retval := C.WrapSaveFileDialogWithInitialDir(titleToC, filtersToC, fileToC, initialdirToC)
	return bool(retval)
}

// FpsControllerWithKeyUpKeyDownKeyLeftKeyRightBtnDxDyPosRotSpeedDtT Implement a first-person-shooter like controller.  The input position and rotation parameters are returned modified according to the state of the control keys.  This function is usually used by passing the current camera position and rotation then updating the camera transformation with the returned values.
func FpsControllerWithKeyUpKeyDownKeyLeftKeyRightBtnDxDyPosRotSpeedDtT(keyup bool, keydown bool, keyleft bool, keyright bool, btn bool, dx float32, dy float32, pos *Vec3, rot *Vec3, speed float32, dtt int64) {
	keyupToC := C.bool(keyup)
	keydownToC := C.bool(keydown)
	keyleftToC := C.bool(keyleft)
	keyrightToC := C.bool(keyright)
	btnToC := C.bool(btn)
	dxToC := C.float(dx)
	dyToC := C.float(dy)
	posToC := pos.h
	rotToC := rot.h
	speedToC := C.float(speed)
	dttToC := C.int64_t(dtt)
	C.WrapFpsControllerWithKeyUpKeyDownKeyLeftKeyRightBtnDxDyPosRotSpeedDtT(keyupToC, keydownToC, keyleftToC, keyrightToC, btnToC, dxToC, dyToC, posToC, rotToC, speedToC, dttToC)
}

// FpsController Implement a first-person-shooter like controller.  The input position and rotation parameters are returned modified according to the state of the control keys.  This function is usually used by passing the current camera position and rotation then updating the camera transformation with the returned values.
func FpsController(keyboard *Keyboard, mouse *Mouse, pos *Vec3, rot *Vec3, speed float32, dt int64) {
	keyboardToC := keyboard.h
	mouseToC := mouse.h
	posToC := pos.h
	rotToC := rot.h
	speedToC := C.float(speed)
	dtToC := C.int64_t(dt)
	C.WrapFpsController(keyboardToC, mouseToC, posToC, rotToC, speedToC, dtToC)
}

// Sleep Sleep the caller thread, this function will resume execution after waiting for at least the specified amount of time.
func Sleep(duration int64) {
	durationToC := C.int64_t(duration)
	C.WrapSleep(durationToC)
}

// AudioInit Initialize the audio system.
func AudioInit() bool {
	retval := C.WrapAudioInit()
	return bool(retval)
}

// AudioShutdown Shutdown the audio system.
func AudioShutdown() {
	C.WrapAudioShutdown()
}

// LoadWAVSoundFile Load a sound in WAV format from the local filesystem and return a reference to it.
func LoadWAVSoundFile(path string) int32 {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadWAVSoundFile(pathToC)
	return int32(retval)
}

// LoadWAVSoundAsset Load a sound in WAV format from the assets system and return a reference to it.  See [harfang.man.Assets].
func LoadWAVSoundAsset(name string) int32 {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapLoadWAVSoundAsset(nameToC)
	return int32(retval)
}

// LoadOGGSoundFile ...
func LoadOGGSoundFile(path string) int32 {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	retval := C.WrapLoadOGGSoundFile(pathToC)
	return int32(retval)
}

// LoadOGGSoundAsset ...
func LoadOGGSoundAsset(name string) int32 {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapLoadOGGSoundAsset(nameToC)
	return int32(retval)
}

// UnloadSound Unload a sound from the audio system.
func UnloadSound(snd SoundRef) {
	sndToC := C.int32_t(snd)
	C.WrapUnloadSound(sndToC)
}

// SetListener Set the listener transformation and velocity for spatialization by the audio system.
func SetListener(world *Mat4, velocity *Vec3) {
	worldToC := world.h
	velocityToC := velocity.h
	C.WrapSetListener(worldToC, velocityToC)
}

// PlayStereo Start playing a stereo sound. Return a handle to the started source.
func PlayStereo(snd SoundRef, state *StereoSourceState) int32 {
	sndToC := C.int32_t(snd)
	stateToC := state.h
	retval := C.WrapPlayStereo(sndToC, stateToC)
	return int32(retval)
}

// PlaySpatialized Start playing a spatialized sound. Return a handle to the started source.
func PlaySpatialized(snd SoundRef, state *SpatializedSourceState) int32 {
	sndToC := C.int32_t(snd)
	stateToC := state.h
	retval := C.WrapPlaySpatialized(sndToC, stateToC)
	return int32(retval)
}

// StreamWAVFileStereo Start an audio stream from a WAV file on the local filesystem.  See [harfang.man.Assets].
func StreamWAVFileStereo(path string, state *StereoSourceState) int32 {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	stateToC := state.h
	retval := C.WrapStreamWAVFileStereo(pathToC, stateToC)
	return int32(retval)
}

// StreamWAVAssetStereo Start an audio stream from a WAV file from the assets system.  See [harfang.man.Assets].
func StreamWAVAssetStereo(name string, state *StereoSourceState) int32 {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	stateToC := state.h
	retval := C.WrapStreamWAVAssetStereo(nameToC, stateToC)
	return int32(retval)
}

// StreamWAVFileSpatialized Start an audio stream from a WAV file on the local filesystem.  See [harfang.SetSourceTransform].
func StreamWAVFileSpatialized(path string, state *SpatializedSourceState) int32 {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	stateToC := state.h
	retval := C.WrapStreamWAVFileSpatialized(pathToC, stateToC)
	return int32(retval)
}

// StreamWAVAssetSpatialized Start an audio stream from a WAV file from the assets system.  See [harfang.SetSourceTransform] and [harfang.man.Assets].
func StreamWAVAssetSpatialized(name string, state *SpatializedSourceState) int32 {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	stateToC := state.h
	retval := C.WrapStreamWAVAssetSpatialized(nameToC, stateToC)
	return int32(retval)
}

// StreamOGGFileStereo ...
func StreamOGGFileStereo(path string, state *StereoSourceState) int32 {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	stateToC := state.h
	retval := C.WrapStreamOGGFileStereo(pathToC, stateToC)
	return int32(retval)
}

// StreamOGGAssetStereo ...
func StreamOGGAssetStereo(name string, state *StereoSourceState) int32 {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	stateToC := state.h
	retval := C.WrapStreamOGGAssetStereo(nameToC, stateToC)
	return int32(retval)
}

// StreamOGGFileSpatialized ...
func StreamOGGFileSpatialized(path string, state *SpatializedSourceState) int32 {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	stateToC := state.h
	retval := C.WrapStreamOGGFileSpatialized(pathToC, stateToC)
	return int32(retval)
}

// StreamOGGAssetSpatialized ...
func StreamOGGAssetSpatialized(name string, state *SpatializedSourceState) int32 {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	stateToC := state.h
	retval := C.WrapStreamOGGAssetSpatialized(nameToC, stateToC)
	return int32(retval)
}

// GetSourceDuration  Return the duration of an audio source.
func GetSourceDuration(source SourceRef) int64 {
	sourceToC := C.int32_t(source)
	retval := C.WrapGetSourceDuration(sourceToC)
	return int64(retval)
}

// GetSourceTimecode Return the current timecode of a playing audio source.
func GetSourceTimecode(source SourceRef) int64 {
	sourceToC := C.int32_t(source)
	retval := C.WrapGetSourceTimecode(sourceToC)
	return int64(retval)
}

// SetSourceTimecode Set timecode of the audio source.
func SetSourceTimecode(source SourceRef, t int64) bool {
	sourceToC := C.int32_t(source)
	tToC := C.int64_t(t)
	retval := C.WrapSetSourceTimecode(sourceToC, tToC)
	return bool(retval)
}

// SetSourceVolume Set audio source volume.
func SetSourceVolume(source SourceRef, volume float32) {
	sourceToC := C.int32_t(source)
	volumeToC := C.float(volume)
	C.WrapSetSourceVolume(sourceToC, volumeToC)
}

// SetSourcePanning Set a playing audio source panning.
func SetSourcePanning(source SourceRef, panning float32) {
	sourceToC := C.int32_t(source)
	panningToC := C.float(panning)
	C.WrapSetSourcePanning(sourceToC, panningToC)
}

// SetSourceRepeat Set audio source repeat mode.
func SetSourceRepeat(source SourceRef, repeat SourceRepeat) {
	sourceToC := C.int32_t(source)
	repeatToC := C.int32_t(repeat)
	C.WrapSetSourceRepeat(sourceToC, repeatToC)
}

// SetSourceTransform Set a playing spatialized audio source transformation.
func SetSourceTransform(source SourceRef, world *Mat4, velocity *Vec3) {
	sourceToC := C.int32_t(source)
	worldToC := world.h
	velocityToC := velocity.h
	C.WrapSetSourceTransform(sourceToC, worldToC, velocityToC)
}

// GetSourceState Return the state of an audio source.
func GetSourceState(source SourceRef) SourceState {
	sourceToC := C.int32_t(source)
	retval := C.WrapGetSourceState(sourceToC)
	return SourceState(retval)
}

// PauseSource Pause a playing audio source.  See [harfang.PlayStereo] and [harfang.PlaySpatialized].
func PauseSource(source SourceRef) {
	sourceToC := C.int32_t(source)
	C.WrapPauseSource(sourceToC)
}

// StopSource Stop a playing audio source.
func StopSource(source SourceRef) {
	sourceToC := C.int32_t(source)
	C.WrapStopSource(sourceToC)
}

// StopAllSources Stop all playing audio sources.
func StopAllSources() {
	C.WrapStopAllSources()
}

// OpenVRInit Initialize OpenVR. Start the device display, its controllers and trackers.
func OpenVRInit() bool {
	retval := C.WrapOpenVRInit()
	return bool(retval)
}

// OpenVRShutdown Shutdown OpenVR.
func OpenVRShutdown() {
	C.WrapOpenVRShutdown()
}

// OpenVRCreateEyeFrameBuffer Creates and returns an [harfang.man.VR] eye framebuffer, with the desired level of anti-aliasing. This function must be invoked twice, for the left and right eyes.
func OpenVRCreateEyeFrameBuffer() *OpenVREyeFrameBuffer {
	retval := C.WrapOpenVRCreateEyeFrameBuffer()
	retvalGO := &OpenVREyeFrameBuffer{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *OpenVREyeFrameBuffer) {
		C.WrapOpenVREyeFrameBufferFree(cleanval.h)
	})
	return retvalGO
}

// OpenVRCreateEyeFrameBufferWithAa Creates and returns an [harfang.man.VR] eye framebuffer, with the desired level of anti-aliasing. This function must be invoked twice, for the left and right eyes.
func OpenVRCreateEyeFrameBufferWithAa(aa OpenVRAA) *OpenVREyeFrameBuffer {
	aaToC := C.int32_t(aa)
	retval := C.WrapOpenVRCreateEyeFrameBufferWithAa(aaToC)
	retvalGO := &OpenVREyeFrameBuffer{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *OpenVREyeFrameBuffer) {
		C.WrapOpenVREyeFrameBufferFree(cleanval.h)
	})
	return retvalGO
}

// OpenVRDestroyEyeFrameBuffer Destroy an eye framebuffer.
func OpenVRDestroyEyeFrameBuffer(eyefb *OpenVREyeFrameBuffer) {
	eyefbToC := eyefb.h
	C.WrapOpenVRDestroyEyeFrameBuffer(eyefbToC)
}

// OpenVRGetState Returns the current OpenVR state including the body, head and eye transformations.
func OpenVRGetState(body *Mat4, znear float32, zfar float32) *OpenVRState {
	bodyToC := body.h
	znearToC := C.float(znear)
	zfarToC := C.float(zfar)
	retval := C.WrapOpenVRGetState(bodyToC, znearToC, zfarToC)
	retvalGO := &OpenVRState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *OpenVRState) {
		C.WrapOpenVRStateFree(cleanval.h)
	})
	return retvalGO
}

// OpenVRStateToViewState Compute the left and right eye view states from an OpenVR state.  See [harfang.OpenVRGetState].
func OpenVRStateToViewState(state *OpenVRState) (*ViewState, *ViewState) {
	stateToC := state.h
	left := NewViewState()
	leftToC := left.h
	right := NewViewState()
	rightToC := right.h
	C.WrapOpenVRStateToViewState(stateToC, leftToC, rightToC)
	return left, right
}

// OpenVRSubmitFrame Submit the left and right eye textures to the OpenVR compositor.  See [harfang.OpenVRCreateEyeFrameBuffer].
func OpenVRSubmitFrame(left *OpenVREyeFrameBuffer, right *OpenVREyeFrameBuffer) {
	leftToC := left.h
	rightToC := right.h
	C.WrapOpenVRSubmitFrame(leftToC, rightToC)
}

// OpenVRPostPresentHandoff Signal to the OpenVR compositor that it can immediatly start processing the current frame.
func OpenVRPostPresentHandoff() {
	C.WrapOpenVRPostPresentHandoff()
}

// OpenVRGetColorTexture Return the color texture attached to an eye framebuffer.
func OpenVRGetColorTexture(eye *OpenVREyeFrameBuffer) *Texture {
	eyeToC := eye.h
	retval := C.WrapOpenVRGetColorTexture(eyeToC)
	retvalGO := &Texture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Texture) {
		C.WrapTextureFree(cleanval.h)
	})
	return retvalGO
}

// OpenVRGetDepthTexture Return the depth texture attached to an eye framebuffer.
func OpenVRGetDepthTexture(eye *OpenVREyeFrameBuffer) *Texture {
	eyeToC := eye.h
	retval := C.WrapOpenVRGetDepthTexture(eyeToC)
	retvalGO := &Texture{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Texture) {
		C.WrapTextureFree(cleanval.h)
	})
	return retvalGO
}

// OpenVRGetFrameBufferSize ...
func OpenVRGetFrameBufferSize() *IVec2 {
	retval := C.WrapOpenVRGetFrameBufferSize()
	retvalGO := &IVec2{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVec2) {
		C.WrapIVec2Free(cleanval.h)
	})
	return retvalGO
}

// SRanipalInit Initial the SRanipal eye detection SDK.
func SRanipalInit() bool {
	retval := C.WrapSRanipalInit()
	return bool(retval)
}

// SRanipalShutdown Shutdown the SRanipal eye detection SDK.
func SRanipalShutdown() {
	C.WrapSRanipalShutdown()
}

// SRanipalLaunchEyeCalibration Launch the eye detection calibration sequence.
func SRanipalLaunchEyeCalibration() {
	C.WrapSRanipalLaunchEyeCalibration()
}

// SRanipalIsViveProEye Return `true` if the eye detection device in use is Vive Pro Eye.
func SRanipalIsViveProEye() bool {
	retval := C.WrapSRanipalIsViveProEye()
	return bool(retval)
}

// SRanipalGetState Return the current SRanipal device state.
func SRanipalGetState() *SRanipalState {
	retval := C.WrapSRanipalGetState()
	retvalGO := &SRanipalState{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SRanipalState) {
		C.WrapSRanipalStateFree(cleanval.h)
	})
	return retvalGO
}

// MakeVertex ...
func MakeVertex(pos *Vec3) *Vertex {
	posToC := pos.h
	retval := C.WrapMakeVertex(posToC)
	retvalGO := &Vertex{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vertex) {
		C.WrapVertexFree(cleanval.h)
	})
	return retvalGO
}

// MakeVertexWithNrm ...
func MakeVertexWithNrm(pos *Vec3, nrm *Vec3) *Vertex {
	posToC := pos.h
	nrmToC := nrm.h
	retval := C.WrapMakeVertexWithNrm(posToC, nrmToC)
	retvalGO := &Vertex{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vertex) {
		C.WrapVertexFree(cleanval.h)
	})
	return retvalGO
}

// MakeVertexWithNrmUv0 ...
func MakeVertexWithNrmUv0(pos *Vec3, nrm *Vec3, uv0 *Vec2) *Vertex {
	posToC := pos.h
	nrmToC := nrm.h
	uv0ToC := uv0.h
	retval := C.WrapMakeVertexWithNrmUv0(posToC, nrmToC, uv0ToC)
	retvalGO := &Vertex{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vertex) {
		C.WrapVertexFree(cleanval.h)
	})
	return retvalGO
}

// MakeVertexWithNrmUv0Color0 ...
func MakeVertexWithNrmUv0Color0(pos *Vec3, nrm *Vec3, uv0 *Vec2, color0 *Color) *Vertex {
	posToC := pos.h
	nrmToC := nrm.h
	uv0ToC := uv0.h
	color0ToC := color0.h
	retval := C.WrapMakeVertexWithNrmUv0Color0(posToC, nrmToC, uv0ToC, color0ToC)
	retvalGO := &Vertex{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Vertex) {
		C.WrapVertexFree(cleanval.h)
	})
	return retvalGO
}

// SaveGeometryToFile Save a geometry to the local filesystem.  Note that in order to render a geometry it must have been converted to model by the asset compiler.  See [harfang.GeometryBuilder] and [harfang.ModelBuilder].
func SaveGeometryToFile(path string, geo *Geometry) bool {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	geoToC := geo.h
	retval := C.WrapSaveGeometryToFile(pathToC, geoToC)
	return bool(retval)
}

// NewIsoSurface Return a new iso-surface object.  See [harfang.IsoSurfaceSphere] to draw to an iso-surface and [harfang.IsoSurfaceToModel] to draw it.
func NewIsoSurface(width int32, height int32, depth int32) *IsoSurface {
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	depthToC := C.int32_t(depth)
	retval := C.WrapNewIsoSurface(widthToC, heightToC, depthToC)
	retvalGO := &IsoSurface{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IsoSurface) {
		C.WrapIsoSurfaceFree(cleanval.h)
	})
	return retvalGO
}

// IsoSurfaceSphere Output a sphere to an iso-surface.
func IsoSurfaceSphere(surface *IsoSurface, width int32, height int32, depth int32, x float32, y float32, z float32, radius float32) {
	surfaceToC := surface.h
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	depthToC := C.int32_t(depth)
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	radiusToC := C.float(radius)
	C.WrapIsoSurfaceSphere(surfaceToC, widthToC, heightToC, depthToC, xToC, yToC, zToC, radiusToC)
}

// IsoSurfaceSphereWithValue Output a sphere to an iso-surface.
func IsoSurfaceSphereWithValue(surface *IsoSurface, width int32, height int32, depth int32, x float32, y float32, z float32, radius float32, value float32) {
	surfaceToC := surface.h
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	depthToC := C.int32_t(depth)
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	radiusToC := C.float(radius)
	valueToC := C.float(value)
	C.WrapIsoSurfaceSphereWithValue(surfaceToC, widthToC, heightToC, depthToC, xToC, yToC, zToC, radiusToC, valueToC)
}

// IsoSurfaceSphereWithValueExponent Output a sphere to an iso-surface.
func IsoSurfaceSphereWithValueExponent(surface *IsoSurface, width int32, height int32, depth int32, x float32, y float32, z float32, radius float32, value float32, exponent float32) {
	surfaceToC := surface.h
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	depthToC := C.int32_t(depth)
	xToC := C.float(x)
	yToC := C.float(y)
	zToC := C.float(z)
	radiusToC := C.float(radius)
	valueToC := C.float(value)
	exponentToC := C.float(exponent)
	C.WrapIsoSurfaceSphereWithValueExponent(surfaceToC, widthToC, heightToC, depthToC, xToC, yToC, zToC, radiusToC, valueToC, exponentToC)
}

// GaussianBlurIsoSurface Apply a Gaussian blur to an iso-surface.
func GaussianBlurIsoSurface(surface *IsoSurface, width int32, height int32, depth int32) *IsoSurface {
	surfaceToC := surface.h
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	depthToC := C.int32_t(depth)
	retval := C.WrapGaussianBlurIsoSurface(surfaceToC, widthToC, heightToC, depthToC)
	retvalGO := &IsoSurface{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IsoSurface) {
		C.WrapIsoSurfaceFree(cleanval.h)
	})
	return retvalGO
}

// IsoSurfaceToModel Convert an iso-surface to a render model, this function is geared toward efficiency and meant for realtime.
func IsoSurfaceToModel(builder *ModelBuilder, surface *IsoSurface, width int32, height int32, depth int32) bool {
	builderToC := builder.h
	surfaceToC := surface.h
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	depthToC := C.int32_t(depth)
	retval := C.WrapIsoSurfaceToModel(builderToC, surfaceToC, widthToC, heightToC, depthToC)
	return bool(retval)
}

// IsoSurfaceToModelWithMaterial Convert an iso-surface to a render model, this function is geared toward efficiency and meant for realtime.
func IsoSurfaceToModelWithMaterial(builder *ModelBuilder, surface *IsoSurface, width int32, height int32, depth int32, material uint16) bool {
	builderToC := builder.h
	surfaceToC := surface.h
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	depthToC := C.int32_t(depth)
	materialToC := C.ushort(material)
	retval := C.WrapIsoSurfaceToModelWithMaterial(builderToC, surfaceToC, widthToC, heightToC, depthToC, materialToC)
	return bool(retval)
}

// IsoSurfaceToModelWithMaterialIsolevel Convert an iso-surface to a render model, this function is geared toward efficiency and meant for realtime.
func IsoSurfaceToModelWithMaterialIsolevel(builder *ModelBuilder, surface *IsoSurface, width int32, height int32, depth int32, material uint16, isolevel float32) bool {
	builderToC := builder.h
	surfaceToC := surface.h
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	depthToC := C.int32_t(depth)
	materialToC := C.ushort(material)
	isolevelToC := C.float(isolevel)
	retval := C.WrapIsoSurfaceToModelWithMaterialIsolevel(builderToC, surfaceToC, widthToC, heightToC, depthToC, materialToC, isolevelToC)
	return bool(retval)
}

// IsoSurfaceToModelWithMaterialIsolevelScaleXScaleYScaleZ Convert an iso-surface to a render model, this function is geared toward efficiency and meant for realtime.
func IsoSurfaceToModelWithMaterialIsolevelScaleXScaleYScaleZ(builder *ModelBuilder, surface *IsoSurface, width int32, height int32, depth int32, material uint16, isolevel float32, scalex float32, scaley float32, scalez float32) bool {
	builderToC := builder.h
	surfaceToC := surface.h
	widthToC := C.int32_t(width)
	heightToC := C.int32_t(height)
	depthToC := C.int32_t(depth)
	materialToC := C.ushort(material)
	isolevelToC := C.float(isolevel)
	scalexToC := C.float(scalex)
	scaleyToC := C.float(scaley)
	scalezToC := C.float(scalez)
	retval := C.WrapIsoSurfaceToModelWithMaterialIsolevelScaleXScaleYScaleZ(builderToC, surfaceToC, widthToC, heightToC, depthToC, materialToC, isolevelToC, scalexToC, scaleyToC, scalezToC)
	return bool(retval)
}

// CreateBloomFromFile ...
func CreateBloomFromFile(path string, ratio BackbufferRatio) *Bloom {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	ratioToC := C.int32_t(ratio)
	retval := C.WrapCreateBloomFromFile(pathToC, ratioToC)
	retvalGO := &Bloom{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Bloom) {
		C.WrapBloomFree(cleanval.h)
	})
	return retvalGO
}

// CreateBloomFromAssets ...
func CreateBloomFromAssets(path string, ratio BackbufferRatio) *Bloom {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	ratioToC := C.int32_t(ratio)
	retval := C.WrapCreateBloomFromAssets(pathToC, ratioToC)
	retvalGO := &Bloom{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *Bloom) {
		C.WrapBloomFree(cleanval.h)
	})
	return retvalGO
}

// DestroyBloom Destroy a bloom post process object and all associated resources.
func DestroyBloom(bloom *Bloom) {
	bloomToC := bloom.h
	C.WrapDestroyBloom(bloomToC)
}

// ApplyBloom Process `input` texture and generate a bloom overlay on top of `output`, input and output must be of the same size.  Use [harfang.CreateBloomFromFile]/[harfang.CreateBloomFromAssets] to create a [harfang.Bloom] object and [harfang.DestroyBloom] to destroy its internal resources after usage.
func ApplyBloom(viewid *uint16, rect *IntRect, input *Texture, output *FrameBufferHandle, bloom *Bloom, threshold float32, smoothness float32, intensity float32) {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	rectToC := rect.h
	inputToC := input.h
	outputToC := output.h
	bloomToC := bloom.h
	thresholdToC := C.float(threshold)
	smoothnessToC := C.float(smoothness)
	intensityToC := C.float(intensity)
	C.WrapApplyBloom(viewidToC, rectToC, inputToC, outputToC, bloomToC, thresholdToC, smoothnessToC, intensityToC)
}

// CreateSAOFromFile ...
func CreateSAOFromFile(path string, ratio BackbufferRatio) *SAO {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	ratioToC := C.int32_t(ratio)
	retval := C.WrapCreateSAOFromFile(pathToC, ratioToC)
	retvalGO := &SAO{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SAO) {
		C.WrapSAOFree(cleanval.h)
	})
	return retvalGO
}

// CreateSAOFromAssets ...
func CreateSAOFromAssets(path string, ratio BackbufferRatio) *SAO {
	pathToC, idFinpathToC := wrapString(path)
	defer idFinpathToC()
	ratioToC := C.int32_t(ratio)
	retval := C.WrapCreateSAOFromAssets(pathToC, ratioToC)
	retvalGO := &SAO{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *SAO) {
		C.WrapSAOFree(cleanval.h)
	})
	return retvalGO
}

// DestroySAO Destroy an ambient occlusion post process object and its resources.
func DestroySAO(sao *SAO) {
	saoToC := sao.h
	C.WrapDestroySAO(saoToC)
}

// ComputeSAO ...
func ComputeSAO(viewid *uint16, rect *IntRect, attr0 *Texture, attr1 *Texture, noise *Texture, output *FrameBufferHandle, sao *SAO, projection *Mat44, bias float32, radius float32, samplecount int32, sharpness float32) {
	viewidToC := (*C.ushort)(unsafe.Pointer(viewid))
	rectToC := rect.h
	attr0ToC := attr0.h
	attr1ToC := attr1.h
	noiseToC := noise.h
	outputToC := output.h
	saoToC := sao.h
	projectionToC := projection.h
	biasToC := C.float(bias)
	radiusToC := C.float(radius)
	samplecountToC := C.int32_t(samplecount)
	sharpnessToC := C.float(sharpness)
	C.WrapComputeSAO(viewidToC, rectToC, attr0ToC, attr1ToC, noiseToC, outputToC, saoToC, projectionToC, biasToC, radiusToC, samplecountToC, sharpnessToC)
}

// BeginProfilerSection Begin a named profiler section. Call [harfang.EndProfilerSection] to end the section.
func BeginProfilerSection(name string) int32 {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	retval := C.WrapBeginProfilerSection(nameToC)
	return int32(retval)
}

// BeginProfilerSectionWithSectionDetails Begin a named profiler section. Call [harfang.EndProfilerSection] to end the section.
func BeginProfilerSectionWithSectionDetails(name string, sectiondetails string) int32 {
	nameToC, idFinnameToC := wrapString(name)
	defer idFinnameToC()
	sectiondetailsToC, idFinsectiondetailsToC := wrapString(sectiondetails)
	defer idFinsectiondetailsToC()
	retval := C.WrapBeginProfilerSectionWithSectionDetails(nameToC, sectiondetailsToC)
	return int32(retval)
}

// EndProfilerSection End a named profiler section. Call [harfang.BeginProfilerSection] to begin a new section.
func EndProfilerSection(sectionidx int32) {
	sectionidxToC := C.size_t(sectionidx)
	C.WrapEndProfilerSection(sectionidxToC)
}

// EndProfilerFrame End a profiler frame and return it.  See [harfang.PrintProfilerFrame] to print a profiler frame to the console.
func EndProfilerFrame() *ProfilerFrame {
	retval := C.WrapEndProfilerFrame()
	retvalGO := &ProfilerFrame{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ProfilerFrame) {
		C.WrapProfilerFrameFree(cleanval.h)
	})
	return retvalGO
}

// CaptureProfilerFrame Capture the current profiler frame but do not end it. See [harfang.EndProfilerFrame] to capture and end the current profiler frame.  See [harfang.PrintProfilerFrame] to print a profiler frame to the console.
func CaptureProfilerFrame() *ProfilerFrame {
	retval := C.WrapCaptureProfilerFrame()
	retvalGO := &ProfilerFrame{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *ProfilerFrame) {
		C.WrapProfilerFrameFree(cleanval.h)
	})
	return retvalGO
}

// PrintProfilerFrame Print a profiler frame to the console. Print all sections in the frame, their duration and event count.
func PrintProfilerFrame(profilerframe *ProfilerFrame) {
	profilerframeToC := profilerframe.h
	C.WrapPrintProfilerFrame(profilerframeToC)
}

// MakeVideoStreamer ...
func MakeVideoStreamer(modulepath string) *IVideoStreamer {
	modulepathToC, idFinmodulepathToC := wrapString(modulepath)
	defer idFinmodulepathToC()
	retval := C.WrapMakeVideoStreamer(modulepathToC)
	retvalGO := &IVideoStreamer{h: retval}
	runtime.SetFinalizer(retvalGO, func(cleanval *IVideoStreamer) {
		C.WrapIVideoStreamerFree(cleanval.h)
	})
	return retvalGO
}

// UpdateTexture ...
func UpdateTexture(streamer *IVideoStreamer, handle *uintptr, texture *Texture, size *IVec2, format *int32) bool {
	streamerToC := streamer.h
	handleToC := (*C.intptr_t)(unsafe.Pointer(handle))
	textureToC := texture.h
	sizeToC := size.h
	formatToC := (*C.int32_t)(unsafe.Pointer(format))
	retval := C.WrapUpdateTexture(streamerToC, handleToC, textureToC, sizeToC, formatToC)
	return bool(retval)
}

// UpdateTextureWithDestroy ...
func UpdateTextureWithDestroy(streamer *IVideoStreamer, handle *uintptr, texture *Texture, size *IVec2, format *int32, destroy bool) bool {
	streamerToC := streamer.h
	handleToC := (*C.intptr_t)(unsafe.Pointer(handle))
	textureToC := texture.h
	sizeToC := size.h
	formatToC := (*C.int32_t)(unsafe.Pointer(format))
	destroyToC := C.bool(destroy)
	retval := C.WrapUpdateTextureWithDestroy(streamerToC, handleToC, textureToC, sizeToC, formatToC, destroyToC)
	return bool(retval)
}

// CastForwardPipelineToPipeline ...
func CastForwardPipelineToPipeline(o *ForwardPipeline) *Pipeline {
	oToC := o.h
	retval := C.WrapCastForwardPipelineToPipeline(oToC)
	var retvalGO *Pipeline
	if retval != nil {
		retvalGO = &Pipeline{h: retval}
		runtime.SetFinalizer(retvalGO, func(cleanval *Pipeline) {
			C.WrapPipelineFree(cleanval.h)
		})
	}
	return retvalGO
}

// CastPipelineToForwardPipeline ...
func CastPipelineToForwardPipeline(o *Pipeline) *ForwardPipeline {
	oToC := o.h
	retval := C.WrapCastPipelineToForwardPipeline(oToC)
	var retvalGO *ForwardPipeline
	if retval != nil {
		retvalGO = &ForwardPipeline{h: retval}
		runtime.SetFinalizer(retvalGO, func(cleanval *ForwardPipeline) {
			C.WrapForwardPipelineFree(cleanval.h)
		})
	}
	return retvalGO
}

// ResetFlags ...
type ResetFlags uint32

// RFNone ...
var RFNone = ResetFlags(C.WrapGetRFNone())

// RFMSAA2X ...
var RFMSAA2X = ResetFlags(C.WrapGetRFMSAA2X())

// RFMSAA4X ...
var RFMSAA4X = ResetFlags(C.WrapGetRFMSAA4X())

// RFMSAA8X ...
var RFMSAA8X = ResetFlags(C.WrapGetRFMSAA8X())

// RFMSAA16X ...
var RFMSAA16X = ResetFlags(C.WrapGetRFMSAA16X())

// RFVSync ...
var RFVSync = ResetFlags(C.WrapGetRFVSync())

// RFMaxAnisotropy ...
var RFMaxAnisotropy = ResetFlags(C.WrapGetRFMaxAnisotropy())

// RFCapture ...
var RFCapture = ResetFlags(C.WrapGetRFCapture())

// RFFlushAfterRender ...
var RFFlushAfterRender = ResetFlags(C.WrapGetRFFlushAfterRender())

// RFFlipAfterRender ...
var RFFlipAfterRender = ResetFlags(C.WrapGetRFFlipAfterRender())

// RFSRGBBackBuffer ...
var RFSRGBBackBuffer = ResetFlags(C.WrapGetRFSRGBBackBuffer())

// RFHDR10 ...
var RFHDR10 = ResetFlags(C.WrapGetRFHDR10())

// RFHiDPI ...
var RFHiDPI = ResetFlags(C.WrapGetRFHiDPI())

// RFDepthClamp ...
var RFDepthClamp = ResetFlags(C.WrapGetRFDepthClamp())

// RFSuspend ...
var RFSuspend = ResetFlags(C.WrapGetRFSuspend())

// DebugFlags ...
type DebugFlags uint32

// DFIFH ...
var DFIFH = DebugFlags(C.WrapGetDFIFH())

// DFProfiler ...
var DFProfiler = DebugFlags(C.WrapGetDFProfiler())

// DFStats ...
var DFStats = DebugFlags(C.WrapGetDFStats())

// DFText ...
var DFText = DebugFlags(C.WrapGetDFText())

// DFWireframe ...
var DFWireframe = DebugFlags(C.WrapGetDFWireframe())

// ClearFlags ...
type ClearFlags uint16

// CFNone ...
var CFNone = ClearFlags(C.WrapGetCFNone())

// CFColor ...
var CFColor = ClearFlags(C.WrapGetCFColor())

// CFDepth ...
var CFDepth = ClearFlags(C.WrapGetCFDepth())

// CFStencil ...
var CFStencil = ClearFlags(C.WrapGetCFStencil())

// CFDiscardColor0 ...
var CFDiscardColor0 = ClearFlags(C.WrapGetCFDiscardColor0())

// CFDiscardColor1 ...
var CFDiscardColor1 = ClearFlags(C.WrapGetCFDiscardColor1())

// CFDiscardColor2 ...
var CFDiscardColor2 = ClearFlags(C.WrapGetCFDiscardColor2())

// CFDiscardColor3 ...
var CFDiscardColor3 = ClearFlags(C.WrapGetCFDiscardColor3())

// CFDiscardColor4 ...
var CFDiscardColor4 = ClearFlags(C.WrapGetCFDiscardColor4())

// CFDiscardColor5 ...
var CFDiscardColor5 = ClearFlags(C.WrapGetCFDiscardColor5())

// CFDiscardColor6 ...
var CFDiscardColor6 = ClearFlags(C.WrapGetCFDiscardColor6())

// CFDiscardColor7 ...
var CFDiscardColor7 = ClearFlags(C.WrapGetCFDiscardColor7())

// CFDiscardDepth ...
var CFDiscardDepth = ClearFlags(C.WrapGetCFDiscardDepth())

// CFDiscardStencil ...
var CFDiscardStencil = ClearFlags(C.WrapGetCFDiscardStencil())

// CFDiscardColorAll ...
var CFDiscardColorAll = ClearFlags(C.WrapGetCFDiscardColorAll())

// CFDiscardAll ...
var CFDiscardAll = ClearFlags(C.WrapGetCFDiscardAll())

// TextureFlags ...
type TextureFlags uint64

// TFUMirror ...
var TFUMirror = TextureFlags(C.WrapGetTFUMirror())

// TFUClamp ...
var TFUClamp = TextureFlags(C.WrapGetTFUClamp())

// TFUBorder ...
var TFUBorder = TextureFlags(C.WrapGetTFUBorder())

// TFVMirror ...
var TFVMirror = TextureFlags(C.WrapGetTFVMirror())

// TFVClamp ...
var TFVClamp = TextureFlags(C.WrapGetTFVClamp())

// TFVBorder ...
var TFVBorder = TextureFlags(C.WrapGetTFVBorder())

// TFWMirror ...
var TFWMirror = TextureFlags(C.WrapGetTFWMirror())

// TFWClamp ...
var TFWClamp = TextureFlags(C.WrapGetTFWClamp())

// TFWBorder ...
var TFWBorder = TextureFlags(C.WrapGetTFWBorder())

// TFSamplerMinPoint ...
var TFSamplerMinPoint = TextureFlags(C.WrapGetTFSamplerMinPoint())

// TFSamplerMinAnisotropic ...
var TFSamplerMinAnisotropic = TextureFlags(C.WrapGetTFSamplerMinAnisotropic())

// TFSamplerMagPoint ...
var TFSamplerMagPoint = TextureFlags(C.WrapGetTFSamplerMagPoint())

// TFSamplerMagAnisotropic ...
var TFSamplerMagAnisotropic = TextureFlags(C.WrapGetTFSamplerMagAnisotropic())

// TFBlitDestination ...
var TFBlitDestination = TextureFlags(C.WrapGetTFBlitDestination())

// TFReadBack ...
var TFReadBack = TextureFlags(C.WrapGetTFReadBack())

// TFRenderTarget ...
var TFRenderTarget = TextureFlags(C.WrapGetTFRenderTarget())

// LoadSaveSceneFlags ...
type LoadSaveSceneFlags uint32

// LSSFNodes ...
var LSSFNodes = LoadSaveSceneFlags(C.WrapGetLSSFNodes())

// LSSFScene ...
var LSSFScene = LoadSaveSceneFlags(C.WrapGetLSSFScene())

// LSSFAnims ...
var LSSFAnims = LoadSaveSceneFlags(C.WrapGetLSSFAnims())

// LSSFKeyValues ...
var LSSFKeyValues = LoadSaveSceneFlags(C.WrapGetLSSFKeyValues())

// LSSFPhysics ...
var LSSFPhysics = LoadSaveSceneFlags(C.WrapGetLSSFPhysics())

// LSSFScripts ...
var LSSFScripts = LoadSaveSceneFlags(C.WrapGetLSSFScripts())

// LSSFAll ...
var LSSFAll = LoadSaveSceneFlags(C.WrapGetLSSFAll())

// LSSFQueueTextureLoads ...
var LSSFQueueTextureLoads = LoadSaveSceneFlags(C.WrapGetLSSFQueueTextureLoads())

// LSSFFreezeMatrixToTransformOnSave ...
var LSSFFreezeMatrixToTransformOnSave = LoadSaveSceneFlags(C.WrapGetLSSFFreezeMatrixToTransformOnSave())

// LSSFQueueModelLoads ...
var LSSFQueueModelLoads = LoadSaveSceneFlags(C.WrapGetLSSFQueueModelLoads())

// LSSFDoNotChangeCurrentCameraIfValid ...
var LSSFDoNotChangeCurrentCameraIfValid = LoadSaveSceneFlags(C.WrapGetLSSFDoNotChangeCurrentCameraIfValid())

// SoundRef ...
type SoundRef int32

// SNDInvalid ...
var SNDInvalid = SoundRef(C.WrapGetSNDInvalid())

// SourceRef ...
type SourceRef int32

// SRCInvalid ...
var SRCInvalid = SourceRef(C.WrapGetSRCInvalid())

// InvalidFrameBufferHandle ...
var InvalidFrameBufferHandle = FrameBufferHandle{h: C.WrapGetInvalidFrameBufferHandle()}

// InvalidModelRef ...
var InvalidModelRef = ModelRef{h: C.WrapGetInvalidModelRef()}

// InvalidTextureRef ...
var InvalidTextureRef = TextureRef{h: C.WrapGetInvalidTextureRef()}

// InvalidMaterialRef ...
var InvalidMaterialRef = MaterialRef{h: C.WrapGetInvalidMaterialRef()}

// InvalidPipelineProgramRef ...
var InvalidPipelineProgramRef = PipelineProgramRef{h: C.WrapGetInvalidPipelineProgramRef()}

// InvalidSceneAnimRef ...
var InvalidSceneAnimRef = SceneAnimRef{h: C.WrapGetInvalidSceneAnimRef()}

// UnspecifiedAnimTime ...
var UnspecifiedAnimTime = int64(C.WrapGetUnspecifiedAnimTime())

// NullNode ...
var NullNode = Node{h: C.WrapGetNullNode()}

// OnTextInput ...
var OnTextInput = SignalReturningVoidTakingConstCharPtr{h: C.WrapGetOnTextInput()}

// InvalidAudioStreamRef ...
var InvalidAudioStreamRef = int32(C.WrapGetInvalidAudioStreamRef())
