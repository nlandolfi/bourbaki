# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make directed_acyclic_graph_ancestry_partial_order.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

directed_acyclic_graph_ancestry_partial_order.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > directed_acyclic_graph_ancestry_partial_order.tex

directed_acyclic_graph_ancestry_partial_order.pdf: ../../*.tex ../../trademark.pdf *.tex directed_acyclic_graph_ancestry_partial_order.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode directed_acyclic_graph_ancestry_partial_order.tex
	../../latexrun directed_acyclic_graph_ancestry_partial_order.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: directed_acyclic_graph_ancestry_partial_order.pdf
	open directed_acyclic_graph_ancestry_partial_order.pdf

o:
	make open

terms: directed_acyclic_graph_ancestry_partial_order.tex
	bbk sheet -mode ts directed_acyclic_graph_ancestry_partial_order.tex

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
	rm -f directed_acyclic_graph_ancestry_partial_order.*
	make reset
	bbk sheet -mode mk > Makefile