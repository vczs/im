package dao

type Room struct {
	Rid      string `json:"rid"`
	Uid      string `json:"uid"`
	Number   int    `json:"number"`
	Name     string `json:"name"`
	Describe string `json:"describe"`
	Ct       int64  `json:"ct"`
	Ut       int64  `json:"ut"`
}

func (Room) CollectionName() string {
	return "room"
}