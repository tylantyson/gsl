package gsl

// Object
type Object struct {
	ArrayNumber     int    `json:"ArrayNumber"`
	Checksum        string `json:"Checksum"`
	ContentType     string `json:"ContentType"`
	DateCreated     string `json:"DateCreated"`
	Guid            string `json:"Guid"`
	IsDirectory     bool   `json:"IsDirectory"`
	LastChanged     string `json:"LastChanged"`
	Length          int    `json:"Length"`
	ObjectName      string `json:"ObjectName"`
	Path            string `json:"Path"`
	ReplicatedZones string `json:"ReplicatedZones"`
	ServerId        int    `json:"ServerId"`
	StorageZoneId   int    `json:"StorageZoneId"`
	StorageZoneName string `json:"StorageZoneName"`
	UserId          string `json:"UserId"`
}
