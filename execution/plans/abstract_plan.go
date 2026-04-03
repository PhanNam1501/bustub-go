// ===----------------------------------------------------------------------===//
//
//                         BusTub (Go port)
//
// abstract_plan.go
//
// ===----------------------------------------------------------------------===//

package plans

// Plan is an executable logical/physical plan node.
// TODO: unify plan typing and expression trees.
type Plan interface {
	isPlan()
}

// AbstractPlan is a shared base for plan nodes.
// TODO: add common fields (schema, output columns, etc.).
type AbstractPlan struct{}

func (p *AbstractPlan) isPlan() {}

// SeqScanPlan is a sequential scan plan.
type SeqScanPlan struct {
	AbstractPlan
	// TODO: table and predicate info.
}

// InsertPlan inserts tuples.
type InsertPlan struct {
	AbstractPlan
	// TODO: target table and values.
}

// UpdatePlan updates tuples.
type UpdatePlan struct {
	AbstractPlan
	// TODO: target table and update expressions.
}

// DeletePlan deletes tuples (mark delete).
type DeletePlan struct {
	AbstractPlan
	// TODO: target table and delete predicate.
}

// HashJoinPlan performs hash join.
type HashJoinPlan struct {
	AbstractPlan
	// TODO: join keys and children plans.
}

// NestedLoopJoinPlan performs nested loop join.
type NestedLoopJoinPlan struct {
	AbstractPlan
	// TODO: join keys and children plans.
}

// AggregationPlan performs GROUP BY aggregations.
type AggregationPlan struct {
	AbstractPlan
	// TODO: group keys and aggregation expressions.
}

// LimitPlan applies LIMIT.
type LimitPlan struct {
	AbstractPlan
	// TODO: limit/offset values.
}

// SortPlan applies ORDER BY.
type SortPlan struct {
	AbstractPlan
	// TODO: sort keys.
}

// TopNPlan applies TOP N.
type TopNPlan struct {
	AbstractPlan
	// TODO: top-n values.
}

