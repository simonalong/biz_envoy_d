#!/usr/bin/env bash
/app/server &
envoy -c /etc/envoy/envoy.yaml --service-cluster biz-d-service
