package main

import (
	"strings"
)

type Language string

const (
	Undefined Language = "undefined"
	C				   = "c"
	CPP				   = "c++"
	Go                 = "go"
	Python             = "python"
	Rust               = "rust"
	Haskell            = "haskell"
	Javascript         = "javascript"
	Java               = "java"
)

type Code struct {
	Lang Language
	Code string
}

var acceptedPrefixes = map[string]Language{
	"c":       C,
	"c++":     CPP,
	"python":  Python,
	"haskell": Haskell,
}

func ParseCodeBlock(str string) *Code {
	// Se if it is inside a code block.
	lang := Undefined
	if strings.HasPrefix(str, "```") && strings.HasSuffix(str, "```") {
		inner := str[3:len(str) - 3];
		for key, pref := range acceptedPrefixes {
			if strings.HasPrefix(inner, key) {
				lang = pref
				inner = inner[len(key):]
				break
			}
		}
		if lang == Undefined {
			return nil
		}
		return &Code {
			Lang: lang,
			Code: inner,
		}
	}
	return nil;
}






