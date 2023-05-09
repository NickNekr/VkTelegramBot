package helper

type myUpdate struct {
	Ok     bool `json:"ok"`
	Result struct {
		Id                      int64  `json:"id"`
		IsBot                   bool   `json:"is_bot"`
		FirstName               string `json:"first_name"`
		Username                string `json:"username"`
		CanJoinGroups           bool   `json:"can_join_groups"`
		CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
		SupportsInlineQueries   bool   `json:"supports_inline_queries"`
	} `json:"result"`
}

type message struct {
	MessageID int      `json:"message_id"`
	From      user     `json:"from"`
	Chat      chat     `json:"chat"`
	Date      int      `json:"date"`
	Text      string   `json:"text"`
	Entities  []entity `json:"entities"`
}

type user struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type chat struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type entity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

type callbackQuery struct {
	Id      int      `json:"update_id"`
	From    user     `json:"from"`
	Message *message `json:"message"`
	Data    string   `json:"data"`
}

type update struct {
	UpdateID      int            `json:"update_id"`
	Message       *message       `json:"message"`
	CallbackQuery *callbackQuery `json:"callback_query"`
}

type Response struct {
	Ok     bool     `json:"ok"`
	Result []update `json:"result"`
}

type MessageResponse struct {
	Ok      bool     `json:"ok"`
	Message *message `json:"result"`
}

type BodyMessage struct {
	ChatId      int                  `json:"chat_id"`
	Text        string               `json:"text"`
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

const (
	WELCOME_MESSAGE = "Hi, I'm a bot that doesn't use libraries to work with the telegram Api, but uses their Api directly!\n" +
		"I also have as many as 6 buttons!"
	PRESSED_MESSAGE = "You pressed "
)

var (
	Markup = InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{
		{
			{Text: "Button 1", CallbackData: "Button 1"},
			{Text: "Button 2", CallbackData: "Button 2"},
			{Text: "Button 3", CallbackData: "Button 3"},
			{Text: "Button 4", CallbackData: "Button 4"},
		},
		{
			{Text: "Button 5", CallbackData: "Button 5"},
			{Text: "Button 6", CallbackData: "Button 6"},
		},
	}}
)
