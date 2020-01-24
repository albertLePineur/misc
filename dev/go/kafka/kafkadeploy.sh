#!/bin/bash

ZK_CHROOT='kafka'
ZK_PORT='2181'
ZK_HOST="localhost"
SCRAM_USER="admin"
SCRAM_PWD="azerty123"

./bin/kafka-config.sh --zookeeper ${ZK_HOST}:${ZK_PORT}/${ZK_CHROOT} --alter --add-config "SCRAM-SHA-512=[iterations=8192,password=${SCRAM_PWD}]" --entity-type users --entity-name ${SCRAM_USER}


Server { org.apache.zookeeper.server.auth.DigestLoginModule required user_admin="azerty"; };