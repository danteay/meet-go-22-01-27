#!/usr/bin/env bash

VERSION="v0.0.1"

MESSAGE=""
TIMES=1
START_TIME=1

POSITIONAL_ARGS=()

function help_text {
    echo "Available commands for bashcli: "
    echo
    echo "  -h, --help"
    echo "    Print this help text."
    echo
    echo "  -v, --version"
    echo "    Print the version of bashcli.sh."
    echo
    echo "  -m, --message"
    echo "    Message that will be printed."
    echo
    echo "  -t, --times"
    echo "    Number of times that the message should be printed (default 1)."
    echo
}

function validate_args {
    if [ -z "${MESSAGE}" ]; then
        echo "Error: No message was provided."
        echo
        help_text
        exit 1
    fi
}

function main {
  for (( c=START_TIME; c<=TIMES; c++ ))
  do
     echo "$MESSAGE"
  done
}

if [ $# -eq 0 ]; then
    help_text
    exit 0
fi

while [[ $# -gt 0 ]]; do
  case $1 in
    -m|--message)
      MESSAGE="$2"
      shift # past argument
      shift # past value
      ;;
    -t|--times)
      TIMES="$2"
      TIMES=$((TIMES + 0))
      shift # past argument
      shift # past value
      ;;
    -h|--help)
      help_text
      exit 0
      ;;
    -v|--version)
      echo "$VERSION"
      exit 0
      ;;
    *)
      POSITIONAL_ARGS+=("$1") # save positional arg
      shift # past argument
      ;;
  esac
done

set -- "${POSITIONAL_ARGS[@]}" # restore positional parameters

validate_args

main
