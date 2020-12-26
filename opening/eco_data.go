package opening

var ecoData = []byte(`Code,Name,PGN
A00,Anderssen's Opening,1.a2a3
A00,"Polish Gambit, Anderssen's Opening",1.a2a3 a7a5 2.b2b4
A00,Creepy Crawly Formation,1.a2a3 e7e5 2.h2h3 d7d5
A00,Andersspike,1.a2a3 g7g6 2.g2g4
A00,Ware; Meadow Hay; Crab,1.a2a4
A00,"Wing Gambit, Ware Opening",1.a2a4 b7b5
A00,"Cologne Gambit, Ware Opening",1.a2a4 b7b6 2.d2d4 d7d5 3.b1c3 b8d7
A00,Ware Gambit,1.a2a4 e7e5 2.a4a5 d7d5 3.e2e3 f7f5 4.a5a6
A00,Polish; Orangutan; Sokolsky; Hunt,1.b2b4
A00,"Birmingham Gambit, Polish",1.b2b4 c7c5
A00,"Schuhler Gambit, Polish",1.b2b4 c7c6 2.c1b2 a7a5 3.b4b5 c6b5 4.e2e4
A00,"Myers Variation, Polish",1.b2b4 d7d5 2.c1b2 c7c6 3.a2a4
A00,"Bugayev Attack, Polish",1.b2b4 e7e5 2.a2a3
A00,"Wolferts Gambit, Polish",1.b2b4 e7e5 2.c1b2 c7c5
A00,Schiffler-Sokolsky; Tartakower Gambit,1.b2b4 e7e5 2.c1b2 f7f6 3.e2e4 f8b4
A00,"Brinckmann Variation, Polish",1.b2b4 e7e5 2.c1b2 f7f6 3.e2e4 f8b4 4.f1c4 b8c6 5.f2f4 d8e7 6.f4f5 g7g6
A00,"Bucker Defense, Polish",1.b2b4 b8a6
A00,"Grigorian Variation, Polish",1.b2b4 b8c6
A00,Polish Spike,1.b2b4 g8f6 2.c1b2 g7g6 3.g2g4
A00,"Karniewski; Tubingen Variation, Polish",1.b2b4 g8h6
A00,Saragossa; Hempel's Opening,1.c2c3
A00,Hanham; Hayward,1.c2c3 e7e5 2.d1c2
A00,Venezolana,1.d2d3 c7c5 2.b1c3 b8c6 3.g2g3
A00,Reversed Rat,1.d2d3 e7e5
A00,Spike Deferred; Myers Spike Attack,1.d2d3 g7g6 2.g2g4
A00,Keoni-Hiva Gambit,1.e2e3 d7d5 2.b1c3 g8f6 3.a2a3 e7e5 4.f2f4 e5f4 5.g1f3
A00,Van't Kruijs Opening,1.e2e3 e7e5
A00,Bouncing Bishop Variation,1.e2e3 e7e5 2.f1c4 b7b5 3.c4b3
A00,Amsterdam Attack,1.e2e3 e7e5 2.c2c4 d7d6 3.b1c3 b8c6 4.b2b3 g8f6
A00,Ekolu Variation,1.e2e3 e7e5 2.b1c3 d7d5 3.f2f4 e5f4 4.g1f3
A00,Alua Variation,1.e2e3 e7e5 2.b1c3 b8c6 3.f2f4 e5f4 4.g1f3
A00,Akahi Variation,1.e2e3 e7e5 2.b1c3 g8f6 3.f2f4 e5f4 4.g1f3
A00,King's Head Opening,1.f2f3 e7e5 2.e2e4
A00,Hammerschlag; Pork Chop Opening,1.f2f3 e7e5 2.e1f2
A00,Gambit 59,1.g2g3 d7d5 2.f1g2 e7e5 3.b2b4
A00,Regan Variation,1.g2g3 d7d5 2.f1g2 e7e5 3.d2d4
A00,Reversed Alekhine,1.g2g3 e7e5 2.g1f3
A00,Lasker Simul Special,1.g2g3 h7h5
A00,Grob; Ahlhausen; Colibri; Fric; Genoa; Spike,1.g2g4
A00,Grob Gambit,1.g2g4 d7d5 2.f1g2 c8g4
A00,"Romford Counter-Gambit, Grob",1.g2g4 d7d5 2.f1g2 c8g4 3.c2c4 d5d4
A00,"Spike, Grob",1.g2g4 d7d5 2.f1g2 c7c6 3.g4g5
A00,"Short Spike, Grob",1.g2g4 d7d5 2.f1g2 c7c6 3.h2h3
A00,"Zilbermints Gambit, Grob",1.g2g4 d7d5 2.e2e4
A00,"Keene Defense, Grob",1.g2g4 d7d5 2.h2h3 e7e5
A00,"Hurst Attack, Grob",1.g2g4 e7e5 2.f1g2 d7d5 3.c2c4
A00,"Alessi Gambit, Grob",1.g2g4 f7f5
A00,Double Grob,1.g2g4 g7g5
A00,Spike Deferred,1.g2g4 g7g6 2.d2d3
A00,"Bucker Defense, Grob",1.g2g4 g8h6
A00,Clemenz Opening,1.h2h3
A00,Kadas Opening,1.h2h4
A00,Despres Opening,1.h2h4
A00,Koola-Koola,1.h2h4 a7a5
A00,Wulumulu,1.h2h4 e7e5 2.d2d4
A00,Durkin's Opening; Sodium Attack,1.b1a3
A00,Dunst; Heinrichsen; Will Bull; Kotrc,1.b1c3
A00,"St. George Defense, Dunst",1.b1c3 a7a6
A00,"Novosibirsk Variation, Dunst",1.b1c3 c7c5 2.d2d4 c5d4 3.d1d4 b8c6 4.d4h4
A00,"Twyble Attack, Dunst",1.b1c3 c7c5 2.a1b1
A00,Dunst Gambit,1.b1c3 d7d5 2.e2e4 d5e4 3.d2d3
A00,"Pseudo-Diemer Gambit, Dunst",1.b1c3 d7d5 2.e2e4 d5e4 3.f2f3 e4f3 4.d1f3
A00,"Amazon, Dunst",1.b1c3 d7d5 2.e2e4 d5e4 3.c3e4 d8d5
A00,"Battambang, Dunst",1.b1c3 e7e5 2.a2a3
A00,Irish Attack; Polish Attack,1.b1c3 e7e5 2.f2f4
A00,"Napoleon Attack, Dunst",1.b1c3 e7e5 2.g1f3
A00,"Glasscoe Gambit, Dunst",1.b1c3 f7f5 2.g2g4
A00,"Anti-Pirc Variation, Dunst",1.b1c3 g7g6 2.h2h4
A00,"Tubingen Gambit, Dunst",1.b1c3 g8f6 2.g2g4
A00,Amar; Paris Opening,1.g1h3
A00,Drunken Knight Opening,1.g1h3
A00,Krazy Kat,1.g1h3 e7e5 2.f2f3 d7d5 3.h3f2
A01,Larsen; Queen's Fianchetto; Nimzo-Larsen Attack A01,1.b2b3
A01,"Symmetrical Variation, Larsen",1.b2b3 b7b6
A01,"English Variation, Larsen",1.b2b3 c7c5
A01,"Classical Variation, Larsen",1.b2b3 d7d5
A01,"Modern Variation, Larsen",1.b2b3 e7e5
A01,"Ringelbach Gambit, Larsen",1.b2b3 e7e5 2.c1b2 f7f5 3.e2e4
A01,"Paschmann Gambit, Larsen",1.b2b3 e7e5 2.c1b2 b8c6 3.f2f4
A01,"Dutch Variation, Larsen",1.b2b3 f7f5
A01,"Spike Variation, Larsen",1.b2b3 g8f6 2.c1b2 g7g6 3.g2g4
A02,ï¿½Bird's Opening; Dutch Attack,1.f2f4
A02,"From's Gambit, Bird",1.f2f4 e7e5
A02,Canard Formation; Double Duck,1.f2f4 f7f5 2.d2d4 d7d5
A02,"Pawn Thrust Attack, Bird",1.f2f4 f7f5 2.d2d4 d7d5
A02,"Swiss Gambit, Bird",1.f2f4 f7f5 2.e2e4 f5e4 3.b1c3 g8f6 4.g2g4
A02,"Hobbs Gambit, Bird",1.f2f4 g7g5
A02,"Batavo Polish Attack, Bird",1.f2f4 g8f6 2.g1f3 g7g6 3.b2b4
A03,A03,1.f2f4 d7d5
A03,Sturm; Mujannahï¿½,1.f2f4 d7d5 2.c2c4
A03,"Prokofiev Gambit, Bird",1.f2f4 d7d5 2.e2e4 d5e4 3.d2d3
A03,"Williams Gambit, Bird",1.f2f4 d7d5 2.e2e4 d5e4 3.b1c3
A03,"Horsefly Defense, Bird",1.f2f4 g8h6
A04,Wade Defense,1.g1f3 d7d6
A04,Polish G; Dalesio G; Pirc-Lisitsin G,1.g1f3 f7f5 2.e2e4 f5e4 3.f3g5
A04,"Herrstrom Gambit, Reti",1.g1f3 g7g5
A04,Oberndorfer Gambit; Reti,1.g1f3 g7g5
A04,Lessing Defense,1.g1f3 b8c6
A05,King's Indian Attack; Reti; Zukertort; A05,1.g1f3 g8f6
A06,A06,1.g1f3 d7d5
A06,"Nimzovich Opening, Reti",1.g1f3 d7d5 2.b2b3
A06,"Bloodgood Gambit, Reti",1.g1f3 d7d5 2.b2b3 c7c5 3.e2e4 d5e4 4.f3e5
A06,"Santasiere's Folly, Reti",1.g1f3 d7d5 2.b2b4
A06,Old Indian Attack,1.g1f3 d7d5 2.d2d3
A06,Tennison G; Lemberg G; Abonyi G,1.g1f3 d7d5 2.e2e4
A06,"Reversed Mexican Defense, Reti",1.g1f3 d7d5 2.b1c3
A06,"Ampel Variation, Reti",1.g1f3 d7d5 2.h1g1
A06,"Improved Lisitsin G, Reti",1.g1f3 f7f5 2.d2d3
A07,"Barcza Opening, Reti; A07",1.g1f3 d7d5 2.g2g3
A08,A08,1.g1f3 d7d5 2.g2g3 c7c5
A09,Reti Polonaise,1.g1f3 d7d5 2.c2c4 d5d4 3.b2b4
A09,"Penguin Variation, Reti",1.g1f3 d7d5 2.c2c4 d5d4 3.h1g1
A09,Reti Gambit Accepted,1.g1f3 d7d5 2.c2c4 d5c4
A09,"Neo-Catalan, Reti",1.g1f3 d7d5 2.c2c4 e7e6 3.g2g3
A10,English Opening; A10,1.c2c4
A10,"Halibut Gambit, English",1.c2c4 b7b5
A10,"Anglo Scandanavian Defense, English",1.c2c4 d7d5
A10,"Loehn Gambit, English",1.c2c4 d7d5 2.c4d5 e7e6
A10,"Schulz Gambit, English",1.c2c4 d7d5 2.c4d5 g8f6
A10,"Malvinas Variation, English",1.c2c4 d7d5 2.c4d5 d8d5 3.b1c3 d5a5
A10,Anglo-Dutch Defense,1.c2c4 f7f5
A10,Anglo-Polish Dutch,1.c2c4 f7f5 2.b2b4
A10,"Hickmann Gambit, English",1.c2c4 f7f5 2.e2e4
A10,"Hickmann Gambit, English",1.c2c4 f7f5 2.e2e4 f5e4 3.d2d3
A10,"Porcupine Variation, English",1.c2c4 f7f5 2.e2e4 f5e4 3.b1c3 g8f6 4.g2g4
A10,"Wade Gambit, English",1.c2c4 f7f5 2.g2g4
A10,English Spike,1.c2c4 f7f5 2.g2g4
A10,"Chabanon Gambit, English",1.c2c4 f7f5 2.g1f3 d7d6 3.e2e4
A10,"Myers Defense, English",1.c2c4 g7g5 2.d2d4 f8g7
A10,"Zilbermints Gambit, English",1.c2c4 g7g5 2.d2d4 e7e5
A10,"Adorjan System, English",1.c2c4 g7g6 2.e2e4 e7e5
A10,Beefeater,1.c2c4 g7g6 2.b1c3 f8g7 3.d2d4 c7c5 4.d4d5 g7c3 5.b2c3 f7f5
A10,"Anglo-Lithuanian Variation, English",1.c2c4 b8c6
A11,"Anglo-Slav, English; A11",1.c2c4 c7c6
A12,A12,1.c2c4 c7c6 2.g1f3 d7d5 3.b2b3
A12,"Bled Variation, English",1.c2c4 c7c6 2.g1f3 d7d5 3.b2b3 g8f6 4.g2g3 g7g6
A13,"Agincourt; Kurajica Defense, English; A13",1.c2c4 e7e6
A13,English Defense,1.c2c4 e7e6 2.d2d4 b7b6
A13,"King's Knight Variation, English",1.c2c4 e7e6 2.g1f3 d7d5
A13,"Wimpy System, English",1.c2c4 e7e6 2.g1f3 d7d5 3.b2b3 g8f6 4.c1b2 c7c5 5.e2e3
A13,"Afincourt Variation, English",1.c2c4 e7e6 2.g1f3 d7d5 3.g2g3
A13,"Kurajica Defense, English",1.c2c4 e7e6 2.g1f3 d7d5 3.g2g3 c7c6
A13,"Neo Catalan, English",1.c2c4 e7e6 2.g1f3 d7d5 3.g2g3 g8f6
A13,"Neo Catalan Accepted, English",1.c2c4 e7e6 2.g1f3 d7d5 3.g2g3 g8f6 4.f1g2 d5c4
A13,"Romanishin Gambit, English",1.c2c4 e7e6 2.g1f3 g8f6 3.g2g3 a7a6 4.f1g2 b7b5
A14,A14,1.c2c4 e7e6 2.g1f3 d7d5 3.g2g3 g8f6 4.f1g2 f8e7
A14,"Neo Catalan Declined, English",1.c2c4 e7e6 2.g1f3 d7d5 3.g2g3 g8f6 4.f1g2 f8e7 5.e1g1
A14,"Keres Defense, English",1.c2c4 e7e6 2.g1f3 d7d5 3.g2g3 g8f6 4.f1g2 f8e7 5.e1g1 c7c5 6.c4d5 f6d5 7.b1c3 b8c6
A15,"Anglo System, English; A15",1.c2c4 g8f6
A15,"Mikenas-Carls Variation, English",1.c2c4 g8f6
A15,English Orangutan,1.c2c4 g8f6 2.b2b4
A15,Shirazi Shove,1.c2c4 g8f6 2.g2g3 h7h5
A15,"King's Knight Variation, English",1.c2c4 g8f6 2.g1f3
A15,"Romanishin Variation, English",1.c2c4 g8f6 2.g1f3 e7e6 3.g2g3 a7a6
A16,A16,1.c2c4 g8f6 2.b1c3
A16,"Golombek Defense, English; Anti English",1.c2c4 g8f6 2.b1c3 d7d5
A16,"Czech Variation, English",1.c2c4 g8f6 2.b1c3 d7d5 3.c4d5 f6d5 4.g2g3 g7g6 5.f1g2 d5b6
A16,"Korchnoi's Variation, English",1.c2c4 g8f6 2.b1c3 d7d5 3.c4d5 f6d5 4.g1f3 g7g6 5.g2g3 f8g7 6.f1g2 e7e5
A16,"Stein Attack, English",1.c2c4 g8f6 2.b1c3 d7d5 3.c4d5 f6d5 4.g1f3 g7g6 5.d1a4
A16,Anti-anti-grunfeld,1.c2c4 g8f6 2.b1c3 g7g6 3.g1f3 f8g7 4.e2e4
A17,"Smyslov Defense, English",1.c2c4 g8f6 2.b1c3 d7d5 3.c4d5 f6d5 4.g2g3 g7g6 5.f1g2 d5c3
A17,"Hegehog System, Anglo-Indian English; A17",1.c2c4 g8f6 2.b1c3 e7e6
A17,Nimzo-English Opening,1.c2c4 g8f6 2.b1c3 e7e6 3.g1f3 f8b4
A17,"Zviagintsev-Krasenkov Attack, English",1.c2c4 g8f6 2.b1c3 e7e6 3.g1f3 f8b4 4.g2g4
A18,"Flohr-Mikenas-Carls Variation, English; A18",1.c2c4 g8f6 2.b1c3 e7e6 3.e2e4
A18,"Kevitz Defense, English",1.c2c4 g8f6 2.b1c3 e7e6 3.e2e4 b8c6
A18,"Queen's Indian Formation, English",1.c2c4 g8f6 2.b1c3 e7e6 3.g1f3 b7b6
A19,"Floh-Mikenas-Carls Variation, English; A19",1.c2c4 g8f6 2.b1c3 e7e6 3.e2e4 c7c5
A19,"Nei Gambit, English",1.c2c4 g8f6 2.b1c3 e7e6 3.e2e4 c7c5 4.e4e5 f6g8
A20,"Kahiko-Hula Gambit, English",1.c2c4 e7e5 2.e2e3 g8f6 3.f2f4 e5f4 4.g1f3
A20,"Double Whammy, English",1.c2c4 e7e5 2.f2f4
A20,"Drill Variation, English",1.c2c4 e7e5 2.g2g3 h7h5
A20,"Nimzovich Variation, English",1.c2c4 e7e5 2.g1f3
A20,"Nimzovich-Flohr Variation, English",1.c2c4 e7e5 2.g1f3 e5e4
A20,English Queen,1.c2c4 e7e5 2.d1a4
A21,"Reversed Sicilian, English; A21",1.c2c4 e7e5 2.b1c3
A21,"Kramnik-Shirov Variation, English",1.c2c4 e7e5 2.b1c3 f8b4
A21,"Troger Variation, English",1.c2c4 e7e5 2.b1c3 d7d6 3.g2g3 c8e6 4.f1g2 b8c6
A21,"Keres Defense, English",1.c2c4 e7e5 2.b1c3 d7d6 3.g2g3 c7c6
A21,"General Variation, King's English",1.c2c4 e7e5 2.b1c3 d7d6 3.g1f3
A21,"Smyslov Variation, English",1.c2c4 e7e5 2.b1c3 d7d6 3.g1f3 c8g4
A21,Anglo-Dutch,1.c2c4 e7e5 2.b1c3 f7f5
A22,"Two Knights Variation, English; A22",1.c2c4 e7e5 2.b1c3 g8f6
A22,"Mazedonisch Variation, English",1.c2c4 e7e5 2.b1c3 g8f6 3.f2f4
A22,"Bremen; Carls Variation, English",1.c2c4 e7e5 2.b1c3 g8f6 3.g2g3
A22,"Kapengut; Smyslov; Modern V, English",1.c2c4 e7e5 2.b1c3 g8f6 3.g2g3 f8b4
A22,"Bellon Gambit, English",1.c2c4 e7e5 2.b1c3 g8f6 3.g1f3 e5e4 4.f3g5 b7b5
A22,"Erbenheimer Gambit, English",1.c2c4 e7e5 2.b1c3 g8f6 3.g1f3 e5e4 4.f3g5 f6g4
A23,"Keres Variation, English; A23",1.c2c4 e7e5 2.b1c3 g8f6 3.g2g3 c7c6
A23,"Reversed Dragon, English",1.c2c4 e7e5 2.b1c3 g8f6 3.g2g3 d7d5
A24,"Fianchetto Lines, Two Knights Variation, English; A24",1.c2c4 e7e5 2.b1c3 g8f6 3.g2g3 g7g6
A24,"Prickly Pawn Pass System, English",1.c2c4 e7e5 2.b1c3 g8f6 3.g2g3 g7g6 4.f1g2 f8g7 5.d2d3 d7d6 6.e2e4 e8g8 7.g1e2 c7c6 8.e1g1 a7a6
A25,"Reversed Closed Sicilian, English; A25",1.c2c4 e7e5 2.b1c3 b8c6
A25,"Taimanov Variation, English",1.c2c4 e7e5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7
A25,"Hungarian Attack, English",1.c2c4 e7e5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.a1b1
A25,"Closed System, Full Symmetry, King's English",1.c2c4 e7e5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.d2d3 d7d6
A25,"Botvinnik System, King's English",1.c2c4 e7e5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.d2d3 d7d6 6.e2e4
A25,"Bremen-Hort Variation, English",1.c2c4 e7e5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.e2e3 d7d6 6.g1e2 c8e6
A25,"Botvinnik Variation, English",1.c2c4 e7e5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.e2e4
A26,"Closed System, King's English; A26",1.c2c4 e7e5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.d2d3
A27,"Three Knights System, English; A27",1.c2c4 e7e5 2.b1c3 b8c6 3.g1f3
A28,"Four Knights Variation, English; A28",1.c2c4 e7e5 2.b1c3 b8c6 3.g1f3 g8f6
A28,"Korchnoi Line, English",1.c2c4 e7e5 2.b1c3 b8c6 3.g1f3 g8f6 4.a2a3
A28,"Flexible Line, English",1.c2c4 e7e5 2.b1c3 b8c6 3.g1f3 g8f6 4.d2d3
A28,"Bradley Beach Variation, English",1.c2c4 e7e5 2.b1c3 b8c6 3.g1f3 g8f6 4.d2d4 e5e4
A28,"Quiet Line, English",1.c2c4 e7e5 2.b1c3 b8c6 3.g1f3 g8f6 4.e2e3
A28,"Romanishin Variation, English",1.c2c4 e7e5 2.b1c3 b8c6 3.g1f3 g8f6 4.e2e3 f8b4 5.d1c2 b4c3
A28,"Keene; Stean Variation, English",1.c2c4 e7e5 2.b1c3 b8c6 3.g1f3 g8f6 4.e2e3 f8b4 5.d1c2 e8g8 6.c3d5 f8e8 7.c2f5
A28,"Botvinnik Line, English",1.c2c4 e7e5 2.b1c3 b8c6 3.g1f3 g8f6 4.e2e4
A29,"Fianchetto Line, Four Knights English; A29",1.c2c4 e7e5 2.b1c3 b8c6 3.g1f3 g8f6 4.g2g3
A30,"Wing Variation, English",1.c2c4 c7c5 2.b2b4
A30,"Napolitano Gambit, Symmetrical English",1.c2c4 c7c5 2.g1f3 g8f6 3.b2b4
A30,"Hedgehog Variation, Symmetrical English",1.c2c4 c7c5 2.g1f3 g8f6 3.g2g3 b7b6 4.f1g2 c8b7 5.e1g1 e7e6 6.b1c3 f8e7
A30,"Double Fianchetto Defense, English",1.c2c4 c7c5 2.g1f3 g8f6 3.g2g3 b7b6 4.f1g2 c8b7 5.e1g1 g7g6
A31,"Anti-Benoni Variation, Symmetrical English; A31",1.c2c4 c7c5 2.g1f3 g8f6 3.d2d4
A31,"Two Knights Variation, English",1.c2c4 c7c5 2.g1f3 g8f6 3.d2d4 c5d4 4.f3d4
A32,"Spielmann Defense, Symmetrical English; A32",1.c2c4 c7c5 2.g1f3 g8f6 3.d2d4 c5d4 4.f3d4 e7e6
A33,A33,1.c2c4 c7c5 2.g1f3 g8f6 3.d2d4 c5d4 4.f3d4 e7e6 5.b1c3 b8c6
A33,"Geller Variation, Symmetrical English",1.c2c4 c7c5 2.g1f3 g8f6 3.d2d4 c5d4 4.f3d4 e7e6 5.b1c3 b8c6 6.g2g3 d8b6
A34,"Normal Variation, Symmetrical English; A34",1.c2c4 c7c5 2.b1c3
A34,"Fianchetto Variation, Symmetrical Englishï¿½ï¿½",1.c2c4 c7c5 2.b1c3 g8f6 3.g2g3
A34,"Rubinstein Variation, Symmetrical English",1.c2c4 c7c5 2.b1c3 g8f6 3.g2g3 d7d5 4.c4d5 f6d5 5.f1g2 d5c7
A34,"Asymmetrical Variation, English",1.c2c4 c7c5 2.b1c3 g8f6 3.g1f3 d7d5 4.c4d5 f6d5
A34,"Keres-Parma System, English",1.c2c4 c7c5 2.b1c3 g8f6 3.g1f3 e7e6 4.g2g3 b8c6
A36,"Fianchetto Variation, Symmetrical English; A36",1.c2c4 c7c5 2.b1c3 b8c6 3.g2g3
A36,"Symmetrical Variation, Symmetrical English",1.c2c4 c7c5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7
A36,"Botvinnik System Reversed, Symmetrical English",1.c2c4 c7c5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.e2e3 e7e5
A36,"Botvinnik Variation, Symmetrical English",1.c2c4 c7c5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.e2e4
A37,"Two Knights Line, Symmetrical English; A37",1.c2c4 c7c5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.g1f3
A38,"Full Symmetry Line, Symmetrical English; A38",1.c2c4 c7c5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.g1f3 g8f6
A38,"Double Fianchetto, Symmetrical English",1.c2c4 c7c5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.g1f3 g8f6 6.e1g1 e8g8 7.b2b3
A38,"Duchamp Variation, Symmetrical English",1.c2c4 c7c5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.g1f3 g8f6 6.e1g1 e8g8 7.d2d3
A39,"Mecking Variation, Symmetrical English; A39",1.c2c4 c7c5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.g1f3 g8f6 6.e1g1 e8g8 7.d2d4
A40,"Hartlaub Gambit, English",1.c2c4 b7b6 2.d2d4 c8b7 3.b1c3 e7e6 4.e2e4 f7f5 5.e4f5 g8f6
A40,Queen Pawn Opening; A40,1.d2d4
A40,St George's Gambit,1.d2d4 a7a6 2.c2c4 b7b5 3.e2e4 e7e6 4.c4b5 a6b5
A40,"Spassky Gambit, Polish Defense",1.d2d4 b7b5 2.e2e4 c8b7 3.f1b5
A40,"Queen Fianchetto, Queen Pawn Game",1.d2d4 b7b6
A40,Charlick Gambit; Englund Gambit,1.d2d4 e7e5
A40,Englund Gambit Declined,1.d2d4 e7e5 2.d4d5
A40,Diemer Counterattack,1.d2d4 e7e5 2.d4d5 f8c5 3.e2e4 d8h4
A40,Thuring Gambit; Hartlaub-Charlick Gambit,1.d2d4 e7e5 2.d4e5 d7d6
A40,Soller; Purchas Gambit,1.d2d4 e7e5 2.d4e5 f7f6
A40,Felbecker Gambit,1.d2d4 e7e5 2.d4e5 b8c6 3.g1f3 f8c5
A40,Soller Gambit Deferred,1.d2d4 e7e5 2.d4e5 b8c6 3.g1f3 f7f6
A40,"Zilbermints Gambit II, Englund Gambit",1.d2d4 e7e5 2.d4e5 b8c6 3.g1f3 h7h6
A40,"Zilbermints Gambit, Englund Gambit",1.d2d4 e7e5 2.d4e5 b8c6 3.g1f3 g8e7
A40,Englund Gambit,1.d2d4 e7e5 2.d4e5 b8c6 3.g1f3 d8e7
A40,"Stockhom Variation, Englund Gambit",1.d2d4 e7e5 2.d4e5 b8c6 3.g1f3 d8e7 4.d1d5
A40,"Mosquito Gambit, Englund",1.d2d4 e7e5 2.d4e5 d8h4
A40,Horwitz Defense,1.d2d4 e7e6
A40,Franco-Indian Defense,1.d2d4 e7e6
A40,"Seneschaud Variation, Dutch",1.d2d4 e7e6 2.c1f4 f7f5 3.g2g4
A40,English Defense,1.d2d4 e7e6 2.c2c4 b7b6
A40,Perrin Variation,1.d2d4 e7e6 2.c2c4 b7b6 3.e2e4 c8b7 4.f1d3 b8c6
A40,Poli Gambit,1.d2d4 e7e6 2.c2c4 b7b6 3.e2e4 c8b7 4.f2f3 f7f5 5.e4f5 g8h6
A40,Hartlaub Gambit Declined,1.d2d4 e7e6 2.c2c4 b7b6 3.b1c3 c8b7 4.e2e4 f7f5 5.d4d5
A40,Hartlaub Gambit Accepted,1.d2d4 e7e6 2.c2c4 b7b6 3.b1c3 c8b7 4.e2e4 f7f5 5.e4f5 g8f6
A40,"Keres Defense, Indian; Kangaroo Defense",1.d2d4 e7e6 2.c2c4 f8b4
A40,"Transpositional Variation, Keres Defense",1.d2d4 e7e6 2.c2c4 f8b4 3.b1c3
A40,Zilbermints Gambit,1.d2d4 e7e6 2.c2c4 e6e5
A40,Borg Gambit,1.d2d4 g7g5
A40,Queen Pawn Fianchetto,1.d2d4 g7g6
A40,Kotov-Robatsch Defense,1.d2d4 g7g6 2.c2c4 f8g7
A40,Pseudo-Saemisch,1.d2d4 g7g6 2.c2c4 f8g7 3.e2e4 d7d6 4.c1e3 g8f6 5.f2f3
A40,Rossolimo Variation,1.d2d4 g7g6 2.c2c4 f8g7 3.e2e4 d7d6 4.g1f3 c8g4
A40,Neo-Modern Defense,1.d2d4 g7g6 2.c2c4 f8g7 3.e2e4 e7e5
A40,Beefeater Variation,1.d2d4 g7g6 2.c2c4 f8g7 3.b1c3 c7c5 4.d4d5 g7c3 5.b2c3 f7f5
A40,Modern Defense,1.d2d4 g7g6 2.c2c4 f8g7 3.b1c3 d7d6
A40,"Kotov Variation, Modern",1.d2d4 g7g6 2.c2c4 f8g7 3.b1c3 d7d6 4.e2e4 b8c6
A40,Lizard Defense; Pirc-Diemer Gambit,1.d2d4 g7g6 2.h2h4 g8f6 3.h4h5
A40,Bogoljubov; Mikenas-Van Geet Defense,1.d2d4 b8c6
A40,Lundin Defense,1.d2d4 b8c6
A40,"Lithuanian Variation, Mikenas",1.d2d4 b8c6 2.c2c4 e7e5 3.d4d5 c6e7
A40,"Cannstatter Variation, Mikenas",1.d2d4 b8c6 2.c2c4 e7e5 3.d4d5 c6d4
A40,Pozarek Gambit,1.d2d4 b8c6 2.c2c4 e7e5 3.d4e5 c6e5 4.b1c3 e5c4
A40,Montevideo Defense,1.d2d4 b8c6 2.d4d5 c6b8
A40,Zaire Defense,1.d2d4 b8c6 2.d4d5 c6b8 3.e2e4 g8f6 4.e4e5 f6g8
A40,Polish Variation,1.e2e4 g7g6 2.d2d4 f8g7 3.c2c4 c7c5 4.g1f3 d8b6
A40,"Dunworth Variation, Modern Defense",1.e2e4 g7g6 2.d2d4 f8g7 3.c2c4 d7d5 4.e4d5 c7c6 5.d5c6 g7d4
A41,Anglo-Slav Opening,1.d2d4 c7c6 2.c2c4 d7d6
A41,Wade Defense,1.d2d4 d7d6 2.g1f3 c8g4
A41,"Averbakh Variation, Modern",1.d2d4 g7g6 2.c2c4 f8g7 3.b1c3 d7d6 4.e2e4
A41,"Pterodactyl Variation, Modern",1.d2d4 g7g6 2.c2c4 f8g7 3.b1c3 d7d6 4.e2e4 c7c5 5.g1f3 d8a5
A42,A42,1.d2d4 d7d6 2.c2c4
A42,English Rat,1.d2d4 d7d6 2.c2c4 e7e5
A42,Lisbon Gambit,1.d2d4 d7d6 2.c2c4 e7e5 3.d4e5 b8c6
A42,Randspringer Variation,1.d2d4 g7g6 2.c2c4 f8g7 3.b1c3 d7d6 4.e2e4 f7f5
A43,Staunton Defense; Benoni C-G; A43,1.d2d4 c7c5
A43,"Zilbermints Gambit; Nakamura Gambit, Benoni",1.d2d4 c7c5 2.b2b4 c5b4
A43,Old Benoni,1.d2d4 c7c5 2.d4d5 d7d6
A43,"Schmid Variation, Benoni",1.d2d4 c7c5 2.d4d5 d7d6 3.b1c3 g7g6
A43,Franco-Benoni; Franco-Sicilian,1.d2d4 c7c5 2.d4d5 e7e6 3.e2e4
A43,"Clarendon Court Variation, Benoni",1.d2d4 c7c5 2.d4d5 f7f5
A43,Benoni-Staunton Gambit,1.d2d4 c7c5 2.d4d5 f7f5 3.e2e4
A43,"Snail Variation, Benoni",1.d2d4 c7c5 2.d4d5 b8a6
A43,Benoni-Indian Defense,1.d2d4 c7c5 2.d4d5 g8f6
A43,"Woozle Defense, Benoni",1.d2d4 c7c5 2.d4d5 g8f6 3.b1c3 d8a5
A43,"Kingside move order, Benoni",1.d2d4 c7c5 2.d4d5 g8f6 3.g1f3
A43,"Habichd; Hawk Variation, Benoni",1.d2d4 c7c5 2.d4d5 g8f6 3.g1f3 c5c4
A43,Benoni Gambit Accepted,1.d2d4 c7c5 2.d4c5
A43,"Cormorant Gambit, Benoni",1.d2d4 c7c5 2.d4c5 b7b6
A43,"Schlenker Defense, Benoni",1.d2d4 c7c5 2.d4c5 b8a6
A43,Zilbermints-Benoni Gambit,1.d2d4 c7c5 2.g1f3 c5d4 3.b2b4
A43,"Tamarkin Countergambit, Benoni",1.d2d4 c7c5 2.g1f3 c5d4 3.b2b4 e7e5
A43,Franco-Sicilian Hybrid,1.d2d4 e7e6 2.c2c4 c7c5 3.d4d5 e6d5 4.c4d5 d7d6 5.b1c3 g7g6 6.e2e4 f8g7 7.g1f3 g8e7
A44,"Pawn Thrust, Benoni; A44",1.d2d4 c7c5 2.d4d5 e7e5
A44,Blockade Variation; Semi-Benoni; Russian Variation,1.d2d4 c7c5 2.d4d5 e7e5 3.e2e4 d7d6
A45,Indian Defense; A45,1.d2d4 g8f6
A45,Basque Opening,1.d2d4 g8f6 2.b2b3
A45,Trompowsky; Opocensky; Ruth Opening,1.d2d4 g8f6 2.c1g5
A45,"Poisoned Pawn Variation, Trompowsky",1.d2d4 g8f6 2.c1g5 c7c5 3.d4d5 d8b6 4.b1c3
A45,Chepukaitis Gambit,1.d2d4 g8f6 2.c1g5 c7c5 3.d4d5 d8b6 4.b1c3 b6b2 5.g5d2 b2b6
A45,"Borg Variation, Trompowsky",1.d2d4 g8f6 2.c1g5 f6e4 3.g5f4 g7g5
A45,"Edge Variation, Trompowsky",1.d2d4 g8f6 2.c1g5 f6e4 3.g5h4
A45,"Hergert Gambit #2, Trompowsky",1.d2d4 g8f6 2.c1g5 f6e4 3.g5h4 c7c6 4.b1d2 d8a5 5.c2c3 e4d2 6.d1d2 d7d5 7.e2e4
A45,"Hergert Gambit #1, Trompowsky",1.d2d4 g8f6 2.c1g5 f6e4 3.g5h4 d7d5 4.f2f3 e4f6 5.b1c3 c8f5 6.e2e4
A45,"Raptor Variation, Trompowsky",1.d2d4 g8f6 2.c1g5 f6e4 3.h2h4
A45,"Hergert Gambit, Trompowsky",1.d2d4 g8f6 2.c1g5 f6e4 3.h2h4 e4g5 4.h4g5 e7e5
A45,Pawn Push Variation,1.d2d4 g8f6 2.d4d5
A45,Omega Gambit,1.d2d4 g8f6 2.e2e4
A45,Arafat Gambit,1.d2d4 g8f6 2.e2e4 f6e4 3.f1d3 e4f6 4.c1g5
A45,Paleface Attack,1.d2d4 g8f6 2.f2f3
A45,Webster Gambit,1.d2d4 g8f6 2.f2f3 c7c5 3.d4d5 e7e6 4.e2e4 e6d5 5.e4e5
A45,Kraus-Muhlherr Counter Gambit,1.d2d4 g8f6 2.f2f3 d7d5 3.e2e4 e7e5 4.d4e5 f6e4
A45,Gedult Attack,1.d2d4 g8f6 2.f2f3 d7d5 3.g2g4
A45,Gedult-Gunderam System,1.d2d4 g8f6 2.f2f3 d7d5 3.g2g4
A45,Linder Counter Gambit,1.d2d4 g8f6 2.f2f3 e7e6 3.e2e4 f6e4
A45,Canard Attack,1.d2d4 g8f6 2.f2f4
A45,Ueberlinger Gambit,1.d2d4 g8f6 2.g2g3 e7e5
A45,Bullfrog; Weidenhagen; Krabbe; Ochsenfrosch,1.d2d4 g8f6 2.g2g4
A45,Bronstein Gambit,1.d2d4 g8f6 2.g2g4
A45,Oshima Defense,1.d2d4 g8f6 2.g2g4 e7e5
A45,Gibbins-Wiedehagen Gambit Accepted,1.d2d4 g8f6 2.g2g4 f6g4
A45,Stummer Gambit,1.d2d4 g8f6 2.g2g4 f6g4 3.e2e4 d7d6 4.f1e2 g4f6 5.b1c3
A45,Maltese Falcon,1.d2d4 g8f6 2.g2g4 f6g4 3.f2f3 g4f6 4.e2e4
A45,Reversed Chigorin Defense,1.d2d4 g8f6 2.b1c3 c7c5
A45,Von Beivorseen,1.d2d4 g8f6 2.b1c3 c7c5 3.d4c5
A45,Malich Gambit,1.d2d4 g8f6 2.b1c3 d7d5 3.c1g5 c7c5 4.g5f6 g7f6 5.e2e4 d5e4 6.d4d5
A45,Huebsch Gambit,1.d2d4 g8f6 2.b1c3 d7d5 3.e2e4 f6e4
A45,Maddigan Gambit,1.d2d4 g8f6 2.b1c3 e7e5
A45,Lazard Gambit; Maddigan Gambit,1.d2d4 g8f6 2.b1d2 e7e5
A46,"Knights Variation, Indian Game; A46",1.d2d4 g8f6 2.g1f3
A46,Polish Defense,1.d2d4 g8f6 2.g1f3 b7b5
A46,Spielmann-Indian,1.d2d4 g8f6 2.g1f3 c7c5
A46,Pseudo-Benko Gambit,1.d2d4 g8f6 2.g1f3 c7c5 3.d4d5 b7b5
A46,Czech-Indian,1.d2d4 g8f6 2.g1f3 c7c6
A46,London System,1.d2d4 g8f6 2.g1f3 e7e6 3.c1f4
A46,Torre Attack,1.d2d4 g8f6 2.g1f3 e7e6 3.c1g5
A46,Petrosian Gambit,1.d2d4 g8f6 2.g1f3 e7e6 3.c1g5 c7c5 4.e2e3 b7b6 5.d4d5
A46,Swiss; Wagner Gambit,1.d2d4 g8f6 2.g1f3 e7e6 3.c1g5 c7c5 4.e2e4
A46,"Nimzovich Variation, Torre Attack",1.d2d4 g8f6 2.g1f3 e7e6 3.c1g5 h7h6
A46,Yusupov-Rubinstein System,1.d2d4 g8f6 2.g1f3 e7e6 3.e2e3
A46,Dory Defense,1.d2d4 g8f6 2.g1f3 f6e4
A47,"Capablanca Variation, Indian Game; A47",1.d2d4 g8f6 2.g1f3 b7b6
A47,Torre Attack,1.d2d4 g8f6 2.g1f3 b7b6 3.c1g5
A47,Schnepper Gambit,1.d2d4 g8f6 2.g1f3 b7b6 3.c2c3 e7e5
A47,Marienbad Variation,1.d2d4 g8f6 2.g1f3 b7b6 3.g2g3 c8b7 4.f1g2 c7c5
A47,Berg Variation,1.d2d4 g8f6 2.g1f3 b7b6 3.g2g3 c8b7 4.f1g2 c7c5 5.c2c4 c5d4 6.d1d4
A48,Pseudo-King's Indian Variation; A48,1.d2d4 g8f6 2.g1f3 g7g6
A48,"London System, Torre Attack",1.d2d4 g8f6 2.g1f3 g7g6 3.c1f4
A48,London Variation,1.d2d4 g8f6 2.g1f3 g7g6 3.c1f4 f8g7 4.b1d2
A48,"Fianchetto Defense, Torre Attack",1.d2d4 g8f6 2.g1f3 g7g6 3.c1g5
A48,"Euwe Variation, Torre Attack",1.d2d4 g8f6 2.g1f3 g7g6 3.c1g5 f8g7 4.b1d2 c7c5
A48,King's Indian Variation,1.d2d4 g8f6 2.g1f3 g7g6 3.e2e3 f8g7 4.f1d3 d7d6
A49,Przepiorka Variation; A49,1.d2d4 g8f6 2.g1f3 g7g6 3.g2g3
A50,"Normal Variation, Indian Game; A50",1.d2d4 g8f6 2.c2c4
A50,Pyrenees Gambit,1.d2d4 g8f6 2.c2c4 b7b5
A50,Queen's Indian Defense; Saemisch-Indian,1.d2d4 g8f6 2.c2c4 b7b6
A50,Medusa Gambit,1.d2d4 g8f6 2.c2c4 g7g5
A50,Van Geet; Kevitz; Trajkovic; Mexican,1.d2d4 g8f6 2.c2c4 b8c6
A50,Horsefly Gambit,1.d2d4 g8f6 2.c2c4 b8c6 3.d4d5 c6e5 4.f2f4
A51,Budapest Defense; A51,1.d2d4 g8f6 2.c2c4 e7e5
A51,"Fajarowicz Variation, Budapest",1.d2d4 g8f6 2.c2c4 e7e5 3.d4e5 f6e4
A51,"Bonsdorf Variation, Budapest",1.d2d4 g8f6 2.c2c4 e7e5 3.d4e5 f6e4 4.a2a3 b7b6
A51,"Steiner Variation, Budapest",1.d2d4 g8f6 2.c2c4 e7e5 3.d4e5 f6e4 4.d1c2
A52,"Normal Variation, Budapest; A52",1.d2d4 g8f6 2.c2c4 e7e5 3.d4e5 f6g4
A52,"Rubinstein Variation, Budapest",1.d2d4 g8f6 2.c2c4 e7e5 3.d4e5 f6g4 4.c1f4
A52,"Kornl Richter Gambit, Budapest",1.d2d4 g8f6 2.c2c4 e7e5 3.d4e5 f6g4 4.c1f4 b8c6 5.g1f3 f8b4 6.b1c3 b4c3
A52,"Center Play Variation, Budapest; Alekhine Variation",1.d2d4 g8f6 2.c2c4 e7e5 3.d4e5 f6g4 4.e2e4
A52,"Balogh; Tartakower C-G, Budapest",1.d2d4 g8f6 2.c2c4 e7e5 3.d4e5 f6g4 4.e2e4 d7d6
A52,"Barmina Gambit, Budapest",1.d2d4 g8f6 2.c2c4 e7e5 3.d4e5 f6g4 4.e2e4 g4e5 5.f2f4 f8c5
A52,"Wikstrom Gambit, Budapest",1.d2d4 g8f6 2.c2c4 e7e5 3.d4e5 f6g4 4.e2e4 g4e5 5.f2f4 b8c6
A52,"Abonyi Variation, Budapest",1.d2d4 g8f6 2.c2c4 e7e5 3.d4e5 f6g4 4.e2e4 g4e5 5.f2f4 e5c6
A52,"Adler Variation, Budapest",1.d2d4 g8f6 2.c2c4 e7e5 3.d4e5 f6g4 4.g1f3
A53,Old Indian Defense; Chigorin Indian Defense; A53,1.d2d4 g8f6 2.c2c4 d7d6
A53,Aged Gibbon Gambit,1.d2d4 g8f6 2.c2c4 d7d6 3.g2g4
A53,"Janowski Variation, Old Indian",1.d2d4 g8f6 2.c2c4 d7d6 3.b1c3 c8f5
A53,"Grinberg Gambit, Old Indian",1.d2d4 g8f6 2.c2c4 d7d6 3.b1c3 c8f5 4.e2e4
A53,Ukranian Variation,1.d2d4 g8f6 2.c2c4 d7d6 3.b1c3 e7e5
A53,Duz-Khotimirsky Variation,1.d2d4 g8f6 2.c2c4 d7d6 3.b1c3 b8d7 4.e2e3 e7e5 5.f1d3
A53,Tartakower-Indian,1.d2d4 g8f6 2.c2c4 d7d6 3.g1f3 c8g4
A54,"Two Knights Variation, Old Indian; A54",1.d2d4 g8f6 2.c2c4 d7d6 3.b1c3 e7e5 4.g1f3
A55,"Normal Variation, Old Indian; A55",1.d2d4 g8f6 2.c2c4 d7d6 3.b1c3 e7e5 4.g1f3 b8d7 5.e2e4
A56,"Vulture Defense, Benoni",1.d2d4 c7c5 2.d4d5 g8f6 3.c2c4 f6e4
A56,Hromadka Defense; Loose Gambit: A56,1.d2d4 g8f6 2.c2c4 c7c5
A56,"Hromadka System, Benoni",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 d7d6
A56,Lundin; Bronstein Counter-Gambit,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 d7d6 4.b1c3 g7g6 5.e2e4 b7b5
A56,Petrosian System; Czech Benoni Defense,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e5
A56,"Czech Variation, Benoni",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e5 4.b1c3 d7d6 5.e2e4 f8e7
A56,"King's Indian System, Benoni",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e5 4.b1c3 d7d6 5.e2e4 g7g6
A56,Vulture Variation; Der Geier,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 f6e4
A56,"Weenink Variation, Benoni",1.d2d4 g8f6 2.c2c4 c7c5 3.d4c5 e7e6
A57,Benko Gambit; Volga; Opocensky; Vaitonis; A57,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5
A57,"Sosonko Variation, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.a2a4
A57,"Bishop Attack, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.c1g5
A57,Benko Gambit Accepted,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.c4b5
A57,"Pawn Return Variation, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.c4b5 a7a6 5.b5b6
A57,"Central Storming Variation, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.c4b5 a7a6 5.b5a6 c8a6 6.b1c3 d7d6 7.f2f4
A57,"Modern Variation, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.c4b5 a7a6 5.e2e3
A57,"Dlugy Variation, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.c4b5 a7a6 5.f2f3
A57,"Nescafe Frappe Attack, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.c4b5 a7a6 5.b1c3 a6b5 6.e2e4 b5b4 7.c3b5
A57,"Zaitsev Variation, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.c4b5 a7a6 5.b1c3 a6b5 6.e2e4 b5b4 7.c3b5 d7d6 8.g1f3 g7g6 9.f1c4
A57,Benko Gambit Declined; Hjoerring Countergambit,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.e2e4
A57,"Pseudo-Saemisch, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.f2f3
A57,"Mutkin Countergambit, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.g2g4
A57,"Quiet Line, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.b1d2
A57,"Main Line, Benko Gambit Declined",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.g1f3
A58,"Fully Accepted Variation, Benko Gambit; A58",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.c4b5 a7a6 5.b5a6
A58,"Fianchetto Variation, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.c4b5 a7a6 5.b5a6 c8a6 6.b1c3 g7g6 7.g1f3 d7d6 8.g2g3 f8g7 9.f1g2
A59,A59,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.c4b5 a7a6 5.b5a6 c8a6 6.b1c3 d7d6 7.e2e4
A59,"King Walk Variation, Benko Gambit",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 b7b5 4.c4b5 a7a6 5.b5a6 c8a6 6.b1c3 d7d6 7.e2e4 a6f1 8.e1f1 g7g6 9.g2g3 f8g7 10.f1g2 e8g8
A60,"Tal Variation, Modern Variation, Benoni; A60",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6
A60,Modern Benoni,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6
A61,A61,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3
A61,"Snake Variation, Benoni",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 f8d6
A61,"Uhlmann Variation, Benoni",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.g1f3 g7g6 7.c1g5
A61,"Nimzovich Variation, Knight's Tour Variation, Benoni",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.g1f3 g7g6 7.f3d2
A62,A62,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.g1f3 g7g6 7.g2g3
A62,"Fianchetto Variation, Benoni",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.g1f3 g7g6 7.g2g3 f8g7 8.f1g2
A62,"Hastings Defense, Benoni",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.g1f3 g7g6 7.g2g3 f8g7 8.f1g2 e8g8 9.e1g1 b8d7
A63,A63,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.g1f3 g7g6 7.g2g3 f8g7 8.f1g2 e8g8 9.e1g1 a7a6
A64,A64,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.g1f3 g7g6 7.g2g3 f8g7 8.f1g2 e8g8 9.e1g1 a7a6 10.a2a4 b8d7 11.f3d2
A65,"King Pawn lines, Benoni; A65",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4
A66,"Pawn Storm Variation, Benoni; A66",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.f2f4
A66,"Mikenas V, Benoni",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.f2f4 f8g7 8.e4e5
A67,"Taimanov Variation, Benoni; A67",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.f2f4 f8g7 8.f1b5
A68,A68,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.f2f4 f8g7 8.g1f3
A69,A69,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.f2f4 f8g7 8.g1f3 e8g8 9.f1e2 f8e8
A70,"Classical Variation, Benoni; A70",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.g1f3
A71,A71,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.g1f3 f8g7 8.c1g5
A72,A72,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.g1f3 f8g7 8.f1e2
A73,A73,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.g1f3 f8g7 8.f1e2 e8g8 9.e1g1
A74,A74,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.g1f3 f8g7 8.f1e2 e8g8 9.e1g1 a7a6
A75,"Argentine Attack, Benoni; A75",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.g1f3 f8g7 8.f1e2 e8g8 9.e1g1 a7a6 10.a2a4 c8g4
A76,"Czerniak Defesne, Benoni; A76",1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.g1f3 f8g7 8.f1e2 e8g8 9.e1g1 f8e8
A77,A77,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.g1f3 f8g7 8.f1e2 e8g8 9.e1g1 f8e8 10.f3d2
A78,A78,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.g1f3 f8g7 8.f1e2 e8g8 9.e1g1 f8e8 10.f3d2 b8a6
A79,A79,1.d2d4 g8f6 2.c2c4 c7c5 3.d4d5 e7e6 4.b1c3 e6d5 5.c4d5 d7d6 6.e2e4 g7g6 7.g1f3 f8g7 8.f1e2 e8g8 9.e1g1 f8e8 10.f3d2 b8a6 11.f2f3
A80,Dutch Defense; A80,1.d2d4 f7f5
A80,"Seneschaud Gambit, Dutch",1.d2d4 f7f5 2.c1f4 e7e6 3.g2g4
A80,"Hopton Attack, Dutch",1.d2d4 f7f5 2.c1g5
A80,"Blackmar's Second Gambit, Dutch",1.d2d4 f7f5 2.e2e4 f5e4 3.b1c3 g8f6 4.f2f3
A80,"Lisbon Gambit, Dutch; Krejcik Gambit",1.d2d4 f7f5 2.g2g4
A80,"Hevendehl Gambit, Dutch",1.d2d4 f7f5 2.g2g4 e7e5
A80,"Tate Gambit, Dutch",1.d2d4 f7f5 2.g2g4 f5g4 3.e2e4 d7d5 4.b1c3
A80,"Korchnoi's Variation, Dutch",1.d2d4 f7f5 2.h2h3
A80,"Janzen-Korchnoi Gambit, Dutch",1.d2d4 f7f5 2.h2h3 g8f6 3.g2g4
A80,"Raphael Variation, Dutch",1.d2d4 f7f5 2.b1c3
A80,"Kingfisher Gambit, Dutch",1.d2d4 f7f5 2.b1c3 d7d5 3.e2e4
A80,"Spielmann Gambit, Dutch",1.d2d4 f7f5 2.b1c3 g8f6 3.g2g4
A80,"Barcza Variation, Dutch",1.d2d4 f7f5 2.g1f3 g8f6 3.c2c3
A80,"Anti-Stonewall, Dutch",1.d2d4 f7f5 2.d1d3 d7d5 3.g2g4
A80,"Anti-Modern, Dutch",1.d2d4 f7f5 2.d1d3 d7d6 3.g2g4
A80,"Anti-Classical Line, Dutch",1.d2d4 f7f5 2.d1d3 e7e6 3.g2g4
A80,"Anti-Leningrad Variation, Dutch",1.d2d4 f7f5 2.d1d3 g7g6 3.g2g4
A80,"Manhattan Gambit, Dutch",1.d2d4 f7f5 2.d1d3 g8f6 3.g2g4
A80,"Pretzel Gambit, Dutch",1.d2d4 f7f5 2.d1d3 g8f6 3.g2g4
A81,"Dutch Indian; Leningrad Variation, Dutch; A81",1.d2d4 f7f5 2.g2g3
A81,"Basman Variation, Dutch",1.d2d4 f7f5 2.g2g3 g7g6 3.f1g2 f8g7 4.g1f3 c7c6 5.e1g1 g8h6
A81,"Carlsbad Variation, Dutch",1.d2d4 f7f5 2.g2g3 g7g6 3.f1g2 f8g7 4.g1h3
A81,"Antoshin-Hort Variation, Dutch",1.d2d4 f7f5 2.g2g3 g8f6 3.f1g2 d7d6
A81,"Ilyin-Genevsky Variation, Dutch",1.d2d4 f7f5 2.g2g3 g8f6 3.f1g2 e7e6 4.g1f3 f8e7 5.e1g1 e8g8 6.c2c4 d7d6
A81,"Blackburne Variation, Dutch",1.d2d4 f7f5 2.g2g3 g8f6 3.f1g2 e7e6 4.g1h3
A81,"Semi-Leningrad Variation, Dutch",1.d2d4 f7f5 2.g2g3 g8f6 3.f1g2 g7g6
A82,"Balogh Defense, Dutch",1.d2d4 f7f5 2.e2e4 d7d6
A82,"Staunton Gambit Accepted, Dutch",1.d2d4 f7f5 2.e2e4 f5e4
A82,"Tartakower Variation, Dutch",1.d2d4 f7f5 2.e2e4 f5e4 3.b1c3 g8f6 4.g2g4
A82,"American Gambit, Dutch",1.d2d4 f7f5 2.e2e4 f5e4 3.b1d2
A83,"General Variation, Staunton Gambit, Dutch; A83",1.d2d4 f7f5 2.e2e4 f5e4 3.b1c3 g8f6 4.c1g5
A83,"Nimzovich Variation, Dutch",1.d2d4 f7f5 2.e2e4 f5e4 3.b1c3 g8f6 4.c1g5 b7b6
A83,"Chigorin Variation, Dutch",1.d2d4 f7f5 2.e2e4 f5e4 3.b1c3 g8f6 4.c1g5 c7c6
A83,"Lasker Variation, Dutch",1.d2d4 f7f5 2.e2e4 f5e4 3.b1c3 g8f6 4.c1g5 g7g6 5.f2f3
A84,"Slav Stonewall, Dutch",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c6 4.e2e3 f7f5
A84,Bellon Gambit,1.d2d4 e7e6 2.c2c4 f7f5 3.e2e4
A84,A84,1.d2d4 f7f5 2.c2c4
A84,"Classical Variation, Dutch",1.d2d4 f7f5 2.c2c4 e7e6
A84,"Bellon Gambit, Dutch",1.d2d4 f7f5 2.c2c4 e7e6 3.e2e4
A84,"Rubinstein Variation, Dutch",1.d2d4 f7f5 2.c2c4 e7e6 3.b1c3
A84,"Queen's Knight Variation, Dutch",1.d2d4 f7f5 2.c2c4 g8f6 3.b1c3
A84,"Krause Variation, Dutch",1.d2d4 f7f5 2.c2c4 g8f6 3.b1c3 d7d6 4.g1f3 b8c6
A85," ",1.d2d4 f7f5 2.c2c4 g7g6 3.b1c3
A85,"Bladel Variation, Dutch",1.d2d4 f7f5 2.c2c4 g7g6 3.b1c3 g8h6
A85,"Normal Variation, Dutch",1.d2d4 f7f5 2.c2c4 g8f6
A85,"Ozol Variation, Dutch",1.d2d4 f7f5 2.c2c4 g8f6 3.b1c3 g7g6 4.g1f3 f8g7 5.c1g5
A86,"Fianchetto Variation, Dutch; A86",1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3
A86,"Leningrad Variation, Dutch",1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 g7g6
A87,A87,1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 g7g6 4.f1g2 f8g7 5.g1f3
A87,"Fluid Formation, Dutch",1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 g7g6 4.f1g2 f8g7 5.g1f3 e8g8 6.e1g1 d7d6
A88,"Warsaw Variation, Dutch; A88",1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 g7g6 4.f1g2 f8g7 5.g1f3 e8g8 6.e1g1 d7d6 7.b1c3 c7c6
A89,"Matulovic Variation, Dutch; A89",1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 g7g6 4.f1g2 f8g7 5.g1f3 e8g8 6.e1g1 d7d6 7.b1c3 b8c6
A90,A90,1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6
A90,"Classical Variation, Dutchï¿½",1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2
A90,Nimzo-Dutch Variation,1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2 f8b4
A91,A91,1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2 f8e7
A91,"Botvinnik-Bronstein Variation, Dutch",1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2 f8e7 5.b1c3 e8g8 6.e2e3
A92,A92,1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2 f8e7 5.g1f3
A92,"Alekhine Variation, Dutch",1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 f6e4
A93,"Botvinnik Variation, Dutch; A93",1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 d7d5 7.b2b3
A94,A94,1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 d7d5 7.b2b3 c7c6
A95,A95,1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 d7d5 7.b1c3
A96,A96,1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 d7d6
A97,"Ilyin-Zhenevsky Variation, Dutch; A97",1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 d7d6 7.b1c3 d8e8
A98,A98,1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 d7d6 7.b1c3 d8e8 8.d1c2
A99,A99,1.d2d4 f7f5 2.c2c4 g8f6 3.g2g3 e7e6 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 d7d6 7.b1c3 d8e8 8.b2b3
B00,Ware Defense,1.e2e4 a7a5
B00,Corn Stalks,1.e2e4 a7a5
B00,Snagkeouss Defense,1.e2e4 a7a5 2.d2d4 b8c6
B00,San Jorge Variation,1.e2e4 a7a6 2.d2d4 b7b5 3.g1f3 c8b7 4.f1d3 d7d6 5.e1g1 g7g6 6.c2c3 f8g7
B00,"Polish Variation, St George Defense",1.e2e4 a7a6 2.d2d4 b7b5 3.g1f3 c8b7 4.f1d3 e7e6
B00,"Zilbermints Gambit, St George Defense",1.e2e4 a7a6 2.d2d4 e7e5
B00,San Jorge Variation,1.e2e4 a7a6 2.d2d4 e7e6
B00,"Three Pawn Attack, St George",1.e2e4 a7a6 2.d2d4 e7e6 3.c2c4
B00,"Sanky-Georg Gambit, St George",1.e2e4 a7a6 2.d2d4 e7e6 3.c2c4 b7b5
B00,"Woodchuck Variation, St George",1.e2e4 a7a6 2.d2d4 b8c6
B00,Owen's Defense; Greek Defense,1.e2e4 b7b6
B00,Guatemalan Defense,1.e2e4 b7b6 2.d2d4 c8a6
B00,Matovinsky Gambit,1.e2e4 b7b6 2.d2d4 c8b7 3.f1d3 f7f5
B00,Matovinsky Gambit,1.e2e4 b7b6 2.d2d4 c8b7 3.f1d3 f7f5 4.e4f5 b7g2 5.d1h5 g7g6
B00,"Naselwaus Variation, Owen's Defense",1.e2e4 b7b6 2.d2d4 c8b7 3.c1g5
B00,"Wind Gambit, Owen Defense",1.e2e4 b7b6 2.d2d4 c8b7 3.f2f3 e7e5
B00,"Smith Gambit, Owen Defense",1.e2e4 b7b6 2.d2d4 c8b7 3.g1f3
B00,"Hekili-Loa Gambit, Owen Defense",1.e2e4 b7b6 2.d2d4 c7c5 3.d4c5 b8c6
B00,Mao Tse Tung; Duras,1.e2e4 f7f5
B00,Fred; From Reversed,1.e2e4 f7f5
B00,Tiers Counter-Gambit,1.e2e4 f7f5
B00,Barnes Defense,1.e2e4 f7f6
B00,Basman Defense; Macho Grob,1.e2e4 g7g5
B00,Borg Gambit,1.e2e4 g7g5 2.d2d4 f8g7
B00,"Zilbermints Gambit, Borg",1.e2e4 g7g5 2.d2d4 e7e5
B00,Borg Defense; Troon Gambit,1.e2e4 g7g5 2.d2d4 h7h6 3.h2h4 g5g4
B00,Pickering (Goldsmith) Defense,1.e2e4 h7h5
B00,Picklepuss Defense,1.e2e4 h7h5 2.d2d4 g8f6
B00,Carr Defense,1.e2e4 h7h6
B00,"Zilbermints Gambit, Carr Defense",1.e2e4 h7h6 2.d2d4 e7e5
B00,Lemming Defense,1.e2e4 b8a6
B00,Nimzovich Defense,1.e2e4 b8c6
B00,"Wheeler Gambit, Nimzovich Defense",1.e2e4 b8c6 2.b2b4 c6b4 3.c2c3 b4c6 4.d2d4
B00,Pseudo-Spanish Variation,1.e2e4 b8c6 2.f1b5
B00,"Scandanavian Variation, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 d7d5
B00,"Hornung Gambit, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 d7d5 3.c1e3
B00,"Aachen Gambit, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 d7d5 3.e4d5 c6b4
B00,"Marshall Gambit, Nimzovich",1.e2e4 b8c6 2.d2d4 d7d5 3.e4d5 d8d5 4.b1c3
B00,"Bogoljubow Variation, Nizovich Defense",1.e2e4 b8c6 2.d2d4 d7d5 3.b1c3
B00,"Brandics Gambit, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 d7d5 3.b1c3 a7a6
B00,"Richter Gambit, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 d7d5 3.b1c3 d5e4 4.d4d5 c6b8 5.f2f3
B00,"Nimzovich Gambit, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 d7d5 3.b1c3 d5e4 4.d4d5 c6e5
B00,"Alekseev Variation, Nimzovich; Heinola-Deppe Gambit",1.e2e4 b8c6 2.d2d4 d7d5 3.b1c3 e7e5
B00,"Erben Gambit, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 d7d5 3.b1c3 g7g6
B00,"Vehre Variation, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 d7d5 3.b1c3 g8f6
B00,"Mikenas Variation, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 d7d6
B00,"Kennedy Variation, Nimzovich",1.e2e4 b8c6 2.d2d4 e7e5
B00,"Linksspringer Variation, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 e7e5 3.d4d5
B00,"Bileefelder Gambit, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 e7e5 3.d4e5 f8c5
B00,"Burgar Gambit, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 e7e5 3.d4e5 d7d6
B00,"Hammer Gambit, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 e7e5 3.d4e5 f7f6
B00,"Riemann Defense, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 e7e5 3.d4e5 c6e5 4.f2f4 e5c6
B00,"Keres Attack, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 e7e5 3.d4e5 c6e5 4.b1c3
B00,"Paulsen Attack, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 e7e5 3.d4e5 c6e5 4.g1f3
B00,"Herford Gambit, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 e7e5 3.d4e5 d8h4
B00,"French Connection, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 e7e6
B00,"Neo-Mongoloid Defense, Nimzovich",1.e2e4 b8c6 2.d2d4 f7f6
B00,"Pirc Connection, Nimzovich Defense",1.e2e4 b8c6 2.d2d4 g7g6
B00,"Breyer Defense, Nimzovich",1.e2e4 b8c6 2.b1c3 g8f6
B00,"Williams Variation, Nimzovich",1.e2e4 b8c6 2.g1f3 d7d6
B00,Franco-Nimzovich Variation,1.e2e4 b8c6 2.g1f3 e7e6
B00,"Colorado Defense, Nimzovich; Lean Variation",1.e2e4 b8c6 2.g1f3 f7f5
B00,"Colorado Counter Accepted, Nimzovich Defense",1.e2e4 b8c6 2.g1f3 f7f5 3.e4f5
B00,"El Columpio Defense, Nimzovich Defense",1.e2e4 b8c6 2.g1f3 g8f6 3.e4e5 f6g4
B00,Adams Defense,1.e2e4 g8h6
B00,Wild Bull Defense,1.e2e4 g8h6
B01,Center Counter; Scandanavian; B01,1.e2e4 d7d5
B01,"Zilbermints Gambit, Center Counter",1.e2e4 d7d5 2.b2b4
B01,"Blackbunre-Kloosterboer Gambit, Center Counter",1.e2e4 d7d5 2.e4d5 c7c6 3.d5c6 e7e5
B01,"Atkinson (Kloosterboer) Variation, Center Counter",1.e2e4 d7d5 2.e4d5 c7c6 3.d5c6 e7e5
B01,"Blackburne Gambit, Center Counter",1.e2e4 d7d5 2.e4d5 c7c6 3.d5c6 b8c6
B01,"Boehnke Gambit, Center Counter",1.e2e4 d7d5 2.e4d5 e7e5
B01,"Boehnke Gambit Accepted, Center Counter",1.e2e4 d7d5 2.e4d5 e7e5 3.d5e6 c8e6
B01,"Panov Transfer, Center Counter",1.e2e4 d7d5 2.e4d5 g8f6 3.c2c4 c7c6
B01,"Icelandic-Palme Gambit, CC",1.e2e4 d7d5 2.e4d5 g8f6 3.c2c4 e7e6
B01,"Portuguese Variation, Center Counter",1.e2e4 d7d5 2.e4d5 g8f6 3.d2d4 c8g4
B01,"Portuguese Gambit, Center Counter",1.e2e4 d7d5 2.e4d5 g8f6 3.d2d4 c8g4 4.f2f3 g4f5 5.f1b5 b8d7
B01,"Kadas Gambit, Center Counter",1.e2e4 d7d5 2.e4d5 g8f6 3.d2d4 c7c6 4.d5c6 e7e5
B01,"Wing Gambit, Center Counter",1.e2e4 d7d5 2.e4d5 g8f6 3.d2d4 g7g6 4.c2c4 b7b5
B01,"Marshall Variation, Center Counter",1.e2e4 d7d5 2.e4d5 g8f6 3.d2d4 f6d5
B01,"Kiel Variation, Center Counter",1.e2e4 d7d5 2.e4d5 g8f6 3.d2d4 f6d5 4.c2c4 d5b4
B01,"Tomalty Gambit, Center Counter",1.e2e4 d7d5 2.e4d5 g8f6 3.d2d4 f6d5 4.c2c4 d5f6 5.b1c3 c7c5
B01,"Gipslis Variation, Center Counter",1.e2e4 d7d5 2.e4d5 g8f6 3.d2d4 f6d5 4.g1f3 c8g4
B01,"Richter Variation, Center Counter",1.e2e4 d7d5 2.e4d5 g8f6 3.d2d4 f6d5 4.g1f3 g7g6
B01,"Mieses-Kotroc Variation, Center Counter",1.e2e4 d7d5 2.e4d5 d8d5
B01,"Leonhardt (Rosen) Gambit, Center Counter",1.e2e4 d7d5 2.e4d5 d8d5 3.b1c3 d5a5 4.b2b4
B01,"Anderssen Variation, Center Counter",1.e2e4 d7d5 2.e4d5 d8d5 3.b1c3 d5a5 4.d2d4 e7e5
B01,"Collijn Variation, Center Counter",1.e2e4 d7d5 2.e4d5 d8d5 3.b1c3 d5a5 4.d2d4 e7e5 5.g1f3 c8g4
B01,"Mieses Variation, Center Counter",1.e2e4 d7d5 2.e4d5 d8d5 3.b1c3 d5a5 4.d2d4 g8f6
B01,"Gruenfeld Variation, Center Counter",1.e2e4 d7d5 2.e4d5 d8d5 3.b1c3 d5a5 4.d2d4 g8f6 5.g1f3 c8f5 6.f3e5 c7c6 7.g2g4
B01,"Lasker Variation, Center Counter",1.e2e4 d7d5 2.e4d5 d8d5 3.b1c3 d5a5 4.d2d4 g8f6 5.g1f3 c8g4 6.h2h3
B01,"Gubinsky-Melts Defense, Center Counter",1.e2e4 d7d5 2.e4d5 d8d5 3.b1c3 d5d6
B01,"Schiller-Pytel Variation, Center Counter",1.e2e4 d7d5 2.e4d5 d8d5 3.b1c3 d5d6 4.d2d4 c7c6
B01,"Bronstein Variation, Center Counter",1.e2e4 d7d5 2.e4d5 d8d5 3.b1c3 d5d6 4.d2d4 g8f6 5.g1f3 a7a6
B01,Closed Scandavian,1.e2e4 d7d5 2.b1c3
B02,"Fried Fox, Alekhine; Krejcik Gambit",1.e2e4 g8f6 2.f1c4 f6e4 3.c4f7
B02,"Maroczy Variation, Alekhine",1.e2e4 g8f6 2.d2d3
B02,"Welling Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.b2b3
B02,"Kmoch Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.f1c4 d5b6 4.c4b3 c7c5 5.d2d3
B02,"Hunt Gambit, Alekhine; Two Pawn Attack",1.e2e4 g8f6 2.e4e5 f6d5 3.c2c4
B02,"Steiner Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.c2c4 d5b6 4.b2b3
B02,"Lasker Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.c2c4 d5b6 4.c4c5
B02,"Mikenas Gambit, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.c2c4 d5b6 4.c4c5 b6d5 5.f1c4 e7e6 6.b1c3 d7d6 7.c3d5 e6d5 8.c4d5
B02,"Matsukevich Gambit, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.c2c4 d5b6 4.c4c5 b6d5 5.b1c3 d5c3 6.d2c3 d7d6 7.c1g5
B02,"Buntin Gambit, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.e5e6
B02,"Buckley Attack, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.b1a3
B02,"Saemisch Attack, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.b1c3
B02,Mokele Mbembe; Knight's Tour,1.e2e4 g8f6 2.e4e5 f6e4
B02,"Retreat; Brooklyn Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6g8
B02,"Everglades Defense, Alekhine",1.e2e4 g8f6 2.e4e5 f6g8 3.d2d4 f7f5
B02,"Scandanavian Variation, Alekhine",1.e2e4 g8f6 2.b1c3 d7d5
B02,"Myers Gambit, Alekhine",1.e2e4 g8f6 2.b1c3 d7d5 3.d2d3 d5e4 4.c1g5
B02,"Spielmann Variation, Alekhine",1.e2e4 g8f6 2.b1c3 d7d5 3.e4e5 f6d7 4.e5e6
B02,"Noch Gambit, Alekhine",1.e2e4 g8f6 2.b1c3 d7d5 3.e4d5
B02,"Geschev Gambit, Alekhine",1.e2e4 g8f6 2.b1c3 d7d5 3.e4d5 c7c6
B02,"John Tracy Gambit, Alekhine",1.e2e4 g8f6 2.g1f3
B03,B03,1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4
B03,"Wall Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 c7c5
B03,"Balogh Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.f1c4
B03,"Hunt Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.c2c4 d5b6 5.c4c5
B03,Four Pawns Attack; Chase V,1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.c2c4 d5b6 5.f2f4
B03,"Trifunovic Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.c2c4 d5b6 5.f2f4 c8f5
B03,"Cambridge Gambit, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.c2c4 d5b6 5.f2f4 g7g5
B04,B04,1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.g1f3
B04,"Panov Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.g1f3 c8g4 5.h2h3
B04,"Larsen Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.g1f3 d6e5
B04,"Alburt Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.g1f3 g7g6
B04,"Keres Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.g1f3 g7g6 5.f1c4 d5b6 6.c4b3 f8g7 7.a2a4
B04,"Karpov Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.g1f3 g7g6 5.c2c4 d5b6 6.h2h3 f8g7 7.f1e2
B04,"Schmid Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.g1f3 d5b6
B05,B05,1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.g1f3 c8g4
B05,"Flohr Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.g1f3 c8g4 5.f1e2 c7c6
B05,"Alekhine Variation, Alekhine",1.e2e4 g8f6 2.e4e5 f6d5 3.d2d4 d7d6 4.g1f3 c8g4 5.c2c4
B06,Pirc; Robatsch; Modern; King Pawn Fianchetto; B06,1.e2e4 g7g6
B06,"Monkey's Bum, Pirc",1.e2e4 g7g6 2.f1c4 f8g7 3.d1f3 e7e6 4.d2d4
B06,Bishop Attack,1.e2e4 g7g6 2.d2d4 f8g7 3.f1c4
B06,Buecker Gambit,1.e2e4 g7g6 2.d2d4 f8g7 3.f1c4 b7b5
B06,Westermann Gambit,1.e2e4 g7g6 2.d2d4 f8g7 3.c1d2
B06,Wind Gambit,1.e2e4 g7g6 2.d2d4 f8g7 3.f1d3
B06,Three Pawns Attack,1.e2e4 g7g6 2.d2d4 f8g7 3.f2f4
B06,"Pterodactyl, Pirc",1.e2e4 g7g6 2.d2d4 f8g7 3.b1c3 c7c5
B06,"Gurgenidze Defense, Modern Defense",1.e2e4 g7g6 2.d2d4 f8g7 3.b1c3 c7c6 4.f2f4 d7d5 5.e4e5 h7h5
B06,"Mittenberger Gambit, Modern Defense",1.e2e4 g7g6 2.d2d4 f8g7 3.b1c3 d7d5
B06,Anti-Modern,1.e2e4 g7g6 2.d2d4 f8g7 3.b1c3 d7d6 4.f1c4 c7c6 5.d1e2
B06,Pseudo-Austrian Attack,1.e2e4 g7g6 2.d2d4 f8g7 3.b1c3 d7d6 4.f2f4
B06,"Two Knight Variation, Modern Defense",1.e2e4 g7g6 2.d2d4 f8g7 3.b1c3 d7d6 4.g1f3
B06,"Suttles Variation, Modern Defense",1.e2e4 g7g6 2.d2d4 f8g7 3.b1c3 d7d6 4.g1f3 c7c6
B06,"Mongredien Defense, Pirc",1.e2e4 g7g6 2.d2d4 f8g7 3.g1f3 b7b6
B06,"Geller's System, Modern Defense",1.e2e4 g7g6 2.d2d4 f8g7 3.g1f3 d7d6 4.c2c3
B06,"Rossolimo Variation, Modern Defense",1.e2e4 g7g6 2.d2d4 f8g7 3.g1f3 d7d6 4.c2c4 c8g4
B06,Accelerated Gurgenidze,1.e2e4 g7g6 2.d2d4 f8g7 3.g1f3 d7d6 4.b1c3 c7c6
B06,"Small Center Variation, Pirc",1.e2e4 g7g6 2.d2d4 f8g7 3.g1f3 d7d6 4.b1c3 e7e6
B06,"Fianchetto Gambit, Modern Defense",1.e2e4 g7g6 2.d2d4 f7f5
B06,Norwegian Defense,1.e2e4 g7g6 2.d2d4 g8f6
B06,Masur Gambit,1.e2e4 g7g6 2.d2d4 g8h6 3.b1c3 f7f5 4.c1h6 f8h6 5.e4f5 e8g8
B07,B07,1.e2e4 d7d6
B07,"Maroczy Defense, Pirc",1.e2e4 d7d6 2.d2d4 e7e5
B07,"Philidor Gambit, Pirc",1.e2e4 d7d6 2.d2d4 e7e5 3.d4e5 c8d7
B07,Small Center Defense,1.e2e4 d7d6 2.d2d4 e7e6
B07,"Balogh Defense, Pirc",1.e2e4 d7d6 2.d2d4 f7f5
B07,Pirc Defense,1.e2e4 d7d6 2.d2d4 g8f6
B07,"Lion's Jaw, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.f2f3
B07,Czech Defense,1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 c7c6
B07,Pirc; Ufimtsev; Antal Defense,1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6
B07,"Kholmov System, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.f1c4
B07,"Classical System, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.f1e2
B07,"Chinese Variation, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.f1e2 f8g7 5.g2g4
B07,"Bayonet Attack, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.f1e2 f8g7 5.h2h4
B07,"Sveshnikov-Jansa Attack, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.c1e3 c7c6 5.h2h3
B07,150 Attack,1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.c1e3 c7c6 5.d1d2
B07,"Panov; Opatijav; Byrne Variation, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.c1g5
B07,Anti-Philidor,1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 b8d7 4.f2f4
B07,"Roscher Gambit, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.g1f3
B08,B08,1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.g1f3
B08,"Quiet System, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.g1f3 f8g7 5.f1e2
B08,"Parma Defense, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.g1f3 f8g7 5.f1e2 e8g8 6.e1g1 c8g4
B08,"Czech Defense, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.g1f3 f8g7 5.f1e2 e8g8 6.e1g1 c7c6
B08,"Chigorin Line, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.g1f3 f8g7 5.f1e2 e8g8 6.e1g1 b8c6
B08,"Schlechter Variation, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.g1f3 f8g7 5.h2h3
B09,"Austrian Attack, Pirc; B09",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.f2f4
B09,"Ljubojevic Variation, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.f2f4 f8g7 5.f1c4
B09,"Austrian; Yugoslav Defense, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.f2f4 f8g7 5.g1f3
B09,"Dragon Formation, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.f2f4 f8g7 5.g1f3 c7c5
B09,"Weiss Variation, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.f2f4 f8g7 5.g1f3 e8g8 6.f1d3
B09,"Kurajica Variation, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.f2f4 f8g7 5.g1f3 e8g8 6.c1e3
B09,"Unzicker Attack, Pirc",1.e2e4 d7d6 2.d2d4 g8f6 3.b1c3 g7g6 4.f2f4 f8g7 5.g1f3 e8g8 6.e4e5
B10,"Hillbilly Attack, Caro-Kann",1.e2e4 c7c6 2.f1c4
B10,"Modern; English Variation, Accelerated Panov; Caro-Kann",1.e2e4 c7c6 2.c2c4
B10,"Breyer Variation, Caro-Kann",1.e2e4 c7c6 2.d2d3
B10,"Stein Attack, Caro-Kann",1.e2e4 c7c6 2.d2d3 d7d5 3.b1d2 g7g6 4.g2g3 f8g7 5.f1g2 e7e5 6.g1f3 g8e7 7.e1g1 e8g8 8.b2b4
B10,"Massachussets Defense, Caro-Kann",1.e2e4 c7c6 2.d2d4 f7f5
B10,"Masi Variation, Caro-Kann",1.e2e4 c7c6 2.d2d4 g8f6
B10,"Scorpion-Horus Gambit, Caro-Kann",1.e2e4 c7c6 2.b1c3 d7d5 3.d2d3 d5e4 4.c1g5
B10,"Goldman Variation, Caro",1.e2e4 c7c6 2.b1c3 d7d5 3.d1f3
B10,"Two Knights Variation, Caro",1.e2e4 c7c6 2.g1f3 d7d5 3.b1c3
B11,B11,1.e2e4 c7c6 2.b1c3
B11,"Two Knights Attack, Caro Kann",1.e2e4 c7c6 2.b1c3 d7d5 3.g1f3
B11,"Mindeno Variation, Caro-Kann",1.e2e4 c7c6 2.b1c3 d7d5 3.g1f3 c8g4
B11,"Retreat Line, Mindeno Variation, Caro-Kann",1.e2e4 c7c6 2.b1c3 d7d5 3.g1f3 c8g4 4.h2h3 g4h5
B12,B12,1.e2e4 c7c6 2.d2d4
B12,"Landau Gambit, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.f1d3 g8f6 4.e4e5 f6d7 5.e5e6
B12,"Mieses Gambit, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.c1e3
B12,"Diemer-Duhm Gambit, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.c2c4
B12,"Prins Attack, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.e4e5 c8f5 4.b2b4
B12,"Bayonet Variation, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.e4e5 c8f5 4.g2g4
B12,"Tal Variation, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.e4e5 c8f5 4.h2h4
B12,"Van der Wiel Attack, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.e4e5 c8f5 4.b1c3
B12,"Dreyev Defense, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.e4e5 c8f5 4.b1c3 d8b6
B12,"Bronstein Variation, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.e4e5 c8f5 4.g1e2
B12,"Short Variation, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.e4e5 c8f5 4.g1f3
B12,"Botvinnik-Carls Defense, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.e4e5 c6c5
B12,"Maroczy Variation, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.f2f3
B12,"Fantasy; Lilienfisch Variation, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.f2f3
B12,"Maroczy Gambit, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.f2f3 d5e4 4.f3e4 e7e5 5.g1f3 e5d4 6.f1c4
B12,"Modern Variation, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.b1d2
B12,New Caro-Kann,1.e2e4 c7c6 2.d2d4 d7d5 3.b1d2 g7g6
B12,"Edinburgh Variation, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.b1d2 d8b6
B12,"Ulysses Gambit, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.g1f3 d5e4 4.f3g5
B12,"De Bruycker Defense, Caro",1.e2e4 c7c6 2.d2d4 b8a6
B12,"Hector Gambit, Caro",1.e2e4 c7c6 2.b1c3 d7d5 3.g1f3 d5e4 4.f3g5
B13,"Exchange Variation, Caro; B13",1.e2e4 c7c6 2.d2d4 d7d5 3.e4d5
B13,Panov-Botvinnik; Gunderam Attack,1.e2e4 c7c6 2.d2d4 d7d5 3.e4d5 c6d5 4.c2c4
B14,B14,1.e2e4 c7c6 2.d2d4 d7d5 3.e4d5 c6d5 4.c2c4 g8f6 5.b1c3
B14,"Carlsbad Line, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.e4d5 c6d5 4.c2c4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6
B14,"Czerniak Line, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.e4d5 c6d5 4.c2c4 g8f6 5.b1c3 b8c6 6.c1g5 d8a5
B14,"Reifir-Spielmann Line, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.e4d5 c6d5 4.c2c4 g8f6 5.b1c3 b8c6 6.c1g5 d8b6
B15,B15,1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3
B15,"Gurgenidze Variation, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 b7b5
B15,"Von Hennig Gambit, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.f1c4
B15,"Milner-Barry Gambit, Caro; Rasa-Studier Gambit",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.f2f3
B15,"Knight Variation, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 g8f6
B15,"Tarrasch (Alekhine) Gambit, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 g8f6 5.f1d3
B15,"Tartakower Variation, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 g8f6 5.e4f6 e7f6
B15,"Forgacs Variation, Caro Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 g8f6 5.e4f6 e7f6 6.f1c4
B15,"Gurgenidze System, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 g7g6
B15,"Gurgenidze Variation, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 g7g6 4.e4e5 f8g7 5.f2f4 h7h5
B15,"Campomanes Attack, Caro",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 g8f6
B16,"Finnish Variation, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 h7h6
B16,"Nimzovich; Bronstein-Larsen, Caro; B16",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 g8f6 5.e4f6 g7f6
B17,"Karpov Variation, Caro-Kann; B17",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 b8d7
B17,"Smyslov Variation, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 b8d7 5.f1c4 g8f6 6.e4g5 e7e6 7.d1e2 d7b6
B17,"Tiviakov-Fischer Attack, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 b8d7 5.f1c4 g8f6 6.e4f6 d7f6
B17,"Kasparov Attack, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 b8d7 5.g1f3 g8f6 6.e4g3
B17,"Ivanchuk Defense, Caro Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 b8d7 5.e4g5 d7f6
B18,"Classical Variation, Caro; B18",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 c8f5
B18,"Flohr Variation, Caro-Kann",1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 c8f5 5.e4g3 f5g6 6.g1h3
B19,B19,1.e2e4 c7c6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 c8f5 5.e4g3 f5g6 6.h2h4 h7h6 7.g1f3
B20,"Mengarini Gambit, Sicilian",1.e2e4 c7c5 2.a2a3 b8c6 3.b2b4
B20,"Snyder Variation, Sicilian",1.e2e4 c7c5 2.b2b3
B20,"Queen Fianchetto Variation, Sicilian",1.e2e4 c7c5 2.b2b3 b7b6
B20,"Wing Gambit, Sicilian",1.e2e4 c7c5 2.b2b4
B20,"Wing; Carlsbad Variation, Sicilian",1.e2e4 c7c5 2.b2b4 c5b4 3.a2a3 b4a3
B20,"Marienbad Variation, Sicilian",1.e2e4 c7c5 2.b2b4 c5b4 3.a2a3 d7d5 4.e4d5 d8d5 5.c1b2
B20,"Abrahams Variation, Sicilian Wing Gambit",1.e2e4 c7c5 2.b2b4 c5b4 3.c1b2
B20,"Santasiere Wing Gambit, Sicilian",1.e2e4 c7c5 2.b2b4 c5b4 3.c2c4
B20,"Bowlder Attack, Sicilian",1.e2e4 c7c5 2.f1c4
B20,Slow Sicilian,1.e2e4 c7c5 2.f1e2
B20,"Gloria Variation, Sicilian",1.e2e4 c7c5 2.c2c4 d7d6 3.b1c3 b8c6 4.g2g3 h7h5
B20,"Chameleon Variation, Sicilian",1.e2e4 c7c5 2.d2d3
B20,"Big Clamp Variation, Sicilian",1.e2e4 c7c5 2.d2d3 b8c6 3.c2c3 d7d6 4.f2f4
B20,Sicilian Center Game,1.e2e4 c7c5 2.d2d4 c5d4
B20,"Smith-Morra; Fleissig Gambit, Sic",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3
B20,"Rivadavia Gambit, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3
B20,"Push Variation, Smith-Morra Gambit, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4d3
B20,"Dubois Variation, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4d3 4.c3c4
B20,"Scandanavian Formation, Smith-Morra, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d7d5
B20,"Smith-Morra Accepted, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3
B20,"Chicago Defense, Smith-Morra Accepted, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3 4.b1c3 d7d6 5.f1c4 e7e6 6.g1f3 g8f6 7.e1g1 a7a6
B20,"Taimanov Formation, Smith-Morra Accepted, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3 4.b1c3 e7e6 5.f1c4 a7a6 6.g1f3 g8e7
B20,"Kan Formation, Smith-Morra Accepted, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3 4.b1c3 e7e6 5.g1f3 a7a6
B20,"Fianchetto Defense, Smith-Morra Accepted, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3 4.b1c3 e7e6 5.g1f3 g7g6
B20,"Classical Formation, Smith-Morra Accepted, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3 4.b1c3 b8c6 5.g1f3 d7d6 6.f1c4 a7a6 7.e1g1 g8f6
B20,"Scheveningen Formation, Smith-Morra Accepted",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3 4.b1c3 b8c6 5.g1f3 d7d6 6.f1c4 e7e6
B20,"Paulsen Formation, Smith-Morra Accepted, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3 4.b1c3 b8c6 5.g1f3 e7e6 6.f1c4 a7a6
B20,"Morphy Defense Deferred, Smith-Morra Accepted",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3 4.b1c3 b8c6 5.g1f3 e7e6 6.f1c4 a7a6 7.e1g1 b7b5 8.c4b3 f8c5
B20,"Larsen Defense, Smith-Morra Accepted, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3 4.b1c3 b8c6 5.g1f3 e7e6 6.f1c4 a7a6 7.e1g1 d8c7 8.d1e2 f8d6
B20,"Pin Defense, Smith-Morra Accepted, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3 4.b1c3 b8c6 5.g1f3 e7e6 6.f1c4 f8b4
B20,"Morphy Defense, Smith-Morra Accepted, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3 4.b1c3 b8c6 5.g1f3 e7e6 6.f1c4 f8c5
B20,"Sozin Formation, Smith-Morra Accepted, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d4c3 4.b1c3 b8c6 5.g1f3 e7e6 6.f1c4 d7d6 7.e1g1 a7a6 8.d1e2 b7b5
B20,"Center Formation, Smith-Morra Declined, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 e7e5
B20,"Alapin Formation, Smith-Morra Declined, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 g8f6
B20,"Wing Formation, Smith-Morra Declined, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.c2c3 d8a5
B20,"Halasz Gambit, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.f2f4
B20,"Morphy Gambit, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.g1f3
B20,"Andreaschek Gambit, Sicilian",1.e2e4 c7c5 2.d2d4 c5d4 3.g1f3 e7e5 4.c2c3
B20,"Steinitz; Tartakower, Lasker-Dunne Variation, Sicilian",1.e2e4 c7c5 2.g2g3
B20,"Myers Attack, Sicilian",1.e2e4 c7c5 2.h2h4
B20,King David's Opening,1.e2e4 c7c5 2.e1e2
B20,"Keres; Billium Attack, Sicilian; Chameleon",1.e2e4 c7c5 2.g1e2
B20,"Brick Variation, Sicilian",1.e2e4 c7c5 2.g1h3
B21,"Larsen-Santasiere (McDonnell Attack) Variation, Sicilian; B21",1.e2e4 c7c5 2.f2f4
B21,"Tal Gambit, Sicilian",1.e2e4 c7c5 2.f2f4 d7d5 3.e4d5 g8f6
B22,"Alapin; Sveshnikov Var, Sicilian; B22",1.e2e4 c7c5 2.c2c3
B22,"Barmen Defense, Sicilian",1.e2e4 c7c5 2.c2c3 d7d5 3.e4d5 d8d5
B22,"Sherzer Variation, Sicilian",1.e2e4 c7c5 2.c2c3 g8f6 3.e4e5 f6d5 4.d2d4 e7e6 5.g1f3 b8c6
B22,"Stoltz Attack, Sicilian",1.e2e4 c7c5 2.c2c3 g8f6 3.e4e5 f6d5 4.g1f3 b8c6 5.f1c4 d5b6 6.c4b3
B23,"Closed; Chigorin Variation, Sicilian; B23",1.e2e4 c7c5 2.b1c3
B23,"Grand Prix Variation, Sicilian",1.e2e4 c7c5 2.b1c3 d7d6 3.f2f4
B23,"Marshall Gambit, Sicilian",1.e2e4 c7c5 2.b1c3 e7e6 3.d2d4 d7d5
B23,"Korchnoi Defense, Sicilian",1.e2e4 c7c5 2.b1c3 e7e6 3.g2g3 d7d5
B23,Traditional Closed Sicilian,1.e2e4 c7c5 2.b1c3 b8c6
B23,"Portland Attack, Sicilian",1.e2e4 c7c5 2.b1c3 b8c6 3.d2d3 g7g6 4.g2g4
B23,"Vinken System, Sicilian; Grand Prix Attack",1.e2e4 c7c5 2.b1c3 b8c6 3.f2f4
B23,"Schofman Variation, Sicilian",1.e2e4 c7c5 2.b1c3 b8c6 3.f2f4 g7g6 4.g1f3 f8g7 5.f1c4 e7e6 6.f4f5
B23,Chameleon Sicilian,1.e2e4 c7c5 2.b1c3 b8c6 3.g1e2
B24,"Fianchetto Variation, Sicilian Closed; B24",1.e2e4 c7c5 2.b1c3 b8c6 3.g2g3
B25,B25,1.e2e4 c7c5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.d2d3 d7d6
B25,"Botvinnik Defense, Sicilian Closed",1.e2e4 c7c5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.d2d3 d7d6 6.f2f4 e7e5
B25,"Edge Variation, Sicilian Closed",1.e2e4 c7c5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.d2d3 d7d6 6.f2f4 e7e5 7.g1h3 g8e7
B26,B26,1.e2e4 c7c5 2.b1c3 b8c6 3.g2g3 g7g6 4.f1g2 f8g7 5.d2d3 d7d6 6.c1e3
B27,B27,1.e2e4 c7c5 2.g1f3
B27,"Polish Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 b7b5
B27,"Katalimov Variation, Sicilian",1.e2e4 c7c5 2.g1f3 b7b6
B27,"Jalalabad Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e5
B27,"Brussels Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 f7f5
B27,"Double-Dutch Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 f7f5 3.e4f5 g8h6
B27,"Hyperaccelerated Dragon, Sicilian",1.e2e4 c7c5 2.g1f3 g7g6
B27,"Action Extension, Sicilian",1.e2e4 c7c5 2.g1f3 g7g6 3.c2c4 f8h6
B27,"Hyperaccelerated Pterodactyl, Sic",1.e2e4 c7c5 2.g1f3 g7g6 3.d2d4 f8g7
B27,"Frederico Variation, Sicilian",1.e2e4 c7c5 2.g1f3 g7g6 3.d2d4 f7f5
B27,Sicilian Fred,1.e2e4 c7c5 2.g1f3 g7g6 3.d2d4 f7f5
B27,"Buecker Variation, Sicilian",1.e2e4 c7c5 2.g1f3 h7h6
B27,"Mongoose Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d8a5
B27,"Althouse; Stiletto Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d8a5
B27,"Accelerated Paulsen, Sicilian",1.e2e4 c7c5 2.g1f3 d8c7
B27,"Quinteros Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d8c7
B28,"O'Kelly Variation, Sicilian; B28",1.e2e4 c7c5 2.g1f3 a7a6
B28,"Kieseritzky System, Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.b2b3
B28,"Wing Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.b2b4
B28,"Aronin System, Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.f1e2
B28,"Venice System, Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.c2c3
B28,"Ljubojevic Line, Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.c2c3 b7b5
B28,"Gambit Line, O'Kelly Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.c2c3 d7d5 4.e4d5 g8f6
B28,"Steiner Line, Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.c2c3 d7d6
B28,"Barcza Line, Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.c2c3 g8f6
B28,"Maroczy Bind, O'Kelly Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.c2c4
B28,"Robatsch Line, Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.c2c4 d7d6
B28,"Paulsen Line, Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.c2c4 e7e6
B28,"Geller Line, Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.c2c4 b8c6 4.d2d4 c5d4 5.f3d4 e7e5
B28,"Quiet System, O'Kelly Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.d2d3
B28,"Normal System, O'Kelly Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.d2d4
B28,"Cortlever Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.d2d4 c5d4 4.f1c4
B28,"Smith-Morra Line, O'Kelly Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.d2d4 c5d4 4.c2c3
B28,"Taimanov Line, O'Kelly Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.d2d4 c5d4 4.f3d4 e7e5
B28,"Kan Line, O'Kelly Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.d2d4 c5d4 4.f3d4 e7e6
B28,"Zagorovsky Line, O'Kelly Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.d2d4 c5d4 4.d1d4
B28,"Reti System, O'Kelly Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.g2g3
B28,"Yerevan System, O'Kelly Sicilian",1.e2e4 c7c5 2.g1f3 a7a6 3.b1c3
B29,"Nimzovich; Rubinstein Var, Sicilian; B29",1.e2e4 c7c5 2.g1f3 g8f6
B29,"Advance Variation, Sicilian",1.e2e4 c7c5 2.g1f3 g8f6 3.e4e5
B30,Old Sicilian,1.e2e4 c7c5 2.g1f3 b8c6
B30,"Rossolimo (Nezmetdinov) Variation, Sicilian; B30",1.e2e4 c7c5 2.g1f3 b8c6 3.f1b5
B31,B31,1.e2e4 c7c5 2.g1f3 b8c6 3.f1b5 g7g6
B31,"Gufeld Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.f1b5 g7g6 4.e1g1 f8g7 5.c2c3 e7e5 6.d2d4
B31,"Lutikov Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.f1b5 g7g6 4.e1g1 f8g7 5.c2c3 g8f6 6.d2d4
B31,"Gurgenidze Variation, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.f1b5 g7g6 4.e1g1 f8g7 5.f1e1 e7e5 6.b2b4
B31,"San Francisco Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.f1b5 c6a5 4.b2b4
B32,B32,1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4
B32,"Nimzovich Variation, Sicilian; Nimzo-American Variation",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 d7d5
B32,"Bourdonnais; Lowenthal Var, Sic",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 e7e5
B32,"Kalashnikov Variation, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 e7e5 5.d4b5 d7d6
B32,"Retreat Variation, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e5 6.d4f3
B32,"Godiva Variation, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 d8b6
B32,"Flohr Variation, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 d8c7
B33,B33,1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g8f6
B33,"Four Knights Variation, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3
B33,"Lasker-Pelikan Variation, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e5
B33,"Pelikan Variation, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e5 6.d4b5 d7d6
B33,"Chelyabinsk Variation, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e5 6.d4b5 d7d6 7.c1g5 a7a6 8.b5a3 b7b5
B33,"Sveshnikov Variation, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e5 6.d4b5 d7d6 7.c1g5 a7a6 8.b5a3 b7b5 9.g5f6 g7f6 10.c3d5 f6f5
B33,"Bird Variation, Siclian",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e5 6.d4b5 d7d6 7.c1g5 a7a6 8.b5a3 c8e6
B34,Accelerated Dragon; Fianchetto Var; B34,1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g7g6
B35,B35,1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g7g6 5.b1c3 f8g7 6.c1e3 g8f6 7.f1c4
B36,"Maroczy Bind, Sicilian; B36",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g7g6 5.c2c4
B36,"Gurgenidze Variation, Sicilian",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g7g6 5.c2c4 g8f6 6.b1c3 c6d4 7.d1d4 d7d6
B37,B37,1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g7g6 5.c2c4 f8g7
B38,B38,1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g7g6 5.c2c4 f8g7 6.c1e3
B39,"Breyer Variation, Sicilian; B39",1.e2e4 c7c5 2.g1f3 b8c6 3.d2d4 c5d4 4.f3d4 g7g6 5.c2c4 f8g7 6.c1e3 g8f6 7.b1c3 f6g4
B40,"French Variation, Sicilian; B40",1.e2e4 c7c5 2.g1f3 e7e6
B40,"Wing Gambit Deferred, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.b2b4
B40,"Kramnik Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.c2c4
B40,"Gaw-Paw Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 d8b6
B40,"Kveinis Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 d8b6
B41,B41,1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4
B41,"Accelerated Wimpy Variation, Sic",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4
B41,"Paulsen; Kan Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 a7a6
B41,"Reti Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 a7a6 5.c2c4
B41,"Hedgehog Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 a7a6 5.c2c4 g7g6
B41,"Bronstein Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 a7a6 5.c2c4 g8f6 6.b1c3 f8b4 7.f1d3 b8c6
B41,"Anderssen; Counterattack Var, Sic",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 g8f6
B41,Pin Variation; Sicilian CounterAttack,1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 f8b4
B41,"Jaffe Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 f8b4 6.f1d3 e6e5
B41,"Koch Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 f8b4 6.e4e5
B41,"Marshall Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 d7d5
B42,B42,1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 a7a6 5.f1d3
B42,"Polugaevsky Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 a7a6 5.f1d3 f8c5
B42,"Swiss Cheese Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 a7a6 5.f1d3 g7g6
B43,B43,1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 a7a6 5.b1c3
B44,"Taimanov; Bastrikov, Paulsen Variation, Sicilian; B44",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 b8c6
B44,"Szen Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 b8c6 5.d4b5
B44,"Gary Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 b8c6 5.d4b5 d7d6 6.c2c4 g8f6 7.b1c3 a7a6 8.b5a3 d6d5
B45,B45,1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 b8c6 5.b1c3
B45,"Paulsen Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 b8c6 5.b1c3 g8f6 6.d4b5 f8b4
B45,"American Attack, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 b8c6 5.b1c3 g8f6 6.d4b5 f8b4 7.b5d6
B45,"Four Knights Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6
B45,"Cobra Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.d4b5 f8c5
B46,B46,1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 b8c6 5.b1c3 a7a6
B46,"Taimanov Variation, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 b8c6 5.b1c3 a7a6 6.f1e2 g8e7
B47,"Bastrikov Variation, Sicilian; B47",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 b8c6 5.b1c3 d8c7
B47,"Ponomariov Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 b8c6 5.b1c3 d8c7 6.d4b5 c7b8 7.c1e3 a7a6 8.e3b6
B48,B48,1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 b8c6 5.b1c3 d8c7 6.c1e3
B49,B49,1.e2e4 c7c5 2.g1f3 e7e6 3.d2d4 c5d4 4.f3d4 b8c6 5.b1c3 d8c7 6.c1e3 a7a6 7.f1e2
B50,"Modern Variation, Sicilian; B50",1.e2e4 c7c5 2.g1f3 d7d6
B50,"Kopec System, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.f1d3
B50,"Delayed Alapin, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.c2c3
B50,Basman-Palatnik Gambit,1.e2e4 c7c5 2.g1f3 d7d6 3.c2c3 g8f6 4.f1e2 b8c6 5.d2d4 c5d4 6.c3d4 f6e4
B50,"Kotov Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.g2g3 b7b5
B51,"Wing Gambit Deferred, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.b2b4
B51,"Moscow Variation, Sicilian; Canal Attack; B51",1.e2e4 c7c5 2.g1f3 d7d6 3.f1b5
B51,"Moscow Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.f1b5 b8c6 4.e1g1 c8d7 5.c2c3 a7a6 6.b5c6 d7c6 7.f1e1 g8f6 8.d2d4 c6e4 9.c1g5
B51,"Dorfman Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.f1b5 b8c6 4.e1g1 c8d7 5.d1e2 g7g6 6.e4e5
B52,"Bronstein (Haag) Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.f1b5 c8d7 4.b5d7 d8d7 5.e1g1 b8c6 6.c2c3 g8f6 7.d2d4
B53,B53,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4
B53,"Tartakower Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.c2c3
B53,"Chekhover Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.d1d4
B53,"Zaitsev Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.d1d4 b8c6 5.f1b5 d8d7
B53,"Lazy Knight Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 b8d7
B53,"Cortlever Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 g8f6 4.d4c5 f6e4 5.c5d6
B54,B54,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4
B55,B55,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6
B55,"Ginsberg Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.f1c4
B55,"Moscow (Prins) Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.f2f3
B55,"Venice Attack, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.f2f3 e7e5 6.f1b5
B56,B56,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3
B56,"Kupreichik Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 c8d7
B56,"Venice Attack, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e5 6.f1b5
B56,"Classical Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6
B56,"Yates Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.f1d3
B56,"Spielmann Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.d4e2
B57,"Sozin; Velimirovic Attack, Sicilian; B57",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.f1c4
B57,"Anti-Sozin Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.f1c4 d8b6
B58,B58,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.f1e2
B58,"Boleslavsky Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.f1e2 e7e5
B58,"Louma Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.f1e2 e7e5 7.d4c6
B58,"Fianchetto Variation, Richter-Rauzer Sicilizn",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.g2g3
B59,B59,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.f1e2 e7e5 7.d4b3
B60,"Richter-Rauzer Variation, Sicilian; B60",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5
B60,"Dragon Variation, Richter-Rauzer Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 g7g6
B61,B61,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.d1d2
B62,B62,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6
B62,"Margate (Vitolins) Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6 7.f1b5
B62,"Podebrady Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6 7.d4b3
B62,"Exchange Variation, Richter-Rauzer Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6 7.d4c6
B63,"Rauzer Variation, Sicilian; B63",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6 7.d1d2
B63,"Ivanov Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6 7.d1d2 d8b6
B64,B64,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6 7.d1d2 f8e7 8.e1c1 e8g8 9.f2f4
B65,B65,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6 7.d1d2 f8e7 8.e1c1 e8g8 9.f2f4 c6d4 10.d2d4
B66,B66,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6 7.d1d2 a7a6
B67,B67,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6 7.d1d2 a7a6 8.e1c1 c8d7
B68,B68,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6 7.d1d2 a7a6 8.e1c1 c8d7 9.f2f4 f8e7
B69,B69,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 b8c6 6.c1g5 e7e6 7.d1d2 a7a6 8.e1c1 c8d7 9.f2f4 f8e7 10.d4f3 b7b5 11.g5f6
B70,"Dragon Variation, Sicilian; B70",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6
B70,"Alekhine Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.f1e2 f8g7 7.e1g1 e8g8 8.d4b3
B70,"Fianchetto Variation, Sicilian Dragon",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.g2g3
B71,"Levenfish Variation, Sicilian; B71",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.f2f4
B72,B72,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3
B72,"Nottingham Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f1e2 b8c6 8.d4b3
B72,"Amsterdam (Battery) Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f1e2 b8c6 8.d1d2
B72,"Grigoriev Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f1e2 b8c6 8.d1d2 e8g8 9.e1c1
B73,B73,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f1e2 b8c6 8.e1g1
B73,"Zollner Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f1e2 b8c6 8.e1g1 e8g8 9.f2f4 d8b6
B74,B74,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f1e2 b8c6 8.e1g1 e8g8 9.d4b3
B74,"Alekhine Line, Sicilian Dragon",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f1e2 b8c6 8.e1g1 e8g8 9.d4b3 a7a5
B74,"Maroczy Line, Sicilian Dragon",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f1e2 b8c6 8.e1g1 e8g8 9.d4b3 c8e6 10.f2f4 c6a5
B74,"Tartakower Line, Sicilian Dragon",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f1e2 b8c6 8.e1g1 e8g8 9.d4b3 c8e6 10.f2f4 d8c8
B75,B75,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f2f3
B76,"Yugoslav Attack, Sicilian Dragon; B76",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f2f3 e8g8
B76,"Panov Variation, Sicilian Dragon",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f2f3 e8g8 8.d1d2 b8c6 9.g2g4
B77,B77,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f2f3 e8g8 8.d1d2 b8c6 9.f1c4
B77,"Sosonko Variation, Sicilian Dragon",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f2f3 e8g8 8.d1d2 b8c6 9.f1c4 f6d7
B78,B78,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f2f3 e8g8 8.d1d2 b8c6 9.f1c4 c8d7 10.e1c1
B79,B79,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 g7g6 6.c1e3 f8g7 7.f2f3 e8g8 8.d1d2 b8c6 9.f1c4 c8d7 10.e1c1 d8a5
B80,"Scheveingen Variation, Sicilian; B80",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6
B80,"Vitolins Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.f1b5
B80,"English Attack, Scheveningen Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.c1e3 a7a6 7.f2f3
B80,"Fianchetto Variation, Scheveningen Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.g2g3
B81,"Delayed Keres Attack, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.c1e3 a7a6 7.g2g4
B81,"Perenyi Gambit, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.c1e3 a7a6 7.g2g4 e6e5 8.d4f5 g7g6 9.g4g5
B81,"Keres; Panov Variation, Sicilian; B81",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.g2g4
B82,"Matanovic Attack, Sicilian; B82",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.f2f4
B82,"Tal Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.f2f4 b8c6 7.c1e3 f8e7 8.d1f3
B84,"Modern Paulsen, Sicilian; B84",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.f1e2 a7a6
B85,B85,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.f1e2 a7a6 7.e1g1 f8e7
B85,"Paulsen Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.f1e2 a7a6 7.e1g1 d8c7 8.f2f4 b8c6
B86,"Fischer Attack, Sozin Attack, Sicilian; B86",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.f1c4
B87,"Flank Variation, Sicilian; B87",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.f1c4 a7a6 7.c4b3 b7b5
B88,"Leonhardt Variation, Sicilian; B88",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.f1c4 b8c6
B89,"Sherbakov Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.f1c4 b8c6 7.c1e3 f8e7 8.c4b3 e8g8 9.e1g1 c6a5 10.f2f4 b7b6
B89,"Velimirovic Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 e7e6 6.f1c4 b8c6 7.c1e3 f8e7 8.d1e2
B90,"Najdorf Variation, Sicilian; B90",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6
B90,"Lipnitzky Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.f1c4
B90,"Byrne Variation, Sicilian; English Attack",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1e3
B90,"Anti-English, Najdorf Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1e3 f6g4
B90,"Adams Attack, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.h2h3
B90,"Freak Attack, Siicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.h1g1
B91,"Zagreb Variation, Sicilian; B91",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.g2g3
B92,"Opocensky Variation, Sicilian; B92",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.f1e2
B92,"Opocensky Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.f1e2 e7e5
B93,"Amsterdam Variation, Sicilian; B93",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.f2f4
B94,B94,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1g5
B95,B95,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1g5 e7e6
B96,B96,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1g5 e7e6 7.f2f4
B96,Polugaevsky; Accelerated Najdorf,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1g5 e7e6 7.f2f4 b7b5
B96,"Simagin Line, Najdorf Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1g5 e7e6 7.f2f4 b7b5 8.e4e5 d6e5 9.f4e5 d8c7 10.d1e2
B97,"Poisoned Pawn Variation, Najdorf Sicilian; B97",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1g5 e7e6 7.f2f4 d8b6
B97,"Poisoned Pawn Variation, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1g5 e7e6 7.f2f4 d8b6 8.d1d2
B97,"Poisoned Pawn Accepted, Najdorf Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1g5 e7e6 7.f2f4 d8b6 8.d1d2 b6b2
B98,B98,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1g5 e7e6 7.f2f4 f8e7
B98,"Argenitne; Goteborg Var, Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1g5 e7e6 7.f2f4 f8e7 8.d1f3 h7h6 9.g5h4 g7g5
B98,"Browne Variation, Najdorf Sicilian",1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1g5 e7e6 7.f2f4 f8e7 8.d1f3 h7h6 9.g5h4 d8c7
B99,B99,1.e2e4 c7c5 2.g1f3 d7d6 3.d2d4 c5d4 4.f3d4 g8f6 5.b1c3 a7a6 6.c1g5 e7e6 7.f2f4 f8e7 8.d1f3 d8c7 9.e1c1 b8d7
C00,French Defense; C00,1.e2e4 e7e6
C00,"Reti (Horwitz) Variation, French",1.e2e4 e7e6 2.b2b3
C00,"Papa-Ticulat Gambit, French",1.e2e4 e7e6 2.b2b3 d7d5 3.c1b2
C00,"Banzai-Leong Gambit, French",1.e2e4 e7e6 2.b2b4
C00,Pinova Gambit,1.e2e4 e7e6 2.b2b4 f8b4 3.e4e5
C00,"Bird Invitation, French",1.e2e4 e7e6 2.f1b5
C00,"Orthoschapp Variation, French",1.e2e4 e7e6 2.c2c4 d7d5 3.c4d5 e6d5 4.d1b3
C00,"Reversed Philidor, French",1.e2e4 e7e6 2.d2d3 d7d5 3.b1d2 g8f6 4.g1f3 b8c6 5.f1e2
C00,"Franco-Hiva Gambit, French",1.e2e4 e7e6 2.d2d3 f7f5
C00,"Alapin; Witaker Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.c1e3
C00,"Alapin-Diemer Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.c1e3 d5e4 4.f2f3
C00,"Diemer-Duhm Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.c2c4
C00,"Steinitz Attack, French",1.e2e4 e7e6 2.e4e5
C00,"Reuter Gambit, French",1.e2e4 e7e6 2.f2f4 d7d5 3.g1f3 d5e4 4.f3g5
C00,"Reti-Spielmann Attack, French",1.e2e4 e7e6 2.g2g3
C00,"Queen's Knight, French",1.e2e4 e7e6 2.b1c3
C00,"Pelikan Variation, French",1.e2e4 e7e6 2.b1c3 d7d5 3.f2f4
C00,"Two Knight Variation, French",1.e2e4 e7e6 2.b1c3 d7d5 3.g1f3
C00,"Knight Variation, French",1.e2e4 e7e6 2.g1f3
C00,"Wing Gambit, French",1.e2e4 e7e6 2.g1f3 d7d5 3.e4e5 c7c5 4.b2b4
C00,"Two Knights Variation, French",1.e2e4 e7e6 2.g1f3 d7d5 3.b1c3
C00,"Chigorin Variation, French",1.e2e4 e7e6 2.d1e2
C01,C01,1.e2e4 e7e6 2.d2d4
C01,"Baeuerle Gambit, French",1.e2e4 e7e6 2.d2d4 b7b5
C01,"Schlechter Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.f1d3
C01,"Exchange Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4d5 e6d5
C01,"Monte Carlo Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4d5 e6d5 4.c2c4
C01,"Svenius Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4d5 e6d5 4.b1c3 g8f6 5.c1g5
C01,"Canal Attack, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4d5 e6d5 5.f1d3 g8e7 6.d1h5
C01,"Perseus Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.g1f3
C01,"Carlson Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.g1f3 d5e4 4.f3e5
C01,"Morphy Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.g1h3
C01,Hoffman Gambit,1.e2e4 e7e6 2.d2d4 d7d5 3.d1e2 e6e5 4.f2f4 e5f4
C01,"Franco-Hiva Gambit I, French",1.e2e4 e7e6 2.d2d4 f7f5
C01,"Mediterranean Defense, French",1.e2e4 e7e6 2.d2d4 g8f6
C02,"Advance Variation, French; C02",1.e2e4 e7e6 2.d2d4 d7d5 3.e4e5
C02,"Extended Bishop Swap, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4e5 c8d7
C02,"Advance; Nimzovich; Steinitz Var, Fr",1.e2e4 e7e6 2.d2d4 d7d5 3.e4e5 c8d7
C02,"Frenkel Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4e5 c7c5 4.b2b4
C02,"Paulsen Attack, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4e5 c7c5 4.c2c3 b8c6 5.g1f3
C02,"Euwe Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4e5 c7c5 4.c2c3 b8c6 5.g1f3 c8d7
C02,"Lputian Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4e5 c7c5 4.c2c3 b8c6 5.g1f3 d8b6 6.a2a3 g8h6
C02,"Milner-Barry Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4e5 c7c5 4.c2c3 b8c6 5.g1f3 d8b6 6.f1d3
C02,"Wade Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4e5 c7c5 4.c2c3 d8b6 5.g1f3 c8d7
C02,"Nimzovich Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4e5 c7c5 4.g1f3
C02,"Ruisdonk Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4e5 c7c5 4.g1f3 c5d4 5.f1d3
C02,"Nimzovich Attack, French",1.e2e4 e7e6 2.d2d4 d7d5 3.e4e5 c7c5 4.d1g4
C03,"Tarrasch Variation, French; C03",1.e2e4 e7e6 2.d2d4 d7d5 3.b1d2
C03,"Morozevich Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1d2 f8e7
C03,"Haberditz Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1d2 f7f5
C04,C04,1.e2e4 e7e6 2.d2d4 d7d5 3.b1d2 b8c6
C05,C05,1.e2e4 e7e6 2.d2d4 d7d5 3.b1d2 g8f6
C06,C06,1.e2e4 e7e6 2.d2d4 d7d5 3.b1d2 g8f6 4.e4e5 f6d7 5.f1d3
C07,Open System Tarrasch French; C07,1.e2e4 e7e6 2.d2d4 d7d5 3.b1d2 c7c5
C07,"Shaposhnikov Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1d2 c7c5 4.e4d5 g8f6
C07,"Chistiakov Defense, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1d2 c7c5 4.e4d5 d8d5
C08,C08,1.e2e4 e7e6 2.d2d4 d7d5 3.b1d2 c7c5 4.e4d5 e6d5
C09,C09,1.e2e4 e7e6 2.d2d4 d7d5 3.b1d2 c7c5 4.e4d5 e6d5 5.g1f3 b8c6
C10,"Normal Variation, French; C10",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3
C10,"Marshall Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 c7c5
C10,"Rubinstein Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 d5e4
C10,"Rasa-Studier Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 d5e4 4.c1e3
C10,"Ellis Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 e6e5
C10,"Blackburne Defense, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 b8d7
C10,"Capablanca Line, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 b8d7 5.g1f3 g8f6 6.e4f6 d7f6 7.f3e5
C10,"Frere; Becker; Katilimov, Maric Var, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 d5e4 4.c3e4 d8d5
C10,"Svenonius Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.e4d5
C11,"Classical Variation, French; C11",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6
C11,"Swiss Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.f1d3
C11,"Henneberger Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1e3
C11,"Burn Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 d5e4
C11,"Steinitz Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.e4e5
C11,"Bradford Attack, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.e4e5 f6d7 5.f2f4 c7c5 6.d4c5 f8c5 7.d1g4
C11,"Brodsky-Jones Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.e4e5 f6d7 5.f2f4 c7c5 6.d4c5 b8c6 7.a2a3 f8c5 8.d1g4 e8g8 9.g1f3 f7f6
C11,"Boleslavsky Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.e4e5 f6d7 5.f2f4 c7c5 6.g1f3 b8c6 7.c1e3
C11,"Gledhill Attack, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.e4e5 f6d7 5.d1g4
C12,C12,1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5
C12,"McCutcheon Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8b4
C12,"Olland (Dutch) Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8b4 5.e4e5 h7h6 6.g5c1
C12,"Lasker Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8b4 5.e4e5 h7h6 6.g5d2 b4c3
C12,"Duras Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8b4 5.e4e5 h7h6 6.g5d2 b4c3 7.b2c3 f6e4 8.d1g4 e8f8 9.d2c1
C12,"Tartakower Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8b4 5.e4e5 h7h6 6.g5d2 f6d7
C12,"Janowski Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8b4 5.e4e5 h7h6 6.g5e3
C12,"Bernstein Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8b4 5.e4e5 h7h6 6.g5h4
C12,"Chigorin Variation, Frenchï¿½",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8b4 5.e4e5 h7h6 6.e5f6
C12,"Grigoriev Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8b4 5.e4e5 h7h6 6.e5f6 h6g5 7.f6g7 h8g8 8.h2h4 g5h4 9.d1g4
C13,"Vistaneckis (Nimzovich) Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6g8
C13,"Frankfurt Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6g8 6.g5e3 b7b6
C14,C14,1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7
C14,"Anderssen (Richter) Attack, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.g5f6
C14,"Tartakower Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6e4
C14,"Classical Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6d7 6.g5e7 d8e7
C14,"Tarrasch Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6d7 6.g5e7 d8e7 7.f1d3
C14,"Steinitz Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6d7 6.g5e7 d8e7 7.f2f4
C14,"Alapin Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6d7 6.g5e7 d8e7 7.c3b5
C14,"Rubinstein Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6d7 6.g5e7 d8e7 7.d1d2
C14,"Pollock Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6d7 6.g5e7 d8e7 7.d1g4
C14,"Chatard-Alekhine; Albin Attack, Fr",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6d7 6.h2h4
C14,"Maroczy Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6d7 6.h2h4 a7a6
C14,"Albin-Chatard Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6d7 6.h2h4 e7g5 7.h4g5 d8g5
C14,"Breyer Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6d7 6.h2h4 c7c5
C14,"Teichmann Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6d7 6.h2h4 f7f6
C14,"Spielmann Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 g8f6 4.c1g5 f8e7 5.e4e5 f6d7 6.h2h4 e8g8
C15,"Winckelmann-Riemer Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.a2a3
C15,"Fingerslip Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.c1d2
C15,"Kunin Double Gambit, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.c1d2 d5e4 5.d1g4 d8d4
C15,"Schwarz's Line, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.c1d2 g8e7 5.c3b1
C15,"Kondratiyev Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.f1d3 c7c5 5.e4d5 d8d5 6.c1d2
C15,"Delayed Exchange Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4d5
C15,"Alekhine (Maroczy) Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.g1e2
C15,"Kan Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.g1e2 d5e4 5.a2a3 b4c3 6.e2c3 b8c6
C16,C16,1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4e5
C16,"Petrosian Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4e5 d8d7
C17,C17,1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4e5 c7c5
C17,"Bogoljubow Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4e5 c7c5 5.c1d2
C17,"Icelandic Defense, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4e5 c7c5 5.c1d2 g8e7 6.f2f4
C17,"Moscow Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4e5 c7c5 5.d1g4
C18,"Swiss (Retreat) Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4e5 c7c5 5.a2a3 b4a5
C18,"Armenian Line, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4e5 c7c5 5.a2a3 b4a5 6.b2b4 c5d4
C18,"Maroczy-Willis Variation, French",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4e5 c7c5 5.a2a3 c5d4 6.a3b4 d4c3
C19,C19,1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4e5 c7c5 5.a2a3 b4c3 6.b2c3 g8e7
C19,"Poisoned Pawn Variation, French Winawer",1.e2e4 e7e6 2.d2d4 d7d5 3.b1c3 f8b4 4.e4e5 c7c5 5.a2a3 b4c3 6.b2c3 g8e7 7.d1g4
C20,King Pawn Game; C20,1.e2e4 e7e5
C20,Mengarini Opening,1.e2e4 e7e5 2.a2a3
C20,Portuguese Opening,1.e2e4 e7e5 2.f1b5
C20,Miguel Gambit,1.e2e4 e7e5 2.f1b5 f8c5 3.b2b4
C20,Tortoise Opening,1.e2e4 e7e5 2.f1d3
C20,Macleod Attack,1.e2e4 e7e5 2.c2c3
C20,Norwalder Gambit,1.e2e4 e7e5 2.c2c3 d7d5 3.d1h5 f8d6
C20,Lasa Gambit,1.e2e4 e7e5 2.c2c3 f7f5
C20,Whale Opening,1.e2e4 e7e5 2.c2c4
C20,Bavarian Gambit,1.e2e4 e7e5 2.c2c4 d7d5
C20,Clam Variation; Leonardis Variation,1.e2e4 e7e5 2.d2d3
C20,Indian Opening,1.e2e4 e7e5 2.d2d3
C20,Weber Gambit,1.e2e4 e7e5 2.d2d3 d7d5 3.e4d5 c7c6 4.d5c6 b8c6
C20,King's Gambit Reversedï¿½,1.e2e4 e7e5 2.d2d3 f7f5
C20,Radisch Gambit,1.e2e4 e7e5 2.d2d3 g8f6 3.f2f4 f8c5
C20,Krejcik Gambit,1.e2e4 e7e5 2.f2f3 g8f6 3.f1c4 f8c5 4.g1e2 b8c6 5.b2b4
C20,King's Head Opening,1.e2e4 e7e5 2.f2f3 g8f6 3.b1c3
C20,Alapin Opening,1.e2e4 e7e5 2.g1e2
C20,Napoleon Attack,1.e2e4 e7e5 2.d1f3
C20,Wayward Queen; Danvers; Parnham,1.e2e4 e7e5 2.d1h5
C20,Mellon Gambit,1.e2e4 e7e5 2.d1h5 b8c6 3.f1c4 g8h6 4.d2d3 g7g6 5.h5f3 f7f6 6.g1e2 d7d5
C20,Kiddie Countergambit,1.e2e4 e7e5 2.d1h5 g8f6
C21,Beyer Gambit,1.e2e4 e7e5 2.d2d4 d7d5
C21,Philidor Gambit,1.e2e4 e7e5 2.d2d4 d7d6 3.d4e5 c8d7
C21,Center Game Accepted,1.e2e4 e7e5 2.d2d4 e5d4
C21,Von der Lasa Gambit,1.e2e4 e7e5 2.d2d4 e5d4 3.f1c4
C21,"Ross Gambit, Center",1.e2e4 e7e5 2.d2d4 e5d4 3.f1d3
C21,Danish Gambit,1.e2e4 e7e5 2.d2d4 e5d4 3.c2c3
C21,"Sorensen Defense, Center",1.e2e4 e7e5 2.d2d4 e5d4 3.c2c3 d7d5
C21,Danish Gambit Accepted,1.e2e4 e7e5 2.d2d4 e5d4 3.c2c3 d4c3 4.f1c4 c3b2 5.c1b2
C21,"Copenhagen Defense, Danish Gambit",1.e2e4 e7e5 2.d2d4 e5d4 3.c2c3 d4c3 4.f1c4 c3b2 5.c1b2 f8b4
C21,"Schlechter Defense, Center",1.e2e4 e7e5 2.d2d4 e5d4 3.c2c3 d4c3 4.f1c4 c3b2 5.c1b2 d7d5
C21,"Classical Defense, Danish Gambit",1.e2e4 e7e5 2.d2d4 e5d4 3.c2c3 d4c3 4.f1c4 c3b2 5.c1b2 g8f6
C21,"Chigorin Defense, Danish Gambit",1.e2e4 e7e5 2.d2d4 e5d4 3.c2c3 d4c3 4.f1c4 c3b2 5.c1b2 d8e7
C21,"Svenonius Defense, Center",1.e2e4 e7e5 2.d2d4 e5d4 3.c2c3 g8e7
C21,Halasz-McDonnell Gambit,1.e2e4 e7e5 2.d2d4 e5d4 3.f2f4
C21,"Crocodile Variation, Center Game",1.e2e4 e7e5 2.d2d4 e5d4 3.f2f4 f8c5 4.g1f3 b8c6 5.c2c3
C21,"Lanc-Arnold Gambit, Center",1.e2e4 e7e5 2.d2d4 e5d4 3.g1f3 f8c5 4.c2c3
C21,"Scippler Gambit, Center",1.e2e4 e7e5 2.d2d4 e5d4 3.g1f3 f8c5 4.c2c3 d4c3 5.f1c4
C22,C22,1.e2e4 e7e5 2.d2d4 e5d4 3.d1d4
C22,"Hall Variation, Center",1.e2e4 e7e5 2.d2d4 e5d4 3.d1d4 b8c6 4.d4c4
C22,"Kupreichik Variation, Center",1.e2e4 e7e5 2.d2d4 e5d4 3.d1d4 b8c6 4.d4e3 g8f6 5.b1c3 f8b4 6.c1d2 e8g8 7.e1c1 f8e8 8.f1c4
C23,Bishop's Opening; C23,1.e2e4 e7e5 2.f1c4
C23,Anderssen Gambit,1.e2e4 e7e5 2.f1c4 b7b5 3.c4b5 c7c6
C23,Thorold Gambit,1.e2e4 e7e5 2.f1c4 b7b5 3.c4b5 f7f5
C23,"MacDonnell Gambit, Bishop",1.e2e4 e7e5 2.f1c4 f8c5 3.b2b4
C23,"Labourdonnais-Denker Gambit, Bishop",1.e2e4 e7e5 2.f1c4 f8c5 3.b2b4 c5b4 4.c2c3
C23,"McDonnell Double Gambit, Bishop",1.e2e4 e7e5 2.f1c4 f8c5 3.b2b4 c5b4 4.f2f4
C23,"Philidor Variation, Bishop",1.e2e4 e7e5 2.f1c4 f8c5 3.c2c3
C23,"Lewis Counter-Gambit, Bishop",1.e2e4 e7e5 2.f1c4 f8c5 3.c2c3 d7d5
C23,"Lewis Countergambit, Bishop",1.e2e4 e7e5 2.f1c4 f8c5 3.c2c3 d7d5 4.c4d5 g8f6
C23,"Walker Variation, Bishop",1.e2e4 e7e5 2.f1c4 f8c5 3.c2c3 d7d5 4.c4d5 g8f6 5.d2d4
C23,"Lewis Gambit, Bishop",1.e2e4 e7e5 2.f1c4 f8c5 3.d2d4
C23,"Stein Gambit, Bishop",1.e2e4 e7e5 2.f1c4 f8c5 3.f2f4
C23,"Lopez Variation, Bishop",1.e2e4 e7e5 2.f1c4 f8c5 3.d1e2
C23,"Lopez Gambit, Bishop",1.e2e4 e7e5 2.f1c4 f8c5 3.d1e2 g8f6 4.f2f4
C23,"Khan Gambit, Bishop",1.e2e4 e7e5 2.f1c4 d7d5
C23,"King's Gambit Reversed, Bishop",1.e2e4 e7e5 2.f1c4 f7f5
C23,Calabrese Counter-Gambit,1.e2e4 e7e5 2.f1c4 f7f5
C23,"Kreijcik Gambit, Bishop",1.e2e4 e7e5 2.f1c4 g8f6 3.f2f3 f8c5 4.g1e2 b8c6 5.b2b4
C23,"Greco Gambit, Bishop",1.e2e4 e7e5 2.f1c4 g8f6 3.f2f4
C23,"Horwitz Gambit, Bishop",1.e2e4 e7e5 2.f1c4 g8f6 3.b1c3 b7b5
C23,"Blanel Gambit, Bishop",1.e2e4 e7e5 2.f1c4 g8f6 3.b1c3 f6e4
C23,"Paschmann Gambit, Bishop",1.e2e4 e7e5 2.f1c4 g8f6 3.g1e2 f6e4 4.e2c3
C23,"Meitner-Mieses Gambit, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f1c4 f8c5 4.d1g4 d8f6 5.c3d5
C23,"Bronstein Gambit, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f1c4 g8f6 4.f2f4 f6e4 5.g1f3
C24,"Spielmann Attack, Bishop",1.e2e4 e7e5 2.f1c4 g8f6 3.d2d3 f8c5 4.b1c3
C24,"Kitchner Folly, Bishop",1.e2e4 e7e5 2.f1c4 g8f6 3.d2d3 f8e7 4.g1f3 e8g8
C24,"Vienna Hybrid, Bishop",1.e2e4 e7e5 2.f1c4 g8f6 3.d2d3 b8c6 4.b1c3
C24,"Hromadka Variation, Bishop",1.e2e4 e7e5 2.f1c4 g8f6 3.d2d3 b8c6 4.b1c3 f8b4 5.g1e2
C24,"Warsaw Gambit, Bishop",1.e2e4 e7e5 2.f1c4 g8f6 3.d2d4 e5d4 4.c2c3
C24,Urusov; Ouroussoff; Keidansky G,1.e2e4 e7e5 2.f1c4 g8f6 3.d2d4 e5d4 4.g1f3
C24,"Keidansky Gambit, Bishop",1.e2e4 e7e5 2.f1c4 g8f6 3.d2d4 e5d4 4.g1f3 f6e4 5.d1d4
C25,"Zhuravlev Counter-Gambit, Vienna",1.e2e4 e7e5 2.b1c3 f8b4 3.d1g4
C25,"Zhuraviev Countergambit, Vienna",1.e2e4 e7e5 2.b1c3 f8b4 3.d1g4 g8f6
C25,"Anderssen Defense, Vienna",1.e2e4 e7e5 2.b1c3 f8c5
C25,"Hamppe-Mietner Variation, Vienna",1.e2e4 e7e5 2.b1c3 f8c5 3.c3a4
C25,"Giraffe Attack, Vienna",1.e2e4 e7e5 2.b1c3 f8c5 3.d1g4
C25,"Omaha Gambit, Vienna",1.e2e4 e7e5 2.b1c3 d7d6 3.f2f4
C25,"Fyfe Gambit, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.d2d4
C25,"Philidor Countergambit, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.d2d4 f7f5
C25,"Alapin Variation, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4
C25,Vienna Gambit,1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4
C25,"Quelle Gambit, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4 f8c5 4.f4e5 d7d6
C25,"Steinitz Gambit, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4 e5f4 4.d2d4
C25,"Fraser-Minckwitz Variation, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4 e5f4 4.d2d4 d8h4 5.e1e2 b7b6
C25,"Zukertort Defense, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4 e5f4 4.d2d4 d8h4 5.e1e2 d7d5
C25,"Paulsen Defense, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4 e5f4 4.d2d4 d8h4 5.e1e2 d7d6
C25,"Soerensen Defense, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4 e5f4 4.d2d4 d8h4 5.e1e2 g7g5
C25,"Knight Variation, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4 e5f4 4.g1f3
C25,"Cunningham Defense, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4 e5f4 4.g1f3 f8e7
C25,"Hamppe-Muzio Gambit, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4 e5f4 4.g1f3 g7g5 5.f1c4 g5g4 6.e1g1
C25,"Pierce Gambit, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4 e5f4 4.g1f3 g7g5 5.d2d4
C25,"Hamppe-Allgaier Gambit, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4 e5f4 4.g1f3 g7g5 5.h2h4 g5g4 6.f3g5
C25,Hamppe-Allgaier-Thorold Gambit,1.e2e4 e7e5 2.b1c3 b8c6 3.f2f4 e5f4 4.g1f3 g7g5 5.h2h4 g5g4 6.f3g5 h7h6 7.g5f7 e8f7 8.d2d4
C25,"Paulsen Variation, Vienna",1.e2e4 e7e5 2.b1c3 b8c6 3.g2g3
C25,Mariotti Gambit,1.e2e4 e7e5 2.b1c3 b8c6 3.g2g3 f8c5 4.f1g2 h7h5 5.g1f3 h5h4
C25,"Polluck Gambit, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.g2g3 f8c5 4.f1g2 b8c6 5.g1e2 d7d5 6.e4d5
C25,"Erben Gambit, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.g2g3 d7d5 4.e4d5 c7c6
C26,"Falkbeer Variation, Vienna; C26",1.e2e4 e7e5 2.b1c3 g8f6
C26,"Mengarini Variation, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.a2a3
C26,"Oxford Variation, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f2f4 d7d5 4.f4e5 f6e4 5.d2d3
C26,"Mieses Variation, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.g2g3
C27,"Stanley Variation, Vienna; C27",1.e2e4 e7e5 2.b1c3 g8f6 3.f1c4
C27,"Reversed Spanish, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f1c4 f8b4
C27,"Bishop Variation, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f1c4 f8c5
C27,"Eifel Gambit, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f1c4 f8c5 4.g1e2 b7b5
C27,"Modern Variation, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f1c4 f6e4
C27,"Boden-Kieseritzky Gambit, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f1c4 f6e4 4.g1f3
C27,"Monster Variation, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f1c4 f6e4 4.d1h5 e4d6 5.c4b3 f8e7
C27,"Alekhine Variation, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f1c4 f6e4 4.d1h5 e4d6 5.c4b3 f8e7 6.g1f3 b8c6 7.f3e5
C27,"Adams Gambit, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f1c4 f6e4 4.d1h5 e4d6 5.c4b3 b8c6 6.d2d4
C27,"Frankenstein-Dracula Var, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f1c4 f6e4 4.d1h5 e4d6 5.c4b3 b8c6 6.c3b5 g7g6 7.h5f3 f7f5 8.f3d5 d8e7
C28,"Three Knights Variation, Vienna; C28",1.e2e4 e7e5 2.b1c3 g8f6 3.f1c4 b8c6
C29,"Vienna Gambit, Vienna; C29",1.e2e4 e7e5 2.b1c3 g8f6 3.f2f4
C29,"Steinitz Variation, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f2f4 d7d5 4.d2d3
C29,"Breyer Variation, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f2f4 d7d5 4.f4e5 f6e4 5.g1f3 f8e7
C29,"Kaufmann Variation, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f2f4 d7d5 4.f4e5 f6e4 5.g1f3 c8g4 6.d1e2
C29,"Paulsen Attack, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f2f4 d7d5 4.f4e5 f6e4 5.d1f3
C29,"Bardeleben Variation, Vienna",1.e2e4 e7e5 2.b1c3 g8f6 3.f2f4 d7d5 4.f4e5 f6e4 5.d1f3 f7f5
C30,King's Gambit; C30,1.e2e4 e7e5 2.f2f4
C30,"Heath Variation, KGD; Rotlewi Countergambit",1.e2e4 e7e5 2.f2f4 f8c5 3.g1f3 d7d6 4.b2b4
C30,"Marshall (Euwe) Attack, KGD",1.e2e4 e7e5 2.f2f4 f8c5 3.g1f3 d7d6 4.c2c3 c8g4 5.f4e5 d6e5 6.d1a4
C30,"Rubinstein Countergambit, KGD",1.e2e4 e7e5 2.f2f4 f8c5 3.g1f3 d7d6 4.c2c3 f7f5
C30,"Svenonius Variation, KGD",1.e2e4 e7e5 2.f2f4 f8c5 3.g1f3 d7d6 4.b1c3 g8f6 5.f1c4 b8c6 6.d2d3 c8g4
C30,"Senechaud Counter-Gambit, KGD",1.e2e4 e7e5 2.f2f4 f8c5 3.g1f3 g7g5
C30,"Mafia Defense, KGD",1.e2e4 e7e5 2.f2f4 c7c5
C30,"Panteldakis Countergambit, KGD",1.e2e4 e7e5 2.f2f4 f7f5
C30,"Soller-Zilbermints Gambit, KGD",1.e2e4 e7e5 2.f2f4 f7f6 3.f4e5 b8c6
C30,Zilbermints Double Countergambit; KGD,1.e2e4 e7e5 2.f2f4 g7g5
C30,"Queen's Knight Defense, KGD",1.e2e4 e7e5 2.f2f4 b8c6
C30,"Miles Defense, KGD",1.e2e4 e7e5 2.f2f4 b8c6 3.g1f3 f7f5
C30,"Zilbermints Double Gambit I, KGD",1.e2e4 e7e5 2.f2f4 b8c6 3.g1f3 g7g5
C30,"Petrov's Defense, KGD",1.e2e4 e7e5 2.f2f4 g8f6
C30,"Norwalde (Norwalder) Variation, KGD",1.e2e4 e7e5 2.f2f4 d8f6
C30,"Schubert Variation, KGD",1.e2e4 e7e5 2.f2f4 d8f6 3.b1c3 f6f4 4.d2d4
C30,"Keene Defense, KGD",1.e2e4 e7e5 2.f2f4 d8h4 3.g2g3 h4e7
C31,"Hinrichsen Gambit, KGD",1.e2e4 e7e5 2.f2f4 d7d5 3.d2d4
C31,"Falkbeer Countergambit Accepted, KGD",1.e2e4 e7e5 2.f2f4 d7d5 3.e4d5
C31,"Miles Gambit, KGD",1.e2e4 e7e5 2.f2f4 d7d5 3.e4d5 f8c5
C31,"Nimzovich-Marshall Variation, KGD",1.e2e4 e7e5 2.f2f4 d7d5 3.e4d5 c7c6
C31,"Marshall C-G, KGD",1.e2e4 e7e5 2.f2f4 d7d5 3.e4d5 c7c6
C31,"Pickler Gambit, KGD",1.e2e4 e7e5 2.f2f4 d7d5 3.e4d5 c7c6 4.d5c6 f8c5
C31,"Nimzovich Variation, KGD; Anderssen Attack",1.e2e4 e7e5 2.f2f4 d7d5 3.e4d5 e5e4 4.f1b5
C31,"Blackbunre Attack, KGD",1.e2e4 e7e5 2.f2f4 d7d5 3.g1f3
C32,"Staunton Line, KGD; C32",1.e2e4 e7e5 2.f2f4 d7d5 3.e4d5 e5e4
C32,"Morphy Defense, KGD",1.e2e4 e7e5 2.f2f4 d7d5 3.e4d5 e5e4 4.d2d3 g8f6 5.b1c3 f8b4 6.c1d2 e4e3
C32,"Keres Variation, KGD",1.e2e4 e7e5 2.f2f4 d7d5 3.e4d5 e5e4 4.d2d3 g8f6 5.b1d2
C32,"Old Line, Falkbeer Countergambit, KGD",1.e2e4 e7e5 2.f2f4 d7d5 3.e4d5 e5e4 4.d2d3 g8f6 5.d1e2
C32,"Modern Transfer, Falkbeer Countergambit, KGD",1.e2e4 e7e5 2.f2f4 d7d5 3.e4d5 e5f4
C33,King's Gambit Accepted; C33,1.e2e4 e7e5 2.f2f4 e5f4
C33,"Orsini Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.b2b3
C33,"Schurig Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1b5
C33,"Bishop's Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4
C33,"Kieseritzky Gambit, Bishop's Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 b7b5
C33,"Lopez Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 c7c6
C33,"Bledow; Morphy Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d7d5
C33,"Bledow Countergambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d7d5 4.c4d5 g8f6
C33,"Gianutio Counter-Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 f7f5
C33,"Anderssen Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 g7g5
C33,"Maurian Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 b8c6
C33,Cozio; Hanneken; Prussian; Lichtenhein,1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 g8f6
C33,"Bogoljubow Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 g8f6 4.b1c3
C33,"Bogoljubov; Jaenisch Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 g8f6 4.b1c3 c7c6
C33,"Bryan; Kieseritzky Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d8h4 4.e1f1 b7b5
C33,"Greco Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d8h4 4.e1f1 f8c5
C33,"Cozio Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d8h4 4.e1f1 d7d6
C33,"Classical Defense, KGA; Lopez Variation",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d8h4 4.e1f1 g7g5
C33,"Grimm Attack, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d8h4 4.e1f1 g7g5 5.b1c3 f8g7 6.d2d4 d7d6 7.e4e5
C33,"MacDonnell Attack, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d8h4 4.e1f1 g7g5 5.b1c3 f8g7 6.d2d4 g8e7 7.g2g3
C33,"Fraser Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d8h4 4.e1f1 g7g5 5.b1c3 f8g7 6.g2g3 f4g3 7.d1f3
C33,"Cozio Attack, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d8h4 4.e1f1 g7g5 5.d1f3
C33,"Boden Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d8h4 4.e1f1 b8c6
C33,"Jaenisch Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d8h4 4.e1f1 g8f6
C33,"Jaenisch Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1c4 d8h4 4.e1f1 h4f6
C33,Lesser Bishop; Tartakower,1.e2e4 e7e5 2.f2f4 e5f4 3.f1e2
C33,"Weiss Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.f1e2 f7f5 4.e4f5 d7d6
C33,"Central; Polerio; Willemson, Steinitz",1.e2e4 e7e5 2.f2f4 e5f4 3.d2d4
C33,"Gaga Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g2g3
C33,"Gama Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g2g3 f4g3 4.g1f3
C33,"Leonardo (Stamma) Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.h2h4
C33,"Drunken King, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.e1f2
C33,"Keres, Mason, Parnu, Requena",1.e2e4 e7e5 2.f2f4 e5f4 3.b1c3
C33,"Requena Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.b1c3 d8h4
C33,"Paris Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1e2
C33,"Eisenberg Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1h3
C33,"Basman Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.d1e2
C33,"Bilguer-Mirage Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.d1e2
C33,"Breyer; Hungarian; Carrera, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.d1f3
C33,"Dodo Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.d1g4
C33,"Carrera Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.d1h5
C34,"King's Knight Gambit, KGA; C34",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3
C34,"Fischer Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 d7d6
C34,"Schulder Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 d7d6 4.b2b4
C34,"Gianutio Countergambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 f7f5
C34,"Becker; Anti-Kieseritzky Var, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 h7h6
C34,"MacLeod Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 b8c6
C34,"Bonsch-Osmolovsky Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g8e7
C34,"Schallop Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g8f6
C34,"Tashkent Attack, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g8f6 4.e4e5 f6h5 5.g2g4
C35,"Cunningham Gambit, KGA; C35",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 f8e7
C35,"Bertin; Three Pawns Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 f8e7 4.f1c4 e7h4 5.g2g3
C35,"Neo-Cunningham Gambit, KGA; McCormick Defense",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 f8e7 4.f1c4 g8f6
C36,"Scandinavian Variation, KGA; Modern Defense, KGA; C36",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 d7d5
C36,"Abbazia Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 d7d5 4.e4d5 g8f6
C37,C37,1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5
C37,"Ghulam Kassim; Koch Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4
C37,"Middleton Counter-Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 d7d6 5.e1g1 c8g4 6.h2h3 h7h5 7.h3g4
C37,"Greco-Lolli; Wild Muzio Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.c4f7
C37,"Young Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.c4f7 e8f7 6.e1g1 g4f3 7.d1f3 d8f6 8.d2d4 f6d4 9.c1e3
C37,"Koch; Ghulam Kassim Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.d2d4
C37,Kotov Gambit,1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.d2d4 g4f3 6.c1f4
C37,"Ghulam-Kassim Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.d2d4 g4f3 6.d1f3
C37,"Australian Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.h2h4
C37,"MacDonnell Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.b1c3
C37,"Anderssen Counter-Attack, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.f3e5 d8h4 6.e1f1 g8h6 7.d2d4 d7d6
C37,"Muzio-Polerio Gambit, KGA; Wild Muzio",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.e1g1
C37,"Brentano Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.e1g1 d7d5
C37,"Holloway Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.e1g1 g4f3 6.d1f3 b8c6
C37,"From Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.e1g1 g4f3 6.d1f3 d8e7
C37,"Sarratt Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.e1g1 g4f3 6.d1f3 d8f6
C37,"Young Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.e1g1 g4f3 6.d1f3 d8f6 7.c4f7 e8f7 8.d2d4 f6d4 9.c1e3 d4f6
C37,"Double Muzio Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.e1g1 g4f3 6.d1f3 d8f6 7.e4e5 f6e5 8.c4f7
C37,"Paulsen Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.e1g1 g4f3 6.d1f3 d8f6 7.e4e5 f6e5 8.d2d3 f8h6 9.b1c3 g8e7 10.c1d2
C37,"Bello Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.e1g1 g4f3 6.d1f3 d8f6 7.b1c3
C37,"Baldwin Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.e1g1 g4f3 6.d1f3 d8f6 7.b1c3 f6d4 8.g1h1 d4c4 9.c3d5
C37,"Kling & Horwitz Counter-Attack, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 g5g4 5.e1g1 d8e7
C37,"Blachy Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 b8c6
C37,"Rosentreter Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.d2d4
C37,"Rosentreter-Testa Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.d2d4 g5g4 5.c1f4
C37,"Quaade; Taylor Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.b1c3
C38,"Traditional Variation, KGA; C38",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 f8g7
C38,"Mayet Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 f8g7 5.d2d4 d7d6 6.c2c3
C38,"Philidor Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 f8g7 5.h2h4
C38,"Greco Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 f8g7 5.h2h4 h7h6 6.d2d4 d7d6
C38,"Schultz Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 f8g7 5.h2h4 h7h6 6.d2d4 d7d6 7.d1d3
C38,"Hanstein Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.f1c4 f8g7 5.e1g1
C39,C39,1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4
C39,"Kieseritzky Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5
C39,"Paulsen Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 f8g7
C39,"Campbell; Brentano; Morphy Var, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 d7d5
C39,"Green; Kolisch Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 d7d6
C39,"Long Whip Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 h7h5
C39,"Cotter Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 h7h6 6.e5f7
C39,"Neumann Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 b8c6
C39,"Berlin Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 g8f6
C39,"Anderssen Defense, Kieseritsky Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 g8f6 6.f1c4 d7d5 7.e4d5 f8d6
C39,"Anderssen-Cordel Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 g8f6 6.f1c4 d7d5 7.e4d5 f8d6 8.d2d4 f6h5 9.c1f4 h5f4
C39,"Rice Gambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 g8f6 6.f1c4 d7d5 7.e4d5 f8d6 8.e1g1
C39,"Rice Gambit Accepted, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 g8f6 6.f1c4 d7d5 7.e4d5 f8d6 8.e1g1 d6e5
C39,"Paulsen Defense Deferred, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 g8f6 6.f1c4 d7d5 7.e4d5 f8g7
C39,"Rubinstein Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 g8f6 6.d2d4
C39,"Riviere Variation, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 g8f6 6.e5g4 d7d5
C39,"Rosenthal; Stockwhip Var, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3e5 d8e7
C39,"Allgaier; Cotter Cambit, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3g5
C39,"Walker (Urusov) Attack, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3g5 h7h6 6.g5f7 e8f7 7.f1c4
C39,"Urusov Attack, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3g5 h7h6 6.g5f7 e8f7 7.f1c4 d7d5 8.c4d5 f7e8 9.d2d4
C39,"Thorold Attack, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3g5 h7h6 6.g5f7 e8f7 7.d2d4
C39,"Horny Defense, KGA",1.e2e4 e7e5 2.f2f4 e5f4 3.g1f3 g7g5 4.h2h4 g5g4 5.f3g5 h7h6 6.g5f7 e8f7 7.d1g4 g8f6 8.g4f4
C40,King's Knight Opening; C40,1.e2e4 e7e5 2.g1f3
C40,Glatz Gambit; Busch-Gass Gambit,1.e2e4 e7e5 2.g1f3 f8c5
C40,Chiodini Gambit,1.e2e4 e7e5 2.g1f3 f8c5 3.f3e5 b8c6
C40,Jalalabad Defense,1.e2e4 e7e5 2.g1f3 c7c5
C40,Gunderam Gambit,1.e2e4 e7e5 2.g1f3 c7c6
C40,Queen's Pawn Counter-Gambit; Elephant Gambit,1.e2e4 e7e5 2.g1f3 d7d5
C40,Maroczy; Elephant Gambit,1.e2e4 e7e5 2.g1f3 d7d5 3.e4d5 f8d6
C40,Halasz-Jacobsen Gambit,1.e2e4 e7e5 2.g1f3 d7d5 3.e4d5 f8d6
C40,Maroczy Gambit,1.e2e4 e7e5 2.g1f3 d7d5 3.e4d5 f8d6
C40,Paulsen Counter-Gambit,1.e2e4 e7e5 2.g1f3 d7d5 3.e4d5 e5e4
C40,Wasp Variation,1.e2e4 e7e5 2.g1f3 d7d5 3.f3e5 d5e4 4.f1c4 d8g5
C40,Latvian; Greco Counter-Gambit,1.e2e4 e7e5 2.g1f3 f7f5
C40,"Senechaud Gambit, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.b2b4
C40,"Strautins Gambit, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.f1c4 b7b5
C40,"Poisoned Pawn Variation, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.f1c4 f5e4 4.f3e5 d8g5 5.d2d4 g5g2
C40,"Morgado Defense, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.f1c4 g8f6
C40,"Diepstraten Countergambit, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.c2c4
C40,"Clam Gambit, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.d2d3 b8c6 4.e4f5
C40,"Mason Countergambit, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.d2d4
C40,Latvian Gambit Accepted,1.e2e4 e7e5 2.g1f3 f7f5 3.e4f5
C40,"Lobster Gambit, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.g2g4
C40,"Mlotkowski Variation, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.b1c3
C40,"Fraser Defense, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.f3e5 b8c6
C40,"Corkscrew Counter-Gambit, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.f3e5 g8f6 4.f1c4 f5e4 5.e5f7 d8e7 6.f7h8 d7d5
C40,"Greco Variation, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.f3e5 d8e7
C40,"Mail Line, Latvian Gambit Accepted",1.e2e4 e7e5 2.g1f3 f7f5 3.f3e5 d8f6 4.d2d4
C40,"Bilguer Variation, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.f3e5 d8f6 4.d2d4 d7d6 5.e5c4
C40,"Bronstein Attack, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.f3e5 d8f6 4.d2d4 d7d6 5.e5c4 f5e4 6.f1e2
C40,"Nimzovich Attack, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.f3e5 d8f6 4.d2d4 d7d6 5.e5c4 f5e4 6.c4e3
C40,"Bronstein Gambit, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.f3e5 d8f6 4.d2d4 d7d6 5.e5c4 f5e4 6.d1h5 g7g6 7.h5e2
C40,"Foltys-Leonhardt Variation, Latvian Gambit",1.e2e4 e7e5 2.g1f3 f7f5 3.f3e5 d8f6 4.e5c4
C40,"Foltys Variation, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.f3e5 d8f6 4.e5c4 f5e4 5.d2d3
C40,"Leonhardt Variation, Latvian",1.e2e4 e7e5 2.g1f3 f7f5 3.f3e5 d8f6 4.e5c4 f5e4 5.b1c3
C40,Damiano Defense,1.e2e4 e7e5 2.g1f3 f7f6
C40,Damiano Gambit,1.e2e4 e7e5 2.g1f3 f7f6 3.f3e5 f6e5 4.d1h5 g7g6 5.h5e5 d8e7 6.e5h8
C40,"Chigorin Gambit, Damiano",1.e2e4 e7e5 2.g1f3 f7f6 3.f3e5 d8e7 4.e5f3 d7d5
C40,Levenstein; Queen's Defense; Brazilian; Gunderam,1.e2e4 e7e5 2.g1f3 d8e7
C40,Gunderam Gambit,1.e2e4 e7e5 2.g1f3 d8e7 3.f1c4 f7f5
C40,McConnell Defense,1.e2e4 e7e5 2.g1f3 d8f6
C40,Greco Defense,1.e2e4 e7e5 2.g1f3 d8f6
C40,Labourdonnais Gambit,1.e2e4 e7e5 2.g1f3 d8f6 3.f1c4 f6g6 4.e1g1
C41,Philidor Defense; C41,1.e2e4 e7e5 2.g1f3 d7d6
C41,"Lopez Counter-Gambit, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.f1c4 f7f5
C41,Philidor Gambit,1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 c8d7
C41,"Alapin-Blackburne Gambit, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 c8g4 4.d4e5 b8d7
C41,"Exchange Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 e5d4
C41,"Morphy Gambit, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 e5d4 4.f1c4
C41,"Bird Gambit, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 e5d4 4.c2c3
C41,"Paulsen Attack, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 e5d4 4.f3d4 d6d5 5.e4d5
C41,"Larsen Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 e5d4 4.f3d4 g7g6
C41,"Boden Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 e5d4 4.d1d4 c8d7
C41,Philidor Counter Attack,1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 f7f5
C41,"Zukertort Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 f7f5 4.b1c3
C41,"Hanham; Lord Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 b8d7
C41,"Delmar Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 b8d7 4.f1c4 c7c6 5.c2c3
C41,"Schlechter Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 b8d7 4.f1c4 c7c6 5.b1c3
C41,"Krause Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 b8d7 4.f1c4 c7c6 5.e1g1
C41,"Steiner Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 b8d7 4.f1c4 c7c6 5.e1g1 f8e7 6.d4e5
C41,"Sharp Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 b8d7 4.f1c4 d7b6
C41,"Jaenisch Counter-Attack, Philidor; Nimzovich Variation",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 g8f6
C41,"Klein Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 g8f6 4.f1c4
C41,"Sokolsky Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 g8f6 4.d4e5 f6e4 5.b1d2
C41,"Rellstab Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 g8f6 4.d4e5 f6e4 5.d1d5
C41,"Larobok Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 g8f6 4.b1c3 b8d7
C41,"Locock Variation, Philidor",1.e2e4 e7e5 2.g1f3 d7d6 3.d2d4 g8f6 4.f3g5 h7h6 5.g5f7
C42,Petroff Defense; Russian Defense; C42,1.e2e4 e7e5 2.g1f3 g8f6
C42,"Urosov Gambit, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f1c4
C42,"Boden-Kieseritzky Gambit, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f1c4 f6e4 4.b1c3
C42,"Lichtenhein Defense, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f1c4 f6e4 4.b1c3 d7d5
C42,"Russian Three Knights Var, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.b1c3
C42,"Paulsen Attack, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5c4
C42,"Karklins-Martinovsky Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5d3
C42,"Millenium Attack, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.f1d3
C42,"Kaufmann Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.c2c4
C42,"French Attack, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d2d3
C42,"Classical Attack, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d2d4
C42,"Marshall Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d2d4 d6d5 6.f1d3 f8d6
C42,"Tarrasch Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d2d4 d6d5 6.f1d3 f8d6 7.e1g1 e8g8 8.c2c4 c8g4
C42,"Staunton Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d2d4 d6d5 6.f1d3 f8d6 7.e1g1 e8g8 8.c2c4 c7c6
C42,"Jaenisch Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d2d4 d6d5 6.f1d3 f8e7 7.e1g1 b8c6 8.c2c4
C42,"Browne Attack, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d2d4 d6d5 6.f1d3 f8e7 7.e1g1 b8c6 8.c2c4 c6b4 9.c4d5
C42,"Chigorin Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d2d4 d6d5 6.f1d3 f8e7 7.e1g1 b8c6 8.f1e1
C42,"Berger Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d2d4 d6d5 6.f1d3 f8e7 7.e1g1 b8c6 8.f1e1 c8g4 9.c2c3 f7f5 10.b1d2
C42,"Mason Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d2d4 d6d5 6.f1d3 f8e7 7.e1g1 e8g8
C42,"Mason-Showalter Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d2d4 d6d5 6.f1d3 b8c6
C42,"Nimzovich Attack, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.b1c3
C42,"Lasker Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d1e2
C42,"Cozio (Lasker) Attack, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f3 f6e4 5.d1e2
C42,"Cochrane Gambit, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.e5f7
C42,"Stafford Gambit, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 b8c6
C42,"Damiano Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 f6e4
C42,"Kholmov Gambit, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 f6e4 4.d1e2 d8e7
C42,"Moody Gambit, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.d1e2 b8c6 4.d2d4
C43,"Modern Attack, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.d2d4
C43,"Symmetrical Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.d2d4 d7d5
C43,"Suchting Gambit, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.d2d4 e5d4 4.f1c4
C43,"Tal Gambit, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.d2d4 e5d4 4.e4e5 f6e4 5.f1b5
C43,"Center Attack, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.d2d4 e5d4 4.e4e5 f6e4 5.d1d4
C43,"Trifunovic Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.d2d4 f6e4 4.f1d3 d7d5 5.f3e5 f8d6 6.e1g1 e8g8 7.c2c4 d6e5
C43,"Murret Variation, Petroff",1.e2e4 e7e5 2.g1f3 g8f6 3.d2d4 f6e4 4.f1d3 b8c6
C43,"Steinitz Variation, Petoff; C43",1.e2e4 e7e5 2.g1f3 g8f6 3.f3e5 d7d6 4.d2d4
C44,C44,1.e2e4 e7e5 2.g1f3 b8c6
C44,Paschman Wing Gambit,1.e2e4 e7e5 2.g1f3 b8c6 3.b2b4
C44,Taylor Opening; Inverted Hungarian,1.e2e4 e7e5 2.g1f3 b8c6 3.f1e2
C44,Inverted Hanham Opening,1.e2e4 e7e5 2.g1f3 b8c6 3.f1e2 g8f6 4.d2d3 d7d5 5.b1d2
C44,Basman Gambit,1.e2e4 e7e5 2.g1f3 b8c6 3.f1e2 g8f6 4.d2d4 e5d4 5.e4e5
C44,"Spanish Variation, Ponziani",1.e2e4 e7e5 2.g1f3 b8c6 3.c2c3 d7d5 4.f1b5
C44,"Nikitin Gambit, Ponziani",1.e2e4 e7e5 2.g1f3 b8c6 3.c2c3 d7d5 4.f1b5 d5e4 5.f3e5 d8d5 6.d1a4
C44,"Caro Variation, Ponziani",1.e2e4 e7e5 2.g1f3 b8c6 3.c2c3 d7d5 4.d1a4 c8d7
C44,"Steinitz Variation, Ponziani",1.e2e4 e7e5 2.g1f3 b8c6 3.c2c3 d7d5 4.d1a4 f7f6
C44,"Leonhardt Variation, Ponziani",1.e2e4 e7e5 2.g1f3 b8c6 3.c2c3 d7d5 4.d1a4 g8f6
C44,Ponziani Counter-Gambit,1.e2e4 e7e5 2.g1f3 b8c6 3.c2c3 f7f5
C44,"Schmidt Attack, Ponziani",1.e2e4 e7e5 2.g1f3 b8c6 3.c2c3 f7f5 4.d2d4 d7d6 5.d4d5
C44,"Jaenisch Counterattack, Ponziani",1.e2e4 e7e5 2.g1f3 b8c6 3.c2c3 g8f6
C44,"Neumann Gambit, Ponziani",1.e2e4 e7e5 2.g1f3 b8c6 3.c2c3 g8f6 4.f1c4
C44,"Breyer Opening, Ponziani",1.e2e4 e7e5 2.g1f3 b8c6 3.c2c3 g8f6 4.d2d3 f8e7 5.b2b4
C44,"Vukovic Gambit, Ponziani",1.e2e4 e7e5 2.g1f3 b8c6 3.c2c3 g8f6 4.d2d4 f6e4 5.d4d5 f8c5
C44,Dresden Opening,1.e2e4 e7e5 2.g1f3 b8c6 3.c2c4
C44,Scotch Game,1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4
C44,"MacLopez; Relfsson Gambit, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f1b5
C44,Scotch Gambit,1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f1c4
C44,"London Defense, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f1c4 f8b4
C44,"Vitzhum Attack, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f1c4 f8c5 5.f3g5 g8h6 6.d1h5
C44,"Paulsen; Anderssen C-A, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f1c4 f8c5 5.e1g1 d7d6 6.c2c3 c8g4
C44,"Benima Defense, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f1c4 f8e7
C44,"Goering Gambit, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.c2c3
C44,"Goering Gambit Declined, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.c2c3 d7d5
C44,"Double Pawn Sacrifice, Goering Gambit",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.c2c3 d4c3 5.f1c4
C44,"Bardeleben Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.c2c3 d4c3 5.b1c3 f8b4 6.f1c4 g8f6
C44,"Lolli Attack, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 c6d4
C44,"Napoleon Gambit, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 c6d4 4.f3d4 e5d4 5.f1c4
C44,Konstantinopolsky Opening,1.e2e4 e7e5 2.g1f3 b8c6 3.g2g3
C44,Chicago; Irish; Schulze-Muller Gambit,1.e2e4 e7e5 2.g1f3 b8c6 3.f3e5 c6e5 4.d2d4
C45,C45,1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4
C45,"Malaniuk Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 f8b4
C45,"Classical Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 f8c5
C45,"Gunsberg Defense, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 f8c5 5.c1e3 d8f6 6.c2c3 g8e7 7.f1b5 c6d8
C45,"Fleissig; Meitner Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 f8c5 5.c1e3 d8f6 6.c2c3 g8e7 7.d4c2
C45,"Blackburne Attack, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 f8c5 5.c1e3 d8f6 6.c2c3 g8e7 7.d1d2
C45,"Millennium Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 f8c5 5.c1e3 d8f6 6.c2c3 f6g6
C45,"Blumenfeld Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 f8c5 5.c1e3 d8f6 6.d4b5
C45,"Potter Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 f8c5 5.d4b3
C45,"Romanishin Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 f8c5 5.d4b3 c5b4
C45,"Geller Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 f8c5 5.d4b3 c5b6 6.a2a4 a7a6 7.b1c3 g8f6
C45,"Intermezzo Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 f8c5 5.d4c6 d8f6
C45,"Alekhine Gambit, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 g8f6 5.e4e5
C45,Schmidt Var; Scotch Four Knights,1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 g8f6 5.b1c3
C45,"Mieses Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 g8f6 5.d4c6
C45,"Tartakower Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 g8f6 5.d4c6 b7c6 6.b1d2
C45,"Steinitz; Pulling Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 d8h4
C45,"Braune Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 d8h4 5.c1e3
C45,"Horwitz Attack, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 d8h4 5.d4b5
C45,"Steinitz Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 d8h4 5.b1c3
C45,"Modern Defesne, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 d8h4 5.b1c3 f8b4
C45,"Fraser Attack, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 d8h4 5.d4f3
C45,"Paulsen Variation, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 d8h4 5.d4f5
C45,"Paulsen Attack, Scotch",1.e2e4 e7e5 2.g1f3 b8c6 3.d2d4 e5d4 4.f3d4 f8c5 5.c1e3 d8f6 6.c2c3 g8e7 7.f1b5
C46,Three Knights Opening; C46,1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3
C46,"Schlechter Variation, Three Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 f8b4 4.c3d5 g8f6
C46,"Winawer Defense, Three Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 f7f5
C46,"Steinitz Defense, Three Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g7g6
C46,"Rosenthal Variation, Three Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g7g6 4.d2d4 e5d4 5.c3d5
C46,"Italian Variation, Four Knightsï¿½ï¿½",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1c4
C46,"Halloween Gambit, Four Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f3e5
C47,Four Knights Openng; C47,1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6
C47,"Gunsberg Variation, Four Kngiths",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.a2a3
C47,"Provincial Opening, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.a2a3 d7d6 5.h2h3
C47,"Noa Gambit, Four Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1c4 f6e4 5.c4f7
C47,"Bogoljubow (Scotch) Variation, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.d2d4
C47,"Oxford Gambit, Four Kngihts",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.d2d4 f8b4 5.d4d5 c6d4
C47,"Krause Gambit, Four Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.d2d4 f8b4 5.f3e5
C47,"Leonhardt Defense, Four Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.d2d4 f8b4 5.f3e5 d8e7
C47,"Belgrade Gambit, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.d2d4 e5d4 5.c3d5
C47,"Schmid Gambit, Four Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.d2d4 e5d4 5.f3d4 f6e4
C48,Spanish Four Knights Game; C48,1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5
C48,"Ranken Variation, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 a7a6 5.b5c6
C48,"Classical Variation, Four Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 f8c5
C48,"Marshall Gambit, Four Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 f8c5 5.f3e5 c6d4 6.b5a4 e8g8
C48,"Rubinstein Variation, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 c6d4
C49,"Double Ruy Lopez Opening, 4 Knights; C49",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 f8b4
C49,"Paulsen Variation, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 f8b4 5.e1g1 e8g8 6.b5c6
C49,"Svenonius Variation, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 f8b4 5.e1g1 e8g8 6.d2d3 b4c3 7.b2c3 d7d5
C49,"Janowski Variation, Four Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 f8b4 5.e1g1 e8g8 6.d2d3 b4c3 7.b2c3 d7d6 8.f1e1
C49,"Symmetrical Variation, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 f8b4 5.e1g1 e8g8 6.d2d3 d7d6
C49,"Bernstein Variation, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 f8b4 5.e1g1 e8g8 6.d2d3 d7d6 7.c1g5 b4c3 8.b2c3 c6e7
C49,"Metger Unpin Variation, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 f8b4 5.e1g1 e8g8 6.d2d3 d7d6 7.c1g5 b4c3 8.b2c3 d8e7 9.f1e1 c6d8
C49,"Pillsbury Variation, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 f8b4 5.e1g1 e8g8 6.d2d3 d7d6 7.c1g5 c6e7
C49,"Maroczy Variation, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 f8b4 5.e1g1 e8g8 6.d2d3 d7d6 7.c3e2
C49,"Blackburne Attack, 4 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.b1c3 g8f6 4.f1b5 f8b4 5.e1g1 e8g8 6.c3d5 b4c5 7.d2d4
C50,Italian Game; C50,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4
C50,Giuoco Piano; Italian Opening,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5
C50,Lucchini Gambit; Italian,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.d2d3 f7f5
C50,"Canal Variation, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.d2d3 g8f6 5.b1c3 d7d6 6.c1g5
C50,"Rosentreter Gambit, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.d2d4
C50,"Italian Variation, Four Knights Game",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b1c3 g8f6
C50,"Deutz Gambit, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.e1g1 g8f6 5.d2d4
C50,Hungarian Defense,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8e7
C50,Half Giuoco Piano; Semi-Italian Opening,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 d7d6
C50,Rousseau Gambit,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f7f5
C50,"Shilling (Kostic) Gambit, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 c6d4
C51,Evans Gambit; C51,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4
C51,"Lange Variation, Evans Gambit",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b6 5.b4b5 c6a5 6.f3e5 g8h6
C51,Evans Gambit Accepted,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4
C51,"Normal Position, Evans Gambit",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4c5 6.d2d4 e5d4 7.e1g1 d7d6 8.c3d4 c5b6
C51,"Ulvestad Variation, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4c5 6.d2d4 e5d4 7.e1g1 d7d6 8.c3d4 c5b6 9.d4d5 c6a5 10.c1b2
C51,"Paulsen Variation, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4c5 6.d2d4 e5d4 7.e1g1 d7d6 8.c3d4 c5b6 9.d4d5 c6a5 10.c1b2 g8e7
C51,"Fraser Attack; Mortimer, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4c5 6.d2d4 e5d4 7.e1g1 d7d6 8.c3d4 c5b6 9.b1c3 c8g4 10.d1a4
C51,"Goering Attack, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4c5 6.d2d4 e5d4 7.e1g1 d7d6 8.c3d4 c5b6 9.b1c3 c6a5 10.c1g5
C51,"Stone-Ware Defense, Evans Gambit",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4d6
C51,"Anderssen Variation, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4e7
C51,"Cordel Line, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4e7 6.d2d4 c6a5
C51,"Hein Countergambit, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 d7d5
C52,"Main Line, Evans Gambit; C52",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4a5
C52,"Leonhardt Variation, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4a5 6.d2d4 b7b5
C52,"Sokolsky Variation, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4a5 6.d2d4 d7d6 7.c1g5
C52,"Lasker Defense, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4a5 6.d2d4 d7d6 7.e1g1 a5b6
C52,"Tartakower Attack, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4a5 6.d2d4 d7d6 7.d1b3
C52,"Pierce Defense, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4a5 6.d2d4 e5d4
C52,"Schulen Defense, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4a5 6.d2d4 e5d4 7.e1g1 b7b5
C52,"Waller Attack, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4a5 6.d2d4 e5d4 7.e1g1 d7d6 8.d1b3
C52,"Compromised Defense, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4a5 6.d2d4 e5d4 7.e1g1 d4c3
C52,"Mieses Defense, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4a5 6.d2d4 e5d4 7.e1g1 g8e7
C52,"Slow Variation, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4a5 6.e1g1
C52,"Richardson Attack, Evans",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.b2b4 c5b4 5.c2c3 b4a5 6.e1g1 g8f6 7.d2d4 e8g8 8.f3e5
C53,"Classical Variation, Italian; C53",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3
C53,"Labourdonnais Variation, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 d7d6 5.d2d4 e5d4 6.c3d4 c5b6
C53,"Alexandre Gambit, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 f7f5
C53,Giuoco Pianissimo,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 g8f6 5.d2d3
C53,"Center Attack, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 g8f6 5.d2d4
C53,"Grevo Gambit, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 g8f6 5.d2d4 e5d4 6.c3d4
C53,"Albin Gambit, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 g8f6 5.e1g1
C53,"Lewis Defense, Italian; Closed Variation",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 d8e7
C53,"Mestel Variation, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 d8e7 5.d2d4 c5b6 6.c1g5
C54,C54,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 g8f6
C54,"Bird Attack, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 g8f6 5.b2b4
C54,"Cracow Variation, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 g8f6 5.d2d4 e5d4 6.c3d4 c5b4 7.e1f1
C54,"Moeller (Therkatz) Attack, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 g8f6 5.d2d4 e5d4 6.c3d4 c5b4 7.b1c3 f6e4 8.e1g1 b4c3 9.d4d5
C54,"Therkatz-Herzog, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 g8f6 5.d2d4 e5d4 6.c3d4 c5b4 7.b1c3 f6e4 8.e1g1 b4c3 9.d4d5 c3f6 10.f1e1 c6e7
C54,"Greco Variation, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 g8f6 5.d2d4 e5d4 6.c3d4 c5b4 7.b1c3 f6e4 8.e1g1 e4c3
C54,"Mason Gambit, Italian",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 f8c5 4.c2c3 g8f6 5.d2d4 e5d4 6.e1g1
C55,Two Knights Defense; C55,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6
C55,De Riviere Gambit,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.c2c3
C55,Modern Bishop's Opening,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.d2d3
C56,"Open Variation, Two Knights Defense; C56",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.d2d4
C56,"Perreux Variation, Two Knights Defense",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.d2d4 e5d4 5.f3g5
C56,Scotch Gambit,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.d2d4 e5d4 5.e1g1
C56,Max Lange Attack,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.d2d4 e5d4 5.e1g1 f8c5 6.e4e5
C56,"Loman Defense, 2 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.d2d4 e5d4 5.e1g1 f8c5 6.e4e5 d7d5 7.e5f6 d5c4 8.f1e1 c8e6 9.f3g5 g7g6
C56,"Marshall Variation, 2 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.d2d4 e5d4 5.e1g1 f8c5 6.e4e5 d7d5 7.e5f6 d5c4 8.f1e1 c5e7 9.f3g5 d8d5 10.b1c3 d5f5
C56,"Double Gambit Accepted, Two Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.d2d4 e5d4 5.e1g1 f6e4
C56,"Nachmanson Gambit, 2 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.d2d4 e5d4 5.e1g1 f6e4 6.b1c3
C57,"Prussian Game; Knight Attack, Two Knights; C57",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5
C57,"Traxler; Wilkes-Barre Var, 2 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 f8c5
C57,"Trencianske-Teplice Gambit, Two Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 f8c5 5.c4f7 e8e7 6.d2d4
C57,"Ulvestad Variation, 2 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 b7b5
C57,"Kurkin Gambit, Two Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 b7b5 6.c4f1 h7h6 7.g5f7
C57,"Kloss Gambit, Two Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 c6b4
C57,"Fritz Variation, 2 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 c6d4
C57,"Berliner Variation, 2 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 c6d4 6.c2c3 b7b5 7.c4f1 f6d5 8.g5e4 d8h4
C57,"Lolli Attack, 2 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 f6d5 6.d2d4
C57,"Fried Liver; Fegatello Attack, 2 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 f6d5 6.g5f7
C57,"Ponziani-Steintz Gambit, Two Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 f6e4
C58,"Polerio Defense, Two Knights; C58",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 c6a5
C58,"Bogoljubow Variation, Two Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 c6a5 6.c4b5 c7c6 7.d5c6 b7c6 8.d1f3
C58,"Colman Variation, 2 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 c6a5 6.c4b5 c7c6 7.d5c6 b7c6 8.d1f3 a8b8
C58,"Kieseritzky; Morphy Var, 2 Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 c6a5 6.d2d3
C58,"Yankovich Variation, Two Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 c6a5 6.d2d3 h7h6 7.g5f3 e5e4 8.d1e2 a5c4 9.d3c4 f8c5 10.f3d2
C59,C59,1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 c6a5 6.c4b5 c7c6 7.d5c6 b7c6 8.b5e2
C59,"Suhle Defense, Two Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 c6a5 6.c4b5 c7c6 7.d5c6 b7c6 8.b5e2 h7h6
C59,"Goering Variation, Two Knights",1.e2e4 e7e5 2.g1f3 b8c6 3.f1c4 g8f6 4.f3g5 d7d5 5.e4d5 c6a5 6.c4b5 c7c6 7.d5c6 b7c6 8.b5e2 h7h6 9.g5f3 e5e4 10.f3e5
C60,Ruy Lopez; Spanish Opening; C60,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5
C60,"Bulgarian Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a5
C60,"Rotary-Albany Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 b7b6
C60,"Alapin Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f8b4
C60,"Alapin Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f8b4 4.c2c3 b4a5 5.b5c6 d7c6
C60,"Lucena Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f8e7
C60,"Sawyer's Gambit, Ruy Lopez; Spanish Countergambit",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 d7d5
C60,"Harding Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 d7d5 4.f3e5 d8g5 5.e5c6
C60,"Fricke Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 d7d5 4.f3e5 d8g5 5.e1g1
C60,"Nuremberg Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f7f6
C60,"Brentano Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g7g5
C60,"Barnes; Smyslov; Fianchetto Var, Ruy",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g7g6
C60,"Cozio Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8e7
C60,"Tartakower Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8e7 4.d2d4 e5d4 5.f3d4 g7g6 6.b1c3 f8g7 7.c1e3 e8g8 8.d1d2 d7d5
C60,"Paulsen Variation, Cozio Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8e7 4.b1c3 g7g6
C60,"Vinogradov Var, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 d8e7
C60,"Frankfort; Gunderam Var, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 d8f6
C61,"Bird Defense, Ruy Lopez; C61",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 c6d4
C61,"Paulsen Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 c6d4 4.f3d4 e5d4 5.e1g1 g8e7
C62,"Steinitz Defense, Ruy Lopez; C62",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 d7d6
C62,"Nimzovich Attack, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 d7d6 4.d2d4 c8d7 5.b1c3 g8f6 6.b5c6
C62,"Center Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 d7d6 4.d2d4 e5d4 5.e1g1
C63,"Schliemann; Jaenisch Var, Ruy Lopez; C63",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f7f5
C63,"Exchange Variation, Schliemann Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f7f5 4.b5c6
C63,"Schoenemann Attack, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f7f5 4.d2d4
C63,"Jaenisch Gambit Accepted, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f7f5 4.e4f5
C63,"Kostic Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f7f5 4.b1c3 f5e4 5.c3e4 f8e7
C63,"Moehring Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f7f5 4.b1c3 f5e4 5.c3e4 d7d5 6.f3e5 d5e4 7.e5c6 d8d5
C63,"Tartakower Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f7f5 4.b1c3 f5e4 5.c3e4 g8f6
C63,"Soviet Youth Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f7f5 4.b1c3 c6d4 5.e1g1
C63,"Kostic Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f7f5 4.b1c3 g8f6 5.e4f5 f8c5
C64,"Classical; Cordel Var, Ruy Lopez; C64",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f8c5
C64,"Spanish Wing Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f8c5 4.b2b4
C64,"Charousek Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f8c5 4.c2c3 c5b6
C64,"Konikowski Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f8c5 4.c2c3 d7d5
C64,"Cordel Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f8c5 4.c2c3 f7f5
C64,"Benelux Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f8c5 4.c2c3 g8f6 5.d2d4 c5b6 6.e1g1 e8g8
C64,"Boden Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f8c5 4.c2c3 d8e7
C64,"Zaitsev Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f8c5 4.e1g1 c6d4 5.b2b4
C64,"Zukertort Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 f8c5 4.e1g1 g8f6 5.c2c3
C65,"Berlin Defense, Ruy Lopez; C65",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6
C65,"Kaufmann Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.d2d3 f8c5 5.c1e3
C65,"Anderssen Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.d2d3 d7d6 5.b5c6
C65,"Duras Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.d2d3 d7d6 5.c2c4
C65,"Mortimer Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.d2d3 c6e7
C65,"Mortimer Trap, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.d2d3 c6e7 5.f3e5 c7c6
C65,"Nyholm Attack, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.d2d4 e5d4 5.e1g1
C65,"Beverwijk Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f8c5
C66,"Improved Steinitz Defense, Ruy Lopez; C66",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 d7d6
C66,"Hedgehog Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 d7d6 5.d2d4 c8d7 6.b1c3 f8e7
C66,"Closed Bernstein Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 d7d6 5.d2d4 c8d7 6.b1c3 f8e7 7.c1g5
C66,"Closed Showalter Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 d7d6 5.d2d4 c8d7 6.b1c3 f8e7 7.b5c6
C66,"Tarrasch Trap, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 d7d6 5.d2d4 c8d7 6.b1c3 f8e7 7.f1e1 e8g8
C66,"Closed Wolf Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 d7d6 5.d2d4 c8d7 6.b1c3 e5d4
C66,"Fishing Pole Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6g4
C67,"Rio Gambit Accepted, Ruy Lopez; C67",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6e4
C67,"Rosenthal Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6e4 5.d2d4 a7a6
C67,"Rio de Janeiro Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6e4 5.d2d4 f8e7
C67,"Minckwitz Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6e4 5.d2d4 f8e7 6.d4e5
C67,"Trifunovic Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6e4 5.d2d4 f8e7 6.d1e2 d7d5
C67,"Pillsbury Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6e4 5.d2d4 f8e7 6.d1e2 e4d6 7.b5c6 b7c6 8.d4e5 d6b7 9.b2b3
C67,"Zukertort Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6e4 5.d2d4 f8e7 6.d1e2 e4d6 7.b5c6 b7c6 8.d4e5 d6b7 9.c2c4
C67,"Berlin; Rio de Janerio Var, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6e4 5.d2d4 f8e7 6.d1e2 e4d6 7.b5c6 b7c6 8.d4e5 d6b7 9.b1c3 e8g8 10.f1e1
C67,"Winawer Attack, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6e4 5.d2d4 f8e7 6.d1e2 e4d6 7.b5c6 b7c6 8.d4e5 d6b7 9.f3d4 e8g8
C67,"Cordel Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6e4 5.d2d4 f8e7 6.d1e2 e4d6 7.b5c6 b7c6 8.d4e5 d6f5
C67,"L'Hermet Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6e4 5.d2d4 e4d6
C67,"Westerinen Line, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 g8f6 4.e1g1 f6e4 5.d2d4 e4d6 6.b5c6 d7c6 7.d4e5 d6e4
C68,"Morphy Defense, Ruy Lopez; C68",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6
C68,"Exchange Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5c6
C68,"Lutikov Variation, Ruy Lopez Exchange",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5c6 b7c6
C68,"Keres Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5c6 b7c6 5.b1c3
C68,"Romanovsky Variation, Ruy Lopez Exchange",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5c6 b7c6 5.b1c3 f7f6 6.d2d3
C68,"Alekhjne Variation, Ruy Lopez Exchange",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5c6 d7c6 5.d2d4 e5d4 6.d1d4 d8d4 7.f3d4 f8d6
C68,"Barendregt-Fischer Var, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5c6 d7c6 5.e1g1
C68,"King's Bishop Variation, Ruy Lopez Exchange",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5c6 d7c6 5.e1g1 f8d6
C68,"Alapin Gambit, Ruy Lopez Exchange",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5c6 d7c6 5.e1g1 c8g4 6.h2h3 h7h5
C69,"Gligoric Variation, Ruy Lopez; C69",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5c6 d7c6 5.e1g1 f7f6
C69,"Bronstein Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5c6 d7c6 5.e1g1 d8d6
C70,"Morphy Defense, Ruy Lopez; C70",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4
C70,"Caro Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 b7b5
C70,"Graz Variation, Ruy",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 b7b5 5.a4b3 f8c5
C70,"Furman-Taimanov; Wing Var, Ruy; Norwegian Variation",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 b7b5 5.a4b3 c6a5
C70,"Nightingale Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 b7b5 5.a4b3 c6a5 6.b3f7
C70,"Alapin's Defense Deferred, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 f8b4
C70,"Classical Defense Deferred, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 f8c5
C70,"Schliemann Defense Deferred, Ruy",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 f7f5
C70,"Jaenisch Gambit Deferred, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 f7f5 5.e4f5
C70,"Cozio Defense Deferred, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8e7
C71,Steinitz Defense Deferred; Ruy Lopez; C71,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6
C71,"Modern Steinitz Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6
C71,"Keres Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6 5.c2c4
C71,"Noah's Ark Trap, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6 5.d2d4 b7b5 6.a4b3 e5d4 7.f3d4 c6d4 8.d1d4 c7c5 9.d4d5 c8e6 10.d5c6 e6d7
C72,C72,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6 5.e1g1
C73,C73,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6 5.a4c6 b7c6
C73,"Richter Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6 5.a4c6 b7c6 6.d2d4
C73,"Alapin Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6 5.a4c6 b7c6 6.d2d4 f7f6
C74,C74,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6 5.c2c3
C74,"Siesta Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6 5.c2c3 f7f5
C74,"Kopaiev Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6 5.c2c3 f7f5 6.e4f5 c8f5 7.e1g1
C75,C75,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6 5.c2c3 c8d7
C76,C76,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 d7d6 5.c2c3 c8d7 6.d2d4 g7g6
C77,C77,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6
C77,Treybel; Flohr; Exchange Var Deferred,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.a4c6
C77,"Bayreuth Variation, ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.a4c6
C77,"Jaffe Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.c2c3
C77,"Anderssen Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.d2d3
C77,"Duras Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.d2d3 d7d6 6.c2c4
C77,"Mackenzie Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.d2d4
C77,"Center Attack, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.d2d4 e5d4
C77,"Nimzovich; Tarrasch Var, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.b1c3
C77,"Wormald Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.d1e2
C78,"Normal Variation, Ruy Lopez; C78",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1
C78,"Shannon Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 b7b5 6.a4b3 c8b7 7.f3g5 d7d5
C78,"Wing Attack, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 b7b5 6.a4b3 f8e7 7.a2a4
C78,"Neo-Steinitz Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 b7b5 6.a4b3 d7d6
C78,"Moeller Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8c5
C79,"Central Countergambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 d7d5
C79,"Russian Defense, Ruy Lopez; C79; Lasker Defense",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 d7d6
C79,"Brix Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 g7g6
C80,"Tarrasch; Open Variation, Ruy Lopez; C80",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4
C80,"Schlechter Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.a2a4 c6d4
C80,"Berger Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.a2a4 c6d4 9.f3d4 e5d4 10.b1c3
C80,"Harksen Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.c2c4
C80,"Bernstein Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.b1d2
C80,"Karpov Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.b1d2 e4c5 10.c2c3 d5d4
C80,"Zukertort Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c6e7
C80,"Richter Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.d4d5
C80,"Friess Attack, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.f3e5
C80,"Riga Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 e5d4
C80,"Knoore Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.b1c3
C80,"Tartakower Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d1e2
C80,"Skipworth Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.f1e1 d7d5
C81,Keres; Estonian; Howell; Moscow Var,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6
C81,"Howell Attack, Ruy Lopez; C81",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.d1e2
C81,"Adam Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.d1e2 f8e7 10.c2c4
C81,"Ekstrom Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.d1e2 f8e7 10.f1d1 e8g8 11.c2c4
C82,C82,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.c2c3
C82,"Italian Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.c2c3 f8c5
C82,"St Petersburg Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.c2c3 f8c5 10.b1d2
C82,"Dilworth; Kleczynski Var, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.c2c3 f8c5 10.b1d2 e8g8 11.b3c2
C82,"Motzko Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.c2c3 f8c5 10.d1d3
C82,"Nenarokov Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.c2c3 f8c5 10.d1d3 c6e7
C82,"Berlin Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.c2c3 e4c5
C83,"Malkin Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5
C83,"Breslau Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f6e4 6.d2d4 b7b5 7.a4b3 d7d5 8.d4e5 c8e6 9.c2c3 f8e7 10.f1e1 e8g8 11.f3d4
C84,"Closed Defense, Ruy Lopez; C84",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7
C84,"Center Attack, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.d2d4
C84,"Basque Gambit, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.d2d4 e5d4 7.e4e5 f6e4 8.c2c3
C84,"Lopez Four Knights, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.b1c3
C85,Delayed Exchange Ruy Lopez Deferred; C85,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.a4c6
C86,"Worrall Attack, Ruy Lopez; C86",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.d1e2
C86,"Schlechter Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.d1e2 b7b5 7.a4b3 e8g8 8.c2c3 d7d5
C86,"English Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.d1e2 b7b5 7.a4b3 e8g8 8.c2c3 d7d6 9.d2d4 c8g4
C87,C87,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1
C87,"Averbakh Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 d7d6
C87,"Kecskemet; Timbuktu Var, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 d7d6 7.c2c3 e8g8 8.d2d4 c8d7 9.b1d2 d7e8
C87,"Romanovsky Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 d7d6 7.c2c3 e8g8 8.d2d4 c8d7 9.b1d2 e5d4 10.c3d4 c6b4
C88,C88,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3
C88,"Trajkovic Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 c8b7
C88,"Pseudo-Marshall Var, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 c8b7 8.c2c3 d7d5 9.e4d5 f6d5 10.f3e5 c6e5
C88,"Rosen Attack, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.d2d4
C88,"Anti-Marshall Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 e8g8 8.a2a4
C89,C89,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 e8g8 8.c2c3
C89,"Marshall Attack, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 e8g8 8.c2c3 d7d5
C89,"Steiner Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 e8g8 8.c2c3 d7d5 9.e4d5 e5e4
C89,"Kevitz Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 e8g8 8.c2c3 d7d5 9.e4d5 f6d5 10.f3e5 c6e5 11.e1e5 c7c6 12.b3d5 c6d5 13.d2d4 e7d6 14.e5e3
C89,"Geller Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 e8g8 8.c2c3 d7d5 9.e4d5 f6d5 10.f3e5 c6e5 11.e1e5 c7c6 12.d2d4 e7d6 13.e5e1 d8h4 14.g2g3 h4h3
C89,"Fischer Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 e8g8 8.c2c3 d7d5 9.e4d5 f6d5 10.f3e5 c6e5 11.e1e5 c7c6 12.g2g3
C90,C90,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6
C90,"Balla Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 c6a5 9.b3c2 c7c5 10.d2d4 d8c7
C90,"Suetin Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.a2a3
C90,"Lutikov Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.b3c2
C90,"Pilnik Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.d2d3
C90,"Suetin Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 e8g8 8.c2c3 d7d6 9.a2a3
C90,"Leonhardt Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 e8g8 8.c2c3 c6a5 9.b3c2 c7c5 10.d2d4 d8c7 11.h2h3 a5c6 12.d4d5 c6d8 13.b1d2 g7g5
C91,"Yates Variation, Ruy; C91",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.d2d4
C91,"Bogoljubow; Counterthrust Var, Ruy",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.d2d4 c8g4
C92,C92,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3
C92,"Keres Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 a6a5
C92,"Flohr; Zaitsev; Lenzerheide Var, Ruy",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c8b7
C92,"Kholmov Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c8e6
C92,"Spanish Benoni, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6a5 10.b3c2 c7c5 11.d2d4 d8c7 12.d4d5
C92,"Keres Defense, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 f6d7
C92,"Zaitsev System, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 f8e8
C93,"Smyslov Variation, Ruy Lopez; C93",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 h7h6
C94,"Breyer; Retreat Variation, Ruy Lopez; C94",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6b8
C95,C95,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6b8 10.d2d4
C95,"Borisenko Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6b8 10.d2d4 b8d7
C95,"Arseniev Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6b8 10.d2d4 b8d7 11.c3c4
C95,"Romanishin Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6b8 10.d2d4 b8d7 11.b1d2 c8b7 12.b3c2 f8e8 13.d2f1 e7f8 14.c1g5
C95,"Simagin Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6b8 10.d2d4 b8d7 11.f3h4
C96,C96,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6a5
C96,"Borisenko Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6a5 10.b3c2 c7c5 11.d2d4 a5c6
C96,"Keres Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6a5 10.b3c2 c7c5 11.d2d4 f6d7
C97,"Classical Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6a5 10.b3c2 c7c5 11.d2d4 d8c7 12.b1d2
C98,"Rauzer Variation, Ruy Lopez",1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6a5 10.b3c2 c7c5 11.d2d4 d8c7 12.b1d2 a5c6 13.d4e5 d6e5 14.a2a4
C99,C99,1.e2e4 e7e5 2.g1f3 b8c6 3.f1b5 a7a6 4.b5a4 g8f6 5.e1g1 f8e7 6.f1e1 b7b5 7.a4b3 d7d6 8.c2c3 e8g8 9.h2h3 c6a5 10.b3c2 c7c5 11.d2d4 d8c7 12.b1d2 c5d4
D00,Queen Pawn Game; Close Opening; D00,1.d2d4 d7d5
D00,Sarratt Attack,1.d2d4 d7d5 2.c1f4
D00,"Steinitz Countergambit, Queen Pawn Game",1.d2d4 d7d5 2.c1f4 c7c5
D00,Morris Counter-Gambit,1.d2d4 d7d5 2.c1f4 c7c5 3.e2e4 d5e4
D00,"Bishop Attack; Levitsky Attack, Queen Pawn Game",1.d2d4 d7d5 2.c1g5
D00,Zot,1.d2d4 d7d5 2.c1g5
D00,"Welling Variation, Queen Pawn Game",1.d2d4 d7d5 2.c1g5 c8g4
D00,"Euwe Variation, Queen Pawn Game",1.d2d4 d7d5 2.c1g5 h7h6 3.g5h4 c7c6 4.g1f3 d8b6
D00,"Stonewall Attack, Queen Pawn Game",1.d2d4 d7d5 2.e2e3 g8f6 3.f1d3
D00,Stonewall Attack,1.d2d4 d7d5 2.e2e3 g8f6 3.f1d3 c7c5 4.c2c3
D00,Blackmar-Diemer Gambit,1.d2d4 d7d5 2.e2e4
D00,"Fritz Gambit, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.f1c4
D00,"Diemer-Rosenberg Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.c1e3
D00,"Pohlmann Variation, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.f2f3 f7f5
D00,"Mieses; Nimzovich Var, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.f2f3 b8c6
D00,"Polish Gambit, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3
D00,"Grosshans Defense, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 c8d7
D00,"Zeller Defesne, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 c8f5
D00,"Zeller Defesne, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 c8f5 4.f2f3 e4f3 5.d1f3 f5c8
D00,"Soller Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 c8f5 4.f2f3 g8f6 5.f1c4
D00,"Dries Counter-Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 c7c5
D00,"Lemberger Counter-Gambit, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 e7e5
D00,"Soller; Lemberger Var, BDG; Diemer Attack",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 e7e5 4.c1e3
D00,"Endgame Variation, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 e7e5 4.d4e5
D00,"Rasmussen Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 e7e5 4.g1e2
D00,Lange Gambit; Alfred Lange Gambit; Simple Variation,1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 e7e5 4.c3e4
D00,"Sneiders Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 e7e5 4.d1h5
D00,"Rasa-Studier Variation, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 e7e6 4.c1e3
D00,"Giertz-Peters Gambit, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 e7e6 4.f2f3
D00,"Netherlands Variation, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 f7f5
D00,"Vienna Variation, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6
D00,"Studier-Rasa Gambit, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.c1e3
D00,Von Popiel Gambit,1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.c1g5
D00,"Zilbermints Variation, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.c1g5 c8f5 5.g5f6 e7f6 6.g2g4 f5g6 7.d1e2 f8b4 8.e2b5
D00,Blackmar-Diemer Gambit,1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3
D00,Hans Mullers; Vienna V,1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 c8f5
D00,"Soller; Vienna Var, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 c8f5 5.f1c4
D00,"Tejler Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 c8f5 5.c1f4
D00,"Sperling Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 c8f5 5.d4d5
D00,"Diemer Gambit, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 c8f5 5.f3e4
D00,"Kampars Gambit, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 c8f5 5.g2g4 f5g6 6.g4g5 f6d5 7.f3e4
D00,"Gunderam Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 c8f5 5.g2g4 f5g6 6.h2h4
D00,"Brombacher; Daberitz Var, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 c7c5
D00,"Langeheinecke Variation, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4e3
D00,"Elbert Counter-Gambit, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e7e5
D00,"Weinsbach Declination, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e7e6
D00,"Pfrang Gambit, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e7e6 5.f3e4 f8b4 6.f1d3 b4c3 7.b2c3 f6e4 8.c1e3
D00,"Gunderam; Tartakower Var, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 c8f5
D00,"Teichmann Variation, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 c8g4
D00,"Kaulich Variation, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 c7c5
D00,"Ziegler; Gunderam's 2nd Defense, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 c7c6
D00,"Duthilleuk Gambit, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 e7e6 6.f1d3 b8c6 7.c1g5 f8e7 8.c3e4
D00,"Bogoljubow Variation, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 g7g6
D00,"Mad Dog Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 g7g6 6.f1c4 f8g7 7.h2h4
D00,"Nimzovich Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 g7g6 6.f1c4 f8g7 7.f3e5
D00,"Kloss Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 g7g6 6.f1c4 f8g7 7.e1g1 e8g8 8.g1h1
D00,"Studier Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 g7g6 6.f1c4 f8g7 7.e1g1 e8g8 8.d1e1
D00,"Diemer Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 g7g6 6.c1f4
D00,"Peters Attack, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 g7g6 6.f3e5
D00,"Gunderam Defense, BDG; Rook Pawn Defense",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 h7h5
D00,"Schuttler Defense, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 b8d7
D00,"Pietrowsky Variation, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.g1f3 b8c6
D00,"Ryder Variation, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.d1f3
D00,"Tautvaisis Defense, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.d1f3 b8c6
D00,"Tejler Gambit, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 e4f3 5.d1f3 d8d4 6.c1e3 d4g4 7.f3f2 f6e4 8.c3e4 g4e4 9.e1c1 b8c6 10.f1d3
D00,"Albrecht; Karlshruhle Defense, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 g7g6
D00,"Lamb Defense, BDG",1.d2d4 d7d5 2.e2e4 d5e4 3.b1c3 g8f6 4.f2f3 b8c6
D00,"Mason Attack, Queen Pawn Game",1.d2d4 d7d5 2.f2f4
D00,Zurich Gambit,1.d2d4 d7d5 2.g2g4
D00,Chigorin; Ponziani; Veresov,1.d2d4 d7d5 2.b1c3
D00,"Alburt Defense, QP",1.d2d4 d7d5 2.b1c3 c8f5
D00,Anti-Veresov,1.d2d4 d7d5 2.b1c3 c8g4
D00,"Irish Gambit, QP",1.d2d4 d7d5 2.b1c3 c7c5
D00,"Shaviliuk Gambit, QP",1.d2d4 d7d5 2.b1c3 e7e5
D00,"Dutch System, QP",1.d2d4 d7d5 2.b1c3 f7f5
D00,"Fianchetto Defense, QP",1.d2d4 d7d5 2.b1c3 g7g6 3.g1f3 f8g7
D00,Hubsch Gambit,1.d2d4 d7d5 2.b1c3 g8f6 3.e2e4 f6e4
D00,Amazon Attack,1.d2d4 d7d5 2.d1d3
D00,Siberian Attack,1.d2d4 d7d5 2.d1d3 b8c6 3.b1c3
D01,Betbeder; Veresov; Levitzky; Parisian,1.d2d4 d7d5 2.c2c4 d5c4 3.e2e4 f7f5
D01,Richter-Veresov Opening; D01,1.d2d4 d7d5 2.b1c3 g8f6 3.c1g5
D01,"Veresov Variation, QP",1.d2d4 d7d5 2.b1c3 g8f6 3.c1g5 c8f5 4.g5f6
D01,"Richter Variation, Veresov",1.d2d4 d7d5 2.b1c3 g8f6 3.c1g5 c8f5 4.f2f3
D01,"Classical Defense, Veresov",1.d2d4 d7d5 2.b1c3 g8f6 3.c1g5 e7e6 4.g1f3
D01,"Two Knights System, Veresov",1.d2d4 d7d5 2.b1c3 g8f6 3.c1g5 b8d7 4.g1f3
D01,"Grunfeld Defense, Veresov",1.d2d4 d7d5 2.b1c3 g8f6 3.c1g5 b8d7 4.g1f3 g7g6
D01,"Boyce Defense, Veresov",1.d2d4 d7d5 2.b1c3 g8f6 3.c1g5 f6e4
D02,"Zukertort Variation, QP; D02",1.d2d4 d7d5 2.g1f3
D02,"Anti-Torre, QP",1.d2d4 d7d5 2.g1f3 c8g4
D02,Anti-Torre,1.d2d4 d7d5 2.g1f3 c8g4 3.f3e5 g4f5 4.g2g4
D02,Queen's Gambit Reversed; Krause Variation,1.d2d4 d7d5 2.g1f3 c7c5
D02,"Chandler Gambit, QP",1.d2d4 d7d5 2.g1f3 c7c5 3.g2g3 c5d4 4.f1g2
D02,"Chigorin Variation, QP",1.d2d4 d7d5 2.g1f3 b8c6
D02,"Symmetrical Variation, QP",1.d2d4 d7d5 2.g1f3 g8f6
D02,London System; Queen's Bishop Opening,1.d2d4 d7d5 2.g1f3 g8f6 3.c1f4
D02,"Poisoned Pawn Variation, QP",1.d2d4 d7d5 2.g1f3 g8f6 3.c1f4 c7c5 4.e2e3 d8b6 5.b1c3
D02,"Zilbermints Variation, QP",1.d2d4 d7d5 2.g1f3 g8f6 3.c2c4 b7b5
D02,Pseudo-Catalan,1.d2d4 d7d5 2.g1f3 g8f6 3.g2g3
D02,Barry Attack,1.d2d4 g8f6 2.g1f3 g7g6 3.b1c3 d7d5 4.c1f4 f8g7 5.e2e3 e8g8 6.f1e2
D03,"Torre Attack, QP; D03",1.d2d4 d7d5 2.g1f3 g8f6 3.c1g5
D03,"Breyer Variation, QP",1.d2d4 d7d5 2.g1f3 g8f6 3.c1g5 e7e6 4.e2e3 c7c5 5.c2c3 d8b6
D03,"Grunfeld Variation, QP",1.d2d4 d7d5 2.g1f3 g8f6 3.c1g5 g7g6
D03,"Gossip Variation, QP",1.d2d4 d7d5 2.g1f3 g8f6 3.c1g5 f6e4
D04,Colle System; D04,1.d2d4 d7d5 2.g1f3 g8f6 3.e2e3
D04,"Bishop Defense, QP; Anti-Colle",1.d2d4 d7d5 2.g1f3 g8f6 3.e2e3 c8f5
D05,"Classical Defense, Colle System; D05",1.d2d4 d7d5 2.g1f3 g8f6 3.e2e3 e7e6
D05,"Zukertort Variation, QP",1.d2d4 d7d5 2.g1f3 g8f6 3.e2e3 e7e6 4.f1d3
D05,"Rubinstein Attack, Colle",1.d2d4 d7d5 2.g1f3 g8f6 3.e2e3 e7e6 4.f1d3 c7c5 5.b2b3
D06,Queen's Gambit; Aleppo Gambit; D06,1.d2d4 d7d5 2.c2c4
D06,"Zilbermints Gambit, QGD",1.d2d4 d7d5 2.c2c4 b7b5
D06,"Brieger Defense; Sahovic Defense, QGD; Baltic Defense",1.d2d4 d7d5 2.c2c4 c8f5
D06,"Pseudo-Slav, QGD",1.d2d4 d7d5 2.c2c4 c8f5 3.b1c3 e7e6 4.g1f3 c7c6
D06,"Pseudo-Chigorin, QGD",1.d2d4 d7d5 2.c2c4 c8f5 3.b1c3 e7e6 4.g1f3 b8c6
D06,"Queen Attack, Baltic Defens, QGD",1.d2d4 d7d5 2.c2c4 c8f5 3.b1c3 e7e6 4.d1b3
D06,"Austrian Defense; Summetrical Defense, QGD",1.d2d4 d7d5 2.c2c4 c7c5
D06,"Gusev Countergambit, QGD",1.d2d4 d7d5 2.c2c4 c7c5 3.c4d5 g8f6
D06,"Haberditz Variation, QGD",1.d2d4 d7d5 2.c2c4 c7c5 3.c4d5 g8f6 4.e2e4 f6e4 5.d4c5 d8a5
D06,"Salvio Countergambit, QGD",1.d2d4 d7d5 2.c2c4 c7c5 3.d4c5 d5d4
D06,"Tartakower Gambit, QGD",1.d2d4 d7d5 2.c2c4 b8c6 3.b1c3 e7e5
D06,"Modern Gambit, Chigorin Defense",1.d2d4 d7d5 2.c2c4 b8c6 3.b1c3 g8f6 4.g1f3 d5c4
D06,"Alekhine Variation, Chigorin Defense",1.d2d4 d7d5 2.c2c4 b8c6 3.g1f3 c8g4 4.d1a4
D06,"Lazard Gambit, Chigorin Defense",1.d2d4 d7d5 2.c2c4 b8c6 3.g1f3 e7e5
D06,"Cocquio Counter Gambit, QGD; Tan Gambit",1.d2d4 d7d5 2.c2c4 g8f6 3.c4d5 c7c6
D07,"Chigorin Defense, QGD; D07",1.d2d4 d7d5 2.c2c4 b8c6
D07,"Exchange Variation, Chigorin Defense",1.d2d4 d7d5 2.c2c4 b8c6 3.c4d5 d8d5
D07,Costa's Line,1.d2d4 d7d5 2.c2c4 b8c6 3.c4d5 d8d5 4.e2e3 e7e5 5.b1c3 f8b4 6.c1d2 b4c3 7.d2c3 e5d4 8.g1e2
D07,"Janowski Variation, QGD",1.d2d4 d7d5 2.c2c4 b8c6 3.b1c3 d5c4 4.g1f3
D08,"Albin Countergambit; Cavalotti Var, QGD; D08",1.d2d4 d7d5 2.c2c4 e7e5
D08,"Lasker Trap, QGD",1.d2d4 d7d5 2.c2c4 e7e5 3.d4e5 d5d4 4.e2e3 f8b4 5.c1d2 d4e3
D08,"Normal Line, Albin Countergambit, QGD",1.d2d4 d7d5 2.c2c4 e7e5 3.d4e5 d5d4 4.g1f3
D08,"Tartakower Defense, QGD",1.d2d4 d7d5 2.c2c4 e7e5 3.d4e5 d5d4 4.g1f3 c7c5
D08,"Alapin Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e5 3.d4e5 d5d4 4.g1f3 b8c6 5.b1d2
D09,"Fianchetto Variation, Albin Countergambit; D09",1.d2d4 d7d5 2.c2c4 e7e5 3.d4e5 d5d4 4.g1f3 b8c6 5.g2g3
D10,"Slav Defense; Czech Defense, QGD; D10",1.d2d4 d7d5 2.c2c4 c7c6
D10,"Exchange Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.c4d5
D10,Slav Pawn Game,1.d2d4 d7d5 2.c2c4 c7c6 3.e2e3
D10,"Diemer Gambit, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.e2e4
D10,"Alekhine Attack, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.b1c3 d5c4 4.e2e4
D10,"Winawer Counter Gambit, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.b1c3 e7e5
D10,"Anti-Winawer Gambit, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.b1c3 e7e5 4.e2e4
D11,"Modern Line, Slav; D11",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3
D11,Slav Bishop's Game,1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 c8f5
D11,"Bonet Gambit, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.c1g5
D11,"Quiet Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.e2e3
D11,"Breyer Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1d2
D12,"Schallopp Defense, Slav; D12",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.e2e3 c8f5
D12,"Amsterdam Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.e2e3 c8f5 5.c4d5 c6d5 6.b1c3
D12,"Landau Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.e2e3 c8f5 5.c4d5 c6d5 6.d1b3 d8c8 7.c1d2 e7e6 8.b1a3
D12,"Pin Defense, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.e2e3 c8g4
D13,"Exchange Variation, Slav; D13",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.c4d5
D14,"Symmetrical Line, Slav; D14",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.c4d5 c6d5 5.b1c3 b8c6 6.c1f4 c8f5
D14,"Trifunovic Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.c4d5 c6d5 5.b1c3 b8c6 6.c1f4 c8f5 7.e2e3 e7e6 8.d1b3 f8b4
D15,"Three Knights Variation, Slav; D15",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3
D15,"Chameleon Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 a7a6
D15,"Advance System, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 c6c5
D15,"Two Knights Attack, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4
D15,"Alekhine Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.e2e3
D15,"Alekhine Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.e2e3 b7b5 6.a2a4 b5b4
D15,"Geller Variation; Geller-Tolush Var, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.e2e4
D15,"Schlechter Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 g7g6
D15,"Suchting Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d8b6
D16,"Alapin Variation, Slav; D16",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4
D16,"Steiner Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 c8g4
D16,"Soultanbeieff Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 e7e6
D16,"Smyslov Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 b8a6
D17,Czech Defense; Open Slav Defense; D17,1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 c8f5
D17,"Krause Attack, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 c8f5 6.f3e5
D17,"Wiesbaden Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 c8f5 6.f3e5 e7e6
D17,"Sharp Line, Wiesbaden Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 c8f5 6.f3e5 e7e6 7.f2f3 f8b4 8.e2e4
D17,"Fazekas Gambit, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 c8f5 6.f3e5 b8a6 7.e2e4
D17,"Carlsbad Variation, Slav",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 c8f5 6.f3e5 b8d7 7.e5c4 d8c7 8.g2g3 e7e5
D17,"Baarn Variation, Slav; Bled Attack",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 c8f5 6.f3h4
D18,"Dutch Variation, Slav; D18",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 c8f5 6.e2e3
D18,"Mareschall Variation, Slav; Lasker Variation",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 c8f5 6.e2e3 b8a6
D19,"Main Line, Czech Variation, Slav; D19",1.d2d4 d7d5 2.c2c4 c7c6 3.g1f3 g8f6 4.b1c3 d5c4 5.a2a4 c8f5 6.e2e3 e7e6 7.f1c4 f8b4 8.e1g1 e8g8 9.d1e2
D20,"Old Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.e2e3
D20,"Saduleto Variation, QGA; Central Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.e2e4
D20,"Greco Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.e2e4 b7b5
D20,"Rubinstein Variation, QGAï¿½ï¿½",1.d2d4 d7d5 2.c2c4 d5c4 3.e2e4 c7c5
D20,"Yefimov Gambit, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.e2e4 c7c5 4.d4d5 b7b5
D20,"Linares Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.e2e4 c7c5 4.d4d5 g8f6 5.b1c3 b7b5
D20,"McDonnell Defense, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.e2e4 e7e5
D20,"Somov Gambit, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.e2e4 e7e5 4.f1c4
D20,"Versov; Betbeder; Schwartz Defense, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.e2e4 f7f5
D20,"Modern Defense, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.e2e4 b8c6
D20,"Alekhine System, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.e2e4 g8f6
D21,"Blackburne's Move, QGA; Normal Variation, QGA; D21",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3
D21,"Alekhine Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 a7a6
D21,"Slav Gambit, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 b7b5
D21,"Gunsberg Defense, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 c7c5
D21,"Prianishenmo Gambit, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 c7c5 4.d4d5 g8f6 5.b1c3 e7e6 6.e2e4 e6d5 7.e4e5
D21,"Rosenthal Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 e7e6
D21,"Godes Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 b8d7
D22,D22,1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 a7a6 4.e2e3
D22,"Haberditz Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 a7a6 4.e2e3 b7b5
D22,"Alekhine Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 a7a6 4.e2e3 c8g4
D22,"Borisenko-Furman Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 a7a6 4.e2e4
D23,D23,1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6
D23,"Mannheim Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.d1a4
D24,"Showalter Variation, QGA; D24",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.b1c3
D24,"Bogoljubow Defense, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.b1c3 a7a6 5.e2e4
D25,D25,1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3
D25,"Flohr; Winawer Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 c8e6
D25,"Larsen; Janowski; Modern Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 c8g4
D25,"Smyslov Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 g7g6
D25,Queen's Gambit Accepted Deferred,1.d2d4 d7d5 2.g1f3 g8f6 3.c2c4 d5c4
D26,Traditional System QGA; D26,1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6
D26,"Classical Defense, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6 5.f1c4 c7c5
D26,"Steinitz Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6 5.f1c4 c7c5 6.e1g1 c5d4
D26,"Development Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6 5.f1c4 c7c5 6.e1g1 b8c6
D27,"Rubinstein Variation, QGA; D27",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6 5.f1c4 c7c5 6.e1g1 a7a6
D27,"Furman Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6 5.f1c4 c7c5 6.e1g1 a7a6 7.d4c5 f8c5
D27,"Russian Gambit, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6 5.f1c4 c7c5 6.e1g1 a7a6 7.e3e4
D28,"Alekhine System, Classical Defense, QGA; D28",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6 5.f1c4 c7c5 6.e1g1 a7a6 7.d1e2
D28,"Accelerated Fianchetto, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6 5.f1c4 c7c5 6.e1g1 a7a6 7.d1e2 b7b5
D28,"Counter Fianchetto Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6 5.f1c4 c7c5 6.e1g1 a7a6 7.d1e2 b8c6 8.d4c5
D28,"Four Knights Variation, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6 5.f1c4 c7c5 6.e1g1 a7a6 7.d1e2 b8c6 8.b1c3
D29,D29,1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6 5.f1c4 c7c5 6.e1g1 a7a6 7.d1e2 b7b5 8.c4b3 c8b7
D29,"Smyslov Variation, Classical Defense, QGA",1.d2d4 d7d5 2.c2c4 d5c4 3.g1f3 g8f6 4.e2e3 e7e6 5.f1c4 c7c5 6.e1g1 a7a6 7.d1e2 b7b5 8.c4b3 c8b7 9.f1d1 b8d7 10.b1c3 f8d6
D30,D30,1.d2d4 d7d5 2.c2c4 e7e6
D30,"Pseudo-Tarrasch, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.g1f3 c7c5
D30,"Bishop Attack, Tarrasch",1.d2d4 d7d5 2.c2c4 e7e6 3.g1f3 c7c5 4.c4d5 e6d5 5.c1g5
D30,"Anti-Noteboom Gambit, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.g1f3 c7c6 4.b1c3 d5c4 5.g2g3
D30,"Traditional Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.g1f3 g8f6 4.c1g5
D30,"Vienna Variation, QGDï¿½ï¿½",1.d2d4 d7d5 2.c2c4 e7e6 3.g1f3 g8f6 4.c1g5 f8b4
D30,"Duras Variation, QGD; Capablanca Variation",1.d2d4 d7d5 2.c2c4 e7e6 3.g1f3 g8f6 4.c1g5 h7h6
D30,"Hastings Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.g1f3 g8f6 4.c1g5 h7h6 5.g5f6 d8f6 6.b1c3 c7c6 7.d1b3
D30,"Capablanca Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.g1f3 g8f6 4.c1g5 b8d7 5.e2e3 c7c6 6.b1d2
D30,"Pseudo-Meran; Semi-Slav, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.g1f3 g8f6 4.e2e3 c7c6
D30,"Quiet Variation, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.g1f3 g8f6 4.e2e3 c7c6 5.b1d2
D30,"Semmering Variation, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.g1f3 g8f6 4.e2e3 c7c6 5.b1d2 b8d7 6.f1d3 c6c5
D31,"Baltic Defense, QGD",1.d2d4 d7d5 2.c2c4 c8f5 3.c4d5 f5b1 4.d1a4 c7c6 5.d5c6 b8c6
D31,"Argentinian Gambit, QGD",1.d2d4 d7d5 2.c2c4 c8f5 3.c4d5 f5b1 4.d1a4 c7c6 5.d5c6 b8c6 6.a1b1 d8d4
D31,"Queen's Knight Variation, QGD; D31",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3
D31,"Janowski Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 a7a6
D31,"Alapin Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 b7b6
D31,"Alatortsev; Petrosian, Charousek Var, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 f8e7
D31,"Miladinovic Gambit, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 f8e7 4.e2e4 d5e4 5.f2f3
D31,"Portisch Gambit, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c6 4.e2e3 f7f5 5.g2g4
D31,"Gunderam Gambit, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c6 4.e2e4 d5e4 5.f2f3
D31,"Noteboom Variation, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c6 4.g1f3 d5c4
D31,"Abrahams; Noteboom; Klaus Junge Var, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c6 4.g1f3 d5c4 5.a2a4 f8b4 6.e2e3 b7b5 7.c1d2 a7a5
D31,"Anti-Noteboom Variation, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c6 4.g1f3 d5c4 5.c1g5
D31,"Beliavsky Line, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c6 4.g1f3 d5c4 5.c1g5 f7f6
D32,"Tarrasch Defense, QGD; D32",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5
D32,"Schara-Hennig; Duisburg Gambit, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 c5d4
D32,"Von Hennig Gambit, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 c5d4 5.d1d4 b8c6 6.d4d1 e6d5 7.d1d5 c8e6
D32,"Tarrasch Gambit, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.d4c5 d5d4 6.c3a4 b7b5
D32,"Marshall Gambit, Tarrasch",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.e2e4
D32,"Two Knights Variation, Tarrasch",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3
D32,"Grunfeld Gambit, Tarrasch",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.d4c5 d5d4 7.c3a4 b7b5
D32,"Symmetrical Variation, Tarrasch",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.e2e3 g8f6 5.g1f3 b8c6
D33,"Lodz; Schlecter; Rubinstein Var, QGD; D33",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3
D33,"Folkestone; Swedish Var, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3 c5c4
D33,"Prague Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3 g8f6
D33,"Wagner Variation, Tarrasch",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3 g8f6 7.f1g2 c8g4
D34,D34,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3 g8f6 7.f1g2 f8e7
D34,"Rubinstein; Normal Position, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3 g8f6 7.f1g2 f8e7 8.e1g1 e8g8
D34,"Carlsbad Variation, Tarrasch",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3 g8f6 7.f1g2 f8e7 8.e1g1 e8g8 9.c1g5
D34,"Endgame Variation, Tarrasch",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3 g8f6 7.f1g2 f8e7 8.e1g1 e8g8 9.c1g5 c8e6
D34,"Bogoljubow Variation, Tarrasch",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3 g8f6 7.f1g2 f8e7 8.e1g1 e8g8 9.c1g5 c8e6 10.a1c1 c5c4
D34,"Advance Variation, Tarrasch",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3 g8f6 7.f1g2 f8e7 8.e1g1 e8g8 9.c1g5 c5c4
D34,"Spassky Variation, Tarrasch",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3 g8f6 7.f1g2 f8e7 8.e1g1 e8g8 9.c1g5 c5d4 10.f3d4
D34,"Reti Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3 g8f6 7.f1g2 f8e7 8.e1g1 e8g8 9.d4c5 e7c5 10.c3a4
D34,Classical Tarrasch Gambit,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 c7c5 4.c4d5 e6d5 5.g1f3 b8c6 6.g2g3 g8f6 7.f1g2 f8e7 8.e1g1 e8g8 9.d4c5 d5d4
D35,D35,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6
D35,"Harrwitz; Rubinstein Var, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1f4
D35,"Exchange Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c4d5
D35,"Positional Variation, Exchange Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c4d5 e6d5 5.c1g5
D35,"Saemisch Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c4d5 e6d5 5.g1f3 b8d7 6.c1f4
D36,"Reshevsky Variation, QGD; D36",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c4d5 e6d5 5.c1g5 c7c6 6.d1c2
D37,"Three Knights Variation, QGD; D37",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3
D37,"Harrwitz Attack, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 f8e7 5.c1f4
D38,"Ragozin Variation, QGD; D38",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 f8b4
D38,"Alekhine Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 f8b4 5.d1a4
D39,"Vienna Variation, QGD; D39",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 f8b4 5.c1g5 d5c4
D40,"Semi-Tarrasch Variation, QGD; D40",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c5
D40,"Pillsbury Variation, Semi-Tarrasch",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c5 5.c1g5
D41,"San Sebastian Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c5 5.c4d5 f6d5 6.e2e4 d5c3 7.b2c3 c5d4 8.c3d4 f8b4 9.c1d2 b4d2
D42,D42,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c5 5.c4d5 f6d5 6.e2e3 b8c6 7.f1d3
D43,D43,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6
D43,"Anti-Meran Gambit, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.c1g5
D44,Semi-Slav Defense Accepted; D44,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.c1g5 d5c4
D44,"Botvinnik Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.c1g5 d5c4 6.e2e4 b7b5
D44,"Ekstrom Variation, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.c1g5 d5c4 6.e2e4 b7b5 7.e4e5 h7h6 8.g5h4 g7g5 9.e5f6 g5h4 10.f3e5
D44,"Lilienthal Variation, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.c1g5 d5c4 6.e2e4 b7b5 7.e4e5 h7h6 8.g5h4 g7g5 9.f3g5 h6g5 10.h4g5 b8d7
D44,"Alatortsev System, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.c1g5 d5c4 6.e2e4 b7b5 7.e4e5 h7h6 8.g5h4 g7g5 9.f3g5 f6d5
D44,Anti-Moscow Gambit,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.c1g5 h7h6 6.g5h4
D44,"Stonewall Defense, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 f6e4 6.f1d3 f7f5
D44,"Vienna Variation, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 d5c4
D45,D45,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3
D45,"Acclerated Meran, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 a7a6
D45,"Normal Variation, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7
D45,Rubinstein; anti-Meran,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f3e5
D45,"Stoltz Variation, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.d1c2
D45,"Center Variation, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.d1c2 f8d6 7.e3e4
D45,"Mikhalchishin Line, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.d1c2 f8d6 7.e3e4 d5e4 8.c3e4 f6e4 9.c2e4 e6e5 10.d4e5
D45,"Shabalov Attack, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.d1c2 f8d6 7.g2g4
D46,D46,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3
D46,"Romi Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 f8b4
D46,"Chigorin Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 f8d6
D46,"Bogoljubow Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 f8e7
D47,"Meran Variation, QGD; D47",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 d5c4 7.d3c4
D47,Lundin; Neo-Meran,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 d5c4 7.d3c4 b7b5 8.c4d3 b5b4
D47,"Wade Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 d5c4 7.d3c4 b7b5 8.c4d3 c8b7
D47,"Larsen Variation, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 d5c4 7.d3c4 b7b5 8.c4d3 c8b7 9.e3e4 b5b4 10.c3a4 c6c5
D48,"Reynolds Variation, QGD; D48",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 d5c4 7.d3c4 b7b5 8.c4d3 a7a6
D48,"Pirc Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 d5c4 7.d3c4 b7b5 8.c4d3 a7a6 9.e3e4 b5b4
D48,"Old Variation, Semi-Slav",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 d5c4 7.d3c4 b7b5 8.c4d3 a7a6 9.e3e4 c6c5 10.e4e5
D48,"Old Meran Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 d5c4 7.d3c4 b7b5 8.c4d3 a7a6 9.e1g1
D49,"Blumenfeld Variation, QGD; D49",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 d5c4 7.d3c4 b7b5 8.c4d3 a7a6 9.e3e4 c6c5 10.e4e5 c5d4 11.c3b5
D49,"Sozin Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 d5c4 7.d3c4 b7b5 8.c4d3 a7a6 9.e3e4 c6c5 10.e4e5 c5d4 11.c3b5 d7e5
D49,"Stahlberg Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.g1f3 c7c6 5.e2e3 b8d7 6.f1d3 d5c4 7.d3c4 b7b5 8.c4d3 a7a6 9.e3e4 c6c5 10.e4e5 c5d4 11.c3b5 d7e5 12.f3e5 a6b5
D50,"Modern Variation, QGD; D50",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5
D50,"Uhlmann Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.g1f3 h7h6 6.g5h4 e8g8 7.a1c1 d5c4
D50,Been & Koomen; Dutch & Peruvian,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 c7c5
D50,"Canal; Venice Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 c7c5 5.c4d5 d8b6
D50,"Primitive Pillsbury Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 c7c5 5.g1f3 c5d4 6.d1d4
D51,"Knight Defense, QGD; D51",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 b8d7
D51,"Manhattan; Westphalia Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 b8d7 5.e2e3 f8b4
D51,"Westphalian Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 b8d7 5.e2e3 f8b4 6.g1f3 c7c5
D51,"Alekhine Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 b8d7 5.g1f3 c7c6 6.e2e4
D51,"Alekhine Gambit, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 b8d7 5.g1f3 h7h6 6.g5h4 d5c4
D52,D52,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 b8d7 5.e2e3 c7c6 6.g1f3
D52,"Cambridge Springs; Pillsbury Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 b8d7 5.e2e3 c7c6 6.g1f3 d8a5
D52,"Bogoljubov Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 b8d7 5.e2e3 c7c6 6.g1f3 d8a5 7.f3d2 f8b4 8.d1c2
D52,"Argentine Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 b8d7 5.e2e3 c7c6 6.g1f3 d8a5 7.f3d2 f8b4 8.d1c2 e8g8 9.g5h4
D53,D53,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7
D53,"Heral Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.g5f6
D53,"Lasker Defense, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 f6e4
D54,"Neo-Orthodox Variation, QGD; D54",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.a1c1
D55,D55,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3
D55,"Pillsbury Attack, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b7b6 7.f1d3 c8b7 8.c4d5 e6d5 9.f3e5
D55,"Neo Orthodox Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 h7h6
D55,"Anti-Tartakower Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 h7h6 7.g5f6
D55,"Petrosian Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 h7h6 7.g5f6 e7f6 8.a1c1 c7c6 9.f1d3 b8d7 10.e1g1 d5c4
D56,D56,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 h7h6 7.g5h4
D56,"Lasker Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 h7h6 7.g5h4 f6e4 8.h4e7 d8e7
D56,"Teichmann Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 h7h6 7.g5h4 f6e4 8.h4e7 d8e7 9.d1c2
D57,D57,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 h7h6 7.g5h4 f6e4 8.h4e7 d8e7 9.c4d5
D57,"Bernstein Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 h7h6 7.g5h4 f6e4 8.h4e7 d8e7 9.c4d5 e4c3 10.b2c3
D58,"Tartakower Variation, QGD; D58",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 h7h6 7.g5h4 b7b6
D59,"Exchange Variation, Tartakower Defense, QGD; D59",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 h7h6 7.g5h4 b7b6 8.c4d5 f6d5
D60,"Orthodox Variation, QGD; D60",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7
D60,"Botvinnik Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.f1d3
D60,"Rauzer Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.d1b3
D61,"Rubinstein Variation, Orthodox Defense, QGD; D61",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.d1c2
D62,D62,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.d1c2 c7c5 8.c4d5
D63,D63,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1
D63,"Henneberger; Swiss Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 a7a6
D63,"Capablanca Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 b7b6 8.c4d5 e6d5 9.f1b5
D63,"Pillsbury Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 b7b6 8.c4d5 e6d5 9.f1d3
D63,"Rubinstein Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 c7c6
D64,"Janowski Variation, Orthodox Defense, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 c7c6 8.f1d3 d5c4 9.d3c4 f6d5 10.h2h4
D64,"Rubinstein Attack, Orthodox Defense, QGD;ï¿½ D64",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 c7c6 8.d1c2
D65,D65,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 c7c6 8.d1c2 a7a6 9.c4d5
D66,D66,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 c7c6 8.f1d3
D66,"Fianchetto Variation, Orthodox Defense, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 c7c6 8.f1d3 d5c4 9.d3c4 b7b5
D67,"Capablanca Freeing Maneuver, QGD; D67",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 c7c6 8.f1d3 d5c4 9.d3c4 f6d5
D67,"Classical Variation, QGD",1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 c7c6 8.f1d3 d5c4 9.d3c4 f6d5 10.g5e7 d8e7 11.e1g1 d5c3 12.c1c3
D68,D68,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 c7c6 8.f1d3 d5c4 9.d3c4 f6d5 10.g5e7 d8e7 11.e1g1 d5c3 12.c1c3 e6e5
D69,D69,1.d2d4 d7d5 2.c2c4 e7e6 3.b1c3 g8f6 4.c1g5 f8e7 5.e2e3 e8g8 6.g1f3 b8d7 7.a1c1 c7c6 8.f1d3 d5c4 9.d3c4 f6d5 10.g5e7 d8e7 11.e1g1 d5c3 12.c1c3 e6e5 13.d4e5
D70,King's Indian Defense; West Indian Defense; D70,1.d2d4 g8f6 2.c2c4 g7g6
D70,Danube Gambit; Adorjan Gambit,1.d2d4 g8f6 2.c2c4 g7g6 3.d4d5 b7b5
D70,Danube; Donau Gambit,1.d2d4 g8f6 2.c2c4 g7g6 3.d4d5 b7b5 4.c4b5 a7a6 5.b5a6 c7c6
D70,Nimzovich Attack; Anti-Grunfeld; Alekhine Variation,1.d2d4 g8f6 2.c2c4 g7g6 3.f2f3
D70,Goglidze Attack,1.d2d4 g8f6 2.c2c4 g7g6 3.f2f3 d7d5
D70,Leko Gambit,1.d2d4 g8f6 2.c2c4 g7g6 3.f2f3 e7e5
D70,Kemeri; Neo-Grunfeldï¿½,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5
D70,"Lutikov Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.f2f3
D70,"Murey Attack, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.f2f3 c7c5 5.c4d5 f6d5 6.c3a4
D70,Neo-Grunfeld Defense,1.d2d4 g8f6 2.c2c4 g7g6 3.g1f3 d7d5
D70,Iago Defense,1.d2d4 g8f6 2.c2c4 b8c6
D71,Counterthrust Variation,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7
D71,"Exchange Variation, Neo-Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7 5.c4d5 f6d5
D71,D71,1.d2d4 g8f6 2.c2c4 b8c6 3.g2g3 d7d5
D72,D72,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7 5.c4d5 f6d5 6.e2e4 d5b6
D73,D73,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7 5.g1f3
D74,"Delayed Exchange Variation, Neo-Grunfeld; D74",1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7 5.g1f3 e8g8 6.c4d5 f6d5 7.e1g1
D75,D75,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7 5.g1f3 e8g8 6.c4d5 f6d5 7.e1g1 c7c5
D76,D76,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7 5.g1f3 e8g8 6.c4d5 f6d5 7.e1g1 d5b6
D77,"Classical Variation, Neo-Grunfeld; D77",1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7 5.g1f3 e8g8 6.e1g1
D78,"Original Defesne, Neo-Grunfeld; D78",1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7 5.g1f3 e8g8 6.e1g1 c7c6
D78,"Modern Defense, Neo-Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7 5.g1f3 e8g8 6.e1g1 d5c4
D78,"Polgar Variation, Neo-Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7 5.g1f3 e8g8 6.e1g1 b8c6
D79,D79,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7 5.g1f3 e8g8 6.e1g1 c7c6 7.c4d5
D79,"Ultra-delayed Exchange Variation, Neo-Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 d7d5 4.f1g2 f8g7 5.g1f3 e8g8 6.e1g1 c7c6 7.c4d5 c6d5
D80,Grunfeld Defense; D80,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5
D80,"Stockholm Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c1g5
D80,"Makogonov Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.e2e3 f8g7 5.g1f3 e8g8 6.b2b4
D80,"Opocensky Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.e2e3 f8g7 5.g1f3 e8g8 6.c1d2
D80,"Paris Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.e2e3 f8g7 5.g1f3 e8g8 6.f1d3
D80,"Schlechter Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.e2e3 f8g7 5.g1f3 e8g8 6.f1d3 c7c6
D80,"Flohr Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.e2e3 f8g7 5.g1f3 e8g8 6.f1d3 c7c6 7.e1g1 c8f5
D80,"Smyslov Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.e2e3 f8g7 5.g1f3 e8g8 6.f1d3 c7c6 7.e1g1 c8g4
D80,"Vienna Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.e2e3 f8g7 5.g1f3 e8g8 6.d1b3
D80,"Pachman Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.e2e3 f8g7 5.g1f3 e8g8 6.d1b3 d5c4 7.f1c4 b8d7 8.f3g5
D80,"Botvinnik Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.e2e3 f8g7 5.g1f3 e8g8 6.d1b3 e7e6
D80,"Schlechter Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.e2e3 c7c6
D80,"Kemeri Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g2g3
D80,Gibbon Gambit; Grunfeld Spike,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g2g4
D80,"Zaitsev Gambit, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.h2h4
D81,"Accelerated Russian Variation, Grunfeld; D81",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.d1b3
D83,Grunfeld Gambit; D83,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c1f4 f8g7 5.e2e3 e8g8
D83,"Capablanca Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c1f4 f8g7 5.e2e3 e8g8 6.a1c1
D83,"Botvinnik Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c1f4 f8g7 5.e2e3 e8g8 6.a1c1 c7c5 7.d4c5 c8e6
D83,"Reshevsky Gambit, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c1f4 f8g7 5.e2e3 e8g8 6.a1c1 c7c5 7.d4c5 d8a5
D84,D84,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c1f4 f8g7 5.e2e3 e8g8 6.c4d5 f6d5
D85,"Exchange Variation, Grunfeld; D85",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5
D85,"Exchange Variation, Grunfeld;",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3
D85,"Modern Main Line, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3 f8g7 7.g1f3
D85,"Kramnik's Line, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3 f8g7 7.g1f3 c7c5 8.h2h3
D85,"Kemeri Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.g2g3
D85,Nadanian Attack,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.c3a4
D86,"Classical Main Line, Grunfeld; D86",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3 f8g7 7.f1c4
D86,"Simagin Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3 f8g7 7.f1c4 e8g8 8.g1e2 b7b6 9.h2h4 c8a6
D86,"Simagin Improved Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3 f8g7 7.f1c4 e8g8 8.g1e2 b8c6
D86,"Larsen Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3 f8g7 7.f1c4 e8g8 8.g1e2 d8d7
D86,"Larsen Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3 f8g7 7.f1c4 e8g8 8.g1e2 d8d7 9.e1g1 b7b6
D87,"Spassky Variation, Grunfeld; D87",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3 f8g7 7.f1c4 e8g8 8.g1e2 c7c5
D87,"Seville Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3 f8g7 7.f1c4 e8g8 8.g1e2 c7c5 9.e1g1 b8c6 10.c1e3
D88,D88,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3 f8g7 7.f1c4 e8g8 8.g1e2 c7c5 9.e1g1 b8c6 10.c1e3 c5d4
D89,D89,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3 f8g7 7.f1c4 e8g8 8.g1e2 c7c5 9.e1g1 b8c6 10.c1e3 c5d4 11.c3d4 c8g4 12.f2f3 c6a5 13.c4d3
D89,"Sokolsky Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.c4d5 f6d5 5.e2e4 d5c3 6.b2c3 f8g7 7.f1c4 e8g8 8.g1e2 c7c5 9.e1g1 b8c6 10.c1e3 c5d4 11.c3d4 c8g4 12.f2f3 c6a5 13.c4d3 g4e6 14.d4d5
D90,"Three Knights Variation, Grunfeld; D90",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3
D90,"Flohr Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.d1a4
D91,"Petrosian System, Grunfeld; D91",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.c1g5
D93,D93,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.c1f4 e8g8 6.e2e3
D94,"Burille Variation, Grunfeld; D94",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.e2e3
D94,"Makogonov Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.e2e3 e8g8 6.b2b4
D94,"Opocensky Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.e2e3 e8g8 6.c1d2
D95,D95,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.e2e3 e8g8 6.d1b3
D96,"Russian Variation, Grunfeld; D96",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.d1b3
D97,D97,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.d1b3 d5c4 6.b3c4 e8g8 7.e2e4
D97,"Hungarian Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.d1b3 d5c4 6.b3c4 e8g8 7.e2e4 a7a6
D97,"Levenfish Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.d1b3 d5c4 6.b3c4 e8g8 7.e2e4 b7b6
D97,"Szabo (Boleslavsky) Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.d1b3 d5c4 6.b3c4 e8g8 7.e2e4 c7c6
D97,"Prins Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.d1b3 d5c4 6.b3c4 e8g8 7.e2e4 b8a6
D97,"Simagin (Byrne) Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.d1b3 d5c4 6.b3c4 e8g8 7.e2e4 b8c6
D98,"Smyslov Variation, Grunfeld; D98",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.d1b3 d5c4 6.b3c4 e8g8 7.e2e4 c8g4
D99,D99,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.d1b3 d5c4 6.b3c4 e8g8 7.e2e4 c8g4 8.c1e3 f6d7 9.c4b3
D99,"Yugoslav Variation, Grunfeld",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 d7d5 4.g1f3 f8g7 5.d1b3 d5c4 6.b3c4 e8g8 7.e2e4 c8g4 8.c1e3 f6d7 9.c4b3 c7c5
E00,East Indian Defense; E00,1.d2d4 g8f6 2.c2c4 e7e6
E00,Seirawan Attack,1.d2d4 g8f6 2.c2c4 e7e6 3.c1g5
E00,Catalan Opening,1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3
E00,"Hungarian Gambit, Catalan",1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 e6e5
E00,"Devin Gambit, Catalan",1.d2d4 g8f6 2.c2c4 e7e6 3.g2g4
E01,E01,1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5
E02,E02,1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 d5c4
E03,E03,1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 d5c4 5.d1a4 b8d7 6.a4c4
E03,"Alekhine Variation, Catalan",1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 d5c4 5.d1a4 b8d7 6.a4c4 a7a6 7.c4c2
E04,"Open Defense, Catalan; E04",1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 d5c4 5.g1f3
E05,"Classical Line, Catalan; E05",1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 d5c4 5.g1f3 f8e7
E05,"Modern Sharp Variation, Catalan",1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 d5c4 5.g1f3 b8c6 6.d1a4 f8b4
E06,"Closed Variation, Catalan",1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2
E06,E06,1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 f8e7
E06,"Tarrasch Defense, Catalan",1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 c7c5 5.g1f3 b8c6
E07,E07,1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 b8d7
E07,"Botvinnik Variation, Catalan",1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 b8d7 7.b1c3 c7c6 8.d1d3
E08,E08,1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 b8d7 7.d1c2
E08,"Zagoriansky Variation, Catalan",1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 b8d7 7.d1c2 c7c6 8.f1d1 b7b6 9.a2a4
E09,E09,1.d2d4 g8f6 2.c2c4 e7e6 3.g2g3 d7d5 4.f1g2 f8e7 5.g1f3 e8g8 6.e1g1 b8d7 7.d1c2 c7c6 8.b1d2
E10,Anti-Nimzo-Indian; E10,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3
E10,Dzindzi-Indian Defense,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 a7a6
E10,Blumenfeld Variation,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 c7c5
E10,Blumenfeld Countergambit,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 c7c5 4.d4d5 b7b5
E10,Dus-Chotimursky Variation,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 c7c5 4.d4d5 b7b5 5.c1g5
E10,"Spielmann Variation, Blumenfeld",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 c7c5 4.d4d5 b7b5 5.c1g5 e6d5 6.c4d5 h7h6
E10,Blumenfeld Countergambit Accepted,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 c7c5 4.d4d5 b7b5 5.d5e6 f7e6 6.c4b5 d7d5
E10,"Rubinstein Variation, Blumenfeld",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 c7c5 4.d4d5 b7b5 5.e2e4
E10,Dory (Doery) Defense,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 f6e4
E11,Bogoljubow Defense; Bogo-Indian Defense; E11,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 f8b4
E11,"Nimzovich Variation, Bogo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 f8b4 4.c1d2 d8e7
E11,"Grunfeld Variation, Bogo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 f8b4 4.b1d2
E12,Queen's Indian Defense; E12,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6
E12,"Petrosian Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.a2a3
E12,"Farago Defense, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.a2a3 c8a6 5.d1c2 a6b7
E12,"Murey Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.a2a3 c8b7 5.b1c3 d7d5 6.c4d5 f6d5 7.d1c2 c7c5 8.e2e4 d5c3 9.b2c3 b8c6 10.c1b2 c5d4 11.c3d4 a8c8 12.a1d1 f8d6
E12,"Miles Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.c1f4
E12,"Kasparov Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.b1c3
E12,"Kasparov-Petrosian Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.b1c3 c8b7 5.a2a3
E12,"Polovodin Gambit, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.b1c3 c8b7 5.a2a3 d7d5 6.c4d5 f6d5 7.e2e4
E12,"Hedgehog Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.b1c3 c8b7 5.a2a3 g7g6
E12,"Botvinnik Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.b1c3 c8b7 5.c1g5 h7h6 6.g5h4 g7g5 7.h4g3 f6h5
E13,E13,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.b1c3 c8b7 5.c1g5 h7h6 6.g5h4 f8b4
E14,"Spassky System, QID; E14",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.e2e3
E14,"Averbakh Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.e2e3 c8b7 5.f1d3 c7c5 6.e1g1 f8e7 7.b2b3 e8g8 8.c1b2 c5d4 9.f3d4
E15,"Fianchetto System, QID; E15",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3
E15,"Nimzovich Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8a6
E15,"Quiet Line, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8a6 5.b2b3
E15,"Nimzovich Attack, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8a6 5.d1a4
E15,"Timman's Line, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8a6 5.d1b3
E15,"Saemisch Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 c7c5
E16,"Fianchetto Traditional, QID; E16",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7
E16,"Capablanca Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8b4
E16,"Yates Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8b4 6.c1d2 a7a5
E16,"Riumin Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8b4 6.c1d2 b4e7
E16,"Buerger Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 c7c5 6.d4d5 e6d5 7.f3g5
E16,"Rubinstein Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 c7c5 6.d4d5 e6d5 7.f3h4
E16,"Fine Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 d8c8 6.e1g1 c7c5 7.d4d5
E17,"Traditional Variation, QID; E17",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8e7
E17,Anti-Queen's Indian System,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8e7 6.b1c3
E17,"Opocensky Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8e7 6.b1c3 f6e4 7.c1d2
E17,"Classical Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8e7 6.e1g1
E17,"Euwe Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8e7 6.e1g1 e8g8 7.b2b3
E17,"Taimanov Gambit, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8e7 6.e1g1 e8g8 7.d4d5 e6d5 8.f3d4
E17,"Polugaevsky Gambit, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8e7 6.e1g1 e8g8 7.d4d5 e6d5 8.f3h4
E17,"Kramnik Variation, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8e7 6.e1g1 e8g8 7.f1e1
E18,E18,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8e7 6.e1g1 e8g8 7.b1c3
E18,"Tiviakov Defense, QID",1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8e7 6.e1g1 e8g8 7.b1c3 b8a6
E19,E19,1.d2d4 g8f6 2.c2c4 e7e6 3.g1f3 b7b6 4.g2g3 c8b7 5.f1g2 f8e7 6.e1g1 e8g8 7.b1c3 f6e4 8.d1c2
E20,E20,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3
E20,Nimzo-Indian Defense,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4
E20,"Kmoch Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.f2f3
E20,"Steiner Variation, Nimzo-Indian; Romanishin Variation",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.g2g3
E20,"Romanishin Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.g1f3 c7c5 5.g2g3 e8g8 6.f1g2
E20,"English Hybrid, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.g1f3 c7c5 5.g2g3 e8g8 6.f1g2 c5d4 7.f3d4 d7d5 8.c4d5 f6d5
E20,"Ragozin Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.g1f3 d7d5
E21,"Three Knights Variation, Nimzo-Indian; E21",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.g1f3
E21,"Korchnoi Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.g1f3 c7c5 5.d4d5
E21,"Shocron Gambit, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.g1f3 c7c5 5.d4d5 b7b5
E21,"Euwe Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.g1f3 c7c5 5.d4d5 f6e4
E22,"Spielmann Variation, Nimzo-Indian; E22",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1b3
E23,"Romanovsky Gambit, Nimzo-Indian; E23",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1b3 c7c5 5.d4c5 b8c6
E23,"San Remo Variation, Nimzo-Indian; Stahlberg Variation",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1b3 c7c5 5.d4c5 b8c6 6.g1f3 f6e4 7.c1d2 e4c5
E23,"Stahlberg Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1b3 c7c5 5.d4c5 b8c6 6.g1f3 f6e4 7.c1d2 e4c5 8.b3c2 f7f5 9.g2g3
E23,"Carlsbad Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1b3 c7c5 5.d4c5 b8c6 6.g1f3 f6e4 7.c1d2 e4d2
E23,"Stahlberg Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1b3 c7c5 5.d4c5 b8c6 6.g1f3 f6e4 7.c1d2 e4d2 8.f3d2 e8g8 9.e1c1
E24,"Saemisch Variation, Nimzo-Indian; E24",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3
E24,"Accelerated Saemisch Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3 b4c3 5.b2c3
E24,"Botvinnik Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3 b4c3 5.b2c3 c7c5 6.f2f3 d7d5 7.e2e3 e8g8 8.c4d5 f6d5
E25,E25,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3 b4c3 5.b2c3 c7c5 6.f2f3 d7d5 7.c4d5
E25,"Keres Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3 b4c3 5.b2c3 c7c5 6.f2f3 d7d5 7.c4d5 f6d5 8.d4c5
E25,"Romanovsky Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3 b4c3 5.b2c3 c7c5 6.f2f3 d7d5 7.c4d5 f6d5 8.d4c5 f7f5
E26,"Saemisch Variation, Nimzo-Indian; E26",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3 b4c3 5.b2c3 c7c5 6.e2e3
E26,"O'Kelly Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3 b4c3 5.b2c3 c7c5 6.e2e3 b7b6
E26,"Capablanca Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3 b4c3 5.b2c3 c7c5 6.e2e3 b8c6 7.f1d3 e8g8 8.g1e2 b7b6 9.e3e4 f6e8
E27,E27,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3 b4c3 5.b2c3 e8g8
E28,E28,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3 b4c3 5.b2c3 e8g8 6.e2e3
E29,E29,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3 b4c3 5.b2c3 e8g8 6.e2e3 c7c5 7.f1d3 b8c6
E29,"Capablanca Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.a2a3 b4c3 5.b2c3 e8g8 6.e2e3 c7c5 7.f1d3 b8c6 8.g1e2 b7b6 9.e3e4 f6e8
E30,"Leningrad Variation, Nimzo-Indian; E30",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.c1g5
E30,"Averbakh Gambit, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.c1g5 h7h6 5.g5h4 c7c5 6.d4d5 b7b5
E31,"Benoni Defense, Nimzo-Indian; E31",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.c1g5 h7h6 5.g5h4 c7c5 6.d4d5 d7d6
E32,"Classical Variation, Nimzo-Indian; E32",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2
E33,"Milner-Barry, Zurich Variation, Nimzo-Indian; E33",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 b8c6
E33,"Milner-Barry Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 b8c6 5.g1f3 d7d6
E33,"Old Zurich Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 b8c6 5.g1f3 d7d6 6.a2a3
E33,"Vitolinsh-Adorjan Gambit, Nimzo Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 e8g8 5.a2a3 b4c3 6.c2c3 b7b5
E33,"Keres Defense, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 e8g8 5.a2a3 b4c3 6.c2c3 b7b6
E34,"Noa Variation, Nimzo-Indian; E34",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 d7d5
E35,E35,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 d7d5 5.c4d5 e6d5
E36,E36,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 d7d5 5.a2a3
E36,"Botvinnik Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 d7d5 5.a2a3 b4c3 6.c2c3 b8c6
E37,E37,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 d7d5 5.a2a3 b4c3 6.c2c3 f6e4
E37,"San Remo Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 d7d5 5.a2a3 b4c3 6.c2c3 f6e4 7.c3c2 b8c6 8.e2e3 e6e5
E38,"Berlin Variation, Nimzo-Indian; E38",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 c7c5
E38,"Steiner Variation, Nimzo-Indianï¿½ï¿½",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 c7c5 5.d4c5 b4c3
E39,"Pirc Variation, Nimzo-Indian; E39",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.d1c2 c7c5 5.d4c5 e8g8
E40,"Rubinstein Variation, Nimzo-Indian; E40",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3
E40,"Keres Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 d7d5 5.a2a3 b4c3 6.b2c3 c7c5 7.c4d5 e6d5 8.f1d3 e8g8 9.g1e2 b7b6
E40,"Taimanov Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 b8c6
E40,"Ragozin Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 b8c6 5.g1f3 e8g8 6.f1d3 d7d5 7.e1g1 d5c4
E41,"Huebner Variation, Nimzo-Indian; E41",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 c7c5
E41,"Huebner Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 c7c5 5.f1d3 b8c6 6.g1f3 b4c3
E41,"Panov Attack, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 c7c5 5.g1f3 c5d4 6.e3d4 d7d5
E42,"Rubinstein Variation, Nimzo-Indian; E42",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 c7c5 5.g1e2
E42,"Sherbakov Attack, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 c7c5 5.g1e2 c5d4 6.e3d4 e8g8 7.c4c5
E43,"St. Petersburg Variation, Nimzo-Indian; E43",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 b7b6
E44,"Fischer Variation, Nimzo-Indian; E44",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 b7b6 5.g1e2
E45,"Bronstein Variation, Nimzo-Indian; E45; Byrne Variation",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 b7b6 5.g1e2 c8a6
E46,"Normal Variation, Nimzo-Indian; E46",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8
E46,"Reshevsky Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1e2
E46,"Simagin Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1e2 d7d5 6.a2a3 b4d6
E47,"Bishop Attack, Nimzo-Indian; E47",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.f1d3
E48,"Classical Defense, Nimzo-Indian; E48",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.f1d3 d7d5
E48,"Olafsson Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.f1d3 d7d5 6.g1f3 c7c5 7.e1g1 d5c4 8.d3c4 b7b6 9.d1e2 c8b7 10.f1d1 d8c8
E48,"Bronstein Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.f1d3 d7d5 6.g1f3 c7c5 7.e1g1 d5c4 8.d3c4 c8d7
E48,"Bronstein Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.f1d3 d7d5 6.g1f3 c7c5 7.e1g1 d5c4 8.d3c4 b8d7 9.c3e2
E48,"Smyslov Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.f1d3 d7d5 6.g1f3 c7c5 7.e1g1 d5c4 8.d3c4 d8e7
E49,"Botvinnik System, Nimzo-Indian; E49",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.f1d3 d7d5 6.a2a3 b4c3 7.b2c3
E50,E50,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3
E50,"Huebner Deferred, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 c7c5
E51,"Ragozin Variation, Nimzo-Indian; E51",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 d7d5
E51,"Saemisch Deferred, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 d7d5 6.a2a3
E52,"Schlechter Defense, Nimzo-Indian; E52",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 d7d5 6.f1d3 b7b6
E53,Bernstein Defense; Gligoric System; E53,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 d7d5 6.f1d3 c7c5
E53,"Keres Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 d7d5 6.f1d3 c7c5 7.e1g1 b7b6
E54,E54,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 d7d5 6.f1d3 c7c5 7.e1g1 d5c4 8.d3c4
E55,"Bronstein Variation, Nimzo-Indian; E55",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 d7d5 6.f1d3 c7c5 7.e1g1 d5c4 8.d3c4 b8d7
E55,"Smyslov Variation, Nimzo-Indian",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 d7d5 6.f1d3 c7c5 7.e1g1 d5c4 8.d3c4 d8e7
E56,"Bernstein Variation, Nimzo-Indian; E56",1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 d7d5 6.f1d3 c7c5 7.e1g1 b8c6
E57,E57,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 d7d5 6.f1d3 c7c5 7.e1g1 b8c6 8.a2a3 d5c4 9.d3c4 c5d4
E58,E58,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 d7d5 6.f1d3 c7c5 7.e1g1 b8c6 8.a2a3 b4c3 9.b2c3
E59,E59,1.d2d4 g8f6 2.c2c4 e7e6 3.b1c3 f8b4 4.e2e3 e8g8 5.g1f3 d7d5 6.f1d3 c7c5 7.e1g1 b8c6 8.a2a3 b4c3 9.b2c3 d5c4 10.d3c4
E60,King's Indian Defense; West Indian Defense; E60,1.d2d4 g8f6 2.c2c4 g7g6
E60,"Fianchetto Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3
E60,Counterthrust Variation,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 f8g7 4.f1g2 d7d5
E60,Kemeri Variation,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 f8g7 4.f1g2 d7d5 5.c4d5 f6d5 6.g1f3 e8g8 7.e1g1 c7c5
E60,Yugoslav Variation,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 f8g7 4.f1g2 d7d6 5.g1f3 e8g8 6.b1c3 c7c5
E60,Bronstein Variation,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 f8g7 4.f1g2 d7d6 5.g1f3 e8g8 6.b1c3 c7c6 7.e1g1 d8a5
E60,Panno Variation,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 f8g7 4.f1g2 d7d6 5.g1f3 e8g8 6.b1c3 b8c6 7.e1g1 a7a6 8.h2h3 a8b8
E60,Spassky Variation,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 f8g7 4.f1g2 d7d6 5.g1f3 e8g8 6.b1c3 b8c6 7.e1g1 c8f5
E60,Simagin Variation,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 f8g7 4.f1g2 d7d6 5.g1f3 e8g8 6.b1c3 b8c6 7.e1g1 c8g4
E60,Szabo Variation,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 f8g7 4.f1g2 d7d6 5.g1f3 e8g8 6.b1c3 b8c6 7.e1g1 e7e5
E60,Panno System,1.d2d4 g8f6 2.c2c4 g7g6 3.g2g3 f8g7 4.f1g2 d7d6 5.g1f3 e8g8 6.e1g1 b8c6 7.d4d5 c6a5
E60,"King's Knight Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.g1f3
E60,"Santassiere Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.g1f3 f8g7 4.b2b4
E61,E61,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3
E61,"Smyslov Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.c1g5
E61,"Smyslov Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.c1g5
E62,"Delayed Fianchetto, KID; E62",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3
E62,"Larsen Defense, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 c7c6 7.e1g1 c8f5
E62,"Kavalek Defense, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 c7c6 7.e1g1 d8a5
E62,"Benjamin Defense, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 c7c6 7.e1g1 d8b6
E62,"Carlsbad Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 b8c6
E62,"Lesser Simagin (Spassky) Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 b8c6 7.e1g1 c8f5
E62,"Simagin Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 b8c6 7.e1g1 c8g4
E62,"Uhlmann-Szabo System, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 b8c6 7.e1g1 e7e5
E63,"Panno Variation, KID; E63",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 b8c6 7.e1g1 a7a6
E64,"Yugoslav Variation, KID; E64",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 c7c5
E65,E65,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 c7c5 7.e1g1
E66,E66,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 c7c5 7.e1g1 b8c6 8.d4d5
E68,E68,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 b8d7 7.e1g1 e7e5 8.e2e4
E69,E69,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.g1f3 d7d6 5.g2g3 e8g8 6.f1g2 b8d7 7.e1g1 e7e5 8.e2e4 c7c6 9.h2h3
E70,"Normal Variation, KID; E70",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4
E70,"Accelerated Averbakh Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.c1g5
E71,"Makogonov Variation, KID; E71",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.h2h3
E72,"Deferred Fianchetto, KID; E72",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g2g3
E72,"Pomar System, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g2g3 e8g8 6.f1g2 e7e5 7.g1e2
E73,E73,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f1e2
E73,"Semi-Averbakh System, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f1e2 e8g8 6.c1e3
E73,"Averbakh Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f1e2 e8g8 6.c1g5
E73,"Spanish Defense, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f1e2 e8g8 6.c1g5 a7a6
E73,"Flexible Defense, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f1e2 e8g8 6.c1g5 h7h6
E73,"Burgess Line, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f1e2 e8g8 6.c1g5 b8a6 7.d1d2 c7c6
E73,"Geller Defense, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f1e2 e8g8 6.c1g5 b8d7
E74,"Benoni Defense, KID; E74",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f1e2 e8g8 6.c1g5 c7c5
E76,"Four Pawns Attack, KID; E76",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f4
E76,"Six Pawns Attack, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f4 c7c5 6.d4d5 e8g8 7.f1e2 e7e6 8.d5e6 f7e6 9.g2g4 b8c6 10.h2h4
E76,"Modern Defense, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f4 b8a6
E76,"Dynamic Attack, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f4 e8g8 6.g1f3 c7c5 7.d4d5
E77,"Six Pawns Attack, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6
E77,E77,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f4 e8g8 6.f1e2
E77,"Florentine Gambit, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f4 e8g8 6.f1e2 c7c5 7.d4d5 e7e6 8.g1f3 e6d5 9.e4e5
E78,"Fluid Attack, KID; E78",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f4 e8g8 6.f1e2 c7c5 7.g1f3
E79,"Exchange Variation, Four Pawns Attack, KID; E79",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f4 e8g8 6.f1e2 c7c5 7.g1f3 c5d4 8.f3d4 b8c6 9.c1e3
E80,"Saemisch Variation, KID; E80",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3
E80,"Bronstein Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e7e5 6.d4d5 f6h5 7.c1e3 b8a6 8.d1d2 d8h4 9.g2g3 h5g3 10.d2f2 g3f1 11.f2h4 f1e3 12.e1f2 e3c4
E81,"Normal Defesne, Sameisch Variation, KID; E81",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8
E81,"Byrne Defense, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1e3 c7c6 7.f1d3 a7a6
E81,"Steiner Attack, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1g5
E81,"Bobotsov-Korchnoi-Petrosian Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.g1e2
E82,"Double Fianchetto, KID; E82",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1e3 b7b6
E83,"Yates Defense, KID; E83",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1e3 b8c6
E83,"Panno Formation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1e3 b8c6 7.g1e2 a7a6
E83,"Ruban Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1e3 b8c6 7.g1e2 a8b8
E84,E84,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1e3 b8c6 7.g1e2 a7a6 8.d1d2 a8b8
E85,"Orthodox Variation, KID; E85",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1e3 e7e5
E86,E86,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1e3 e7e5 7.g1e2 c7c6
E87,"Closed Variation, KID; E87",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1e3 e7e5 7.d4d5
E87,"Bronstein Defense, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1e3 e7e5 7.d4d5 f6h5 8.d1d2 d8h4 9.g2g3 h5g3 10.d2f2 g3f1
E88,E88,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1e3 e7e5 7.d4d5 c7c6
E89,E89,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.f2f3 e8g8 6.c1e3 e7e5 7.d4d5 c7c6 8.g1e2
E90,E90,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3
E90,"Larsen Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.c1e3
E90,"Zinnowitz Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.c1g5
E91,"Orthodox Variation, KID; E91",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2
E92,E92,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5
E92,"Gligoric-Taimanov Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.c1e3
E92,"Petrosian Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.d4d5
E92,"Stein Defense, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.d4d5 a7a5
E93,E93,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.d4d5 b8d7
E93,"Keres Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.d4d5 b8d7 8.c1g5 h7h6 9.g5h4 g6g5 10.h4g3 f6h5 11.h2h4
E94,E94,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1
E94,"Ukranian Defense, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 a7a5
E94,"Donner Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 c7c6
E94,"Glek Defense, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8a6
E95,E95,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8d7 8.f1e1
E96,E96,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8d7 8.f1e1 c7c6 9.e2f1 a7a5
E97,"Aronin-Taimanov Variation, KID; E97",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8c6
E97,"Normal Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8c6 8.d4d5 c6e7
E97,"Bayonet Attack, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8c6 8.d4d5 c6e7 9.b2b4
E97,"Yepishin's Line, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8c6 8.d4d5 c6e7 9.b2b4 f6h5 10.d1c2
E97,"Sokolov's Line, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8c6 8.d4d5 c6e7 9.b2b4 f6h5 10.f1e1
E97,"Korchnoi Attack, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8c6 8.d4d5 c6e7 9.c1d2
E98,E98,1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8c6 8.d4d5 c6e7 9.f3e1
E98,"Kozul Gambit, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8c6 8.d4d5 c6e7 9.f3e1 f6d7 10.c1e3 f7f5
E99,"Benko Attack, KID; E99",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8c6 8.d4d5 c6e7 9.f3e1 f6d7 10.f2f3
E99,"Benko Variation, KID",1.d2d4 g8f6 2.c2c4 g7g6 3.b1c3 f8g7 4.e2e4 d7d6 5.g1f3 e8g8 6.f1e2 e7e5 7.e1g1 b8c6 8.d4d5 c6e7 9.f3e1 f6d7 10.f2f3 f7f5 11.g2g4`)
