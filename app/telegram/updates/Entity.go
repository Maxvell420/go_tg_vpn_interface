package updates

type Entity struct {
	Type            string `json:"type"`
	Offset          *int
	Lenght          *int
	Language        *string
	Url             *string
	Custom_emoji_id *int
}
