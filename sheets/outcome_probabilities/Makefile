# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make outcome_probabilities.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

outcome_probabilities.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > outcome_probabilities.tex

outcome_probabilities.pdf: ../../*.tex ../../trademark.pdf *.tex outcome_probabilities.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode outcome_probabilities.tex
	../../latexrun outcome_probabilities.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: outcome_probabilities.pdf
	open outcome_probabilities.pdf

o:
	make open

terms: outcome_probabilities.tex
	bbk sheet -mode ts outcome_probabilities.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f outcome_probabilities.*
	make reset
	bbk sheet -mode mk > Makefile