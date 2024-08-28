package cliente

// Abstracción a través de estructuras
type Cliente struct {
	id              int
	primerNombre    string
	segundoNombre   string
	primerApellido  string
	segundoApellido string
}

// Encapsulación a través de indentificador no exportado

// Funcion constructora
func NewCliente(id *int, primerNombre string, segundoNombre string, primerApellido string, segundoApellido string) (cliente Cliente) {
	*id++
	cliente.id = *id
	cliente.primerNombre = primerNombre
	cliente.segundoNombre = segundoNombre
	cliente.primerApellido = primerApellido
	cliente.segundoApellido = segundoApellido
	return cliente
}

// Metodos SETTER
func (c *Cliente) SetId(id int)                              { c.id = id }
func (c *Cliente) SetPrimerNombre(primerNombre string)       { c.primerNombre = primerNombre }
func (c *Cliente) SetSegundoNombre(segundoNombre string)     { c.segundoNombre = segundoNombre }
func (c *Cliente) SetPrimerApellido(primerApellido string)   { c.primerApellido = primerApellido }
func (c *Cliente) SetSegundoApellido(segundoApellido string) { c.segundoApellido = segundoApellido }

// Metodos GETTER
func (c *Cliente) Id() int                 { return c.id }
func (c *Cliente) PrimerNombre() string    { return c.primerNombre }
func (c *Cliente) SegundoNombre() string   { return c.segundoNombre }
func (c *Cliente) PrimerApellido() string  { return c.primerApellido }
func (c *Cliente) SegundoApellido() string { return c.segundoApellido }
