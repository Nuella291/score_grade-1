package grader

func DisplayGrade(score float64) string {
	if score < 0 || score > 100 {
		return "Invalid score. Please enter a score between 0 and 100."
	}
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else {
		return "F"
	}
}
