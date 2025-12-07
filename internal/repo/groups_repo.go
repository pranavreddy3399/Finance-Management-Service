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
	GetMembersByGroupID(ctx context.Context, groupID string) ([]*model.GroupMemberDetailEntity, error)
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
	const query = "INSERT INTO `group_members` (`user_id`, `group_id`, `role`, `status`, `joined_at`) " +
		"VALUES (?, ?, ?, ?, NOW())"

	_, err := gr.db.ExecContext(
		ctx,
		query,
		memberData.UserID,
		memberData.GroupID,
		memberData.Role,
		memberData.Status,
	)
	if err != nil {
		return fmt.Errorf("AddGroupMember | insert error: %w", err)
	}
	return nil
}

func (gr *GroupRepo) GetMembersByGroupID(ctx context.Context, groupID string) ([]*model.GroupMemberDetailEntity, error) {
	const query = `
		SELECT
			gm.user_id,
			gm.group_id,
			gm.role,
			gm.status,
			u.name,
			u.email,
			u.phno
		FROM group_members gm
		JOIN user1 u ON gm.user_id = u.id
		WHERE gm.group_id = ? AND gm.status = 'ACTIVE'
	`

	var members []*model.GroupMemberDetailEntity
	if err := gr.db.SelectContext(ctx, &members, query, groupID); err != nil {
		return nil, fmt.Errorf("GetMembersByGroupID | select error: %w", err)
	}
	return members, nil
}
