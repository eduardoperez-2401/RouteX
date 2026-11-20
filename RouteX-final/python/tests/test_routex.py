import sys, unittest
from pathlib import Path
sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "src"))
from routex import RouteX, load_data

class RouteXTests(unittest.TestCase):
    def setUp(self):
        l,e,v,_ = load_data()
        self.app = RouteX(l,e,v)
    def test_shortest_distance_exists(self):
        self.assertIsNotNone(self.app.shortest("A","F","distance"))
    def test_fastest_route(self):
        r = self.app.shortest("A","F","time")
        self.assertEqual(r["optimized_for"], "time")
    def test_cheapest_route(self):
        self.assertGreater(self.app.shortest("A","F","cost")["cost"], 0)
    def test_disconnected(self):
        self.assertIsNone(self.app.shortest("A","Z","distance")) if False else self.assertRaises(ValueError, self.app.shortest, "A","Z","distance")
    def test_vehicle_assignment(self):
        v = self.app.assign_vehicle(6)
        self.assertIsNotNone(v)
        self.assertFalse(v["available"])
    def test_no_capacity(self):
        self.assertIsNone(self.app.assign_vehicle(99))

if __name__ == "__main__": unittest.main()
