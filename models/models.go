package models

import (
	"net/mail"

	// SMTPServer is used to provide a default SMTP server preference.
	"time"
)

type SMTPServer struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Config represents the configuration information.
type Config struct {
	URL    string     `json:"url"`
	SMTP   SMTPServer `json:"smtp"`
	DBPath string     `json:"dbpath"`
}

// User represents the user model for gophish.
type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Hash     string `json:"-"`
	APIKey   string `json:"api_key" db:"api_key"`
}

// Flash is used to hold flash information for use in templates.
type Flash struct {
	Type    string
	Message string
}

//Campaign is a struct representing a created campaign
type Campaign struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name"`
	CreatedDate   time.Time `json:"created_date" db:"created_date"`
	CompletedDate time.Time `json:"completed_date" db:"completed_date"`
	Template      string    `json:"template"` //This may change
	Status        string    `json:"status"`
	Uid           int64     `json:"-"`
}

type UserCampaigns struct {
	CampaignId int64
	UserId     int64
}

type Result struct {
	Id       int64
	TargetId int64
	Status   string `json:"status"`
}

type CampaignResults struct {
	CampaignId int64
	TargetId   int64
}

type Target struct {
	Id    int64        `json:"-"`
	Email mail.Address `json:"email"`
}