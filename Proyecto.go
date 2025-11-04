// Josue and JUlian
package main

import "fmt"

// --- STRUCTS BÁSICOS ---

// Producto representa un ítem del catálogo.
type Producto struct {
	ID          int
	Nombre      string
	Descripcion string
	Precio      float64
	Stock       int
}

// Constructor (función para crear productos)
func NuevoProducto(id int, nombre, descripcion string, precio float64, stock int) Producto {
	return Producto{
		ID:          id,
		Nombre:      nombre,
		Descripcion: descripcion,
		Precio:      precio,
		Stock:       stock,
	}
}

// Getter: obtiene el nombre del producto
func (p Producto) GetNombre() string {
	return p.Nombre
}

// Getter: obtiene el precio
func (p Producto) GetPrecio() float64 {
	return p.Precio
}

// --- CLIENTE ---

type Cliente struct {
	ID     int
	Nombre string
	Email  string
}

// Constructor
func NuevoCliente(id int, nombre, email string) Cliente {
	return Cliente{
		ID:     id,
		Nombre: nombre,
		Email:  email,
	}
}

// Getter
func (c Cliente) GetNombre() string {
	return c.Nombre
}

// --- PEDIDO ---

type Pedido struct {
	ID        int
	Cliente   Cliente
	Productos []Producto // slice de productos
	Total     float64
}

// Agregar producto al pedido
func (p *Pedido) AgregarProducto(prod Producto) {
	p.Productos = append(p.Productos, prod)
	p.Total += prod.Precio
}

// Mostrar resumen del pedido
func (p Pedido) MostrarResumen() {
	fmt.Println(" Pedido N°", p.ID)
	fmt.Println("Cliente:", p.Cliente.Nombre)
	fmt.Println("Productos:")
	for _, prod := range p.Productos {
		fmt.Printf("- %s ($%.2f)\n", prod.Nombre, prod.Precio)
	}
	fmt.Printf("Total: $%.2f\n", p.Total)
}

// --- PROGRAMA PRINCIPAL ---

func main() {
	// Crear productos
	p1 := NuevoProducto(1, "Laptop", "16GB RAM", 850.50, 5)
	p2 := NuevoProducto(2, "Mouse", "Inalámbrico", 25.00, 10)
	p3 := NuevoProducto(3, "Teclado", "Mecánico", 70.00, 7)
	p4 := NuevoProducto(4, "Monitor", "27 pulgadas", 300.00, 3)

	// Crear cliente
	c := NuevoCliente(1, "Josue", "josue@example.com")
	c2 := NuevoCliente(2, "Julian", "julianm@gmail.com")

	// Mostrar nombres de clientes
	fmt.Println("Cliente 1:", c.GetNombre())
	fmt.Println("Cliente 2:", c2.GetNombre())

	// Crear pedido
	pedido := Pedido{ID: 1001, Cliente: c}
	pedido.AgregarProducto(p1)
	pedido.AgregarProducto(p2)
	pedido.AgregarProducto(p3)

	pedido2 := Pedido{ID: 1002, Cliente: c2}
	pedido2.AgregarProducto(p2)
	pedido2.AgregarProducto(p4)
	// Mostrar pedido
	pedido.MostrarResumen()
	pedido2.MostrarResumen()
}
