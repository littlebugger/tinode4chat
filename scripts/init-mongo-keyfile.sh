#!/bin/bash

# Copy the keyFile to the appropriate location and set permissions
cp /.mongodb-keyfile/mongo-keyfile /data/configdb/mongo-keyfile
chown mongodb:mongodb /data/configdb/mongo-keyfile
chmod 400 /data/configdb/mongo-keyfile

# Start MongoDB with the specified command
exec mongod --replSet rs0 --bind_ip_all --keyFile /data/configdb/mongo-keyfile