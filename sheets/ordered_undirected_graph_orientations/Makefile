# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make ordered_undirected_graph_orientations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

ordered_undirected_graph_orientations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > ordered_undirected_graph_orientations.tex

ordered_undirected_graph_orientations.pdf: ../../*.tex ../../trademark.pdf *.tex ordered_undirected_graph_orientations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode ordered_undirected_graph_orientations.tex
	../../latexrun ordered_undirected_graph_orientations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: ordered_undirected_graph_orientations.pdf
	open ordered_undirected_graph_orientations.pdf

o:
	make open

terms: ordered_undirected_graph_orientations.tex
	bbk_sheet -mode ts ordered_undirected_graph_orientations.tex

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
	rm -f ordered_undirected_graph_orientations.*
	make reset
	bbk_sheet -mode mk > Makefile