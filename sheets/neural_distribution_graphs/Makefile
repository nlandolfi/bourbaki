# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make neural_distribution_graphs.pdf

neural_distribution_graphs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > neural_distribution_graphs.tex

neural_distribution_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex neural_distribution_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode neural_distribution_graphs.tex
	../../latexrun neural_distribution_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: neural_distribution_graphs.pdf
	open neural_distribution_graphs.pdf

o:
	make open

terms: neural_distribution_graphs.tex
	bbk_sheet -mode ts neural_distribution_graphs.tex

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
	rm -f neural_distribution_graphs.*
	make reset
	bbk_sheet -mode mk > Makefile