package KCPDeal

import (
	"container/list"
	"net"
	"protocol/msg"
)

//var IndexOfCreate int32
//var IndexOfClient int
//=======================
var Clients map[string]ConnStruct = make(map[string]ConnStruct, 0)
var Clients_UDP map[string]ConnStruct_UDP = make(map[string]ConnStruct_UDP, 0)

//var Frame_Cache map[int32] msg.PlayerInfo = make(map[int32] msg.PlayerInfo,0)
//var Frame_Cache []ConnStruct_UDP = make([]ConnStruct_UDP,0)
//var Frame_Cache *EsQueue = NewQueue(300)

//var Frame_Cache *list.List = list.New()xxxxx
//var Frame_Cache_Remove *list.List = list.New()xxxxx

//var AllPlayer map[int32] msg.PlayerInfo = make(map[int32] msg.PlayerInfo,0)
type ConnStruct struct {
	Conn        net.Conn
	Accountdata Accountdata
	Playerinfo  *msg.PlayerInfo
}
type ConnStruct_UDP struct {
	Frame_Cache        *list.List // = list.New()
	Frame_Cache_Remove *list.List
	Conn_udp           net.UDPConn
	//Conn_tcp    net.Conn
	//Accountdata Accountdata
	StatusInfo *msg.StatusInfo
	Addr       *net.UDPAddr
}
type Accountdata struct {
	AccountID *int32
	Password  *int32
	//Level    int32
	// sessionid string
}

//============ enum ===========
//type Enum int
const (
	//_ Enum = iota
	//Enum            	= iota
	TCP_Enum_HeartBeat            = 110
	TCP_Enum_CreateSelf           = 10000
	TCP_Enum_CreateOtherPlayer    = 10001
	UDP_Enum_UpdateStatus         = 20002
	UDP_Enum_UpdateStatus_Confirm = 20003
	TCP_Enum_OthersLineOff        = 10003
	TCP_Enum_AddRoboot            = 99999
)
