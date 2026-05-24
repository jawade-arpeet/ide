package dto

type CreateWorkspacePayload struct {
	Name string `json:"name" binding:"required"`
}

type AddWorkspaceMemberPayload struct {
	AccountID string `json:"account_id" binding:"required"`
	Role      string `json:"role" binding:"required"`
}
