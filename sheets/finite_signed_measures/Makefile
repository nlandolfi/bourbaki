# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make finite_signed_measures.pdf
	make graph.pdf

finite_signed_measures.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > finite_signed_measures.tex

finite_signed_measures.pdf: ../../*.tex ../../trademark.pdf *.tex finite_signed_measures.tex
	pdflatex --file-line-error -interaction=nonstopmode finite_signed_measures.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: finite_signed_measures.pdf
	open finite_signed_measures.pdf

o:
	make open

defs: finite_signed_measures.tex
	defs finite_signed_measures.tex

spell:
	aspell -c finite_signed_measures.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f finite_signed_measures.*
	make reset
	bbk_sheet -mode mk > Makefile