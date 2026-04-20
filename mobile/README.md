# Mobile

This directory is the monorepo home for the Flutter mobile app.

Use this folder for future mobile implementation work, including the server-driven UI rendering layer and app-specific presentation code.

## Run Locally

```bash
flutter pub get
flutter run --dart-define=API_BASE_URL=http://localhost:8080
```

## Tests

```bash
flutter test
flutter test --coverage
```

Coverage report is written to `coverage/lcov.info`.
