identity.tex: sheet.tex
	sheet -mode c -in sheet.tex > identity.tex

identity.pdf: ../../*.tex ../../trademark.pdf *.tex identity.tex
	pdflatex identity.tex
	make defs

make:
	make identity.pdf

open: identity.pdf
	open identity.pdf

o:
	make open

defs: identity.tex
	defs identity.tex

spell:
	aspell -c identity.tex

s:
	make spell