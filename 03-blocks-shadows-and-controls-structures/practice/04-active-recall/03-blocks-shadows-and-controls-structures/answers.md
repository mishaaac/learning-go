# Answers

## Question 1

El bloque universo contiene los identificadores predeclarados de Go, como `int`, `string`, `true`, `nil`, `len` y `append`. Están disponibles en todos los archivos, salvo que una declaración más cercana use el mismo nombre y los oculte.

El bloque de paquete contiene las declaraciones de nivel superior de todos los archivos que pertenecen al mismo paquete: tipos, variables, constantes y funciones. Esos nombres pueden utilizarse desde los archivos del paquete según sus reglas de exportación hacia otros paquetes.

Cada archivo tiene además un bloque de archivo. Allí se encuentran principalmente los nombres introducidos por `import`; un paquete importado en un archivo no queda automáticamente disponible en los demás archivos del mismo paquete.

El bloque de una función contiene sus parámetros, resultados nombrados y variables locales declaradas en su cuerpo. Solo pueden usarse dentro de esa función y a partir del punto donde comienza su alcance.

Las estructuras como `if`, `switch`, `for` y cada bloque escrito con llaves introducen bloques anidados. Las variables declaradas allí existen únicamente dentro de ese bloque y sus bloques internos, desde el comienzo de su alcance hasta el cierre correspondiente.

## Question 2

Una declaración corta con `:=` puede reutilizar un nombre solamente cuando ese identificador ya fue declarado en el mismo bloque, tiene un tipo compatible con el nuevo valor y la declaración corta introduce al menos otro nombre nuevo que no sea `_`. En ese caso el nombre existente recibe una asignación y los nombres nuevos se declaran.

Si el nombre visible pertenece a un bloque exterior, `:=` dentro del bloque actual declara un identificador nuevo. Ese identificador hace shadowing: mientras el bloque interno está activo, las referencias al nombre resuelven a la declaración interna. La variable exterior sigue existiendo, pero queda oculta y no cambia por las asignaciones a la variable interna. Al terminar el bloque, la declaración interna sale de alcance y el nombre vuelve a referirse a la variable exterior.

Para modificar deliberadamente la variable exterior desde un bloque interno se usa `=`, siempre que no se necesite declarar otro identificador en esa misma instrucción.

## Question 3

Un valor declarado en la instrucción inicial de un `if` está en alcance dentro de la condición y en todas sus ramas, incluidos `else if` y `else`, pero deja de existir al terminar la instrucción completa. En un `switch`, el valor de la instrucción inicial está disponible en la expresión del switch y en todos sus `case`, hasta que termina el switch.

Una variable declarada dentro de una rama de `if` solo pertenece al bloque de esa rama. De forma similar, una declaración dentro de un `case` no está disponible fuera del bloque correspondiente.

Un `if` evalúa sus condiciones en orden y ejecuta únicamente la primera rama verdadera; si ninguna coincide, puede ejecutar `else`. Un `switch` examina los casos en orden y ejecuta el primero que coincide, o `default` si ninguno coincide. Un switch sin expresión trata cada condición de sus casos como un booleano. Go no continúa automáticamente al siguiente caso; `fallthrough` tendría que solicitarse de manera explícita.

## Question 4

Al recorrer un array o slice con dos variables, `for-range` produce primero el índice y después una copia del elemento ubicado allí. Los índices se visitan en orden ascendente desde cero. Modificar la variable que recibe la copia no modifica por sí mismo el elemento original; para hacerlo se utiliza el índice.

Al recorrer un mapa, `for-range` produce una clave y una copia de su valor. Go no garantiza el orden de recorrido, por lo que no debe usarse para generar un reporte ordenado. La ausencia de orden también significa que ejecuciones distintas pueden visitar las mismas entradas en órdenes diferentes.

Al recorrer un `string`, `for-range` decodifica UTF-8. La primera variable es el desplazamiento en bytes donde comienza la runa y la segunda es su valor `rune`. El desplazamiento puede saltar varios bytes y no debe confundirse con una posición de caracteres. Para una codificación inválida se produce `utf8.RuneError` y el recorrido avanza según las reglas de decodificación.

En todos los casos se puede usar una sola variable para recibir únicamente el primer valor, o `_` para ignorar expresamente uno de los valores producidos.

## Question 5

Un `break` sin etiqueta termina la instrucción `for`, `switch` o `select` más interna que lo contiene. Si aparece dentro de un `switch` que está anidado en un `for`, la instrucción más interna es el switch, por lo que el loop continúa con su flujo normal.

Una etiqueta permite indicar una instrucción envolvente concreta. Si el loop tiene la etiqueta `outer`, `break outer` termina ese loop completo aun cuando la instrucción se encuentre dentro de un switch o de otro loop anidado. De manera equivalente, `continue outer` inicia la siguiente iteración del loop etiquetado.

## Question 6

Una etiqueta de `goto` pertenece a una función, pero el salto no puede entrar desde fuera a un bloque más interno. Hacerlo permitiría comenzar la ejecución en una región sin haber pasado por la entrada normal de ese bloque y sin respetar el alcance previsto de sus declaraciones.

Go tampoco permite que un `goto` salte por encima de una declaración si la variable declarada estaría en alcance en la etiqueta de destino. De otro modo, el programa podría usar una variable cuyo inicializador nunca se ejecutó. Estas restricciones se comprueban al compilar y mantienen coherentes el alcance y la inicialización, aunque un camino específico pareciera seguro durante la ejecución.

## Question 7

- El `for` de tres partes, `for initialization; condition; update`, comunica una progresión controlada, como recorrer números del 0 al 9. Sería menos claro que `for-range` cuando solo se necesita procesar cada elemento de una colección.
- El `for` con una sola condición comunica una repetición que continúa mientras cambia un estado, igual que un `while` en otros lenguajes. Un `for` de tres partes sería más claro si la inicialización, el límite y el avance forman una secuencia numérica evidente.
- El `for` infinito, escrito `for`, comunica que la salida ocurre desde el cuerpo mediante `break`, `return` u otro evento. Una condición visible en la cabecera sería más clara cuando existe una regla de terminación sencilla que puede expresarse allí.
- `for-range` comunica el recorrido de los valores de un array, slice, map, string u otro operando permitido. Un loop de tres partes sería más claro cuando se necesita un salto de índices específico, recorrer solo una ventana numérica o controlar el avance manualmente.

## Question 8

Los contadores finales se declaran antes de ambos loops para que permanezcan en alcance cuando termine el recorrido. Un `for-range` exterior visita los strings en el orden del slice y un `for-range` interior decodifica cada string como runas. La variable de la runa pertenece al recorrido interno, mientras que cualquier estado exclusivo de un string se declara dentro de la iteración exterior.

El loop exterior recibe una etiqueta, por ejemplo `auditLoop`. Cuando el loop interior encuentra `!`, incrementa el total de rechazados y ejecuta `continue auditLoop`. Eso termina inmediatamente el análisis del string actual: las runas posteriores al marcador no se inspeccionan, el string no se reporta como aceptado y el siguiente string comienza.

Cuando aparece `X`, un `break auditLoop` termina el loop exterior completo. No se inspecciona el resto del string actual ni ningún string posterior. Un `break` sin etiqueta no serviría si estuviera dentro de un `switch`, porque solo terminaría el switch, ni bastaría para salir de los dos loops.

Las runas ordinarias se cuentan únicamente en la rama que no corresponde a ningún marcador. El reporte de aceptación y el incremento de strings aceptados se colocan después del loop interior; ese código solo se alcanza cuando se examinó el string completo. Un `switch` sobre la runa permite que `!`, `X` y el caso ordinario sean resultados mutuamente excluyentes.

No hubo preguntas dudosas: todas las respuestas cubren los puntos solicitados.
