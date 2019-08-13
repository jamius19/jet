package jet

// OrderByClause
type OrderByClause interface {
	serializeForOrderBy(statement StatementType, out *SqlBuilder)
}

type orderByClauseImpl struct {
	expression Expression
	ascent     bool
}

func (o *orderByClauseImpl) serializeForOrderBy(statement StatementType, out *SqlBuilder) {
	if o.expression == nil {
		panic("jet: nil expression in ORDER BY clause")
	}

	o.expression.serializeForOrderBy(statement, out)

	if o.ascent {
		out.WriteString("ASC")
	} else {
		out.WriteString("DESC")
	}
}

func newOrderByClause(expression Expression, ascent bool) OrderByClause {
	return &orderByClauseImpl{expression: expression, ascent: ascent}
}
