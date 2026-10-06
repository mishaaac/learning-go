# Comparación y conversión de structs en Go

La igualdad de structs depende de la comparabilidad de todos sus campos, mientras que la conversión depende de la relación entre los tipos struct de origen y destino. Los structs con nombre y anónimos siguen reglas de asignabilidad diferentes aunque sus campos visibles parezcan iguales.

## 1. Qué debo entender

- Un struct admite `==` y `!=` solo cuando todos sus campos son comparables.
- La igualdad compara los campos correspondientes.
- Los tipos struct con nombre definidos por separado siguen siendo tipos diferentes.
- Los tipos struct con nombre compatibles pueden necesitar una conversión explícita.
- Un struct con nombre y uno anónimo pueden ser asignables directamente cuando sus tipos subyacentes son idénticos y se cumplen las reglas de asignabilidad.

## 2. Conceptos clave

| Situación | Asignación/comparación directa |
|---|---|
| Mismo tipo con nombre y comparable | Permitida |
| El struct contiene slice, map o función | Igualdad no permitida |
| Dos tipos con nombre distintos y estructura coincidente | Requiere conversión explícita |
| Tipos con nombre y anónimo con estructura subyacente idéntica | Puede permitirse la asignación directa |
| Semántica de igualdad personalizada | Escribe una función; `==` no puede redefinirse |

## 3. Cómo funciona en Go

```go
type FirstPerson struct {
    Name string
    Age  int
}

type SecondPerson struct {
    Name string
    Age  int
}

first := FirstPerson{Name: "Bob", Age: 50}
second := SecondPerson(first) // explicit conversion

var anonymous struct {
    Name string
    Age  int
}
anonymous = first

fmt.Println(first == anonymous) // true
fmt.Println(second)
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** dos structs con nombre y campos coincidentes siguen siendo tipos con nombre distintos; la similitud estructural por sí sola no permite asignación ni comparación directa.

Los nombres, el orden, los tipos y el estado embedded de los campos determinan la identidad estructural; los tags también importan para la identidad de tipo y la asignabilidad directa. Para una conversión explícita de structs, Go permite tipos struct subyacentes idénticos en los demás aspectos ignorando los tags.

> ⚠️ **Corrección importante:** los channels son comparables, por lo que un campo channel no vuelve por sí solo no comparable a un struct. Los slices, maps y funciones son los tipos de campo habituales que impiden comparar structs.

Go no ofrece un método que reemplace `==`. Si la igualdad debe ignorar o transformar campos, expresa esa regla en una función separada.

## 5. Chuleta rápida

- Struct comparable → todos los campos son comparables
- Igualdad → compara campos correspondientes
- Campo slice/map/función → struct no comparable
- Campo channel → sigue siendo comparable
- Tipos con nombre distintos siguen siendo distintos
- Structs con nombre coincidentes → puede ser posible la conversión explícita
- La conversión necesita estructura y tipos de campos coincidentes
- Los struct tags se ignoran en la conversión explícita
- Named ↔ anonymous compatible puede asignarse directamente
- `==` no puede personalizarse
- Semántica propia → escribe una función de comparación

### Comprueba que realmente lo sabes

1. ¿Por qué un campo slice impide comparar structs mientras que un campo channel no?
2. ¿Por qué se necesita una conversión explícita entre dos structs con nombres distintos pero estructura coincidente?
3. ¿Cómo afectan los struct tags a la identidad de tipo frente a la conversión explícita?
