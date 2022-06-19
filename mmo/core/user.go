package core

import (
	"fmt"
	"math/rand"
	"sync"
	//"time"

	"github.com/golang/protobuf/proto"
	"github.com/huahearts/kyubia/kiface"
	"github.com/huahearts/kyubia/mmo/pb"
)

type User struct {
	Pid  uint32
	Conn kiface.IConnection
	X    float32
	Y    float32
	Z    float32
	V    float32
}

var PidGen uint32 = 1
var IdLock sync.Mutex

func NewUser(conn ziface.IConnection) *User {
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()
	return &User{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(134 + rand.Intn(17)),
		V:    0,
	}
}

func (u *User) SendMsg(msgId uint32, data proto.Message) {
	fmt.Printf("before Marshal data=%+v\n", data)
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal msg err :", err)
		return
	}

	fmt.Printf("after marshal data=%v\n", msg)
	if u.Conn == nil {
		fmt.Println("connection in player is nil")
		return
	}

	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		fmt.Println("usr sendmsg err!")
		return
	}
	return
}

func (u *User) SyncPid() {
	data := &pb.SyncPid{
		Pid: u.Pid,
	}

	u.SendMsg(1, data)
}

func (p *User) BroadCastStartPosition() {

	msg := &pb.BroadCast{
		PID: p.Pid,
		Tp:  2, //TP2 代表广播坐标
		Data: &pb.BroadCast_P{
			&pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	p.SendMsg(200, msg)
}
