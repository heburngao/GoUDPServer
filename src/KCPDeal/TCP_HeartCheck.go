package KCPDeal

import (
	"net"
	"time"
	//"fmt"
)

//=================== check status ====================
//检测是否有通信字节流
func TCP_GravelChannel(buff []byte, msg chan byte) {
	for _, v := range buff {
		msg <- v
	}
	close(msg)
}

//延时处理与掉线判别
func TCP_HeartCheck(conn net.Conn, msg chan byte) {
	select {
	case <-time.After(time.Second * 20): //超时会断开
		//xxxfmt.Println("time is out 27 seconds!")
		conn.Close()

		//========================
		for _, con := range Clients {
			if con.Conn == conn {
				//utils.Clients = utils.Remove(utils.Clients, con)
				//delete(dataMgr.Clients, con.Accountdata.AccountID)
				//dataMgr.Log("断开一个链接, 剩余连接数:", len(dataMgr.Clients))
				//dataMgr.IndexOfClient--

				go TCP_OthersLineOff(con)

			}
		}
	case <-msg:
		//xxxfmt.Println(">>>>>> 后延超时 27秒 确认，connecting !", <-msg)
		conn.SetDeadline(time.Now().Add(18 * time.Second)) //设置时限//此处需略大于客户端的设定

		break
	}
}
