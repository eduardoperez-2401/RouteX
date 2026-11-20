import fs from "node:fs";
import { RouteX } from "./routex.js";
const dataDir = new URL("../../datos/", import.meta.url);
const read = name => JSON.parse(fs.readFileSync(new URL(name, dataDir), "utf8"));
const app = new RouteX(read("ubicaciones.json"), read("conexiones.json"), read("vehiculos.json"));
console.log("ROUTEX | Consulta A -> F");
for (const [criterion,label] of [["distance","Más corta"],["time","Más rápida"],["cost","Menor costo"]]) {
  console.log(`\n${label}:`, app.shortest("A","F",criterion));
}
console.log("\nRutas alternativas:", app.alternatives("A","F"));
console.log("\nVehículo asignado:", app.assignVehicle(6)?.id ?? "Solicitud pendiente");
