package game

import "github.com/teyvat-helper/hk4e-proto/pb"

// send PlayerStoreNotify
//
//	flow:
//		SEND ··> PlayerStoreNotify
func (s *Server) SendPlayerStoreNotify(ctx *Context) error {
	player := ctx.Session().GetPlayer()
	var notify pb.PlayerStoreNotify
	notify.StoreType = pb.StoreType_STORE_PACK
	notify.WeightLimit = 30000
	notify.ItemList = []*pb.Item{}
	for _, v := range player.Item().GetPackStore().GetItemList() {
		var item pb.Item
		item.ItemId = v.GetItemId()
		item.Guid = v.GetGuid()
		switch detail := v.GetDetail().(type) {
		case *pb.ItemBin_Equip:
			var equip pb.Equip
			equip.IsLocked = detail.Equip.GetIsLocked()
			switch detail := detail.Equip.GetDetail().(type) {
			case *pb.EquipBin_Reliquary:
				equip.Detail = &pb.Equip_Reliquary{Reliquary: &pb.Reliquary{
					Level:            detail.Reliquary.GetLevel(),
					Exp:              detail.Reliquary.GetExp(),
					MainPropId:       detail.Reliquary.GetMainPropId(),
					AppendPropIdList: detail.Reliquary.GetAppendPropIdList(),
				}}
			case *pb.EquipBin_Weapon:
				equip.Detail = &pb.Equip_Weapon{Weapon: &pb.Weapon{
					Level:        detail.Weapon.GetLevel(),
					Exp:          detail.Weapon.GetExp(),
					PromoteLevel: detail.Weapon.GetPromoteLevel(),
					AffixMap:     detail.Weapon.GetAffixMap(),
				}}
			}
			item.Detail = &pb.Item_Equip{Equip: &equip}
		case *pb.ItemBin_Furniture:
			item.Detail = &pb.Item_Furniture{Furniture: &pb.Furniture{
				Count: detail.Furniture.GetCount(),
			}}
		case *pb.ItemBin_Material:
			item.Detail = &pb.Item_Material{Material: &pb.Material{
				Count: detail.Material.GetCount(),
			}}
		}
		notify.ItemList = append(notify.ItemList, &item)
	}
	return s.Send(ctx, &notify)
}

// send StoreWeightLimitNotify
//
//	flow:
//		SEND ··> StoreWeightLimitNotify
func (s *Server) SendStoreWeightLimitNotify(ctx *Context) error {
	var notify pb.StoreWeightLimitNotify
	notify.StoreType = pb.StoreType_STORE_PACK
	notify.WeightLimit = 30000
	notify.WeaponCountLimit = 2000
	notify.MaterialCountLimit = 2000
	notify.ReliquaryCountLimit = 1500
	notify.FurnitureCountLimit = 2000
	return s.Send(ctx, &notify)
}

func (s *Server) SendStoreItemChangeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendStoreItemDelNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendItemAddHintNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleUseItemReq(ctx *Context, req *pb.UseItemReq) error {
	panic("not implement")
}

func (s *Server) SendUseItemRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleDropItemReq(ctx *Context, req *pb.DropItemReq) error {
	panic("not implement")
}

func (s *Server) SendDropItemRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleWearEquipReq(ctx *Context, req *pb.WearEquipReq) error {
	panic("not implement")
}

func (s *Server) SendWearEquipRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleTakeoffEquipReq(ctx *Context, req *pb.TakeoffEquipReq) error {
	panic("not implement")
}

func (s *Server) SendTakeoffEquipRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarEquipChangeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleWeaponUpgradeReq(ctx *Context, req *pb.WeaponUpgradeReq) error {
	panic("not implement")
}

func (s *Server) SendWeaponUpgradeRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleWeaponPromoteReq(ctx *Context, req *pb.WeaponPromoteReq) error {
	panic("not implement")
}

func (s *Server) SendWeaponPromoteRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleReliquaryUpgradeReq(ctx *Context, req *pb.ReliquaryUpgradeReq) error {
	panic("not implement")
}

func (s *Server) SendReliquaryUpgradeRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleReliquaryPromoteReq(ctx *Context, req *pb.ReliquaryPromoteReq) error {
	panic("not implement")
}

func (s *Server) SendReliquaryPromoteRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarCardChangeReq(ctx *Context, req *pb.AvatarCardChangeReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarCardChangeRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendGrantRewardNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleWeaponAwakenReq(ctx *Context, req *pb.WeaponAwakenReq) error {
	panic("not implement")
}

func (s *Server) SendWeaponAwakenRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendItemCdGroupTimeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendDropHintNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleCombineReq(ctx *Context, req *pb.CombineReq) error {
	panic("not implement")
}

func (s *Server) SendCombineRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendForgeQueueDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleForgeGetQueueDataReq(ctx *Context, req *pb.ForgeGetQueueDataReq) error {
	panic("not implement")
}

func (s *Server) SendForgeGetQueueDataRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleForgeStartReq(ctx *Context, req *pb.ForgeStartReq) error {
	panic("not implement")
}

func (s *Server) SendForgeStartRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleForgeQueueManipulateReq(ctx *Context, req *pb.ForgeQueueManipulateReq) error {
	panic("not implement")
}

func (s *Server) SendForgeQueueManipulateRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendResinChangeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleBuyResinReq(ctx *Context, req *pb.BuyResinReq) error {
	panic("not implement")
}

func (s *Server) SendBuyResinRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendMaterialDeleteReturnNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleTakeMaterialDeleteReturnReq(ctx *Context, req *pb.TakeMaterialDeleteReturnReq) error {
	panic("not implement")
}

func (s *Server) SendTakeMaterialDeleteReturnRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendMaterialDeleteUpdateNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleMcoinExchangeHcoinReq(ctx *Context, req *pb.McoinExchangeHcoinReq) error {
	panic("not implement")
}

func (s *Server) SendMcoinExchangeHcoinRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleDestroyMaterialReq(ctx *Context, req *pb.DestroyMaterialReq) error {
	panic("not implement")
}

func (s *Server) SendDestroyMaterialRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetEquipLockStateReq(ctx *Context, req *pb.SetEquipLockStateReq) error {
	panic("not implement")
}

func (s *Server) SendSetEquipLockStateRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleCalcWeaponUpgradeReturnItemsReq(ctx *Context, req *pb.CalcWeaponUpgradeReturnItemsReq) error {
	panic("not implement")
}

func (s *Server) SendCalcWeaponUpgradeReturnItemsRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendForgeDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendForgeFormulaDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCombineDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCombineFormulaDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendClosedItemNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCheckAddItemExceedLimitNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetIsAutoUnlockSpecificEquipReq(ctx *Context, req *pb.SetIsAutoUnlockSpecificEquipReq) error {
	panic("not implement")
}

func (s *Server) SendSetIsAutoUnlockSpecificEquipRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleReliquaryDecomposeReq(ctx *Context, req *pb.ReliquaryDecomposeReq) error {
	panic("not implement")
}

func (s *Server) SendReliquaryDecomposeRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendReliquaryFilterStateSaveNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendReliquaryFilterStateNotify(ctx *Context) error {
	panic("not implement")
}
