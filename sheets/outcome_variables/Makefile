# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make outcome_variables.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

outcome_variables.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > outcome_variables.tex

outcome_variables.pdf: ../../*.tex ../../trademark.pdf *.tex outcome_variables.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode outcome_variables.tex
	../../latexrun outcome_variables.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: outcome_variables.pdf
	open outcome_variables.pdf

o:
	make open

terms: outcome_variables.tex
	bbk sheet -mode ts outcome_variables.tex

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
	rm -f outcome_variables.*
	make reset
	bbk sheet -mode mk > Makefile