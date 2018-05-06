# How to run this example

1. Start docker containers : Zookeeper, Kafka, Postgresql
    ```
    docker-compose up -d
    ```
2. Start Bank micro service
    ```
    go run BankService/main.go 
    ```
3. Open other terminal
    ```
    micro call BankService BankService.CreateAccount '{"name": "Rocker Goox", "balance": 50} 
    ```
4. Use PGAdmin to open bank database at localhost:5432, user: postgres, pass: @123-
 OR
 
   ```
   $ docker exec -it  -u postgres restgrpckafka_db_1 psql -d bank
   bank=# table accounts;
   bank=# table events
   ```