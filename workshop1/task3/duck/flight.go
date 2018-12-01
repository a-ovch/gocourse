package duck

type flightBehavior interface {
	fly() string
}

type noFlight struct{}

func (fb noFlight) fly() string {
	return "I can't fly!"
}

type backwardsFlightBehavior struct{}

func (fb backwardsFlightBehavior) fly() string {
	return "It's very strange but I can fly only backwards"
}

type reactiveFlightBehavior struct{}

func (fb reactiveFlightBehavior) fly() string {
	return "🚀"
}
