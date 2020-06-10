#!/usr/bin/env bash

cd languages/mf

antlr -Dlanguage=Go -listener Manifest.g4 -o ../manifest
