package clause

// PreWhere prewhere clause
type PreWhere struct {
	Exprs []Expression
}

// Name where clause name
func (prewhere PreWhere) Name() string {
	return "PREWHERE"
}

// Build build PreWhere clause
func (prewhere PreWhere) Build(builder Builder) {
	// Switch position if the first query expression is a single Or condition
	for idx, expr := range prewhere.Exprs {
		if v, ok := expr.(OrConditions); !ok || len(v.Exprs) > 1 {
			if idx != 0 {
				prewhere.Exprs[0], prewhere.Exprs[idx] = prewhere.Exprs[idx], prewhere.Exprs[0]
			}
			break
		}
	}

	buildExprs(prewhere.Exprs, builder, " AND ")
}

// MergeClause merge prewhere clauses
func (prewhere PreWhere) MergeClause(clause *Clause) {
	if w, ok := clause.Expression.(PreWhere); ok {
		exprs := make([]Expression, len(w.Exprs)+len(prewhere.Exprs))
		copy(exprs, w.Exprs)
		copy(exprs[len(w.Exprs):], prewhere.Exprs)
		prewhere.Exprs = exprs
	}

	clause.Expression = prewhere
}
