package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	data := map[[2]string]bool{
		{"colou?", "colo"}:     true,
		{"colou?", "colou"}:    true,
		{"colou?r", "color"}:   true,
		{"colou?r", "colour"}:  true,
		{"colou*r", "color"}:   true,
		{"colou*r", "colour"}:  true,
		{"colou*r", "colouur"}: true,
		{"col.*r", "color"}:    true,
		{"col.*r", "colour"}:   true,
		{"col.*r", "colr"}:     true,
		{"col.*r", "collar"}:   true,
		{"colou+r", "colour"}:  true,
		{"\\.$", "end."}:       true,
		{"h.llo", "hello"}:     true,
		{"h.llo", "hallo"}:     true,
		{"h.llo", "hxllo"}:     true,
		{"h.llo", "hilo"}:      false,
		{"he?llo", "hello"}:    true,
		{"he?llo", "hllo"}:     true,
		{"he?llo", "helo"}:     false,
		{"ab*c", "ac"}:         true,
		{"ab*c", "abc"}:        true,
		{"ab*c", "abbc"}:       true,
		{"ab*c", "abbbbbc"}:    true,
		{"ab+c", "ac"}:         false,
		{"ab+c", "abc"}:        true,
		{"ab+c", "abbc"}:       true,
		{"ab+c", "abbbbbc"}:    true,
		{"ab?c", "ac"}:         true,
		{"ab?c", "abc"}:        true,
		{"ab?c", "abbc"}:       false,
		{"hel+o", "helo"}:      true,
		{"hel+o", "hello"}:     true,
		{"hel+o", "hellllo"}:   true,
		{"col.*r", "colxxxr"}:  true,
		{"col.*r", "colr"}:     true,
		{"col.*r", "collar"}:   true,
		{"col.*r", "colxxr"}:   true,
		{"a\\*b", "a*b"}:       true,
		{"a\\*b", "ab"}:        false,
		{"a\\.b", "a.b"}:       true,
		{"a\\.b", "aab"}:       false,
		{"c\\+d", "c+d"}:       true,
		{"c\\+d", "cd"}:        false,
		{"a\\\\b", "a\\b"}:     true,
		{"a\\\\b", "ab"}:       false,
		{"", ""}:               true,
		{"a", ""}:              false,
		{"", "a"}:              false,
		{".*", ""}:             true,
		{".*", "anything"}:     true,
		{"a*", ""}:             true,
		{"a*", "aaa"}:          true,
		{"a+", ""}:             false,
		{"a+", "a"}:            true,
		{"a+", "aaa"}:          true,
		{"a?", ""}:             true,
		{"a?", "a"}:            true,
		{"a?", "aa"}:           true,
		{"a?b", "b"}:           true,
		{"a?b", "ab"}:          true,
		{"a?b", "aab"}:         true,
		{"a*b*c*", ""}:         true,
		{"a*b*c*", "aaabccc"}:  true,
		{"a*b*c*", "abcc"}:     true,
		{"a.*b", "ab"}:         true,
		{"a.*b", "axyzb"}:      true,
		{"a.*b", "a b"}:        true,
		{"a.*b", "a\nb"}:       true,
		{"^a", "a"}:            true,
		{"^a", "ba"}:           false,
		{"a$", "a"}:            true,
		{"a$", "ab"}:           false,
		{"^a$", "a"}:           true,
		{"^a$", "aa"}:          false,
		{"a\\*", "a*"}:         true,
		{"a\\*", "a"}:          false,
		{"a\\.", "a."}:         true,
		{"a\\.", "ab"}:         false,
		{"\\$", "$"}:           true,
		{"\\$", "a$"}:          true,
		{"a\\?b", "a?b"}:       true,
		{"a\\?b", "ab"}:        false,
	}

	for in, expected := range data {
		res := match(in[0], in[1])
		if res != expected {
			t.Errorf("Regexp : %s Target: %s Expected: %v Result: %v", in[0], in[1], expected, res)
		}
	}
}

func BenchmarkMatchComplex(b *testing.B) {
	reg := "^a.*b$"
	target := "axxxbyyyz"
	for i := 0; i < b.N; i++ {
		match(reg, target)
	}
}

func BenchmarkRegexPackageComplex(b *testing.B) {
	reg := "^a.*b$"
	re, err := regexp.Compile(reg)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	target := "axxxbyyyz"
	for i := 0; i < b.N; i++ {
		re.MatchString(target)
	}
}

func BenchmarkMatchLargeInput(b *testing.B) {
	reg := "^a.*b$"
	target := "a" + string(make([]byte, 10000)) + "b"
	for i := 0; i < b.N; i++ {
		match(reg, target)
	}
}

func BenchmarkRegexPackageLargeInput(b *testing.B) {
	reg := "^a.*b$"
	re, err := regexp.Compile(reg)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	target := "axxxbyyyz"
	for i := 0; i < b.N; i++ {
		re.MatchString(target)
	}
}
