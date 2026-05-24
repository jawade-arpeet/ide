CREATE TYPE enum_workspace_member_role AS ENUM ('owner', 'member');

CREATE TABLE IF NOT EXISTS workspace_members (
    workspace_id UUID NOT NULL REFERENCES workspaces(id),
    account_id UUID NOT NULL REFERENCES accounts(id),
    role enum_workspace_member_role NOT NULL DEFAULT 'member',
    joined_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (workspace_id, account_id)
);
