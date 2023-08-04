package model

import (
	"strings"
)

type CommitMessage struct {
	Subject     string
	Body        []string
	InvalidBody bool
}

func CreateCommitMessageFrom(messageLines []string) *CommitMessage {
	cm := &CommitMessage{InvalidBody: false}
	cm.addSubject(messageLines)
	cm.addBody(messageLines)

	return cm
}

func (cm *CommitMessage) ValidateSubject() int {
	currentSubjectLength := len(cm.Subject)

	if strings.HasPrefix(cm.Subject, "Merge ") {
		return 0
	}

	if currentSubjectLength > 50 {
		return currentSubjectLength - 50
	}

	return 0
}

func (cm *CommitMessage) addSubject(messageLines []string) {
	if len(messageLines) >= 1 {
		cm.Subject = messageLines[0]
	}
}

func (cm *CommitMessage) addBody(messageLines []string) {
	if len(messageLines) > 1 {
		cm.Body = messageLines[2:]
		if messageLines[1] != "" {
			cm.InvalidBody = true
			cm.Body = messageLines[1:]
		}
	}
}
