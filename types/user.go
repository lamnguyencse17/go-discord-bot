package types

type User struct {
	Verified bool
	Username string
	Id string
	Email *string
	Discriminator string
	Bot bool
	Avatar *string
	Session_id string
}