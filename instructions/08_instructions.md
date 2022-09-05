Checke die Übung aus:

    git checkout exercises -- stats/stats_inc_test.go

Für unseren `Stats` Typen möchten wir eine Funktion namens `Inc` implementieren
welche als Argument einen String-Parameter mit dem Namen `source` bekommt.

Die Funktion soll jeweils die Property `Total` für die jeweilige `Source` erhöhen.

Folgende Aufrufe


	s := Stats{}

	s.Inc("Pexels")
	s.Inc("Pexels")
	s.Inc("Pexels")
	s.Inc("Unsplash")

	fmt.Printf("%s", s)

sollten folgendes Ausgeben:

```
Pexels: 3
Unsplash: 1
```
Teste die Lösung mit:

    go test ./...

Eine (mögliche) Lösung findest Du hier:

    git checkout solutions -- stats/stats_inc_solution.go