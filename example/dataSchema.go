package main

type SlackUser struct {
	Ginger_Created int32  `json:"ginger_created"`
	Ginger_Id      int32  `json:"ginger_id" gorm:"primary_key"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	Deleted        bool   `json:"deleted"`
	Status         string `json:"status"`
	Color          string `json:"color"`
	RealName       string `json:"real_name"`
	Tz             string `json:"tz"`
	TzLabel        string `json:"tz_label"`
	TzOffset       int    `json:"tz_offset"`
	//	Profile           ProfileType `json:"profile"`
	ProfileId         int64
	IsAdmin           bool `json:"is_admin"`
	IsOwner           bool `json:"is_owner"`
	IsPrimaryOwner    bool `json:"is_primary_owner"`
	IsRestricted      bool `json:"is_restricted"`
	IsUltraRestricted bool `json:"is_ultra_restricted"`
	IsBot             bool `json:"is_bot"`
	HasFiles          bool `json:"has_files"`
	Has2Fa            bool `json:"has_2fa"`
}

type ProfileType struct {
	Id                 int64
	Image24            string `json:"image_24"`
	Image32            string `json:"image_32"`
	Image48            string `json:"image_48"`
	Image72            string `json:"image_72"`
	Image192           string `json:"image_192"`
	ImageOriginal      string `json:"image_original"`
	RealName           string `json:"real_name"`
	RealNameNormalized string `json:"real_name_normalized"`
	Email              string `json:"email"`
}
