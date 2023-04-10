#!/bin/bash

# General Config ⚙️
CURR_DIR=$PWD

# Hubla Style 🖌️
source "$CURR_DIR"/shell/variables.sh

# Github Hooks 🪝
GITHUB_PACKAGE_PATH="$CURR_DIR"/github/hooks
GITHUB_PATH="$CURR_DIR"/.git

# Set executable permissions 🆗
chmod +x "$PWD/backend/backend.sh"
chmod +x "$PWD/.git/hooks/pre-commit"

# Backend.sh 🛠
make_backend() {
  source "$CURR_DIR"/backend/backend.sh
}

makeCmd() {
  echo "${bold}${blue}$LOGO${normal}"
 # make_backend
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

checkHooks() {
    err=0
    echo "Checking github hooks..."
    for FILE in "$GITHUB_PACKAGE_PATH"/*; do
      f="$(basename -- $FILE)"
      FILE2="$GITHUB_PATH"/hooks/$f
      if [ -f "$FILE2" ]; then
        if cmp -s "$FILE" "$FILE2"; then
          lineOk "Hook file ${underline}$f${normal} installed and updated"
        else
          lineError "Hook file ${underline}$f${normal} ${red}installed but out-of-date [OUT-OF-DATE]"
          err=1
        fi
      else
        lineError "Hook file ${underline}$f${normal} ${red}not installed [NOT INSTALLED]"
        err=1
      fi
      if [ $err -ne 0 ]; then
        echo -e "\nRun ${bold}make setup-env${normal} to setup your development environment, then try again.\n"
        exit 1
      fi
    done
}

echo -e "\n\n"
title1 "STARTING PRE-COMMIT SCRIPT"

checkHooks

if [ "$1" == "lint" ]; then
  lint
elif [ "$1" == "logs" ]; then
  logs
elif [ "$1" == "format" ]; then
  format
else
  echo "Executing with parameter $1"
  makeCmd "$1"
fi

if [ "$1" != "clean" -a "$1" != "lint" -a "$1" != "format" -a "$1" != "logs" ]; then
  lint
  logs
  format
fi