# E-commerce Manager · Go

Ejercicio de modelado de productos, clientes y pedidos con **Go**. La demostración crea dos pedidos en memoria, agrega productos y muestra el detalle y el total en la consola.

## Ejecutar

Con Go instalado, desde la carpeta del repositorio:

```powershell
go run Proyecto.go
```

No requiere librerías externas, una API ni una base de datos.

## Conceptos que practica

- Estructuras `Producto`, `Cliente` y `Pedido`.
- Funciones constructoras y métodos.
- Slices para agrupar productos.
- Receptores por puntero para modificar un pedido.
- Recorrido de productos y salida formateada.

## Resultado de la demostración

Los datos del programa generan el pedido **1001** con un total de **945.50** y el pedido **1002** con un total de **325.00**. Los importes sirven para comprobar la suma de los productos del ejemplo.

## Alcance y siguientes pasos

Es un programa de consola con datos predefinidos: no es una tienda web ni procesa pagos. El stock se representa como un dato del producto, pero agregarlo a un pedido no descuenta existencias. Para ampliar el ejercicio se pueden añadir cantidades, validaciones y persistencia.

## Autoría

El código original identifica a **Josue y Julian**. Se conserva ese reconocimiento.
