package model

import (
	"github.com/google/uuid"
)

func GetGroups(id uuid.UUID) ([]Group, error) {
	var userGroups []UserGroup
	err := db.Where("user_id = ?", id).Find(&userGroups).Error
	if err != nil {
		return []Group{}, err
	}
	var groups []Group
	for _, userGroup := range userGroups {
		var group Group
		err := db.Where("id = ?", userGroup.GroupID).First(&group).Error
		if err != nil {
			return []Group{}, err
		}
		groups = append(groups, group)
	}
	return groups, err
}

func CreateGroup(id uuid.UUID, group GroupRequest) (Group, error) {
	newId := uuid.New()
	newGroup := Group{
		ID:   newId,
		Name: group.Name,
	}
	newUserGroup := UserGroup{
		UserID:  id,
		GroupID: newId,
	}

	err := db.Create(&newUserGroup).Error
	if err != nil {
		return Group{}, err
	}
	err = db.Create(&newGroup).Error
	if err != nil {
		return Group{}, err
	}

	return newGroup, nil
}

func PutGroup(id uuid.UUID, groupID uuid.UUID, group GroupRequest) (Group, error) {
	// このgroupIDかつこのuserIDのuserGroupが存在するかを確認する
	var targetUserGroup UserGroup
	err := db.Where("user_id = ? AND group_id = ?", id, groupID).First(&targetUserGroup).Error
	if err != nil {
		return Group{}, err
	}

	// 存在したら、このgroupIDのgroupを更新する
	var targetGroup Group
	err = db.Where("id = ?", groupID).First(&targetGroup).Error
	if err != nil {
		return Group{}, err
	}
	// 変更を加える
	targetGroup.Name = group.Name
	db.Save(&targetGroup)

	return targetGroup, nil
}

func DeleteGroup(id uuid.UUID, groupID uuid.UUID) error {
	// このgroupIDのuserGroupが存在するかを確認する
	var targetUserGroup UserGroup
	err := db.Where("id = ?", groupID).First(&targetUserGroup).Error
	if err == nil {
		// このuserGroupを削除
		db.Delete(&targetUserGroup)
	} else {
		return err
	}

	// このgroupIDのグループがあるか確認する
	var targetGroup Group
	err = db.Where("id = ?", groupID).First(&targetGroup).Error
	if err == nil {
		// このグループを削除
		db.Delete(&targetGroup)
	}

	// このgroupIDのtodoが存在するかを確認する
	var targetTodos []Todo
	err = db.Where("group_id = ?", groupID).Find(&targetTodos).Error
	if err == nil {
		// このtodoを削除
		db.Delete(&targetTodos)
	}

	return nil
}
