# Answers

## Question 1

Un array representa una secuencia de longitud fija. Su longitud forma parte del tipo, por lo que `[3]int` y `[4]int` son tipos diferentes. El valor cero de un array contiene el valor cero del tipo de sus elementos en todas sus posiciones. Un array se puede comparar con `==` cuando su tipo de elemento es comparable; la comparación revisa todos los elementos y su orden. Asignar un array a otra variable copia su contenido.

Un slice representa una vista de longitud variable sobre un array subyacente. Su longitud no forma parte del tipo: dos valores `[]int` tienen el mismo tipo aunque sus longitudes sean distintas. El valor cero de un slice es `nil`, con longitud y capacidad cero. Los slices no se pueden comparar entre sí con `==`; solamente se puede comparar un slice con `nil`.

Al obtener un slice desde un array no se copian los elementos. El slice comparte el almacenamiento del array, por lo que una escritura mediante cualquiera de los dos valores se puede observar mediante el otro mientras ambos señalen las mismas posiciones.

## Question 2

La longitud de un slice indica cuántos elementos contiene y determina sus índices válidos, desde `0` hasta `len(s)-1`. La capacidad indica cuántos elementos puede abarcar desde su posición inicial hasta el final del array subyacente antes de necesitar otro almacenamiento. El slice conserva una referencia a ese almacenamiento, además de su longitud y capacidad.

- `var s []int` crea un slice `nil` con longitud y capacidad cero.
- `s := []int{}` crea un slice vacío pero no `nil`, también con longitud y capacidad cero.
- Un literal como `[]int{10, 20}` crea dos elementos; su longitud y capacidad son dos.
- `make([]int, length)` crea `length` elementos con sus valores cero y una capacidad igual a la longitud.
- `make([]int, length, capacity)` permite elegir ambas medidas, siempre que la longitud no supere la capacidad. `make([]int, 0, 5)` reserva espacio para crecer, pero todavía no crea índices válidos.

`append` devuelve un nuevo valor de slice con mayor longitud. Si existe capacidad suficiente, puede reutilizar el array subyacente; si no existe, reserva otro array y copia los elementos. Por eso siempre se debe conservar el slice devuelto y no se debe depender de un factor específico de crecimiento de la capacidad.

`copy(destination, source)` copia como máximo el menor valor entre las longitudes de ambos slices. No aumenta la longitud del destino y su capacidad adicional no permite copiar hacia índices que todavía no existen. El destino conserva su propio almacenamiento si fue creado de forma independiente.

## Question 3

Un `string` es una secuencia inmutable de bytes. Normalmente esos bytes contienen texto codificado en UTF-8, pero Go no exige que todo `string` contenga UTF-8 válido.

- `len(text)` cuenta bytes, no caracteres ni runas.
- `text[index]` devuelve un byte de tipo `uint8`, aunque ese byte sea solamente una parte de un punto de código con varios bytes.
- `text[low:high]` selecciona un rango de bytes. Si los límites cortan la codificación de una runa, el resultado puede no ser UTF-8 válido.
- `[]byte(text)` produce una secuencia mutable con los bytes de la cadena.
- `[]rune(text)` decodifica el texto en puntos de código Unicode; cada posición representa una runa completa.
- `range` sobre un `string` entrega el desplazamiento en bytes de cada runa y el valor `rune` decodificado. Los índices no necesariamente avanzan de uno en uno cuando aparecen runas de varios bytes.

Por ejemplo, `世` ocupa tres bytes en UTF-8, pero cuenta como una sola runa.

## Question 4

Las claves de un mapa deben ser comparables con `==`. Se permiten, por ejemplo, booleanos, números, strings, punteros, canales, arrays comparables y structs cuyos campos sean comparables. No se permiten slices, maps ni funciones como claves.

El valor cero de un mapa es `nil`. Un mapa `nil` tiene longitud cero y se puede leer, consultar con comma-ok, recorrer, eliminar o limpiar sin producir un `panic`, pero una asignación como `m[key] = value` falla porque todavía no existe almacenamiento inicializado. Un mapa creado con un literal o con `make` sí permite escrituras.

Leer una clave ausente devuelve el valor cero del tipo del valor. Como ese mismo valor cero podría estar almacenado legítimamente, comma-ok devuelve además un booleano: `value, exists := m[key]`. Así se distingue una clave presente con valor cero de una clave ausente.

El número opcional pasado a `make(map[K]V, sizeHint)` es una estimación inicial de elementos, no una longitud, una capacidad observable ni un límite. El mapa puede crecer cuando sea necesario. Go tampoco garantiza un orden de iteración estable; un reporte determinista debe utilizar otro orden conocido, como un array o un slice de claves.

## Question 5

Un sub-slice conserva una referencia al mismo array subyacente y normalmente también conserva capacidad después de su longitud visible. Si `append` agrega un elemento sin superar esa capacidad, escribe en la siguiente posición del array compartido. Esa posición puede pertenecer a la longitud visible del slice original, por lo que el cambio aparece como una sobrescritura en ese otro valor.

Cuando un `append` supera la capacidad disponible, Go reserva otro array y copia los elementos. A partir de ese momento, las escrituras realizadas en el nuevo almacenamiento ya no afectan al slice original. El crecimiento exacto de la nueva capacidad no está especificado.

Una expresión completa como `source[low:high:max]` permite limitar la capacidad a `max-low`. Si la longitud ya ocupa toda esa capacidad, el siguiente `append` debe usar almacenamiento nuevo. Sin embargo, la expresión no copia los elementos que ya estaban dentro del rango: antes del `append`, ambos slices todavía comparten esas posiciones y una asignación indexada sobre ellas continúa siendo visible desde ambos.

## Question 6

Un `struct` admite `==` solamente cuando todos sus campos son comparables. Los campos `string`, `int`, `bool` y los arrays comparables no causan problemas, pero un campo slice hace que el `struct` completo deje de ser comparable. Que dos slices contengan los mismos elementos no cambia esta regla; su contenido debe compararse de otra forma, por ejemplo convirtiéndolos a arrays cuando la longitud requerida es conocida.

Dos tipos `struct` definidos con nombres diferentes tienen identidades diferentes, aunque sus campos tengan los mismos nombres, tipos y orden. Esa separación evita mezclar accidentalmente valores que representan conceptos distintos. Cuando sus tipos subyacentes son compatibles, se puede cruzar el límite mediante una conversión explícita.

Un `struct` anónimo no introduce otra identidad de tipo con nombre. Si su estructura subyacente coincide con la del tipo nombrado y se cumplen las reglas de asignabilidad, el valor puede asignarse directamente porque al menos uno de los tipos no tiene nombre. Las etiquetas de campos también forman parte de la identidad estructural relevante.

## Question 7

Para un slice:

- `clear(s)` reemplaza con valores cero todos los elementos de su longitud actual. No cambia su longitud, capacidad ni estado `nil`, y conserva el array subyacente. Si la longitud es mayor que cero, sus índices siguen disponibles para escritura inmediata.
- `s = s[:0]` establece la longitud en cero y conserva la capacidad y el almacenamiento. No borra los valores anteriores; pueden volver a observarse al extender el slice dentro de su capacidad. Ya no se puede escribir mediante `s[0]`, pero se puede usar `append` inmediatamente. Un slice no `nil` continúa sin ser `nil`.
- `s = nil` deja longitud y capacidad cero y hace que el slice sea `nil`. La variable deja de referenciar el array anterior, pero no borra los datos que todavía puedan observar otros aliases. No permite una escritura indexada, aunque `append` funciona sobre un slice `nil`.

Para un mapa:

- `clear(m)` elimina todas las entradas y deja longitud cero. Si el mapa estaba inicializado, sigue sin ser `nil` y acepta nuevas asignaciones inmediatamente. `clear` también es seguro sobre un mapa `nil`, aunque ese mapa continúa sin aceptar escrituras.
- `m = nil` elimina la referencia de esa variable al mapa y deja longitud cero. No borra las entradas visibles mediante otras variables que compartan el mismo mapa. Las lecturas siguen siendo seguras, pero una asignación directa produce un `panic` hasta volver a inicializar el mapa.

## Question 8

Usaría un array como `[3]string{"build", "test", "deploy"}` para representar la lista fija de categorías. El array expresa que la cantidad no cambia y también ofrece un orden estable para imprimir los conteos. A partir de él construiría un conjunto `map[string]struct{}` para comprobar pertenencia mediante comma-ok sin recorrer toda la lista.

Modelaría cada registro aceptado con un `struct` nombrado que contenga `Identifier`, `Category`, `Label`, `ByteCount` y `RuneCount`. Guardaría los registros aceptados en un slice creado con longitud cero y capacidad reservada para la cantidad de entradas. `append` conservaría su orden de aceptación sin crear registros cero iniciales.

Usaría otro `map[string]struct{}` como conjunto de identificadores aceptados. Primero comprobaría la categoría y después el identificador; solo insertaría el identificador en el conjunto cuando el registro haya pasado ambas validaciones. Así un registro rechazado no bloquea por error un identificador posterior.

Mantendría los conteos en `map[string]int`. Una consulta con comma-ok distinguiría una categoría ausente de una categoría almacenada con conteo cero. Para obtener un reporte determinista recorrería el array de categorías, no el mapa.

Los bytes de cada etiqueta se obtendrían con `len(label)` o `len([]byte(label))`, mientras que las runas se contarían con `len([]rune(label))`. Acumularía esos valores a partir de cada registro aceptado, sin escribir totales fijos. Las conversiones a `[]byte` y `[]rune` producen representaciones diferentes del mismo texto: unidades UTF-8 frente a puntos de código Unicode.

Antes del reset crearía un snapshot independiente con un slice de la misma longitud y `copy`. Como el `struct` propuesto solo contiene strings e ints, la copia de sus campos es suficiente: los strings son inmutables y no hay slices o maps internos que requieran otra copia. El slice de trabajo y el snapshot tendrían arrays subyacentes distintos, por lo que `clear(working)` pondría en cero únicamente los registros de trabajo y conservaría su longitud; el snapshot mantendría los datos originales.

Los slices no se compararían directamente con `==`. Los registros individuales sí serían comparables porque todos sus campos lo son; para comparar catálogos ordenados habría que revisar su longitud y cada registro, o convertirlos a arrays cuando la longitud exacta sea parte conocida del problema. También evitaría crear sub-slices innecesarios del catálogo porque compartirían almacenamiento. Si necesitara archivar un mapa, tendría que crear otro mapa y copiar sus entradas, ya que asignar un mapa a otra variable mantendría almacenamiento compartido.

No hubo preguntas dudosas: todas las respuestas cubren los puntos solicitados.
