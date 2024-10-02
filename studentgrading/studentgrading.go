package studentgrading

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSVLineToStudent(line string) student {
	lineItems := strings.Split(line, ",")
	t1, _ := strconv.Atoi(lineItems[3])
	t2, _ := strconv.Atoi(lineItems[4])
	t3, _ := strconv.Atoi(lineItems[5])
	t4, _ := strconv.Atoi(lineItems[6])
	return student{
		firstName:  lineItems[0],
		lastName:   lineItems[1],
		university: lineItems[2],
		test1Score: t1,
		test2Score: t2,
		test3Score: t3,
		test4Score: t4,
	}
}

func parseCSV(filePath string) []student {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileLines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	students := []student{}
	for _, element := range fileLines[1:] {
		students = append(students, parseCSVLineToStudent(element))
	}
	return students
}

func getGrade(finalScore float32) Grade {
	if finalScore < 35 {
		return Grade(F)
	} else if finalScore >= 35 && finalScore < 50 {
		return Grade(C)
	} else if finalScore >= 50 && finalScore < 70 {
		return Grade(B)
	} else {
		return Grade(A)
	}
}

func calculateGrade(students []student) []studentStat {
	studentStats := []studentStat{}
	for _, student := range students {
		finalScore := float32(student.test1Score+student.test2Score+student.test3Score+student.test4Score) / 4
		grade := getGrade(finalScore)
		studentStat := studentStat{
			student:    student,
			finalScore: finalScore,
			grade:      grade,
		}
		studentStats = append(studentStats, studentStat)
	}
	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	topStudent := gradedStudents[0]
	for _, studentStat := range gradedStudents {
		if studentStat.finalScore > topStudent.finalScore {
			topStudent = studentStat
		}
	}
	return topStudent
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {

	studentPerUniMap := make(map[string][]studentStat)
	for _, currentStudent := range gs {
		studentPerUniMap[currentStudent.university] = append(studentPerUniMap[currentStudent.university], currentStudent)
	}

	topperPerUniMap := make(map[string]studentStat)
	for key, value := range studentPerUniMap {
		topperPerUniMap[key] = findOverallTopper(value)
	}
	return topperPerUniMap
}
