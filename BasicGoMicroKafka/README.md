# Basic PubSub between two microservices using Kafka

This is an example of pubsub via the client/server interfaces.

PubSub at the client/server level works much like RPC but for async comms. It uses the same encoding but 
rather than using the transport interface it uses the broker for messaging. This includes the ability 
to encode metadata into context which is passed through with messages.

## Contents

- srv - contains the message consumer
- cli - contains the message producer

## Step by step instruction

1. Start Consul
    ```
    $ consul agent -dev -ui
    ```
2. Change to directory
    ```
    $ cd BasicGoMicroKafka
    ```
3. Kick start two Docker containers: zookeeper and kafka
    ```
    $ docker-compose up -d
    ```
    Look at detail of [docker-compose.yml](docker-compose.yml)

4. Start micro service that is message consumer
    ```
    $ go run srv/main.go
    ```
5. Start micro service that is message producer
    ```
    $ go run cli/main.go
    ```

6. We can consume message with topic 'Foo' by starting a temporary Docker container that runs kafka-console-consumer
    ```
    $ docker run --net=host --rm -it confluentinc/cp-kafka:latest \
      kafka-console-consumer --bootstrap-server localhost:29092 --topic Foo  
    ```

