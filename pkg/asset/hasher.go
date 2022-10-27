package asset

type AbilityNameHash = uint32

func NewAbilityNameHash(name string) AbilityNameHash {
	var v uint32
	for i := 0; i < len(name); i++ {
		v = v*131 + uint32(name[i])
	}
	return AbilityNameHash(v)
}
