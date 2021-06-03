# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make probability_distributions.pdf

probability_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probability_distributions.tex

probability_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex probability_distributions.tex
	# pdflatex --file-line-error -interaction=nonstopmode probability_distributions.tex
	../../latexrun probability_distributions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probability_distributions.pdf
	open probability_distributions.pdf

o:
	make open

terms: probability_distributions.tex
	bbk_sheet -mode ts probability_distributions.tex

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
	rm -f probability_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile