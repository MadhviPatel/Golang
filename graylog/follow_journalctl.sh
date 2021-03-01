#!/bin/bash

# This script will ssh into a machine and telnet all the jounalctl logs
# into graylog via telnet. Every system event logged will be sent to graylog,
# there is no filter.
#
# Usage:   ./follow_journalctl TC_NAME MACHINE_NAME SERVICE_NAME
# Example: ./follow_journalctl tc4 wapi00 wapi

ENV="$1"
MACHINE="$2"
APP="${MACHINE%??}"

echo "Following journalctl for $APP in $ENV to Graylog"
ssh -t $MACHINE.$ENV journalctl -u $APP -f | nc localhost 5555
