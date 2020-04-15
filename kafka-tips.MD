TOPIC="my_topic"
ZK="1.2.3.4:2181/my_chroot"
KAFKA="1.2.3.4:9092"
BROKER_ID=1

### Basic admin stuff
./kafka-topics.sh [--zookeeper $ZK|--boostrap-server $KAFKA] --create --topic $TOPIC \
  --partitions 10 \
  --replication-factor 3
./kafka-topics.sh [--zookeeper $ZK|--boostrap-server $KAFKA] --delete --topic $TOPIC
./kafka-topics.sh [--zookeeper $ZK|--boostrap-server $KAFKA] --describe [--topic $TOPIC]
./kafka-topics.sh [--zookeeper $ZK|--boostrap-server $KAFKA] --list

### Manage cfg on broker level
# List cfg
./kafka-configs.sh --bootstrap-server $KAFKA --entity-type brokers --entity-name $BROKER_ID --describe
# Add cfg
./kafka-configs.sh --bootstrap-server $KAFKA --entity-type brokers --entity-name $BROKER_ID --alter --add-config 'compression.type=gzip'
# Del cfg
./kafka-configs.sh --bootstrap-server $KAFKA --entity-type brokers --entity-name $BROKER_ID --alter --delete-config 'compression.type=gzip'

### Manage cfg on topics level
# List cfg
./kafka-configs.sh --zookeeper $ZK --entity-type topics --entity-name $TOPIC --describe
# Add cfg
./kafka-configs.sh --zookeeper $ZK --entity-type topics --entity-name $TOPIC --alter --add-config 'cleanup.policy=compact'
# Del cfg
./kafka-configs.sh --zookeeper $ZK --entity-type topics --entity-name $TOPIC --alter --delete-config 'compression.type=gzip'

