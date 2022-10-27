package game

import "github.com/teyvat-helper/hk4e-proto/pb"

func (s *Server) HandleGetPlayerFriendListReq(ctx *Context, req *pb.GetPlayerFriendListReq) error {
	panic("not implement")
}

func (s *Server) SendGetPlayerFriendListRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAskAddFriendReq(ctx *Context, req *pb.AskAddFriendReq) error {
	panic("not implement")
}

func (s *Server) SendAskAddFriendRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleDealAddFriendReq(ctx *Context, req *pb.DealAddFriendReq) error {
	panic("not implement")
}

func (s *Server) SendDealAddFriendRsp(ctx *Context) error {
	panic("not implement")
}

// handle GetPlayerSocialDetailReq
//
//	flow:
//		*RECV <·· GetPlayerSocialDetailReq
//		*SEND ··> GetPlayerSocialDetailRsp
func (s *Server) HandleGetPlayerSocialDetailReq(ctx *Context, req *pb.GetPlayerSocialDetailReq) error {
	return s.SendGetPlayerSocialDetailRsp(ctx, req.GetUid())
}
func (s *Server) SendGetPlayerSocialDetailRsp(ctx *Context, id uint32) error {
	player := ctx.Session().GetPlayer()
	basic := player.Basic()
	social := player.Social()
	var resp pb.GetPlayerSocialDetailRsp
	resp.DetailData = &pb.SocialDetail{
		Uid:                   uint32(player.ID),
		Level:                 basic.GetLevel(),
		WorldLevel:            1,
		Nickname:              basic.GetNickname(),
		Signature:             social.GetSignature(),
		Birthday:              social.GetBirthday(),
		NameCardId:            social.GetNameCardId(),
		IsShowAvatar:          social.GetIsShowAvatar(),
		ShowNameCardIdList:    social.GetShowNameCardIdList(),
		FriendEnterHomeOption: pb.FriendEnterHomeOption_FRIEND_ENTER_HOME_OPTION_DIRECT,
		ProfilePicture: &pb.ProfilePicture{
			AvatarId:  basic.GetHeadImageAvatarId(),
			CostumeId: basic.GetProfilePictureCostumeId(),
		},
	}
	resp.DetailData.OnlineState = pb.FriendOnlineState_FRIEND_ONLINE
	resp.DetailData.IsFriend = true
	resp.DetailData.IsMpModeAvailable = true
	resp.DetailData.FinishAchievementNum = 0
	return s.Send(ctx, &resp)
}

func (s *Server) HandleDeleteFriendReq(ctx *Context, req *pb.DeleteFriendReq) error {
	panic("not implement")
}

func (s *Server) SendDeleteFriendRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetPlayerBirthdayReq(ctx *Context, req *pb.SetPlayerBirthdayReq) error {
	panic("not implement")
}

func (s *Server) SendSetPlayerBirthdayRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetPlayerSignatureReq(ctx *Context, req *pb.SetPlayerSignatureReq) error {
	panic("not implement")
}

func (s *Server) SendSetPlayerSignatureRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetPlayerHeadImageReq(ctx *Context, req *pb.SetPlayerHeadImageReq) error {
	panic("not implement")
}

func (s *Server) SendSetPlayerHeadImageRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendUpdatePS4FriendListNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendDeleteFriendNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAddFriendNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAskAddFriendNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetNameCardReq(ctx *Context, req *pb.SetNameCardReq) error {
	panic("not implement")
}

func (s *Server) SendSetNameCardRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetAllUnlockNameCardReq(ctx *Context, req *pb.GetAllUnlockNameCardReq) error {
	panic("not implement")
}

func (s *Server) SendGetAllUnlockNameCardRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAddBlacklistReq(ctx *Context, req *pb.AddBlacklistReq) error {
	panic("not implement")
}

func (s *Server) SendAddBlacklistRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleRemoveBlacklistReq(ctx *Context, req *pb.RemoveBlacklistReq) error {
	panic("not implement")
}

func (s *Server) SendRemoveBlacklistRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendUnlockNameCardNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetRecentMpPlayerListReq(ctx *Context, req *pb.GetRecentMpPlayerListReq) error {
	panic("not implement")
}

func (s *Server) SendGetRecentMpPlayerListRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSocialDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleTakeFirstShareRewardReq(ctx *Context, req *pb.TakeFirstShareRewardReq) error {
	panic("not implement")
}

func (s *Server) SendTakeFirstShareRewardRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleUpdatePS4BlockListReq(ctx *Context, req *pb.UpdatePS4BlockListReq) error {
	panic("not implement")
}

func (s *Server) SendUpdatePS4BlockListRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetPlayerBlacklistReq(ctx *Context, req *pb.GetPlayerBlacklistReq) error {
	panic("not implement")
}

func (s *Server) SendGetPlayerBlacklistRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandlePlayerReportReq(ctx *Context, req *pb.PlayerReportReq) error {
	panic("not implement")
}

func (s *Server) SendPlayerReportRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetFriendRemarkNameReq(ctx *Context, req *pb.SetFriendRemarkNameReq) error {
	panic("not implement")
}

func (s *Server) SendSetFriendRemarkNameRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleUpdatePlayerShowAvatarListReq(ctx *Context, req *pb.UpdatePlayerShowAvatarListReq) error {
	panic("not implement")
}

func (s *Server) SendUpdatePlayerShowAvatarListRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetFriendShowAvatarInfoReq(ctx *Context, req *pb.GetFriendShowAvatarInfoReq) error {
	panic("not implement")
}

func (s *Server) SendGetFriendShowAvatarInfoRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleUpdatePlayerShowNameCardListReq(ctx *Context, req *pb.UpdatePlayerShowNameCardListReq) error {
	panic("not implement")
}

func (s *Server) SendUpdatePlayerShowNameCardListRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetFriendShowNameCardInfoReq(ctx *Context, req *pb.GetFriendShowNameCardInfoReq) error {
	panic("not implement")
}

func (s *Server) SendGetFriendShowNameCardInfoRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleForceAddPlayerFriendReq(ctx *Context, req *pb.ForceAddPlayerFriendReq) error {
	panic("not implement")
}

func (s *Server) SendForceAddPlayerFriendRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendProfilePictureChangeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPSNFriendListNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPSNBlackListNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetPlayerAskFriendListReq(ctx *Context, req *pb.GetPlayerAskFriendListReq) error {
	panic("not implement")
}

func (s *Server) SendGetPlayerAskFriendListRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetChatEmojiCollectionReq(ctx *Context, req *pb.GetChatEmojiCollectionReq) error {
	panic("not implement")
}

func (s *Server) SendGetChatEmojiCollectionRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetChatEmojiCollectionReq(ctx *Context, req *pb.SetChatEmojiCollectionReq) error {
	panic("not implement")
}

func (s *Server) SendSetChatEmojiCollectionRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleUpdatePS4FriendListReq(ctx *Context, req *pb.UpdatePS4FriendListReq) error {
	panic("not implement")
}

func (s *Server) SendUpdatePS4FriendListRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendFriendInfoChangeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerSignatureAuditDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerSignatureNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSignatureAuditConfigNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleReadSignatureAuditReq(ctx *Context, req *pb.ReadSignatureAuditReq) error {
	panic("not implement")
}

func (s *Server) SendReadSignatureAuditRsp(ctx *Context) error {
	panic("not implement")
}
