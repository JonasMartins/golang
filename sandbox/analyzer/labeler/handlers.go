package labeler

import (
	drawer "analyzer/drawer"
	"errors"
	"fmt"
)

func GatherData(size uint, drawer *drawer.Drawer) error {

	contests, err := drawer.GetSome(int(size), 0)
	if err != nil {
		return err
	}

	sumQuadrantsLabelResults(contests)

	return errors.New("continue")

}

func sumQuadrantsLabelResults(contests []*drawer.Contest) {

	var quadrantsPatterns = make(map[string]uint)
	var q *QuadrantsArrangement
	var pattern string
	for _, c := range contests {
		q = labelerContestQuadrants(c)
		pattern = fmt.Sprintf("Q1:%d-Q2:%d-Q3:%d-Q4:%d", q.Q1, q.Q2, q.Q3, q.Q4)
		quadrantsPatterns[pattern] = quadrantsPatterns[pattern] + 1
	}

	for index, pattern := range quadrantsPatterns {
		fmt.Printf("-> %s : %d\n", index, pattern)
	}

}

func labelerContestQuadrants(c *drawer.Contest) *QuadrantsArrangement {
	q := QuadrantsArrangement{}
	q.Q1, q.Q2, q.Q3, q.Q4 = 0, 0, 0, 0

	classifyContestQuadrant(c.Bola_1, &q)
	classifyContestQuadrant(c.Bola_2, &q)
	classifyContestQuadrant(c.Bola_3, &q)
	classifyContestQuadrant(c.Bola_4, &q)
	classifyContestQuadrant(c.Bola_5, &q)
	classifyContestQuadrant(c.Bola_6, &q)

	return &q

}

func classifyContestQuadrant(n int, q *QuadrantsArrangement) {

	if (n >= 1 && n <= 5) || (n >= 11 && n <= 15) || (n >= 21 && n <= 25) {
		q.Q1++
	}

	if (n >= 6 && n <= 10) || (n >= 16 && n <= 20) || (n >= 26 && n <= 30) {
		q.Q2++
	}

	if (n >= 31 && n <= 35) || (n >= 41 && n <= 45) || (n >= 51 && n <= 55) {
		q.Q3++
	}

	if (n >= 36 && n <= 40) || (n >= 46 && n <= 50) || (n >= 56 && n <= 60) {
		q.Q4++
	}
}
