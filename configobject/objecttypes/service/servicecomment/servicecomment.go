package servicecomment

import (
	"fmt"
	"git.icinga.com/icingadb/icingadb-main/configobject"
	"git.icinga.com/icingadb/icingadb-main/connection"
	"git.icinga.com/icingadb/icingadb-main/utils"
)

var (
	ObjectInformation configobject.ObjectInformation
	Fields         = []string{
		"id",
		"environment_id",
		"service_id",
		"name_checksum",
		"properties_checksum",
		"name",
		"author",
		"text",
		"entry_type",
		"entry_time",
		"is_persistent",
		"expire_time",
		"zone_id",
	}
)

type ServiceComment struct {
	Id                  string  `json:"id"`
	EnvId               string  `json:"environment_id"`
	ServiceId           string  `json:"service_id"`
	NameChecksum        string  `json:"name_checksum"`
	PropertiesChecksum  string  `json:"checksum"`
	Name                string  `json:"name"`
	Author              string  `json:"author"`
	Text               	string  `json:"text"`
	EntryType           float64	`json:"entry_type"`
	EntryTime           float64 `json:"entry_time"`
	IsPersistent      	bool  	`json:"is_persistent"`
	ExpireTime          float64 `json:"expire_time"`
	ZoneId              string	`json:"zone_id"`
}

func NewServiceComment() connection.Row {
	s := ServiceComment{}

	return &s
}

func (s *ServiceComment) InsertValues() []interface{} {
	v := s.UpdateValues()

	return append([]interface{}{utils.EncodeChecksum(s.Id)}, v...)
}

func (s *ServiceComment) UpdateValues() []interface{} {
	v := make([]interface{}, 0)

	v = append(
		v,
		utils.EncodeChecksum(s.EnvId),
		utils.EncodeChecksum(s.ServiceId),
		utils.EncodeChecksum(s.NameChecksum),
		utils.EncodeChecksum(s.PropertiesChecksum),
		s.Name,
		s.Author,
		s.Text,
		utils.CommentEntryTypes[fmt.Sprintf("%.0f", s.EntryType)],
		s.EntryTime,
		utils.Bool[s.IsPersistent],
		s.ExpireTime,
		utils.EncodeChecksum(s.ZoneId),
	)

	return v
}

func (s *ServiceComment) GetId() string {
	return s.Id
}

func (s *ServiceComment) SetId(id string) {
	s.Id = id
}

func (s *ServiceComment) GetFinalRows() ([]connection.Row, error) {
	return []connection.Row{s}, nil
}

func init() {
	name := "service_comment"
	ObjectInformation = configobject.ObjectInformation{
		ObjectType: name,
		RedisKey: "servicecomment",
		PrimaryMySqlField: "id",
		Factory: NewServiceComment,
		HasChecksum: true,
		BulkInsertStmt: connection.NewBulkInsertStmt(name, Fields),
		BulkDeleteStmt: connection.NewBulkDeleteStmt(name,  "id"),
		BulkUpdateStmt: connection.NewBulkUpdateStmt(name, Fields),
		NotificationListenerType: "servicecomment",
	}
}