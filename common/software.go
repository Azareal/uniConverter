package common

type Software interface {
	Version() string
	DbEngines() []string

	Users() UserList
	Roles() RoleList
	Forums() ForumList
	Topics() TopicList
	Posts() PostList

	AddUser(user *User) error
	AddRole(role *Role) error
	AddForum(forum *Forum) error
	AddTopic(topic *Topic) error
	AddPost(post *Post) error

	Producer() PostProducer
	Consumer() PostConsumer
}

type User struct {
	ID      int
	IsAdmin bool
}

type Role struct {
	ID       int
	IsBanned bool
	IsMod    bool
	IsAdmin  bool
}

type Forum struct {
}

type Topic struct {
}

type Post struct {
}

type UserList interface {
	Next() *User
}

type RoleList interface {
	Next() *Role
}

type ForumList interface {
	Next() *Forum
}

type TopicList interface {
	Next() *Topic
}

type PostList interface {
	Next() *Post
}

// Producer consumes a string and spits out an AST
type PostProducer interface {
	Run(msg string) *Ast
}

// Consumer consumes an AST and spits out a string
type PostConsumer interface {
	Run(tree *Ast) string
}

type SQLUserList struct {
	Table string

	IDRow   string
	NameRow string
	// TODO: Add an interface for the password algorithm, some software (e.g. phpBB) like doing wacky things like eight layers of md5, etc. so it'll be hard to cover this
	PasswordRow string
	SaltRow     string
	EmailRow    string
	RoleRow     string // Should hold the column holding a unique identifier pointing to a Role
	JoinedAtRow string
}
