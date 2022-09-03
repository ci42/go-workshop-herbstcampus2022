Checke die Übung aus:

    git checkout exercises -- web/tmpl_caching_test.go web/templates/index.html 'web/static/*'

Wir wollen das Template unserer Webapp im Speicher cachen:

Erzeuge in eine globale Variable `indexTemplate` vom Type `*template.Template` im Package `web`

Lese das template für unsere webapp via `init()` Funktion unter Benutzung von `NewTemplateFromFile`
ein. Die Template Datei heißt `./web/templates/index.html`. Speicher das Template in der globalen `indexTemplate`
Variable.

Eine (mögliche) Lösung findest Du hier:

    git checkout solutions -- web/tmpl_caching_solution.go