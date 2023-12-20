* Go is case sensitive
* Variable type needs to be declared in advance
* Everything is type.  

# Variable declaration
* Variables can be declared in 3 ways explicit, implicit and the walrus operator
	* Explicit, var a int = 4
	* Implicit, var a = 4
	* Using walrus operator, a := 4
* There are various data types type and the types available are
	* String
	* Bool
	* Integer - > uint, uint64, int8, int64 uintptr (aliases are involved)
	* Floating -> float32, float64
	* Complex
	* Arrays
	* Slices
	* Maps
	* Structs
	* Pointers
	* Functions are also a types
	* Channels
	* mutex
# User Input
* bufio/reader and os are the packages that are used for reading the input from the sure
	* reader := bufio.NewReader(os.Stdin) -> And the values is read from the standard input
	* input, _ := reader.ReadString('\n')
	* And the reader will read the input until a new line is entered
# Type conversion
* The type of the data are implemented through function
	* strconv is the package that needs to be used to convert a datatype
	* And it has various function like FormatFloat, FormatInt, FormatUint etc
* To trim the new line character from the string string package is required
	* strconv.ParseFloat(strings.TrimSpace(input), 64)
	* strings package has all the string related functions

# Time in golang
- time is a package in golang
- For formatting you need to always use "01-02-2006 15:04:05 Monday" as a parameter to the time.Format() function - https://pkg.go.dev/time#Time.Format
# Memory Management in golang
* Memory allocation and deallocation happens automatically
* new()
	* Allocate memory but no INIT
	* you will get a memory address
	* zeroed storage
* make()
	* Memory is allocated and INIT
	* you will get a memory address
	* non-zeroed storage
- GC happens automatically
- Anything Out of Scope or nil is eligible for GC.
- https://pkg.go.dev/runtime

# Pointers
* Pointers in golang is more like c
* Pointer declaration
	* `var ptr *int
	* `var ptr = &myNumber
* and * is used to access the value from a pointer
	*  `*ptr = *ptr * 2
# Arrays 
* Declaring an array gets, initialised with default values of the datatypes
* Initialization 
```
var vegList = [3]string{"potato", "beans", "mushroom"}
var fruitList [4]string
fruitList[0] = "Apple"
fruitList[1] = "Tomato"
fruitList[2] = "Peach"
```
# Slices
* Declaring Slices gets initialised with nil as the default value
* The difference between declaring arrays and slices in go is we omit the length parameter
```
var vegList = []string{}
slice := []int{1, 2, 3}
```
* You can also use make function to declare a slice
	* with capacity
		*` slice1 := make([]int, 5, 10)
	* without capacity
		* `slice2 := make([]int, 5)
# Loops

```
for key, value := range slice {
	%% Statement %%
}
```

# Structs
* Structs in go are like classes in python
* But it does not support inheritance
```
type User struct {
	Name string
	Email string
	Status bool
	Age int
}
```
