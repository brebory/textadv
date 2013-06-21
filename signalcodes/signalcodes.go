package signalcodes

type SignalCode int

const(
  OK SignalCode = iota
  FAIL
  ERROR
  STOP
  CONTINUE
)

