package maxprofitassignment

import (
	"fmt"
	"sort"

	"github.com/ishanchopra90/Practice/practice/dsa_practice/common"
)

/*
map difficulty to profit
sort on difficulty
find maxsofar profit in sorted difficulty ary using bin search
find maxPossible profit

note: difficulty to profit map can get tricky for duplicate difficulties
need diffculty array index to profit map
so instead create a Job object to avoid all this confusion
*/

type Job struct {
	Difficulty int
	Profit     int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	var jobs []Job
	for i := 0; i < len(difficulty); i++ {
		job := Job{
			Difficulty: difficulty[i],
			Profit:     profit[i],
		}

		jobs = append(jobs, job)
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Difficulty < jobs[j].Difficulty
	})

	/*
		    // uncomment to debug
			fmt.Println("jobs")
			for i := 0; i < len(jobs); i++ {
				fmt.Printf("job difficulty %v profit %v\n", jobs[i].Difficulty, jobs[i].Profit)
			}
	*/

	maxProfitAry := make([]int, len(difficulty))
	maxProfitSoFar := 0
	for i := 0; i < len(jobs); i++ {
		curProfit := jobs[i].Profit
		if maxProfitSoFar < curProfit {
			maxProfitSoFar = curProfit
		}

		fmt.Printf("idx %v cur profit %v maxprofitsofar %v\n", i, curProfit, maxProfitSoFar)
		maxProfitAry[i] = maxProfitSoFar
	}

	totalProfit := 0
	for i := 0; i < len(worker); i++ {
		workerMaxDifficulty := worker[i]

		// binary search for this difficulty in sorted difficulty array
		// find idx of greatest element lesser than or equal to workerMaxDifficulty
		jobToFind := Job{
			Difficulty: workerMaxDifficulty,
		}

		foundIdxs := common.BinarySearchIdxLesserThanDuplicatesTypeGeneric(jobToFind, jobs, compareJobs)
		if foundIdxs[0] == -1 {
			continue
		}

		foundMaxProfit := 0
		for i := foundIdxs[0]; i <= foundIdxs[1]; i++ {
			if foundMaxProfit < maxProfitAry[i] {
				foundMaxProfit = maxProfitAry[i]
			}
		}

		totalProfit += foundMaxProfit
	}

	return totalProfit
}

func compareJobs(a Job, b Job) int {
	if a.Difficulty == b.Difficulty {
		return 0
	}

	if a.Difficulty < b.Difficulty {
		return -1
	}

	return 1
}
