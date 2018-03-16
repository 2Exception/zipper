package consensus

import (
	"github.com/zipper-project/zipper/types"
	"github.com/zipper-project/zipper/common/mpool"
)

type ConsensusWorker struct {
	consenter Consenter
}

func (worker *ConsensusWorker) VmJob(data interface{}) (interface{}, error) {
	workerData := data.(*types.WorkerData)
	msg := workerData.GetMsg()

	worker.consenter.RecvConsensus(msg.Payload)
	return nil, nil
}

func (worker *ConsensusWorker) VmReady() bool {
	return true
}

func NewConsensusWorker(consenter Consenter) *ConsensusWorker {
	return &ConsensusWorker{
		consenter:consenter,
	}
}

func GetConsensusWorkers(workerNums int, consenter Consenter) []mpool.VmWorker {
	cssWorkers := make([]mpool.VmWorker, 0)
	for i := 0; i < workerNums; i++ {
		cssWorkers = append(cssWorkers, NewConsensusWorker(consenter))
	}

	return cssWorkers
}