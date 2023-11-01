package syntax

import "fmt"

/*
Use of Interface: You can use interface when in methods or functions you want to pass
different types of argument in them just like Println () function. Or you can also use
interface when multiple types implement same interface.
*/

// map are key values types
// keys can be
/*
boolean
numeric
string,
pointer
channel
interface types
structs – if all it’s field type is comparable
array – if the type of value of array element is comparable
//////////////
and this types can't be as types
Slice
Map
Function
*/

// values can be :
//Value can be of any type in a map

type ErrorsInterface interface {
	Length() string
	badKeyType() string
	badValueType() string
}

type myerror struct {
}

func (i myerror) Length() string {
	return "lenght problem ..."
}

func (i myerror) badKeyType() string {
	return "bad key types ..."
}

func (i myerror) badValueType() string {
	return "values has bad types ..."
}

func MapIntArrayAndStringValue(keyArray []int, valueArray []string) (returnedMap map[int]string, err string) {

	var ei ErrorsInterface
	ei = myerror{}

	if len(keyArray) != len(valueArray) {
		return nil, ei.Length()
	}

	var myMap = make(map[int]string)
	for _, key := range keyArray {
		for i := 0; i < len(valueArray); i++ {
			myMap[key] = valueArray[i]
		}
	}
	return myMap, ""
}

// Check TypeOf an Interface
func WeCanCheckTypeOfInterfaceLikeThis(a interface{}) {
	////////////////////////////////////////////////////////////////////////
	//we can check if ei is same type as ErroreiInterfacer
	val, ok := a.(ErrorsInterface)
	fmt.Printf("val of interface is : %v , ok :%v", val, ok)
	//we can also check interface type in a switch with ei.(type) command
	CheckTypeOfInterface("string")
	CheckTypeOfInterface(56.7)
	CheckTypeOfInterface(true)
	////////////////////////////////////////////////////////////////////////
}

func CheckTypeOfInterface(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("intercae is int type")
	case float64:
		fmt.Println("type is float 64")
	case string:
		fmt.Println("interface is string type")
	case ErrorsInterface:
		fmt.Println("interface is custom type (errorINterface)")
	}
}
