package main
import (
 "encoding/json";"fmt";"os";"path/filepath";"sync"
 "routex/internal/routex"
)
func read[T any](name string) []T {
 b,err:=os.ReadFile(filepath.Join("..","..","datos",name))
 if err!=nil { b,err=os.ReadFile(filepath.Join("..","datos",name)) }
 if err!=nil { panic(err) };var v []T;if err=json.Unmarshal(b,&v);err!=nil{panic(err)};return v
}
func main(){
 loc:=read[routex.Location]("ubicaciones.json");edges:=read[routex.Edge]("conexiones.json");vehicles:=read[routex.Vehicle]("vehiculos.json")
 app,err:=routex.New(loc,edges,vehicles);if err!=nil{panic(err)}
 fmt.Println("ROUTEX | Consulta A -> F")
 criteria:=[]string{"distance","time","cost"};labels:=[]string{"Más corta","Más rápida","Menor costo"}
 var wg sync.WaitGroup
 for i,c:=range criteria {wg.Add(1);go func(i int,c string){defer wg.Done();r,e:=app.Shortest("A","F",c);if e!=nil{fmt.Println(e);return};fmt.Printf("\n%s: %+v\n",labels[i],r)}(i,c)}
 wg.Wait()
 v:=app.AssignVehicle(6);if v==nil{fmt.Println("\nVehículo asignado: solicitud pendiente")}else{fmt.Println("\nVehículo asignado:",v.ID)}
}
