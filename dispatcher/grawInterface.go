package dispatcher

import (
	"errors"
	"github.com/turnage/graw/botfaces"
	"github.com/turnage/graw/reddit"
)

type GrawUnionType struct {
	botfaces.Loader
	botfaces.Tearer
	botfaces.PostHandler
	botfaces.CommentHandler
	botfaces.MessageHandler
	botfaces.CommentReplyHandler
	botfaces.MentionHandler
	botfaces.UserHandler
	loaderHandlers [](func()error)
	tearerHandlers [](func())
	postHandlers [](func(*reddit.Post)error)
	commentHandlers [](func(*reddit.Comment)error)
	messageHandlers [](func(*reddit.Message)error)
	commentReplyHandlers [](func(*reddit.Message)error)
	mentionHandlers [](func(*reddit.Message)error)
	userPostHandlers [](func(*reddit.Post)error)
	userCommentHandlers [](func(*reddit.Comment)error)
}


// Aggregates several post handlers
func GetUnionType(ifaces ...interface{}) (* GrawUnionType) {
	ret := GrawUnionType{
		loaderHandlers: [](func()error){},
		tearerHandlers: [](func()){},
		postHandlers: [](func(*reddit.Post)error){},
		commentHandlers: [](func(*reddit.Comment)error){},
		messageHandlers: [](func(*reddit.Message)error){},
		commentReplyHandlers: [](func(*reddit.Message)error){},
		mentionHandlers: [](func(*reddit.Message)error){},
		userPostHandlers: [](func(*reddit.Post)error){},
		userCommentHandlers: [](func(*reddit.Comment)error){},
	}
	for _,iface := range ifaces {
		if c,ok :=  iface.(botfaces.Loader); ok {
			ret.loaderHandlers = append(ret.loaderHandlers, c.SetUp)
		}
		if c,ok :=  iface.(botfaces.Tearer); ok {
			ret.tearerHandlers = append(ret.tearerHandlers, c.TearDown)
		}
		if c,ok :=  iface.(botfaces.PostHandler); ok {
			ret.postHandlers = append(ret.postHandlers, c.Post)
		}
		if c,ok :=  iface.(botfaces.CommentHandler); ok {
			ret.commentHandlers = append(ret.commentHandlers, c.Comment)
		}
		if c,ok :=  iface.(botfaces.MessageHandler); ok {
			ret.messageHandlers = append(ret.messageHandlers, c.Message)
		}
		if c,ok :=  iface.(botfaces.CommentReplyHandler); ok {
			ret.commentReplyHandlers = append(ret.commentReplyHandlers, c.CommentReply)
		}
		if c,ok :=  iface.(botfaces.MentionHandler); ok {
			ret.mentionHandlers = append(ret.mentionHandlers, c.Mention)
		}
		if c,ok :=  iface.(botfaces.UserHandler); ok {
			ret.userPostHandlers = append(ret.userPostHandlers, c.UserPost)
			ret.userCommentHandlers = append(ret.userCommentHandlers, c.UserComment)
		}
	}

	return &ret
}

func errorsToError(errs []error ) error{
	var ret string = ""
	for _, err := range errs {
		if (err != nil) {
			ret += "," + err.Error();
		}
	}
	if (ret == "") {
		return nil
	} else {
		return errors.New(ret)
	}
}

func (union *GrawUnionType) SetUp() error {
	errors := []error{}
	for _, handler := range union.loaderHandlers {
		if (handler != nil){
			errors = append(errors, handler())
		}
	}
	return errorsToError(errors)
}

func (union *GrawUnionType) TearDown() error {
	errors := []error{}
	for _, handler := range union.tearerHandlers {
		if (handler != nil){
			errors = append(errors, union.TearDown())
		}
	}
	return errorsToError(errors)
}

func (union *GrawUnionType) Post(p *reddit.Post) error {
	errors := []error{}
	for _, handler := range union.postHandlers {
		if (handler != nil){
			errors = append(errors, handler(p))
		}
	}
	return errorsToError(errors)
}

func (union *GrawUnionType) Comment(c *reddit.Comment) error {
	errors := []error{}
	for _, handler := range union.commentHandlers {
		if (handler != nil){
			errors = append(errors, handler(c))
		}
	}
	return errorsToError(errors)
}

func (union *GrawUnionType) Message(m *reddit.Message) error {
	errors := []error{}
	for _, handler := range union.messageHandlers {
		if (handler != nil){
			errors = append(errors, handler(m))
		}
	}
	return errorsToError(errors)
}

func (union *GrawUnionType) CommentReply(m *reddit.Message) error {
	errors := []error{}
	for _, handler := range union.commentReplyHandlers {
		if (handler != nil){
			errors = append(errors, handler(m))
		}
	}
	return errorsToError(errors)
}

func (union *GrawUnionType) Mention(m *reddit.Message) error {
	errors := []error{}
	for _, handler := range union.mentionHandlers {
		if (handler != nil){
			errors = append(errors, handler(m))
		}
	}
	return errorsToError(errors)
}

func (union *GrawUnionType) UserPost(p *reddit.Post) error {
	errors := []error{}
	for _, handler := range union.userPostHandlers {
		if (handler != nil){
			errors = append(errors, handler(p))
		}
	}
	return errorsToError(errors)
}

func (union *GrawUnionType) UserComment(c *reddit.Comment) error {
	errors := []error{}
	for _, handler := range union.userCommentHandlers {
		if (handler != nil){
			errors = append(errors, handler(c))
		}
	}
	return errorsToError(errors)
}
