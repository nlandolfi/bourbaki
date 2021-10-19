# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make elimination_graphs.pdf

elimination_graphs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > elimination_graphs.tex

elimination_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex elimination_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode elimination_graphs.tex
	../../latexrun elimination_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: elimination_graphs.pdf
	open elimination_graphs.pdf

o:
	make open

terms: elimination_graphs.tex
	bbk_sheet -mode ts elimination_graphs.tex

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
	rm -f elimination_graphs.*
	make reset
	bbk_sheet -mode mk > Makefile