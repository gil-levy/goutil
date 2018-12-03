package tests

type Describe struct {
	description string
}

func (d *Describe) It(s string) string {
	return d.description + ":\nit should " + s
}
