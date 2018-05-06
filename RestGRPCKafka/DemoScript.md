micro call BankService BankService.CreateAccount '{"name": "Jim", "balance": 100}'

docker run --net=host --rm -it confluentinc/cp-kafka:latest \
kafka-console-consumer --bootstrap-server localhost:29092 --topic BankAccount --from-beginning


micro call BankService BankService.CreateAccount '{"name": "Alex Young", "balance": 50}'
