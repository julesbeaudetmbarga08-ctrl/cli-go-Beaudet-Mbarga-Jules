Projet CLI Go
 Installation
Prérequis : Go 1.22 ou plus récent.

git clone https://github.com/<user>/cli-go-Beaudet-Mbarga-Jules.git
cd cli-go-Beaudet-Mbarga-Jules/trace

 Utilisation
Commande générale :
go run . <entree> <sortie>
go run . tests/T01-entree.txt sortie.txt

 Règles du projet

(up)	met le mot en MAJUSCULES
(low)	met le mot en minuscules
(cap)	met la première lettre en majuscule
(hex)	interprète le mot comme un nombre hexadécimal
(bin)	interprète le mot comme un nombre binaire


Architecture du projet
main.go
  ├── vérifie les arguments
  ├── appelle fileio.Lire
  ├── appelle transform.Texte
  └── appelle fileio.Ecrire

fileio/
  ├── Lire(chemin string) (string, error)
  └── Ecrire(chemin string, texte string) error

transform/
  ├── Texte(texte string) (string, error)
  ├── Ligne(ligne string) (string, error)
  └── autres fonctions de transformation

Tests
Les tests sont dans le dossier tests/.

go run . tests/T3-entree.txt sortie.txt
diff sortie.txt tests/T3-attendu.txt

