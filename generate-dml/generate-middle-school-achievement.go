package main

import (
	"bytes"
	"fmt"
	"math/rand"
	_type "themoment-team/hellogsm-notice-server/generate-dml/type"
)

// 랜덤 배열 생성 함수
func randomIntArray(length int, min int, max int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(max-min+1) + min
	}
	return arr
}

// GraduateStatus에 따라 중학교 성취도 데이터를 생성하는 함수
func GenerateMiddleSchoolAchievementInsertQuery(rows int, graduateStatuses []_type.GraduateStatus) string {
	var buffer bytes.Buffer

	buffer.WriteString("-- tb_middle_school_achievement" + "\n\n")

	artsPhysicalSubjects := `["체육","미술","음악"]`
	generalSubjects := `["국어","도덕","사회","역사","수학","과학","기술가정","영어"]`
	newSubjects := `["프로그래밍"]`

	for i := 1; i <= rows; i++ {
		graduateStatus := graduateStatuses[i-1]
		var gedTotalScore string
		var absentDays, achievement_2_1, achievement_2_2, achievement_3_1, achievement_3_2, artsPhysicalAchievement, attendanceDays, volunteerTime string
		var freeSemester, liberalSystem string

		switch graduateStatus {
		case _type.CANDIDATE:
			gedTotalScore = "NULL"
			absentDays = fmt.Sprintf("%v", randomIntArray(3, 0, 3))
			achievement_2_1 = fmt.Sprintf("%v", randomIntArray(9, 1, 5))
			achievement_2_2 = fmt.Sprintf("%v", randomIntArray(9, 1, 5))
			achievement_3_1 = fmt.Sprintf("%v", randomIntArray(9, 1, 5))
			achievement_3_2 = "NULL"
			artsPhysicalAchievement = fmt.Sprintf("%v", randomIntArray(9, 3, 5))
			attendanceDays = fmt.Sprintf("%v", randomIntArray(9, 1, 5))
			volunteerTime = fmt.Sprintf("%v", randomIntArray(3, 0, 5))
			freeSemester = "NULL"
			liberalSystem = "자유학년제"

		case _type.GRADUATE:
			gedTotalScore = "NULL"
			absentDays = fmt.Sprintf("%v", randomIntArray(3, 0, 3))
			achievement_2_1 = fmt.Sprintf("%v", randomIntArray(9, 1, 5))
			achievement_2_2 = fmt.Sprintf("%v", randomIntArray(9, 1, 5))
			achievement_3_1 = fmt.Sprintf("%v", randomIntArray(9, 1, 5))
			achievement_3_2 = fmt.Sprintf("%v", randomIntArray(9, 1, 5))
			artsPhysicalAchievement = fmt.Sprintf("%v", randomIntArray(9, 3, 5))
			attendanceDays = fmt.Sprintf("%v", randomIntArray(9, 1, 5))
			volunteerTime = fmt.Sprintf("%v", randomIntArray(3, 0, 5))
			freeSemester = "NULL"
			liberalSystem = "자유학년제"

		case _type.GED:
			gedTotalScore = fmt.Sprintf("%d", rand.Intn(201)+400) // 400 ~ 600 사이의 랜덤 값
			absentDays = "NULL"
			achievement_2_1 = "NULL"
			achievement_2_2 = "NULL"
			achievement_3_1 = "NULL"
			achievement_3_2 = "NULL"
			attendanceDays = "NULL"
			artsPhysicalAchievement = "NULL"
			volunteerTime = "NULL"
			freeSemester = "NULL"
			liberalSystem = "NULL"
			artsPhysicalSubjects = "NULL"
			generalSubjects = "NULL"
			newSubjects = "NULL"
		}

		query := fmt.Sprintf(
			"INSERT INTO tb_middle_school_achievement (oneseo_id, ged_total_score, absent_days, achievement_1_2, achievement_2_1, achievement_2_2, achievement_3_1, achievement_3_2, arts_physical_achievement, arts_physical_subjects, attendance_days, free_semester, general_subjects, liberal_system, new_subjects, volunteer_time) "+
				"VALUES (%d, '%s', '%s', 'NULL', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');",
			i, gedTotalScore, absentDays, achievement_2_1, achievement_2_2, achievement_3_1, achievement_3_2, artsPhysicalAchievement, artsPhysicalSubjects,
			attendanceDays, freeSemester, generalSubjects, liberalSystem, newSubjects, volunteerTime,
		)

		buffer.WriteString(query + "\n")
	}

	return buffer.String()
}
