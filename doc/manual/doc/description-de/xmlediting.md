title: Bearbeitung des Layoutregelwerks
---
Bearbeitung des Layoutregelwerks
================================

Durch die Benutzung von Standard XML (utf-8) lässt sich das
Layoutregelwerk mit jedem Texteditor bearbeiten. Mithilfe des
mitgelieferten RelaxNG-Schemas wird die Eingabe des Layoutregelwerks
deutlich erleichtert. Dazu muss ein XML-Editor benutzt werden, der
RelaxNG Schemas verarbeiten kann:

-   [OxygenXML](https://www.oxygenxml.com) (Mac, Windows, Linux)
-   [XMLSpy](http://www.altova.com/xml-editor/) (Windows)
-   [XML Blueprint](https://www.xmlblueprint.com/) (nur Windows)
-   [GNU Emacs](https://www.gnu.org/software/emacs/) mit [nxml-mode](http://www.thaiopensource.com/nxml-mode/) (alle Betriebssysteme, kostenlos)
-   [jEdit](http://www.jedit.org) (Mac, Windows, Linux, kostenlos)

Das RelaxNG-Schema für das Layoutregelwerk liegt in der ZIP-Datei im
Verzeichnis `share/schema/` unter dem Dateinamen `layoutschema-en.rng` für das Regelwerk.

Hinweis
-------

Der XML-Namensraum des Layoutregelwerks ist
`urn:speedata.de:2009/publisher/en`. Daher muss das Layoutregelwerk,
sofern es gegen das RelaxNG Schema validiert wird, wie folgt aussehen:

    <Layout xmlns="urn:speedata.de:2009/publisher/en">
     ...
    </Layout>

Falls layoutspezifische XPath-Funktionen benutzt werden, muss der
Namensraum `urn:speedata:2009/publisher/functions/en` angegeben werden:

    <Layout xmlns="urn:speedata.de:2009/publisher/en"
            xmlns:sd="urn:speedata:2009/publisher/functions/en">
     ...
    </Layout>

Verknüpfung von Layoutregelwerk und RelaxNG Schema im XML-Editor
----------------------------------------------------------------

Beispielhaft soll hier anhand des OxygenXML Editors gezeigt werden, wie
das Schema den Dokumenten zugewiesen werden kann. Neben der Möglichkeit,
das Schema jedem Dokument einzeln zuzuweisen:

{{ img . "oxygen-schemazuweisen.png" }}

gibt es die dauerhafte Lösung, allen Dokumenten im Namensraum
`urn:speedata.de:2009/publisher/en` das RelaxNG Schema zuzuweisen.

Dazu erstellt man in den Einstellungen unter „Dokumenttypen Zuordnung“
eine neue Regel.

{{ img . "oxygen-doczuordnung1.png" }}

Füllen Sie den Dialog wie in der nächsten Abbildung aus:

{{ img . "oxygen-doczuordnung.png" }}

und klicken Sie auf auf den `+`-Button und erstellen die Regel wie in
dieser Abbildung:

{{ img . "namensraumzuordnung.png" }}


Anschließend muss noch der Pfad zur Schemadatei angegeben werden
(`.../schema/layoutschema-en.rng`):


{{ img . "schemazuweisung-oxygen.png" }}

Die Dialoge müssen mit `OK` bestätigt werden. Sobald ein Layoutregelwerk
geöffnet wird, ist dem Dokument ein Schema zugeordnet man erhält die
volle Editorunterstützung mit Autovervollständigung und Tooltips.

Wenn der XML-Editor RelaxNG mit eingebetteten Schematron-Regeln unterstützt, ist es hilfreich, diese zu aktivieren. Dadurch werden zusätzliche Fehlerquellen während der Eingabe erkannt und gemeldet.

