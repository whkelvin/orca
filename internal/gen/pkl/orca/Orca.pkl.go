// Code generated from Pkl module `orca`. DO NOT EDIT.
package orca

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Orca struct {
	Name string `pkl:"name"`

	Arguments []*Argument `pkl:"arguments"`

	Script []string `pkl:"script"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Orca
func LoadFromPath(ctx context.Context, path string) (ret *Orca, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Orca
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Orca, error) {
	var ret Orca
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
