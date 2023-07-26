#! /bin/bash
# Script to hot reload when backend files are changed

if ! command -v entr
then
    printf "entr(1) is not installlled in this machine, please install it before running this script\n"
fi

# Select all files without:
# Folders: node_modules, dist, src, .git, public, scripts, test, bin
# Files: .gitignore, all .js root files, all .md root files, .gitignore, Makefile, package and package-lock .json, go .mod and .sum
# to exclude a path add: -path PATH_FROM_ROOT -o \
# to exclude a file add: ! -name "FILE_SCHEMA" \
find . -type d \(               \
    -path ./node_modules -o     \
    -path ./dist -o             \
    -path ./src -o              \
    -path ./.git -o             \
    -path ./public -o           \
    -path ./test -o             \
    -path ./bin -o              \
    -path ./scripts \)          \
    -prune -o                   \
    ! -name ".gitignore"        \
    ! -name "Makefile"          \
    ! -name "package-lock.json" \
    ! -name "package.json"      \
    ! -name "*.md"              \
    ! -name "*.js"              \
    ! -name "go.mod"            \
    ! -name "go.sum"            \
    -print | entr -r make dev