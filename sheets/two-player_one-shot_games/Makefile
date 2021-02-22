# generated automatically by sheet.go

make:
	make two-player_one-shot_games.pdf

two-player_one-shot_games.tex: sheet.tex
	sheet -mode c -in sheet.tex > two-player_one-shot_games.tex

two-player_one-shot_games.pdf: ../../*.tex ../../trademark.pdf *.tex two-player_one-shot_games.tex
	pdflatex two-player_one-shot_games.tex
	make defs


open: two-player_one-shot_games.pdf
	open two-player_one-shot_games.pdf

o:
	make open

defs: two-player_one-shot_games.tex
	defs two-player_one-shot_games.tex

spell:
	aspell -c two-player_one-shot_games.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile