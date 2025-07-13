package snowvillage

type userOrigin string

const (
	INSTAGRAM        userOrigin = "instagram"
	GOOGLE           userOrigin = "google"
	FRIENDSANDFAMILY userOrigin = "friends-and-family"
	NEWSPAPERAD      userOrigin = "newspaper-ad"
)

func (u userOrigin) IsValid() bool {
	switch u {
	case INSTAGRAM, GOOGLE, FRIENDSANDFAMILY, NEWSPAPERAD:
		return true
	default:
		return false
	}
}
