# generated automatically by sheet.go

make:
	make subset_algebras.pdf

subset_algebras.tex: sheet.tex
	sheet -mode c -in sheet.tex > subset_algebras.tex

subset_algebras.pdf: ../../*.tex ../../trademark.pdf *.tex subset_algebras.tex
	pdflatex subset_algebras.tex
	make defs


open: subset_algebras.pdf
	open subset_algebras.pdf

o:
	make open

defs: subset_algebras.tex
	defs subset_algebras.tex

spell:
	aspell -c subset_algebras.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile