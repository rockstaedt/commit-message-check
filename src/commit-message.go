package src

import (
	"errors"
	"regexp"
)

type CommitMessage struct {
	Subject     string
	Body        []string
	InvalidBody bool
}

func CreateCommitMessageFrom(messageLines []string) (*CommitMessage, error) {
	cm := &CommitMessage{InvalidBody: false}
	cm.addSubject(messageLines)
	cm.addBody(messageLines)

	return cm, nil
}

func (cm *CommitMessage) ValidateSubject() (bool, error) {
	if len(cm.Subject) > 72 {
		return false, errors.New("subject length exceeds 72 characters")
	}

	if len(cm.Subject) > 50 {
		re := regexp.MustCompile(`^#\d+ -\s*(.*)$`)
		trimmedSubject := re.ReplaceAllString(cm.Subject, `$1`)

		return len(trimmedSubject) < 51, nil
	}

	return true, nil
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
