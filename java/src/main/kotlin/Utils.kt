val osName: String = System.getProperty("os.name")

fun isWindows(): Boolean {
    return osName.startsWith("Windows")
}

fun isLinux(): Boolean {
    return osName.startsWith("Linux")
}