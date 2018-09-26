package main

import (
	util "KCPUtil"
	dataMgr "KCPDeal"
	"fmt"
	"net"
	"os"
	"time"
	"sync"
	ikcp "KCP"
	"log"
)
var PR = fmt.Println




func main() {



	//socket, err := net.ListenUDP("udp4", &net.UDPAddr{
	//	IP:   net.IPv4(0, 0, 0, 0),
	//	Port: 8080,
	//})

	port := ":2020"
	addr := dataMgr.GetLocalIp()
	//addr := "0.0.0.0"
	//###### tcp #########
	//netListen, err := net.Listen("tcp",addr+port) xxxxx
	//dataMgr.CheckError(err)xxxxx
	//defer closeLisener(netListen) xxxxx
	PR("tcp init")

	//###### udp #########
	udp_addr, err := net.ResolveUDPAddr("udp", addr+port)
	fmt.Println("address : ", udp_addr)
	dataMgr.CheckError(err)
	fmt.Println("recv:1")
/*
	addr := net.UDPAddr{
		Port: 2000,
		IP: net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
*/
	connUdp, err := net.ListenUDP("udp", udp_addr)
	//conn,err := net.Listen("udp",addr+":2020")
	fmt.Println("[ UDP ] recv:2")
	defer connUdp.Close()
	fmt.Println("[ UDP ] recv:3")
	dataMgr.CheckError(err)

	fmt.Println("[ UDP ] recv:4 " )


	recvUDPmsg(connUdp)
	 //recvTCPmsg(netListen)xxx
	fmt.Println("[ UDP ] recv:5")
}
////###### tcp #########
//func recvTCPmsg(listener net.Listener)  {
//	if len(dataMgr.Clients) >= 200 {
//
//		return
//	}
//	PT("[ TCP ] waitting for clients @ TCP:: ")
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			continue
//		}
//		PT(conn.RemoteAddr().String(), "[ TCP ] @@@@@@@@@@ tcp connect success", "剩余连接数:", len(dataMgr.Clients))
//		go TCP_handler(conn)
//	}
//}
//
//func TCP_handler(conn net.Conn)  {
//	//处理粘包
//	for {
//
//		//size, err := conn.Read(buffer) //此方法是阻塞的，一直等待客户端的消息
//		body := &util.CombineBody{0,make([]byte, util.BUF_MAX), nil, make([]byte, 0), 0}
//		sizeRead, err := conn.Read(body.Buffer) //此方法是阻塞的，一直等待客户端的消息
//		//==========处理错误==============
//		if err != nil {
//			//if err == io.EOF {
//
//			conn.Close()
//			PT(conn.RemoteAddr().String(), "[ TCP ] [ goUDP:: ]########## 连接断开 connection error: ", err)
//
//			for _, con := range dataMgr.Clients {
//				if con.Conn == conn {
//					//utils.Clients = utils.Remove(utils.Clients, con)
//					//delete(dataMgr.Clients, con.Accountdata.AccountID)
//					//dataMgr.Log("断开一个链接, 剩余连接数:", len(dataMgr.Clients))
//					//dataMgr.IndexOfClient--
//
//					go dataMgr.TCP_OthersLineOff(con)
//
//				}
//			}
//
//			//}else{
//			//	Log("########## normal error ! ", err)
//			//}
//
//			return
//		}
//		//检测状态 ===================================
//		msg := make(chan byte)
//		go dataMgr.TCP_GravelChannel(body.Buffer[:sizeRead], msg)
//		go dataMgr.TCP_HeartCheck(conn, msg)
//		//==========处理错误==============
//		go Unpack(sizeRead,body,conn,err)
//
//	} //end for
//}
//func Unpack(sizeRead int , body *util.CombineBody,conn net.Conn, err error)  {
//	//lockUpack.Lock()
//	//==========处理粘包==============
//	if sizeRead > 0 {
//		size, real := body.GetRealReadBuf(sizeRead) //收集字节序
//		fmt.Println("[ TCP ] size:", size, " real:", real , " len(real):" , len(real))
//		//real := util.CombinedBuff
//		var combinedLen int = len(real)
//		//var index_getCombineSize int = 0
//		body.Index_getCombineSize = 0
//		for {
//			if combinedLen < 1 {
//				break
//			}
//
//			if combinedLen < PACK_HEADER {
//				body.TooShortReceive(combinedLen)
//				//接收头部不完整
//				break
//			} else {
//				fmt.Println("[ TCP ] ### ### index_getCombineSize:",body.Index_getCombineSize)
//				msgSettingSize := body.GetMsgSize2(int32(body.Index_getCombineSize))
//				fmt.Println("[ TCP ] GetMsgSize2 ::## " , msgSettingSize)
//				if combinedLen < int(msgSettingSize) {
//					body.TooShortReceive(combinedLen)
//					//接收消息体不完整
//					break
//				} else {
//					usefulBuff := body.GetUsefulBuff(int(msgSettingSize)) //最终合并
//					fmt.Println("[ TCP ] 正式处理############## 当前包长:" , combinedLen , " usefulBuff: " , usefulBuff , " ||| size: " , size)
//					//body := &util.ReadBody{make([]byte,1024),0}
//					//body.SetBuff_read(readlReadBuffItem)
//					//msgOpcode := body.ReadShort()
//					//msgsize := util.GetMsgSize()
//					go dealwith_TCP(err, conn, usefulBuff, int(msgSettingSize))
//					combinedLen -= int(msgSettingSize)
//					if combinedLen > 0 {
//						fmt.Println("[ TCP ] 有粘包############## 剩余包长:" , combinedLen)
//						//body.Index_getCombineSize = int(msgSettingSize)
//						body.Index_getCombineSize += int(msgSettingSize)
//					}
//				}
//			}
//		}
//	} else {
//		conn.Close()
//		return
//	}
//	//lockUpack.Unlock()
//}
//func dealwith_TCP(err error,conn net.Conn,buffer []byte , size int)  {
//
//	//xxx打印
//	//xxxfor i := 0; i < size; i++ {
//	//xxx	fmt.Println(i, buffer[i])
//	//xxx}
//
//	// fmt.Println(buffer);
//	// fmt.Println(Bytes2Int(buffer[0:2]))
//	////检测状态 ===================================
//	//msg := make(chan byte)
//	//go dataMgr.GravelChannel(buffer[:size], msg)
//	//go dataMgr.HeartBeatingChecking(conn, msg)
//	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>> 读取client数据 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
//
//	cmd, pb := dataMgr.ReadClient(buffer)
//
//	//cmd, pb := dataMgr.ReadClient(tmpBuffer)zzzzz
//
//	pb = dataMgr.GetCloneArry(pb)
//	// ===============  正式按条件解读 客户端业务逻辑  ====================
//	fmt.Println("[ TCP ]>>>>>>>  处理  >>>>>>", cmd)
//	go dataMgr.TCP_Receive(cmd, pb, conn)
//	//========================================================================
//
//}
//func closeLisener(listener net.Listener) {
//	dataMgr.Log("[ TCP ] 关闭监听 listener")
//	listener.Close()
//
//}
const PACK_HEADER = 6




//###### udp ########
var llock sync.Mutex
func recvUDPmsg(connUDP *net.UDPConn) {
	defer connUDP.Close();
	if len(dataMgr.Clients_UDP) >= 200 {

		return
	}
	//go UDP_TimerWaitting() //xxxxx instead by  kcp as below  暂去
	PR("[ UDP ] waitting for clients @ [KCP]:: ")


	llock.Lock()
	//==== kcp ====

	ikcp.SndToClient = func(id int,buf []byte, _len int) {
		 fmt.Println("SndToClient:: 处理 kcp 发送 ##### ", buf)
		for _, d := range dataMgr.Clients_UDP {

			//go dataMgr.UDP_TimerStatusCast(d) xxxxx
			addr := d.Addr
			i, err := d.Conn_udp.WriteToUDP(buf, addr) //把其他玩家的数据列表发给当前连接者 本体方法
			if err != nil {
				log.Fatal("[ UDP ] d.Write Error", err, i)
				//d.Conn_udp.Close()
				//delete(Clients_UDP, *d.StatusInfo.Userid) //d.Accountdata.AccountID)
				//fmt.Errorf("[ UDP ] ################## 删除 udp ##################")
				//
				//closeudp := Clients[*d.StatusInfo.Userid].Conn
				//closeudp.Close()
				//delete(Clients, *d.StatusInfo.Userid)
				//return
				break
			}
		}
	}


	  	go ikcp.Use(2)
	llock.Unlock()
		 rcv(connUDP)
	fmt.Println("rcv new data :6:")
	//go ikcp.Use(2)
	fmt.Println("rcv new data :7:")
	//==== kcp ====

	//for {
	//
	//	//read udp
	//	//		buf := make([]byte,4096)
	//	//		readLen, remoteAddr, err := conn.ReadFromUDP(buf[0:])
	//	body := &util.CombineBody{0, make([]byte, util.BUF_MAX), nil, make([]byte, 0), 0}
	//	PR("[ UDP ] recv:1, 等待client...")
	//	readLen, clientAddr, err := connUDP.ReadFromUDP(body.Buffer) //此方法是阻塞的，一直等待客户端的消息 connUDP收到 -> body.Buffer中
	//	PR("[ UDP ] recv:################## 收到新一轮  ##################")
	//
	//	PR("[ UDP ] recv:2 udp address: " , connUDP.RemoteAddr(),"地址: ", clientAddr)                                  //收到客户端包数据，才走到这里，然后往下走
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "[ UDP ] ######读取失败#####，Fatal error: %s", err.Error())
	//		return
	//		 //continue
	//	}
	//
	//	//fmt.Println(" udp msg is : ", string(buf[0:readLen]), " 接到长度: " , readLen, "地址: " , remoteAddr)
	//	//fmt.Println(" udp msg is : ", string(body.Buffer[0:readLen]), " 接到长度: ", readLen, "地址: ", remoteAddr)
	//	fmt.Println("[ UDP ]  udp msg is : ", body.Buffer[0:readLen], " 接到长度: ", readLen, "地址: ", clientAddr)
	//	go UDP_handler(err, connUDP, body.Buffer, readLen ,clientAddr)
	//
	//
	//	//write to udp #######################
	//	//udpmsg := make([]byte, 1024)
	//	//_, err = conn.WriteToUDP(udpmsg, remoteAddr)
	//	//dataMgr.CheckError(err)
	//
	//}

}
func rcv(connUDP *net.UDPConn){
	for {
		body := &util.CombineBody{0, make([]byte, util.BUF_MAX), nil, make([]byte, 0), 0}
		readLen, clientAddr, err := connUDP.ReadFromUDP(body.Buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ UDP ] ######读取失败#####，Fatal error: %s", err.Error() , clientAddr)
			//return
			continue
		}
		fmt.Println("==================== rcv new data len:1:",readLen)
		//fmt.Println("rcv new data buffer:2:##### ",body.Buffer)
		if ikcp.RcvFromClient != nil {

			//fmt.Println("rcv new data len:3:",readLen)
			ikcp.DealRcvFn = func(buf []byte) {

				fmt.Println("rcv new data len:2:",len(buf))

				fmt.Println("rcv new data buffer:3:##### ",buf[24:])

				cmd, pb := dataMgr.ReadClient(buf[24:])
				fmt.Println("rcv new data len:4:",readLen)
				//=========================================================================================================
				pb = dataMgr.GetCloneArry(pb)
				// ===============  正式按条件解读 客户端业务逻辑  ====================
				fmt.Println("[ UDP ]>>>>>>>  处理  >>>>>>", cmd)
				go dataMgr.UDP_Receive(cmd, pb[0:], *connUDP, clientAddr)
				fmt.Println("rcv new data len:5:",readLen)

			}

			ikcp.RcvFromClient(body.Buffer,readLen, ikcp.DealRcvFn)//readLen, body.Buffer)
			//ikcp.RcvFromClient(buff,readLen, ikcp.DealRcvFn)
		}else{
			fmt.Println("ikcp.RcvFromClient == nil ")
		}
	}
}
var lockTimer sync.Mutex
func UDP_TimerWaitting(){
	timerTick := time.Tick(time.Millisecond * 66) //每隔66毫秒执行一次把本轮收到的所有包列表下发
	for t := range timerTick {
		//lockTimer.Lock()
		//fmt.Println("TTTTTTTTTTTTT 每隔66毫秒执行一次 , t :", t)
		for _, d := range dataMgr.Clients_UDP {
		//	if d.Playerinfo == nil { // && *d.Playerinfo.Userid == *pbObj.Player.Userid {
		//		continue
		//	}
		//
		//	//if *d.Playerinfo.Userid == *pbObj.Player.Userid { //对指针的值比较
		//	//	fmt.Println("[ UDP ] timer ####################==> 下发 userid:" , *d.Playerinfo.Userid , " addr:",d.Addr , " Clients_UDP个数:",len(dataMgr.Clients_UDP))
		//
		//		//同时下发广播出去
		//		go dataMgr.StatusCast_UDP(d , d.Addr)
		//		break
		//	//}
			go dataMgr.UDP_TimerStatusCast(d)
		}

		//dataMgr.UDP_TimerStatusCast()//旧方式
		t.Clock()
		//lockTimer.Unlock()
	}
}
func UDP_handler(err error, connUDP *net.UDPConn, buffer []byte, size int , clientAddr *net.UDPAddr) {

	//xxx打印
	// for i := 0; i < size; i++ {
	// 	fmt.Println("[ UDP ] :",i, buffer[i])
	// }

	//检测状态 ===================================
	//msg := make(chan byte) tcp
	//go dataMgr.GravelChannel(buffer[:size], msg) tcp
	//go dataMgr.HeartBeatingChecking_UDP(*conn, msg) for tcp
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>> 读取client数据 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
//=========================================================================================================
	cmd, pb := dataMgr.ReadClient(buffer)
//=========================================================================================================
	pb = dataMgr.GetCloneArry(pb)
	// ===============  正式按条件解读 客户端业务逻辑  ====================
	fmt.Println("[ UDP ]>>>>>>>  处理  >>>>>>", cmd)
	go dataMgr.UDP_Receive(cmd, pb[0:], *connUDP , clientAddr)

}
