# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make bipartite_graphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

bipartite_graphs.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > bipartite_graphs.tex

bipartite_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex bipartite_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode bipartite_graphs.tex
	../../latexrun bipartite_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: bipartite_graphs.pdf
	open bipartite_graphs.pdf

o:
	make open

terms: bipartite_graphs.tex
	bbk_sheet -mode ts bipartite_graphs.tex

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
	rm -f bipartite_graphs.*
	make reset
	bbk_sheet -mode mk > Makefile