package RoleType

type Role string

const (
	Authenticated       Role = "Authenticated"
	Administrator       Role = "Administrator"
	DepartmentHead      Role = "DepartmentHead"
	Executor            Role = "Executor"
	ResponsibleExecutor Role = "ResponsibleExecutor"
)
