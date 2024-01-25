package room

type Room struct {
	ID       int64  `json:"id"`
	RoomName string `json:"roomName"`
	Title    string `json:"title"`
}

func NewRoom(roomName, title string) *Room {
	return &Room{
		RoomName: roomName,
		Title:    title,
	}
}
