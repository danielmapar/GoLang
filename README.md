# Go - Tutorial

## Packages and Modules

* By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

* Note that functionName starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).

* Go executes init functions automatically at program startup, after global variables have been initialized. For more about init functions, see Effective Go.

* Ending a file's name with _test.go tells the go test command that this file contains test functions.

## Functions

* A return statement without arguments returns the named return values. This is known as a "naked" return.
    * ```golang
        func split(sum int) (x, y int) {
            x = sum * 4 / 9
            y = sum - x
            return
        }
        ```
* A var declaration can include initializers, one per variable. If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
    * `var i, j int = 1, 2`

## Commands

* `go mod tidy`
* `go run example.go`: Runs a go program
* `go mod init example.com/mymodule`: Initializes a module
* `go test . -v`: The go test command executes test functions (whose names begin with Test) in test files (whose names end with _test.go). You can add the -v flag to get verbose output that lists all of the tests and their results. Implement test functions in the same package as the code you're testing.
* `go build`: command compiles the packages, along with their dependencies, but it doesn't install the results.
* `go install`: command compiles and installs the packages.

## Types

* Go's basic types are:

    * ```golang
        bool

        string

        int  int8  int16  int32  int64
        uint uint8 uint16 uint32 uint64 uintptr

        byte // alias for uint8

        rune // alias for int32
            // represents a Unicode code point

        float32 float64

        complex64 complex128
        ```
* The zero value is:
    * `0` for numeric types,
    * `false` for the boolean type, and
    * `""` (the empty string) for strings.

* The expression `T(v)` converts the value `v` to the type `T`.

* Constants are declared like variables, but with the `const` keyword.

* A slice is like an array, except that its size changes dynamically as you add and remove items. The slice is one of Go's most useful types.

* In Go, you initialize a map with the following syntax: make(map[key-type]value-type). You have the Hellos function return this map to the caller. For more about maps, see Go maps in action on the Go blog.

* In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value. You don't need the index, so you use the Go blank identifier (an underscore) to ignore it.
    * `for _, name := range names {`

* Test functions take a pointer to the testing package's testing.T type as a parameter..
    * `func TestHelloName(t *testing.T) {`

## Operators

* `:=`: operator is a shortcut for declaring and initializing a variable in one line.
    * ```golang
        var message string
        message = fmt.Sprintf("Hi, %v. Welcome!", name)

        // or

        message := fmt.Sprintf("Hi, %v. Welcome!", name)
        ```

## Directives

* `replace module-path [module-version] => replacement-path [replacement-version]`
    * Replaces the content of a module at a specific version (or all versions) with another module version or with a local directory. Go tools will use the replacement path when resolving the dependency.

* To reference a published module, a go.mod file would typically omit the replace directive and use a require directive with a tagged version number at the end.
    * `require example.com/greetings v1.1.0`