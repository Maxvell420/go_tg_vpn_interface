package up_types

import (
	"GO/app/tel/updates/struct/up_types/update_particles"
)

type MyChatMember struct {
	Chat update_particles.Chat
	From update_particles.Chat
}
