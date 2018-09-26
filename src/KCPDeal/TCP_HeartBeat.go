package KCPDeal

//发送>>>>>>================================================
import (
	//cgolib "BuffUtil"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"protocol/msg"
)

func TCP_HeartBeat(conn ConnStruct) {
	go beatheart(conn)
}
func beatheart(conn ConnStruct) {
	//===== protobuf as below ========================
	state := int32(1)
	//ss := int32(0)
	//
	//xx := int32(0)
	//yy := int32(0)
	//zz := int32(0)
	//tt := msg.Vect3{X:&xx,Y:&yy,Z:&zz}
	//####################################
	//收集所有玩家的当前状态下发同步
	var AllplayerList = make([]*msg.PlayerInfo, 0)
	if 0 < len(Clients) {
		//找到全部玩家
		for _, d := range Clients {
			if d.Playerinfo != nil && *d.Playerinfo.Userid != *conn.Playerinfo.Userid { //排除自己的信息
				//spawnPos := d.Playerinfo.SpawnPos//拉回到坐标
				//state := d.Playerinfo.Status

				//xxxfmt.Println("@@@@@@ other玩家 userid:" , *d.Playerinfo.Userid , "当前玩家 userid:" , *conn.Playerinfo.Userid , "pos:" , *d.Playerinfo.Force.X,*d.Playerinfo.Force.Y,*d.Playerinfo.Force.Z)
				//status := msg.StatusInfo{d.Playerinfo.Userid,&tt,&ss,spawnPos,state}
				//*status.SpawnPos = *d.Playerinfo.SpawnPos
				//*status.Userid = *d.Playerinfo.Userid
				//*status.TargetPos = tt
				//*status.TranSpeed = ss
				//*status.Status = *d.Playerinfo.Status
				AllplayerList = append(AllplayerList, d.Playerinfo) //新角色 ， 加入
				fmt.Println("[ TCP ] ===============================userid: " + *d.Playerinfo.Userid + "================================5更新角色出生点 ，也可用于偏差时的拉回: ======!!!!!")
				fmt.Println(" ###### new pos:", *d.Playerinfo.SpawnPos.X, *d.Playerinfo.SpawnPos.Y)
			}
		}
		//}else {
		//	AllplayerList = nil
	}
	//####################################
	//position := conn.Playerinfo.SpawnPos
	dataPB := &msg.Rspn_HeartBeating{
		Status:     &state,
		PlayerList: AllplayerList, //偏差拉回，在每次心跳返回时
		//PlayerList:nil,//不拉回
	}
	fmt.Println("################## 下发心跳包 ##################")
	//encode protobuf ========================
	pb_encodede, err := proto.Marshal(dataPB)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}
	sendBuff := WriteClient(TCP_Enum_HeartBeat, 1, pb_encodede)
	//xxxlog.Println("beatheart back")
	//==================== 直接处理 as below =======================
	////cgolib.ResetIndex_Write()
	//body := &cgolib.WriteBody{make([]byte,1024),0}
	//body.WriteInt(1)
	//writeBuf := body.GetBuff_Write_HasData()
	//bf := GetCloneArry(writeBuf)
	//Log(" ===== callback 心跳 pb len: ", len(bf), " pb: ", bf)
	//sendBuff := WriteClient(HeartBeat, 1, bf)
	//================================================================================================
	i, err := conn.Conn.Write(sendBuff)
	if err != nil {
		log.Fatal("conn.Write Error", err, i)
		conn.Conn.Close()
		closeudp := Clients_UDP[*conn.Playerinfo.Userid].Conn_udp
		closeudp.Close()
		//Clients = Remove(Clients, conn)
		delete(Clients, *conn.Playerinfo.Userid)
		delete(Clients_UDP, *conn.Playerinfo.Userid)
		fmt.Errorf("################## 删除 udp ##################")
		return
	}
	//time.Sleep(time.Second * 10)
}

/*
//#### udp ######
func CallbackHeartBeating_UDP(conn ConnStruct_UDP) {
	go beatheart_UDP(conn)
}
func beatheart_UDP(conn ConnStruct_UDP)  {
	//===== protobuf as below ========================
	state := int32(1)
	dataPB := &msg.Rspn_HeartBeating{
		Status:&state,
	}

	//encode protobuf ========================
	pb_encodede, err := proto.Marshal(dataPB)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}
	sendBuff := WriteClient(HeartBeatCase, 1, pb_encodede)
	//xxxlog.Println("beatheart back")
	//==================== 直接处理 as below =======================
	////cgolib.ResetIndex_Write()
	//body := &cgolib.WriteBody{make([]byte,1024),0}
	//body.WriteInt(1)
	//writeBuf := body.GetBuff_Write_HasData()
	//bf := GetCloneArry(writeBuf)
	//Log(" ===== callback 心跳 pb len: ", len(bf), " pb: ", bf)
	//sendBuff := WriteClient(HeartBeat, 1, bf)
	//================================================================================================
	i, err := conn.Conn_udp.Write(sendBuff)
	if err != nil {
		log.Fatal("conn.Write Error", err,i)
		conn.Conn_udp.Close()
		//Clients = Remove(Clients, conn)
		delete(Clients,conn.Accountdata.AccountID)
		return
	}
	//time.Sleep(time.Second * 10)
}
*/
