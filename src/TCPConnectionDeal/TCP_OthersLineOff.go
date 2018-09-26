package TCPConnectionDeal

//其他玩家的离线

import (
	//cgolib "BuffUtil"
	"fmt"
	"log"
	//"strconv"
	"github.com/golang/protobuf/proto"
	"protocol/msg"
	//"container/list"
)

func TCP_OthersLineOff(construct ConnStruct) {
	//if len(Clients) < 1 {
	//	fmt.Println("===============================================================没有玩家: ======!!!!!")
	//	return
	//}
	var leveOffList = make([]*msg.PlayerInfo, 0)
	for _, d := range Clients {
		if *d.Playerinfo.Userid == *construct.Playerinfo.Userid { // && *d.Playerinfo.Userid != *construct.Playerinfo.Userid{
			leveOffList = append(leveOffList, d.Playerinfo) //离线角色
			fmt.Println("### 离线角色 ,accid: ", *d.Accountdata.AccountID)
		}
	}

	fmt.Println("A###### 有离线的,tcpClient剩余: ", len(Clients), "被离线个数: ", len(leveOffList), " udpClient 剩余:", len(Clients_UDP))

	delete(Clients, *construct.Playerinfo.Userid)

	//if connStructUdp,ok := Clients_UDP[*construct.Playerinfo.Userid];ok { 不能调用此 udp close 会让下次连接无响应
	//	//closeudp := Clients_UDP[*construct.Playerinfo.Userid].Conn_udp
	//	//closeudp.Close()
	//	connStructUdp.Conn_udp.Close()
	//}else {
	//
	//}
	delete(Clients_UDP, *construct.Playerinfo.Userid)
	fmt.Errorf("################## 删除 udp ##################")
	if len(Clients) < 1 {
		//Frame_Cache = make([]ConnStruct_UDP,0)
		//Frame_Cache = NewQueue(300)
		//Frame_Cache = list.New()旧方式
		fmt.Println("### 清理列表")
	}

	fmt.Println("B###### 有离线的,tcpClient剩余: ", len(Clients), "被离线个数: ", len(leveOffList), " udpClient 剩余:", len(Clients_UDP))
	//告知 剩余在线的玩家 ，我掉线了
	for _, d := range Clients {

		dataPB := &msg.Rspn_LeaveOffOthers{
			Player: leveOffList,
		}

		//encode protobuf ========================
		pb_encodede, err := proto.Marshal(dataPB)
		if err != nil {
			log.Fatal("marshaling error:", err)
		}
		sendBuff := WriteClient(TCP_Enum_OthersLineOff, 1, pb_encodede)
		fmt.Println("下发离线")
		i, err := d.Conn.Write(sendBuff) //把其他玩家的数据列表发给当前连接者
		if err != nil {
			log.Fatal("conn.Write Error", err, i)
			d.Conn.Close()
			closeudp := Clients_UDP[*d.Playerinfo.Userid].Conn_tcp
			closeudp.Close()
			//Clients = Remove(Clients, d)
			delete(Clients, *d.Playerinfo.Userid)
			delete(Clients_UDP, *d.Playerinfo.Userid)
			fmt.Errorf("################## 删除 udp ##################")
			return
		}
	}

}

/*
//##### UDP ######
//因为udp是无持续连接特性， 所以专门设计了一个tcp用于心跳包，并在此处理离线逻辑
func LeaveOffothers_UDP(construct ConnStruct_UDP) {
	//if len(Clients) < 1 {
	//	fmt.Println("===============================================================没有玩家: ======!!!!!")
	//	return
	//}
	var leveOffList = make([] *msg.PlayerInfo,0)
	for _,d := range Clients_UDP{
		if *d.Playerinfo.Userid == *construct.Playerinfo.Userid {// && *d.Playerinfo.Userid != *construct.Playerinfo.Userid{
			leveOffList = append(leveOffList,d.Playerinfo) //离线角色
			fmt.Println("离线角色 ,accid: ",d.Accountdata.AccountID)
		}
	}
	delete(Clients_UDP, construct.Accountdata.AccountID)
	fmt.Println("###### 有离线的,剩余client: ",len(Clients_UDP) , "被离线个数: ", len(leveOffList))
	//告知 剩余在线的玩家 ，我掉线了
	for _,d := range Clients_UDP{

		dataPB := &msg.Rspn_LeaveOffOthers{
			Player: leveOffList,
		}

		//encode protobuf ========================
		pb_encodede, err := proto.Marshal(dataPB)
		if err != nil {
			log.Fatal("marshaling error:", err)
		}
		sendBuff := WriteClient(LeaveOffOthers, 1, pb_encodede)
		fmt.Println("下发离线")
		i, err := d.Conn_udp.Write(sendBuff) //把其他玩家的数据列表发给当前连接者
		if err != nil {
			log.Fatal("conn.Write Error", err,i)
			d.Conn_udp.Close()

			//Clients = Remove(Clients, d)
			delete(Clients_UDP,d.Accountdata.AccountID)
			return
		}
	}


}*/
