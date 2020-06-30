// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY BELGIAN ROBOTS.

// WARNING: This file has automatically been generated on Mon, 29 Jun 2020 19:56:25 MDT.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package openal_go

/*
#cgo pkg-config: openal
#include "AL/al.h"
#include "AL/alc.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// __GO__ function as declared in project_openal_go/<predefine>:36
// func __GO__(arg0 []byte) {
// 	carg0, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&arg0)).Data)), cgoAllocsUnknown
// 	C.__GO__(carg0)
// }

// AlDopplerFactor function as declared in AL/al.h:401
func AlDopplerFactor(value ALfloat) {
	cvalue, _ := (C.ALfloat)(value), cgoAllocsUnknown
	C.alDopplerFactor(cvalue)
}

// AlDopplerVelocity function as declared in AL/al.h:409
func AlDopplerVelocity(value ALfloat) {
	cvalue, _ := (C.ALfloat)(value), cgoAllocsUnknown
	C.alDopplerVelocity(cvalue)
}

// AlSpeedOfSound function as declared in AL/al.h:421
func AlSpeedOfSound(value ALfloat) {
	cvalue, _ := (C.ALfloat)(value), cgoAllocsUnknown
	C.alSpeedOfSound(cvalue)
}

// AlDistanceModel function as declared in AL/al.h:442
func AlDistanceModel(distanceModel ALenum) {
	cdistanceModel, _ := (C.ALenum)(distanceModel), cgoAllocsUnknown
	C.alDistanceModel(cdistanceModel)
}

// AlEnable function as declared in AL/al.h:453
func AlEnable(capability ALenum) {
	ccapability, _ := (C.ALenum)(capability), cgoAllocsUnknown
	C.alEnable(ccapability)
}

// AlDisable function as declared in AL/al.h:454
func AlDisable(capability ALenum) {
	ccapability, _ := (C.ALenum)(capability), cgoAllocsUnknown
	C.alDisable(ccapability)
}

// AlIsEnabled function as declared in AL/al.h:455
func AlIsEnabled(capability ALenum) ALboolean {
	ccapability, _ := (C.ALenum)(capability), cgoAllocsUnknown
	__ret := C.alIsEnabled(ccapability)
	__v := (ALboolean)(__ret)
	return __v
}

// AlGetString function as declared in AL/al.h:458
// REMOVED by DORIAN due to incompatibility
// func AlGetString(param ALenum) *ALchar {
// 	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
// 	__ret := C.alGetString(cparam)
	
// 	__v := packPALcharString(__ret)
// 	return __v
// }

// AlGetBooleanv function as declared in AL/al.h:459
func AlGetBooleanv(param ALenum, values []ALboolean) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALboolean)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alGetBooleanv(cparam, cvalues)
}

// AlGetIntegerv function as declared in AL/al.h:460
func AlGetIntegerv(param ALenum, values []ALint) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alGetIntegerv(cparam, cvalues)
}

// AlGetFloatv function as declared in AL/al.h:461
func AlGetFloatv(param ALenum, values []ALfloat) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alGetFloatv(cparam, cvalues)
}

// AlGetDoublev function as declared in AL/al.h:462
func AlGetDoublev(param ALenum, values []ALdouble) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALdouble)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alGetDoublev(cparam, cvalues)
}

// AlGetBoolean function as declared in AL/al.h:463
func AlGetBoolean(param ALenum) ALboolean {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	__ret := C.alGetBoolean(cparam)
	__v := (ALboolean)(__ret)
	return __v
}

// AlGetInteger function as declared in AL/al.h:464
func AlGetInteger(param ALenum) ALint {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	__ret := C.alGetInteger(cparam)
	__v := (ALint)(__ret)
	return __v
}

// AlGetFloat function as declared in AL/al.h:465
func AlGetFloat(param ALenum) ALfloat {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	__ret := C.alGetFloat(cparam)
	__v := (ALfloat)(__ret)
	return __v
}

// AlGetDouble function as declared in AL/al.h:466
func AlGetDouble(param ALenum) ALdouble {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	__ret := C.alGetDouble(cparam)
	__v := (ALdouble)(__ret)
	return __v
}

// AlGetError function as declared in AL/al.h:473
func AlGetError() ALenum {
	__ret := C.alGetError()
	__v := (ALenum)(__ret)
	return __v
}

// AlIsExtensionPresent function as declared in AL/al.h:481
func AlIsExtensionPresent(extname []ALchar) ALboolean {
	cextname, _ := unpackArgSALchar(extname)
	__ret := C.alIsExtensionPresent(cextname)
	packSALchar(extname, cextname)
	__v := (ALboolean)(__ret)
	return __v
}

// AlGetProcAddress function as declared in AL/al.h:482
func AlGetProcAddress(fname []ALchar) unsafe.Pointer {
	cfname, _ := unpackArgSALchar(fname)
	__ret := C.alGetProcAddress(cfname)
	packSALchar(fname, cfname)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// AlGetEnumValue function as declared in AL/al.h:483
func AlGetEnumValue(ename []ALchar) ALenum {
	cename, _ := unpackArgSALchar(ename)
	__ret := C.alGetEnumValue(cename)
	packSALchar(ename, cename)
	__v := (ALenum)(__ret)
	return __v
}

// AlListenerf function as declared in AL/al.h:487
func AlListenerf(param ALenum, value ALfloat) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue, _ := (C.ALfloat)(value), cgoAllocsUnknown
	C.alListenerf(cparam, cvalue)
}

// AlListener3f function as declared in AL/al.h:488
func AlListener3f(param ALenum, value1 ALfloat, value2 ALfloat, value3 ALfloat) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue1, _ := (C.ALfloat)(value1), cgoAllocsUnknown
	cvalue2, _ := (C.ALfloat)(value2), cgoAllocsUnknown
	cvalue3, _ := (C.ALfloat)(value3), cgoAllocsUnknown
	C.alListener3f(cparam, cvalue1, cvalue2, cvalue3)
}

// AlListenerfv function as declared in AL/al.h:489
func AlListenerfv(param ALenum, values []ALfloat) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alListenerfv(cparam, cvalues)
}

// AlListeneri function as declared in AL/al.h:490
func AlListeneri(param ALenum, value ALint) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue, _ := (C.ALint)(value), cgoAllocsUnknown
	C.alListeneri(cparam, cvalue)
}

// AlListener3i function as declared in AL/al.h:491
func AlListener3i(param ALenum, value1 ALint, value2 ALint, value3 ALint) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue1, _ := (C.ALint)(value1), cgoAllocsUnknown
	cvalue2, _ := (C.ALint)(value2), cgoAllocsUnknown
	cvalue3, _ := (C.ALint)(value3), cgoAllocsUnknown
	C.alListener3i(cparam, cvalue1, cvalue2, cvalue3)
}

// AlListeneriv function as declared in AL/al.h:492
func AlListeneriv(param ALenum, values []ALint) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alListeneriv(cparam, cvalues)
}

// AlGetListenerf function as declared in AL/al.h:495
func AlGetListenerf(param ALenum, value []ALfloat) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value)).Data)), cgoAllocsUnknown
	C.alGetListenerf(cparam, cvalue)
}

// AlGetListener3f function as declared in AL/al.h:496
func AlGetListener3f(param ALenum, value1 []ALfloat, value2 []ALfloat, value3 []ALfloat) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue1, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value1)).Data)), cgoAllocsUnknown
	cvalue2, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value2)).Data)), cgoAllocsUnknown
	cvalue3, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value3)).Data)), cgoAllocsUnknown
	C.alGetListener3f(cparam, cvalue1, cvalue2, cvalue3)
}

// AlGetListenerfv function as declared in AL/al.h:497
func AlGetListenerfv(param ALenum, values []ALfloat) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alGetListenerfv(cparam, cvalues)
}

// AlGetListeneri function as declared in AL/al.h:498
func AlGetListeneri(param ALenum, value []ALint) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value)).Data)), cgoAllocsUnknown
	C.alGetListeneri(cparam, cvalue)
}

// AlGetListener3i function as declared in AL/al.h:499
func AlGetListener3i(param ALenum, value1 []ALint, value2 []ALint, value3 []ALint) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue1, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value1)).Data)), cgoAllocsUnknown
	cvalue2, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value2)).Data)), cgoAllocsUnknown
	cvalue3, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value3)).Data)), cgoAllocsUnknown
	C.alGetListener3i(cparam, cvalue1, cvalue2, cvalue3)
}

// AlGetListeneriv function as declared in AL/al.h:500
func AlGetListeneriv(param ALenum, values []ALint) {
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alGetListeneriv(cparam, cvalues)
}

// AlGenSources function as declared in AL/al.h:504
func AlGenSources(n ALsizei, sources []ALuint) {
	cn, _ := (C.ALsizei)(n), cgoAllocsUnknown
	csources, _ := (*C.ALuint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&sources)).Data)), cgoAllocsUnknown
	C.alGenSources(cn, csources)
}

// AlDeleteSources function as declared in AL/al.h:506
func AlDeleteSources(n ALsizei, sources []ALuint) {
	cn, _ := (C.ALsizei)(n), cgoAllocsUnknown
	csources, _ := (*C.ALuint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&sources)).Data)), cgoAllocsUnknown
	C.alDeleteSources(cn, csources)
}

// AlIsSource function as declared in AL/al.h:508
func AlIsSource(source ALuint) ALboolean {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	__ret := C.alIsSource(csource)
	__v := (ALboolean)(__ret)
	return __v
}

// AlSourcef function as declared in AL/al.h:511
func AlSourcef(source ALuint, param ALenum, value ALfloat) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue, _ := (C.ALfloat)(value), cgoAllocsUnknown
	C.alSourcef(csource, cparam, cvalue)
}

// AlSource3f function as declared in AL/al.h:512
func AlSource3f(source ALuint, param ALenum, value1 ALfloat, value2 ALfloat, value3 ALfloat) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue1, _ := (C.ALfloat)(value1), cgoAllocsUnknown
	cvalue2, _ := (C.ALfloat)(value2), cgoAllocsUnknown
	cvalue3, _ := (C.ALfloat)(value3), cgoAllocsUnknown
	C.alSource3f(csource, cparam, cvalue1, cvalue2, cvalue3)
}

// AlSourcefv function as declared in AL/al.h:513
func AlSourcefv(source ALuint, param ALenum, values []ALfloat) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alSourcefv(csource, cparam, cvalues)
}

// AlSourcei function as declared in AL/al.h:514
func AlSourcei(source ALuint, param ALenum, value ALint) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue, _ := (C.ALint)(value), cgoAllocsUnknown
	C.alSourcei(csource, cparam, cvalue)
}

// AlSource3i function as declared in AL/al.h:515
func AlSource3i(source ALuint, param ALenum, value1 ALint, value2 ALint, value3 ALint) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue1, _ := (C.ALint)(value1), cgoAllocsUnknown
	cvalue2, _ := (C.ALint)(value2), cgoAllocsUnknown
	cvalue3, _ := (C.ALint)(value3), cgoAllocsUnknown
	C.alSource3i(csource, cparam, cvalue1, cvalue2, cvalue3)
}

// AlSourceiv function as declared in AL/al.h:516
func AlSourceiv(source ALuint, param ALenum, values []ALint) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alSourceiv(csource, cparam, cvalues)
}

// AlGetSourcef function as declared in AL/al.h:519
func AlGetSourcef(source ALuint, param ALenum, value []ALfloat) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value)).Data)), cgoAllocsUnknown
	C.alGetSourcef(csource, cparam, cvalue)
}

// AlGetSource3f function as declared in AL/al.h:520
func AlGetSource3f(source ALuint, param ALenum, value1 []ALfloat, value2 []ALfloat, value3 []ALfloat) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue1, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value1)).Data)), cgoAllocsUnknown
	cvalue2, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value2)).Data)), cgoAllocsUnknown
	cvalue3, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value3)).Data)), cgoAllocsUnknown
	C.alGetSource3f(csource, cparam, cvalue1, cvalue2, cvalue3)
}

// AlGetSourcefv function as declared in AL/al.h:521
func AlGetSourcefv(source ALuint, param ALenum, values []ALfloat) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alGetSourcefv(csource, cparam, cvalues)
}

// AlGetSourcei function as declared in AL/al.h:522
func AlGetSourcei(source ALuint, param ALenum, value []ALint) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value)).Data)), cgoAllocsUnknown
	C.alGetSourcei(csource, cparam, cvalue)
}

// AlGetSource3i function as declared in AL/al.h:523
func AlGetSource3i(source ALuint, param ALenum, value1 []ALint, value2 []ALint, value3 []ALint) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue1, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value1)).Data)), cgoAllocsUnknown
	cvalue2, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value2)).Data)), cgoAllocsUnknown
	cvalue3, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value3)).Data)), cgoAllocsUnknown
	C.alGetSource3i(csource, cparam, cvalue1, cvalue2, cvalue3)
}

// AlGetSourceiv function as declared in AL/al.h:524
func AlGetSourceiv(source ALuint, param ALenum, values []ALint) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alGetSourceiv(csource, cparam, cvalues)
}

// AlSourcePlayv function as declared in AL/al.h:528
func AlSourcePlayv(n ALsizei, sources []ALuint) {
	cn, _ := (C.ALsizei)(n), cgoAllocsUnknown
	csources, _ := (*C.ALuint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&sources)).Data)), cgoAllocsUnknown
	C.alSourcePlayv(cn, csources)
}

// AlSourceStopv function as declared in AL/al.h:530
func AlSourceStopv(n ALsizei, sources []ALuint) {
	cn, _ := (C.ALsizei)(n), cgoAllocsUnknown
	csources, _ := (*C.ALuint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&sources)).Data)), cgoAllocsUnknown
	C.alSourceStopv(cn, csources)
}

// AlSourceRewindv function as declared in AL/al.h:532
func AlSourceRewindv(n ALsizei, sources []ALuint) {
	cn, _ := (C.ALsizei)(n), cgoAllocsUnknown
	csources, _ := (*C.ALuint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&sources)).Data)), cgoAllocsUnknown
	C.alSourceRewindv(cn, csources)
}

// AlSourcePausev function as declared in AL/al.h:534
func AlSourcePausev(n ALsizei, sources []ALuint) {
	cn, _ := (C.ALsizei)(n), cgoAllocsUnknown
	csources, _ := (*C.ALuint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&sources)).Data)), cgoAllocsUnknown
	C.alSourcePausev(cn, csources)
}

// AlSourcePlay function as declared in AL/al.h:537
func AlSourcePlay(source ALuint) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	C.alSourcePlay(csource)
}

// AlSourceStop function as declared in AL/al.h:539
func AlSourceStop(source ALuint) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	C.alSourceStop(csource)
}

// AlSourceRewind function as declared in AL/al.h:541
func AlSourceRewind(source ALuint) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	C.alSourceRewind(csource)
}

// AlSourcePause function as declared in AL/al.h:543
func AlSourcePause(source ALuint) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	C.alSourcePause(csource)
}

// AlSourceQueueBuffers function as declared in AL/al.h:546
func AlSourceQueueBuffers(source ALuint, nb ALsizei, buffers []ALuint) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cnb, _ := (C.ALsizei)(nb), cgoAllocsUnknown
	cbuffers, _ := (*C.ALuint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffers)).Data)), cgoAllocsUnknown
	C.alSourceQueueBuffers(csource, cnb, cbuffers)
}

// AlSourceUnqueueBuffers function as declared in AL/al.h:548
func AlSourceUnqueueBuffers(source ALuint, nb ALsizei, buffers []ALuint) {
	csource, _ := (C.ALuint)(source), cgoAllocsUnknown
	cnb, _ := (C.ALsizei)(nb), cgoAllocsUnknown
	cbuffers, _ := (*C.ALuint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffers)).Data)), cgoAllocsUnknown
	C.alSourceUnqueueBuffers(csource, cnb, cbuffers)
}

// AlGenBuffers function as declared in AL/al.h:552
func AlGenBuffers(n ALsizei, buffers []ALuint) {
	cn, _ := (C.ALsizei)(n), cgoAllocsUnknown
	cbuffers, _ := (*C.ALuint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffers)).Data)), cgoAllocsUnknown
	C.alGenBuffers(cn, cbuffers)
}

// AlDeleteBuffers function as declared in AL/al.h:554
func AlDeleteBuffers(n ALsizei, buffers []ALuint) {
	cn, _ := (C.ALsizei)(n), cgoAllocsUnknown
	cbuffers, _ := (*C.ALuint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffers)).Data)), cgoAllocsUnknown
	C.alDeleteBuffers(cn, cbuffers)
}

// AlIsBuffer function as declared in AL/al.h:556
func AlIsBuffer(buffer ALuint) ALboolean {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	__ret := C.alIsBuffer(cbuffer)
	__v := (ALboolean)(__ret)
	return __v
}

// AlBufferData function as declared in AL/al.h:559
func AlBufferData(buffer ALuint, format ALenum, data *ALvoid, size ALsizei, freq ALsizei) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cformat, _ := (C.ALenum)(format), cgoAllocsUnknown
	cdata, _ := unsafe.Pointer(data), cgoAllocsUnknown
	csize, _ := (C.ALsizei)(size), cgoAllocsUnknown
	cfreq, _ := (C.ALsizei)(freq), cgoAllocsUnknown
	C.alBufferData(cbuffer, cformat, cdata, csize, cfreq)
}

// AlBufferf function as declared in AL/al.h:562
func AlBufferf(buffer ALuint, param ALenum, value ALfloat) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue, _ := (C.ALfloat)(value), cgoAllocsUnknown
	C.alBufferf(cbuffer, cparam, cvalue)
}

// AlBuffer3f function as declared in AL/al.h:563
func AlBuffer3f(buffer ALuint, param ALenum, value1 ALfloat, value2 ALfloat, value3 ALfloat) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue1, _ := (C.ALfloat)(value1), cgoAllocsUnknown
	cvalue2, _ := (C.ALfloat)(value2), cgoAllocsUnknown
	cvalue3, _ := (C.ALfloat)(value3), cgoAllocsUnknown
	C.alBuffer3f(cbuffer, cparam, cvalue1, cvalue2, cvalue3)
}

// AlBufferfv function as declared in AL/al.h:564
func AlBufferfv(buffer ALuint, param ALenum, values []ALfloat) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alBufferfv(cbuffer, cparam, cvalues)
}

// AlBufferi function as declared in AL/al.h:565
func AlBufferi(buffer ALuint, param ALenum, value ALint) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue, _ := (C.ALint)(value), cgoAllocsUnknown
	C.alBufferi(cbuffer, cparam, cvalue)
}

// AlBuffer3i function as declared in AL/al.h:566
func AlBuffer3i(buffer ALuint, param ALenum, value1 ALint, value2 ALint, value3 ALint) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue1, _ := (C.ALint)(value1), cgoAllocsUnknown
	cvalue2, _ := (C.ALint)(value2), cgoAllocsUnknown
	cvalue3, _ := (C.ALint)(value3), cgoAllocsUnknown
	C.alBuffer3i(cbuffer, cparam, cvalue1, cvalue2, cvalue3)
}

// AlBufferiv function as declared in AL/al.h:567
func AlBufferiv(buffer ALuint, param ALenum, values []ALint) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alBufferiv(cbuffer, cparam, cvalues)
}

// AlGetBufferf function as declared in AL/al.h:570
func AlGetBufferf(buffer ALuint, param ALenum, value []ALfloat) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value)).Data)), cgoAllocsUnknown
	C.alGetBufferf(cbuffer, cparam, cvalue)
}

// AlGetBuffer3f function as declared in AL/al.h:571
func AlGetBuffer3f(buffer ALuint, param ALenum, value1 []ALfloat, value2 []ALfloat, value3 []ALfloat) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue1, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value1)).Data)), cgoAllocsUnknown
	cvalue2, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value2)).Data)), cgoAllocsUnknown
	cvalue3, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value3)).Data)), cgoAllocsUnknown
	C.alGetBuffer3f(cbuffer, cparam, cvalue1, cvalue2, cvalue3)
}

// AlGetBufferfv function as declared in AL/al.h:572
func AlGetBufferfv(buffer ALuint, param ALenum, values []ALfloat) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALfloat)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alGetBufferfv(cbuffer, cparam, cvalues)
}

// AlGetBufferi function as declared in AL/al.h:573
func AlGetBufferi(buffer ALuint, param ALenum, value []ALint) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value)).Data)), cgoAllocsUnknown
	C.alGetBufferi(cbuffer, cparam, cvalue)
}

// AlGetBuffer3i function as declared in AL/al.h:574
func AlGetBuffer3i(buffer ALuint, param ALenum, value1 []ALint, value2 []ALint, value3 []ALint) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalue1, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value1)).Data)), cgoAllocsUnknown
	cvalue2, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value2)).Data)), cgoAllocsUnknown
	cvalue3, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value3)).Data)), cgoAllocsUnknown
	C.alGetBuffer3i(cbuffer, cparam, cvalue1, cvalue2, cvalue3)
}

// AlGetBufferiv function as declared in AL/al.h:575
func AlGetBufferiv(buffer ALuint, param ALenum, values []ALint) {
	cbuffer, _ := (C.ALuint)(buffer), cgoAllocsUnknown
	cparam, _ := (C.ALenum)(param), cgoAllocsUnknown
	cvalues, _ := (*C.ALint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alGetBufferiv(cbuffer, cparam, cvalues)
}

// AlcCreateContext function as declared in AL/alc.h:170
func AlcCreateContext(device []ALCdevice, attrlist []ALCint) *ALCcontext {
	cdevice, _ := (*C.ALCdevice)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&device)).Data)), cgoAllocsUnknown
	cattrlist, _ := (*C.ALCint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&attrlist)).Data)), cgoAllocsUnknown
	__ret := C.alcCreateContext(cdevice, cattrlist)
	__v := *(**ALCcontext)(unsafe.Pointer(&__ret))
	return __v
}

// AlcMakeContextCurrent function as declared in AL/alc.h:171
func AlcMakeContextCurrent(context []ALCcontext) ALCboolean {
	ccontext, _ := (*C.ALCcontext)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&context)).Data)), cgoAllocsUnknown
	__ret := C.alcMakeContextCurrent(ccontext)
	__v := (ALCboolean)(__ret)
	return __v
}

// AlcProcessContext function as declared in AL/alc.h:172
func AlcProcessContext(context []ALCcontext) {
	ccontext, _ := (*C.ALCcontext)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&context)).Data)), cgoAllocsUnknown
	C.alcProcessContext(ccontext)
}

// AlcSuspendContext function as declared in AL/alc.h:173
func AlcSuspendContext(context []ALCcontext) {
	ccontext, _ := (*C.ALCcontext)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&context)).Data)), cgoAllocsUnknown
	C.alcSuspendContext(ccontext)
}

// AlcDestroyContext function as declared in AL/alc.h:174
func AlcDestroyContext(context []ALCcontext) {
	ccontext, _ := (*C.ALCcontext)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&context)).Data)), cgoAllocsUnknown
	C.alcDestroyContext(ccontext)
}

// AlcGetCurrentContext function as declared in AL/alc.h:175
func AlcGetCurrentContext() *ALCcontext {
	__ret := C.alcGetCurrentContext()
	__v := *(**ALCcontext)(unsafe.Pointer(&__ret))
	return __v
}

// AlcGetContextsDevice function as declared in AL/alc.h:176
func AlcGetContextsDevice(context []ALCcontext) *ALCdevice {
	ccontext, _ := (*C.ALCcontext)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&context)).Data)), cgoAllocsUnknown
	__ret := C.alcGetContextsDevice(ccontext)
	__v := *(**ALCdevice)(unsafe.Pointer(&__ret))
	return __v
}

// AlcOpenDevice function as declared in AL/alc.h:179
func AlcOpenDevice(devicename []ALCchar) *ALCdevice {
	cdevicename, _ := unpackArgSALCchar(devicename)
	__ret := C.alcOpenDevice(cdevicename)
	packSALCchar(devicename, cdevicename)
	__v := *(**ALCdevice)(unsafe.Pointer(&__ret))
	return __v
}

// AlcCloseDevice function as declared in AL/alc.h:180
func AlcCloseDevice(device []ALCdevice) ALCboolean {
	cdevice, _ := (*C.ALCdevice)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&device)).Data)), cgoAllocsUnknown
	__ret := C.alcCloseDevice(cdevice)
	__v := (ALCboolean)(__ret)
	return __v
}

// AlcGetError function as declared in AL/alc.h:188
func AlcGetError(device []ALCdevice) ALCenum {
	cdevice, _ := (*C.ALCdevice)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&device)).Data)), cgoAllocsUnknown
	__ret := C.alcGetError(cdevice)
	__v := (ALCenum)(__ret)
	return __v
}

// AlcIsExtensionPresent function as declared in AL/alc.h:196
func AlcIsExtensionPresent(device []ALCdevice, extname []ALCchar) ALCboolean {
	cdevice, _ := (*C.ALCdevice)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&device)).Data)), cgoAllocsUnknown
	cextname, _ := unpackArgSALCchar(extname)
	__ret := C.alcIsExtensionPresent(cdevice, cextname)
	packSALCchar(extname, cextname)
	__v := (ALCboolean)(__ret)
	return __v
}

// AlcGetProcAddress function as declared in AL/alc.h:197
func AlcGetProcAddress(device []ALCdevice, funcname []ALCchar) unsafe.Pointer {
	cdevice, _ := (*C.ALCdevice)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&device)).Data)), cgoAllocsUnknown
	cfuncname, _ := unpackArgSALCchar(funcname)
	__ret := C.alcGetProcAddress(cdevice, cfuncname)
	packSALCchar(funcname, cfuncname)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// AlcGetEnumValue function as declared in AL/alc.h:198
func AlcGetEnumValue(device []ALCdevice, enumname []ALCchar) ALCenum {
	cdevice, _ := (*C.ALCdevice)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&device)).Data)), cgoAllocsUnknown
	cenumname, _ := unpackArgSALCchar(enumname)
	__ret := C.alcGetEnumValue(cdevice, cenumname)
	packSALCchar(enumname, cenumname)
	__v := (ALCenum)(__ret)
	return __v
}

// AlcGetString function as declared in AL/alc.h:201
// REMOVED by DORIAN due to incompatibility
// func AlcGetString(device []ALCdevice, param ALCenum) *ALCchar {
// 	cdevice, _ := (*C.ALCdevice)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&device)).Data)), cgoAllocsUnknown
// 	cparam, _ := (C.ALCenum)(param), cgoAllocsUnknown
// 	__ret := C.alcGetString(cdevice, cparam)
// 	__v := packPALCcharString(__ret)
// 	return __v
// }

// AlcGetIntegerv function as declared in AL/alc.h:202
func AlcGetIntegerv(device []ALCdevice, param ALCenum, size ALCsizei, values []ALCint) {
	cdevice, _ := (*C.ALCdevice)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&device)).Data)), cgoAllocsUnknown
	cparam, _ := (C.ALCenum)(param), cgoAllocsUnknown
	csize, _ := (C.ALCsizei)(size), cgoAllocsUnknown
	cvalues, _ := (*C.ALCint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&values)).Data)), cgoAllocsUnknown
	C.alcGetIntegerv(cdevice, cparam, csize, cvalues)
}

// AlcCaptureOpenDevice function as declared in AL/alc.h:205
func AlcCaptureOpenDevice(devicename []ALCchar, frequency ALCuint, format ALCenum, buffersize ALCsizei) *ALCdevice {
	cdevicename, _ := unpackArgSALCchar(devicename)
	cfrequency, _ := (C.ALCuint)(frequency), cgoAllocsUnknown
	cformat, _ := (C.ALCenum)(format), cgoAllocsUnknown
	cbuffersize, _ := (C.ALCsizei)(buffersize), cgoAllocsUnknown
	__ret := C.alcCaptureOpenDevice(cdevicename, cfrequency, cformat, cbuffersize)
	packSALCchar(devicename, cdevicename)
	__v := *(**ALCdevice)(unsafe.Pointer(&__ret))
	return __v
}

// AlcCaptureCloseDevice function as declared in AL/alc.h:206
func AlcCaptureCloseDevice(device []ALCdevice) ALCboolean {
	cdevice, _ := (*C.ALCdevice)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&device)).Data)), cgoAllocsUnknown
	__ret := C.alcCaptureCloseDevice(cdevice)
	__v := (ALCboolean)(__ret)
	return __v
}

// AlcCaptureStart function as declared in AL/alc.h:207
func AlcCaptureStart(device []ALCdevice) {
	cdevice, _ := (*C.ALCdevice)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&device)).Data)), cgoAllocsUnknown
	C.alcCaptureStart(cdevice)
}

// AlcCaptureStop function as declared in AL/alc.h:208
func AlcCaptureStop(device []ALCdevice) {
	cdevice, _ := (*C.ALCdevice)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&device)).Data)), cgoAllocsUnknown
	C.alcCaptureStop(cdevice)
}

// AlcCaptureSamples function as declared in AL/alc.h:209
func AlcCaptureSamples(device []ALCdevice, buffer *ALCvoid, samples ALCsizei) {
	cdevice, _ := (*C.ALCdevice)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&device)).Data)), cgoAllocsUnknown
	cbuffer, _ := unsafe.Pointer(buffer), cgoAllocsUnknown
	csamples, _ := (C.ALCsizei)(samples), cgoAllocsUnknown
	C.alcCaptureSamples(cdevice, cbuffer, csamples)
}
