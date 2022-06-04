#include "cw.h"

const char* cw_GetStringUTFChars(JNIEnv * env, jstring str, jboolean * isCopy) {
    return (*env)->GetStringUTFChars(env, str, isCopy);
}

void cw_ReleaseStringUTFChars(JNIEnv *env, jstring str, const char *utf) {
    (*env)->ReleaseStringUTFChars(env, str, utf);
}

const jstring* cw_NewStringUTF(JNIEnv *env, const char *utf) {
    return (*env)->NewStringUTF(env, utf);
}