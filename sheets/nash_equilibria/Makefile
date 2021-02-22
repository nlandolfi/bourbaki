# generated automatically by sheet.go

make:
	make nash_equilibria.pdf

nash_equilibria.tex: sheet.tex
	sheet -mode c -in sheet.tex > nash_equilibria.tex

nash_equilibria.pdf: ../../*.tex ../../trademark.pdf *.tex nash_equilibria.tex
	pdflatex nash_equilibria.tex
	make defs


open: nash_equilibria.pdf
	open nash_equilibria.pdf

o:
	make open

defs: nash_equilibria.tex
	defs nash_equilibria.tex

spell:
	aspell -c nash_equilibria.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile