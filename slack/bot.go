package slack

import (
  "github.com/nlopes/slack"

  db "github.com/l1fescape/cryptic.town/db"
  util "github.com/l1fescape/cryptic.town/util"
)

type Bot struct {
  api *slack.Client
  rtm *slack.RTM
  store  *db.Store
}

const (
  privMessageChannelID = "D04GCSJUU"
)

var postMessageParams = slack.PostMessageParameters{
  Username: "cryptic.town",
  IconEmoji: ":house:",
}

func New(token string, store *db.Store) *Bot {
  if token == "" {
    panic("no slack token provided")
  }
  bot := &Bot{
    api: slack.New(token),
    store: store,
  }

  slack.SetLogger(util.Log)

  return bot
}

func (b *Bot) Start() {
  b.rtm = b.api.NewRTM()
  go b.rtm.ManageConnection()

  for msg := range b.rtm.IncomingEvents {
    switch ev := msg.Data.(type) {
    case *slack.MessageEvent:
      err := b.HandleMessage(ev.Msg.Channel, ev.Msg.User, ev.Msg.Text)
      if err != nil {
        util.Log.Printf("slack parse error: %v", err)
      }
      break

    default:
      //
    }
  }
}

func (b *Bot) Stop() {
  if err := b.rtm.Disconnect(); err != nil {
   util.Log.Println(err)
  }
}

func (b *Bot) HandleMessage(channel string, userID string, text string) (error) {
  if channel == privMessageChannelID && userID != "" {
    return b.HandlePrivateMessage(channel, userID, text)
  }
  return nil
}

func (b *Bot) HandlePrivateMessage(channel string, userID string, text string) error {
  util.Log.Printf("slack priv message: %s %s %s", channel, userID, text)
  user, err := b.api.GetUserInfo(userID)
  if err != nil {
    return err
  }

  if text == "token" {
    return b.SendUserToken(channel, user.Name)
  }
  if text == "new token" {
    return b.ResetAndSendUserToken(channel, user.Name)
  }
  if text == "help" {
    return b.SendCommandsList(channel)
  }

  _, err = b.store.CreateOrSetHome(user.Name, text)
  if err != nil {
    return err
  }

  return nil
}

func (b *Bot) SendCommandsList(channel string) error {
  message := `
Commands:
	token					Returns your auth token
	new token			Generates a new token
`
  _, _, err := b.api.PostMessage(channel, message, postMessageParams)
  return err
}

func (b *Bot) SendUserToken(channel string, username string) error {
  token, err := b.store.GetUserToken(username)
  if err != nil {
    return err
  }
  _, _, err = b.api.PostMessage(channel, "`"+token+"`", postMessageParams)
  return err
}

func (b *Bot) ResetAndSendUserToken(channel string, username string) error {
  token, err := b.store.ResetUserToken(username)
  if err != nil {
    return err
  }
  _, _, err = b.api.PostMessage(channel, "`"+token+"`", postMessageParams)
  return err
}
