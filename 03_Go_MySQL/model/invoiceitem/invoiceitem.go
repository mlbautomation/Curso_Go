package invoiceitem

import (
	"fmt"
	"strings"
	"time"
)

// Model of invoiceitem
type Model struct {
	ID        uint
	ClienteID uint
	ProductID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Models []*Model

func NewModelInvoiceItem(clienteId uint, productId uint) *Model {
	return &Model{
		ClienteID: clienteId,
		ProductID: productId,
		CreatedAt: time.Now(),
	}
}

// Función para que la información se vea mejor
func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-10d | %-10d | %10s | %10s",
		m.ID, m.ClienteID, m.ProductID,
		m.CreatedAt.Format("2006-01-02"),
		m.UpdatedAt.Format("2006-01-02"))
}

// Función para que la información se vea mejor
func (m Models) String() string {
	builder := strings.Builder{} //Usamos el paquete strings, tipo Builder
	builder.WriteString(fmt.Sprintf("%-02s | %-10s | %-10s | %-10s | %-10s\n",
		"id", "cliente_id", "product_id", "created_at", "updated_at"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}
