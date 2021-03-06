package generator

// Generator is the interface that wraps the method Generate to create
// a mask for a given id.
type Generator interface {
	Generate(id string) (mask string)
}
