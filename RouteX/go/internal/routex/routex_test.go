package routex
import "testing"
func sample()*Network{
 l:=[]Location{{"A","A"},{"B","B"},{"C","C"}}
 e:=[]Edge{{"A","B",2,3,10,"bajo"},{"B","C",2,3,10,"medio"}}
 v:=[]Vehicle{{"V1","Van",5,true,"Ana"}}
 n,_:=New(l,e,v);return n
}
func TestShortest(t *testing.T){r,e:=sample().Shortest("A","C","distance");if e!=nil||r==nil||r.Distance!=4{t.Fatal("ruta por distancia incorrecta")}}
func TestTime(t *testing.T){r,_:=sample().Shortest("A","C","time");if r.Time!=6.9{t.Fatalf("tiempo inesperado: %v",r.Time)}}
func TestCost(t *testing.T){r,_:=sample().Shortest("A","C","cost");if r.Cost!=20{t.Fatal("costo incorrecto")}}
func TestInvalidNode(t *testing.T){if _,e:=sample().Shortest("A","Z","distance");e==nil{t.Fatal("debió rechazar nodo")}}
func TestVehicle(t *testing.T){n:=sample();v:=n.AssignVehicle(4);if v==nil||v.Available{t.Fatal("asignación incorrecta")}}
func TestNoVehicle(t *testing.T){if sample().AssignVehicle(20)!=nil{t.Fatal("no debe asignar vehículo")}}
