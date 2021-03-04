# generated automatically by sheet.go

make:
	make inductions.pdf

inductions.tex: sheet.tex
	sheet -mode c -in sheet.tex > inductions.tex

inductions.pdf: ../../*.tex ../../trademark.pdf *.tex inductions.tex
	pdflatex inductions.tex
	make defs


open: inductions.pdf
	open inductions.pdf

o:
	make open

defs: inductions.tex
	defs inductions.tex

spell:
	aspell -c inductions.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile