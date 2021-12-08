package chat

import "fmt"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

var (
	// messages = []Message{}
	group = Group{name: "test group"}
)

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			// Register the client.
			h.clients[client.name] = client

			// Add the client to the group.
			group.Add(client)

			// Send the messages to the client in the group.
			for _, oldClient := range group.clients {
				if oldClient.name != client.name {
					oldClient.send <- Message{
						From:    client.name,
						To:      oldClient.name,
						Content: fmt.Sprintf("%s joined the group", client.name),
					}
				}
			}

			// Retrive the old messages for current client.
			// for _, message := range messages {
			// 	if message.To == client.name {
			// 		client.send <- message
			// 	}
			// }
		case client := <-h.unregister:
			if _, ok := h.clients[client.name]; ok {
				delete(h.clients, client.name)
				close(client.send)
			}

			for index, oldClient := range group.clients {
				if oldClient.name != client.name {
					oldClient.send <- Message{
						From:    client.name,
						To:      oldClient.name,
						Content: fmt.Sprintf("%s left the group", client.name),
					}
				} else {
					group.clients = append(group.clients[:index], group.clients[index+1:]...)
				}
			}

		case message := <-h.broadcast:
			if message.ChannelType == "group" {
				for _, client := range group.clients {
					if client.name != message.From {
						client.send <- message
					}
				}
			} else if message.ChannelType == "user" {
				if client, ok := h.clients[message.To]; ok {
					client.send <- message
				}
			}
		}
	}
}

func (h *Hub) GetUsers() []string {
	users := []string{}
	for _, client := range h.clients {
		users = append(users, client.name)
	}
	return users
}
