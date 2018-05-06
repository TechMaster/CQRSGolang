docker run --net=host --rm confluentinc/cp-kafka:latest \
kafka-topics --create --topic foo --partitions 1 \
--replication-factor 1 --if-not-exists --zookeeper localhost:32181


docker run --net=host --rm confluentinc/cp-kafka:latest \
kafka-topics --list --zookeeper localhost:32181

docker run --net=host --rm -it confluentinc/cp-kafka:latest \
kafka-console-producer --request-required-acks 1 --broker-list 192.168.1.108:29092 --topic foo

docker run --net=host --rm -it confluentinc/cp-kafka:latest \
kafka-console-consumer --bootstrap-server 192.168.1.108:29092 --topic foo

docker run --net=host --rm -it confluentinc/cp-kafka:latest \
kafka-console-consumer --bootstrap-server localhost:29092 --topic Foo