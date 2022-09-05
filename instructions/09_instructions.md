Checke die Übung aus:

    git checkout exercises -- internal/concurrency/deadlock_test.go internal/concurrency/deadlock.go

Fixe den Deadlock in der Funktion helloTony in der Datei `race_condition.go`.
Nutze zur Lösung eine Go-Routine, nicht einen gepufferten Channel.

Teste die Lösung mit:

    go test ./...

Eine (mögliche) Lösung findest Du hier:

    git checkout solutions -- internal/concurrency/deadlock_solution.go