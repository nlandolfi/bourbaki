# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make undirected_subgraphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

undirected_subgraphs.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > undirected_subgraphs.tex

undirected_subgraphs.pdf: ../../*.tex ../../trademark.pdf *.tex undirected_subgraphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode undirected_subgraphs.tex
	../../latexrun undirected_subgraphs.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: undirected_subgraphs.pdf
	open undirected_subgraphs.pdf

o:
	make open

terms: undirected_subgraphs.tex
	bbk sheet -mode ts undirected_subgraphs.tex

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
	rm -f undirected_subgraphs.*
	make reset
	bbk sheet -mode mk > Makefile