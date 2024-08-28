package factura

import (
	"fmt"

	"github.com/mlbautomation/Curso_GO/02_Go_P00/00_proyecto/cliente"
	"github.com/mlbautomation/Curso_GO/02_Go_P00/00_proyecto/item"
)

type Factura struct {
	cliente cliente.Cliente
	items   item.Items
	total   float64
}

func New(cliente cliente.Cliente, items item.Items) (factura Factura) {
	var total float64

	for _, v := range items {
		total += v.GetPrecio()
	}

	factura.cliente = cliente
	factura.items = items
	factura.total = total
	return factura
}

// Metodos SETTER
func (f *Factura) SetCliente(cliente cliente.Cliente) { f.cliente = cliente }
func (f *Factura) SetItems(items item.Items)          { f.items = items }

// Metodos GETTER
func (f *Factura) Cliente() cliente.Cliente { return f.cliente }
func (f *Factura) Items() item.Items        { return f.items }

// Metodos
func (f *Factura) Imprimir() {
	fmt.Println("Item Cliente:", f.cliente.Id())
	fmt.Println("Nombre Cliente:", f.cliente.PrimerNombre(), f.cliente.SegundoNombre(), f.cliente.PrimerApellido(), f.cliente.SegundoApellido())
	for i := range f.items {
		fmt.Println("Item Factura:", i+1, "Item Producto:", f.items[i].GetId(), "Descripción:", f.items[i].GetDescripcion(), "Precio:", f.items[i].GetPrecio())
	}
	fmt.Println("Total Factura:", f.total)
}

func Facturar() {

	var N_Producto int = 0
	var N_Cliente int = 0

	Color_Camisas := []string{"Blanco", "Negro", "Celeste", "Rosado", "Azul"}
	Tallas_Camisas := []string{"XXL", "XL", "L", "M", "S"}

	cliente_01 := cliente.NewCliente(&N_Cliente, "Marlon", "Moises", "Ly", "Bellido")
	cliente_02 := cliente.NewCliente(&N_Cliente, "Romina", "", "Sanchez", "Arrospide")

	item_01 := item.NewCamisa(&N_Producto, Color_Camisas[4], Tallas_Camisas[1], 40.00)
	item_02 := item.NewTerno(&N_Producto, Color_Camisas[1], Tallas_Camisas[0], 100.00)
	item_03 := item.NewTerno(&N_Producto, Color_Camisas[2], Tallas_Camisas[4], 120.00)
	item_04 := item.NewCamisa(&N_Producto, Color_Camisas[3], Tallas_Camisas[2], 80.00)

	//factura_01 := factura.New(cliente_01, []item.Item{&item_01, &item_02})
	factura_01 := New(cliente_01, []item.Item{&item_01, &item_02})
	//factura_02 := factura.New(cliente_02, []item.Item{&item_01, &item_03, &item_04})
	factura_02 := New(cliente_02, []item.Item{&item_01, &item_03, &item_04})

	fmt.Println("*****************************************")
	factura_01.Imprimir()
	fmt.Println("*****************************************")
	factura_02.Imprimir()
	fmt.Println("*****************************************")

}
