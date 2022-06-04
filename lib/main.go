package main

import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/lucaspape/idblib"
)

var api *idblib.Api

//export Init
func Init(databasePathC *C.char) {
	databasePath := C.GoString(databasePathC)

	api = idblib.NewApi(databasePath)

	fmt.Println(databasePath)
}

func resultToC(m map[string]interface{}, err error) *C.char {
	rM := make(map[string]interface{})
	rM["result"] = m
	rM["error"] = fmt.Sprint(err)

	j, _ := json.Marshal(rM)
	jS := C.CString(string(j))

	return jS
}

//export GetDatabases
func GetDatabases() *C.char {
	return resultToC(api.GetDatabases())
}

//export CreateDatabase
func CreateDatabase(name *C.char) *C.char {
	return resultToC(api.CreateDatabase(C.GoString(name)))
}

//export LoadDatabase
func LoadDatabase(name *C.char) *C.char {
	return resultToC(make(map[string]interface{}), api.LoadDatabase(C.GoString(name)))
}

//export GetDatabase
func GetDatabase(name *C.char) *C.char {
	return resultToC(api.GetDatabase(C.GoString(name)))
}

//export GetDatabaseTables
func GetDatabaseTables(name *C.char) *C.char {
	return resultToC(api.GetDatabaseTables(C.GoString(name)))
}

//export CreateTableInDatabase
func CreateTableInDatabase(name *C.char, tableName *C.char, fields *C.char) *C.char {
	fieldsI, err := cStringToInterface(fields)

	if err != nil {
		return resultToC(make(map[string]interface{}), err)
	}

	return resultToC(api.CreateTableInDatabase(C.GoString(name), C.GoString(tableName), *fieldsI))
}

//export GetFromDatabaseTable
func GetFromDatabaseTable(name *C.char, tableName *C.char, request *C.char) *C.char {
	requestI, err := cStringToInterface(request)

	if err != nil {
		return resultToC(make(map[string]interface{}), err)
	}

	return resultToC(api.GetFromDatabaseTable(C.GoString(name), C.GoString(tableName), *requestI))
}

//export InsertToDatabaseTable
func InsertToDatabaseTable(name *C.char, tableName *C.char, object *C.char) *C.char {
	objectI, err := cStringToInterface(object)

	if err != nil {
		return resultToC(make(map[string]interface{}), err)
	}

	return resultToC(api.InsertToDatabaseTable(C.GoString(name), C.GoString(tableName), *objectI))
}

//export RemoveFromDatabaseTable
func RemoveFromDatabaseTable(name *C.char, tableName *C.char, request *C.char) *C.char {
	requestI, err := cStringToInterface(request)

	if err != nil {
		return resultToC(make(map[string]interface{}), err)
	}

	return resultToC(api.RemoveFromDatabaseTable(C.GoString(name), C.GoString(tableName), *requestI))
}

//export UpdateInDatabaseTable
func UpdateInDatabaseTable(name *C.char, tableName *C.char, object *C.char) *C.char {
	objectI, err := cStringToInterface(object)

	if err != nil {
		return resultToC(make(map[string]interface{}), err)
	}

	return resultToC(api.UpdateInDatabaseTable(C.GoString(name), C.GoString(tableName), *objectI))
}

func cStringToInterface(c *C.char) (*interface{}, error) {
	i := new(interface{})
	err := json.Unmarshal([]byte(C.GoString(c)), &i)

	return i, err
}

func main() {}
