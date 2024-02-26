package services

import "net/smtp"

func SendEmail(sendingEmail string) {
	auth := smtp.PlainAuth("", "noreply@speedat.site", "pwd", "speedat.site")

	from := "noreply@speedat.site"
	to := []string{sendingEmail}

	headerSubject := "Subject: 테스트/r/n"
	headerBlank := "\r\n"
	body := "메일 테스트입니다\r\n"
	msg := []byte(headerSubject + headerBlank + body)

	// 메일 보내기
	err := smtp.SendMail("noreply@speedat.site", auth, from, to, msg)
	if err != nil {
		panic(err)
	}
}
