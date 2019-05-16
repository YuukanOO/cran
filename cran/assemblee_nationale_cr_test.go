package cran

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="fr" lang="fr">
<head>
</head>
<body>
<div id="englobe">
	<h2 class="entete">Assembl&eacute;e nationale<br/>
	XV<sup>e</sup> l&eacute;gislature<br/>
	Session ordinaire de 2018-2019<br/><br/>
	Compte rendu<br/>
	int&eacute;gral</h2>
	<div class="SYCERON">
		<h1 class="seance">Deuxième séance du lundi 13 mai 2019</h1>
		<div id="somjo"></div>
		<div class="ouverture_seance">
			<h5 class="presidence">
				<a name="P1714018" id="P1714018"><span></span></a>Présidence de Mme&nbsp;Annie Genevard<br><br>vice-présidente
			</h5>
			<p>
				<b>
					<a href="http://www2.assemblee-nationale.fr/deputes/fiche/OMC_PA605991">Mme la présidente</a>
				</b>. La séance est ouverte.
			</p>
			<p> <i>(La séance est ouverte à vingt et une heures trente.)</i></p>
		</div>

		<div class="Point">
			<h5 class="numencad">1</h5>
			<h2 class="titre1">Transformation de la fonction publique<i></i></h2>
		</div>
		
		<div class="Point">
			<h2 class="titre2">Discussion générale</h2>
			<div class="intervention">
				<p>
					<b>
						<a href="http://www2.assemblee-nationale.fr/deputes/fiche/OMC_PA1874">M. le président</a>
					</b>.  La parole est à Mme&nbsp;Frédérique Dumas.
				</p>
			</div>

			<div class="Point">
				<h2 class="titre99">
					<i>Suspension et reprise de la séance</i>
				</h2>
				<p>
					<b>
						<a href="http://www2.assemblee-nationale.fr/deputes/fiche/OMC_PA1874">M. le président</a>
					</b>.  La séance est suspendue.
				</p>
			</div>
			<p>
				<b>
					<a href="http://www2.assemblee-nationale.fr/deputes/fiche/OMC_PA718868">M. Michel Larive</a>
				</b>.  Avec toutes les citations énumérées depuis ce matin, si les œuvres de Victor Hugo étaient toujours soumises à droits d’auteur, et si –&nbsp;comme il l’aurait souhaité&nbsp;– la manne financière participait au financement de la création d’artistes vivants, nous aurions ce soir suffisamment de subsides pour aider un grand nombre de créateurs français&nbsp;! 
				<i>(Sourires.)</i>
			</p>
		</div>
	</div>
</div>
</body>
</html>`))
	})

	return httptest.NewServer(mux)
}

func TestAssembleeNationaleCRParsing(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	p := &assembleNationaleCRProvider{}
	p.Fetch(ts.URL, func (report *Report, err error) {
		
	})
}

func TestNodes(t *testing.T) {
	r := newRoot()
	s1 := newSection("One section")
	s2 := newSection("Another one")
	s1.Append(s2)
	r.Append(s1)

	if len(r.Children()) != 1 {
		t.Error("Should be 1")
	}

	s := r.Children()[0].(*SectionNode)

	if s.Title != "One section" {
		t.Error("Title should be One section")
	}
}