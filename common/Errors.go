package common

import "errors"

var (
	// 连接丢失
	ERR_CONNECTION_LOSS = errors.New("ERR_CONNECTION_LOSS")
    // 发送消息满
	ERR_SEND_MESSAGE_FULL = errors.New("ERR_SEND_MESSAGE_FULL")
    // 加入房间2次
	ERR_JOIN_ROOM_TWICE = errors.New("ERR_JOIN_ROOM_TWICE")
    // 未在该房间
	ERR_NOT_IN_ROOM = errors.New("ERR_NOT_IN_ROOM")
    // 房间id无效
	ERR_ROOM_ID_INVALID = errors.New("ERR_ROOM_ID_INVALID")
    // 分发渠道已满
	ERR_DISPATCH_CHANNEL_FULL = errors.New("ERR_DISPATCH_CHANNEL_FULL")
	//  管理器渠道已满
	ERR_MERGE_CHANNEL_FULL = errors.New("ERR_MERGE_CHANNEL_FULL")
	// 验证无效
	ERR_CERT_INVALID = errors.New("ERR_CERT_INVALID")
	// 逻辑分发渠道已满
	ERR_LOGIC_DISPATCH_CHANNEL_FULL = errors.New("ERR_LOGIC_DISPATCH_CHANNEL_FULL")
)
