package main

// #include "cw.h"
import "C"

//export Java_IdbJNI_Init
func Java_IdbJNI_Init(env *C.JNIEnv, jObject C.jobject, databasePath C.jstring) {
	databasePathC := getCString(env, databasePath)

	Init(databasePathC)

	releaseJString(env, databasePath, databasePathC)
}

//export Java_IdbJNI_GetDatabases
func Java_IdbJNI_GetDatabases(env *C.JNIEnv, jObject C.jobject) *C.jstring {
	result := GetDatabases()

	return getJString(env, result)
}

//export Java_IdbJNI_CreateDatabase
func Java_IdbJNI_CreateDatabase(env *C.JNIEnv, jObject C.jobject, name C.jstring) *C.jstring {
	nameC := getCString(env, name)

	result := CreateDatabase(nameC)

	releaseJString(env, name, nameC)

	return getJString(env, result)
}

//export Java_IdbJNI_LoadDatabase
func Java_IdbJNI_LoadDatabase(env *C.JNIEnv, jObject C.jobject, name C.jstring) *C.jstring {
	nameC := getCString(env, name)

	result := LoadDatabase(nameC)

	releaseJString(env, name, nameC)

	return getJString(env, result)
}

//export Java_IdbJNI_GetDatabase
func Java_IdbJNI_GetDatabase(env *C.JNIEnv, jObject C.jobject, name C.jstring) *C.jstring {
	nameC := getCString(env, name)

	result := GetDatabase(nameC)

	releaseJString(env, name, nameC)

	return getJString(env, result)
}

//export Java_IdbJNI_GetDatabaseTables
func Java_IdbJNI_GetDatabaseTables(env *C.JNIEnv, jObject C.jobject, name C.jstring) *C.jstring {
	nameC := getCString(env, name)

	result := GetDatabaseTables(nameC)

	releaseJString(env, name, nameC)

	return getJString(env, result)
}

//export Java_IdbJNI_CreateTableInDatabase
func Java_IdbJNI_CreateTableInDatabase(env *C.JNIEnv, jObject C.jobject, name C.jstring, tableName C.jstring, fields C.jstring) *C.jstring {
	nameC := getCString(env, name)
	tableNameC := getCString(env, tableName)
	fieldsC := getCString(env, fields)

	result := CreateTableInDatabase(nameC, tableNameC, fieldsC)

	releaseJString(env, name, nameC)
	releaseJString(env, tableName, tableNameC)
	releaseJString(env, fields, fieldsC)

	return getJString(env, result)
}

//export Java_IdbJNI_GetFromDatabaseTable
func Java_IdbJNI_GetFromDatabaseTable(env *C.JNIEnv, jObject C.jobject, name C.jstring, tableName C.jstring, request C.jstring) *C.jstring {
	nameC := getCString(env, name)
	tableNameC := getCString(env, tableName)
	requestC := getCString(env, request)

	result := GetFromDatabaseTable(nameC, tableNameC, requestC)

	releaseJString(env, name, nameC)
	releaseJString(env, tableName, tableNameC)
	releaseJString(env, request, requestC)

	return getJString(env, result)
}

//export Java_IdbJNI_InsertToDatabaseTable
func Java_IdbJNI_InsertToDatabaseTable(env *C.JNIEnv, jObject C.jobject, name C.jstring, tableName C.jstring, object C.jstring) *C.jstring {
	nameC := getCString(env, name)
	tableNameC := getCString(env, tableName)
	objectC := getCString(env, object)

	result := InsertToDatabaseTable(nameC, tableNameC, objectC)

	releaseJString(env, name, nameC)
	releaseJString(env, tableName, tableNameC)
	releaseJString(env, object, objectC)

	return getJString(env, result)
}

//export Java_IdbJNI_RemoveFromDatabaseTable
func Java_IdbJNI_RemoveFromDatabaseTable(env *C.JNIEnv, jObject C.jobject, name C.jstring, tableName C.jstring, request C.jstring) *C.jstring {
	nameC := getCString(env, name)
	tableNameC := getCString(env, tableName)
	requestC := getCString(env, request)

	result := RemoveFromDatabaseTable(nameC, tableNameC, requestC)

	releaseJString(env, name, nameC)
	releaseJString(env, tableName, tableNameC)
	releaseJString(env, request, requestC)

	return getJString(env, result)
}

//export Java_IdbJNI_UpdateInDatabaseTable
func Java_IdbJNI_UpdateInDatabaseTable(env *C.JNIEnv, jObject C.jobject, name C.jstring, tableName C.jstring, object C.jstring) *C.jstring {
	nameC := getCString(env, name)
	tableNameC := getCString(env, tableName)
	objectC := getCString(env, object)

	result := UpdateInDatabaseTable(nameC, tableNameC, objectC)

	releaseJString(env, name, nameC)
	releaseJString(env, tableName, tableNameC)
	releaseJString(env, object, objectC)

	return getJString(env, result)
}

func getCString(env *C.JNIEnv, j C.jstring) *C.char {
	return C.cw_GetStringUTFChars(env, j, (*C.jboolean)(nil))
}

func releaseJString(env *C.JNIEnv, j C.jstring, s *C.char) {
	defer C.cw_ReleaseStringUTFChars(env, j, s)
}

func getJString(env *C.JNIEnv, c *C.char) *C.jstring {
	return C.cw_NewStringUTF(env, c)
}
