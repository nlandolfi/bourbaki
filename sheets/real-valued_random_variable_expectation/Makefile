# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real-valued_random_variable_expectation.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real-valued_random_variable_expectation.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > real-valued_random_variable_expectation.tex

real-valued_random_variable_expectation.pdf: ../../*.tex ../../trademark.pdf *.tex real-valued_random_variable_expectation.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real-valued_random_variable_expectation.tex
	../../latexrun real-valued_random_variable_expectation.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real-valued_random_variable_expectation.pdf
	open real-valued_random_variable_expectation.pdf

o:
	make open

terms: real-valued_random_variable_expectation.tex
	bbk sheet -mode ts real-valued_random_variable_expectation.tex

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
	rm -f real-valued_random_variable_expectation.*
	make reset
	bbk sheet -mode mk > Makefile