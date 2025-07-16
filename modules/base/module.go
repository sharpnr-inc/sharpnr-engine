package base

type InitialiseBaseModule struct {
	orgId InitBaseModule
}

func AddBaseModuleInOrg(orgid InitBaseModule) *InitialiseBaseModule {
	return &InitialiseBaseModule{orgId: orgid}
}
