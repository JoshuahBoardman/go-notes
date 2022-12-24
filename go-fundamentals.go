// ~ Variables ~

// Explicit variable declaration 
var explicitVariableName string = "String Value"  // doesnt have to be a string, it could be any other variable type.

// Implicit varible declaration 
var implictVaribleName = "Implicit varible value"

// Implicit varible declaration with the walrus operator
walrusImplicitVariableName := "Walrus implicit variable value"


// ~ Data Types ~

// Character
// the character type is one letter wrapped in single quotes
var characterVariable rune = 'a'

// String
// the string type is a group of characters wrapped in double quotes
var stringVariable string = "String value"

// Integer Types:

// Unsigned Integer
// unsigned integers are always posative
var uint8Variable uint8 = 255 // uint8 can be 0 - 255
var uint16Variable uint16 = 65535 // uint16 can be 0 - 65535
var uint32Variable uint32 = 4294967295 // uint32 can be 0 - 4294967295
var uint64Variable uint64 = 18446744073709551615 // uint64 can be 0 - 18446744073709551615

// Signed Integer
// signed integers can be posative or negative
var int8Variable int8 = 127 // int8 can be -128 - 127
var int16Variable int16 = 32767 // in16 can be -32768 - 32767
var int32Variable int32 = 2147483647 // int32 can be -2147483648 - 2147483647
var int64Variable int64 = 9223372036854775807 // int64 can be -9223372036854775808 - 9223372036854775807

// Machine Dependent Integer
// machine dependent integers can be negative or posative
// machine dependent integers ajust depending on the machine the code is being ran on 
var uintVariable uint = 18446744073709551615 // uint can be either uint32 or uint64 depending on system
// int is often times what you will use to represent numbers unless you have a specific reason not to
var intVariable int = 9223372036854775807 // int can be either a int32 or int64 depending on system
var uintpterVariable uintptr = 0xc82000c2900 // uintptr is another machine dependent integer type that is used for arithmetic on unsafe pointers

// Float Types:
var float32Variable float32 = 98.7 // float32 is for smaller floats
var float64Variable float64 = 1.7e+308 // float64 is for much larger numbers

// Complex Types:
// Both complex32 and complex64 are both data types for high level mathematics

// Booleans
// booleans are true or false values
var booleanVariable bool = true 


// ~ Data Structures ~

// Array:
// arrays have a fixed length 

// Explicit array initialization 
var explicitArrayVariable [4]string = [4]string {"An", "array", "of", "strings",}

// Implicit array initialization
var implicitArrayVariable := [4]string {"An", "array", "of", "strings",}

// Slice:

// Creating a slice from an already declared array
var sliceOfPreDeclaredArray[]string = explicitArrayVariable[0:3]

// Creating a slice without a pre declared array
var sliceVariable[]string = []string {"A", "slice", "without", "a", "pre", "declared", "array",}

// Increasing slices length with a new element (technically making a new slice)
var sliceWithNewElement := append(sliceVariable, "plus")

// Creating a slice with make()
var sliceMadeWithMake := make([]string, 3)

// Reslicing an array
var sliceOfPreDeclaredArray[0:] // Reslicing to get the whole pre declared array 

// Map:
// maps store values in key data pairs

// Map literal 
var mapLiteral map[string]string = map[string]string {"firstKey": "firstValue", "secondKey": "secondValue", "thirdKey": "thirdValue",}

// Implicit map initialization
var implicitMapInitalization := map[string]int {"firstKey": 1, "secondKey": 2, "thirdKey": 3,}

// Creat a map with make()
var mapMadeWithMake := make(map[string]string) // Empty map

// Access map values
mapLiteral["firstKey"]

// Change map values
mapLiteral["secondKey"] = "forthValue"

// Add a value to a map
mapLiteral["fourthKey"] = "fifthValue"

// Delete a map value
delete(mapLiteral, "fourthKey")

// Check if a map contains a value
keysValueIfExists, boolIfKeyExists := mapLiteral["fristKey"]


// ~ Functions ~
// functions are blocks of reusable code 

// Creating a function
// You are not required to take parameters or retrun a value with functions
func functionName(parameter string) string {
    parameter = "Functions can take in paramaters to perform operations on data and return that data to be used in the program"
    return parameter
}

// Call a function
// You can reference a function without calling it by not adding "()" at the end
functionName("random string")

// Assign a functions return value to a variable 
var functionsReturnValue = functionName(explicitVariableName)

// Defer keyword in functions
// The defer keyword will run some code after the function has returned
// The defer keyword can be great for clean up
func deferKeyword(valueToBeDefered int) int {
    defer explicitVariableName = valueToBeDefered
    return 1 + valueToBeDefered
}


// ~ Pointers & Dereference Operator ~

// Pointers
// pointers are the memory address that a value is stored at
var pointerOfExplicitVariableName = &explictVariableName // this stroed the memory address of explictVariableName

// Pointer as a parameter
func dereferencePointerParam(pointerParam *string) string {
   // Dereference Operator
   // Dereferencing gets the value stored a a pointers memory address
   dereferencedPointer := *pointerParam 
   return dereferencedPointer
}

// ~ Structs ~
// Structs allow you to make custom types which can contain multiple fields of different typed data

// Creating a struct
type structName struct {
    firstField string
    secondField int
    thirdField bool
}

// Initializing a variable of a struct type
var structVariable structName = structName {"firstValue", 2, true}

// Embedded Struct 
// embedded structs are structs types used inside another struct (often times they will be a pointer so you can mutate the orginal value instead of creating a copy)
type containsStructName struct {
    embeddedStruct structName
    secondField float32
    thirdField rune
}

// Initializing an embedded struct
var containsEmbeddedStruct containsStructName = containsStructName {structName {"embedded struct first field", 1947, true}, 3.14, 'e'}

// Accessing Embedded Struct
containsEmbeddedStruct.embeddedStruct 

// Accessing Embedded Struct Value
containsEmbeddedStruct.embeddedStruct.firstFiled


// ~ Methods ~ 
// methods are functions that can only be accessed on specified types

// Method Creation
// you will often want to pass the reciver type as a pointer so you can mutate the original value and not a copy
func (s *structName) structMethod(concatFirstField string) {
    s.firstField = s.firstField + " " + concatFirstField
}

// Calling a method 
structVaribale.structMethod("has been mutated by a method")


// ~ Interfaces ~
// interfaces are defining a group of types that impliment some form of shared functionality (methods)

// Create an Interface 
type interfaceName interface {
   structMethod(concatFirsjutField string) // Any type that has a method called structMethod that takes in the same arguments impliments the interfaceName interface
} 

// Using an Interface
// interfaces can be used anywhere you can use a type
func useInterfaceName(implimentsInterface interfaceName) {
    implimentsInterface.structMethod(" by a function that only takes data that impliments the interfaceName interface")
}

