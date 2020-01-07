#!/bin/bash

CNF_FILE="/opt/kafka/config/kafka.properties"
if [ ! -f "${CNF_FILE}" ]; then
  echo "Le fichier ${CNF_FILE} n'est pas present" 
fi

echo "Set variables from env vars (if a variable wasn't set, a default parameter will be used)"
BROKER_ID="${BROKER_ID:-0}"
BROKER_RACK="${BROKER_RACK:-"MyRack"}"
ZOOKEEPER_CONNECT="${ZOOKEEPER_CONNECT:-"localhost:2181"}"
LISTENERS="${LISTENERS:-"PLAINTEXT://localhost:9092"}"
LOGS_DIR="${LOGS_DIR:-"/opt/kafka/data"}"
UNCLEAN_LEADER_ELECTION_FALSE="${UNCLEAN_LEADER_ELECTION_FALSE:-False}"
NUM_PARTITION="${NUM_PARTITION:-1}"

ARRAY_PARAMETER=("BROKER_ID" "BROKER_RACK" "ZOOKEEPER_CONNECT" "LISTENERS" "LOGS_DIR" "UNCLEAN_LEADER_ELECTION_FALSE" "NUM_PARTITION")
ARRAY_VALUE=("${BROKER_ID}" "${BROKER_RACK}" "${ZOOKEEPER_CONNECT}" "${LISTENERS}" "${LOGS_DIR}" "${UNCLEAN_LEADER_ELECTION_FALSE}" "${NUM_PARTITION}")

for (( compteur=0; compteur<${#ARRAY_PARAMETER[@]}; compteur++ )); do
  parameter="${ARRAY_PARAMETER[$compteur]}"
  value="${ARRAY_VALUE[$compteur]}"
  echo "${parameter} => ${value}"
  sed -i "s|$parameter|$value|g" "${CNF_FILE}"
done

/opt/kafka/bin/kafka-server-start.sh /opt/kafka/config/kafka.properties