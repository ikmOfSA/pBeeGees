package leaderrotation

import (
	"fmt"

	"github.com/relab/hotstuff"
	"github.com/relab/hotstuff/consensus"
)

type repBased struct {
	mods *consensus.Modules
}

//InitConsensusModule gives the module a reference to the Modules object.
//It also allows the module to set module options using the OptionsBuilder
func (r *repBased) InitConsensusModule(mods *consensus.Modules, _ *consensus.OptionsBuilder) {
	r.mods = mods
}

//GetLeader returns the id of the leader in the given view
func (r repBased) GetLeader(view consensus.View) hotstuff.ID {
	//assume IDS start at 1'
	fmt.Println("The leader now is: ", hotstuff.ID(view%consensus.View(r.mods.Configuration().Len())+1))
	return hotstuff.ID(view%consensus.View(r.mods.Configuration().Len()) + 1)
}

//NewRepBased returns a new random reputation-based leader rotation implementation
func NewRepBased() consensus.LeaderRotation {
	return &repBased{}
}
