package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"import_data/config"
)

var DB *sql.DB

func Initialize(cfg *config.Config) error {
	// 先连接到mysql服务器，不指定数据库
	dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
	)

	var err error
	DB, err = sql.Open("mysql", dsnWithoutDB)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	// 创建数据库
	_, err = DB.Exec("CREATE DATABASE IF NOT EXISTS import_data CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")
	if err != nil {
		return err
	}

	// 关闭当前连接
	DB.Close()

	// 重新连接到指定数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	// 创建表
	createTableSQL := `CREATE TABLE IF NOT EXISTS articles (
		id VARCHAR(32) PRIMARY KEY COMMENT 'URL的MD5值',
		title VARCHAR(500) NOT NULL COMMENT '标题',
		content TEXT NOT NULL COMMENT '正文',
		url VARCHAR(500) NOT NULL UNIQUE COMMENT '原始URL',
		publish_date DATE COMMENT '发布日期',
		summary TEXT COMMENT '摘要',
		tags VARCHAR(500) COMMENT '标签，多个标签用逗号分隔',
		author VARCHAR(200) COMMENT '作者',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
	)`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}