package duck

type quackBehavior func() string

var noQuack quackBehavior = func() string {
	return "**silence**"
}

var quietQuack quackBehavior = func() string {
	return "quack..."
}

var loudQuack quackBehavior = func() string {
	return "QUACK!QUACK!QUACK"
}
