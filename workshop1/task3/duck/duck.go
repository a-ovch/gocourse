package duck

type Duck interface {
	Fly() string
	Quack() string
}

type duck struct {
	flightBehavior flightBehavior
	quackBehavior  quackBehavior
}

func (d duck) Fly() string {
	return d.flightBehavior.fly()
}

func (d duck) Quack() string {
	return d.quackBehavior()
}

type DeafMuteDuck struct {
	duck
}

func NewDeafMuteDuck() *DeafMuteDuck {
	return &DeafMuteDuck{duck{noFlight{}, noQuack}}
}

type ReactiveLoudDuck struct {
	duck
}

func NewReactiveLoudDuck() *ReactiveLoudDuck {
	d := ReactiveLoudDuck{}
	d.flightBehavior = reactiveFlightBehavior{}
	d.quackBehavior = loudQuack
	return &d
}

type BackwardsQuietDuck struct {
	duck
}

func NewBackwardsQuietDuck() *BackwardsQuietDuck {
	d := BackwardsQuietDuck{}
	d.flightBehavior = backwardsFlightBehavior{}
	d.quackBehavior = quietQuack
	return &d
}
