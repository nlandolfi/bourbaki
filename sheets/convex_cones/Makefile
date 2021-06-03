# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make convex_cones.pdf

convex_cones.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > convex_cones.tex

convex_cones.pdf: ../../*.tex ../../trademark.pdf *.tex convex_cones.tex
	# pdflatex --file-line-error -interaction=nonstopmode convex_cones.tex
	../../latexrun convex_cones.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: convex_cones.pdf
	open convex_cones.pdf

o:
	make open

terms: convex_cones.tex
	bbk_sheet -mode ts convex_cones.tex

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
	rm -f convex_cones.*
	make reset
	bbk_sheet -mode mk > Makefile