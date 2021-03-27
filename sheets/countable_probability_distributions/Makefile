# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make countable_probability_distributions.pdf
	make graph.pdf

countable_probability_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > countable_probability_distributions.tex

countable_probability_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex countable_probability_distributions.tex
	pdflatex --file-line-error -interaction=nonstopmode countable_probability_distributions.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: countable_probability_distributions.pdf
	open countable_probability_distributions.pdf

o:
	make open

defs: countable_probability_distributions.tex
	defs countable_probability_distributions.tex

spell:
	aspell -c countable_probability_distributions.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f countable_probability_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile