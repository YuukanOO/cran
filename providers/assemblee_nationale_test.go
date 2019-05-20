package providers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"cran/domain"
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
		<h1 class="seance">Titre de la séance</h1>
		<div id="somjo"></div>
		<div class="ouverture_seance">
			<h5 class="presidence">
				<a name="P1714018" id="P1714018"><span></span></a>Civilité président<br><br>poste président
			</h5>
			<p>
				<b>
					<a href="http://www2.assemblee-nationale.fr/deputes/fiche/OMC_PA605991">John Doe</a>
				</b>. La séance est ouverte.
			</p>
			<p> <i>(date ouverture séance)</i></p>
		</div>

		<div class="Point">
			<h5 class="numencad">1</h5>
			<h2 class="titre1">Titre 1<i></i></h2>
		</div>
		
		<div class="Point">
			<h2 class="titre2">Sous Titre 1</h2>
			<div class="intervention">
				<p>
					<b>
						<a href="http://www2.assemblee-nationale.fr/deputes/fiche/OMC_PA1874">John Doe</a>
					</b>.  La parole est à Jean Dupont.
				</p>
			</div>

			<div class="Point">
				<h2 class="titre99">
					<i>Titre 99</i>
				</h2>
				<p>
					<b>
						<a href="http://www2.assemblee-nationale.fr/deputes/fiche/OMC_PA1874">John Doe</a>
					</b>.  La séance est suspendue.
				</p>
			</div>
			<p>
				<b>
					<a href="http://www2.assemblee-nationale.fr/deputes/fiche/OMC_PA718868">Jean Dupont</a>
				</b>.  J'ai des choses à dire !
			</p>
		</div>

		<div class="Point">
			<h5 class="numencad">2</h5>
			<h2 class="titre1">Titre 2<i></i></h2>
			<p>
				<b>
					<a href="http://www2.assemblee-nationale.fr/deputes/fiche/OMC_PA718868">Jean Dupont</a>
				</b>.  Et encore d'autres ici !
			</p>
		</div>
	</div>
</div>
</body>
</html>`))
	})

	return httptest.NewServer(mux)
}

func TestAssembleeNationaleProvider(t *testing.T) {
	assert := assert.New(t)
	ts := newTestServer()
	defer ts.Close()
	p := &assemblee_nationale{}

	p.Fetch(ts.URL, func(report *domain.Report, err error) {
		assert.NotNil(report, "Report should not be nil")
		assert.Equal("Titre de la séance", report.Title, "Titles should match")

		children := report.Children()

		if assert.Len(children, 3, "It should have 3 sections") {
			section, _ := children[0].(*domain.Section)

			if assert.NotNil(section) {
				assert.Equal("Civilité présidentposte président", section.Title)

				sectionChildren := section.Children()

				if assert.Len(sectionChildren, 2, "Section should have 2 children") {
					intervention, _ := sectionChildren[0].(*domain.Intervention)

					if assert.NotNil(intervention) {
						assert.Equal("John Doe", intervention.SpeakerID)
						assert.Equal("La séance est ouverte.", intervention.Content)
					}

					notice, _ := sectionChildren[1].(*domain.Notice)

					if assert.NotNil(notice) {
						assert.Equal("(date ouverture séance)", notice.Content)
					}
				}
			}

			section, _ = children[1].(*domain.Section)

			if assert.NotNil(section) {
				assert.Equal("Titre 1", section.Title)
			}

			section, _ = children[2].(*domain.Section)

			if assert.NotNil(section) {
				assert.Equal("Titre 2", section.Title)
			}
		}
	})
}