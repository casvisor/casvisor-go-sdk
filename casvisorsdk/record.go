package casvisorsdk

import (
	"encoding/json"
)

type Record struct {
	Id int `xorm:"int notnull pk autoincr" json:"id"`

	Owner       string `xorm:"varchar(100) index" json:"owner"`
	Name        string `xorm:"varchar(100) index" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	Organization string `xorm:"varchar(100)" json:"organization"`
	ClientIp     string `xorm:"varchar(100)" json:"clientIp"`
	User         string `xorm:"varchar(100)" json:"user"`
	Method       string `xorm:"varchar(100)" json:"method"`
	RequestUri   string `xorm:"varchar(1000)" json:"requestUri"`
	Action       string `xorm:"varchar(1000)" json:"action"`

	Object string `xorm:"-" json:"object"`
	// ExtendedUser *User  `xorm:"-" json:"extendedUser"`

	IsTriggered bool `json:"isTriggered"`
}

func (c *Client) AddRecord(record *Record) (bool, error) {
	if record.Owner == "" {
		record.Owner = c.OrganizationName
	}
	if record.Organization == "" {
		record.Organization = c.OrganizationName
	}

	postBytes, err := json.Marshal(record)
	if err != nil {
		return false, err
	}

	resp, err := c.DoPost("add-record", nil, postBytes, false, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}
