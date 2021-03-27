# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make almost_everywhere_measurability.pdf
	make graph.pdf

almost_everywhere_measurability.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > almost_everywhere_measurability.tex

almost_everywhere_measurability.pdf: ../../*.tex ../../trademark.pdf *.tex almost_everywhere_measurability.tex
	pdflatex --file-line-error -interaction=nonstopmode almost_everywhere_measurability.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: almost_everywhere_measurability.pdf
	open almost_everywhere_measurability.pdf

o:
	make open

defs: almost_everywhere_measurability.tex
	defs almost_everywhere_measurability.tex

spell:
	aspell -c almost_everywhere_measurability.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f almost_everywhere_measurability.*
	make reset
	bbk_sheet -mode mk > Makefile