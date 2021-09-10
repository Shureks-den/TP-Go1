package uniq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testFlags Flags
var in string = "default"
var out = "default"
var c = false
var d = false
var u = false
var i = false
var s = 0
var f = 0

func TestDefault(t *testing.T) {
	testFlags.output = &out
	testFlags.c = &c
	testFlags.d = &d
	testFlags.u = &u
	testFlags.i = &i
	testFlags.s = &s
	testFlags.f = &f
	text := "test1.txt"
	testFlags.input = &text
	var data []string
	Read(&data, &testFlags)
	assert.Equal(t, []string{"I love music.\n", "\n", "I love music of Kartik.\n", "Thanks.\n", "I love music of Kartik.\n"}, data)
}

func TestCFlag(t *testing.T) {
	c = true
	testFlags.output = &out
	testFlags.c = &c
	testFlags.d = &d
	testFlags.u = &u
	testFlags.i = &i
	testFlags.s = &s
	testFlags.f = &f
	text := "test1.txt"
	testFlags.input = &text
	var data []string
	Read(&data, &testFlags)
	assert.Equal(t, []string{"3 I love music.\n",
		"1 \n",
		"2 I love music of Kartik.\n",
		"1 Thanks.\n",
		"2 I love music of Kartik.\n"}, data)
	c = false
}

func TestDFlag(t *testing.T) {
	d = true
	testFlags.output = &out
	testFlags.c = &c
	testFlags.d = &d
	testFlags.u = &u
	testFlags.i = &i
	testFlags.s = &s
	testFlags.f = &f
	text := "test1.txt"
	testFlags.input = &text
	var data []string
	Read(&data, &testFlags)
	assert.Equal(t, []string{"I love music.\n", "I love music of Kartik.\n", "I love music of Kartik.\n"}, data)
	d = false
}

func TestUFlag(t *testing.T) {
	u = true
	testFlags.output = &out
	testFlags.c = &c
	testFlags.d = &d
	testFlags.u = &u
	testFlags.i = &i
	testFlags.s = &s
	testFlags.f = &f
	text := "test1.txt"
	testFlags.input = &text
	var data []string
	Read(&data, &testFlags)
	assert.Equal(t, []string{"\n", "Thanks.\n"}, data)
	u = false
}

func TestIFlag(t *testing.T) {
	i = true
	testFlags.output = &out
	testFlags.c = &c
	testFlags.d = &d
	testFlags.u = &u
	testFlags.i = &i
	testFlags.s = &s
	testFlags.f = &f
	text := "test2.txt"
	testFlags.input = &text
	var data []string
	Read(&data, &testFlags)
	assert.Equal(t, []string{"I LOVE MUSIC.\n",
		"\n",
		"I love MuSIC of Kartik.\n",
		"Thanks.\n",
		"I love music of kartik.\n"}, data)
	i = false
}

func TestFFlag(t *testing.T) {
	f = 1
	testFlags.output = &out
	testFlags.c = &c
	testFlags.d = &d
	testFlags.u = &u
	testFlags.i = &i
	testFlags.s = &s
	testFlags.f = &f
	text := "test3.txt"
	testFlags.input = &text
	var data []string
	Read(&data, &testFlags)
	assert.Equal(t, []string{"We love music.\n",
		"\n",
		"I love music of Kartik.\n",
		"Thanks.\n"}, data)
	f = 0
}

func TestSFlag(t *testing.T) {
	s = 1
	testFlags.output = &out
	testFlags.c = &c
	testFlags.d = &d
	testFlags.u = &u
	testFlags.i = &i
	testFlags.s = &s
	testFlags.f = &f
	text := "test4.txt"
	testFlags.input = &text
	var data []string
	Read(&data, &testFlags)
	assert.Equal(t, []string{"I love music.\n",
		"\n",
		"I love music of Kartik.\n",
		"We love music of Kartik.\n",
		"Thanks.\n"}, data)
	s = 0
}

func TestMultipleFlags(t *testing.T) {
	s = 1
	f = 1
	d = true
	i = true
	testFlags.output = &out
	testFlags.c = &c
	testFlags.d = &d
	testFlags.u = &u
	testFlags.i = &i
	testFlags.s = &s
	testFlags.f = &f
	text := "test5.txt"
	testFlags.input = &text
	var data []string
	Read(&data, &testFlags)
	assert.Equal(t, []string{"We love music.\n",
		"I love muSic of Kartik.\n"}, data)
	s = 0
	f = 0
	d = false
	i = false
}
