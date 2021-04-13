package models

type Word struct {
	Word     string `json:"word"`
	Meaning  string `json:"meaning"`
	Learnt   int    `json:"learnt"`
	LearntAt string `json:"learntAt"`
}

type Word2 struct {
	Word string `json:"word"`
	/*Meaning  sql.NullString `json:"meaning"`
	Learnt   sql.NullInt64  `json:"learnt"`
	LearntAt sql.NullString `json:"learntAt"`*/
}

type ShowWord struct {
	SWord    string `json:"word"`
	SMeaning string `json:"meaning"`
}

type SearchWord struct {
	Word string `json:"word"`
}
