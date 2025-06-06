package assets

type ImageType string

const (
	ImageTypeBackground ImageType = "background"
	ImageTypeCharacter  ImageType = "character"
	ImageTypeProjectile ImageType = "projectile"
)

func (it ImageType) String() string {
	return string(it)
}
