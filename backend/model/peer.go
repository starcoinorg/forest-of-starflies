package model

import "time"

// Peer info
type Peer struct {
	ID             int       `orm:"column(id)"`
	HashID         string    `orm:"column(hash_id)"`
	Address        string    `orm:"column(address)"`
	OnlineDuration int       `orm:"column(online_duration)"`
	Claimed        int       `orm:"column(claimed)"`
	Network        string    `orm:"column(network)"`
	CreatedAt      time.Time `json:"create_at"` // todo beego orm has an issue about auto_now_add
	UpdatedAt      time.Time `json:"update_at" orm:"auto_now;type(datetime)"`
}

// TableUnique muilti key unique
func (p *Peer) TableUnique() [][]string {
	return [][]string{
		{"HashID", "Claimed", "Network"},
	}
}

// TableName peer table name
func (p *Peer) TableName() string {
	return "peer"
}
