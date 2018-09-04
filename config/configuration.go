package config

type Config struct {
  Host string `json:"phpipam_server"`
  AppID string `json:"phpipam_appid"`
  User string `json:"phpipam_user"`
  Password string `json:"phpipam_password"`
}
