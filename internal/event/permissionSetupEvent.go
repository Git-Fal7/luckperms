package event

import (
	"context"

	"github.com/git-fal7/luckperms/internal/database"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/pkg/util/permission"
)

func permSetupEvent() func(*proxy.PermissionsSetupEvent) {
	return func(e *proxy.PermissionsSetupEvent) {
		e.SetFunc(func(perm string) permission.TriState {
			player, ok := e.Subject().(proxy.Player)
			if !ok { // Means that its not a player (the console).
				return permission.True
			}
			result, err := database.DB.UserHasPermission(context.Background(), database.UserHasPermissionParams{
				Uuid:       player.ID().String(),
				Permission: perm,
			})
			if err != nil || result == 0 {
				return permission.False
			}
			return permission.True
		})
	}
}
