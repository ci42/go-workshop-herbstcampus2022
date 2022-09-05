Checke die Übung aus:

    git checkout exercises -- web/tmpl_test.go web/templates/test_template.txt

Wir wollen eine Datei einlesen (das html template für unsere Suchmachine):
Erstelle eine Funktion mit der Signatur

    func NewTemplateFromFile(fileName string) (*template.Template, error)

im Package `web` welche eine Datei mit dem namen `fileName` einliest und daraus ein Template erstellt.
Sie soll einen Fehler zurückliefern, falls ein Fehler bei dem einlesen der Datei
oder bei dem Erstellen des Templates auftritt.

Nützliche Doku:

https://go.dev/pkg/io/ioutil/

https://go.dev/pkg/text/template/

Teste die Lösung mit:

    go test ./...

Eine (mögliche) Lösung findest Du hier:

    git checkout solutions -- web/tmpl_solution.go
