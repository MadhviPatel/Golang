#!/bin/bash

# This script will ssh into a log00 machine and telnet all the logs of the
# chosen app into graylog via telnet. Everything inside the folder will be
# sent to graylog, there is no filter.
# For this reason, it is recommended to ssh into the log machine and delete
# any unwanted files, like tars and *.gz. These files often have nothing
# important and while graylog will still log them, they will appear as
# garbage that can't be searched for. They will also clutter the elasticsearch
# indexes for no good reason.
#
# Usage:   ./watch TC_NAME APP_NAME
# Example: ./watch tc4 wapi


ENV="$1"
APP="$2"

echo "Tailing $APP in $ENV to Graylog"
ssh -t log00.$ENV tail -n +1  -f /log/swix/$ENV/$APP/* | nc localhost 5555
