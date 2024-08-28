package item

// Abstracción a través de estructuras
type camisa struct {
	id     int
	talla  string
	color  string
	precio float64
}

// Encapsulación a través de indentificador no exportado

// Funcion constructora
func NewCamisa(id *int, talla string, color string, precio float64) (camisa camisa) {
	*id++
	camisa.id = *id
	camisa.talla = talla
	camisa.color = color
	camisa.precio = precio
	return camisa
}

// Metodos SETTER
func (c *camisa) SetId(id int)             { c.id = id }
func (c *camisa) SetTalla(talla string)    { c.talla = talla }
func (c *camisa) SetColor(color string)    { c.color = color }
func (c *camisa) SetPrecio(precio float64) { c.precio = precio }

// Metodos GETTER
func (c *camisa) Id() int         { return c.id }
func (c *camisa) Talla() string   { return c.talla }
func (c *camisa) Color() string   { return c.color }
func (c *camisa) Precio() float64 { return c.precio }

// Metodos de la interfaz
func (c *camisa) GetId() int             { return c.id }
func (c *camisa) GetDescripcion() string { return "CAM" + c.Talla() + c.Color() }
func (c *camisa) GetPrecio() float64     { return c.Precio() }
