# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make convex_hulls.pdf

convex_hulls.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > convex_hulls.tex

convex_hulls.pdf: ../../*.tex ../../trademark.pdf *.tex convex_hulls.tex
	# pdflatex --file-line-error -interaction=nonstopmode convex_hulls.tex
	../../latexrun convex_hulls.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: convex_hulls.pdf
	open convex_hulls.pdf

o:
	make open

terms: convex_hulls.tex
	bbk_sheet -mode ts convex_hulls.tex

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
	rm -f convex_hulls.*
	make reset
	bbk_sheet -mode mk > Makefile