#!/bin/bash

echo "Waiting for MongoDB to start..."

until mongosh --host mongodb --eval "print('MongoDB is up')" &>/dev/null; do
  sleep 2
done

echo "MongoDB started. Initiating replica set..."

mongosh --host mongodb --eval "rs.initiate({_id: 'rs0', members: [{ _id: 0, host: 'mongodb:27017' }]})"

echo "Waiting for the replica set to become PRIMARY..."

until mongosh --host mongodb --eval "rs.isMaster().ismaster" | grep "true" &>/dev/null; do
  sleep 2
done

echo "Replica set is PRIMARY. Initialization complete."