package src

import (
	"rockstaedt/commit-message-check/src/utils"
)

type CommitMessage struct {
	Subject     string
	Body        []string
	InvalidBody bool
}

func CreateCommitMessageFrom(inputPath string) (*CommitMessage, error) {
	messageLines, err := utils.GetLinesFromTextFile(inputPath)
	if err != nil {
		return nil, err
	}

	cm := &CommitMessage{InvalidBody: false}
	cm.addSubject(messageLines)
	cm.addBody(messageLines)

	return cm, nil
}

func (cm *CommitMessage) ValidateSubject() bool {
	if len(cm.Subject) > 50 {
		return false
	}

	return true
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
