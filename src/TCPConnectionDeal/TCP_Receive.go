package TCPConnectionDeal

import (
	msg "protocol/msg"
	//cgolib "BuffUtil"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"sync"
	//"time"
	"strconv"
	"container/list"
)

var LLock sync.Mutex

var robootIndex int32 = int32(10000)

func AddRoBoot(constru ConnStruct) {
	for i := 0; i < 3; i++ {
		fmt.Println("添加机器人", i)
		userid := strconv.Itoa(int(robootIndex)) //int32(1122)
		nickname := "高贺兵"                        //strconv.Itoa(i)
		level := int32(1)
		fmt.Println("userid: ", userid)

		vec2X := float32(index)
		vec2Z := float32(0)
		index += 10000
		state := msg.State_IDLE
		frameIndex := int32(0)
		spawnPos := msg.Vect2{
			X: &vec2X,
			Y: &vec2Z,
		}
		pl := msg.PlayerInfo{
			Userid:     &userid, //&userid,
			Nickname:   &nickname,
			Level:      &level,
			Status:     &state,
			FrameIndex: &frameIndex,
			SpawnPos:   &spawnPos,
		}
		//================================================================
		//############# 机器人 add roboot ####################
		robootConstruct := ConnStruct{constru.Conn, Accountdata{&robootIndex, nil}, &pl}
		//############# 机器人 add roboot ####################
		//robootConstruct.Playerinfo = &pl
		//######  加入玩家列表 #######
		if _, ok := Clients[*robootConstruct.Playerinfo.Userid]; ok {
			// 有
		} else {
			// 无
			Clients[*robootConstruct.Playerinfo.Userid] = robootConstruct // map 处理方式， 新的
		}
		//######  加入玩家列表 #######
		go TCP_CreateOthers(robootConstruct)
		robootIndex++
	}
}

//##### tcp/ip ######
//var AllplayerList = make([] *msg.PlayerInfo,0)
func TCP_Receive(cmd int16, pb []byte, conn net.Conn) {
	//LLock.Lock()
	//defer LLock.Unlock()
	//fmt.Println("[ TCP ] 处理 DealWithCMD")
	//var (
	//	userid   int32
	//	nickname string
	//	level    int32
	//)
	// cgolib.ExchangeBytes(pb)//字节翻转
	switch cmd {
	case UDP_Enum_UpdateStatus: //接收客户端每66毫秒一次的udp数据
		pbObj := &msg.Rqst_UpdateStatus{} //值的指针
		err := proto.Unmarshal(pb, pbObj)
		if err != nil {
			log.Fatal("[ UDP ] a update status error:", err)
			return
		}
		//fmt.Println("[ UDP ] |||||||||||||||a 收到所有人, 同步坐标 , userid:", *pbObj.Info.Userid)
		//########## 将第一次转入的udp唯一对象存入map中，以作后续广播用 ##########
		if connStructUdp, ok := Clients_UDP[*pbObj.Info.Userid]; ok {
			*connStructUdp.StatusInfo = *pbObj.Info
			//缓存行为帧
			//Frame_Cache = append(Frame_Cache, v)
			//Frame_Cache.Put(v)

			/******
			Try(func() {
				fmt.Println("[ UDP ] @@@@@@@@@@@@@@@@@@@@@@@@@@1 Key Has Found,旧成员", connStructUdp)
				panic("Frame_Cache.PushBack Error....")
				fmt.Println("[ UDP ] @@@@@@@@@@@@@@@@@@@@@@@@@@2 Key Has Found,旧成员", connStructUdp)
			}, func(e interface{}) {
				print(e)
			})*/
			//Frame_Cache.PushBack(connStructUdp)//旧方式
			connStructUdp.Frame_Cache.PushBack(*connStructUdp.StatusInfo)
			//fmt.Println("[ UDP ] @@@@@@@@@@@@@@@@@@@@@@@@@@3 Key Has Found,旧成员 , 新增：", connStructUdp.Frame_Cache.Len(), "待删: ", connStructUdp.Frame_Cache_Remove.Len())
		} else {
			if _, ok := Clients[*pbObj.Info.Userid]; ok { //如果TCP连接存在，则加入到udp列表中，排除已断开的udp反复通信

				constrUDP := ConnStruct_UDP{list.New(), list.New(), conn, pbObj.Info ,nil}
				Clients_UDP[*pbObj.Info.Userid] = constrUDP
				//fmt.Println("[ UDP ] @@@@@@@@@@@@@@@@@@@@@@@@@@ Key Not Found, 新成员 , 新增：", constrUDP.Frame_Cache.Len(), "待删: ", constrUDP.Frame_Cache_Remove.Len())
				//缓存行为帧
				//Frame_Cache = append(Frame_Cache, constrUDP)
				//Frame_Cache.Put(constrUDP)

				//Frame_Cache.PushBack(constrUDP)//旧方式
				constrUDP.Frame_Cache.PushBack(*constrUDP.StatusInfo)
			} else {
				fmt.Println("[ UDP ] what?")
			}
		}
	/*case UpdateStatusCase:  走UDP
	pbObj := &msg.Rqst_UpdateStatus{} //值的指针
	err := proto.Unmarshal(pb, pbObj)
	if err != nil {
		log.Fatal("a update status error:", err)
		return
	}
	//xxxfmt.Println("a 收到所有人, 同步坐标 , userid:", *pbObj.Player.Userid)
	//for _,d := range AllplayerList{
	for _, d := range Clients {
		if d.Playerinfo == nil { // && *d.Playerinfo.Userid == *pbObj.Player.Userid {
			continue
		}

		//xxxfmt.Println("b 收到，同步别人坐标 , a::userid:", *pbObj.Player.Userid, " b::userid:", *d.Playerinfo.Userid)
		if *d.Playerinfo.Userid == *pbObj.Player.Userid { //对指针的值比较
			fmt.Println("ok")
			//更新相关玩家信息记录
			*d.Playerinfo = *pbObj.Player
			//同时下发广播出去
			go UpdateStatus(d)
			break
		}
	}
	*/
	//case TCP_Enum_AddRoboot:

	case TCP_Enum_CreateSelf:

		//============用protobuf 反序列化 ================================================

		pbObj := &msg.Rqst_CreateSelf{}
		err := proto.Unmarshal(pb, pbObj)
		if err != nil {
			log.Fatal("[ TCP ] a unmarshal error: ", err)
			return
		}
		accountid := pbObj.Account
		password := pbObj.Password
		//xxxfmt.Println("////////正确收到protobuf: accountid :", accountid, "password:", password)
		// ===== 直接字节序解读 ======= as below =====================================================
		//
		//	//cgolib.SetBuff_read(pb)
		//	body := &cgolib.ReadBody{pb,0}
		//accountid := body.ReadInt()
		//password := body.ReadInt()
		//	//nickname := body.ReadString(11)//易引起bug，注意
		//	//level := body.ReadInt()

		//========================================================================

		//fmt.Println("//////////////////////////////////////////正确收到stream: nickname :", nickname, "level:", level, "id:", userid)

		// clients[indexOfClient] = conn
		connStruct := ConnStruct{conn, Accountdata{accountid, password}, nil}

		//Clients = append(Clients, connStruct) // list 处理方式 旧的
		//Clients[connStruct.Accountdata.AccountID] = connStruct // map 处理方式， 新的

		//xxxfmt.Println("append: ", connStruct) //, " index: ", IndexOfClient)
		Log("[ TCP ] 剩余连接数:", len(Clients))
		//IndexOfClient++

		//ss := Clients[0].PlayerData.NickName
		//fmt.Println("昵称:" , ss )
		//===================== 推送给客户端 ====================================================================================
		//for _, con := range Clients {
		//
		//	if con.PlayerData.AccountID == connStruct.PlayerData.AccountID {
		//		//go CallbackMatching(Clients[i])//把所收到的玩家列表传出去，生成新玩家
		//		//fmt.Println("()()()()()()()我的userid: " , con.PlayerData.Userid , " index: " , i)
		//		go CreateSelf(con) //把所收到的玩家列表传出去，生成新玩家
		//		continue
		//	}
		//	//time.Sleep(time.Second * 1) xxx
		//	//fmt.Println("()()()()()()()其他人的userid: " , con.PlayerData.Userid , " index: " , i)
		//	//go CreateOthers(con) xxx
		//}
		//================================================================================
		go TCP_CreateSelf(connStruct)
		//================================================================================
		//
		//for i, con := range Clients {
		//	if con.Conn == conn {
		//		//time.Sleep(time.Second*1)//等待3秒发出，以免粘包
		//		go UpdateOtherPlayer(Clients[i])//同步其他玩家数据
		//		break
		//	}
		//}
		//for _,con := range Clients{
		//	 go UpdateOtherPlayer(con)
		//}

		//go UpdateOtherPlayer()//同步其他玩家数据

		//for i, con := range clients {
		//	if con.conn != nil {
		//
		//		fmt.Println("callback  index:", i, " ||  conn: ", con)
		//		// go callbackMatching(clients[i])
		//		go callbackMatching(con)
		//	}
		//}
	case TCP_Enum_HeartBeat: //每15秒才收发一次
		// ==== protobuf as below =====
		pbObj := &msg.Rqst_HeartBeating{}
		err := proto.Unmarshal(pb, pbObj)
		if err != nil {
			log.Fatal("[ TCP ] unmarshal error: ", err)
			return
		}
		statusCode := pbObj.Status
		//playerinfo := pbObj.Player

		// ########## ########## ######### 每次心跳包中， 坐标等数据 记录下来  ???????????? 不这么弄，太差了，表现
		//if v ,ok := Clients[*pbObj.Player.Userid]; ok {
		//	// 有
		//	//if *playerinfo.Userid != "" { //只记下心跳上报的，排除初次上报
		//		*v.Playerinfo.SpawnPos = *playerinfo.SpawnPos //update status 更新角色出生点
		//		fmt.Println("非初次心跳包，同步数据下发回去")
		//	//}
		//}else {
		//	// 无
		//}
		//########## ########## #########
		fmt.Println("[ TCP ] ================================================================================")
		fmt.Println("[ TCP ] 收到心跳包", statusCode, " clients.Count:", len(Clients), " clients_UDP.count:", len(Clients_UDP))
		//for i,d := range  Clients_UDP {
		//	fmt.Println("######### Clients_UDP: i:",i , " d:" , d )
		//}
		// ===== 直接字节序解读 ======= as below =====================================================
		////cgolib.SetBuff_read(pb)
		//body := &cgolib.ReadBody{pb, 0}
		//statusCode := body.ReadInt()
		//fmt.Println("收到心跳包", statusCode, " clients.Count:", len(Clients))
		//=========
		for _, con := range Clients {

			if con.Conn == conn {
				go TCP_HeartBeat(con)
				break
			}
			//	if con.conn != nil {
			//
			//		fmt.Println("callback  index:", i, " ||  conn: ", con)
			//		 //go callbackHeartBeating(clients[i])
			//		go callbackHeartBeating(con)
			//	}
		}

	default:
		fmt.Println("[ TCP ] 接收客户端信息错误!")
	}

}
