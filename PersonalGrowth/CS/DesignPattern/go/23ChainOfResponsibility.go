package designpattern

import "fmt"

// Handler 定义处理者接口
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string) string
}

// BaseHandler 提供基础实现
type BaseHandler struct {
	nextHandler Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.nextHandler = handler
	return handler
}

func (h *BaseHandler) Handle(request string) string {
	if h.nextHandler != nil {
		return h.nextHandler.Handle(request)
	}
	return ""
}

// ConcreteHandlerA 具体处理者A
type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) Handle(request string) string {
	if request == "A" {
		return "Handler A handled " + request
	}
	return h.BaseHandler.Handle(request)
}

// ConcreteHandlerB 具体处理者B
type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) Handle(request string) string {
	if request == "B" {
		return "Handler B handled " + request
	}
	return h.BaseHandler.Handle(request)
}

// ConcreteHandlerC 具体处理者C
type ConcreteHandlerC struct {
	BaseHandler
}

func (h *ConcreteHandlerC) Handle(request string) string {
	if request == "C" {
		return "Handler C handled " + request
	}
	return h.BaseHandler.Handle(request)
}

// 实际应用示例：请假审批系统
type LeaveRequest struct {
	name   string
	days   int
	reason string
}

func NewLeaveRequest(name string, days int, reason string) *LeaveRequest {
	return &LeaveRequest{
		name:   name,
		days:   days,
		reason: reason,
	}
}

type LeaveHandler interface {
	SetNext(handler LeaveHandler) LeaveHandler
	HandleRequest(request *LeaveRequest) string
}

type BaseLeaveHandler struct {
	nextHandler LeaveHandler
}

func (h *BaseLeaveHandler) SetNext(handler LeaveHandler) LeaveHandler {
	h.nextHandler = handler
	return handler
}

// TeamLeader 团队领导（可以批准3天以内的假期）
type TeamLeader struct {
	BaseLeaveHandler
}

func (h *TeamLeader) HandleRequest(request *LeaveRequest) string {
	if request.days <= 3 {
		return fmt.Sprintf("TeamLeader approved %s's leave request for %d days",
			request.name, request.days)
	}
	if h.nextHandler != nil {
		return h.nextHandler.HandleRequest(request)
	}
	return "Leave request denied"
}

// Manager 经理（可以批准7天以内的假期）
type Manager struct {
	BaseLeaveHandler
}

func (h *Manager) HandleRequest(request *LeaveRequest) string {
	if request.days <= 7 {
		return fmt.Sprintf("Manager approved %s's leave request for %d days",
			request.name, request.days)
	}
	if h.nextHandler != nil {
		return h.nextHandler.HandleRequest(request)
	}
	return "Leave request denied"
}

// Director 总监（可以批准14天以内的假期）
type DirectorLeaveHandler struct {
	BaseLeaveHandler
}

func (h *DirectorLeaveHandler) HandleRequest(request *LeaveRequest) string {
	if request.days <= 14 {
		return fmt.Sprintf("Director approved %s's leave request for %d days",
			request.name, request.days)
	}
	if h.nextHandler != nil {
		return h.nextHandler.HandleRequest(request)
	}
	return "Leave request denied"
}
