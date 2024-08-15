package main

import (
	"fmt"
	"math"
)

type Input struct {
	SubjectsInAMonth     map[string]float64
	GradesSumPerSubjects map[string]float64
}

type School struct {
	Grades map[string]float64
	Areas  map[string]map[string]float64
}

var Settings = School{
	Grades: map[string]float64{
		"5/5": 9,
		"5":   7,
		"4":   3,
		"3":   1,
		"2":   0,
		"gm":  0,
	},
	Areas: map[string]map[string]float64{
		"tebigy": {
			Astronomy: 20,
			Chemistry: 25,
			Biology:   45,
			Physics:   50,
		},
		"takyk": {
			Math:        60,
			Informatics: 40,
			Geometry:    50,
		},
		"ynsanperwer": {
			English:    25,
			History:    30,
			Literature: 35,
		},
	},
}

const (
	Astronomy   = "astronomy"
	Chemistry   = "chemistry"
	Biology     = "biology"
	Physics     = "physics"
	Math        = "math"
	Informatics = "informatics"
	Geometry    = "geometry"
	English     = "english"
	History     = "history"
	Literature  = "literature"
)

var Subjects = []string{
	Astronomy,
	Chemistry,
	Biology,
	Physics,
	Math,
	Informatics,
	Geometry,
	English,
	History,
	Literature,
}

func getSchool() School {
	return Settings
}

func getInput(school School) Input {
	return Input{
		SubjectsInAMonth: map[string]float64{
			Astronomy:   4,
			Chemistry:   4,
			Biology:     4,
			Physics:     12,
			Math:        12,
			Informatics: 8,
			Geometry:    4,
			English:     12,
			History:     4,
			Literature:  8,
		},
		GradesSumPerSubjects: map[string]float64{
			Astronomy:   float64(school.Grades["5/5"]*1 + school.Grades["4"]*1),
			Chemistry:   float64(school.Grades["5"] * 2),
			Biology:     float64(school.Grades["4"] * 2),
			Physics:     float64(school.Grades["5"]*6 + school.Grades["4"]*1),
			Math:        float64(school.Grades["5"] * 9),
			Informatics: float64(school.Grades["5"] * 6),
			Geometry:    float64(school.Grades["5"]*1 + school.Grades["3"]*1),
			English:     float64(school.Grades["5"] * 7),
			History:     float64(school.Grades["5"] * 1),
			Literature:  float64(school.Grades["2"] * 1),
		},
	}
}

func main() {
	school := getSchool()
	input := getInput(school)

	subjectsCapacity := make(map[string]float64)
	subjectsProgressPercentage := make(map[string]float64)

	for _, subject := range Subjects {
		subjectsCapacity[subject] = input.SubjectsInAMonth[subject] * school.Grades["5"]
		subjectsProgressPercentage[subject] = round(input.GradesSumPerSubjects[subject]/subjectsCapacity[subject]*100, 2)
	}

	areasProgress := make(map[string]float64)
	for area, subjectsRequiredProgress := range school.Areas {
		areaProgress := []float64{}
		for subject, requiredProgress := range subjectsRequiredProgress {
			if subjectsProgressPercentage[subject] < requiredProgress {
				successPercentage := subjectsProgressPercentage[subject] / requiredProgress
				areaProgress = append(areaProgress, successPercentage)
			} else {
				areaProgress = append(areaProgress, 1)
			}
		}

		areasProgress[area] = round(sum(areaProgress)/float64(len(areaProgress))*100, 2)
	}

	areasProgressPercentage := make(map[string]float64)
	totalProgress := sum(values(areasProgress))
	for area, progress := range areasProgress {
		if progress > 0 {
			areasProgressPercentage[area] = round(progress/totalProgress*100, 2)
		}
	}

	printResults("Subjects Success Percentage", subjectsProgressPercentage)
	printResults("Area Progress", areasProgress)
	printResults("Analytics", areasProgressPercentage)
}

func printResults(title string, data map[string]float64) {
	fmt.Printf("[%s]\n", title)
	for key, value := range data {
		fmt.Printf("%s %.2f%%\n", key, value)
	}
	fmt.Println()
}

func round(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func sum(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total
}

func values(m map[string]float64) []float64 {
	arr := []float64{}
	for _, v := range m {
		arr = append(arr, v)
	}
	return arr
}
