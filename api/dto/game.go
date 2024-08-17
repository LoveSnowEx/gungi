package dto

type Player struct {
	Name string `json:"name"`
}

type Game struct {
	Players []Player `json:"players"`
}
