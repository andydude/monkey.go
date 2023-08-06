#!/bin/sh
# SPDX-FileCopyrightText: 2023 Andrew Robbins
# SPDX-License-Identifier: MIT
ANTLR4=$HOME/.local/share/java/antlr-4.13.0-complete.jar
java -cp $ANTLR4:$CLASSPATH org.antlr.v4.gui.TestRig "$@"
