Checke die Übung aus:

    git checkout exercises -- web/server/http_server_test.go

Wir bauen uns einen http-Server:

Schreibe eine Funktion `startHTTPServer` die einen Argument namens `port` vom Type `int` hat
und einen HTTP-Server startet, welcher den Dateibaum unter `.web/static` ausliefert.

Diese Funktion soll am Ende der `main()` Funktion mit dem Port `4242` aufgerufen werden.

  curl http://localhost:4242/index.html

sollte ein HTML-Template zurückliefern

Dokumentationen über HTTP Server und Client sind im Paket `net/http` zu finden.

Teste die Lösung mit:

    go test ./...

Eine (mögliche) Lösung findest Du hier:

    git checkout solutions -- web/server/http_server_solution.go
