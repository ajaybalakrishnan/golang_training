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

for i := 0; i < len(days); i++ {
	fmt.Println(days[i])
}

for i := range days {
	fmt.Println(days[i])
}

for index, day := range days {
	fmt.Printf("Index is %v and value is %v\n", index, day)
}

rougueValue := 1
for rougueValue < 10 {
	if rougueValue == 5 {
		rougueValue++
		continue
	} else if rougueValue == 6 {
		break
	}
	fmt.Println("Value is: ", rougueValue)
	rougueValue++
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
# Conditional Statement
* If else is like it is in any other programming language
```
if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out!"
	} else {
		result = "Exactly 10 loging"
	}
```

# Switch Case

* Syntax for switch case

```
switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough // next case is also executed
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot")
	default:
		fmt.Println("That was impossible!")
	}
}
```
* The `fallthrough` keyword executed the following case as well irrespective of the condition
# Control Statements / Loops

* loop can have multiple syntax in go
	* it can use range keyword, 
	* it can use var declaration and incremental statement int he for statement
	* It can have just the control statement alone
```
for i := 0; i < len(days); i++ {
 	fmt.Println(days[i])
}

for i := range days {
 	fmt.Println(days[i])
}

for index, day := range days {
	fmt.Printf("Index is %v and value is %v\n", index, day)
}
```
```
for rougueValue < 10 {
		if rougueValue == 4 {
			goto flag
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		} else if rougueValue == 6 {
			break
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}
```

# Functions
* Functions are declared with the keyword `func` the data type for the arguments needs to be declared in the function definition.
* The return type needs to  be mentioned in the function definition if the function is returning something. It can return either single data type or multiple data types.
```
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	sum := 0
	for _, val := range values {
		sum = sum + val
	}
	return sum, "Pro result function"
}
```

# Methods
* Go does not have classes. However, you can define methods on types.
* A method is a function with a special _receiver_ argument.
* The receiver appears in its own argument list between the `func` keyword and the method name.

```
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active", u.Status)
}

func (u User) ResetMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email Reseted to:", u.Email)
}
```

# Defer Statement
* Defer statements executes the control line reaches the end of the function body.
* When the multiple defer statements are invoked, the last defered statement will be executed first. i.e LIFO (Last In First Out).
```
func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	myDefer()
}

func myDefer() {
	defer fmt.Print("\n")
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
```
**Output**
```
[ajayb@localhost 17defer]$ go run main.go 
Hello
43210
Two
One
World
[ajayb@localhost 17defer]$
```

# Files
* File operations is more like any it is in any other programming languages
* In go packages `io` and `os` provides all the necessary method to modify create and remove a file
* `defer` keyword is mostly used with `file.Close()` statement
```
content := "This needs to go in a file."
file, err := os.Create(("./myfile.txt"))
length, err := io.WriteString(file, content)

databyte, err := os.ReadFile(filename)
// err needs to be handled
fmt.Println("Text data inside teh file is \n", string(databyte))
```

# Handling web request

* Web Requests handling is done via [http](https://pkg.go.dev/net/http) package.
* Any of the request returns a [response](https://pkg.go.dev/net/http#Response) object
	* This response object has all the 
	* And this response needs to be closed by the client as it is not closed by default
	```
	defer response.Body.Close()
	```
	* The content of the body can be read like the following
```
	databytes, err := io.ReadAll(response.Body)
	content := string(databytes)
	fmt.Println(content)
```
-  If the network connection fails or the server terminates the response, Body.Read calls return an error.

# Handling URLs
- Urls can be handled using the url module
- Simply pass the url as a parameter to the `url.Parse()` function.
```
const myurl = "https://lco.dev:3000/learn?course=golang&paymentid=nsjfiuefn"
result, _ := url.Parse(myurl)
```
- `Parse` function returns `*URL`. And the URL type has the below fields
```
type URL struct {
	Scheme      [string](https://pkg.go.dev/builtin#string)
	Opaque      [string](https://pkg.go.dev/builtin#string)    // encoded opaque data
	User        *[Userinfo](https://pkg.go.dev/net/url#Userinfo) // username and password information
	Host        [string](https://pkg.go.dev/builtin#string)    // host or host:port
	Path        [string](https://pkg.go.dev/builtin#string)    // path (relative paths may omit leading slash)
	RawPath     [string](https://pkg.go.dev/builtin#string)    // encoded path hint (see EscapedPath method)
	OmitHost    [bool](https://pkg.go.dev/builtin#bool)      // do not emit empty host (authority)
	ForceQuery  [bool](https://pkg.go.dev/builtin#bool)      // append a query ('?') even if RawQuery is empty
	RawQuery    [string](https://pkg.go.dev/builtin#string)    // encoded query values, without '?'
	Fragment    [string](https://pkg.go.dev/builtin#string)    // fragment for references, without '#'
	RawFragment [string](https://pkg.go.dev/builtin#string)    // encoded fragment hint (see EscapedFragment method)
}
```
- Similarly URLs can be constructed using the `url` module. 
```
partsOfUrl := &url.URL{
	Scheme:   "https",
	Host:     "lco.dev",
	Path:     "learn",
	RawQuery: "course=golang&paymentid=nsjfiuefn",
}
anotherUrl := partsOfUrl.String()
```
