package repo

import (
	"context"
	"fms/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type GroupRepoInt interface {
	CreateGroup(ctx context.Context, groupData *model.GroupEntity) (string, error)
	AddGroupMember(ctx context.Context, memberData *model.GroupMemberEntity) error
}

type GroupRepo struct {
	db *sqlx.DB
}

func NewGroupRepo(db *sqlx.DB) GroupRepoInt {
	return &GroupRepo{
		db: db,
	}
}

func (gr *GroupRepo) CreateGroup(ctx context.Context, groupData *model.GroupEntity) (string, error) {
	// Use double-quoted Go string so we can include MySQL backticks to avoid reserved-word issues.
	const query = "INSERT INTO `groups` (`group_id`, `group_name`, `description`, `created_by`, `created_at`, `updated_at`) " +
		"VALUES (?, ?, ?, ?, NOW(), NOW())"

	_, err := gr.db.ExecContext(
		ctx,
		query,
		groupData.GroupID,
		groupData.GroupName,
		groupData.Description,
		groupData.CreatedBy,
	)
	if err != nil {
		return "", fmt.Errorf("CreateGroup | insert error: %w", err)
	}

	return groupData.GroupID, nil
}

func (gr *GroupRepo) AddGroupMember(ctx context.Context, memberData *model.GroupMemberEntity) error {
	const query = "INSERT INTO `group_members` (`group_member_id`, `group_id`, `role`, `status`, `joined_at`) " +
		"VALUES (?, ?, ?, ?, NOW())"

	_, err := gr.db.ExecContext(
		ctx,
		query,
		memberData.GroupMemberID,
		memberData.GroupID,
		memberData.Role,
		memberData.Status,
	)
	if err != nil {
		return fmt.Errorf("AddGroupMember | insert error: %w", err)
	}
	return nil
}
