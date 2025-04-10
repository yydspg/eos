
package msg

import (
	"errors"
	"fmt"
	"time"
)

func (x *GetMaxAndMinSeqReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	return nil
}

func (x *SendMsgReq) Check() error {
    // 检查 MsgData 是否为空
    if x.MsgData == nil {
        // 如果 MsgData 为空，返回一个新的错误，错误信息为 "MsgData is empty"
        return errors.New("MsgData is empty")
    }
    // 调用 MsgData 的 Check 方法进行检查
    if err := x.MsgData.Check(); err != nil {
        // 如果 MsgData 的 Check 方法返回错误，则直接返回该错误
        return err
    }
    // 如果所有检查都通过，返回 nil 表示没有错误
    return nil
}

func (x *SetSendMsgStatusReq) Check() error {
	if x.Status < 0 || x.Status > 3 {
		return errors.New("status is invalid")
	}
	return nil
}

func (x *GetSendMsgStatusReq) Check() error {
	return nil
}



func (x *DelMsgsReq) Check() error {
	return nil
}

func (x *RevokeMsgReq) Check() error {
	if x.ConversationID == "" {
		return errors.New("conversationID is empty")
	}
	if x.Seq < 1 {
		return errors.New("seq is invalid")
	}
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	return nil
}

func (x *MarkMsgsAsReadReq) Check() error {
	if x.ConversationID == "" {
		return errors.New("conversationID is empty")
	}
	if x.Seqs == nil {
		return errors.New("seqs is empty")
	}
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	for _, seq := range x.Seqs {
		if seq == 0 {
			return errors.New("seqs has 0 value is invalid")
		}
	}
	return nil
}

func (x *MarkConversationAsReadReq) Check() error {
	if x.ConversationID == "" {
		return errors.New("conversationID is empty")
	}
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	if x.HasReadSeq < 1 {
		return errors.New("hasReadSeq is invalid")
	}
	for _, seq := range x.Seqs {
		if seq == 0 {
			return errors.New("seqs has 0 value is invalid")
		}
	}
	return nil
}

func (x *SetConversationHasReadSeqReq) Check() error {
	if x.ConversationID == "" {
		return errors.New("conversationID is empty")
	}
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	if x.HasReadSeq < 1 {
		return errors.New("hasReadSeq is invalid")
	}
	return nil
}

func (x *ClearConversationsMsgReq) Check() error {
	if x.ConversationIDs == nil {
		return errors.New("conversationIDs is empty")
	}
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	return nil
}

func (x *UserClearAllMsgReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	return nil
}

func (x *DeleteMsgsReq) Check() error {
	if x.ConversationID == "" {
		return errors.New("conversationID is empty")
	}
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	if x.Seqs == nil {
		return errors.New("seqs is empty")
	}
	return nil
}

func (x *DeleteMsgPhysicalReq) Check() error {
	if x.ConversationIDs == nil {
		return errors.New("conversationIDs is empty")
	}
	return nil
}

func (x *GetConversationMaxSeqReq) Check() error {
	if x.ConversationID == "" {
		return errors.New("conversationID is empty")
	}
	return nil
}

func (x *GetConversationsHasReadAndMaxSeqReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	return nil
}

func (x *GetConversationMaxSeqResp) Format() any {
	if x.MaxSeq > 50 {
		return fmt.Sprintf("len is %v", x.MaxSeq)
	}
	return x
}

func (x *GetConversationsHasReadAndMaxSeqResp) Format() any {
	if len(x.Seqs) > 50 {
		return fmt.Sprintf("len is %v", len(x.Seqs))
	}
	return x
}

func (x *SeqsInfoResp) Format() any {
	if len(x.MaxSeqs) > 50 {
		return fmt.Sprintf("len is %v", len(x.MaxSeqs))
	}
	return x
}

func (x *DestructMsgsReq) Check() error {
	if x.Timestamp > time.Now().UnixMilli() {
		return errors.New("request millisecond timestamp error")
	}
	if x.Limit <= 0 {
		return errors.New("request limit error")
	}
	return nil
}
