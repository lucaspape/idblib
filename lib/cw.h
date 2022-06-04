#include "jni.h"

const char* cw_GetStringUTFChars(JNIEnv *, jstring, jboolean *);
void cw_ReleaseStringUTFChars(JNIEnv *, jstring, const char *);
const jstring* cw_NewStringUTF(JNIEnv *, const char *);