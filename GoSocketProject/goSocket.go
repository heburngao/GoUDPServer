package main

import (
	//cgolib "BuffUtil" //引入src目录下的 BuffUtil目录
	dataMgr "ConnectionDeal"
	//"ConnectionDeal"
	 util "BuffUtil"
	//"bytes"
	//"encoding/binary"
	//"fmt"
	//"github.com/golang/protobuf/proto"
	//"log"
	"net"
	//"os"
	//"protocol"
	//rqst "protocol/Rqs"
	//resp "protocol/Rsp"
	//"reflect"
	"time"
	//"io"
	//"strconv"
	"fmt"
	//"runtime"
	//"sync"
)

//===================================================================================================================
func main() {
	//utils.Test()
	//================================================ 建立socket，监听端口 ==================================================================

	 //netListen, err := net.Listen("tcp", "172.20.4.41:2020")
	//netListen, err := net.Listen("tcp", "127.0.0.1:2020")
	addr := dataMgr.GetLocalIp()
	netListen, err := net.Listen("tcp", addr+":2020")
	net.DialTimeout("tcp", netListen.Addr().String(), time.Second*30)

	dataMgr.CheckError(err)
	//defer netListen.Close()
	defer closeLisener(netListen)
	dataMgr.Log("Waiting for clients >>>>>>>>>>>>>>>>>")
	for {

		conn, err := netListen.Accept()
		//conn.SetDeadline(time.Now().Add(15 * time.Second))//设置时限
		if err != nil {

			continue
		}
		dataMgr.Log("@@@@@@ @@@@@@ receive new connect")

		dataMgr.Log(conn.RemoteAddr().String(), "@@@@@@@@@@ tcp connect success", "剩余连接数:", len(dataMgr.Clients))
		go handleConnection(conn)

	}

}
func closeLisener(listener net.Listener) {
	dataMgr.Log("关闭监听 listener")
	listener.Close()

}
func closeConnect(conn net.Conn) {
	dataMgr.Log("关闭监听 connection")
	conn.Close()
}

//var plyaers = make([]PlayerData,0)
const PACK_HEADER = 6
//处理连接
func handleConnection(conn net.Conn) {
	//lockConn.Lock()
	if len(dataMgr.Clients) >= 200 {
		conn.Close() //限定200连接
		return
	}
	//声明一个临时缓冲区，用来存储被截断的数据
	//tmpBuffer := make([]byte, 0)//zzzzz

	//声明一个管道用于接收解包的数据
	//readerChannel := make(chan []byte, 16)zzzzz
	//go reader(readerChannel)zzzzz

	//buffer := make([]byte, 1024)


	//处理粘包
	for {

		//size, err := conn.Read(buffer) //此方法是阻塞的，一直等待客户端的消息
		body := &util.CombineBody{0,make([]byte, util.BUF_MAX), nil, make([]byte, 0), 0}
		sizeRead, err := conn.Read(body.Buffer) //此方法是阻塞的，一直等待客户端的消息
		//==========处理错误==============
		if err != nil {
			//if err == io.EOF {

			conn.Close()
			dataMgr.Log(conn.RemoteAddr().String(), "########## 连接断开connection error: ", err)
			for _, con := range dataMgr.Clients {
				if con.Conn == conn {
					//utils.Clients = utils.Remove(utils.Clients, con)
					//delete(dataMgr.Clients, con.Accountdata.AccountID)
					dataMgr.Log("断开一个链接, 剩余连接数:", len(dataMgr.Clients))
					//dataMgr.IndexOfClient--

					go dataMgr.LeaveOffothers(con)

				}
			}

			//}else{
			//	Log("########## normal error ! ", err)
			//}

			return
		}
		//==========处理错误==============
		go Unpack(sizeRead,body,conn,err)

	} //end for

	// var ticker *time.Ticker = time.NewTicker(time.Second)
	//lockConn.Unlock()
}
//var lockUpack sync.Mutex
//var lockConn sync.Mutex
func Unpack(sizeRead int , body *util.CombineBody,conn net.Conn, err error)  {
	//lockUpack.Lock()
	//==========处理粘包==============
	if sizeRead > 0 {
		size, real := body.GetRealReadBuf(sizeRead) //收集字节序
		fmt.Println("size:", size, " real:", real , " len(real):" , len(real))
		//real := util.CombinedBuff
		var combinedLen int = len(real)
		//var index_getCombineSize int = 0
		body.Index_getCombineSize = 0
		for {
			if combinedLen < 1 {
				break
			}

			if combinedLen < PACK_HEADER {
				body.TooShortReceive(combinedLen)
				//接收头部不完整
				break
			} else {
				fmt.Println("### ### index_getCombineSize:",body.Index_getCombineSize)
				msgSettingSize := body.GetMsgSize2(int32(body.Index_getCombineSize))
				fmt.Println("GetMsgSize2 ::## " , msgSettingSize)
				if combinedLen < int(msgSettingSize) {
					body.TooShortReceive(combinedLen)
					//接收消息体不完整
					break
				} else {
					usefulBuff := body.GetUsefulBuff(int(msgSettingSize)) //最终合并
					fmt.Println("正式处理############## 当前包长:" , combinedLen , " usefulBuff: " , usefulBuff , " ||| size: " , size)
					//body := &util.ReadBody{make([]byte,1024),0}
					//body.SetBuff_read(readlReadBuffItem)
					//msgOpcode := body.ReadShort()
					//msgsize := util.GetMsgSize()
					go dealwith(err, conn, usefulBuff, int(msgSettingSize))
					combinedLen -= int(msgSettingSize)
					if combinedLen > 0 {
						fmt.Println("有粘包############## 剩余包长:" , combinedLen)
						//body.Index_getCombineSize = int(msgSettingSize)
						body.Index_getCombineSize += int(msgSettingSize)
					}
				}
			}
		}
	} else {
		conn.Close()
		return
	}
	//lockUpack.Unlock()
}
func dealwith(err error,conn net.Conn,buffer []byte , size int)  {

	//xxx打印
	//xxxfor i := 0; i < size; i++ {
	//xxx	fmt.Println(i, buffer[i])
	//xxx}

	// fmt.Println(buffer);
	// fmt.Println(Bytes2Int(buffer[0:2]))
	//检测状态 ===================================
	msg := make(chan byte)
	go dataMgr.GravelChannel(buffer[:size], msg)
	go dataMgr.HeartBeatingChecking(conn, msg)
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>> 读取client数据 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>


	//tmpBuffer = util.Unpack(append(tmpBuffer, buffer[:size]...), readerChannel)zzzzz

	cmd, pb := dataMgr.ReadClient(buffer)

	//cmd, pb := dataMgr.ReadClient(tmpBuffer)zzzzz

	pb = dataMgr.GetCloneArry(pb)
	// ===============  正式按条件解读 客户端业务逻辑  ====================
	fmt.Println(">>>>>>>  处理  >>>>>>", cmd)
	go dataMgr.DealWithCMD(cmd, pb, conn)
	//========================================================================
	//单次计时
	//timer := time.NewTimer(time.Second * 1) //一秒后触发
	//<-timer.C                               //此处在等待channel中的信号，执行此段代码时会阻塞一秒
	//timer.Stop()
	//=====
	//for range time.Tick(time.Second*1){
	//	fmt.Println("每隔一秒执行一次 , t :" , t)
	//}
	// 或者 =====
	//timerTick := time.Tick(time.Second * 1) //每隔1秒执行一次
	//for t := range timerTick {
	//	fmt.Println("每隔一秒执行一次 , t :", t)
	//	//t.Clock()
	//}

	// 或者 ====
	//for {
	//	time.Sleep(time.Second * 1)
	//	fmt.Println("每隔一秒执行一次")
	//}
	//========================================================================
	// fmt.Println("读取payload:",cgolib.ReadBytes())

	// fmt.Println(string(buffer[0:2]))
	// fmt.Println(get_Int32(buffer,2))
	// fmt.Println(string(buffer[6:n]))
	// Log(conn.RemoteAddr().String(), "<<<<<<<<收到receive data string:\n", string(buffer[:n]))

}
//func reader(readerChannel chan []byte) {
//	for {
//		select {
//		case data := <-readerChannel:
//			Log(string(data))
//		}
//	}
//}
//
//func Log(v ...interface{}) {
//	fmt.Println(v...)
//}
//===================================================
// [] net.Conn  conns = make([] net.Conn)
