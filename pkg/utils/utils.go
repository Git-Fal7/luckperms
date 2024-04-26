package utils

import (
	"context"

	"github.com/git-fal7/luckperms/internal/database"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/pkg/util/uuid"
)

func GetPrefixOfPlayer(player proxy.Player) string {
	result, err := database.DB.GetPrefixOfUser(context.Background(), player.ID().String())
	if err != nil {
		return ""
	}
	return result[0].Prefix
}

func GetSuffixOfPlayer(player proxy.Player) string {
	result, err := database.DB.GetSuffixOfUser(context.Background(), player.ID().String())
	if err != nil {
		return ""
	}
	return result[0].Suffix
}

func PlayerHasPermission(playerUUID uuid.UUID, perm string) bool {
	result, err := database.DB.UserHasPermission(context.Background(), database.UserHasPermissionParams{
		Uuid:       playerUUID.String(),
		Permission: perm,
	})
	if err != nil || result == 0 {
		return false
	}
	return true
}
