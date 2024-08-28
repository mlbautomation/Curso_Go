package client

import (
	"fmt"
	"strings"
	"time"
)

// Model of invoiceheader
type Model struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Models []*Model

func NewModelClient(name string) *Model {
	return &Model{
		Name:      name,
		CreatedAt: time.Now(),
	}
}

// Función para que la información se vea mejor
func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %10s | %10s",
		m.ID, m.Name,
		m.CreatedAt.Format("2006-01-02"),
		m.UpdatedAt.Format("2006-01-02"))
}

// Función para que la información se vea mejor
func (m Models) String() string {
	builder := strings.Builder{} //Usamos el paquete strings, tipo Builder
	builder.WriteString(fmt.Sprintf("%-02s | %-20s | %-10s | %-10s\n",
		"id", "name", "created_at", "updated_at"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}
