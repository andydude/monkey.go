#!/bin/sh
# SPDX-FileCopyrightText: 2023 Andrew Robbins
# SPDX-License-Identifier: MIT
ANTLR4FMT=$HOME/.local/share/java/antlr4-formatter-standalone-1.2.2-SNAPSHOT.jar
java -jar $ANTLR4FMT "$@"
