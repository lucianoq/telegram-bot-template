package main

type MessageToSend struct {
	ID   int64 `json:"chat_id"`
	Text string `json:"text"`
}

type Message struct {
	ID                int64 `json:"message_id"`
	Sender            User `json:"from"`
	Unixtime          int64 `json:"date"`
	Chat              Chat `json:"chat"`
	OriginalSender    User `json:"forward_from"`
	OriginalUnixtime  int64 `json:"forward_date"`
	ReplyTo           *Message `json:"reply_to_message"`
	Text              string `json:"text"`
	Entities          []MessageEntity `json:"entities"`
	Audio             Audio `json:"audio"`
	Document          Document `json:"document"`
	Photo             []PhotoSize `json:"photo"`
	Sticker           Sticker `json:"sticker"`
	Video             Video `json:"video"`
	Voice             Voice `json:"voice"`
	Caption           string    `json:"caption"`
	Contact           Contact `json:"contact"`
	Location          Location `json:"location"`
	Venue             Venue `json:"venue"`
	UserJoined        User `json:"new_chat_member"`
	UserLeft          User `json:"left_chat_member"`
	NewChatTitle      string `json:"new_chat_title"`
	NewChatPhoto      []PhotoSize `json:"new_chat_photo"`
	ChatPhotoDeleted  bool `json:"delete_chat_photo"`
	ChatCreated       bool `json:"group_chat_created"`
	SuperGroupCreated bool `json:"supergroup_chat_created"`
	ChannelCreated    bool `json:"channel_chat_created"`
	MigrateTo         int64 `json:"migrate_to_chat_id"`
	MigrateFrom       int64 `json:"migrate_from_chat_id"`
	Pinned            *Message `json:"pinned_message"`
}

type User struct {
	ID        int64    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type Chat struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type File struct {
	FileID   string `json:"file_id"`
	FileSize int64    `json:"file_size"`
	FileName string
}

type Update struct {
	ID      int64     `json:"update_id"`
	Message Message `json:"message"`
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int64  `json:"offset"`
	Length int64 `json:"length"`
	Url    string `json:"url"`
}

type PhotoSize struct {
	File
	Width  int64  `json:"width"`
	Height int64 `json:"height"`
}

type Voice struct {
	File
	Duration int64  `json:"duration"`
	Mime     string  `json:"mime_type"`
}

type Venue struct {
	Location   Location `json:"location"`
	Title      string `json:"title"`
	Address    string `json:"address"`
	Foursquare string `json:"foursquare_id"`
}

type Document struct {
	File
	Thumb    PhotoSize `json:"thumb"`
	FileName string `json:"file_name"`
	Mime     string `json:"mime_type"`
}

type Sticker struct {
	File
	Width  int64  `json:"width"`
	Height int64 `json:"height"`
	Thumb  PhotoSize `json:"thumb"`
}

type Audio struct {
	File
	Duration  int64  `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	Mime      string `json:"mime_type"`
}

type Video struct {
	File
	Width    int64 `json:"width"`
	Height   int64 `json:"height"`
	Duration int64  `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	Mime     string `json:"mime_type"`
}

type Contact struct {
	UserID      int64    `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}