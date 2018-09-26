package ConnectionDeal

import (
msg "protocol/msg"
//cgolib "BuffUtil"
"fmt"
"github.com/golang/protobuf/proto"
"log"
"net"
"sync"
"container/list"
//"time"

)

var LLock_UDP sync.Mutex



//####### udp ###### 接收客户端操作
var removeFlag int64
var removeCount int64
func UDP_Receive(cmd int16, pb []byte, connUdp net.UDPConn  ,remoteAddr *net.UDPAddr) {
	//LLock_UDP.Lock()
	//defer  LLock_UDP.Unlock()
	fmt.Println("[ UDP ]处理 StatusDeal_UDP , len(Clients_UDP):",len(Clients_UDP) , " len(Clients):", len(Clients))
	//defer func() {
	//	if r := recover(); r != nil {
	//
	//			  fmt.Errorf("panicing: %v", r)
	//
	//	}
	//}()
	//panic("error")
	//var (
	//	userid   int32
	//	nickname string
	//	level    int32
	//)

	// cgolib.ExchangeBytes(pb)//字节翻转


	switch cmd {
	case UDP_Enum_UpdateStatus_Confirm:
		pbObj := &msg.Rqst_UpdateStatus_Confirm{} //值的指针
		err := proto.Unmarshal(pb, pbObj)
		if err != nil {
			log.Fatal("[ UDP ] a confirm status error:", err)
			return
		}
		fmt.Println("[ UDP ] ###+++++++++++++++++++++  1 收到 UDP confirm >> Frameid:",*pbObj.FrameID)
		for { //收到client 回调指令后， 将需要清理的行为队列清理
			if connStructUdp, ok := Clients_UDP[*pbObj.UserID]; ok {

				if 1 > connStructUdp.Frame_Cache_Remove.Len() {
					break
				}
				ele := connStructUdp.Frame_Cache_Remove.Front()
				if ele == nil {
					fmt.Println("[ UDP ] ###+++++++++++++++++++++  2 收到 UDP confirm >>")
					continue
				}
				fmt.Println("[ UDP ] ###+++++++++++++++++++++  3 收到 UDP confirm >> Frameid:",*pbObj.FrameID)
				status, ok := (ele.Value).(msg.StatusInfo)
				if ok {
					if *status.FrameIndex == *pbObj.FrameID { //如果收到的确认中frameId 是当前的，则删除
						connStructUdp.Frame_Cache_Remove.Remove(ele)
						//connStructUdp.Frame_Cache.Remove(ele) //收到CLIENT回调后消理
						fmt.Println("[ UDP ] ###+++++++++++++++++++++##### 如果收到的确认中 userid:",*pbObj.UserID,"frameId:",*pbObj.FrameID," 是当前的，则删除" , " len(Frame_Cache_Remove))",connStructUdp.Frame_Cache_Remove.Len())
					}
				}
			}else {
				break
			}
		}

	case UDP_Enum_UpdateStatus:           //接收客户端每66毫秒一次的udp数据
		pbObj := &msg.Rqst_UpdateStatus{} //值的指针
		err := proto.Unmarshal(pb, pbObj)
		if err != nil {
			log.Fatal("[ UDP ] a update status error:", err)
			return
		}
		fmt.Println("[ UDP ] |||||||||||||||a 收到所有人, 同步坐标 , userid:", *pbObj.Info.Userid)
		//########## 将第一次转入的udp唯一对象存入map中，以作后续广播用 ##########
		if connStructUdp,ok := Clients_UDP[*pbObj.Info.Userid];ok {
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
				fmt.Println("[ UDP ] @@@@@@@@@@@@@@@@@@@@@@@@@@3 Key Has Found,旧成员 , 新增：", connStructUdp.Frame_Cache.Len(), "待删: ", connStructUdp.Frame_Cache_Remove.Len())
		}else {
			if _, ok := Clients[*pbObj.Info.Userid]; ok {//如果TCP连接存在，则加入到udp列表中，排除已断开的udp反复通信

				constrUDP := ConnStruct_UDP{list.New(),list.New(),connUdp, pbObj.Info, remoteAddr}
				Clients_UDP[*pbObj.Info.Userid] = constrUDP
				fmt.Println("[ UDP ] @@@@@@@@@@@@@@@@@@@@@@@@@@ Key Not Found, 新成员 , 新增：", constrUDP.Frame_Cache.Len(), "待删: ", constrUDP.Frame_Cache_Remove.Len())
				//缓存行为帧
				//Frame_Cache = append(Frame_Cache, constrUDP)
				//Frame_Cache.Put(constrUDP)

				//Frame_Cache.PushBack(constrUDP)//旧方式
				constrUDP.Frame_Cache.PushBack(*constrUDP.StatusInfo)
			}else{
				fmt.Println("[ UDP ] what?")
			}
		}
		//########## ########## ################### ########## ################### ########## #########
		//go StatusCast_UDP()//下发包


		// ########## ########## ######### 更新状态后， 坐标等数据 记录下来 ########## ########## #########
		//if connstruct, ok := Clients[*pbObj.Info.Userid]; ok {
		//	// 有
		//	//没办法拉回，如果没有上报的数据
		//	*connstruct.Playerinfo.SpawnPos = *pbObj.Info.SpawnPos //update status 更新角色出生点 ，也可用于偏差时的拉回
		//	fmt.Println("[ UDP ] ===============================================================4更新角色出生点 ，也可用于偏差时的拉回: ======!!!!!")
		//} else {
		//	// 无
		//}
		//########## ########## ################### ########## ################### ########## #########
		//for _,d := range AllplayerList{
		/*for _, d := range Clients_UDP { 先不用，搬到timer中
			if d.Playerinfo == nil { // && *d.Playerinfo.Userid == *pbObj.Player.Userid {
				continue
			}

			fmt.Println("b 收到，同步别人坐标 , a::userid:", *pbObj.Player.Userid, " b::userid:", *d.Playerinfo.Userid)
			if *d.Playerinfo.Userid == *pbObj.Player.Userid { //对指针的值比较
				fmt.Println("[ UDP ]ok")
				//更新相关玩家信息记录
				*d.Playerinfo = *pbObj.Player
				//同时下发广播出去
				go UpdateStatus_UDP(d , remoteAddr)
				break
			}
		}*/
/* udp 只处理状态同步，其他的走tcp
	case CreateSelfCase:

		//============用protobuf 反序列化 ================================================

		pbObj := &msg.Rqst_CreateSelf{}
		err := proto.Unmarshal(pb, pbObj)
		if err != nil {
			log.Fatal("a unmarshal error: ", err)
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
		connStruct := ConnStruct_UDP{connUdp, nil,Accountdata{accountid, password}, nil,nil}
		fmt.Println("remoteip: ",connUdp.RemoteAddr() , "localip:", connUdp.LocalAddr() , "remoteAddr:",*remoteAddr)
		//Clients = append(Clients, connStruct) // list 处理方式 旧的
		//Clients[connStruct.Accountdata.AccountID] = connStruct // map 处理方式， 新的

		//xxxfmt.Println("append: ", connStruct) //, " index: ", IndexOfClient)
		//Log("剩余连接数:", len(Clients_UDP))
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
		go CreateSelf_UDP(connStruct , remoteAddr)
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
	case HeartBeatCase:
		// ==== protobuf as below =====
		pbObj := &msg.Rqst_HeartBeating{}
		err := proto.Unmarshal(pb, pbObj)
		if err != nil {
			log.Fatal("unmarshal error: ", err)
			return
		}
		statusCode := pbObj.Status
		fmt.Println("收到心跳包", statusCode, " clients.Count:", len(Clients_UDP))
		// ===== 直接字节序解读 ======= as below =====================================================
		////cgolib.SetBuff_read(pb)
		//body := &cgolib.ReadBody{pb, 0}
		//statusCode := body.ReadInt()
		//fmt.Println("收到心跳包", statusCode, " clients.Count:", len(Clients))
		//=========
		for _, con := range Clients_UDP {

			if con.Conn_udp == connUdp {
				go CallbackHeartBeating_UDP(con)
				break
			}
			//	if con.conn != nil {
			//
			//		fmt.Println("callback  index:", i, " ||  conn: ", con)
			//		 //go callbackHeartBeating(clients[i])
			//		go callbackHeartBeating(con)
			//	}
		}
*/
	default:
		fmt.Println("[ UDP ] 接收客户端信息错误!")
	}
	//LLock_UDP.Unlock()
}