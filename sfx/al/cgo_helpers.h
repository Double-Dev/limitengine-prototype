// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY BELGIAN ROBOTS.

// WARNING: This file has automatically been generated on Wed, 1 Jul 2020 12:00:25 MDT.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

#include "AL/al.h"
#include "AL/alc.h"
#include <stdlib.h>
#pragma once

#define __CGOGEN 1

// LPALENABLE_c158bd47 is a proxy for callback LPALENABLE.
void LPALENABLE_c158bd47(int capability);

// LPALDISABLE_5476105f is a proxy for callback LPALDISABLE.
void LPALDISABLE_5476105f(int capability);

// LPALISENABLED_dc30ce46 is a proxy for callback LPALISENABLED.
char LPALISENABLED_dc30ce46(int capability);

// LPALGETSTRING_54e8aeb0 is a proxy for callback LPALGETSTRING.
char* LPALGETSTRING_54e8aeb0(int param);

// LPALGETBOOLEANV_96c6c0d is a proxy for callback LPALGETBOOLEANV.
void LPALGETBOOLEANV_96c6c0d(int param, char* values);

// LPALGETINTEGERV_dd07d6cc is a proxy for callback LPALGETINTEGERV.
void LPALGETINTEGERV_dd07d6cc(int param, int* values);

// LPALGETFLOATV_219ee508 is a proxy for callback LPALGETFLOATV.
void LPALGETFLOATV_219ee508(int param, float* values);

// LPALGETDOUBLEV_4c4fd8b is a proxy for callback LPALGETDOUBLEV.
void LPALGETDOUBLEV_4c4fd8b(int param, double* values);

// LPALGETBOOLEAN_d54a567c is a proxy for callback LPALGETBOOLEAN.
char LPALGETBOOLEAN_d54a567c(int param);

// LPALGETINTEGER_5c0751f is a proxy for callback LPALGETINTEGER.
int LPALGETINTEGER_5c0751f(int param);

// LPALGETFLOAT_257bcd44 is a proxy for callback LPALGETFLOAT.
float LPALGETFLOAT_257bcd44(int param);

// LPALGETDOUBLE_10b1eef6 is a proxy for callback LPALGETDOUBLE.
double LPALGETDOUBLE_10b1eef6(int param);

// LPALGETERROR_b1032fa0 is a proxy for callback LPALGETERROR.
int LPALGETERROR_b1032fa0();

// LPALISEXTENSIONPRESENT_2de7d519 is a proxy for callback LPALISEXTENSIONPRESENT.
char LPALISEXTENSIONPRESENT_2de7d519(char* extname);

// LPALGETPROCADDRESS_f3ca5dca is a proxy for callback LPALGETPROCADDRESS.
void* LPALGETPROCADDRESS_f3ca5dca(char* fname);

// LPALGETENUMVALUE_5dc5dce is a proxy for callback LPALGETENUMVALUE.
int LPALGETENUMVALUE_5dc5dce(char* ename);

// LPALLISTENERF_8695e9ae is a proxy for callback LPALLISTENERF.
void LPALLISTENERF_8695e9ae(int param, float value);

// LPALLISTENER3F_2e1d9d9d is a proxy for callback LPALLISTENER3F.
void LPALLISTENER3F_2e1d9d9d(int param, float value1, float value2, float value3);

// LPALLISTENERFV_61e2004a is a proxy for callback LPALLISTENERFV.
void LPALLISTENERFV_61e2004a(int param, float* values);

// LPALLISTENERI_162af43f is a proxy for callback LPALLISTENERI.
void LPALLISTENERI_162af43f(int param, int value);

// LPALLISTENER3I_bea2800c is a proxy for callback LPALLISTENER3I.
void LPALLISTENER3I_bea2800c(int param, int value1, int value2, int value3);

// LPALLISTENERIV_e67a1c85 is a proxy for callback LPALLISTENERIV.
void LPALLISTENERIV_e67a1c85(int param, int* values);

// LPALGETLISTENERF_42098513 is a proxy for callback LPALGETLISTENERF.
void LPALGETLISTENERF_42098513(int param, float* value);

// LPALGETLISTENER3F_9b09cec0 is a proxy for callback LPALGETLISTENER3F.
void LPALGETLISTENER3F_9b09cec0(int param, float* value1, float* value2, float* value3);

// LPALGETLISTENERFV_d4f65317 is a proxy for callback LPALGETLISTENERFV.
void LPALGETLISTENERFV_d4f65317(int param, float* values);

// LPALGETLISTENERI_d2b69882 is a proxy for callback LPALGETLISTENERI.
void LPALGETLISTENERI_d2b69882(int param, int* value);

// LPALGETLISTENER3I_bb6d351 is a proxy for callback LPALGETLISTENER3I.
void LPALGETLISTENER3I_bb6d351(int param, int* value1, int* value2, int* value3);

// LPALGETLISTENERIV_536e4fd8 is a proxy for callback LPALGETLISTENERIV.
void LPALGETLISTENERIV_536e4fd8(int param, int* values);

// LPALGENSOURCES_60e99605 is a proxy for callback LPALGENSOURCES.
void LPALGENSOURCES_60e99605(int n, unsigned int* sources);

// LPALDELETESOURCES_9e261641 is a proxy for callback LPALDELETESOURCES.
void LPALDELETESOURCES_9e261641(int n, unsigned int* sources);

// LPALISSOURCE_45834e0d is a proxy for callback LPALISSOURCE.
char LPALISSOURCE_45834e0d(unsigned int source);

// LPALSOURCEF_a4b3dc17 is a proxy for callback LPALSOURCEF.
void LPALSOURCEF_a4b3dc17(unsigned int source, int param, float value);

// LPALSOURCE3F_9c82b080 is a proxy for callback LPALSOURCE3F.
void LPALSOURCE3F_9c82b080(unsigned int source, int param, float value1, float value2, float value3);

// LPALSOURCEFV_d37d2d57 is a proxy for callback LPALSOURCEFV.
void LPALSOURCEFV_d37d2d57(unsigned int source, int param, float* values);

// LPALSOURCEI_340cc186 is a proxy for callback LPALSOURCEI.
void LPALSOURCEI_340cc186(unsigned int source, int param, int value);

// LPALSOURCE3I_c3dad11 is a proxy for callback LPALSOURCE3I.
void LPALSOURCE3I_c3dad11(unsigned int source, int param, int value1, int value2, int value3);

// LPALSOURCEIV_54e53198 is a proxy for callback LPALSOURCEIV.
void LPALSOURCEIV_54e53198(unsigned int source, int param, int* values);

// LPALGETSOURCEF_e04f5f0d is a proxy for callback LPALGETSOURCEF.
void LPALGETSOURCEF_e04f5f0d(unsigned int source, int param, float* value);

// LPALGETSOURCE3F_61a4b579 is a proxy for callback LPALGETSOURCE3F.
void LPALGETSOURCE3F_61a4b579(unsigned int source, int param, float* value1, float* value2, float* value3);

// LPALGETSOURCEFV_2e5b28ae is a proxy for callback LPALGETSOURCEFV.
void LPALGETSOURCEFV_2e5b28ae(unsigned int source, int param, float* values);

// LPALGETSOURCEI_70f0429c is a proxy for callback LPALGETSOURCEI.
void LPALGETSOURCEI_70f0429c(unsigned int source, int param, int* value);

// LPALGETSOURCE3I_f11ba8e8 is a proxy for callback LPALGETSOURCE3I.
void LPALGETSOURCE3I_f11ba8e8(unsigned int source, int param, int* value1, int* value2, int* value3);

// LPALGETSOURCEIV_a9c33461 is a proxy for callback LPALGETSOURCEIV.
void LPALGETSOURCEIV_a9c33461(unsigned int source, int param, int* values);

// LPALSOURCEPLAYV_d0019668 is a proxy for callback LPALSOURCEPLAYV.
void LPALSOURCEPLAYV_d0019668(int n, unsigned int* sources);

// LPALSOURCESTOPV_d950058b is a proxy for callback LPALSOURCESTOPV.
void LPALSOURCESTOPV_d950058b(int n, unsigned int* sources);

// LPALSOURCEREWINDV_fe90ec74 is a proxy for callback LPALSOURCEREWINDV.
void LPALSOURCEREWINDV_fe90ec74(int n, unsigned int* sources);

// LPALSOURCEPAUSEV_5820338d is a proxy for callback LPALSOURCEPAUSEV.
void LPALSOURCEPAUSEV_5820338d(int n, unsigned int* sources);

// LPALSOURCEPLAY_6eeaef95 is a proxy for callback LPALSOURCEPLAY.
void LPALSOURCEPLAY_6eeaef95(unsigned int source);

// LPALSOURCESTOP_89352799 is a proxy for callback LPALSOURCESTOP.
void LPALSOURCESTOP_89352799(unsigned int source);

// LPALSOURCEREWIND_4b9d72e8 is a proxy for callback LPALSOURCEREWIND.
void LPALSOURCEREWIND_4b9d72e8(unsigned int source);

// LPALSOURCEPAUSE_4715ec4d is a proxy for callback LPALSOURCEPAUSE.
void LPALSOURCEPAUSE_4715ec4d(unsigned int source);

// LPALSOURCEQUEUEBUFFERS_4c6e4ca3 is a proxy for callback LPALSOURCEQUEUEBUFFERS.
void LPALSOURCEQUEUEBUFFERS_4c6e4ca3(unsigned int source, int nb, unsigned int* buffers);

// LPALSOURCEUNQUEUEBUFFERS_c16dec16 is a proxy for callback LPALSOURCEUNQUEUEBUFFERS.
void LPALSOURCEUNQUEUEBUFFERS_c16dec16(unsigned int source, int nb, unsigned int* buffers);

// LPALGENBUFFERS_fa38f53c is a proxy for callback LPALGENBUFFERS.
void LPALGENBUFFERS_fa38f53c(int n, unsigned int* buffers);

// LPALDELETEBUFFERS_4f77578 is a proxy for callback LPALDELETEBUFFERS.
void LPALDELETEBUFFERS_4f77578(int n, unsigned int* buffers);

// LPALISBUFFER_2b53c18c is a proxy for callback LPALISBUFFER.
char LPALISBUFFER_2b53c18c(unsigned int buffer);

// LPALBUFFERDATA_92ec6ef is a proxy for callback LPALBUFFERDATA.
void LPALBUFFERDATA_92ec6ef(unsigned int buffer, int format, void* data, int size, int freq);

// LPALBUFFERF_3e62bf2e is a proxy for callback LPALBUFFERF.
void LPALBUFFERF_3e62bf2e(unsigned int buffer, int param, float value);

// LPALBUFFER3F_c31de9eb is a proxy for callback LPALBUFFER3F.
void LPALBUFFER3F_c31de9eb(unsigned int buffer, int param, float value1, float value2, float value3);

// LPALBUFFERFV_8ce2743c is a proxy for callback LPALBUFFERFV.
void LPALBUFFERFV_8ce2743c(unsigned int buffer, int param, float* values);

// LPALBUFFERI_aedda2bf is a proxy for callback LPALBUFFERI.
void LPALBUFFERI_aedda2bf(unsigned int buffer, int param, int value);

// LPALBUFFER3I_53a2f47a is a proxy for callback LPALBUFFER3I.
void LPALBUFFER3I_53a2f47a(unsigned int buffer, int param, int value1, int value2, int value3);

// LPALBUFFERIV_b7a68f3 is a proxy for callback LPALBUFFERIV.
void LPALBUFFERIV_b7a68f3(unsigned int buffer, int param, int* values);

// LPALGETBUFFERF_7a9e3c34 is a proxy for callback LPALGETBUFFERF.
void LPALGETBUFFERF_7a9e3c34(unsigned int buffer, int param, float* value);

// LPALGETBUFFER3F_3e3bec12 is a proxy for callback LPALGETBUFFER3F.
void LPALGETBUFFER3F_3e3bec12(unsigned int buffer, int param, float* value1, float* value2, float* value3);

// LPALGETBUFFERFV_71c471c5 is a proxy for callback LPALGETBUFFERFV.
void LPALGETBUFFERFV_71c471c5(unsigned int buffer, int param, float* values);

// LPALGETBUFFERI_ea2121a5 is a proxy for callback LPALGETBUFFERI.
void LPALGETBUFFERI_ea2121a5(unsigned int buffer, int param, int* value);

// LPALGETBUFFER3I_ae84f183 is a proxy for callback LPALGETBUFFER3I.
void LPALGETBUFFER3I_ae84f183(unsigned int buffer, int param, int* value1, int* value2, int* value3);

// LPALGETBUFFERIV_f65c6d0a is a proxy for callback LPALGETBUFFERIV.
void LPALGETBUFFERIV_f65c6d0a(unsigned int buffer, int param, int* values);

// LPALDOPPLERFACTOR_6e5e34c4 is a proxy for callback LPALDOPPLERFACTOR.
void LPALDOPPLERFACTOR_6e5e34c4(float value);

// LPALDOPPLERVELOCITY_37ac1909 is a proxy for callback LPALDOPPLERVELOCITY.
void LPALDOPPLERVELOCITY_37ac1909(float value);

// LPALSPEEDOFSOUND_c05fbd29 is a proxy for callback LPALSPEEDOFSOUND.
void LPALSPEEDOFSOUND_c05fbd29(float value);

// LPALDISTANCEMODEL_69926daa is a proxy for callback LPALDISTANCEMODEL.
void LPALDISTANCEMODEL_69926daa(int distanceModel);

// LPALCCREATECONTEXT_b16fc55a is a proxy for callback LPALCCREATECONTEXT.
ALCcontext* LPALCCREATECONTEXT_b16fc55a(ALCdevice* device, int* attrlist);

// LPALCMAKECONTEXTCURRENT_c13d1dfc is a proxy for callback LPALCMAKECONTEXTCURRENT.
char LPALCMAKECONTEXTCURRENT_c13d1dfc(ALCcontext* context);

// LPALCPROCESSCONTEXT_80cc7c79 is a proxy for callback LPALCPROCESSCONTEXT.
void LPALCPROCESSCONTEXT_80cc7c79(ALCcontext* context);

// LPALCSUSPENDCONTEXT_9429c059 is a proxy for callback LPALCSUSPENDCONTEXT.
void LPALCSUSPENDCONTEXT_9429c059(ALCcontext* context);

// LPALCDESTROYCONTEXT_a63f2985 is a proxy for callback LPALCDESTROYCONTEXT.
void LPALCDESTROYCONTEXT_a63f2985(ALCcontext* context);

// LPALCGETCURRENTCONTEXT_9ef7ba4c is a proxy for callback LPALCGETCURRENTCONTEXT.
ALCcontext* LPALCGETCURRENTCONTEXT_9ef7ba4c();

// LPALCGETCONTEXTSDEVICE_348eacee is a proxy for callback LPALCGETCONTEXTSDEVICE.
ALCdevice* LPALCGETCONTEXTSDEVICE_348eacee(ALCcontext* context);

// LPALCOPENDEVICE_45b1d80f is a proxy for callback LPALCOPENDEVICE.
ALCdevice* LPALCOPENDEVICE_45b1d80f(char* devicename);

// LPALCCLOSEDEVICE_188afbc4 is a proxy for callback LPALCCLOSEDEVICE.
char LPALCCLOSEDEVICE_188afbc4(ALCdevice* device);

// LPALCGETERROR_aeb8780a is a proxy for callback LPALCGETERROR.
int LPALCGETERROR_aeb8780a(ALCdevice* device);

// LPALCISEXTENSIONPRESENT_9fa44568 is a proxy for callback LPALCISEXTENSIONPRESENT.
char LPALCISEXTENSIONPRESENT_9fa44568(ALCdevice* device, char* extname);

// LPALCGETPROCADDRESS_1c77514f is a proxy for callback LPALCGETPROCADDRESS.
void* LPALCGETPROCADDRESS_1c77514f(ALCdevice* device, char* funcname);

// LPALCGETENUMVALUE_c4c0e61c is a proxy for callback LPALCGETENUMVALUE.
int LPALCGETENUMVALUE_c4c0e61c(ALCdevice* device, char* enumname);

// LPALCGETSTRING_62f45f11 is a proxy for callback LPALCGETSTRING.
char* LPALCGETSTRING_62f45f11(ALCdevice* device, int param);

// LPALCGETINTEGERV_a0a1aff2 is a proxy for callback LPALCGETINTEGERV.
void LPALCGETINTEGERV_a0a1aff2(ALCdevice* device, int param, int size, int* values);

// LPALCCAPTUREOPENDEVICE_e864b189 is a proxy for callback LPALCCAPTUREOPENDEVICE.
ALCdevice* LPALCCAPTUREOPENDEVICE_e864b189(char* devicename, unsigned int frequency, int format, int buffersize);

// LPALCCAPTURECLOSEDEVICE_1cfc08b8 is a proxy for callback LPALCCAPTURECLOSEDEVICE.
char LPALCCAPTURECLOSEDEVICE_1cfc08b8(ALCdevice* device);

// LPALCCAPTURESTART_4f64c48c is a proxy for callback LPALCCAPTURESTART.
void LPALCCAPTURESTART_4f64c48c(ALCdevice* device);

// LPALCCAPTURESTOP_c4d1e273 is a proxy for callback LPALCCAPTURESTOP.
void LPALCCAPTURESTOP_c4d1e273(ALCdevice* device);

// LPALCCAPTURESAMPLES_1b5d7311 is a proxy for callback LPALCCAPTURESAMPLES.
void LPALCCAPTURESAMPLES_1b5d7311(ALCdevice* device, void* buffer, int samples);
