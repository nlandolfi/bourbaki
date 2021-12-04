# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make smooth_manifolds.pdf

smooth_manifolds.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > smooth_manifolds.tex

smooth_manifolds.pdf: ../../*.tex ../../trademark.pdf *.tex smooth_manifolds.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode smooth_manifolds.tex
	../../latexrun smooth_manifolds.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: smooth_manifolds.pdf
	open smooth_manifolds.pdf

o:
	make open

terms: smooth_manifolds.tex
	bbk_sheet -mode ts smooth_manifolds.tex

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
	rm -f smooth_manifolds.*
	make reset
	bbk_sheet -mode mk > Makefile