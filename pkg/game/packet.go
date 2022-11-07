package game

import (
	"fmt"
	"reflect"
	"runtime/debug"

	"github.com/rs/zerolog/log"
	"github.com/teyvat-helper/hk4e-proto/pb"
)

type Packet struct {
	session *PlayerSession
	head    *pb.PacketHead
	message pb.ProtoMessage
	command pb.ProtoCommand
	rawData []byte
}

func (s *Server) onPacket(packet *Packet) {
	defer func() {
		if r := recover(); r != nil {
			log.Trace().Err(fmt.Errorf("panic: %s", r)).Str("stack", string(debug.Stack())).Str("type", reflect.TypeOf(packet.message).Elem().String()).Msg("Failed to handle packet")
		}
	}()

	var err error
	ctx := s.Context(packet)
	if packet.message != nil {
		name := packet.message.ProtoMessageType()
		if block := s.filter.Map(ctx, name); block {
			log.Warn().Uint16("cmdid", uint16(packet.command)).Str("name", string(name)).Msg("Blocked recv packet")
			return
		}
	}

	switch message := packet.message.(type) {

	// handle cmd avatar
	case *pb.SetUpAvatarTeamReq:
		err = s.HandleSetUpAvatarTeamReq(ctx, message)
	case *pb.ChooseCurAvatarTeamReq:
		err = s.HandleChooseCurAvatarTeamReq(ctx, message)
	case *pb.ChangeAvatarReq:
		err = s.HandleChangeAvatarReq(ctx, message)
	case *pb.AvatarPromoteReq:
		err = s.HandleAvatarPromoteReq(ctx, message)
	case *pb.SpringUseReq:
		err = s.HandleSpringUseReq(ctx, message)
	case *pb.RefreshBackgroundAvatarReq:
		err = s.HandleRefreshBackgroundAvatarReq(ctx, message)
	case *pb.AvatarUpgradeReq:
		err = s.HandleAvatarUpgradeReq(ctx, message)
	case *pb.AvatarDieAnimationEndReq:
		err = s.HandleAvatarDieAnimationEndReq(ctx, message)
	case *pb.AvatarChangeElementTypeReq:
		err = s.HandleAvatarChangeElementTypeReq(ctx, message)
	case *pb.AvatarExpeditionAllDataReq:
		err = s.HandleAvatarExpeditionAllDataReq(ctx, message)
	case *pb.AvatarExpeditionStartReq:
		err = s.HandleAvatarExpeditionStartReq(ctx, message)
	case *pb.AvatarExpeditionCallBackReq:
		err = s.HandleAvatarExpeditionCallBackReq(ctx, message)
	case *pb.AvatarExpeditionGetRewardReq:
		err = s.HandleAvatarExpeditionGetRewardReq(ctx, message)
	case *pb.ChangeMpTeamAvatarReq:
		err = s.HandleChangeMpTeamAvatarReq(ctx, message)
	case *pb.ChangeTeamNameReq:
		err = s.HandleChangeTeamNameReq(ctx, message)
	case *pb.FocusAvatarReq:
		err = s.HandleFocusAvatarReq(ctx, message)
	case *pb.AvatarWearFlycloakReq:
		err = s.HandleAvatarWearFlycloakReq(ctx, message)
	case *pb.AvatarFetterLevelRewardReq:
		err = s.HandleAvatarFetterLevelRewardReq(ctx, message)
	case *pb.AvatarPromoteGetRewardReq:
		err = s.HandleAvatarPromoteGetRewardReq(ctx, message)
	case *pb.AvatarChangeCostumeReq:
		err = s.HandleAvatarChangeCostumeReq(ctx, message)
	case *pb.AvatarChangeAnimHashReq:
		err = s.HandleAvatarChangeAnimHashReq(ctx, message)
	case *pb.PersistentDungeonSwitchAvatarReq:
		err = s.HandlePersistentDungeonSwitchAvatarReq(ctx, message)
	case *pb.AddBackupAvatarTeamReq:
		err = s.HandleAddBackupAvatarTeamReq(ctx, message)
	case *pb.DelBackupAvatarTeamReq:
		err = s.HandleDelBackupAvatarTeamReq(ctx, message)
	case *pb.ItemRenameAvatarReq:
		err = s.HandleItemRenameAvatarReq(ctx, message)

	// handle item misc
	case *pb.UseItemReq:
		err = s.HandleUseItemReq(ctx, message)
	case *pb.DropItemReq:
		err = s.HandleDropItemReq(ctx, message)
	case *pb.WearEquipReq:
		err = s.HandleWearEquipReq(ctx, message)
	case *pb.TakeoffEquipReq:
		err = s.HandleTakeoffEquipReq(ctx, message)
	case *pb.WeaponUpgradeReq:
		err = s.HandleWeaponUpgradeReq(ctx, message)
	case *pb.WeaponPromoteReq:
		err = s.HandleWeaponPromoteReq(ctx, message)
	case *pb.ReliquaryUpgradeReq:
		err = s.HandleReliquaryUpgradeReq(ctx, message)
	case *pb.ReliquaryPromoteReq:
		err = s.HandleReliquaryPromoteReq(ctx, message)
	case *pb.AvatarCardChangeReq:
		err = s.HandleAvatarCardChangeReq(ctx, message)
	case *pb.WeaponAwakenReq:
		err = s.HandleWeaponAwakenReq(ctx, message)
	case *pb.CombineReq:
		err = s.HandleCombineReq(ctx, message)
	case *pb.ForgeGetQueueDataReq:
		err = s.HandleForgeGetQueueDataReq(ctx, message)
	case *pb.ForgeStartReq:
		err = s.HandleForgeStartReq(ctx, message)
	case *pb.ForgeQueueManipulateReq:
		err = s.HandleForgeQueueManipulateReq(ctx, message)
	case *pb.BuyResinReq:
		err = s.HandleBuyResinReq(ctx, message)
	case *pb.TakeMaterialDeleteReturnReq:
		err = s.HandleTakeMaterialDeleteReturnReq(ctx, message)
	case *pb.McoinExchangeHcoinReq:
		err = s.HandleMcoinExchangeHcoinReq(ctx, message)
	case *pb.DestroyMaterialReq:
		err = s.HandleDestroyMaterialReq(ctx, message)
	case *pb.SetEquipLockStateReq:
		err = s.HandleSetEquipLockStateReq(ctx, message)
	case *pb.CalcWeaponUpgradeReturnItemsReq:
		err = s.HandleCalcWeaponUpgradeReturnItemsReq(ctx, message)
	case *pb.SetIsAutoUnlockSpecificEquipReq:
		err = s.HandleSetIsAutoUnlockSpecificEquipReq(ctx, message)
	case *pb.ReliquaryDecomposeReq:
		err = s.HandleReliquaryDecomposeReq(ctx, message)

	// handle cmd misc
	case *pb.GmTalkReq:
		err = s.HandleGmTalkReq(ctx, message)
	case *pb.PingReq:
		err = s.HandlePingReq(ctx, message)
	case *pb.GetOnlinePlayerListReq:
		err = s.HandleGetOnlinePlayerListReq(ctx, message)
	case *pb.GetOnlinePlayerInfoReq:
		err = s.HandleGetOnlinePlayerInfoReq(ctx, message)
	case *pb.CheckSegmentCRCReq:
		err = s.HandleCheckSegmentCRCReq(ctx, message)

	// handle cmd player
	case *pb.GetPlayerTokenReq:
		err = s.HandleGetPlayerTokenReq(ctx, message)
	case *pb.PlayerLoginReq:
		err = s.HandlePlayerLoginReq(ctx, message)
	case *pb.PlayerLogoutReq:
		err = s.HandlePlayerLogoutReq(ctx, message)
	case *pb.ChangeGameTimeReq:
		err = s.HandleChangeGameTimeReq(ctx, message)
	case *pb.SetPlayerPropReq:
		err = s.HandleSetPlayerPropReq(ctx, message)
	case *pb.SetPlayerBornDataReq:
		err = s.HandleSetPlayerBornDataReq(ctx, message)
	case *pb.SetPlayerNameReq:
		err = s.HandleSetPlayerNameReq(ctx, message)
	case *pb.SetOpenStateReq:
		err = s.HandleSetOpenStateReq(ctx, message)
	case *pb.PlayerCookReq:
		err = s.HandlePlayerCookReq(ctx, message)
	case *pb.PlayerRandomCookReq:
		err = s.HandlePlayerRandomCookReq(ctx, message)
	case *pb.PlayerCompoundMaterialReq:
		err = s.HandlePlayerCompoundMaterialReq(ctx, message)
	case *pb.TakeCompoundOutputReq:
		err = s.HandleTakeCompoundOutputReq(ctx, message)
	case *pb.GetCompoundDataReq:
		err = s.HandleGetCompoundDataReq(ctx, message)
	case *pb.PlayerSetPauseReq:
		err = s.HandlePlayerSetPauseReq(ctx, message)
	case *pb.PlayerSetLanguageReq:
		err = s.HandlePlayerSetLanguageReq(ctx, message)
	case *pb.TakePlayerLevelRewardReq:
		err = s.HandleTakePlayerLevelRewardReq(ctx, message)
	case *pb.ItemGivingReq:
		err = s.HandleItemGivingReq(ctx, message)
	case *pb.PlayerCookArgsReq:
		err = s.HandlePlayerCookArgsReq(ctx, message)
	case *pb.PlayerForceExitReq:
		err = s.HandlePlayerForceExitReq(ctx, message)
	case *pb.GetNextResourceInfoReq:
		err = s.HandleGetNextResourceInfoReq(ctx, message)
	case *pb.AdjustWorldLevelReq:
		err = s.HandleAdjustWorldLevelReq(ctx, message)
	case *pb.DailyTaskFilterCityReq:
		err = s.HandleDailyTaskFilterCityReq(ctx, message)
	case *pb.GetGameplayRecommendationReq:
		err = s.HandleGetGameplayRecommendationReq(ctx, message)
	case *pb.TakeBackGivingItemReq:
		err = s.HandleTakeBackGivingItemReq(ctx, message)
	case *pb.ReadNicknameAuditReq:
		err = s.HandleReadNicknameAuditReq(ctx, message)
	case *pb.PlayerCompoundMaterialBoostReq:
		err = s.HandlePlayerCompoundMaterialBoostReq(ctx, message)

	// handle cmd scene
	case *pb.LeaveSceneReq:
		err = s.HandleLeaveSceneReq(ctx, message)
	case *pb.SceneInitFinishReq:
		err = s.HandleSceneInitFinishReq(ctx, message)
	case *pb.SceneEntityMoveReq:
		err = s.HandleSceneEntityMoveReq(ctx, message)
	case *pb.SceneAvatarStaminaStepReq:
		err = s.HandleSceneAvatarStaminaStepReq(ctx, message)
	case *pb.GetScenePointReq:
		err = s.HandleGetScenePointReq(ctx, message)
	case *pb.SceneTransToPointReq:
		err = s.HandleSceneTransToPointReq(ctx, message)
	case *pb.GetSceneAreaReq:
		err = s.HandleGetSceneAreaReq(ctx, message)
	case *pb.SceneEntityDrownReq:
		err = s.HandleSceneEntityDrownReq(ctx, message)
	case *pb.SceneCreateEntityReq:
		err = s.HandleSceneCreateEntityReq(ctx, message)
	case *pb.SceneDestroyEntityReq:
		err = s.HandleSceneDestroyEntityReq(ctx, message)
	case *pb.EnterWorldAreaReq:
		err = s.HandleEnterWorldAreaReq(ctx, message)
	case *pb.EntityForceSyncReq:
		err = s.HandleEntityForceSyncReq(ctx, message)
	case *pb.GetAreaExplorePointReq:
		err = s.HandleGetAreaExplorePointReq(ctx, message)
	case *pb.ClientTransmitReq:
		err = s.HandleClientTransmitReq(ctx, message)
	case *pb.ExecuteGroupTriggerReq:
		err = s.HandleExecuteGroupTriggerReq(ctx, message)
	case *pb.LevelupCityReq:
		err = s.HandleLevelupCityReq(ctx, message)
	case *pb.PersonalSceneJumpReq:
		err = s.HandlePersonalSceneJumpReq(ctx, message)
	case *pb.JoinPlayerSceneReq:
		err = s.HandleJoinPlayerSceneReq(ctx, message)
	case *pb.SceneKickPlayerReq:
		err = s.HandleSceneKickPlayerReq(ctx, message)
	case *pb.BackMyWorldReq:
		err = s.HandleBackMyWorldReq(ctx, message)
	case *pb.SeeMonsterReq:
		err = s.HandleSeeMonsterReq(ctx, message)
	case *pb.EnterSceneReadyReq:
		err = s.HandleEnterSceneReadyReq(ctx, message)
	case *pb.EnterSceneDoneReq:
		err = s.HandleEnterSceneDoneReq(ctx, message)
	case *pb.WorldPlayerReviveReq:
		err = s.HandleWorldPlayerReviveReq(ctx, message)
	case *pb.SetSceneWeatherAreaReq:
		err = s.HandleSetSceneWeatherAreaReq(ctx, message)
	case *pb.ExecuteGadgetLuaReq:
		err = s.HandleExecuteGadgetLuaReq(ctx, message)
	case *pb.SceneEntitiesMovesReq:
		err = s.HandleSceneEntitiesMovesReq(ctx, message)
	case *pb.UnlockTransPointReq:
		err = s.HandleUnlockTransPointReq(ctx, message)
	case *pb.SceneWeatherForcastReq:
		err = s.HandleSceneWeatherForcastReq(ctx, message)
	case *pb.MarkMapReq:
		err = s.HandleMarkMapReq(ctx, message)
	case *pb.PostEnterSceneReq:
		err = s.HandlePostEnterSceneReq(ctx, message)
	case *pb.PlayerChatReq:
		err = s.HandlePlayerChatReq(ctx, message)
	case *pb.GetDungeonEntryExploreConditionReq:
		err = s.HandleGetDungeonEntryExploreConditionReq(ctx, message)
	case *pb.GetMapMarkTipsReq:
		err = s.HandleGetMapMarkTipsReq(ctx, message)
	case *pb.ChangeWorldToSingleModeReq:
		err = s.HandleChangeWorldToSingleModeReq(ctx, message)
	case *pb.GetWorldMpInfoReq:
		err = s.HandleGetWorldMpInfoReq(ctx, message)
	case *pb.GetScenePerformanceReq:
		err = s.HandleGetScenePerformanceReq(ctx, message)
	case *pb.GetMapAreaReq:
		err = s.HandleGetMapAreaReq(ctx, message)
	case *pb.CheckGroupReplacedReq:
		err = s.HandleCheckGroupReplacedReq(ctx, message)
	case *pb.WidgetQuickHitTreeReq:
		err = s.HandleWidgetQuickHitTreeReq(ctx, message)

	// handle cmd social
	case *pb.GetPlayerFriendListReq:
		err = s.HandleGetPlayerFriendListReq(ctx, message)
	case *pb.AskAddFriendReq:
		err = s.HandleAskAddFriendReq(ctx, message)
	case *pb.DealAddFriendReq:
		err = s.HandleDealAddFriendReq(ctx, message)
	case *pb.GetPlayerSocialDetailReq:
		err = s.HandleGetPlayerSocialDetailReq(ctx, message)
	case *pb.DeleteFriendReq:
		err = s.HandleDeleteFriendReq(ctx, message)
	case *pb.SetPlayerBirthdayReq:
		err = s.HandleSetPlayerBirthdayReq(ctx, message)
	case *pb.SetPlayerSignatureReq:
		err = s.HandleSetPlayerSignatureReq(ctx, message)
	case *pb.SetPlayerHeadImageReq:
		err = s.HandleSetPlayerHeadImageReq(ctx, message)
	case *pb.SetNameCardReq:
		err = s.HandleSetNameCardReq(ctx, message)
	case *pb.GetAllUnlockNameCardReq:
		err = s.HandleGetAllUnlockNameCardReq(ctx, message)
	case *pb.AddBlacklistReq:
		err = s.HandleAddBlacklistReq(ctx, message)
	case *pb.RemoveBlacklistReq:
		err = s.HandleRemoveBlacklistReq(ctx, message)
	case *pb.GetRecentMpPlayerListReq:
		err = s.HandleGetRecentMpPlayerListReq(ctx, message)
	case *pb.TakeFirstShareRewardReq:
		err = s.HandleTakeFirstShareRewardReq(ctx, message)
	case *pb.UpdatePS4BlockListReq:
		err = s.HandleUpdatePS4BlockListReq(ctx, message)
	case *pb.GetPlayerBlacklistReq:
		err = s.HandleGetPlayerBlacklistReq(ctx, message)
	case *pb.PlayerReportReq:
		err = s.HandlePlayerReportReq(ctx, message)
	case *pb.SetFriendRemarkNameReq:
		err = s.HandleSetFriendRemarkNameReq(ctx, message)
	case *pb.UpdatePlayerShowAvatarListReq:
		err = s.HandleUpdatePlayerShowAvatarListReq(ctx, message)
	case *pb.GetFriendShowAvatarInfoReq:
		err = s.HandleGetFriendShowAvatarInfoReq(ctx, message)
	case *pb.UpdatePlayerShowNameCardListReq:
		err = s.HandleUpdatePlayerShowNameCardListReq(ctx, message)
	case *pb.GetFriendShowNameCardInfoReq:
		err = s.HandleGetFriendShowNameCardInfoReq(ctx, message)
	case *pb.ForceAddPlayerFriendReq:
		err = s.HandleForceAddPlayerFriendReq(ctx, message)
	case *pb.GetPlayerAskFriendListReq:
		err = s.HandleGetPlayerAskFriendListReq(ctx, message)
	case *pb.GetChatEmojiCollectionReq:
		err = s.HandleGetChatEmojiCollectionReq(ctx, message)
	case *pb.SetChatEmojiCollectionReq:
		err = s.HandleSetChatEmojiCollectionReq(ctx, message)
	case *pb.UpdatePS4FriendListReq:
		err = s.HandleUpdatePS4FriendListReq(ctx, message)
	case *pb.ReadSignatureAuditReq:
		err = s.HandleReadSignatureAuditReq(ctx, message)

	}
	if err != nil {
		log.Printf("[GAME] Failed to handle %s: %v", packet.message.ProtoMessageType(), err)
	}
}
