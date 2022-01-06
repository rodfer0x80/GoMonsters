#!/bin/bash
PROGNAME="CryptoMonster"
kill $(pgrep $PROGNAME) && echo -e "Program terminated" || echo "Program not running"
