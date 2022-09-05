Checke die Übung aus:

    git checkout exercises -- internal/metasearch/metasearch_test.go internal/metasearch/api_keys.go

Erstelle in der Datei `search.go` eine Funktion mit der Signatur

    func MetaImageSearch(q string, timeout time.Duration) ([]Image, Stats, error)

Diese soll folgende Ergebnisse zurückliefern:

- `[]Image` : eine Liste von Bildern die verschiedene Suchmachinen zum Suchbegriff `q` gefunden haben

    Füge dazu die Bibliothek scraper ein: `go get github.com/ci42/ImgScraper/scraper` und nutze den Typ `scraper`
    um die verschiedenen Bildersuchen anzusprechen. Trage die passende API-Keys in die Datei `internal/metasearch/api_keys.go` ein. 
  
- `Stats`: Statistiken welche Suchmachine wieviel Bilder zurückgeliefert hat

- `error`: Fehler-Objekt falls ein Fehler aufgetreten ist, `nil` sonst

Optialerweise sollen alle drei Bildersuchmachinen (pixabay, pexels und unsplash) gleichzeitig abgefragt werden. Wird
das übergebene Timeout überschritten, soll sowohl ein Fehler als auch die bisher eingesammelten Bilder und Stats
zurückgeliefert werden.

Teste die Lösung mit:

    go test ./...

Eine (mögliche) Lösung findest Du hier:

    git checkout solutions -- internal/metasearch/metasearch_solution.go



