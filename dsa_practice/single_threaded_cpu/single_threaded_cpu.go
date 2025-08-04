package singlethreadedcpu

import (
	"sort"

	"github.com/ishanchopra90/Practice/practice/dsa_practice/common"
)

/*
sort tasks by enqueue time, and move from enqueue time to enqeue time
start at earliest enqueue time, cpu picks up shortest processing task at this time - goes to next cpu pos
go to next enqueue time, check if cpu pos lies between prev and this or at this enqueue time - then eval next cpu pos
continue.
track shortest processing tasks with a heap
*/
type Task struct {
	Idx            int
	EnqueueTime    int
	ProcessingTime int
}

func getOrder(tasks [][]int) []int {
	var taskAry []Task
	for i, taskElement := range tasks {
		task := Task{
			Idx:            i,
			EnqueueTime:    taskElement[0],
			ProcessingTime: taskElement[1],
		}

		taskAry = append(taskAry, task)
	}

	sort.Slice(taskAry, func(i, j int) bool {
		if taskAry[i].EnqueueTime == taskAry[j].EnqueueTime {
			return taskAry[i].ProcessingTime < taskAry[j].ProcessingTime
		}

		return taskAry[i].EnqueueTime < taskAry[j].EnqueueTime
	})

	curCpuPos := 0
	var taskOrder []int // order of task indices
	minHeap := common.MinHeap[Task]{
		Compare: taskComparator,
	}

	for _, task := range taskAry {
		for curCpuPos < task.EnqueueTime {
			if len(minHeap.Heap) > 0 {
				nextTaskForCpu := minHeap.RemoveMin()
				curCpuPos = Max(curCpuPos, nextTaskForCpu.EnqueueTime) + nextTaskForCpu.ProcessingTime
				taskOrder = append(taskOrder, nextTaskForCpu.Idx)
			} else {
				curCpuPos = task.EnqueueTime
			}
		}

		minHeap.Add(task)
	}

	for len(minHeap.Heap) > 0 {
		nextTaskForCpu := minHeap.RemoveMin()
		taskOrder = append(taskOrder, nextTaskForCpu.Idx)
	}

	return taskOrder
}

func Max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func taskComparator(a Task, b Task) int {
	if a.ProcessingTime == b.ProcessingTime {
		if a.Idx == b.Idx {
			return 0
		}

		if a.Idx < b.Idx {
			return -1
		}

		if a.Idx > b.Idx {
			return 1
		}
	}

	if a.ProcessingTime > b.ProcessingTime {
		return 1
	}

	return -1
}
