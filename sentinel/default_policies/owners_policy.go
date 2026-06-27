package sentinelDefaultpolicies

type SentinalDefaultPolices string

const (
	OwnerPolcy SentinalDefaultPolices = `permit (
  principal in Role::"owner",
  action,
  resource
) when { resource.orgId == principal.orgId };`
)
