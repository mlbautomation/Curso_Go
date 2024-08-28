package tipos

import "fmt"

func Tipos() {

	const PI = 3.14159
	var cat, dog string = "🐈", "🐕"
	fmt.Printf("El %s se peleó con el %s por %v minutos\n", dog, cat, PI) // no realiza el salto de linea
	fmt.Println("El", dog, "se peleó con el", cat, "por", PI, "minutos")  //realiza el salto de linea
	fmt.Printf("El tipo de variable de %v es %T\n", PI, PI)

	//Tipos de datos
	var var_1 bool = true        //true/false
	var var_2 int16 = 500        //no esta permitido comilla simple para string '
	var var_3 rune = '?'         //con esto podemos tener el valor de un caracter unicode
	var var_4 complex64 = 1 + 1i //complejo
	//GO es fuertemente tipado, por lo que no se puede hacer operaciones con diferente tipos de datos, tenemos de usar convertidores
	var var_5 = int32(var_2) + var_3
	fmt.Println("var_1:", var_1)
	fmt.Println("var_2:", var_2)
	fmt.Println("var_3:", var_3)
	fmt.Println("var_4:", var_4)
	fmt.Println("var_5:", var_5)

	//identificador blank para reservar variables que no usamos
	_ = 1.73
	var _ string = "vacio"

	//el valor cero es asignado inmediatamente
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println("Valor cero es asignado inmediatamente i:", i)
	fmt.Println("Valor cero es asignado inmediatamente f:", f)
	fmt.Println("Valor cero es asignado inmediatamente b:", b)
	fmt.Println("Valor cero es asignado inmediatamente s:", s)
}

/*
fmt - Go Packages
https://pkg.go.dev/fmt

General:
--------
%v	the value in a default format when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value

Boolean:
--------
%t	the word true or false

Integer:
--------
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"

Floating-point and complex constituents:
----------------------------------------
%b	decimalless scientific notation with exponent a power of two, in the manner of strconv.FormatFloat with the 'b' format, e.g. -123456p-78
%e	scientific notation, e.g. -1.234456e+78
%E	scientific notation, e.g. -1.234456E+78
%f	decimal point but no exponent, e.g. 123.456
%F	synonym for %f
%g	%e for large exponents, %f otherwise. Precision is discussed below.
%G	%E for large exponents, %F otherwise
%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20
String and slice of bytes (treated equivalently with these verbs):
%s	the uninterpreted bytes of the string or slice
%q	a double-quoted string safely escaped with Go syntax
%x	base 16, lower-case, two characters per byte
%X	base 16, upper-case, two characters per byte

Slice:
------
%p	address of 0th element in base 16 notation, with leading 0x

Pointer:
--------
%p	base 16 notation, with leading 0x
The %b, %d, %o, %x and %X verbs also work with pointers, formatting the value exactly as if it were an integer.

The default format for %v is:
-----------------------------
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p

For compound objects, the elements are printed using these rules, recursively, laid out like this:
--------------------------------------------------------------------------------------------------
struct:             {field0 field1 ...}
array, slice:       [elem0 elem1 ...]
maps:               map[key1:value1 key2:value2 ...]
pointer to above:   &{}, &[], &map[]

... seguir la documentación de goland: fmt - Go Packages
https://pkg.go.dev/fmt
*/
