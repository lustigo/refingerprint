# ReFingerprint

![Survey Frontend](https://github.com/lustigo/refingerprint/workflows/Survey%20Frontend/badge.svg)
![Survey Backend](https://github.com/lustigo/refingerprint/workflows/Survey%20Backend/badge.svg)
![Extension](https://github.com/lustigo/refingerprint/workflows/Extension/badge.svg)
![Processing](https://github.com/lustigo/refingerprint/workflows/Processing/badge.svg)

Dies ist eine Bachelorarbeit, in der Daten gesammelt werden, um durch die Nutzung von ReCaptcha einen Nutzer eindeutig identifizieren zu können.

Dies beinhaltet eine Umfragen-Applikation, welche beim Ausfüllen einer Umfrage zwischendurch ein ReCaptcha anzeigt.
Der Nutzer sollte nicht wissen, dass diese Umfrage dazu dient, Daten über die Benutzung von ReCaptcha zu sammeln.
Die Umfrage-Applikation befindet sich im `survey` Verzeichnis.

Während der Lösung des ReCaptchas werden einige Daten durch eine WebExtension aufgezeichnet und auf dem lokalen Computer abgelegt.
Diese liegt im `ext` Verzeichnis.

Diese Daten werden dann durch ein weiteres Programm umgewandelt um es für die Verwendung in einem neuronalen Netzwerk vorzubereiten.
Dieses liegt im `proc` Verzeichnis.
