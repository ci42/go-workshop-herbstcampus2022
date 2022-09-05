Checke die Übung aus:

    git checkout exercises -- stats/stats_stringer_test.go

Schreibe die Methode `String()` für den Typ `Stats`. Sie soll pro Cnt-Eintrag einen String zurückgeben.
Folgender Aufruf:

```
println(
  Stats{
    Cnt{"Pexels", 20},
    Cnt{"Unsplash", 100},
  }.String()
)
```

soll folgenden Output produzieren:

```
Pexels: 20
Unsplash: 100
```

Teste die Lösung mit:

    go test ./...

Eine (mögliche) Lösung findest Du hier:

    git checkout solutions -- stats/stats_stringer_solution.go




