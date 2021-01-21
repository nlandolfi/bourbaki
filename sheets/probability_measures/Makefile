probability_measures.tex: sheet.tex
	sheet -mode c -in sheet.tex > probability_measures.tex

probability_measures.pdf: ../../*.tex ../../trademark.pdf *.tex probability_measures.tex
	pdflatex probability_measures.tex
	make defs

make:
	make probability_measures.pdf

open: probability_measures.pdf
	open probability_measures.pdf

o:
	make open

defs: probability_measures.tex
	defs probability_measures.tex

spell:
	aspell -c probability_measures.tex

s:
	make spell