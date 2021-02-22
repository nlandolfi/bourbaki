# generated automatically by sheet.go

make:
	make probability_outcomes.pdf

probability_outcomes.tex: sheet.tex
	sheet -mode c -in sheet.tex > probability_outcomes.tex

probability_outcomes.pdf: ../../*.tex ../../trademark.pdf *.tex probability_outcomes.tex
	pdflatex probability_outcomes.tex
	make defs


open: probability_outcomes.pdf
	open probability_outcomes.pdf

o:
	make open

defs: probability_outcomes.tex
	defs probability_outcomes.tex

spell:
	aspell -c probability_outcomes.tex

s:
	make spell