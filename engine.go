// textadv is an interactive fiction game engine
package textadv

import (
  "textadv/signalcodes"
)

type Describer interface {
  Describe() string
}

type DescriptionEditor interface {
  EditDescription(string)
}

type Engine interface {
  Run() signalcodes.SignalCode
}

type Manager interface {
  Signal(signalcodes.SignalCode, chan<- signalcodes.SignalCode)
  HandleSignals(<-chan signalcodes.SignalCode) signalcodes.SignalCode
  Stop() signalcodes.SignalCode
}

type EngineManager struct {
  signals chan signalcodes.SignalCode
}

type TextAdvEngine struct {
  entities []Entity
  EngineManager
}

type Entity struct {
  description string
}

type EntityCollection struct {
  description string
  entities []Entity
}

func CreateEntity(description string) (*Entity) {
  e := Entity{ description: description }
  return &e
}

func (self *Entity) Describe() string {
  return self.description
}

func (self *Entity) EditDescription(newDescription string) {
  self.description = newDescription
}

func (self *EntityCollection) Describe() string {
  return self.description
}

func (self *EntityCollection) EditDescription(newDescription string) {
  self.description = newDescription
}
