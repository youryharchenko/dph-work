package julog

type ClauseSubtable map[string][]AbstractClause
type ClauseTable map[string]ClauseSubtable

func indexClauses(clauses []AbstractClause) ClauseTable {
	return insertClauses(ClauseTable{}, clauses)
}

func insertClauses(table ClauseTable, clauses []AbstractClause) ClauseTable {
	// Ensure no duplicates are added
	clauses = unique(clauses)
	if len(table) > 0 {
		clauses = setdiff(clauses, deindexClauses(table))
	}
	// Iterate over clauses and insert into table
	for _, c := range clauses {
		insertClause(table, c)
	}
	return table
}

func insertClause(table ClauseTable, clause AbstractClause) {
	subtable, ok := table[clause.Head().Name()]
	if !ok {
		subtable = ClauseSubtable{}
		table[clause.Head().Name()] = subtable
	}
	switch head := clause.Head().(type) {
	case Compound:
		if len(head.Args()) > 0 {
			arg := head.Args()[0]
			switch a := arg.(type) {
			case Var:
				v, ok := subtable["__var__"]
				if !ok {
					v = []AbstractClause{}
					subtable["__var__"] = v
				}
				v = append(v, clause)
			default:
				v, ok := subtable[a.Name()]
				if !ok {
					v = []AbstractClause{}
					subtable[a.Name()] = v
				}
				v = append(v, clause)
			}
			v, ok := subtable["__all__"]
			if !ok {
				v = []AbstractClause{}
				subtable["__all__"] = v
			}
			v = append(v, clause)
		} else {
			v, ok := subtable["____no_args____"]
			if !ok {
				v = []AbstractClause{}
				subtable["____no_args____"] = v
			}
			v = append(v, clause)
		}
	default:
		v, ok := subtable["____no_args____"]
		if !ok {
			v = []AbstractClause{}
			subtable["____no_args____"] = v
		}
		v = append(v, clause)
	}
}

func unique(clauses []AbstractClause) []AbstractClause {
	res := []AbstractClause{}
	set := map[AbstractClause]any{}
	for _, c := range clauses {
		_, ok := set[c]
		if !ok {
			set[c] = nil
			res = append(res, c)
		}
	}
	return res
}

func setdiff(clauses []AbstractClause, diff []AbstractClause) []AbstractClause {
	res := []AbstractClause{}
	set := map[AbstractClause]any{}
	for _, c := range diff {
		set[c] = nil
	}
	for _, c := range clauses {
		_, ok := set[c]
		if !ok {
			res = append(res, c)
		}
	}
	return res
}

func deindexClauses(table ClauseTable) []AbstractClause {
	res := []AbstractClause{}

	for _, s := range table {
		v, ok := s["__no_args__"]
		if ok {
			res = append(res, v...)
		} else {
			res = append(res, s["__all__"]...)
		}
	}
	return res
}
