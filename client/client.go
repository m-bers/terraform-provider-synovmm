package client

type Actor struct {
	Account        string `json:"account"`
	Device_ID      string `json:"device_id"`
	IK_Message     string `json:"ik_message"`
	Is_Portal_Port string `json:"is_portal_port"`
	SID            string `json:"sid"`
	Synotoken      string `json:"synotoken"`
}

type PingResponse struct {
	Actor Actor `json:"data"`
}
