create-topic() {
  docker exec kafka kafka-topics --create --topic $1 --partitions 3 --replication-factor 1 --bootstrap-server localhost:9092
}

create-topic user-request-topic
create-topic request-topic
create-topic request-loan-topic
create-topic response-topic