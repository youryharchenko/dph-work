package julog

import (
	"log"
	"maps"
)

type GoalTree struct {
	term     Term      // Term to be proven
	parent   *GoalTree // Parent goal
	children []Term    // List of subgoals
	active   int       // Index of active subgoal
	env      Subst     // Dictionary of variable mappings for the current goal
	vmap     Subst     // Variables inherited from parent
}

func Resolve(goals []Term, clauses []AbstractClause, options map[string]any) (bool, []Subst) {
	return resolve(goals, indexClauses(clauses), options)
}

func resolve(goals []Term, table ClauseTable, options map[string]any) (bool, []Subst) {
	env := get(options, "env", Subst{})
	funcs := get(options, "funcs", map[string]func(a ...any) bool{})
	mode := get(options, "mode", "all")
	search := get(options, "search", "bfs")
	occurs_check := get(options, "occurs_check", false)
	vcount := get(options, "vcount", uint(0))

	queue := []GoalTree{{term: Const{name: false}, parent: nil, children: goals,
		active: 0, env: env.(Subst), vmap: Subst{}}}
	subst := []Subst{}

	for len(queue) > 0 {
		goal := GoalTree{}
		if search == "dfs" {
			goal = queue[len(queue)-1]
			queue = queue[:len(queue)-1]
		} else {
			goal = queue[0]
			queue = queue[1:]
		}

		log.Println("Goal:", Clause{head: goal.term, body: goal.children})
		log.Println("Env", goal.env)

		if goal.active >= len(goal.children) {
			if goal.parent == nil {

				log.Println("Success: ", goal.env)

				if !containsSubst(subst, goal.env) {
					subst = append(subst, goal.env)
				}
				if mode == "all" {
					continue
				} else if mode == "any" || len(queue) == 0 {
					break
				} else {
					log.Println("Interactive not implemented")
					break
				}

			}

			log.Println("Done, returning to parent.")

			parent := goal.parent
			vmap := Subst{}

			for k, v := range goal.vmap {
				switch cvar := v.(type) {
				case Var:
					vv, ok := goal.env[cvar]
					if !ok {
						continue
					}
					if k == vv {
						continue
					}
					vmap[k] = vv
				}

			}
			parent.env = compose(parent.env, vmap)
			queue = append(queue, *parent)

			continue
		}

		term := goal.children[goal.active]

		log.Println("Subgoal: ", term)

		// Handle built-in special terms ToDo

		vmap := Subst{}
		term = freshen(substitute(term, goal.env), vmap, &vcount)

		matched_clauses := retrieveClauses(table, term, funcs)
		matched := false

		for _, c := range matched_clauses {
			c = freshen(c, Subst{}, &vcount)

			unifier := unify(term, c.head, occurs_check, funcs)
			if unifier == nil {
				continue
			}

			child := GoalTree{term: c.head, parent: &goal, children: copy(c.body), active: 0, env: unifier, vmap: vmap}
			queue = append(queue, child)
			matched = true
		}
		if !matched {
			log.Println("Failed, no matching clauses.")
		}

	}
	return len(subst) > 0, subst
}

func get(options map[string]any, key string, def any) any {
	v, ok := options[key]
	if ok {
		return v
	}
	return def

}

func containsSubst(substs []Subst, subst Subst) bool {
	for _, s := range substs {
		if maps.Equal(s, subst) {
			return true
		}
	}
	return false
}
