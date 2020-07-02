// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY BELGIAN ROBOTS.

// WARNING: This file has automatically been generated on Wed, 1 Jul 2020 12:00:25 MDT.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package al

/*
#cgo pkg-config: openal
#include "AL/al.h"
#include "AL/alc.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// __const as defined in project_openal_go/<predefine>:8
	__const = 0
	// __STDC_HOSTED__ as defined in project_openal_go/<predefine>:24
	__STDC_HOSTED__ = 1
	// __STDC_VERSION__ as defined in project_openal_go/<predefine>:25
	__STDC_VERSION__ = 199901
	// __STDC__ as defined in project_openal_go/<predefine>:26
	__STDC__ = 1
	// __GNUC__ as defined in project_openal_go/<predefine>:27
	__GNUC__ = 4
	// __FLT_MIN__ as defined in project_openal_go/<predefine>:32
	__FLT_MIN__ = 0
	// __DBL_MIN__ as defined in project_openal_go/<predefine>:33
	__DBL_MIN__ = 0
	// __LDBL_MIN__ as defined in project_openal_go/<predefine>:34
	__LDBL_MIN__ = 0
	// __x86_64__ as defined in project_openal_go/<predefine>:38
	__x86_64__ = 1
	// AL_API as defined in AL/al.h:14
	AL_API = 0
	// ALAPI as defined in AL/al.h:27
	ALAPI = 0
	// ALAPIENTRY as defined in AL/al.h:28
	// ALAPIENTRY = AL_APIENTRY
	// AL_INVALID as defined in AL/al.h:29
	AL_INVALID = -1
	// AL_ILLEGAL_ENUM as defined in AL/al.h:30
	AL_ILLEGAL_ENUM = 40962
	// AL_ILLEGAL_COMMAND as defined in AL/al.h:31
	AL_ILLEGAL_COMMAND = 40964
	// AL_NONE as defined in AL/al.h:80
	AL_NONE = 0
	// AL_FALSE as defined in AL/al.h:83
	AL_FALSE = 0
	// AL_TRUE as defined in AL/al.h:86
	AL_TRUE = 1
	// AL_SOURCE_RELATIVE as defined in AL/al.h:97
	AL_SOURCE_RELATIVE = 514
	// AL_CONE_INNER_ANGLE as defined in AL/al.h:108
	AL_CONE_INNER_ANGLE = 4097
	// AL_CONE_OUTER_ANGLE as defined in AL/al.h:118
	AL_CONE_OUTER_ANGLE = 4098
	// AL_PITCH as defined in AL/al.h:128
	AL_PITCH = 4099
	// AL_POSITION as defined in AL/al.h:144
	AL_POSITION = 4100
	// AL_DIRECTION as defined in AL/al.h:154
	AL_DIRECTION = 4101
	// AL_VELOCITY as defined in AL/al.h:163
	AL_VELOCITY = 4102
	// AL_LOOPING as defined in AL/al.h:173
	AL_LOOPING = 4103
	// AL_BUFFER as defined in AL/al.h:182
	AL_BUFFER = 4105
	// AL_GAIN as defined in AL/al.h:196
	AL_GAIN = 4106
	// AL_MIN_GAIN as defined in AL/al.h:206
	AL_MIN_GAIN = 4109
	// AL_MAX_GAIN as defined in AL/al.h:216
	AL_MAX_GAIN = 4110
	// AL_ORIENTATION as defined in AL/al.h:228
	AL_ORIENTATION = 4111
	// AL_SOURCE_STATE as defined in AL/al.h:235
	AL_SOURCE_STATE = 4112
	// AL_INITIAL as defined in AL/al.h:238
	AL_INITIAL = 4113
	// AL_PLAYING as defined in AL/al.h:239
	AL_PLAYING = 4114
	// AL_PAUSED as defined in AL/al.h:240
	AL_PAUSED = 4115
	// AL_STOPPED as defined in AL/al.h:241
	AL_STOPPED = 4116
	// AL_BUFFERS_QUEUED as defined in AL/al.h:250
	AL_BUFFERS_QUEUED = 4117
	// AL_BUFFERS_PROCESSED as defined in AL/al.h:262
	AL_BUFFERS_PROCESSED = 4118
	// AL_REFERENCE_DISTANCE as defined in AL/al.h:274
	AL_REFERENCE_DISTANCE = 4128
	// AL_ROLLOFF_FACTOR as defined in AL/al.h:286
	AL_ROLLOFF_FACTOR = 4129
	// AL_CONE_OUTER_GAIN as defined in AL/al.h:297
	AL_CONE_OUTER_GAIN = 4130
	// AL_MAX_DISTANCE as defined in AL/al.h:309
	AL_MAX_DISTANCE = 4131
	// AL_SEC_OFFSET as defined in AL/al.h:312
	AL_SEC_OFFSET = 4132
	// AL_SAMPLE_OFFSET as defined in AL/al.h:314
	AL_SAMPLE_OFFSET = 4133
	// AL_BYTE_OFFSET as defined in AL/al.h:316
	AL_BYTE_OFFSET = 4134
	// AL_SOURCE_TYPE as defined in AL/al.h:331
	AL_SOURCE_TYPE = 4135
	// AL_STATIC as defined in AL/al.h:334
	AL_STATIC = 4136
	// AL_STREAMING as defined in AL/al.h:335
	AL_STREAMING = 4137
	// AL_UNDETERMINED as defined in AL/al.h:336
	AL_UNDETERMINED = 4144
	// AL_FORMAT_MONO8 as defined in AL/al.h:339
	AL_FORMAT_MONO8 = 4352
	// AL_FORMAT_MONO16 as defined in AL/al.h:340
	AL_FORMAT_MONO16 = 4353
	// AL_FORMAT_STEREO8 as defined in AL/al.h:341
	AL_FORMAT_STEREO8 = 4354
	// AL_FORMAT_STEREO16 as defined in AL/al.h:342
	AL_FORMAT_STEREO16 = 4355
	// AL_FREQUENCY as defined in AL/al.h:345
	AL_FREQUENCY = 8193
	// AL_BITS as defined in AL/al.h:347
	AL_BITS = 8194
	// AL_CHANNELS as defined in AL/al.h:349
	AL_CHANNELS = 8195
	// AL_SIZE as defined in AL/al.h:351
	AL_SIZE = 8196
	// AL_UNUSED as defined in AL/al.h:358
	AL_UNUSED = 8208
	// AL_PENDING as defined in AL/al.h:359
	AL_PENDING = 8209
	// AL_PROCESSED as defined in AL/al.h:360
	AL_PROCESSED = 8210
	// AL_NO_ERROR as defined in AL/al.h:364
	AL_NO_ERROR = 0
	// AL_INVALID_NAME as defined in AL/al.h:367
	AL_INVALID_NAME = 40961
	// AL_INVALID_ENUM as defined in AL/al.h:370
	AL_INVALID_ENUM = 40962
	// AL_INVALID_VALUE as defined in AL/al.h:373
	AL_INVALID_VALUE = 40963
	// AL_INVALID_OPERATION as defined in AL/al.h:376
	AL_INVALID_OPERATION = 40964
	// AL_OUT_OF_MEMORY as defined in AL/al.h:379
	AL_OUT_OF_MEMORY = 40965
	// AL_VENDOR as defined in AL/al.h:383
	AL_VENDOR = 45057
	// AL_VERSION as defined in AL/al.h:385
	AL_VERSION = 45058
	// AL_RENDERER as defined in AL/al.h:387
	AL_RENDERER = 45059
	// AL_EXTENSIONS as defined in AL/al.h:389
	AL_EXTENSIONS = 45060
	// AL_DOPPLER_FACTOR as defined in AL/al.h:400
	AL_DOPPLER_FACTOR = 49152
	// AL_DOPPLER_VELOCITY as defined in AL/al.h:408
	AL_DOPPLER_VELOCITY = 49153
	// AL_SPEED_OF_SOUND as defined in AL/al.h:420
	AL_SPEED_OF_SOUND = 49155
	// AL_DISTANCE_MODEL as defined in AL/al.h:441
	AL_DISTANCE_MODEL = 53248
	// AL_INVERSE_DISTANCE as defined in AL/al.h:445
	AL_INVERSE_DISTANCE = 53249
	// AL_INVERSE_DISTANCE_CLAMPED as defined in AL/al.h:446
	AL_INVERSE_DISTANCE_CLAMPED = 53250
	// AL_LINEAR_DISTANCE as defined in AL/al.h:447
	AL_LINEAR_DISTANCE = 53251
	// AL_LINEAR_DISTANCE_CLAMPED as defined in AL/al.h:448
	AL_LINEAR_DISTANCE_CLAMPED = 53252
	// AL_EXPONENT_DISTANCE as defined in AL/al.h:449
	AL_EXPONENT_DISTANCE = 53253
	// AL_EXPONENT_DISTANCE_CLAMPED as defined in AL/al.h:450
	AL_EXPONENT_DISTANCE_CLAMPED = 53254
	// ALC_API as defined in AL/alc.h:14
	ALC_API = 0
	// ALCAPI as defined in AL/alc.h:26
	ALCAPI = 0
	// ALCAPIENTRY as defined in AL/alc.h:27
	// ALCAPIENTRY = ALC_APIENTRY
	// ALC_INVALID as defined in AL/alc.h:28
	ALC_INVALID = 0
	// ALC_VERSION_0_1 as defined in AL/alc.h:31
	ALC_VERSION_0_1 = 1
	// ALC_FALSE as defined in AL/alc.h:81
	ALC_FALSE = 0
	// ALC_TRUE as defined in AL/alc.h:84
	ALC_TRUE = 1
	// ALC_FREQUENCY as defined in AL/alc.h:87
	ALC_FREQUENCY = 4103
	// ALC_REFRESH as defined in AL/alc.h:90
	ALC_REFRESH = 4104
	// ALC_SYNC as defined in AL/alc.h:93
	ALC_SYNC = 4105
	// ALC_MONO_SOURCES as defined in AL/alc.h:96
	ALC_MONO_SOURCES = 4112
	// ALC_STEREO_SOURCES as defined in AL/alc.h:99
	ALC_STEREO_SOURCES = 4113
	// ALC_NO_ERROR as defined in AL/alc.h:102
	ALC_NO_ERROR = 0
	// ALC_INVALID_DEVICE as defined in AL/alc.h:105
	ALC_INVALID_DEVICE = 40961
	// ALC_INVALID_CONTEXT as defined in AL/alc.h:108
	ALC_INVALID_CONTEXT = 40962
	// ALC_INVALID_ENUM as defined in AL/alc.h:111
	ALC_INVALID_ENUM = 40963
	// ALC_INVALID_VALUE as defined in AL/alc.h:114
	ALC_INVALID_VALUE = 40964
	// ALC_OUT_OF_MEMORY as defined in AL/alc.h:117
	ALC_OUT_OF_MEMORY = 40965
	// ALC_MAJOR_VERSION as defined in AL/alc.h:121
	ALC_MAJOR_VERSION = 4096
	// ALC_MINOR_VERSION as defined in AL/alc.h:122
	ALC_MINOR_VERSION = 4097
	// ALC_ATTRIBUTES_SIZE as defined in AL/alc.h:125
	ALC_ATTRIBUTES_SIZE = 4098
	// ALC_ALL_ATTRIBUTES as defined in AL/alc.h:126
	ALC_ALL_ATTRIBUTES = 4099
	// ALC_DEFAULT_DEVICE_SPECIFIER as defined in AL/alc.h:129
	ALC_DEFAULT_DEVICE_SPECIFIER = 4100
	// ALC_DEVICE_SPECIFIER as defined in AL/alc.h:136
	ALC_DEVICE_SPECIFIER = 4101
	// ALC_EXTENSIONS as defined in AL/alc.h:138
	ALC_EXTENSIONS = 4102
	// ALC_EXT_CAPTURE as defined in AL/alc.h:142
	ALC_EXT_CAPTURE = 1
	// ALC_CAPTURE_DEVICE_SPECIFIER as defined in AL/alc.h:149
	ALC_CAPTURE_DEVICE_SPECIFIER = 784
	// ALC_CAPTURE_DEFAULT_DEVICE_SPECIFIER as defined in AL/alc.h:151
	ALC_CAPTURE_DEFAULT_DEVICE_SPECIFIER = 785
	// ALC_CAPTURE_SAMPLES as defined in AL/alc.h:153
	ALC_CAPTURE_SAMPLES = 786
	// ALC_ENUMERATE_ALL_EXT as defined in AL/alc.h:157
	ALC_ENUMERATE_ALL_EXT = 1
	// ALC_DEFAULT_ALL_DEVICES_SPECIFIER as defined in AL/alc.h:159
	ALC_DEFAULT_ALL_DEVICES_SPECIFIER = 4114
	// ALC_ALL_DEVICES_SPECIFIER as defined in AL/alc.h:166
	ALC_ALL_DEVICES_SPECIFIER = 4115
)
