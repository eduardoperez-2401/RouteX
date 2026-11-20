import json, heapq, argparse
from pathlib import Path

FACTORS = {"bajo": 1.0, "medio": 1.3, "alto": 1.8}

class RouteX:
    def __init__(self, locations, edges, vehicles):
        self.locations = {x["id"]: x for x in locations}
        self.graph = {x: [] for x in self.locations}
        for e in edges:
            if e["from"] not in self.graph or e["to"] not in self.graph:
                raise ValueError("Conexión con ubicación inexistente")
            if e["distance_km"] < 0 or e["base_minutes"] < 0 or e["cost"] < 0:
                raise ValueError("Los pesos no pueden ser negativos")
            self.graph[e["from"]].append((e["to"], e))
            self.graph[e["to"]].append((e["from"], e))
        self.vehicles = vehicles

    def edge_weight(self, edge, criterion):
        if criterion == "distance":
            return edge["distance_km"]
        if criterion == "time":
            return edge["base_minutes"] * FACTORS[edge["traffic"]]
        if criterion == "cost":
            return edge["cost"]
        raise ValueError("Criterio debe ser distance, time o cost")

    def shortest(self, origin, destination, criterion="distance"):
        if origin not in self.graph or destination not in self.graph:
            raise ValueError("Origen o destino inexistente")
        dist, prev = {n: float("inf") for n in self.graph}, {}
        dist[origin] = 0
        queue = [(0, origin)]
        while queue:
            current, node = heapq.heappop(queue)
            if current != dist[node]:
                continue
            if node == destination:
                break
            for neighbor, edge in self.graph[node]:
                candidate = current + self.edge_weight(edge, criterion)
                if candidate < dist[neighbor]:
                    dist[neighbor], prev[neighbor] = candidate, (node, edge)
                    heapq.heappush(queue, (candidate, neighbor))
        if dist[destination] == float("inf"):
            return None
        nodes, used = [destination], []
        cursor = destination
        while cursor != origin:
            parent, edge = prev[cursor]
            used.append(edge)
            cursor = parent
            nodes.append(cursor)
        nodes.reverse(); used.reverse()
        return self.summarize(nodes, used, criterion)

    def summarize(self, nodes, edges, criterion):
        return {
            "nodes": nodes,
            "names": [self.locations[n]["name"] for n in nodes],
            "distance_km": round(sum(e["distance_km"] for e in edges), 2),
            "time_minutes": round(sum(e["base_minutes"] * FACTORS[e["traffic"]] for e in edges), 2),
            "cost": round(sum(e["cost"] for e in edges), 2),
            "optimized_for": criterion
        }

    def alternatives(self, origin, destination, limit=5):
        if origin not in self.graph or destination not in self.graph:
            raise ValueError("Origen o destino inexistente")
        found = []
        def visit(node, path, used):
            if len(found) >= 100:
                return
            if node == destination:
                found.append(self.summarize(path, used, "distance"))
                return
            for nxt, edge in self.graph[node]:
                if nxt not in path:
                    visit(nxt, path + [nxt], used + [edge])
        visit(origin, [origin], [])
        return sorted(found, key=lambda x: (x["distance_km"], x["time_minutes"]))[:limit]

    def assign_vehicle(self, passengers):
        candidates = [v for v in self.vehicles if v["available"] and v["capacity"] >= passengers]
        if not candidates:
            return None
        chosen = min(candidates, key=lambda v: v["capacity"])
        chosen["available"] = False
        return chosen

def load_data(root=None):
    root = Path(root) if root else Path(__file__).resolve().parents[2] / "datos"
    def read(name): return json.loads((root / name).read_text(encoding="utf-8"))
    return read("ubicaciones.json"), read("conexiones.json"), read("vehiculos.json"), read("solicitudes.json")

def main():
    parser = argparse.ArgumentParser(description="RouteX - gestión y optimización de rutas")
    parser.add_argument("--origen", default="A")
    parser.add_argument("--destino", default="F")
    args = parser.parse_args()
    locations, edges, vehicles, _ = load_data()
    system = RouteX(locations, edges, vehicles)
    print("ROUTEX | Consulta A -> F")
    for criterion, label in [("distance","Más corta"),("time","Más rápida"),("cost","Menor costo")]:
        result = system.shortest(args.origen, args.destino, criterion)
        print(f"\n{label}:")
        print(json.dumps(result, ensure_ascii=False, indent=2))
    print("\nRutas alternativas:")
    for route in system.alternatives(args.origen, args.destino):
        print(" -> ".join(route["nodes"]), f'| {route["distance_km"]} km | {route["time_minutes"]} min | ${route["cost"]}')
    vehicle = system.assign_vehicle(6)
    print("\nVehículo asignado:", vehicle["id"] if vehicle else "Solicitud pendiente")

if __name__ == "__main__":
    main()
