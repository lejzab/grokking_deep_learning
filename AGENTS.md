# coding style
- Dodawaj do funkcji komentarze po polsku.
- Funkcje nie powinny panikować, powinny zwracać błędy.
- Sprawdzaj błędy.
- Zachowuj strukturę projektu zgodnie z [golang-standards/project-layout](https://github.com/golang-standards/project-layout):
    - `cmd/` - główne aplikacje (np. `cmd/grokking/main.go`).
    - `internal/` - logika specyficzna dla projektu (prywatna).
    - `pkg/` - biblioteki ogólnego przeznaczenia (publiczne, np. `pkg/math`).
- Używaj konkretnych nazw pakietów zamiast generycznych typu `utils`.

# tests
- Utrzymuj testy. Każda funkcje powinno być przetestowane.
- Pliki testowe powinny znajdować się w tym samym pakiecie co kod źródłowy.
- Każda nowa funkcjonalność musi posiadać testy jednostkowe.

