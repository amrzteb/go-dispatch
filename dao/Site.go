package dao

type Site struct {
	ID 			int `json:"id" orm:"column(id);auto"`
	appkey 		string
	channel 	string
	name 		string
	location 	string
	isp			string
}
