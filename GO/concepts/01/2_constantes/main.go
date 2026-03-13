package main

import "fmt"

// Las ocnstantes no pueden modificarse durnate toda la ejeucion del programa
// No puedo usar el declarador de variable corta := ya que asume como variables

const (
	os     = "linux"
	domain = "meli"
)

const (
	Jan = iota + 1 // iota asigna valores de forma incremntal
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

//const os, domain string = "linux", "meli.com" //tambien esa valido

//const os, domain = "linux", "ed.team" //esto tambien esa valido

func main() {
	fmt.Println(os, domain)
	fmt.Println(Jan, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec)
}
