package interfaces

// Repo definition.
type Repo interface {
	Save(interface{}, interface{})
	SaveSet(interface{}, ...interface{})
	GetASet(interface{}) []string
}
