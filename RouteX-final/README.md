# RouteX – Sistema de gestión y optimización de rutas

Prototipo académico que resuelve el mismo problema mediante Python, JavaScript (Node.js) y Go.

## Objetivo
Administrar ubicaciones, conexiones, vehículos y solicitudes; representar la red como grafo y calcular rutas óptimas por distancia, tiempo y costo.

## Reglas comunes
- Grafo no dirigido: cada conexión se puede recorrer en ambos sentidos.
- Dijkstra se ejecuta con pesos no negativos.
- Tiempo real estimado = tiempo base × factor de tráfico.
- Factores: bajo 1.0, medio 1.3, alto 1.8.
- El vehículo elegido es el disponible con capacidad suficiente más pequeña.
- Si no hay ruta o vehículo adecuado, se informa la situación.

## Requisitos
- Python 3.10 o superior.
- Node.js 18 o superior.
- Go 1.21 o superior.

## Ejecutar
Desde la carpeta raíz:

### Python
```bash
python python/src/routex.py
python -m unittest discover -s python/tests -v
```

### JavaScript
```bash
cd javascript
npm start
npm test
cd ..
```

### Go
```bash
cd go
go run ./cmd/routex
go test ./...
cd ..
```

### Visualización
Abrir `visualizacion/index.html` en un navegador. La página muestra la red y permite consultar las tres métricas con los datos de demostración.

## Datos
Los archivos JSON de `datos/` son la fuente compartida de ubicaciones, conexiones, vehículos y solicitudes.

## Estructura
- `python/`: implementación con clases, funciones, estructuras y excepciones.
- `javascript/`: implementación multiparadigma en Node.js.
- `go/`: implementación compilada y concurrencia con goroutines.
- `visualizacion/`: tablero demostrativo.
- `documentos/`: análisis comparativo y guía de exposición.

## Alcance y limitaciones
Es un prototipo académico con datos ficticios. No utiliza mapas reales, GPS, API de tráfico ni persistencia transaccional. La búsqueda de rutas alternativas enumera caminos simples y limita el recorrido para evitar explosión combinatoria. El programa Go demuestra concurrencia al ejecutar en paralelo las tres consultas de optimización.

## Captura de datos desde la interfaz visual

Abre `visualizacion/index.html` en Chrome o Edge. En **Administración de datos** puedes registrar ubicaciones, conexiones, vehículos y solicitudes. La interfaz recalcula las rutas con los datos capturados.

Los cambios se guardan en el almacenamiento local del navegador. Para conservarlos o trasladarlos, utiliza **Importar / exportar → Descargar todos los datos (JSON)**. El respaldo contiene las colecciones `locations`, `connections`, `vehicles` y `requests`. Para que las versiones de consola utilicen esos datos, actualiza los archivos equivalentes dentro de `datos/` (`ubicaciones.json`, `conexiones.json`, `vehiculos.json` y `solicitudes.json`) con la estructura correspondiente.

La visualización es una demostración local: no utiliza GPS, mapas reales ni tráfico en tiempo real.
