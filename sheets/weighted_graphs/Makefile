# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make weighted_graphs.pdf
	make graph.pdf

weighted_graphs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > weighted_graphs.tex

weighted_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex weighted_graphs.tex
	# pdflatex --file-line-error -interaction=nonstopmode weighted_graphs.tex
	../../latexrun weighted_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: weighted_graphs.pdf
	open weighted_graphs.pdf

o:
	make open

terms: weighted_graphs.tex
	bbk_sheet -mode ts weighted_graphs.tex

spell:
	aspell -c weighted_graphs.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f weighted_graphs.*
	make reset
	bbk_sheet -mode mk > Makefile