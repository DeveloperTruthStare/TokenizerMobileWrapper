# Kagome v2 Mobile Wrapper

This is a wrapper for [this](https://github.com/ikawaha/kagome/tree/v2/tokenizer) repository and is designed to be built with gomobile for android or iOS. Currently only instructions and tests for Android are provided, but theoretically the iOS build process is similar.

### Supported Features

Only the `Tokenize(string) Token` method is implemented currently. The signature does differ slightly as I could not return a "complicated" class and instead return a jsonified string of the class. The class signature can be found in [tokenize.go](/tokenizer.go).

# Installation

## Step 1 - Getting the .aar and -sources.jar

### Option 1

Download the latest from the releases

### Option 2

Install gomobile with

`$ go install golang.org/x/mobile/cmd/gomobile@latest`

Then run

`$ gomobile init`

Install an ndk from Android Studio. Go to `tools > SDK Manager` Under `SDK Tools` Check `NDK (Side by Side)` (if given the option to choose, choose most recent, you can eneable this with the "Show Package Details" Toggle in the bottom Right)

Find your project's minSdk found in `app/build.gradle.kts` This may vary depending on the type of project.

Ensure gomobile/bind is downloaded

`go get golang.org/x/gomobile/bind`

**This will be removed with go mod tidy, so you need to do this every time you run that command.**

Then compile the build

`gomobile -o build/tokenizer.aar -target=android -androidapi 34 github.com/DeveloperTruthStare/tokenizer`

## Step 2 - Using this in Android Studio

I created a file `app/libs` that I copied both tokenizer.aar and tokenizer-source.jar into

Go to `File > Project Structure` Under `Dependencies` select `<All Modules>` click the `+` button and select `2 JAR/AAR Dependency`

Step 1. Provide a path to the library file or directory to add.

`libs/tokenizer.aar'

Step 2. Assign your dependency to a configuration by selecting one of the configurations below.

`implementation`

Select `OK`

This line `implementation(files("libs/tokenizer.aar"))` should have been added to your build.gradle.kts and you should now be able to `import tokenizer.Tokenizer` and use as the following

```
val tokenizer = tokenizer.newTokenizer()

val tokenJson = tokenizer.tokenize("寿司が食べたいんですが")
```

To use this effectively I use the gson package

Add the following dependency under dependencies in `build.gradle.kts` and sync gradle

`implementation(libs.gson)`

After the above code you can do the following

```kt
data class Token (
    var surface: String,
    var features: List<String>
)

fun jsonToTokens(jsonString: String): List<Token> {
    val gson = Gson()
    val type = object : TypeToken<List<Token>>() {}.type
    return gson.fromJson(jsonString, type)
}

fun tokenize(japaneseText: String) {
    val tokenizer = tokenizer.newTokenizer()
    val tokenJson = tokenizer.tokenize(japaneseText)
    val tokens = jsonToTokens(tokenJson)
    for (token in tokens) {
        Log.d(token.surface, token.features.toString()) // idk if you can .toString() a List<String> tbh
    }
}
```
