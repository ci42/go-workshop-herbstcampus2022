Checke die Übung aus:

    git checkout exercises -- stats/cnt_test.go

Wir wollen unseren ersten eigenen Typ kreieren: Er soll dazu dienen bei der
Meta-Image-Suche eine kleine Statistik abzubilden. Der Name soll
`Cnt` sein und er soll zusammengesetzt sein und zwei Properties enthalten:
Eine mit dem Namen `Source` vom Typ `string` und eine weitere mit dem Namen
`Total` vom Typ `int`.

Erzeuge den Typen in einer neuen Datei mit dem Namen `cnt.go`. Diese
Datei gehört zu dem package `stats`.

Teste die Lösung mit:

    go test ./...

Eine (mögliche) Lösung findest Du hier:

    git checkout solutions -- stats/cnt_solution.go

