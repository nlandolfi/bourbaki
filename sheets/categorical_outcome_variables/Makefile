# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make categorical_outcome_variables.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

categorical_outcome_variables.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > categorical_outcome_variables.tex

categorical_outcome_variables.pdf: ../../*.tex ../../trademark.pdf *.tex categorical_outcome_variables.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode categorical_outcome_variables.tex
	../../latexrun categorical_outcome_variables.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: categorical_outcome_variables.pdf
	open categorical_outcome_variables.pdf

o:
	make open

terms: categorical_outcome_variables.tex
	bbk sheet -mode ts categorical_outcome_variables.tex

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
	rm -f categorical_outcome_variables.*
	make reset
	bbk sheet -mode mk > Makefile