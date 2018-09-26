package ConnectionDeal

import (
	//cgolib "BuffUtil"
	"fmt"
	"log"
	//"strconv"
	"github.com/golang/protobuf/proto"
	"protocol/msg"
	"sync"

)





var LockUpdatestatus sync.Mutex
/*
func UpdateStatus(conn ConnStruct) {
	LockUpdatestatus.Lock()
	//if len(Clients) < 1 {
	//	fmt.Println("===============================================================没有玩家: ======!!!!!")
	//	return
	//}

	var AllplayerList = make([] *msg.PlayerInfo,0)
	//找到全部玩家
	for _,d := range Clients{
		if d.Playerinfo != nil {
			//xxxfmt.Println("@@@@@@ other玩家 userid:" , *d.Playerinfo.Userid , "当前玩家 userid:" , *conn.Playerinfo.Userid , "pos:" , *d.Playerinfo.Force.X,*d.Playerinfo.Force.Y,*d.Playerinfo.Force.Z)
			AllplayerList = append(AllplayerList,d.Playerinfo) //新角色 ， 加入
		}
	}



	//var info *msg.PlayerInfo = conn.Playerinfo
	for _, d := range Clients {
		dataPB := &msg.Rspn_UpdateStatus{
			Player: AllplayerList,//info
		}
		fmt.Println("########## ##### updatestatus inputX,Y: ", *d.Playerinfo.X, *d.Playerinfo.Y)
		//*d.Playerinfo = *info
		//d.Playerinfo = info
		//Clients[d.Accountdata.AccountID] = d

		//xxxfmt.Println("########## ##### updatestatus position: ", *d.Playerinfo.Force.X, *d.Playerinfo.Force.Y, *d.Playerinfo.Force.Z)
		//xxxfmt.Println("########## ##### updatestatus rotate: ", *d.Playerinfo.Rotate.X, *d.Playerinfo.Rotate.Y, *d.Playerinfo.Rotate.Z)
		//encode protobuf ========================
		pb_encodede, err := proto.Marshal(dataPB)
		if err != nil {
			log.Fatal("marshaling error:", err)
		}
		sendBuff := WriteClient(UpdateStatusCase, 1, pb_encodede)
		fmt.Println("下发同步坐标")
		i, err := d.Conn.Write(sendBuff) //把其他玩家的数据列表发给当前连接者
		if err != nil {
			log.Fatal("d.Write Error", err, i)
			//d.Conn.Close()
			d.Conn.Close()
			//Clients = Remove(Clients, d)
			delete(Clients, *d.Accountdata.AccountID) //d.Accountdata.AccountID)
			return
		}

	}
	LockUpdatestatus.Unlock()
}
*/
//##### udp ##### 每66毫秒下发同步数据包列表
func UDP_TimerStatusCast(conn ConnStruct_UDP){//, addr *net.UDPAddr) {
	//for {//xxx
	//	if 1 > Frame_Cache.Len() {
	//		//fmt.Println("===========================1没有玩家更新信息: ======!!!!!", " len(Client):", len(Clients), " len(Client_UDP):", len(Clients_UDP))
	//		return
	//	}
	//
	//	//下发，释放缓存行为帧
	//	var AllplayerList = make([] *msg.StatusInfo, 0)
	//		ele := Frame_Cache.Front()
	//		Frame_Cache.Remove(ele)
	//		runItem, ok := (ele.Value).(ConnStruct_UDP)
	//		//if hasRemove == false {
	//		if !ok {
	//			fmt.Println("[ UDP ] ===============================================================3没有玩家更新信息: ======!!!!!")
	//			break
	//		}
	//		fmt.Println("[ UDP ] $$$$$$$$$$$$$$$$$$$$$$$$$$ len(Frame_Cache): ", Frame_Cache.Len(), " len(Client):", len(Clients), " len(Client_UDP):", len(Clients_UDP))
	//		// ########## ########## ######### 更新状态后， 坐标等数据 记录下来 ########## ########## #########
	//		//if connstruct, ok := Clients[*runItem.StatusInfo.Userid]; ok {
	//		//	// 有
	//		//	*connstruct.Playerinfo.SpawnPos = *runItem.StatusInfo.SpawnPos //update status 更新角色出生点
	//		//} else {
	//		//	// 无
	//		//}
	//		//########## ########## ################### ########## ################### ########## #########
	//		AllplayerList = append(AllplayerList, runItem.StatusInfo) //新角色 ， 加入
	//
	//	for _, d := range Clients_UDP {
	//		dataPB := &msg.Rspn_UpdateStatus{
	//			PlayerInfo: AllplayerList, //info
	//		}
	//		fmt.Println("[ UDP ]########## ##### updatestatus inputX,Z: ", *d.StatusInfo.TargetPos.X, *d.StatusInfo.TargetPos.Z, " len(Clients_UDP):", len(Clients_UDP))
	//
	//		//xxxfmt.Println("########## ##### updatestatus position: ", *d.Playerinfo.Force.X, *d.Playerinfo.Force.Y, *d.Playerinfo.Force.Z)
	//		//xxxfmt.Println("########## ##### updatestatus rotate: ", *d.Playerinfo.Rotate.X, *d.Playerinfo.Rotate.Y, *d.Playerinfo.Rotate.Z)
	//		//encode protobuf ========================
	//		pb_encodede, err := proto.Marshal(dataPB)
	//		if err != nil {
	//			log.Fatal("marshaling error:", err)
	//		}
	//		sendBuff := WriteClient(UpdateStatusCase, 1, pb_encodede)
	//		addr := d.Addr
	//		fmt.Println("[ UDP ]下发同步坐标 addr: ", addr)
	//		i, err := d.Conn_udp.WriteToUDP(sendBuff, addr) //把其他玩家的数据列表发给当前连接者
	//		if err != nil {
	//			log.Fatal("[ UDP ] d.Write Error", err, i)
	//			//d.Conn.Close()
	//			d.Conn_udp.Close()
	//			//Clients = Remove(Clients, d)
	//			delete(Clients_UDP, *d.StatusInfo.Userid) //d.Accountdata.AccountID)
	//			fmt.Errorf("[ UDP ] ################## 删除 udp ##################")
	//
	//			closeudp := Clients[*d.StatusInfo.Userid].Conn
	//			closeudp.Close()
	//			delete(Clients, *d.StatusInfo.Userid)
	//			//LockUpdatestatus.Unlock()
	//			return
	//		}
	//
	//	} //end for
	//}//end for xxx

	//暂时不处理 2018.3.8
	//if 0 < conn.Frame_Cache_Remove.Len() {
	//	ele := conn.Frame_Cache_Remove.Front()
	//	conn.Frame_Cache_Remove.Remove(ele)
	//	conn.Frame_Cache.PushBack(ele)
	//}

	//############################
	if 1 > conn.Frame_Cache.Len() {
		//fmt.Println("===========================1没有玩家更新信息: ======!!!!!", " len(Client):", len(Clients), " len(Client_UDP):", len(Clients_UDP))
		return
	}

	//下发，释放缓存行为帧
	var AllplayerList = make([] *msg.StatusInfo, 0)
	for { //aaaa
		if 1 > conn.Frame_Cache.Len() {
			fmt.Println("===========================2没有玩家更新信息: ======!!!!!", " len(Client):", len(Clients), " len(Client_UDP):", len(Clients_UDP))
			break//此处不能为return ,会导致掉线
		}
		//if 0 < conn.Frame_Cache.Len() {

			ele := conn.Frame_Cache.Front()
			//fmt.Println("[ UDP ]是什么:", ele.Value)
			status, ok := (ele.Value).(msg.StatusInfo) //ConnStruct_UDP)//接口强转为ConnStruct_UDP
			if !ok {
				fmt.Println("[ UDP ] ===============================================================3没有玩家更新信息: ======!!!!!")
				return
			}
			//conn.Frame_Cache_Remove.PushBack(status)//暂时不处理 2018.3.8
			conn.Frame_Cache.Remove(ele) //TODO 收到client回调后，消除

			//fmt.Println("[ UDP ] 装载 >>>>> to remove $$$$$$$$$$$$$$$$$$$$$$$$$$ len(conn.Frame_Cache): ", conn.Frame_Cache.Len(), " len(Client):", len(Clients), " len(Client_UDP):", len(Clients_UDP))
			//// ########## ########## ######### 更新状态后， 坐标等数据 记录下来 ########## ########## #########
			if connstruct, ok := Clients[*status.Userid]; ok {
				// 有
				//没办法拉回，如果没有上报的数据
				*connstruct.Playerinfo.SpawnPos = *status.SpawnPos //update status 更新角色出生点 ，也可用于偏差时的拉回
				//fmt.Println("[ UDP ] ===============================userid: "+ *connstruct.Playerinfo.Userid +"================================4更新角色出生点 ，也可用于偏差时的拉回: ======!!!!!", )
				//fmt.Println(" old pos:", *connstruct.Playerinfo.SpawnPos.X,*connstruct.Playerinfo.SpawnPos.Y , " ###### new pos:" ,*status.SpawnPos.X,*status.SpawnPos.Y)
			} else {
				// 无
				//fmt.Println("[ UDP ] 无")
			}
			////########## ########## ################### ########## ################### ########## #########
			AllplayerList = append(AllplayerList, &status) //新角色 ， 加入
		//}
	}//end for aaaa
	//fmt.Println("[ UDP ] xx$$$$$$$$$$$$$$$$$$$$$$$$$$ len(conn.Frame_Cache): ", conn.Frame_Cache.Len(), " len(Client):", len(Clients), " len(Client_UDP):", len(Clients_UDP))
	for _, d := range Clients_UDP {
		if 1 > len(AllplayerList) {
			break
		}
		//if *d.StatusInfo.Userid == *conn.StatusInfo.Userid{
		//	continue//排除自己
		//}

		dataPB := &msg.Rspn_UpdateStatus{
			Info: AllplayerList,
		}
		fmt.Println("[ UDP ] 下发 >>>>> ########## ##### updatestatus inputX,Z: ", *d.StatusInfo.TargetPos.X, *d.StatusInfo.TargetPos.Z, " len(Clients_UDP):", len(Clients_UDP), " len(AllplayerList):", len(AllplayerList))

		//xxxfmt.Println("########## ##### updatestatus position: ", *d.Playerinfo.Force.X, *d.Playerinfo.Force.Y, *d.Playerinfo.Force.Z)
		//xxxfmt.Println("########## ##### updatestatus rotate: ", *d.Playerinfo.Rotate.X, *d.Playerinfo.Rotate.Y, *d.Playerinfo.Rotate.Z)
		//encode protobuf ========================
		pb_encodede, err := proto.Marshal(dataPB)
		if err != nil {
			log.Fatal("marshaling error:", err)
		}
		sendBuff := WriteClient(UDP_Enum_UpdateStatus, 1, pb_encodede)
		addr := d.Addr
		fmt.Println("[ UDP ]下发同步坐标 addr: ", addr)
		i, err := d.Conn_udp.WriteToUDP(sendBuff, addr) //把其他玩家的数据列表发给当前连接者
		if err != nil {
			log.Fatal("[ UDP ] d.Write Error", err, i)
			//d.Conn.Close()
			d.Conn_udp.Close()
			//Clients = Remove(Clients, d)
			delete(Clients_UDP, *d.StatusInfo.Userid) //d.Accountdata.AccountID)
			fmt.Errorf("[ UDP ] ################## 删除 udp ##################")

			closeudp := Clients[*d.StatusInfo.Userid].Conn
			closeudp.Close()
			delete(Clients, *d.StatusInfo.Userid)
			return
		}

	} //end for
	//}//end for xxx
	//############################

}
func RemoveT(slice []ConnStruct_UDP, elems ...ConnStruct_UDP)  ([]ConnStruct_UDP , ConnStruct_UDP , bool) {
	//LockRemove.Lock()
	isInElems := make(map[ConnStruct_UDP]bool)
	//标记当前要remove对象
	for _, elem := range elems {
		isInElems[elem] = true
	}

	var item ConnStruct_UDP
	//重组
	index := 0
	hasRemove := false
	for _, elem := range slice {
		if !isInElems[elem] {//把不在标记中的对名象重组
			slice[index] = elem
			index++
		}else{
			item = elem
			hasRemove = true
		}
	}
	//LockRemove.Unlock()
	return slice[0:index],item ,hasRemove
}
func RemoveI(slc []ConnStruct_UDP, indexRemove int) ([]ConnStruct_UDP, ConnStruct_UDP, bool) {
	hasDelete := false

	var ele ConnStruct_UDP
	ele = slc[indexRemove]

	slc = append(slc[0:indexRemove], slc[indexRemove+1:]...)
	fmt.Printf("[ UDP ] Inside Remove = %s\n", slc)
	if len(slc) > indexRemove {
		hasDelete = true
	}
	return slc,ele,hasDelete
}
