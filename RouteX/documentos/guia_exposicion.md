# Guía rápida para presentar RouteX

## Apertura
“RouteX es un prototipo de gestión y optimización de rutas para una empresa de transporte. El objetivo académico es resolver el mismo problema en Python, JavaScript y Go, conservando datos y reglas de negocio para comparar sus paradigmas y ejecución.”

## Problema
Una empresa necesita administrar ubicaciones, caminos, vehículos y solicitudes. Debe encontrar rutas según distancia, tiempo o costo, considerar tráfico y asignar unidades disponibles.

## Solución técnica
La red se modela como un grafo no dirigido: las ubicaciones son vértices y las conexiones son aristas ponderadas. Dijkstra encuentra el camino óptimo para cada criterio. El tiempo se ajusta mediante factores de tráfico: bajo 1.0, medio 1.3 y alto 1.8.

## Comparación de lenguajes
- Python: alto nivel, tipado dinámico, multiparadigma; clases y funciones; rápido de desarrollar.
- JavaScript con Node.js: alto nivel, tipado dinámico, orientado a eventos y asíncrono; adecuado para servicios y aplicaciones web.
- Go: alto nivel, tipado estático y compilado a código máquina; goroutines y canales facilitan concurrencia.

## Demostración sugerida
1. Abrir la visualización y explicar los nodos y conexiones.
2. Ejecutar Python: `python python/src/routex.py`.
3. Ejecutar JavaScript desde `javascript`: `npm start`.
4. Ejecutar Go desde `go`: `go run ./cmd/routex`.
5. Mostrar las pruebas automatizadas.
6. Explicar que las tres versiones usan los mismos archivos JSON.

## Conclusión
“El proyecto demuestra que un mismo problema puede resolverse con distintos lenguajes y paradigmas. La equivalencia funcional permite comparar claridad, organización, tipado, ejecución y concurrencia sin cambiar las reglas del problema.”

## Transparencia para la exposición
Los datos son ficticios y el sistema no consulta tráfico real ni mapas comerciales. La página visual es demostrativa. Los tiempos de ejecución deben medirse en el mismo equipo y bajo las mismas condiciones antes de presentar conclusiones de rendimiento.
