To run the program:
You first need to run a server, then you need to join with a client.

To run the server:
First naviagte to the "server" folder in your terminal.
Then run "go run server.go" in this terminal.
You are now hosting a server.

To join with a client:
Open another terminal.
In this terminal navigate to the "client" folder.
Then run "go run client.go"
This client has now joined the server.

To send a message:
From the client terminal type "send" followed by the message you want to send, then hit enter.
You have now send a message to the server which has sent it to all connected clients.

To disconnect:
From the client terminal type "disconnect" then hit enter.
This client is now disconnected from the server.

It is possible to connect with multiple clients at the same time, simply follow the join with a client step once again.
