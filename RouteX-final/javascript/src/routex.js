export const FACTORS = { bajo: 1.0, medio: 1.3, alto: 1.8 };

export class RouteX {
  constructor(locations, edges, vehicles) {
    this.locations = new Map(locations.map(x => [x.id, x]));
    this.graph = new Map(locations.map(x => [x.id, []]));
    for (const e of edges) {
      if (!this.graph.has(e.from) || !this.graph.has(e.to)) throw new Error("Ubicación inexistente");
      if (e.distance_km < 0 || e.base_minutes < 0 || e.cost < 0) throw new Error("Peso negativo");
      this.graph.get(e.from).push([e.to, e]);
      this.graph.get(e.to).push([e.from, e]);
    }
    this.vehicles = vehicles;
  }
  weight(e, criterion) {
    if (criterion === "distance") return e.distance_km;
    if (criterion === "time") return e.base_minutes * FACTORS[e.traffic];
    if (criterion === "cost") return e.cost;
    throw new Error("Criterio inválido");
  }
  shortest(origin, destination, criterion = "distance") {
    if (!this.graph.has(origin) || !this.graph.has(destination)) throw new Error("Origen o destino inexistente");
    const dist = new Map([...this.graph.keys()].map(k => [k, Infinity]));
    const prev = new Map(); const pending = new Set(this.graph.keys());
    dist.set(origin, 0);
    while (pending.size) {
      let node = null, best = Infinity;
      for (const candidate of pending) if (dist.get(candidate) < best) [best, node] = [dist.get(candidate), candidate];
      if (node === null || best === Infinity) break;
      pending.delete(node);
      if (node === destination) break;
      for (const [next, edge] of this.graph.get(node)) {
        const alt = best + this.weight(edge, criterion);
        if (pending.has(next) && alt < dist.get(next)) { dist.set(next, alt); prev.set(next, [node, edge]); }
      }
    }
    if (dist.get(destination) === Infinity) return null;
    const nodes = [destination], used = []; let cursor = destination;
    while (cursor !== origin) { const [parent, edge] = prev.get(cursor); used.push(edge); cursor = parent; nodes.push(cursor); }
    nodes.reverse(); used.reverse();
    return this.summarize(nodes, used, criterion);
  }
  summarize(nodes, edges, criterion) {
    return { nodes, names: nodes.map(n => this.locations.get(n).name),
      distance_km: Number(edges.reduce((s,e) => s + e.distance_km, 0).toFixed(2)),
      time_minutes: Number(edges.reduce((s,e) => s + e.base_minutes * FACTORS[e.traffic], 0).toFixed(2)),
      cost: Number(edges.reduce((s,e) => s + e.cost, 0).toFixed(2)), optimized_for: criterion };
  }
  alternatives(origin, destination, limit = 5) {
    if (!this.graph.has(origin) || !this.graph.has(destination)) throw new Error("Ubicación inexistente");
    const found = [];
    const visit = (node, path, used) => {
      if (found.length >= 100) return;
      if (node === destination) { found.push(this.summarize(path, used, "distance")); return; }
      for (const [next, edge] of this.graph.get(node)) if (!path.includes(next)) visit(next, [...path,next], [...used,edge]);
    };
    visit(origin, [origin], []);
    return found.sort((a,b) => a.distance_km-b.distance_km || a.time_minutes-b.time_minutes).slice(0,limit);
  }
  assignVehicle(passengers) {
    const candidates = this.vehicles.filter(v => v.available && v.capacity >= passengers).sort((a,b) => a.capacity-b.capacity);
    if (!candidates.length) return null;
    candidates[0].available = false; return candidates[0];
  }
}
