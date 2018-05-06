.PHONY: lang

install:
	go get github.com/antlr/antlr4/runtime/Go/antlr

lang:
	sh lang.sh