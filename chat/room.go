package main 
type room struct {
    forward chan []byte
    join chan *client
    leave chan *client
    clients map[*client]bool
}

func newRoom() *room {
    return &room{
        forward: make(chan []byte),
        join   : make(chan *client),
        leave  : make(chan *client),
        clients: make(ma[*client]bool),
    }
}

func (r *room) run() {
    for {
        select {
        case client := <-r.join
            r.clients[client] = true
        case client := <-r.leave
            delete(r.clients, client)
            close(client.send)
        case msg := <-r.forward:
            for client := range r.clients {
                client.send <- msg
            }
        }
    }
}

const (
    socketBufferSize  = 1024
    messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, 
    WriteBufferSize: socketBufferSize}
func (r *romm) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    socket, err := upgrader.Upgrade(w, req, nil)
    if err != nil {
        log.Fatal("ServeHTTP:", err)
        return
    }
    client := &client {
        socket: socket,
        send: make(chan []byte, messageBufferSize),
        room: r,
    }
    r.join <- client
    defer func() {r.leave <- client} ()
    go client.write()
    client.read()
}