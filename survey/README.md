# ReFingerprint Survey

Dies ist das Hauptverzeichnis der Umfragen-Applikation.

Der Frontend-Teil wird mit dem Vue-Framework im Material Design erstellt.

Der Backend-Teil wird mit Node.js umgesetzt.

Beide Teile verwenden Typescript.

## Inhaltsverzeichnis

- Installation
- Konfiguration
  - Server
  - Umfragen
    - Widgets
- Benutzung

## Installation

Container bauen:

```bash
DOCKER_BUILDKIT=1 docker build -t refingerprint .
```

Erstes Starten:

```bash
docker run -d -p 443:443 --env RECAPTCHA_SECRET=<Secret einfügen> -v "$(pwd)"<relativer Datenpfad>:/data --name refingerprint refingerprint
```

Stoppen:

```bash
docker stop refingerprint
```

Starten:

```bash
docker start refingerprint
```

## Konfiguration
## Benutzung