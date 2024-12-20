# go-0mq

## server 
```~/dev/go-0mq/server$ go run main.go
2024/12/20 15:50:49 listening on port 5555
2024/12/20 15:50:52 Received hello
2024/12/20 15:50:53 Received hello
2024/12/20 15:50:54 Received hello
2024/12/20 15:50:55 Received hello
2024/12/20 15:50:56 Received hello
2024/12/20 15:50:57 Received hello
2024/12/20 15:50:58 Received hello
2024/12/20 15:50:59 Received hello
2024/12/20 15:51:00 Received hello
2024/12/20 15:51:01 Received hello
```

## client 
```
~/dev/go-0mq/client$ go run main.go
Connecting to the server...

2024/12/20 15:50:52 Sending request 0...
2024/12/20 15:50:53 Received reply 0 [world]
2024/12/20 15:50:53 Sending request 1...
2024/12/20 15:50:54 Received reply 1 [world]
2024/12/20 15:50:54 Sending request 2...
2024/12/20 15:50:55 Received reply 2 [world]
2024/12/20 15:50:55 Sending request 3...
2024/12/20 15:50:56 Received reply 3 [world]
2024/12/20 15:50:56 Sending request 4...
2024/12/20 15:50:57 Received reply 4 [world]
2024/12/20 15:50:57 Sending request 5...
2024/12/20 15:50:58 Received reply 5 [world]
2024/12/20 15:50:58 Sending request 6...
2024/12/20 15:50:59 Received reply 6 [world]
2024/12/20 15:50:59 Sending request 7...
2024/12/20 15:51:00 Received reply 7 [world]
2024/12/20 15:51:00 Sending request 8...
2024/12/20 15:51:01 Received reply 8 [world]
2024/12/20 15:51:01 Sending request 9...
2024/12/20 15:51:02 Received reply 9 [world]
```
## Note
Using request reply pattern. Other patterns are possible.
On Ubuntu, install 0mq library. 
```
sudo apt-get install libzmq3-dev
```
