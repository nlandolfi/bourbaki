# generated automatically by sheet.go

make:
	make positive_definite_matrices.pdf

positive_definite_matrices.tex: sheet.tex
	sheet -mode c -in sheet.tex > positive_definite_matrices.tex

positive_definite_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex positive_definite_matrices.tex
	pdflatex positive_definite_matrices.tex
	make defs


open: positive_definite_matrices.pdf
	open positive_definite_matrices.pdf

o:
	make open

defs: positive_definite_matrices.tex
	defs positive_definite_matrices.tex

spell:
	aspell -c positive_definite_matrices.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile