# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make canonical_maps.pdf

canonical_maps.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > canonical_maps.tex

canonical_maps.pdf: ../../*.tex ../../trademark.pdf *.tex canonical_maps.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode canonical_maps.tex
	../../latexrun canonical_maps.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: canonical_maps.pdf
	open canonical_maps.pdf

o:
	make open

terms: canonical_maps.tex
	bbk_sheet -mode ts canonical_maps.tex

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
	rm -f canonical_maps.*
	make reset
	bbk_sheet -mode mk > Makefile