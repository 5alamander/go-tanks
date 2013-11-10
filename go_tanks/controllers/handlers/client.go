package go_tanks

import (
  //log "../log"
  i "../../interfaces"
)

var messageHandlers = map[string]func( i.World, i.Client, *i.Message )error{
  "Client" :ClientMessageHandler,
}

func HandleMessage( w i.World, c i.Client, m *i.Message ) error {
  message := *m

  return messageHandlers[message["Type"].(string)](w, c, m)
}

func ClientMessageHandler( w i.World, c i.Client, m *i.Message ) error {
  message := *m

  if message["WorldRecieveDisabled"] != nil {
    c.SetWorldRecieveDisabled(message["WorldRecieveDisabled"].(bool))
  }

  return nil
}
