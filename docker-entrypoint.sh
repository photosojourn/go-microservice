#!/bin/bash
set -e

eval $(ssm-env)

exec go-micrsoservice