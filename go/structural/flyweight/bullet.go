package flyweight

import (
	"fmt"
)

// BulletType enum
type BulletType int

const (
	NINE_MM BulletType = iota
	ELEVEN_MM
	ACP
)

// Bullet struct representing the flyweight object
type Bullet struct {
	image  string
	radius float64
	weight float64
	bType  BulletType
}

func NewBullet(image string, radius, weight float64, bType BulletType) *Bullet {
	return &Bullet{
		image:  image,
		radius: radius,
		weight: weight,
		bType:  bType,
	}
}

func (b *Bullet) Image() string {
	return b.image
}

func (b *Bullet) Radius() float64 {
	return b.radius
}

func (b *Bullet) Weight() float64 {
	return b.weight
}

func (b *Bullet) Type() BulletType {
	return b.bType
}

// BulletRegistry to manage shared Bullet instances
type BulletRegistry struct {
	bullets map[BulletType]*Bullet
}

func NewBulletRegistry() *BulletRegistry {
	return &BulletRegistry{
		bullets: make(map[BulletType]*Bullet),
	}
}

func (br *BulletRegistry) AddBullet(bullet *Bullet) {
	br.bullets[bullet.Type()] = bullet
}

func (br *BulletRegistry) GetBullet(bType BulletType) *Bullet {
	return br.bullets[bType]
}

// FlyingBullet struct representing the extrinsic state
type FlyingBullet struct {
	x, y, z   float64
	direction float64
	bullet    *Bullet
}

func NewFlyingBullet(x, y, z, direction float64, bullet *Bullet) *FlyingBullet {
	return &FlyingBullet{
		x:         x,
		y:         y,
		z:         z,
		direction: direction,
		bullet:    bullet,
	}
}

func (fb *FlyingBullet) Details() string {
	return fmt.Sprintf("FlyingBullet details: x=%.1f, y=%.1f, z=%.1f, direction=%.1f, bullet_type=%d, bullet_image=%s, bullet_radius=%.1f, bullet_weight=%.1f",
		fb.x, fb.y, fb.z, fb.direction, fb.bullet.Type(), fb.bullet.Image(), fb.bullet.Radius(), fb.bullet.Weight())
}
