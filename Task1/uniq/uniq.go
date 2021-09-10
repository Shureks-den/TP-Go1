package uniq

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Flags struct {
	c      *bool
	d      *bool
	u      *bool
	f      *int
	s      *int
	i      *bool
	input  *string
	output *string
}

func ignReg(str string, fl *Flags) string {
	if *fl.i {
		return strings.ToLower(str)
	}
	return str
}

func slice(str string, fl *Flags) string {
	if len(str)-1 > *fl.s {
		return str[*fl.s:]
	}
	return str
}

func ignWrds(str string, fl *Flags) string {
	res := strings.Split(str, " ")
	if len(res)-1 > *fl.f {
		return strings.Join(res[*fl.f:], " ")
	}
	return str
}

func check(sum int, fl *Flags) bool {
	switch {
	case *fl.c:
		return sum != 0
	case *fl.d:
		return sum > 1
	case *fl.u:
		return sum == 1
	default:
		return sum >= 1
	}
}

func writeString(flag *bool, str string, sum int) string {
	if *flag {
		str = strconv.Itoa(sum) + " " + str + "\n"
	} else {
		str = str + "\n"
	}
	return str
}

func Read(uniqData *[]string, fl *Flags) {
	// choosing input
	var sc *bufio.Scanner
	if *fl.input == "default" {
		sc = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(*fl.input)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		sc = bufio.NewScanner(file)
	}

	if (*fl.c && *fl.d) || (*fl.c && *fl.u) || (*fl.u && *fl.d) {
		log.Fatal("Either one of c, d or u can be used in one call")
	}

	var prev string
	var first string
	sum := 0
	for ; sc.Scan(); sum++ {
		txt := sc.Text()
		// добавление строки при соблюдении всех условий
		if slice(ignWrds(ignReg(prev, fl), fl), fl) != slice(ignWrds(ignReg(txt, fl), fl), fl) {
			if check(sum, fl) {
				// так как у флага -c отличается запись, то нужно учитывать этот флаг
				*uniqData = append(*uniqData, writeString(fl.c, first, sum))
			}
			sum = 0
		}

		if sum == 0 {
			first = txt
		}
		prev = txt
	}
	// добавление последней строки
	if check(sum, fl) {
		*uniqData = append(*uniqData, writeString(fl.c, first, sum))
	}
}

func Write(uniqData []string, fl *Flags) {
	if *fl.output == "default" {
		for i := range uniqData {
			fmt.Printf("%s", uniqData[i])
		}
		return
	}
	file, err := os.Create(*fl.output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for i := range uniqData {
		file.WriteString(uniqData[i])
	}
}

func ParseFlags(fl *Flags) {
	fl.input = flag.String("input_file", "default", "File to read data from")
	fl.output = flag.String("output_file", "default", "File to write data to")
	fl.c = flag.Bool("c", false, "Print string num")
	fl.d = flag.Bool("d", false, "Print strings that occure more than once")
	fl.u = flag.Bool("u", false, "Print strings that occure once")
	fl.i = flag.Bool("i", false, "Ignore letter register")
	fl.s = flag.Int("s", 0, "Ignore first n symbols")
	fl.f = flag.Int("f", 0, "Ignore first n words")
	flag.Parse()
}
