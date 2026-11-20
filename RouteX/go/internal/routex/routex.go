package routex

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type Location struct { ID string `json:"id"`; Name string `json:"name"` }
type Edge struct {
	From string `json:"from"`; To string `json:"to"`
	Distance float64 `json:"distance_km"`; BaseMinutes float64 `json:"base_minutes"`
	Cost float64 `json:"cost"`; Traffic string `json:"traffic"`
}
type Vehicle struct {
	ID string `json:"id"`; Name string `json:"name"`; Capacity int `json:"capacity"`
	Available bool `json:"available"`; Driver string `json:"driver"`
}
type Result struct {
	Nodes []string `json:"nodes"`; Names []string `json:"names"`
	Distance float64 `json:"distance_km"`; Time float64 `json:"time_minutes"`
	Cost float64 `json:"cost"`; OptimizedFor string `json:"optimized_for"`
}
type link struct { to string; edge Edge }
type Network struct { Names map[string]string; Graph map[string][]link; Vehicles []Vehicle }
var factors = map[string]float64{"bajo":1,"medio":1.3,"alto":1.8}

func New(locations []Location, edges []Edge, vehicles []Vehicle) (*Network,error) {
	n:=&Network{Names:map[string]string{},Graph:map[string][]link{},Vehicles:vehicles}
	for _,l:=range locations { n.Names[l.ID]=l.Name; n.Graph[l.ID]=[]link{} }
	for _,e:=range edges {
		if _,ok:=n.Graph[e.From]; !ok { return nil,fmt.Errorf("origen inexistente") }
		if _,ok:=n.Graph[e.To]; !ok { return nil,fmt.Errorf("destino inexistente") }
		if e.Distance<0 || e.BaseMinutes<0 || e.Cost<0 { return nil,fmt.Errorf("peso negativo") }
		n.Graph[e.From]=append(n.Graph[e.From],link{e.To,e})
		n.Graph[e.To]=append(n.Graph[e.To],link{e.From,e})
	}
	return n,nil
}
func weight(e Edge, criterion string) (float64,error) {
	switch criterion {
	case "distance": return e.Distance,nil
	case "time": f,ok:=factors[e.Traffic]; if !ok{return 0,fmt.Errorf("tráfico inválido")}; return e.BaseMinutes*f,nil
	case "cost": return e.Cost,nil
	default: return 0,fmt.Errorf("criterio inválido")
	}
}
type item struct { node string; cost float64; index int }
type priorityQueue []*item
func (p priorityQueue) Len() int{return len(p)}
func (p priorityQueue) Less(i,j int)bool{return p[i].cost<p[j].cost}
func (p priorityQueue) Swap(i,j int){p[i],p[j]=p[j],p[i];p[i].index=i;p[j].index=j}
func (p *priorityQueue) Push(x any){v:=x.(*item);v.index=len(*p);*p=append(*p,v)}
func (p *priorityQueue) Pop() any {old:=*p;n:=len(old);v:=old[n-1];*p=old[:n-1];return v}

func (n *Network) Shortest(origin,destination,criterion string)(*Result,error){
	if _,ok:=n.Graph[origin];!ok{return nil,fmt.Errorf("origen inexistente")}
	if _,ok:=n.Graph[destination];!ok{return nil,fmt.Errorf("destino inexistente")}
	dist:=map[string]float64{};prev:=map[string]link{};parent:=map[string]string{}
	for id:=range n.Graph {dist[id]=math.Inf(1)}
	dist[origin]=0
	pq:=&priorityQueue{};heap.Init(pq);heap.Push(pq,&item{node:origin,cost:0})
	for pq.Len()>0 {
		cur:=heap.Pop(pq).(*item)
		if cur.cost!=dist[cur.node]{continue};if cur.node==destination{break}
		for _,ln:=range n.Graph[cur.node] {
			w,err:=weight(ln.edge,criterion);if err!=nil{return nil,err}
			alt:=cur.cost+w
			if alt<dist[ln.to] {dist[ln.to]=alt;parent[ln.to]=cur.node;prev[ln.to]=ln;heap.Push(pq,&item{node:ln.to,cost:alt})}
		}
	}
	if math.IsInf(dist[destination],1){return nil,nil}
	nodes:=[]string{destination};edges:=[]Edge{};cursor:=destination
	for cursor!=origin {ln:=prev[cursor];edges=append(edges,ln.edge);cursor=parent[cursor];nodes=append(nodes,cursor)}
	for i,j:=0,len(nodes)-1;i<j;i,j=i+1,j-1 {nodes[i],nodes[j]=nodes[j],nodes[i]}
	for i,j:=0,len(edges)-1;i<j;i,j=i+1,j-1 {edges[i],edges[j]=edges[j],edges[i]}
	r:=&Result{Nodes:nodes,OptimizedFor:criterion}
	for _,id:=range nodes {r.Names=append(r.Names,n.Names[id])}
	for _,e:=range edges {r.Distance+=e.Distance;r.Time+=e.BaseMinutes*factors[e.Traffic];r.Cost+=e.Cost}
	return r,nil
}
func (n *Network) AssignVehicle(passengers int)*Vehicle {
	indices:=[]int{}
	for i,v:=range n.Vehicles {if v.Available && v.Capacity>=passengers {indices=append(indices,i)}}
	if len(indices)==0{return nil}
	sort.Slice(indices,func(i,j int)bool{return n.Vehicles[indices[i]].Capacity<n.Vehicles[indices[j]].Capacity})
	i:=indices[0];n.Vehicles[i].Available=false;return &n.Vehicles[i]
}
