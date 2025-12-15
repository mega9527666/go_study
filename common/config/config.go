package config

type envType_struct struct {
	Dev    int
	Test   int
	Online int
}

var EnvironmentType = envType_struct{
	Dev:    1,
	Test:   2,
	Online: 3,
}

var Environment = EnvironmentType.Dev

type serverType_struct struct {
	WebServer      string
	Account_server string
}

var ServerType_List = serverType_struct{
	WebServer:      "[WebServer]",
	Account_server: "[account_server]",
}

var ServerType = ServerType_List.WebServer
