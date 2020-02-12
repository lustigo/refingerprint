# ReFingerprint Survey

Dies ist das Hauptverzeichnis der Umfragen-Applikation.

Der Frontend-Teil wurde mit dem Vue-Framework im Material Design erstellt.

Der Backend-Teil wurde mit Node.js umgesetzt.

Beide Teile verwenden Typescript.

## Inhaltsverzeichnis

- [Installation](#installation)
- [Konfiguration](#konfiguration)
  - [Server](#konfiguration-server)
  - [Umfragen](#konfiguration-umfragen)
    - [Widgets](#konfiguration-widgets)
      - [Captcha](#konfiguration-captcha)
      - [Checkbox](#konfiguration-checkbox)
      - [Dropdown](#konfiguration-dropdown)
      - [Matrix](#konfiguration-matrix)
      - [Textbox](#konfiguration-textbox)
      - [Textfield](#konfiguration-textfield)
      - [Text](#konfiguration-text)
      - [Vector](#konfiguration-vector)
- [Benutzung](#benutzung)

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

- [Server](#konfiguration-server)
- [Umfragen](#konfiguration-survey)
  - [Widgets](#konfiguration-widgets)

### Konfiguration Server

Für das Starten des Servers benötigen Sie einen API-Schlüssel von ReCaptcha, welchen Sie [hier](https://www.google.com/recaptcha/intro/v3.html) in der Admin Console erhalten. Bitte achten Sie darauf, dass Sie die die reCAPTCHA Version 2 mit dem Kästchen "Ich bin kein Roboter." auswählen.

Sie erhalten einen Websitenschlüssel, welchen Sie [später](#konfiguration-captcha) benötigen, und einen Geheimen Schlüssel, welchen Sie in den obigen Befehl für das erste Starten (anstatt `<Secret einfügen>`) einfügen.

Zudem müssen Sie einen Ordner erstellen und den relativen Pfad anstatt `<relativer Pfad>` in denselben Befehl einfügen.
Dieser Ordner ist der Hauptordner für die zu erstellenden Umfragen-Ordner, welche die Beschreibung und anschließend die Rückmeldungen enthalten werden.

In diesem Ordner müssen zudem von Ihnen noch der TLS-Schlüssel und das TLS-Zertifikat abgelegt werden.

Wenn Sie das Programm lokal testen wollen, können Sie mit dem folgendem Befehl ein selbst-signiertes Zertifikat erstellen:

```bash
openssl req -newkey rsa:4096 -nodes -keyout key.pem -x509 -days 365 -out crt.pem
```

Bitte speichern Sie die beiden generierten Dateien in dem zuvor erstellten Ordner ab.

Falls Sie den Dateinamen ändern, oder die Dateien in einem Unterverzeichnis speichern möchten, müssen Sie die Umgebungsvariablen setzen.
Im Folgenden sehen Sie die Umgebungsvariablen für die TLS-Konfiguration mit deren Standardwerten:

```TOML
TLS_KEY_PATH=/data/key.pem
TLS_CERT_PATH=/data/crt.pem
```

Für den realen Einsatz fragen Sie bitte Ihren IT-Administrator oder erzeugen Sie ein Zertifikat für Ihre Domain und lassen dieses von einer CA unterschreiben (zum Beispiel von Let's Encrypt).

### Konfiguration Umfragen

Um eine neue Umfrage anzulegen, müssen Sie einen neuen Ordner in Ihrem Datenpfad anlegen.

Der Ordnername ist die eindeutige Bezeichnung der Umfrage und muss eine sogenannte UUID sein.
Um eine UUID zu generieren [klicken Sie hier](https://www.uuidgenerator.net/version4).

In diesem Verzeichnis werden die Antworten der Teilnehmer abgespeichert und die Beschreibung der Umfrage.

Die Beschreibung der Umfrage muss in dem Umfragenordner in der Datei `structure.json` liegen und ein valides `json`-Dokument sein.

Die Beschreibung hat folgende Form:

```json
{
  "survey": "<Hier die soeben generierte UUID einfügen>",
  "pages": [
    []
    ...
  ],
}
```

Der Pages-Wert ist ein Mehrdimensionales Array, er listet die verschiedenen Seiten auf.

Auf jeder Seite können mehrere Widgets platziert werden, deshalb ist jede Seite ein Array.

#### Konfiguration Widgets

Die Beschreibung eines Widgets ist folgendermaßen aufgebaut:

```json
{
  "type": "<Typbezeichnung>",
  "required": true|false,
  "structure": {}
}
```

Mit dem `required` Wert können Sie bestimmen, ob dieses Widget ausgefüllt werden muss oder nicht.

Mit dem `type` Wert bestimmen Sie, welches Widget Sie anzeigen möchten.
Die Werte für diesen Wert finden Sie bei der Dokumentation des jeweiligen Widgets.

Mit dem `structure` Wert konfigurieren Sie ein Widget näher.
Diese Konfiguration hängt aber von dem jeweiligen Widget ab.
Bitte entnehmen Sie das Konfigurationschema der Dokumentation des jeweiligen Widgets.

Folgende Widgets sind möglich:

- [Captcha](#konfiguration-captcha)
- [Checkbox](#konfiguration-checkbox)
- [Dropdown](#konfiguration-dropdown)
- [Matrix](#konfiguration-matrix)
- [Textbox](#konfiguration-textbox)
- [Textfield](#konfiguration-textfield)
- [Text](#konfiguration-text)
- [Vector](#konfiguration-vector)

##### Konfiguration Captcha

##### Konfiguration Checkbox

##### Konfiguration Dropdown

##### Konfiguration Matrix

##### Konfiguration Textbox

##### Konfiguration Textfield

##### Konfiguration Text

##### Konfiguration Vector

## Benutzung

Ihren Teilnehmern müssen Sie folgenden Link geben:

```URL
https://<IhreDomain>/?<Ihre Umfragen-UUID>
```

In dem Datenpfad unter Ihrer Umfragen-UUID finden Sie die Rückmeldungen Ihrer Teilnehmer in jeweiligen `json`-Dateien.