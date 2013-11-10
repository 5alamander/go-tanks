package go_tanks

type Message            map[string]interface{}
type MessageChan        chan *Message

type Client interface {
  SendMessage ( *Message ) error
  ReadMessage () ( *Message, error )

  SetAuthCredentials ( login, password string)
  SetTankId ( int )

  Login () *string
  Password () *string

  OutBox () MessageChan
  InBox () MessageChan

  ReadInBox () *Message
  WriteInBox ( *Message )
  ReadOutBox () *Message
  WriteOutBox ( *Message )

  SendWorld ( *Message )
  SetWorldRecieveDisabled( bool )
}

func ( m *Message ) GetType () interface{} {
  return (*m)["Type"].(int)
}

func ErrorMessage ( message string ) *Message {
  return &Message{"Type":"Error", "Message": message}
}
