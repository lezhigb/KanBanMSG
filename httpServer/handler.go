package httpServer

import (
	"errors"
	"fmt"
	"github.com/lezhigb/KanBanMSG/dingCaller"
)

//Cards
func onCreateCard(paras []string) error {
	if len(paras) != 7 {
		return errors.New("unknown para")
	}
	go dingCaller.SendDingMsg("创建卡片", fmt.Sprintf("```%s```在阶段:```%s```的```%s```添加了```%s```", paras[1], paras[4], paras[3], paras[2]), paras[6])
	return nil
}
func onMoveCard(paras []string) error {

	return nil
}
func onArchivedCard(paras []string) error {

	return nil
}
func onRestoredCard(paras []string) error {

	return nil
}

//Card content
func onAddComment(paras []string) error {

	return nil
}
func onAddedLabel(paras []string) error {

	return nil
}
func onJoinMember(paras []string) error {

	return nil
}
func onSetCustomField(paras []string) error {

	return nil
}
func onAddAttachment(paras []string) error {

	return nil
}
func onDeleteAttachment(paras []string) error {

	return nil
}
func onAddChecklist(paras []string) error {

	return nil
}
func onRemoveChecklist(paras []string) error {

	return nil
}
func onUncompleteChecklist(paras []string) error {

	return nil
}
func onAddChecklistItem(paras []string) error {

	return nil
}
func onCheckedItem(paras []string) error {

	return nil
}
func onRemovedChecklistItem(paras []string) error {

	return nil
}

//Board
func onCreateCustomField(paras []string) error {

	return nil
}

//Lists
func onCreateList(paras []string) error {

	return nil
}
func onArchivedList(paras []string) error {

	return nil
}
func onRemoveList(paras []string) error {

	return nil
}

//Swimlane
func onCreateSwimlane(paras []string) error {

	return nil
}
func onArchivedSwimlane(paras []string) error {

	return nil
}
func onRemoveSwimlane(paras []string) error {

	return nil
}
