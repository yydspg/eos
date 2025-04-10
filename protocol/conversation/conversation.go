package conversation

import (

	"errors"
	"fmt"
	"github.com/eos/protocol/constant"
)

func (q *ConversationReq) Check() error {
	if q.ConversationID == "" {
		return errors.New("conversationID is empty")
	}
	return nil
}
func (q *Conversation) Check() error {
	if q.OwnerUserID == "" {
		return errors.New("OwnUserID is empty")
	}
	if q.ConversationID == "" {
		return errors.New("ConversationID is empty")
	}
	if q.ConversationType < 1 || q.ConversationType > 4 {
		return errors.New("ConversationType is invalid")
	}
	if q.RecvMsgOpt < 0 || q.RecvMsgOpt > 2 {
		return errors.New("RecvMsgOpt is invalid")
	}
	return nil
}
func (q *SetConversationReq) Check() error {
	if q.Conversation == nil {
		return errors.New("Conversation is empty")
	}
	if q.Conversation.ConversationID == "" {
		return errors.New("conversation is empty")
	}
	return nil
}  

func (q *GetConversationsReq) Check() error {
	if q.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if q.ConversationIDs == nil {
		return errors.New("conversationIDs is empty")
	}
	return nil
}

func (q *GetAllConversationsReq) Check() error {
	if q.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	return nil
}

func (q *GetRecvMsgNotNotifyUserIDsReq) Check() error {
	if q.GroupID == "" {
		return errors.New("groupID is empty")
	} 
	return nil
}

func (q *CreateGroupChatConversationsReq) Check() error {
	if q.GroupID == "" {
		return errors.New("groupID is empty")
	}
	return nil
}

func (q *SetConversationMaxSeqReq) Check() error {
	if q.ConversationID == "" {
		return errors.New("conversationID is empty")
	}
	if q.OwnerUserID == nil {
		return errors.New("ownerUserID is empty")
	}
	if q.MaxSeq == 0 {
		return errors.New("maxReq is invalid")
	}
	return nil
}

func (q *SetConversationsReq) Check() error {
	if q.UserIDs == nil {
		return errors.New("userID is empty")
	}
	if q.Conversation == nil {
		return errors.New("conversation is empty")
	}
	if q.Conversation.ConversationType == 0 {
		return errors.New("conversationType is invalid")
	}
	if q.Conversation.ConversationType == constant.SingleChatType && q.Conversation.UserID == "" { 
		return errors.New("userID is empty")
	}
	if q.Conversation.ConversationType == constant.ReadGroupChatType && q.Conversation.GroupID == "" {
		return errors.New("groupID is empty")
	}

	return nil
}
func (q *GetUserConversationIDsHashReq) Check() error {
	if q.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	return nil
}
func (q *GetConversationsByConversationIDReq) Check() error {
	if q.ConversationIDs == nil {
		return errors.New("conversationIDs is empty")
	}
	return nil
}
func (q *GetSortedConversationListReq) Check() error {
	if q.UserID == "" {
		return errors.New("userID is empty")
	}
	if q.Pagination == nil {
		return errors.New("pagination is empty")
	}
	if q.Pagination.PageNumber < 1 {
		return errors.New("pageNumber is invalid")
	}
	return nil
}
func (q *GetConversationIDsResp) Format() any {
	if len(q.ConversationIDs) > 50 {
		return fmt.Sprintf("len is %v", len(q.ConversationIDs))
	}
	return q
}

func (q *GetOwnerConversationResp) Format() any {
	if q.Total > 50 {
		return fmt.Sprintf("len is %v", q.Total)
	}
	return q
}

func (q *GetAllConversationsResp) Format() any {
	if len(q.Conversations) > 50 {
		return fmt.Sprintf("len is %v", len(q.Conversations))
	}
	return q
}

func (q *GetFullOwnerConversationIDsResp) Format() any {
	if len(q.ConversationIDs) > 20 {
		return fmt.Sprintf("len is %v", len(q.ConversationIDs))
	}
	return q
}

func (q *ClearUserConversationMsgReq) Check() error {
	if q.Limit <= 0 {
		return errors.New("limit is invalid")
	}
	return nil
}