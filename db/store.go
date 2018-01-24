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

func (s *Store) GetUsers() ([]string, error) {
  users, err := s.client.HKeys("users").Result()
  if err != nil {
    return []string{}, err
  }
  return users, nil
}

func (s *Store) CreateHome(userID string, username string, body string) (*Home, error) {
  dbuser, err := s.getUser(username)
  if err != nil && err != redis.Nil {
    return nil, err
  }
  if dbuser != nil {
    return nil, errors.New("user exists")
  }

  user := &User{
    Name: username,
    Token: userID,
    Body: body,
  }
  _, err = s.client.HSet("users", username, user).Result()
  if err != nil {
    return nil, err
  }

  return &Home{ Name: username, Body: body }, nil
}

func (s *Store) SetHome(userID string, username string, body string) (*Home, error) {
  home, err := s.GetUserHome(username)
  if home == nil {
    return nil, errors.New("user does not exist")
  }
  if err != nil {
    return nil, err
  }

  token, err := s.GetUserToken(username)
  if err != nil {
    return nil, err
  }
  if token != userID {
    return nil, errors.New("user token incorrect")
  }

  user := &User{
    Name: home.Name,
    Token: token,
    Body: body,
  }
  _, err = s.client.HSet("users", username, user).Result()
  if err != nil {
    return nil, err
  }

  return &Home{ Name: home.Name, Body: body }, nil
}

func (s *Store) CreateOrSetHome(userID string, userName string, body string) (*Home, error) {
  user, err := s.getUser(userName)
  if user != nil {
    return s.SetHome(userID, userName, body)
  }
  if err != nil && err != redis.Nil {
    return nil, err
  }

  return s.CreateHome(userID, userName, body)
}

func (s *Store) Quit() {
  // s.client.Quit()
}
