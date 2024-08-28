package product

import (
	"fmt"
	"strings"
	"time"
)

// Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Models []*Model

func NewModelProduct(name string, price uint) *Model {
	return &Model{
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
}

// Función para que la información se vea mejor
func (m *Model) String() string {
	return fmt.Sprintf("%-02d | %-20s | %-20s | %-10d | %-10s | %-10s",
		m.ID, m.Name, m.Observations, m.Price,
		m.CreatedAt.Format("2006-01-02"),
		m.UpdatedAt.Format("2006-01-02"))
}

// Función para que la información se vea mejor
func (m Models) String() string {
	builder := strings.Builder{} //Usamos el paquete strings, tipo Builder
	builder.WriteString(fmt.Sprintf("%02s | %-20s | %-20s | %-10s | %-10s | %-10s\n",
		"id", "name", "observations", "price", "created_at", "updated_at"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}
