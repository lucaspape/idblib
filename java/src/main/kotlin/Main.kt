fun main(){
    val ldb = IdbJNI()

    ldb.Init("databases/")

    println(ldb.CreateDatabase("testDB"))
    println(ldb.LoadDatabase("testDB"))
    println(ldb.GetDatabases())
    println(ldb.GetDatabase("testDB"))
    println(ldb.GetDatabaseTables("testDB"))
    println(ldb.CreateTableInDatabase("testDB", "testTable", "{\"testField\": {\"type\": \"text\", \"indexed\": true}}"))
    println(ldb.InsertToDatabaseTable("testDB", "testTable", "{\"testField\": \"hello\"}"))
    println(ldb.GetFromDatabaseTable("testDB", "testTable", "{ \"where\": { \"field\": \"testField\", \"value\": \"hello\", \"operator\": \"=\" } }"))
}