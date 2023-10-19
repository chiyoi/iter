#!/bin/sh
cd $(dirname $(realpath $0)) || return
usage() {
    pwd
    echo "Scripts:"
    echo "$0 tidy"
    echo "    Go mod tidy."
    echo "$0 run_test"
    echo "    Go test."
}

tidy() {
    go mod tidy
}

run_test() {
    go test -v .
}

if test -z "$1" -o -n "$(echo "$1" | grep -Ex '\-{0,2}h(elp)?')"; then
usage
exit
fi

case "$1" in
tidy|run_test) ;;
*)
usage
exit 1
;;
esac

$@

