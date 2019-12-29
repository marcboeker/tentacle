package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Node is a single machine that belongs to an user.
type Node struct {
	ID          int `gorm:"PRIMARY_KEY,AUTO_INCREMENT"`
	Name        string
	Token       string
	SubnetNodes []SubnetNode
}

// A Subnet Subnets multiple nodes under a single name and C.I.D.R.
type Subnet struct {
	ID   int `gorm:"PRIMARY_KEY,AUTO_INCREMENT"`
	Name string
	CIDR string
}

// SubnetNode connects a node to a subnet.
type SubnetNode struct {
	SubnetID int
	NodeID   int
	IP       string
}

// DB represents a database connection.
type DB struct {
	*gorm.DB
}

// Close closes the underlying database connection.
func (db *DB) Close() error {
	return db.DB.Close()
}

// NewDB creates a database instance and creates all the missing tables.
func NewDB(db *gorm.DB) *DB {
	dbi := DB{DB: db}
	db.AutoMigrate(&Node{})
	db.AutoMigrate(&Subnet{})
	db.AutoMigrate(&SubnetNode{})
	return &dbi
}

// NewSQLiteDB creates a SQLite database instance.
func NewSQLiteDB(path string) (*DB, error) {
	sqlite3DB, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return NewDB(sqlite3DB), nil
}
