# generated automatically by sheet.go

make:
	make identity_matrices.pdf

identity_matrices.tex: sheet.tex
	sheet -mode c -in sheet.tex > identity_matrices.tex

identity_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex identity_matrices.tex
	pdflatex identity_matrices.tex
	make defs


open: identity_matrices.pdf
	open identity_matrices.pdf

o:
	make open

defs: identity_matrices.tex
	defs identity_matrices.tex

spell:
	aspell -c identity_matrices.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile