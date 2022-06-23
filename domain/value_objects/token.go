package value_objects

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUUID   string `json:"access_uuid"`
	RefreshUUID  string `json:"refresh_uuid"`
	ATExpires    int64  `json:"at_expires"`
	ATDuration   int    `json:"at_duration"`
	RTExpires    int64  `json:"rt_expires"`
	RTDuration   int    `json:"rt_duration"`
	Username     string `json:"username"`
}
