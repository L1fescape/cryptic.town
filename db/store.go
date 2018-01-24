package db

import (
  "encoding/json"
  "errors"

  "github.com/go-redis/redis"
  util "github.com/l1fescape/cryptic.town/util"
)

type Options struct {
  Addr      string
  Password  string
}

type Store struct {
  client *redis.Client
}

type User struct {
  Name  string `json:"Name"`
  Token string `json:"Token"`
  Body  string `json:"Body"`
}

func (u *User) MarshalBinary() ([]byte, error) {
  b, err := json.Marshal(u)
  if err != nil {
    return nil, err
  }
  return b, nil
}

func (u *User) UnmarshalBinary(data []byte) error {
  return json.Unmarshal(data, u)
}

type Home struct {
  Name string
  Body string
}

func New(options *Options) *Store {
  client := redis.NewClient(&redis.Options{
    Addr:     options.Addr,
    Password: options.Password,
    DB:       0,
  })

  return &Store{
    client: client,
  }
}

func (s *Store) getUser(username string) (*User, error) {
  dbUser, err := s.client.HGet("users", username).Result()
  if err != nil {
    return nil, err
  }

  user := &User{}
  user.UnmarshalBinary([]byte(dbUser))

  return user, nil
}

func (s *Store) GetUserHome(username string) (*Home, error) {
  user, err := s.getUser(username)
  if err != nil {
    return nil, err
  }

  return &Home{ Name: user.Name, Body: user.Body }, nil
}

func (s *Store) GetUserToken(username string) (string, error) {
  user, err := s.getUser(username)
  if err != nil {
    return "", err
  }

  return user.Token, nil
}

func (s *Store) ResetUserToken(username string) (string, error) {
  var user *User
  dbuser, err := s.getUser(username)
  if err != nil && err != redis.Nil {
    return "", err
  }

  token := util.GenToken()
  if dbuser == nil {
    user = &User{
      Name: username,
      Token: token,
      Body: "",
    }
  } else {
    user = &User{
      Name: dbuser.Name,
      Token: token,
      Body: dbuser.Body,
    }
  }

  _, err = s.client.HSet("users", username, user).Result()
  if err != nil {
    return "", err
  }

  return token, nil
}

func (s *Store) GetUsers() ([]string, error) {
  users, err := s.client.HKeys("users").Result()
  if err != nil {
    return []string{}, err
  }
  return users, nil
}

func (s *Store) CreateHome(username string, body string) (*Home, error) {
  dbuser, err := s.getUser(username)
  if err != nil && err != redis.Nil {
    return nil, err
  }
  if dbuser != nil {
    return nil, errors.New("user exists")
  }

  user := &User{
    Name: username,
    Token: util.GenToken(),
    Body: body,
  }
  _, err = s.client.HSet("users", username, user).Result()
  if err != nil {
    return nil, err
  }

  return &Home{ Name: username, Body: body }, nil
}

func (s *Store) SetHome(username string, token string, body string) (*Home, error) {
  home, err := s.GetUserHome(username)
  if home == nil {
    return nil, errors.New("user does not exist")
  }
  if err != nil {
    return nil, err
  }

  dbToken, err := s.GetUserToken(username)
  if err != nil {
    return nil, err
  }
  if dbToken != token {
    return nil, errors.New("unauthorized")
  }

  user := &User{
    Name: home.Name,
    Token: dbToken,
    Body: body,
  }
  _, err = s.client.HSet("users", username, user).Result()
  if err != nil {
    return nil, err
  }

  return &Home{ Name: home.Name, Body: body }, nil
}

func (s *Store) CreateOrSetHome(username string, body string) (*Home, error) {
  user, err := s.getUser(username)
  if user != nil {
    return s.SetHome(username, user.Token, body)
  }
  if err != nil && err != redis.Nil {
    return nil, err
  }

  return s.CreateHome(username, body)
}

func (s *Store) Quit() {
  // https://github.com/go-redis/redis/blob/cfed9ab470998d048dcd221ad916cef784e9fa21/commands.go#L310
  // s.client.Quit()
}
