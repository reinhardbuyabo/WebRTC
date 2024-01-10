# WebRTC

- The 3 main components are:

- [ ] 1. Routes
    - [ ] Fiber
        - [ ] Room
- [ ] 2. Websocket
    - [ ] Websocket / Fiber
        - [ ] Chat
- [ ] 3. Stream
    - [ ] WebRTC
        - [ ] Stream

- All operations of the room are done via the routes, handled by the help of Fiber.
- Streams are going to be powered by WebRTC.
- Rooms and Streams will require WebSockets.

- [ ] Room
    - [ ] Websocket
        - [ ] Chat
        - [ ] Viewer
    - [ ] Stream

Fiber 
`new()`

Config:
1. Serve html files
2. Creates an app instance
3. Cross Origin Resource Sharing (middleware)
4. Logger (middleware)

- Every route has a handler to Create a Room `Room Create`.

```go
// Route: room/create
func RoomCreate(){
    // creates a URL for a new room
    // room / newID
    // new id created with gguid
}
```

Room
- Getting the room or creating it if doesn't exist.
```go
// Route room/{uuid}
func Room(){

}
```

```go
func create() {

}
```

```go
func getRoom(){

}
```

`create()` or `getRoom()` creates a new hub.

| Hub |
| :---: | 
| clients |
| broadcast |
| register | 
| unregister |

- Rendering the room, sending info like:
1. room link
2. websocket addr
3. stream link

=> use layouts (html) to render info.

Room Web Socket
Create connection to all the peers in the room, with web socket
```go
// Route: room/{uuid}/websocket
func createConnection(){

}
```

```go
// sends message to all peers in the room
func broadcast(){

}
```

```go
// adds peers to the channel
func register(){

}
```

```go
// removes peers from the channel
func unregister(){

}
```

- [ ] main.go
    - [ ] server.go
    - [ ] chat.go
    - [ ] room.go
    - [ ] stream.go
    - [ ] welcome.go
    - [ ] client.go
    - [ ] hub.go
    - [ ] peers.go


So what exactly is sent between 2 clients and how is it sent?
1. SDP
A Session Description Protocol is an object containing information about the session connection such as the codec, address, media type, audio and video and so on.

2. ICE Candidate
An ICE candidate is a public IP address and port that could be potentially used as an address that receives data.

- Each peer will typically have multiple ICE candidates that are gathered by making requests to a STUN Server, and they are exchanged between the 2 peers.

## Exchanging Session Description Protocols and ICE Candidates.
