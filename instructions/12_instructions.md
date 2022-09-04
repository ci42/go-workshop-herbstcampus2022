Checke die Übung aus:

    git checkout exercises -- cmd/flags_test.go

Die Meta-Suche soll beim Start einen Parameter entegegennhemen können: mit `--timeout`
soll bestimmt werden können, wie lange maximal nach Bildern gesucht werden soll.

Es soll dabei möglich sein, Angaben wie: `--timeout=0.1s` oder `--timeout=400ms` zu machen.

Das Paket `flag` ist dein Freund.

Eine (mögliche) Lösung findest Du hier:

    git checkout solutions -- cmd/flags_solution.go
