package leaderrotation

import (
	"github.com/relab/hotstuff"
	"github.com/relab/hotstuff/modules"
)

type roundRobin struct {
	mod *modules.Modules
}

func (rr *roundRobin) InitModule(hs *modules.Modules, _ *modules.OptionsBuilder) {
	rr.mod = hs
}

// GetLeader returns the id of the leader in the given view
func (rr roundRobin) GetLeader(view hotstuff.View) hotstuff.ID {
	// TODO: does not support reconfiguration
	// assume IDs start at 1
	return hotstuff.ID(view%hotstuff.View(rr.mod.Configuration().Len()) + 1)
}

// NewRoundRobin returns a new round-robin leader rotation implementation.
func NewRoundRobin() modules.LeaderRotation {
	return &roundRobin{}
}
