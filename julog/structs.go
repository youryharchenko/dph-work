package julog

type Term interface {
	Name() string
	Args() []Term
}

type Const struct {
	Term
	name any
}

type Var struct {
	Term
	name any
}

type Compound struct {
	Term
	name string
	args []Term
}

type AbstractClause interface {
	Head() Term
}

type Clause struct {
	AbstractClause
	head Term
	body []Term
}

func (c *Clause) Head() Term {
	return c.head
}

type Subst map[Var]Term
