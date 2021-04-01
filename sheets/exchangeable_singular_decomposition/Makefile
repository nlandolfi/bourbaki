# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make exchangeable_singular_decomposition.pdf
	make graph.pdf

exchangeable_singular_decomposition.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > exchangeable_singular_decomposition.tex

exchangeable_singular_decomposition.pdf: ../../*.tex ../../trademark.pdf *.tex exchangeable_singular_decomposition.tex
	pdflatex --file-line-error -interaction=nonstopmode exchangeable_singular_decomposition.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: exchangeable_singular_decomposition.pdf
	open exchangeable_singular_decomposition.pdf

o:
	make open

terms: exchangeable_singular_decomposition.tex
	bbk_sheet -mode ts exchangeable_singular_decomposition.tex

spell:
	aspell -c exchangeable_singular_decomposition.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f exchangeable_singular_decomposition.*
	make reset
	bbk_sheet -mode mk > Makefile