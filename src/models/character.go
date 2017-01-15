package models

import ()

type Character struct {
	id               int
	imageUrl         string
	name             string
	descriptionLines []string
}

func (this *Character) ImageUrl() string {
	return this.imageUrl
}

func (this *Character) Name() string {
	return this.name
}

func (this *Character) DescriptionLines() []string {
	return this.descriptionLines
}

func (this *Character) Id() int {
	return this.id
}

func (this *Character) SetImageUrl(value string) {
	this.imageUrl = value
}

func (this *Character) SetName(value string) {
	this.name = value
}

func (this *Character) SetDescriptionLines(value []string) {
	this.descriptionLines = value
}

func (this *Character) SetId(value int) {
	this.id = value
}

func GetCharacters() []Character {
	result := []Character{
		Character{
			id:       1,
			imageUrl: "weerwolf",
			name:     "Gewone Weerwolf",
			descriptionLines: []string{
				"Elke nacht worden alle weerwolven op teken van de spelleider wakker en wijzen gezamenlijk een slachtoffer aan.",
				"De weerwolven winnen het spel wanneer alle niet-weerwolven zijn uitgeschakeld"},
		},
		Character{
			id:       2,
			imageUrl: "grote_boze_wolf",
			name:     "De Grote Boze Wolf",
			descriptionLines: []string{
				"De Grote Boze Wolf wordt elke nacht samen met de andere weerwolven wakker en kiest met de rest van de weerwolven een slachtoffer.",
				"Vervolgens wordt de Grote Boze Wolf nog een keer alleen wakker en maakt hij een tweede slachtoffer.",
				"Zodra een van de weerwolven (inclusief de kleine wilde en de wolfshond) sterft verliest de Grote Boze Wolf zijn gave en speelt hij als een normale weerwolf verder.",
				"De Grote Boze Wolf is een weerwolf en wint samen met de weerwolven"},
		},
		Character{
			id:       3,
			imageUrl: "besmettelijke_oerwolf",
			name:     "De Besmettelijke Oerwolf",
			descriptionLines: []string{
				"De Besmettelijke Oerwolf wordt elke nacht samen met de andere weerwolven wakker en kiest met de rest van de weerwolven een slachtoffer.",
				"Wanneer de wolven weer gaan slapen mag de besmettelijke oerwolf zijn poot opsteken. In plaats van dat het weerwolven slachtoffer sterft wordt deze nu besmet.",
				"De besmette speler krijgt van de spelleider een klopje op zijn schouder. De besmette speler verandert vanaf de volgende nacht in een weerwolf en blijft dit tot de rest van het spel. Als de speler speciale krachten heeft, heeft hij deze nog steeds.",
				"De Besmettelijke Oerwolf is een weerwolf en wint samen met de weerwolven"},
		},
		Character{
			id:       4,
			imageUrl: "dief",
			name:     "De Dief",
			descriptionLines: []string{
				"Wanneer de Dief in het spel zit worden er twee karakter kaarten meer in het spel gedaan dan het aantal spelers. Bij het uitdelen van de karakter kaarten blijven er dus twee kaarten over.",
				"De Dief wordt in de eerste nacht door de spelleider opgeroepen en mag één van de twee overgebleven kaarten kiezen. Vanaf nu speelt de Dief zijn gekozen karakter.",
				"De Dief mag ook besluiten om geen kaart te kiezen en zijn Dief kaart te houden. Hij speelt dan als Dief verder in plaats van het gekozen karakter.",
				"Wanneer allebij de overgebleven twee weerwolven zijn moet de Dief een van deze kaarten kiezen",
				"De Dief is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       5,
			imageUrl: "toegewijde_dienstmeisje",
			name:     "Het Toegewijde dienstmeisje",
			descriptionLines: []string{
				"Het Toegewijde Dienstmeisje mag vlak voor de kaart van het slachtoffer van de burgers bekend gemaakt wordt zeggen dat ze het dienstmeisje is. Zei krijgt nu de identiteit van de net gestorven speler en speelt met deze identiteit verder. Haar nieuwe identiteit wordt aan niemand bekend gemaakt.",
				"Het Toegewijde Dienstmeisje is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld, mits ze van karakter veranderd, dan is haar doel gelijk aan haar nieuwe karakter.",
			},
		},
		Character{
			id:       6,
			imageUrl: "toneelspeler",
			name:     "De Toneelspeler",
			descriptionLines: []string{
				"Wanneer de Toneelspeler in het spel zit kiest de spelleider drie karakters (geen weerwolven) uit met een speciale gave. De kaarten van deze karakters worden open op tafel gelegd.",
				"Elke nacht wordt de Toneelspeler wakker en mag hij één van de op tafel liggende rollen vervullen. Hij speelt deze rol tot de volgende nacht.",
				"Wanneer de Toneelspeler een rol heeft vervuld, wordt dit karakter van de tafel verwijderd en kan hij dit karakter niet meer spelen.",
				"Als de Toneelspeler elke rol gespeeld heeft, wordt hij 's nachts niet meer opgeroepen.",
				"De Toneelspeler is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       7,
			imageUrl: "kleine_wilde",
			name:     "De Kleine Wilde",
			descriptionLines: []string{
				"In de eerste nacht wordt de Kleine Wilde opgeroepen en kiest hij een speler. Deze speler is vanaf nu zijn grote voorbeeld.",
				"Zolang zijn grote voorbeeld leeft is de kleine Wilde een burger.",
				"Als het grote voorbeeld van de Kleine Wilde sterft wordt de Kleine Wilde een weerwolf en wordt hij vanaf nu elke nacht met de weerwolven wakker.",
				"Totdat de Kleine Wilde een weerwolf is en 's nachts met de weerwolven mee speelt weten de weerwolven niet welke speler de Kleine Wilde speelt.",
				"Is De Kleine Wilde een burger dan wint hij samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
				"Is De Kleine Wilde een weerwolf dan wint hij samen met de weerwolven wanneer alle niet-weerwolven zijn uitgeschakeld.",
			},
		},
		Character{
			id:       8,
			imageUrl: "wolfshond",
			name:     "De Wolfshond",
			descriptionLines: []string{
				"De Wolfshond mag aan het begin van het spel kiezen of hij met de weerwolven of met de burgers mee speelt.",
				"Als de Wolfshond kiest om verder te spelen als weerwolf doet hij in de eerste nacht samen met de weerolf zijn ogen open.",
				"Als de Wolfshond kiest om te spelen als burger dan houdt hij wanneer de weerwolven wakker worden zijn ogen gesloten.",
				"De Wolfshond speelt voor de rest van het spel met zijn gekozen team mee.",
				"Is de Wolfshond een burger dan wint hij samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
				"Is de wolfshond een weerwolf dan wint hij samen met de weerwolven wanneer alle niet-weerwolven zijn uitgeschakeld.",
			},
		},
		Character{
			id:       9,
			imageUrl: "witte_weerwolf",
			name:     "De Witte Weerwolf",
			descriptionLines: []string{
				"De Witte Weerwolf wordt elke nacht samen met de andere weerwolven wakker en maakt met de rest van de weerwolven een slachtoffer.",
				"Om de nacht (alleen de even nachten) wordt de Witte Weerwolven een extra keer opgeroepen en mag hij een weerwolf vermoorden",
				"De Witte Weerwolf is niet verplicht om een weerwolven slachtoffer te maken",
				"De Witte Weerwolf wint alleen het spel als hij als laatste overblijft.",
			},
		},
		Character{
			id:       10,
			imageUrl: "engel",
			name:     "De Engel",
			descriptionLines: []string{
				"Wanneer de Engel in het spel zit begint het spel met een dag in plaats van een nacht.",
				"Het doel van de engel is om de eerste dag of de eerste nacht het leven te laten. Als dit lukt heeft De Engel per direct het spel gewonnen.",
				"Wanneer de Engel de eerste dag en nacht overleefd speel hij verder als een burger en wint hij samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       11,
			imageUrl: "fluitspeler",
			name:     "De Fluitspeler",
			descriptionLines: []string{
				"De Fluitspeler wordt elke nacht door de spelleider opgeroepen en mag twee spelers aanwijzen. De aangewezen spelers zijn vanaf nu betoverd",
				"De Fluitspeler sluit de ogen en de twee betoverde spelers worden aangetikt door de spelleider. Vervolgens mogen deze twee nieuw betoverde spelers hun ogen openen. Ze zien elkaar en gaan daarna weer slapen.",
				"De Fluitspeler mag zichzelf niet betoveren.",
				"De Fluitspeler wint het spel wanneer alle overgebleven spelers betoverd zijn.",
			},
		},
		Character{
			id:       12,
			imageUrl: "verschrikkelijke_sektarier",
			name:     "De Verschrikkelijke Sektariër",
			descriptionLines: []string{
				"Wanneer de Verschrikkelijke Sektariër in het spel zit worden de spelers van te voren verdeeld in twee groepen.",
				"De verdeling gebeurt op basis van een specifiek kenmerk (bijv. man/vrouw of wel/geen bril). Dit kenmerk is voor alle spelers bekend.",
				"De Verschrikkelijke Sektariër behoort net als de rest van de spelers tot één van deze twee groepen.",
				"De Verschrikkelijke Sektariër wint het spel wanneer er alleen nog spelers over zijn van de groep waar hij zelf in zit.",
			},
		},
		Character{
			id:       13,
			imageUrl: "gewone_burger",
			name:     "Gewone Burger",
			descriptionLines: []string{
				"De Gewone Burger heeft geen bijzondere gave.",
				"De burgers winnen het spel wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       14,
			imageUrl: "zieneres",
			name:     "De Ziener",
			descriptionLines: []string{
				"De Ziener wordt elke nacht op teken van de spelleider wakker en wijst iemand aan. De spelleider laat vervolgens de kaart van de gekozen persoon aan de Ziener zien.",
				"De Ziener is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld",
			},
		},
		Character{
			id:       15,
			imageUrl: "cupido",
			name:     "Cupido",
			descriptionLines: []string{
				"Alleen aan het begin van de eerste nacht wordt Cupido wakker.",
				"Cupido wijst twee spelers aan. Deze twee spelers zijn vanaf dat moment de geliefden en zijn smoorverliefd op elkaar.",
				"Cupido mag ook zichzelf aanwijzen om zo een van de geliefden te worden.",
				"Nadat Cupido gaat slapen maakt de spelleider de geliefden wakker, zodat de geliefden weten wie de andere geliefde is.",
				"Wanneer een van de geliefden sterft, sterft de andere geliefde direct aan liefdesverdriet.",
				"Een geliefde mag nooit tegen zijn geliefde stemmen.",
				"Wanneer de geliefden niet in hetzelfde team zitten (bijvoorbeeld een burger en een weerwolf/fluitspeler/sektariër) winnen de geliefden alleen het spel als ze alle andere spelers weten te elimineren",
				"Cupido is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       16,
			imageUrl: "heks",
			name:     "De Heks",
			descriptionLines: []string{
				"Wanneer de heks 's snachts opgeroepen wordt, wijst de spelleider aan de heks aan wie de weerwolven als slachtoffer hebben gekozen.",
				"De heks heef twee toverdrankjes. Een vergif en een levenselixer",
				"Met het vergif kan ze 's nachts een speler naar keuze elimineren.",
				"Met het levenselixer kan ze tijdens de nacht het weerwolven slachtoffer van die nacht tot leven wekken.",
				"De heks mag het levenselixer ook op zichzelf gebruiken.",
				"De heks mag beide drankjes niet op dezelfde speler gebruiken.",
				"De heks mag beide drankjes in dezelfde nacht gebruiken.",
				"Beide drankjes mogen maar één keer per spel gebruikt worden.",
				"De heks is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       17,
			imageUrl: "jager",
			name:     "De Jager",
			descriptionLines: []string{
				"Wanneer de jager sterft mag hij onmiddelijk een mede speler naar keuze elimineren.",
				"Wanneer de jager de laatst levende speler neerschiet, wint niemand had spel.",
				"De jager is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       18,
			imageUrl: "onschuldige_meisje",
			name:     "De Slet",
			descriptionLines: []string{
				"De slet wijst elke nacht een speler aan om vervolgens bij deze speler te slapen.",
				"Wanneer de slet door de weerwolven wordt aangewezen is ze niet thuis en overleeft ze de nacht.",
				"Wanneer de speler bij wie de slet slaapt wordt aangewezen door de weerwolven, sterven zowel de aangewezen speler als de slet.",
				"De slet mag niet twee nachten achter elkaar bij dezelfde speler slapen.",
				"De slet kan niet beschermd worden door de genezer of gered worden door de heks.",
				"De slet kan wel door de heks worden geëlimineerd.",
				"De slet is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       19,
			imageUrl: "genezer",
			name:     "De Genezer",
			descriptionLines: []string{
				"Elke nacht wordt de genezer wakker en wijst een speler aan. Deze speler is voor die nacht beschermd tegen de weerwolven.",
				"De genezer mag niet twee nachten achter elkaar dezelfde speler beschermen.",
				"De genezer mag ook zichzelf beschermen.",
				"De bescherming van de genezer beschermt niet tegen de fluitspeler of tegen de besmettelijke oerwolf.",
				"De Genezer is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       20,
			imageUrl: "dorpsoudste",
			name:     "De Dorpsoudste",
			descriptionLines: []string{
				"De dorpsoudste heeft uitsluitend een extra leven tegen de weerwolven en kan dus één keer door de weerwolven aangevallen worden zonder te sterven.",
				"Wanneer de dorpsoudste vermoord wordt door de heks, de jager of overdag gelynchet wordt door zijn mede-burgers, verliezen alle burgers direct al hun speciale krachten.",
				"Wanneer de heks de dorpsoudste beschermt tegen de weerwolven krijgt hij nooit meer dan één leven terug.",
				"De Dorpsoudste is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       21,
			imageUrl: "zondebok",
			name:     "De Zondebok",
			descriptionLines: []string{
				"Als het stemmen van de dorpsbewoners met een gelijke stand eindigt, sterft de Zondebok. De dagelijkse stemming eindigt hierbij ook direct.",
				"Wanneer de Zondebok sterft op welke wijze dan ook, mag hij eenmalig beslissen wie er de eerstvolgende stemronde wel en wie niet mogen stemmen.",
				"Het kan zijn dat bij de volgende stemronde iedereen die nog mocht stemmen is overleden. In dit geval vind er geen stemming plaats.",
				"De Zondebok is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       22,
			imageUrl: "dorpsgek",
			name:     "De Dorpsgek",
			descriptionLines: []string{
				"Wanneer de dorpsgek tijdens de stemronde door de burgers als schuldige wordt aangewezen, wordt de indentiteit van de dorpsgek door de spelleider aan iedereen bekend gemaakt. De burgers besluiten de Dorpsgek niet te vermoorden. De Dorpsgek doet dus nog steeds mee aan het spel, maar hij verliest wel zijn stemrecht.",
				"De Dorpsgek kan op elke andere wijze nog steeds geëlimineerd worden.",
				"Als de identiteit van de dorpsgek openbaar is, mag hij geen burgemeester (meer) zijn.",
				"De Dorpsgek is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       23,
			imageUrl: "burgemeester",
			name:     "De Burgemeester",
			descriptionLines: []string{
				"De Burgemeester wordt aangesteld door middel van een extra stemming die overdag plaats vindt.",
				"De Burgemeester is een aanvulling op een rol en is dus geen vervanging van een karakter.",
				"De Burgemeester heeft bij de dagelijkse stemming een doorslaggevende, oftewel anderhalve, stem.",
				"Wanneer de Burgemeester sterft kiest hij zelf zonder verder overleg zijn opvolger uit.",
			},
		},
		Character{
			id:       24,
			imageUrl: "twee_gezusters",
			name:     "De Twee Gezusters",
			descriptionLines: []string{
				"Beide Gezusters worden de eerste nacht opgeroepen, zodat ze van elkaar weten dat ze de Twee Gezusters zijn.",
				"De Twee Gezusters zijn burgers en winnen samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       25,
			imageUrl: "drie_gebroeders",
			name:     "De Drie Gebroeders",
			descriptionLines: []string{
				"Alle drie de Gebroeders worden de eerste nacht opgeroepen, zodat ze van elkaar weten dat ze de Drie Gebroeders zijn.",
				"De Drie Gebroeders zijn burgers en winnen samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       26,
			imageUrl: "vos",
			name:     "De Vos",
			descriptionLines: []string{
				"Elke nacht kiest de Vos een speler. Deze speler vormt samen met zijn buren een groep van drie.",
				"Als er zich in deze groep minstend één weerwolf bevindt, dan steekt de spelleider zijn duim op.",
				"Als er zich in deze groep geen weerwolven bevinden, dan schudt de spelleider zijn hoofd en verliest de Vos voor de rest van het spel zijn gave.",
				"De Vos is niet verplicht om in de nacht zijn gave te gebruiken.",
				"De Vos is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       27,
			imageUrl: "titus_en_zijn_dansende_beer",
			name:     "Titus en zijn Dansende Beer",
			descriptionLines: []string{
				"Elke ochtend gromt de beer van Titus wanneer Titus de buurman is van minstends één weerwolf.",
				"Spelers die uit het spel zijn worden niet in acht genomen.",
				"Als Titus besmet is door de Besmettelijke Oerwolf zal de beer ook grommen.",
				"Titus is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       28,
			imageUrl: "stotterende_raadsheer",
			name:     "De Stotterende Raadsheer",
			descriptionLines: []string{
				"De Stotterende Raadsheer wordt in de eesrte nacht wakker om aan de spelleider een zelf bedacht teken te laten zien.",
				"Wanneer de Sotterende Raadsheer tijdens de dagelijks stemming dit teken geeft aan de spelleider, vindt er direct na deze stemming een tweede stemming plaats om een tweede speler te elimineren. Dit mag hij maar één keer in het hele spel doen.",
				"De Stotterende Raadsheer is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       29,
			imageUrl: "ridder_met_het_roestige_zwaard",
			name:     "De Ridder met het Roestige Zwaard",
			descriptionLines: []string{
				"Als bij de bekendmaking van de slachtoffers in de nacht de Ridder met het Roestige Zwaard is gestorven, gaat de eerste weerwolf aan de linker hand van de Ridder de volgende nacht dood.",
				"De Ridder met het Roestige Zwaard is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
		Character{
			id:       30,
			imageUrl: "zigeunerin",
			name:     "De Zigeunerin",
			descriptionLines: []string{
				"Wanneer de Zigeunerin in het spel zit houdt de spelleider de 5 \"Spiritisme\" kaarten uit de volle maan uitbereiding bij zich.",
				"Iedere nacht mag de Zigeunerin haar gave gebruiken.",
				"Als de Zigeneurin haar gave wil gebruiken leest de spelleider de 4 vragen op één van de Spiritisme kaarten voor.",
				"De Zigeunerin kiest zonder zichzelf prijs te geven één van de vragen uit en kiest een speler aan die deze vraag moet stellen.",
				"De volgende morgen wordt de door de Zigeunerin gekozen speler bezeten en stelt hij de door de Zigeunerin uitgekozen vraag.",
				"De speler die als is eerste is gestorven verschijnt als geest en beantwoordt de gestelde vraag.",
				"De Zigeunerin is een burger en wint samen met de burgers wanneer alle niet-burgers zijn uitgeschakeld.",
			},
		},
	}
	return result
}
