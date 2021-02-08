# generated automatically by sheet.go

make:
	make distributions.pdf

distributions.tex: sheet.tex
	sheet -mode c -in sheet.tex > distributions.tex

distributions.pdf: ../../*.tex ../../trademark.pdf *.tex distributions.tex
	pdflatex distributions.tex
	make defs


open: distributions.pdf
	open distributions.pdf

o:
	make open

defs: distributions.tex
	defs distributions.tex

spell:
	aspell -c distributions.tex

s:
	make spell
