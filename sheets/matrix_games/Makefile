# generated automatically by sheet.go

make:
	make matrix_games.pdf

matrix_games.tex: sheet.tex
	sheet -mode c -in sheet.tex > matrix_games.tex

matrix_games.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_games.tex
	pdflatex matrix_games.tex
	make defs


open: matrix_games.pdf
	open matrix_games.pdf

o:
	make open

defs: matrix_games.tex
	defs matrix_games.tex

spell:
	aspell -c matrix_games.tex

s:
	make spell