package notion

type UserType string

const (
	UserTypePerson UserType = "person"
	UserTypeBot    UserType = "bot"
)

type User struct {
	Object    string   `json:"object"`
	Id        string   `json:"id"`
	Type      UserType `json:"type,omitempty"`
	Name      string   `json:"name,omitempty"`
	AvatarUrl string   `json:"avatar_url,omitempty"`
	Person    struct {
		Email string `json:"email"`
	} `json:"person,omitempty"`
}
