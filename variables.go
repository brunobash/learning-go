//Explicit declaration
var name string
var age int

//Declaration with inicialization
var name string = "Alice"
var age int = 20

//Short declaration
name := "Alice"
age := 20

/*
# Primitive data types in Go:

- Integers: int, int8, int16, int32, int64
- Floating-points: float32, float64
- Boolean: bool
- String: string

# Derived data types in Go:
- Arrays: A collection of elements of the same type with fixed size
- Slices: A "slice" of an array with dinamic size
- Maps: Collection of key-value pairs
- Structs: Composite data type that group variable under a sigle name
*/

//Integers:
var x int = 14
var y int32 = 25

//Floating-points:
var tx float64
var pi float64 = 3.1415
nap := 2.71828

//Booleans:
var a bool
var b bool = false
c := true

//Strings:
var k string
var j string = "LetterJ"
w := "LetterW"

//Array
var numbers [5]int = [5]int{2, 4, 6, 8, 10}
var names [7]string = [7]string{"Peter", "James", "John", "Andrew", "Philip", "Matthew", "Thomas"}

//Slice
var fruits []string = []string{"Pear", "Apple", "Strawberry"}
var numbers []int = []int{1, 2, 3, 5, 7, 11}


//Map
var countryCapital map[string]string = map[string]string{
	"Brazil":  "Brasília",
	"Canada": "Ottawa",
	"Germany": "Berlin",
}

//Struct
type Person struct {
	Name string
	Age int
}

var p Person = Person{"Alice", 20}

/*
If you don't declare a varible without initializing it, it assumes a default value
known as "zero value"
*/

var counter int //couter is 0
var active bool //active is false

/* 
## Good Practices:

- Prefer short statements: Use := for a clean syntax, but be careful with scope and visibility
- Const: Use const for imutable values. When a value must not change, use const instead of var.
*/

//Const:
const Gold = 1.61803


