# generated automatically by sheet.go

make:
	make zero_sum_matrix_games.pdf

zero_sum_matrix_games.tex: sheet.tex
	sheet -mode c -in sheet.tex > zero_sum_matrix_games.tex

zero_sum_matrix_games.pdf: ../../*.tex ../../trademark.pdf *.tex zero_sum_matrix_games.tex
	pdflatex zero_sum_matrix_games.tex
	make defs


open: zero_sum_matrix_games.pdf
	open zero_sum_matrix_games.pdf

o:
	make open

defs: zero_sum_matrix_games.tex
	defs zero_sum_matrix_games.tex

spell:
	aspell -c zero_sum_matrix_games.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile