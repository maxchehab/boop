#!/bin/bash

alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.7.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
alias grun='java org.antlr.v4.gui.TestRig'

mkdir -p lang
cp JSON.g4 lang/JSON.g4
cd lang

antlr4 -Dlanguage=Go JSON.g4