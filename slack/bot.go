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


func New(token string, store *db.Store) *Bot {
  if token == "" {
    err := "no slack token provided"
    util.Log.Println(err)
    panic(err)
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
      // todo implement commands
      util.Log.Printf("%s - %s", ev.Msg.User, ev.Msg.Text)
      user, err := b.api.GetUserInfo(ev.Msg.User)
      if err != nil {
        util.Log.Printf("error getting user %s", err)
        continue
      }
      if ev.Msg.Channel == "D04GCSJUU" {
        _, err := b.store.CreateOrSetHome(ev.Msg.User, user.Name, ev.Msg.Text)
        if err != nil {
          util.Log.Printf("err %v", err)
        }
      }

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
