package database

// Line ;

type UserLine struct {
	LineUID   string  `json:"LineId"`
	BotStatus *string `json:"BotStatus"`
	//	UID     string  `json:"UserID"`
}

type Data struct {
	UserID string `json:"userid"`
	LineID string `json:"lineid"`
	CarID  string `json:"carid"`
}
