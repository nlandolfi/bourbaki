# generated automatically by sheet.go

make:
	make orthogonal_complements.pdf

orthogonal_complements.tex: sheet.tex
	sheet -mode c -in sheet.tex > orthogonal_complements.tex

orthogonal_complements.pdf: ../../*.tex ../../trademark.pdf *.tex orthogonal_complements.tex
	pdflatex orthogonal_complements.tex
	make defs


open: orthogonal_complements.pdf
	open orthogonal_complements.pdf

o:
	make open

defs: orthogonal_complements.tex
	defs orthogonal_complements.tex

spell:
	aspell -c orthogonal_complements.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile