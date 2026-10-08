# `goto` en Go

`goto` transfiere el control a un label dentro de la misma función. Go conserva esta funcionalidad, pero restringe los saltos que violarían el scope, por lo que debe seguir siendo una herramienta excepcional y no un flujo de control habitual.

## 1. Qué debo entender

- La sintaxis es `goto label`, con `label:` marcando el destino.
- El label de destino debe estar en la misma función.
- Un salto no puede omitir declaraciones cuyas variables estarían en scope en el destino.
- Un salto no puede entrar desde fuera en un bloque interior.
- `break` y `continue` etiquetados suelen ser más claros para bucles anidados.

## 2. Conceptos clave

| Salto | ¿Permitido? |
|---|---:|
| A un label válido de la misma función | A veces |
| A través de código sin violar el scope | Posiblemente |
| Sobre una declaración hacia su scope | No |
| Desde fuera hacia un bloque interior | No |
| Hacia otra función | No |

## 3. Cómo funciona en Go

```go
value := readValue()
if value < 0 {
    goto done
}

process(value)

done:
fmt.Println("finished")
```

La siguiente forma es ilegal porque el salto omite una declaración que estaría en scope en el label:

```go
goto done
result := compute()
done:
fmt.Println(result)
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** un label no elimina las reglas normales de scope; solo da nombre a un posible destino del flujo de control.

Saltar desde fuera de un `if` u otro bloque hacia un label interior es inválido. Un `goto` tampoco puede cruzar los límites de una función.

Un uso excepcional puede hacer converger varios caminos complejos en una sección común de limpieza o procesamiento final, evitando flags artificiales o duplicación considerable. Considera primero `return`, extraer una función o usar `break`/`continue` etiquetados.

> ⚠️ **Corrección importante:** “evitar `goto`” es una recomendación de diseño, no una prohibición del lenguaje. El compilador acepta saltos legales y la standard library contiene casos poco frecuentes donde un salto restringido simplifica el flujo.

## 5. Chuleta rápida

- Sintaxis → `goto label`
- Destino → `label:`
- El label debe estar en la misma función
- No puede saltar una declaración hacia su scope
- No puede entrar en un bloque interior
- No puede saltar hacia otra función
- Para bucles anidados, prefiere `break`/`continue` etiquetados
- Considera primero `return` o extraer una función
- Uso poco frecuente → converger en lógica final común
- Que sea legal no significa automáticamente que sea legible

### Comprueba que realmente lo sabes

1. ¿Por qué es ilegal saltar una declaración cuando la variable estaría en scope en el destino?
2. ¿Cuándo resulta más claro un `break` etiquetado que `goto`?
3. ¿Qué características podrían justificar un salto poco frecuente hacia una lógica final común?
