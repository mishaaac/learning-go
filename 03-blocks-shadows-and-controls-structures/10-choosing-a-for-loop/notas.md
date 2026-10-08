# Elegir el `for` adecuado en Go

Go usa una sola keyword para varios patrones de iteración. La forma más clara es la que expresa directamente si el código recorre contenido, sigue límites explícitos, repite según una condición o sale desde el cuerpo.

## 1. Qué debo entender

- Prefiere `range` al recorrer todos los elementos de un valor compatible.
- Usa un `for` de tres partes cuando importen los límites de inicio, fin y actualización.
- Usa `for condition` cuando la repetición siga una condición cambiante.
- Usa `for {}` cuando las decisiones de salida ocurran naturalmente dentro del cuerpo.
- Para strings, prefiere `range` al procesar Unicode code points en vez de bytes.

## 2. Conceptos clave

| Necesidad | Forma natural |
|---|---|
| Recorrer todos los elementos de un slice/map | `for ... := range value` |
| Decodificar runes de un string | `for offset, r := range text` |
| Visitar un intervalo de índices concreto | `for i := start; i < end; i++` |
| Repetir mientras se cumpla una condición | `for condition` |
| Ejecutar primero y decidir la salida dentro | `for { ... break }` |

## 3. Cómo funciona en Go

Recorrer todo:

```go
for index, value := range values {
    fmt.Println(index, value)
}
```

Recorrer solo el centro de un slice:

```go
for index := 1; index < len(values)-1; index++ {
    fmt.Println(index, values[index])
}
```

Decodificar texto correctamente por rune:

```go
for _, r := range text {
    fmt.Printf("%c", r)
}
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** indexar un string itera bytes; aplicar range sobre un string decodifica runes UTF-8.

Usar `range` junto con varias comprobaciones `continue` y `break` para expresar límites numéricos sencillos puede ocultar el intervalo buscado. A la inversa, indexar manualmente todo un slice añade trabajo de control que `range` ya expresa.

Un bucle infinito normalmente debe mostrar un camino deliberado mediante `break`, `return`, cancelación u otra salida. Si ninguna forma se lee con claridad, el cuerpo del bucle puede estar realizando demasiadas tareas y beneficiarse de una refactorización.

## 5. Chuleta rápida

- Todos los elementos → `range`
- Runes de un string → `range`
- Inicio/fin/actualización explícitos → `for` de tres partes
- Repetición por condición → `for condition`
- Salida decidida dentro del cuerpo → `for {}`
- Las posiciones de un slice son índices de elementos
- Los índices de un string son posiciones de bytes
- Evita indexar manualmente cuando `range` expresa la intención
- Evita complicar `range` para límites numéricos sencillos
- Los bucles infinitos necesitan un diseño deliberado de salida

### Comprueba que realmente lo sabes

1. ¿Por qué un bucle de tres partes es más claro para visitar solo el centro de un slice?
2. ¿Por qué el recorrido de strings Unicode debería usar normalmente `range`?
3. ¿Qué señal de diseño ofrece un bucle infinito sin una salida evidente?
