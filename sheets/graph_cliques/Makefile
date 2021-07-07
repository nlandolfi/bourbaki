# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make graph_cliques.pdf

graph_cliques.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > graph_cliques.tex

graph_cliques.pdf: ../../*.tex ../../trademark.pdf *.tex graph_cliques.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode graph_cliques.tex
	../../latexrun graph_cliques.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: graph_cliques.pdf
	open graph_cliques.pdf

o:
	make open

terms: graph_cliques.tex
	bbk_sheet -mode ts graph_cliques.tex

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
	rm -f graph_cliques.*
	make reset
	bbk_sheet -mode mk > Makefile