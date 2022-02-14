package uniq

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
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

func ignSymb(str string, fl *Flags) string {
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

func checkFlgs(sum int, fl *Flags) bool {
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

func formatString(flag *bool, str string, sum int) string {
	if *flag {
		str = strconv.Itoa(sum) + " " + str
	}
	str = str + "\n"
	return str
}

func Read(uniqData *[]string, fl *Flags) error {
	// выбираем откуда читать
	var sc *bufio.Scanner
	if *fl.input == "default" {
		sc = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(*fl.input)
		if err != nil {
			return err
		}
		defer file.Close()
		sc = bufio.NewScanner(file)
	}

	if (*fl.c && *fl.d) || (*fl.c && *fl.u) || (*fl.u && *fl.d) {
		return errors.New("Either one of c, d or u can be used in one call")
	}

	var prevStr string
	var firstStr string
	sum := 0
	for ; sc.Scan(); sum++ {
		curStr := sc.Text()
		// добавление строки при соблюдении всех условий
		if ignSymb(ignWrds(ignReg(prevStr, fl), fl), fl) != ignSymb(ignWrds(ignReg(curStr, fl), fl), fl) {
			if checkFlgs(sum, fl) {
				// так как у флага -c отличается запись, то нужно учитывать этот флаг
				*uniqData = append(*uniqData, formatString(fl.c, firstStr, sum))
			}
			sum = 0
		}

		if sum == 0 {
			firstStr = curStr
		}
		prevStr = curStr
	}
	// добавление последней строки
	if checkFlgs(sum, fl) {
		*uniqData = append(*uniqData, formatString(fl.c, firstStr, sum))
	}
	return nil
}

func Write(uniqData []string, fl *Flags) error {
	if *fl.output == "default" {
		for i := range uniqData {
			fmt.Printf("%s", uniqData[i])
		}
		return nil
	}
	file, err := os.Create(*fl.output)
	if err != nil {
		return err
	}
	defer file.Close()
	for i := range uniqData {
		file.WriteString(uniqData[i])
	}
	return nil
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
