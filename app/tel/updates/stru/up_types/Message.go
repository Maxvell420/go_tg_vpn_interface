package up_types

import (
	"GO/app/tel/updates/stru/up_types/update_particles"
)

type Message struct {
	Message_id int
	From       update_particles.From
	Chat       update_particles.Chat
	Entities   []update_particles.Entity
	Date       int
	Text       *string
}
