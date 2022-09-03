Checke die Übung aus:

    git checkout exercises -- toolusage/fmt_test.go toolusage/toolusage.go

Führe Tests mit

    go test ./...

aus und behebe die Fehler. Es wird hier lediglich die korrekte Formatierung
der Dateien überprüft - der Code an sich kann erst einmal ignoriert werden.

Go bringt ein Kommando zum automatischen Formatieren von Dateien mit und es
ist üblich, dies zu nutzen. Rufe

    go fmt

auf, prüfe, dass die Tests durchlaufen.
