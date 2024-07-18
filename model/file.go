package model

type FileAttachment struct {
}

type ReqFileDeleteParam struct {
	Filename string `json:"filename" binding:"required"`
}
