permit (
  // 1. Change the namespace to Org
  principal in Org::Role::"super_admin_role_id",
  action,
  resource
)
// 2. Change the namespace here too
when { resource in Org::Organization::"org_uuid" };
