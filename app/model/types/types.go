package types

/** Telegram types */
type Message struct {
	Update_id uint64      `json:"update_id"`
	Message   MessageBody `json:"message"`
}

type MessageBody struct {
	Date       uint64 `json:"date"`
	Chat       Chat   `json:"chat"`
	Message_id uint64 `json:"message_id"`
	From       User   `json:"from"`
	Text       string `json:"text"`
}

type Chat struct {
	Id         uint64 `json:"id"`
	Last_name  string `json:"last_name"`
	First_name string `json:"first_name"`
	Username   string `json:"username"`
}

type User struct {
	Id         uint64 `json:"id"`
	Last_name  string `json:"last_name"`
	First_name string `json:"first_name"`
	Username   string `json:"username"`
}

type ShortMessage struct {
	Chat_id uint64 `json:"chat_id"`
	Text    string `json:"text"`
}

type Audio struct {
	Chat_id uint64 `json:"chat_id"`
	Audio   string `json:"audio"`
}
