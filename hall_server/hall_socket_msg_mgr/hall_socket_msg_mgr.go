package hall_socket_msg_mgr

import (
	"mega/engine/dateutil"
	"mega/engine/logger"
	"mega/engine/socket_connection"
	"mega/proto/go/jhaoproto"

	"google.golang.org/protobuf/proto"
)

func Login_resp(s *socket_connection.Socket_Connection) {

	// str := "{'acount':'abcd', 'scrore':1}"

	loginData := &jhaoproto.RespLogin{
		UserInfo: &jhaoproto.UserInfo{
			UserId:    9527,
			UserName:  "abcd",
			UserPhoto: "http://aweilrjer",
			Sex:       1,
		},
	}

	loginByte, err := proto.Marshal(loginData)

	sendData := &jhaoproto.BaseMsg{
		CmdOrder:   9527,
		CmdIndex:   int32(jhaoproto.CmdIndex_Login),
		TimeUpload: dateutil.Now_UnixMicro(),
		Data:       loginByte,
	}

	data, err := proto.Marshal(sendData)
	if err != nil {
		logger.Warn("login_resp proto error", err)
	} else {
		s.Send(data)
	}
}
