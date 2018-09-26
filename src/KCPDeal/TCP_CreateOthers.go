package KCPDeal

import (
	//cgolib "BuffUtil"
	"fmt"
	"log"
	//"strconv"
	"github.com/golang/protobuf/proto"
	"protocol/msg"
)

func TCP_CreateOthers(conn ConnStruct) {

	//if len(Clients) < 2 {
	//	fmt.Println("===============================================================没有其他玩家: ======!!!!!")
	//	return
	//}
	fmt.Println("===============================================================其他玩家: ======!!!!!")
	var AllplayerList = make([]*msg.PlayerInfo, 0)
	//找到全部玩家
	for _, d := range Clients {
		if d.Playerinfo != nil { // && *d.Playerinfo.Userid != *conn.Playerinfo.Userid {
			//xxxfmt.Println("@@@@@@ other玩家 userid:" , *d.Playerinfo.Userid , "当前玩家 userid:" , *conn.Playerinfo.Userid , "pos:" , *d.Playerinfo.Force.X,*d.Playerinfo.Force.Y,*d.Playerinfo.Force.Z)
			AllplayerList = append(AllplayerList, d.Playerinfo) //新角色 ， 加入
		}
	}

	//xxxfmt.Println("@@@@@@ createOthers 玩家，个数:",len(AllplayerList))
	//xxxfor _,dX := range AllplayerList {
	//xxx	fmt.Println("@@@@@@ createOthers other玩家 userid:" , *dX.Userid )//, "当前玩家 userid:" , *conn.Playerinfo.Userid)
	//xxx}
	//if len(AllplayerList) < 1 {
	//	return
	//}

	for _, d := range Clients {
		//if *conn.Accountdata.AccountID == *d.Accountdata.AccountID{
		//	fmt.Println("$$$$$ 排除给自己: " , *d.Accountdata.AccountID , " accid : " , *d.Accountdata.AccountID )
		//	continue//
		//}

		//userid := int32(1123)
		//nickname := "高贺兵2"//strconv.Itoa(i)
		//level := int32(1)
		//
		//X := float32(0)
		//Y := float32(0)
		//Z := float32(0)
		//vec := msg.Vect3{
		//	X:&X,
		//	Y:&Y,
		//	Z:&Z,
		//}
		//state := msg.State_IDLE
		//
		//player := msg.PlayerInfo{
		//	Userid: &userid,
		//	Nickname: &nickname,
		//	Level:&level,
		//	Pos:&vec,
		//	Status:&state,
		//
		//}
		//allplayerList := [] *msg.PlayerInfo{
		//	&player,
		//}

		dataPB := &msg.Rspn_CreateOthers{
			Player: AllplayerList,
		}
		//xxxfmt.Println("@@@@@ 下发其他玩家 " , *d.Accountdata.AccountID)
		//encode protobuf ========================
		pb_encodede, err := proto.Marshal(dataPB)
		if err != nil {
			log.Fatal("marshaling error:", err)
		}
		sendBuff := WriteClient(TCP_Enum_CreateOtherPlayer, 1, pb_encodede)
		////====================直接处理 as below ====================================================
		////func UpdateOtherPlayer() {
		////cgolib.ResetIndex_Write()
		//body := &cgolib.WriteBody{make([]byte,1024),0}
		////TODO 其他玩家的列表
		////clients
		////body.WriteInt(len(Clients)-1)//写入其他玩家数
		//body.WriteInt(len(Clients))//写入玩家数
		//for _,d := range Clients{
		//	if conn.PlayerData.AccountID == d.PlayerData.AccountID{
		//		fmt.Println("$$$$$ 排除自己: " ,  " accountid : " , d.PlayerData.AccountID)
		//		continue// 排除自己
		//	}
		//
		//	fmt.Println(" =====其他玩家 =================从属于：",conn.PlayerData.AccountID,"============================================== userid:" , d.PlayerData.AccountID)
		//	body.WriteInt(int(d.PlayerData.AccountID))//userid
		//	body.WriteShort(len("其他玩家"))//名字长度
		//	body.WriteString("其他玩家")//玩家名
		//	body.WriteInt(int(1))//玩家等级
		//
		//
		//
		//	writeBuf := body.GetBuff_Write_HasData()
		//	sendbf := GetCloneArry(writeBuf)
		//	Log(" =====其他玩家 callback other pb len: ", len(sendbf), " pb: ", sendbf)
		//	sendBuff := WriteClient(CreateOtherPlayer, 1, sendbf)
		//===============================================================
		i, err := d.Conn.Write(sendBuff) //把其他玩家的数据列表发给当前连接者
		if err != nil {
			ch := Clients_UDP[*d.Playerinfo.Userid]
			log.Fatal("conn.Write Error", err, i)
			d.Conn.Close()
			closeudp := ch.Conn_udp //Clients_UDP[*d.Playerinfo.Userid].Conn_udp
			closeudp.Close()
			//Clients = Remove(Clients, d)
			delete(Clients, *d.Playerinfo.Userid)
			delete(Clients_UDP, *d.Playerinfo.Userid)
			fmt.Errorf("################## 删除 udp ##################")
			return
		}
	}

	//writeBuf := cgolib.GetBuff_Write_HasData()
	//sendbf := GetCloneArry(writeBuf)
	//Log(" =====callback pb len: ", len(sendbf), " pb: ", sendbf)
	//sendBuff := WriteClient(CreateOtherPlayer, 1, sendbf)
	////===============================================================
	//_, err := conn.conn.Write(sendBuff)//把其他玩家的数据列表发给当前连接者
	//if err != nil {
	//	log.Fatal("conn.Write Error", err)
	//	conn.conn.Close()
	//
	//	clients = remove(clients, conn)
	//	return
	//}

}

/*
//##### udp #####
func CreateOthers_UDP(conn ConnStruct_UDP) {

	//if len(Clients) < 2 {
	//	fmt.Println("===============================================================没有其他玩家: ======!!!!!")
	//	return
	//}
	fmt.Println("===============================================================其他玩家: ======!!!!!")
	var AllplayerList = make([] *msg.PlayerInfo,0)
	//找到全部玩家
	for _,d := range Clients_UDP{
		if d.Playerinfo != nil {
			//xxxfmt.Println("@@@@@@ other玩家 userid:" , *d.Playerinfo.Userid , "当前玩家 userid:" , *conn.Playerinfo.Userid , "pos:" , *d.Playerinfo.Force.X,*d.Playerinfo.Force.Y,*d.Playerinfo.Force.Z)
			AllplayerList = append(AllplayerList,d.Playerinfo) //新角色 ， 加入
		}
	}
	//xxxfmt.Println("@@@@@@ createOthers 玩家，个数:",len(AllplayerList))
	//xxxfor _,dX := range AllplayerList {
	//xxx	fmt.Println("@@@@@@ createOthers other玩家 userid:" , *dX.Userid )//, "当前玩家 userid:" , *conn.Playerinfo.Userid)
	//xxx}
	if len(AllplayerList) < 1 {
		return
	}

	for _,d := range Clients_UDP{
		//if *conn.Accountdata.AccountID == *d.Accountdata.AccountID{
		//	fmt.Println("$$$$$ 排除给自己: " , *d.Accountdata.AccountID , " accid : " , *d.Accountdata.AccountID )
		//	continue//
		//}

		//userid := int32(1123)
		//nickname := "高贺兵2"//strconv.Itoa(i)
		//level := int32(1)
		//
		//X := float32(0)
		//Y := float32(0)
		//Z := float32(0)
		//vec := msg.Vect3{
		//	X:&X,
		//	Y:&Y,
		//	Z:&Z,
		//}
		//state := msg.State_IDLE
		//
		//player := msg.PlayerInfo{
		//	Userid: &userid,
		//	Nickname: &nickname,
		//	Level:&level,
		//	Pos:&vec,
		//	Status:&state,
		//
		//}
		//allplayerList := [] *msg.PlayerInfo{
		//	&player,
		//}

		dataPB := &msg.Rspn_CreateOthers{
			Player: AllplayerList,
		}
		//xxxfmt.Println("@@@@@ 下发其他玩家 " , *d.Accountdata.AccountID)
		//encode protobuf ========================
		pb_encodede, err := proto.Marshal(dataPB)
		if err != nil {
			log.Fatal("marshaling error:", err)
		}
		sendBuff := WriteClient(CreateOtherPlayer, 1, pb_encodede)
		////====================直接处理 as below ====================================================

		//===============================================================
		i, err := d.Conn_udp.WriteToUDP(sendBuff,conn.Addr) //把其他玩家的数据列表发给当前连接者
		if err != nil {
			log.Fatal("conn.Write Error", err,i)
			d.Conn_udp.Close()

			//Clients = Remove(Clients, d)
			delete(Clients_UDP,d.Accountdata.AccountID)
			return
		}
	}


}*/
