#!/bin/bash
set -e
source dev.env
make $TEST_SUITE
mvn -f java/gorpc/pom.xml install
mvn -f java/vtgate-client/pom.xml install
set +e
