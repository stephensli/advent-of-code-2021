package main

import (
	"strconv"
	"strings"
	"sync"

	"github.com/life4/genesis/slices"
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/cache"
	"github.com/stephensli/aoc/helpers/file"
)

var (
	OperationalCondition string = "."
	DamagedCondition     string = "#"
	UnknownCondition     string = "?"
)

type Spring struct {
	conditions               []string
	contiguousDamagedSprings []int
}

type CacheRecord struct {
	Conditions []string
	Runs       []int
}

func (c CacheRecord) String() string {
	output := strings.Join(c.Conditions, "")
	for _, v := range c.Runs {
		output += strconv.Itoa(v)
	}

	return output
}

func ComputeCombinations(cache cache.Cache[string, int], conditions []string, runs []int) (result int) {
	key := CacheRecord{Conditions: conditions, Runs: runs}.String()

	if value, ok := cache.Get(key); ok {
		return value
	}

	defer func() {
		cache.Set(key, result)
	}()

	// We have run out of possible conditions to check through but we can be in
	// two different situations, either we have no other possible runs left, which
	// means we have found a combination or we have runs left and we have not
	// found a valid combination.
	if len(conditions) == 0 {
		if len(runs) == 0 {
			return 1
		}
		return 0
	}

	// If we have no more runs left and the remaining conditions are all
	// non-broken values, then we have found another possible combination
	// otherwise its invalid.
	if len(runs) == 0 {
		for _, c := range conditions {
			if c == DamagedCondition {
				return 0
			}
		}
		return 1
	}

	// If the line is not long enough to complete the remaining run then we also
	// cannot count this as a valid combination and can exit out early.
	if len(conditions) < slices.Sum(runs)+len(runs)-1 {
		return 0
	}

	// If the next value is a fixed string, then just continue into the next loop
	// iteration.
	if conditions[0] == OperationalCondition {
		return ComputeCombinations(cache, conditions[1:], runs)
	}

	// If the next value of the run is going to be a Damaged, lets go and see if
	// we can find a valid range of springs that meet the requirement for the run,
	// otherwise return out a invalid combination.
	if conditions[0] == DamagedCondition {
		nextRun := runs[0]
		remainingRuns := runs[1:]

		for i := 0; i < nextRun; i++ {
			// During the run duration, a valid spring is blocking the entire length
			// and we cannot complete this step.
			if conditions[i] == OperationalCondition {
				return 0
			}
		}

		if len(conditions) > nextRun && conditions[nextRun] == DamagedCondition {
			return 0
		}

		if nextRun+1 < len(conditions) {
			// Continue with the checks to see if the next runs can be actioned.
			return ComputeCombinations(cache, conditions[nextRun+1:], remainingRuns)

		}

		return ComputeCombinations(cache, []string{}, remainingRuns)
	}

	return ComputeCombinations(cache, append([]string{DamagedCondition}, conditions[1:]...), runs) +
		ComputeCombinations(cache, append([]string{OperationalCondition}, conditions[1:]...), runs)
}

func parse(p2 bool, input []string) []Spring {
	springs := make([]Spring, len(input))

	for i, v := range input {
		sections := strings.Split(v, " ")

		conditions := strings.Split(sections[0], "")
		contiguousDamagedSprings := strings.Split(sections[1], ",")

		spring := Spring{
			conditions:               []string{},
			contiguousDamagedSprings: []int{},
		}

		count := 1

		if p2 {
			count = 5
		}

		for i := 0; i < count; i++ {
			for _, condition := range conditions {
				spring.conditions = append(spring.conditions, string(condition))
			}

			if p2 && i != count-1 {
				spring.conditions = append(spring.conditions, UnknownCondition)
			}
		}

		for i := 0; i < count; i++ {
			for _, damagedArrangement := range contiguousDamagedSprings {
				value, _ := strconv.Atoi(damagedArrangement)
				spring.contiguousDamagedSprings = append(spring.contiguousDamagedSprings, value)
			}
		}

		springs[i] = spring
	}

	return springs
}

func main() {
	path, complete := aoc.Setup(2023, 12, false)
	defer complete()

	springs := parse(false, file.ToTextLines(path))
	springs2 := parse(true, file.ToTextLines(path))

	var partOne int
	var partTwo int

	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	for i, s := range springs {
		wg.Add(2)

		i := i

		go func() {
			defer wg.Done()
			cacheOne := cache.New[string, int]()
			value := ComputeCombinations(cacheOne, s.conditions, s.contiguousDamagedSprings)

			mx.Lock()
			defer mx.Unlock()
			partOne += value
		}()

		go func() {
			defer wg.Done()
			cacheTwo := cache.New[string, int]()
			value := ComputeCombinations(cacheTwo, springs2[i].conditions, springs2[i].contiguousDamagedSprings)

			mx.Lock()
			defer mx.Unlock()
			partTwo += value
		}()
	}

	wg.Wait()

	aoc.PrintAnswer(1, partOne)
	aoc.PrintAnswer(2, partTwo)
}

