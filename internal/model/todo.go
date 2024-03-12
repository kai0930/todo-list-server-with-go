package model

import (
	"github.com/google/uuid"
	"time"
)

func GetTodos(id uuid.UUID, groupId uuid.UUID) ([]Todo, error) {
	// このidかつこのgroupIdのuserGroupを取得する
	var userGroup UserGroup
	err := db.Where("user_id = ? AND group_id = ?", id, groupId).First(&userGroup).Error
	if err != nil {
		return []Todo{}, err
	}
	// 存在したら、このgroupIdのtodoを取得する
	var todos []Todo
	err = db.Where("group_id = ?", groupId).Find(&todos).Error
	if err != nil {
		return []Todo{}, err
	}
	return todos, nil
}

func CreateTodo(id uuid.UUID, todo CreateTodoRequest) (Todo, error) {
	// このidかつこのtodoのgroupに所属しているかを確認する
	var userGroup UserGroup
	err := db.Where("user_id = ? AND group_id = ?", id, todo.GroupID).First(&userGroup).Error
	if err != nil {
		return Todo{}, err
	}
	// 存在したら、新しいtodoを作成する
	newId := uuid.New()
	newTodo := Todo{
		ID:      newId,
		GroupID: todo.GroupID,
		Title: func() string {
			if todo.Title != nil {
				return *todo.Title
			} else {
				return ""
			}
		}(),
		Description: func() string {
			if todo.Description != nil {
				return *todo.Description
			} else {
				return ""
			}
		}(),
		DueDate: func() time.Time {
			if todo.DueDate != nil {
				return *todo.DueDate
			} else {
				return time.Time{}
			}
		}(),
		IsCompleted: false,
	}
	err = db.Create(&newTodo).Error
	if err != nil {
		return Todo{}, err
	}
	return newTodo, nil
}

func PutTodo(id uuid.UUID, todoID uuid.UUID, todo PutTodoRequest) (Todo, error) {
	// このtodoIDのtodoを取得する
	var targetTodo Todo
	err := db.Where("id = ? AND group_id = ?", todoID, todo.GroupID).First(&targetTodo).Error
	if err != nil {
		return Todo{}, err
	}
	// このユーザーがこのtodoのgroupに所属しているかを確認する
	var userGroup UserGroup
	err = db.Where("user_id = ? AND group_id = ?", id, todo.GroupID).First(&userGroup).Error
	if err != nil {
		return Todo{}, err
	}
	// 存在したら、このtodoを更新する
	targetTodo.Title = func() string {
		if todo.Title != nil {
			return *todo.Title
		} else {
			return targetTodo.Title
		}
	}()
	targetTodo.Description = func() string {
		if todo.Description != nil {
			return *todo.Description
		} else {
			return targetTodo.Description
		}
	}()
	targetTodo.DueDate = func() time.Time {
		if todo.DueDate != nil {
			return *todo.DueDate
		} else {
			return targetTodo.DueDate
		}
	}()
	targetTodo.IsCompleted = func() bool {
		if todo.IsCompleted != nil {
			return *todo.IsCompleted
		} else {
			return targetTodo.IsCompleted
		}
	}()
	err = db.Save(&targetTodo).Error
	if err != nil {
		return Todo{}, err
	}
	return targetTodo, nil
}

func DeleteTodo(id uuid.UUID, todoID uuid.UUID) error {
	// このidかつこのtodoIDのtodoを取得する
	var targetTodo Todo
	err := db.Where("id = ?", todoID).First(&targetTodo).Error
	if err != nil {
		return err
	}
	// このユーザーがこのtodoのgroupに所属しているかを確認する
	var userGroup UserGroup
	err = db.Where("user_id = ? AND group_id = ?", id, targetTodo.GroupID).First(&userGroup).Error
	if err != nil {
		return err
	}
	// 存在したら、このtodoを削除する
	err = db.Delete(&targetTodo).Error
	if err != nil {
		return err
	}
	return nil
}
