#!/bin/sh
# SPDX-FileCopyrightText: 2023 Andrew Robbins
# SPDX-License-Identifier: MIT
# Run ANTLR4 tool
./antlr4_tool.sh \
    -Dlanguage=Go \
    -no-listener \
    -visitor \
    MonkeyParser.g4 \
    MonkeyLexer.g4

rm -f *.interp *.tokens
