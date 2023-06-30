package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRole int64

const (
	BannedUser UserRole = -10
	SuperUser  UserRole = 0
	Admin      UserRole = 1
	Auditor    UserRole = 2
	Guest      UserRole = 5
	NormalUser UserRole = 10
)

type User struct {
	ID            int64  `gorm:"primaryKey;autoIncrement;not null"`
	Name          string `gorm:"name"`
	EncrptedEmail string `gorm:"index;type:varchar(500) NOT NULL"`
	Role          UserRole
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type HashedEmail struct {
	HashedEmail string `gorm:"primaryKey;type:char(64) NOT NULL"`
}

type Post struct {
	ID          int64 `gorm:"primaryKey;autoIncrement;not null"`
	UserID      int64
	Content     string `gorm:"index:,class:FULLTEXT,option:WITH PARSER ngram;type: varchar(10000) NOT NULL"`
	LikeCount   int64  `gorm:"index"`
	ReplyCount  int64  `gorm:"index"`
	ReportCount int64
	CreatedAt   time.Time      `gorm:"index"`
	UpdatedAt   time.Time      `gorm:"index"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Attention struct {
	UserID int64 `gorm:"index"`
	User   User
	PostID int64 `gorm:"index"`
	Post   Post
}

type Comment struct {
	ID          int64  `gorm:"primaryKey;autoIncrement;not null"`
	NickName    string `gorm:"type:varchar(60) NOT NULL"`
	ReplyTo     int64  `gorm:"index"`
	PostID      int64  `gorm:"index"`
	UserID      int64
	Content     string `gorm:"index:,class:FULLTEXT,option:WITH PARSER ngram;type: varchar(10000) NOT NULL"`
	ReportCount int64
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Report struct {
	ID                int64 `gorm:"primaryKey;autoIncrement;not null"`
	UserID            int64
	ReportedUserID    int64
	ReportedPostID    int64
	ReportedCommentID int64
	Content           string    `gorm:"type: varchar(1000) NOT NULL"`
	CreatedAt         time.Time `gorm:"index"`
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type Banned struct {
	ID        int64 `gorm:"primaryKey;autoIncrement;not null"`
	UserID    int64
	ExpireAt  time.Time
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
