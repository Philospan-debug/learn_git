package main

import (
	"database/sql"
	"fmt"
	"math"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	xerrors "github.com/pkg/errors"
)

var db *sqlx.DB

type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`                            // 帖子id
	AuthorID    int64     `json:"author_id" db:"author_id"`                          // 作者id
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"` // 社区id
	Status      int32     `json:"status" db:"status"`                                // 帖子状态
	Title       string    `json:"title" db:"title" binding:"required"`               // 帖子标题
	Content     string    `json:"content" db:"content" binding:"required"`           // 帖子内容
	CreateTime  time.Time `json:"create_time" db:"create_time"`                      // 帖子创建时间
}

func Init() (err error) {
	dsn := `root:root123@tcp(127.0.0.1:3306)/%s?parseTime=true&loc=Local`
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		xerrors.Wrapf(err, "sqlx.Connect failed")
	}
	return
}

func GetPostById(pid int64) (post Post, err error) {
	sqlStr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?
	`
	err = db.Get(post, sqlStr, pid)
	if err == sql.ErrNoRows {
		xerrors.Wrapf(err, "errnorows:pid:%d not exist", pid)
	}
	return
}
func main() {
	err := Init()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	var post Post
	post, err = GetPostById(math.MaxInt64)
	if err != nil {
		if xerrors.Unwrap(err) == sql.ErrNoRows {
			fmt.Printf("pid:%v not exist\n", math.MaxInt64)
		}
	}
	fmt.Println(post)
}
