package interfaces

// Repo definition.
type Repo interface {
	Save(interface{}, interface{})
	SaveSet(interface{}, ...interface{})
	Get(interface{}) string
	GetASet(interface{}) []string
}
