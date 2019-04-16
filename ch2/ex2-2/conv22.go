// Package tempconv performs Celsius and Fahrenheit conversions.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/kevdonk/gopl/ch2/ex2-1/tempconv"
)

type Kilogram float64
type Pound float64
type Kilometer float64
type Mile float64

func KgToLb(k Kilogram) Pound { return Pound(k * 2.2) }
func LbToKg(l Pound) Kilogram { return Kilogram(l / 2.2) }
func KmToM(k Kilometer) Mile  { return Mile(k / 1.6) }
func MToKm(m Mile) Kilometer  { return Kilometer(m * 1.6) }

func (k Kilometer) String() string { return fmt.Sprintf("%gkm", k) }
func (m Mile) String() string      { return fmt.Sprintf("%gmi", m) }
func (lb Pound) String() string    { return fmt.Sprintf("%glbs", lb) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			convert(arg)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			convert(input.Text())
		}
	}
}

func convert(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	kg := Kilogram(t)
	lb := Pound(t)
	km := Kilometer(t)
	m := Mile(t)
	fmt.Printf("temperature\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c), c, tempconv.CToK(c), f, tempconv.FToK(f))
	fmt.Printf("distance\n%s = %s\n%s = %s\n",
		km, KmToM(km), m, MToKm(m))
	fmt.Printf("weight\n%s = %s\n%s = %s\n",
		lb, LbToKg(lb), kg, KgToLb(kg))
}
