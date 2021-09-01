// Code generated by entc, DO NOT EDIT.

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// CloudEvents is the predicate function for cloudevents builders.
type CloudEvents func(*sql.Selector)

// Namespace is the predicate function for namespace builders.
type Namespace func(*sql.Selector)

// Workflow is the predicate function for workflow builders.
type Workflow func(*sql.Selector)

// WorkflowEvents is the predicate function for workflowevents builders.
type WorkflowEvents func(*sql.Selector)

// WorkflowEventsWait is the predicate function for workfloweventswait builders.
type WorkflowEventsWait func(*sql.Selector)

// WorkflowInstance is the predicate function for workflowinstance builders.
type WorkflowInstance func(*sql.Selector)
