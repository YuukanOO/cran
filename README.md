Compte Rendu de l'Assemblée Nationale ✨ [![Build Status](https://travis-ci.org/YuukanOO/cran.svg?branch=master)](https://travis-ci.org/YuukanOO/cran) [![codecov](https://codecov.io/gh/YuukanOO/cran/branch/master/graph/badge.svg)](https://codecov.io/gh/YuukanOO/cran) [![Go Report Card](https://goreportcard.com/badge/github.com/YuukanOO/cran)](https://goreportcard.com/report/github.com/YuukanOO/cran)
===

Récemment, je me suis retrouvé à lire un compte rendu sur le site de l'assemblée nationale. J'ai trouvé la mise en page peu adaptée au contenu et le manque d'informations sur les participants fort dommage.

Comme ça faisait un moment que je voulais testé le langage **Go**, je me suis dis que ça ferait un bon petit TP !

Et voici le résultat avec à gauche la version originale et à droite la version sortie par cet outil :

<div align="center">
  <img src="source.png" width="280px"></img>
  <img src="pretty.png" width="280px"></img>
</div>

*Une démo est hébergée sur [heroku](https://powerful-scrubland-26285.herokuapp.com/) !*

## Récupération et lancement

Il vous faudra une version de **Go** installée sur votre machine (1.11 ou ultérieur).

La commande `go run` ci-dessous ira chercher les dépendances nécessaires au projet puis lancera le serveur web.

```bash
$ git clone https://github.com/YuukanOO/cran.git
$ cd cran
$ go run cmd/cran-web/main.go
```

## Tests

Pour lancer les tests, lancer la commande suivante :

```bash
$ go test ./... --cover
```