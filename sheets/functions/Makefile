# generated automatically by sheet.go

make:
	make functions.pdf

functions.tex: sheet.tex
	sheet -mode c -in sheet.tex > functions.tex

functions.pdf: ../../*.tex ../../trademark.pdf *.tex functions.tex
	pdflatex functions.tex
	make defs


open: functions.pdf
	open functions.pdf

o:
	make open

defs: functions.tex
	defs functions.tex

spell:
	aspell -c functions.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile