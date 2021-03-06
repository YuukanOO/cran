package assembleenationale

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"cran/pkg/generating"

	"github.com/stretchr/testify/assert"
)

func TestProviderAccept(t *testing.T) {
	p := provider{}

	assert.False(t, p.Accept("http://some.url/report"))
	assert.True(t, p.Accept(accept+"/a_report"))
}

func TestProviderName(t *testing.T) {
	p := provider{}

	assert.Equal(t, name, p.Name())
}

func TestProviderContentParsing(t *testing.T) {
	assert := assert.New(t)
	ts := newTestServer()
	defer ts.Close()
	p := &provider{}

	p.Fetch(ts.URL, func(report *generating.Report, err error) {
		assert.NotNil(report, "Report should not be nil")
		assert.Equal("Titre de la séance", report.Title, "Titles should match")

		children := report.Children()

		if assert.Len(children, 3, "It should have 3 sections") {
			section, _ := children[0].(*generating.Section)

			if assert.NotNil(section) {
				assert.Equal("Civilité présidentposte président", section.Title)
				assert.Equal(1, section.Level)
				assert.Equal("section-1", section.ID())

				sectionChildren := section.Children()

				if assert.Len(sectionChildren, 2, "Section should have 2 children") {
					intervention, _ := sectionChildren[0].(*generating.Intervention)

					if assert.NotNil(intervention) {
						assert.Equal("John Doe", intervention.SpeakerID)
						assert.Equal("La séance est ouverte.", intervention.Content)
						assert.Equal("intervention-2", intervention.ID())
					}

					notice, _ := sectionChildren[1].(*generating.Notice)

					if assert.NotNil(notice) {
						assert.Equal("<i>(date ouverture séance)</i>", notice.Content)
						assert.Equal("notice-3", notice.ID())
					}
				}
			}

			section, _ = children[1].(*generating.Section)

			if assert.NotNil(section) {
				assert.Equal("Titre 1", section.Title)
				assert.Equal(1, section.Level)
				assert.Equal("section-4", section.ID())

				sectionChildren := section.Children()

				if assert.Len(sectionChildren, 1, "First section should have 1 child") {
					section, _ = sectionChildren[0].(*generating.Section)

					if assert.NotNil(section, "It should have a subsection") {
						assert.Equal("Sous Titre 1", section.Title)
						assert.Equal(2, section.Level)
						assert.Equal("section-5", section.ID())

						sectionChildren := section.Children()

						if assert.Len(sectionChildren, 2, "Sub section should have 2 children") {
							intervention, _ := sectionChildren[0].(*generating.Intervention)

							if assert.NotNil(intervention, "One intervention") {
								assert.Equal("John Doe", intervention.SpeakerID)
								assert.Equal("La parole est à Jean Dupont.", intervention.Content)
								assert.Equal("intervention-6", intervention.ID())
							}

							section, _ := sectionChildren[1].(*generating.Section)

							if assert.NotNil(section, "And one section") {
								assert.Equal("Titre 99", section.Title)
								assert.Equal(3, section.Level)
								assert.Equal("section-7", section.ID())

								sectionChildren := section.Children()

								if assert.Len(sectionChildren, 2, "Last section should have 2 interventions") {
									intervention, _ := sectionChildren[0].(*generating.Intervention)

									if assert.NotNil(intervention, "First intervention") {
										assert.Equal("John Doe", intervention.SpeakerID)
										assert.Equal("La séance est suspendue.", intervention.Content)
										assert.Equal("intervention-8", intervention.ID())

									}

									intervention, _ = sectionChildren[1].(*generating.Intervention)

									if assert.NotNil(intervention, "Last intervention") {
										assert.Equal("Jean Dupont", intervention.SpeakerID)
										assert.Equal("J&#39;ai des choses à dire !", intervention.Content)
										assert.Equal("intervention-9", intervention.ID())
									}
								}
							}
						}
					}
				}
			}

			section, _ = children[2].(*generating.Section)

			if assert.NotNil(section) {
				assert.Equal("Titre 2", section.Title)
				assert.Equal(1, section.Level)
				assert.Equal("section-10", section.ID())
			}
		}
	})
}

func TestProviderSpeakersParsing(t *testing.T) {
	assert := assert.New(t)
	ts := newTestServer()
	defer ts.Close()
	p := &provider{}

	p.Fetch(ts.URL, func(report *generating.Report, err error) {
		if assert.Len(report.Speakers, 2, "It should contains 2 speakers") {
			speaker := report.Speakers["John Doe"]

			if assert.NotNil(speaker, "It should contains John Doe profile") {
				assert.Equal("John Doe", speaker.ID)
				assert.Equal("M. John Doe", speaker.Name)
				assert.Equal(ts.URL+"/fiche/1", speaker.ProfileURL)
				assert.Equal(ts.URL+"/photos/1.jpg", speaker.PictureURL)
				assert.Equal("Doubs\u00a0(5e circonscription)", speaker.Location)
				assert.Equal("Les Républicains", speaker.Side)
			}

			speaker = report.Speakers["Jean Dupont"]

			if assert.NotNil(speaker, "It should contains Jean Dupont profile") {
				assert.Equal("Jean Dupont", speaker.ID)
				assert.Equal("M. Jean Dupont", speaker.Name)
				assert.Equal(ts.URL+"/fiche/2", speaker.ProfileURL)
				assert.Equal(ts.URL+"/photos/2.jpg", speaker.PictureURL)
				assert.Equal("Bouches-du-Rhône\u00a0(4e circonscription)", speaker.Location)
				assert.Equal("Les Républicains", speaker.Side)
			}
		}
	})
}

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/fiche/1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="fr" lang="fr">
<head>
</head>
<body>
		<div id="page-fond">
			<div id="page">
				<div>
					<div id="deputes-fiche">
						<section>
							<div id="haut-contenu-page">
								<article>
									<div class="titre-bandeau-bleu clearfix">
										<h1>M. John Doe</h1>
										<p class="deputy-healine-sub-title">Doubs&nbsp;(5<sup>e</sup> circonscription)</p>
										<p class="deputy-healine-sub-title orange">Mandat en cours</p>
									</div>
									<div class="contenu-principal en-direct-commission clearfix">
										<div class="interieur-contenu-principal">
											<div id="deputes-illustration">
												<div class="deputes-image">
													<img src="/photos/1.jpg">
												</div>
												<span><a title="Accédez à la composition du groupe" href="/compo">Les Républicains</a></span>
											</div>
										</div>
									</div>
								</article>
							</div>
						</section>
					</div>
				</div>
			</div>
		</div>
</body>
</html>`))
	})

	mux.HandleFunc("/fiche/2", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!DOCTYPE html>
		<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="fr" lang="fr">
		<head>
		</head>
		<body>
				<div id="page-fond">
					<div id="page">
						<div>
							<div id="deputes-fiche">
								<section>
									<div id="haut-contenu-page">
										<article>
											<div class="titre-bandeau-bleu clearfix">
												<h1>M. Jean Dupont</h1>
												<p class="deputy-healine-sub-title">Bouches-du-Rhône&nbsp;(4<sup>e</sup> circonscription)</p>
												<p class="deputy-healine-sub-title orange">Mandat en cours</p>
											</div>
											<div class="contenu-principal en-direct-commission clearfix">
												<div class="interieur-contenu-principal">
													<div id="deputes-illustration">
														<div class="deputes-image">
															<img src="/photos/2.jpg">
														</div>
														<span class="pres-groupe">Président du groupe</span>															
														<span><a title="Accédez à la composition du groupe" href="/compo">Les Républicains</a></span>
													</div>
												</div>
											</div>
										</article>
									</div>
								</section>
							</div>
						</div>
					</div>
				</div>
		</body>
		</html>`))
	})

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
					<a href="` + server.URL + `/fiche/1">John Doe</a>
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
						<a href="` + server.URL + `/fiche/1">John Doe</a>						
					</b>.  La parole est à Jean Dupont.
				</p>
			</div>

			<div class="Point">
				<h2 class="titre99">
					<i>Titre 99</i>
				</h2>
				<p>
					<b>
						<a href="` + server.URL + `/fiche/1">John Doe</a>						
					</b>.  La séance est suspendue.
				</p>
			</div>
			<p>
				<b>
					<a href="` + server.URL + `/fiche/2">Jean Dupont</a>
				</b>.  J'ai des choses à dire !
			</p>
		</div>

		<div class="Point">
			<h5 class="numencad">2</h5>
			<h2 class="titre1">Titre 2<i></i></h2>
			<p>
				<b>
					<a href="` + server.URL + `/fiche/2>Jean Dupont</a>
				</b>.  Et encore d'autres ici !
			</p>
		</div>
	</div>
</div>
</body>
</html>`))
	})

	return server
}
