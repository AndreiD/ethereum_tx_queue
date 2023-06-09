# :loop: Ethereum Tx Queue

The Ethereum Tx Queue is a small utility that allows you to queue transactions to be sent sequentially to an Ethereum RPC endpoint. It provides a simple way to retry failed transactions, and saves successful transactions and errors in separate log files. 

It has disk persistance by boltdb and will continue to process if the server is shutdown while the queue has elements.

:horse: Attention: if the server is down, it won't receive requests to enqueue new transactions, to address this some google pub/sub / AWS SQS or some other mechanism should be used, but there's no reason why would the server go down ...right? :) :rabbit2:

### Configuration

To use the Ethereum Tx Queue, you'll need to configure it by setting the following environment variables:

- `SERVER_PORT`: The port that the Ethereum Tx Queue server will run on (default is 5000).
- `RPC_URL`: The URL of the Ethereum RPC endpoint that transactions will be sent to.
- `MAX_QUEUE_SIZE`: The maximum number of transactions that can be queued at once (default is 10000).
- `RETRY_COUNT`: The number of times to retry a failed transaction (default is 2).
- `RETRY_DELAY_MS`: The number of milliseconds to wait between retries (default is 7000).
- `SLEEP_AFTER_TX_MS`: The number of milliseconds to wait after sending a transaction (default is 5000).
- `LOCAL_DB_PATH`: The path to a local database file used to store transaction data (default is "/tmp/ethtxqueue.db").

Note that the Ethereum Tx Queue server only runs on the local machine by default.

### Usage

To use the Ethereum Tx Queue, you'll need to send a POST request to `http://localhost:<SERVER_PORT>/push` with a JSON payload containing the following data:

```
{
    "rawTx": "<raw transaction data>"
}
```

The `rawTx` field should contain the raw, signed transaction data. Private keys are not required to send transactions through the Ethereum Tx Queue.

When you send a transaction through the Ethereum Tx Queue, it will be added to the queue and sent to the Ethereum RPC endpoint in the order it was received. If the transaction fails, this "program" will retry it according to the `RETRY_COUNT` and `RETRY_DELAY_SEC` settings. If the transaction is successful, it will be logged in the `success.log` file. If an error occurs, it will be logged in the `error.log` file.

After sending a transaction, the Ethereum Tx Queue will wait for `SLEEP_AFTER_TX_MS` milliseconds before accepting another transaction. This can be used to prevent rate-limiting issues with the Ethereum RPC endpoint.

:bulb: The server also has a GET /health endpoint if you want to monitor the health of application


### Security

Raw txes are already signed by the private key so you shouldn't worry about them.
You should secure this API by doing some allowlist

:bulb: eg: in nginx you could do something like this:

~~~~
location /my-endpoint {
    allow 1.2.3.4;
    allow 192.168.1.200;
    deny all;
    # your configuration for the endpoint here
}
~~~~

### Testing

Look under /tests/ folder. Publish the Vote.sol contract to some testnet or just use this one

~~~~
https://testnet.snowtrace.io/address/0x721A8a1B6f313532c74e74C7e5Df3268f9B23917
~~~~


Edit the .env file here for the RPC URL and the Private Key
Private Key is needed for the raw tx.... 

Edit the contract address in the create_raw_tx.js script

Run the "server" on port 5000. Open another terminal

Run it with: 
~~~~
go run .
~~~~

It will auto-send raw tx to your server

### License

The Ethereum Tx Queue is licensed under the Beerware license.