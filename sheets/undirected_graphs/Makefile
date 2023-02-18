# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make undirected_graphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

undirected_graphs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > undirected_graphs.tex

undirected_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex undirected_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode undirected_graphs.tex
	../../latexrun undirected_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: undirected_graphs.pdf
	open undirected_graphs.pdf

o:
	make open

terms: undirected_graphs.tex
	bbk_sheet -mode ts undirected_graphs.tex

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
	rm -f undirected_graphs.*
	make reset
	bbk_sheet -mode mk > Makefile