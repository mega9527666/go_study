package hall_socket_msg_mgr

import (
	"mega/engine/dateutil"
	"mega/engine/logger"
	"mega/engine/socket_connection"
	"mega/proto/go/jhaoproto"

	"google.golang.org/protobuf/proto"
)

func Login_resp(s *socket_connection.Socket_Connection) {

	str := "{'acount':'abcd', 'scrore':1}"

	sendData := &jhaoproto.BaseMsg{
		CmdOrder:   1,
		CmdIndex:   int32(jhaoproto.CmdIndex_Ping),
		TimeUpload: dateutil.Now_UnixMicro(),
		Data:       []byte(str),
	}

	data, err := proto.Marshal(sendData)
	if err != nil {
		logger.Warn("login_resp proto error", err)
	} else {
		s.Send(data)
	}
}
