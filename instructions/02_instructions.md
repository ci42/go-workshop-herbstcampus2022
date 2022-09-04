Check die Übung aus:

    git checkout exercises -- helloworld/hello_world_test.go helloworld/hello_world.go

Wir schreiben unsere erste Funktion: Hello World.

Implementiere in der Datei helloworld/hello_world.go eine Funktion

    func helloWorld() {
      // ...
    }

welche den String "Hello Herbstcampus 2022!" in einer Zeile ausgibt.
Mache dich dazu mit der Dokumentation von go vertraut:

https://go.dev/pkg/fmt/

Teste die Funktion mittels:

    go test ./...

Eine (mögliche) Lösung findest Du hier:

    git checkout solutions -- helloworld/hello_world_solution.go

