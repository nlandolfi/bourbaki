# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make subspace_distance.pdf

subspace_distance.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > subspace_distance.tex

subspace_distance.pdf: ../../*.tex ../../trademark.pdf *.tex subspace_distance.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode subspace_distance.tex
	../../latexrun subspace_distance.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subspace_distance.pdf
	open subspace_distance.pdf

o:
	make open

terms: subspace_distance.tex
	bbk_sheet -mode ts subspace_distance.tex

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
	rm -f subspace_distance.*
	make reset
	bbk_sheet -mode mk > Makefile