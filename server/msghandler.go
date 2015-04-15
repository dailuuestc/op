package main

import "github.com/golang/protobuf/proto"

type msg struct {
	t       uint32
	session uint32
	data    []byte
}

type msgCB func(*Agent, proto.Message)
type msgHandler struct {
	p  proto.Message
	cb msgCB
}

var handlers = make(map[uint32]msgHandler)

func registerHandler(t uint32, p proto.Message, cb msgCB) {
	handlers[t] = msgHandler{p, cb}
}

func dispatchOutsideMsg(agent *Agent, m *msg) {
	log(DEBUG, "dispatchOutsideMsg\n")
	if m.session != agent.session+1 {
		log(ERROR, "session not equal, cli[%d], svr[%d]\n", m.session, agent.session+1)
		return
	}
	agent.session++

	h, ok := handlers[m.t]
	if ok != true {
		log(ERROR, "msg[%d] handler not found\n", m.t)
		return
	}

	if err := proto.Unmarshal(m.data, h.p); err != nil {
		log(ERROR, "msg[%d] Unmarshal failed: %s\n", m.t, err)
		return
	}

	h.cb(agent, h.p)
}

func replyMsg(agent *Agent, t uint32, p proto.Message) {
	data, err := proto.Marshal(p)
	if err != nil {
		log(ERROR, "proto[%d] marshal failed: %s", err)
		return
	}
	m := &msg{t, agent.session, data}
	send(agent, m)
}

func dispatchInnerMsg(agent *Agent, msg interface{}) {
}
