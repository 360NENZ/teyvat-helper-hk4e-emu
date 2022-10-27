package asset

import (
	"reflect"
	"testing"
)

func TestNewAbilityNameHash(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want AbilityNameHash
	}{{
		name: "Avatar_PlayerBoy_NormalAttack_DamageHandler",
		args: args{name: "Avatar_PlayerBoy_NormalAttack_DamageHandler"},
		want: 2957764605,
	}, {
		name: "Avatar_Player_FlyingBomber",
		args: args{name: "Avatar_Player_FlyingBomber"},
		want: 1410219662,
	}, {
		name: "Avatar_Player_CamCtrl",
		args: args{name: "Avatar_Player_CamCtrl"},
		want: 1474894886,
	}, {
		name: "Avatar_PlayerBoy_FallingAnthem",
		args: args{name: "Avatar_PlayerBoy_FallingAnthem"},
		want: 937205334,
	}, {
		name: "GrapplingHookSkill_Ability",
		args: args{name: "GrapplingHookSkill_Ability"},
		want: 1771196189,
	}, {
		name: "Avatar_DefaultAbility_VisionReplaceDieInvincible",
		args: args{name: "Avatar_DefaultAbility_VisionReplaceDieInvincible"},
		want: 2306062007,
	}, {
		name: "Avatar_DefaultAbility_AvartarInShaderChange",
		args: args{name: "Avatar_DefaultAbility_AvartarInShaderChange"},
		want: 3105629177,
	}, {
		name: "Avatar_SprintBS_Invincible",
		args: args{name: "Avatar_SprintBS_Invincible"},
		want: 3771526669,
	}, {
		name: "Avatar_Freeze_Duration_Reducer",
		args: args{name: "Avatar_Freeze_Duration_Reducer"},
		want: 100636247,
	}, {
		name: "Avatar_Attack_ReviveEnergy",
		args: args{name: "Avatar_Attack_ReviveEnergy"},
		want: 1564404322,
	}, {
		name: "Avatar_Component_Initializer",
		args: args{name: "Avatar_Component_Initializer"},
		want: 497711942,
	}, {
		name: "Avatar_HDMesh_Controller",
		args: args{name: "Avatar_HDMesh_Controller"},
		want: 3531639848,
	}, {
		name: "Avatar_Trampoline_Jump_Controller",
		args: args{name: "Avatar_Trampoline_Jump_Controller"},
		want: 4255783285,
	}, {
		name: "Avatar_PlayerBoy_ExtraAttack_Common",
		args: args{name: "Avatar_PlayerBoy_ExtraAttack_Common"},
		want: 1042696700,
	}, {
		name: "Avatar_FallAnthem_Achievement_Listener",
		args: args{name: "Avatar_FallAnthem_Achievement_Listener"},
		want: 825255509,
	}, {
		name: "Avatar_PlayerGirl_NormalAttack_DamageHandler",
		args: args{name: "Avatar_PlayerGirl_NormalAttack_DamageHandler"},
		want: 4291357363,
	}, {
		name: "Avatar_PlayerGirl_FallingAnthem",
		args: args{name: "Avatar_PlayerGirl_FallingAnthem"},
		want: 3832178184,
	}, {
		name: "Avatar_PlayerGirl_ExtraAttack_Common",
		args: args{name: "Avatar_PlayerGirl_ExtraAttack_Common"},
		want: 3374327026,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAbilityNameHash(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAbilityNameHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
