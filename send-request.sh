if [ -z $1 ]; then
    echo "Usage: send-request.sh <file>"
    exit 1
fi
docker exec kafka kafka-console-producer --topic user-request-topic --broker-list localhost:9092 < $1