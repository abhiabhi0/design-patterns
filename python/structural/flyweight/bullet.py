from enum import Enum
from typing import Dict

class BulletType(Enum):
    NINE_MM = 1
    ELEVEN_MM = 2
    ACP = 3

class Bullet:
    def __init__(self, image: str, radius: float, weight: float, bullet_type: BulletType):
        self.__image = image
        self.__radius = radius
        self.__weight = weight
        self.__type = bullet_type

    @property
    def image(self) -> str:
        return self.__image

    @property
    def radius(self) -> float:
        return self.__radius

    @property
    def weight(self) -> float:
        return self.__weight

    @property
    def type(self) -> BulletType:
        return self.__type

class BulletRegistry:
    def __init__(self):
        self.__bullets: Dict[BulletType, Bullet] = {}

    def add_bullet(self, bullet: Bullet):
        self.__bullets[bullet.type] = bullet

    def get_bullet(self, bullet_type: BulletType) -> Bullet:
        return self.__bullets.get(bullet_type)

class FlyingBullet:
    def __init__(self, x: float, y: float, z: float, direction: float, bullet: Bullet):
        self.__x = x
        self.__y = y
        self.__z = z
        self.__direction = direction
        self.__bullet = bullet

if __name__ == "__main__":
    # Create a BulletRegistry instance
    registry = BulletRegistry()

    # Add bullets to the registry
    registry.add_bullet(Bullet("9mm.png", 9.0, 7.5, BulletType.NINE_MM))
    registry.add_bullet(Bullet("11mm.png", 11.0, 8.0, BulletType.ELEVEN_MM))
    registry.add_bullet(Bullet("acp.png", 12.0, 9.0, BulletType.ACP))

    # Retrieve a bullet from the registry
    bullet = registry.get_bullet(BulletType.NINE_MM)

    # Create a FlyingBullet instance
    flying_bullet = FlyingBullet(0.0, 0.0, 0.0, 90.0, bullet)

    # Print the details of the FlyingBullet instance
    print(f"FlyingBullet details: x={flying_bullet._FlyingBullet__x}, "
          f"y={flying_bullet._FlyingBullet__y}, z={flying_bullet._FlyingBullet__z}, "
          f"direction={flying_bullet._FlyingBullet__direction}, "
          f"bullet_type={flying_bullet._FlyingBullet__bullet.type.name}, "
          f"bullet_image={flying_bullet._FlyingBullet__bullet.image}, "
          f"bullet_radius={flying_bullet._FlyingBullet__bullet.radius}, "
          f"bullet_weight={flying_bullet._FlyingBullet__bullet.weight}")
