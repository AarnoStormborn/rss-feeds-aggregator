#!/bin/bash

cd /home/harshsingh/Documents/coding/golang/rss-aggregator/sql/schema
export PATH=$PATH:$(go env GOPATH)/bin
goose postgres postgres://postgres:postgres@localhost:5432/rssagg up
echo "Up Migration Completed"