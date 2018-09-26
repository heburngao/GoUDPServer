package KCPDeal

//发送>>>>>>================================================

import (
	//cgolib "BuffUtil"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"protocol/msg"
	"strconv"
	//"net"
)

var index = int(0)

func TCP_CreateSelf(conn ConnStruct) { //, clients map[*int32]ConnStruct) {
	//defer conn.conn.Close() //xxxxxx
	//defer closeConnect(conn.conn) 干嘛关闭?
	// daytime := time.LocalTime().String()

	// opcode := "DB"
	// opcodebyte := []byte(opcode)
	// opcode := binary.BigEndian.Uint16([]byte("DB"))
	// opcodebyte := []byte("DB")
	// opcodebyte2:=Int162Bytes(opcode)
	// length:=Int2Bytes(10)
	// ret:=Short2Bytes(100)
	// ret:=Int162Bytes(100)
	// payload:=binary.BigEndian.Uint16([]byte("OK"))

	// sendBuff := make([]byte,1024)
	// sendBuff = BytesCombine(sendBuff,[]byte("DB"),[]byte(10),[]byte(100),[]byte("OK"))
	fmt.Println("===============================================================新玩家: ======!!!!!")
	//===== protobuf as below ========================

	uuid := int(*conn.Accountdata.AccountID)
	//IndexOfCreate ++
	userid := strconv.Itoa(uuid) //int32(1122)
	nickname := "高贺兵"            //strconv.Itoa(i)
	level := int32(1)
	fmt.Println("userid: ", userid)
	//X := float32(25)
	//Y := float32(0)
	//Z := float32(25)
	//X := int32(0)//(250000)
	//Y := int32(0)
	//Z := int32(0)//(25000)

	//R := float32(0)
	//G := float32(0)
	//B := float32(0)
	//spawnPos := msg.Vect3{
	//	X:&X,
	//	Y:&Y,
	//	Z:&Z,
	//}
	//rotate := msg.Vect3{
	//	X:&R,
	//	Y:&G,
	//	Z:&B,
	//}

	index += 10000
	index1 := index % 30000
	index2 := index / 30000
	vec2X := float32(index1)
	vec2Z := float32(index2 * 10000)
	//出生点坐标，阵列排布

	state := msg.State_IDLE
	//vec2 := msg.Vect2{
	//	X:&vec2X,
	//	Y:&vec2Y,
	//}
	//speed := float32(0)
	frameIndex := int32(0)
	spawnPos := msg.Vect2{
		X: &vec2X,
		Y: &vec2Z,
	}
	pl := msg.PlayerInfo{
		Userid:   &userid, //&userid,
		Nickname: &nickname,
		Level:    &level,
		//Force:&vec,
		Status: &state,
		//Rotate:&rotate,
		//Speed:&speed,
		//InputForce:&vec2,
		//X:&vec2X,
		//Y:&vec2Y,
		FrameIndex: &frameIndex,
		SpawnPos:   &spawnPos,
	}
	//================================================================
	//AllplayerList = append(AllplayerList,&pl) //新角色 ， 加入
	conn.Playerinfo = &pl

	//if v ,ok := AllPlayer[*construct.Playerinfo.Userid]; ok {

	//######  加入玩家列表 #######
	if _, ok := Clients[*conn.Playerinfo.Userid]; ok {
		// 有
		//v = construct.Playerinfo
	} else {
		// 无
		//AllPlayer[*construct.Playerinfo.Userid] = v
		Clients[*conn.Playerinfo.Userid] = conn // map 处理方式， 新的
	}
	//######  加入玩家列表 #######
	//Clients_UDP[*conn.Accountdata.AccountID] = conn // map 处理方式， 新的
	//for _,d := range AllplayerList{
	//	fmt.Println("### 遍历 PlayerInfoList userid: " , *d.Userid , "current userid:", *pl.Userid, " conn::accountid: " , *conn.Accountdata.AccountID ,  " 个数:" ,len(Clients))
	//}
	//xxxfmt.Println("### new player userid: ", *pl.Userid, " conn::accountid: " , *conn.Accountdata.AccountID ,  " 个数:" ,len(Clients))
	//================================================================
	//for _,d := range Clients{
	//	if d.Playerinfo == nil {
	//		continue
	//	}
	//	fmt.Println("@@@  userid: " , *d.Playerinfo.Userid)
	//	if *conn.Accountdata.AccountID == *d.Accountdata.AccountID {
	//		d.Playerinfo = &pl
	//		fmt.Println("@@@ new  userid: " , *d.Playerinfo.Userid)
	//	}
	//}
	dataPB := &msg.Rspn_CreateSelf{
		Player: &pl,
	}

	//encode protobuf ========================
	pb_encodede, err := proto.Marshal(dataPB)
	if err != nil {
		log.Fatal("[ TCP ] marshaling error:", err)
	}
	sendBuff := WriteClient(TCP_Enum_CreateSelf, 1, pb_encodede)
	//=======用上面的protobuff or 直接处理成字节序 ==========================================
	////cgolib.ResetIndex_Write()
	//body := &cgolib.WriteBody{make([]byte,1024),0}
	//
	//str :=  "你好，高大东"+strconv.Itoa(IndexOfClient)
	//Log("====================创角 str:",str," len= ", len(str)  )
	//body.WriteShort(len(str)) //字符串前先写一short的长度，按客户端要求,解读string前，先读一short 长度
	//body.WriteString(str) // 玩家名
	//
	//body.WriteInt(int(conn.PlayerData.AccountID))//9999) //玩家id
	//Log("====================创角 accountid:",conn.PlayerData.AccountID)
	//
	//body.WriteInt(int(1))//1)//玩家等级
	//Log("====================创角 level:",1)
	//
	////TODO 其他玩家的列表
	////clients
	////cgolib.WriteInt(len(clients)-1)//写入其他玩家数
	////for _,d := range clients{
	////	if conn.conn == d.conn{
	////		continue
	////	}
	////
	////	cgolib.WriteInt(int(d.playerData.userid))//userid
	////	cgolib.WriteShort(len(d.playerData.nickName))//名字长度
	////	cgolib.WriteString(d.playerData.nickName)//玩家名
	////	cgolib.WriteInt(int(d.playerData.level))//玩家等级
	////}
	//
	//
	//
	////str = "a session"
	////Log("str len= ", len(str))
	////cgolib.WriteShort(len(str)) //字符串前先写一short的长度，按客户端要求
	////cgolib.WriteString(str)  //session 不需要
	//
	//writeBuf := body.GetBuff_Write_HasData()
	//sendbf := GetCloneArry(writeBuf)
	//Log(" ====================创角 callback pb len: ", len(sendbf), " pb: ", sendbf)
	//=======================
	//sendBuff := WriteClient(CreateNewPlayer, 1, sendbf)
	//==============================================================================================================================
	i, err := conn.Conn.Write(sendBuff) //把玩家的数据列表发给当前连接者
	if err != nil {
		log.Fatal("conn.Write Error", err, i)
		conn.Conn.Close()
		closeudp := Clients_UDP[*conn.Playerinfo.Userid].Conn_udp
		closeudp.Close()
		//Clients = Remove(Clients, conn)
		delete(Clients, *conn.Playerinfo.Userid)
		delete(Clients_UDP, *conn.Playerinfo.Userid)
		return
	}
	//====== 创建完自己的角色， 再处理别的角色 ========
	go TCP_CreateOthers(conn)
	//添加机器人
	//go AddRoBoot(conn)

	// conn.Write(sendBuff)
	// conn.Write(sendBuff)

	// Log(">>>>>>>>>>发送 send back to clients:\n", sendBuff, "\n 发送的pb len: ", len(sendbf), " pb: ", sendbf)
	//time.Sleep(time.Second * 10)

	//conn.conn.Close()
	//conn.conn.Close()
	//conn.Close()
	//处理客户端源源不断地收到字节序为0的问题
	// if len(sendBuff) < 1 {
	// 	fmt.Println("sendBuff is 0 byte  len: ", len(sendBuff))
	// 	conn.Close()
	// 	return
	// }

	// totalbytes:=BytesCombine(opcodebyte2,length,ret,payload)
	// conn.Write(totalbytes)
	// Log(">>>>>>>>>>发送 send back to clients:\n", sendBuff, " pb len: ", len(sendbf)) //string([]byte(daytime)))
	// Log(">>>>>>>>>>发送 send back to clients:\n", cgolib.GetBuff_write())
}

/*
//##### udp #####

func CreateSelf_UDP(conn ConnStruct_UDP ,remoteAddr *net.UDPAddr){//, clients map[*int32]ConnStruct) {

	fmt.Println("======================CreateSelf_UDP=========================================新玩家: ======!!!!!")
	//===== protobuf as below ========================

	//IndexOfCreate ++
	userid := conn.Accountdata.AccountID //int32(1122)
	nickname := "高贺兵"//strconv.Itoa(i)
	level := int32(1)
	fmt.Println("userid: " ,*userid)
	//X := float32(25)
	//Y := float32(0)
	//Z := float32(25)
	//
	//R := float32(0)
	//G := float32(0)
	//B := float32(0)
	//vec := msg.Vect3{
	//	X:&X,
	//	Y:&Y,
	//	Z:&Z,
	//}
	//rotate := msg.Vect3{
	//	X:&R,
	//	Y:&G,
	//	Z:&B,
	//}
	vec2X := float32(0)
	vec2Y := float32(0)
	state := msg.State_IDLE
	//vec2 := msg.Vect2{
	//	X:&vec2X,
	//	Y:&vec2Y,
	//}
	//speed := float32(0)
	pl := msg.PlayerInfo{
		Userid: userid,//&userid,
		Nickname: &nickname,
		Level:&level,
		//Force:&vec,
		Status:&state,
		//Rotate:&rotate,
		//Speed:&speed,
		//InputForce:&vec2,
		X:&vec2X,
		Y:&vec2Y,
	}
	//================================================================
	//AllplayerList = append(AllplayerList,&pl) //新角色 ， 加入
	conn.Playerinfo = &pl
	conn.Addr = remoteAddr
	Clients_UDP[*conn.Accountdata.AccountID] = conn // map 处理方式， 新的
	//for _,d := range AllplayerList{
	//	fmt.Println("### 遍历 PlayerInfoList userid: " , *d.Userid , "current userid:", *pl.Userid, " conn::accountid: " , *conn.Accountdata.AccountID ,  " 个数:" ,len(Clients))
	//}
	//xxxfmt.Println("### new player userid: ", *pl.Userid, " conn::accountid: " , *conn.Accountdata.AccountID ,  " 个数:" ,len(Clients))
	//================================================================
	//for _,d := range Clients{
	//	if d.Playerinfo == nil {
	//		continue
	//	}
	//	fmt.Println("@@@  userid: " , *d.Playerinfo.Userid)
	//	if *conn.Accountdata.AccountID == *d.Accountdata.AccountID {
	//		d.Playerinfo = &pl
	//		fmt.Println("@@@ new  userid: " , *d.Playerinfo.Userid)
	//	}
	//}
	dataPB := &msg.Rspn_CreateSelf{
		Player:&pl,
	}

	//encode protobuf ========================
	pb_encodede, err := proto.Marshal(dataPB)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}
	sendBuff := WriteClient(CreateSelfCase, 1, pb_encodede)
	//=======用上面的protobuff or 直接处理成字节序 ==========================================

	//==============================================================================================================================
	i, err := conn.Conn_udp.WriteToUDP(sendBuff,remoteAddr) //把玩家的数据列表发给当前连接者
	if err != nil {
		log.Fatal("conn.Write Error", err,i)
		conn.Conn_udp.Close()

		//Clients = Remove(Clients, conn)
		delete(Clients_UDP,conn.Accountdata.AccountID)
		return
	}
	//====== 创建完自己的角色， 再处理别的角色 ========
	go CreateOthers_UDP(conn)



}*/
