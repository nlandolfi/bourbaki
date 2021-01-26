sets.tex: sheet.tex
	sheet -mode c -in sheet.tex > sets.tex

sets.pdf: ../../*.tex ../../trademark.pdf *.tex sets.tex
	pdflatex sets.tex
	make defs

make:
	make sets.pdf

open: sets.pdf
	open sets.pdf

o:
	make open

defs: sets.tex
	defs sets.tex

spell:
	aspell -c sets.tex

s:
	make spell