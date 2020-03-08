# ReFingerprint Extension

Dies ist die WebExtension, welche Daten aufzeichnet, die bei der Lösung eines ReCaptchas, welches in der Umfragen-Applikation angezeigt wird, aufzeichnet.

In den Chrome Einstellungen sollte die (erweiterte) Option "Downloads -> Vor dem Download von Dateien nach dem Speicherort fragen" ausgeschaltet sein.
Zusätzlich muss Chrome mit den Argumenten `--disable-web-security`, `--disable-site-isolation-trials` und `--user-data-dir="<beliebiger Ordner>` gestartet werden, um die Same Origin Policy zu unterdrücken, sodass die Extension auch auf das ReCaptcha IFrame zugreifen kann.
Alle anderen Chrome oder Chromium Instanzen müssen davor geschlossen werden.

Bei Firefox ist es nicht möglich, diese Policy zu unterdrücken, deshalb wird Firefox nicht unterstützt.
