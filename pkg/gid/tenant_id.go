package gid

type TenantID [8]byte

var (
	NilTenant = TenantID{}
)
