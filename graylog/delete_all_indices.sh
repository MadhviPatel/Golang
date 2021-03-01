#!/bin/bash

# This script will delete all the logs from graylog and cleanup its indexes.
# It is recommended to do this after using the watch script, mainly because
# since the watch script downloads all the information everytime it is run,
# if you don't cleanup, next time you log into the app, you will have repeated
# logs.
#
# Usage:   ./delete_all_indices

curl -XDELETE 'http://localhost:9200/*'
curl -XPOST http://localhost:9000/api/system/indices/ranges/rebuild
