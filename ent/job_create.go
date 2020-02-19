// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kcarretto/paragon/ent/job"
	"github.com/kcarretto/paragon/ent/tag"
	"github.com/kcarretto/paragon/ent/task"
	"github.com/kcarretto/paragon/ent/user"
)

// JobCreate is the builder for creating a Job entity.
type JobCreate struct {
	config
	Name         *string
	CreationTime *time.Time
	Content      *string
	Staged       *bool
	tasks        map[int]struct{}
	tags         map[int]struct{}
	prev         map[int]struct{}
	next         map[int]struct{}
	owner        map[int]struct{}
}

// SetName sets the Name field.
func (jc *JobCreate) SetName(s string) *JobCreate {
	jc.Name = &s
	return jc
}

// SetCreationTime sets the CreationTime field.
func (jc *JobCreate) SetCreationTime(t time.Time) *JobCreate {
	jc.CreationTime = &t
	return jc
}

// SetNillableCreationTime sets the CreationTime field if the given value is not nil.
func (jc *JobCreate) SetNillableCreationTime(t *time.Time) *JobCreate {
	if t != nil {
		jc.SetCreationTime(*t)
	}
	return jc
}

// SetContent sets the Content field.
func (jc *JobCreate) SetContent(s string) *JobCreate {
	jc.Content = &s
	return jc
}

// SetStaged sets the Staged field.
func (jc *JobCreate) SetStaged(b bool) *JobCreate {
	jc.Staged = &b
	return jc
}

// SetNillableStaged sets the Staged field if the given value is not nil.
func (jc *JobCreate) SetNillableStaged(b *bool) *JobCreate {
	if b != nil {
		jc.SetStaged(*b)
	}
	return jc
}

// AddTaskIDs adds the tasks edge to Task by ids.
func (jc *JobCreate) AddTaskIDs(ids ...int) *JobCreate {
	if jc.tasks == nil {
		jc.tasks = make(map[int]struct{})
	}
	for i := range ids {
		jc.tasks[ids[i]] = struct{}{}
	}
	return jc
}

// AddTasks adds the tasks edges to Task.
func (jc *JobCreate) AddTasks(t ...*Task) *JobCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return jc.AddTaskIDs(ids...)
}

// AddTagIDs adds the tags edge to Tag by ids.
func (jc *JobCreate) AddTagIDs(ids ...int) *JobCreate {
	if jc.tags == nil {
		jc.tags = make(map[int]struct{})
	}
	for i := range ids {
		jc.tags[ids[i]] = struct{}{}
	}
	return jc
}

// AddTags adds the tags edges to Tag.
func (jc *JobCreate) AddTags(t ...*Tag) *JobCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return jc.AddTagIDs(ids...)
}

// SetPrevID sets the prev edge to Job by id.
func (jc *JobCreate) SetPrevID(id int) *JobCreate {
	if jc.prev == nil {
		jc.prev = make(map[int]struct{})
	}
	jc.prev[id] = struct{}{}
	return jc
}

// SetNillablePrevID sets the prev edge to Job by id if the given value is not nil.
func (jc *JobCreate) SetNillablePrevID(id *int) *JobCreate {
	if id != nil {
		jc = jc.SetPrevID(*id)
	}
	return jc
}

// SetPrev sets the prev edge to Job.
func (jc *JobCreate) SetPrev(j *Job) *JobCreate {
	return jc.SetPrevID(j.ID)
}

// SetNextID sets the next edge to Job by id.
func (jc *JobCreate) SetNextID(id int) *JobCreate {
	if jc.next == nil {
		jc.next = make(map[int]struct{})
	}
	jc.next[id] = struct{}{}
	return jc
}

// SetNillableNextID sets the next edge to Job by id if the given value is not nil.
func (jc *JobCreate) SetNillableNextID(id *int) *JobCreate {
	if id != nil {
		jc = jc.SetNextID(*id)
	}
	return jc
}

// SetNext sets the next edge to Job.
func (jc *JobCreate) SetNext(j *Job) *JobCreate {
	return jc.SetNextID(j.ID)
}

// SetOwnerID sets the owner edge to User by id.
func (jc *JobCreate) SetOwnerID(id int) *JobCreate {
	if jc.owner == nil {
		jc.owner = make(map[int]struct{})
	}
	jc.owner[id] = struct{}{}
	return jc
}

// SetOwner sets the owner edge to User.
func (jc *JobCreate) SetOwner(u *User) *JobCreate {
	return jc.SetOwnerID(u.ID)
}

// Save creates the Job in the database.
func (jc *JobCreate) Save(ctx context.Context) (*Job, error) {
	if jc.Name == nil {
		return nil, errors.New("ent: missing required field \"Name\"")
	}
	if err := job.NameValidator(*jc.Name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"Name\": %v", err)
	}
	if jc.CreationTime == nil {
		v := job.DefaultCreationTime()
		jc.CreationTime = &v
	}
	if jc.Content == nil {
		return nil, errors.New("ent: missing required field \"Content\"")
	}
	if err := job.ContentValidator(*jc.Content); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"Content\": %v", err)
	}
	if jc.Staged == nil {
		v := job.DefaultStaged
		jc.Staged = &v
	}
	if len(jc.prev) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"prev\"")
	}
	if len(jc.next) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"next\"")
	}
	if len(jc.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	if jc.owner == nil {
		return nil, errors.New("ent: missing required edge \"owner\"")
	}
	return jc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (jc *JobCreate) SaveX(ctx context.Context) *Job {
	v, err := jc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jc *JobCreate) sqlSave(ctx context.Context) (*Job, error) {
	var (
		j     = &Job{config: jc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: job.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: job.FieldID,
			},
		}
	)
	if value := jc.Name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: job.FieldName,
		})
		j.Name = *value
	}
	if value := jc.CreationTime; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: job.FieldCreationTime,
		})
		j.CreationTime = *value
	}
	if value := jc.Content; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: job.FieldContent,
		})
		j.Content = *value
	}
	if value := jc.Staged; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: job.FieldStaged,
		})
		j.Staged = *value
	}
	if nodes := jc.tasks; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   job.TasksTable,
			Columns: []string{job.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jc.tags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   job.TagsTable,
			Columns: job.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jc.prev; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   job.PrevTable,
			Columns: []string{job.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jc.next; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   job.NextTable,
			Columns: []string{job.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jc.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   job.OwnerTable,
			Columns: []string{job.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, jc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	j.ID = int(id)
	return j, nil
}
