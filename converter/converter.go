package converter

import "../common"
import "../systems/gosora"

func Lookup(software string, version string) (soft common.Software, exists bool) {
	switch software {
	case "gosora":
		return gosora.Lookup(version)
	}
	return soft, false
}

type Converter struct {
	From common.Software
	To   common.Software
}

func NewConverter() *Converter {
	return &Converter{}
}

func (conv *Converter) From(soft common.Software) {
	conv.From = soft
}

func (conv *Converter) To(soft common.Software) {
	conv.To = soft
}

func (conv *Converter) Convert() error {
	fromRoles := conv.From.Roles()
	for {
		role := fromRoles.Next()
		if role == nil {
			break
		}
		err := conv.To.AddRole(role)
		if err != nil {
			return err
		}
	}

	fromUsers := conv.From.Users()
	for {
		user := fromUsers.Next()
		if user == nil {
			break
		}
		err := conv.To.AddUser(user)
		if err != nil {
			return err
		}
	}

	fromForums := conv.From.Forums()
	for {
		forum := fromForums.Next()
		if forum == nil {
			break
		}
		err := conv.To.AddForum(forum)
		if err != nil {
			return err
		}
	}

	fromTopics := conv.From.Topics()
	for {
		topic := fromTopics.Next()
		if topic == nil {
			break
		}
		err := conv.To.AddTopic(topic)
		if err != nil {
			return err
		}
	}

	fromPosts := conv.From.Posts()
	for {
		post := fromPosts.Next()
		if post == nil {
			break
		}
		err := conv.To.AddPost(post)
		if err != nil {
			return err
		}
	}
	return nil
}
