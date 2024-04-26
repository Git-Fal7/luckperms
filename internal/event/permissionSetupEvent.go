package event

import (
	"github.com/git-fal7/luckperms/pkg/utils"
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
			if utils.PlayerHasPermission(player.ID(), perm) {
				return permission.True
			}
			return permission.False
		})
	}
}
