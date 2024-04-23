package event

import (
	"context"
	"log"

	"github.com/git-fal7/luckperms/internal/database"
	"github.com/google/uuid"
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
			_, err := database.DB.GetPermission(context.Background(), database.GetPermissionParams{
				Uuid:       uuid.UUID(player.ID()),
				Permission: perm,
			})
			if err != nil {
				log.Printf("Couldnt get permission %s of player %s", perm, player.Username())
				return permission.False
			}
			return permission.True
		})
	}
}
