package room

type Room struct {
	ID       int64  `json:"id"`
	RoomName string `json:"roomName"`
	Title    string `json:"title"`
}
