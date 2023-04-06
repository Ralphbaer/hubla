#!/bin/bash

CURR_DIR=$PWD
LOGO=$(cat "$CURR_DIR"/shell/logo.txt)

echo CURR_DIR
source "$CURR_DIR"/shell/colors.sh
source "$CURR_DIR"/shell/ascii.sh

echo "${bold}${blue}$LOGO${normal}"

makeCmd() {
  cmd=$1
  for DIR in "$CURR_DIR"/*; do
    FILE="$DIR"/Makefile
    if [ -f "$FILE" ]; then
      if grep -q "$cmd:" "$FILE"; then
        (
          cd "$DIR" || exit
          echo ""
          border "########### Executing ${magenta}make $1${normal} command in package ${bold}${blue}$DIR${normal} ###########"
          make $cmd
        )
        err=$?
        if [ $err -ne 0 ]; then
          echo -e "\n${bold}${red}An error has occurred during test process ${bold}[FAIL]${norma}\n"
          exit 1
        fi
      fi
    fi
  done
}

lint() {
  title1 "STARTING LINT"
  out=$(golint $CURR_DIR/... | tee /dev/tty)
  out_err=$?
  err=0
  if [ $out_err -ne 0 ]; then
    echo -e "\n${bold}${red}An error has occurred during lint process\n"
    err=1
    exit 1
  fi
  if [ -n "$out" ]; then
    echo -e "\n${red}Some lint rules are broken ${bold}[WARNING]${normal}\n"
    err=1
    exit 1
  fi
  if [ ! $err -ne 0 ]; then
    lineOk "\nAll lint rules passed"
  fi
}

format() {
  title1 "Formatting all golang source code"
  go fmt ./...
  lineOk "All go files formatted"
}

checkLogs() {
  title1 "STARTING LOGS ANALYZER"
  allFiles=$(find . -type f -path '*usecase*/*' -name '*.go')
  err=0
  while IFS= read -r path; do
     file=$( awk 'f && f-- {print} /err != nil/ { f = 1 }' $path | column)
     if [[ ! -z "$file" && $path != *"_test"* && $path != *"./common"* ]]; then
          while IFS= read -r line; do
              if [[ "$line" != *"logger.Error"* && "$line" != *"log.Error"* ]]; then
              err=1
              echo $path
              fi
          done <<< "$file"
     fi
  done <<< "$allFiles"
  if [ $err -ne 0 ]; then
      echo -e "\n${yellow}You need to log all errors inside usecases after they are handled. ${bold}[WARNING]${normal}\n"
  fi
}

echo -e "\n\n"
title1 "STARTING PRE-COMMIT SCRIPT"

if [ "$1" == "lint" ]; then
  lint
elif [ "$1" == "format" ]; then
  format
else
  echo "Executing with parameter $1"
  makeCmd "$1"
fi

if [ "$1" != "clean" -a "$1" != "lint" -a "$1" != "checkEnvs" -a "$1" != "docs" -a "$1" != "format" ]; then
  format
  lint
 # checkLogs
fi
