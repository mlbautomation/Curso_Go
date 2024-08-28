package item

// Abstracción a través de estructuras
type terno struct {
	id     int
	talla  string
	color  string
	precio float64
}

// Encapsulación a través de indentificador no exportado

// Funcion constructora
func NewTerno(id *int, talla string, color string, precio float64) (terno terno) {
	*id++
	terno.id = *id
	terno.talla = talla
	terno.color = color
	terno.precio = precio
	return terno
}

// Metodos SETTER
func (t *terno) SetId(id int)             { t.id = id }
func (t *terno) SetTalla(talla string)    { t.talla = talla }
func (t *terno) SetColor(color string)    { t.color = color }
func (t *terno) SetPrecio(precio float64) { t.precio = precio }

// Metodos GETTER
func (t *terno) Id() int         { return t.id }
func (t *terno) Talla() string   { return t.talla }
func (t *terno) Color() string   { return t.color }
func (t *terno) Precio() float64 { return t.precio }

// Metodos de la interfaz
func (t *terno) GetId() int             { return t.id }
func (t *terno) GetDescripcion() string { return "TER" + t.Talla() + t.Color() }
func (t *terno) GetPrecio() float64     { return t.Precio() }
