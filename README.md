# Ethereum Tx Queue

Small utility that queues transactions, to be sent sequantially to an RPC.

It will retry the transaction for RETRY_COUNT with RETRY_DELAY

in case there's an error, it wiill write it to error.txt file

Note:

- the raw tx is needed
- rawtx is already signed, no private keys are required for this to send the tx

# License: Beerware