package go_tanks

import (
  //log "../log"
  //i "../interfaces"
)

type GameController struct {
  Controller
}

func ( c *GameController ) JoinToGame () error {
  c.addToWorld()

  c.World.NewTank( c.Client )
  message := *(c.Client.ReadInBox())

  message["Message"] = "Your tank id"

  c.Client.SetTankId( message["Id"].(int) )
  c.Client.SendMessage( &message )

  //for {
    //_, err := c.Client.ReadMessage()
    //if ( err != nil ) { continue }

  //}

  return nil
}

func ( c *GameController ) addToWorld () {
  c.World.AttachClient( c.Client )
}
