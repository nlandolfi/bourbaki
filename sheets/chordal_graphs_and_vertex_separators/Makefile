# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make chordal_graphs_and_vertex_separators.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

chordal_graphs_and_vertex_separators.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > chordal_graphs_and_vertex_separators.tex

chordal_graphs_and_vertex_separators.pdf: ../../*.tex ../../trademark.pdf *.tex chordal_graphs_and_vertex_separators.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode chordal_graphs_and_vertex_separators.tex
	../../latexrun chordal_graphs_and_vertex_separators.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: chordal_graphs_and_vertex_separators.pdf
	open chordal_graphs_and_vertex_separators.pdf

o:
	make open

terms: chordal_graphs_and_vertex_separators.tex
	bbk_sheet -mode ts chordal_graphs_and_vertex_separators.tex

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
	rm -f chordal_graphs_and_vertex_separators.*
	make reset
	bbk_sheet -mode mk > Makefile