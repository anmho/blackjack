// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: proto/blackjack/messages.proto

package pb_blackjack

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlackjackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*BlackjackRequest_JoinGameRequest
	//	*BlackjackRequest_StartGameRequest
	//	*BlackjackRequest_BlackjackMove
	//	*BlackjackRequest_ViewGamesRequest
	//	*BlackjackRequest_LeaveGameRequest
	Request isBlackjackRequest_Request `protobuf_oneof:"request"`
}

func (x *BlackjackRequest) Reset() {
	*x = BlackjackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blackjack_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlackjackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlackjackRequest) ProtoMessage() {}

func (x *BlackjackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blackjack_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlackjackRequest.ProtoReflect.Descriptor instead.
func (*BlackjackRequest) Descriptor() ([]byte, []int) {
	return file_proto_blackjack_messages_proto_rawDescGZIP(), []int{0}
}

func (m *BlackjackRequest) GetRequest() isBlackjackRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *BlackjackRequest) GetJoinGameRequest() *JoinGameRequest {
	if x, ok := x.GetRequest().(*BlackjackRequest_JoinGameRequest); ok {
		return x.JoinGameRequest
	}
	return nil
}

func (x *BlackjackRequest) GetStartGameRequest() *StartGameRequest {
	if x, ok := x.GetRequest().(*BlackjackRequest_StartGameRequest); ok {
		return x.StartGameRequest
	}
	return nil
}

func (x *BlackjackRequest) GetBlackjackMove() *PlayGameRequest {
	if x, ok := x.GetRequest().(*BlackjackRequest_BlackjackMove); ok {
		return x.BlackjackMove
	}
	return nil
}

func (x *BlackjackRequest) GetViewGamesRequest() *ViewGamesRequest {
	if x, ok := x.GetRequest().(*BlackjackRequest_ViewGamesRequest); ok {
		return x.ViewGamesRequest
	}
	return nil
}

func (x *BlackjackRequest) GetLeaveGameRequest() *LeaveGameRequest {
	if x, ok := x.GetRequest().(*BlackjackRequest_LeaveGameRequest); ok {
		return x.LeaveGameRequest
	}
	return nil
}

type isBlackjackRequest_Request interface {
	isBlackjackRequest_Request()
}

type BlackjackRequest_JoinGameRequest struct {
	JoinGameRequest *JoinGameRequest `protobuf:"bytes,1,opt,name=joinGameRequest,proto3,oneof"`
}

type BlackjackRequest_StartGameRequest struct {
	StartGameRequest *StartGameRequest `protobuf:"bytes,2,opt,name=startGameRequest,proto3,oneof"`
}

type BlackjackRequest_BlackjackMove struct {
	BlackjackMove *PlayGameRequest `protobuf:"bytes,3,opt,name=blackjackMove,proto3,oneof"`
}

type BlackjackRequest_ViewGamesRequest struct {
	ViewGamesRequest *ViewGamesRequest `protobuf:"bytes,4,opt,name=viewGamesRequest,proto3,oneof"`
}

type BlackjackRequest_LeaveGameRequest struct {
	LeaveGameRequest *LeaveGameRequest `protobuf:"bytes,5,opt,name=leaveGameRequest,proto3,oneof"`
}

func (*BlackjackRequest_JoinGameRequest) isBlackjackRequest_Request() {}

func (*BlackjackRequest_StartGameRequest) isBlackjackRequest_Request() {}

func (*BlackjackRequest_BlackjackMove) isBlackjackRequest_Request() {}

func (*BlackjackRequest_ViewGamesRequest) isBlackjackRequest_Request() {}

func (*BlackjackRequest_LeaveGameRequest) isBlackjackRequest_Request() {}

type BlackjackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*BlackjackResponse_JoinGameResult
	//	*BlackjackResponse_StartGameResult
	//	*BlackjackResponse_PlayGameResult
	//	*BlackjackResponse_ViewGamesResult
	//	*BlackjackResponse_LeaveGameResult
	Result isBlackjackResponse_Result `protobuf_oneof:"result"`
}

func (x *BlackjackResponse) Reset() {
	*x = BlackjackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blackjack_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlackjackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlackjackResponse) ProtoMessage() {}

func (x *BlackjackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blackjack_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlackjackResponse.ProtoReflect.Descriptor instead.
func (*BlackjackResponse) Descriptor() ([]byte, []int) {
	return file_proto_blackjack_messages_proto_rawDescGZIP(), []int{1}
}

func (m *BlackjackResponse) GetResult() isBlackjackResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *BlackjackResponse) GetJoinGameResult() *JoinGameResult {
	if x, ok := x.GetResult().(*BlackjackResponse_JoinGameResult); ok {
		return x.JoinGameResult
	}
	return nil
}

func (x *BlackjackResponse) GetStartGameResult() *StartGameResult {
	if x, ok := x.GetResult().(*BlackjackResponse_StartGameResult); ok {
		return x.StartGameResult
	}
	return nil
}

func (x *BlackjackResponse) GetPlayGameResult() *PlayGameResult {
	if x, ok := x.GetResult().(*BlackjackResponse_PlayGameResult); ok {
		return x.PlayGameResult
	}
	return nil
}

func (x *BlackjackResponse) GetViewGamesResult() *ViewGamesResult {
	if x, ok := x.GetResult().(*BlackjackResponse_ViewGamesResult); ok {
		return x.ViewGamesResult
	}
	return nil
}

func (x *BlackjackResponse) GetLeaveGameResult() *LeaveGameResult {
	if x, ok := x.GetResult().(*BlackjackResponse_LeaveGameResult); ok {
		return x.LeaveGameResult
	}
	return nil
}

type isBlackjackResponse_Result interface {
	isBlackjackResponse_Result()
}

type BlackjackResponse_JoinGameResult struct {
	JoinGameResult *JoinGameResult `protobuf:"bytes,1,opt,name=joinGameResult,proto3,oneof"`
}

type BlackjackResponse_StartGameResult struct {
	StartGameResult *StartGameResult `protobuf:"bytes,2,opt,name=startGameResult,proto3,oneof"`
}

type BlackjackResponse_PlayGameResult struct {
	PlayGameResult *PlayGameResult `protobuf:"bytes,3,opt,name=playGameResult,proto3,oneof"`
}

type BlackjackResponse_ViewGamesResult struct {
	ViewGamesResult *ViewGamesResult `protobuf:"bytes,4,opt,name=viewGamesResult,proto3,oneof"`
}

type BlackjackResponse_LeaveGameResult struct {
	LeaveGameResult *LeaveGameResult `protobuf:"bytes,5,opt,name=leaveGameResult,proto3,oneof"`
}

func (*BlackjackResponse_JoinGameResult) isBlackjackResponse_Result() {}

func (*BlackjackResponse_StartGameResult) isBlackjackResponse_Result() {}

func (*BlackjackResponse_PlayGameResult) isBlackjackResponse_Result() {}

func (*BlackjackResponse_ViewGamesResult) isBlackjackResponse_Result() {}

func (*BlackjackResponse_LeaveGameResult) isBlackjackResponse_Result() {}

type ViewGamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ViewGamesRequest) Reset() {
	*x = ViewGamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blackjack_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewGamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewGamesRequest) ProtoMessage() {}

func (x *ViewGamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blackjack_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewGamesRequest.ProtoReflect.Descriptor instead.
func (*ViewGamesRequest) Descriptor() ([]byte, []int) {
	return file_proto_blackjack_messages_proto_rawDescGZIP(), []int{2}
}

type ViewGamesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Games []*Game `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
}

func (x *ViewGamesResult) Reset() {
	*x = ViewGamesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blackjack_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewGamesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewGamesResult) ProtoMessage() {}

func (x *ViewGamesResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blackjack_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewGamesResult.ProtoReflect.Descriptor instead.
func (*ViewGamesResult) Descriptor() ([]byte, []int) {
	return file_proto_blackjack_messages_proto_rawDescGZIP(), []int{3}
}

func (x *ViewGamesResult) GetGames() []*Game {
	if x != nil {
		return x.Games
	}
	return nil
}

type PlayGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Move:
	//
	//	*PlayGameRequest_HitMove
	//	*PlayGameRequest_StandMove
	//	*PlayGameRequest_DoubleMove
	//	*PlayGameRequest_SplitMove
	//	*PlayGameRequest_SurrenderMove
	Move isPlayGameRequest_Move `protobuf_oneof:"move"`
}

func (x *PlayGameRequest) Reset() {
	*x = PlayGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blackjack_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayGameRequest) ProtoMessage() {}

func (x *PlayGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blackjack_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayGameRequest.ProtoReflect.Descriptor instead.
func (*PlayGameRequest) Descriptor() ([]byte, []int) {
	return file_proto_blackjack_messages_proto_rawDescGZIP(), []int{4}
}

func (m *PlayGameRequest) GetMove() isPlayGameRequest_Move {
	if m != nil {
		return m.Move
	}
	return nil
}

func (x *PlayGameRequest) GetHitMove() *HitMove {
	if x, ok := x.GetMove().(*PlayGameRequest_HitMove); ok {
		return x.HitMove
	}
	return nil
}

func (x *PlayGameRequest) GetStandMove() *StandMove {
	if x, ok := x.GetMove().(*PlayGameRequest_StandMove); ok {
		return x.StandMove
	}
	return nil
}

func (x *PlayGameRequest) GetDoubleMove() *DoubleMove {
	if x, ok := x.GetMove().(*PlayGameRequest_DoubleMove); ok {
		return x.DoubleMove
	}
	return nil
}

func (x *PlayGameRequest) GetSplitMove() *SplitMove {
	if x, ok := x.GetMove().(*PlayGameRequest_SplitMove); ok {
		return x.SplitMove
	}
	return nil
}

func (x *PlayGameRequest) GetSurrenderMove() *SurrenderMove {
	if x, ok := x.GetMove().(*PlayGameRequest_SurrenderMove); ok {
		return x.SurrenderMove
	}
	return nil
}

type isPlayGameRequest_Move interface {
	isPlayGameRequest_Move()
}

type PlayGameRequest_HitMove struct {
	HitMove *HitMove `protobuf:"bytes,1,opt,name=hitMove,proto3,oneof"`
}

type PlayGameRequest_StandMove struct {
	StandMove *StandMove `protobuf:"bytes,2,opt,name=standMove,proto3,oneof"`
}

type PlayGameRequest_DoubleMove struct {
	DoubleMove *DoubleMove `protobuf:"bytes,3,opt,name=doubleMove,proto3,oneof"`
}

type PlayGameRequest_SplitMove struct {
	SplitMove *SplitMove `protobuf:"bytes,4,opt,name=splitMove,proto3,oneof"`
}

type PlayGameRequest_SurrenderMove struct {
	SurrenderMove *SurrenderMove `protobuf:"bytes,5,opt,name=surrenderMove,proto3,oneof"`
}

func (*PlayGameRequest_HitMove) isPlayGameRequest_Move() {}

func (*PlayGameRequest_StandMove) isPlayGameRequest_Move() {}

func (*PlayGameRequest_DoubleMove) isPlayGameRequest_Move() {}

func (*PlayGameRequest_SplitMove) isPlayGameRequest_Move() {}

func (*PlayGameRequest_SurrenderMove) isPlayGameRequest_Move() {}

type PlayGameResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayGameResult) Reset() {
	*x = PlayGameResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blackjack_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayGameResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayGameResult) ProtoMessage() {}

func (x *PlayGameResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blackjack_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayGameResult.ProtoReflect.Descriptor instead.
func (*PlayGameResult) Descriptor() ([]byte, []int) {
	return file_proto_blackjack_messages_proto_rawDescGZIP(), []int{5}
}

type JoinGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId string `protobuf:"bytes,1,opt,name=gameId,proto3" json:"gameId,omitempty"`
}

func (x *JoinGameRequest) Reset() {
	*x = JoinGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blackjack_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGameRequest) ProtoMessage() {}

func (x *JoinGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blackjack_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGameRequest.ProtoReflect.Descriptor instead.
func (*JoinGameRequest) Descriptor() ([]byte, []int) {
	return file_proto_blackjack_messages_proto_rawDescGZIP(), []int{6}
}

func (x *JoinGameRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

type JoinGameResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *JoinGameResult) Reset() {
	*x = JoinGameResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blackjack_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGameResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGameResult) ProtoMessage() {}

func (x *JoinGameResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blackjack_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGameResult.ProtoReflect.Descriptor instead.
func (*JoinGameResult) Descriptor() ([]byte, []int) {
	return file_proto_blackjack_messages_proto_rawDescGZIP(), []int{7}
}

func (x *JoinGameResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type StartGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartGameRequest) Reset() {
	*x = StartGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blackjack_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameRequest) ProtoMessage() {}

func (x *StartGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blackjack_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameRequest.ProtoReflect.Descriptor instead.
func (*StartGameRequest) Descriptor() ([]byte, []int) {
	return file_proto_blackjack_messages_proto_rawDescGZIP(), []int{8}
}

type StartGameResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *StartGameResult) Reset() {
	*x = StartGameResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blackjack_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameResult) ProtoMessage() {}

func (x *StartGameResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blackjack_messages_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameResult.ProtoReflect.Descriptor instead.
func (*StartGameResult) Descriptor() ([]byte, []int) {
	return file_proto_blackjack_messages_proto_rawDescGZIP(), []int{9}
}

func (x *StartGameResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type LeaveGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaveGameRequest) Reset() {
	*x = LeaveGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blackjack_messages_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveGameRequest) ProtoMessage() {}

func (x *LeaveGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blackjack_messages_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveGameRequest.ProtoReflect.Descriptor instead.
func (*LeaveGameRequest) Descriptor() ([]byte, []int) {
	return file_proto_blackjack_messages_proto_rawDescGZIP(), []int{10}
}

type LeaveGameResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaveGameResult) Reset() {
	*x = LeaveGameResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blackjack_messages_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveGameResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveGameResult) ProtoMessage() {}

func (x *LeaveGameResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blackjack_messages_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveGameResult.ProtoReflect.Descriptor instead.
func (*LeaveGameResult) Descriptor() ([]byte, []int) {
	return file_proto_blackjack_messages_proto_rawDescGZIP(), []int{11}
}

var File_proto_blackjack_messages_proto protoreflect.FileDescriptor

var file_proto_blackjack_messages_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63,
	0x6b, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x1a, 0x1a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x2f, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a, 0x10, 0x42, 0x6c, 0x61, 0x63,
	0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0f,
	0x6a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63,
	0x6b, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x0f, 0x6a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x42, 0x0a, 0x0d, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x4d, 0x6f, 0x76, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61,
	0x63, 0x6b, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x4d,
	0x6f, 0x76, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x76, 0x69, 0x65, 0x77, 0x47, 0x61, 0x6d, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x47, 0x61,
	0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x76, 0x69,
	0x65, 0x77, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49,
	0x0a, 0x10, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x6a, 0x61, 0x63, 0x6b, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0xff, 0x02, 0x0a, 0x11, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x6a, 0x6f,
	0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x2e, 0x4a,
	0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52,
	0x0e, 0x6a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x46, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x6a, 0x61, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x43, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x70, 0x6c,
	0x61, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x46, 0x0a, 0x0f,
	0x76, 0x69, 0x65, 0x77, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63,
	0x6b, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x48, 0x00, 0x52, 0x0f, 0x76, 0x69, 0x65, 0x77, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x46, 0x0a, 0x0f, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x6c, 0x65, 0x61,
	0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x08, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x56, 0x69, 0x65, 0x77, 0x47, 0x61,
	0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x0f, 0x56, 0x69,
	0x65, 0x77, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x25, 0x0a,
	0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62,
	0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x67,
	0x61, 0x6d, 0x65, 0x73, 0x22, 0xb0, 0x02, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x68, 0x69, 0x74, 0x4d,
	0x6f, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x61, 0x63,
	0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x2e, 0x48, 0x69, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x48, 0x00, 0x52,
	0x07, 0x68, 0x69, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x4d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6c,
	0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x76,
	0x65, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x37,
	0x0a, 0x0a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x2e, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x73, 0x70, 0x6c, 0x69, 0x74,
	0x4d, 0x6f, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6c, 0x61,
	0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x4d, 0x6f, 0x76, 0x65,
	0x48, 0x00, 0x52, 0x09, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x40, 0x0a,
	0x0d, 0x73, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b,
	0x2e, 0x53, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x48, 0x00,
	0x52, 0x0d, 0x73, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x42,
	0x06, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x79, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x4a, 0x6f, 0x69,
	0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x12, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x62,
	0x5f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6a, 0x61, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_blackjack_messages_proto_rawDescOnce sync.Once
	file_proto_blackjack_messages_proto_rawDescData = file_proto_blackjack_messages_proto_rawDesc
)

func file_proto_blackjack_messages_proto_rawDescGZIP() []byte {
	file_proto_blackjack_messages_proto_rawDescOnce.Do(func() {
		file_proto_blackjack_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_blackjack_messages_proto_rawDescData)
	})
	return file_proto_blackjack_messages_proto_rawDescData
}

var file_proto_blackjack_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_blackjack_messages_proto_goTypes = []interface{}{
	(*BlackjackRequest)(nil),  // 0: blackjack.BlackjackRequest
	(*BlackjackResponse)(nil), // 1: blackjack.BlackjackResponse
	(*ViewGamesRequest)(nil),  // 2: blackjack.ViewGamesRequest
	(*ViewGamesResult)(nil),   // 3: blackjack.ViewGamesResult
	(*PlayGameRequest)(nil),   // 4: blackjack.PlayGameRequest
	(*PlayGameResult)(nil),    // 5: blackjack.PlayGameResult
	(*JoinGameRequest)(nil),   // 6: blackjack.JoinGameRequest
	(*JoinGameResult)(nil),    // 7: blackjack.JoinGameResult
	(*StartGameRequest)(nil),  // 8: blackjack.StartGameRequest
	(*StartGameResult)(nil),   // 9: blackjack.StartGameResult
	(*LeaveGameRequest)(nil),  // 10: blackjack.LeaveGameRequest
	(*LeaveGameResult)(nil),   // 11: blackjack.LeaveGameResult
	(*Game)(nil),              // 12: blackjack.Game
	(*HitMove)(nil),           // 13: blackjack.HitMove
	(*StandMove)(nil),         // 14: blackjack.StandMove
	(*DoubleMove)(nil),        // 15: blackjack.DoubleMove
	(*SplitMove)(nil),         // 16: blackjack.SplitMove
	(*SurrenderMove)(nil),     // 17: blackjack.SurrenderMove
}
var file_proto_blackjack_messages_proto_depIdxs = []int32{
	6,  // 0: blackjack.BlackjackRequest.joinGameRequest:type_name -> blackjack.JoinGameRequest
	8,  // 1: blackjack.BlackjackRequest.startGameRequest:type_name -> blackjack.StartGameRequest
	4,  // 2: blackjack.BlackjackRequest.blackjackMove:type_name -> blackjack.PlayGameRequest
	2,  // 3: blackjack.BlackjackRequest.viewGamesRequest:type_name -> blackjack.ViewGamesRequest
	10, // 4: blackjack.BlackjackRequest.leaveGameRequest:type_name -> blackjack.LeaveGameRequest
	7,  // 5: blackjack.BlackjackResponse.joinGameResult:type_name -> blackjack.JoinGameResult
	9,  // 6: blackjack.BlackjackResponse.startGameResult:type_name -> blackjack.StartGameResult
	5,  // 7: blackjack.BlackjackResponse.playGameResult:type_name -> blackjack.PlayGameResult
	3,  // 8: blackjack.BlackjackResponse.viewGamesResult:type_name -> blackjack.ViewGamesResult
	11, // 9: blackjack.BlackjackResponse.leaveGameResult:type_name -> blackjack.LeaveGameResult
	12, // 10: blackjack.ViewGamesResult.games:type_name -> blackjack.Game
	13, // 11: blackjack.PlayGameRequest.hitMove:type_name -> blackjack.HitMove
	14, // 12: blackjack.PlayGameRequest.standMove:type_name -> blackjack.StandMove
	15, // 13: blackjack.PlayGameRequest.doubleMove:type_name -> blackjack.DoubleMove
	16, // 14: blackjack.PlayGameRequest.splitMove:type_name -> blackjack.SplitMove
	17, // 15: blackjack.PlayGameRequest.surrenderMove:type_name -> blackjack.SurrenderMove
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_proto_blackjack_messages_proto_init() }
func file_proto_blackjack_messages_proto_init() {
	if File_proto_blackjack_messages_proto != nil {
		return
	}
	file_proto_blackjack_game_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_blackjack_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlackjackRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blackjack_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlackjackResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blackjack_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewGamesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blackjack_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewGamesResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blackjack_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayGameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blackjack_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayGameResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blackjack_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blackjack_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGameResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blackjack_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blackjack_messages_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGameResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blackjack_messages_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveGameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_blackjack_messages_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveGameResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_blackjack_messages_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BlackjackRequest_JoinGameRequest)(nil),
		(*BlackjackRequest_StartGameRequest)(nil),
		(*BlackjackRequest_BlackjackMove)(nil),
		(*BlackjackRequest_ViewGamesRequest)(nil),
		(*BlackjackRequest_LeaveGameRequest)(nil),
	}
	file_proto_blackjack_messages_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*BlackjackResponse_JoinGameResult)(nil),
		(*BlackjackResponse_StartGameResult)(nil),
		(*BlackjackResponse_PlayGameResult)(nil),
		(*BlackjackResponse_ViewGamesResult)(nil),
		(*BlackjackResponse_LeaveGameResult)(nil),
	}
	file_proto_blackjack_messages_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*PlayGameRequest_HitMove)(nil),
		(*PlayGameRequest_StandMove)(nil),
		(*PlayGameRequest_DoubleMove)(nil),
		(*PlayGameRequest_SplitMove)(nil),
		(*PlayGameRequest_SurrenderMove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_blackjack_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_blackjack_messages_proto_goTypes,
		DependencyIndexes: file_proto_blackjack_messages_proto_depIdxs,
		MessageInfos:      file_proto_blackjack_messages_proto_msgTypes,
	}.Build()
	File_proto_blackjack_messages_proto = out.File
	file_proto_blackjack_messages_proto_rawDesc = nil
	file_proto_blackjack_messages_proto_goTypes = nil
	file_proto_blackjack_messages_proto_depIdxs = nil
}
