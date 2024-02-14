package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() { 
    if len(os.Args) > 1 {
	args := os.Args[1:]
	parseInput(args)
	os.Exit(0)
    }
    reader := bufio.NewReader(os.Stdin)
    for {
	input,err := reader.ReadString('\n')
	if err == io.EOF {
	    break
	}
	if err != nil {
	    fmt.Printf("ERROR: %s", err)
	    continue
	} 
	values := strings.Split(input, " ")
	parseInput(values)
    }
}

func parseInput(values []string) {
    for _, elt := range values {
	// When values come from a pipe the last 
	// number as a "\n" attach to it.
	if strings.Contains(elt,"\n") {
	    elt = strings.TrimSuffix(elt,"\n")
	}
	number,err := strconv.ParseInt(elt,10,0)
	if err != nil {
	    fmt.Printf("ERROR: %s\n",err)
	    continue
	}
	if number < 0 {
	    fmt.Println("Negative number are not supported")
	    continue
	}
	fmt.Printf("%s\n",simpleNumber(int(number)))
    }
}

//export convert
func convert(number int) *byte {
    bytes := []byte(simpleNumber(number))
    return &(bytes[0])
}

func simpleNumber(number int) string {
    numbers := []string{"zéro", "un", "deux", "trois", "quatre", "cinq", "six", "sept", "huit", "neuf", "dix", "onze", "douze", "treize", "quatorze", "quinze", "seize"}
    if number >= 0 && number < len(numbers) {
	return numbers[number]
    }
    switch number {
    case 20:
	return "vingt"
    case 30:
	return "trente"
    case 40:
	return "quarante"
    case 50:
	return "cinquante"
    case 60:
	return "soixante"
    case 100:
	return "cent"
    case 1000:
	return "mille"
    default:
	return parseComplexe(number)
}
}

func parseComplexe(number int) (string) {
    if number == 81 {
	return ("quatre-vingt-un")
    }
    if number == 91 {
	return "quatre-vingt-onze"
    }

    if number < 100 {
	divider := number / 10
	reminder := number % 10
	switch divider {
	case 7:
	    return handle70s(number)
	case 8:
	    return handle80s(number)
	case 9:
	    return handle90s(number)
	default:
	    var sep string
	    if number%10 != 1 {
		sep = "-"
	    } else {
		sep = " et "
	    }
	    return simpleNumber(divider*10) + sep + simpleNumber(reminder)
	}
    }
    if number < 1000 {
	return handle100(number)
    } else if number < 1_000_000 {
	return handle1000(number)
    } else if number < 1_000_000_000 {
	return handle1M(number)
    } else {
	return "nil"
    }
}

func handle1M(number int) string {
    head := number / 1_000_000
    tail := number % 1_000_000
    var res string
    if head == 1 {
	res = "un million"
    } else {
	res = simpleNumber(head) + " millions"
    }
    if tail != 0 {
	res = res + " " + simpleNumber(tail)
    }
    return res
}

func handle1000(number int) string {
    head := number / 1000
    tail := number % 1000
    var res = "mille"
    if head != 1 {
	res = simpleNumber(head) + " " + res
    }
    if tail != 0 {
	res = res + " " + simpleNumber(tail)
    }
    return res
}

func handle100(number int) string {
    head := number / 100
    tail := number % 100
    var res = "cent"
    if head != 1 {
	res = simpleNumber(head) + " " + res
    }
    // If it's not followed by another number
    // 1_000_000 doesn't count
    if tail == 0 && number < 1000 || number >= 1_000_000 {
	res += "s"
    }
    if tail != 0 {
	res += " " + simpleNumber(tail)
    }
    return res
}

func handle90s(number int) string {
    var tail = number - 80
    return "quatre-vingt" + "-" + simpleNumber(tail)
}

func handle80s(number int) string {
    tail := number - 60
    res := "quatre-" + simpleNumber(tail)
    // If it's not followed by another number
    if tail%20 == 0 {
	return res + "s"
    }
    return res
}

func handle70s(number int) string {
    var tail = number - 60
    var sep string
    if number%10 != 1 {
	sep = "-"
    } else {
	sep = " et "
    }
    return "soixante" + sep + simpleNumber(tail)
}
