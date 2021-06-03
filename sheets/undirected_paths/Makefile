# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make undirected_paths.pdf

undirected_paths.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > undirected_paths.tex

undirected_paths.pdf: ../../*.tex ../../trademark.pdf *.tex undirected_paths.tex
	# pdflatex --file-line-error -interaction=nonstopmode undirected_paths.tex
	../../latexrun undirected_paths.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: undirected_paths.pdf
	open undirected_paths.pdf

o:
	make open

terms: undirected_paths.tex
	bbk_sheet -mode ts undirected_paths.tex

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
	rm -f undirected_paths.*
	make reset
	bbk_sheet -mode mk > Makefile