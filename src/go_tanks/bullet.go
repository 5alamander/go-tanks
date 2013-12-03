package go_tanks

import (
  "math"
)

type Bullet struct {
  Id            int
  TankId        int
  Coords        *Coords
  Direction     float64
}

func NewBullet ( tank *Tank ) *Bullet {
  direction := tank.Direction + tank.Gun.Direction; 

  if direction < 0 { direction += 360 }
  if direction > 360 { direction -= 360 }

  return &Bullet{
    TankId: tank.Id,
    Coords: &Coords{ tank.Coords.X, tank.Coords.Y },
    Direction: direction,
  }
}

func ( b *Bullet ) CalculateMove ( speed int ) (*Coords, float64) {

  radDirection := (math.Pi * b.Direction) / 180
  x := b.Coords.X + int( math.Cos( radDirection ) * float64(speed) )
  y := b.Coords.Y + int( math.Sin( radDirection ) * float64(speed) )

  return &Coords{X: x, Y: y}, b.Direction
}

func ( b *Bullet ) ApplyMove ( c *Coords, d float64 ) {
  b.Coords = c
  b.Direction = d
}


