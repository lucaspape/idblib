import com.google.gson.Gson
import com.google.gson.JsonElement
import com.google.gson.JsonObject
import com.google.gson.JsonParser

class Idb(private var databasePath: String) {

    private val idbJNI = IdbJNI()
    private val jsonParser = JsonParser()
    private val gson = Gson()

    init {
        idbJNI.Init(databasePath)
    }

    fun getDatabases(): JsonObject {
        return parseResponse(idbJNI.GetDatabases())
    }

    fun createDatabases(name: String): JsonObject {
        return parseResponse(idbJNI.CreateDatabase(name))
    }

    fun loadDatabase(name: String): JsonObject {
        return parseResponse(idbJNI.LoadDatabase(name))
    }

    fun getDatabase(name: String): JsonObject {
        return parseResponse(idbJNI.GetDatabase(name))
    }

    fun getDatabaseTables(name: String): JsonObject {
        return parseResponse(idbJNI.GetDatabaseTables(name))
    }

    fun createTableInDatabase(name: String, tableName: String, fields: JsonObject): JsonObject {
        return parseResponse(idbJNI.CreateTableInDatabase(name, tableName, jsonObjectToString(fields)))
    }

    fun getFromDatabaseTable(name: String, tableName: String, request: JsonObject): JsonObject {
        return parseResponse(idbJNI.GetFromDatabaseTable(name, tableName, jsonObjectToString(request)))
    }

    fun insertToDatabaseTable(name: String, tableName: String, dbObject: JsonObject): JsonObject {
        return parseResponse(idbJNI.InsertToDatabaseTable(name, tableName, jsonObjectToString(dbObject)))
    }

    fun removeFromDatabaseTable(name: String, tableName: String, request: JsonObject): JsonObject {
        return parseResponse(idbJNI.RemoveFromDatabaseTable(name, tableName, jsonObjectToString(request)))
    }

    fun updateInDatabaseTable(name: String, tableName: String, dbObject: JsonObject): JsonObject {
        return parseResponse(idbJNI.UpdateInDatabaseTable(name, tableName, jsonObjectToString(dbObject)))
    }

    private fun parseResponse(response: String): JsonObject {
        val json = jsonParser.parse(response).asJsonObject

        val error = json.get("error")

        if (error != null) {
            throw IdbException(error.asString)
        }

        return json.getAsJsonObject("result")
    }

    private fun jsonObjectToString(jsonObject: JsonObject): String {
        return gson.toJson(jsonObject)
    }
}