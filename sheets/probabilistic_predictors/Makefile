# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make probabilistic_predictors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

probabilistic_predictors.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > probabilistic_predictors.tex

probabilistic_predictors.pdf: ../../*.tex ../../trademark.pdf *.tex probabilistic_predictors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode probabilistic_predictors.tex
	../../latexrun probabilistic_predictors.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probabilistic_predictors.pdf
	open probabilistic_predictors.pdf

o:
	make open

terms: probabilistic_predictors.tex
	bbk sheet -mode ts probabilistic_predictors.tex

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
	rm -f probabilistic_predictors.*
	make reset
	bbk sheet -mode mk > Makefile