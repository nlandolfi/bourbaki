sequences.tex: sheet.tex
	sheet -mode c -in sheet.tex > sequences.tex

sequences.pdf: ../../*.tex ../../trademark.pdf *.tex sequences.tex
	pdflatex sequences.tex
	make defs

make:
	make sequences.pdf

open: sequences.pdf
	open sequences.pdf

o:
	make open

defs: sequences.tex
	defs sequences.tex

spell:
	aspell -c sequences.tex

s:
	make spell