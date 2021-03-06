import java.io.File

class IdbJNI {
    init {
        if(isWindows()){
            System.load(File("idblib.dll").canonicalPath)
        }else if(isLinux()){
            System.load(File("idblib.so").canonicalPath)
        }else{
            throw java.lang.Exception("Operating system not supported by idblib")
        }
    }

    external fun Init(databasePath: String): String
    external fun GetDatabases(): String
    external fun CreateDatabase(name: String): String
    external fun LoadDatabase(name: String): String
    external fun GetDatabase(name: String): String
    external fun GetDatabaseTables(name: String): String
    external fun CreateTableInDatabase(name: String, tableName: String, fields: String): String
    external fun GetFromDatabaseTable(name: String, tableName: String, request: String): String
    external fun InsertToDatabaseTable(name: String, tableName: String, dbObject: String): String
    external fun RemoveFromDatabaseTable(name: String, tableName: String, request: String): String
    external fun UpdateInDatabaseTable(name: String, tableName: String, dbObject: String): String
}